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

// BACnetConfirmedServiceRequestAcknowledgeAlarm is the data-structure of this message
type BACnetConfirmedServiceRequestAcknowledgeAlarm struct {
	*BACnetConfirmedServiceRequest
	AcknowledgingProcessIdentifier *BACnetContextTagUnsignedInteger
	EventObjectIdentifier          *BACnetContextTagObjectIdentifier
	EventStateAcknowledged         *BACnetContextTagEventState
	Timestamp                      *BACnetTimeStampEnclosed
	AcknowledgmentSource           *BACnetContextTagCharacterString
	TimeOfAcknowledgment           *BACnetTimeStampEnclosed

	// Arguments.
	ServiceRequestLength uint16
}

// IBACnetConfirmedServiceRequestAcknowledgeAlarm is the corresponding interface of BACnetConfirmedServiceRequestAcknowledgeAlarm
type IBACnetConfirmedServiceRequestAcknowledgeAlarm interface {
	IBACnetConfirmedServiceRequest
	// GetAcknowledgingProcessIdentifier returns AcknowledgingProcessIdentifier (property field)
	GetAcknowledgingProcessIdentifier() *BACnetContextTagUnsignedInteger
	// GetEventObjectIdentifier returns EventObjectIdentifier (property field)
	GetEventObjectIdentifier() *BACnetContextTagObjectIdentifier
	// GetEventStateAcknowledged returns EventStateAcknowledged (property field)
	GetEventStateAcknowledged() *BACnetContextTagEventState
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() *BACnetTimeStampEnclosed
	// GetAcknowledgmentSource returns AcknowledgmentSource (property field)
	GetAcknowledgmentSource() *BACnetContextTagCharacterString
	// GetTimeOfAcknowledgment returns TimeOfAcknowledgment (property field)
	GetTimeOfAcknowledgment() *BACnetTimeStampEnclosed
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

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ACKNOWLEDGE_ALARM
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) InitializeParent(parent *BACnetConfirmedServiceRequest) {
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) GetParent() *BACnetConfirmedServiceRequest {
	return m.BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) GetAcknowledgingProcessIdentifier() *BACnetContextTagUnsignedInteger {
	return m.AcknowledgingProcessIdentifier
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) GetEventObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.EventObjectIdentifier
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) GetEventStateAcknowledged() *BACnetContextTagEventState {
	return m.EventStateAcknowledged
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) GetTimestamp() *BACnetTimeStampEnclosed {
	return m.Timestamp
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) GetAcknowledgmentSource() *BACnetContextTagCharacterString {
	return m.AcknowledgmentSource
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) GetTimeOfAcknowledgment() *BACnetTimeStampEnclosed {
	return m.TimeOfAcknowledgment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestAcknowledgeAlarm factory function for BACnetConfirmedServiceRequestAcknowledgeAlarm
func NewBACnetConfirmedServiceRequestAcknowledgeAlarm(acknowledgingProcessIdentifier *BACnetContextTagUnsignedInteger, eventObjectIdentifier *BACnetContextTagObjectIdentifier, eventStateAcknowledged *BACnetContextTagEventState, timestamp *BACnetTimeStampEnclosed, acknowledgmentSource *BACnetContextTagCharacterString, timeOfAcknowledgment *BACnetTimeStampEnclosed, serviceRequestLength uint16) *BACnetConfirmedServiceRequestAcknowledgeAlarm {
	_result := &BACnetConfirmedServiceRequestAcknowledgeAlarm{
		AcknowledgingProcessIdentifier: acknowledgingProcessIdentifier,
		EventObjectIdentifier:          eventObjectIdentifier,
		EventStateAcknowledged:         eventStateAcknowledged,
		Timestamp:                      timestamp,
		AcknowledgmentSource:           acknowledgmentSource,
		TimeOfAcknowledgment:           timeOfAcknowledgment,
		BACnetConfirmedServiceRequest:  NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConfirmedServiceRequestAcknowledgeAlarm(structType interface{}) *BACnetConfirmedServiceRequestAcknowledgeAlarm {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAcknowledgeAlarm); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAcknowledgeAlarm); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestAcknowledgeAlarm(casted.Child)
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequest); ok {
		return CastBACnetConfirmedServiceRequestAcknowledgeAlarm(casted.Child)
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAcknowledgeAlarm"
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (acknowledgingProcessIdentifier)
	lengthInBits += m.AcknowledgingProcessIdentifier.GetLengthInBits()

	// Simple field (eventObjectIdentifier)
	lengthInBits += m.EventObjectIdentifier.GetLengthInBits()

	// Simple field (eventStateAcknowledged)
	lengthInBits += m.EventStateAcknowledged.GetLengthInBits()

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits()

	// Simple field (acknowledgmentSource)
	lengthInBits += m.AcknowledgmentSource.GetLengthInBits()

	// Simple field (timeOfAcknowledgment)
	lengthInBits += m.TimeOfAcknowledgment.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestAcknowledgeAlarmParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (*BACnetConfirmedServiceRequestAcknowledgeAlarm, error) {
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAcknowledgeAlarm"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (acknowledgingProcessIdentifier)
	if pullErr := readBuffer.PullContext("acknowledgingProcessIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_acknowledgingProcessIdentifier, _acknowledgingProcessIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _acknowledgingProcessIdentifierErr != nil {
		return nil, errors.Wrap(_acknowledgingProcessIdentifierErr, "Error parsing 'acknowledgingProcessIdentifier' field")
	}
	acknowledgingProcessIdentifier := CastBACnetContextTagUnsignedInteger(_acknowledgingProcessIdentifier)
	if closeErr := readBuffer.CloseContext("acknowledgingProcessIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (eventObjectIdentifier)
	if pullErr := readBuffer.PullContext("eventObjectIdentifier"); pullErr != nil {
		return nil, pullErr
	}
	_eventObjectIdentifier, _eventObjectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _eventObjectIdentifierErr != nil {
		return nil, errors.Wrap(_eventObjectIdentifierErr, "Error parsing 'eventObjectIdentifier' field")
	}
	eventObjectIdentifier := CastBACnetContextTagObjectIdentifier(_eventObjectIdentifier)
	if closeErr := readBuffer.CloseContext("eventObjectIdentifier"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (eventStateAcknowledged)
	if pullErr := readBuffer.PullContext("eventStateAcknowledged"); pullErr != nil {
		return nil, pullErr
	}
	_eventStateAcknowledged, _eventStateAcknowledgedErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_EVENT_STATE))
	if _eventStateAcknowledgedErr != nil {
		return nil, errors.Wrap(_eventStateAcknowledgedErr, "Error parsing 'eventStateAcknowledged' field")
	}
	eventStateAcknowledged := CastBACnetContextTagEventState(_eventStateAcknowledged)
	if closeErr := readBuffer.CloseContext("eventStateAcknowledged"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (timestamp)
	if pullErr := readBuffer.PullContext("timestamp"); pullErr != nil {
		return nil, pullErr
	}
	_timestamp, _timestampErr := BACnetTimeStampEnclosedParse(readBuffer, uint8(uint8(3)))
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field")
	}
	timestamp := CastBACnetTimeStampEnclosed(_timestamp)
	if closeErr := readBuffer.CloseContext("timestamp"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (acknowledgmentSource)
	if pullErr := readBuffer.PullContext("acknowledgmentSource"); pullErr != nil {
		return nil, pullErr
	}
	_acknowledgmentSource, _acknowledgmentSourceErr := BACnetContextTagParse(readBuffer, uint8(uint8(4)), BACnetDataType(BACnetDataType_CHARACTER_STRING))
	if _acknowledgmentSourceErr != nil {
		return nil, errors.Wrap(_acknowledgmentSourceErr, "Error parsing 'acknowledgmentSource' field")
	}
	acknowledgmentSource := CastBACnetContextTagCharacterString(_acknowledgmentSource)
	if closeErr := readBuffer.CloseContext("acknowledgmentSource"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (timeOfAcknowledgment)
	if pullErr := readBuffer.PullContext("timeOfAcknowledgment"); pullErr != nil {
		return nil, pullErr
	}
	_timeOfAcknowledgment, _timeOfAcknowledgmentErr := BACnetTimeStampEnclosedParse(readBuffer, uint8(uint8(5)))
	if _timeOfAcknowledgmentErr != nil {
		return nil, errors.Wrap(_timeOfAcknowledgmentErr, "Error parsing 'timeOfAcknowledgment' field")
	}
	timeOfAcknowledgment := CastBACnetTimeStampEnclosed(_timeOfAcknowledgment)
	if closeErr := readBuffer.CloseContext("timeOfAcknowledgment"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAcknowledgeAlarm"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConfirmedServiceRequestAcknowledgeAlarm{
		AcknowledgingProcessIdentifier: CastBACnetContextTagUnsignedInteger(acknowledgingProcessIdentifier),
		EventObjectIdentifier:          CastBACnetContextTagObjectIdentifier(eventObjectIdentifier),
		EventStateAcknowledged:         CastBACnetContextTagEventState(eventStateAcknowledged),
		Timestamp:                      CastBACnetTimeStampEnclosed(timestamp),
		AcknowledgmentSource:           CastBACnetContextTagCharacterString(acknowledgmentSource),
		TimeOfAcknowledgment:           CastBACnetTimeStampEnclosed(timeOfAcknowledgment),
		BACnetConfirmedServiceRequest:  &BACnetConfirmedServiceRequest{},
	}
	_child.BACnetConfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAcknowledgeAlarm"); pushErr != nil {
			return pushErr
		}

		// Simple Field (acknowledgingProcessIdentifier)
		if pushErr := writeBuffer.PushContext("acknowledgingProcessIdentifier"); pushErr != nil {
			return pushErr
		}
		_acknowledgingProcessIdentifierErr := m.AcknowledgingProcessIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("acknowledgingProcessIdentifier"); popErr != nil {
			return popErr
		}
		if _acknowledgingProcessIdentifierErr != nil {
			return errors.Wrap(_acknowledgingProcessIdentifierErr, "Error serializing 'acknowledgingProcessIdentifier' field")
		}

		// Simple Field (eventObjectIdentifier)
		if pushErr := writeBuffer.PushContext("eventObjectIdentifier"); pushErr != nil {
			return pushErr
		}
		_eventObjectIdentifierErr := m.EventObjectIdentifier.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("eventObjectIdentifier"); popErr != nil {
			return popErr
		}
		if _eventObjectIdentifierErr != nil {
			return errors.Wrap(_eventObjectIdentifierErr, "Error serializing 'eventObjectIdentifier' field")
		}

		// Simple Field (eventStateAcknowledged)
		if pushErr := writeBuffer.PushContext("eventStateAcknowledged"); pushErr != nil {
			return pushErr
		}
		_eventStateAcknowledgedErr := m.EventStateAcknowledged.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("eventStateAcknowledged"); popErr != nil {
			return popErr
		}
		if _eventStateAcknowledgedErr != nil {
			return errors.Wrap(_eventStateAcknowledgedErr, "Error serializing 'eventStateAcknowledged' field")
		}

		// Simple Field (timestamp)
		if pushErr := writeBuffer.PushContext("timestamp"); pushErr != nil {
			return pushErr
		}
		_timestampErr := m.Timestamp.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timestamp"); popErr != nil {
			return popErr
		}
		if _timestampErr != nil {
			return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
		}

		// Simple Field (acknowledgmentSource)
		if pushErr := writeBuffer.PushContext("acknowledgmentSource"); pushErr != nil {
			return pushErr
		}
		_acknowledgmentSourceErr := m.AcknowledgmentSource.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("acknowledgmentSource"); popErr != nil {
			return popErr
		}
		if _acknowledgmentSourceErr != nil {
			return errors.Wrap(_acknowledgmentSourceErr, "Error serializing 'acknowledgmentSource' field")
		}

		// Simple Field (timeOfAcknowledgment)
		if pushErr := writeBuffer.PushContext("timeOfAcknowledgment"); pushErr != nil {
			return pushErr
		}
		_timeOfAcknowledgmentErr := m.TimeOfAcknowledgment.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timeOfAcknowledgment"); popErr != nil {
			return popErr
		}
		if _timeOfAcknowledgmentErr != nil {
			return errors.Wrap(_timeOfAcknowledgmentErr, "Error serializing 'timeOfAcknowledgment' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAcknowledgeAlarm"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConfirmedServiceRequestAcknowledgeAlarm) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
