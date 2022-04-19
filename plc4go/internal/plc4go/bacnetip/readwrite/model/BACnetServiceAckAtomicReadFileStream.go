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

// BACnetServiceAckAtomicReadFileStream is the data-structure of this message
type BACnetServiceAckAtomicReadFileStream struct {
	*BACnetServiceAckAtomicReadFileStreamOrRecord
	FileStartPosition *BACnetApplicationTagSignedInteger
	FileData          *BACnetApplicationTagOctetString
}

// IBACnetServiceAckAtomicReadFileStream is the corresponding interface of BACnetServiceAckAtomicReadFileStream
type IBACnetServiceAckAtomicReadFileStream interface {
	IBACnetServiceAckAtomicReadFileStreamOrRecord
	// GetFileStartPosition returns FileStartPosition (property field)
	GetFileStartPosition() *BACnetApplicationTagSignedInteger
	// GetFileData returns FileData (property field)
	GetFileData() *BACnetApplicationTagOctetString
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetServiceAckAtomicReadFileStream) InitializeParent(parent *BACnetServiceAckAtomicReadFileStreamOrRecord, peekedTagHeader *BACnetTagHeader, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetServiceAckAtomicReadFileStreamOrRecord.PeekedTagHeader = peekedTagHeader
	m.BACnetServiceAckAtomicReadFileStreamOrRecord.OpeningTag = openingTag
	m.BACnetServiceAckAtomicReadFileStreamOrRecord.ClosingTag = closingTag
}

func (m *BACnetServiceAckAtomicReadFileStream) GetParent() *BACnetServiceAckAtomicReadFileStreamOrRecord {
	return m.BACnetServiceAckAtomicReadFileStreamOrRecord
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetServiceAckAtomicReadFileStream) GetFileStartPosition() *BACnetApplicationTagSignedInteger {
	return m.FileStartPosition
}

func (m *BACnetServiceAckAtomicReadFileStream) GetFileData() *BACnetApplicationTagOctetString {
	return m.FileData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckAtomicReadFileStream factory function for BACnetServiceAckAtomicReadFileStream
func NewBACnetServiceAckAtomicReadFileStream(fileStartPosition *BACnetApplicationTagSignedInteger, fileData *BACnetApplicationTagOctetString, peekedTagHeader *BACnetTagHeader, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) *BACnetServiceAckAtomicReadFileStream {
	_result := &BACnetServiceAckAtomicReadFileStream{
		FileStartPosition: fileStartPosition,
		FileData:          fileData,
		BACnetServiceAckAtomicReadFileStreamOrRecord: NewBACnetServiceAckAtomicReadFileStreamOrRecord(peekedTagHeader, openingTag, closingTag),
	}
	_result.Child = _result
	return _result
}

func CastBACnetServiceAckAtomicReadFileStream(structType interface{}) *BACnetServiceAckAtomicReadFileStream {
	if casted, ok := structType.(BACnetServiceAckAtomicReadFileStream); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetServiceAckAtomicReadFileStream); ok {
		return casted
	}
	if casted, ok := structType.(BACnetServiceAckAtomicReadFileStreamOrRecord); ok {
		return CastBACnetServiceAckAtomicReadFileStream(casted.Child)
	}
	if casted, ok := structType.(*BACnetServiceAckAtomicReadFileStreamOrRecord); ok {
		return CastBACnetServiceAckAtomicReadFileStream(casted.Child)
	}
	return nil
}

func (m *BACnetServiceAckAtomicReadFileStream) GetTypeName() string {
	return "BACnetServiceAckAtomicReadFileStream"
}

func (m *BACnetServiceAckAtomicReadFileStream) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetServiceAckAtomicReadFileStream) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (fileStartPosition)
	lengthInBits += m.FileStartPosition.GetLengthInBits()

	// Simple field (fileData)
	lengthInBits += m.FileData.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetServiceAckAtomicReadFileStream) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckAtomicReadFileStreamParse(readBuffer utils.ReadBuffer) (*BACnetServiceAckAtomicReadFileStream, error) {
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicReadFileStream"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (fileStartPosition)
	if pullErr := readBuffer.PullContext("fileStartPosition"); pullErr != nil {
		return nil, pullErr
	}
	_fileStartPosition, _fileStartPositionErr := BACnetApplicationTagParse(readBuffer)
	if _fileStartPositionErr != nil {
		return nil, errors.Wrap(_fileStartPositionErr, "Error parsing 'fileStartPosition' field")
	}
	fileStartPosition := CastBACnetApplicationTagSignedInteger(_fileStartPosition)
	if closeErr := readBuffer.CloseContext("fileStartPosition"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (fileData)
	if pullErr := readBuffer.PullContext("fileData"); pullErr != nil {
		return nil, pullErr
	}
	_fileData, _fileDataErr := BACnetApplicationTagParse(readBuffer)
	if _fileDataErr != nil {
		return nil, errors.Wrap(_fileDataErr, "Error parsing 'fileData' field")
	}
	fileData := CastBACnetApplicationTagOctetString(_fileData)
	if closeErr := readBuffer.CloseContext("fileData"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicReadFileStream"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckAtomicReadFileStream{
		FileStartPosition: CastBACnetApplicationTagSignedInteger(fileStartPosition),
		FileData:          CastBACnetApplicationTagOctetString(fileData),
		BACnetServiceAckAtomicReadFileStreamOrRecord: &BACnetServiceAckAtomicReadFileStreamOrRecord{},
	}
	_child.BACnetServiceAckAtomicReadFileStreamOrRecord.Child = _child
	return _child, nil
}

func (m *BACnetServiceAckAtomicReadFileStream) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicReadFileStream"); pushErr != nil {
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

		// Simple Field (fileData)
		if pushErr := writeBuffer.PushContext("fileData"); pushErr != nil {
			return pushErr
		}
		_fileDataErr := m.FileData.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("fileData"); popErr != nil {
			return popErr
		}
		if _fileDataErr != nil {
			return errors.Wrap(_fileDataErr, "Error serializing 'fileData' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicReadFileStream"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckAtomicReadFileStream) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
