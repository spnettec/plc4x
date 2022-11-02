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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetDestination is the corresponding interface of BACnetDestination
type BACnetDestination interface {
	utils.LengthAware
	utils.Serializable
	// GetValidDays returns ValidDays (property field)
	GetValidDays() BACnetDaysOfWeekTagged
	// GetFromTime returns FromTime (property field)
	GetFromTime() BACnetApplicationTagTime
	// GetToTime returns ToTime (property field)
	GetToTime() BACnetApplicationTagTime
	// GetRecipient returns Recipient (property field)
	GetRecipient() BACnetRecipient
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() BACnetApplicationTagUnsignedInteger
	// GetIssueConfirmedNotifications returns IssueConfirmedNotifications (property field)
	GetIssueConfirmedNotifications() BACnetApplicationTagBoolean
	// GetTransitions returns Transitions (property field)
	GetTransitions() BACnetEventTransitionBitsTagged
}

// BACnetDestinationExactly can be used when we want exactly this type and not a type which fulfills BACnetDestination.
// This is useful for switch cases.
type BACnetDestinationExactly interface {
	BACnetDestination
	isBACnetDestination() bool
}

// _BACnetDestination is the data-structure of this message
type _BACnetDestination struct {
        ValidDays BACnetDaysOfWeekTagged
        FromTime BACnetApplicationTagTime
        ToTime BACnetApplicationTagTime
        Recipient BACnetRecipient
        ProcessIdentifier BACnetApplicationTagUnsignedInteger
        IssueConfirmedNotifications BACnetApplicationTagBoolean
        Transitions BACnetEventTransitionBitsTagged
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDestination) GetValidDays() BACnetDaysOfWeekTagged {
	return m.ValidDays
}

func (m *_BACnetDestination) GetFromTime() BACnetApplicationTagTime {
	return m.FromTime
}

func (m *_BACnetDestination) GetToTime() BACnetApplicationTagTime {
	return m.ToTime
}

func (m *_BACnetDestination) GetRecipient() BACnetRecipient {
	return m.Recipient
}

func (m *_BACnetDestination) GetProcessIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.ProcessIdentifier
}

func (m *_BACnetDestination) GetIssueConfirmedNotifications() BACnetApplicationTagBoolean {
	return m.IssueConfirmedNotifications
}

func (m *_BACnetDestination) GetTransitions() BACnetEventTransitionBitsTagged {
	return m.Transitions
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetDestination factory function for _BACnetDestination
func NewBACnetDestination( validDays BACnetDaysOfWeekTagged , fromTime BACnetApplicationTagTime , toTime BACnetApplicationTagTime , recipient BACnetRecipient , processIdentifier BACnetApplicationTagUnsignedInteger , issueConfirmedNotifications BACnetApplicationTagBoolean , transitions BACnetEventTransitionBitsTagged ) *_BACnetDestination {
return &_BACnetDestination{ ValidDays: validDays , FromTime: fromTime , ToTime: toTime , Recipient: recipient , ProcessIdentifier: processIdentifier , IssueConfirmedNotifications: issueConfirmedNotifications , Transitions: transitions }
}

// Deprecated: use the interface for direct cast
func CastBACnetDestination(structType interface{}) BACnetDestination {
    if casted, ok := structType.(BACnetDestination); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDestination); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDestination) GetTypeName() string {
	return "BACnetDestination"
}

func (m *_BACnetDestination) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetDestination) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (validDays)
	lengthInBits += m.ValidDays.GetLengthInBits()

	// Simple field (fromTime)
	lengthInBits += m.FromTime.GetLengthInBits()

	// Simple field (toTime)
	lengthInBits += m.ToTime.GetLengthInBits()

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits()

	// Simple field (processIdentifier)
	lengthInBits += m.ProcessIdentifier.GetLengthInBits()

	// Simple field (issueConfirmedNotifications)
	lengthInBits += m.IssueConfirmedNotifications.GetLengthInBits()

	// Simple field (transitions)
	lengthInBits += m.Transitions.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetDestination) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetDestinationParse(readBuffer utils.ReadBuffer) (BACnetDestination, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDestination"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDestination")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (validDays)
	if pullErr := readBuffer.PullContext("validDays"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for validDays")
	}
_validDays, _validDaysErr := BACnetDaysOfWeekTaggedParse(readBuffer , uint8( uint8(0) ) , TagClass( TagClass_APPLICATION_TAGS ) )
	if _validDaysErr != nil {
		return nil, errors.Wrap(_validDaysErr, "Error parsing 'validDays' field of BACnetDestination")
	}
	validDays := _validDays.(BACnetDaysOfWeekTagged)
	if closeErr := readBuffer.CloseContext("validDays"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for validDays")
	}

	// Simple Field (fromTime)
	if pullErr := readBuffer.PullContext("fromTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fromTime")
	}
_fromTime, _fromTimeErr := BACnetApplicationTagParse(readBuffer)
	if _fromTimeErr != nil {
		return nil, errors.Wrap(_fromTimeErr, "Error parsing 'fromTime' field of BACnetDestination")
	}
	fromTime := _fromTime.(BACnetApplicationTagTime)
	if closeErr := readBuffer.CloseContext("fromTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fromTime")
	}

	// Simple Field (toTime)
	if pullErr := readBuffer.PullContext("toTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for toTime")
	}
_toTime, _toTimeErr := BACnetApplicationTagParse(readBuffer)
	if _toTimeErr != nil {
		return nil, errors.Wrap(_toTimeErr, "Error parsing 'toTime' field of BACnetDestination")
	}
	toTime := _toTime.(BACnetApplicationTagTime)
	if closeErr := readBuffer.CloseContext("toTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for toTime")
	}

	// Simple Field (recipient)
	if pullErr := readBuffer.PullContext("recipient"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for recipient")
	}
_recipient, _recipientErr := BACnetRecipientParse(readBuffer)
	if _recipientErr != nil {
		return nil, errors.Wrap(_recipientErr, "Error parsing 'recipient' field of BACnetDestination")
	}
	recipient := _recipient.(BACnetRecipient)
	if closeErr := readBuffer.CloseContext("recipient"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for recipient")
	}

	// Simple Field (processIdentifier)
	if pullErr := readBuffer.PullContext("processIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for processIdentifier")
	}
_processIdentifier, _processIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _processIdentifierErr != nil {
		return nil, errors.Wrap(_processIdentifierErr, "Error parsing 'processIdentifier' field of BACnetDestination")
	}
	processIdentifier := _processIdentifier.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("processIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for processIdentifier")
	}

	// Simple Field (issueConfirmedNotifications)
	if pullErr := readBuffer.PullContext("issueConfirmedNotifications"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for issueConfirmedNotifications")
	}
_issueConfirmedNotifications, _issueConfirmedNotificationsErr := BACnetApplicationTagParse(readBuffer)
	if _issueConfirmedNotificationsErr != nil {
		return nil, errors.Wrap(_issueConfirmedNotificationsErr, "Error parsing 'issueConfirmedNotifications' field of BACnetDestination")
	}
	issueConfirmedNotifications := _issueConfirmedNotifications.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("issueConfirmedNotifications"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for issueConfirmedNotifications")
	}

	// Simple Field (transitions)
	if pullErr := readBuffer.PullContext("transitions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transitions")
	}
_transitions, _transitionsErr := BACnetEventTransitionBitsTaggedParse(readBuffer , uint8( uint8(0) ) , TagClass( TagClass_APPLICATION_TAGS ) )
	if _transitionsErr != nil {
		return nil, errors.Wrap(_transitionsErr, "Error parsing 'transitions' field of BACnetDestination")
	}
	transitions := _transitions.(BACnetEventTransitionBitsTagged)
	if closeErr := readBuffer.CloseContext("transitions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transitions")
	}

	if closeErr := readBuffer.CloseContext("BACnetDestination"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDestination")
	}

	// Create the instance
	return &_BACnetDestination{
			ValidDays: validDays,
			FromTime: fromTime,
			ToTime: toTime,
			Recipient: recipient,
			ProcessIdentifier: processIdentifier,
			IssueConfirmedNotifications: issueConfirmedNotifications,
			Transitions: transitions,
		}, nil
}

func (m *_BACnetDestination) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDestination) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetDestination"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDestination")
	}

	// Simple Field (validDays)
	if pushErr := writeBuffer.PushContext("validDays"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for validDays")
	}
	_validDaysErr := writeBuffer.WriteSerializable(m.GetValidDays())
	if popErr := writeBuffer.PopContext("validDays"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for validDays")
	}
	if _validDaysErr != nil {
		return errors.Wrap(_validDaysErr, "Error serializing 'validDays' field")
	}

	// Simple Field (fromTime)
	if pushErr := writeBuffer.PushContext("fromTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for fromTime")
	}
	_fromTimeErr := writeBuffer.WriteSerializable(m.GetFromTime())
	if popErr := writeBuffer.PopContext("fromTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for fromTime")
	}
	if _fromTimeErr != nil {
		return errors.Wrap(_fromTimeErr, "Error serializing 'fromTime' field")
	}

	// Simple Field (toTime)
	if pushErr := writeBuffer.PushContext("toTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for toTime")
	}
	_toTimeErr := writeBuffer.WriteSerializable(m.GetToTime())
	if popErr := writeBuffer.PopContext("toTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for toTime")
	}
	if _toTimeErr != nil {
		return errors.Wrap(_toTimeErr, "Error serializing 'toTime' field")
	}

	// Simple Field (recipient)
	if pushErr := writeBuffer.PushContext("recipient"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for recipient")
	}
	_recipientErr := writeBuffer.WriteSerializable(m.GetRecipient())
	if popErr := writeBuffer.PopContext("recipient"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for recipient")
	}
	if _recipientErr != nil {
		return errors.Wrap(_recipientErr, "Error serializing 'recipient' field")
	}

	// Simple Field (processIdentifier)
	if pushErr := writeBuffer.PushContext("processIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for processIdentifier")
	}
	_processIdentifierErr := writeBuffer.WriteSerializable(m.GetProcessIdentifier())
	if popErr := writeBuffer.PopContext("processIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for processIdentifier")
	}
	if _processIdentifierErr != nil {
		return errors.Wrap(_processIdentifierErr, "Error serializing 'processIdentifier' field")
	}

	// Simple Field (issueConfirmedNotifications)
	if pushErr := writeBuffer.PushContext("issueConfirmedNotifications"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for issueConfirmedNotifications")
	}
	_issueConfirmedNotificationsErr := writeBuffer.WriteSerializable(m.GetIssueConfirmedNotifications())
	if popErr := writeBuffer.PopContext("issueConfirmedNotifications"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for issueConfirmedNotifications")
	}
	if _issueConfirmedNotificationsErr != nil {
		return errors.Wrap(_issueConfirmedNotificationsErr, "Error serializing 'issueConfirmedNotifications' field")
	}

	// Simple Field (transitions)
	if pushErr := writeBuffer.PushContext("transitions"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for transitions")
	}
	_transitionsErr := writeBuffer.WriteSerializable(m.GetTransitions())
	if popErr := writeBuffer.PopContext("transitions"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for transitions")
	}
	if _transitionsErr != nil {
		return errors.Wrap(_transitionsErr, "Error serializing 'transitions' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDestination"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDestination")
	}
	return nil
}


func (m *_BACnetDestination) isBACnetDestination() bool {
	return true
}

func (m *_BACnetDestination) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



