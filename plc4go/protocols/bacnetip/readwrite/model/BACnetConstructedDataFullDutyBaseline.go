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

// BACnetConstructedDataFullDutyBaseline is the corresponding interface of BACnetConstructedDataFullDutyBaseline
type BACnetConstructedDataFullDutyBaseline interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFullDutyBaseLine returns FullDutyBaseLine (property field)
	GetFullDutyBaseLine() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
	// IsBACnetConstructedDataFullDutyBaseline is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataFullDutyBaseline()
}

// _BACnetConstructedDataFullDutyBaseline is the data-structure of this message
type _BACnetConstructedDataFullDutyBaseline struct {
	BACnetConstructedDataContract
	FullDutyBaseLine BACnetApplicationTagReal
}

var _ BACnetConstructedDataFullDutyBaseline = (*_BACnetConstructedDataFullDutyBaseline)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataFullDutyBaseline)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFullDutyBaseline) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFullDutyBaseline) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FULL_DUTY_BASELINE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFullDutyBaseline) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFullDutyBaseline) GetFullDutyBaseLine() BACnetApplicationTagReal {
	return m.FullDutyBaseLine
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFullDutyBaseline) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetFullDutyBaseLine())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFullDutyBaseline factory function for _BACnetConstructedDataFullDutyBaseline
func NewBACnetConstructedDataFullDutyBaseline(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, fullDutyBaseLine BACnetApplicationTagReal, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFullDutyBaseline {
	if fullDutyBaseLine == nil {
		panic("fullDutyBaseLine of type BACnetApplicationTagReal for BACnetConstructedDataFullDutyBaseline must not be nil")
	}
	_result := &_BACnetConstructedDataFullDutyBaseline{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FullDutyBaseLine:              fullDutyBaseLine,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFullDutyBaseline(structType any) BACnetConstructedDataFullDutyBaseline {
	if casted, ok := structType.(BACnetConstructedDataFullDutyBaseline); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFullDutyBaseline); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFullDutyBaseline) GetTypeName() string {
	return "BACnetConstructedDataFullDutyBaseline"
}

func (m *_BACnetConstructedDataFullDutyBaseline) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (fullDutyBaseLine)
	lengthInBits += m.FullDutyBaseLine.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataFullDutyBaseline) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataFullDutyBaseline) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataFullDutyBaseline BACnetConstructedDataFullDutyBaseline, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFullDutyBaseline"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFullDutyBaseline")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	fullDutyBaseLine, err := ReadSimpleField[BACnetApplicationTagReal](ctx, "fullDutyBaseLine", ReadComplex[BACnetApplicationTagReal](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagReal](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'fullDutyBaseLine' field"))
	}
	m.FullDutyBaseLine = fullDutyBaseLine

	actualValue, err := ReadVirtualField[BACnetApplicationTagReal](ctx, "actualValue", (*BACnetApplicationTagReal)(nil), fullDutyBaseLine)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFullDutyBaseline"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFullDutyBaseline")
	}

	return m, nil
}

func (m *_BACnetConstructedDataFullDutyBaseline) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFullDutyBaseline) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFullDutyBaseline"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFullDutyBaseline")
		}

		if err := WriteSimpleField[BACnetApplicationTagReal](ctx, "fullDutyBaseLine", m.GetFullDutyBaseLine(), WriteComplex[BACnetApplicationTagReal](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'fullDutyBaseLine' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFullDutyBaseline"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFullDutyBaseline")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFullDutyBaseline) IsBACnetConstructedDataFullDutyBaseline() {}

func (m *_BACnetConstructedDataFullDutyBaseline) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
