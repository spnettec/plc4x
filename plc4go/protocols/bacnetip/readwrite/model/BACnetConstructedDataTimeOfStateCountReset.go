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

// BACnetConstructedDataTimeOfStateCountReset is the corresponding interface of BACnetConstructedDataTimeOfStateCountReset
type BACnetConstructedDataTimeOfStateCountReset interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTimeOfStateCountReset returns TimeOfStateCountReset (property field)
	GetTimeOfStateCountReset() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// BACnetConstructedDataTimeOfStateCountResetExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTimeOfStateCountReset.
// This is useful for switch cases.
type BACnetConstructedDataTimeOfStateCountResetExactly interface {
	BACnetConstructedDataTimeOfStateCountReset
	isBACnetConstructedDataTimeOfStateCountReset() bool
}

// _BACnetConstructedDataTimeOfStateCountReset is the data-structure of this message
type _BACnetConstructedDataTimeOfStateCountReset struct {
	*_BACnetConstructedData
	TimeOfStateCountReset BACnetDateTime
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeOfStateCountReset) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_OF_STATE_COUNT_RESET
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeOfStateCountReset) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimeOfStateCountReset) GetTimeOfStateCountReset() BACnetDateTime {
	return m.TimeOfStateCountReset
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimeOfStateCountReset) GetActualValue() BACnetDateTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDateTime(m.GetTimeOfStateCountReset())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimeOfStateCountReset factory function for _BACnetConstructedDataTimeOfStateCountReset
func NewBACnetConstructedDataTimeOfStateCountReset(timeOfStateCountReset BACnetDateTime, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimeOfStateCountReset {
	_result := &_BACnetConstructedDataTimeOfStateCountReset{
		TimeOfStateCountReset:  timeOfStateCountReset,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeOfStateCountReset(structType any) BACnetConstructedDataTimeOfStateCountReset {
	if casted, ok := structType.(BACnetConstructedDataTimeOfStateCountReset); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeOfStateCountReset); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) GetTypeName() string {
	return "BACnetConstructedDataTimeOfStateCountReset"
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (timeOfStateCountReset)
	lengthInBits += m.TimeOfStateCountReset.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataTimeOfStateCountResetParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimeOfStateCountReset, error) {
	return BACnetConstructedDataTimeOfStateCountResetParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataTimeOfStateCountResetParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimeOfStateCountReset, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeOfStateCountReset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeOfStateCountReset")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeOfStateCountReset)
	if pullErr := readBuffer.PullContext("timeOfStateCountReset"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeOfStateCountReset")
	}
	_timeOfStateCountReset, _timeOfStateCountResetErr := BACnetDateTimeParseWithBuffer(ctx, readBuffer)
	if _timeOfStateCountResetErr != nil {
		return nil, errors.Wrap(_timeOfStateCountResetErr, "Error parsing 'timeOfStateCountReset' field of BACnetConstructedDataTimeOfStateCountReset")
	}
	timeOfStateCountReset := _timeOfStateCountReset.(BACnetDateTime)
	if closeErr := readBuffer.CloseContext("timeOfStateCountReset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeOfStateCountReset")
	}

	// Virtual field
	_actualValue := timeOfStateCountReset
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeOfStateCountReset"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeOfStateCountReset")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTimeOfStateCountReset{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		TimeOfStateCountReset: timeOfStateCountReset,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeOfStateCountReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeOfStateCountReset")
		}

		// Simple Field (timeOfStateCountReset)
		if pushErr := writeBuffer.PushContext("timeOfStateCountReset"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeOfStateCountReset")
		}
		_timeOfStateCountResetErr := writeBuffer.WriteSerializable(ctx, m.GetTimeOfStateCountReset())
		if popErr := writeBuffer.PopContext("timeOfStateCountReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeOfStateCountReset")
		}
		if _timeOfStateCountResetErr != nil {
			return errors.Wrap(_timeOfStateCountResetErr, "Error serializing 'timeOfStateCountReset' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeOfStateCountReset"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeOfStateCountReset")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) isBACnetConstructedDataTimeOfStateCountReset() bool {
	return true
}

func (m *_BACnetConstructedDataTimeOfStateCountReset) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
