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

// AdsReadRequest is the corresponding interface of AdsReadRequest
type AdsReadRequest interface {
	utils.LengthAware
	utils.Serializable
	AdsData
	// GetIndexGroup returns IndexGroup (property field)
	GetIndexGroup() uint32
	// GetIndexOffset returns IndexOffset (property field)
	GetIndexOffset() uint32
	// GetLength returns Length (property field)
	GetLength() uint32
}

// AdsReadRequestExactly can be used when we want exactly this type and not a type which fulfills AdsReadRequest.
// This is useful for switch cases.
type AdsReadRequestExactly interface {
	AdsReadRequest
	isAdsReadRequest() bool
}

// _AdsReadRequest is the data-structure of this message
type _AdsReadRequest struct {
	*_AdsData
	IndexGroup  uint32
	IndexOffset uint32
	Length      uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsReadRequest) GetCommandId() CommandId {
	return CommandId_ADS_READ
}

func (m *_AdsReadRequest) GetResponse() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsReadRequest) InitializeParent(parent AdsData) {}

func (m *_AdsReadRequest) GetParent() AdsData {
	return m._AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsReadRequest) GetIndexGroup() uint32 {
	return m.IndexGroup
}

func (m *_AdsReadRequest) GetIndexOffset() uint32 {
	return m.IndexOffset
}

func (m *_AdsReadRequest) GetLength() uint32 {
	return m.Length
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsReadRequest factory function for _AdsReadRequest
func NewAdsReadRequest(indexGroup uint32, indexOffset uint32, length uint32) *_AdsReadRequest {
	_result := &_AdsReadRequest{
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		Length:      length,
		_AdsData:    NewAdsData(),
	}
	_result._AdsData._AdsDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsReadRequest(structType interface{}) AdsReadRequest {
	if casted, ok := structType.(AdsReadRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsReadRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsReadRequest) GetTypeName() string {
	return "AdsReadRequest"
}

func (m *_AdsReadRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsReadRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (indexGroup)
	lengthInBits += 32

	// Simple field (indexOffset)
	lengthInBits += 32

	// Simple field (length)
	lengthInBits += 32

	return lengthInBits
}

func (m *_AdsReadRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsReadRequestParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (AdsReadRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsReadRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (indexGroup)
	_indexGroup, _indexGroupErr := readBuffer.ReadUint32("indexGroup", 32)
	if _indexGroupErr != nil {
		return nil, errors.Wrap(_indexGroupErr, "Error parsing 'indexGroup' field of AdsReadRequest")
	}
	indexGroup := _indexGroup

	// Simple Field (indexOffset)
	_indexOffset, _indexOffsetErr := readBuffer.ReadUint32("indexOffset", 32)
	if _indexOffsetErr != nil {
		return nil, errors.Wrap(_indexOffsetErr, "Error parsing 'indexOffset' field of AdsReadRequest")
	}
	indexOffset := _indexOffset

	// Simple Field (length)
	_length, _lengthErr := readBuffer.ReadUint32("length", 32)
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field of AdsReadRequest")
	}
	length := _length

	if closeErr := readBuffer.CloseContext("AdsReadRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadRequest")
	}

	// Create a partially initialized instance
	_child := &_AdsReadRequest{
		IndexGroup:  indexGroup,
		IndexOffset: indexOffset,
		Length:      length,
		_AdsData:    &_AdsData{},
	}
	_child._AdsData._AdsDataChildRequirements = _child
	return _child, nil
}

func (m *_AdsReadRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadRequest")
		}

		// Simple Field (indexGroup)
		indexGroup := uint32(m.GetIndexGroup())
		_indexGroupErr := writeBuffer.WriteUint32("indexGroup", 32, (indexGroup))
		if _indexGroupErr != nil {
			return errors.Wrap(_indexGroupErr, "Error serializing 'indexGroup' field")
		}

		// Simple Field (indexOffset)
		indexOffset := uint32(m.GetIndexOffset())
		_indexOffsetErr := writeBuffer.WriteUint32("indexOffset", 32, (indexOffset))
		if _indexOffsetErr != nil {
			return errors.Wrap(_indexOffsetErr, "Error serializing 'indexOffset' field")
		}

		// Simple Field (length)
		length := uint32(m.GetLength())
		_lengthErr := writeBuffer.WriteUint32("length", 32, (length))
		if _lengthErr != nil {
			return errors.Wrap(_lengthErr, "Error serializing 'length' field")
		}

		if popErr := writeBuffer.PopContext("AdsReadRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_AdsReadRequest) isAdsReadRequest() bool {
	return true
}

func (m *_AdsReadRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
