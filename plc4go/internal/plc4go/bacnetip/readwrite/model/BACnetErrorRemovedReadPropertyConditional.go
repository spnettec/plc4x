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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetErrorRemovedReadPropertyConditional struct {
	*BACnetError
}

// The corresponding interface
type IBACnetErrorRemovedReadPropertyConditional interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetErrorRemovedReadPropertyConditional) ServiceChoice() uint8 {
	return 0x0D
}

func (m *BACnetErrorRemovedReadPropertyConditional) InitializeParent(parent *BACnetError, errorClass *BACnetApplicationTagEnumerated, errorCode *BACnetApplicationTagEnumerated) {
	m.ErrorClass = errorClass
	m.ErrorCode = errorCode
}

func NewBACnetErrorRemovedReadPropertyConditional(errorClass *BACnetApplicationTagEnumerated, errorCode *BACnetApplicationTagEnumerated) *BACnetError {
	child := &BACnetErrorRemovedReadPropertyConditional{
		BACnetError: NewBACnetError(errorClass, errorCode),
	}
	child.Child = child
	return child.BACnetError
}

func CastBACnetErrorRemovedReadPropertyConditional(structType interface{}) *BACnetErrorRemovedReadPropertyConditional {
	castFunc := func(typ interface{}) *BACnetErrorRemovedReadPropertyConditional {
		if casted, ok := typ.(BACnetErrorRemovedReadPropertyConditional); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetErrorRemovedReadPropertyConditional); ok {
			return casted
		}
		if casted, ok := typ.(BACnetError); ok {
			return CastBACnetErrorRemovedReadPropertyConditional(casted.Child)
		}
		if casted, ok := typ.(*BACnetError); ok {
			return CastBACnetErrorRemovedReadPropertyConditional(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetErrorRemovedReadPropertyConditional) GetTypeName() string {
	return "BACnetErrorRemovedReadPropertyConditional"
}

func (m *BACnetErrorRemovedReadPropertyConditional) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetErrorRemovedReadPropertyConditional) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetErrorRemovedReadPropertyConditional) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetErrorRemovedReadPropertyConditionalParse(readBuffer utils.ReadBuffer) (*BACnetError, error) {
	if pullErr := readBuffer.PullContext("BACnetErrorRemovedReadPropertyConditional"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetErrorRemovedReadPropertyConditional"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetErrorRemovedReadPropertyConditional{
		BACnetError: &BACnetError{},
	}
	_child.BACnetError.Child = _child
	return _child.BACnetError, nil
}

func (m *BACnetErrorRemovedReadPropertyConditional) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetErrorRemovedReadPropertyConditional"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetErrorRemovedReadPropertyConditional"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetErrorRemovedReadPropertyConditional) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
