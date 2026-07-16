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
import org.apache.plc4x.java.api.authentication.PlcAuthentication;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.opcua.config.OpcuaConfiguration;
import org.apache.plc4x.java.opcua.context.OpcuaDriverContext;
import org.apache.plc4x.java.opcua.tag.OpcuaTag;
import org.apache.plc4x.java.spi.config.Configuration;
import org.apache.plc4x.java.spi.drivers.ConnectionBase;
import org.apache.plc4x.java.spi.drivers.DriverBase;
import org.apache.plc4x.java.spi.transports.api.TransportInstance;
import org.apache.plc4x.java.utils.auditlog.api.AuditLog;

import java.util.List;
import java.util.Optional;
import java.util.regex.Matcher;

public class OpcuaPlcDriver extends DriverBase {

    @Override
    public String getProtocolCode() {
        return "opcua";
    }

    @Override
    public String getProtocolName() {
        return "Opcua";
    }

    @Override
    protected Class<? extends Configuration> getConfigurationClass() {
        return OpcuaConfiguration.class;
    }

    @Override
    public Optional<String> getDefaultTransportCode() {
        return Optional.of("tcp");
    }

    @Override
    public List<String> getSupportedTransportCodes() {
        return List.of("tcp");
    }

    @Override
    protected boolean canRead() {
        return true;
    }

    @Override
    protected boolean canWrite() {
        return true;
    }

    @Override
    protected boolean canSubscribe() {
        return true;
    }

    /**
     * Extracts the transport endpoint path (e.g. {@code "/milo"}) from the
     * connection URL and passes it to the {@link OpcuaConnection} so the
     * OPC UA HELLO message carries the correct endpoint URL.
     *
     * <p>The new SPI only hands {@code host:port} to the TCP transport —
     * the path part is lost.  The pre-merge code parsed it inside
     * {@link OpcuaDriverContext#setConfiguration} which received the full
     * URL; we replicate that logic here.</p>
     */
    @Override
    public PlcConnection getConnection(String connectionString, PlcAuthentication authentication)
            throws PlcConnectionException {
        // Parse transport endpoint path before the parent strips it
        String transportEndpoint = "";
        Matcher matcher = OpcuaDriverContext.URI_PATTERN.matcher(connectionString);
        if (matcher.matches()) {
            String te = matcher.group("transportEndpoint");
            if (te != null) {
                transportEndpoint = te;
            }
        }

        PlcConnection connection = super.getConnection(connectionString, authentication);
        if (connection instanceof OpcuaConnection opcua) {
            opcua.setTransportEndpoint(transportEndpoint);
            if (authentication != null) {
                opcua.setPlcAuthentication(authentication);
            }
        }
        return connection;
    }

    @Override
    protected boolean canBrowse() {
        return true;
    }

    @Override
    protected ConnectionBase<?> getConnection(Configuration configuration,
                                              TransportInstance<?> transportInstance,
                                              AuditLog auditLog) {
        return new OpcuaConnection((OpcuaConfiguration) configuration, transportInstance, auditLog);
    }

    @Override
    public OpcuaTag prepareTag(String tagAddress) {
        return OpcuaTag.of(tagAddress);
    }

}
