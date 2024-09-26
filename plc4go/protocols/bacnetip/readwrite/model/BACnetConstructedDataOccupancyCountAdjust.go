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

// BACnetConstructedDataOccupancyCountAdjust is the corresponding interface of BACnetConstructedDataOccupancyCountAdjust
type BACnetConstructedDataOccupancyCountAdjust interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetOccupancyCountAdjust returns OccupancyCountAdjust (property field)
	GetOccupancyCountAdjust() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataOccupancyCountAdjust is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataOccupancyCountAdjust()
}

// _BACnetConstructedDataOccupancyCountAdjust is the data-structure of this message
type _BACnetConstructedDataOccupancyCountAdjust struct {
	BACnetConstructedDataContract
	OccupancyCountAdjust BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataOccupancyCountAdjust = (*_BACnetConstructedDataOccupancyCountAdjust)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataOccupancyCountAdjust)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOccupancyCountAdjust) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OCCUPANCY_COUNT_ADJUST
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOccupancyCountAdjust) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyCountAdjust) GetOccupancyCountAdjust() BACnetApplicationTagBoolean {
	return m.OccupancyCountAdjust
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyCountAdjust) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetOccupancyCountAdjust())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataOccupancyCountAdjust factory function for _BACnetConstructedDataOccupancyCountAdjust
func NewBACnetConstructedDataOccupancyCountAdjust(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, occupancyCountAdjust BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataOccupancyCountAdjust {
	if occupancyCountAdjust == nil {
		panic("occupancyCountAdjust of type BACnetApplicationTagBoolean for BACnetConstructedDataOccupancyCountAdjust must not be nil")
	}
	_result := &_BACnetConstructedDataOccupancyCountAdjust{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		OccupancyCountAdjust:          occupancyCountAdjust,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOccupancyCountAdjust(structType any) BACnetConstructedDataOccupancyCountAdjust {
	if casted, ok := structType.(BACnetConstructedDataOccupancyCountAdjust); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOccupancyCountAdjust); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) GetTypeName() string {
	return "BACnetConstructedDataOccupancyCountAdjust"
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (occupancyCountAdjust)
	lengthInBits += m.OccupancyCountAdjust.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataOccupancyCountAdjust BACnetConstructedDataOccupancyCountAdjust, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOccupancyCountAdjust"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOccupancyCountAdjust")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	occupancyCountAdjust, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "occupancyCountAdjust", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'occupancyCountAdjust' field"))
	}
	m.OccupancyCountAdjust = occupancyCountAdjust

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), occupancyCountAdjust)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOccupancyCountAdjust"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOccupancyCountAdjust")
	}

	return m, nil
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOccupancyCountAdjust"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOccupancyCountAdjust")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "occupancyCountAdjust", m.GetOccupancyCountAdjust(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'occupancyCountAdjust' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOccupancyCountAdjust"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOccupancyCountAdjust")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataOccupancyCountAdjust) IsBACnetConstructedDataOccupancyCountAdjust() {}

func (m *_BACnetConstructedDataOccupancyCountAdjust) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
