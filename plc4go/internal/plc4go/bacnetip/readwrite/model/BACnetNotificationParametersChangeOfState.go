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

// BACnetNotificationParametersChangeOfState is the data-structure of this message
type BACnetNotificationParametersChangeOfState struct {
	*BACnetNotificationParameters
	InnerOpeningTag *BACnetOpeningTag
	ChangeOfState   *BACnetPropertyStates
	StatusFlags     *BACnetStatusFlags
	InnerClosingTag *BACnetClosingTag

	// Arguments.
	TagNumber  uint8
	ObjectType BACnetObjectType
}

// IBACnetNotificationParametersChangeOfState is the corresponding interface of BACnetNotificationParametersChangeOfState
type IBACnetNotificationParametersChangeOfState interface {
	IBACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() *BACnetOpeningTag
	// GetChangeOfState returns ChangeOfState (property field)
	GetChangeOfState() *BACnetPropertyStates
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() *BACnetStatusFlags
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() *BACnetClosingTag
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

func (m *BACnetNotificationParametersChangeOfState) InitializeParent(parent *BACnetNotificationParameters, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetNotificationParameters.OpeningTag = openingTag
	m.BACnetNotificationParameters.PeekedTagHeader = peekedTagHeader
	m.BACnetNotificationParameters.ClosingTag = closingTag
}

func (m *BACnetNotificationParametersChangeOfState) GetParent() *BACnetNotificationParameters {
	return m.BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetNotificationParametersChangeOfState) GetInnerOpeningTag() *BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *BACnetNotificationParametersChangeOfState) GetChangeOfState() *BACnetPropertyStates {
	return m.ChangeOfState
}

func (m *BACnetNotificationParametersChangeOfState) GetStatusFlags() *BACnetStatusFlags {
	return m.StatusFlags
}

func (m *BACnetNotificationParametersChangeOfState) GetInnerClosingTag() *BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfState factory function for BACnetNotificationParametersChangeOfState
func NewBACnetNotificationParametersChangeOfState(innerOpeningTag *BACnetOpeningTag, changeOfState *BACnetPropertyStates, statusFlags *BACnetStatusFlags, innerClosingTag *BACnetClosingTag, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, objectType BACnetObjectType) *BACnetNotificationParametersChangeOfState {
	_result := &BACnetNotificationParametersChangeOfState{
		InnerOpeningTag:              innerOpeningTag,
		ChangeOfState:                changeOfState,
		StatusFlags:                  statusFlags,
		InnerClosingTag:              innerClosingTag,
		BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectType),
	}
	_result.Child = _result
	return _result
}

func CastBACnetNotificationParametersChangeOfState(structType interface{}) *BACnetNotificationParametersChangeOfState {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfState); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfState); ok {
		return casted
	}
	if casted, ok := structType.(BACnetNotificationParameters); ok {
		return CastBACnetNotificationParametersChangeOfState(casted.Child)
	}
	if casted, ok := structType.(*BACnetNotificationParameters); ok {
		return CastBACnetNotificationParametersChangeOfState(casted.Child)
	}
	return nil
}

func (m *BACnetNotificationParametersChangeOfState) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfState"
}

func (m *BACnetNotificationParametersChangeOfState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersChangeOfState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Simple field (changeOfState)
	lengthInBits += m.ChangeOfState.GetLengthInBits()

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits()

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersChangeOfState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersChangeOfStateParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, peekedTagNumber uint8) (*BACnetNotificationParametersChangeOfState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfState"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, pullErr
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetContextTagParse(readBuffer, uint8(peekedTagNumber), BACnetDataType(BACnetDataType_OPENING_TAG))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field")
	}
	innerOpeningTag := CastBACnetOpeningTag(_innerOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (changeOfState)
	if pullErr := readBuffer.PullContext("changeOfState"); pullErr != nil {
		return nil, pullErr
	}
	_changeOfState, _changeOfStateErr := BACnetPropertyStatesParse(readBuffer, uint8(uint8(0)))
	if _changeOfStateErr != nil {
		return nil, errors.Wrap(_changeOfStateErr, "Error parsing 'changeOfState' field")
	}
	changeOfState := CastBACnetPropertyStates(_changeOfState)
	if closeErr := readBuffer.CloseContext("changeOfState"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (statusFlags)
	if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
		return nil, pullErr
	}
	_statusFlags, _statusFlagsErr := BACnetStatusFlagsParse(readBuffer, uint8(uint8(1)))
	if _statusFlagsErr != nil {
		return nil, errors.Wrap(_statusFlagsErr, "Error parsing 'statusFlags' field")
	}
	statusFlags := CastBACnetStatusFlags(_statusFlags)
	if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, pullErr
	}
	_innerClosingTag, _innerClosingTagErr := BACnetContextTagParse(readBuffer, uint8(peekedTagNumber), BACnetDataType(BACnetDataType_CLOSING_TAG))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field")
	}
	innerClosingTag := CastBACnetClosingTag(_innerClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfState"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersChangeOfState{
		InnerOpeningTag:              CastBACnetOpeningTag(innerOpeningTag),
		ChangeOfState:                CastBACnetPropertyStates(changeOfState),
		StatusFlags:                  CastBACnetStatusFlags(statusFlags),
		InnerClosingTag:              CastBACnetClosingTag(innerClosingTag),
		BACnetNotificationParameters: &BACnetNotificationParameters{},
	}
	_child.BACnetNotificationParameters.Child = _child
	return _child, nil
}

func (m *BACnetNotificationParametersChangeOfState) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfState"); pushErr != nil {
			return pushErr
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return pushErr
		}
		_innerOpeningTagErr := m.InnerOpeningTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return popErr
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (changeOfState)
		if pushErr := writeBuffer.PushContext("changeOfState"); pushErr != nil {
			return pushErr
		}
		_changeOfStateErr := m.ChangeOfState.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("changeOfState"); popErr != nil {
			return popErr
		}
		if _changeOfStateErr != nil {
			return errors.Wrap(_changeOfStateErr, "Error serializing 'changeOfState' field")
		}

		// Simple Field (statusFlags)
		if pushErr := writeBuffer.PushContext("statusFlags"); pushErr != nil {
			return pushErr
		}
		_statusFlagsErr := m.StatusFlags.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("statusFlags"); popErr != nil {
			return popErr
		}
		if _statusFlagsErr != nil {
			return errors.Wrap(_statusFlagsErr, "Error serializing 'statusFlags' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return pushErr
		}
		_innerClosingTagErr := m.InnerClosingTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return popErr
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfState"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersChangeOfState) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
