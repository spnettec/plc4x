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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// NPDUNetworkPriorityTagged is the corresponding interface of NPDUNetworkPriorityTagged
type NPDUNetworkPriorityTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() NPDUNetworkPriority
}

// NPDUNetworkPriorityTaggedExactly can be used when we want exactly this type and not a type which fulfills NPDUNetworkPriorityTagged.
// This is useful for switch cases.
type NPDUNetworkPriorityTaggedExactly interface {
	NPDUNetworkPriorityTagged
	isNPDUNetworkPriorityTagged() bool
}

// _NPDUNetworkPriorityTagged is the data-structure of this message
type _NPDUNetworkPriorityTagged struct {
        Header BACnetTagHeader
        Value NPDUNetworkPriority

	// Arguments.
	TagNumber uint8
	TagClass TagClass
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NPDUNetworkPriorityTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_NPDUNetworkPriorityTagged) GetValue() NPDUNetworkPriority {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewNPDUNetworkPriorityTagged factory function for _NPDUNetworkPriorityTagged
func NewNPDUNetworkPriorityTagged( header BACnetTagHeader , value NPDUNetworkPriority , tagNumber uint8 , tagClass TagClass ) *_NPDUNetworkPriorityTagged {
return &_NPDUNetworkPriorityTagged{ Header: header , Value: value , TagNumber: tagNumber , TagClass: tagClass }
}

// Deprecated: use the interface for direct cast
func CastNPDUNetworkPriorityTagged(structType interface{}) NPDUNetworkPriorityTagged {
    if casted, ok := structType.(NPDUNetworkPriorityTagged); ok {
		return casted
	}
	if casted, ok := structType.(*NPDUNetworkPriorityTagged); ok {
		return *casted
	}
	return nil
}

func (m *_NPDUNetworkPriorityTagged) GetTypeName() string {
	return "NPDUNetworkPriorityTagged"
}

func (m *_NPDUNetworkPriorityTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.GetHeader().GetActualLength()) * int32(int32(8)))

	return lengthInBits
}


func (m *_NPDUNetworkPriorityTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NPDUNetworkPriorityTaggedParse(theBytes []byte, tagNumber uint8, tagClass TagClass) (NPDUNetworkPriorityTagged, error) {
	return NPDUNetworkPriorityTaggedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func NPDUNetworkPriorityTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (NPDUNetworkPriorityTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NPDUNetworkPriorityTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NPDUNetworkPriorityTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
_header, _headerErr := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of NPDUNetworkPriorityTagged")
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
	_value, _valueErr := ReadEnumGenericFailing(readBuffer, header.GetActualLength(), NPDUNetworkPriority_LIFE_SAVETY_MESSAGE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of NPDUNetworkPriorityTagged")
	}
	var value NPDUNetworkPriority
	if _value != nil {
            value = _value.(NPDUNetworkPriority)
	}

	if closeErr := readBuffer.CloseContext("NPDUNetworkPriorityTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NPDUNetworkPriorityTagged")
	}

	// Create the instance
	return &_NPDUNetworkPriorityTagged{
            TagNumber: tagNumber,
            TagClass: tagClass,
			Header: header,
			Value: value,
		}, nil
}

func (m *_NPDUNetworkPriorityTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NPDUNetworkPriorityTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("NPDUNetworkPriorityTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for NPDUNetworkPriorityTagged")
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

	if popErr := writeBuffer.PopContext("NPDUNetworkPriorityTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for NPDUNetworkPriorityTagged")
	}
	return nil
}


////
// Arguments Getter

func (m *_NPDUNetworkPriorityTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_NPDUNetworkPriorityTagged) GetTagClass() TagClass {
	return m.TagClass
}
//
////

func (m *_NPDUNetworkPriorityTagged) isNPDUNetworkPriorityTagged() bool {
	return true
}

func (m *_NPDUNetworkPriorityTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



