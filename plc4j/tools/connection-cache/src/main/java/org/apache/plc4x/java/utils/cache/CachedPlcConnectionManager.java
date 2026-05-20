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

import org.apache.plc4x.java.DefaultPlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.PlcConnectionManager;
import org.apache.plc4x.java.api.PlcDriverManager;
import org.apache.plc4x.java.api.authentication.PlcAuthentication;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.utils.cache.exceptions.PlcConnectionManagerClosedException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.io.IOException;
import java.nio.file.Paths;
import java.time.Duration;
import java.util.*;
import java.util.concurrent.ExecutionException;
import java.util.concurrent.Future;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;
import java.util.concurrent.atomic.AtomicBoolean;

public class CachedPlcConnectionManager implements PlcConnectionManager, AutoCloseable {

    private static final Logger LOG = LoggerFactory.getLogger(CachedPlcConnectionManager.class);

    private final PlcConnectionManager connectionManager;
    private final Duration maxLeaseTime;
    private final Duration maxWaitTime;

    private final Map<String, ConnectionContainer> connectionContainers;

    private final AtomicBoolean closed = new AtomicBoolean(false);

    public static Builder getBuilder() {
        return new Builder(new DefaultPlcDriverManager());
    }

    public static Builder getBuilder(PlcConnectionManager connectionManager) {
        return new Builder(connectionManager);
    }

    public CachedPlcConnectionManager(PlcConnectionManager connectionManager, Duration maxLeaseTime, Duration maxWaitTime) {
        this.connectionManager = connectionManager;
        this.maxLeaseTime = maxLeaseTime;
        this.maxWaitTime = maxWaitTime;
        this.connectionContainers = new HashMap<>();
    }

    @Override
    public PlcConnection getConnection(String url) throws PlcConnectionException {
        // If the connection manager is already closed, abort.
        if(closed.get()) {
            throw new PlcConnectionManagerClosedException();
        }

        // For serial transports, two different URLs (e.g. /dev/ttyUSB0 vs /dev/serial/by-path/...)
        // can resolve to the same underlying device. Without normalization the cache would create
        // two ConnectionContainers each trying to grab the same exclusive serial fd, and the second
        // open fails with EBUSY (manifesting as "Error creating channel"). Normalize the path
        // portion via toRealPath() so symlink and canonical-path variants share one container.
        String cacheKey = normalizeCacheKey(url);

        // Get a connection container for the given url.
        ConnectionContainer connectionContainer;
        synchronized (connectionContainers) {
            connectionContainer = connectionContainers.get(cacheKey);
            if (connectionContainer == null) {
                LOG.debug("Creating new connection");

                // Crate a connection container to manage handling this connection.
                // Keep the user-supplied URL inside the container so error messages and downstream
                // driver logs preserve the literal string the caller wrote.
                connectionContainer = new ConnectionContainer(connectionManager, url, maxLeaseTime);
                connectionContainers.put(cacheKey, connectionContainer);
            } else {
                LOG.debug("Reusing exising connection");
            }
        }

        // Get a lease (a future for a connection)
        Future<PlcConnection> leaseFuture = connectionContainer.lease();
        try {
            return leaseFuture.get(this.maxWaitTime.toMillis(), TimeUnit.MILLISECONDS);
        } catch (ExecutionException e) {
            throw new PlcConnectionException(e);
        } catch ( TimeoutException e) {
            throw new PlcConnectionException("Error acquiring lease for connection cause TimeoutException", e);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            throw new PlcConnectionException("Error acquiring lease for connection cause InterruptedException", e);
        }
    }

    /**
     * Normalize a serial transport URL so symlink and canonical-path forms collapse to the same
     * cache key. For non-serial transports (or any error resolving the path) the original URL is
     * returned unchanged.
     */
    static String normalizeCacheKey(String url) {
        if (url == null) {
            return url;
        }
        int sep = url.indexOf("://");
        if (sep < 0) {
            return url;
        }
        // Scheme is "<protocol>" or "<protocol>:<transport>"; only canonicalize when transport == serial.
        String scheme = url.substring(0, sep);
        if (!scheme.endsWith(":serial") && !scheme.equals("serial")) {
            return url;
        }
        int q = url.indexOf('?', sep + 3);
        String path = (q < 0) ? url.substring(sep + 3) : url.substring(sep + 3, q);
        String tail = (q < 0) ? "" : url.substring(q);
        try {
            String canonical = Paths.get(path).toRealPath().toString();
            return url.substring(0, sep + 3) + canonical + tail;
        } catch (IOException | RuntimeException e) {
            // Device not yet present, no permission, or path invalid — leave key as-is so the
            // downstream open() can produce its own error.
            return url;
        }
    }

    @Override
    public PlcDriverManager getDriverManager() {
        return connectionManager.getDriverManager();
    }

    @Override
    public PlcConnection getConnection(String url, PlcAuthentication authentication) throws PlcConnectionException {
        throw new PlcConnectionException("the cached driver manager currently doesn't support authentication");
    }

    /**
     * Drop the cached container for {@code url} and close its underlying connection.
     *
     * <p>Use this when a caller knows a connection-string is no longer in use (typically on OSGi
     * config change or component deactivate). Without an explicit drop the container stays in the
     * map holding the underlying fd until either a read errors (invalidating it through
     * {@link ConnectionContainer#returnConnection}) or the whole manager is closed.
     *
     * <p>If a lease is currently outstanding on the container, its underlying connection is closed
     * synchronously; the lease holder will observe a closed connection on its next operation. This
     * is intentional — invalidation is the caller's signal that the URL must be released now.
     *
     * <p>No-op if {@code url} was never cached.
     */
    public void invalidate(String url) {
        if (closed.get() || url == null) {
            return;
        }
        String cacheKey = normalizeCacheKey(url);
        ConnectionContainer container;
        synchronized (connectionContainers) {
            container = connectionContainers.remove(cacheKey);
        }
        if (container != null) {
            LOG.debug("Invalidating cached connection for {}", url);
            container.close();
        }
    }

    @Override
    public void close() throws Exception {
        // Set the cache to "closed" so no new connections can be requested.
        closed.set(true);

        // Tell all connections to close themselves.
        connectionContainers.forEach((connectionString, connectionContainer) -> {
            connectionContainer.close();
        });
    }

    public static class Builder {

        private final PlcConnectionManager connectionManager;
        private Duration maxLeaseTime;
        private Duration maxWaitTime;

        public Builder(PlcConnectionManager connectionManager) {
            this.connectionManager = connectionManager;
            this.maxLeaseTime = Duration.ofSeconds(4);
            this.maxWaitTime = Duration.ofSeconds(20);
        }

        public CachedPlcConnectionManager build() {
            return new CachedPlcConnectionManager(
                this.connectionManager, this.maxLeaseTime, this.maxWaitTime);
        }

        public CachedPlcConnectionManager.Builder withMaxLeaseTime(Duration maxLeaseTime) {
            this.maxLeaseTime = maxLeaseTime;
            return this;
        }

        public CachedPlcConnectionManager.Builder withMaxWaitTime(Duration maxWaitTime) {
            this.maxWaitTime = maxWaitTime;
            return this;
        }}

}
