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
	"github.com/rs/zerolog/log"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type GroupObjectDescriptorRealisationType1 struct {
	DataPointer           uint8
	TransmitEnable        bool
	SegmentSelectorEnable bool
	WriteEnable           bool
	ReadEnable            bool
	CommunicationEnable   bool
	Priority              CEMIPriority
	ValueType             ComObjectValueType
}

// The corresponding interface
type IGroupObjectDescriptorRealisationType1 interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewGroupObjectDescriptorRealisationType1(dataPointer uint8, transmitEnable bool, segmentSelectorEnable bool, writeEnable bool, readEnable bool, communicationEnable bool, priority CEMIPriority, valueType ComObjectValueType) *GroupObjectDescriptorRealisationType1 {
	return &GroupObjectDescriptorRealisationType1{DataPointer: dataPointer, TransmitEnable: transmitEnable, SegmentSelectorEnable: segmentSelectorEnable, WriteEnable: writeEnable, ReadEnable: readEnable, CommunicationEnable: communicationEnable, Priority: priority, ValueType: valueType}
}

func CastGroupObjectDescriptorRealisationType1(structType interface{}) *GroupObjectDescriptorRealisationType1 {
	castFunc := func(typ interface{}) *GroupObjectDescriptorRealisationType1 {
		if casted, ok := typ.(GroupObjectDescriptorRealisationType1); ok {
			return &casted
		}
		if casted, ok := typ.(*GroupObjectDescriptorRealisationType1); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *GroupObjectDescriptorRealisationType1) GetTypeName() string {
	return "GroupObjectDescriptorRealisationType1"
}

func (m *GroupObjectDescriptorRealisationType1) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *GroupObjectDescriptorRealisationType1) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (dataPointer)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (transmitEnable)
	lengthInBits += 1

	// Simple field (segmentSelectorEnable)
	lengthInBits += 1

	// Simple field (writeEnable)
	lengthInBits += 1

	// Simple field (readEnable)
	lengthInBits += 1

	// Simple field (communicationEnable)
	lengthInBits += 1

	// Simple field (priority)
	lengthInBits += 2

	// Simple field (valueType)
	lengthInBits += 8

	return lengthInBits
}

func (m *GroupObjectDescriptorRealisationType1) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func GroupObjectDescriptorRealisationType1Parse(readBuffer utils.ReadBuffer) (*GroupObjectDescriptorRealisationType1, error) {
	if pullErr := readBuffer.PullContext("GroupObjectDescriptorRealisationType1"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (dataPointer)
	dataPointer, _dataPointerErr := readBuffer.ReadUint8("dataPointer", 8)
	if _dataPointerErr != nil {
		return nil, errors.Wrap(_dataPointerErr, "Error parsing 'dataPointer' field")
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 1)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x1) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x1),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (transmitEnable)
	transmitEnable, _transmitEnableErr := readBuffer.ReadBit("transmitEnable")
	if _transmitEnableErr != nil {
		return nil, errors.Wrap(_transmitEnableErr, "Error parsing 'transmitEnable' field")
	}

	// Simple Field (segmentSelectorEnable)
	segmentSelectorEnable, _segmentSelectorEnableErr := readBuffer.ReadBit("segmentSelectorEnable")
	if _segmentSelectorEnableErr != nil {
		return nil, errors.Wrap(_segmentSelectorEnableErr, "Error parsing 'segmentSelectorEnable' field")
	}

	// Simple Field (writeEnable)
	writeEnable, _writeEnableErr := readBuffer.ReadBit("writeEnable")
	if _writeEnableErr != nil {
		return nil, errors.Wrap(_writeEnableErr, "Error parsing 'writeEnable' field")
	}

	// Simple Field (readEnable)
	readEnable, _readEnableErr := readBuffer.ReadBit("readEnable")
	if _readEnableErr != nil {
		return nil, errors.Wrap(_readEnableErr, "Error parsing 'readEnable' field")
	}

	// Simple Field (communicationEnable)
	communicationEnable, _communicationEnableErr := readBuffer.ReadBit("communicationEnable")
	if _communicationEnableErr != nil {
		return nil, errors.Wrap(_communicationEnableErr, "Error parsing 'communicationEnable' field")
	}

	// Simple Field (priority)
	if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
		return nil, pullErr
	}
	priority, _priorityErr := CEMIPriorityParse(readBuffer)
	if _priorityErr != nil {
		return nil, errors.Wrap(_priorityErr, "Error parsing 'priority' field")
	}
	if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (valueType)
	if pullErr := readBuffer.PullContext("valueType"); pullErr != nil {
		return nil, pullErr
	}
	valueType, _valueTypeErr := ComObjectValueTypeParse(readBuffer)
	if _valueTypeErr != nil {
		return nil, errors.Wrap(_valueTypeErr, "Error parsing 'valueType' field")
	}
	if closeErr := readBuffer.CloseContext("valueType"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("GroupObjectDescriptorRealisationType1"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewGroupObjectDescriptorRealisationType1(dataPointer, transmitEnable, segmentSelectorEnable, writeEnable, readEnable, communicationEnable, priority, valueType), nil
}

func (m *GroupObjectDescriptorRealisationType1) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("GroupObjectDescriptorRealisationType1"); pushErr != nil {
		return pushErr
	}

	// Simple Field (dataPointer)
	dataPointer := uint8(m.DataPointer)
	_dataPointerErr := writeBuffer.WriteUint8("dataPointer", 8, (dataPointer))
	if _dataPointerErr != nil {
		return errors.Wrap(_dataPointerErr, "Error serializing 'dataPointer' field")
	}

	// Reserved Field (reserved)
	{
		_err := writeBuffer.WriteUint8("reserved", 1, uint8(0x1))
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (transmitEnable)
	transmitEnable := bool(m.TransmitEnable)
	_transmitEnableErr := writeBuffer.WriteBit("transmitEnable", (transmitEnable))
	if _transmitEnableErr != nil {
		return errors.Wrap(_transmitEnableErr, "Error serializing 'transmitEnable' field")
	}

	// Simple Field (segmentSelectorEnable)
	segmentSelectorEnable := bool(m.SegmentSelectorEnable)
	_segmentSelectorEnableErr := writeBuffer.WriteBit("segmentSelectorEnable", (segmentSelectorEnable))
	if _segmentSelectorEnableErr != nil {
		return errors.Wrap(_segmentSelectorEnableErr, "Error serializing 'segmentSelectorEnable' field")
	}

	// Simple Field (writeEnable)
	writeEnable := bool(m.WriteEnable)
	_writeEnableErr := writeBuffer.WriteBit("writeEnable", (writeEnable))
	if _writeEnableErr != nil {
		return errors.Wrap(_writeEnableErr, "Error serializing 'writeEnable' field")
	}

	// Simple Field (readEnable)
	readEnable := bool(m.ReadEnable)
	_readEnableErr := writeBuffer.WriteBit("readEnable", (readEnable))
	if _readEnableErr != nil {
		return errors.Wrap(_readEnableErr, "Error serializing 'readEnable' field")
	}

	// Simple Field (communicationEnable)
	communicationEnable := bool(m.CommunicationEnable)
	_communicationEnableErr := writeBuffer.WriteBit("communicationEnable", (communicationEnable))
	if _communicationEnableErr != nil {
		return errors.Wrap(_communicationEnableErr, "Error serializing 'communicationEnable' field")
	}

	// Simple Field (priority)
	if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
		return pushErr
	}
	_priorityErr := m.Priority.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("priority"); popErr != nil {
		return popErr
	}
	if _priorityErr != nil {
		return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
	}

	// Simple Field (valueType)
	if pushErr := writeBuffer.PushContext("valueType"); pushErr != nil {
		return pushErr
	}
	_valueTypeErr := m.ValueType.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("valueType"); popErr != nil {
		return popErr
	}
	if _valueTypeErr != nil {
		return errors.Wrap(_valueTypeErr, "Error serializing 'valueType' field")
	}

	if popErr := writeBuffer.PopContext("GroupObjectDescriptorRealisationType1"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *GroupObjectDescriptorRealisationType1) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
