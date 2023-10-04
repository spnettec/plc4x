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

// BACnetAccessThreatLevel is the corresponding interface of BACnetAccessThreatLevel
type BACnetAccessThreatLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetThreatLevel returns ThreatLevel (property field)
	GetThreatLevel() BACnetApplicationTagUnsignedInteger
}

// BACnetAccessThreatLevelExactly can be used when we want exactly this type and not a type which fulfills BACnetAccessThreatLevel.
// This is useful for switch cases.
type BACnetAccessThreatLevelExactly interface {
	BACnetAccessThreatLevel
	isBACnetAccessThreatLevel() bool
}

// _BACnetAccessThreatLevel is the data-structure of this message
type _BACnetAccessThreatLevel struct {
	ThreatLevel BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccessThreatLevel) GetThreatLevel() BACnetApplicationTagUnsignedInteger {
	return m.ThreatLevel
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAccessThreatLevel factory function for _BACnetAccessThreatLevel
func NewBACnetAccessThreatLevel(threatLevel BACnetApplicationTagUnsignedInteger) *_BACnetAccessThreatLevel {
	return &_BACnetAccessThreatLevel{ThreatLevel: threatLevel}
}

// Deprecated: use the interface for direct cast
func CastBACnetAccessThreatLevel(structType any) BACnetAccessThreatLevel {
	if casted, ok := structType.(BACnetAccessThreatLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccessThreatLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccessThreatLevel) GetTypeName() string {
	return "BACnetAccessThreatLevel"
}

func (m *_BACnetAccessThreatLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (threatLevel)
	lengthInBits += m.ThreatLevel.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetAccessThreatLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessThreatLevelParse(ctx context.Context, theBytes []byte) (BACnetAccessThreatLevel, error) {
	return BACnetAccessThreatLevelParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetAccessThreatLevelParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetAccessThreatLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetAccessThreatLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccessThreatLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (threatLevel)
	if pullErr := readBuffer.PullContext("threatLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for threatLevel")
	}
	_threatLevel, _threatLevelErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _threatLevelErr != nil {
		return nil, errors.Wrap(_threatLevelErr, "Error parsing 'threatLevel' field of BACnetAccessThreatLevel")
	}
	threatLevel := _threatLevel.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("threatLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for threatLevel")
	}

	if closeErr := readBuffer.CloseContext("BACnetAccessThreatLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccessThreatLevel")
	}

	// Create the instance
	return &_BACnetAccessThreatLevel{
		ThreatLevel: threatLevel,
	}, nil
}

func (m *_BACnetAccessThreatLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccessThreatLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetAccessThreatLevel"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccessThreatLevel")
	}

	// Simple Field (threatLevel)
	if pushErr := writeBuffer.PushContext("threatLevel"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for threatLevel")
	}
	_threatLevelErr := writeBuffer.WriteSerializable(ctx, m.GetThreatLevel())
	if popErr := writeBuffer.PopContext("threatLevel"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for threatLevel")
	}
	if _threatLevelErr != nil {
		return errors.Wrap(_threatLevelErr, "Error serializing 'threatLevel' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAccessThreatLevel"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccessThreatLevel")
	}
	return nil
}

func (m *_BACnetAccessThreatLevel) isBACnetAccessThreatLevel() bool {
	return true
}

func (m *_BACnetAccessThreatLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
