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

// BACnetFaultParameterFaultExtendedParametersEntry is the data-structure of this message
type BACnetFaultParameterFaultExtendedParametersEntry struct {
	PeekedTagHeader *BACnetTagHeader
	Child           IBACnetFaultParameterFaultExtendedParametersEntryChild
}

// IBACnetFaultParameterFaultExtendedParametersEntry is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntry
type IBACnetFaultParameterFaultExtendedParametersEntry interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() *BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetPeekedIsContextTag returns PeekedIsContextTag (virtual field)
	GetPeekedIsContextTag() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetFaultParameterFaultExtendedParametersEntryParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetFaultParameterFaultExtendedParametersEntry, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetFaultParameterFaultExtendedParametersEntryChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetFaultParameterFaultExtendedParametersEntry, peekedTagHeader *BACnetTagHeader)
	GetParent() *BACnetFaultParameterFaultExtendedParametersEntry

	GetTypeName() string
	IBACnetFaultParameterFaultExtendedParametersEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetFaultParameterFaultExtendedParametersEntry) GetPeekedTagHeader() *BACnetTagHeader {
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

func (m *BACnetFaultParameterFaultExtendedParametersEntry) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (m *BACnetFaultParameterFaultExtendedParametersEntry) GetPeekedIsContextTag() bool {
	return bool(bool((m.GetPeekedTagHeader().GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtendedParametersEntry factory function for BACnetFaultParameterFaultExtendedParametersEntry
func NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader *BACnetTagHeader) *BACnetFaultParameterFaultExtendedParametersEntry {
	return &BACnetFaultParameterFaultExtendedParametersEntry{PeekedTagHeader: peekedTagHeader}
}

func CastBACnetFaultParameterFaultExtendedParametersEntry(structType interface{}) *BACnetFaultParameterFaultExtendedParametersEntry {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntry); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntry); ok {
		return casted
	}
	if casted, ok := structType.(IBACnetFaultParameterFaultExtendedParametersEntryChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *BACnetFaultParameterFaultExtendedParametersEntry) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntry"
}

func (m *BACnetFaultParameterFaultExtendedParametersEntry) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetFaultParameterFaultExtendedParametersEntry) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetFaultParameterFaultExtendedParametersEntry) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetFaultParameterFaultExtendedParametersEntry) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultExtendedParametersEntryParse(readBuffer utils.ReadBuffer) (*BACnetFaultParameterFaultExtendedParametersEntry, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntry"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntry")
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

	// Virtual field
	_peekedIsContextTag := bool((peekedTagHeader.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))
	peekedIsContextTag := bool(_peekedIsContextTag)
	_ = peekedIsContextTag

	// Validation
	if !(bool(bool(!(peekedIsContextTag))) || bool(bool(bool(bool(peekedIsContextTag) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x6)))) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x7)))))) {
		return nil, errors.WithStack(utils.ParseValidationError{"unexpected opening or closing tag"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetFaultParameterFaultExtendedParametersEntryChild interface {
		InitializeParent(*BACnetFaultParameterFaultExtendedParametersEntry, *BACnetTagHeader)
		GetParent() *BACnetFaultParameterFaultExtendedParametersEntry
	}
	var _child BACnetFaultParameterFaultExtendedParametersEntryChild
	var typeSwitchError error
	switch {
	case peekedTagNumber == 0x0 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryNull
		_child, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryNullParse(readBuffer)
	case peekedTagNumber == 0x4 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryReal
		_child, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryRealParse(readBuffer)
	case peekedTagNumber == 0x2 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryUnsigned
		_child, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryUnsignedParse(readBuffer)
	case peekedTagNumber == 0x1 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryBoolean
		_child, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryBooleanParse(readBuffer)
	case peekedTagNumber == 0x3 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryInteger
		_child, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryIntegerParse(readBuffer)
	case peekedTagNumber == 0x5 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryDouble
		_child, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryDoubleParse(readBuffer)
	case peekedTagNumber == 0x6 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryOctetString
		_child, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryOctetStringParse(readBuffer)
	case peekedTagNumber == 0x7 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryCharacterString
		_child, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryCharacterStringParse(readBuffer)
	case peekedTagNumber == 0x8 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryBitString
		_child, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryBitStringParse(readBuffer)
	case peekedTagNumber == 0x9 && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryEnumerated
		_child, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryEnumeratedParse(readBuffer)
	case peekedTagNumber == 0xA && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryDate
		_child, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryDateParse(readBuffer)
	case peekedTagNumber == 0xB && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryTime
		_child, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryTimeParse(readBuffer)
	case peekedTagNumber == 0xC && peekedIsContextTag == bool(false): // BACnetFaultParameterFaultExtendedParametersEntryObjectidentifier
		_child, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryObjectidentifierParse(readBuffer)
	case peekedTagNumber == uint8(0) && peekedIsContextTag == bool(true): // BACnetFaultParameterFaultExtendedParametersEntryReference
		_child, typeSwitchError = BACnetFaultParameterFaultExtendedParametersEntryReferenceParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntry"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntry")
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), peekedTagHeader)
	return _child.GetParent(), nil
}

func (m *BACnetFaultParameterFaultExtendedParametersEntry) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetFaultParameterFaultExtendedParametersEntry) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetFaultParameterFaultExtendedParametersEntry, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntry"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntry")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}
	// Virtual field
	if _peekedIsContextTagErr := writeBuffer.WriteVirtual("peekedIsContextTag", m.GetPeekedIsContextTag()); _peekedIsContextTagErr != nil {
		return errors.Wrap(_peekedIsContextTagErr, "Error serializing 'peekedIsContextTag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntry"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntry")
	}
	return nil
}

func (m *BACnetFaultParameterFaultExtendedParametersEntry) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
