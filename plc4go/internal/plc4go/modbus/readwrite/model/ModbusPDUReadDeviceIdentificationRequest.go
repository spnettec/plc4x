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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const ModbusPDUReadDeviceIdentificationRequest_MEITYPE uint8 = 0x0E

// The data-structure of this message
type ModbusPDUReadDeviceIdentificationRequest struct {
	*ModbusPDU
	Level    ModbusDeviceInformationLevel
	ObjectId uint8
}

// The corresponding interface
type IModbusPDUReadDeviceIdentificationRequest interface {
	IModbusPDU
	// GetLevel returns Level (property field)
	GetLevel() ModbusDeviceInformationLevel
	// GetObjectId returns ObjectId (property field)
	GetObjectId() uint8
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
func (m *ModbusPDUReadDeviceIdentificationRequest) GetErrorFlag() bool {
	return bool(false)
}

func (m *ModbusPDUReadDeviceIdentificationRequest) GetFunctionFlag() uint8 {
	return 0x2B
}

func (m *ModbusPDUReadDeviceIdentificationRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ModbusPDUReadDeviceIdentificationRequest) InitializeParent(parent *ModbusPDU) {}

func (m *ModbusPDUReadDeviceIdentificationRequest) GetParent() *ModbusPDU {
	return m.ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *ModbusPDUReadDeviceIdentificationRequest) GetLevel() ModbusDeviceInformationLevel {
	return m.Level
}

func (m *ModbusPDUReadDeviceIdentificationRequest) GetObjectId() uint8 {
	return m.ObjectId
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUReadDeviceIdentificationRequest factory function for ModbusPDUReadDeviceIdentificationRequest
func NewModbusPDUReadDeviceIdentificationRequest(level ModbusDeviceInformationLevel, objectId uint8) *ModbusPDUReadDeviceIdentificationRequest {
	_result := &ModbusPDUReadDeviceIdentificationRequest{
		Level:     level,
		ObjectId:  objectId,
		ModbusPDU: NewModbusPDU(),
	}
	_result.Child = _result
	return _result
}

func CastModbusPDUReadDeviceIdentificationRequest(structType interface{}) *ModbusPDUReadDeviceIdentificationRequest {
	if casted, ok := structType.(ModbusPDUReadDeviceIdentificationRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUReadDeviceIdentificationRequest); ok {
		return casted
	}
	if casted, ok := structType.(ModbusPDU); ok {
		return CastModbusPDUReadDeviceIdentificationRequest(casted.Child)
	}
	if casted, ok := structType.(*ModbusPDU); ok {
		return CastModbusPDUReadDeviceIdentificationRequest(casted.Child)
	}
	return nil
}

func (m *ModbusPDUReadDeviceIdentificationRequest) GetTypeName() string {
	return "ModbusPDUReadDeviceIdentificationRequest"
}

func (m *ModbusPDUReadDeviceIdentificationRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUReadDeviceIdentificationRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Const Field (meiType)
	lengthInBits += 8

	// Simple field (level)
	lengthInBits += 8

	// Simple field (objectId)
	lengthInBits += 8

	return lengthInBits
}

func (m *ModbusPDUReadDeviceIdentificationRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadDeviceIdentificationRequestParse(readBuffer utils.ReadBuffer, response bool) (*ModbusPDUReadDeviceIdentificationRequest, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUReadDeviceIdentificationRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Const Field (meiType)
	meiType, _meiTypeErr := readBuffer.ReadUint8("meiType", 8)
	if _meiTypeErr != nil {
		return nil, errors.Wrap(_meiTypeErr, "Error parsing 'meiType' field")
	}
	if meiType != ModbusPDUReadDeviceIdentificationRequest_MEITYPE {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", ModbusPDUReadDeviceIdentificationRequest_MEITYPE) + " but got " + fmt.Sprintf("%d", meiType))
	}

	// Simple Field (level)
	if pullErr := readBuffer.PullContext("level"); pullErr != nil {
		return nil, pullErr
	}
	_level, _levelErr := ModbusDeviceInformationLevelParse(readBuffer)
	if _levelErr != nil {
		return nil, errors.Wrap(_levelErr, "Error parsing 'level' field")
	}
	level := _level
	if closeErr := readBuffer.CloseContext("level"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (objectId)
	_objectId, _objectIdErr := readBuffer.ReadUint8("objectId", 8)
	if _objectIdErr != nil {
		return nil, errors.Wrap(_objectIdErr, "Error parsing 'objectId' field")
	}
	objectId := _objectId

	if closeErr := readBuffer.CloseContext("ModbusPDUReadDeviceIdentificationRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &ModbusPDUReadDeviceIdentificationRequest{
		Level:     level,
		ObjectId:  objectId,
		ModbusPDU: &ModbusPDU{},
	}
	_child.ModbusPDU.Child = _child
	return _child, nil
}

func (m *ModbusPDUReadDeviceIdentificationRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadDeviceIdentificationRequest"); pushErr != nil {
			return pushErr
		}

		// Const Field (meiType)
		_meiTypeErr := writeBuffer.WriteUint8("meiType", 8, 0x0E)
		if _meiTypeErr != nil {
			return errors.Wrap(_meiTypeErr, "Error serializing 'meiType' field")
		}

		// Simple Field (level)
		if pushErr := writeBuffer.PushContext("level"); pushErr != nil {
			return pushErr
		}
		_levelErr := m.Level.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("level"); popErr != nil {
			return popErr
		}
		if _levelErr != nil {
			return errors.Wrap(_levelErr, "Error serializing 'level' field")
		}

		// Simple Field (objectId)
		objectId := uint8(m.ObjectId)
		_objectIdErr := writeBuffer.WriteUint8("objectId", 8, (objectId))
		if _objectIdErr != nil {
			return errors.Wrap(_objectIdErr, "Error serializing 'objectId' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUReadDeviceIdentificationRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ModbusPDUReadDeviceIdentificationRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
