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

// BVLCReadForeignDeviceTableAck is the data-structure of this message
type BVLCReadForeignDeviceTableAck struct {
	*BVLC
	Table []*BVLCForeignDeviceTableEntry

	// Arguments.
	BvlcPayloadLength uint16
}

// IBVLCReadForeignDeviceTableAck is the corresponding interface of BVLCReadForeignDeviceTableAck
type IBVLCReadForeignDeviceTableAck interface {
	IBVLC
	// GetTable returns Table (property field)
	GetTable() []*BVLCForeignDeviceTableEntry
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

func (m *BVLCReadForeignDeviceTableAck) GetBvlcFunction() uint8 {
	return 0x07
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BVLCReadForeignDeviceTableAck) InitializeParent(parent *BVLC) {}

func (m *BVLCReadForeignDeviceTableAck) GetParent() *BVLC {
	return m.BVLC
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BVLCReadForeignDeviceTableAck) GetTable() []*BVLCForeignDeviceTableEntry {
	return m.Table
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBVLCReadForeignDeviceTableAck factory function for BVLCReadForeignDeviceTableAck
func NewBVLCReadForeignDeviceTableAck(table []*BVLCForeignDeviceTableEntry, bvlcPayloadLength uint16) *BVLCReadForeignDeviceTableAck {
	_result := &BVLCReadForeignDeviceTableAck{
		Table: table,
		BVLC:  NewBVLC(),
	}
	_result.Child = _result
	return _result
}

func CastBVLCReadForeignDeviceTableAck(structType interface{}) *BVLCReadForeignDeviceTableAck {
	if casted, ok := structType.(BVLCReadForeignDeviceTableAck); ok {
		return &casted
	}
	if casted, ok := structType.(*BVLCReadForeignDeviceTableAck); ok {
		return casted
	}
	if casted, ok := structType.(BVLC); ok {
		return CastBVLCReadForeignDeviceTableAck(casted.Child)
	}
	if casted, ok := structType.(*BVLC); ok {
		return CastBVLCReadForeignDeviceTableAck(casted.Child)
	}
	return nil
}

func (m *BVLCReadForeignDeviceTableAck) GetTypeName() string {
	return "BVLCReadForeignDeviceTableAck"
}

func (m *BVLCReadForeignDeviceTableAck) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BVLCReadForeignDeviceTableAck) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Table) > 0 {
		for _, element := range m.Table {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BVLCReadForeignDeviceTableAck) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BVLCReadForeignDeviceTableAckParse(readBuffer utils.ReadBuffer, bvlcPayloadLength uint16) (*BVLCReadForeignDeviceTableAck, error) {
	if pullErr := readBuffer.PullContext("BVLCReadForeignDeviceTableAck"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Array field (table)
	if pullErr := readBuffer.PullContext("table", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Length array
	table := make([]*BVLCForeignDeviceTableEntry, 0)
	{
		_tableLength := bvlcPayloadLength
		_tableEndPos := readBuffer.GetPos() + uint16(_tableLength)
		for readBuffer.GetPos() < _tableEndPos {
			_item, _err := BVLCForeignDeviceTableEntryParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'table' field")
			}
			table = append(table, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("table", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BVLCReadForeignDeviceTableAck"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BVLCReadForeignDeviceTableAck{
		Table: table,
		BVLC:  &BVLC{},
	}
	_child.BVLC.Child = _child
	return _child, nil
}

func (m *BVLCReadForeignDeviceTableAck) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BVLCReadForeignDeviceTableAck"); pushErr != nil {
			return pushErr
		}

		// Array Field (table)
		if m.Table != nil {
			if pushErr := writeBuffer.PushContext("table", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.Table {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'table' field")
				}
			}
			if popErr := writeBuffer.PopContext("table", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BVLCReadForeignDeviceTableAck"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BVLCReadForeignDeviceTableAck) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
