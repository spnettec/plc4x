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

// BACnetShedLevelPercent is the corresponding interface of BACnetShedLevelPercent
type BACnetShedLevelPercent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetShedLevel
	// GetPercent returns Percent (property field)
	GetPercent() BACnetContextTagUnsignedInteger
	// IsBACnetShedLevelPercent is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetShedLevelPercent()
}

// _BACnetShedLevelPercent is the data-structure of this message
type _BACnetShedLevelPercent struct {
	BACnetShedLevelContract
	Percent BACnetContextTagUnsignedInteger
}

var _ BACnetShedLevelPercent = (*_BACnetShedLevelPercent)(nil)
var _ BACnetShedLevelRequirements = (*_BACnetShedLevelPercent)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetShedLevelPercent) GetParent() BACnetShedLevelContract {
	return m.BACnetShedLevelContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetShedLevelPercent) GetPercent() BACnetContextTagUnsignedInteger {
	return m.Percent
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetShedLevelPercent factory function for _BACnetShedLevelPercent
func NewBACnetShedLevelPercent(peekedTagHeader BACnetTagHeader, percent BACnetContextTagUnsignedInteger) *_BACnetShedLevelPercent {
	if percent == nil {
		panic("percent of type BACnetContextTagUnsignedInteger for BACnetShedLevelPercent must not be nil")
	}
	_result := &_BACnetShedLevelPercent{
		BACnetShedLevelContract: NewBACnetShedLevel(peekedTagHeader),
		Percent:                 percent,
	}
	_result.BACnetShedLevelContract.(*_BACnetShedLevel)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetShedLevelPercent(structType any) BACnetShedLevelPercent {
	if casted, ok := structType.(BACnetShedLevelPercent); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetShedLevelPercent); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetShedLevelPercent) GetTypeName() string {
	return "BACnetShedLevelPercent"
}

func (m *_BACnetShedLevelPercent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetShedLevelContract.(*_BACnetShedLevel).getLengthInBits(ctx))

	// Simple field (percent)
	lengthInBits += m.Percent.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetShedLevelPercent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetShedLevelPercent) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetShedLevel) (__bACnetShedLevelPercent BACnetShedLevelPercent, err error) {
	m.BACnetShedLevelContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetShedLevelPercent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetShedLevelPercent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	percent, err := ReadSimpleField[BACnetContextTagUnsignedInteger](ctx, "percent", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(0)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'percent' field"))
	}
	m.Percent = percent

	if closeErr := readBuffer.CloseContext("BACnetShedLevelPercent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetShedLevelPercent")
	}

	return m, nil
}

func (m *_BACnetShedLevelPercent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetShedLevelPercent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetShedLevelPercent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetShedLevelPercent")
		}

		if err := WriteSimpleField[BACnetContextTagUnsignedInteger](ctx, "percent", m.GetPercent(), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'percent' field")
		}

		if popErr := writeBuffer.PopContext("BACnetShedLevelPercent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetShedLevelPercent")
		}
		return nil
	}
	return m.BACnetShedLevelContract.(*_BACnetShedLevel).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetShedLevelPercent) IsBACnetShedLevelPercent() {}

func (m *_BACnetShedLevelPercent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
