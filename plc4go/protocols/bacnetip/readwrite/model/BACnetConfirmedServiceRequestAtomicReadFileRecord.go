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

// BACnetConfirmedServiceRequestAtomicReadFileRecord is the corresponding interface of BACnetConfirmedServiceRequestAtomicReadFileRecord
type BACnetConfirmedServiceRequestAtomicReadFileRecord interface {
	BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
	// GetFileStartRecord returns FileStartRecord (property field)
	GetFileStartRecord() BACnetApplicationTagSignedInteger
	// GetRequestRecordCount returns RequestRecordCount (property field)
	GetRequestRecordCount() BACnetApplicationTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetConfirmedServiceRequestAtomicReadFileRecord is the data-structure of this message
type _BACnetConfirmedServiceRequestAtomicReadFileRecord struct {
	*_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
	FileStartRecord    BACnetApplicationTagSignedInteger
	RequestRecordCount BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) InitializeParent(parent BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord, peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) {
	m.PeekedTagHeader = peekedTagHeader
	m.OpeningTag = openingTag
	m.ClosingTag = closingTag
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetParent() BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord {
	return m._BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetFileStartRecord() BACnetApplicationTagSignedInteger {
	return m.FileStartRecord
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetRequestRecordCount() BACnetApplicationTagUnsignedInteger {
	return m.RequestRecordCount
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestAtomicReadFileRecord factory function for _BACnetConfirmedServiceRequestAtomicReadFileRecord
func NewBACnetConfirmedServiceRequestAtomicReadFileRecord(fileStartRecord BACnetApplicationTagSignedInteger, requestRecordCount BACnetApplicationTagUnsignedInteger, peekedTagHeader BACnetTagHeader, openingTag BACnetOpeningTag, closingTag BACnetClosingTag) *_BACnetConfirmedServiceRequestAtomicReadFileRecord {
	_result := &_BACnetConfirmedServiceRequestAtomicReadFileRecord{
		FileStartRecord:    fileStartRecord,
		RequestRecordCount: requestRecordCount,
		_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord: NewBACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord(peekedTagHeader, openingTag, closingTag),
	}
	_result._BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord._BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestAtomicReadFileRecord(structType interface{}) BACnetConfirmedServiceRequestAtomicReadFileRecord {
	if casted, ok := structType.(BACnetConfirmedServiceRequestAtomicReadFileRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestAtomicReadFileRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetTypeName() string {
	return "BACnetConfirmedServiceRequestAtomicReadFileRecord"
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (fileStartRecord)
	lengthInBits += m.FileStartRecord.GetLengthInBits()

	// Simple field (requestRecordCount)
	lengthInBits += m.RequestRecordCount.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestAtomicReadFileRecordParse(readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestAtomicReadFileRecord, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestAtomicReadFileRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestAtomicReadFileRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (fileStartRecord)
	if pullErr := readBuffer.PullContext("fileStartRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for fileStartRecord")
	}
	_fileStartRecord, _fileStartRecordErr := BACnetApplicationTagParse(readBuffer)
	if _fileStartRecordErr != nil {
		return nil, errors.Wrap(_fileStartRecordErr, "Error parsing 'fileStartRecord' field")
	}
	fileStartRecord := _fileStartRecord.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("fileStartRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for fileStartRecord")
	}

	// Simple Field (requestRecordCount)
	if pullErr := readBuffer.PullContext("requestRecordCount"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestRecordCount")
	}
	_requestRecordCount, _requestRecordCountErr := BACnetApplicationTagParse(readBuffer)
	if _requestRecordCountErr != nil {
		return nil, errors.Wrap(_requestRecordCountErr, "Error parsing 'requestRecordCount' field")
	}
	requestRecordCount := _requestRecordCount.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("requestRecordCount"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestRecordCount")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestAtomicReadFileRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestAtomicReadFileRecord")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestAtomicReadFileRecord{
		FileStartRecord:    fileStartRecord,
		RequestRecordCount: requestRecordCount,
		_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord: &_BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord{},
	}
	_child._BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecord._BACnetConfirmedServiceRequestAtomicReadFileStreamOrRecordChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestAtomicReadFileRecord"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestAtomicReadFileRecord")
		}

		// Simple Field (fileStartRecord)
		if pushErr := writeBuffer.PushContext("fileStartRecord"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for fileStartRecord")
		}
		_fileStartRecordErr := writeBuffer.WriteSerializable(m.GetFileStartRecord())
		if popErr := writeBuffer.PopContext("fileStartRecord"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for fileStartRecord")
		}
		if _fileStartRecordErr != nil {
			return errors.Wrap(_fileStartRecordErr, "Error serializing 'fileStartRecord' field")
		}

		// Simple Field (requestRecordCount)
		if pushErr := writeBuffer.PushContext("requestRecordCount"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestRecordCount")
		}
		_requestRecordCountErr := writeBuffer.WriteSerializable(m.GetRequestRecordCount())
		if popErr := writeBuffer.PopContext("requestRecordCount"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestRecordCount")
		}
		if _requestRecordCountErr != nil {
			return errors.Wrap(_requestRecordCountErr, "Error serializing 'requestRecordCount' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestAtomicReadFileRecord"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestAtomicReadFileRecord")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestAtomicReadFileRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
