/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString is the data-structure of this message
type BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString struct {
	*BACnetNotificationParametersChangeOfDiscreteValueNewValue
	OctetStringValue *BACnetApplicationTagOctetString

	// Arguments.
	TagNumber uint8
}

// IBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString
type IBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString interface {
	IBACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetOctetStringValue returns OctetStringValue (property field)
	GetOctetStringValue() *BACnetApplicationTagOctetString
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString) InitializeParent(parent *BACnetNotificationParametersChangeOfDiscreteValueNewValue, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValue.OpeningTag = openingTag
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValue.PeekedTagHeader = peekedTagHeader
	m.BACnetNotificationParametersChangeOfDiscreteValueNewValue.ClosingTag = closingTag
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString) GetParent() *BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	return m.BACnetNotificationParametersChangeOfDiscreteValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString) GetOctetStringValue() *BACnetApplicationTagOctetString {
	return m.OctetStringValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString factory function for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString(octetStringValue *BACnetApplicationTagOctetString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString {
	_result := &BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString{
		OctetStringValue: octetStringValue,
		BACnetNotificationParametersChangeOfDiscreteValueNewValue: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString(structType interface{}) *BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString); ok {
		return casted
	}
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValue); ok {
		return CastBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString(casted.Child)
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValue); ok {
		return CastBACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString(casted.Child)
	}
	return nil
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString"
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (octetStringValue)
	lengthInBits += m.OctetStringValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetStringParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (octetStringValue)
	if pullErr := readBuffer.PullContext("octetStringValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for octetStringValue")
	}
	_octetStringValue, _octetStringValueErr := BACnetApplicationTagParse(readBuffer)
	if _octetStringValueErr != nil {
		return nil, errors.Wrap(_octetStringValueErr, "Error parsing 'octetStringValue' field")
	}
	octetStringValue := CastBACnetApplicationTagOctetString(_octetStringValue)
	if closeErr := readBuffer.CloseContext("octetStringValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for octetStringValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString")
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString{
		OctetStringValue: CastBACnetApplicationTagOctetString(octetStringValue),
		BACnetNotificationParametersChangeOfDiscreteValueNewValue: &BACnetNotificationParametersChangeOfDiscreteValueNewValue{},
	}
	_child.BACnetNotificationParametersChangeOfDiscreteValueNewValue.Child = _child
	return _child, nil
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString")
		}

		// Simple Field (octetStringValue)
		if pushErr := writeBuffer.PushContext("octetStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for octetStringValue")
		}
		_octetStringValueErr := m.OctetStringValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("octetStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for octetStringValue")
		}
		if _octetStringValueErr != nil {
			return errors.Wrap(_octetStringValueErr, "Error serializing 'octetStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersChangeOfDiscreteValueNewValueOctetString) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
