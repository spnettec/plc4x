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
import org.apache.plc4x.java.api.PlcConnectionManager;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.listener.EventListener;
import org.apache.plc4x.java.utils.cache.exceptions.PlcConnectionManagerClosedException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.LinkedList;
import java.util.Queue;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.Future;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicLong;

class ConnectionContainer {
    private static final Logger LOGGER = LoggerFactory.getLogger(ConnectionContainer.class);
    private static final long LEASE_WARN_THROTTLE_NANOS = TimeUnit.SECONDS.toNanos(30);

    private final PlcConnectionManager connectionManager;
    private final String connectionUrl;
    private final Duration maxLeaseTime;
    private final Queue<CompletableFuture<PlcConnection>> queue;

    private PlcConnection connection;
    private LeasedPlcConnection leasedConnection;

    // First-failure WARN with full stack trace is the operationally useful signal
    // ("PLC just went unreachable"); a wire-driven 100ms retry loop turns it into
    // ~10/s with 30-line stack traces, which buries every other log. Keep the first
    // event in each 30s window at WARN; demote the rest to DEBUG and report the
    // suppressed count on the next WARN.
    private final AtomicLong leaseWarnLastNanos = new AtomicLong(0L);
    private final AtomicLong leaseWarnSuppressed = new AtomicLong(0L);

    public ConnectionContainer(PlcConnectionManager connectionManager, String connectionUrl,
                               Duration maxLeaseTime) {
        this.connectionManager = connectionManager;
        this.connectionUrl = connectionUrl;
        this.maxLeaseTime = maxLeaseTime;
        this.queue = new LinkedList<>();
        this.connection = null;
        this.leasedConnection = null;
    }

    public synchronized void close() {
        // Close all waiting clients exceptionally.
        queue.forEach(plcConnectionCompletableFuture ->
            plcConnectionCompletableFuture.completeExceptionally(new PlcConnectionManagerClosedException()));

        // Clear the queue.
        queue.clear();

        // If the connection is currently used, close it.
        if(leasedConnection != null) {
            try {
                leasedConnection.closeConnection();
                leasedConnection = null;
            } catch (Exception e) {
                // Ignore this ...
            }
        } else {
            try {
                connection.close();
            } catch (Exception e) {
                // Ignore this ...
            }
        }
    }

    public synchronized Future<PlcConnection> lease() {
        CompletableFuture<PlcConnection> connectionFuture = new CompletableFuture<>();

        // Try to get a new connection, if we haven't got one yet.
        if(connection == null || !connection.isConnected()) {
            try {
                connection = connectionManager.getConnection(connectionUrl);
            } catch (PlcConnectionException e) {
                logLeaseFailure(e);
                connectionFuture.completeExceptionally(e);
                return connectionFuture;
            }
        }

        // If the connection is currently idle, return the connection immediately.
        if (leasedConnection == null) {
            leasedConnection = new LeasedPlcConnection(this, connection, maxLeaseTime);
            connectionFuture.complete(leasedConnection);
        }
        // Otherwise queue the future up for completion as soon as the connection is returned.
        else {
            queue.add(connectionFuture);
        }
        return connectionFuture;
    }

    public synchronized void returnConnection(LeasedPlcConnection returnedLeasedConnection, boolean invalidateConnection) {
        if(returnedLeasedConnection != leasedConnection) {
            LOGGER.error("Error trying to return lease from invalid connection: returned={} leased={}",
                returnedLeasedConnection, leasedConnection);
            throw new PlcRuntimeException("Error trying to return lease from invalid connection");
        }

        // If something happened while using the connection, invalidate this one. Reconnect
        // lazily — only when someone is actually waiting in the queue. Serial transports in
        // particular need the kernel to release the underlying fd before reopen succeeds;
        // eagerly reopening here racing with our own async close() yields a flood of
        // "Unable to open the com port" failures under wire-driven 100ms poll loops.
        if(invalidateConnection) {
            if (connection != null) {
                try {
                    connection.close();
                } catch (Exception e) {
                    // We're ignoring this as we have no idea, what state the connection is in.
                    // Nevertheless, it is polite to say something in logs about this situation.
                    LOGGER.warn("Exception while closing connection", e);
                }
                connection = null;
            }
            if(returnedLeasedConnection == null){
                return;
            }
        }

        // If the queue is empty, defer any reconnect to the next lease() call.
        if(queue.isEmpty()) {
            leasedConnection = null;
            return;
        }

        // Someone is waiting — (re)establish the connection now if needed.
        if(connection == null || !connection.isConnected()) {
            try {
                connection = connectionManager.getConnection(connectionUrl);
            } catch (PlcConnectionException e) {
                logLeaseFailure(e);
                queue.forEach(future -> future.completeExceptionally(e));
                queue.clear();
                leasedConnection = null;
                connection = null;
                return;
            }
        }

        // Create a new lease and complete the next future in the queue with this.
        leasedConnection = new LeasedPlcConnection(this, connection, maxLeaseTime);
        CompletableFuture<PlcConnection> leaseFuture = queue.poll();
        if(leaseFuture != null) {
            leaseFuture.complete(leasedConnection);
        }
    }


    private void logLeaseFailure(PlcConnectionException e) {
        long now = System.nanoTime();
        long last = leaseWarnLastNanos.get();
        if (now - last >= LEASE_WARN_THROTTLE_NANOS && leaseWarnLastNanos.compareAndSet(last, now)) {
            long suppressed = leaseWarnSuppressed.getAndSet(0L);
            if (suppressed == 0L) {
                LOGGER.warn("Exception while getting connection for lease", e);
            } else {
                LOGGER.warn("Exception while getting connection for lease (suppressed {} similar in last {}s)",
                    suppressed, TimeUnit.NANOSECONDS.toSeconds(LEASE_WARN_THROTTLE_NANOS), e);
            }
        } else {
            leaseWarnSuppressed.incrementAndGet();
            LOGGER.debug("Exception while getting connection for lease", e);
        }
    }

    public void addEventListener(EventListener listener) {
        if((connection != null) && (connection instanceof EventPlcConnection)) {
            ((EventPlcConnection) connection).addEventListener(listener);
        }
    }

    public void removeEventListener(EventListener listener) {
        if((connection != null) && (connection instanceof EventPlcConnection)) {
            ((EventPlcConnection) connection).removeEventListener(listener);
        }
    }

}
