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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataNetworkInterfaceName is the corresponding interface of BACnetConstructedDataNetworkInterfaceName
type BACnetConstructedDataNetworkInterfaceName interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNetworkInterfaceName returns NetworkInterfaceName (property field)
	GetNetworkInterfaceName() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataNetworkInterfaceNameExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataNetworkInterfaceName.
// This is useful for switch cases.
type BACnetConstructedDataNetworkInterfaceNameExactly interface {
	BACnetConstructedDataNetworkInterfaceName
	isBACnetConstructedDataNetworkInterfaceName() bool
}

// _BACnetConstructedDataNetworkInterfaceName is the data-structure of this message
type _BACnetConstructedDataNetworkInterfaceName struct {
	*_BACnetConstructedData
	NetworkInterfaceName BACnetApplicationTagCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNetworkInterfaceName) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataNetworkInterfaceName) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_NETWORK_INTERFACE_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNetworkInterfaceName) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataNetworkInterfaceName) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkInterfaceName) GetNetworkInterfaceName() BACnetApplicationTagCharacterString {
	return m.NetworkInterfaceName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkInterfaceName) GetActualValue() BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetNetworkInterfaceName())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataNetworkInterfaceName factory function for _BACnetConstructedDataNetworkInterfaceName
func NewBACnetConstructedDataNetworkInterfaceName(networkInterfaceName BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataNetworkInterfaceName {
	_result := &_BACnetConstructedDataNetworkInterfaceName{
		NetworkInterfaceName:   networkInterfaceName,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNetworkInterfaceName(structType interface{}) BACnetConstructedDataNetworkInterfaceName {
	if casted, ok := structType.(BACnetConstructedDataNetworkInterfaceName); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNetworkInterfaceName); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNetworkInterfaceName) GetTypeName() string {
	return "BACnetConstructedDataNetworkInterfaceName"
}

func (m *_BACnetConstructedDataNetworkInterfaceName) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataNetworkInterfaceName) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (networkInterfaceName)
	lengthInBits += m.NetworkInterfaceName.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataNetworkInterfaceName) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataNetworkInterfaceNameParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataNetworkInterfaceName, error) {
	return BACnetConstructedDataNetworkInterfaceNameParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataNetworkInterfaceNameParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataNetworkInterfaceName, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNetworkInterfaceName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNetworkInterfaceName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (networkInterfaceName)
	if pullErr := readBuffer.PullContext("networkInterfaceName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for networkInterfaceName")
	}
	_networkInterfaceName, _networkInterfaceNameErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _networkInterfaceNameErr != nil {
		return nil, errors.Wrap(_networkInterfaceNameErr, "Error parsing 'networkInterfaceName' field of BACnetConstructedDataNetworkInterfaceName")
	}
	networkInterfaceName := _networkInterfaceName.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("networkInterfaceName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for networkInterfaceName")
	}

	// Virtual field
	_actualValue := networkInterfaceName
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNetworkInterfaceName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNetworkInterfaceName")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataNetworkInterfaceName{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NetworkInterfaceName: networkInterfaceName,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataNetworkInterfaceName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataNetworkInterfaceName) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNetworkInterfaceName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNetworkInterfaceName")
		}

		// Simple Field (networkInterfaceName)
		if pushErr := writeBuffer.PushContext("networkInterfaceName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for networkInterfaceName")
		}
		_networkInterfaceNameErr := writeBuffer.WriteSerializable(m.GetNetworkInterfaceName())
		if popErr := writeBuffer.PopContext("networkInterfaceName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for networkInterfaceName")
		}
		if _networkInterfaceNameErr != nil {
			return errors.Wrap(_networkInterfaceNameErr, "Error serializing 'networkInterfaceName' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNetworkInterfaceName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNetworkInterfaceName")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataNetworkInterfaceName) isBACnetConstructedDataNetworkInterfaceName() bool {
	return true
}

func (m *_BACnetConstructedDataNetworkInterfaceName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
