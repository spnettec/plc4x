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


// ChangeListRemoveError is the corresponding interface of ChangeListRemoveError
type ChangeListRemoveError interface {
	utils.LengthAware
	utils.Serializable
	BACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetFirstFailedElementNumber returns FirstFailedElementNumber (property field)
	GetFirstFailedElementNumber() BACnetContextTagUnsignedInteger
}

// ChangeListRemoveErrorExactly can be used when we want exactly this type and not a type which fulfills ChangeListRemoveError.
// This is useful for switch cases.
type ChangeListRemoveErrorExactly interface {
	ChangeListRemoveError
	isChangeListRemoveError() bool
}

// _ChangeListRemoveError is the data-structure of this message
type _ChangeListRemoveError struct {
	*_BACnetError
        ErrorType ErrorEnclosed
        FirstFailedElementNumber BACnetContextTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ChangeListRemoveError)  GetErrorChoice() BACnetConfirmedServiceChoice {
return BACnetConfirmedServiceChoice_REMOVE_LIST_ELEMENT}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ChangeListRemoveError) InitializeParent(parent BACnetError ) {}

func (m *_ChangeListRemoveError)  GetParent() BACnetError {
	return m._BACnetError
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ChangeListRemoveError) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_ChangeListRemoveError) GetFirstFailedElementNumber() BACnetContextTagUnsignedInteger {
	return m.FirstFailedElementNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewChangeListRemoveError factory function for _ChangeListRemoveError
func NewChangeListRemoveError( errorType ErrorEnclosed , firstFailedElementNumber BACnetContextTagUnsignedInteger ) *_ChangeListRemoveError {
	_result := &_ChangeListRemoveError{
		ErrorType: errorType,
		FirstFailedElementNumber: firstFailedElementNumber,
    	_BACnetError: NewBACnetError(),
	}
	_result._BACnetError._BACnetErrorChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastChangeListRemoveError(structType interface{}) ChangeListRemoveError {
    if casted, ok := structType.(ChangeListRemoveError); ok {
		return casted
	}
	if casted, ok := structType.(*ChangeListRemoveError); ok {
		return *casted
	}
	return nil
}

func (m *_ChangeListRemoveError) GetTypeName() string {
	return "ChangeListRemoveError"
}

func (m *_ChangeListRemoveError) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ChangeListRemoveError) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits()

	// Simple field (firstFailedElementNumber)
	lengthInBits += m.FirstFailedElementNumber.GetLengthInBits()

	return lengthInBits
}


func (m *_ChangeListRemoveError) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ChangeListRemoveErrorParse(readBuffer utils.ReadBuffer, errorChoice BACnetConfirmedServiceChoice) (ChangeListRemoveError, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ChangeListRemoveError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ChangeListRemoveError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (errorType)
	if pullErr := readBuffer.PullContext("errorType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for errorType")
	}
_errorType, _errorTypeErr := ErrorEnclosedParse(readBuffer , uint8( uint8(0) ) )
	if _errorTypeErr != nil {
		return nil, errors.Wrap(_errorTypeErr, "Error parsing 'errorType' field of ChangeListRemoveError")
	}
	errorType := _errorType.(ErrorEnclosed)
	if closeErr := readBuffer.CloseContext("errorType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for errorType")
	}

	// Simple Field (firstFailedElementNumber)
	if pullErr := readBuffer.PullContext("firstFailedElementNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for firstFailedElementNumber")
	}
_firstFailedElementNumber, _firstFailedElementNumberErr := BACnetContextTagParse(readBuffer , uint8( uint8(1) ) , BACnetDataType( BACnetDataType_UNSIGNED_INTEGER ) )
	if _firstFailedElementNumberErr != nil {
		return nil, errors.Wrap(_firstFailedElementNumberErr, "Error parsing 'firstFailedElementNumber' field of ChangeListRemoveError")
	}
	firstFailedElementNumber := _firstFailedElementNumber.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("firstFailedElementNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for firstFailedElementNumber")
	}

	if closeErr := readBuffer.CloseContext("ChangeListRemoveError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ChangeListRemoveError")
	}

	// Create a partially initialized instance
	_child := &_ChangeListRemoveError{
		_BACnetError: &_BACnetError{
		},
		ErrorType: errorType,
		FirstFailedElementNumber: firstFailedElementNumber,
	}
	_child._BACnetError._BACnetErrorChildRequirements = _child
	return _child, nil
}

func (m *_ChangeListRemoveError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ChangeListRemoveError) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ChangeListRemoveError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ChangeListRemoveError")
		}

	// Simple Field (errorType)
	if pushErr := writeBuffer.PushContext("errorType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for errorType")
	}
	_errorTypeErr := writeBuffer.WriteSerializable(m.GetErrorType())
	if popErr := writeBuffer.PopContext("errorType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for errorType")
	}
	if _errorTypeErr != nil {
		return errors.Wrap(_errorTypeErr, "Error serializing 'errorType' field")
	}

	// Simple Field (firstFailedElementNumber)
	if pushErr := writeBuffer.PushContext("firstFailedElementNumber"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for firstFailedElementNumber")
	}
	_firstFailedElementNumberErr := writeBuffer.WriteSerializable(m.GetFirstFailedElementNumber())
	if popErr := writeBuffer.PopContext("firstFailedElementNumber"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for firstFailedElementNumber")
	}
	if _firstFailedElementNumberErr != nil {
		return errors.Wrap(_firstFailedElementNumberErr, "Error serializing 'firstFailedElementNumber' field")
	}

		if popErr := writeBuffer.PopContext("ChangeListRemoveError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ChangeListRemoveError")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_ChangeListRemoveError) isChangeListRemoveError() bool {
	return true
}

func (m *_ChangeListRemoveError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



