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


// TriggerControlDataTriggerEvent is the corresponding interface of TriggerControlDataTriggerEvent
type TriggerControlDataTriggerEvent interface {
	utils.LengthAware
	utils.Serializable
	TriggerControlData
	// GetActionSelector returns ActionSelector (property field)
	GetActionSelector() byte
}

// TriggerControlDataTriggerEventExactly can be used when we want exactly this type and not a type which fulfills TriggerControlDataTriggerEvent.
// This is useful for switch cases.
type TriggerControlDataTriggerEventExactly interface {
	TriggerControlDataTriggerEvent
	isTriggerControlDataTriggerEvent() bool
}

// _TriggerControlDataTriggerEvent is the data-structure of this message
type _TriggerControlDataTriggerEvent struct {
	*_TriggerControlData
        ActionSelector byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TriggerControlDataTriggerEvent) InitializeParent(parent TriggerControlData , commandTypeContainer TriggerControlCommandTypeContainer , triggerGroup byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.TriggerGroup = triggerGroup
}

func (m *_TriggerControlDataTriggerEvent)  GetParent() TriggerControlData {
	return m._TriggerControlData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TriggerControlDataTriggerEvent) GetActionSelector() byte {
	return m.ActionSelector
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewTriggerControlDataTriggerEvent factory function for _TriggerControlDataTriggerEvent
func NewTriggerControlDataTriggerEvent( actionSelector byte , commandTypeContainer TriggerControlCommandTypeContainer , triggerGroup byte ) *_TriggerControlDataTriggerEvent {
	_result := &_TriggerControlDataTriggerEvent{
		ActionSelector: actionSelector,
    	_TriggerControlData: NewTriggerControlData(commandTypeContainer, triggerGroup),
	}
	_result._TriggerControlData._TriggerControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTriggerControlDataTriggerEvent(structType interface{}) TriggerControlDataTriggerEvent {
    if casted, ok := structType.(TriggerControlDataTriggerEvent); ok {
		return casted
	}
	if casted, ok := structType.(*TriggerControlDataTriggerEvent); ok {
		return *casted
	}
	return nil
}

func (m *_TriggerControlDataTriggerEvent) GetTypeName() string {
	return "TriggerControlDataTriggerEvent"
}

func (m *_TriggerControlDataTriggerEvent) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_TriggerControlDataTriggerEvent) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (actionSelector)
	lengthInBits += 8;

	return lengthInBits
}


func (m *_TriggerControlDataTriggerEvent) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TriggerControlDataTriggerEventParse(readBuffer utils.ReadBuffer) (TriggerControlDataTriggerEvent, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TriggerControlDataTriggerEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TriggerControlDataTriggerEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (actionSelector)
_actionSelector, _actionSelectorErr := readBuffer.ReadByte("actionSelector")
	if _actionSelectorErr != nil {
		return nil, errors.Wrap(_actionSelectorErr, "Error parsing 'actionSelector' field of TriggerControlDataTriggerEvent")
	}
	actionSelector := _actionSelector

	if closeErr := readBuffer.CloseContext("TriggerControlDataTriggerEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TriggerControlDataTriggerEvent")
	}

	// Create a partially initialized instance
	_child := &_TriggerControlDataTriggerEvent{
		_TriggerControlData: &_TriggerControlData{
		},
		ActionSelector: actionSelector,
	}
	_child._TriggerControlData._TriggerControlDataChildRequirements = _child
	return _child, nil
}

func (m *_TriggerControlDataTriggerEvent) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TriggerControlDataTriggerEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TriggerControlDataTriggerEvent")
		}

	// Simple Field (actionSelector)
	actionSelector := byte(m.GetActionSelector())
	_actionSelectorErr := writeBuffer.WriteByte("actionSelector", (actionSelector))
	if _actionSelectorErr != nil {
		return errors.Wrap(_actionSelectorErr, "Error serializing 'actionSelector' field")
	}

		if popErr := writeBuffer.PopContext("TriggerControlDataTriggerEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TriggerControlDataTriggerEvent")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_TriggerControlDataTriggerEvent) isTriggerControlDataTriggerEvent() bool {
	return true
}

func (m *_TriggerControlDataTriggerEvent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



