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

package model

import (
	"context"
	"fmt"

	"github.com/apache/plc4x/plc4go/spi/utils"

	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// OpcuaNodeIdServicesVariableOrdered is an enum
type OpcuaNodeIdServicesVariableOrdered int32

type IOpcuaNodeIdServicesVariableOrdered interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	OpcuaNodeIdServicesVariableOrdered_OrderedListType_OrderedObject_Placeholder_NumberInList OpcuaNodeIdServicesVariableOrdered = 23521
	OpcuaNodeIdServicesVariableOrdered_OrderedListType_NodeVersion OpcuaNodeIdServicesVariableOrdered = 23525
)

var OpcuaNodeIdServicesVariableOrderedValues []OpcuaNodeIdServicesVariableOrdered

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableOrderedValues = []OpcuaNodeIdServicesVariableOrdered {
		OpcuaNodeIdServicesVariableOrdered_OrderedListType_OrderedObject_Placeholder_NumberInList,
		OpcuaNodeIdServicesVariableOrdered_OrderedListType_NodeVersion,
	}
}

func OpcuaNodeIdServicesVariableOrderedByValue(value int32) (enum OpcuaNodeIdServicesVariableOrdered, ok bool) {
	switch value {
		case 23521:
			return OpcuaNodeIdServicesVariableOrdered_OrderedListType_OrderedObject_Placeholder_NumberInList, true
		case 23525:
			return OpcuaNodeIdServicesVariableOrdered_OrderedListType_NodeVersion, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableOrderedByName(value string) (enum OpcuaNodeIdServicesVariableOrdered, ok bool) {
	switch value {
	case "OrderedListType_OrderedObject_Placeholder_NumberInList":
		return OpcuaNodeIdServicesVariableOrdered_OrderedListType_OrderedObject_Placeholder_NumberInList, true
	case "OrderedListType_NodeVersion":
		return OpcuaNodeIdServicesVariableOrdered_OrderedListType_NodeVersion, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableOrderedKnows(value int32)  bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableOrderedValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastOpcuaNodeIdServicesVariableOrdered(structType any) OpcuaNodeIdServicesVariableOrdered {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableOrdered {
		if sOpcuaNodeIdServicesVariableOrdered, ok := typ.(OpcuaNodeIdServicesVariableOrdered); ok {
			return sOpcuaNodeIdServicesVariableOrdered
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableOrdered) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableOrdered) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableOrderedParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableOrdered, error) {
	return OpcuaNodeIdServicesVariableOrderedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableOrderedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableOrdered, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableOrdered", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableOrdered")
	}
	if enum, ok := OpcuaNodeIdServicesVariableOrderedByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableOrdered")
		return OpcuaNodeIdServicesVariableOrdered(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableOrdered) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableOrdered) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableOrdered", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableOrdered) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableOrdered_OrderedListType_OrderedObject_Placeholder_NumberInList:
		return "OrderedListType_OrderedObject_Placeholder_NumberInList"
	case OpcuaNodeIdServicesVariableOrdered_OrderedListType_NodeVersion:
		return "OrderedListType_NodeVersion"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableOrdered) String() string {
	return e.PLC4XEnumName()
}

