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

// BACnetConstructedDataEventLogLogBuffer is the corresponding interface of BACnetConstructedDataEventLogLogBuffer
type BACnetConstructedDataEventLogLogBuffer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFloorText returns FloorText (property field)
	GetFloorText() []BACnetEventLogRecord
	// IsBACnetConstructedDataEventLogLogBuffer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataEventLogLogBuffer()
}

// _BACnetConstructedDataEventLogLogBuffer is the data-structure of this message
type _BACnetConstructedDataEventLogLogBuffer struct {
	BACnetConstructedDataContract
	FloorText []BACnetEventLogRecord
}

var _ BACnetConstructedDataEventLogLogBuffer = (*_BACnetConstructedDataEventLogLogBuffer)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataEventLogLogBuffer)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventLogLogBuffer) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_EVENT_LOG
}

func (m *_BACnetConstructedDataEventLogLogBuffer) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOG_BUFFER
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventLogLogBuffer) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEventLogLogBuffer) GetFloorText() []BACnetEventLogRecord {
	return m.FloorText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEventLogLogBuffer factory function for _BACnetConstructedDataEventLogLogBuffer
func NewBACnetConstructedDataEventLogLogBuffer(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, floorText []BACnetEventLogRecord, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEventLogLogBuffer {
	_result := &_BACnetConstructedDataEventLogLogBuffer{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FloorText:                     floorText,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventLogLogBuffer(structType any) BACnetConstructedDataEventLogLogBuffer {
	if casted, ok := structType.(BACnetConstructedDataEventLogLogBuffer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventLogLogBuffer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventLogLogBuffer) GetTypeName() string {
	return "BACnetConstructedDataEventLogLogBuffer"
}

func (m *_BACnetConstructedDataEventLogLogBuffer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.FloorText) > 0 {
		for _, element := range m.FloorText {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataEventLogLogBuffer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataEventLogLogBuffer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataEventLogLogBuffer BACnetConstructedDataEventLogLogBuffer, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventLogLogBuffer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventLogLogBuffer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	floorText, err := ReadTerminatedArrayField[BACnetEventLogRecord](ctx, "floorText", ReadComplex[BACnetEventLogRecord](BACnetEventLogRecordParseWithBuffer, readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'floorText' field"))
	}
	m.FloorText = floorText

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventLogLogBuffer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventLogLogBuffer")
	}

	return m, nil
}

func (m *_BACnetConstructedDataEventLogLogBuffer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEventLogLogBuffer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventLogLogBuffer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventLogLogBuffer")
		}

		if err := WriteComplexTypeArrayField(ctx, "floorText", m.GetFloorText(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'floorText' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventLogLogBuffer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventLogLogBuffer")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEventLogLogBuffer) IsBACnetConstructedDataEventLogLogBuffer() {}

func (m *_BACnetConstructedDataEventLogLogBuffer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
