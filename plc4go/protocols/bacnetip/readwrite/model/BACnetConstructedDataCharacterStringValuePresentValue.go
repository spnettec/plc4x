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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataCharacterStringValuePresentValue is the corresponding interface of BACnetConstructedDataCharacterStringValuePresentValue
type BACnetConstructedDataCharacterStringValuePresentValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPresentValue returns PresentValue (property field)
	GetPresentValue() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// IsBACnetConstructedDataCharacterStringValuePresentValue is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCharacterStringValuePresentValue()
}

// _BACnetConstructedDataCharacterStringValuePresentValue is the data-structure of this message
type _BACnetConstructedDataCharacterStringValuePresentValue struct {
	BACnetConstructedDataContract
	PresentValue BACnetApplicationTagCharacterString
}

var _ BACnetConstructedDataCharacterStringValuePresentValue = (*_BACnetConstructedDataCharacterStringValuePresentValue)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCharacterStringValuePresentValue)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCharacterStringValuePresentValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_CHARACTERSTRING_VALUE
}

func (m *_BACnetConstructedDataCharacterStringValuePresentValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESENT_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCharacterStringValuePresentValue) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCharacterStringValuePresentValue) GetPresentValue() BACnetApplicationTagCharacterString {
	return m.PresentValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCharacterStringValuePresentValue) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetPresentValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCharacterStringValuePresentValue factory function for _BACnetConstructedDataCharacterStringValuePresentValue
func NewBACnetConstructedDataCharacterStringValuePresentValue(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, presentValue BACnetApplicationTagCharacterString, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCharacterStringValuePresentValue {
	if presentValue == nil {
		panic("presentValue of type BACnetApplicationTagCharacterString for BACnetConstructedDataCharacterStringValuePresentValue must not be nil")
	}
	_result := &_BACnetConstructedDataCharacterStringValuePresentValue{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		PresentValue:                  presentValue,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCharacterStringValuePresentValue(structType any) BACnetConstructedDataCharacterStringValuePresentValue {
	if casted, ok := structType.(BACnetConstructedDataCharacterStringValuePresentValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCharacterStringValuePresentValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCharacterStringValuePresentValue) GetTypeName() string {
	return "BACnetConstructedDataCharacterStringValuePresentValue"
}

func (m *_BACnetConstructedDataCharacterStringValuePresentValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (presentValue)
	lengthInBits += m.PresentValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCharacterStringValuePresentValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCharacterStringValuePresentValue) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCharacterStringValuePresentValue BACnetConstructedDataCharacterStringValuePresentValue, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCharacterStringValuePresentValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCharacterStringValuePresentValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	presentValue, err := ReadSimpleField[BACnetApplicationTagCharacterString](ctx, "presentValue", ReadComplex[BACnetApplicationTagCharacterString](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagCharacterString](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'presentValue' field"))
	}
	m.PresentValue = presentValue

	actualValue, err := ReadVirtualField[BACnetApplicationTagCharacterString](ctx, "actualValue", (*BACnetApplicationTagCharacterString)(nil), presentValue)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCharacterStringValuePresentValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCharacterStringValuePresentValue")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCharacterStringValuePresentValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCharacterStringValuePresentValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCharacterStringValuePresentValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCharacterStringValuePresentValue")
		}

		if err := WriteSimpleField[BACnetApplicationTagCharacterString](ctx, "presentValue", m.GetPresentValue(), WriteComplex[BACnetApplicationTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'presentValue' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCharacterStringValuePresentValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCharacterStringValuePresentValue")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCharacterStringValuePresentValue) IsBACnetConstructedDataCharacterStringValuePresentValue() {
}

func (m *_BACnetConstructedDataCharacterStringValuePresentValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
