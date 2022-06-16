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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetUnconfirmedServiceRequestUnconfirmedTextMessage is the corresponding interface of BACnetUnconfirmedServiceRequestUnconfirmedTextMessage
type BACnetUnconfirmedServiceRequestUnconfirmedTextMessage interface {
	BACnetUnconfirmedServiceRequest
	// GetTextMessageSourceDevice returns TextMessageSourceDevice (property field)
	GetTextMessageSourceDevice() BACnetContextTagObjectIdentifier
	// GetMessageClass returns MessageClass (property field)
	GetMessageClass() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	// GetMessagePriority returns MessagePriority (property field)
	GetMessagePriority() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
	// GetMessage returns Message (property field)
	GetMessage() BACnetContextTagCharacterString
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetUnconfirmedServiceRequestUnconfirmedTextMessage is the data-structure of this message
type _BACnetUnconfirmedServiceRequestUnconfirmedTextMessage struct {
	*_BACnetUnconfirmedServiceRequest
	TextMessageSourceDevice BACnetContextTagObjectIdentifier
	MessageClass            BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	MessagePriority         BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged
	Message                 BACnetContextTagCharacterString

	// Arguments.
	ServiceRequestLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return BACnetUnconfirmedServiceChoice_UNCONFIRMED_TEXT_MESSAGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) InitializeParent(parent BACnetUnconfirmedServiceRequest) {
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetParent() BACnetUnconfirmedServiceRequest {
	return m._BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetTextMessageSourceDevice() BACnetContextTagObjectIdentifier {
	return m.TextMessageSourceDevice
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetMessageClass() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass {
	return m.MessageClass
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetMessagePriority() BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged {
	return m.MessagePriority
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetMessage() BACnetContextTagCharacterString {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestUnconfirmedTextMessage factory function for _BACnetUnconfirmedServiceRequestUnconfirmedTextMessage
func NewBACnetUnconfirmedServiceRequestUnconfirmedTextMessage(textMessageSourceDevice BACnetContextTagObjectIdentifier, messageClass BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass, messagePriority BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged, message BACnetContextTagCharacterString, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage {
	_result := &_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage{
		TextMessageSourceDevice:          textMessageSourceDevice,
		MessageClass:                     messageClass,
		MessagePriority:                  messagePriority,
		Message:                          message,
		_BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestUnconfirmedTextMessage(structType interface{}) BACnetUnconfirmedServiceRequestUnconfirmedTextMessage {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUnconfirmedTextMessage); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUnconfirmedTextMessage); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedTextMessage"
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (textMessageSourceDevice)
	lengthInBits += m.TextMessageSourceDevice.GetLengthInBits()

	// Optional Field (messageClass)
	if m.MessageClass != nil {
		lengthInBits += m.MessageClass.GetLengthInBits()
	}

	// Simple field (messagePriority)
	lengthInBits += m.MessagePriority.GetLengthInBits()

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestUnconfirmedTextMessageParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestUnconfirmedTextMessage, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedTextMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestUnconfirmedTextMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (textMessageSourceDevice)
	if pullErr := readBuffer.PullContext("textMessageSourceDevice"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for textMessageSourceDevice")
	}
	_textMessageSourceDevice, _textMessageSourceDeviceErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _textMessageSourceDeviceErr != nil {
		return nil, errors.Wrap(_textMessageSourceDeviceErr, "Error parsing 'textMessageSourceDevice' field")
	}
	textMessageSourceDevice := _textMessageSourceDevice.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("textMessageSourceDevice"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for textMessageSourceDevice")
	}

	// Optional Field (messageClass) (Can be skipped, if a given expression evaluates to false)
	var messageClass BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("messageClass"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for messageClass")
		}
		_val, _err := BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassParse(readBuffer, uint8(1))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'messageClass' field")
		default:
			messageClass = _val.(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass)
			if closeErr := readBuffer.CloseContext("messageClass"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for messageClass")
			}
		}
	}

	// Simple Field (messagePriority)
	if pullErr := readBuffer.PullContext("messagePriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for messagePriority")
	}
	_messagePriority, _messagePriorityErr := BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTaggedParse(readBuffer, uint8(uint8(2)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _messagePriorityErr != nil {
		return nil, errors.Wrap(_messagePriorityErr, "Error parsing 'messagePriority' field")
	}
	messagePriority := _messagePriority.(BACnetConfirmedServiceRequestConfirmedTextMessageMessagePriorityTagged)
	if closeErr := readBuffer.CloseContext("messagePriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for messagePriority")
	}

	// Simple Field (message)
	if pullErr := readBuffer.PullContext("message"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for message")
	}
	_message, _messageErr := BACnetContextTagParse(readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_CHARACTER_STRING))
	if _messageErr != nil {
		return nil, errors.Wrap(_messageErr, "Error parsing 'message' field")
	}
	message := _message.(BACnetContextTagCharacterString)
	if closeErr := readBuffer.CloseContext("message"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for message")
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedTextMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestUnconfirmedTextMessage")
	}

	// Create a partially initialized instance
	_child := &_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage{
		TextMessageSourceDevice:          textMessageSourceDevice,
		MessageClass:                     messageClass,
		MessagePriority:                  messagePriority,
		Message:                          message,
		_BACnetUnconfirmedServiceRequest: &_BACnetUnconfirmedServiceRequest{},
	}
	_child._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedTextMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestUnconfirmedTextMessage")
		}

		// Simple Field (textMessageSourceDevice)
		if pushErr := writeBuffer.PushContext("textMessageSourceDevice"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for textMessageSourceDevice")
		}
		_textMessageSourceDeviceErr := writeBuffer.WriteSerializable(m.GetTextMessageSourceDevice())
		if popErr := writeBuffer.PopContext("textMessageSourceDevice"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for textMessageSourceDevice")
		}
		if _textMessageSourceDeviceErr != nil {
			return errors.Wrap(_textMessageSourceDeviceErr, "Error serializing 'textMessageSourceDevice' field")
		}

		// Optional Field (messageClass) (Can be skipped, if the value is null)
		var messageClass BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass = nil
		if m.GetMessageClass() != nil {
			if pushErr := writeBuffer.PushContext("messageClass"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for messageClass")
			}
			messageClass = m.GetMessageClass()
			_messageClassErr := writeBuffer.WriteSerializable(messageClass)
			if popErr := writeBuffer.PopContext("messageClass"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for messageClass")
			}
			if _messageClassErr != nil {
				return errors.Wrap(_messageClassErr, "Error serializing 'messageClass' field")
			}
		}

		// Simple Field (messagePriority)
		if pushErr := writeBuffer.PushContext("messagePriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for messagePriority")
		}
		_messagePriorityErr := writeBuffer.WriteSerializable(m.GetMessagePriority())
		if popErr := writeBuffer.PopContext("messagePriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for messagePriority")
		}
		if _messagePriorityErr != nil {
			return errors.Wrap(_messagePriorityErr, "Error serializing 'messagePriority' field")
		}

		// Simple Field (message)
		if pushErr := writeBuffer.PushContext("message"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for message")
		}
		_messageErr := writeBuffer.WriteSerializable(m.GetMessage())
		if popErr := writeBuffer.PopContext("message"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for message")
		}
		if _messageErr != nil {
			return errors.Wrap(_messageErr, "Error serializing 'message' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedTextMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestUnconfirmedTextMessage")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestUnconfirmedTextMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
