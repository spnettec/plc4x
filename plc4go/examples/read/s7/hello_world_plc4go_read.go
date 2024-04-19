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
	var connectionString = "s7://10.80.41.47"
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
		AddTagAddress("bool-value-1", "%DB1:0.0:BOOL"). // true
		AddTagAddress("bool-value-2", "%DB1:0.1:BOOL"). // false
		// It seems S7 PLCs ignores the array notation for BOOL
		//AddTagAddress("bool-array", "%DB1:2:BOOL[16]").
		AddTagAddress("byte-value", "%DB1:4:BYTE").
		AddTagAddress("byte-array", "%DB1:6:BYTE[2]").
		AddTagAddress("word-value", "%DB1:8:WORD").
		AddTagAddress("word-array", "%DB1:10:WORD[2]").
		AddTagAddress("dword-value", "%DB1:14:DWORD").
		AddTagAddress("dword-array", "%DB1:18:DWORD[2]").
		AddTagAddress("int-value", "%DB1:26:INT").                     // 23
		AddTagAddress("int-array", "%DB1:28:INT[2]").                  // 123, -142
		AddTagAddress("dint-value", "%DB1:32:DINT").                   // 24
		AddTagAddress("dint-array", "%DB1:36:DINT[2]").                // 1234, -2345
		AddTagAddress("real-value", "%DB1:44:REAL").                   // 3.14159
		AddTagAddress("real-array", "%DB1:48:REAL[2]").                // 12.345, 12.345
		AddTagAddress("string-value", "%DB1:56:STRING").               // "Hurz"
		AddTagAddress("string-array", "%DB1:312:STRING[2]").           // "Wolf", "Lamm"
		AddTagAddress("time-value", "%DB1:824:TIME").                  // 1234ms
		AddTagAddress("time-array", "%DB1:828:TIME[2]").               // 123ms, 234ms
		AddTagAddress("date-value", "%DB1:836:DATE").                  // D#2020-08-20
		AddTagAddress("date-array", "%DB1:838:DATE[2]").               // D#1990-03-28, D#2020-10-25
		AddTagAddress("time-of-day-value", "%DB1:842:TIME_OF_DAY").    // TOD#12:34:56
		AddTagAddress("time-of-day-array", "%DB1:846:TIME_OF_DAY[2]"). // TOD#16:34:56, TOD#08:15:00
		AddTagAddress("date-and-time-value", "%DB1:854:DTL").          // DTL#1978-03-28-12:34:56
		AddTagAddress("date-and-time-array", "%DB1:866:DTL[2]").       // DTL#1978-03-28-12:34:56, DTL#1978-03-28-12:34:56
		AddTagAddress("char-value", "%DB1:890:CHAR").                  // "H"
		AddTagAddress("char-array", "%DB1:892:CHAR[2]").               // "H", "u", "r", "z"
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
	for _, value := range rrr.GetResponse().GetTagNames() {
		if rrr.GetResponse().GetResponseCode(value) != model.PlcResponseCode_OK {
			fmt.Printf("error an non-ok,tag %s return code: %s\n", value, rrr.GetResponse().GetResponseCode(value).GetName())
			return
		}
	}
	for _, value := range rrr.GetResponse().GetTagNames() {
		fmt.Printf("Got  tag:%s, result:%s\n", value, rrr.GetResponse().GetValue(value).GetString())
	}
}
