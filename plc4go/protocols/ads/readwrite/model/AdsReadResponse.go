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

// AdsReadResponse is the data-structure of this message
type AdsReadResponse struct {
	*AdsData
	Result ReturnCode
	Data   []byte
}

// IAdsReadResponse is the corresponding interface of AdsReadResponse
type IAdsReadResponse interface {
	IAdsData
	// GetResult returns Result (property field)
	GetResult() ReturnCode
	// GetData returns Data (property field)
	GetData() []byte
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

func (m *AdsReadResponse) GetCommandId() CommandId {
	return CommandId_ADS_READ
}

func (m *AdsReadResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *AdsReadResponse) InitializeParent(parent *AdsData) {}

func (m *AdsReadResponse) GetParent() *AdsData {
	return m.AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *AdsReadResponse) GetResult() ReturnCode {
	return m.Result
}

func (m *AdsReadResponse) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsReadResponse factory function for AdsReadResponse
func NewAdsReadResponse(result ReturnCode, data []byte) *AdsReadResponse {
	_result := &AdsReadResponse{
		Result:  result,
		Data:    data,
		AdsData: NewAdsData(),
	}
	_result.Child = _result
	return _result
}

func CastAdsReadResponse(structType interface{}) *AdsReadResponse {
	if casted, ok := structType.(AdsReadResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*AdsReadResponse); ok {
		return casted
	}
	if casted, ok := structType.(AdsData); ok {
		return CastAdsReadResponse(casted.Child)
	}
	if casted, ok := structType.(*AdsData); ok {
		return CastAdsReadResponse(casted.Child)
	}
	return nil
}

func (m *AdsReadResponse) GetTypeName() string {
	return "AdsReadResponse"
}

func (m *AdsReadResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AdsReadResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (result)
	lengthInBits += 32

	// Implicit Field (length)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *AdsReadResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsReadResponseParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (*AdsReadResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsReadResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (result)
	if pullErr := readBuffer.PullContext("result"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for result")
	}
	_result, _resultErr := ReturnCodeParse(readBuffer)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field")
	}
	result := _result
	if closeErr := readBuffer.CloseContext("result"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for result")
	}

	// Implicit Field (length) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	length, _lengthErr := readBuffer.ReadUint32("length", 32)
	_ = length
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(length)
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("AdsReadResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadResponse")
	}

	// Create a partially initialized instance
	_child := &AdsReadResponse{
		Result:  result,
		Data:    data,
		AdsData: &AdsData{},
	}
	_child.AdsData.Child = _child
	return _child, nil
}

func (m *AdsReadResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadResponse")
		}

		// Simple Field (result)
		if pushErr := writeBuffer.PushContext("result"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for result")
		}
		_resultErr := m.Result.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("result"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for result")
		}
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		length := uint32(uint32(len(m.GetData())))
		_lengthErr := writeBuffer.WriteUint32("length", 32, (length))
		if _lengthErr != nil {
			return errors.Wrap(_lengthErr, "Error serializing 'length' field")
		}

		// Array Field (data)
		if m.Data != nil {
			// Byte Array field (data)
			_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
			}
		}

		if popErr := writeBuffer.PopContext("AdsReadResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsReadResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
