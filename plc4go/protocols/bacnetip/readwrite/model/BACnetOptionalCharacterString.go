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

// BACnetOptionalCharacterString is the corresponding interface of BACnetOptionalCharacterString
type BACnetOptionalCharacterString interface {
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetOptionalCharacterStringExactly can be used when we want exactly this type and not a type which fulfills BACnetOptionalCharacterString.
// This is useful for switch cases.
type BACnetOptionalCharacterStringExactly interface {
	BACnetOptionalCharacterString
	isBACnetOptionalCharacterString() bool
}

// _BACnetOptionalCharacterString is the data-structure of this message
type _BACnetOptionalCharacterString struct {
	_BACnetOptionalCharacterStringChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetOptionalCharacterStringChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}

type BACnetOptionalCharacterStringParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child BACnetOptionalCharacterString, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetOptionalCharacterStringChild interface {
	utils.Serializable
	InitializeParent(parent BACnetOptionalCharacterString, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetOptionalCharacterString

	GetTypeName() string
	BACnetOptionalCharacterString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalCharacterString) GetPeekedTagHeader() BACnetTagHeader {
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

func (m *_BACnetOptionalCharacterString) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetOptionalCharacterString factory function for _BACnetOptionalCharacterString
func NewBACnetOptionalCharacterString(peekedTagHeader BACnetTagHeader) *_BACnetOptionalCharacterString {
	return &_BACnetOptionalCharacterString{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetOptionalCharacterString(structType interface{}) BACnetOptionalCharacterString {
	if casted, ok := structType.(BACnetOptionalCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalCharacterString) GetTypeName() string {
	return "BACnetOptionalCharacterString"
}

func (m *_BACnetOptionalCharacterString) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetOptionalCharacterString) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetOptionalCharacterStringParse(readBuffer utils.ReadBuffer) (BACnetOptionalCharacterString, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalCharacterString")
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
	type BACnetOptionalCharacterStringChildSerializeRequirement interface {
		BACnetOptionalCharacterString
		InitializeParent(BACnetOptionalCharacterString, BACnetTagHeader)
		GetParent() BACnetOptionalCharacterString
	}
	var _childTemp interface{}
	var _child BACnetOptionalCharacterStringChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetOptionalCharacterStringNull
		_childTemp, typeSwitchError = BACnetOptionalCharacterStringNullParse(readBuffer)
	case true: // BACnetOptionalCharacterStringValue
		_childTemp, typeSwitchError = BACnetOptionalCharacterStringValueParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}
	_child = _childTemp.(BACnetOptionalCharacterStringChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetOptionalCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalCharacterString")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetOptionalCharacterString) SerializeParent(writeBuffer utils.WriteBuffer, child BACnetOptionalCharacterString, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetOptionalCharacterString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetOptionalCharacterString")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetOptionalCharacterString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetOptionalCharacterString")
	}
	return nil
}

func (m *_BACnetOptionalCharacterString) isBACnetOptionalCharacterString() bool {
	return true
}

func (m *_BACnetOptionalCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
