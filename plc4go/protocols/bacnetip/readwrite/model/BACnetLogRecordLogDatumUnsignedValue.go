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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetLogRecordLogDatumUnsignedValue is the corresponding interface of BACnetLogRecordLogDatumUnsignedValue
type BACnetLogRecordLogDatumUnsignedValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetLogRecordLogDatum
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetContextTagUnsignedInteger
}

// BACnetLogRecordLogDatumUnsignedValueExactly can be used when we want exactly this type and not a type which fulfills BACnetLogRecordLogDatumUnsignedValue.
// This is useful for switch cases.
type BACnetLogRecordLogDatumUnsignedValueExactly interface {
	BACnetLogRecordLogDatumUnsignedValue
	isBACnetLogRecordLogDatumUnsignedValue() bool
}

// _BACnetLogRecordLogDatumUnsignedValue is the data-structure of this message
type _BACnetLogRecordLogDatumUnsignedValue struct {
	*_BACnetLogRecordLogDatum
        UnsignedValue BACnetContextTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogRecordLogDatumUnsignedValue) InitializeParent(parent BACnetLogRecordLogDatum , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetLogRecordLogDatumUnsignedValue)  GetParent() BACnetLogRecordLogDatum {
	return m._BACnetLogRecordLogDatum
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecordLogDatumUnsignedValue) GetUnsignedValue() BACnetContextTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetLogRecordLogDatumUnsignedValue factory function for _BACnetLogRecordLogDatumUnsignedValue
func NewBACnetLogRecordLogDatumUnsignedValue( unsignedValue BACnetContextTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 ) *_BACnetLogRecordLogDatumUnsignedValue {
	_result := &_BACnetLogRecordLogDatumUnsignedValue{
		UnsignedValue: unsignedValue,
    	_BACnetLogRecordLogDatum: NewBACnetLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogRecordLogDatumUnsignedValue(structType interface{}) BACnetLogRecordLogDatumUnsignedValue {
    if casted, ok := structType.(BACnetLogRecordLogDatumUnsignedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecordLogDatumUnsignedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecordLogDatumUnsignedValue) GetTypeName() string {
	return "BACnetLogRecordLogDatumUnsignedValue"
}

func (m *_BACnetLogRecordLogDatumUnsignedValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLogRecordLogDatumUnsignedValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetLogRecordLogDatumUnsignedValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogRecordLogDatumUnsignedValueParse(theBytes []byte, tagNumber uint8) (BACnetLogRecordLogDatumUnsignedValue, error) {
	return BACnetLogRecordLogDatumUnsignedValueParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetLogRecordLogDatumUnsignedValueParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLogRecordLogDatumUnsignedValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogRecordLogDatumUnsignedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecordLogDatumUnsignedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unsignedValue)
	if pullErr := readBuffer.PullContext("unsignedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unsignedValue")
	}
_unsignedValue, _unsignedValueErr := BACnetContextTagParseWithBuffer(readBuffer , uint8( uint8(4) ) , BACnetDataType( BACnetDataType_UNSIGNED_INTEGER ) )
	if _unsignedValueErr != nil {
		return nil, errors.Wrap(_unsignedValueErr, "Error parsing 'unsignedValue' field of BACnetLogRecordLogDatumUnsignedValue")
	}
	unsignedValue := _unsignedValue.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("unsignedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unsignedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogRecordLogDatumUnsignedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecordLogDatumUnsignedValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogRecordLogDatumUnsignedValue{
		_BACnetLogRecordLogDatum: &_BACnetLogRecordLogDatum{
			TagNumber: tagNumber,
		},
		UnsignedValue: unsignedValue,
	}
	_child._BACnetLogRecordLogDatum._BACnetLogRecordLogDatumChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogRecordLogDatumUnsignedValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecordLogDatumUnsignedValue) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogRecordLogDatumUnsignedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogRecordLogDatumUnsignedValue")
		}

	// Simple Field (unsignedValue)
	if pushErr := writeBuffer.PushContext("unsignedValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for unsignedValue")
	}
	_unsignedValueErr := writeBuffer.WriteSerializable(m.GetUnsignedValue())
	if popErr := writeBuffer.PopContext("unsignedValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for unsignedValue")
	}
	if _unsignedValueErr != nil {
		return errors.Wrap(_unsignedValueErr, "Error serializing 'unsignedValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetLogRecordLogDatumUnsignedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogRecordLogDatumUnsignedValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetLogRecordLogDatumUnsignedValue) isBACnetLogRecordLogDatumUnsignedValue() bool {
	return true
}

func (m *_BACnetLogRecordLogDatumUnsignedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



