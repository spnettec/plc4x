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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyAccessResultAccessResultPropertyAccessError is the corresponding interface of BACnetPropertyAccessResultAccessResultPropertyAccessError
type BACnetPropertyAccessResultAccessResultPropertyAccessError interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyAccessResultAccessResult
	// GetPropertyAccessError returns PropertyAccessError (property field)
	GetPropertyAccessError() ErrorEnclosed
	// IsBACnetPropertyAccessResultAccessResultPropertyAccessError is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyAccessResultAccessResultPropertyAccessError()
}

// _BACnetPropertyAccessResultAccessResultPropertyAccessError is the data-structure of this message
type _BACnetPropertyAccessResultAccessResultPropertyAccessError struct {
	BACnetPropertyAccessResultAccessResultContract
	PropertyAccessError ErrorEnclosed
}

var _ BACnetPropertyAccessResultAccessResultPropertyAccessError = (*_BACnetPropertyAccessResultAccessResultPropertyAccessError)(nil)
var _ BACnetPropertyAccessResultAccessResultRequirements = (*_BACnetPropertyAccessResultAccessResultPropertyAccessError)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) GetParent() BACnetPropertyAccessResultAccessResultContract {
	return m.BACnetPropertyAccessResultAccessResultContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) GetPropertyAccessError() ErrorEnclosed {
	return m.PropertyAccessError
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyAccessResultAccessResultPropertyAccessError factory function for _BACnetPropertyAccessResultAccessResultPropertyAccessError
func NewBACnetPropertyAccessResultAccessResultPropertyAccessError(peekedTagHeader BACnetTagHeader, propertyAccessError ErrorEnclosed, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetPropertyAccessResultAccessResultPropertyAccessError {
	if propertyAccessError == nil {
		panic("propertyAccessError of type ErrorEnclosed for BACnetPropertyAccessResultAccessResultPropertyAccessError must not be nil")
	}
	_result := &_BACnetPropertyAccessResultAccessResultPropertyAccessError{
		BACnetPropertyAccessResultAccessResultContract: NewBACnetPropertyAccessResultAccessResult(peekedTagHeader, objectTypeArgument, propertyIdentifierArgument, propertyArrayIndexArgument),
		PropertyAccessError:                            propertyAccessError,
	}
	_result.BACnetPropertyAccessResultAccessResultContract.(*_BACnetPropertyAccessResultAccessResult)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyAccessResultAccessResultPropertyAccessError(structType any) BACnetPropertyAccessResultAccessResultPropertyAccessError {
	if casted, ok := structType.(BACnetPropertyAccessResultAccessResultPropertyAccessError); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyAccessResultAccessResultPropertyAccessError); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) GetTypeName() string {
	return "BACnetPropertyAccessResultAccessResultPropertyAccessError"
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyAccessResultAccessResultContract.(*_BACnetPropertyAccessResultAccessResult).getLengthInBits(ctx))

	// Simple field (propertyAccessError)
	lengthInBits += m.PropertyAccessError.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyAccessResultAccessResult, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, propertyArrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetPropertyAccessResultAccessResultPropertyAccessError BACnetPropertyAccessResultAccessResultPropertyAccessError, err error) {
	m.BACnetPropertyAccessResultAccessResultContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyAccessResultAccessResultPropertyAccessError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyAccessResultAccessResultPropertyAccessError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	propertyAccessError, err := ReadSimpleField[ErrorEnclosed](ctx, "propertyAccessError", ReadComplex[ErrorEnclosed](ErrorEnclosedParseWithBufferProducer((uint8)(uint8(5))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'propertyAccessError' field"))
	}
	m.PropertyAccessError = propertyAccessError

	if closeErr := readBuffer.CloseContext("BACnetPropertyAccessResultAccessResultPropertyAccessError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyAccessResultAccessResultPropertyAccessError")
	}

	return m, nil
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyAccessResultAccessResultPropertyAccessError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyAccessResultAccessResultPropertyAccessError")
		}

		if err := WriteSimpleField[ErrorEnclosed](ctx, "propertyAccessError", m.GetPropertyAccessError(), WriteComplex[ErrorEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'propertyAccessError' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyAccessResultAccessResultPropertyAccessError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyAccessResultAccessResultPropertyAccessError")
		}
		return nil
	}
	return m.BACnetPropertyAccessResultAccessResultContract.(*_BACnetPropertyAccessResultAccessResult).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) IsBACnetPropertyAccessResultAccessResultPropertyAccessError() {
}

func (m *_BACnetPropertyAccessResultAccessResultPropertyAccessError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
