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

// BACnetApplicationTagCharacterString is the corresponding interface of BACnetApplicationTagCharacterString
type BACnetApplicationTagCharacterString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadCharacterString
	// GetValue returns Value (virtual field)
	GetValue() string
	// IsBACnetApplicationTagCharacterString is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetApplicationTagCharacterString()
}

// _BACnetApplicationTagCharacterString is the data-structure of this message
type _BACnetApplicationTagCharacterString struct {
	BACnetApplicationTagContract
	Payload BACnetTagPayloadCharacterString
}

var _ BACnetApplicationTagCharacterString = (*_BACnetApplicationTagCharacterString)(nil)
var _ BACnetApplicationTagRequirements = (*_BACnetApplicationTagCharacterString)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetApplicationTagCharacterString) GetParent() BACnetApplicationTagContract {
	return m.BACnetApplicationTagContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagCharacterString) GetPayload() BACnetTagPayloadCharacterString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetApplicationTagCharacterString) GetValue() string {
	ctx := context.Background()
	_ = ctx
	return fmt.Sprintf("%v", m.GetPayload().GetValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetApplicationTagCharacterString factory function for _BACnetApplicationTagCharacterString
func NewBACnetApplicationTagCharacterString(header BACnetTagHeader, payload BACnetTagPayloadCharacterString) *_BACnetApplicationTagCharacterString {
	if payload == nil {
		panic("payload of type BACnetTagPayloadCharacterString for BACnetApplicationTagCharacterString must not be nil")
	}
	_result := &_BACnetApplicationTagCharacterString{
		BACnetApplicationTagContract: NewBACnetApplicationTag(header),
		Payload:                      payload,
	}
	_result.BACnetApplicationTagContract.(*_BACnetApplicationTag)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagCharacterString(structType any) BACnetApplicationTagCharacterString {
	if casted, ok := structType.(BACnetApplicationTagCharacterString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagCharacterString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagCharacterString) GetTypeName() string {
	return "BACnetApplicationTagCharacterString"
}

func (m *_BACnetApplicationTagCharacterString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetApplicationTagContract.(*_BACnetApplicationTag).getLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetApplicationTagCharacterString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetApplicationTagCharacterString) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetApplicationTag, header BACnetTagHeader) (__bACnetApplicationTagCharacterString BACnetApplicationTagCharacterString, err error) {
	m.BACnetApplicationTagContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagCharacterString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagCharacterString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	payload, err := ReadSimpleField[BACnetTagPayloadCharacterString](ctx, "payload", ReadComplex[BACnetTagPayloadCharacterString](BACnetTagPayloadCharacterStringParseWithBufferProducer((uint32)(header.GetActualLength())), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'payload' field"))
	}
	m.Payload = payload

	value, err := ReadVirtualField[string](ctx, "value", (*string)(nil), payload.GetValue())
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'value' field"))
	}
	_ = value

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagCharacterString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagCharacterString")
	}

	return m, nil
}

func (m *_BACnetApplicationTagCharacterString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagCharacterString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagCharacterString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagCharacterString")
		}

		if err := WriteSimpleField[BACnetTagPayloadCharacterString](ctx, "payload", m.GetPayload(), WriteComplex[BACnetTagPayloadCharacterString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'payload' field")
		}
		// Virtual field
		value := m.GetValue()
		_ = value
		if _valueErr := writeBuffer.WriteVirtual(ctx, "value", m.GetValue()); _valueErr != nil {
			return errors.Wrap(_valueErr, "Error serializing 'value' field")
		}

		if popErr := writeBuffer.PopContext("BACnetApplicationTagCharacterString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagCharacterString")
		}
		return nil
	}
	return m.BACnetApplicationTagContract.(*_BACnetApplicationTag).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetApplicationTagCharacterString) IsBACnetApplicationTagCharacterString() {}

func (m *_BACnetApplicationTagCharacterString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
