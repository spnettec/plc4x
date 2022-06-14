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

// ApduDataExtWriteRouterMemoryRequest is the data-structure of this message
type ApduDataExtWriteRouterMemoryRequest struct {
	*ApduDataExt

	// Arguments.
	Length uint8
}

// IApduDataExtWriteRouterMemoryRequest is the corresponding interface of ApduDataExtWriteRouterMemoryRequest
type IApduDataExtWriteRouterMemoryRequest interface {
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

func (m *ApduDataExtWriteRouterMemoryRequest) GetExtApciType() uint8 {
	return 0x0A
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataExtWriteRouterMemoryRequest) InitializeParent(parent *ApduDataExt) {}

func (m *ApduDataExtWriteRouterMemoryRequest) GetParent() *ApduDataExt {
	return m.ApduDataExt
}

// NewApduDataExtWriteRouterMemoryRequest factory function for ApduDataExtWriteRouterMemoryRequest
func NewApduDataExtWriteRouterMemoryRequest(length uint8) *ApduDataExtWriteRouterMemoryRequest {
	_result := &ApduDataExtWriteRouterMemoryRequest{
		ApduDataExt: NewApduDataExt(length),
	}
	_result.Child = _result
	return _result
}

func CastApduDataExtWriteRouterMemoryRequest(structType interface{}) *ApduDataExtWriteRouterMemoryRequest {
	if casted, ok := structType.(ApduDataExtWriteRouterMemoryRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataExtWriteRouterMemoryRequest); ok {
		return casted
	}
	if casted, ok := structType.(ApduDataExt); ok {
		return CastApduDataExtWriteRouterMemoryRequest(casted.Child)
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return CastApduDataExtWriteRouterMemoryRequest(casted.Child)
	}
	return nil
}

func (m *ApduDataExtWriteRouterMemoryRequest) GetTypeName() string {
	return "ApduDataExtWriteRouterMemoryRequest"
}

func (m *ApduDataExtWriteRouterMemoryRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataExtWriteRouterMemoryRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtWriteRouterMemoryRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtWriteRouterMemoryRequestParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExtWriteRouterMemoryRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtWriteRouterMemoryRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtWriteRouterMemoryRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtWriteRouterMemoryRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtWriteRouterMemoryRequest")
	}

	// Create a partially initialized instance
	_child := &ApduDataExtWriteRouterMemoryRequest{
		ApduDataExt: &ApduDataExt{},
	}
	_child.ApduDataExt.Child = _child
	return _child, nil
}

func (m *ApduDataExtWriteRouterMemoryRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtWriteRouterMemoryRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtWriteRouterMemoryRequest")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtWriteRouterMemoryRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtWriteRouterMemoryRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtWriteRouterMemoryRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
