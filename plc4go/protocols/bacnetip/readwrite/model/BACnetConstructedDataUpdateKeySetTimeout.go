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

// BACnetConstructedDataUpdateKeySetTimeout is the data-structure of this message
type BACnetConstructedDataUpdateKeySetTimeout struct {
	*BACnetConstructedData
	UpdateKeySetTimeout *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataUpdateKeySetTimeout is the corresponding interface of BACnetConstructedDataUpdateKeySetTimeout
type IBACnetConstructedDataUpdateKeySetTimeout interface {
	IBACnetConstructedData
	// GetUpdateKeySetTimeout returns UpdateKeySetTimeout (property field)
	GetUpdateKeySetTimeout() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataUpdateKeySetTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataUpdateKeySetTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_UPDATE_KEY_SET_TIMEOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataUpdateKeySetTimeout) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataUpdateKeySetTimeout) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataUpdateKeySetTimeout) GetUpdateKeySetTimeout() *BACnetApplicationTagUnsignedInteger {
	return m.UpdateKeySetTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataUpdateKeySetTimeout factory function for BACnetConstructedDataUpdateKeySetTimeout
func NewBACnetConstructedDataUpdateKeySetTimeout(updateKeySetTimeout *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataUpdateKeySetTimeout {
	_result := &BACnetConstructedDataUpdateKeySetTimeout{
		UpdateKeySetTimeout:   updateKeySetTimeout,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataUpdateKeySetTimeout(structType interface{}) *BACnetConstructedDataUpdateKeySetTimeout {
	if casted, ok := structType.(BACnetConstructedDataUpdateKeySetTimeout); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUpdateKeySetTimeout); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataUpdateKeySetTimeout(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataUpdateKeySetTimeout(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataUpdateKeySetTimeout) GetTypeName() string {
	return "BACnetConstructedDataUpdateKeySetTimeout"
}

func (m *BACnetConstructedDataUpdateKeySetTimeout) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataUpdateKeySetTimeout) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (updateKeySetTimeout)
	lengthInBits += m.UpdateKeySetTimeout.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataUpdateKeySetTimeout) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataUpdateKeySetTimeoutParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataUpdateKeySetTimeout, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUpdateKeySetTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUpdateKeySetTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (updateKeySetTimeout)
	if pullErr := readBuffer.PullContext("updateKeySetTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for updateKeySetTimeout")
	}
	_updateKeySetTimeout, _updateKeySetTimeoutErr := BACnetApplicationTagParse(readBuffer)
	if _updateKeySetTimeoutErr != nil {
		return nil, errors.Wrap(_updateKeySetTimeoutErr, "Error parsing 'updateKeySetTimeout' field")
	}
	updateKeySetTimeout := CastBACnetApplicationTagUnsignedInteger(_updateKeySetTimeout)
	if closeErr := readBuffer.CloseContext("updateKeySetTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for updateKeySetTimeout")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUpdateKeySetTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUpdateKeySetTimeout")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataUpdateKeySetTimeout{
		UpdateKeySetTimeout:   CastBACnetApplicationTagUnsignedInteger(updateKeySetTimeout),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataUpdateKeySetTimeout) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUpdateKeySetTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUpdateKeySetTimeout")
		}

		// Simple Field (updateKeySetTimeout)
		if pushErr := writeBuffer.PushContext("updateKeySetTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for updateKeySetTimeout")
		}
		_updateKeySetTimeoutErr := writeBuffer.WriteSerializable(m.UpdateKeySetTimeout)
		if popErr := writeBuffer.PopContext("updateKeySetTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for updateKeySetTimeout")
		}
		if _updateKeySetTimeoutErr != nil {
			return errors.Wrap(_updateKeySetTimeoutErr, "Error serializing 'updateKeySetTimeout' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUpdateKeySetTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUpdateKeySetTimeout")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataUpdateKeySetTimeout) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
