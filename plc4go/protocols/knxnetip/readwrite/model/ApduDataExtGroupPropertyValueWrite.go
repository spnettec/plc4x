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

// ApduDataExtGroupPropertyValueWrite is the corresponding interface of ApduDataExtGroupPropertyValueWrite
type ApduDataExtGroupPropertyValueWrite interface {
	ApduDataExt
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _ApduDataExtGroupPropertyValueWrite is the data-structure of this message
type _ApduDataExtGroupPropertyValueWrite struct {
	*_ApduDataExt

	// Arguments.
	Length uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtGroupPropertyValueWrite) GetExtApciType() uint8 {
	return 0x2A
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtGroupPropertyValueWrite) InitializeParent(parent ApduDataExt) {}

func (m *_ApduDataExtGroupPropertyValueWrite) GetParent() ApduDataExt {
	return m._ApduDataExt
}

// NewApduDataExtGroupPropertyValueWrite factory function for _ApduDataExtGroupPropertyValueWrite
func NewApduDataExtGroupPropertyValueWrite(length uint8) *_ApduDataExtGroupPropertyValueWrite {
	_result := &_ApduDataExtGroupPropertyValueWrite{
		_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtGroupPropertyValueWrite(structType interface{}) ApduDataExtGroupPropertyValueWrite {
	if casted, ok := structType.(ApduDataExtGroupPropertyValueWrite); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtGroupPropertyValueWrite); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtGroupPropertyValueWrite) GetTypeName() string {
	return "ApduDataExtGroupPropertyValueWrite"
}

func (m *_ApduDataExtGroupPropertyValueWrite) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtGroupPropertyValueWrite) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_ApduDataExtGroupPropertyValueWrite) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtGroupPropertyValueWriteParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtGroupPropertyValueWrite, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtGroupPropertyValueWrite"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtGroupPropertyValueWrite")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtGroupPropertyValueWrite"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtGroupPropertyValueWrite")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtGroupPropertyValueWrite{
		_ApduDataExt: &_ApduDataExt{},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtGroupPropertyValueWrite) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtGroupPropertyValueWrite"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtGroupPropertyValueWrite")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtGroupPropertyValueWrite"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtGroupPropertyValueWrite")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_ApduDataExtGroupPropertyValueWrite) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
