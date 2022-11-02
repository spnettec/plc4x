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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetTimerStateChangeValueInteger is the corresponding interface of BACnetTimerStateChangeValueInteger
type BACnetTimerStateChangeValueInteger interface {
	utils.LengthAware
	utils.Serializable
	BACnetTimerStateChangeValue
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetApplicationTagSignedInteger
}

// BACnetTimerStateChangeValueIntegerExactly can be used when we want exactly this type and not a type which fulfills BACnetTimerStateChangeValueInteger.
// This is useful for switch cases.
type BACnetTimerStateChangeValueIntegerExactly interface {
	BACnetTimerStateChangeValueInteger
	isBACnetTimerStateChangeValueInteger() bool
}

// _BACnetTimerStateChangeValueInteger is the data-structure of this message
type _BACnetTimerStateChangeValueInteger struct {
	*_BACnetTimerStateChangeValue
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

func (m *_BACnetTimerStateChangeValueInteger) InitializeParent(parent BACnetTimerStateChangeValue , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetTimerStateChangeValueInteger)  GetParent() BACnetTimerStateChangeValue {
	return m._BACnetTimerStateChangeValue
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerStateChangeValueInteger) GetIntegerValue() BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetTimerStateChangeValueInteger factory function for _BACnetTimerStateChangeValueInteger
func NewBACnetTimerStateChangeValueInteger( integerValue BACnetApplicationTagSignedInteger , peekedTagHeader BACnetTagHeader , objectTypeArgument BACnetObjectType ) *_BACnetTimerStateChangeValueInteger {
	_result := &_BACnetTimerStateChangeValueInteger{
		IntegerValue: integerValue,
    	_BACnetTimerStateChangeValue: NewBACnetTimerStateChangeValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetTimerStateChangeValueInteger(structType interface{}) BACnetTimerStateChangeValueInteger {
    if casted, ok := structType.(BACnetTimerStateChangeValueInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerStateChangeValueInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerStateChangeValueInteger) GetTypeName() string {
	return "BACnetTimerStateChangeValueInteger"
}

func (m *_BACnetTimerStateChangeValueInteger) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetTimerStateChangeValueInteger) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetTimerStateChangeValueInteger) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTimerStateChangeValueIntegerParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetTimerStateChangeValueInteger, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerStateChangeValueInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerStateChangeValueInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (integerValue)
	if pullErr := readBuffer.PullContext("integerValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for integerValue")
	}
_integerValue, _integerValueErr := BACnetApplicationTagParse(readBuffer)
	if _integerValueErr != nil {
		return nil, errors.Wrap(_integerValueErr, "Error parsing 'integerValue' field of BACnetTimerStateChangeValueInteger")
	}
	integerValue := _integerValue.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("integerValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for integerValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetTimerStateChangeValueInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerStateChangeValueInteger")
	}

	// Create a partially initialized instance
	_child := &_BACnetTimerStateChangeValueInteger{
		_BACnetTimerStateChangeValue: &_BACnetTimerStateChangeValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		IntegerValue: integerValue,
	}
	_child._BACnetTimerStateChangeValue._BACnetTimerStateChangeValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetTimerStateChangeValueInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerStateChangeValueInteger) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetTimerStateChangeValueInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetTimerStateChangeValueInteger")
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

		if popErr := writeBuffer.PopContext("BACnetTimerStateChangeValueInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetTimerStateChangeValueInteger")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetTimerStateChangeValueInteger) isBACnetTimerStateChangeValueInteger() bool {
	return true
}

func (m *_BACnetTimerStateChangeValueInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



