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
type ComObjectTableRealisationType1 struct {
	*ComObjectTable
	NumEntries           uint8
	RamFlagsTablePointer uint8
	ComObjectDescriptors []*GroupObjectDescriptorRealisationType1
}

// The corresponding interface
type IComObjectTableRealisationType1 interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *ComObjectTableRealisationType1) FirmwareType() FirmwareType {
	return FirmwareType_SYSTEM_1
}

func (m *ComObjectTableRealisationType1) InitializeParent(parent *ComObjectTable) {}

func NewComObjectTableRealisationType1(numEntries uint8, ramFlagsTablePointer uint8, comObjectDescriptors []*GroupObjectDescriptorRealisationType1) *ComObjectTable {
	child := &ComObjectTableRealisationType1{
		NumEntries:           numEntries,
		RamFlagsTablePointer: ramFlagsTablePointer,
		ComObjectDescriptors: comObjectDescriptors,
		ComObjectTable:       NewComObjectTable(),
	}
	child.Child = child
	return child.ComObjectTable
}

func CastComObjectTableRealisationType1(structType interface{}) *ComObjectTableRealisationType1 {
	castFunc := func(typ interface{}) *ComObjectTableRealisationType1 {
		if casted, ok := typ.(ComObjectTableRealisationType1); ok {
			return &casted
		}
		if casted, ok := typ.(*ComObjectTableRealisationType1); ok {
			return casted
		}
		if casted, ok := typ.(ComObjectTable); ok {
			return CastComObjectTableRealisationType1(casted.Child)
		}
		if casted, ok := typ.(*ComObjectTable); ok {
			return CastComObjectTableRealisationType1(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *ComObjectTableRealisationType1) GetTypeName() string {
	return "ComObjectTableRealisationType1"
}

func (m *ComObjectTableRealisationType1) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *ComObjectTableRealisationType1) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (numEntries)
	lengthInBits += 8

	// Simple field (ramFlagsTablePointer)
	lengthInBits += 8

	// Array field
	if len(m.ComObjectDescriptors) > 0 {
		for i, element := range m.ComObjectDescriptors {
			last := i == len(m.ComObjectDescriptors)-1
			lengthInBits += element.LengthInBitsConditional(last)
		}
	}

	return lengthInBits
}

func (m *ComObjectTableRealisationType1) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func ComObjectTableRealisationType1Parse(readBuffer utils.ReadBuffer, firmwareType FirmwareType) (*ComObjectTable, error) {
	if pullErr := readBuffer.PullContext("ComObjectTableRealisationType1"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (numEntries)
	_numEntries, _numEntriesErr := readBuffer.ReadUint8("numEntries", 8)
	if _numEntriesErr != nil {
		return nil, errors.Wrap(_numEntriesErr, "Error parsing 'numEntries' field")
	}
	numEntries := _numEntries

	// Simple Field (ramFlagsTablePointer)
	_ramFlagsTablePointer, _ramFlagsTablePointerErr := readBuffer.ReadUint8("ramFlagsTablePointer", 8)
	if _ramFlagsTablePointerErr != nil {
		return nil, errors.Wrap(_ramFlagsTablePointerErr, "Error parsing 'ramFlagsTablePointer' field")
	}
	ramFlagsTablePointer := _ramFlagsTablePointer

	// Array field (comObjectDescriptors)
	if pullErr := readBuffer.PullContext("comObjectDescriptors", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Count array
	comObjectDescriptors := make([]*GroupObjectDescriptorRealisationType1, numEntries)
	{
		for curItem := uint16(0); curItem < uint16(numEntries); curItem++ {
			_item, _err := GroupObjectDescriptorRealisationType1Parse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'comObjectDescriptors' field")
			}
			comObjectDescriptors[curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("comObjectDescriptors", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("ComObjectTableRealisationType1"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ComObjectTableRealisationType1{
		NumEntries:           numEntries,
		RamFlagsTablePointer: ramFlagsTablePointer,
		ComObjectDescriptors: comObjectDescriptors,
		ComObjectTable:       &ComObjectTable{},
	}
	_child.ComObjectTable.Child = _child
	return _child.ComObjectTable, nil
}

func (m *ComObjectTableRealisationType1) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ComObjectTableRealisationType1"); pushErr != nil {
			return pushErr
		}

		// Simple Field (numEntries)
		numEntries := uint8(m.NumEntries)
		_numEntriesErr := writeBuffer.WriteUint8("numEntries", 8, (numEntries))
		if _numEntriesErr != nil {
			return errors.Wrap(_numEntriesErr, "Error serializing 'numEntries' field")
		}

		// Simple Field (ramFlagsTablePointer)
		ramFlagsTablePointer := uint8(m.RamFlagsTablePointer)
		_ramFlagsTablePointerErr := writeBuffer.WriteUint8("ramFlagsTablePointer", 8, (ramFlagsTablePointer))
		if _ramFlagsTablePointerErr != nil {
			return errors.Wrap(_ramFlagsTablePointerErr, "Error serializing 'ramFlagsTablePointer' field")
		}

		// Array Field (comObjectDescriptors)
		if m.ComObjectDescriptors != nil {
			if pushErr := writeBuffer.PushContext("comObjectDescriptors", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.ComObjectDescriptors {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'comObjectDescriptors' field")
				}
			}
			if popErr := writeBuffer.PopContext("comObjectDescriptors", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("ComObjectTableRealisationType1"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ComObjectTableRealisationType1) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
