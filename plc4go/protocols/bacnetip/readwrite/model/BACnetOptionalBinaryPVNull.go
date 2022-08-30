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


// BACnetOptionalBinaryPVNull is the corresponding interface of BACnetOptionalBinaryPVNull
type BACnetOptionalBinaryPVNull interface {
	utils.LengthAware
	utils.Serializable
	BACnetOptionalBinaryPV
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
}

// BACnetOptionalBinaryPVNullExactly can be used when we want exactly this type and not a type which fulfills BACnetOptionalBinaryPVNull.
// This is useful for switch cases.
type BACnetOptionalBinaryPVNullExactly interface {
	BACnetOptionalBinaryPVNull
	isBACnetOptionalBinaryPVNull() bool
}

// _BACnetOptionalBinaryPVNull is the data-structure of this message
type _BACnetOptionalBinaryPVNull struct {
	*_BACnetOptionalBinaryPV
        NullValue BACnetApplicationTagNull
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetOptionalBinaryPVNull) InitializeParent(parent BACnetOptionalBinaryPV , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetOptionalBinaryPVNull)  GetParent() BACnetOptionalBinaryPV {
	return m._BACnetOptionalBinaryPV
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetOptionalBinaryPVNull) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetOptionalBinaryPVNull factory function for _BACnetOptionalBinaryPVNull
func NewBACnetOptionalBinaryPVNull( nullValue BACnetApplicationTagNull , peekedTagHeader BACnetTagHeader ) *_BACnetOptionalBinaryPVNull {
	_result := &_BACnetOptionalBinaryPVNull{
		NullValue: nullValue,
    	_BACnetOptionalBinaryPV: NewBACnetOptionalBinaryPV(peekedTagHeader),
	}
	_result._BACnetOptionalBinaryPV._BACnetOptionalBinaryPVChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetOptionalBinaryPVNull(structType interface{}) BACnetOptionalBinaryPVNull {
    if casted, ok := structType.(BACnetOptionalBinaryPVNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetOptionalBinaryPVNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetOptionalBinaryPVNull) GetTypeName() string {
	return "BACnetOptionalBinaryPVNull"
}

func (m *_BACnetOptionalBinaryPVNull) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetOptionalBinaryPVNull) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetOptionalBinaryPVNull) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetOptionalBinaryPVNullParse(readBuffer utils.ReadBuffer) (BACnetOptionalBinaryPVNull, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetOptionalBinaryPVNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetOptionalBinaryPVNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nullValue)
	if pullErr := readBuffer.PullContext("nullValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nullValue")
	}
_nullValue, _nullValueErr := BACnetApplicationTagParse(readBuffer)
	if _nullValueErr != nil {
		return nil, errors.Wrap(_nullValueErr, "Error parsing 'nullValue' field of BACnetOptionalBinaryPVNull")
	}
	nullValue := _nullValue.(BACnetApplicationTagNull)
	if closeErr := readBuffer.CloseContext("nullValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nullValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetOptionalBinaryPVNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetOptionalBinaryPVNull")
	}

	// Create a partially initialized instance
	_child := &_BACnetOptionalBinaryPVNull{
		_BACnetOptionalBinaryPV: &_BACnetOptionalBinaryPV{
		},
		NullValue: nullValue,
	}
	_child._BACnetOptionalBinaryPV._BACnetOptionalBinaryPVChildRequirements = _child
	return _child, nil
}

func (m *_BACnetOptionalBinaryPVNull) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetOptionalBinaryPVNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetOptionalBinaryPVNull")
		}

	// Simple Field (nullValue)
	if pushErr := writeBuffer.PushContext("nullValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for nullValue")
	}
	_nullValueErr := writeBuffer.WriteSerializable(m.GetNullValue())
	if popErr := writeBuffer.PopContext("nullValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for nullValue")
	}
	if _nullValueErr != nil {
		return errors.Wrap(_nullValueErr, "Error serializing 'nullValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetOptionalBinaryPVNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetOptionalBinaryPVNull")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetOptionalBinaryPVNull) isBACnetOptionalBinaryPVNull() bool {
	return true
}

func (m *_BACnetOptionalBinaryPVNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



