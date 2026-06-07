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

/**
 * Caching {@link PlcConnectionManager} — establishes a single underlying connection
 * per URL and leases it to concurrent callers via {@link ConnectionContainer}.
 * Connection lifecycle (creation, health check, teardown) is managed here;
 * {@link ConnectionContainer} is a pure lease wrapper.
 */
public class CachedPlcConnectionManager implements PlcConnectionManager {

    private static final Logger LOG = LoggerFactory.getLogger(CachedPlcConnectionManager.class);

    private final PlcConnectionManager connectionManager;
    private final Duration maxLeaseTime;
    private final Duration maxWaitTime;

    private final Map<String, ConnectionContainer> connectionContainers;

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
        // For serial transports, normalize symlink/canonical-path variants to share one container.
        String cacheKey = normalizeCacheKey(url);

        ConnectionContainer connectionContainer;
        synchronized (connectionContainers) {
            connectionContainer = connectionContainers.get(cacheKey);
            if (connectionContainer == null || connectionContainer.isClosed()) {
                LOG.debug("Creating new cached connection for {}", url);
                PlcConnection connection = connectionManager.getConnection(url);
                connectionContainer = new ConnectionContainer(connection, maxLeaseTime);
                connectionContainers.put(cacheKey, connectionContainer);
            } else if (connectionContainer.getRawConnection() == null
                    || !connectionContainer.getRawConnection().isConnected()) {
                LOG.debug("Cached connection for {} is dead, recreating...", url);
                connectionContainer.close();
                PlcConnection connection = connectionManager.getConnection(url);
                connectionContainer = new ConnectionContainer(connection, maxLeaseTime);
                connectionContainers.put(cacheKey, connectionContainer);
            } else {
                LOG.debug("Reusing existing cached connection for {}", url);
            }
        }

        Future<PlcConnection> leaseFuture = connectionContainer.lease();
        try {
            return leaseFuture.get(this.maxWaitTime.toMillis(), TimeUnit.MILLISECONDS);
        } catch (ExecutionException e) {
            connectionContainer.close();
            connectionContainers.remove(cacheKey);
            throw new PlcConnectionException(e);
        } catch (TimeoutException e) {
            connectionContainer.close();
            connectionContainers.remove(cacheKey);
            throw new PlcConnectionException("Error acquiring lease for connection cause TimeoutException", e);
        } catch (InterruptedException e) {
            Thread.currentThread().interrupt();
            connectionContainer.close();
            connectionContainers.remove(cacheKey);
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
     * Drop the cached container for {@code url} and close its underlying connection
     * (fd / socket / Netty EventLoopGroup are reclaimed).
     */
    @Override
    public void invalidate(String url) {
        if (url == null) {
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

    /**
     * Close all cached connections and clear the container map. After this call the
     * manager is still usable — the next {@link #getConnection(String)} creates fresh
     * connections.
     */
    public void destroy() {
        synchronized (connectionContainers) {
            connectionContainers.values().forEach(ConnectionContainer::close);
            connectionContainers.clear();
        }
    }

    /**
     * Alias for {@link #destroy()} — follows the conventional Java resource-release
     * naming so callers that expect {@code close()} (e.g. OSGi deactivate) work
     * without special-casing.
     */
    public void close() {
        destroy();
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
        }
    }

}
