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

// BACnetConstructedDataActivationTime is the corresponding interface of BACnetConstructedDataActivationTime
type BACnetConstructedDataActivationTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetActivationTime returns ActivationTime (property field)
	GetActivationTime() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// BACnetConstructedDataActivationTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataActivationTime.
// This is useful for switch cases.
type BACnetConstructedDataActivationTimeExactly interface {
	BACnetConstructedDataActivationTime
	isBACnetConstructedDataActivationTime() bool
}

// _BACnetConstructedDataActivationTime is the data-structure of this message
type _BACnetConstructedDataActivationTime struct {
	*_BACnetConstructedData
	ActivationTime BACnetDateTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataActivationTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataActivationTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACTIVATION_TIME
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataActivationTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataActivationTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataActivationTime) GetActivationTime() BACnetDateTime {
	return m.ActivationTime
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataActivationTime) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetActivationTime())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataActivationTime factory function for _BACnetConstructedDataActivationTime
func NewBACnetConstructedDataActivationTime(activationTime BACnetDateTime, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataActivationTime {
	_result := &_BACnetConstructedDataActivationTime{
		ActivationTime:         activationTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataActivationTime(structType any) BACnetConstructedDataActivationTime {
	if casted, ok := structType.(BACnetConstructedDataActivationTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataActivationTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataActivationTime) GetTypeName() string {
	return "BACnetConstructedDataActivationTime"
}

func (m *_BACnetConstructedDataActivationTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (activationTime)
	lengthInBits += m.ActivationTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataActivationTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataActivationTimeParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataActivationTime, error) {
	return BACnetConstructedDataActivationTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataActivationTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataActivationTime, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataActivationTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataActivationTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (activationTime)
	if pullErr := readBuffer.PullContext("activationTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for activationTime")
	}
	_activationTime, _activationTimeErr := BACnetDateTimeParseWithBuffer(ctx, readBuffer)
	if _activationTimeErr != nil {
		return nil, errors.Wrap(_activationTimeErr, "Error parsing 'activationTime' field of BACnetConstructedDataActivationTime")
	}
	activationTime := _activationTime.(BACnetDateTime)
	if closeErr := readBuffer.CloseContext("activationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for activationTime")
	}

	// Virtual field
	_actualValue := activationTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataActivationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataActivationTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataActivationTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ActivationTime: activationTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataActivationTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataActivationTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataActivationTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataActivationTime")
		}

		// Simple Field (activationTime)
		if pushErr := writeBuffer.PushContext("activationTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for activationTime")
		}
		_activationTimeErr := writeBuffer.WriteSerializable(ctx, m.GetActivationTime())
		if popErr := writeBuffer.PopContext("activationTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for activationTime")
		}
		if _activationTimeErr != nil {
			return errors.Wrap(_activationTimeErr, "Error serializing 'activationTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataActivationTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataActivationTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataActivationTime) isBACnetConstructedDataActivationTime() bool {
	return true
}

func (m *_BACnetConstructedDataActivationTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
