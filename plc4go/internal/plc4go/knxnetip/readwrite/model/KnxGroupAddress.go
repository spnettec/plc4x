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
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type KnxGroupAddress struct {
	Child IKnxGroupAddressChild
}

// The corresponding interface
type IKnxGroupAddress interface {
	// NumLevels returns NumLevels
	NumLevels() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IKnxGroupAddressParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IKnxGroupAddress, serializeChildFunction func() error) error
	GetTypeName() string
}

type IKnxGroupAddressChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *KnxGroupAddress)
	GetTypeName() string
	IKnxGroupAddress
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewKnxGroupAddress factory function for KnxGroupAddress
func NewKnxGroupAddress() *KnxGroupAddress {
	return &KnxGroupAddress{}
}

func CastKnxGroupAddress(structType interface{}) *KnxGroupAddress {
	castFunc := func(typ interface{}) *KnxGroupAddress {
		if casted, ok := typ.(KnxGroupAddress); ok {
			return &casted
		}
		if casted, ok := typ.(*KnxGroupAddress); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *KnxGroupAddress) GetTypeName() string {
	return "KnxGroupAddress"
}

func (m *KnxGroupAddress) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *KnxGroupAddress) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *KnxGroupAddress) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *KnxGroupAddress) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func KnxGroupAddressParse(readBuffer utils.ReadBuffer, numLevels uint8) (*KnxGroupAddress, error) {
	if pullErr := readBuffer.PullContext("KnxGroupAddress"); pullErr != nil {
		return nil, pullErr
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *KnxGroupAddress
	var typeSwitchError error
	switch {
	case numLevels == uint8(1): // KnxGroupAddressFreeLevel
		_parent, typeSwitchError = KnxGroupAddressFreeLevelParse(readBuffer, numLevels)
	case numLevels == uint8(2): // KnxGroupAddress2Level
		_parent, typeSwitchError = KnxGroupAddress2LevelParse(readBuffer, numLevels)
	case numLevels == uint8(3): // KnxGroupAddress3Level
		_parent, typeSwitchError = KnxGroupAddress3LevelParse(readBuffer, numLevels)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("KnxGroupAddress"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *KnxGroupAddress) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *KnxGroupAddress) SerializeParent(writeBuffer utils.WriteBuffer, child IKnxGroupAddress, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("KnxGroupAddress"); pushErr != nil {
		return pushErr
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("KnxGroupAddress"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *KnxGroupAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
