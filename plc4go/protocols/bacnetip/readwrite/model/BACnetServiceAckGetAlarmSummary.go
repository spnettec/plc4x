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

// BACnetServiceAckGetAlarmSummary is the corresponding interface of BACnetServiceAckGetAlarmSummary
type BACnetServiceAckGetAlarmSummary interface {
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetApplicationTagObjectIdentifier
	// GetEventState returns EventState (property field)
	GetEventState() BACnetEventStateTagged
	// GetAcknowledgedTransitions returns AcknowledgedTransitions (property field)
	GetAcknowledgedTransitions() BACnetEventTransitionBitsTagged
}

// BACnetServiceAckGetAlarmSummaryExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckGetAlarmSummary.
// This is useful for switch cases.
type BACnetServiceAckGetAlarmSummaryExactly interface {
	BACnetServiceAckGetAlarmSummary
	isBACnetServiceAckGetAlarmSummary() bool
}

// _BACnetServiceAckGetAlarmSummary is the data-structure of this message
type _BACnetServiceAckGetAlarmSummary struct {
	*_BACnetServiceAck
	ObjectIdentifier        BACnetApplicationTagObjectIdentifier
	EventState              BACnetEventStateTagged
	AcknowledgedTransitions BACnetEventTransitionBitsTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckGetAlarmSummary) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_GET_ALARM_SUMMARY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckGetAlarmSummary) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckGetAlarmSummary) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckGetAlarmSummary) GetObjectIdentifier() BACnetApplicationTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetServiceAckGetAlarmSummary) GetEventState() BACnetEventStateTagged {
	return m.EventState
}

func (m *_BACnetServiceAckGetAlarmSummary) GetAcknowledgedTransitions() BACnetEventTransitionBitsTagged {
	return m.AcknowledgedTransitions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckGetAlarmSummary factory function for _BACnetServiceAckGetAlarmSummary
func NewBACnetServiceAckGetAlarmSummary(objectIdentifier BACnetApplicationTagObjectIdentifier, eventState BACnetEventStateTagged, acknowledgedTransitions BACnetEventTransitionBitsTagged, serviceAckLength uint32) *_BACnetServiceAckGetAlarmSummary {
	_result := &_BACnetServiceAckGetAlarmSummary{
		ObjectIdentifier:        objectIdentifier,
		EventState:              eventState,
		AcknowledgedTransitions: acknowledgedTransitions,
		_BACnetServiceAck:       NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckGetAlarmSummary(structType interface{}) BACnetServiceAckGetAlarmSummary {
	if casted, ok := structType.(BACnetServiceAckGetAlarmSummary); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckGetAlarmSummary); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckGetAlarmSummary) GetTypeName() string {
	return "BACnetServiceAckGetAlarmSummary"
}

func (m *_BACnetServiceAckGetAlarmSummary) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetServiceAckGetAlarmSummary) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	// Simple field (eventState)
	lengthInBits += m.EventState.GetLengthInBits()

	// Simple field (acknowledgedTransitions)
	lengthInBits += m.AcknowledgedTransitions.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetServiceAckGetAlarmSummary) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckGetAlarmSummaryParse(theBytes []byte, serviceAckLength uint32) (BACnetServiceAckGetAlarmSummary, error) {
	return BACnetServiceAckGetAlarmSummaryParseWithBuffer(utils.NewReadBufferByteBased(theBytes), serviceAckLength)
}

func BACnetServiceAckGetAlarmSummaryParseWithBuffer(readBuffer utils.ReadBuffer, serviceAckLength uint32) (BACnetServiceAckGetAlarmSummary, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckGetAlarmSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckGetAlarmSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field of BACnetServiceAckGetAlarmSummary")
	}
	objectIdentifier := _objectIdentifier.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	// Simple Field (eventState)
	if pullErr := readBuffer.PullContext("eventState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for eventState")
	}
	_eventState, _eventStateErr := BACnetEventStateTaggedParseWithBuffer(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _eventStateErr != nil {
		return nil, errors.Wrap(_eventStateErr, "Error parsing 'eventState' field of BACnetServiceAckGetAlarmSummary")
	}
	eventState := _eventState.(BACnetEventStateTagged)
	if closeErr := readBuffer.CloseContext("eventState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for eventState")
	}

	// Simple Field (acknowledgedTransitions)
	if pullErr := readBuffer.PullContext("acknowledgedTransitions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for acknowledgedTransitions")
	}
	_acknowledgedTransitions, _acknowledgedTransitionsErr := BACnetEventTransitionBitsTaggedParseWithBuffer(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _acknowledgedTransitionsErr != nil {
		return nil, errors.Wrap(_acknowledgedTransitionsErr, "Error parsing 'acknowledgedTransitions' field of BACnetServiceAckGetAlarmSummary")
	}
	acknowledgedTransitions := _acknowledgedTransitions.(BACnetEventTransitionBitsTagged)
	if closeErr := readBuffer.CloseContext("acknowledgedTransitions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for acknowledgedTransitions")
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckGetAlarmSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckGetAlarmSummary")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckGetAlarmSummary{
		_BACnetServiceAck: &_BACnetServiceAck{
			ServiceAckLength: serviceAckLength,
		},
		ObjectIdentifier:        objectIdentifier,
		EventState:              eventState,
		AcknowledgedTransitions: acknowledgedTransitions,
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckGetAlarmSummary) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckGetAlarmSummary) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckGetAlarmSummary"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckGetAlarmSummary")
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

		if popErr := writeBuffer.PopContext("BACnetServiceAckGetAlarmSummary"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckGetAlarmSummary")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetServiceAckGetAlarmSummary) isBACnetServiceAckGetAlarmSummary() bool {
	return true
}

func (m *_BACnetServiceAckGetAlarmSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
