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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// S7Address is the corresponding interface of S7Address
type S7Address interface {
	utils.LengthAware
	utils.Serializable
	// GetAddressType returns AddressType (discriminator field)
	GetAddressType() uint8
}

// S7AddressExactly can be used when we want exactly this type and not a type which fulfills S7Address.
// This is useful for switch cases.
type S7AddressExactly interface {
	S7Address
	isS7Address() bool
}

// _S7Address is the data-structure of this message
type _S7Address struct {
	_S7AddressChildRequirements
}

type _S7AddressChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetAddressType() uint8
}


type S7AddressParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child S7Address, serializeChildFunction func() error) error
	GetTypeName() string
}

type S7AddressChild interface {
	utils.Serializable
InitializeParent(parent S7Address )
	GetParent() *S7Address

	GetTypeName() string
	S7Address
}


// NewS7Address factory function for _S7Address
func NewS7Address( ) *_S7Address {
return &_S7Address{ }
}

// Deprecated: use the interface for direct cast
func CastS7Address(structType interface{}) S7Address {
    if casted, ok := structType.(S7Address); ok {
		return casted
	}
	if casted, ok := structType.(*S7Address); ok {
		return *casted
	}
	return nil
}

func (m *_S7Address) GetTypeName() string {
	return "S7Address"
}



func (m *_S7Address) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (addressType)
	lengthInBits += 8;

	return lengthInBits
}

func (m *_S7Address) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func S7AddressParse(theBytes []byte) (S7Address, error) {
	return S7AddressParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func S7AddressParseWithBuffer(readBuffer utils.ReadBuffer) (S7Address, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("S7Address"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7Address")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (addressType) (Used as input to a switch field)
	addressType, _addressTypeErr := readBuffer.ReadUint8("addressType", 8)
	if _addressTypeErr != nil {
		return nil, errors.Wrap(_addressTypeErr, "Error parsing 'addressType' field of S7Address")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type S7AddressChildSerializeRequirement interface {
		S7Address
		InitializeParent(S7Address )
		GetParent() S7Address
	}
	var _childTemp interface{}
	var _child S7AddressChildSerializeRequirement
	var typeSwitchError error
	switch {
case addressType == 0x10 : // S7AddressAny
		_childTemp, typeSwitchError = S7AddressAnyParseWithBuffer(readBuffer, )
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [addressType=%v]", addressType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of S7Address")
	}
	_child = _childTemp.(S7AddressChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("S7Address"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7Address")
	}

	// Finish initializing
_child.InitializeParent(_child )
	return _child, nil
}

func (pm *_S7Address) SerializeParent(writeBuffer utils.WriteBuffer, child S7Address, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("S7Address"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for S7Address")
	}

	// Discriminator Field (addressType) (Used as input to a switch field)
	addressType := uint8(child.GetAddressType())
	_addressTypeErr := writeBuffer.WriteUint8("addressType", 8, (addressType))

	if _addressTypeErr != nil {
		return errors.Wrap(_addressTypeErr, "Error serializing 'addressType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("S7Address"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for S7Address")
	}
	return nil
}


func (m *_S7Address) isS7Address() bool {
	return true
}

func (m *_S7Address) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



