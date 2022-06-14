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

// BACnetConstructedDataMACAddress is the data-structure of this message
type BACnetConstructedDataMACAddress struct {
	*BACnetConstructedData
	MacAddress *BACnetApplicationTagOctetString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataMACAddress is the corresponding interface of BACnetConstructedDataMACAddress
type IBACnetConstructedDataMACAddress interface {
	IBACnetConstructedData
	// GetMacAddress returns MacAddress (property field)
	GetMacAddress() *BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagOctetString
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

func (m *BACnetConstructedDataMACAddress) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataMACAddress) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAC_ADDRESS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMACAddress) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMACAddress) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMACAddress) GetMacAddress() *BACnetApplicationTagOctetString {
	return m.MacAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataMACAddress) GetActualValue() *BACnetApplicationTagOctetString {
	return CastBACnetApplicationTagOctetString(m.GetMacAddress())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMACAddress factory function for BACnetConstructedDataMACAddress
func NewBACnetConstructedDataMACAddress(macAddress *BACnetApplicationTagOctetString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataMACAddress {
	_result := &BACnetConstructedDataMACAddress{
		MacAddress:            macAddress,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMACAddress(structType interface{}) *BACnetConstructedDataMACAddress {
	if casted, ok := structType.(BACnetConstructedDataMACAddress); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMACAddress); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMACAddress(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMACAddress(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMACAddress) GetTypeName() string {
	return "BACnetConstructedDataMACAddress"
}

func (m *BACnetConstructedDataMACAddress) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMACAddress) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (macAddress)
	lengthInBits += m.MacAddress.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataMACAddress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMACAddressParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataMACAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMACAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMACAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (macAddress)
	if pullErr := readBuffer.PullContext("macAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for macAddress")
	}
	_macAddress, _macAddressErr := BACnetApplicationTagParse(readBuffer)
	if _macAddressErr != nil {
		return nil, errors.Wrap(_macAddressErr, "Error parsing 'macAddress' field")
	}
	macAddress := CastBACnetApplicationTagOctetString(_macAddress)
	if closeErr := readBuffer.CloseContext("macAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for macAddress")
	}

	// Virtual field
	_actualValue := macAddress
	actualValue := CastBACnetApplicationTagOctetString(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMACAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMACAddress")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMACAddress{
		MacAddress:            CastBACnetApplicationTagOctetString(macAddress),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMACAddress) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMACAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMACAddress")
		}

		// Simple Field (macAddress)
		if pushErr := writeBuffer.PushContext("macAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for macAddress")
		}
		_macAddressErr := writeBuffer.WriteSerializable(m.MacAddress)
		if popErr := writeBuffer.PopContext("macAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for macAddress")
		}
		if _macAddressErr != nil {
			return errors.Wrap(_macAddressErr, "Error serializing 'macAddress' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMACAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMACAddress")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMACAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
