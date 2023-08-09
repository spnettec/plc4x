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
	"github.com/rs/zerolog"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetLiftCarDriveStatusTagged is the corresponding interface of BACnetLiftCarDriveStatusTagged
type BACnetLiftCarDriveStatusTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetLiftCarDriveStatus
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
}

// BACnetLiftCarDriveStatusTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetLiftCarDriveStatusTagged.
// This is useful for switch cases.
type BACnetLiftCarDriveStatusTaggedExactly interface {
	BACnetLiftCarDriveStatusTagged
	isBACnetLiftCarDriveStatusTagged() bool
}

// _BACnetLiftCarDriveStatusTagged is the data-structure of this message
type _BACnetLiftCarDriveStatusTagged struct {
        Header BACnetTagHeader
        Value BACnetLiftCarDriveStatus
        ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass TagClass
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLiftCarDriveStatusTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetLiftCarDriveStatusTagged) GetValue() BACnetLiftCarDriveStatus {
	return m.Value
}

func (m *_BACnetLiftCarDriveStatusTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetLiftCarDriveStatusTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetLiftCarDriveStatus_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetLiftCarDriveStatusTagged factory function for _BACnetLiftCarDriveStatusTagged
func NewBACnetLiftCarDriveStatusTagged( header BACnetTagHeader , value BACnetLiftCarDriveStatus , proprietaryValue uint32 , tagNumber uint8 , tagClass TagClass ) *_BACnetLiftCarDriveStatusTagged {
return &_BACnetLiftCarDriveStatusTagged{ Header: header , Value: value , ProprietaryValue: proprietaryValue , TagNumber: tagNumber , TagClass: tagClass }
}

// Deprecated: use the interface for direct cast
func CastBACnetLiftCarDriveStatusTagged(structType any) BACnetLiftCarDriveStatusTagged {
    if casted, ok := structType.(BACnetLiftCarDriveStatusTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLiftCarDriveStatusTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLiftCarDriveStatusTagged) GetTypeName() string {
	return "BACnetLiftCarDriveStatusTagged"
}

func (m *_BACnetLiftCarDriveStatusTagged) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any {return int32(int32(0))}, func() any {return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8))))}).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() any {return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8))))}, func() any {return int32(int32(0))}).(int32))

	return lengthInBits
}


func (m *_BACnetLiftCarDriveStatusTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLiftCarDriveStatusTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetLiftCarDriveStatusTagged, error) {
	return BACnetLiftCarDriveStatusTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetLiftCarDriveStatusTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetLiftCarDriveStatusTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetLiftCarDriveStatusTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLiftCarDriveStatusTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
_header, _headerErr := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetLiftCarDriveStatusTagged")
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
	_value, _valueErr := ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetLiftCarDriveStatus_VENDOR_PROPRIETARY_VALUE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetLiftCarDriveStatusTagged")
	}
	var value BACnetLiftCarDriveStatus
	if _value != nil {
            value = _value.(BACnetLiftCarDriveStatus)
	}

	// Virtual field
	_isProprietary := bool((value) == (BACnetLiftCarDriveStatus_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)
	_ = isProprietary

	// Manual Field (proprietaryValue)
	_proprietaryValue, _proprietaryValueErr := ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)
	if _proprietaryValueErr != nil {
		return nil, errors.Wrap(_proprietaryValueErr, "Error parsing 'proprietaryValue' field of BACnetLiftCarDriveStatusTagged")
	}
	var proprietaryValue uint32
	if _proprietaryValue != nil {
            proprietaryValue = _proprietaryValue.(uint32)
	}

	if closeErr := readBuffer.CloseContext("BACnetLiftCarDriveStatusTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLiftCarDriveStatusTagged")
	}

	// Create the instance
	return &_BACnetLiftCarDriveStatusTagged{
            TagNumber: tagNumber,
            TagClass: tagClass,
			Header: header,
			Value: value,
			ProprietaryValue: proprietaryValue,
		}, nil
}

func (m *_BACnetLiftCarDriveStatusTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLiftCarDriveStatusTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("BACnetLiftCarDriveStatusTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLiftCarDriveStatusTagged")
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
	_valueErr := WriteEnumGeneric(ctx, writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}
	// Virtual field
	isProprietary := m.GetIsProprietary()
	_ = isProprietary
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	// Manual Field (proprietaryValue)
	_proprietaryValueErr := WriteProprietaryEnumGeneric(ctx, writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	if _proprietaryValueErr != nil {
		return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLiftCarDriveStatusTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLiftCarDriveStatusTagged")
	}
	return nil
}


////
// Arguments Getter

func (m *_BACnetLiftCarDriveStatusTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetLiftCarDriveStatusTagged) GetTagClass() TagClass {
	return m.TagClass
}
//
////

func (m *_BACnetLiftCarDriveStatusTagged) isBACnetLiftCarDriveStatusTagged() bool {
	return true
}

func (m *_BACnetLiftCarDriveStatusTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



