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
)

// Code generated by code-generation. DO NOT EDIT.

// ParameterValueBaudRateSelector is the corresponding interface of ParameterValueBaudRateSelector
type ParameterValueBaudRateSelector interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ParameterValue
	// GetValue returns Value (property field)
	GetValue() BaudRateSelector
	// GetData returns Data (property field)
	GetData() []byte
}

// ParameterValueBaudRateSelectorExactly can be used when we want exactly this type and not a type which fulfills ParameterValueBaudRateSelector.
// This is useful for switch cases.
type ParameterValueBaudRateSelectorExactly interface {
	ParameterValueBaudRateSelector
	isParameterValueBaudRateSelector() bool
}

// _ParameterValueBaudRateSelector is the data-structure of this message
type _ParameterValueBaudRateSelector struct {
	*_ParameterValue
	Value BaudRateSelector
	Data  []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ParameterValueBaudRateSelector) GetParameterType() ParameterType {
	return ParameterType_BAUD_RATE_SELECTOR
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ParameterValueBaudRateSelector) InitializeParent(parent ParameterValue) {}

func (m *_ParameterValueBaudRateSelector) GetParent() ParameterValue {
	return m._ParameterValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ParameterValueBaudRateSelector) GetValue() BaudRateSelector {
	return m.Value
}

func (m *_ParameterValueBaudRateSelector) GetData() []byte {
	return m.Data
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewParameterValueBaudRateSelector factory function for _ParameterValueBaudRateSelector
func NewParameterValueBaudRateSelector(value BaudRateSelector, data []byte, numBytes uint8) *_ParameterValueBaudRateSelector {
	_result := &_ParameterValueBaudRateSelector{
		Value:           value,
		Data:            data,
		_ParameterValue: NewParameterValue(numBytes),
	}
	_result._ParameterValue._ParameterValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastParameterValueBaudRateSelector(structType any) ParameterValueBaudRateSelector {
	if casted, ok := structType.(ParameterValueBaudRateSelector); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValueBaudRateSelector); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValueBaudRateSelector) GetTypeName() string {
	return "ParameterValueBaudRateSelector"
}

func (m *_ParameterValueBaudRateSelector) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ParameterValueBaudRateSelector) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ParameterValueBaudRateSelectorParse(theBytes []byte, parameterType ParameterType, numBytes uint8) (ParameterValueBaudRateSelector, error) {
	return ParameterValueBaudRateSelectorParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), parameterType, numBytes)
}

func ParameterValueBaudRateSelectorParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, parameterType ParameterType, numBytes uint8) (ParameterValueBaudRateSelector, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ParameterValueBaudRateSelector"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValueBaudRateSelector")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((numBytes) >= (1))) {
		return nil, errors.WithStack(utils.ParseValidationError{"BaudRateSelector has exactly one byte"})
	}

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
	_value, _valueErr := BaudRateSelectorParseWithBuffer(ctx, readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of ParameterValueBaudRateSelector")
	}
	value := _value
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(uint16(numBytes) - uint16(uint16(1)))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of ParameterValueBaudRateSelector")
	}

	if closeErr := readBuffer.CloseContext("ParameterValueBaudRateSelector"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValueBaudRateSelector")
	}

	// Create a partially initialized instance
	_child := &_ParameterValueBaudRateSelector{
		_ParameterValue: &_ParameterValue{
			NumBytes: numBytes,
		},
		Value: value,
		Data:  data,
	}
	_child._ParameterValue._ParameterValueChildRequirements = _child
	return _child, nil
}

func (m *_ParameterValueBaudRateSelector) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ParameterValueBaudRateSelector) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ParameterValueBaudRateSelector"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ParameterValueBaudRateSelector")
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

		if popErr := writeBuffer.PopContext("ParameterValueBaudRateSelector"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ParameterValueBaudRateSelector")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ParameterValueBaudRateSelector) isParameterValueBaudRateSelector() bool {
	return true
}

func (m *_ParameterValueBaudRateSelector) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
