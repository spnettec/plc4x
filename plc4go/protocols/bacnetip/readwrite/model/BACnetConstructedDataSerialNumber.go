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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataSerialNumber is the corresponding interface of BACnetConstructedDataSerialNumber
type BACnetConstructedDataSerialNumber interface {
	BACnetConstructedData
	// GetSerialNumber returns SerialNumber (property field)
	GetSerialNumber() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetConstructedDataSerialNumber is the data-structure of this message
type _BACnetConstructedDataSerialNumber struct {
	*_BACnetConstructedData
	SerialNumber BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSerialNumber) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSerialNumber) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SERIAL_NUMBER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSerialNumber) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSerialNumber) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSerialNumber) GetSerialNumber() BACnetApplicationTagCharacterString {
	return m.SerialNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSerialNumber) GetActualValue() BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetSerialNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSerialNumber factory function for _BACnetConstructedDataSerialNumber
func NewBACnetConstructedDataSerialNumber(serialNumber BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSerialNumber {
	_result := &_BACnetConstructedDataSerialNumber{
		SerialNumber:           serialNumber,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSerialNumber(structType interface{}) BACnetConstructedDataSerialNumber {
	if casted, ok := structType.(BACnetConstructedDataSerialNumber); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSerialNumber); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSerialNumber) GetTypeName() string {
	return "BACnetConstructedDataSerialNumber"
}

func (m *_BACnetConstructedDataSerialNumber) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataSerialNumber) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (serialNumber)
	lengthInBits += m.SerialNumber.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSerialNumber) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSerialNumberParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSerialNumber, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSerialNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSerialNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (serialNumber)
	if pullErr := readBuffer.PullContext("serialNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for serialNumber")
	}
	_serialNumber, _serialNumberErr := BACnetApplicationTagParse(readBuffer)
	if _serialNumberErr != nil {
		return nil, errors.Wrap(_serialNumberErr, "Error parsing 'serialNumber' field")
	}
	serialNumber := _serialNumber.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("serialNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for serialNumber")
	}

	// Virtual field
	_actualValue := serialNumber
	actualValue := _actualValue.(BACnetApplicationTagCharacterString)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSerialNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSerialNumber")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSerialNumber{
		SerialNumber:           serialNumber,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSerialNumber) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSerialNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSerialNumber")
		}

		// Simple Field (serialNumber)
		if pushErr := writeBuffer.PushContext("serialNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for serialNumber")
		}
		_serialNumberErr := writeBuffer.WriteSerializable(m.GetSerialNumber())
		if popErr := writeBuffer.PopContext("serialNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for serialNumber")
		}
		if _serialNumberErr != nil {
			return errors.Wrap(_serialNumberErr, "Error serializing 'serialNumber' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSerialNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSerialNumber")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSerialNumber) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
