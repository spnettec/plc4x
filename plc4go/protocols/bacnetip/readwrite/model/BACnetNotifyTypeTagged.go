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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetNotifyTypeTagged is the corresponding interface of BACnetNotifyTypeTagged
type BACnetNotifyTypeTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetNotifyType
}

// BACnetNotifyTypeTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetNotifyTypeTagged.
// This is useful for switch cases.
type BACnetNotifyTypeTaggedExactly interface {
	BACnetNotifyTypeTagged
	isBACnetNotifyTypeTagged() bool
}

// _BACnetNotifyTypeTagged is the data-structure of this message
type _BACnetNotifyTypeTagged struct {
	Header BACnetTagHeader
	Value  BACnetNotifyType

	// Arguments.
	TagNumber uint8
	TagClass  TagClass
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotifyTypeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetNotifyTypeTagged) GetValue() BACnetNotifyType {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotifyTypeTagged factory function for _BACnetNotifyTypeTagged
func NewBACnetNotifyTypeTagged(header BACnetTagHeader, value BACnetNotifyType, tagNumber uint8, tagClass TagClass) *_BACnetNotifyTypeTagged {
	return &_BACnetNotifyTypeTagged{Header: header, Value: value, TagNumber: tagNumber, TagClass: tagClass}
}

// Deprecated: use the interface for direct cast
func CastBACnetNotifyTypeTagged(structType interface{}) BACnetNotifyTypeTagged {
	if casted, ok := structType.(BACnetNotifyTypeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotifyTypeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotifyTypeTagged) GetTypeName() string {
	return "BACnetNotifyTypeTagged"
}

func (m *_BACnetNotifyTypeTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetNotifyTypeTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}

func (m *_BACnetNotifyTypeTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotifyTypeTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetNotifyTypeTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotifyTypeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotifyTypeTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
	_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field")
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
	if !(bool(bool(bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool(bool(bool((header.GetActualTagNumber()) == (tagNumber))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetNotifyType_ALARM)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value.(BACnetNotifyType)

	if closeErr := readBuffer.CloseContext("BACnetNotifyTypeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotifyTypeTagged")
	}

	// Create the instance
	return NewBACnetNotifyTypeTagged(header, value, tagNumber, tagClass), nil
}

func (m *_BACnetNotifyTypeTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetNotifyTypeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNotifyTypeTagged")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(m.GetHeader())
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

	if popErr := writeBuffer.PopContext("BACnetNotifyTypeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNotifyTypeTagged")
	}
	return nil
}

func (m *_BACnetNotifyTypeTagged) isBACnetNotifyTypeTagged() bool {
	return true
}

func (m *_BACnetNotifyTypeTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
