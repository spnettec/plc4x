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
type AdsWriteControlResponse struct {
	*AdsData
	Result ReturnCode
}

// The corresponding interface
type IAdsWriteControlResponse interface {
	// GetResult returns Result
	GetResult() ReturnCode
	// LengthInBytes returns the length in bytes
	LengthInBytes() uint16
	// LengthInBits returns the length in bits
	LengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsWriteControlResponse) CommandId() CommandId {
	return CommandId_ADS_WRITE_CONTROL
}

func (m *AdsWriteControlResponse) GetCommandId() CommandId {
	return CommandId_ADS_WRITE_CONTROL
}

func (m *AdsWriteControlResponse) Response() bool {
	return bool(true)
}

func (m *AdsWriteControlResponse) GetResponse() bool {
	return bool(true)
}

func (m *AdsWriteControlResponse) InitializeParent(parent *AdsData) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *AdsWriteControlResponse) GetResult() ReturnCode {
	return m.Result
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewAdsWriteControlResponse(result ReturnCode) *AdsData {
	child := &AdsWriteControlResponse{
		Result:  result,
		AdsData: NewAdsData(),
	}
	child.Child = child
	return child.AdsData
}

func CastAdsWriteControlResponse(structType interface{}) *AdsWriteControlResponse {
	castFunc := func(typ interface{}) *AdsWriteControlResponse {
		if casted, ok := typ.(AdsWriteControlResponse); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsWriteControlResponse); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsWriteControlResponse(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsWriteControlResponse(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsWriteControlResponse) GetTypeName() string {
	return "AdsWriteControlResponse"
}

func (m *AdsWriteControlResponse) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsWriteControlResponse) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (result)
	lengthInBits += 32

	return lengthInBits
}

func (m *AdsWriteControlResponse) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsWriteControlResponseParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (*AdsData, error) {
	if pullErr := readBuffer.PullContext("AdsWriteControlResponse"); pullErr != nil {
		return nil, pullErr
	}

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

	if closeErr := readBuffer.CloseContext("AdsWriteControlResponse"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsWriteControlResponse{
		Result:  result,
		AdsData: &AdsData{},
	}
	_child.AdsData.Child = _child
	return _child.AdsData, nil
}

func (m *AdsWriteControlResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsWriteControlResponse"); pushErr != nil {
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

		if popErr := writeBuffer.PopContext("AdsWriteControlResponse"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsWriteControlResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
