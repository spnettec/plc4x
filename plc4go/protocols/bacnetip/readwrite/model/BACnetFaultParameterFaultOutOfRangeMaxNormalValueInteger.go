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


// BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger interface {
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetApplicationTagSignedInteger
}

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueIntegerExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger.
// This is useful for switch cases.
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueIntegerExactly interface {
	BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger
	isBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger() bool
}

// _BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger struct {
	*_BACnetFaultParameterFaultOutOfRangeMaxNormalValue
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

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) InitializeParent(parent BACnetFaultParameterFaultOutOfRangeMaxNormalValue , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger)  GetParent() BACnetFaultParameterFaultOutOfRangeMaxNormalValue {
	return m._BACnetFaultParameterFaultOutOfRangeMaxNormalValue
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) GetIntegerValue() BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger factory function for _BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger
func NewBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger( integerValue BACnetApplicationTagSignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 ) *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger {
	_result := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger{
		IntegerValue: integerValue,
    	_BACnetFaultParameterFaultOutOfRangeMaxNormalValue: NewBACnetFaultParameterFaultOutOfRangeMaxNormalValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetFaultParameterFaultOutOfRangeMaxNormalValue._BACnetFaultParameterFaultOutOfRangeMaxNormalValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger(structType interface{}) BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger {
    if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger"
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultOutOfRangeMaxNormalValueIntegerParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (integerValue)
	if pullErr := readBuffer.PullContext("integerValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for integerValue")
	}
_integerValue, _integerValueErr := BACnetApplicationTagParse(readBuffer)
	if _integerValueErr != nil {
		return nil, errors.Wrap(_integerValueErr, "Error parsing 'integerValue' field of BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger")
	}
	integerValue := _integerValue.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("integerValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for integerValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger{
		_BACnetFaultParameterFaultOutOfRangeMaxNormalValue: &_BACnetFaultParameterFaultOutOfRangeMaxNormalValue{
			TagNumber: tagNumber,
		},
		IntegerValue: integerValue,
	}
	_child._BACnetFaultParameterFaultOutOfRangeMaxNormalValue._BACnetFaultParameterFaultOutOfRangeMaxNormalValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger")
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

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) isBACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger() bool {
	return true
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



