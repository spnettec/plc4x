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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataTimerAll is the data-structure of this message
type BACnetConstructedDataTimerAll struct {
	*BACnetConstructedData

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataTimerAll is the corresponding interface of BACnetConstructedDataTimerAll
type IBACnetConstructedDataTimerAll interface {
	IBACnetConstructedData
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

func (m *BACnetConstructedDataTimerAll) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TIMER
}

func (m *BACnetConstructedDataTimerAll) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataTimerAll) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataTimerAll) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

// NewBACnetConstructedDataTimerAll factory function for BACnetConstructedDataTimerAll
func NewBACnetConstructedDataTimerAll(openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataTimerAll {
	_result := &BACnetConstructedDataTimerAll{
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataTimerAll(structType interface{}) *BACnetConstructedDataTimerAll {
	if casted, ok := structType.(BACnetConstructedDataTimerAll); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimerAll); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataTimerAll(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataTimerAll(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataTimerAll) GetTypeName() string {
	return "BACnetConstructedDataTimerAll"
}

func (m *BACnetConstructedDataTimerAll) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataTimerAll) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *BACnetConstructedDataTimerAll) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTimerAllParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataTimerAll, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimerAll"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, utils.ParseValidationError{"TODO: implement me BACnetConstructedData ...TIMER..."}
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimerAll"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataTimerAll{
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataTimerAll) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimerAll"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimerAll"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataTimerAll) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
