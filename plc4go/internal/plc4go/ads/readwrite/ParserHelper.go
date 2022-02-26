/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

package readwrite

import (
	"github.com/apache/plc4x/plc4go/internal/plc4go/ads/readwrite/model"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type AdsParserHelper struct {
}

func (m AdsParserHelper) Parse(typeName string, arguments []string, io utils.ReadBuffer) (interface{}, error) {
	switch typeName {
	case "AmsSerialFrame":
		return model.AmsSerialFrameParse(io)
	case "DataItem":
		dataFormatName, err := utils.StrToString(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		stringLength, err := utils.StrToInt32(arguments[1])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		stringEncoding, err := utils.StrToString(arguments[2])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.DataItemParse(io, dataFormatName, stringLength, stringEncoding)
	case "AdsMultiRequestItem":
		indexGroup, err := utils.StrToUint32(arguments[0])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.AdsMultiRequestItemParse(io, indexGroup)
	case "AmsSerialAcknowledgeFrame":
		return model.AmsSerialAcknowledgeFrameParse(io)
	case "AdsData":
		var commandId model.CommandId
		response, err := utils.StrToBool(arguments[1])
		if err != nil {
			return nil, errors.Wrap(err, "Error parsing")
		}
		return model.AdsDataParse(io, commandId, response)
	case "AmsNetId":
		return model.AmsNetIdParse(io)
	case "AdsStampHeader":
		return model.AdsStampHeaderParse(io)
	case "AmsSerialResetFrame":
		return model.AmsSerialResetFrameParse(io)
	case "AdsNotificationSample":
		return model.AdsNotificationSampleParse(io)
	case "AmsTCPPacket":
		return model.AmsTCPPacketParse(io)
	case "State":
		return model.StateParse(io)
	case "AmsPacket":
		return model.AmsPacketParse(io)
	}
	return nil, errors.Errorf("Unsupported type %s", typeName)
}
