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

// S7MessageUserData is the corresponding interface of S7MessageUserData
type S7MessageUserData interface {
	utils.LengthAware
	utils.Serializable
	S7Message
}

// S7MessageUserDataExactly can be used when we want exactly this type and not a type which fulfills S7MessageUserData.
// This is useful for switch cases.
type S7MessageUserDataExactly interface {
	S7MessageUserData
	isS7MessageUserData() bool
}

// _S7MessageUserData is the data-structure of this message
type _S7MessageUserData struct {
	*_S7Message
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7MessageUserData) GetMessageType() uint8 {
	return 0x07
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7MessageUserData) InitializeParent(parent S7Message, tpduReference uint16, parameter S7Parameter, payload S7Payload) {
	m.TpduReference = tpduReference
	m.Parameter = parameter
	m.Payload = payload
}

func (m *_S7MessageUserData) GetParent() S7Message {
	return m._S7Message
}

// NewS7MessageUserData factory function for _S7MessageUserData
func NewS7MessageUserData(tpduReference uint16, parameter S7Parameter, payload S7Payload) *_S7MessageUserData {
	_result := &_S7MessageUserData{
		_S7Message: NewS7Message(tpduReference, parameter, payload),
	}
	_result._S7Message._S7MessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7MessageUserData(structType interface{}) S7MessageUserData {
	if casted, ok := structType.(S7MessageUserData); ok {
		return casted
	}
	if casted, ok := structType.(*S7MessageUserData); ok {
		return *casted
	}
	return nil
}

func (m *_S7MessageUserData) GetTypeName() string {
	return "S7MessageUserData"
}

func (m *_S7MessageUserData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_S7MessageUserData) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_S7MessageUserData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7MessageUserDataParse(readBuffer utils.ReadBuffer) (S7MessageUserData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7MessageUserData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7MessageUserData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("S7MessageUserData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7MessageUserData")
	}

	// Create a partially initialized instance
	_child := &_S7MessageUserData{
		_S7Message: &_S7Message{},
	}
	_child._S7Message._S7MessageChildRequirements = _child
	return _child, nil
}

func (m *_S7MessageUserData) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7MessageUserData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7MessageUserData")
		}

		if popErr := writeBuffer.PopContext("S7MessageUserData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7MessageUserData")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_S7MessageUserData) isS7MessageUserData() bool {
	return true
}

func (m *_S7MessageUserData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
