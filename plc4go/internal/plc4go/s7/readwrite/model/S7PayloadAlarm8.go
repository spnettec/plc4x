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

// The data-structure of this message
type S7PayloadAlarm8 struct {
	*S7PayloadUserDataItem
	AlarmMessage *AlarmMessagePushType
}

// The corresponding interface
type IS7PayloadAlarm8 interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7PayloadAlarm8) CpuFunctionType() uint8 {
	return 0x00
}

func (m *S7PayloadAlarm8) CpuSubfunction() uint8 {
	return 0x05
}

func (m *S7PayloadAlarm8) DataLength() uint16 {
	return 0
}

func (m *S7PayloadAlarm8) InitializeParent(parent *S7PayloadUserDataItem, returnCode DataTransportErrorCode, transportSize DataTransportSize) {
	m.S7PayloadUserDataItem.ReturnCode = returnCode
	m.S7PayloadUserDataItem.TransportSize = transportSize
}

func NewS7PayloadAlarm8(alarmMessage *AlarmMessagePushType, returnCode DataTransportErrorCode, transportSize DataTransportSize) *S7PayloadUserDataItem {
	child := &S7PayloadAlarm8{
		AlarmMessage:          alarmMessage,
		S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize),
	}
	child.Child = child
	return child.S7PayloadUserDataItem
}

func CastS7PayloadAlarm8(structType interface{}) *S7PayloadAlarm8 {
	castFunc := func(typ interface{}) *S7PayloadAlarm8 {
		if casted, ok := typ.(S7PayloadAlarm8); ok {
			return &casted
		}
		if casted, ok := typ.(*S7PayloadAlarm8); ok {
			return casted
		}
		if casted, ok := typ.(S7PayloadUserDataItem); ok {
			return CastS7PayloadAlarm8(casted.Child)
		}
		if casted, ok := typ.(*S7PayloadUserDataItem); ok {
			return CastS7PayloadAlarm8(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7PayloadAlarm8) GetTypeName() string {
	return "S7PayloadAlarm8"
}

func (m *S7PayloadAlarm8) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7PayloadAlarm8) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (alarmMessage)
	lengthInBits += m.AlarmMessage.LengthInBits()

	return lengthInBits
}

func (m *S7PayloadAlarm8) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7PayloadAlarm8Parse(readBuffer utils.ReadBuffer, cpuFunctionType uint8, cpuSubfunction uint8) (*S7PayloadUserDataItem, error) {
	if pullErr := readBuffer.PullContext("S7PayloadAlarm8"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (alarmMessage)
	if pullErr := readBuffer.PullContext("alarmMessage"); pullErr != nil {
		return nil, pullErr
	}
	_alarmMessage, _alarmMessageErr := AlarmMessagePushTypeParse(readBuffer)
	if _alarmMessageErr != nil {
		return nil, errors.Wrap(_alarmMessageErr, "Error parsing 'alarmMessage' field")
	}
	alarmMessage := CastAlarmMessagePushType(_alarmMessage)
	if closeErr := readBuffer.CloseContext("alarmMessage"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7PayloadAlarm8"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadAlarm8{
		AlarmMessage:          CastAlarmMessagePushType(alarmMessage),
		S7PayloadUserDataItem: &S7PayloadUserDataItem{},
	}
	_child.S7PayloadUserDataItem.Child = _child
	return _child.S7PayloadUserDataItem, nil
}

func (m *S7PayloadAlarm8) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadAlarm8"); pushErr != nil {
			return pushErr
		}

		// Simple Field (alarmMessage)
		if pushErr := writeBuffer.PushContext("alarmMessage"); pushErr != nil {
			return pushErr
		}
		_alarmMessageErr := m.AlarmMessage.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("alarmMessage"); popErr != nil {
			return popErr
		}
		if _alarmMessageErr != nil {
			return errors.Wrap(_alarmMessageErr, "Error serializing 'alarmMessage' field")
		}

		if popErr := writeBuffer.PopContext("S7PayloadAlarm8"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadAlarm8) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
