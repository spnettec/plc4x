/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.utils.cache;

import org.apache.plc4x.java.api.EventPlcConnection;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.listener.EventListener;

import java.time.Duration;
import java.util.LinkedList;
import java.util.Queue;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.Future;

/**
 * Pure lease manager — wraps a single {@link PlcConnection} and serializes
 * concurrent access through a {@link LeasedPlcConnection}. Connection lifecycle
 * (creation, health check, teardown) is handled by {@link CachedPlcConnectionManager}.
 */
class ConnectionContainer {

    private PlcConnection connection;
    private final Duration maxLeaseTime;
    private final Queue<CompletableFuture<PlcConnection>> queue;

    private LeasedPlcConnection leasedConnection;
    private boolean closed;

    ConnectionContainer(PlcConnection connection, Duration maxLeaseTime) {
        this.connection = connection;
        this.maxLeaseTime = maxLeaseTime;
        this.queue = new LinkedList<>();
    }

    PlcConnection getRawConnection() {
        return connection;
    }

    boolean isClosed() {
        return closed;
    }

    /**
     * Close the container and the underlying connection. Pending lease futures
     * are completed exceptionally.
     */
    synchronized void close() {
        if (closed) {
            return;
        }
        closed = true;
        queue.forEach(f -> f.completeExceptionally(new PlcRuntimeException("Container closed")));
        queue.clear();
        if (leasedConnection != null) {
            try {
                leasedConnection.closeConnection();
                leasedConnection = null;
            } catch (Exception ignored) {
            }
        } else {
            try {
                connection.close();
            } catch (Exception ignored) {
            }
        }
    }

    synchronized Future<PlcConnection> lease() {
        if (closed) {
            CompletableFuture<PlcConnection> future = new CompletableFuture<>();
            future.completeExceptionally(new PlcRuntimeException("Container closed"));
            return future;
        }

        CompletableFuture<PlcConnection> connectionFuture = new CompletableFuture<>();

        if (leasedConnection == null) {
            leasedConnection = new LeasedPlcConnection(this, connection, maxLeaseTime);
            connectionFuture.complete(leasedConnection);
        } else {
            queue.add(connectionFuture);
        }
        return connectionFuture;
    }

    void addEventListener(EventListener listener) {
        if (connection instanceof EventPlcConnection) {
            ((EventPlcConnection) connection).addEventListener(listener);
        }
    }

    void removeEventListener(EventListener listener) {
        if (connection instanceof EventPlcConnection) {
            ((EventPlcConnection) connection).removeEventListener(listener);
        }
    }

    synchronized void returnConnection(LeasedPlcConnection returnedLeasedConnection, boolean connectionError) {
        if (returnedLeasedConnection != leasedConnection) {
            throw new PlcRuntimeException("Error trying to return lease from invalid connection");
        }

        // If an I/O error occurred while using the connection, close it now so
        // CachedPlcConnectionManager.getConnection() creates a fresh one next time.
        if (connectionError) {
            try {
                connection.close();
            } catch (Exception ignored) {
            }
            connection = null;
        }

        if (queue.isEmpty()) {
            leasedConnection = null;
            return;
        }

        if (connection == null) {
            // Connection is gone — fail all waiting futures.
            CompletableFuture<PlcConnection> f = queue.poll();
            while (f != null) {
                f.completeExceptionally(new PlcRuntimeException("Connection invalidated due to I/O error"));
                f = queue.poll();
            }
            leasedConnection = null;
            return;
        }

        leasedConnection = new LeasedPlcConnection(this, connection, maxLeaseTime);
        CompletableFuture<PlcConnection> leaseFuture = queue.poll();
        if (leaseFuture != null) {
            leaseFuture.complete(leasedConnection);
        }
    }
}
