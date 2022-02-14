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
type ModbusPDUReadFileRecordRequest struct {
	*ModbusPDU
	Items []*ModbusPDUReadFileRecordRequestItem
}

// The corresponding interface
type IModbusPDUReadFileRecordRequest interface {
	// GetItems returns Items
	GetItems() []*ModbusPDUReadFileRecordRequestItem
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ModbusPDUReadFileRecordRequest) ErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadFileRecordRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadFileRecordRequest) FunctionFlag() uint8 {
	return 0x14
}

func (m *ModbusPDUReadFileRecordRequest) GetFunctionFlag() uint8 {
	return 0x14
}

func (m *ModbusPDUReadFileRecordRequest) Response() bool {
	return bool(false)
}

func (m *ModbusPDUReadFileRecordRequest) GetResponse() bool {
	return bool(false)
}

func (m *ModbusPDUReadFileRecordRequest) InitializeParent(parent *ModbusPDU) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ModbusPDUReadFileRecordRequest) GetItems() []*ModbusPDUReadFileRecordRequestItem {
	return m.Items
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewModbusPDUReadFileRecordRequest factory function for ModbusPDUReadFileRecordRequest
func NewModbusPDUReadFileRecordRequest(items []*ModbusPDUReadFileRecordRequestItem) *ModbusPDU {
	child := &ModbusPDUReadFileRecordRequest{
		Items:     items,
		ModbusPDU: NewModbusPDU(),
	}
	child.Child = child
	return child.ModbusPDU
}

func CastModbusPDUReadFileRecordRequest(structType interface{}) *ModbusPDUReadFileRecordRequest {
	castFunc := func(typ interface{}) *ModbusPDUReadFileRecordRequest {
		if casted, ok := typ.(ModbusPDUReadFileRecordRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*ModbusPDUReadFileRecordRequest); ok {
			return casted
		}
		if casted, ok := typ.(ModbusPDU); ok {
			return CastModbusPDUReadFileRecordRequest(casted.Child)
		}
		if casted, ok := typ.(*ModbusPDU); ok {
			return CastModbusPDUReadFileRecordRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ModbusPDUReadFileRecordRequest) GetTypeName() string {
	return "ModbusPDUReadFileRecordRequest"
}

func (m *ModbusPDUReadFileRecordRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUReadFileRecordRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Items) > 0 {
		for _, element := range m.Items {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *ModbusPDUReadFileRecordRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadFileRecordRequestParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDU, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUReadFileRecordRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field")
	}

	// Array field (items)
	if pullErr := readBuffer.PullContext("items", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	items := make([]*ModbusPDUReadFileRecordRequestItem, 0)
	{
		_itemsLength := byteCount
		_itemsEndPos := readBuffer.GetPos() + uint16(_itemsLength)
		for readBuffer.GetPos() < _itemsEndPos {
			_item, _err := ModbusPDUReadFileRecordRequestItemParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'items' field")
			}
			items = append(items, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("items", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadFileRecordRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadFileRecordRequest{
		Items:     items,
		ModbusPDU: &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child.ModbusPDU, nil
}

func (m *ModbusPDUReadFileRecordRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	itemsArraySizeInBytes := func(items []*ModbusPDUReadFileRecordRequestItem) uint32 {
		var sizeInBytes uint32 = 0
		for _, v := range items {
			sizeInBytes += uint32(v.GetLengthInBytes())
		}
		return sizeInBytes
	}
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadFileRecordRequest"); pushErr != nil {
			return pushErr
		}

		// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		byteCount := uint8(uint8(itemsArraySizeInBytes(m.GetItems())))
		_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
		if _byteCountErr != nil {
			return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
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

		if popErr := writeBuffer.PopContext("ModbusPDUReadFileRecordRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUReadFileRecordRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
