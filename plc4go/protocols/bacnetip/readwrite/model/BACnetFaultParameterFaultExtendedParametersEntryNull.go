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


// BACnetFaultParameterFaultExtendedParametersEntryNull is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryNull
type BACnetFaultParameterFaultExtendedParametersEntryNull interface {
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetNullValue returns NullValue (property field)
	GetNullValue() BACnetApplicationTagNull
}

// BACnetFaultParameterFaultExtendedParametersEntryNullExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultExtendedParametersEntryNull.
// This is useful for switch cases.
type BACnetFaultParameterFaultExtendedParametersEntryNullExactly interface {
	BACnetFaultParameterFaultExtendedParametersEntryNull
	isBACnetFaultParameterFaultExtendedParametersEntryNull() bool
}

// _BACnetFaultParameterFaultExtendedParametersEntryNull is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryNull struct {
	*_BACnetFaultParameterFaultExtendedParametersEntry
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

func (m *_BACnetFaultParameterFaultExtendedParametersEntryNull) InitializeParent(parent BACnetFaultParameterFaultExtendedParametersEntry , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryNull)  GetParent() BACnetFaultParameterFaultExtendedParametersEntry {
	return m._BACnetFaultParameterFaultExtendedParametersEntry
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryNull) GetNullValue() BACnetApplicationTagNull {
	return m.NullValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetFaultParameterFaultExtendedParametersEntryNull factory function for _BACnetFaultParameterFaultExtendedParametersEntryNull
func NewBACnetFaultParameterFaultExtendedParametersEntryNull( nullValue BACnetApplicationTagNull , peekedTagHeader BACnetTagHeader ) *_BACnetFaultParameterFaultExtendedParametersEntryNull {
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryNull{
		NullValue: nullValue,
    	_BACnetFaultParameterFaultExtendedParametersEntry: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
	}
	_result._BACnetFaultParameterFaultExtendedParametersEntry._BACnetFaultParameterFaultExtendedParametersEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryNull(structType interface{}) BACnetFaultParameterFaultExtendedParametersEntryNull {
    if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryNull) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryNull"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryNull) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryNull) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (nullValue)
	lengthInBits += m.NullValue.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetFaultParameterFaultExtendedParametersEntryNull) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultExtendedParametersEntryNullParse(readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultExtendedParametersEntryNull, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nullValue)
	if pullErr := readBuffer.PullContext("nullValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nullValue")
	}
_nullValue, _nullValueErr := BACnetApplicationTagParse(readBuffer)
	if _nullValueErr != nil {
		return nil, errors.Wrap(_nullValueErr, "Error parsing 'nullValue' field of BACnetFaultParameterFaultExtendedParametersEntryNull")
	}
	nullValue := _nullValue.(BACnetApplicationTagNull)
	if closeErr := readBuffer.CloseContext("nullValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nullValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryNull")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultExtendedParametersEntryNull{
		_BACnetFaultParameterFaultExtendedParametersEntry: &_BACnetFaultParameterFaultExtendedParametersEntry{
		},
		NullValue: nullValue,
	}
	_child._BACnetFaultParameterFaultExtendedParametersEntry._BACnetFaultParameterFaultExtendedParametersEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryNull) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryNull")
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

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryNull")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetFaultParameterFaultExtendedParametersEntryNull) isBACnetFaultParameterFaultExtendedParametersEntryNull() bool {
	return true
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



