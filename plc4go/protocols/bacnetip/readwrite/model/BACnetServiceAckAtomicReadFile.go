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

// BACnetServiceAckAtomicReadFile is the data-structure of this message
type BACnetServiceAckAtomicReadFile struct {
	*BACnetServiceAck
	EndOfFile    *BACnetApplicationTagBoolean
	AccessMethod *BACnetServiceAckAtomicReadFileStreamOrRecord

	// Arguments.
	ServiceAckLength uint16
}

// IBACnetServiceAckAtomicReadFile is the corresponding interface of BACnetServiceAckAtomicReadFile
type IBACnetServiceAckAtomicReadFile interface {
	IBACnetServiceAck
	// GetEndOfFile returns EndOfFile (property field)
	GetEndOfFile() *BACnetApplicationTagBoolean
	// GetAccessMethod returns AccessMethod (property field)
	GetAccessMethod() *BACnetServiceAckAtomicReadFileStreamOrRecord
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

func (m *BACnetServiceAckAtomicReadFile) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_ATOMIC_READ_FILE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetServiceAckAtomicReadFile) InitializeParent(parent *BACnetServiceAck) {}

func (m *BACnetServiceAckAtomicReadFile) GetParent() *BACnetServiceAck {
	return m.BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetServiceAckAtomicReadFile) GetEndOfFile() *BACnetApplicationTagBoolean {
	return m.EndOfFile
}

func (m *BACnetServiceAckAtomicReadFile) GetAccessMethod() *BACnetServiceAckAtomicReadFileStreamOrRecord {
	return m.AccessMethod
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckAtomicReadFile factory function for BACnetServiceAckAtomicReadFile
func NewBACnetServiceAckAtomicReadFile(endOfFile *BACnetApplicationTagBoolean, accessMethod *BACnetServiceAckAtomicReadFileStreamOrRecord, serviceAckLength uint16) *BACnetServiceAckAtomicReadFile {
	_result := &BACnetServiceAckAtomicReadFile{
		EndOfFile:        endOfFile,
		AccessMethod:     accessMethod,
		BACnetServiceAck: NewBACnetServiceAck(serviceAckLength),
	}
	_result.Child = _result
	return _result
}

func CastBACnetServiceAckAtomicReadFile(structType interface{}) *BACnetServiceAckAtomicReadFile {
	if casted, ok := structType.(BACnetServiceAckAtomicReadFile); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetServiceAckAtomicReadFile); ok {
		return casted
	}
	if casted, ok := structType.(BACnetServiceAck); ok {
		return CastBACnetServiceAckAtomicReadFile(casted.Child)
	}
	if casted, ok := structType.(*BACnetServiceAck); ok {
		return CastBACnetServiceAckAtomicReadFile(casted.Child)
	}
	return nil
}

func (m *BACnetServiceAckAtomicReadFile) GetTypeName() string {
	return "BACnetServiceAckAtomicReadFile"
}

func (m *BACnetServiceAckAtomicReadFile) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetServiceAckAtomicReadFile) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (endOfFile)
	lengthInBits += m.EndOfFile.GetLengthInBits()

	// Simple field (accessMethod)
	lengthInBits += m.AccessMethod.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetServiceAckAtomicReadFile) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetServiceAckAtomicReadFileParse(readBuffer utils.ReadBuffer, serviceAckLength uint16) (*BACnetServiceAckAtomicReadFile, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetServiceAckAtomicReadFile"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckAtomicReadFile")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (endOfFile)
	if pullErr := readBuffer.PullContext("endOfFile"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for endOfFile")
	}
	_endOfFile, _endOfFileErr := BACnetApplicationTagParse(readBuffer)
	if _endOfFileErr != nil {
		return nil, errors.Wrap(_endOfFileErr, "Error parsing 'endOfFile' field")
	}
	endOfFile := CastBACnetApplicationTagBoolean(_endOfFile)
	if closeErr := readBuffer.CloseContext("endOfFile"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for endOfFile")
	}

	// Simple Field (accessMethod)
	if pullErr := readBuffer.PullContext("accessMethod"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessMethod")
	}
	_accessMethod, _accessMethodErr := BACnetServiceAckAtomicReadFileStreamOrRecordParse(readBuffer)
	if _accessMethodErr != nil {
		return nil, errors.Wrap(_accessMethodErr, "Error parsing 'accessMethod' field")
	}
	accessMethod := CastBACnetServiceAckAtomicReadFileStreamOrRecord(_accessMethod)
	if closeErr := readBuffer.CloseContext("accessMethod"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessMethod")
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckAtomicReadFile"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckAtomicReadFile")
	}

	// Create a partially initialized instance
	_child := &BACnetServiceAckAtomicReadFile{
		EndOfFile:        CastBACnetApplicationTagBoolean(endOfFile),
		AccessMethod:     CastBACnetServiceAckAtomicReadFileStreamOrRecord(accessMethod),
		BACnetServiceAck: &BACnetServiceAck{},
	}
	_child.BACnetServiceAck.Child = _child
	return _child, nil
}

func (m *BACnetServiceAckAtomicReadFile) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckAtomicReadFile"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckAtomicReadFile")
		}

		// Simple Field (endOfFile)
		if pushErr := writeBuffer.PushContext("endOfFile"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for endOfFile")
		}
		_endOfFileErr := m.EndOfFile.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("endOfFile"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for endOfFile")
		}
		if _endOfFileErr != nil {
			return errors.Wrap(_endOfFileErr, "Error serializing 'endOfFile' field")
		}

		// Simple Field (accessMethod)
		if pushErr := writeBuffer.PushContext("accessMethod"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessMethod")
		}
		_accessMethodErr := m.AccessMethod.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("accessMethod"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessMethod")
		}
		if _accessMethodErr != nil {
			return errors.Wrap(_accessMethodErr, "Error serializing 'accessMethod' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckAtomicReadFile"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckAtomicReadFile")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetServiceAckAtomicReadFile) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
