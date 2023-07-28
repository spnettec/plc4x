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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

	// Code generated by code-generation. DO NOT EDIT.


// CreateObjectError is the corresponding interface of CreateObjectError
type CreateObjectError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetError
	// GetErrorType returns ErrorType (property field)
	GetErrorType() ErrorEnclosed
	// GetFirstFailedElementNumber returns FirstFailedElementNumber (property field)
	GetFirstFailedElementNumber() BACnetContextTagUnsignedInteger
}

// CreateObjectErrorExactly can be used when we want exactly this type and not a type which fulfills CreateObjectError.
// This is useful for switch cases.
type CreateObjectErrorExactly interface {
	CreateObjectError
	isCreateObjectError() bool
}

// _CreateObjectError is the data-structure of this message
type _CreateObjectError struct {
	*_BACnetError
        ErrorType ErrorEnclosed
        FirstFailedElementNumber BACnetContextTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CreateObjectError)  GetErrorChoice() BACnetConfirmedServiceChoice {
return BACnetConfirmedServiceChoice_CREATE_OBJECT}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CreateObjectError) InitializeParent(parent BACnetError ) {}

func (m *_CreateObjectError)  GetParent() BACnetError {
	return m._BACnetError
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CreateObjectError) GetErrorType() ErrorEnclosed {
	return m.ErrorType
}

func (m *_CreateObjectError) GetFirstFailedElementNumber() BACnetContextTagUnsignedInteger {
	return m.FirstFailedElementNumber
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewCreateObjectError factory function for _CreateObjectError
func NewCreateObjectError( errorType ErrorEnclosed , firstFailedElementNumber BACnetContextTagUnsignedInteger ) *_CreateObjectError {
	_result := &_CreateObjectError{
		ErrorType: errorType,
		FirstFailedElementNumber: firstFailedElementNumber,
    	_BACnetError: NewBACnetError(),
	}
	_result._BACnetError._BACnetErrorChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCreateObjectError(structType any) CreateObjectError {
    if casted, ok := structType.(CreateObjectError); ok {
		return casted
	}
	if casted, ok := structType.(*CreateObjectError); ok {
		return *casted
	}
	return nil
}

func (m *_CreateObjectError) GetTypeName() string {
	return "CreateObjectError"
}

func (m *_CreateObjectError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (errorType)
	lengthInBits += m.ErrorType.GetLengthInBits(ctx)

	// Simple field (firstFailedElementNumber)
	lengthInBits += m.FirstFailedElementNumber.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_CreateObjectError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CreateObjectErrorParse(ctx context.Context, theBytes []byte, errorChoice BACnetConfirmedServiceChoice) (CreateObjectError, error) {
	return CreateObjectErrorParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), errorChoice)
}

func CreateObjectErrorParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, errorChoice BACnetConfirmedServiceChoice) (CreateObjectError, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CreateObjectError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CreateObjectError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (errorType)
	if pullErr := readBuffer.PullContext("errorType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for errorType")
	}
_errorType, _errorTypeErr := ErrorEnclosedParseWithBuffer(ctx, readBuffer , uint8( uint8(0) ) )
	if _errorTypeErr != nil {
		return nil, errors.Wrap(_errorTypeErr, "Error parsing 'errorType' field of CreateObjectError")
	}
	errorType := _errorType.(ErrorEnclosed)
	if closeErr := readBuffer.CloseContext("errorType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for errorType")
	}

	// Simple Field (firstFailedElementNumber)
	if pullErr := readBuffer.PullContext("firstFailedElementNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for firstFailedElementNumber")
	}
_firstFailedElementNumber, _firstFailedElementNumberErr := BACnetContextTagParseWithBuffer(ctx, readBuffer , uint8( uint8(1) ) , BACnetDataType( BACnetDataType_UNSIGNED_INTEGER ) )
	if _firstFailedElementNumberErr != nil {
		return nil, errors.Wrap(_firstFailedElementNumberErr, "Error parsing 'firstFailedElementNumber' field of CreateObjectError")
	}
	firstFailedElementNumber := _firstFailedElementNumber.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("firstFailedElementNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for firstFailedElementNumber")
	}

	if closeErr := readBuffer.CloseContext("CreateObjectError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CreateObjectError")
	}

	// Create a partially initialized instance
	_child := &_CreateObjectError{
		_BACnetError: &_BACnetError{
		},
		ErrorType: errorType,
		FirstFailedElementNumber: firstFailedElementNumber,
	}
	_child._BACnetError._BACnetErrorChildRequirements = _child
	return _child, nil
}

func (m *_CreateObjectError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CreateObjectError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CreateObjectError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CreateObjectError")
		}

	// Simple Field (errorType)
	if pushErr := writeBuffer.PushContext("errorType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for errorType")
	}
	_errorTypeErr := writeBuffer.WriteSerializable(ctx, m.GetErrorType())
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
	_firstFailedElementNumberErr := writeBuffer.WriteSerializable(ctx, m.GetFirstFailedElementNumber())
	if popErr := writeBuffer.PopContext("firstFailedElementNumber"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for firstFailedElementNumber")
	}
	if _firstFailedElementNumberErr != nil {
		return errors.Wrap(_firstFailedElementNumberErr, "Error serializing 'firstFailedElementNumber' field")
	}

		if popErr := writeBuffer.PopContext("CreateObjectError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CreateObjectError")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_CreateObjectError) isCreateObjectError() bool {
	return true
}

func (m *_CreateObjectError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



