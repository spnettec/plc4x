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


// BACnetPropertyStatesTimerState is the corresponding interface of BACnetPropertyStatesTimerState
type BACnetPropertyStatesTimerState interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetTimerState returns TimerState (property field)
	GetTimerState() BACnetTimerStateTagged
}

// BACnetPropertyStatesTimerStateExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesTimerState.
// This is useful for switch cases.
type BACnetPropertyStatesTimerStateExactly interface {
	BACnetPropertyStatesTimerState
	isBACnetPropertyStatesTimerState() bool
}

// _BACnetPropertyStatesTimerState is the data-structure of this message
type _BACnetPropertyStatesTimerState struct {
	*_BACnetPropertyStates
        TimerState BACnetTimerStateTagged
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesTimerState) InitializeParent(parent BACnetPropertyStates , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesTimerState)  GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesTimerState) GetTimerState() BACnetTimerStateTagged {
	return m.TimerState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetPropertyStatesTimerState factory function for _BACnetPropertyStatesTimerState
func NewBACnetPropertyStatesTimerState( timerState BACnetTimerStateTagged , peekedTagHeader BACnetTagHeader ) *_BACnetPropertyStatesTimerState {
	_result := &_BACnetPropertyStatesTimerState{
		TimerState: timerState,
    	_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesTimerState(structType interface{}) BACnetPropertyStatesTimerState {
    if casted, ok := structType.(BACnetPropertyStatesTimerState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesTimerState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesTimerState) GetTypeName() string {
	return "BACnetPropertyStatesTimerState"
}

func (m *_BACnetPropertyStatesTimerState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesTimerState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timerState)
	lengthInBits += m.TimerState.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetPropertyStatesTimerState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesTimerStateParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesTimerState, error) {
	return BACnetPropertyStatesTimerStateParseWithBuffer(utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesTimerStateParseWithBuffer(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesTimerState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesTimerState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesTimerState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timerState)
	if pullErr := readBuffer.PullContext("timerState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timerState")
	}
_timerState, _timerStateErr := BACnetTimerStateTaggedParseWithBuffer(readBuffer , uint8( peekedTagNumber ) , TagClass( TagClass_CONTEXT_SPECIFIC_TAGS ) )
	if _timerStateErr != nil {
		return nil, errors.Wrap(_timerStateErr, "Error parsing 'timerState' field of BACnetPropertyStatesTimerState")
	}
	timerState := _timerState.(BACnetTimerStateTagged)
	if closeErr := readBuffer.CloseContext("timerState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timerState")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesTimerState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesTimerState")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesTimerState{
		_BACnetPropertyStates: &_BACnetPropertyStates{
		},
		TimerState: timerState,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesTimerState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesTimerState) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesTimerState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesTimerState")
		}

	// Simple Field (timerState)
	if pushErr := writeBuffer.PushContext("timerState"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timerState")
	}
	_timerStateErr := writeBuffer.WriteSerializable(m.GetTimerState())
	if popErr := writeBuffer.PopContext("timerState"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timerState")
	}
	if _timerStateErr != nil {
		return errors.Wrap(_timerStateErr, "Error serializing 'timerState' field")
	}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesTimerState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesTimerState")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetPropertyStatesTimerState) isBACnetPropertyStatesTimerState() bool {
	return true
}

func (m *_BACnetPropertyStatesTimerState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



