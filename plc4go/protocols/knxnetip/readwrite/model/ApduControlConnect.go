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

// ApduControlConnect is the data-structure of this message
type ApduControlConnect struct {
	*ApduControl
}

// IApduControlConnect is the corresponding interface of ApduControlConnect
type IApduControlConnect interface {
	IApduControl
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

func (m *ApduControlConnect) GetControlType() uint8 {
	return 0x0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduControlConnect) InitializeParent(parent *ApduControl) {}

func (m *ApduControlConnect) GetParent() *ApduControl {
	return m.ApduControl
}

// NewApduControlConnect factory function for ApduControlConnect
func NewApduControlConnect() *ApduControlConnect {
	_result := &ApduControlConnect{
		ApduControl: NewApduControl(),
	}
	_result.Child = _result
	return _result
}

func CastApduControlConnect(structType interface{}) *ApduControlConnect {
	if casted, ok := structType.(ApduControlConnect); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduControlConnect); ok {
		return casted
	}
	if casted, ok := structType.(ApduControl); ok {
		return CastApduControlConnect(casted.Child)
	}
	if casted, ok := structType.(*ApduControl); ok {
		return CastApduControlConnect(casted.Child)
	}
	return nil
}

func (m *ApduControlConnect) GetTypeName() string {
	return "ApduControlConnect"
}

func (m *ApduControlConnect) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduControlConnect) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduControlConnect) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduControlConnectParse(readBuffer utils.ReadBuffer) (*ApduControlConnect, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduControlConnect"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduControlConnect")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduControlConnect"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduControlConnect")
	}

	// Create a partially initialized instance
	_child := &ApduControlConnect{
		ApduControl: &ApduControl{},
	}
	_child.ApduControl.Child = _child
	return _child, nil
}

func (m *ApduControlConnect) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduControlConnect"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduControlConnect")
		}

		if popErr := writeBuffer.PopContext("ApduControlConnect"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduControlConnect")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduControlConnect) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
