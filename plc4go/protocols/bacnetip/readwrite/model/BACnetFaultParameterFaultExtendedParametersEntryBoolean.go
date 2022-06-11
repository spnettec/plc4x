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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetFaultParameterFaultExtendedParametersEntryBoolean is the data-structure of this message
type BACnetFaultParameterFaultExtendedParametersEntryBoolean struct {
	*BACnetFaultParameterFaultExtendedParametersEntry
	BooleanValue *BACnetApplicationTagBoolean
}

// IBACnetFaultParameterFaultExtendedParametersEntryBoolean is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryBoolean
type IBACnetFaultParameterFaultExtendedParametersEntryBoolean interface {
	IBACnetFaultParameterFaultExtendedParametersEntry
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() *BACnetApplicationTagBoolean
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetFaultParameterFaultExtendedParametersEntryBoolean) InitializeParent(parent *BACnetFaultParameterFaultExtendedParametersEntry, peekedTagHeader *BACnetTagHeader) {
	m.BACnetFaultParameterFaultExtendedParametersEntry.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryBoolean) GetParent() *BACnetFaultParameterFaultExtendedParametersEntry {
	return m.BACnetFaultParameterFaultExtendedParametersEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetFaultParameterFaultExtendedParametersEntryBoolean) GetBooleanValue() *BACnetApplicationTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtendedParametersEntryBoolean factory function for BACnetFaultParameterFaultExtendedParametersEntryBoolean
func NewBACnetFaultParameterFaultExtendedParametersEntryBoolean(booleanValue *BACnetApplicationTagBoolean, peekedTagHeader *BACnetTagHeader) *BACnetFaultParameterFaultExtendedParametersEntryBoolean {
	_result := &BACnetFaultParameterFaultExtendedParametersEntryBoolean{
		BooleanValue: booleanValue,
		BACnetFaultParameterFaultExtendedParametersEntry: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetFaultParameterFaultExtendedParametersEntryBoolean(structType interface{}) *BACnetFaultParameterFaultExtendedParametersEntryBoolean {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryBoolean); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryBoolean); ok {
		return casted
	}
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntry); ok {
		return CastBACnetFaultParameterFaultExtendedParametersEntryBoolean(casted.Child)
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntry); ok {
		return CastBACnetFaultParameterFaultExtendedParametersEntryBoolean(casted.Child)
	}
	return nil
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryBoolean) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryBoolean"
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryBoolean) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryBoolean) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryBoolean) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultExtendedParametersEntryBooleanParse(readBuffer utils.ReadBuffer) (*BACnetFaultParameterFaultExtendedParametersEntryBoolean, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (booleanValue)
	if pullErr := readBuffer.PullContext("booleanValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for booleanValue")
	}
	_booleanValue, _booleanValueErr := BACnetApplicationTagParse(readBuffer)
	if _booleanValueErr != nil {
		return nil, errors.Wrap(_booleanValueErr, "Error parsing 'booleanValue' field")
	}
	booleanValue := CastBACnetApplicationTagBoolean(_booleanValue)
	if closeErr := readBuffer.CloseContext("booleanValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for booleanValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryBoolean")
	}

	// Create a partially initialized instance
	_child := &BACnetFaultParameterFaultExtendedParametersEntryBoolean{
		BooleanValue: CastBACnetApplicationTagBoolean(booleanValue),
		BACnetFaultParameterFaultExtendedParametersEntry: &BACnetFaultParameterFaultExtendedParametersEntry{},
	}
	_child.BACnetFaultParameterFaultExtendedParametersEntry.Child = _child
	return _child, nil
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryBoolean) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryBoolean")
		}

		// Simple Field (booleanValue)
		if pushErr := writeBuffer.PushContext("booleanValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for booleanValue")
		}
		_booleanValueErr := m.BooleanValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("booleanValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for booleanValue")
		}
		if _booleanValueErr != nil {
			return errors.Wrap(_booleanValueErr, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryBoolean")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetFaultParameterFaultExtendedParametersEntryBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
