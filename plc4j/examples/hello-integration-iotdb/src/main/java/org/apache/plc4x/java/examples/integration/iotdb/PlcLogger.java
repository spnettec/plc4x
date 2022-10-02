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
package org.apache.plc4x.java.examples.integration.iotdb;

import org.apache.edgent.providers.direct.DirectProvider;
import org.apache.edgent.topology.Topology;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

/**
 * using this example, you can store data one by one into IoTDB.
 *
 * modified according to hello-integration-edgent
 *
 * arguments example:
 * --connection-string simulated://127.0.0.1 --field-address RANDOM/foo:Integer  --polling-interval 1000
 * --iotdb-address 127.0.0.1:6667 --iotdb-user-name root --iotdb-user-password root --iotdb-sg mi
 * --iotdb-device d1 --iotdb-datatype INT32
 */
public class PlcLogger {

    private static final Logger LOGGER = LoggerFactory.getLogger(PlcLogger.class);

    //Time series ID
    static String timeSeries;

    //device ID
    static String deviceId;

    //sensor ID
    static String sensor;

    static String dataType;

    static IIoTDBWriter ioTDBWriter = null;

    static boolean useJDBC;

    public static void main(String[] args) throws Exception {
    }

}
