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

// BACnetConstructedDataTimeSynchronizationInterval is the corresponding interface of BACnetConstructedDataTimeSynchronizationInterval
type BACnetConstructedDataTimeSynchronizationInterval interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTimeSynchronization returns TimeSynchronization (property field)
	GetTimeSynchronization() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataTimeSynchronizationIntervalExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTimeSynchronizationInterval.
// This is useful for switch cases.
type BACnetConstructedDataTimeSynchronizationIntervalExactly interface {
	BACnetConstructedDataTimeSynchronizationInterval
	isBACnetConstructedDataTimeSynchronizationInterval() bool
}

// _BACnetConstructedDataTimeSynchronizationInterval is the data-structure of this message
type _BACnetConstructedDataTimeSynchronizationInterval struct {
	*_BACnetConstructedData
	TimeSynchronization BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeSynchronizationInterval) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TIME_SYNCHRONIZATION_INTERVAL
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeSynchronizationInterval) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimeSynchronizationInterval) GetTimeSynchronization() BACnetApplicationTagUnsignedInteger {
	return m.TimeSynchronization
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimeSynchronizationInterval) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetTimeSynchronization())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTimeSynchronizationInterval factory function for _BACnetConstructedDataTimeSynchronizationInterval
func NewBACnetConstructedDataTimeSynchronizationInterval(timeSynchronization BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTimeSynchronizationInterval {
	_result := &_BACnetConstructedDataTimeSynchronizationInterval{
		TimeSynchronization:    timeSynchronization,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeSynchronizationInterval(structType any) BACnetConstructedDataTimeSynchronizationInterval {
	if casted, ok := structType.(BACnetConstructedDataTimeSynchronizationInterval); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeSynchronizationInterval); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) GetTypeName() string {
	return "BACnetConstructedDataTimeSynchronizationInterval"
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (timeSynchronization)
	lengthInBits += m.TimeSynchronization.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataTimeSynchronizationIntervalParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimeSynchronizationInterval, error) {
	return BACnetConstructedDataTimeSynchronizationIntervalParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataTimeSynchronizationIntervalParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimeSynchronizationInterval, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeSynchronizationInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeSynchronizationInterval")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeSynchronization)
	if pullErr := readBuffer.PullContext("timeSynchronization"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeSynchronization")
	}
	_timeSynchronization, _timeSynchronizationErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _timeSynchronizationErr != nil {
		return nil, errors.Wrap(_timeSynchronizationErr, "Error parsing 'timeSynchronization' field of BACnetConstructedDataTimeSynchronizationInterval")
	}
	timeSynchronization := _timeSynchronization.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("timeSynchronization"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeSynchronization")
	}

	// Virtual field
	_actualValue := timeSynchronization
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeSynchronizationInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeSynchronizationInterval")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTimeSynchronizationInterval{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		TimeSynchronization: timeSynchronization,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeSynchronizationInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeSynchronizationInterval")
		}

		// Simple Field (timeSynchronization)
		if pushErr := writeBuffer.PushContext("timeSynchronization"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for timeSynchronization")
		}
		_timeSynchronizationErr := writeBuffer.WriteSerializable(ctx, m.GetTimeSynchronization())
		if popErr := writeBuffer.PopContext("timeSynchronization"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for timeSynchronization")
		}
		if _timeSynchronizationErr != nil {
			return errors.Wrap(_timeSynchronizationErr, "Error serializing 'timeSynchronization' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeSynchronizationInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeSynchronizationInterval")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) isBACnetConstructedDataTimeSynchronizationInterval() bool {
	return true
}

func (m *_BACnetConstructedDataTimeSynchronizationInterval) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
