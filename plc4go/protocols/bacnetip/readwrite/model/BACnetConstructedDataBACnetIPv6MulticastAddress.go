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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetConstructedDataBACnetIPv6MulticastAddress is the corresponding interface of BACnetConstructedDataBACnetIPv6MulticastAddress
type BACnetConstructedDataBACnetIPv6MulticastAddress interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpv6MulticastAddress returns Ipv6MulticastAddress (property field)
	GetIpv6MulticastAddress() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
}

// BACnetConstructedDataBACnetIPv6MulticastAddressExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBACnetIPv6MulticastAddress.
// This is useful for switch cases.
type BACnetConstructedDataBACnetIPv6MulticastAddressExactly interface {
	BACnetConstructedDataBACnetIPv6MulticastAddress
	isBACnetConstructedDataBACnetIPv6MulticastAddress() bool
}

// _BACnetConstructedDataBACnetIPv6MulticastAddress is the data-structure of this message
type _BACnetConstructedDataBACnetIPv6MulticastAddress struct {
	*_BACnetConstructedData
        Ipv6MulticastAddress BACnetApplicationTagOctetString
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_BACNET_IPV6_MULTICAST_ADDRESS}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetIpv6MulticastAddress() BACnetApplicationTagOctetString {
	return m.Ipv6MulticastAddress
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetActualValue() BACnetApplicationTagOctetString {
	return CastBACnetApplicationTagOctetString(m.GetIpv6MulticastAddress())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataBACnetIPv6MulticastAddress factory function for _BACnetConstructedDataBACnetIPv6MulticastAddress
func NewBACnetConstructedDataBACnetIPv6MulticastAddress( ipv6MulticastAddress BACnetApplicationTagOctetString , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataBACnetIPv6MulticastAddress {
	_result := &_BACnetConstructedDataBACnetIPv6MulticastAddress{
		Ipv6MulticastAddress: ipv6MulticastAddress,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBACnetIPv6MulticastAddress(structType interface{}) BACnetConstructedDataBACnetIPv6MulticastAddress {
    if casted, ok := structType.(BACnetConstructedDataBACnetIPv6MulticastAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPv6MulticastAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPv6MulticastAddress"
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ipv6MulticastAddress)
	lengthInBits += m.Ipv6MulticastAddress.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBACnetIPv6MulticastAddressParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBACnetIPv6MulticastAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPv6MulticastAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPv6MulticastAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipv6MulticastAddress)
	if pullErr := readBuffer.PullContext("ipv6MulticastAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipv6MulticastAddress")
	}
_ipv6MulticastAddress, _ipv6MulticastAddressErr := BACnetApplicationTagParse(readBuffer)
	if _ipv6MulticastAddressErr != nil {
		return nil, errors.Wrap(_ipv6MulticastAddressErr, "Error parsing 'ipv6MulticastAddress' field of BACnetConstructedDataBACnetIPv6MulticastAddress")
	}
	ipv6MulticastAddress := _ipv6MulticastAddress.(BACnetApplicationTagOctetString)
	if closeErr := readBuffer.CloseContext("ipv6MulticastAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipv6MulticastAddress")
	}

	// Virtual field
	_actualValue := ipv6MulticastAddress
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPv6MulticastAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPv6MulticastAddress")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBACnetIPv6MulticastAddress{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Ipv6MulticastAddress: ipv6MulticastAddress,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPv6MulticastAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPv6MulticastAddress")
		}

	// Simple Field (ipv6MulticastAddress)
	if pushErr := writeBuffer.PushContext("ipv6MulticastAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ipv6MulticastAddress")
	}
	_ipv6MulticastAddressErr := writeBuffer.WriteSerializable(m.GetIpv6MulticastAddress())
	if popErr := writeBuffer.PopContext("ipv6MulticastAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ipv6MulticastAddress")
	}
	if _ipv6MulticastAddressErr != nil {
		return errors.Wrap(_ipv6MulticastAddressErr, "Error serializing 'ipv6MulticastAddress' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPv6MulticastAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPv6MulticastAddress")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) isBACnetConstructedDataBACnetIPv6MulticastAddress() bool {
	return true
}

func (m *_BACnetConstructedDataBACnetIPv6MulticastAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



