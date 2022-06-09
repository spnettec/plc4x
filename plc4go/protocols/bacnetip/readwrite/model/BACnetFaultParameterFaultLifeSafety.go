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

// BACnetFaultParameterFaultLifeSafety is the data-structure of this message
type BACnetFaultParameterFaultLifeSafety struct {
	*BACnetFaultParameter
	OpeningTag            *BACnetOpeningTag
	ListOfFaultValues     *BACnetFaultParameterFaultLifeSafetyListOfFaultValues
	ModePropertyReference *BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag            *BACnetClosingTag
}

// IBACnetFaultParameterFaultLifeSafety is the corresponding interface of BACnetFaultParameterFaultLifeSafety
type IBACnetFaultParameterFaultLifeSafety interface {
	IBACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetListOfFaultValues returns ListOfFaultValues (property field)
	GetListOfFaultValues() *BACnetFaultParameterFaultLifeSafetyListOfFaultValues
	// GetModePropertyReference returns ModePropertyReference (property field)
	GetModePropertyReference() *BACnetDeviceObjectPropertyReferenceEnclosed
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

func (m *BACnetFaultParameterFaultLifeSafety) InitializeParent(parent *BACnetFaultParameter, peekedTagHeader *BACnetTagHeader) {
	m.BACnetFaultParameter.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetFaultParameterFaultLifeSafety) GetParent() *BACnetFaultParameter {
	return m.BACnetFaultParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetFaultParameterFaultLifeSafety) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetFaultParameterFaultLifeSafety) GetListOfFaultValues() *BACnetFaultParameterFaultLifeSafetyListOfFaultValues {
	return m.ListOfFaultValues
}

func (m *BACnetFaultParameterFaultLifeSafety) GetModePropertyReference() *BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.ModePropertyReference
}

func (m *BACnetFaultParameterFaultLifeSafety) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultLifeSafety factory function for BACnetFaultParameterFaultLifeSafety
func NewBACnetFaultParameterFaultLifeSafety(openingTag *BACnetOpeningTag, listOfFaultValues *BACnetFaultParameterFaultLifeSafetyListOfFaultValues, modePropertyReference *BACnetDeviceObjectPropertyReferenceEnclosed, closingTag *BACnetClosingTag, peekedTagHeader *BACnetTagHeader) *BACnetFaultParameterFaultLifeSafety {
	_result := &BACnetFaultParameterFaultLifeSafety{
		OpeningTag:            openingTag,
		ListOfFaultValues:     listOfFaultValues,
		ModePropertyReference: modePropertyReference,
		ClosingTag:            closingTag,
		BACnetFaultParameter:  NewBACnetFaultParameter(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetFaultParameterFaultLifeSafety(structType interface{}) *BACnetFaultParameterFaultLifeSafety {
	if casted, ok := structType.(BACnetFaultParameterFaultLifeSafety); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultLifeSafety); ok {
		return casted
	}
	if casted, ok := structType.(BACnetFaultParameter); ok {
		return CastBACnetFaultParameterFaultLifeSafety(casted.Child)
	}
	if casted, ok := structType.(*BACnetFaultParameter); ok {
		return CastBACnetFaultParameterFaultLifeSafety(casted.Child)
	}
	return nil
}

func (m *BACnetFaultParameterFaultLifeSafety) GetTypeName() string {
	return "BACnetFaultParameterFaultLifeSafety"
}

func (m *BACnetFaultParameterFaultLifeSafety) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetFaultParameterFaultLifeSafety) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (listOfFaultValues)
	lengthInBits += m.ListOfFaultValues.GetLengthInBits()

	// Simple field (modePropertyReference)
	lengthInBits += m.ModePropertyReference.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetFaultParameterFaultLifeSafety) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultLifeSafetyParse(readBuffer utils.ReadBuffer) (*BACnetFaultParameterFaultLifeSafety, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultLifeSafety"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(3)))
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
	_listOfFaultValues, _listOfFaultValuesErr := BACnetFaultParameterFaultLifeSafetyListOfFaultValuesParse(readBuffer, uint8(uint8(0)))
	if _listOfFaultValuesErr != nil {
		return nil, errors.Wrap(_listOfFaultValuesErr, "Error parsing 'listOfFaultValues' field")
	}
	listOfFaultValues := CastBACnetFaultParameterFaultLifeSafetyListOfFaultValues(_listOfFaultValues)
	if closeErr := readBuffer.CloseContext("listOfFaultValues"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (modePropertyReference)
	if pullErr := readBuffer.PullContext("modePropertyReference"); pullErr != nil {
		return nil, pullErr
	}
	_modePropertyReference, _modePropertyReferenceErr := BACnetDeviceObjectPropertyReferenceEnclosedParse(readBuffer, uint8(uint8(1)))
	if _modePropertyReferenceErr != nil {
		return nil, errors.Wrap(_modePropertyReferenceErr, "Error parsing 'modePropertyReference' field")
	}
	modePropertyReference := CastBACnetDeviceObjectPropertyReferenceEnclosed(_modePropertyReference)
	if closeErr := readBuffer.CloseContext("modePropertyReference"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(uint8(3)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultLifeSafety"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetFaultParameterFaultLifeSafety{
		OpeningTag:            CastBACnetOpeningTag(openingTag),
		ListOfFaultValues:     CastBACnetFaultParameterFaultLifeSafetyListOfFaultValues(listOfFaultValues),
		ModePropertyReference: CastBACnetDeviceObjectPropertyReferenceEnclosed(modePropertyReference),
		ClosingTag:            CastBACnetClosingTag(closingTag),
		BACnetFaultParameter:  &BACnetFaultParameter{},
	}
	_child.BACnetFaultParameter.Child = _child
	return _child, nil
}

func (m *BACnetFaultParameterFaultLifeSafety) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultLifeSafety"); pushErr != nil {
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

		// Simple Field (modePropertyReference)
		if pushErr := writeBuffer.PushContext("modePropertyReference"); pushErr != nil {
			return pushErr
		}
		_modePropertyReferenceErr := m.ModePropertyReference.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("modePropertyReference"); popErr != nil {
			return popErr
		}
		if _modePropertyReferenceErr != nil {
			return errors.Wrap(_modePropertyReferenceErr, "Error serializing 'modePropertyReference' field")
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

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultLifeSafety"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetFaultParameterFaultLifeSafety) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
