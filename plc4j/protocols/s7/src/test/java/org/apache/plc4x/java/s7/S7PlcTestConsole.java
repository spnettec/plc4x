/*
Licensed to the Apache Software Foundation (ASF) under one
or more contributor license agreements.  See the NOTICE file
distributed with this work for additional information
regarding copyright ownership.  The ASF licenses this file
to you under the Apache License, Version 2.0 (the
"License"); you may not use this file except in compliance
with the License.  You may obtain a copy of the License at

  http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing,
software distributed under the License is distributed on an
"AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
KIND, either express or implied.  See the License for the
specific language governing permissions and limitations
under the License.
*/
package org.apache.plc4x.java.s7;

import org.apache.plc4x.java.PlcDriverManager;
import org.apache.plc4x.java.api.connection.PlcConnection;
import org.apache.plc4x.java.api.connection.PlcReader;
import org.apache.plc4x.java.api.messages.PlcReadRequest;
import org.apache.plc4x.java.api.messages.PlcReadResponse;
import org.apache.plc4x.java.api.model.Address;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.util.List;
import java.util.Optional;
import java.util.Scanner;

public class S7PlcTestConsole {

    private final static Logger logger = LoggerFactory.getLogger(S7PlcTestConsole.class);

    /**
     * Example code do demonstrate using the S7 Plc Driver.
     *
     * @param args ignored.
     * @throws Exception something went wrong.
     */
    public static void main(String[] args) throws Exception {
        // Create a connection to the S7 PLC (s7://{hostname/ip}/{racknumber}/{slotnumber})
        logger.info("Connecting");
        try (PlcConnection plcConnection = new PlcDriverManager().getConnection("s7://192.168.0.1/0/0")){
            logger.info("Connected");

            Optional<PlcReader> reader = plcConnection.getReader();
            // Check if this connection support reading of data.
            if (reader.isPresent()) {
                PlcReader plcReader = reader.get();

                Scanner scanner = new Scanner(System.in);
                String line;
                while(!"exit".equalsIgnoreCase(line = scanner.next())) {
                    try {
                        Address address = plcConnection.parseAddress(line);
                        PlcReadResponse plcReadResponse = plcReader.read(new PlcReadRequest(Byte.class, address)).get();
                        List<Object> data = plcReadResponse.getResponseItems().get(0).getValues();
                        System.out.println("Response: " + data.get(0));
                    } catch(Exception e) {
                        e.printStackTrace();
                    }
                }
            }
        }
        // Catch any exception or the application won't be able to finish if something goes wrong.
        catch (Exception e) {
            e.printStackTrace();
        }
        // The application would cleanly terminate after several seconds ... this just speeds things up.
        System.exit(0);
    }

}
