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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// EiPCommand is an enum
type EiPCommand uint16

type IEiPCommand interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	EiPCommand_RegisterSession   EiPCommand = 0x0065
	EiPCommand_UnregisterSession EiPCommand = 0x0066
	EiPCommand_SendRRData        EiPCommand = 0x006F
)

var EiPCommandValues []EiPCommand

func init() {
	_ = errors.New
	EiPCommandValues = []EiPCommand{
		EiPCommand_RegisterSession,
		EiPCommand_UnregisterSession,
		EiPCommand_SendRRData,
	}
}

func EiPCommandByValue(value uint16) EiPCommand {
	switch value {
	case 0x0065:
		return EiPCommand_RegisterSession
	case 0x0066:
		return EiPCommand_UnregisterSession
	case 0x006F:
		return EiPCommand_SendRRData
	}
	return 0
}

func EiPCommandByName(value string) (enum EiPCommand, ok bool) {
	ok = true
	switch value {
	case "RegisterSession":
		enum = EiPCommand_RegisterSession
	case "UnregisterSession":
		enum = EiPCommand_UnregisterSession
	case "SendRRData":
		enum = EiPCommand_SendRRData
	default:
		enum = 0
		ok = false
	}
	return
}

func EiPCommandKnows(value uint16) bool {
	for _, typeValue := range EiPCommandValues {
		if uint16(typeValue) == value {
			return true
		}
	}
	return false
}

func CastEiPCommand(structType interface{}) EiPCommand {
	castFunc := func(typ interface{}) EiPCommand {
		if sEiPCommand, ok := typ.(EiPCommand); ok {
			return sEiPCommand
		}
		return 0
	}
	return castFunc(structType)
}

func (m EiPCommand) GetLengthInBits() uint16 {
	return 16
}

func (m EiPCommand) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func EiPCommandParse(readBuffer utils.ReadBuffer) (EiPCommand, error) {
	val, err := readBuffer.ReadUint16("EiPCommand", 16)
	if err != nil {
		return 0, nil
	}
	return EiPCommandByValue(val), nil
}

func (e EiPCommand) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint16("EiPCommand", 16, uint16(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e EiPCommand) PLC4XEnumName() string {
	switch e {
	case EiPCommand_RegisterSession:
		return "RegisterSession"
	case EiPCommand_UnregisterSession:
		return "UnregisterSession"
	case EiPCommand_SendRRData:
		return "SendRRData"
	}
	return ""
}

func (e EiPCommand) String() string {
	return e.PLC4XEnumName()
}
