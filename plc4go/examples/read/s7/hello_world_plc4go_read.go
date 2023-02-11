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

package main

import (
	"fmt"

	"github.com/apache/plc4x/plc4go/pkg/api"
	"github.com/apache/plc4x/plc4go/pkg/api/drivers"
	"github.com/apache/plc4x/plc4go/pkg/api/model"
)

func main() {
	driverManager := plc4go.NewPlcDriverManager()
	drivers.RegisterS7Driver(driverManager)
	//var ip = "10.80.41.18"

	// Get a connection to a remote PLC
	var connectionString = "s7://10.166.11.18?remote-rack=0&remote-slot=1"
	crc := driverManager.GetConnection(connectionString)

	// Wait for the driver to connect (or not)
	connectionResult := <-crc
	if connectionResult.GetErr() != nil {
		fmt.Printf("error connecting to PLC: %s", connectionResult.GetErr().Error())
		return
	}
	connection := connectionResult.GetConnection()

	// Make sure the connection is closed at the end
	defer connection.BlockingClose()

	// Prepare a read-request
	readRequest, err := connection.ReadRequestBuilder().
		AddTagAddress("field", "%DB132:124:STRING(40)|GBK").
		Build()
	if err != nil {
		fmt.Printf("error preparing read-request: %s", connectionResult.GetErr().Error())
		return
	}

	// Execute a read-request
	rrc := readRequest.Execute()

	// Wait for the response to finish
	rrr := <-rrc
	if rrr.GetErr() != nil {
		fmt.Printf("error executing read-request: %s", rrr.GetErr().Error())
		return
	}

	// Do something with the response
	if rrr.GetResponse().GetResponseCode("field") != model.PlcResponseCode_OK {
		fmt.Printf("error an non-ok return code: %s", rrr.GetResponse().GetResponseCode("field").GetName())
		return
	}

	value := rrr.GetResponse().GetValue("field")
	fmt.Printf("Got result %s", value.GetString())
}
