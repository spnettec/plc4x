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
	"io"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetCOVSubscription is the corresponding interface of BACnetCOVSubscription
type BACnetCOVSubscription interface {
	utils.LengthAware
	utils.Serializable
	// GetRecipient returns Recipient (property field)
	GetRecipient() BACnetRecipientProcessEnclosed
	// GetMonitoredPropertyReference returns MonitoredPropertyReference (property field)
	GetMonitoredPropertyReference() BACnetObjectPropertyReferenceEnclosed
	// GetIssueConfirmedNotifications returns IssueConfirmedNotifications (property field)
	GetIssueConfirmedNotifications() BACnetContextTagBoolean
	// GetTimeRemaining returns TimeRemaining (property field)
	GetTimeRemaining() BACnetContextTagUnsignedInteger
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() BACnetContextTagReal
}

// BACnetCOVSubscriptionExactly can be used when we want exactly this type and not a type which fulfills BACnetCOVSubscription.
// This is useful for switch cases.
type BACnetCOVSubscriptionExactly interface {
	BACnetCOVSubscription
	isBACnetCOVSubscription() bool
}

// _BACnetCOVSubscription is the data-structure of this message
type _BACnetCOVSubscription struct {
        Recipient BACnetRecipientProcessEnclosed
        MonitoredPropertyReference BACnetObjectPropertyReferenceEnclosed
        IssueConfirmedNotifications BACnetContextTagBoolean
        TimeRemaining BACnetContextTagUnsignedInteger
        CovIncrement BACnetContextTagReal
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCOVSubscription) GetRecipient() BACnetRecipientProcessEnclosed {
	return m.Recipient
}

func (m *_BACnetCOVSubscription) GetMonitoredPropertyReference() BACnetObjectPropertyReferenceEnclosed {
	return m.MonitoredPropertyReference
}

func (m *_BACnetCOVSubscription) GetIssueConfirmedNotifications() BACnetContextTagBoolean {
	return m.IssueConfirmedNotifications
}

func (m *_BACnetCOVSubscription) GetTimeRemaining() BACnetContextTagUnsignedInteger {
	return m.TimeRemaining
}

func (m *_BACnetCOVSubscription) GetCovIncrement() BACnetContextTagReal {
	return m.CovIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetCOVSubscription factory function for _BACnetCOVSubscription
func NewBACnetCOVSubscription( recipient BACnetRecipientProcessEnclosed , monitoredPropertyReference BACnetObjectPropertyReferenceEnclosed , issueConfirmedNotifications BACnetContextTagBoolean , timeRemaining BACnetContextTagUnsignedInteger , covIncrement BACnetContextTagReal ) *_BACnetCOVSubscription {
return &_BACnetCOVSubscription{ Recipient: recipient , MonitoredPropertyReference: monitoredPropertyReference , IssueConfirmedNotifications: issueConfirmedNotifications , TimeRemaining: timeRemaining , CovIncrement: covIncrement }
}

// Deprecated: use the interface for direct cast
func CastBACnetCOVSubscription(structType interface{}) BACnetCOVSubscription {
    if casted, ok := structType.(BACnetCOVSubscription); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCOVSubscription); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCOVSubscription) GetTypeName() string {
	return "BACnetCOVSubscription"
}

func (m *_BACnetCOVSubscription) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetCOVSubscription) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits()

	// Simple field (monitoredPropertyReference)
	lengthInBits += m.MonitoredPropertyReference.GetLengthInBits()

	// Simple field (issueConfirmedNotifications)
	lengthInBits += m.IssueConfirmedNotifications.GetLengthInBits()

	// Simple field (timeRemaining)
	lengthInBits += m.TimeRemaining.GetLengthInBits()

	// Optional Field (covIncrement)
	if m.CovIncrement != nil {
		lengthInBits += m.CovIncrement.GetLengthInBits()
	}

	return lengthInBits
}


func (m *_BACnetCOVSubscription) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetCOVSubscriptionParse(readBuffer utils.ReadBuffer) (BACnetCOVSubscription, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCOVSubscription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCOVSubscription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (recipient)
	if pullErr := readBuffer.PullContext("recipient"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for recipient")
	}
_recipient, _recipientErr := BACnetRecipientProcessEnclosedParse(readBuffer , uint8( uint8(0) ) )
	if _recipientErr != nil {
		return nil, errors.Wrap(_recipientErr, "Error parsing 'recipient' field of BACnetCOVSubscription")
	}
	recipient := _recipient.(BACnetRecipientProcessEnclosed)
	if closeErr := readBuffer.CloseContext("recipient"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for recipient")
	}

	// Simple Field (monitoredPropertyReference)
	if pullErr := readBuffer.PullContext("monitoredPropertyReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for monitoredPropertyReference")
	}
_monitoredPropertyReference, _monitoredPropertyReferenceErr := BACnetObjectPropertyReferenceEnclosedParse(readBuffer , uint8( uint8(1) ) )
	if _monitoredPropertyReferenceErr != nil {
		return nil, errors.Wrap(_monitoredPropertyReferenceErr, "Error parsing 'monitoredPropertyReference' field of BACnetCOVSubscription")
	}
	monitoredPropertyReference := _monitoredPropertyReference.(BACnetObjectPropertyReferenceEnclosed)
	if closeErr := readBuffer.CloseContext("monitoredPropertyReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for monitoredPropertyReference")
	}

	// Simple Field (issueConfirmedNotifications)
	if pullErr := readBuffer.PullContext("issueConfirmedNotifications"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for issueConfirmedNotifications")
	}
_issueConfirmedNotifications, _issueConfirmedNotificationsErr := BACnetContextTagParse(readBuffer , uint8( uint8(2) ) , BACnetDataType( BACnetDataType_BOOLEAN ) )
	if _issueConfirmedNotificationsErr != nil {
		return nil, errors.Wrap(_issueConfirmedNotificationsErr, "Error parsing 'issueConfirmedNotifications' field of BACnetCOVSubscription")
	}
	issueConfirmedNotifications := _issueConfirmedNotifications.(BACnetContextTagBoolean)
	if closeErr := readBuffer.CloseContext("issueConfirmedNotifications"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for issueConfirmedNotifications")
	}

	// Simple Field (timeRemaining)
	if pullErr := readBuffer.PullContext("timeRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeRemaining")
	}
_timeRemaining, _timeRemainingErr := BACnetContextTagParse(readBuffer , uint8( uint8(3) ) , BACnetDataType( BACnetDataType_UNSIGNED_INTEGER ) )
	if _timeRemainingErr != nil {
		return nil, errors.Wrap(_timeRemainingErr, "Error parsing 'timeRemaining' field of BACnetCOVSubscription")
	}
	timeRemaining := _timeRemaining.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("timeRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeRemaining")
	}

	// Optional Field (covIncrement) (Can be skipped, if a given expression evaluates to false)
	var covIncrement BACnetContextTagReal = nil
{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("covIncrement"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for covIncrement")
		}
_val, _err := BACnetContextTagParse(readBuffer , uint8(4) , BACnetDataType_REAL )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'covIncrement' field of BACnetCOVSubscription")
		default:
			covIncrement = _val.(BACnetContextTagReal)
			if closeErr := readBuffer.CloseContext("covIncrement"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for covIncrement")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetCOVSubscription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCOVSubscription")
	}

	// Create the instance
	return &_BACnetCOVSubscription{
			Recipient: recipient,
			MonitoredPropertyReference: monitoredPropertyReference,
			IssueConfirmedNotifications: issueConfirmedNotifications,
			TimeRemaining: timeRemaining,
			CovIncrement: covIncrement,
		}, nil
}

func (m *_BACnetCOVSubscription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCOVSubscription) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetCOVSubscription"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCOVSubscription")
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

	// Simple Field (monitoredPropertyReference)
	if pushErr := writeBuffer.PushContext("monitoredPropertyReference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for monitoredPropertyReference")
	}
	_monitoredPropertyReferenceErr := writeBuffer.WriteSerializable(m.GetMonitoredPropertyReference())
	if popErr := writeBuffer.PopContext("monitoredPropertyReference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for monitoredPropertyReference")
	}
	if _monitoredPropertyReferenceErr != nil {
		return errors.Wrap(_monitoredPropertyReferenceErr, "Error serializing 'monitoredPropertyReference' field")
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

	// Simple Field (timeRemaining)
	if pushErr := writeBuffer.PushContext("timeRemaining"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timeRemaining")
	}
	_timeRemainingErr := writeBuffer.WriteSerializable(m.GetTimeRemaining())
	if popErr := writeBuffer.PopContext("timeRemaining"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timeRemaining")
	}
	if _timeRemainingErr != nil {
		return errors.Wrap(_timeRemainingErr, "Error serializing 'timeRemaining' field")
	}

	// Optional Field (covIncrement) (Can be skipped, if the value is null)
	var covIncrement BACnetContextTagReal = nil
	if m.GetCovIncrement() != nil {
		if pushErr := writeBuffer.PushContext("covIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for covIncrement")
		}
		covIncrement = m.GetCovIncrement()
		_covIncrementErr := writeBuffer.WriteSerializable(covIncrement)
		if popErr := writeBuffer.PopContext("covIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for covIncrement")
		}
		if _covIncrementErr != nil {
			return errors.Wrap(_covIncrementErr, "Error serializing 'covIncrement' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetCOVSubscription"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCOVSubscription")
	}
	return nil
}


func (m *_BACnetCOVSubscription) isBACnetCOVSubscription() bool {
	return true
}

func (m *_BACnetCOVSubscription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



