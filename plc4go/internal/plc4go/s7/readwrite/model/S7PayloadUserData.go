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

// S7PayloadUserData is the data-structure of this message
type S7PayloadUserData struct {
	*S7Payload
	Items []*S7PayloadUserDataItem

	// Arguments.
	Parameter S7Parameter
}

// IS7PayloadUserData is the corresponding interface of S7PayloadUserData
type IS7PayloadUserData interface {
	IS7Payload
	// GetItems returns Items (property field)
	GetItems() []*S7PayloadUserDataItem
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

func (m *S7PayloadUserData) GetParameterParameterType() uint8 {
	return 0x00
}

func (m *S7PayloadUserData) GetMessageType() uint8 {
	return 0x07
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *S7PayloadUserData) InitializeParent(parent *S7Payload) {}

func (m *S7PayloadUserData) GetParent() *S7Payload {
	return m.S7Payload
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *S7PayloadUserData) GetItems() []*S7PayloadUserDataItem {
	return m.Items
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7PayloadUserData factory function for S7PayloadUserData
func NewS7PayloadUserData(items []*S7PayloadUserDataItem, parameter S7Parameter) *S7PayloadUserData {
	_result := &S7PayloadUserData{
		Items:     items,
		S7Payload: NewS7Payload(parameter),
	}
	_result.Child = _result
	return _result
}

func CastS7PayloadUserData(structType interface{}) *S7PayloadUserData {
	if casted, ok := structType.(S7PayloadUserData); ok {
		return &casted
	}
	if casted, ok := structType.(*S7PayloadUserData); ok {
		return casted
	}
	if casted, ok := structType.(S7Payload); ok {
		return CastS7PayloadUserData(casted.Child)
	}
	if casted, ok := structType.(*S7Payload); ok {
		return CastS7PayloadUserData(casted.Child)
	}
	return nil
}

func (m *S7PayloadUserData) GetTypeName() string {
	return "S7PayloadUserData"
}

func (m *S7PayloadUserData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7PayloadUserData) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Items) > 0 {
		for i, element := range m.Items {
			last := i == len(m.Items)-1
			lengthInBits += element.GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *S7PayloadUserData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7PayloadUserDataParse(readBuffer utils.ReadBuffer, messageType uint8, parameter *S7Parameter) (*S7PayloadUserData, error) {
	if pullErr := readBuffer.PullContext("S7PayloadUserData"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	items := make([]*S7PayloadUserDataItem, uint16(len(CastS7ParameterUserData(parameter).GetItems())))
	{
		for curItem := uint16(0); curItem < uint16(uint16(len(CastS7ParameterUserData(parameter).GetItems()))); curItem++ {
			_item, _err := S7PayloadUserDataItemParse(readBuffer, CastS7ParameterUserDataItemCPUFunctions(CastS7ParameterUserData(parameter).GetItems()[0]).GetCpuFunctionType(), CastS7ParameterUserDataItemCPUFunctions(CastS7ParameterUserData(parameter).GetItems()[0]).GetCpuSubfunction())
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field")
			}
			items[curItem] = CastS7PayloadUserDataItem(_item)
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserData"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7PayloadUserData{
		Items:     items,
		S7Payload: &S7Payload{},
	}
	_child.S7Payload.Child = _child
	return _child, nil
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
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7PayloadUserData) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
