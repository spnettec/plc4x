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


// BACnetShedLevelLevel is the corresponding interface of BACnetShedLevelLevel
type BACnetShedLevelLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetShedLevel
	// GetLevel returns Level (property field)
	GetLevel() BACnetContextTagUnsignedInteger
}

// BACnetShedLevelLevelExactly can be used when we want exactly this type and not a type which fulfills BACnetShedLevelLevel.
// This is useful for switch cases.
type BACnetShedLevelLevelExactly interface {
	BACnetShedLevelLevel
	isBACnetShedLevelLevel() bool
}

// _BACnetShedLevelLevel is the data-structure of this message
type _BACnetShedLevelLevel struct {
	*_BACnetShedLevel
        Level BACnetContextTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetShedLevelLevel) InitializeParent(parent BACnetShedLevel , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetShedLevelLevel)  GetParent() BACnetShedLevel {
	return m._BACnetShedLevel
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetShedLevelLevel) GetLevel() BACnetContextTagUnsignedInteger {
	return m.Level
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetShedLevelLevel factory function for _BACnetShedLevelLevel
func NewBACnetShedLevelLevel( level BACnetContextTagUnsignedInteger , peekedTagHeader BACnetTagHeader ) *_BACnetShedLevelLevel {
	_result := &_BACnetShedLevelLevel{
		Level: level,
    	_BACnetShedLevel: NewBACnetShedLevel(peekedTagHeader),
	}
	_result._BACnetShedLevel._BACnetShedLevelChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetShedLevelLevel(structType any) BACnetShedLevelLevel {
    if casted, ok := structType.(BACnetShedLevelLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetShedLevelLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetShedLevelLevel) GetTypeName() string {
	return "BACnetShedLevelLevel"
}

func (m *_BACnetShedLevelLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (level)
	lengthInBits += m.Level.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetShedLevelLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetShedLevelLevelParse(ctx context.Context, theBytes []byte) (BACnetShedLevelLevel, error) {
	return BACnetShedLevelLevelParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetShedLevelLevelParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetShedLevelLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetShedLevelLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetShedLevelLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (level)
	if pullErr := readBuffer.PullContext("level"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for level")
	}
_level, _levelErr := BACnetContextTagParseWithBuffer(ctx, readBuffer , uint8( uint8(1) ) , BACnetDataType( BACnetDataType_UNSIGNED_INTEGER ) )
	if _levelErr != nil {
		return nil, errors.Wrap(_levelErr, "Error parsing 'level' field of BACnetShedLevelLevel")
	}
	level := _level.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("level"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for level")
	}

	if closeErr := readBuffer.CloseContext("BACnetShedLevelLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetShedLevelLevel")
	}

	// Create a partially initialized instance
	_child := &_BACnetShedLevelLevel{
		_BACnetShedLevel: &_BACnetShedLevel{
		},
		Level: level,
	}
	_child._BACnetShedLevel._BACnetShedLevelChildRequirements = _child
	return _child, nil
}

func (m *_BACnetShedLevelLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetShedLevelLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetShedLevelLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetShedLevelLevel")
		}

	// Simple Field (level)
	if pushErr := writeBuffer.PushContext("level"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for level")
	}
	_levelErr := writeBuffer.WriteSerializable(ctx, m.GetLevel())
	if popErr := writeBuffer.PopContext("level"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for level")
	}
	if _levelErr != nil {
		return errors.Wrap(_levelErr, "Error serializing 'level' field")
	}

		if popErr := writeBuffer.PopContext("BACnetShedLevelLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetShedLevelLevel")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetShedLevelLevel) isBACnetShedLevelLevel() bool {
	return true
}

func (m *_BACnetShedLevelLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



