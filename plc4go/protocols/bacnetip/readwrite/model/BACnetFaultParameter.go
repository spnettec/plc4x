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

// BACnetFaultParameter is the data-structure of this message
type BACnetFaultParameter struct {
	PeekedTagHeader *BACnetTagHeader
	Child           IBACnetFaultParameterChild
}

// IBACnetFaultParameter is the corresponding interface of BACnetFaultParameter
type IBACnetFaultParameter interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() *BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetFaultParameterParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetFaultParameter, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetFaultParameterChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetFaultParameter, peekedTagHeader *BACnetTagHeader)
	GetParent() *BACnetFaultParameter

	GetTypeName() string
	IBACnetFaultParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetFaultParameter) GetPeekedTagHeader() *BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetFaultParameter) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameter factory function for BACnetFaultParameter
func NewBACnetFaultParameter(peekedTagHeader *BACnetTagHeader) *BACnetFaultParameter {
	return &BACnetFaultParameter{PeekedTagHeader: peekedTagHeader}
}

func CastBACnetFaultParameter(structType interface{}) *BACnetFaultParameter {
	if casted, ok := structType.(BACnetFaultParameter); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetFaultParameter); ok {
		return casted
	}
	if casted, ok := structType.(IBACnetFaultParameterChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *BACnetFaultParameter) GetTypeName() string {
	return "BACnetFaultParameter"
}

func (m *BACnetFaultParameter) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetFaultParameter) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetFaultParameter) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetFaultParameter) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterParse(readBuffer utils.ReadBuffer) (*BACnetFaultParameter, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameter"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, pullErr
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetFaultParameterChild interface {
		InitializeParent(*BACnetFaultParameter, *BACnetTagHeader)
		GetParent() *BACnetFaultParameter
	}
	var _child BACnetFaultParameterChild
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetFaultParameterNone
		_child, typeSwitchError = BACnetFaultParameterNoneParse(readBuffer)
	case peekedTagNumber == uint8(1): // BACnetFaultParameterFaultCharacterString
		_child, typeSwitchError = BACnetFaultParameterFaultCharacterStringParse(readBuffer)
	case peekedTagNumber == uint8(2): // BACnetFaultParameterFaultExtended
		_child, typeSwitchError = BACnetFaultParameterFaultExtendedParse(readBuffer)
	case peekedTagNumber == uint8(3): // BACnetFaultParameterFaultLifeSafety
		_child, typeSwitchError = BACnetFaultParameterFaultLifeSafetyParse(readBuffer)
	case peekedTagNumber == uint8(4): // BACnetFaultParameterFaultState
		_child, typeSwitchError = BACnetFaultParameterFaultStateParse(readBuffer)
	case peekedTagNumber == uint8(5): // BACnetFaultParameterFaultStatusFlags
		_child, typeSwitchError = BACnetFaultParameterFaultStatusFlagsParse(readBuffer)
	case peekedTagNumber == uint8(6): // BACnetFaultParameterFaultOutOfRange
		_child, typeSwitchError = BACnetFaultParameterFaultOutOfRangeParse(readBuffer)
	case peekedTagNumber == uint8(7): // BACnetFaultParameterFaultListed
		_child, typeSwitchError = BACnetFaultParameterFaultListedParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameter"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), peekedTagHeader)
	return _child.GetParent(), nil
}

func (m *BACnetFaultParameter) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetFaultParameter) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetFaultParameter, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetFaultParameter"); pushErr != nil {
		return pushErr
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetFaultParameter"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetFaultParameter) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
