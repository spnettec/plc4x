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

// BACnetConstructedDataResolution is the data-structure of this message
type BACnetConstructedDataResolution struct {
	*BACnetConstructedData
	Resolution *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataResolution is the corresponding interface of BACnetConstructedDataResolution
type IBACnetConstructedDataResolution interface {
	IBACnetConstructedData
	// GetResolution returns Resolution (property field)
	GetResolution() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataResolution) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataResolution) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RESOLUTION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataResolution) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataResolution) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataResolution) GetResolution() *BACnetApplicationTagReal {
	return m.Resolution
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataResolution factory function for BACnetConstructedDataResolution
func NewBACnetConstructedDataResolution(resolution *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataResolution {
	_result := &BACnetConstructedDataResolution{
		Resolution:            resolution,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataResolution(structType interface{}) *BACnetConstructedDataResolution {
	if casted, ok := structType.(BACnetConstructedDataResolution); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataResolution); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataResolution(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataResolution(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataResolution) GetTypeName() string {
	return "BACnetConstructedDataResolution"
}

func (m *BACnetConstructedDataResolution) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataResolution) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (resolution)
	lengthInBits += m.Resolution.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataResolution) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataResolutionParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataResolution, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataResolution"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (resolution)
	if pullErr := readBuffer.PullContext("resolution"); pullErr != nil {
		return nil, pullErr
	}
	_resolution, _resolutionErr := BACnetApplicationTagParse(readBuffer)
	if _resolutionErr != nil {
		return nil, errors.Wrap(_resolutionErr, "Error parsing 'resolution' field")
	}
	resolution := CastBACnetApplicationTagReal(_resolution)
	if closeErr := readBuffer.CloseContext("resolution"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataResolution"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataResolution{
		Resolution:            CastBACnetApplicationTagReal(resolution),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataResolution) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataResolution"); pushErr != nil {
			return pushErr
		}

		// Simple Field (resolution)
		if pushErr := writeBuffer.PushContext("resolution"); pushErr != nil {
			return pushErr
		}
		_resolutionErr := m.Resolution.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("resolution"); popErr != nil {
			return popErr
		}
		if _resolutionErr != nil {
			return errors.Wrap(_resolutionErr, "Error serializing 'resolution' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataResolution"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataResolution) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
