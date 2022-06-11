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

// BACnetHostAddress is the data-structure of this message
type BACnetHostAddress struct {
	PeekedTagHeader *BACnetTagHeader
	Child           IBACnetHostAddressChild
}

// IBACnetHostAddress is the corresponding interface of BACnetHostAddress
type IBACnetHostAddress interface {
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

type IBACnetHostAddressParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetHostAddress, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetHostAddressChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetHostAddress, peekedTagHeader *BACnetTagHeader)
	GetParent() *BACnetHostAddress

	GetTypeName() string
	IBACnetHostAddress
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetHostAddress) GetPeekedTagHeader() *BACnetTagHeader {
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

func (m *BACnetHostAddress) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetHostAddress factory function for BACnetHostAddress
func NewBACnetHostAddress(peekedTagHeader *BACnetTagHeader) *BACnetHostAddress {
	return &BACnetHostAddress{PeekedTagHeader: peekedTagHeader}
}

func CastBACnetHostAddress(structType interface{}) *BACnetHostAddress {
	if casted, ok := structType.(BACnetHostAddress); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetHostAddress); ok {
		return casted
	}
	if casted, ok := structType.(IBACnetHostAddressChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *BACnetHostAddress) GetTypeName() string {
	return "BACnetHostAddress"
}

func (m *BACnetHostAddress) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetHostAddress) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetHostAddress) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetHostAddress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetHostAddressParse(readBuffer utils.ReadBuffer) (*BACnetHostAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetHostAddressChild interface {
		InitializeParent(*BACnetHostAddress, *BACnetTagHeader)
		GetParent() *BACnetHostAddress
	}
	var _child BACnetHostAddressChild
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetHostAddressNull
		_child, typeSwitchError = BACnetHostAddressNullParse(readBuffer)
	case peekedTagNumber == uint8(1): // BACnetHostAddressIpAddress
		_child, typeSwitchError = BACnetHostAddressIpAddressParse(readBuffer)
	case peekedTagNumber == uint8(2): // BACnetHostAddressName
		_child, typeSwitchError = BACnetHostAddressNameParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetHostAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostAddress")
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), peekedTagHeader)
	return _child.GetParent(), nil
}

func (m *BACnetHostAddress) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetHostAddress) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetHostAddress, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetHostAddress"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetHostAddress")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetHostAddress"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetHostAddress")
	}
	return nil
}

func (m *BACnetHostAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
