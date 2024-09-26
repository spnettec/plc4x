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

// BACnetHostAddressName is the corresponding interface of BACnetHostAddressName
type BACnetHostAddressName interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetHostAddress
	// GetName returns Name (property field)
	GetName() BACnetContextTagCharacterString
	// IsBACnetHostAddressName is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetHostAddressName()
}

// _BACnetHostAddressName is the data-structure of this message
type _BACnetHostAddressName struct {
	BACnetHostAddressContract
	Name BACnetContextTagCharacterString
}

var _ BACnetHostAddressName = (*_BACnetHostAddressName)(nil)
var _ BACnetHostAddressRequirements = (*_BACnetHostAddressName)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetHostAddressName) GetParent() BACnetHostAddressContract {
	return m.BACnetHostAddressContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetHostAddressName) GetName() BACnetContextTagCharacterString {
	return m.Name
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetHostAddressName factory function for _BACnetHostAddressName
func NewBACnetHostAddressName(peekedTagHeader BACnetTagHeader, name BACnetContextTagCharacterString) *_BACnetHostAddressName {
	if name == nil {
		panic("name of type BACnetContextTagCharacterString for BACnetHostAddressName must not be nil")
	}
	_result := &_BACnetHostAddressName{
		BACnetHostAddressContract: NewBACnetHostAddress(peekedTagHeader),
		Name:                      name,
	}
	_result.BACnetHostAddressContract.(*_BACnetHostAddress)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetHostAddressName(structType any) BACnetHostAddressName {
	if casted, ok := structType.(BACnetHostAddressName); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetHostAddressName); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetHostAddressName) GetTypeName() string {
	return "BACnetHostAddressName"
}

func (m *_BACnetHostAddressName) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetHostAddressContract.(*_BACnetHostAddress).getLengthInBits(ctx))

	// Simple field (name)
	lengthInBits += m.Name.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetHostAddressName) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetHostAddressName) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetHostAddress) (__bACnetHostAddressName BACnetHostAddressName, err error) {
	m.BACnetHostAddressContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetHostAddressName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetHostAddressName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	name, err := ReadSimpleField[BACnetContextTagCharacterString](ctx, "name", ReadComplex[BACnetContextTagCharacterString](BACnetContextTagParseWithBufferProducer[BACnetContextTagCharacterString]((uint8)(uint8(2)), (BACnetDataType)(BACnetDataType_CHARACTER_STRING)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'name' field"))
	}
	m.Name = name

	if closeErr := readBuffer.CloseContext("BACnetHostAddressName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetHostAddressName")
	}

	return m, nil
}

func (m *_BACnetHostAddressName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetHostAddressName) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetHostAddressName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetHostAddressName")
		}

		if err := WriteSimpleField[BACnetContextTagCharacterString](ctx, "name", m.GetName(), WriteComplex[BACnetContextTagCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'name' field")
		}

		if popErr := writeBuffer.PopContext("BACnetHostAddressName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetHostAddressName")
		}
		return nil
	}
	return m.BACnetHostAddressContract.(*_BACnetHostAddress).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetHostAddressName) IsBACnetHostAddressName() {}

func (m *_BACnetHostAddressName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
