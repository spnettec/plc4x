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
package org.apache.plc4x.java.utils.connectionpool2;

import org.apache.plc4x.java.DefaultPlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.PlcConnectionManager;
import org.apache.plc4x.java.api.PlcDriverManager;
import org.apache.plc4x.java.api.authentication.PlcAuthentication;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;

import java.time.Duration;
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.ConcurrentHashMap;
import java.util.function.Function;

/**
 * Pool that sits on top of the {@link CachedDriverManager}.
 * <p>
 * This class is thread safe!
 */
public class CachedPlcConnectionManager implements PlcConnectionManager {
    private PlcConnectionManager connectionManager;
    private int maxLeaseTime;
    private int timeoutMillis;
    private Function<String, CachedDriverManager> factory;
    private final Map<String, CachedDriverManager> cachedManagers = new ConcurrentHashMap<>();

    public static CachedPlcConnectionManager.Builder getBuilder() {
        return new CachedPlcConnectionManager.Builder(new DefaultPlcDriverManager());
    }

    public static CachedPlcConnectionManager.Builder getBuilder(PlcConnectionManager connectionManager) {
        return new CachedPlcConnectionManager.Builder(connectionManager);
    }
    public CachedPlcConnectionManager(Function<String, CachedDriverManager> factory,PlcConnectionManager connectionManager, int maxLeaseTime, int timeoutMillis) {
        this.connectionManager = connectionManager;
        this.maxLeaseTime = maxLeaseTime;
        this.timeoutMillis = timeoutMillis;
        this.factory = factory;
    }

    CachedPlcConnectionManager(Function<String, CachedDriverManager> factory) {
        this.connectionManager = new DefaultPlcDriverManager();
        this.maxLeaseTime = 4_000;
        this.timeoutMillis = 5_000;
        this.factory = factory;
    }
    CachedPlcConnectionManager() {
        this.connectionManager = new DefaultPlcDriverManager();
        this.maxLeaseTime = 4_000;
        this.timeoutMillis = 5_000;
        this.factory = key -> new CachedDriverManager(key, () -> {
            try {
                return connectionManager.getConnection(key);
            } catch (PlcConnectionException e) {
                throw new RuntimeException(e);
            }
        },timeoutMillis);
    }

    @Override
    public PlcConnection getConnection(String url) throws PlcConnectionException {
        return cachedManagers.computeIfAbsent(url, this.factory).getConnection(url);
    }

    @Override
    public PlcConnection getConnection(String url, PlcAuthentication authentication) throws PlcConnectionException {
        throw new UnsupportedOperationException();
    }

    @Override
    public PlcDriverManager getDriverManager() {
        return connectionManager.getDriverManager();
    }

    Map<String, CachedDriverManager> getCachedManagers() {
        return this.cachedManagers;
    }
    public static class Builder {

        private final PlcConnectionManager connectionManager;
        private int maxLeaseTime;
        private int timeoutMillis;
        private Function<String, CachedDriverManager> factory;

        public Builder(PlcConnectionManager connectionManager) {
            this.connectionManager = connectionManager;
            this.maxLeaseTime = 4_000;
            this.timeoutMillis = 5_000;
            this.factory = key -> new CachedDriverManager(key, () -> {
                try {
                    return connectionManager.getConnection(key);
                } catch (PlcConnectionException e) {
                    throw new RuntimeException(e);
                }
            },timeoutMillis);
        }
        public Builder(Function<String, CachedDriverManager> factory,PlcConnectionManager connectionManager) {
            this.connectionManager = connectionManager;
            this.maxLeaseTime = 4_000;
            this.timeoutMillis = 5_000;
            this.factory = factory;
        }
        public Builder(Function<String, CachedDriverManager> factory) {
            this.connectionManager = new DefaultPlcDriverManager();
            this.maxLeaseTime = 4_000;
            this.timeoutMillis = 5_000;
            this.factory = factory;
        }

        public CachedPlcConnectionManager build() {

            return new CachedPlcConnectionManager(this.factory,this.connectionManager, this.maxLeaseTime, this.timeoutMillis);
        }

        public CachedPlcConnectionManager.Builder withMaxLeaseTime(int maxLeaseTime) {
            this.maxLeaseTime = maxLeaseTime;
            return this;
        }

        public CachedPlcConnectionManager.Builder withTimeoutMillis(int timeoutMillis) {
            this.timeoutMillis = timeoutMillis;
            return this;
        }
        public CachedPlcConnectionManager.Builder withFactory(Function<String, CachedDriverManager> factory) {
            this.factory = factory;
            return this;
        }
    }
}
