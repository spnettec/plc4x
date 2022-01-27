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
type BACnetErrorGetEventInformation struct {
	*BACnetError
}

// The corresponding interface
type IBACnetErrorGetEventInformation interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetErrorGetEventInformation) ServiceChoice() uint8 {
	return 0x1D
}

func (m *BACnetErrorGetEventInformation) InitializeParent(parent *BACnetError, errorClass *BACnetApplicationTagEnumerated, errorCode *BACnetApplicationTagEnumerated) {
	m.BACnetError.ErrorClass = errorClass
	m.BACnetError.ErrorCode = errorCode
}

func NewBACnetErrorGetEventInformation(errorClass *BACnetApplicationTagEnumerated, errorCode *BACnetApplicationTagEnumerated) *BACnetError {
	child := &BACnetErrorGetEventInformation{
		BACnetError: NewBACnetError(errorClass, errorCode),
	}
	child.Child = child
	return child.BACnetError
}

func CastBACnetErrorGetEventInformation(structType interface{}) *BACnetErrorGetEventInformation {
	castFunc := func(typ interface{}) *BACnetErrorGetEventInformation {
		if casted, ok := typ.(BACnetErrorGetEventInformation); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetErrorGetEventInformation); ok {
			return casted
		}
		if casted, ok := typ.(BACnetError); ok {
			return CastBACnetErrorGetEventInformation(casted.Child)
		}
		if casted, ok := typ.(*BACnetError); ok {
			return CastBACnetErrorGetEventInformation(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetErrorGetEventInformation) GetTypeName() string {
	return "BACnetErrorGetEventInformation"
}

func (m *BACnetErrorGetEventInformation) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *BACnetErrorGetEventInformation) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	return lengthInBits
}

func (m *BACnetErrorGetEventInformation) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func BACnetErrorGetEventInformationParse(readBuffer utils.ReadBuffer) (*BACnetError, error) {
	if pullErr := readBuffer.PullContext("BACnetErrorGetEventInformation"); pullErr != nil {
		return nil, pullErr
	}

	if closeErr := readBuffer.CloseContext("BACnetErrorGetEventInformation"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetErrorGetEventInformation{
		BACnetError: &BACnetError{},
	}
	_child.BACnetError.Child = _child
	return _child.BACnetError, nil
}

func (m *BACnetErrorGetEventInformation) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetErrorGetEventInformation"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("BACnetErrorGetEventInformation"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetErrorGetEventInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
