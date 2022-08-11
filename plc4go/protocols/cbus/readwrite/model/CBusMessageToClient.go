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

// CBusMessageToClient is the corresponding interface of CBusMessageToClient
type CBusMessageToClient interface {
	utils.LengthAware
	utils.Serializable
	CBusMessage
	// GetReply returns Reply (property field)
	GetReply() ReplyOrConfirmation
}

// CBusMessageToClientExactly can be used when we want exactly this type and not a type which fulfills CBusMessageToClient.
// This is useful for switch cases.
type CBusMessageToClientExactly interface {
	CBusMessageToClient
	isCBusMessageToClient() bool
}

// _CBusMessageToClient is the data-structure of this message
type _CBusMessageToClient struct {
	*_CBusMessage
	Reply ReplyOrConfirmation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CBusMessageToClient) GetIsResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CBusMessageToClient) InitializeParent(parent CBusMessage) {}

func (m *_CBusMessageToClient) GetParent() CBusMessage {
	return m._CBusMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusMessageToClient) GetReply() ReplyOrConfirmation {
	return m.Reply
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusMessageToClient factory function for _CBusMessageToClient
func NewCBusMessageToClient(reply ReplyOrConfirmation, requestContext RequestContext, cBusOptions CBusOptions) *_CBusMessageToClient {
	_result := &_CBusMessageToClient{
		Reply:        reply,
		_CBusMessage: NewCBusMessage(requestContext, cBusOptions),
	}
	_result._CBusMessage._CBusMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCBusMessageToClient(structType interface{}) CBusMessageToClient {
	if casted, ok := structType.(CBusMessageToClient); ok {
		return casted
	}
	if casted, ok := structType.(*CBusMessageToClient); ok {
		return *casted
	}
	return nil
}

func (m *_CBusMessageToClient) GetTypeName() string {
	return "CBusMessageToClient"
}

func (m *_CBusMessageToClient) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CBusMessageToClient) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (reply)
	lengthInBits += m.Reply.GetLengthInBits()

	return lengthInBits
}

func (m *_CBusMessageToClient) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusMessageToClientParse(readBuffer utils.ReadBuffer, isResponse bool, requestContext RequestContext, cBusOptions CBusOptions) (CBusMessageToClient, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusMessageToClient"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusMessageToClient")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (reply)
	if pullErr := readBuffer.PullContext("reply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for reply")
	}
	_reply, _replyErr := ReplyOrConfirmationParse(readBuffer, cBusOptions, requestContext)
	if _replyErr != nil {
		return nil, errors.Wrap(_replyErr, "Error parsing 'reply' field of CBusMessageToClient")
	}
	reply := _reply.(ReplyOrConfirmation)
	if closeErr := readBuffer.CloseContext("reply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for reply")
	}

	if closeErr := readBuffer.CloseContext("CBusMessageToClient"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusMessageToClient")
	}

	// Create a partially initialized instance
	_child := &_CBusMessageToClient{
		_CBusMessage: &_CBusMessage{
			RequestContext: requestContext,
			CBusOptions:    cBusOptions,
		},
		Reply: reply,
	}
	_child._CBusMessage._CBusMessageChildRequirements = _child
	return _child, nil
}

func (m *_CBusMessageToClient) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusMessageToClient"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusMessageToClient")
		}

		// Simple Field (reply)
		if pushErr := writeBuffer.PushContext("reply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for reply")
		}
		_replyErr := writeBuffer.WriteSerializable(m.GetReply())
		if popErr := writeBuffer.PopContext("reply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for reply")
		}
		if _replyErr != nil {
			return errors.Wrap(_replyErr, "Error serializing 'reply' field")
		}

		if popErr := writeBuffer.PopContext("CBusMessageToClient"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusMessageToClient")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CBusMessageToClient) isCBusMessageToClient() bool {
	return true
}

func (m *_CBusMessageToClient) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
