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

// ApduDataExtGroupPropertyValueResponse is the data-structure of this message
type ApduDataExtGroupPropertyValueResponse struct {
	*ApduDataExt

	// Arguments.
	Length uint8
}

// IApduDataExtGroupPropertyValueResponse is the corresponding interface of ApduDataExtGroupPropertyValueResponse
type IApduDataExtGroupPropertyValueResponse interface {
	IApduDataExt
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

func (m *ApduDataExtGroupPropertyValueResponse) GetExtApciType() uint8 {
	return 0x29
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataExtGroupPropertyValueResponse) InitializeParent(parent *ApduDataExt) {}

func (m *ApduDataExtGroupPropertyValueResponse) GetParent() *ApduDataExt {
	return m.ApduDataExt
}

// NewApduDataExtGroupPropertyValueResponse factory function for ApduDataExtGroupPropertyValueResponse
func NewApduDataExtGroupPropertyValueResponse(length uint8) *ApduDataExtGroupPropertyValueResponse {
	_result := &ApduDataExtGroupPropertyValueResponse{
		ApduDataExt: NewApduDataExt(length),
	}
	_result.Child = _result
	return _result
}

func CastApduDataExtGroupPropertyValueResponse(structType interface{}) *ApduDataExtGroupPropertyValueResponse {
	if casted, ok := structType.(ApduDataExtGroupPropertyValueResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtGroupPropertyValueResponse); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtGroupPropertyValueResponse(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtGroupPropertyValueResponse(casted.Child)
	}
	return nil
}

func (m *ApduDataExtGroupPropertyValueResponse) GetTypeName() string {
	return "ApduDataExtGroupPropertyValueResponse"
}

func (m *ApduDataExtGroupPropertyValueResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtGroupPropertyValueResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtGroupPropertyValueResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtGroupPropertyValueResponseParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExtGroupPropertyValueResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtGroupPropertyValueResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtGroupPropertyValueResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtGroupPropertyValueResponse{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child, nil
}

func (m *ApduDataExtGroupPropertyValueResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtGroupPropertyValueResponse"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataExtGroupPropertyValueResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtGroupPropertyValueResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
