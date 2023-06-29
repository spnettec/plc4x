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

// BaudRateSelector is an enum
type BaudRateSelector uint8

type IBaudRateSelector interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BaudRateSelector_SELECTED_4800_BAUD BaudRateSelector = 0x01
	BaudRateSelector_SELECTED_2400_BAUD BaudRateSelector = 0x02
	BaudRateSelector_SELECTED_1200_BAUD BaudRateSelector = 0x03
	BaudRateSelector_SELECTED_600_BAUD  BaudRateSelector = 0x04
	BaudRateSelector_SELECTED_300_BAUD  BaudRateSelector = 0x05
	BaudRateSelector_SELECTED_9600_BAUD BaudRateSelector = 0xFF
)

var BaudRateSelectorValues []BaudRateSelector

func init() {
	_ = errors.New
	BaudRateSelectorValues = []BaudRateSelector{
		BaudRateSelector_SELECTED_4800_BAUD,
		BaudRateSelector_SELECTED_2400_BAUD,
		BaudRateSelector_SELECTED_1200_BAUD,
		BaudRateSelector_SELECTED_600_BAUD,
		BaudRateSelector_SELECTED_300_BAUD,
		BaudRateSelector_SELECTED_9600_BAUD,
	}
}

func BaudRateSelectorByValue(value uint8) (enum BaudRateSelector, ok bool) {
	switch value {
	case 0x01:
		return BaudRateSelector_SELECTED_4800_BAUD, true
	case 0x02:
		return BaudRateSelector_SELECTED_2400_BAUD, true
	case 0x03:
		return BaudRateSelector_SELECTED_1200_BAUD, true
	case 0x04:
		return BaudRateSelector_SELECTED_600_BAUD, true
	case 0x05:
		return BaudRateSelector_SELECTED_300_BAUD, true
	case 0xFF:
		return BaudRateSelector_SELECTED_9600_BAUD, true
	}
	return 0, false
}

func BaudRateSelectorByName(value string) (enum BaudRateSelector, ok bool) {
	switch value {
	case "SELECTED_4800_BAUD":
		return BaudRateSelector_SELECTED_4800_BAUD, true
	case "SELECTED_2400_BAUD":
		return BaudRateSelector_SELECTED_2400_BAUD, true
	case "SELECTED_1200_BAUD":
		return BaudRateSelector_SELECTED_1200_BAUD, true
	case "SELECTED_600_BAUD":
		return BaudRateSelector_SELECTED_600_BAUD, true
	case "SELECTED_300_BAUD":
		return BaudRateSelector_SELECTED_300_BAUD, true
	case "SELECTED_9600_BAUD":
		return BaudRateSelector_SELECTED_9600_BAUD, true
	}
	return 0, false
}

func BaudRateSelectorKnows(value uint8) bool {
	for _, typeValue := range BaudRateSelectorValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBaudRateSelector(structType any) BaudRateSelector {
	castFunc := func(typ any) BaudRateSelector {
		if sBaudRateSelector, ok := typ.(BaudRateSelector); ok {
			return sBaudRateSelector
		}
		return 0
	}
	return castFunc(structType)
}

func (m BaudRateSelector) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BaudRateSelector) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BaudRateSelectorParse(ctx context.Context, theBytes []byte) (BaudRateSelector, error) {
	return BaudRateSelectorParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BaudRateSelectorParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BaudRateSelector, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BaudRateSelector", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BaudRateSelector")
	}
	if enum, ok := BaudRateSelectorByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return BaudRateSelector(val), nil
	} else {
		return enum, nil
	}
}

func (e BaudRateSelector) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BaudRateSelector) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BaudRateSelector", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BaudRateSelector) PLC4XEnumName() string {
	switch e {
	case BaudRateSelector_SELECTED_4800_BAUD:
		return "SELECTED_4800_BAUD"
	case BaudRateSelector_SELECTED_2400_BAUD:
		return "SELECTED_2400_BAUD"
	case BaudRateSelector_SELECTED_1200_BAUD:
		return "SELECTED_1200_BAUD"
	case BaudRateSelector_SELECTED_600_BAUD:
		return "SELECTED_600_BAUD"
	case BaudRateSelector_SELECTED_300_BAUD:
		return "SELECTED_300_BAUD"
	case BaudRateSelector_SELECTED_9600_BAUD:
		return "SELECTED_9600_BAUD"
	}
	return ""
}

func (e BaudRateSelector) String() string {
	return e.PLC4XEnumName()
}
