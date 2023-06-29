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

// BACnetLiftCarDoorCommand is an enum
type BACnetLiftCarDoorCommand uint8

type IBACnetLiftCarDoorCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetLiftCarDoorCommand_NONE  BACnetLiftCarDoorCommand = 0
	BACnetLiftCarDoorCommand_OPEN  BACnetLiftCarDoorCommand = 1
	BACnetLiftCarDoorCommand_CLOSE BACnetLiftCarDoorCommand = 2
)

var BACnetLiftCarDoorCommandValues []BACnetLiftCarDoorCommand

func init() {
	_ = errors.New
	BACnetLiftCarDoorCommandValues = []BACnetLiftCarDoorCommand{
		BACnetLiftCarDoorCommand_NONE,
		BACnetLiftCarDoorCommand_OPEN,
		BACnetLiftCarDoorCommand_CLOSE,
	}
}

func BACnetLiftCarDoorCommandByValue(value uint8) (enum BACnetLiftCarDoorCommand, ok bool) {
	switch value {
	case 0:
		return BACnetLiftCarDoorCommand_NONE, true
	case 1:
		return BACnetLiftCarDoorCommand_OPEN, true
	case 2:
		return BACnetLiftCarDoorCommand_CLOSE, true
	}
	return 0, false
}

func BACnetLiftCarDoorCommandByName(value string) (enum BACnetLiftCarDoorCommand, ok bool) {
	switch value {
	case "NONE":
		return BACnetLiftCarDoorCommand_NONE, true
	case "OPEN":
		return BACnetLiftCarDoorCommand_OPEN, true
	case "CLOSE":
		return BACnetLiftCarDoorCommand_CLOSE, true
	}
	return 0, false
}

func BACnetLiftCarDoorCommandKnows(value uint8) bool {
	for _, typeValue := range BACnetLiftCarDoorCommandValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetLiftCarDoorCommand(structType any) BACnetLiftCarDoorCommand {
	castFunc := func(typ any) BACnetLiftCarDoorCommand {
		if sBACnetLiftCarDoorCommand, ok := typ.(BACnetLiftCarDoorCommand); ok {
			return sBACnetLiftCarDoorCommand
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetLiftCarDoorCommand) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetLiftCarDoorCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLiftCarDoorCommandParse(ctx context.Context, theBytes []byte) (BACnetLiftCarDoorCommand, error) {
	return BACnetLiftCarDoorCommandParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLiftCarDoorCommandParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLiftCarDoorCommand, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetLiftCarDoorCommand", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetLiftCarDoorCommand")
	}
	if enum, ok := BACnetLiftCarDoorCommandByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetLiftCarDoorCommand(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetLiftCarDoorCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetLiftCarDoorCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetLiftCarDoorCommand", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetLiftCarDoorCommand) PLC4XEnumName() string {
	switch e {
	case BACnetLiftCarDoorCommand_NONE:
		return "NONE"
	case BACnetLiftCarDoorCommand_OPEN:
		return "OPEN"
	case BACnetLiftCarDoorCommand_CLOSE:
		return "CLOSE"
	}
	return ""
}

func (e BACnetLiftCarDoorCommand) String() string {
	return e.PLC4XEnumName()
}
