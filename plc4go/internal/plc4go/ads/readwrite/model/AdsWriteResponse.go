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

// AdsWriteResponse is the data-structure of this message
type AdsWriteResponse struct {
	*AdsData
	Result ReturnCode
}

// IAdsWriteResponse is the corresponding interface of AdsWriteResponse
type IAdsWriteResponse interface {
	IAdsData
	// GetResult returns Result (property field)
	GetResult() ReturnCode
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

func (m *AdsWriteResponse) GetCommandId() CommandId {
	return CommandId_ADS_WRITE
}

func (m *AdsWriteResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *AdsWriteResponse) InitializeParent(parent *AdsData) {}

func (m *AdsWriteResponse) GetParent() *AdsData {
	return m.AdsData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *AdsWriteResponse) GetResult() ReturnCode {
	return m.Result
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsWriteResponse factory function for AdsWriteResponse
func NewAdsWriteResponse(result ReturnCode) *AdsWriteResponse {
	_result := &AdsWriteResponse{
		Result:  result,
		AdsData: NewAdsData(),
	}
	_result.Child = _result
	return _result
}

func CastAdsWriteResponse(structType interface{}) *AdsWriteResponse {
	if casted, ok := structType.(AdsWriteResponse); ok {
		return &casted
	}
	if casted, ok := structType.(*AdsWriteResponse); ok {
		return casted
	}
	if casted, ok := structType.(AdsData); ok {
		return CastAdsWriteResponse(casted.Child)
	}
	if casted, ok := structType.(*AdsData); ok {
		return CastAdsWriteResponse(casted.Child)
	}
	return nil
}

func (m *AdsWriteResponse) GetTypeName() string {
	return "AdsWriteResponse"
}

func (m *AdsWriteResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AdsWriteResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (result)
	lengthInBits += 32

	return lengthInBits
}

func (m *AdsWriteResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsWriteResponseParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (*AdsWriteResponse, error) {
	if pullErr := readBuffer.PullContext("AdsWriteResponse"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (result)
	if pullErr := readBuffer.PullContext("result"); pullErr != nil {
		return nil, pullErr
	}
	_result, _resultErr := ReturnCodeParse(readBuffer)
	if _resultErr != nil {
		return nil, errors.Wrap(_resultErr, "Error parsing 'result' field")
	}
	result := _result
	if closeErr := readBuffer.CloseContext("result"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("AdsWriteResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsWriteResponse{
		Result:  result,
		AdsData: &AdsData{},
	}
	_child.AdsData.Child = _child
	return _child, nil
}

func (m *AdsWriteResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsWriteResponse"); pushErr != nil {
			return pushErr
		}

		// Simple Field (result)
		if pushErr := writeBuffer.PushContext("result"); pushErr != nil {
			return pushErr
		}
		_resultErr := m.Result.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("result"); popErr != nil {
			return popErr
		}
		if _resultErr != nil {
			return errors.Wrap(_resultErr, "Error serializing 'result' field")
		}

		if popErr := writeBuffer.PopContext("AdsWriteResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsWriteResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
