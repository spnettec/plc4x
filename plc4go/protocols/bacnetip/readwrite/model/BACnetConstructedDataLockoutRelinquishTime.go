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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataLockoutRelinquishTime is the corresponding interface of BACnetConstructedDataLockoutRelinquishTime
type BACnetConstructedDataLockoutRelinquishTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLockoutRelinquishTime returns LockoutRelinquishTime (property field)
	GetLockoutRelinquishTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataLockoutRelinquishTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLockoutRelinquishTime.
// This is useful for switch cases.
type BACnetConstructedDataLockoutRelinquishTimeExactly interface {
	BACnetConstructedDataLockoutRelinquishTime
	isBACnetConstructedDataLockoutRelinquishTime() bool
}

// _BACnetConstructedDataLockoutRelinquishTime is the data-structure of this message
type _BACnetConstructedDataLockoutRelinquishTime struct {
	*_BACnetConstructedData
	LockoutRelinquishTime BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOCKOUT_RELINQUISH_TIME
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLockoutRelinquishTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetLockoutRelinquishTime() BACnetApplicationTagUnsignedInteger {
	return m.LockoutRelinquishTime
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetLockoutRelinquishTime())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLockoutRelinquishTime factory function for _BACnetConstructedDataLockoutRelinquishTime
func NewBACnetConstructedDataLockoutRelinquishTime(lockoutRelinquishTime BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLockoutRelinquishTime {
	_result := &_BACnetConstructedDataLockoutRelinquishTime{
		LockoutRelinquishTime:  lockoutRelinquishTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLockoutRelinquishTime(structType any) BACnetConstructedDataLockoutRelinquishTime {
	if casted, ok := structType.(BACnetConstructedDataLockoutRelinquishTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLockoutRelinquishTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetTypeName() string {
	return "BACnetConstructedDataLockoutRelinquishTime"
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (lockoutRelinquishTime)
	lengthInBits += m.LockoutRelinquishTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLockoutRelinquishTimeParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLockoutRelinquishTime, error) {
	return BACnetConstructedDataLockoutRelinquishTimeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLockoutRelinquishTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLockoutRelinquishTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLockoutRelinquishTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLockoutRelinquishTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lockoutRelinquishTime)
	if pullErr := readBuffer.PullContext("lockoutRelinquishTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lockoutRelinquishTime")
	}
	_lockoutRelinquishTime, _lockoutRelinquishTimeErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _lockoutRelinquishTimeErr != nil {
		return nil, errors.Wrap(_lockoutRelinquishTimeErr, "Error parsing 'lockoutRelinquishTime' field of BACnetConstructedDataLockoutRelinquishTime")
	}
	lockoutRelinquishTime := _lockoutRelinquishTime.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("lockoutRelinquishTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lockoutRelinquishTime")
	}

	// Virtual field
	_actualValue := lockoutRelinquishTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLockoutRelinquishTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLockoutRelinquishTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLockoutRelinquishTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LockoutRelinquishTime: lockoutRelinquishTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLockoutRelinquishTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLockoutRelinquishTime")
		}

		// Simple Field (lockoutRelinquishTime)
		if pushErr := writeBuffer.PushContext("lockoutRelinquishTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lockoutRelinquishTime")
		}
		_lockoutRelinquishTimeErr := writeBuffer.WriteSerializable(ctx, m.GetLockoutRelinquishTime())
		if popErr := writeBuffer.PopContext("lockoutRelinquishTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lockoutRelinquishTime")
		}
		if _lockoutRelinquishTimeErr != nil {
			return errors.Wrap(_lockoutRelinquishTimeErr, "Error serializing 'lockoutRelinquishTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLockoutRelinquishTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLockoutRelinquishTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) isBACnetConstructedDataLockoutRelinquishTime() bool {
	return true
}

func (m *_BACnetConstructedDataLockoutRelinquishTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
