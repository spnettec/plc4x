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
type AdsWriteControlRequest struct {
	*AdsData
	AdsState    uint16
	DeviceState uint16
	Data        []byte
}

// The corresponding interface
type IAdsWriteControlRequest interface {
	// GetAdsState returns AdsState
	GetAdsState() uint16
	// GetDeviceState returns DeviceState
	GetDeviceState() uint16
	// GetData returns Data
	GetData() []byte
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
func (m *AdsWriteControlRequest) CommandId() CommandId {
	return CommandId_ADS_WRITE_CONTROL
}

func (m *AdsWriteControlRequest) GetCommandId() CommandId {
	return CommandId_ADS_WRITE_CONTROL
}

func (m *AdsWriteControlRequest) Response() bool {
	return bool(false)
}

func (m *AdsWriteControlRequest) GetResponse() bool {
	return bool(false)
}

func (m *AdsWriteControlRequest) InitializeParent(parent *AdsData) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *AdsWriteControlRequest) GetAdsState() uint16 {
	return m.AdsState
}

func (m *AdsWriteControlRequest) GetDeviceState() uint16 {
	return m.DeviceState
}

func (m *AdsWriteControlRequest) GetData() []byte {
	return m.Data
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewAdsWriteControlRequest(adsState uint16, deviceState uint16, data []byte) *AdsData {
	child := &AdsWriteControlRequest{
		AdsState:    adsState,
		DeviceState: deviceState,
		Data:        data,
		AdsData:     NewAdsData(),
	}
	child.Child = child
	return child.AdsData
}

func CastAdsWriteControlRequest(structType interface{}) *AdsWriteControlRequest {
	castFunc := func(typ interface{}) *AdsWriteControlRequest {
		if casted, ok := typ.(AdsWriteControlRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*AdsWriteControlRequest); ok {
			return casted
		}
		if casted, ok := typ.(AdsData); ok {
			return CastAdsWriteControlRequest(casted.Child)
		}
		if casted, ok := typ.(*AdsData); ok {
			return CastAdsWriteControlRequest(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *AdsWriteControlRequest) GetTypeName() string {
	return "AdsWriteControlRequest"
}

func (m *AdsWriteControlRequest) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *AdsWriteControlRequest) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (adsState)
	lengthInBits += 16

	// Simple field (deviceState)
	lengthInBits += 16

	// Implicit Field (length)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *AdsWriteControlRequest) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func AdsWriteControlRequestParse(readBuffer utils.ReadBuffer, commandId CommandId, response bool) (*AdsData, error) {
	if pullErr := readBuffer.PullContext("AdsWriteControlRequest"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (adsState)
	_adsState, _adsStateErr := readBuffer.ReadUint16("adsState", 16)
	if _adsStateErr != nil {
		return nil, errors.Wrap(_adsStateErr, "Error parsing 'adsState' field")
	}
	adsState := _adsState

	// Simple Field (deviceState)
	_deviceState, _deviceStateErr := readBuffer.ReadUint16("deviceState", 16)
	if _deviceStateErr != nil {
		return nil, errors.Wrap(_deviceStateErr, "Error parsing 'deviceState' field")
	}
	deviceState := _deviceState

	// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
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

	if closeErr := readBuffer.CloseContext("AdsWriteControlRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &AdsWriteControlRequest{
		AdsState:    adsState,
		DeviceState: deviceState,
		Data:        data,
		AdsData:     &AdsData{},
	}
	_child.AdsData.Child = _child
	return _child.AdsData, nil
}

func (m *AdsWriteControlRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsWriteControlRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (adsState)
		adsState := uint16(m.AdsState)
		_adsStateErr := writeBuffer.WriteUint16("adsState", 16, (adsState))
		if _adsStateErr != nil {
			return errors.Wrap(_adsStateErr, "Error serializing 'adsState' field")
		}

		// Simple Field (deviceState)
		deviceState := uint16(m.DeviceState)
		_deviceStateErr := writeBuffer.WriteUint16("deviceState", 16, (deviceState))
		if _deviceStateErr != nil {
			return errors.Wrap(_deviceStateErr, "Error serializing 'deviceState' field")
		}

		// Implicit Field (length) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		length := uint32(uint32(len(m.Data)))
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

		if popErr := writeBuffer.PopContext("AdsWriteControlRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *AdsWriteControlRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
