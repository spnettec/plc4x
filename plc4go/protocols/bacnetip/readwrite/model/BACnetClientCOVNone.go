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

// BACnetClientCOVNone is the data-structure of this message
type BACnetClientCOVNone struct {
	*BACnetClientCOV
	DefaultIncrement *BACnetApplicationTagNull
}

// IBACnetClientCOVNone is the corresponding interface of BACnetClientCOVNone
type IBACnetClientCOVNone interface {
	IBACnetClientCOV
	// GetDefaultIncrement returns DefaultIncrement (property field)
	GetDefaultIncrement() *BACnetApplicationTagNull
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

func (m *BACnetClientCOVNone) InitializeParent(parent *BACnetClientCOV, peekedTagHeader *BACnetTagHeader) {
	m.BACnetClientCOV.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetClientCOVNone) GetParent() *BACnetClientCOV {
	return m.BACnetClientCOV
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetClientCOVNone) GetDefaultIncrement() *BACnetApplicationTagNull {
	return m.DefaultIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetClientCOVNone factory function for BACnetClientCOVNone
func NewBACnetClientCOVNone(defaultIncrement *BACnetApplicationTagNull, peekedTagHeader *BACnetTagHeader) *BACnetClientCOVNone {
	_result := &BACnetClientCOVNone{
		DefaultIncrement: defaultIncrement,
		BACnetClientCOV:  NewBACnetClientCOV(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetClientCOVNone(structType interface{}) *BACnetClientCOVNone {
	if casted, ok := structType.(BACnetClientCOVNone); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetClientCOVNone); ok {
		return casted
	}
	if casted, ok := structType.(BACnetClientCOV); ok {
		return CastBACnetClientCOVNone(casted.Child)
	}
	if casted, ok := structType.(*BACnetClientCOV); ok {
		return CastBACnetClientCOVNone(casted.Child)
	}
	return nil
}

func (m *BACnetClientCOVNone) GetTypeName() string {
	return "BACnetClientCOVNone"
}

func (m *BACnetClientCOVNone) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetClientCOVNone) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (defaultIncrement)
	lengthInBits += m.DefaultIncrement.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetClientCOVNone) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetClientCOVNoneParse(readBuffer utils.ReadBuffer) (*BACnetClientCOVNone, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetClientCOVNone"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (defaultIncrement)
	if pullErr := readBuffer.PullContext("defaultIncrement"); pullErr != nil {
		return nil, pullErr
	}
	_defaultIncrement, _defaultIncrementErr := BACnetApplicationTagParse(readBuffer)
	if _defaultIncrementErr != nil {
		return nil, errors.Wrap(_defaultIncrementErr, "Error parsing 'defaultIncrement' field")
	}
	defaultIncrement := CastBACnetApplicationTagNull(_defaultIncrement)
	if closeErr := readBuffer.CloseContext("defaultIncrement"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetClientCOVNone"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetClientCOVNone{
		DefaultIncrement: CastBACnetApplicationTagNull(defaultIncrement),
		BACnetClientCOV:  &BACnetClientCOV{},
	}
	_child.BACnetClientCOV.Child = _child
	return _child, nil
}

func (m *BACnetClientCOVNone) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetClientCOVNone"); pushErr != nil {
			return pushErr
		}

		// Simple Field (defaultIncrement)
		if pushErr := writeBuffer.PushContext("defaultIncrement"); pushErr != nil {
			return pushErr
		}
		_defaultIncrementErr := m.DefaultIncrement.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("defaultIncrement"); popErr != nil {
			return popErr
		}
		if _defaultIncrementErr != nil {
			return errors.Wrap(_defaultIncrementErr, "Error serializing 'defaultIncrement' field")
		}

		if popErr := writeBuffer.PopContext("BACnetClientCOVNone"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetClientCOVNone) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
