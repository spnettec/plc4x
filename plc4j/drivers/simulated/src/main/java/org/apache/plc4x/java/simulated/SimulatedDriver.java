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
package org.apache.plc4x.java.simulated;

import org.apache.plc4x.java.api.PlcConnection;
import org.apache.plc4x.java.api.PlcDriver;
import org.apache.plc4x.java.api.authentication.PlcAuthentication;
import org.apache.plc4x.java.api.exceptions.PlcConnectionException;
import org.apache.plc4x.java.simulated.connection.SimulatedConnection;
import org.apache.plc4x.java.simulated.connection.SimulatedDevice;
import org.apache.plc4x.java.simulated.tag.SimulatedTag;
import org.apache.plc4x.java.spi.configuration.ConfigurationFactory;

import java.util.regex.Matcher;
import java.util.regex.Pattern;

/**
 * Test driver holding its state in the client process.
 * The URL schema is {@code simulated:<device_name>}.
 * Devices are created each time a connection is established and should not be reused.
 * Every device contains a random value generator accessible by address {@code random}.
 * Any value can be stored into test devices, however the state will be gone when connection is closed.
 */
public class SimulatedDriver implements PlcDriver {

    private static final Pattern URI_PATTERN = Pattern.compile(
        "^(?<protocolCode>[a-z0-9\\-]*)(:(?<transportCode>[a-z0-9]*))?:(?<transportConfig>[^?]*)(\\?(?<paramString>.*))?");
    @Override
    public String getProtocolCode() {
        return "simulated";
    }

    @Override
    public String getProtocolName() {
        return "Simulated PLC4X Datasource";
    }

    @Override
    public PlcConnection getConnection(String url) throws PlcConnectionException {
        // TODO: perform further checks
        Matcher matcher = URI_PATTERN.matcher(url);
        if (!matcher.matches()) {
            throw new PlcConnectionException(
                "Connection string doesn't match the format '{protocol-code}:({transport-code})?//{transport-address}(?{parameter-string)?'");
        }
        final String protocolCode = matcher.group("protocolCode");
        final String transportCode = matcher.group("transportCode");
        final String transportConfig = matcher.group("transportConfig");
        final String paramString = matcher.group("paramString");

        // Check if the protocol code matches this driver.
        if (!protocolCode.equals(getProtocolCode())) {
            // Actually this shouldn't happen as the DriverManager should have not used this driver in the first place.
            throw new PlcConnectionException(
                "This driver is not suited to handle this connection string");
        }

        if (transportConfig.isEmpty()) {
            throw new PlcConnectionException("Invalid URL: no device name given.");
        }
        // Create the configuration object.
        SimulatedConfiguration configuration = new ConfigurationFactory().createConfiguration(
            SimulatedConfiguration.class, protocolCode==null?"":protocolCode,
                transportCode == null?"":transportCode, transportConfig, paramString);
        if (configuration == null) {
            throw new PlcConnectionException("Unsupported configuration");
        }
        SimulatedDevice device = new SimulatedDevice(transportConfig,configuration);
        return new SimulatedConnection(device);
    }

    @Override
    public PlcConnection getConnection(String url, PlcAuthentication authentication) throws PlcConnectionException {
        throw new PlcConnectionException("Test driver does not support authentication.");
    }

    @Override
    public SimulatedTag prepareTag(String tagAddress){
        return SimulatedTag.of(tagAddress);
    }

}
