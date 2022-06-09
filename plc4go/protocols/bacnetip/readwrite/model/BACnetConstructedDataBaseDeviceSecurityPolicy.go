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

// BACnetConstructedDataBaseDeviceSecurityPolicy is the data-structure of this message
type BACnetConstructedDataBaseDeviceSecurityPolicy struct {
	*BACnetConstructedData
	BaseDeviceSecurityPolicy *BACnetSecurityLevelTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataBaseDeviceSecurityPolicy is the corresponding interface of BACnetConstructedDataBaseDeviceSecurityPolicy
type IBACnetConstructedDataBaseDeviceSecurityPolicy interface {
	IBACnetConstructedData
	// GetBaseDeviceSecurityPolicy returns BaseDeviceSecurityPolicy (property field)
	GetBaseDeviceSecurityPolicy() *BACnetSecurityLevelTagged
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

func (m *BACnetConstructedDataBaseDeviceSecurityPolicy) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataBaseDeviceSecurityPolicy) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BASE_DEVICE_SECURITY_POLICY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataBaseDeviceSecurityPolicy) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataBaseDeviceSecurityPolicy) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataBaseDeviceSecurityPolicy) GetBaseDeviceSecurityPolicy() *BACnetSecurityLevelTagged {
	return m.BaseDeviceSecurityPolicy
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBaseDeviceSecurityPolicy factory function for BACnetConstructedDataBaseDeviceSecurityPolicy
func NewBACnetConstructedDataBaseDeviceSecurityPolicy(baseDeviceSecurityPolicy *BACnetSecurityLevelTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataBaseDeviceSecurityPolicy {
	_result := &BACnetConstructedDataBaseDeviceSecurityPolicy{
		BaseDeviceSecurityPolicy: baseDeviceSecurityPolicy,
		BACnetConstructedData:    NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataBaseDeviceSecurityPolicy(structType interface{}) *BACnetConstructedDataBaseDeviceSecurityPolicy {
	if casted, ok := structType.(BACnetConstructedDataBaseDeviceSecurityPolicy); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBaseDeviceSecurityPolicy); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataBaseDeviceSecurityPolicy(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataBaseDeviceSecurityPolicy(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataBaseDeviceSecurityPolicy) GetTypeName() string {
	return "BACnetConstructedDataBaseDeviceSecurityPolicy"
}

func (m *BACnetConstructedDataBaseDeviceSecurityPolicy) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataBaseDeviceSecurityPolicy) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (baseDeviceSecurityPolicy)
	lengthInBits += m.BaseDeviceSecurityPolicy.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataBaseDeviceSecurityPolicy) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBaseDeviceSecurityPolicyParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataBaseDeviceSecurityPolicy, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBaseDeviceSecurityPolicy"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (baseDeviceSecurityPolicy)
	if pullErr := readBuffer.PullContext("baseDeviceSecurityPolicy"); pullErr != nil {
		return nil, pullErr
	}
	_baseDeviceSecurityPolicy, _baseDeviceSecurityPolicyErr := BACnetSecurityLevelTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _baseDeviceSecurityPolicyErr != nil {
		return nil, errors.Wrap(_baseDeviceSecurityPolicyErr, "Error parsing 'baseDeviceSecurityPolicy' field")
	}
	baseDeviceSecurityPolicy := CastBACnetSecurityLevelTagged(_baseDeviceSecurityPolicy)
	if closeErr := readBuffer.CloseContext("baseDeviceSecurityPolicy"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBaseDeviceSecurityPolicy"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataBaseDeviceSecurityPolicy{
		BaseDeviceSecurityPolicy: CastBACnetSecurityLevelTagged(baseDeviceSecurityPolicy),
		BACnetConstructedData:    &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataBaseDeviceSecurityPolicy) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBaseDeviceSecurityPolicy"); pushErr != nil {
			return pushErr
		}

		// Simple Field (baseDeviceSecurityPolicy)
		if pushErr := writeBuffer.PushContext("baseDeviceSecurityPolicy"); pushErr != nil {
			return pushErr
		}
		_baseDeviceSecurityPolicyErr := m.BaseDeviceSecurityPolicy.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("baseDeviceSecurityPolicy"); popErr != nil {
			return popErr
		}
		if _baseDeviceSecurityPolicyErr != nil {
			return errors.Wrap(_baseDeviceSecurityPolicyErr, "Error serializing 'baseDeviceSecurityPolicy' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBaseDeviceSecurityPolicy"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataBaseDeviceSecurityPolicy) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
