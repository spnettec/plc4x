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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataLoggingRecord is the corresponding interface of BACnetConstructedDataLoggingRecord
type BACnetConstructedDataLoggingRecord interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLoggingRecord returns LoggingRecord (property field)
	GetLoggingRecord() BACnetAccumulatorRecord
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccumulatorRecord
}

// BACnetConstructedDataLoggingRecordExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLoggingRecord.
// This is useful for switch cases.
type BACnetConstructedDataLoggingRecordExactly interface {
	BACnetConstructedDataLoggingRecord
	isBACnetConstructedDataLoggingRecord() bool
}

// _BACnetConstructedDataLoggingRecord is the data-structure of this message
type _BACnetConstructedDataLoggingRecord struct {
	*_BACnetConstructedData
	LoggingRecord BACnetAccumulatorRecord
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLoggingRecord) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLoggingRecord) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOGGING_RECORD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLoggingRecord) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLoggingRecord) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLoggingRecord) GetLoggingRecord() BACnetAccumulatorRecord {
	return m.LoggingRecord
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLoggingRecord) GetActualValue() BACnetAccumulatorRecord {
	return CastBACnetAccumulatorRecord(m.GetLoggingRecord())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLoggingRecord factory function for _BACnetConstructedDataLoggingRecord
func NewBACnetConstructedDataLoggingRecord(loggingRecord BACnetAccumulatorRecord, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLoggingRecord {
	_result := &_BACnetConstructedDataLoggingRecord{
		LoggingRecord:          loggingRecord,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLoggingRecord(structType interface{}) BACnetConstructedDataLoggingRecord {
	if casted, ok := structType.(BACnetConstructedDataLoggingRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLoggingRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLoggingRecord) GetTypeName() string {
	return "BACnetConstructedDataLoggingRecord"
}

func (m *_BACnetConstructedDataLoggingRecord) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLoggingRecord) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (loggingRecord)
	lengthInBits += m.LoggingRecord.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLoggingRecord) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLoggingRecordParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLoggingRecord, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLoggingRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLoggingRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (loggingRecord)
	if pullErr := readBuffer.PullContext("loggingRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for loggingRecord")
	}
	_loggingRecord, _loggingRecordErr := BACnetAccumulatorRecordParse(readBuffer)
	if _loggingRecordErr != nil {
		return nil, errors.Wrap(_loggingRecordErr, "Error parsing 'loggingRecord' field")
	}
	loggingRecord := _loggingRecord.(BACnetAccumulatorRecord)
	if closeErr := readBuffer.CloseContext("loggingRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for loggingRecord")
	}

	// Virtual field
	_actualValue := loggingRecord
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLoggingRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLoggingRecord")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLoggingRecord{
		LoggingRecord: loggingRecord,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLoggingRecord) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLoggingRecord"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLoggingRecord")
		}

		// Simple Field (loggingRecord)
		if pushErr := writeBuffer.PushContext("loggingRecord"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for loggingRecord")
		}
		_loggingRecordErr := writeBuffer.WriteSerializable(m.GetLoggingRecord())
		if popErr := writeBuffer.PopContext("loggingRecord"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for loggingRecord")
		}
		if _loggingRecordErr != nil {
			return errors.Wrap(_loggingRecordErr, "Error serializing 'loggingRecord' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLoggingRecord"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLoggingRecord")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLoggingRecord) isBACnetConstructedDataLoggingRecord() bool {
	return true
}

func (m *_BACnetConstructedDataLoggingRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
