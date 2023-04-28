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

// BACnetTimerTransitionTagged is the corresponding interface of BACnetTimerTransitionTagged
type BACnetTimerTransitionTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetTimerTransition
}

// BACnetTimerTransitionTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetTimerTransitionTagged.
// This is useful for switch cases.
type BACnetTimerTransitionTaggedExactly interface {
	BACnetTimerTransitionTagged
	isBACnetTimerTransitionTagged() bool
}

// _BACnetTimerTransitionTagged is the data-structure of this message
type _BACnetTimerTransitionTagged struct {
	Header BACnetTagHeader
	Value  BACnetTimerTransition

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTimerTransitionTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetTimerTransitionTagged) GetValue() BACnetTimerTransition {
	return m.Value
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetTimerTransitionTagged factory function for _BACnetTimerTransitionTagged
func NewBACnetTimerTransitionTagged(header BACnetTagHeader, value BACnetTimerTransition, tagNumber uint8, tagClass TagClass) *_BACnetTimerTransitionTagged {
	return &_BACnetTimerTransitionTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetTimerTransitionTagged(structType any) BACnetTimerTransitionTagged {
	if casted, ok := structType.(BACnetTimerTransitionTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTimerTransitionTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTimerTransitionTagged) GetTypeName() string {
	return "BACnetTimerTransitionTagged"
}

func (m *_BACnetTimerTransitionTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetTimerTransitionTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetTimerTransitionTaggedParse(theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetTimerTransitionTagged, error) {
	return BACnetTimerTransitionTaggedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetTimerTransitionTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetTimerTransitionTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTimerTransitionTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTimerTransitionTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetTimerTransitionTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if !(bool((header.GetTagClass()) == (tagClass))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if !(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetTimerTransition_NONE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetTimerTransitionTagged")
	}
	var value BACnetTimerTransition
	if _value != nil {
		value = _value.(BACnetTimerTransition)
	}

	if closeErr := readBuffer.CloseContext("BACnetTimerTransitionTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTimerTransitionTagged")
	}

	// Create the instance
	return &_BACnetTimerTransitionTagged{
		TagNumber: tagNumber,
		TagClass:  tagClass,
		Header:    header,
		Value:     value,
	}, nil
}

func (m *_BACnetTimerTransitionTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTimerTransitionTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetTimerTransitionTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTimerTransitionTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(ctx, m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTimerTransitionTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTimerTransitionTagged")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetTimerTransitionTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetTimerTransitionTagged) GetTagClass() TagClass {
	return m.TagClass
}

//
////

func (m *_BACnetTimerTransitionTagged) isBACnetTimerTransitionTagged() bool {
	return true
}

func (m *_BACnetTimerTransitionTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
