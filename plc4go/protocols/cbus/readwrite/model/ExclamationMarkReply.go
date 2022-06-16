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

// ExclamationMarkReply is the corresponding interface of ExclamationMarkReply
type ExclamationMarkReply interface {
	Reply
	// GetIsA returns IsA (property field)
	GetIsA() ExclamationMark
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _ExclamationMarkReply is the data-structure of this message
type _ExclamationMarkReply struct {
	*_Reply
	IsA ExclamationMark
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ExclamationMarkReply) InitializeParent(parent Reply, magicByte byte) {
	m.MagicByte = magicByte
}

func (m *_ExclamationMarkReply) GetParent() Reply {
	return m._Reply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ExclamationMarkReply) GetIsA() ExclamationMark {
	return m.IsA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewExclamationMarkReply factory function for _ExclamationMarkReply
func NewExclamationMarkReply(isA ExclamationMark, magicByte byte) *_ExclamationMarkReply {
	_result := &_ExclamationMarkReply{
		IsA:    isA,
		_Reply: NewReply(magicByte),
	}
	_result._Reply._ReplyChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastExclamationMarkReply(structType interface{}) ExclamationMarkReply {
	if casted, ok := structType.(ExclamationMarkReply); ok {
		return casted
	}
	if casted, ok := structType.(*ExclamationMarkReply); ok {
		return *casted
	}
	return nil
}

func (m *_ExclamationMarkReply) GetTypeName() string {
	return "ExclamationMarkReply"
}

func (m *_ExclamationMarkReply) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ExclamationMarkReply) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (isA)
	lengthInBits += m.IsA.GetLengthInBits()

	return lengthInBits
}

func (m *_ExclamationMarkReply) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ExclamationMarkReplyParse(readBuffer utils.ReadBuffer) (ExclamationMarkReply, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ExclamationMarkReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ExclamationMarkReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (isA)
	if pullErr := readBuffer.PullContext("isA"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for isA")
	}
	_isA, _isAErr := ExclamationMarkParse(readBuffer)
	if _isAErr != nil {
		return nil, errors.Wrap(_isAErr, "Error parsing 'isA' field")
	}
	isA := _isA.(ExclamationMark)
	if closeErr := readBuffer.CloseContext("isA"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for isA")
	}

	if closeErr := readBuffer.CloseContext("ExclamationMarkReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ExclamationMarkReply")
	}

	// Create a partially initialized instance
	_child := &_ExclamationMarkReply{
		IsA:    isA,
		_Reply: &_Reply{},
	}
	_child._Reply._ReplyChildRequirements = _child
	return _child, nil
}

func (m *_ExclamationMarkReply) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ExclamationMarkReply"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ExclamationMarkReply")
		}

		// Simple Field (isA)
		if pushErr := writeBuffer.PushContext("isA"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for isA")
		}
		_isAErr := writeBuffer.WriteSerializable(m.GetIsA())
		if popErr := writeBuffer.PopContext("isA"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for isA")
		}
		if _isAErr != nil {
			return errors.Wrap(_isAErr, "Error serializing 'isA' field")
		}

		if popErr := writeBuffer.PopContext("ExclamationMarkReply"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ExclamationMarkReply")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ExclamationMarkReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
