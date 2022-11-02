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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// AdsWriteRequest is the corresponding interface of AdsWriteRequest
type AdsWriteRequest interface {
	utils.LengthAware
	utils.Serializable
	AmsPacket
	// GetIndexGroup returns IndexGroup (property field)
	GetIndexGroup() uint32
	// GetIndexOffset returns IndexOffset (property field)
	GetIndexOffset() uint32
	// GetData returns Data (property field)
	GetData() []byte
}

// AdsWriteRequestExactly can be used when we want exactly this type and not a type which fulfills AdsWriteRequest.
// This is useful for switch cases.
type AdsWriteRequestExactly interface {
	AdsWriteRequest
	isAdsWriteRequest() bool
}

// _AdsWriteRequest is the data-structure of this message
type _AdsWriteRequest struct {
	*_AmsPacket
        IndexGroup uint32
        IndexOffset uint32
        Data []byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsWriteRequest)  GetCommandId() CommandId {
return CommandId_ADS_WRITE}

func (m *_AdsWriteRequest)  GetResponse() bool {
return bool(false)}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsWriteRequest) InitializeParent(parent AmsPacket , targetAmsNetId AmsNetId , targetAmsPort uint16 , sourceAmsNetId AmsNetId , sourceAmsPort uint16 , errorCode uint32 , invokeId uint32 ) {	m.TargetAmsNetId = targetAmsNetId
	m.TargetAmsPort = targetAmsPort
	m.SourceAmsNetId = sourceAmsNetId
	m.SourceAmsPort = sourceAmsPort
	m.ErrorCode = errorCode
	m.InvokeId = invokeId
}

func (m *_AdsWriteRequest)  GetParent() AmsPacket {
	return m._AmsPacket
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_AdsWriteRequest) GetIndexGroup() uint32 {
	return m.IndexGroup
}

func (m *_AdsWriteRequest) GetIndexOffset() uint32 {
	return m.IndexOffset
}

func (m *_AdsWriteRequest) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewAdsWriteRequest factory function for _AdsWriteRequest
func NewAdsWriteRequest( indexGroup uint32 , indexOffset uint32 , data []byte , targetAmsNetId AmsNetId , targetAmsPort uint16 , sourceAmsNetId AmsNetId , sourceAmsPort uint16 , errorCode uint32 , invokeId uint32 ) *_AdsWriteRequest {
	_result := &_AdsWriteRequest{
		IndexGroup: indexGroup,
		IndexOffset: indexOffset,
		Data: data,
    	_AmsPacket: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result._AmsPacket._AmsPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsWriteRequest(structType interface{}) AdsWriteRequest {
    if casted, ok := structType.(AdsWriteRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsWriteRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsWriteRequest) GetTypeName() string {
	return "AdsWriteRequest"
}

func (m *_AdsWriteRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsWriteRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (indexGroup)
	lengthInBits += 32;

	// Simple field (indexOffset)
	lengthInBits += 32;

	// Implicit Field (length)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}


func (m *_AdsWriteRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsWriteRequestParse(readBuffer utils.ReadBuffer) (AdsWriteRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsWriteRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsWriteRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (indexGroup)
_indexGroup, _indexGroupErr := readBuffer.ReadUint32("indexGroup", 32)
	if _indexGroupErr != nil {
		return nil, errors.Wrap(_indexGroupErr, "Error parsing 'indexGroup' field of AdsWriteRequest")
	}
	indexGroup := _indexGroup

	// Simple Field (indexOffset)
_indexOffset, _indexOffsetErr := readBuffer.ReadUint32("indexOffset", 32)
	if _indexOffsetErr != nil {
		return nil, errors.Wrap(_indexOffsetErr, "Error parsing 'indexOffset' field of AdsWriteRequest")
	}
	indexOffset := _indexOffset

	// Implicit Field (length) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := readBuffer.ReadUint32("length", 32)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field of AdsWriteRequest")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(length)
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of AdsWriteRequest")
	}

	if closeErr := readBuffer.CloseContext("AdsWriteRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsWriteRequest")
	}

	// Create a partially initialized instance
	_child := &_AdsWriteRequest{
		_AmsPacket: &_AmsPacket{
		},
		IndexGroup: indexGroup,
		IndexOffset: indexOffset,
		Data: data,
	}
	_child._AmsPacket._AmsPacketChildRequirements = _child
	return _child, nil
}

func (m *_AdsWriteRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsWriteRequest) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsWriteRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsWriteRequest")
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

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	length := uint32(uint32(len(m.GetData())))
	_lengthErr := writeBuffer.WriteUint32("length", 32, (length))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Array Field (data)
	// Byte Array field (data)
	if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

		if popErr := writeBuffer.PopContext("AdsWriteRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsWriteRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_AdsWriteRequest) isAdsWriteRequest() bool {
	return true
}

func (m *_AdsWriteRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



