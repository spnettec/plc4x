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

import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.api.exceptions.PlcRuntimeException;
import org.apache.plc4x.java.api.messages.*;
import org.apache.plc4x.java.api.metadata.PlcConnectionMetadata;
import org.apache.plc4x.java.api.model.PlcTag;
import org.apache.plc4x.java.api.value.PlcValue;

import java.time.Duration;
import java.time.LocalDateTime;
import java.time.ZoneId;
import java.util.*;
import java.util.concurrent.CompletableFuture;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.TimeoutException;
import java.util.function.BiFunction;

public class LeasedPlcConnection implements PlcConnection {

    private ConnectionContainer connectionContainer;
    private PlcConnection connection;
    public LeasedPlcConnection(ConnectionContainer connectionContainer, PlcConnection connection) {
        this.connectionContainer = connectionContainer;
        this.connection = connection;
    }

    @Override
    public synchronized void close() {
        if(connectionContainer == null) {
            return;
        }
        // Tell the connection container that the connection is free to be reused.
        connectionContainer.returnConnection(this);
        // Make the connection unusable.
        connection = null;
        connectionContainer = null;
    }
    public void destroy(){
        try {
            connection.close();
        } catch (Exception e) {
        }
        close();
        connectionContainer.close();
    }
    @Override
    public void connect() throws PlcConnectionException {
        throw new PlcConnectionException("Error connecting leased connection");
    }

    @Override
    public boolean isConnected() {
        if(connection == null) {
            return false;
        }
        return connection.isConnected();
    }

    @Override
    public PlcConnectionMetadata getMetadata() {
        if(connection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        return connection.getMetadata();
    }

    @Override
    public CompletableFuture<Void> ping() {
        if(connection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        return connection.ping();
    }

    @Override
    public PlcReadRequest.Builder readRequestBuilder() {
        if(connection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        final PlcReadRequest.Builder innerBuilder = connection.readRequestBuilder();
        return new PlcReadRequest.Builder(){

            @Override
            public PlcReadRequest build() {
                final PlcReadRequest innerPlcReadRequest = innerBuilder.build();
                return new PlcReadRequest(){

                    @Override
                    public CompletableFuture<? extends PlcReadResponse> execute() {
                        CompletableFuture<? extends PlcReadResponse> future = innerPlcReadRequest.execute();
                        final CompletableFuture<PlcReadResponse> responseFuture = new CompletableFuture<>();
                        future.handle((plcReadResponse, throwable) -> {
                            if (plcReadResponse != null) {
                                responseFuture.complete(plcReadResponse);
                            } else {
                                try {
                                    destroy();
                                } catch (Exception e) {
                                }
                                responseFuture.completeExceptionally(throwable);
                            }
                            return null;
                        });
                        return responseFuture;
                    }

                    @Override
                    public int getNumberOfTags() {
                        return innerPlcReadRequest.getNumberOfTags();
                    }

                    @Override
                    public LinkedHashSet<String> getTagNames() {
                        return innerPlcReadRequest.getTagNames();
                    }

                    @Override
                    public PlcTag getTag(String name) {
                        return innerPlcReadRequest.getTag(name);
                    }

                    @Override
                    public List<PlcTag> getTags() {
                        return innerPlcReadRequest.getTags();
                    }
                };
            }

            @Override
            public PlcReadRequest.Builder addTagAddress(String name, String tagAddress) {
                return innerBuilder.addTagAddress(name, tagAddress);
            }

            @Override
            public PlcReadRequest.Builder addTag(String name, PlcTag tag) {
                return innerBuilder.addTag(name,tag);
            }
        };
    }

    @Override
    public PlcWriteRequest.Builder writeRequestBuilder() {
        if(connection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        final PlcWriteRequest.Builder innerBuilder =  connection.writeRequestBuilder();
        return new PlcWriteRequest.Builder() {
            @Override
            public PlcWriteRequest build() {
                PlcWriteRequest innerPlcWriteRequest = innerBuilder.build();
                return new PlcWriteRequest() {
                    @Override
                    public CompletableFuture<? extends PlcWriteResponse> execute() {
                        CompletableFuture<? extends PlcWriteResponse> future = innerPlcWriteRequest.execute();
                        final CompletableFuture<PlcWriteResponse> responseFuture = new CompletableFuture<>();
                        future.handle((plcWriteResponse,throwable)->{
                            if (plcWriteResponse != null) {
                                responseFuture.complete(plcWriteResponse);
                            } else {
                                try {
                                    destroy();
                                } catch (Exception e) {
                                }
                                responseFuture.completeExceptionally(throwable);
                            }
                            return null;
                        });
                        return responseFuture;
                    }

                    @Override
                    public int getNumberOfValues(String name) {
                        return innerPlcWriteRequest.getNumberOfValues(name);
                    }

                    @Override
                    public PlcValue getPlcValue(String name) {
                        return innerPlcWriteRequest.getPlcValue(name);
                    }

                    @Override
                    public int getNumberOfTags() {
                        return innerPlcWriteRequest.getNumberOfTags();
                    }

                    @Override
                    public LinkedHashSet<String> getTagNames() {
                        return innerPlcWriteRequest.getTagNames();
                    }

                    @Override
                    public PlcTag getTag(String name) {
                        return innerPlcWriteRequest.getTag(name);
                    }

                    @Override
                    public List<PlcTag> getTags() {
                        return innerPlcWriteRequest.getTags();
                    }
                };
            }

            @Override
            public PlcWriteRequest.Builder addTagAddress(String name, String tagAddress, Object... values) {
                return innerBuilder.addTagAddress(name,tagAddress,values);
            }

            @Override
            public PlcWriteRequest.Builder addTag(String name, PlcTag tag, Object... values) {
                return innerBuilder.addTag(name,tag,values);
            }
        };
    }

    @Override
    public PlcSubscriptionRequest.Builder subscriptionRequestBuilder() {
        if(connection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        return connection.subscriptionRequestBuilder();
    }

    @Override
    public PlcUnsubscriptionRequest.Builder unsubscriptionRequestBuilder() {
        if(connection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        return connection.unsubscriptionRequestBuilder();
    }

    @Override
    public PlcBrowseRequest.Builder browseRequestBuilder() {
        if(connection == null) {
            throw new PlcRuntimeException("Error using leased connection after returning it to the cache.");
        }
        return connection.browseRequestBuilder();
    }

}
