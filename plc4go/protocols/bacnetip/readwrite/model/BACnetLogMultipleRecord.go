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

// BACnetLogMultipleRecord is the data-structure of this message
type BACnetLogMultipleRecord struct {
	Timestamp *BACnetDateTimeEnclosed
	LogData   *BACnetLogData
}

// IBACnetLogMultipleRecord is the corresponding interface of BACnetLogMultipleRecord
type IBACnetLogMultipleRecord interface {
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() *BACnetDateTimeEnclosed
	// GetLogData returns LogData (property field)
	GetLogData() *BACnetLogData
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

func (m *BACnetLogMultipleRecord) GetTimestamp() *BACnetDateTimeEnclosed {
	return m.Timestamp
}

func (m *BACnetLogMultipleRecord) GetLogData() *BACnetLogData {
	return m.LogData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLogMultipleRecord factory function for BACnetLogMultipleRecord
func NewBACnetLogMultipleRecord(timestamp *BACnetDateTimeEnclosed, logData *BACnetLogData) *BACnetLogMultipleRecord {
	return &BACnetLogMultipleRecord{Timestamp: timestamp, LogData: logData}
}

func CastBACnetLogMultipleRecord(structType interface{}) *BACnetLogMultipleRecord {
	if casted, ok := structType.(BACnetLogMultipleRecord); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetLogMultipleRecord); ok {
		return casted
	}
	return nil
}

func (m *BACnetLogMultipleRecord) GetTypeName() string {
	return "BACnetLogMultipleRecord"
}

func (m *BACnetLogMultipleRecord) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetLogMultipleRecord) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits()

	// Simple field (logData)
	lengthInBits += m.LogData.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetLogMultipleRecord) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogMultipleRecordParse(readBuffer utils.ReadBuffer) (*BACnetLogMultipleRecord, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogMultipleRecord"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timestamp)
	if pullErr := readBuffer.PullContext("timestamp"); pullErr != nil {
		return nil, pullErr
	}
	_timestamp, _timestampErr := BACnetDateTimeEnclosedParse(readBuffer, uint8(uint8(0)))
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field")
	}
	timestamp := CastBACnetDateTimeEnclosed(_timestamp)
	if closeErr := readBuffer.CloseContext("timestamp"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (logData)
	if pullErr := readBuffer.PullContext("logData"); pullErr != nil {
		return nil, pullErr
	}
	_logData, _logDataErr := BACnetLogDataParse(readBuffer, uint8(uint8(1)))
	if _logDataErr != nil {
		return nil, errors.Wrap(_logDataErr, "Error parsing 'logData' field")
	}
	logData := CastBACnetLogData(_logData)
	if closeErr := readBuffer.CloseContext("logData"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetLogMultipleRecord"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetLogMultipleRecord(timestamp, logData), nil
}

func (m *BACnetLogMultipleRecord) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLogMultipleRecord"); pushErr != nil {
		return pushErr
	}

	// Simple Field (timestamp)
	if pushErr := writeBuffer.PushContext("timestamp"); pushErr != nil {
		return pushErr
	}
	_timestampErr := m.Timestamp.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("timestamp"); popErr != nil {
		return popErr
	}
	if _timestampErr != nil {
		return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
	}

	// Simple Field (logData)
	if pushErr := writeBuffer.PushContext("logData"); pushErr != nil {
		return pushErr
	}
	_logDataErr := m.LogData.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("logData"); popErr != nil {
		return popErr
	}
	if _logDataErr != nil {
		return errors.Wrap(_logDataErr, "Error serializing 'logData' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLogMultipleRecord"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetLogMultipleRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
