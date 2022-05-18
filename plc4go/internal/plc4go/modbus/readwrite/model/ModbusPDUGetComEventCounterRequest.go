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

// ModbusPDUGetComEventCounterRequest is the data-structure of this message
type ModbusPDUGetComEventCounterRequest struct {
	*ModbusPDU
}

// IModbusPDUGetComEventCounterRequest is the corresponding interface of ModbusPDUGetComEventCounterRequest
type IModbusPDUGetComEventCounterRequest interface {
	IModbusPDU
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

func (m *ModbusPDUGetComEventCounterRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUGetComEventCounterRequest) GetFunctionFlag() uint8 {
	return 0x0B
}

func (m *ModbusPDUGetComEventCounterRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusPDUGetComEventCounterRequest) InitializeParent(parent *ModbusPDU) {}

func (m *ModbusPDUGetComEventCounterRequest) GetParent() *ModbusPDU {
	return m.ModbusPDU
}

// NewModbusPDUGetComEventCounterRequest factory function for ModbusPDUGetComEventCounterRequest
func NewModbusPDUGetComEventCounterRequest() *ModbusPDUGetComEventCounterRequest {
	_result := &ModbusPDUGetComEventCounterRequest{
		ModbusPDU: NewModbusPDU(),
	}
	_result.Child = _result
	return _result
}

func CastModbusPDUGetComEventCounterRequest(structType interface{}) *ModbusPDUGetComEventCounterRequest {
	if casted, ok := structType.(ModbusPDUGetComEventCounterRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUGetComEventCounterRequest); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUGetComEventCounterRequest(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUGetComEventCounterRequest(casted.Child)
	}
	return nil
}

func (m *ModbusPDUGetComEventCounterRequest) GetTypeName() string {
	return "ModbusPDUGetComEventCounterRequest"
}

func (m *ModbusPDUGetComEventCounterRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUGetComEventCounterRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ModbusPDUGetComEventCounterRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUGetComEventCounterRequestParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDUGetComEventCounterRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUGetComEventCounterRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ModbusPDUGetComEventCounterRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUGetComEventCounterRequest{
		ModbusPDU: &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child, nil
}

func (m *ModbusPDUGetComEventCounterRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUGetComEventCounterRequest"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ModbusPDUGetComEventCounterRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUGetComEventCounterRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
