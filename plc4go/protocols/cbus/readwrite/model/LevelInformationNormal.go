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

// LevelInformationNormal is the corresponding interface of LevelInformationNormal
type LevelInformationNormal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	LevelInformation
	// GetPair1 returns Pair1 (property field)
	GetPair1() LevelInformationNibblePair
	// GetPair2 returns Pair2 (property field)
	GetPair2() LevelInformationNibblePair
	// GetActualLevel returns ActualLevel (virtual field)
	GetActualLevel() uint8
	// GetActualLevelInPercent returns ActualLevelInPercent (virtual field)
	GetActualLevelInPercent() float32
}

// LevelInformationNormalExactly can be used when we want exactly this type and not a type which fulfills LevelInformationNormal.
// This is useful for switch cases.
type LevelInformationNormalExactly interface {
	LevelInformationNormal
	isLevelInformationNormal() bool
}

// _LevelInformationNormal is the data-structure of this message
type _LevelInformationNormal struct {
	*_LevelInformation
	Pair1 LevelInformationNibblePair
	Pair2 LevelInformationNibblePair
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LevelInformationNormal) InitializeParent(parent LevelInformation, raw uint16) {
	m.Raw = raw
}

func (m *_LevelInformationNormal) GetParent() LevelInformation {
	return m._LevelInformation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LevelInformationNormal) GetPair1() LevelInformationNibblePair {
	return m.Pair1
}

func (m *_LevelInformationNormal) GetPair2() LevelInformationNibblePair {
	return m.Pair2
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_LevelInformationNormal) GetActualLevel() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPair2().NibbleValue()<<uint8(4) | m.GetPair1().NibbleValue())
}

func (m *_LevelInformationNormal) GetActualLevelInPercent() float32 {
	ctx := context.Background()
	_ = ctx
	return float32(float32(float32(float32(100))*float32((float32(m.GetActualLevel())+float32(float32(2))))) / float32(float32(255)))
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLevelInformationNormal factory function for _LevelInformationNormal
func NewLevelInformationNormal(pair1 LevelInformationNibblePair, pair2 LevelInformationNibblePair, raw uint16) *_LevelInformationNormal {
	_result := &_LevelInformationNormal{
		Pair1:             pair1,
		Pair2:             pair2,
		_LevelInformation: NewLevelInformation(raw),
	}
	_result._LevelInformation._LevelInformationChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLevelInformationNormal(structType any) LevelInformationNormal {
	if casted, ok := structType.(LevelInformationNormal); ok {
		return casted
	}
	if casted, ok := structType.(*LevelInformationNormal); ok {
		return *casted
	}
	return nil
}

func (m *_LevelInformationNormal) GetTypeName() string {
	return "LevelInformationNormal"
}

func (m *_LevelInformationNormal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (pair1)
	lengthInBits += 8

	// Simple field (pair2)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_LevelInformationNormal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LevelInformationNormalParse(ctx context.Context, theBytes []byte) (LevelInformationNormal, error) {
	return LevelInformationNormalParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func LevelInformationNormalParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LevelInformationNormal, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("LevelInformationNormal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LevelInformationNormal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (pair1)
	if pullErr := readBuffer.PullContext("pair1"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for pair1")
	}
	_pair1, _pair1Err := LevelInformationNibblePairParseWithBuffer(ctx, readBuffer)
	if _pair1Err != nil {
		return nil, errors.Wrap(_pair1Err, "Error parsing 'pair1' field of LevelInformationNormal")
	}
	pair1 := _pair1
	if closeErr := readBuffer.CloseContext("pair1"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for pair1")
	}

	// Simple Field (pair2)
	if pullErr := readBuffer.PullContext("pair2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for pair2")
	}
	_pair2, _pair2Err := LevelInformationNibblePairParseWithBuffer(ctx, readBuffer)
	if _pair2Err != nil {
		return nil, errors.Wrap(_pair2Err, "Error parsing 'pair2' field of LevelInformationNormal")
	}
	pair2 := _pair2
	if closeErr := readBuffer.CloseContext("pair2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for pair2")
	}

	// Virtual field
	_actualLevel := pair2.NibbleValue()<<uint8(4) | pair1.NibbleValue()
	actualLevel := uint8(_actualLevel)
	_ = actualLevel

	// Virtual field
	_actualLevelInPercent := float32(float32(float32(100))*float32((float32(actualLevel)+float32(float32(2))))) / float32(float32(255))
	actualLevelInPercent := float32(_actualLevelInPercent)
	_ = actualLevelInPercent

	if closeErr := readBuffer.CloseContext("LevelInformationNormal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LevelInformationNormal")
	}

	// Create a partially initialized instance
	_child := &_LevelInformationNormal{
		_LevelInformation: &_LevelInformation{},
		Pair1:             pair1,
		Pair2:             pair2,
	}
	_child._LevelInformation._LevelInformationChildRequirements = _child
	return _child, nil
}

func (m *_LevelInformationNormal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LevelInformationNormal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LevelInformationNormal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LevelInformationNormal")
		}

		// Simple Field (pair1)
		if pushErr := writeBuffer.PushContext("pair1"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for pair1")
		}
		_pair1Err := writeBuffer.WriteSerializable(ctx, m.GetPair1())
		if popErr := writeBuffer.PopContext("pair1"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for pair1")
		}
		if _pair1Err != nil {
			return errors.Wrap(_pair1Err, "Error serializing 'pair1' field")
		}

		// Simple Field (pair2)
		if pushErr := writeBuffer.PushContext("pair2"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for pair2")
		}
		_pair2Err := writeBuffer.WriteSerializable(ctx, m.GetPair2())
		if popErr := writeBuffer.PopContext("pair2"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for pair2")
		}
		if _pair2Err != nil {
			return errors.Wrap(_pair2Err, "Error serializing 'pair2' field")
		}
		// Virtual field
		if _actualLevelErr := writeBuffer.WriteVirtual(ctx, "actualLevel", m.GetActualLevel()); _actualLevelErr != nil {
			return errors.Wrap(_actualLevelErr, "Error serializing 'actualLevel' field")
		}
		// Virtual field
		if _actualLevelInPercentErr := writeBuffer.WriteVirtual(ctx, "actualLevelInPercent", m.GetActualLevelInPercent()); _actualLevelInPercentErr != nil {
			return errors.Wrap(_actualLevelInPercentErr, "Error serializing 'actualLevelInPercent' field")
		}

		if popErr := writeBuffer.PopContext("LevelInformationNormal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LevelInformationNormal")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LevelInformationNormal) isLevelInformationNormal() bool {
	return true
}

func (m *_LevelInformationNormal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
