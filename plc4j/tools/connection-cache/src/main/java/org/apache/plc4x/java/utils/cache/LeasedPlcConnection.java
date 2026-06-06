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
import org.apache.plc4x.java.api.listener.EventListener;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.messages.*;
import org.apache.plc4x.java.api.metadata.PlcConnectionMetadata;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.util.*;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicReference;

public class LeasedPlcConnection implements EventPlcConnection {

    private static final Logger log = LoggerFactory.getLogger(LeasedPlcConnection.class);
    private ConnectionContainer connectionContainer;
    private final AtomicReference<PlcConnection> connection;
    private boolean invalidateConnection= false;
    private boolean hasInvalidateConnection= false;
    private final Timer usageTimer;

    LeasedPlcConnection(ConnectionContainer connectionContainer, PlcConnection connection, Duration maxUseTime) {
        this.connectionContainer = connectionContainer;
        this.connection = new AtomicReference<>(connection);
        this.usageTimer = new Timer();
        this.usageTimer.schedule(new TimerTask() {
            @Override
            public void run() {
                close();
            }
        }, Date.from(LocalDateTime.now().plusNanos(maxUseTime.toNanos()).atZone(ZoneId.systemDefault()).toInstant()));
    }

    public synchronized void closeConnection() throws Exception {
        // Get the real connection and close it.
        PlcConnection plcConnection = connection.get();
        if(plcConnection != null) {
            plcConnection.close();
        }

        // Close the LeasedPlcConnection.
        close();
    }

    @Override
    public synchronized void close() {
        // In this case the connection was already closed (possibly by the timer)
        if(connection.get() == null) {
            if(invalidateConnection && !hasInvalidateConnection){
                hasInvalidateConnection = true;
                connectionContainer.returnConnection(null, true);
            }
            return;
        }

        // Cancel automatically timing out.
        usageTimer.cancel();

        // Make the connection unusable.
        connection.set(null);

        if(invalidateConnection){
            hasInvalidateConnection = true;
        }
        // Tell the connection container that the connection is free to be reused.
        connectionContainer.returnConnection(this, invalidateConnection);
    }

    @Override
    public Optional<PlcTag> parseTagAddress(String tagAddress) {
        PlcConnection plcConnection = connection.get();
        if(plcConnection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        return plcConnection.parseTagAddress(tagAddress);
    }

    @Override
    public Optional<PlcValue> parseTagValue(PlcTag tag, Object... values) {
        PlcConnection plcConnection = connection.get();
        if(plcConnection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        return plcConnection.parseTagValue(tag, values);
    }

    @Override
    public void connect() throws PlcConnectionException {
        throw new PlcConnectionException("Error connecting leased connection");
    }

    @Override
    public boolean isConnected() {
        PlcConnection plcConnection = connection.get();
        if(plcConnection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        return plcConnection.isConnected();
    }

    @Override
    public PlcConnectionMetadata getMetadata() {
        PlcConnection plcConnection = connection.get();
        if(plcConnection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        return plcConnection.getMetadata();
    }

    @Override
    public CompletableFuture<? extends PlcPingResponse> ping() {
        PlcConnection plcConnection = connection.get();
        if(plcConnection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        return plcConnection.ping();
    }

    @Override
    public PlcReadRequest.Builder readRequestBuilder() {
        PlcConnection plcConnection = connection.get();
        if(plcConnection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        return LeasedRequest.read(plcConnection.readRequestBuilder(), v -> invalidateConnection = true);
    }

    @Override
    public PlcWriteRequest.Builder writeRequestBuilder() {
        PlcConnection plcConnection = connection.get();
        if(plcConnection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        return LeasedRequest.write(plcConnection.writeRequestBuilder(), v -> invalidateConnection = true);
    }

    @Override
    public PlcSubscriptionRequest.Builder subscriptionRequestBuilder() {
        PlcConnection plcConnection = connection.get();
        if(plcConnection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        return LeasedRequest.subscription(plcConnection.subscriptionRequestBuilder(), v -> invalidateConnection = true);
    }

    @Override
    public PlcUnsubscriptionRequest.Builder unsubscriptionRequestBuilder() {
        PlcConnection plcConnection = connection.get();
        if(plcConnection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        return LeasedRequest.unsubscription(plcConnection.unsubscriptionRequestBuilder(), v -> invalidateConnection = true);
    }

    @Override
    public PlcBrowseRequest.Builder browseRequestBuilder() {
        PlcConnection plcConnection = connection.get();
        if(plcConnection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        return LeasedRequest.browse(plcConnection.browseRequestBuilder(), v -> invalidateConnection = true);
    }

    @Override
    public void addEventListener(EventListener listener) {
        connectionContainer.addEventListener(listener);
    }

    @Override
    public void removeEventListener(EventListener listener) {
        connectionContainer.removeEventListener(listener);
    }

}
