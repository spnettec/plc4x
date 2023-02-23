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

package readwrite

import (
	"context"
	"strings"
	"strconv"

	"github.com/apache/plc4x/plc4go/protocols/ads/readwrite/model"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type AdsXmlParserHelper struct {
}

// Temporary imports to silent compiler warnings (TODO: migrate from static to emission based imports)
func init() {
	_ = strconv.ParseUint
	_ = strconv.ParseInt
	_ = strings.Join
	_ = utils.Dump
}

func (m AdsXmlParserHelper) Parse(typeName string, xmlString string, parserArguments ...string) (interface{}, error) {
    switch typeName {
        case "AmsSerialFrame":
			return model.AmsSerialFrameParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "DataItem":
            plcValueType, _ := model.PlcValueTypeByName(parserArguments[0])
            parsedInt1, err := strconv.ParseInt(parserArguments[1], 10, 32)
            if err!=nil {
                return nil, err
            }
            stringLength := int32(parsedInt1)
			// TODO: find a way to parse the sub types
            var stringEncoding string
            return model.DataItemParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), plcValueType,  stringLength,  stringEncoding  )
        case "AdsTableSizes":
			return model.AdsTableSizesParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "AdsMultiRequestItem":
			parsedUint0, err := strconv.ParseUint(parserArguments[0], 10, 32)
			if err!=nil {
				return nil, err
			}
			indexGroup := uint32(parsedUint0)
            return model.AdsMultiRequestItemParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)), indexGroup  )
        case "AmsSerialAcknowledgeFrame":
			return model.AmsSerialAcknowledgeFrameParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "AdsDataTypeArrayInfo":
			return model.AdsDataTypeArrayInfoParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "AdsDataTypeTableEntry":
			return model.AdsDataTypeTableEntryParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "AmsNetId":
			return model.AmsNetIdParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "AdsStampHeader":
			return model.AdsStampHeaderParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "AmsSerialResetFrame":
			return model.AmsSerialResetFrameParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "AdsDataTypeTableChildEntry":
			return model.AdsDataTypeTableChildEntryParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "AdsConstants":
			return model.AdsConstantsParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "AdsNotificationSample":
			return model.AdsNotificationSampleParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "AdsSymbolTableEntry":
			return model.AdsSymbolTableEntryParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "AmsTCPPacket":
			return model.AmsTCPPacketParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
        case "AmsPacket":
			return model.AmsPacketParseWithBuffer(context.Background(), utils.NewXmlReadBuffer(strings.NewReader(xmlString)))
    }
    return nil, errors.Errorf("Unsupported type %s", typeName)
}
