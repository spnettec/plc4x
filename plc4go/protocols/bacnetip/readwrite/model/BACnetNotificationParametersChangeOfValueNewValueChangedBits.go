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

// BACnetNotificationParametersChangeOfValueNewValueChangedBits is the corresponding interface of BACnetNotificationParametersChangeOfValueNewValueChangedBits
type BACnetNotificationParametersChangeOfValueNewValueChangedBits interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParametersChangeOfValueNewValue
	// GetChangedBits returns ChangedBits (property field)
	GetChangedBits() BACnetContextTagBitString
	// IsBACnetNotificationParametersChangeOfValueNewValueChangedBits is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersChangeOfValueNewValueChangedBits()
}

// _BACnetNotificationParametersChangeOfValueNewValueChangedBits is the data-structure of this message
type _BACnetNotificationParametersChangeOfValueNewValueChangedBits struct {
	BACnetNotificationParametersChangeOfValueNewValueContract
	ChangedBits BACnetContextTagBitString
}

var _ BACnetNotificationParametersChangeOfValueNewValueChangedBits = (*_BACnetNotificationParametersChangeOfValueNewValueChangedBits)(nil)
var _ BACnetNotificationParametersChangeOfValueNewValueRequirements = (*_BACnetNotificationParametersChangeOfValueNewValueChangedBits)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) GetParent() BACnetNotificationParametersChangeOfValueNewValueContract {
	return m.BACnetNotificationParametersChangeOfValueNewValueContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) GetChangedBits() BACnetContextTagBitString {
	return m.ChangedBits
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfValueNewValueChangedBits factory function for _BACnetNotificationParametersChangeOfValueNewValueChangedBits
func NewBACnetNotificationParametersChangeOfValueNewValueChangedBits(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, changedBits BACnetContextTagBitString, tagNumber uint8) *_BACnetNotificationParametersChangeOfValueNewValueChangedBits {
	if changedBits == nil {
		panic("changedBits of type BACnetContextTagBitString for BACnetNotificationParametersChangeOfValueNewValueChangedBits must not be nil")
	}
	_result := &_BACnetNotificationParametersChangeOfValueNewValueChangedBits{
		BACnetNotificationParametersChangeOfValueNewValueContract: NewBACnetNotificationParametersChangeOfValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
		ChangedBits: changedBits,
	}
	_result.BACnetNotificationParametersChangeOfValueNewValueContract.(*_BACnetNotificationParametersChangeOfValueNewValue)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfValueNewValueChangedBits(structType any) BACnetNotificationParametersChangeOfValueNewValueChangedBits {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfValueNewValueChangedBits); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfValueNewValueChangedBits); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfValueNewValueChangedBits"
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersChangeOfValueNewValueContract.(*_BACnetNotificationParametersChangeOfValueNewValue).getLengthInBits(ctx))

	// Simple field (changedBits)
	lengthInBits += m.ChangedBits.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParametersChangeOfValueNewValue, peekedTagNumber uint8, tagNumber uint8) (__bACnetNotificationParametersChangeOfValueNewValueChangedBits BACnetNotificationParametersChangeOfValueNewValueChangedBits, err error) {
	m.BACnetNotificationParametersChangeOfValueNewValueContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfValueNewValueChangedBits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfValueNewValueChangedBits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	changedBits, err := ReadSimpleField[BACnetContextTagBitString](ctx, "changedBits", ReadComplex[BACnetContextTagBitString](BACnetContextTagParseWithBufferProducer[BACnetContextTagBitString]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_BIT_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'changedBits' field"))
	}
	m.ChangedBits = changedBits

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfValueNewValueChangedBits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfValueNewValueChangedBits")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfValueNewValueChangedBits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfValueNewValueChangedBits")
		}

		if err := WriteSimpleField[BACnetContextTagBitString](ctx, "changedBits", m.GetChangedBits(), WriteComplex[BACnetContextTagBitString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'changedBits' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfValueNewValueChangedBits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfValueNewValueChangedBits")
		}
		return nil
	}
	return m.BACnetNotificationParametersChangeOfValueNewValueContract.(*_BACnetNotificationParametersChangeOfValueNewValue).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) IsBACnetNotificationParametersChangeOfValueNewValueChangedBits() {
}

func (m *_BACnetNotificationParametersChangeOfValueNewValueChangedBits) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
