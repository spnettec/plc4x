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
)

// Code generated by code-generation. DO NOT EDIT.

// NotTransmittedToManyReTransmissions is the data-structure of this message
type NotTransmittedToManyReTransmissions struct {
	*Confirmation
}

// INotTransmittedToManyReTransmissions is the corresponding interface of NotTransmittedToManyReTransmissions
type INotTransmittedToManyReTransmissions interface {
	IConfirmation
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

func (m *NotTransmittedToManyReTransmissions) GetConfirmationType() byte {
	return 0x23
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *NotTransmittedToManyReTransmissions) InitializeParent(parent *Confirmation, alpha *Alpha) {
	m.Confirmation.Alpha = alpha
}

func (m *NotTransmittedToManyReTransmissions) GetParent() *Confirmation {
	return m.Confirmation
}

// NewNotTransmittedToManyReTransmissions factory function for NotTransmittedToManyReTransmissions
func NewNotTransmittedToManyReTransmissions(alpha *Alpha) *NotTransmittedToManyReTransmissions {
	_result := &NotTransmittedToManyReTransmissions{
		Confirmation: NewConfirmation(alpha),
	}
	_result.Child = _result
	return _result
}

func CastNotTransmittedToManyReTransmissions(structType interface{}) *NotTransmittedToManyReTransmissions {
	if casted, ok := structType.(NotTransmittedToManyReTransmissions); ok {
		return &casted
	}
	if casted, ok := structType.(*NotTransmittedToManyReTransmissions); ok {
		return casted
	}
	if casted, ok := structType.(Confirmation); ok {
		return CastNotTransmittedToManyReTransmissions(casted.Child)
	}
	if casted, ok := structType.(*Confirmation); ok {
		return CastNotTransmittedToManyReTransmissions(casted.Child)
	}
	return nil
}

func (m *NotTransmittedToManyReTransmissions) GetTypeName() string {
	return "NotTransmittedToManyReTransmissions"
}

func (m *NotTransmittedToManyReTransmissions) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *NotTransmittedToManyReTransmissions) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *NotTransmittedToManyReTransmissions) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func NotTransmittedToManyReTransmissionsParse(readBuffer utils.ReadBuffer) (*NotTransmittedToManyReTransmissions, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NotTransmittedToManyReTransmissions"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("NotTransmittedToManyReTransmissions"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &NotTransmittedToManyReTransmissions{
		Confirmation: &Confirmation{},
	}
	_child.Confirmation.Child = _child
	return _child, nil
}

func (m *NotTransmittedToManyReTransmissions) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NotTransmittedToManyReTransmissions"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("NotTransmittedToManyReTransmissions"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *NotTransmittedToManyReTransmissions) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
