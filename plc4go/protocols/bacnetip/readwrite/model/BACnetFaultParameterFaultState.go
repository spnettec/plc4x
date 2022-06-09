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

// BACnetFaultParameterFaultState is the data-structure of this message
type BACnetFaultParameterFaultState struct {
	*BACnetFaultParameter
	OpeningTag        *BACnetOpeningTag
	ListOfFaultValues *BACnetFaultParameterFaultStateListOfFaultValues
	ClosingTag        *BACnetClosingTag
}

// IBACnetFaultParameterFaultState is the corresponding interface of BACnetFaultParameterFaultState
type IBACnetFaultParameterFaultState interface {
	IBACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetListOfFaultValues returns ListOfFaultValues (property field)
	GetListOfFaultValues() *BACnetFaultParameterFaultStateListOfFaultValues
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetFaultParameterFaultState) InitializeParent(parent *BACnetFaultParameter, peekedTagHeader *BACnetTagHeader) {
	m.BACnetFaultParameter.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetFaultParameterFaultState) GetParent() *BACnetFaultParameter {
	return m.BACnetFaultParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetFaultParameterFaultState) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetFaultParameterFaultState) GetListOfFaultValues() *BACnetFaultParameterFaultStateListOfFaultValues {
	return m.ListOfFaultValues
}

func (m *BACnetFaultParameterFaultState) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultState factory function for BACnetFaultParameterFaultState
func NewBACnetFaultParameterFaultState(openingTag *BACnetOpeningTag, listOfFaultValues *BACnetFaultParameterFaultStateListOfFaultValues, closingTag *BACnetClosingTag, peekedTagHeader *BACnetTagHeader) *BACnetFaultParameterFaultState {
	_result := &BACnetFaultParameterFaultState{
		OpeningTag:           openingTag,
		ListOfFaultValues:    listOfFaultValues,
		ClosingTag:           closingTag,
		BACnetFaultParameter: NewBACnetFaultParameter(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetFaultParameterFaultState(structType interface{}) *BACnetFaultParameterFaultState {
	if casted, ok := structType.(BACnetFaultParameterFaultState); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultState); ok {
		return casted
	}
	if casted, ok := structType.(BACnetFaultParameter); ok {
		return CastBACnetFaultParameterFaultState(casted.Child)
	}
	if casted, ok := structType.(*BACnetFaultParameter); ok {
		return CastBACnetFaultParameterFaultState(casted.Child)
	}
	return nil
}

func (m *BACnetFaultParameterFaultState) GetTypeName() string {
	return "BACnetFaultParameterFaultState"
}

func (m *BACnetFaultParameterFaultState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetFaultParameterFaultState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (listOfFaultValues)
	lengthInBits += m.ListOfFaultValues.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetFaultParameterFaultState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultStateParse(readBuffer utils.ReadBuffer) (*BACnetFaultParameterFaultState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultState"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(4)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (listOfFaultValues)
	if pullErr := readBuffer.PullContext("listOfFaultValues"); pullErr != nil {
		return nil, pullErr
	}
	_listOfFaultValues, _listOfFaultValuesErr := BACnetFaultParameterFaultStateListOfFaultValuesParse(readBuffer, uint8(uint8(0)))
	if _listOfFaultValuesErr != nil {
		return nil, errors.Wrap(_listOfFaultValuesErr, "Error parsing 'listOfFaultValues' field")
	}
	listOfFaultValues := CastBACnetFaultParameterFaultStateListOfFaultValues(_listOfFaultValues)
	if closeErr := readBuffer.CloseContext("listOfFaultValues"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(uint8(4)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultState"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetFaultParameterFaultState{
		OpeningTag:           CastBACnetOpeningTag(openingTag),
		ListOfFaultValues:    CastBACnetFaultParameterFaultStateListOfFaultValues(listOfFaultValues),
		ClosingTag:           CastBACnetClosingTag(closingTag),
		BACnetFaultParameter: &BACnetFaultParameter{},
	}
	_child.BACnetFaultParameter.Child = _child
	return _child, nil
}

func (m *BACnetFaultParameterFaultState) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultState"); pushErr != nil {
			return pushErr
		}

		// Simple Field (openingTag)
		if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
			return pushErr
		}
		_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
			return popErr
		}
		if _openingTagErr != nil {
			return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
		}

		// Simple Field (listOfFaultValues)
		if pushErr := writeBuffer.PushContext("listOfFaultValues"); pushErr != nil {
			return pushErr
		}
		_listOfFaultValuesErr := m.ListOfFaultValues.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("listOfFaultValues"); popErr != nil {
			return popErr
		}
		if _listOfFaultValuesErr != nil {
			return errors.Wrap(_listOfFaultValuesErr, "Error serializing 'listOfFaultValues' field")
		}

		// Simple Field (closingTag)
		if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
			return pushErr
		}
		_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
			return popErr
		}
		if _closingTagErr != nil {
			return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultState"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetFaultParameterFaultState) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
