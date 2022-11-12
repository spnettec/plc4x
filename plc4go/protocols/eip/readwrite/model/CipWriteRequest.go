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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// CipWriteRequest is the corresponding interface of CipWriteRequest
type CipWriteRequest interface {
	utils.LengthAware
	utils.Serializable
	CipService
	// GetRequestPathSize returns RequestPathSize (property field)
	GetRequestPathSize() int8
	// GetTag returns Tag (property field)
	GetTag() []byte
	// GetDataType returns DataType (property field)
	GetDataType() CIPDataTypeCode
	// GetElementNb returns ElementNb (property field)
	GetElementNb() uint16
	// GetData returns Data (property field)
	GetData() []byte
}

// CipWriteRequestExactly can be used when we want exactly this type and not a type which fulfills CipWriteRequest.
// This is useful for switch cases.
type CipWriteRequestExactly interface {
	CipWriteRequest
	isCipWriteRequest() bool
}

// _CipWriteRequest is the data-structure of this message
type _CipWriteRequest struct {
	*_CipService
        RequestPathSize int8
        Tag []byte
        DataType CIPDataTypeCode
        ElementNb uint16
        Data []byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipWriteRequest)  GetService() uint8 {
return 0x4D}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipWriteRequest) InitializeParent(parent CipService ) {}

func (m *_CipWriteRequest)  GetParent() CipService {
	return m._CipService
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipWriteRequest) GetRequestPathSize() int8 {
	return m.RequestPathSize
}

func (m *_CipWriteRequest) GetTag() []byte {
	return m.Tag
}

func (m *_CipWriteRequest) GetDataType() CIPDataTypeCode {
	return m.DataType
}

func (m *_CipWriteRequest) GetElementNb() uint16 {
	return m.ElementNb
}

func (m *_CipWriteRequest) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewCipWriteRequest factory function for _CipWriteRequest
func NewCipWriteRequest( requestPathSize int8 , tag []byte , dataType CIPDataTypeCode , elementNb uint16 , data []byte , serviceLen uint16 ) *_CipWriteRequest {
	_result := &_CipWriteRequest{
		RequestPathSize: requestPathSize,
		Tag: tag,
		DataType: dataType,
		ElementNb: elementNb,
		Data: data,
    	_CipService: NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipWriteRequest(structType interface{}) CipWriteRequest {
    if casted, ok := structType.(CipWriteRequest); ok {
		return casted
	}
	if casted, ok := structType.(*CipWriteRequest); ok {
		return *casted
	}
	return nil
}

func (m *_CipWriteRequest) GetTypeName() string {
	return "CipWriteRequest"
}

func (m *_CipWriteRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CipWriteRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (requestPathSize)
	lengthInBits += 8;

	// Array field
	if len(m.Tag) > 0 {
		lengthInBits += 8 * uint16(len(m.Tag))
	}

	// Simple field (dataType)
	lengthInBits += 16

	// Simple field (elementNb)
	lengthInBits += 16;

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}


func (m *_CipWriteRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CipWriteRequestParse(theBytes []byte, serviceLen uint16) (CipWriteRequest, error) {
	return CipWriteRequestParseWithBuffer(utils.NewReadBufferByteBased(theBytes), serviceLen)
}

func CipWriteRequestParseWithBuffer(readBuffer utils.ReadBuffer, serviceLen uint16) (CipWriteRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CipWriteRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipWriteRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestPathSize)
_requestPathSize, _requestPathSizeErr := readBuffer.ReadInt8("requestPathSize", 8)
	if _requestPathSizeErr != nil {
		return nil, errors.Wrap(_requestPathSizeErr, "Error parsing 'requestPathSize' field of CipWriteRequest")
	}
	requestPathSize := _requestPathSize
	// Byte Array field (tag)
	numberOfBytestag := int(uint16(requestPathSize) * uint16(uint16(2)))
	tag, _readArrayErr := readBuffer.ReadByteArray("tag", numberOfBytestag)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'tag' field of CipWriteRequest")
	}

	// Simple Field (dataType)
	if pullErr := readBuffer.PullContext("dataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dataType")
	}
_dataType, _dataTypeErr := CIPDataTypeCodeParseWithBuffer(readBuffer)
	if _dataTypeErr != nil {
		return nil, errors.Wrap(_dataTypeErr, "Error parsing 'dataType' field of CipWriteRequest")
	}
	dataType := _dataType
	if closeErr := readBuffer.CloseContext("dataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dataType")
	}

	// Simple Field (elementNb)
_elementNb, _elementNbErr := readBuffer.ReadUint16("elementNb", 16)
	if _elementNbErr != nil {
		return nil, errors.Wrap(_elementNbErr, "Error parsing 'elementNb' field of CipWriteRequest")
	}
	elementNb := _elementNb
	// Byte Array field (data)
	numberOfBytesdata := int(uint16(dataType.Size()) * uint16(elementNb))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of CipWriteRequest")
	}

	if closeErr := readBuffer.CloseContext("CipWriteRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipWriteRequest")
	}

	// Create a partially initialized instance
	_child := &_CipWriteRequest{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
		RequestPathSize: requestPathSize,
		Tag: tag,
		DataType: dataType,
		ElementNb: elementNb,
		Data: data,
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_CipWriteRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipWriteRequest) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipWriteRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipWriteRequest")
		}

	// Simple Field (requestPathSize)
	requestPathSize := int8(m.GetRequestPathSize())
	_requestPathSizeErr := writeBuffer.WriteInt8("requestPathSize", 8, (requestPathSize))
	if _requestPathSizeErr != nil {
		return errors.Wrap(_requestPathSizeErr, "Error serializing 'requestPathSize' field")
	}

	// Array Field (tag)
	// Byte Array field (tag)
	if err := writeBuffer.WriteByteArray("tag", m.GetTag()); err != nil {
		return errors.Wrap(err, "Error serializing 'tag' field")
	}

	// Simple Field (dataType)
	if pushErr := writeBuffer.PushContext("dataType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for dataType")
	}
	_dataTypeErr := writeBuffer.WriteSerializable(m.GetDataType())
	if popErr := writeBuffer.PopContext("dataType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for dataType")
	}
	if _dataTypeErr != nil {
		return errors.Wrap(_dataTypeErr, "Error serializing 'dataType' field")
	}

	// Simple Field (elementNb)
	elementNb := uint16(m.GetElementNb())
	_elementNbErr := writeBuffer.WriteUint16("elementNb", 16, (elementNb))
	if _elementNbErr != nil {
		return errors.Wrap(_elementNbErr, "Error serializing 'elementNb' field")
	}

	// Array Field (data)
	// Byte Array field (data)
	if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

		if popErr := writeBuffer.PopContext("CipWriteRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipWriteRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_CipWriteRequest) isCipWriteRequest() bool {
	return true
}

func (m *_CipWriteRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



