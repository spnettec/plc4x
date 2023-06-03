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

package tests

import (
	"testing"

	"github.com/apache/plc4x/plc4go/internal/eip"
	eipIO "github.com/apache/plc4x/plc4go/protocols/eip/readwrite"
	readWriteModel "github.com/apache/plc4x/plc4go/protocols/eip/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/options"
	"github.com/apache/plc4x/plc4go/spi/testutils"
)

func TestEIPDriver(t *testing.T) {
	testutils.SetToTestingLogger(t, readWriteModel.Plc4xModelLog)
	testutils.RunDriverTestsuite(
		t,
		eip.NewDriver(options.WithCustomLogger(testutils.ProduceTestingLogger(t))),
		"assets/testing/protocols/eip/DriverTestsuite.xml",
		eipIO.EipXmlParserHelper{},
	)
}
