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

// InterfaceAdminStatus is an enum
type InterfaceAdminStatus uint32

type IInterfaceAdminStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	InterfaceAdminStatus_interfaceAdminStatusUp InterfaceAdminStatus = 0
	InterfaceAdminStatus_interfaceAdminStatusDown InterfaceAdminStatus = 1
	InterfaceAdminStatus_interfaceAdminStatusTesting InterfaceAdminStatus = 2
)

var InterfaceAdminStatusValues []InterfaceAdminStatus

func init() {
	_ = errors.New
	InterfaceAdminStatusValues = []InterfaceAdminStatus {
		InterfaceAdminStatus_interfaceAdminStatusUp,
		InterfaceAdminStatus_interfaceAdminStatusDown,
		InterfaceAdminStatus_interfaceAdminStatusTesting,
	}
}

func InterfaceAdminStatusByValue(value uint32) (enum InterfaceAdminStatus, ok bool) {
	switch value {
		case 0:
			return InterfaceAdminStatus_interfaceAdminStatusUp, true
		case 1:
			return InterfaceAdminStatus_interfaceAdminStatusDown, true
		case 2:
			return InterfaceAdminStatus_interfaceAdminStatusTesting, true
	}
	return 0, false
}

func InterfaceAdminStatusByName(value string) (enum InterfaceAdminStatus, ok bool) {
	switch value {
	case "interfaceAdminStatusUp":
		return InterfaceAdminStatus_interfaceAdminStatusUp, true
	case "interfaceAdminStatusDown":
		return InterfaceAdminStatus_interfaceAdminStatusDown, true
	case "interfaceAdminStatusTesting":
		return InterfaceAdminStatus_interfaceAdminStatusTesting, true
	}
	return 0, false
}

func InterfaceAdminStatusKnows(value uint32)  bool {
	for _, typeValue := range InterfaceAdminStatusValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastInterfaceAdminStatus(structType any) InterfaceAdminStatus {
	castFunc := func(typ any) InterfaceAdminStatus {
		if sInterfaceAdminStatus, ok := typ.(InterfaceAdminStatus); ok {
			return sInterfaceAdminStatus
		}
		return 0
	}
	return castFunc(structType)
}

func (m InterfaceAdminStatus) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m InterfaceAdminStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func InterfaceAdminStatusParse(ctx context.Context, theBytes []byte) (InterfaceAdminStatus, error) {
	return InterfaceAdminStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func InterfaceAdminStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (InterfaceAdminStatus, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("InterfaceAdminStatus", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading InterfaceAdminStatus")
	}
	if enum, ok := InterfaceAdminStatusByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return InterfaceAdminStatus(val), nil
	} else {
		return enum, nil
	}
}

func (e InterfaceAdminStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e InterfaceAdminStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("InterfaceAdminStatus", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e InterfaceAdminStatus) PLC4XEnumName() string {
	switch e {
	case InterfaceAdminStatus_interfaceAdminStatusUp:
		return "interfaceAdminStatusUp"
	case InterfaceAdminStatus_interfaceAdminStatusDown:
		return "interfaceAdminStatusDown"
	case InterfaceAdminStatus_interfaceAdminStatusTesting:
		return "interfaceAdminStatusTesting"
	}
	return ""
}

func (e InterfaceAdminStatus) String() string {
	return e.PLC4XEnumName()
}

