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
type RelativeTimestamp struct {
	Timestamp uint16
}

// The corresponding interface
type IRelativeTimestamp interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

func NewRelativeTimestamp(timestamp uint16) *RelativeTimestamp {
	return &RelativeTimestamp{Timestamp: timestamp}
}

func CastRelativeTimestamp(structType interface{}) *RelativeTimestamp {
	castFunc := func(typ interface{}) *RelativeTimestamp {
		if casted, ok := typ.(RelativeTimestamp); ok {
			return &casted
		}
		if casted, ok := typ.(*RelativeTimestamp); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *RelativeTimestamp) GetTypeName() string {
	return "RelativeTimestamp"
}

func (m *RelativeTimestamp) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *RelativeTimestamp) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += 16

	return lengthInBits
}

func (m *RelativeTimestamp) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func RelativeTimestampParse(readBuffer utils.ReadBuffer) (*RelativeTimestamp, error) {
	if pullErr := readBuffer.PullContext("RelativeTimestamp"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (timestamp)
	timestamp, _timestampErr := readBuffer.ReadUint16("timestamp", 16)
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field")
	}

	if closeErr := readBuffer.CloseContext("RelativeTimestamp"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewRelativeTimestamp(timestamp), nil
}

func (m *RelativeTimestamp) Serialize(writeBuffer utils.WriteBuffer) error {
	if pushErr := writeBuffer.PushContext("RelativeTimestamp"); pushErr != nil {
		return pushErr
	}

	// Simple Field (timestamp)
	timestamp := uint16(m.Timestamp)
	_timestampErr := writeBuffer.WriteUint16("timestamp", 16, (timestamp))
	if _timestampErr != nil {
		return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
	}

	if popErr := writeBuffer.PopContext("RelativeTimestamp"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *RelativeTimestamp) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
