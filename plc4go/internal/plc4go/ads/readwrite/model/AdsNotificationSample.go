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

// AdsNotificationSample is the data-structure of this message
type AdsNotificationSample struct {
	NotificationHandle uint32
	SampleSize         uint32
	Data               []byte
}

// IAdsNotificationSample is the corresponding interface of AdsNotificationSample
type IAdsNotificationSample interface {
	// GetNotificationHandle returns NotificationHandle (property field)
	GetNotificationHandle() uint32
	// GetSampleSize returns SampleSize (property field)
	GetSampleSize() uint32
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
/////////////////////// Accessors for property fields.
///////////////////////

func (m *AdsNotificationSample) GetNotificationHandle() uint32 {
	return m.NotificationHandle
}

func (m *AdsNotificationSample) GetSampleSize() uint32 {
	return m.SampleSize
}

func (m *AdsNotificationSample) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAdsNotificationSample factory function for AdsNotificationSample
func NewAdsNotificationSample(notificationHandle uint32, sampleSize uint32, data []byte) *AdsNotificationSample {
	return &AdsNotificationSample{NotificationHandle: notificationHandle, SampleSize: sampleSize, Data: data}
}

func CastAdsNotificationSample(structType interface{}) *AdsNotificationSample {
	if casted, ok := structType.(AdsNotificationSample); ok {
		return &casted
	}
	if casted, ok := structType.(*AdsNotificationSample); ok {
		return casted
	}
	return nil
}

func (m *AdsNotificationSample) GetTypeName() string {
	return "AdsNotificationSample"
}

func (m *AdsNotificationSample) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *AdsNotificationSample) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (notificationHandle)
	lengthInBits += 32

	// Simple field (sampleSize)
	lengthInBits += 32

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *AdsNotificationSample) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsNotificationSampleParse(readBuffer utils.ReadBuffer) (*AdsNotificationSample, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsNotificationSample"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (notificationHandle)
	_notificationHandle, _notificationHandleErr := readBuffer.ReadUint32("notificationHandle", 32)
	if _notificationHandleErr != nil {
		return nil, errors.Wrap(_notificationHandleErr, "Error parsing 'notificationHandle' field")
	}
	notificationHandle := _notificationHandle

	// Simple Field (sampleSize)
	_sampleSize, _sampleSizeErr := readBuffer.ReadUint32("sampleSize", 32)
	if _sampleSizeErr != nil {
		return nil, errors.Wrap(_sampleSizeErr, "Error parsing 'sampleSize' field")
	}
	sampleSize := _sampleSize
	// Byte Array field (data)
	numberOfBytesdata := int(sampleSize)
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("AdsNotificationSample"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewAdsNotificationSample(notificationHandle, sampleSize, data), nil
}

func (m *AdsNotificationSample) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("AdsNotificationSample"); pushErr != nil {
		return pushErr
	}

	// Simple Field (notificationHandle)
	notificationHandle := uint32(m.NotificationHandle)
	_notificationHandleErr := writeBuffer.WriteUint32("notificationHandle", 32, (notificationHandle))
	if _notificationHandleErr != nil {
		return errors.Wrap(_notificationHandleErr, "Error serializing 'notificationHandle' field")
	}

	// Simple Field (sampleSize)
	sampleSize := uint32(m.SampleSize)
	_sampleSizeErr := writeBuffer.WriteUint32("sampleSize", 32, (sampleSize))
	if _sampleSizeErr != nil {
		return errors.Wrap(_sampleSizeErr, "Error serializing 'sampleSize' field")
	}

	// Array Field (data)
	if m.Data != nil {
		// Byte Array field (data)
		_writeArrayErr := writeBuffer.WriteByteArray("data", m.Data)
		if _writeArrayErr != nil {
			return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
		}
	}

	if popErr := writeBuffer.PopContext("AdsNotificationSample"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *AdsNotificationSample) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
