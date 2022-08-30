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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetFaultTypeTagged is the corresponding interface of BACnetFaultTypeTagged
type BACnetFaultTypeTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetFaultType
}

// BACnetFaultTypeTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultTypeTagged.
// This is useful for switch cases.
type BACnetFaultTypeTaggedExactly interface {
	BACnetFaultTypeTagged
	isBACnetFaultTypeTagged() bool
}

// _BACnetFaultTypeTagged is the data-structure of this message
type _BACnetFaultTypeTagged struct {
        Header BACnetTagHeader
        Value BACnetFaultType

	// Arguments.
	TagNumber uint8
	TagClass TagClass
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultTypeTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetFaultTypeTagged) GetValue() BACnetFaultType {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetFaultTypeTagged factory function for _BACnetFaultTypeTagged
func NewBACnetFaultTypeTagged( header BACnetTagHeader , value BACnetFaultType , tagNumber uint8 , tagClass TagClass ) *_BACnetFaultTypeTagged {
return &_BACnetFaultTypeTagged{ Header: header , Value: value , TagNumber: tagNumber , TagClass: tagClass }
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultTypeTagged(structType interface{}) BACnetFaultTypeTagged {
    if casted, ok := structType.(BACnetFaultTypeTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultTypeTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultTypeTagged) GetTypeName() string {
	return "BACnetFaultTypeTagged"
}

func (m *_BACnetFaultTypeTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetFaultTypeTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}


func (m *_BACnetFaultTypeTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultTypeTaggedParse(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetFaultTypeTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultTypeTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultTypeTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
_header, _headerErr := BACnetTagHeaderParse(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetFaultTypeTagged")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if (!(bool((header.GetTagClass()) == (tagClass)))) {
		return nil, errors.WithStack(utils.ParseValidationError{"tag class doesn't match"})
	}

	// Validation
	if (!(bool((bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) || bool((bool((header.GetActualTagNumber()) == (tagNumber)))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"tagnumber doesn't match"})
	}

	// Manual Field (value)
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), BACnetFaultType_NONE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetFaultTypeTagged")
	}
	var value BACnetFaultType
	if _value != nil {
            value = _value.(BACnetFaultType)
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultTypeTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultTypeTagged")
	}

	// Create the instance
	return &_BACnetFaultTypeTagged{
            TagNumber: tagNumber,
            TagClass: tagClass,
			Header: header,
			Value: value,
		}, nil
}

func (m *_BACnetFaultTypeTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetFaultTypeTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultTypeTagged")
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

	if popErr := writeBuffer.PopContext("BACnetFaultTypeTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultTypeTagged")
	}
	return nil
}


////
// Arguments Getter

func (m *_BACnetFaultTypeTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetFaultTypeTagged) GetTagClass() TagClass {
	return m.TagClass
}
//
////

func (m *_BACnetFaultTypeTagged) isBACnetFaultTypeTagged() bool {
	return true
}

func (m *_BACnetFaultTypeTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



