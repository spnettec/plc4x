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

// BACnetEventLogRecordLogDatumTimeChange is the data-structure of this message
type BACnetEventLogRecordLogDatumTimeChange struct {
	*BACnetEventLogRecordLogDatum
	TimeChange *BACnetContextTagReal

	// Arguments.
	TagNumber uint8
}

// IBACnetEventLogRecordLogDatumTimeChange is the corresponding interface of BACnetEventLogRecordLogDatumTimeChange
type IBACnetEventLogRecordLogDatumTimeChange interface {
	IBACnetEventLogRecordLogDatum
	// GetTimeChange returns TimeChange (property field)
	GetTimeChange() *BACnetContextTagReal
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetEventLogRecordLogDatumTimeChange) InitializeParent(parent *BACnetEventLogRecordLogDatum, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetEventLogRecordLogDatum.OpeningTag = openingTag
	m.BACnetEventLogRecordLogDatum.PeekedTagHeader = peekedTagHeader
	m.BACnetEventLogRecordLogDatum.ClosingTag = closingTag
}

func (m *BACnetEventLogRecordLogDatumTimeChange) GetParent() *BACnetEventLogRecordLogDatum {
	return m.BACnetEventLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetEventLogRecordLogDatumTimeChange) GetTimeChange() *BACnetContextTagReal {
	return m.TimeChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventLogRecordLogDatumTimeChange factory function for BACnetEventLogRecordLogDatumTimeChange
func NewBACnetEventLogRecordLogDatumTimeChange(timeChange *BACnetContextTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetEventLogRecordLogDatumTimeChange {
	_result := &BACnetEventLogRecordLogDatumTimeChange{
		TimeChange:                   timeChange,
		BACnetEventLogRecordLogDatum: NewBACnetEventLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetEventLogRecordLogDatumTimeChange(structType interface{}) *BACnetEventLogRecordLogDatumTimeChange {
	if casted, ok := structType.(BACnetEventLogRecordLogDatumTimeChange); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventLogRecordLogDatumTimeChange); ok {
		return casted
	}
	if casted, ok := structType.(BACnetEventLogRecordLogDatum); ok {
		return CastBACnetEventLogRecordLogDatumTimeChange(casted.Child)
	}
	if casted, ok := structType.(*BACnetEventLogRecordLogDatum); ok {
		return CastBACnetEventLogRecordLogDatumTimeChange(casted.Child)
	}
	return nil
}

func (m *BACnetEventLogRecordLogDatumTimeChange) GetTypeName() string {
	return "BACnetEventLogRecordLogDatumTimeChange"
}

func (m *BACnetEventLogRecordLogDatumTimeChange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventLogRecordLogDatumTimeChange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (timeChange)
	lengthInBits += m.TimeChange.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetEventLogRecordLogDatumTimeChange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventLogRecordLogDatumTimeChangeParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetEventLogRecordLogDatumTimeChange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventLogRecordLogDatumTimeChange"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeChange)
	if pullErr := readBuffer.PullContext("timeChange"); pullErr != nil {
		return nil, pullErr
	}
	_timeChange, _timeChangeErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_REAL))
	if _timeChangeErr != nil {
		return nil, errors.Wrap(_timeChangeErr, "Error parsing 'timeChange' field")
	}
	timeChange := CastBACnetContextTagReal(_timeChange)
	if closeErr := readBuffer.CloseContext("timeChange"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetEventLogRecordLogDatumTimeChange"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetEventLogRecordLogDatumTimeChange{
		TimeChange:                   CastBACnetContextTagReal(timeChange),
		BACnetEventLogRecordLogDatum: &BACnetEventLogRecordLogDatum{},
	}
	_child.BACnetEventLogRecordLogDatum.Child = _child
	return _child, nil
}

func (m *BACnetEventLogRecordLogDatumTimeChange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventLogRecordLogDatumTimeChange"); pushErr != nil {
			return pushErr
		}

		// Simple Field (timeChange)
		if pushErr := writeBuffer.PushContext("timeChange"); pushErr != nil {
			return pushErr
		}
		_timeChangeErr := m.TimeChange.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timeChange"); popErr != nil {
			return popErr
		}
		if _timeChangeErr != nil {
			return errors.Wrap(_timeChangeErr, "Error serializing 'timeChange' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventLogRecordLogDatumTimeChange"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetEventLogRecordLogDatumTimeChange) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
