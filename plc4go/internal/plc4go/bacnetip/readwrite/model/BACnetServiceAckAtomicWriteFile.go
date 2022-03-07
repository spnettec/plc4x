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
type BACnetServiceAckAtomicWriteFile struct {
	*BACnetServiceAck
	FileStartPosition *BACnetContextTagSignedInteger
}

// The corresponding interface
type IBACnetServiceAckAtomicWriteFile interface {
	IBACnetServiceAck
	// GetFileStartPosition returns FileStartPosition (property field)
	GetFileStartPosition() *BACnetContextTagSignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetServiceAckAtomicWriteFile) ServiceChoice() uint8 {
	return 0x07
}

func (m *BACnetServiceAckAtomicWriteFile) GetServiceChoice() uint8 {
	return 0x07
}

func (m *BACnetServiceAckAtomicWriteFile) InitializeParent(parent *BACnetServiceAck) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetServiceAckAtomicWriteFile) GetFileStartPosition() *BACnetContextTagSignedInteger {
	return m.FileStartPosition
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetServiceAckAtomicWriteFile factory function for BACnetServiceAckAtomicWriteFile
func NewBACnetServiceAckAtomicWriteFile(fileStartPosition *BACnetContextTagSignedInteger) *BACnetServiceAck {
	child := &BACnetServiceAckAtomicWriteFile{
		FileStartPosition: fileStartPosition,
		BACnetServiceAck:  NewBACnetServiceAck(),
	}
	child.Child = child
	return child.BACnetServiceAck
}

func CastBACnetServiceAckAtomicWriteFile(structType interface{}) *BACnetServiceAckAtomicWriteFile {
	if casted, ok := structType.(BACnetServiceAckAtomicWriteFile); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetServiceAckAtomicWriteFile); ok {
		return casted
	}
	if casted, ok := structType.(BACnetServiceAck); ok {
		return CastBACnetServiceAckAtomicWriteFile(casted.Child)
	}
	if casted, ok := structType.(*BACnetServiceAck); ok {
		return CastBACnetServiceAckAtomicWriteFile(casted.Child)
	}
	return nil
}

func (m *BACnetServiceAckAtomicWriteFile) GetTypeName() string {
	return "BACnetServiceAckAtomicWriteFile"
}

func (m *BACnetServiceAckAtomicWriteFile) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetServiceAckAtomicWriteFile) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (fileStartPosition)
	lengthInBits += m.FileStartPosition.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetServiceAckAtomicWriteFile) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckAtomicWriteFileParse(readBuffer utils.ReadBuffer) (*BACnetServiceAck, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicWriteFile"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (fileStartPosition)
	if pullErr := readBuffer.PullContext("fileStartPosition"); pullErr != nil {
		return nil, pullErr
	}
	_fileStartPosition, _fileStartPositionErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_SIGNED_INTEGER))
	if _fileStartPositionErr != nil {
		return nil, errors.Wrap(_fileStartPositionErr, "Error parsing 'fileStartPosition' field")
	}
	fileStartPosition := CastBACnetContextTagSignedInteger(_fileStartPosition)
	if closeErr := readBuffer.CloseContext("fileStartPosition"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicWriteFile"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckAtomicWriteFile{
		FileStartPosition: CastBACnetContextTagSignedInteger(fileStartPosition),
		BACnetServiceAck:  &BACnetServiceAck{},
	}
	_child.BACnetServiceAck.Child = _child
	return _child.BACnetServiceAck, nil
}

func (m *BACnetServiceAckAtomicWriteFile) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicWriteFile"); pushErr != nil {
			return pushErr
		}

		// Simple Field (fileStartPosition)
		if pushErr := writeBuffer.PushContext("fileStartPosition"); pushErr != nil {
			return pushErr
		}
		_fileStartPositionErr := m.FileStartPosition.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("fileStartPosition"); popErr != nil {
			return popErr
		}
		if _fileStartPositionErr != nil {
			return errors.Wrap(_fileStartPositionErr, "Error serializing 'fileStartPosition' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicWriteFile"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckAtomicWriteFile) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
