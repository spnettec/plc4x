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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetDateTimeEnclosed is the corresponding interface of BACnetDateTimeEnclosed
type BACnetDateTimeEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetDateTimeValue returns DateTimeValue (property field)
	GetDateTimeValue() BACnetDateTime
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetDateTimeEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetDateTimeEnclosed.
// This is useful for switch cases.
type BACnetDateTimeEnclosedExactly interface {
	BACnetDateTimeEnclosed
	isBACnetDateTimeEnclosed() bool
}

// _BACnetDateTimeEnclosed is the data-structure of this message
type _BACnetDateTimeEnclosed struct {
	OpeningTag    BACnetOpeningTag
	DateTimeValue BACnetDateTime
	ClosingTag    BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDateTimeEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetDateTimeEnclosed) GetDateTimeValue() BACnetDateTime {
	return m.DateTimeValue
}

func (m *_BACnetDateTimeEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDateTimeEnclosed factory function for _BACnetDateTimeEnclosed
func NewBACnetDateTimeEnclosed(openingTag BACnetOpeningTag, dateTimeValue BACnetDateTime, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetDateTimeEnclosed {
	return &_BACnetDateTimeEnclosed{OpeningTag: openingTag, DateTimeValue: dateTimeValue, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetDateTimeEnclosed(structType any) BACnetDateTimeEnclosed {
	if casted, ok := structType.(BACnetDateTimeEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDateTimeEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDateTimeEnclosed) GetTypeName() string {
	return "BACnetDateTimeEnclosed"
}

func (m *_BACnetDateTimeEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (dateTimeValue)
	lengthInBits += m.DateTimeValue.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetDateTimeEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDateTimeEnclosedParse(theBytes []byte, tagNumber uint8) (BACnetDateTimeEnclosed, error) {
	return BACnetDateTimeEnclosedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetDateTimeEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetDateTimeEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDateTimeEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDateTimeEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetDateTimeEnclosed")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (dateTimeValue)
	if pullErr := readBuffer.PullContext("dateTimeValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateTimeValue")
	}
	_dateTimeValue, _dateTimeValueErr := BACnetDateTimeParseWithBuffer(ctx, readBuffer)
	if _dateTimeValueErr != nil {
		return nil, errors.Wrap(_dateTimeValueErr, "Error parsing 'dateTimeValue' field of BACnetDateTimeEnclosed")
	}
	dateTimeValue := _dateTimeValue.(BACnetDateTime)
	if closeErr := readBuffer.CloseContext("dateTimeValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateTimeValue")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetDateTimeEnclosed")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetDateTimeEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDateTimeEnclosed")
	}

	// Create the instance
	return &_BACnetDateTimeEnclosed{
		TagNumber:     tagNumber,
		OpeningTag:    openingTag,
		DateTimeValue: dateTimeValue,
		ClosingTag:    closingTag,
	}, nil
}

func (m *_BACnetDateTimeEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDateTimeEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetDateTimeEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDateTimeEnclosed")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (dateTimeValue)
	if pushErr := writeBuffer.PushContext("dateTimeValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for dateTimeValue")
	}
	_dateTimeValueErr := writeBuffer.WriteSerializable(ctx, m.GetDateTimeValue())
	if popErr := writeBuffer.PopContext("dateTimeValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for dateTimeValue")
	}
	if _dateTimeValueErr != nil {
		return errors.Wrap(_dateTimeValueErr, "Error serializing 'dateTimeValue' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDateTimeEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDateTimeEnclosed")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetDateTimeEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetDateTimeEnclosed) isBACnetDateTimeEnclosed() bool {
	return true
}

func (m *_BACnetDateTimeEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
