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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// ParameterValueSerialNumber is the corresponding interface of ParameterValueSerialNumber
type ParameterValueSerialNumber interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ParameterValue
	// GetValue returns Value (property field)
	GetValue() SerialNumber
	// GetData returns Data (property field)
	GetData() []byte
}

// ParameterValueSerialNumberExactly can be used when we want exactly this type and not a type which fulfills ParameterValueSerialNumber.
// This is useful for switch cases.
type ParameterValueSerialNumberExactly interface {
	ParameterValueSerialNumber
	isParameterValueSerialNumber() bool
}

// _ParameterValueSerialNumber is the data-structure of this message
type _ParameterValueSerialNumber struct {
	*_ParameterValue
	Value SerialNumber
	Data  []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueSerialNumber) GetParameterType() ParameterType {
	return ParameterType_SERIAL_NUMBER
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueSerialNumber) InitializeParent(parent ParameterValue) {}

func (m *_ParameterValueSerialNumber) GetParent() ParameterValue {
	return m._ParameterValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueSerialNumber) GetValue() SerialNumber {
	return m.Value
}

func (m *_ParameterValueSerialNumber) GetData() []byte {
	return m.Data
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewParameterValueSerialNumber factory function for _ParameterValueSerialNumber
func NewParameterValueSerialNumber(value SerialNumber, data []byte, numBytes uint8) *_ParameterValueSerialNumber {
	_result := &_ParameterValueSerialNumber{
		Value:           value,
		Data:            data,
		_ParameterValue: NewParameterValue(numBytes),
	}
	_result._ParameterValue._ParameterValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastParameterValueSerialNumber(structType any) ParameterValueSerialNumber {
	if casted, ok := structType.(ParameterValueSerialNumber); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueSerialNumber); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueSerialNumber) GetTypeName() string {
	return "ParameterValueSerialNumber"
}

func (m *_ParameterValueSerialNumber) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += m.Value.GetLengthInBits(ctx)

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ParameterValueSerialNumber) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ParameterValueSerialNumberParse(ctx context.Context, theBytes []byte, parameterType ParameterType, numBytes uint8) (ParameterValueSerialNumber, error) {
	return ParameterValueSerialNumberParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), parameterType, numBytes)
}

func ParameterValueSerialNumberParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, parameterType ParameterType, numBytes uint8) (ParameterValueSerialNumber, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ParameterValueSerialNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueSerialNumber")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((numBytes) >= (4))) {
		return nil, errors.WithStack(utils.ParseValidationError{"SerialNumber has exactly four bytes"})
	}

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
	_value, _valueErr := SerialNumberParseWithBuffer(ctx, readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of ParameterValueSerialNumber")
	}
	value := _value.(SerialNumber)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(uint16(numBytes) - uint16(uint16(4)))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of ParameterValueSerialNumber")
	}

	if closeErr := readBuffer.CloseContext("ParameterValueSerialNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueSerialNumber")
	}

	// Create a partially initialized instance
	_child := &_ParameterValueSerialNumber{
		_ParameterValue: &_ParameterValue{
			NumBytes: numBytes,
		},
		Value: value,
		Data:  data,
	}
	_child._ParameterValue._ParameterValueChildRequirements = _child
	return _child, nil
}

func (m *_ParameterValueSerialNumber) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterValueSerialNumber) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueSerialNumber"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueSerialNumber")
		}

		// Simple Field (value)
		if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for value")
		}
		_valueErr := writeBuffer.WriteSerializable(ctx, m.GetValue())
		if popErr := writeBuffer.PopContext("value"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for value")
		}
		if _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		// Array Field (data)
		// Byte Array field (data)
		if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("ParameterValueSerialNumber"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueSerialNumber")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ParameterValueSerialNumber) isParameterValueSerialNumber() bool {
	return true
}

func (m *_ParameterValueSerialNumber) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
