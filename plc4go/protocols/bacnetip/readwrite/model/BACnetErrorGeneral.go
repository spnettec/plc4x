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


// BACnetErrorGeneral is the corresponding interface of BACnetErrorGeneral
type BACnetErrorGeneral interface {
	utils.LengthAware
	utils.Serializable
	BACnetError
	// GetError returns Error (property field)
	GetError() Error
}

// BACnetErrorGeneralExactly can be used when we want exactly this type and not a type which fulfills BACnetErrorGeneral.
// This is useful for switch cases.
type BACnetErrorGeneralExactly interface {
	BACnetErrorGeneral
	isBACnetErrorGeneral() bool
}

// _BACnetErrorGeneral is the data-structure of this message
type _BACnetErrorGeneral struct {
	*_BACnetError
        Error Error
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetErrorGeneral)  GetErrorChoice() BACnetConfirmedServiceChoice {
return 0}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetErrorGeneral) InitializeParent(parent BACnetError ) {}

func (m *_BACnetErrorGeneral)  GetParent() BACnetError {
	return m._BACnetError
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetErrorGeneral) GetError() Error {
	return m.Error
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetErrorGeneral factory function for _BACnetErrorGeneral
func NewBACnetErrorGeneral( error Error ) *_BACnetErrorGeneral {
	_result := &_BACnetErrorGeneral{
		Error: error,
    	_BACnetError: NewBACnetError(),
	}
	_result._BACnetError._BACnetErrorChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetErrorGeneral(structType interface{}) BACnetErrorGeneral {
    if casted, ok := structType.(BACnetErrorGeneral); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetErrorGeneral); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetErrorGeneral) GetTypeName() string {
	return "BACnetErrorGeneral"
}

func (m *_BACnetErrorGeneral) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetErrorGeneral) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (error)
	lengthInBits += m.Error.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetErrorGeneral) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetErrorGeneralParse(readBuffer utils.ReadBuffer, errorChoice BACnetConfirmedServiceChoice) (BACnetErrorGeneral, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetErrorGeneral"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetErrorGeneral")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (error)
	if pullErr := readBuffer.PullContext("error"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for error")
	}
_error, _errorErr := ErrorParse(readBuffer)
	if _errorErr != nil {
		return nil, errors.Wrap(_errorErr, "Error parsing 'error' field of BACnetErrorGeneral")
	}
	error := _error.(Error)
	if closeErr := readBuffer.CloseContext("error"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for error")
	}

	if closeErr := readBuffer.CloseContext("BACnetErrorGeneral"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetErrorGeneral")
	}

	// Create a partially initialized instance
	_child := &_BACnetErrorGeneral{
		_BACnetError: &_BACnetError{
		},
		Error: error,
	}
	_child._BACnetError._BACnetErrorChildRequirements = _child
	return _child, nil
}

func (m *_BACnetErrorGeneral) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetErrorGeneral) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetErrorGeneral"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetErrorGeneral")
		}

	// Simple Field (error)
	if pushErr := writeBuffer.PushContext("error"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for error")
	}
	_errorErr := writeBuffer.WriteSerializable(m.GetError())
	if popErr := writeBuffer.PopContext("error"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for error")
	}
	if _errorErr != nil {
		return errors.Wrap(_errorErr, "Error serializing 'error' field")
	}

		if popErr := writeBuffer.PopContext("BACnetErrorGeneral"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetErrorGeneral")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetErrorGeneral) isBACnetErrorGeneral() bool {
	return true
}

func (m *_BACnetErrorGeneral) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



