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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetEventParameterChangeOfState is the corresponding interface of BACnetEventParameterChangeOfState
type BACnetEventParameterChangeOfState interface {
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetListOfValues returns ListOfValues (property field)
	GetListOfValues() BACnetEventParameterChangeOfStateListOfValues
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterChangeOfStateExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterChangeOfState.
// This is useful for switch cases.
type BACnetEventParameterChangeOfStateExactly interface {
	BACnetEventParameterChangeOfState
	isBACnetEventParameterChangeOfState() bool
}

// _BACnetEventParameterChangeOfState is the data-structure of this message
type _BACnetEventParameterChangeOfState struct {
	*_BACnetEventParameter
        OpeningTag BACnetOpeningTag
        TimeDelay BACnetContextTagUnsignedInteger
        ListOfValues BACnetEventParameterChangeOfStateListOfValues
        ClosingTag BACnetClosingTag
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterChangeOfState) InitializeParent(parent BACnetEventParameter , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetEventParameterChangeOfState)  GetParent() BACnetEventParameter {
	return m._BACnetEventParameter
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterChangeOfState) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterChangeOfState) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterChangeOfState) GetListOfValues() BACnetEventParameterChangeOfStateListOfValues {
	return m.ListOfValues
}

func (m *_BACnetEventParameterChangeOfState) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetEventParameterChangeOfState factory function for _BACnetEventParameterChangeOfState
func NewBACnetEventParameterChangeOfState( openingTag BACnetOpeningTag , timeDelay BACnetContextTagUnsignedInteger , listOfValues BACnetEventParameterChangeOfStateListOfValues , closingTag BACnetClosingTag , peekedTagHeader BACnetTagHeader ) *_BACnetEventParameterChangeOfState {
	_result := &_BACnetEventParameterChangeOfState{
		OpeningTag: openingTag,
		TimeDelay: timeDelay,
		ListOfValues: listOfValues,
		ClosingTag: closingTag,
    	_BACnetEventParameter: NewBACnetEventParameter(peekedTagHeader),
	}
	_result._BACnetEventParameter._BACnetEventParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterChangeOfState(structType interface{}) BACnetEventParameterChangeOfState {
    if casted, ok := structType.(BACnetEventParameterChangeOfState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterChangeOfState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterChangeOfState) GetTypeName() string {
	return "BACnetEventParameterChangeOfState"
}

func (m *_BACnetEventParameterChangeOfState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventParameterChangeOfState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits()

	// Simple field (listOfValues)
	lengthInBits += m.ListOfValues.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetEventParameterChangeOfState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterChangeOfStateParse(theBytes []byte) (BACnetEventParameterChangeOfState, error) {
	return BACnetEventParameterChangeOfStateParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventParameterChangeOfStateParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetEventParameterChangeOfState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterChangeOfState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterChangeOfState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(readBuffer , uint8( uint8(1) ) )
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventParameterChangeOfState")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (timeDelay)
	if pullErr := readBuffer.PullContext("timeDelay"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeDelay")
	}
_timeDelay, _timeDelayErr := BACnetContextTagParseWithBuffer(readBuffer , uint8( uint8(0) ) , BACnetDataType( BACnetDataType_UNSIGNED_INTEGER ) )
	if _timeDelayErr != nil {
		return nil, errors.Wrap(_timeDelayErr, "Error parsing 'timeDelay' field of BACnetEventParameterChangeOfState")
	}
	timeDelay := _timeDelay.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("timeDelay"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeDelay")
	}

	// Simple Field (listOfValues)
	if pullErr := readBuffer.PullContext("listOfValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfValues")
	}
_listOfValues, _listOfValuesErr := BACnetEventParameterChangeOfStateListOfValuesParseWithBuffer(readBuffer , uint8( uint8(1) ) )
	if _listOfValuesErr != nil {
		return nil, errors.Wrap(_listOfValuesErr, "Error parsing 'listOfValues' field of BACnetEventParameterChangeOfState")
	}
	listOfValues := _listOfValues.(BACnetEventParameterChangeOfStateListOfValues)
	if closeErr := readBuffer.CloseContext("listOfValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfValues")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(readBuffer , uint8( uint8(1) ) )
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventParameterChangeOfState")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterChangeOfState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterChangeOfState")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterChangeOfState{
		_BACnetEventParameter: &_BACnetEventParameter{
		},
		OpeningTag: openingTag,
		TimeDelay: timeDelay,
		ListOfValues: listOfValues,
		ClosingTag: closingTag,
	}
	_child._BACnetEventParameter._BACnetEventParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterChangeOfState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterChangeOfState) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterChangeOfState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterChangeOfState")
		}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (timeDelay)
	if pushErr := writeBuffer.PushContext("timeDelay"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timeDelay")
	}
	_timeDelayErr := writeBuffer.WriteSerializable(m.GetTimeDelay())
	if popErr := writeBuffer.PopContext("timeDelay"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timeDelay")
	}
	if _timeDelayErr != nil {
		return errors.Wrap(_timeDelayErr, "Error serializing 'timeDelay' field")
	}

	// Simple Field (listOfValues)
	if pushErr := writeBuffer.PushContext("listOfValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfValues")
	}
	_listOfValuesErr := writeBuffer.WriteSerializable(m.GetListOfValues())
	if popErr := writeBuffer.PopContext("listOfValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfValues")
	}
	if _listOfValuesErr != nil {
		return errors.Wrap(_listOfValuesErr, "Error serializing 'listOfValues' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

		if popErr := writeBuffer.PopContext("BACnetEventParameterChangeOfState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterChangeOfState")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetEventParameterChangeOfState) isBACnetEventParameterChangeOfState() bool {
	return true
}

func (m *_BACnetEventParameterChangeOfState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



