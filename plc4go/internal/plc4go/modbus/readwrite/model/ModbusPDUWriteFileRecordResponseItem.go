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
type ModbusPDUWriteFileRecordResponseItem struct {
	ReferenceType uint8
	FileNumber    uint16
	RecordNumber  uint16
	RecordData    []byte
}

// The corresponding interface
type IModbusPDUWriteFileRecordResponseItem interface {
	// GetReferenceType returns ReferenceType (property field)
	GetReferenceType() uint8
	// GetFileNumber returns FileNumber (property field)
	GetFileNumber() uint16
	// GetRecordNumber returns RecordNumber (property field)
	GetRecordNumber() uint16
	// GetRecordData returns RecordData (property field)
	GetRecordData() []byte
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *ModbusPDUWriteFileRecordResponseItem) GetReferenceType() uint8 {
	return m.ReferenceType
}

func (m *ModbusPDUWriteFileRecordResponseItem) GetFileNumber() uint16 {
	return m.FileNumber
}

func (m *ModbusPDUWriteFileRecordResponseItem) GetRecordNumber() uint16 {
	return m.RecordNumber
}

func (m *ModbusPDUWriteFileRecordResponseItem) GetRecordData() []byte {
	return m.RecordData
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewModbusPDUWriteFileRecordResponseItem factory function for ModbusPDUWriteFileRecordResponseItem
func NewModbusPDUWriteFileRecordResponseItem(referenceType uint8, fileNumber uint16, recordNumber uint16, recordData []byte) *ModbusPDUWriteFileRecordResponseItem {
	return &ModbusPDUWriteFileRecordResponseItem{ReferenceType: referenceType, FileNumber: fileNumber, RecordNumber: recordNumber, RecordData: recordData}
}

func CastModbusPDUWriteFileRecordResponseItem(structType interface{}) *ModbusPDUWriteFileRecordResponseItem {
	if casted, ok := structType.(ModbusPDUWriteFileRecordResponseItem); ok {
		return &casted
	}
	if casted, ok := structType.(*ModbusPDUWriteFileRecordResponseItem); ok {
		return casted
	}
	return nil
}

func (m *ModbusPDUWriteFileRecordResponseItem) GetTypeName() string {
	return "ModbusPDUWriteFileRecordResponseItem"
}

func (m *ModbusPDUWriteFileRecordResponseItem) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ModbusPDUWriteFileRecordResponseItem) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (referenceType)
	lengthInBits += 8

	// Simple field (fileNumber)
	lengthInBits += 16

	// Simple field (recordNumber)
	lengthInBits += 16

	// Implicit Field (recordLength)
	lengthInBits += 16

	// Array field
	if len(m.RecordData) > 0 {
		lengthInBits += 8 * uint16(len(m.RecordData))
	}

	return lengthInBits
}

func (m *ModbusPDUWriteFileRecordResponseItem) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUWriteFileRecordResponseItemParse(readBuffer utils.ReadBuffer) (*ModbusPDUWriteFileRecordResponseItem, error) {
	if pullErr := readBuffer.PullContext("ModbusPDUWriteFileRecordResponseItem"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (referenceType)
	_referenceType, _referenceTypeErr := readBuffer.ReadUint8("referenceType", 8)
	if _referenceTypeErr != nil {
		return nil, errors.Wrap(_referenceTypeErr, "Error parsing 'referenceType' field")
	}
	referenceType := _referenceType

	// Simple Field (fileNumber)
	_fileNumber, _fileNumberErr := readBuffer.ReadUint16("fileNumber", 16)
	if _fileNumberErr != nil {
		return nil, errors.Wrap(_fileNumberErr, "Error parsing 'fileNumber' field")
	}
	fileNumber := _fileNumber

	// Simple Field (recordNumber)
	_recordNumber, _recordNumberErr := readBuffer.ReadUint16("recordNumber", 16)
	if _recordNumberErr != nil {
		return nil, errors.Wrap(_recordNumberErr, "Error parsing 'recordNumber' field")
	}
	recordNumber := _recordNumber

	// Implicit Field (recordLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	recordLength, _recordLengthErr := readBuffer.ReadUint16("recordLength", 16)
	_ = recordLength
	if _recordLengthErr != nil {
		return nil, errors.Wrap(_recordLengthErr, "Error parsing 'recordLength' field")
	}
	// Byte Array field (recordData)
	numberOfBytesrecordData := int(recordLength)
	recordData, _readArrayErr := readBuffer.ReadByteArray("recordData", numberOfBytesrecordData)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'recordData' field")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteFileRecordResponseItem"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewModbusPDUWriteFileRecordResponseItem(referenceType, fileNumber, recordNumber, recordData), nil
}

func (m *ModbusPDUWriteFileRecordResponseItem) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("ModbusPDUWriteFileRecordResponseItem"); pushErr != nil {
		return pushErr
	}

	// Simple Field (referenceType)
	referenceType := uint8(m.ReferenceType)
	_referenceTypeErr := writeBuffer.WriteUint8("referenceType", 8, (referenceType))
	if _referenceTypeErr != nil {
		return errors.Wrap(_referenceTypeErr, "Error serializing 'referenceType' field")
	}

	// Simple Field (fileNumber)
	fileNumber := uint16(m.FileNumber)
	_fileNumberErr := writeBuffer.WriteUint16("fileNumber", 16, (fileNumber))
	if _fileNumberErr != nil {
		return errors.Wrap(_fileNumberErr, "Error serializing 'fileNumber' field")
	}

	// Simple Field (recordNumber)
	recordNumber := uint16(m.RecordNumber)
	_recordNumberErr := writeBuffer.WriteUint16("recordNumber", 16, (recordNumber))
	if _recordNumberErr != nil {
		return errors.Wrap(_recordNumberErr, "Error serializing 'recordNumber' field")
	}

	// Implicit Field (recordLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	recordLength := uint16(uint16(uint16(len(m.GetRecordData()))) / uint16(uint16(2)))
	_recordLengthErr := writeBuffer.WriteUint16("recordLength", 16, (recordLength))
	if _recordLengthErr != nil {
		return errors.Wrap(_recordLengthErr, "Error serializing 'recordLength' field")
	}

	// Array Field (recordData)
	if m.RecordData != nil {
		// Byte Array field (recordData)
		_writeArrayErr := writeBuffer.WriteByteArray("recordData", m.RecordData)
		if _writeArrayErr != nil {
			return errors.Wrap(_writeArrayErr, "Error serializing 'recordData' field")
		}
	}

	if popErr := writeBuffer.PopContext("ModbusPDUWriteFileRecordResponseItem"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *ModbusPDUWriteFileRecordResponseItem) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
