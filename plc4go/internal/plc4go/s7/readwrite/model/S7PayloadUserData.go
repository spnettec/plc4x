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
type S7PayloadUserData struct {
	Items  []*S7PayloadUserDataItem
	Parent *S7Payload
}

// The corresponding interface
type IS7PayloadUserData interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *S7PayloadUserData) ParameterParameterType() uint8 {
	return 0x00
}

func (m *S7PayloadUserData) MessageType() uint8 {
	return 0x07
}

func (m *S7PayloadUserData) InitializeParent(parent *S7Payload) {
}

func NewS7PayloadUserData(items []*S7PayloadUserDataItem) *S7Payload {
	child := &S7PayloadUserData{
		Items:  items,
		Parent: NewS7Payload(),
	}
	child.Parent.Child = child
	return child.Parent
}

func CastS7PayloadUserData(structType interface{}) *S7PayloadUserData {
	castFunc := func(typ interface{}) *S7PayloadUserData {
		if casted, ok := typ.(S7PayloadUserData); ok {
			return &casted
		}
		if casted, ok := typ.(*S7PayloadUserData); ok {
			return casted
		}
		if casted, ok := typ.(S7Payload); ok {
			return CastS7PayloadUserData(casted.Child)
		}
		if casted, ok := typ.(*S7Payload); ok {
			return CastS7PayloadUserData(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *S7PayloadUserData) GetTypeName() string {
	return "S7PayloadUserData"
}

func (m *S7PayloadUserData) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *S7PayloadUserData) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.Parent.ParentLengthInBits())

	// Array field
	if len(m.Items) > 0 {
		for i, element := range m.Items {
			last := i == len(m.Items)-1
			lengthInBits += element.LengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *S7PayloadUserData) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func S7PayloadUserDataParse(readBuffer utils.ReadBuffer, parameter *S7Parameter) (*S7Payload, error) {
	if pullErr := readBuffer.PullContext("S7PayloadUserData"); pullErr != nil {
		return nil, pullErr
	}

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	items := make([]*S7PayloadUserDataItem, uint16(len(CastS7ParameterUserData(parameter).Items)))
	for curItem := uint16(0); curItem < uint16(uint16(len(CastS7ParameterUserData(parameter).Items))); curItem++ {
		_item, _err := S7PayloadUserDataItemParse(readBuffer, CastS7ParameterUserDataItemCPUFunctions(CastS7ParameterUserData(parameter).Items[0]).CpuFunctionType, CastS7ParameterUserDataItemCPUFunctions(CastS7ParameterUserData(parameter).Items[0]).CpuSubfunction)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'items' field")
		}
		items[curItem] = _item
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserData"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadUserData{
		Items:  items,
		Parent: &S7Payload{},
	}
	_child.Parent.Child = _child
	return _child.Parent, nil
}

func (m *S7PayloadUserData) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserData"); pushErr != nil {
			return pushErr
		}

		// Array Field (items)
		if m.Items != nil {
			if pushErr := writeBuffer.PushContext("items", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Items {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'items' field")
				}
			}
			if popErr := writeBuffer.PopContext("items", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("S7PayloadUserData"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.Parent.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadUserData) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
