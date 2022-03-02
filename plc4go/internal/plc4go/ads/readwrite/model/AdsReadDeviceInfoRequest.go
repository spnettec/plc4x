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
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type AdsReadDeviceInfoRequest struct {
	*AdsData
}

// The corresponding interface
type IAdsReadDeviceInfoRequest interface {
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *AdsReadDeviceInfoRequest) CommandId() CommandId {
	return CommandId_ADS_READ_DEVICE_INFO
}

func (m *AdsReadDeviceInfoRequest) GetCommandId() CommandId {
	return CommandId_ADS_READ_DEVICE_INFO
}

func (m *AdsReadDeviceInfoRequest) Response() bool {
	return bool(false)
}

func (m *AdsReadDeviceInfoRequest) GetResponse() bool {
	return bool(false)
}

func (m *AdsReadDeviceInfoRequest) InitializeParent(parent *AdsData) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewAdsReadDeviceInfoRequest factory function for AdsReadDeviceInfoRequest
func NewAdsReadDeviceInfoRequest() *AdsData {
	child := &AdsReadDeviceInfoRequest{
		AdsData: NewAdsData(),
	}
	child.Child = child
	return child.AdsData
}

func CastAdsReadDeviceInfoRequest(structType interface{}) *AdsReadDeviceInfoRequest {
	castFunc := func(typ interface{}) *AdsReadDeviceInfoRequest {
		if casted, ok := typ.(AdsReadDeviceInfoRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsReadDeviceInfoRequest); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsReadDeviceInfoRequest(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsReadDeviceInfoRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsReadDeviceInfoRequest) GetTypeName() string {
	return "AdsReadDeviceInfoRequest"
}

func (m *AdsReadDeviceInfoRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AdsReadDeviceInfoRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *AdsReadDeviceInfoRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsReadDeviceInfoRequestParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (*AdsData, error) {
	if pullErr := readBuffer.PullContext("AdsReadDeviceInfoRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AdsReadDeviceInfoRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsReadDeviceInfoRequest{
		AdsData: &AdsData{},
	}
	_child.AdsData.Child = _child
	return _child.AdsData, nil
}

func (m *AdsReadDeviceInfoRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadDeviceInfoRequest"); pushErr != nil {
			return pushErr
		}

		if popErr := writeBuffer.PopContext("AdsReadDeviceInfoRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsReadDeviceInfoRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
