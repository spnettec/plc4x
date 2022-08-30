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


// BACnetOptionalUnsignedValue is the corresponding interface of BACnetOptionalUnsignedValue
type BACnetOptionalUnsignedValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetOptionalUnsigned
	// GetUnsignedValue returns UnsignedValue (property field)
	GetUnsignedValue() BACnetApplicationTagUnsignedInteger
}

// BACnetOptionalUnsignedValueExactly can be used when we want exactly this type and not a type which fulfills BACnetOptionalUnsignedValue.
// This is useful for switch cases.
type BACnetOptionalUnsignedValueExactly interface {
	BACnetOptionalUnsignedValue
	isBACnetOptionalUnsignedValue() bool
}

// _BACnetOptionalUnsignedValue is the data-structure of this message
type _BACnetOptionalUnsignedValue struct {
	*_BACnetOptionalUnsigned
        UnsignedValue BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetOptionalUnsignedValue) InitializeParent(parent BACnetOptionalUnsigned , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetOptionalUnsignedValue)  GetParent() BACnetOptionalUnsigned {
	return m._BACnetOptionalUnsigned
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalUnsignedValue) GetUnsignedValue() BACnetApplicationTagUnsignedInteger {
	return m.UnsignedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetOptionalUnsignedValue factory function for _BACnetOptionalUnsignedValue
func NewBACnetOptionalUnsignedValue( unsignedValue BACnetApplicationTagUnsignedInteger , peekedTagHeader BACnetTagHeader ) *_BACnetOptionalUnsignedValue {
	_result := &_BACnetOptionalUnsignedValue{
		UnsignedValue: unsignedValue,
    	_BACnetOptionalUnsigned: NewBACnetOptionalUnsigned(peekedTagHeader),
	}
	_result._BACnetOptionalUnsigned._BACnetOptionalUnsignedChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetOptionalUnsignedValue(structType interface{}) BACnetOptionalUnsignedValue {
    if casted, ok := structType.(BACnetOptionalUnsignedValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalUnsignedValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalUnsignedValue) GetTypeName() string {
	return "BACnetOptionalUnsignedValue"
}

func (m *_BACnetOptionalUnsignedValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetOptionalUnsignedValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (unsignedValue)
	lengthInBits += m.UnsignedValue.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetOptionalUnsignedValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetOptionalUnsignedValueParse(readBuffer utils.ReadBuffer) (BACnetOptionalUnsignedValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalUnsignedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalUnsignedValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (unsignedValue)
	if pullErr := readBuffer.PullContext("unsignedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for unsignedValue")
	}
_unsignedValue, _unsignedValueErr := BACnetApplicationTagParse(readBuffer)
	if _unsignedValueErr != nil {
		return nil, errors.Wrap(_unsignedValueErr, "Error parsing 'unsignedValue' field of BACnetOptionalUnsignedValue")
	}
	unsignedValue := _unsignedValue.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("unsignedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for unsignedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetOptionalUnsignedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalUnsignedValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetOptionalUnsignedValue{
		_BACnetOptionalUnsigned: &_BACnetOptionalUnsigned{
		},
		UnsignedValue: unsignedValue,
	}
	_child._BACnetOptionalUnsigned._BACnetOptionalUnsignedChildRequirements = _child
	return _child, nil
}

func (m *_BACnetOptionalUnsignedValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalUnsignedValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetOptionalUnsignedValue")
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

		if popErr := writeBuffer.PopContext("BACnetOptionalUnsignedValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetOptionalUnsignedValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetOptionalUnsignedValue) isBACnetOptionalUnsignedValue() bool {
	return true
}

func (m *_BACnetOptionalUnsignedValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



