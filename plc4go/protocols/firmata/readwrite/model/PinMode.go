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

// PinMode is an enum
type PinMode uint8

type IPinMode interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	PinMode_PinModeInput   PinMode = 0x0
	PinMode_PinModeOutput  PinMode = 0x1
	PinMode_PinModeAnalog  PinMode = 0x2
	PinMode_PinModePwm     PinMode = 0x3
	PinMode_PinModeServo   PinMode = 0x4
	PinMode_PinModeShift   PinMode = 0x5
	PinMode_PinModeI2C     PinMode = 0x6
	PinMode_PinModeOneWire PinMode = 0x7
	PinMode_PinModeStepper PinMode = 0x8
	PinMode_PinModeEncoder PinMode = 0x9
	PinMode_PinModeSerial  PinMode = 0xA
	PinMode_PinModePullup  PinMode = 0xB
)

var PinModeValues []PinMode

func init() {
	_ = errors.New
	PinModeValues = []PinMode{
		PinMode_PinModeInput,
		PinMode_PinModeOutput,
		PinMode_PinModeAnalog,
		PinMode_PinModePwm,
		PinMode_PinModeServo,
		PinMode_PinModeShift,
		PinMode_PinModeI2C,
		PinMode_PinModeOneWire,
		PinMode_PinModeStepper,
		PinMode_PinModeEncoder,
		PinMode_PinModeSerial,
		PinMode_PinModePullup,
	}
}

func PinModeByValue(value uint8) PinMode {
	switch value {
	case 0x0:
		return PinMode_PinModeInput
	case 0x1:
		return PinMode_PinModeOutput
	case 0x2:
		return PinMode_PinModeAnalog
	case 0x3:
		return PinMode_PinModePwm
	case 0x4:
		return PinMode_PinModeServo
	case 0x5:
		return PinMode_PinModeShift
	case 0x6:
		return PinMode_PinModeI2C
	case 0x7:
		return PinMode_PinModeOneWire
	case 0x8:
		return PinMode_PinModeStepper
	case 0x9:
		return PinMode_PinModeEncoder
	case 0xA:
		return PinMode_PinModeSerial
	case 0xB:
		return PinMode_PinModePullup
	}
	return 0
}

func PinModeByName(value string) (enum PinMode, ok bool) {
	ok = true
	switch value {
	case "PinModeInput":
		enum = PinMode_PinModeInput
	case "PinModeOutput":
		enum = PinMode_PinModeOutput
	case "PinModeAnalog":
		enum = PinMode_PinModeAnalog
	case "PinModePwm":
		enum = PinMode_PinModePwm
	case "PinModeServo":
		enum = PinMode_PinModeServo
	case "PinModeShift":
		enum = PinMode_PinModeShift
	case "PinModeI2C":
		enum = PinMode_PinModeI2C
	case "PinModeOneWire":
		enum = PinMode_PinModeOneWire
	case "PinModeStepper":
		enum = PinMode_PinModeStepper
	case "PinModeEncoder":
		enum = PinMode_PinModeEncoder
	case "PinModeSerial":
		enum = PinMode_PinModeSerial
	case "PinModePullup":
		enum = PinMode_PinModePullup
	default:
		enum = 0
		ok = false
	}
	return
}

func PinModeKnows(value uint8) bool {
	for _, typeValue := range PinModeValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastPinMode(structType interface{}) PinMode {
	castFunc := func(typ interface{}) PinMode {
		if sPinMode, ok := typ.(PinMode); ok {
			return sPinMode
		}
		return 0
	}
	return castFunc(structType)
}

func (m PinMode) GetLengthInBits() uint16 {
	return 8
}

func (m PinMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func PinModeParse(readBuffer utils.ReadBuffer) (PinMode, error) {
	val, err := readBuffer.ReadUint8("PinMode", 8)
	if err != nil {
		return 0, nil
	}
	return PinModeByValue(val), nil
}

func (e PinMode) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("PinMode", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e PinMode) PLC4XEnumName() string {
	switch e {
	case PinMode_PinModeInput:
		return "PinModeInput"
	case PinMode_PinModeOutput:
		return "PinModeOutput"
	case PinMode_PinModeAnalog:
		return "PinModeAnalog"
	case PinMode_PinModePwm:
		return "PinModePwm"
	case PinMode_PinModeServo:
		return "PinModeServo"
	case PinMode_PinModeShift:
		return "PinModeShift"
	case PinMode_PinModeI2C:
		return "PinModeI2C"
	case PinMode_PinModeOneWire:
		return "PinModeOneWire"
	case PinMode_PinModeStepper:
		return "PinModeStepper"
	case PinMode_PinModeEncoder:
		return "PinModeEncoder"
	case PinMode_PinModeSerial:
		return "PinModeSerial"
	case PinMode_PinModePullup:
		return "PinModePullup"
	}
	return ""
}

func (e PinMode) String() string {
	return e.PLC4XEnumName()
}
