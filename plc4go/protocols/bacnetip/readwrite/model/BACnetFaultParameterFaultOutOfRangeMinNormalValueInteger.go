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


// BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger
type BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger interface {
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultOutOfRangeMinNormalValue
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetApplicationTagSignedInteger
}

// BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger.
// This is useful for switch cases.
type BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerExactly interface {
	BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger
	isBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger() bool
}

// _BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger struct {
	*_BACnetFaultParameterFaultOutOfRangeMinNormalValue
        IntegerValue BACnetApplicationTagSignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) InitializeParent(parent BACnetFaultParameterFaultOutOfRangeMinNormalValue , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger)  GetParent() BACnetFaultParameterFaultOutOfRangeMinNormalValue {
	return m._BACnetFaultParameterFaultOutOfRangeMinNormalValue
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetIntegerValue() BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger factory function for _BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger
func NewBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger( integerValue BACnetApplicationTagSignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 ) *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger {
	_result := &_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger{
		IntegerValue: integerValue,
    	_BACnetFaultParameterFaultOutOfRangeMinNormalValue: NewBACnetFaultParameterFaultOutOfRangeMinNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetFaultParameterFaultOutOfRangeMinNormalValue._BACnetFaultParameterFaultOutOfRangeMinNormalValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger(structType interface{}) BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger {
    if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerParse(theBytes []byte, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger, error) {
	return BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetFaultParameterFaultOutOfRangeMinNormalValueIntegerParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (integerValue)
	if pullErr := readBuffer.PullContext("integerValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for integerValue")
	}
_integerValue, _integerValueErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _integerValueErr != nil {
		return nil, errors.Wrap(_integerValueErr, "Error parsing 'integerValue' field of BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
	}
	integerValue := _integerValue.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("integerValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for integerValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger{
		_BACnetFaultParameterFaultOutOfRangeMinNormalValue: &_BACnetFaultParameterFaultOutOfRangeMinNormalValue{
			TagNumber: tagNumber,
		},
		IntegerValue: integerValue,
	}
	_child._BACnetFaultParameterFaultOutOfRangeMinNormalValue._BACnetFaultParameterFaultOutOfRangeMinNormalValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
		}

	// Simple Field (integerValue)
	if pushErr := writeBuffer.PushContext("integerValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for integerValue")
	}
	_integerValueErr := writeBuffer.WriteSerializable(m.GetIntegerValue())
	if popErr := writeBuffer.PopContext("integerValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for integerValue")
	}
	if _integerValueErr != nil {
		return errors.Wrap(_integerValueErr, "Error serializing 'integerValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) isBACnetFaultParameterFaultOutOfRangeMinNormalValueInteger() bool {
	return true
}

func (m *_BACnetFaultParameterFaultOutOfRangeMinNormalValueInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



