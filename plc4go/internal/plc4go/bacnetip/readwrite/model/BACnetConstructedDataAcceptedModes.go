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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataAcceptedModes is the data-structure of this message
type BACnetConstructedDataAcceptedModes struct {
	*BACnetConstructedData
	AcceptedModes []*BACnetConstructedDataAcceptedModesEntry

	// Arguments.
	TagNumber                  uint8
	PropertyIdentifierArgument BACnetContextTagPropertyIdentifier
}

// IBACnetConstructedDataAcceptedModes is the corresponding interface of BACnetConstructedDataAcceptedModes
type IBACnetConstructedDataAcceptedModes interface {
	IBACnetConstructedData
	// GetAcceptedModes returns AcceptedModes (property field)
	GetAcceptedModes() []*BACnetConstructedDataAcceptedModesEntry
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

func (m *BACnetConstructedDataAcceptedModes) GetObjectType() BACnetObjectType {
	return 0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAcceptedModes) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAcceptedModes) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAcceptedModes) GetAcceptedModes() []*BACnetConstructedDataAcceptedModesEntry {
	return m.AcceptedModes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAcceptedModes factory function for BACnetConstructedDataAcceptedModes
func NewBACnetConstructedDataAcceptedModes(acceptedModes []*BACnetConstructedDataAcceptedModesEntry, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8, propertyIdentifierArgument BACnetContextTagPropertyIdentifier) *BACnetConstructedDataAcceptedModes {
	_result := &BACnetConstructedDataAcceptedModes{
		AcceptedModes:         acceptedModes,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber, propertyIdentifierArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAcceptedModes(structType interface{}) *BACnetConstructedDataAcceptedModes {
	if casted, ok := structType.(BACnetConstructedDataAcceptedModes); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAcceptedModes); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAcceptedModes(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAcceptedModes(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAcceptedModes) GetTypeName() string {
	return "BACnetConstructedDataAcceptedModes"
}

func (m *BACnetConstructedDataAcceptedModes) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAcceptedModes) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.AcceptedModes) > 0 {
		for _, element := range m.AcceptedModes {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataAcceptedModes) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAcceptedModesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, propertyIdentifierArgument *BACnetContextTagPropertyIdentifier) (*BACnetConstructedDataAcceptedModes, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAcceptedModes"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (acceptedModes)
	if pullErr := readBuffer.PullContext("acceptedModes", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	acceptedModes := make([]*BACnetConstructedDataAcceptedModesEntry, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetConstructedDataAcceptedModesEntryParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'acceptedModes' field")
			}
			acceptedModes = append(acceptedModes, CastBACnetConstructedDataAcceptedModesEntry(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("acceptedModes", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAcceptedModes"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAcceptedModes{
		AcceptedModes:         acceptedModes,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAcceptedModes) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAcceptedModes"); pushErr != nil {
			return pushErr
		}

		// Array Field (acceptedModes)
		if m.AcceptedModes != nil {
			if pushErr := writeBuffer.PushContext("acceptedModes", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.AcceptedModes {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'acceptedModes' field")
				}
			}
			if popErr := writeBuffer.PopContext("acceptedModes", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAcceptedModes"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAcceptedModes) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
