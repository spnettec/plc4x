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

// BACnetConstructedDataAccessEvent is the corresponding interface of BACnetConstructedDataAccessEvent
type BACnetConstructedDataAccessEvent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAccessEvent returns AccessEvent (property field)
	GetAccessEvent() BACnetAccessEventTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccessEventTagged
	// IsBACnetConstructedDataAccessEvent is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAccessEvent()
}

// _BACnetConstructedDataAccessEvent is the data-structure of this message
type _BACnetConstructedDataAccessEvent struct {
	BACnetConstructedDataContract
	AccessEvent BACnetAccessEventTagged
}

var _ BACnetConstructedDataAccessEvent = (*_BACnetConstructedDataAccessEvent)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAccessEvent)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessEvent) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccessEvent) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCESS_EVENT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessEvent) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEvent) GetAccessEvent() BACnetAccessEventTagged {
	return m.AccessEvent
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAccessEvent) GetActualValue() BACnetAccessEventTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAccessEventTagged(m.GetAccessEvent())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessEvent factory function for _BACnetConstructedDataAccessEvent
func NewBACnetConstructedDataAccessEvent(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, accessEvent BACnetAccessEventTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessEvent {
	if accessEvent == nil {
		panic("accessEvent of type BACnetAccessEventTagged for BACnetConstructedDataAccessEvent must not be nil")
	}
	_result := &_BACnetConstructedDataAccessEvent{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AccessEvent:                   accessEvent,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessEvent(structType any) BACnetConstructedDataAccessEvent {
	if casted, ok := structType.(BACnetConstructedDataAccessEvent); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessEvent); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessEvent) GetTypeName() string {
	return "BACnetConstructedDataAccessEvent"
}

func (m *_BACnetConstructedDataAccessEvent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (accessEvent)
	lengthInBits += m.AccessEvent.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessEvent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAccessEvent) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAccessEvent BACnetConstructedDataAccessEvent, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessEvent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessEvent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	accessEvent, err := ReadSimpleField[BACnetAccessEventTagged](ctx, "accessEvent", ReadComplex[BACnetAccessEventTagged](BACnetAccessEventTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'accessEvent' field"))
	}
	m.AccessEvent = accessEvent

	actualValue, err := ReadVirtualField[BACnetAccessEventTagged](ctx, "actualValue", (*BACnetAccessEventTagged)(nil), accessEvent)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessEvent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessEvent")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAccessEvent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAccessEvent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessEvent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessEvent")
		}

		if err := WriteSimpleField[BACnetAccessEventTagged](ctx, "accessEvent", m.GetAccessEvent(), WriteComplex[BACnetAccessEventTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'accessEvent' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessEvent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessEvent")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessEvent) IsBACnetConstructedDataAccessEvent() {}

func (m *_BACnetConstructedDataAccessEvent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
