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

// BACnetConstructedDataSilenced is the data-structure of this message
type BACnetConstructedDataSilenced struct {
	*BACnetConstructedData
	Silenced *BACnetSilencedStateTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataSilenced is the corresponding interface of BACnetConstructedDataSilenced
type IBACnetConstructedDataSilenced interface {
	IBACnetConstructedData
	// GetSilenced returns Silenced (property field)
	GetSilenced() *BACnetSilencedStateTagged
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

func (m *BACnetConstructedDataSilenced) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataSilenced) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SILENCED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataSilenced) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataSilenced) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataSilenced) GetSilenced() *BACnetSilencedStateTagged {
	return m.Silenced
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSilenced factory function for BACnetConstructedDataSilenced
func NewBACnetConstructedDataSilenced(silenced *BACnetSilencedStateTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataSilenced {
	_result := &BACnetConstructedDataSilenced{
		Silenced:              silenced,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataSilenced(structType interface{}) *BACnetConstructedDataSilenced {
	if casted, ok := structType.(BACnetConstructedDataSilenced); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSilenced); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataSilenced(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataSilenced(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataSilenced) GetTypeName() string {
	return "BACnetConstructedDataSilenced"
}

func (m *BACnetConstructedDataSilenced) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataSilenced) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (silenced)
	lengthInBits += m.Silenced.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataSilenced) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSilencedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataSilenced, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSilenced"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (silenced)
	if pullErr := readBuffer.PullContext("silenced"); pullErr != nil {
		return nil, pullErr
	}
	_silenced, _silencedErr := BACnetSilencedStateTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _silencedErr != nil {
		return nil, errors.Wrap(_silencedErr, "Error parsing 'silenced' field")
	}
	silenced := CastBACnetSilencedStateTagged(_silenced)
	if closeErr := readBuffer.CloseContext("silenced"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSilenced"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataSilenced{
		Silenced:              CastBACnetSilencedStateTagged(silenced),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataSilenced) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSilenced"); pushErr != nil {
			return pushErr
		}

		// Simple Field (silenced)
		if pushErr := writeBuffer.PushContext("silenced"); pushErr != nil {
			return pushErr
		}
		_silencedErr := m.Silenced.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("silenced"); popErr != nil {
			return popErr
		}
		if _silencedErr != nil {
			return errors.Wrap(_silencedErr, "Error serializing 'silenced' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSilenced"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataSilenced) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
