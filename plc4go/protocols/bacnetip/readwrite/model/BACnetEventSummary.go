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


// BACnetEventSummary is the corresponding interface of BACnetEventSummary
type BACnetEventSummary interface {
	utils.LengthAware
	utils.Serializable
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetEventState returns EventState (property field)
	GetEventState() BACnetEventStateTagged
	// GetAcknowledgedTransitions returns AcknowledgedTransitions (property field)
	GetAcknowledgedTransitions() BACnetEventTransitionBitsTagged
	// GetEventTimestamps returns EventTimestamps (property field)
	GetEventTimestamps() BACnetEventTimestampsEnclosed
	// GetNotifyType returns NotifyType (property field)
	GetNotifyType() BACnetNotifyTypeTagged
	// GetEventEnable returns EventEnable (property field)
	GetEventEnable() BACnetEventTransitionBitsTagged
	// GetEventPriorities returns EventPriorities (property field)
	GetEventPriorities() BACnetEventPriorities
}

// BACnetEventSummaryExactly can be used when we want exactly this type and not a type which fulfills BACnetEventSummary.
// This is useful for switch cases.
type BACnetEventSummaryExactly interface {
	BACnetEventSummary
	isBACnetEventSummary() bool
}

// _BACnetEventSummary is the data-structure of this message
type _BACnetEventSummary struct {
        ObjectIdentifier BACnetContextTagObjectIdentifier
        EventState BACnetEventStateTagged
        AcknowledgedTransitions BACnetEventTransitionBitsTagged
        EventTimestamps BACnetEventTimestampsEnclosed
        NotifyType BACnetNotifyTypeTagged
        EventEnable BACnetEventTransitionBitsTagged
        EventPriorities BACnetEventPriorities
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventSummary) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetEventSummary) GetEventState() BACnetEventStateTagged {
	return m.EventState
}

func (m *_BACnetEventSummary) GetAcknowledgedTransitions() BACnetEventTransitionBitsTagged {
	return m.AcknowledgedTransitions
}

func (m *_BACnetEventSummary) GetEventTimestamps() BACnetEventTimestampsEnclosed {
	return m.EventTimestamps
}

func (m *_BACnetEventSummary) GetNotifyType() BACnetNotifyTypeTagged {
	return m.NotifyType
}

func (m *_BACnetEventSummary) GetEventEnable() BACnetEventTransitionBitsTagged {
	return m.EventEnable
}

func (m *_BACnetEventSummary) GetEventPriorities() BACnetEventPriorities {
	return m.EventPriorities
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetEventSummary factory function for _BACnetEventSummary
func NewBACnetEventSummary( objectIdentifier BACnetContextTagObjectIdentifier , eventState BACnetEventStateTagged , acknowledgedTransitions BACnetEventTransitionBitsTagged , eventTimestamps BACnetEventTimestampsEnclosed , notifyType BACnetNotifyTypeTagged , eventEnable BACnetEventTransitionBitsTagged , eventPriorities BACnetEventPriorities ) *_BACnetEventSummary {
return &_BACnetEventSummary{ ObjectIdentifier: objectIdentifier , EventState: eventState , AcknowledgedTransitions: acknowledgedTransitions , EventTimestamps: eventTimestamps , NotifyType: notifyType , EventEnable: eventEnable , EventPriorities: eventPriorities }
}

// Deprecated: use the interface for direct cast
func CastBACnetEventSummary(structType interface{}) BACnetEventSummary {
    if casted, ok := structType.(BACnetEventSummary); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventSummary); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventSummary) GetTypeName() string {
	return "BACnetEventSummary"
}

func (m *_BACnetEventSummary) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventSummary) GetLengthInBitsConditional(lastItem bool) uint16 {
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


func (m *_BACnetEventSummary) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventSummaryParse(readBuffer utils.ReadBuffer) (BACnetEventSummary, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
_objectIdentifier, _objectIdentifierErr := BACnetContextTagParse(readBuffer , uint8( uint8(0) ) , BACnetDataType( BACnetDataType_BACNET_OBJECT_IDENTIFIER ) )
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field of BACnetEventSummary")
	}
	objectIdentifier := _objectIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	// Simple Field (eventState)
	if pullErr := readBuffer.PullContext("eventState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventState")
	}
_eventState, _eventStateErr := BACnetEventStateTaggedParse(readBuffer , uint8( uint8(1) ) , TagClass( TagClass_CONTEXT_SPECIFIC_TAGS ) )
	if _eventStateErr != nil {
		return nil, errors.Wrap(_eventStateErr, "Error parsing 'eventState' field of BACnetEventSummary")
	}
	eventState := _eventState.(BACnetEventStateTagged)
	if closeErr := readBuffer.CloseContext("eventState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventState")
	}

	// Simple Field (acknowledgedTransitions)
	if pullErr := readBuffer.PullContext("acknowledgedTransitions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for acknowledgedTransitions")
	}
_acknowledgedTransitions, _acknowledgedTransitionsErr := BACnetEventTransitionBitsTaggedParse(readBuffer , uint8( uint8(2) ) , TagClass( TagClass_CONTEXT_SPECIFIC_TAGS ) )
	if _acknowledgedTransitionsErr != nil {
		return nil, errors.Wrap(_acknowledgedTransitionsErr, "Error parsing 'acknowledgedTransitions' field of BACnetEventSummary")
	}
	acknowledgedTransitions := _acknowledgedTransitions.(BACnetEventTransitionBitsTagged)
	if closeErr := readBuffer.CloseContext("acknowledgedTransitions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for acknowledgedTransitions")
	}

	// Simple Field (eventTimestamps)
	if pullErr := readBuffer.PullContext("eventTimestamps"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventTimestamps")
	}
_eventTimestamps, _eventTimestampsErr := BACnetEventTimestampsEnclosedParse(readBuffer , uint8( uint8(3) ) )
	if _eventTimestampsErr != nil {
		return nil, errors.Wrap(_eventTimestampsErr, "Error parsing 'eventTimestamps' field of BACnetEventSummary")
	}
	eventTimestamps := _eventTimestamps.(BACnetEventTimestampsEnclosed)
	if closeErr := readBuffer.CloseContext("eventTimestamps"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventTimestamps")
	}

	// Simple Field (notifyType)
	if pullErr := readBuffer.PullContext("notifyType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for notifyType")
	}
_notifyType, _notifyTypeErr := BACnetNotifyTypeTaggedParse(readBuffer , uint8( uint8(4) ) , TagClass( TagClass_CONTEXT_SPECIFIC_TAGS ) )
	if _notifyTypeErr != nil {
		return nil, errors.Wrap(_notifyTypeErr, "Error parsing 'notifyType' field of BACnetEventSummary")
	}
	notifyType := _notifyType.(BACnetNotifyTypeTagged)
	if closeErr := readBuffer.CloseContext("notifyType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for notifyType")
	}

	// Simple Field (eventEnable)
	if pullErr := readBuffer.PullContext("eventEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventEnable")
	}
_eventEnable, _eventEnableErr := BACnetEventTransitionBitsTaggedParse(readBuffer , uint8( uint8(5) ) , TagClass( TagClass_CONTEXT_SPECIFIC_TAGS ) )
	if _eventEnableErr != nil {
		return nil, errors.Wrap(_eventEnableErr, "Error parsing 'eventEnable' field of BACnetEventSummary")
	}
	eventEnable := _eventEnable.(BACnetEventTransitionBitsTagged)
	if closeErr := readBuffer.CloseContext("eventEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventEnable")
	}

	// Simple Field (eventPriorities)
	if pullErr := readBuffer.PullContext("eventPriorities"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventPriorities")
	}
_eventPriorities, _eventPrioritiesErr := BACnetEventPrioritiesParse(readBuffer , uint8( uint8(6) ) )
	if _eventPrioritiesErr != nil {
		return nil, errors.Wrap(_eventPrioritiesErr, "Error parsing 'eventPriorities' field of BACnetEventSummary")
	}
	eventPriorities := _eventPriorities.(BACnetEventPriorities)
	if closeErr := readBuffer.CloseContext("eventPriorities"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventPriorities")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventSummary")
	}

	// Create the instance
	return &_BACnetEventSummary{
			ObjectIdentifier: objectIdentifier,
			EventState: eventState,
			AcknowledgedTransitions: acknowledgedTransitions,
			EventTimestamps: eventTimestamps,
			NotifyType: notifyType,
			EventEnable: eventEnable,
			EventPriorities: eventPriorities,
		}, nil
}

func (m *_BACnetEventSummary) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetEventSummary"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventSummary")
	}

	// Simple Field (objectIdentifier)
	if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
	}
	_objectIdentifierErr := writeBuffer.WriteSerializable(m.GetObjectIdentifier())
	if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for objectIdentifier")
	}
	if _objectIdentifierErr != nil {
		return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
	}

	// Simple Field (eventState)
	if pushErr := writeBuffer.PushContext("eventState"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for eventState")
	}
	_eventStateErr := writeBuffer.WriteSerializable(m.GetEventState())
	if popErr := writeBuffer.PopContext("eventState"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for eventState")
	}
	if _eventStateErr != nil {
		return errors.Wrap(_eventStateErr, "Error serializing 'eventState' field")
	}

	// Simple Field (acknowledgedTransitions)
	if pushErr := writeBuffer.PushContext("acknowledgedTransitions"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for acknowledgedTransitions")
	}
	_acknowledgedTransitionsErr := writeBuffer.WriteSerializable(m.GetAcknowledgedTransitions())
	if popErr := writeBuffer.PopContext("acknowledgedTransitions"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for acknowledgedTransitions")
	}
	if _acknowledgedTransitionsErr != nil {
		return errors.Wrap(_acknowledgedTransitionsErr, "Error serializing 'acknowledgedTransitions' field")
	}

	// Simple Field (eventTimestamps)
	if pushErr := writeBuffer.PushContext("eventTimestamps"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for eventTimestamps")
	}
	_eventTimestampsErr := writeBuffer.WriteSerializable(m.GetEventTimestamps())
	if popErr := writeBuffer.PopContext("eventTimestamps"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for eventTimestamps")
	}
	if _eventTimestampsErr != nil {
		return errors.Wrap(_eventTimestampsErr, "Error serializing 'eventTimestamps' field")
	}

	// Simple Field (notifyType)
	if pushErr := writeBuffer.PushContext("notifyType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for notifyType")
	}
	_notifyTypeErr := writeBuffer.WriteSerializable(m.GetNotifyType())
	if popErr := writeBuffer.PopContext("notifyType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for notifyType")
	}
	if _notifyTypeErr != nil {
		return errors.Wrap(_notifyTypeErr, "Error serializing 'notifyType' field")
	}

	// Simple Field (eventEnable)
	if pushErr := writeBuffer.PushContext("eventEnable"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for eventEnable")
	}
	_eventEnableErr := writeBuffer.WriteSerializable(m.GetEventEnable())
	if popErr := writeBuffer.PopContext("eventEnable"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for eventEnable")
	}
	if _eventEnableErr != nil {
		return errors.Wrap(_eventEnableErr, "Error serializing 'eventEnable' field")
	}

	// Simple Field (eventPriorities)
	if pushErr := writeBuffer.PushContext("eventPriorities"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for eventPriorities")
	}
	_eventPrioritiesErr := writeBuffer.WriteSerializable(m.GetEventPriorities())
	if popErr := writeBuffer.PopContext("eventPriorities"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for eventPriorities")
	}
	if _eventPrioritiesErr != nil {
		return errors.Wrap(_eventPrioritiesErr, "Error serializing 'eventPriorities' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventSummary"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventSummary")
	}
	return nil
}


func (m *_BACnetEventSummary) isBACnetEventSummary() bool {
	return true
}

func (m *_BACnetEventSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



