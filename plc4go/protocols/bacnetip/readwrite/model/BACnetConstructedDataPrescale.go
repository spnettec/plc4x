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

// BACnetConstructedDataPrescale is the data-structure of this message
type BACnetConstructedDataPrescale struct {
	*BACnetConstructedData
	Prescale *BACnetPrescale

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataPrescale is the corresponding interface of BACnetConstructedDataPrescale
type IBACnetConstructedDataPrescale interface {
	IBACnetConstructedData
	// GetPrescale returns Prescale (property field)
	GetPrescale() *BACnetPrescale
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

func (m *BACnetConstructedDataPrescale) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataPrescale) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRESCALE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataPrescale) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataPrescale) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataPrescale) GetPrescale() *BACnetPrescale {
	return m.Prescale
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPrescale factory function for BACnetConstructedDataPrescale
func NewBACnetConstructedDataPrescale(prescale *BACnetPrescale, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataPrescale {
	_result := &BACnetConstructedDataPrescale{
		Prescale:              prescale,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataPrescale(structType interface{}) *BACnetConstructedDataPrescale {
	if casted, ok := structType.(BACnetConstructedDataPrescale); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPrescale); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataPrescale(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataPrescale(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataPrescale) GetTypeName() string {
	return "BACnetConstructedDataPrescale"
}

func (m *BACnetConstructedDataPrescale) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataPrescale) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (prescale)
	lengthInBits += m.Prescale.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataPrescale) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataPrescaleParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataPrescale, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPrescale"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (prescale)
	if pullErr := readBuffer.PullContext("prescale"); pullErr != nil {
		return nil, pullErr
	}
	_prescale, _prescaleErr := BACnetPrescaleParse(readBuffer)
	if _prescaleErr != nil {
		return nil, errors.Wrap(_prescaleErr, "Error parsing 'prescale' field")
	}
	prescale := CastBACnetPrescale(_prescale)
	if closeErr := readBuffer.CloseContext("prescale"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPrescale"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataPrescale{
		Prescale:              CastBACnetPrescale(prescale),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataPrescale) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPrescale"); pushErr != nil {
			return pushErr
		}

		// Simple Field (prescale)
		if pushErr := writeBuffer.PushContext("prescale"); pushErr != nil {
			return pushErr
		}
		_prescaleErr := m.Prescale.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("prescale"); popErr != nil {
			return popErr
		}
		if _prescaleErr != nil {
			return errors.Wrap(_prescaleErr, "Error serializing 'prescale' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPrescale"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataPrescale) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
