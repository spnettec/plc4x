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

// BACnetConstructedDataLastCommandTime is the corresponding interface of BACnetConstructedDataLastCommandTime
type BACnetConstructedDataLastCommandTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLastCommandTime returns LastCommandTime (property field)
	GetLastCommandTime() BACnetTimeStamp
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetTimeStamp
}

// BACnetConstructedDataLastCommandTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLastCommandTime.
// This is useful for switch cases.
type BACnetConstructedDataLastCommandTimeExactly interface {
	BACnetConstructedDataLastCommandTime
	isBACnetConstructedDataLastCommandTime() bool
}

// _BACnetConstructedDataLastCommandTime is the data-structure of this message
type _BACnetConstructedDataLastCommandTime struct {
	*_BACnetConstructedData
	LastCommandTime BACnetTimeStamp
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastCommandTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastCommandTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_COMMAND_TIME
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastCommandTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLastCommandTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastCommandTime) GetLastCommandTime() BACnetTimeStamp {
	return m.LastCommandTime
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastCommandTime) GetActualValue() BACnetTimeStamp {
	ctx := context.Background()
	_ = ctx
	return CastBACnetTimeStamp(m.GetLastCommandTime())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLastCommandTime factory function for _BACnetConstructedDataLastCommandTime
func NewBACnetConstructedDataLastCommandTime(lastCommandTime BACnetTimeStamp, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastCommandTime {
	_result := &_BACnetConstructedDataLastCommandTime{
		LastCommandTime:        lastCommandTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastCommandTime(structType any) BACnetConstructedDataLastCommandTime {
	if casted, ok := structType.(BACnetConstructedDataLastCommandTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastCommandTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastCommandTime) GetTypeName() string {
	return "BACnetConstructedDataLastCommandTime"
}

func (m *_BACnetConstructedDataLastCommandTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (lastCommandTime)
	lengthInBits += m.LastCommandTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastCommandTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLastCommandTimeParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastCommandTime, error) {
	return BACnetConstructedDataLastCommandTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLastCommandTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastCommandTime, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastCommandTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastCommandTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lastCommandTime)
	if pullErr := readBuffer.PullContext("lastCommandTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lastCommandTime")
	}
	_lastCommandTime, _lastCommandTimeErr := BACnetTimeStampParseWithBuffer(ctx, readBuffer)
	if _lastCommandTimeErr != nil {
		return nil, errors.Wrap(_lastCommandTimeErr, "Error parsing 'lastCommandTime' field of BACnetConstructedDataLastCommandTime")
	}
	lastCommandTime := _lastCommandTime.(BACnetTimeStamp)
	if closeErr := readBuffer.CloseContext("lastCommandTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lastCommandTime")
	}

	// Virtual field
	_actualValue := lastCommandTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastCommandTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastCommandTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLastCommandTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LastCommandTime: lastCommandTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLastCommandTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLastCommandTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastCommandTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastCommandTime")
		}

		// Simple Field (lastCommandTime)
		if pushErr := writeBuffer.PushContext("lastCommandTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lastCommandTime")
		}
		_lastCommandTimeErr := writeBuffer.WriteSerializable(ctx, m.GetLastCommandTime())
		if popErr := writeBuffer.PopContext("lastCommandTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lastCommandTime")
		}
		if _lastCommandTimeErr != nil {
			return errors.Wrap(_lastCommandTimeErr, "Error serializing 'lastCommandTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastCommandTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastCommandTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastCommandTime) isBACnetConstructedDataLastCommandTime() bool {
	return true
}

func (m *_BACnetConstructedDataLastCommandTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
