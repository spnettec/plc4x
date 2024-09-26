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

// CALDataIdentify is the corresponding interface of CALDataIdentify
type CALDataIdentify interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CALData
	// GetAttribute returns Attribute (property field)
	GetAttribute() Attribute
	// IsCALDataIdentify is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsCALDataIdentify()
}

// _CALDataIdentify is the data-structure of this message
type _CALDataIdentify struct {
	CALDataContract
	Attribute Attribute
}

var _ CALDataIdentify = (*_CALDataIdentify)(nil)
var _ CALDataRequirements = (*_CALDataIdentify)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataIdentify) GetParent() CALDataContract {
	return m.CALDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataIdentify) GetAttribute() Attribute {
	return m.Attribute
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataIdentify factory function for _CALDataIdentify
func NewCALDataIdentify(commandTypeContainer CALCommandTypeContainer, additionalData CALData, attribute Attribute, requestContext RequestContext) *_CALDataIdentify {
	_result := &_CALDataIdentify{
		CALDataContract: NewCALData(commandTypeContainer, additionalData, requestContext),
		Attribute:       attribute,
	}
	_result.CALDataContract.(*_CALData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataIdentify(structType any) CALDataIdentify {
	if casted, ok := structType.(CALDataIdentify); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataIdentify); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataIdentify) GetTypeName() string {
	return "CALDataIdentify"
}

func (m *_CALDataIdentify) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CALDataContract.(*_CALData).getLengthInBits(ctx))

	// Simple field (attribute)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CALDataIdentify) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_CALDataIdentify) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CALData, requestContext RequestContext) (__cALDataIdentify CALDataIdentify, err error) {
	m.CALDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataIdentify"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataIdentify")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	attribute, err := ReadEnumField[Attribute](ctx, "attribute", "Attribute", ReadEnum(AttributeByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'attribute' field"))
	}
	m.Attribute = attribute

	if closeErr := readBuffer.CloseContext("CALDataIdentify"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataIdentify")
	}

	return m, nil
}

func (m *_CALDataIdentify) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataIdentify) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataIdentify"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataIdentify")
		}

		if err := WriteSimpleEnumField[Attribute](ctx, "attribute", "Attribute", m.GetAttribute(), WriteEnum[Attribute, uint8](Attribute.GetValue, Attribute.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'attribute' field")
		}

		if popErr := writeBuffer.PopContext("CALDataIdentify"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataIdentify")
		}
		return nil
	}
	return m.CALDataContract.(*_CALData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CALDataIdentify) IsCALDataIdentify() {}

func (m *_CALDataIdentify) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
