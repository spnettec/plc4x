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

// BACnetAccessRuleTimeRangeSpecifier is an enum
type BACnetAccessRuleTimeRangeSpecifier uint8

type IBACnetAccessRuleTimeRangeSpecifier interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	BACnetAccessRuleTimeRangeSpecifier_SPECIFIED BACnetAccessRuleTimeRangeSpecifier = 0
	BACnetAccessRuleTimeRangeSpecifier_ALWAYS    BACnetAccessRuleTimeRangeSpecifier = 1
)

var BACnetAccessRuleTimeRangeSpecifierValues []BACnetAccessRuleTimeRangeSpecifier

func init() {
	_ = errors.New
	BACnetAccessRuleTimeRangeSpecifierValues = []BACnetAccessRuleTimeRangeSpecifier{
		BACnetAccessRuleTimeRangeSpecifier_SPECIFIED,
		BACnetAccessRuleTimeRangeSpecifier_ALWAYS,
	}
}

func BACnetAccessRuleTimeRangeSpecifierByValue(value uint8) (enum BACnetAccessRuleTimeRangeSpecifier, ok bool) {
	switch value {
	case 0:
		return BACnetAccessRuleTimeRangeSpecifier_SPECIFIED, true
	case 1:
		return BACnetAccessRuleTimeRangeSpecifier_ALWAYS, true
	}
	return 0, false
}

func BACnetAccessRuleTimeRangeSpecifierByName(value string) (enum BACnetAccessRuleTimeRangeSpecifier, ok bool) {
	switch value {
	case "SPECIFIED":
		return BACnetAccessRuleTimeRangeSpecifier_SPECIFIED, true
	case "ALWAYS":
		return BACnetAccessRuleTimeRangeSpecifier_ALWAYS, true
	}
	return 0, false
}

func BACnetAccessRuleTimeRangeSpecifierKnows(value uint8) bool {
	for _, typeValue := range BACnetAccessRuleTimeRangeSpecifierValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetAccessRuleTimeRangeSpecifier(structType any) BACnetAccessRuleTimeRangeSpecifier {
	castFunc := func(typ any) BACnetAccessRuleTimeRangeSpecifier {
		if sBACnetAccessRuleTimeRangeSpecifier, ok := typ.(BACnetAccessRuleTimeRangeSpecifier); ok {
			return sBACnetAccessRuleTimeRangeSpecifier
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetAccessRuleTimeRangeSpecifier) GetLengthInBits(ctx context.Context) uint16 {
	return 8
}

func (m BACnetAccessRuleTimeRangeSpecifier) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessRuleTimeRangeSpecifierParse(ctx context.Context, theBytes []byte) (BACnetAccessRuleTimeRangeSpecifier, error) {
	return BACnetAccessRuleTimeRangeSpecifierParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAccessRuleTimeRangeSpecifierParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessRuleTimeRangeSpecifier, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint8("BACnetAccessRuleTimeRangeSpecifier", 8)
	if err != nil {
		return 0, errors.Wrap(err, "error reading BACnetAccessRuleTimeRangeSpecifier")
	}
	if enum, ok := BACnetAccessRuleTimeRangeSpecifierByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return BACnetAccessRuleTimeRangeSpecifier(val), nil
	} else {
		return enum, nil
	}
}

func (e BACnetAccessRuleTimeRangeSpecifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e BACnetAccessRuleTimeRangeSpecifier) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint8("BACnetAccessRuleTimeRangeSpecifier", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e BACnetAccessRuleTimeRangeSpecifier) PLC4XEnumName() string {
	switch e {
	case BACnetAccessRuleTimeRangeSpecifier_SPECIFIED:
		return "SPECIFIED"
	case BACnetAccessRuleTimeRangeSpecifier_ALWAYS:
		return "ALWAYS"
	}
	return ""
}

func (e BACnetAccessRuleTimeRangeSpecifier) String() string {
	return e.PLC4XEnumName()
}
