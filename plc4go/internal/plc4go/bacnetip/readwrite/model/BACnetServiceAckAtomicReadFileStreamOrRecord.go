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
type BACnetServiceAckAtomicReadFileStreamOrRecord struct {
	PeekedTagHeader *BACnetTagHeader
	OpeningTag      *BACnetOpeningTag
	ClosingTag      *BACnetClosingTag
	Child           IBACnetServiceAckAtomicReadFileStreamOrRecordChild
}

// The corresponding interface
type IBACnetServiceAckAtomicReadFileStreamOrRecord interface {
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() *BACnetTagHeader
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IBACnetServiceAckAtomicReadFileStreamOrRecordParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetServiceAckAtomicReadFileStreamOrRecord, serializeChildFunction func() error) error
	GetTypeName() string
}

type IBACnetServiceAckAtomicReadFileStreamOrRecordChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *BACnetServiceAckAtomicReadFileStreamOrRecord, peekedTagHeader *BACnetTagHeader, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag)
	GetParent() *BACnetServiceAckAtomicReadFileStreamOrRecord

	GetTypeName() string
	IBACnetServiceAckAtomicReadFileStreamOrRecord
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////
func (m *BACnetServiceAckAtomicReadFileStreamOrRecord) GetPeekedTagHeader() *BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *BACnetServiceAckAtomicReadFileStreamOrRecord) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetServiceAckAtomicReadFileStreamOrRecord) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////
func (m *BACnetServiceAckAtomicReadFileStreamOrRecord) GetPeekedTagNumber() uint8 {
	return m.GetPeekedTagHeader().GetActualTagNumber()
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckAtomicReadFileStreamOrRecord factory function for BACnetServiceAckAtomicReadFileStreamOrRecord
func NewBACnetServiceAckAtomicReadFileStreamOrRecord(peekedTagHeader *BACnetTagHeader, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) *BACnetServiceAckAtomicReadFileStreamOrRecord {
	return &BACnetServiceAckAtomicReadFileStreamOrRecord{PeekedTagHeader: peekedTagHeader, OpeningTag: openingTag, ClosingTag: closingTag}
}

func CastBACnetServiceAckAtomicReadFileStreamOrRecord(structType interface{}) *BACnetServiceAckAtomicReadFileStreamOrRecord {
	if casted, ok := structType.(BACnetServiceAckAtomicReadFileStreamOrRecord); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetServiceAckAtomicReadFileStreamOrRecord); ok {
		return casted
	}
	if casted, ok := structType.(IBACnetServiceAckAtomicReadFileStreamOrRecordChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *BACnetServiceAckAtomicReadFileStreamOrRecord) GetTypeName() string {
	return "BACnetServiceAckAtomicReadFileStreamOrRecord"
}

func (m *BACnetServiceAckAtomicReadFileStreamOrRecord) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetServiceAckAtomicReadFileStreamOrRecord) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *BACnetServiceAckAtomicReadFileStreamOrRecord) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetServiceAckAtomicReadFileStreamOrRecord) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckAtomicReadFileStreamOrRecordParse(readBuffer utils.ReadBuffer) (*BACnetServiceAckAtomicReadFileStreamOrRecord, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicReadFileStreamOrRecord"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = readBuffer.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, pullErr
	}
	peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
	readBuffer.Reset(currentPos)

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetContextTagParse(readBuffer, uint8(peekedTagHeader.GetActualTagNumber()), BACnetDataType(BACnetDataType_OPENING_TAG))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetServiceAckAtomicReadFileStreamOrRecordChild interface {
		InitializeParent(*BACnetServiceAckAtomicReadFileStreamOrRecord, *BACnetTagHeader, *BACnetOpeningTag, *BACnetClosingTag)
		GetParent() *BACnetServiceAckAtomicReadFileStreamOrRecord
	}
	var _child BACnetServiceAckAtomicReadFileStreamOrRecordChild
	var typeSwitchError error
	switch {
	case peekedTagNumber == 0x0: // BACnetServiceAckAtomicReadFileStream
		_child, typeSwitchError = BACnetServiceAckAtomicReadFileStreamParse(readBuffer)
	case peekedTagNumber == 0x1: // BACnetServiceAckAtomicReadFileRecord
		_child, typeSwitchError = BACnetServiceAckAtomicReadFileRecordParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetContextTagParse(readBuffer, uint8(peekedTagHeader.GetActualTagNumber()), BACnetDataType(BACnetDataType_CLOSING_TAG))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicReadFileStreamOrRecord"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), peekedTagHeader, openingTag, closingTag)
	return _child.GetParent(), nil
}

func (m *BACnetServiceAckAtomicReadFileStreamOrRecord) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *BACnetServiceAckAtomicReadFileStreamOrRecord) SerializeParent(writeBuffer utils.WriteBuffer, child IBACnetServiceAckAtomicReadFileStreamOrRecord, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicReadFileStreamOrRecord"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicReadFileStreamOrRecord"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetServiceAckAtomicReadFileStreamOrRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
