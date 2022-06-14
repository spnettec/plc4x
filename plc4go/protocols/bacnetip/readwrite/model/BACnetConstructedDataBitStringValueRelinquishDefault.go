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

// BACnetConstructedDataBitStringValueRelinquishDefault is the data-structure of this message
type BACnetConstructedDataBitStringValueRelinquishDefault struct {
	*BACnetConstructedData
	RelinquishDefault *BACnetApplicationTagBitString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataBitStringValueRelinquishDefault is the corresponding interface of BACnetConstructedDataBitStringValueRelinquishDefault
type IBACnetConstructedDataBitStringValueRelinquishDefault interface {
	IBACnetConstructedData
	// GetRelinquishDefault returns RelinquishDefault (property field)
	GetRelinquishDefault() *BACnetApplicationTagBitString
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

func (m *BACnetConstructedDataBitStringValueRelinquishDefault) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BITSTRING_VALUE
}

func (m *BACnetConstructedDataBitStringValueRelinquishDefault) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RELINQUISH_DEFAULT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataBitStringValueRelinquishDefault) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataBitStringValueRelinquishDefault) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataBitStringValueRelinquishDefault) GetRelinquishDefault() *BACnetApplicationTagBitString {
	return m.RelinquishDefault
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBitStringValueRelinquishDefault factory function for BACnetConstructedDataBitStringValueRelinquishDefault
func NewBACnetConstructedDataBitStringValueRelinquishDefault(relinquishDefault *BACnetApplicationTagBitString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataBitStringValueRelinquishDefault {
	_result := &BACnetConstructedDataBitStringValueRelinquishDefault{
		RelinquishDefault:     relinquishDefault,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataBitStringValueRelinquishDefault(structType interface{}) *BACnetConstructedDataBitStringValueRelinquishDefault {
	if casted, ok := structType.(BACnetConstructedDataBitStringValueRelinquishDefault); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBitStringValueRelinquishDefault); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataBitStringValueRelinquishDefault(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataBitStringValueRelinquishDefault(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataBitStringValueRelinquishDefault) GetTypeName() string {
	return "BACnetConstructedDataBitStringValueRelinquishDefault"
}

func (m *BACnetConstructedDataBitStringValueRelinquishDefault) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataBitStringValueRelinquishDefault) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (relinquishDefault)
	lengthInBits += m.RelinquishDefault.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataBitStringValueRelinquishDefault) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBitStringValueRelinquishDefaultParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataBitStringValueRelinquishDefault, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBitStringValueRelinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBitStringValueRelinquishDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (relinquishDefault)
	if pullErr := readBuffer.PullContext("relinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for relinquishDefault")
	}
	_relinquishDefault, _relinquishDefaultErr := BACnetApplicationTagParse(readBuffer)
	if _relinquishDefaultErr != nil {
		return nil, errors.Wrap(_relinquishDefaultErr, "Error parsing 'relinquishDefault' field")
	}
	relinquishDefault := CastBACnetApplicationTagBitString(_relinquishDefault)
	if closeErr := readBuffer.CloseContext("relinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for relinquishDefault")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBitStringValueRelinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBitStringValueRelinquishDefault")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataBitStringValueRelinquishDefault{
		RelinquishDefault:     CastBACnetApplicationTagBitString(relinquishDefault),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataBitStringValueRelinquishDefault) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBitStringValueRelinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBitStringValueRelinquishDefault")
		}

		// Simple Field (relinquishDefault)
		if pushErr := writeBuffer.PushContext("relinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for relinquishDefault")
		}
		_relinquishDefaultErr := writeBuffer.WriteSerializable(m.RelinquishDefault)
		if popErr := writeBuffer.PopContext("relinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for relinquishDefault")
		}
		if _relinquishDefaultErr != nil {
			return errors.Wrap(_relinquishDefaultErr, "Error serializing 'relinquishDefault' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBitStringValueRelinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBitStringValueRelinquishDefault")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataBitStringValueRelinquishDefault) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
