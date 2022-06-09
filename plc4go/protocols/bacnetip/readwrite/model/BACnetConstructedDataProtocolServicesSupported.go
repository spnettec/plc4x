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

// BACnetConstructedDataProtocolServicesSupported is the data-structure of this message
type BACnetConstructedDataProtocolServicesSupported struct {
	*BACnetConstructedData
	ProtocolServicesSupported *BACnetServicesSupportedTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataProtocolServicesSupported is the corresponding interface of BACnetConstructedDataProtocolServicesSupported
type IBACnetConstructedDataProtocolServicesSupported interface {
	IBACnetConstructedData
	// GetProtocolServicesSupported returns ProtocolServicesSupported (property field)
	GetProtocolServicesSupported() *BACnetServicesSupportedTagged
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

func (m *BACnetConstructedDataProtocolServicesSupported) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataProtocolServicesSupported) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROTOCOL_SERVICES_SUPPORTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataProtocolServicesSupported) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataProtocolServicesSupported) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataProtocolServicesSupported) GetProtocolServicesSupported() *BACnetServicesSupportedTagged {
	return m.ProtocolServicesSupported
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProtocolServicesSupported factory function for BACnetConstructedDataProtocolServicesSupported
func NewBACnetConstructedDataProtocolServicesSupported(protocolServicesSupported *BACnetServicesSupportedTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataProtocolServicesSupported {
	_result := &BACnetConstructedDataProtocolServicesSupported{
		ProtocolServicesSupported: protocolServicesSupported,
		BACnetConstructedData:     NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataProtocolServicesSupported(structType interface{}) *BACnetConstructedDataProtocolServicesSupported {
	if casted, ok := structType.(BACnetConstructedDataProtocolServicesSupported); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProtocolServicesSupported); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataProtocolServicesSupported(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataProtocolServicesSupported(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataProtocolServicesSupported) GetTypeName() string {
	return "BACnetConstructedDataProtocolServicesSupported"
}

func (m *BACnetConstructedDataProtocolServicesSupported) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataProtocolServicesSupported) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (protocolServicesSupported)
	lengthInBits += m.ProtocolServicesSupported.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataProtocolServicesSupported) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataProtocolServicesSupportedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataProtocolServicesSupported, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProtocolServicesSupported"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (protocolServicesSupported)
	if pullErr := readBuffer.PullContext("protocolServicesSupported"); pullErr != nil {
		return nil, pullErr
	}
	_protocolServicesSupported, _protocolServicesSupportedErr := BACnetServicesSupportedTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _protocolServicesSupportedErr != nil {
		return nil, errors.Wrap(_protocolServicesSupportedErr, "Error parsing 'protocolServicesSupported' field")
	}
	protocolServicesSupported := CastBACnetServicesSupportedTagged(_protocolServicesSupported)
	if closeErr := readBuffer.CloseContext("protocolServicesSupported"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProtocolServicesSupported"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataProtocolServicesSupported{
		ProtocolServicesSupported: CastBACnetServicesSupportedTagged(protocolServicesSupported),
		BACnetConstructedData:     &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataProtocolServicesSupported) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProtocolServicesSupported"); pushErr != nil {
			return pushErr
		}

		// Simple Field (protocolServicesSupported)
		if pushErr := writeBuffer.PushContext("protocolServicesSupported"); pushErr != nil {
			return pushErr
		}
		_protocolServicesSupportedErr := m.ProtocolServicesSupported.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("protocolServicesSupported"); popErr != nil {
			return popErr
		}
		if _protocolServicesSupportedErr != nil {
			return errors.Wrap(_protocolServicesSupportedErr, "Error serializing 'protocolServicesSupported' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProtocolServicesSupported"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataProtocolServicesSupported) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
