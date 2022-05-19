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

// BACnetEventSummary is the data-structure of this message
type BACnetEventSummary struct {
	ObjectIdentifier        *BACnetContextTagObjectIdentifier
	EventState              *BACnetEventStateTagged
	AcknowledgedTransitions *BACnetEventTransitionBits
	EventTimestamps         *BACnetEventTimestamps
	NotifyType              *BACnetNotifyTypeTagged
	EventEnable             *BACnetEventTransitionBits
	EventPriorities         *BACnetEventProrities
}

// IBACnetEventSummary is the corresponding interface of BACnetEventSummary
type IBACnetEventSummary interface {
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() *BACnetContextTagObjectIdentifier
	// GetEventState returns EventState (property field)
	GetEventState() *BACnetEventStateTagged
	// GetAcknowledgedTransitions returns AcknowledgedTransitions (property field)
	GetAcknowledgedTransitions() *BACnetEventTransitionBits
	// GetEventTimestamps returns EventTimestamps (property field)
	GetEventTimestamps() *BACnetEventTimestamps
	// GetNotifyType returns NotifyType (property field)
	GetNotifyType() *BACnetNotifyTypeTagged
	// GetEventEnable returns EventEnable (property field)
	GetEventEnable() *BACnetEventTransitionBits
	// GetEventPriorities returns EventPriorities (property field)
	GetEventPriorities() *BACnetEventProrities
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetEventSummary) GetObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *BACnetEventSummary) GetEventState() *BACnetEventStateTagged {
	return m.EventState
}

func (m *BACnetEventSummary) GetAcknowledgedTransitions() *BACnetEventTransitionBits {
	return m.AcknowledgedTransitions
}

func (m *BACnetEventSummary) GetEventTimestamps() *BACnetEventTimestamps {
	return m.EventTimestamps
}

func (m *BACnetEventSummary) GetNotifyType() *BACnetNotifyTypeTagged {
	return m.NotifyType
}

func (m *BACnetEventSummary) GetEventEnable() *BACnetEventTransitionBits {
	return m.EventEnable
}

func (m *BACnetEventSummary) GetEventPriorities() *BACnetEventProrities {
	return m.EventPriorities
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventSummary factory function for BACnetEventSummary
func NewBACnetEventSummary(objectIdentifier *BACnetContextTagObjectIdentifier, eventState *BACnetEventStateTagged, acknowledgedTransitions *BACnetEventTransitionBits, eventTimestamps *BACnetEventTimestamps, notifyType *BACnetNotifyTypeTagged, eventEnable *BACnetEventTransitionBits, eventPriorities *BACnetEventProrities) *BACnetEventSummary {
	return &BACnetEventSummary{ObjectIdentifier: objectIdentifier, EventState: eventState, AcknowledgedTransitions: acknowledgedTransitions, EventTimestamps: eventTimestamps, NotifyType: notifyType, EventEnable: eventEnable, EventPriorities: eventPriorities}
}

func CastBACnetEventSummary(structType interface{}) *BACnetEventSummary {
	if casted, ok := structType.(BACnetEventSummary); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventSummary); ok {
		return casted
	}
	return nil
}

func (m *BACnetEventSummary) GetTypeName() string {
	return "BACnetEventSummary"
}

func (m *BACnetEventSummary) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventSummary) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits()

	// Simple field (acknowledgedTransitions)
	lengthInBits += m.AcknowledgedTransitions.GetLengthInBits()

	// Simple field (eventTimestamps)
	lengthInBits += m.EventTimestamps.GetLengthInBits()

	// Simple field (notifyType)
	lengthInBits += m.NotifyType.GetLengthInBits()

	// Simple field (eventEnable)
	lengthInBits += m.EventEnable.GetLengthInBits()

	// Simple field (eventPriorities)
	lengthInBits += m.EventPriorities.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetEventSummary) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventSummaryParse(readBuffer utils.ReadBuffer) (*BACnetEventSummary, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventSummary"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field")
	}
	objectIdentifier := CastBACnetContextTagObjectIdentifier(_objectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (eventState)
	if pullErr := readBuffer.PullContext("eventState"); pullErr != nil {
		return nil, pullErr
	}
	_eventState, _eventStateErr := BACnetEventStateTaggedParse(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _eventStateErr != nil {
		return nil, errors.Wrap(_eventStateErr, "Error parsing 'eventState' field")
	}
	eventState := CastBACnetEventStateTagged(_eventState)
	if closeErr := readBuffer.CloseContext("eventState"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (acknowledgedTransitions)
	if pullErr := readBuffer.PullContext("acknowledgedTransitions"); pullErr != nil {
		return nil, pullErr
	}
	_acknowledgedTransitions, _acknowledgedTransitionsErr := BACnetEventTransitionBitsParse(readBuffer, uint8(uint8(2)))
	if _acknowledgedTransitionsErr != nil {
		return nil, errors.Wrap(_acknowledgedTransitionsErr, "Error parsing 'acknowledgedTransitions' field")
	}
	acknowledgedTransitions := CastBACnetEventTransitionBits(_acknowledgedTransitions)
	if closeErr := readBuffer.CloseContext("acknowledgedTransitions"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (eventTimestamps)
	if pullErr := readBuffer.PullContext("eventTimestamps"); pullErr != nil {
		return nil, pullErr
	}
	_eventTimestamps, _eventTimestampsErr := BACnetEventTimestampsParse(readBuffer, uint8(uint8(3)))
	if _eventTimestampsErr != nil {
		return nil, errors.Wrap(_eventTimestampsErr, "Error parsing 'eventTimestamps' field")
	}
	eventTimestamps := CastBACnetEventTimestamps(_eventTimestamps)
	if closeErr := readBuffer.CloseContext("eventTimestamps"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (notifyType)
	if pullErr := readBuffer.PullContext("notifyType"); pullErr != nil {
		return nil, pullErr
	}
	_notifyType, _notifyTypeErr := BACnetNotifyTypeTaggedParse(readBuffer, uint8(uint8(4)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _notifyTypeErr != nil {
		return nil, errors.Wrap(_notifyTypeErr, "Error parsing 'notifyType' field")
	}
	notifyType := CastBACnetNotifyTypeTagged(_notifyType)
	if closeErr := readBuffer.CloseContext("notifyType"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (eventEnable)
	if pullErr := readBuffer.PullContext("eventEnable"); pullErr != nil {
		return nil, pullErr
	}
	_eventEnable, _eventEnableErr := BACnetEventTransitionBitsParse(readBuffer, uint8(uint8(5)))
	if _eventEnableErr != nil {
		return nil, errors.Wrap(_eventEnableErr, "Error parsing 'eventEnable' field")
	}
	eventEnable := CastBACnetEventTransitionBits(_eventEnable)
	if closeErr := readBuffer.CloseContext("eventEnable"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (eventPriorities)
	if pullErr := readBuffer.PullContext("eventPriorities"); pullErr != nil {
		return nil, pullErr
	}
	_eventPriorities, _eventPrioritiesErr := BACnetEventProritiesParse(readBuffer, uint8(uint8(6)))
	if _eventPrioritiesErr != nil {
		return nil, errors.Wrap(_eventPrioritiesErr, "Error parsing 'eventPriorities' field")
	}
	eventPriorities := CastBACnetEventProrities(_eventPriorities)
	if closeErr := readBuffer.CloseContext("eventPriorities"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetEventSummary"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetEventSummary(objectIdentifier, eventState, acknowledgedTransitions, eventTimestamps, notifyType, eventEnable, eventPriorities), nil
}

func (m *BACnetEventSummary) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetEventSummary"); pushErr != nil {
		return pushErr
	}

	// Simple Field (objectIdentifier)
	if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
		return pushErr
	}
	_objectIdentifierErr := m.ObjectIdentifier.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
		return popErr
	}
	if _objectIdentifierErr != nil {
		return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
	}

	// Simple Field (eventState)
	if pushErr := writeBuffer.PushContext("eventState"); pushErr != nil {
		return pushErr
	}
	_eventStateErr := m.EventState.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("eventState"); popErr != nil {
		return popErr
	}
	if _eventStateErr != nil {
		return errors.Wrap(_eventStateErr, "Error serializing 'eventState' field")
	}

	// Simple Field (acknowledgedTransitions)
	if pushErr := writeBuffer.PushContext("acknowledgedTransitions"); pushErr != nil {
		return pushErr
	}
	_acknowledgedTransitionsErr := m.AcknowledgedTransitions.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("acknowledgedTransitions"); popErr != nil {
		return popErr
	}
	if _acknowledgedTransitionsErr != nil {
		return errors.Wrap(_acknowledgedTransitionsErr, "Error serializing 'acknowledgedTransitions' field")
	}

	// Simple Field (eventTimestamps)
	if pushErr := writeBuffer.PushContext("eventTimestamps"); pushErr != nil {
		return pushErr
	}
	_eventTimestampsErr := m.EventTimestamps.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("eventTimestamps"); popErr != nil {
		return popErr
	}
	if _eventTimestampsErr != nil {
		return errors.Wrap(_eventTimestampsErr, "Error serializing 'eventTimestamps' field")
	}

	// Simple Field (notifyType)
	if pushErr := writeBuffer.PushContext("notifyType"); pushErr != nil {
		return pushErr
	}
	_notifyTypeErr := m.NotifyType.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("notifyType"); popErr != nil {
		return popErr
	}
	if _notifyTypeErr != nil {
		return errors.Wrap(_notifyTypeErr, "Error serializing 'notifyType' field")
	}

	// Simple Field (eventEnable)
	if pushErr := writeBuffer.PushContext("eventEnable"); pushErr != nil {
		return pushErr
	}
	_eventEnableErr := m.EventEnable.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("eventEnable"); popErr != nil {
		return popErr
	}
	if _eventEnableErr != nil {
		return errors.Wrap(_eventEnableErr, "Error serializing 'eventEnable' field")
	}

	// Simple Field (eventPriorities)
	if pushErr := writeBuffer.PushContext("eventPriorities"); pushErr != nil {
		return pushErr
	}
	_eventPrioritiesErr := m.EventPriorities.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("eventPriorities"); popErr != nil {
		return popErr
	}
	if _eventPrioritiesErr != nil {
		return errors.Wrap(_eventPrioritiesErr, "Error serializing 'eventPriorities' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventSummary"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetEventSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
