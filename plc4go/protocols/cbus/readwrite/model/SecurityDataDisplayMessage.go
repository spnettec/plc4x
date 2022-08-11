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
)

// Code generated by code-generation. DO NOT EDIT.

// SecurityDataDisplayMessage is the corresponding interface of SecurityDataDisplayMessage
type SecurityDataDisplayMessage interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetMessage returns Message (property field)
	GetMessage() string
}

// SecurityDataDisplayMessageExactly can be used when we want exactly this type and not a type which fulfills SecurityDataDisplayMessage.
// This is useful for switch cases.
type SecurityDataDisplayMessageExactly interface {
	SecurityDataDisplayMessage
	isSecurityDataDisplayMessage() bool
}

// _SecurityDataDisplayMessage is the data-structure of this message
type _SecurityDataDisplayMessage struct {
	*_SecurityData
	Message string
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataDisplayMessage) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataDisplayMessage) GetParent() SecurityData {
	return m._SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataDisplayMessage) GetMessage() string {
	return m.Message
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataDisplayMessage factory function for _SecurityDataDisplayMessage
func NewSecurityDataDisplayMessage(message string, commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataDisplayMessage {
	_result := &_SecurityDataDisplayMessage{
		Message:       message,
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataDisplayMessage(structType interface{}) SecurityDataDisplayMessage {
	if casted, ok := structType.(SecurityDataDisplayMessage); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataDisplayMessage); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataDisplayMessage) GetTypeName() string {
	return "SecurityDataDisplayMessage"
}

func (m *_SecurityDataDisplayMessage) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataDisplayMessage) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (message)
	lengthInBits += uint16(int32((int32(m.GetCommandTypeContainer().NumBytes()) - int32(int32(1)))) * int32(int32(8)))

	return lengthInBits
}

func (m *_SecurityDataDisplayMessage) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataDisplayMessageParse(readBuffer utils.ReadBuffer, commandTypeContainer SecurityCommandTypeContainer) (SecurityDataDisplayMessage, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataDisplayMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataDisplayMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (message)
	_message, _messageErr := readBuffer.ReadString("message", uint32(((commandTypeContainer.NumBytes())-(1))*(8)))
	if _messageErr != nil {
		return nil, errors.Wrap(_messageErr, "Error parsing 'message' field of SecurityDataDisplayMessage")
	}
	message := _message

	if closeErr := readBuffer.CloseContext("SecurityDataDisplayMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataDisplayMessage")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataDisplayMessage{
		_SecurityData: &_SecurityData{},
		Message:       message,
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataDisplayMessage) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataDisplayMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataDisplayMessage")
		}

		// Simple Field (message)
		message := string(m.GetMessage())
		_messageErr := writeBuffer.WriteString("message", uint32(((m.GetCommandTypeContainer().NumBytes())-(1))*(8)), "UTF-8", (message))
		if _messageErr != nil {
			return errors.Wrap(_messageErr, "Error serializing 'message' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataDisplayMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataDisplayMessage")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SecurityDataDisplayMessage) isSecurityDataDisplayMessage() bool {
	return true
}

func (m *_SecurityDataDisplayMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
