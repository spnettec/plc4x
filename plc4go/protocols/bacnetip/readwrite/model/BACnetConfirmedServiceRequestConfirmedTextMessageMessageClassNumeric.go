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


// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric is the corresponding interface of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
	// GetNumericValue returns NumericValue (property field)
	GetNumericValue() BACnetContextTagUnsignedInteger
}

// BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericExactly interface {
	BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric
	isBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric() bool
}

// _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric is the data-structure of this message
type _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric struct {
	*_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
        NumericValue BACnetContextTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) InitializeParent(parent BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric)  GetParent() BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass {
	return m._BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) GetNumericValue() BACnetContextTagUnsignedInteger {
	return m.NumericValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric factory function for _BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric
func NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric( numericValue BACnetContextTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 ) *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric {
	_result := &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric{
		NumericValue: numericValue,
    	_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass: NewBACnetConfirmedServiceRequestConfirmedTextMessageMessageClass(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass._BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric(structType interface{}) BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric {
    if casted, ok := structType.(BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) GetTypeName() string {
	return "BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric"
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (numericValue)
	lengthInBits += m.NumericValue.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumericParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (numericValue)
	if pullErr := readBuffer.PullContext("numericValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for numericValue")
	}
_numericValue, _numericValueErr := BACnetContextTagParse(readBuffer , uint8( uint8(0) ) , BACnetDataType( BACnetDataType_UNSIGNED_INTEGER ) )
	if _numericValueErr != nil {
		return nil, errors.Wrap(_numericValueErr, "Error parsing 'numericValue' field of BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric")
	}
	numericValue := _numericValue.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("numericValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for numericValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric{
		_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass: &_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass{
			TagNumber: tagNumber,
		},
		NumericValue: numericValue,
	}
	_child._BACnetConfirmedServiceRequestConfirmedTextMessageMessageClass._BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric")
		}

	// Simple Field (numericValue)
	if pushErr := writeBuffer.PushContext("numericValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for numericValue")
	}
	_numericValueErr := writeBuffer.WriteSerializable(m.GetNumericValue())
	if popErr := writeBuffer.PopContext("numericValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for numericValue")
	}
	if _numericValueErr != nil {
		return errors.Wrap(_numericValueErr, "Error serializing 'numericValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) isBACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestConfirmedTextMessageMessageClassNumeric) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



