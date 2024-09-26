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

// BACnetConstructedDataCountBeforeChange is the corresponding interface of BACnetConstructedDataCountBeforeChange
type BACnetConstructedDataCountBeforeChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCountBeforeChange returns CountBeforeChange (property field)
	GetCountBeforeChange() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// IsBACnetConstructedDataCountBeforeChange is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataCountBeforeChange()
}

// _BACnetConstructedDataCountBeforeChange is the data-structure of this message
type _BACnetConstructedDataCountBeforeChange struct {
	BACnetConstructedDataContract
	CountBeforeChange BACnetApplicationTagUnsignedInteger
}

var _ BACnetConstructedDataCountBeforeChange = (*_BACnetConstructedDataCountBeforeChange)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataCountBeforeChange)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCountBeforeChange) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCountBeforeChange) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_COUNT_BEFORE_CHANGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCountBeforeChange) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCountBeforeChange) GetCountBeforeChange() BACnetApplicationTagUnsignedInteger {
	return m.CountBeforeChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCountBeforeChange) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetCountBeforeChange())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCountBeforeChange factory function for _BACnetConstructedDataCountBeforeChange
func NewBACnetConstructedDataCountBeforeChange(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, countBeforeChange BACnetApplicationTagUnsignedInteger, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCountBeforeChange {
	if countBeforeChange == nil {
		panic("countBeforeChange of type BACnetApplicationTagUnsignedInteger for BACnetConstructedDataCountBeforeChange must not be nil")
	}
	_result := &_BACnetConstructedDataCountBeforeChange{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		CountBeforeChange:             countBeforeChange,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCountBeforeChange(structType any) BACnetConstructedDataCountBeforeChange {
	if casted, ok := structType.(BACnetConstructedDataCountBeforeChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCountBeforeChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCountBeforeChange) GetTypeName() string {
	return "BACnetConstructedDataCountBeforeChange"
}

func (m *_BACnetConstructedDataCountBeforeChange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (countBeforeChange)
	lengthInBits += m.CountBeforeChange.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCountBeforeChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataCountBeforeChange) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataCountBeforeChange BACnetConstructedDataCountBeforeChange, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCountBeforeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCountBeforeChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	countBeforeChange, err := ReadSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "countBeforeChange", ReadComplex[BACnetApplicationTagUnsignedInteger](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagUnsignedInteger](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'countBeforeChange' field"))
	}
	m.CountBeforeChange = countBeforeChange

	actualValue, err := ReadVirtualField[BACnetApplicationTagUnsignedInteger](ctx, "actualValue", (*BACnetApplicationTagUnsignedInteger)(nil), countBeforeChange)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCountBeforeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCountBeforeChange")
	}

	return m, nil
}

func (m *_BACnetConstructedDataCountBeforeChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCountBeforeChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCountBeforeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCountBeforeChange")
		}

		if err := WriteSimpleField[BACnetApplicationTagUnsignedInteger](ctx, "countBeforeChange", m.GetCountBeforeChange(), WriteComplex[BACnetApplicationTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'countBeforeChange' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCountBeforeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCountBeforeChange")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCountBeforeChange) IsBACnetConstructedDataCountBeforeChange() {}

func (m *_BACnetConstructedDataCountBeforeChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
