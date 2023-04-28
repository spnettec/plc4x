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

// BACnetConstructedDataTimerMinPresValue is the corresponding interface of BACnetConstructedDataTimerMinPresValue
type BACnetConstructedDataTimerMinPresValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMinPresValue returns MinPresValue (property field)
	GetMinPresValue() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataTimerMinPresValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTimerMinPresValue.
// This is useful for switch cases.
type BACnetConstructedDataTimerMinPresValueExactly interface {
	BACnetConstructedDataTimerMinPresValue
	isBACnetConstructedDataTimerMinPresValue() bool
}

// _BACnetConstructedDataTimerMinPresValue is the data-structure of this message
type _BACnetConstructedDataTimerMinPresValue struct {
	*_BACnetConstructedData
	MinPresValue BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimerMinPresValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TIMER
}

func (m *_BACnetConstructedDataTimerMinPresValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MIN_PRES_VALUE
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimerMinPresValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTimerMinPresValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimerMinPresValue) GetMinPresValue() BACnetApplicationTagUnsignedInteger {
	return m.MinPresValue
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimerMinPresValue) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetMinPresValue())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimerMinPresValue factory function for _BACnetConstructedDataTimerMinPresValue
func NewBACnetConstructedDataTimerMinPresValue(minPresValue BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimerMinPresValue {
	_result := &_BACnetConstructedDataTimerMinPresValue{
		MinPresValue:           minPresValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimerMinPresValue(structType any) BACnetConstructedDataTimerMinPresValue {
	if casted, ok := structType.(BACnetConstructedDataTimerMinPresValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimerMinPresValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimerMinPresValue) GetTypeName() string {
	return "BACnetConstructedDataTimerMinPresValue"
}

func (m *_BACnetConstructedDataTimerMinPresValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (minPresValue)
	lengthInBits += m.MinPresValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimerMinPresValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataTimerMinPresValueParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimerMinPresValue, error) {
	return BACnetConstructedDataTimerMinPresValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataTimerMinPresValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimerMinPresValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimerMinPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimerMinPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (minPresValue)
	if pullErr := readBuffer.PullContext("minPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for minPresValue")
	}
	_minPresValue, _minPresValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _minPresValueErr != nil {
		return nil, errors.Wrap(_minPresValueErr, "Error parsing 'minPresValue' field of BACnetConstructedDataTimerMinPresValue")
	}
	minPresValue := _minPresValue.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("minPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for minPresValue")
	}

	// Virtual field
	_actualValue := minPresValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimerMinPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimerMinPresValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTimerMinPresValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MinPresValue: minPresValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTimerMinPresValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimerMinPresValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimerMinPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimerMinPresValue")
		}

		// Simple Field (minPresValue)
		if pushErr := writeBuffer.PushContext("minPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for minPresValue")
		}
		_minPresValueErr := writeBuffer.WriteSerializable(ctx, m.GetMinPresValue())
		if popErr := writeBuffer.PopContext("minPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for minPresValue")
		}
		if _minPresValueErr != nil {
			return errors.Wrap(_minPresValueErr, "Error serializing 'minPresValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimerMinPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimerMinPresValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimerMinPresValue) isBACnetConstructedDataTimerMinPresValue() bool {
	return true
}

func (m *_BACnetConstructedDataTimerMinPresValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
