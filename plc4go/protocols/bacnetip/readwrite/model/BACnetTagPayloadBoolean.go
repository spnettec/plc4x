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

// BACnetTagPayloadBoolean is the corresponding interface of BACnetTagPayloadBoolean
type BACnetTagPayloadBoolean interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetValue returns Value (virtual field)
	GetValue() bool
	// GetIsTrue returns IsTrue (virtual field)
	GetIsTrue() bool
	// GetIsFalse returns IsFalse (virtual field)
	GetIsFalse() bool
}

// BACnetTagPayloadBooleanExactly can be used when we want exactly this type and not a type which fulfills BACnetTagPayloadBoolean.
// This is useful for switch cases.
type BACnetTagPayloadBooleanExactly interface {
	BACnetTagPayloadBoolean
	isBACnetTagPayloadBoolean() bool
}

// _BACnetTagPayloadBoolean is the data-structure of this message
type _BACnetTagPayloadBoolean struct {

	// Arguments.
	ActualLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetTagPayloadBoolean) GetValue() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.ActualLength) == (1)))
}

func (m *_BACnetTagPayloadBoolean) GetIsTrue() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetValue())
}

func (m *_BACnetTagPayloadBoolean) GetIsFalse() bool {
	ctx := context.Background()
	_ = ctx
	return bool(!(m.GetValue()))
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTagPayloadBoolean factory function for _BACnetTagPayloadBoolean
func NewBACnetTagPayloadBoolean(actualLength uint32) *_BACnetTagPayloadBoolean {
	return &_BACnetTagPayloadBoolean{ActualLength: actualLength}
}

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadBoolean(structType any) BACnetTagPayloadBoolean {
	if casted, ok := structType.(BACnetTagPayloadBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadBoolean) GetTypeName() string {
	return "BACnetTagPayloadBoolean"
}

func (m *_BACnetTagPayloadBoolean) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetTagPayloadBoolean) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTagPayloadBooleanParse(ctx context.Context, theBytes []byte, actualLength uint32) (BACnetTagPayloadBoolean, error) {
	return BACnetTagPayloadBooleanParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), actualLength)
}

func BACnetTagPayloadBooleanParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32) (BACnetTagPayloadBoolean, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetTagPayloadBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_value := bool((actualLength) == (1))
	value := bool(_value)
	_ = value

	// Virtual field
	_isTrue := value
	isTrue := bool(_isTrue)
	_ = isTrue

	// Virtual field
	_isFalse := !(value)
	isFalse := bool(_isFalse)
	_ = isFalse

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadBoolean")
	}

	// Create the instance
	return &_BACnetTagPayloadBoolean{
		ActualLength: actualLength,
	}, nil
}

func (m *_BACnetTagPayloadBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadBoolean) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetTagPayloadBoolean"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadBoolean")
	}
	// Virtual field
	value := m.GetValue()
	_ = value
	if _valueErr := writeBuffer.WriteVirtual(ctx, "value", m.GetValue()); _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}
	// Virtual field
	isTrue := m.GetIsTrue()
	_ = isTrue
	if _isTrueErr := writeBuffer.WriteVirtual(ctx, "isTrue", m.GetIsTrue()); _isTrueErr != nil {
		return errors.Wrap(_isTrueErr, "Error serializing 'isTrue' field")
	}
	// Virtual field
	isFalse := m.GetIsFalse()
	_ = isFalse
	if _isFalseErr := writeBuffer.WriteVirtual(ctx, "isFalse", m.GetIsFalse()); _isFalseErr != nil {
		return errors.Wrap(_isFalseErr, "Error serializing 'isFalse' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadBoolean"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadBoolean")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetTagPayloadBoolean) GetActualLength() uint32 {
	return m.ActualLength
}

//
////

func (m *_BACnetTagPayloadBoolean) isBACnetTagPayloadBoolean() bool {
	return true
}

func (m *_BACnetTagPayloadBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
