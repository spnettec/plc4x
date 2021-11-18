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

// The data-structure of this message
type ApduDataExtWriteRoutingTableRequest struct {
	Parent *ApduDataExt
}

// The corresponding interface
type IApduDataExtWriteRoutingTableRequest interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ApduDataExtWriteRoutingTableRequest) ExtApciType() uint8 {
	return 0x03
}

func (m *ApduDataExtWriteRoutingTableRequest) InitializeParent(parent *ApduDataExt) {
}

func NewApduDataExtWriteRoutingTableRequest() *ApduDataExt {
	child := &ApduDataExtWriteRoutingTableRequest{
		Parent: NewApduDataExt(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastApduDataExtWriteRoutingTableRequest(structType interface{}) *ApduDataExtWriteRoutingTableRequest {
	castFunc := func(typ interface{}) *ApduDataExtWriteRoutingTableRequest {
		if casted, ok := typ.(ApduDataExtWriteRoutingTableRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ApduDataExtWriteRoutingTableRequest); ok {
			return casted
		}
		if casted, ok := typ.(ApduDataExt); ok {
			return CastApduDataExtWriteRoutingTableRequest(casted.Child)
		}
		if casted, ok := typ.(*ApduDataExt); ok {
			return CastApduDataExtWriteRoutingTableRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ApduDataExtWriteRoutingTableRequest) GetTypeName() string {
	return "ApduDataExtWriteRoutingTableRequest"
}

func (m *ApduDataExtWriteRoutingTableRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ApduDataExtWriteRoutingTableRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	return lengthInBits
}

func (m *ApduDataExtWriteRoutingTableRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ApduDataExtWriteRoutingTableRequestParse(readBuffer utils.ReadBuffer, length uint8) (*ApduDataExt, error) {
	if pullErr := readBuffer.PullContext("ApduDataExtWriteRoutingTableRequest"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("ApduDataExtWriteRoutingTableRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ApduDataExtWriteRoutingTableRequest{
		Parent: &ApduDataExt{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *ApduDataExtWriteRoutingTableRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtWriteRoutingTableRequest"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("ApduDataExtWriteRoutingTableRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataExtWriteRoutingTableRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
