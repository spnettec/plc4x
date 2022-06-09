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

// BACnetConstructedDataIPDefaultGateway is the data-structure of this message
type BACnetConstructedDataIPDefaultGateway struct {
	*BACnetConstructedData
	IpDefaultGateway *BACnetApplicationTagOctetString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataIPDefaultGateway is the corresponding interface of BACnetConstructedDataIPDefaultGateway
type IBACnetConstructedDataIPDefaultGateway interface {
	IBACnetConstructedData
	// GetIpDefaultGateway returns IpDefaultGateway (property field)
	GetIpDefaultGateway() *BACnetApplicationTagOctetString
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

func (m *BACnetConstructedDataIPDefaultGateway) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataIPDefaultGateway) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IP_DEFAULT_GATEWAY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataIPDefaultGateway) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataIPDefaultGateway) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataIPDefaultGateway) GetIpDefaultGateway() *BACnetApplicationTagOctetString {
	return m.IpDefaultGateway
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPDefaultGateway factory function for BACnetConstructedDataIPDefaultGateway
func NewBACnetConstructedDataIPDefaultGateway(ipDefaultGateway *BACnetApplicationTagOctetString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataIPDefaultGateway {
	_result := &BACnetConstructedDataIPDefaultGateway{
		IpDefaultGateway:      ipDefaultGateway,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataIPDefaultGateway(structType interface{}) *BACnetConstructedDataIPDefaultGateway {
	if casted, ok := structType.(BACnetConstructedDataIPDefaultGateway); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPDefaultGateway); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataIPDefaultGateway(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataIPDefaultGateway(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataIPDefaultGateway) GetTypeName() string {
	return "BACnetConstructedDataIPDefaultGateway"
}

func (m *BACnetConstructedDataIPDefaultGateway) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataIPDefaultGateway) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ipDefaultGateway)
	lengthInBits += m.IpDefaultGateway.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataIPDefaultGateway) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIPDefaultGatewayParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataIPDefaultGateway, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPDefaultGateway"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipDefaultGateway)
	if pullErr := readBuffer.PullContext("ipDefaultGateway"); pullErr != nil {
		return nil, pullErr
	}
	_ipDefaultGateway, _ipDefaultGatewayErr := BACnetApplicationTagParse(readBuffer)
	if _ipDefaultGatewayErr != nil {
		return nil, errors.Wrap(_ipDefaultGatewayErr, "Error parsing 'ipDefaultGateway' field")
	}
	ipDefaultGateway := CastBACnetApplicationTagOctetString(_ipDefaultGateway)
	if closeErr := readBuffer.CloseContext("ipDefaultGateway"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPDefaultGateway"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataIPDefaultGateway{
		IpDefaultGateway:      CastBACnetApplicationTagOctetString(ipDefaultGateway),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataIPDefaultGateway) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPDefaultGateway"); pushErr != nil {
			return pushErr
		}

		// Simple Field (ipDefaultGateway)
		if pushErr := writeBuffer.PushContext("ipDefaultGateway"); pushErr != nil {
			return pushErr
		}
		_ipDefaultGatewayErr := m.IpDefaultGateway.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("ipDefaultGateway"); popErr != nil {
			return popErr
		}
		if _ipDefaultGatewayErr != nil {
			return errors.Wrap(_ipDefaultGatewayErr, "Error serializing 'ipDefaultGateway' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPDefaultGateway"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataIPDefaultGateway) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
