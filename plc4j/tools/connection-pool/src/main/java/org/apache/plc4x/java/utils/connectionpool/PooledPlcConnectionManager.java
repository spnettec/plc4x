/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */
package org.apache.plc4x.java.utils.connectionpool;

import org.apache.commons.lang3.reflect.FieldUtils;
import org.apache.commons.pool2.KeyedObjectPool;
import org.apache.commons.pool2.impl.GenericKeyedObjectPool;
import org.apache.plc4x.java.DefaultPlcDriverManager;
import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.authentication.PlcAuthentication;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.lang.reflect.InvocationTargetException;
import java.lang.reflect.Proxy;
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.atomic.AtomicBoolean;

/**
 * @author heyoulin
 */
public class PooledPlcConnectionManager extends DefaultPlcDriverManager {

    private static final Logger LOGGER = LoggerFactory.getLogger(PooledPlcConnectionManager.class);

    private KeyedObjectPool<AbstractPoolKey, PlcConnection> keyedObjectPool;

    // Marker class do detected a non null value
    static final NoPlcAuthentication noPlcAuthentication = new NoPlcAuthentication();

    private final PoolKeyFactory poolKeyFactory;

    public PooledPlcConnectionManager() {
        this(GenericKeyedObjectPool::new);
    }

    public PooledPlcConnectionManager(PoolKeyFactory poolKeyFactory) {
        this(GenericKeyedObjectPool::new, poolKeyFactory);
    }

    public PooledPlcConnectionManager(ClassLoader classLoader) {
        this(classLoader, new PoolKeyFactory());
    }

    public PooledPlcConnectionManager(ClassLoader classLoader, PoolKeyFactory poolKeyFactory) {
        super(classLoader);
        setFromPoolCreator(GenericKeyedObjectPool::new);
        this.poolKeyFactory = poolKeyFactory;
    }

    public PooledPlcConnectionManager(PoolCreator poolCreator) {
        this(poolCreator, new PoolKeyFactory());
    }

    public PooledPlcConnectionManager(PoolCreator poolCreator, PoolKeyFactory poolKeyFactory) {
        setFromPoolCreator(poolCreator);
        this.poolKeyFactory = poolKeyFactory;
    }

    public PooledPlcConnectionManager(ClassLoader classLoader, PoolCreator poolCreator) {
        super(classLoader);
        setFromPoolCreator(poolCreator);
        poolKeyFactory = new PoolKeyFactory();
    }

    private void setFromPoolCreator(PoolCreator poolCreator) {
        this.keyedObjectPool = poolCreator.createPool(new PooledPlcConnectionFactory() {
            @Override
            public PlcConnection create(AbstractPoolKey key) throws Exception {
                PlcAuthentication plcAuthentication = key.plcAuthentication;
                String url = key.url;
                if (plcAuthentication == noPlcAuthentication) {
                    LOGGER.debug("getting actual connection for {}", url);
                    return PooledPlcConnectionManager.super.getConnection(url);
                } else {
                    LOGGER.debug("getting actual connection for {} and plcAuthentication {}", url, plcAuthentication);
                    return PooledPlcConnectionManager.super.getConnection(url, plcAuthentication);
                }
            }
        });
    }

    @Override
    public PlcConnection getConnection(String url) throws PlcConnectionException {
        return getConnection(url, noPlcAuthentication);
    }

    @Override
    public PlcConnection getConnection(String url, PlcAuthentication authentication) throws PlcConnectionException {
        AbstractPoolKey poolKey = poolKeyFactory.getPoolKey(url, authentication);
        if (LOGGER.isDebugEnabled()) {
            if (authentication != noPlcAuthentication) {
                LOGGER.debug("Try to borrow an object for url {} and authentication {}", url, authentication);
            } else {
                LOGGER.debug("Try to borrow an object for url {}", url);
            }
        }
        PlcConnection plcConnection;
        try {
            plcConnection = keyedObjectPool.borrowObject(poolKey);
            if (!plcConnection.isConnected()) {
                LOGGER.debug("Attempting to reconnect to device");
                plcConnection.connect();
            }
        } catch (Exception e) {
            throw new PlcConnectionException(e);
        }
        // Used to invalidate a proxy
        AtomicBoolean proxyInvalidated = new AtomicBoolean(false);
        return (PlcConnection) Proxy.newProxyInstance(classLoader, new Class[]{PlcConnection.class}, (proxy, method, args) -> {
            if (proxyInvalidated.get()) {
                throw new IllegalStateException("Proxy not valid anymore");
            }
            if ("close".equals(method.getName())) {
                LOGGER.debug("close called on {}", plcConnection);
                proxyInvalidated.set(true);
                keyedObjectPool.returnObject(poolKey, plcConnection);
                return null;
            } else {
                try {
                    return method.invoke(plcConnection, args);
                } catch (InvocationTargetException e) {
                    if (e.getCause().getClass() == PlcConnectionException.class) {
                        keyedObjectPool.invalidateObject(poolKey, plcConnection);
                        proxyInvalidated.set(true);
                    }
                    throw e;
                }
            }
        });
    }

    @FunctionalInterface
    public interface PoolCreator {
        KeyedObjectPool<AbstractPoolKey, PlcConnection> createPool(PooledPlcConnectionFactory pooledPlcConnectionFactory);
    }

    // TODO: maybe export to jmx // generic poolKey has builtin jmx too
    public Map<String, Number> getStatistics() {
        HashMap<String, Number> statistics = new HashMap<>();
        statistics.put("numActive", keyedObjectPool.getNumActive());
        statistics.put("numIdle", keyedObjectPool.getNumIdle());
        if (keyedObjectPool instanceof GenericKeyedObjectPool) {
            GenericKeyedObjectPool<AbstractPoolKey, PlcConnection> genericKeyedObjectPool = (GenericKeyedObjectPool<AbstractPoolKey, PlcConnection>) this.keyedObjectPool;
            // Thats pretty ugly and we really should't do that...
            try {
                Map poolMap = (Map) FieldUtils.getField(GenericKeyedObjectPool.class, "poolMap", true).get(this.keyedObjectPool);
                statistics.put("pools.count", poolMap.size());
            } catch (IllegalAccessException e) {
                throw new PlcRuntimeException(e);
            }
            Map<String, Integer> numActivePerKey = genericKeyedObjectPool.getNumActivePerKey();
            for (Map.Entry<String, Integer> entry : numActivePerKey.entrySet()) {
                statistics.put(entry.getKey() + ".numActive", entry.getValue());
            }
        }

        return statistics;
    }

    private static final class NoPlcAuthentication implements PlcAuthentication {

    }
}
