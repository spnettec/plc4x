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

// BACnetEventParameterChangeOfLifeSavety is the data-structure of this message
type BACnetEventParameterChangeOfLifeSavety struct {
	*BACnetEventParameter
	OpeningTag                  *BACnetOpeningTag
	TimeDelay                   *BACnetContextTagUnsignedInteger
	ListOfLifeSavetyAlarmValues *BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues
	ListOfAlarmValues           *BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues
	ModePropertyReference       *BACnetDeviceObjectPropertyReferenceEnclosed
	ClosingTag                  *BACnetClosingTag
}

// IBACnetEventParameterChangeOfLifeSavety is the corresponding interface of BACnetEventParameterChangeOfLifeSavety
type IBACnetEventParameterChangeOfLifeSavety interface {
	IBACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() *BACnetContextTagUnsignedInteger
	// GetListOfLifeSavetyAlarmValues returns ListOfLifeSavetyAlarmValues (property field)
	GetListOfLifeSavetyAlarmValues() *BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues
	// GetListOfAlarmValues returns ListOfAlarmValues (property field)
	GetListOfAlarmValues() *BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues
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

func (m *BACnetEventParameterChangeOfLifeSavety) InitializeParent(parent *BACnetEventParameter, peekedTagHeader *BACnetTagHeader) {
	m.BACnetEventParameter.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetEventParameterChangeOfLifeSavety) GetParent() *BACnetEventParameter {
	return m.BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetEventParameterChangeOfLifeSavety) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetEventParameterChangeOfLifeSavety) GetTimeDelay() *BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *BACnetEventParameterChangeOfLifeSavety) GetListOfLifeSavetyAlarmValues() *BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues {
	return m.ListOfLifeSavetyAlarmValues
}

func (m *BACnetEventParameterChangeOfLifeSavety) GetListOfAlarmValues() *BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues {
	return m.ListOfAlarmValues
}

func (m *BACnetEventParameterChangeOfLifeSavety) GetModePropertyReference() *BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.ModePropertyReference
}

func (m *BACnetEventParameterChangeOfLifeSavety) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterChangeOfLifeSavety factory function for BACnetEventParameterChangeOfLifeSavety
func NewBACnetEventParameterChangeOfLifeSavety(openingTag *BACnetOpeningTag, timeDelay *BACnetContextTagUnsignedInteger, listOfLifeSavetyAlarmValues *BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues, listOfAlarmValues *BACnetEventParameterChangeOfLifeSavetyListOfAlarmValues, modePropertyReference *BACnetDeviceObjectPropertyReferenceEnclosed, closingTag *BACnetClosingTag, peekedTagHeader *BACnetTagHeader) *BACnetEventParameterChangeOfLifeSavety {
	_result := &BACnetEventParameterChangeOfLifeSavety{
		OpeningTag:                  openingTag,
		TimeDelay:                   timeDelay,
		ListOfLifeSavetyAlarmValues: listOfLifeSavetyAlarmValues,
		ListOfAlarmValues:           listOfAlarmValues,
		ModePropertyReference:       modePropertyReference,
		ClosingTag:                  closingTag,
		BACnetEventParameter:        NewBACnetEventParameter(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetEventParameterChangeOfLifeSavety(structType interface{}) *BACnetEventParameterChangeOfLifeSavety {
	if casted, ok := structType.(BACnetEventParameterChangeOfLifeSavety); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfLifeSavety); ok {
		return casted
	}
	if casted, ok := structType.(BACnetEventParameter); ok {
		return CastBACnetEventParameterChangeOfLifeSavety(casted.Child)
	}
	if casted, ok := structType.(*BACnetEventParameter); ok {
		return CastBACnetEventParameterChangeOfLifeSavety(casted.Child)
	}
	return nil
}

func (m *BACnetEventParameterChangeOfLifeSavety) GetTypeName() string {
	return "BACnetEventParameterChangeOfLifeSavety"
}

func (m *BACnetEventParameterChangeOfLifeSavety) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventParameterChangeOfLifeSavety) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits()

	// Simple field (listOfLifeSavetyAlarmValues)
	lengthInBits += m.ListOfLifeSavetyAlarmValues.GetLengthInBits()

	// Simple field (listOfAlarmValues)
	lengthInBits += m.ListOfAlarmValues.GetLengthInBits()

	// Simple field (modePropertyReference)
	lengthInBits += m.ModePropertyReference.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetEventParameterChangeOfLifeSavety) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterChangeOfLifeSavetyParse(readBuffer utils.ReadBuffer) (*BACnetEventParameterChangeOfLifeSavety, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfLifeSavety"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(8)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (timeDelay)
	if pullErr := readBuffer.PullContext("timeDelay"); pullErr != nil {
		return nil, pullErr
	}
	_timeDelay, _timeDelayErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _timeDelayErr != nil {
		return nil, errors.Wrap(_timeDelayErr, "Error parsing 'timeDelay' field")
	}
	timeDelay := CastBACnetContextTagUnsignedInteger(_timeDelay)
	if closeErr := readBuffer.CloseContext("timeDelay"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (listOfLifeSavetyAlarmValues)
	if pullErr := readBuffer.PullContext("listOfLifeSavetyAlarmValues"); pullErr != nil {
		return nil, pullErr
	}
	_listOfLifeSavetyAlarmValues, _listOfLifeSavetyAlarmValuesErr := BACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValuesParse(readBuffer, uint8(uint8(1)))
	if _listOfLifeSavetyAlarmValuesErr != nil {
		return nil, errors.Wrap(_listOfLifeSavetyAlarmValuesErr, "Error parsing 'listOfLifeSavetyAlarmValues' field")
	}
	listOfLifeSavetyAlarmValues := CastBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues(_listOfLifeSavetyAlarmValues)
	if closeErr := readBuffer.CloseContext("listOfLifeSavetyAlarmValues"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (listOfAlarmValues)
	if pullErr := readBuffer.PullContext("listOfAlarmValues"); pullErr != nil {
		return nil, pullErr
	}
	_listOfAlarmValues, _listOfAlarmValuesErr := BACnetEventParameterChangeOfLifeSavetyListOfAlarmValuesParse(readBuffer, uint8(uint8(2)))
	if _listOfAlarmValuesErr != nil {
		return nil, errors.Wrap(_listOfAlarmValuesErr, "Error parsing 'listOfAlarmValues' field")
	}
	listOfAlarmValues := CastBACnetEventParameterChangeOfLifeSavetyListOfAlarmValues(_listOfAlarmValues)
	if closeErr := readBuffer.CloseContext("listOfAlarmValues"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (modePropertyReference)
	if pullErr := readBuffer.PullContext("modePropertyReference"); pullErr != nil {
		return nil, pullErr
	}
	_modePropertyReference, _modePropertyReferenceErr := BACnetDeviceObjectPropertyReferenceEnclosedParse(readBuffer, uint8(uint8(4)))
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
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(uint8(8)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfLifeSavety"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetEventParameterChangeOfLifeSavety{
		OpeningTag:                  CastBACnetOpeningTag(openingTag),
		TimeDelay:                   CastBACnetContextTagUnsignedInteger(timeDelay),
		ListOfLifeSavetyAlarmValues: CastBACnetEventParameterChangeOfLifeSavetyListOfLifeSavetyAlarmValues(listOfLifeSavetyAlarmValues),
		ListOfAlarmValues:           CastBACnetEventParameterChangeOfLifeSavetyListOfAlarmValues(listOfAlarmValues),
		ModePropertyReference:       CastBACnetDeviceObjectPropertyReferenceEnclosed(modePropertyReference),
		ClosingTag:                  CastBACnetClosingTag(closingTag),
		BACnetEventParameter:        &BACnetEventParameter{},
	}
	_child.BACnetEventParameter.Child = _child
	return _child, nil
}

func (m *BACnetEventParameterChangeOfLifeSavety) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfLifeSavety"); pushErr != nil {
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

		// Simple Field (timeDelay)
		if pushErr := writeBuffer.PushContext("timeDelay"); pushErr != nil {
			return pushErr
		}
		_timeDelayErr := m.TimeDelay.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timeDelay"); popErr != nil {
			return popErr
		}
		if _timeDelayErr != nil {
			return errors.Wrap(_timeDelayErr, "Error serializing 'timeDelay' field")
		}

		// Simple Field (listOfLifeSavetyAlarmValues)
		if pushErr := writeBuffer.PushContext("listOfLifeSavetyAlarmValues"); pushErr != nil {
			return pushErr
		}
		_listOfLifeSavetyAlarmValuesErr := m.ListOfLifeSavetyAlarmValues.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("listOfLifeSavetyAlarmValues"); popErr != nil {
			return popErr
		}
		if _listOfLifeSavetyAlarmValuesErr != nil {
			return errors.Wrap(_listOfLifeSavetyAlarmValuesErr, "Error serializing 'listOfLifeSavetyAlarmValues' field")
		}

		// Simple Field (listOfAlarmValues)
		if pushErr := writeBuffer.PushContext("listOfAlarmValues"); pushErr != nil {
			return pushErr
		}
		_listOfAlarmValuesErr := m.ListOfAlarmValues.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("listOfAlarmValues"); popErr != nil {
			return popErr
		}
		if _listOfAlarmValuesErr != nil {
			return errors.Wrap(_listOfAlarmValuesErr, "Error serializing 'listOfAlarmValues' field")
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

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfLifeSavety"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetEventParameterChangeOfLifeSavety) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
