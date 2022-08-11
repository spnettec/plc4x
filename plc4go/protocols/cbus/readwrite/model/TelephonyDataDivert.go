/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   https://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// TelephonyDataDivert is the corresponding interface of TelephonyDataDivert
type TelephonyDataDivert interface {
	utils.LengthAware
	utils.Serializable
	TelephonyData
	// GetNumber returns Number (property field)
	GetNumber() string
}

// TelephonyDataDivertExactly can be used when we want exactly this type and not a type which fulfills TelephonyDataDivert.
// This is useful for switch cases.
type TelephonyDataDivertExactly interface {
	TelephonyDataDivert
	isTelephonyDataDivert() bool
}

// _TelephonyDataDivert is the data-structure of this message
type _TelephonyDataDivert struct {
	*_TelephonyData
	Number string
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataDivert) InitializeParent(parent TelephonyData, commandTypeContainer TelephonyCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_TelephonyDataDivert) GetParent() TelephonyData {
	return m._TelephonyData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyDataDivert) GetNumber() string {
	return m.Number
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTelephonyDataDivert factory function for _TelephonyDataDivert
func NewTelephonyDataDivert(number string, commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyDataDivert {
	_result := &_TelephonyDataDivert{
		Number:         number,
		_TelephonyData: NewTelephonyData(commandTypeContainer, argument),
	}
	_result._TelephonyData._TelephonyDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataDivert(structType interface{}) TelephonyDataDivert {
	if casted, ok := structType.(TelephonyDataDivert); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataDivert); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataDivert) GetTypeName() string {
	return "TelephonyDataDivert"
}

func (m *_TelephonyDataDivert) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_TelephonyDataDivert) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (number)
	lengthInBits += uint16(int32((int32(m.GetCommandTypeContainer().NumBytes()) - int32(int32(1)))) * int32(int32(8)))

	return lengthInBits
}

func (m *_TelephonyDataDivert) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TelephonyDataDivertParse(readBuffer utils.ReadBuffer, commandTypeContainer TelephonyCommandTypeContainer) (TelephonyDataDivert, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataDivert"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataDivert")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (number)
	_number, _numberErr := readBuffer.ReadString("number", uint32(((commandTypeContainer.NumBytes())-(1))*(8)))
	if _numberErr != nil {
		return nil, errors.Wrap(_numberErr, "Error parsing 'number' field of TelephonyDataDivert")
	}
	number := _number

	if closeErr := readBuffer.CloseContext("TelephonyDataDivert"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataDivert")
	}

	// Create a partially initialized instance
	_child := &_TelephonyDataDivert{
		_TelephonyData: &_TelephonyData{},
		Number:         number,
	}
	_child._TelephonyData._TelephonyDataChildRequirements = _child
	return _child, nil
}

func (m *_TelephonyDataDivert) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataDivert"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataDivert")
		}

		// Simple Field (number)
		number := string(m.GetNumber())
		_numberErr := writeBuffer.WriteString("number", uint32(((m.GetCommandTypeContainer().NumBytes())-(1))*(8)), "UTF-8", (number))
		if _numberErr != nil {
			return errors.Wrap(_numberErr, "Error serializing 'number' field")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataDivert"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataDivert")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_TelephonyDataDivert) isTelephonyDataDivert() bool {
	return true
}

func (m *_TelephonyDataDivert) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
