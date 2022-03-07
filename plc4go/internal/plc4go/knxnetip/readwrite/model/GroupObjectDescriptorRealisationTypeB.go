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
type GroupObjectDescriptorRealisationTypeB struct {
	UpdateEnable          bool
	TransmitEnable        bool
	SegmentSelectorEnable bool
	WriteEnable           bool
	ReadEnable            bool
	CommunicationEnable   bool
	Priority              CEMIPriority
	ValueType             ComObjectValueType
}

// The corresponding interface
type IGroupObjectDescriptorRealisationTypeB interface {
	// GetUpdateEnable returns UpdateEnable (property field)
	GetUpdateEnable() bool
	// GetTransmitEnable returns TransmitEnable (property field)
	GetTransmitEnable() bool
	// GetSegmentSelectorEnable returns SegmentSelectorEnable (property field)
	GetSegmentSelectorEnable() bool
	// GetWriteEnable returns WriteEnable (property field)
	GetWriteEnable() bool
	// GetReadEnable returns ReadEnable (property field)
	GetReadEnable() bool
	// GetCommunicationEnable returns CommunicationEnable (property field)
	GetCommunicationEnable() bool
	// GetPriority returns Priority (property field)
	GetPriority() CEMIPriority
	// GetValueType returns ValueType (property field)
	GetValueType() ComObjectValueType
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *GroupObjectDescriptorRealisationTypeB) GetUpdateEnable() bool {
	return m.UpdateEnable
}

func (m *GroupObjectDescriptorRealisationTypeB) GetTransmitEnable() bool {
	return m.TransmitEnable
}

func (m *GroupObjectDescriptorRealisationTypeB) GetSegmentSelectorEnable() bool {
	return m.SegmentSelectorEnable
}

func (m *GroupObjectDescriptorRealisationTypeB) GetWriteEnable() bool {
	return m.WriteEnable
}

func (m *GroupObjectDescriptorRealisationTypeB) GetReadEnable() bool {
	return m.ReadEnable
}

func (m *GroupObjectDescriptorRealisationTypeB) GetCommunicationEnable() bool {
	return m.CommunicationEnable
}

func (m *GroupObjectDescriptorRealisationTypeB) GetPriority() CEMIPriority {
	return m.Priority
}

func (m *GroupObjectDescriptorRealisationTypeB) GetValueType() ComObjectValueType {
	return m.ValueType
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewGroupObjectDescriptorRealisationTypeB factory function for GroupObjectDescriptorRealisationTypeB
func NewGroupObjectDescriptorRealisationTypeB(updateEnable bool, transmitEnable bool, segmentSelectorEnable bool, writeEnable bool, readEnable bool, communicationEnable bool, priority CEMIPriority, valueType ComObjectValueType) *GroupObjectDescriptorRealisationTypeB {
	return &GroupObjectDescriptorRealisationTypeB{UpdateEnable: updateEnable, TransmitEnable: transmitEnable, SegmentSelectorEnable: segmentSelectorEnable, WriteEnable: writeEnable, ReadEnable: readEnable, CommunicationEnable: communicationEnable, Priority: priority, ValueType: valueType}
}

func CastGroupObjectDescriptorRealisationTypeB(structType interface{}) *GroupObjectDescriptorRealisationTypeB {
	if casted, ok := structType.(GroupObjectDescriptorRealisationTypeB); ok {
		return &casted
	}
	if casted, ok := structType.(*GroupObjectDescriptorRealisationTypeB); ok {
		return casted
	}
	return nil
}

func (m *GroupObjectDescriptorRealisationTypeB) GetTypeName() string {
	return "GroupObjectDescriptorRealisationTypeB"
}

func (m *GroupObjectDescriptorRealisationTypeB) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *GroupObjectDescriptorRealisationTypeB) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (updateEnable)
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

func (m *GroupObjectDescriptorRealisationTypeB) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func GroupObjectDescriptorRealisationTypeBParse(readBuffer utils.ReadBuffer) (*GroupObjectDescriptorRealisationTypeB, error) {
	if pullErr := readBuffer.PullContext("GroupObjectDescriptorRealisationTypeB"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (updateEnable)
	_updateEnable, _updateEnableErr := readBuffer.ReadBit("updateEnable")
	if _updateEnableErr != nil {
		return nil, errors.Wrap(_updateEnableErr, "Error parsing 'updateEnable' field")
	}
	updateEnable := _updateEnable

	// Simple Field (transmitEnable)
	_transmitEnable, _transmitEnableErr := readBuffer.ReadBit("transmitEnable")
	if _transmitEnableErr != nil {
		return nil, errors.Wrap(_transmitEnableErr, "Error parsing 'transmitEnable' field")
	}
	transmitEnable := _transmitEnable

	// Simple Field (segmentSelectorEnable)
	_segmentSelectorEnable, _segmentSelectorEnableErr := readBuffer.ReadBit("segmentSelectorEnable")
	if _segmentSelectorEnableErr != nil {
		return nil, errors.Wrap(_segmentSelectorEnableErr, "Error parsing 'segmentSelectorEnable' field")
	}
	segmentSelectorEnable := _segmentSelectorEnable

	// Simple Field (writeEnable)
	_writeEnable, _writeEnableErr := readBuffer.ReadBit("writeEnable")
	if _writeEnableErr != nil {
		return nil, errors.Wrap(_writeEnableErr, "Error parsing 'writeEnable' field")
	}
	writeEnable := _writeEnable

	// Simple Field (readEnable)
	_readEnable, _readEnableErr := readBuffer.ReadBit("readEnable")
	if _readEnableErr != nil {
		return nil, errors.Wrap(_readEnableErr, "Error parsing 'readEnable' field")
	}
	readEnable := _readEnable

	// Simple Field (communicationEnable)
	_communicationEnable, _communicationEnableErr := readBuffer.ReadBit("communicationEnable")
	if _communicationEnableErr != nil {
		return nil, errors.Wrap(_communicationEnableErr, "Error parsing 'communicationEnable' field")
	}
	communicationEnable := _communicationEnable

	// Simple Field (priority)
	if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
		return nil, pullErr
	}
	_priority, _priorityErr := CEMIPriorityParse(readBuffer)
	if _priorityErr != nil {
		return nil, errors.Wrap(_priorityErr, "Error parsing 'priority' field")
	}
	priority := _priority
	if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (valueType)
	if pullErr := readBuffer.PullContext("valueType"); pullErr != nil {
		return nil, pullErr
	}
	_valueType, _valueTypeErr := ComObjectValueTypeParse(readBuffer)
	if _valueTypeErr != nil {
		return nil, errors.Wrap(_valueTypeErr, "Error parsing 'valueType' field")
	}
	valueType := _valueType
	if closeErr := readBuffer.CloseContext("valueType"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("GroupObjectDescriptorRealisationTypeB"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewGroupObjectDescriptorRealisationTypeB(updateEnable, transmitEnable, segmentSelectorEnable, writeEnable, readEnable, communicationEnable, priority, valueType), nil
}

func (m *GroupObjectDescriptorRealisationTypeB) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("GroupObjectDescriptorRealisationTypeB"); pushErr != nil {
		return pushErr
	}

	// Simple Field (updateEnable)
	updateEnable := bool(m.UpdateEnable)
	_updateEnableErr := writeBuffer.WriteBit("updateEnable", (updateEnable))
	if _updateEnableErr != nil {
		return errors.Wrap(_updateEnableErr, "Error serializing 'updateEnable' field")
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

	if popErr := writeBuffer.PopContext("GroupObjectDescriptorRealisationTypeB"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *GroupObjectDescriptorRealisationTypeB) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
