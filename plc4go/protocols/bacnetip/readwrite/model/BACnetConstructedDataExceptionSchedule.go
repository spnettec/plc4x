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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataExceptionSchedule is the corresponding interface of BACnetConstructedDataExceptionSchedule
type BACnetConstructedDataExceptionSchedule interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetExceptionSchedule returns ExceptionSchedule (property field)
	GetExceptionSchedule() []BACnetSpecialEvent
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataExceptionScheduleExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataExceptionSchedule.
// This is useful for switch cases.
type BACnetConstructedDataExceptionScheduleExactly interface {
	BACnetConstructedDataExceptionSchedule
	isBACnetConstructedDataExceptionSchedule() bool
}

// _BACnetConstructedDataExceptionSchedule is the data-structure of this message
type _BACnetConstructedDataExceptionSchedule struct {
	*_BACnetConstructedData
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	ExceptionSchedule    []BACnetSpecialEvent
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataExceptionSchedule) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataExceptionSchedule) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EXCEPTION_SCHEDULE
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataExceptionSchedule) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataExceptionSchedule) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataExceptionSchedule) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataExceptionSchedule) GetExceptionSchedule() []BACnetSpecialEvent {
	return m.ExceptionSchedule
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataExceptionSchedule) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataExceptionSchedule factory function for _BACnetConstructedDataExceptionSchedule
func NewBACnetConstructedDataExceptionSchedule(numberOfDataElements BACnetApplicationTagUnsignedInteger, exceptionSchedule []BACnetSpecialEvent, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataExceptionSchedule {
	_result := &_BACnetConstructedDataExceptionSchedule{
		NumberOfDataElements:   numberOfDataElements,
		ExceptionSchedule:      exceptionSchedule,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataExceptionSchedule(structType any) BACnetConstructedDataExceptionSchedule {
	if casted, ok := structType.(BACnetConstructedDataExceptionSchedule); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataExceptionSchedule); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataExceptionSchedule) GetTypeName() string {
	return "BACnetConstructedDataExceptionSchedule"
}

func (m *_BACnetConstructedDataExceptionSchedule) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.ExceptionSchedule) > 0 {
		for _, element := range m.ExceptionSchedule {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataExceptionSchedule) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataExceptionScheduleParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataExceptionSchedule, error) {
	return BACnetConstructedDataExceptionScheduleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataExceptionScheduleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataExceptionSchedule, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataExceptionSchedule"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataExceptionSchedule")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
		_val, _err := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field of BACnetConstructedDataExceptionSchedule")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (exceptionSchedule)
	if pullErr := readBuffer.PullContext("exceptionSchedule", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for exceptionSchedule")
	}
	// Terminated array
	var exceptionSchedule []BACnetSpecialEvent
	{
		for !bool(IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber)) {
			_item, _err := BACnetSpecialEventParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'exceptionSchedule' field of BACnetConstructedDataExceptionSchedule")
			}
			exceptionSchedule = append(exceptionSchedule, _item.(BACnetSpecialEvent))
		}
	}
	if closeErr := readBuffer.CloseContext("exceptionSchedule", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for exceptionSchedule")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataExceptionSchedule"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataExceptionSchedule")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataExceptionSchedule{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NumberOfDataElements: numberOfDataElements,
		ExceptionSchedule:    exceptionSchedule,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataExceptionSchedule) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataExceptionSchedule) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataExceptionSchedule"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataExceptionSchedule")
		}
		// Virtual field
		zero := m.GetZero()
		_ = zero
		if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
		if m.GetNumberOfDataElements() != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.GetNumberOfDataElements()
			_numberOfDataElementsErr := writeBuffer.WriteSerializable(ctx, numberOfDataElements)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		// Array Field (exceptionSchedule)
		if pushErr := writeBuffer.PushContext("exceptionSchedule", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for exceptionSchedule")
		}
		for _curItem, _element := range m.GetExceptionSchedule() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetExceptionSchedule()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'exceptionSchedule' field")
			}
		}
		if popErr := writeBuffer.PopContext("exceptionSchedule", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for exceptionSchedule")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataExceptionSchedule"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataExceptionSchedule")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataExceptionSchedule) isBACnetConstructedDataExceptionSchedule() bool {
	return true
}

func (m *_BACnetConstructedDataExceptionSchedule) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
