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

// S7ParameterModeTransition is the data-structure of this message
type S7ParameterModeTransition struct {
	*S7Parameter
	Method           uint8
	CpuFunctionType  uint8
	CpuFunctionGroup uint8
	CurrentMode      uint8
	SequenceNumber   uint8
}

// IS7ParameterModeTransition is the corresponding interface of S7ParameterModeTransition
type IS7ParameterModeTransition interface {
	IS7Parameter
	// GetMethod returns Method (property field)
	GetMethod() uint8
	// GetCpuFunctionType returns CpuFunctionType (property field)
	GetCpuFunctionType() uint8
	// GetCpuFunctionGroup returns CpuFunctionGroup (property field)
	GetCpuFunctionGroup() uint8
	// GetCurrentMode returns CurrentMode (property field)
	GetCurrentMode() uint8
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() uint8
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

func (m *S7ParameterModeTransition) GetParameterType() uint8 {
	return 0x01
}

func (m *S7ParameterModeTransition) GetMessageType() uint8 {
	return 0x07
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *S7ParameterModeTransition) InitializeParent(parent *S7Parameter) {}

func (m *S7ParameterModeTransition) GetParent() *S7Parameter {
	return m.S7Parameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *S7ParameterModeTransition) GetMethod() uint8 {
	return m.Method
}

func (m *S7ParameterModeTransition) GetCpuFunctionType() uint8 {
	return m.CpuFunctionType
}

func (m *S7ParameterModeTransition) GetCpuFunctionGroup() uint8 {
	return m.CpuFunctionGroup
}

func (m *S7ParameterModeTransition) GetCurrentMode() uint8 {
	return m.CurrentMode
}

func (m *S7ParameterModeTransition) GetSequenceNumber() uint8 {
	return m.SequenceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewS7ParameterModeTransition factory function for S7ParameterModeTransition
func NewS7ParameterModeTransition(method uint8, cpuFunctionType uint8, cpuFunctionGroup uint8, currentMode uint8, sequenceNumber uint8) *S7ParameterModeTransition {
	_result := &S7ParameterModeTransition{
		Method:           method,
		CpuFunctionType:  cpuFunctionType,
		CpuFunctionGroup: cpuFunctionGroup,
		CurrentMode:      currentMode,
		SequenceNumber:   sequenceNumber,
		S7Parameter:      NewS7Parameter(),
	}
	_result.Child = _result
	return _result
}

func CastS7ParameterModeTransition(structType interface{}) *S7ParameterModeTransition {
	if casted, ok := structType.(S7ParameterModeTransition); ok {
		return &casted
	}
	if casted, ok := structType.(*S7ParameterModeTransition); ok {
		return casted
	}
	if casted, ok := structType.(S7Parameter); ok {
		return CastS7ParameterModeTransition(casted.Child)
	}
	if casted, ok := structType.(*S7Parameter); ok {
		return CastS7ParameterModeTransition(casted.Child)
	}
	return nil
}

func (m *S7ParameterModeTransition) GetTypeName() string {
	return "S7ParameterModeTransition"
}

func (m *S7ParameterModeTransition) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *S7ParameterModeTransition) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 16

	// Implicit Field (itemLength)
	lengthInBits += 8

	// Simple field (method)
	lengthInBits += 8

	// Simple field (cpuFunctionType)
	lengthInBits += 4

	// Simple field (cpuFunctionGroup)
	lengthInBits += 4

	// Simple field (currentMode)
	lengthInBits += 8

	// Simple field (sequenceNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *S7ParameterModeTransition) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7ParameterModeTransitionParse(readBuffer utils.ReadBuffer, messageType uint8) (*S7ParameterModeTransition, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7ParameterModeTransition"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint16("reserved", 16)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint16(0x0010) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint16(0x0010),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Implicit Field (itemLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	itemLength, _itemLengthErr := readBuffer.ReadUint8("itemLength", 8)
	_ = itemLength
	if _itemLengthErr != nil {
		return nil, errors.Wrap(_itemLengthErr, "Error parsing 'itemLength' field")
	}

	// Simple Field (method)
	_method, _methodErr := readBuffer.ReadUint8("method", 8)
	if _methodErr != nil {
		return nil, errors.Wrap(_methodErr, "Error parsing 'method' field")
	}
	method := _method

	// Simple Field (cpuFunctionType)
	_cpuFunctionType, _cpuFunctionTypeErr := readBuffer.ReadUint8("cpuFunctionType", 4)
	if _cpuFunctionTypeErr != nil {
		return nil, errors.Wrap(_cpuFunctionTypeErr, "Error parsing 'cpuFunctionType' field")
	}
	cpuFunctionType := _cpuFunctionType

	// Simple Field (cpuFunctionGroup)
	_cpuFunctionGroup, _cpuFunctionGroupErr := readBuffer.ReadUint8("cpuFunctionGroup", 4)
	if _cpuFunctionGroupErr != nil {
		return nil, errors.Wrap(_cpuFunctionGroupErr, "Error parsing 'cpuFunctionGroup' field")
	}
	cpuFunctionGroup := _cpuFunctionGroup

	// Simple Field (currentMode)
	_currentMode, _currentModeErr := readBuffer.ReadUint8("currentMode", 8)
	if _currentModeErr != nil {
		return nil, errors.Wrap(_currentModeErr, "Error parsing 'currentMode' field")
	}
	currentMode := _currentMode

	// Simple Field (sequenceNumber)
	_sequenceNumber, _sequenceNumberErr := readBuffer.ReadUint8("sequenceNumber", 8)
	if _sequenceNumberErr != nil {
		return nil, errors.Wrap(_sequenceNumberErr, "Error parsing 'sequenceNumber' field")
	}
	sequenceNumber := _sequenceNumber

	if closeErr := readBuffer.CloseContext("S7ParameterModeTransition"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &S7ParameterModeTransition{
		Method:           method,
		CpuFunctionType:  cpuFunctionType,
		CpuFunctionGroup: cpuFunctionGroup,
		CurrentMode:      currentMode,
		SequenceNumber:   sequenceNumber,
		S7Parameter:      &S7Parameter{},
	}
	_child.S7Parameter.Child = _child
	return _child, nil
}

func (m *S7ParameterModeTransition) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7ParameterModeTransition"); pushErr != nil {
			return pushErr
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint16("reserved", 16, uint16(0x0010))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Implicit Field (itemLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		itemLength := uint8(uint8(uint8(m.GetLengthInBytes())) - uint8(uint8(2)))
		_itemLengthErr := writeBuffer.WriteUint8("itemLength", 8, (itemLength))
		if _itemLengthErr != nil {
			return errors.Wrap(_itemLengthErr, "Error serializing 'itemLength' field")
		}

		// Simple Field (method)
		method := uint8(m.Method)
		_methodErr := writeBuffer.WriteUint8("method", 8, (method))
		if _methodErr != nil {
			return errors.Wrap(_methodErr, "Error serializing 'method' field")
		}

		// Simple Field (cpuFunctionType)
		cpuFunctionType := uint8(m.CpuFunctionType)
		_cpuFunctionTypeErr := writeBuffer.WriteUint8("cpuFunctionType", 4, (cpuFunctionType))
		if _cpuFunctionTypeErr != nil {
			return errors.Wrap(_cpuFunctionTypeErr, "Error serializing 'cpuFunctionType' field")
		}

		// Simple Field (cpuFunctionGroup)
		cpuFunctionGroup := uint8(m.CpuFunctionGroup)
		_cpuFunctionGroupErr := writeBuffer.WriteUint8("cpuFunctionGroup", 4, (cpuFunctionGroup))
		if _cpuFunctionGroupErr != nil {
			return errors.Wrap(_cpuFunctionGroupErr, "Error serializing 'cpuFunctionGroup' field")
		}

		// Simple Field (currentMode)
		currentMode := uint8(m.CurrentMode)
		_currentModeErr := writeBuffer.WriteUint8("currentMode", 8, (currentMode))
		if _currentModeErr != nil {
			return errors.Wrap(_currentModeErr, "Error serializing 'currentMode' field")
		}

		// Simple Field (sequenceNumber)
		sequenceNumber := uint8(m.SequenceNumber)
		_sequenceNumberErr := writeBuffer.WriteUint8("sequenceNumber", 8, (sequenceNumber))
		if _sequenceNumberErr != nil {
			return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
		}

		if popErr := writeBuffer.PopContext("S7ParameterModeTransition"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *S7ParameterModeTransition) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
