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
package org.apache.plc4x.java.opcua;

import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.messages.*;
import org.apache.plc4x.java.api.model.PlcSubscriptionHandle;
import org.apache.plc4x.java.api.value.PlcValue;
import org.apache.plc4x.java.utils.cache.CachedPlcConnectionManager;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.Duration;
import java.util.Optional;
import java.util.Timer;
import java.util.TimerTask;
import java.util.concurrent.ExecutionException;
import java.util.function.Consumer;

/**
 * This class serves only as a manual entry point for ad-hoc tests of the OPC UA PLC4J driver.
 * <p>
 * <p>
 * The current version is tested against a public server, which is to be replaced later by a separate instance of the Milo framework.
 * Afterwards the code represented here will be used as an example for the introduction page.
 * <p>
 */
public class ManualPLC4XOpcua {
    private static final Logger LOGGER = LoggerFactory.getLogger(ManualPLC4XOpcua.class);
    public static void main(String[] args) throws Exception {
        String connectionString = "opcua:tcp://heyoulin-yofcmini.local:53530/OPCUA/SimulationServer";
        CachedPlcConnectionManager plcConnectionManager = CachedPlcConnectionManager.getBuilder().build();
        try (PlcConnection plcConnection = plcConnectionManager.getConnection(connectionString)) {

            final PlcReadRequest.Builder builder = plcConnection.readRequestBuilder();
            builder.addTagAddress("Counter1","ns=5;s=Counter1");
            builder.addTagAddress("Random1","ns=5;s=Random1");
            final PlcReadRequest readRequest = builder.build();
            final PlcReadResponse readResponse = readRequest.execute().get();
            System.out.println(readResponse.getAsPlcValue());


            final PlcSubscriptionRequest.Builder sbuilder = plcConnection.subscriptionRequestBuilder();

            sbuilder.addChangeOfStateTagAddress("Counter1","ns=5;s=Counter1", Duration.ofMillis(100));
            sbuilder.addChangeOfStateTagAddress("Random1","ns=5;s=Random1",Duration.ofMillis(100));
            PlcSubscriptionRequest subscriptionRequest = sbuilder.build();
            final PlcSubscriptionResponse subscriptionResponse = subscriptionRequest.execute().get();
            Optional<String> optional =  subscriptionResponse.getTagNames().stream().findFirst();
            if(optional.isPresent()){
                final PlcSubscriptionHandle subscriptionHandle = subscriptionResponse.getSubscriptionHandle(optional.get());
                subscriptionHandle.register(new ValueChangeHandler());
            }
            new Timer("time1").schedule(new TimerTask() {

                @Override
                public void run() {
                    try (PlcConnection plcConnection = plcConnectionManager.getConnection(connectionString)) {
                        PlcUnsubscriptionRequest.Builder unBuilder = plcConnection.unsubscriptionRequestBuilder();
                        for (String subscriptionName : subscriptionResponse.getTagNames()) {
                            final PlcSubscriptionHandle subscriptionHandle = subscriptionResponse.getSubscriptionHandle(subscriptionName);
                            unBuilder.addHandles(subscriptionHandle);
                        }
                        try {
                            unBuilder.build().execute().get();
                        } catch (InterruptedException | ExecutionException e) {
                            e.printStackTrace();
                        }
                        System.out.println("Unsubscription request completed");
                    } catch (Exception e){
                        e.printStackTrace();
                    }
                }
            }, 300000);

            new Timer("time2").schedule(new TimerTask() {

                @Override
                public void run() {
                    try (PlcConnection plcConnection = plcConnectionManager.getConnection(connectionString)) {
                        final PlcReadRequest.Builder builder = plcConnection.readRequestBuilder();
                        builder.addTagAddress("Counter1", "ns=5;s=Counter1");
                        builder.addTagAddress("Random1", "ns=5;s=Random1");
                        final PlcReadRequest readRequest = builder.build();
                        try {
                            final PlcReadResponse readResponse = readRequest.execute().get();
                            System.out.println(readResponse.getAsPlcValue());
                        } catch (InterruptedException | ExecutionException e) {
                            e.printStackTrace();
                        }
                    } catch (Exception e) {
                        e.printStackTrace();
                    }
                }
            }, 2000,2000);

        }

    }
    private static class ValueChangeHandler implements Consumer<PlcSubscriptionEvent> {

        @Override
        public void accept(PlcSubscriptionEvent plcSubscriptionEvent) {
            LOGGER.info("Incoming event:");
            for (String fieldName : plcSubscriptionEvent.getTagNames()) {
                final PlcValue plcValue = plcSubscriptionEvent.getPlcValue(fieldName);
                if(plcValue.isList()) {
                    StringBuilder sb = new StringBuilder(String.format("Field '%s' value:", fieldName));
                    for (PlcValue value : plcValue.getList()) {
                        sb.append(" ").append(value.getString());
                    }
                    LOGGER.info(sb.toString());
                } else if (plcValue.isStruct()) {
                    StringBuilder sb = new StringBuilder(String.format("Field '%s' value:", fieldName));
                    plcValue.getStruct().forEach((name, value) ->
                        sb.append(" ").append(name).append("=").append(value.getString())
                    );
                    LOGGER.info(sb.toString());
                } else {
                    LOGGER.info(String.format("Field '%s' value: %s", fieldName, plcValue.getString()));
                }
            }
        }
    }
}
