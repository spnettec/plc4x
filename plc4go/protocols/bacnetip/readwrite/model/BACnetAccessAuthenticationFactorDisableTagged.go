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


// BACnetAccessAuthenticationFactorDisableTagged is the corresponding interface of BACnetAccessAuthenticationFactorDisableTagged
type BACnetAccessAuthenticationFactorDisableTagged interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetAccessAuthenticationFactorDisable
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
}

// BACnetAccessAuthenticationFactorDisableTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetAccessAuthenticationFactorDisableTagged.
// This is useful for switch cases.
type BACnetAccessAuthenticationFactorDisableTaggedExactly interface {
	BACnetAccessAuthenticationFactorDisableTagged
	isBACnetAccessAuthenticationFactorDisableTagged() bool
}

// _BACnetAccessAuthenticationFactorDisableTagged is the data-structure of this message
type _BACnetAccessAuthenticationFactorDisableTagged struct {
        Header BACnetTagHeader
        Value BACnetAccessAuthenticationFactorDisable
        ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass TagClass
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetValue() BACnetAccessAuthenticationFactorDisable {
	return m.Value
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetProprietaryValue() uint32 {
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

func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetIsProprietary() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetValue()) == (BACnetAccessAuthenticationFactorDisable_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetAccessAuthenticationFactorDisableTagged factory function for _BACnetAccessAuthenticationFactorDisableTagged
func NewBACnetAccessAuthenticationFactorDisableTagged( header BACnetTagHeader , value BACnetAccessAuthenticationFactorDisable , proprietaryValue uint32 , tagNumber uint8 , tagClass TagClass ) *_BACnetAccessAuthenticationFactorDisableTagged {
return &_BACnetAccessAuthenticationFactorDisableTagged{ Header: header , Value: value , ProprietaryValue: proprietaryValue , TagNumber: tagNumber , TagClass: tagClass }
}

// Deprecated: use the interface for direct cast
func CastBACnetAccessAuthenticationFactorDisableTagged(structType any) BACnetAccessAuthenticationFactorDisableTagged {
    if casted, ok := structType.(BACnetAccessAuthenticationFactorDisableTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAccessAuthenticationFactorDisableTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetTypeName() string {
	return "BACnetAccessAuthenticationFactorDisableTagged"
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetLengthInBits(ctx context.Context) uint16 {
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


func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetAccessAuthenticationFactorDisableTaggedParse(ctx context.Context, theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetAccessAuthenticationFactorDisableTagged, error) {
	return BACnetAccessAuthenticationFactorDisableTaggedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetAccessAuthenticationFactorDisableTaggedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetAccessAuthenticationFactorDisableTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetAccessAuthenticationFactorDisableTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAccessAuthenticationFactorDisableTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
_header, _headerErr := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetAccessAuthenticationFactorDisableTagged")
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
	_value, _valueErr := ReadEnumGeneric(ctx, readBuffer, header.GetActualLength(), BACnetAccessAuthenticationFactorDisable_VENDOR_PROPRIETARY_VALUE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetAccessAuthenticationFactorDisableTagged")
	}
	var value BACnetAccessAuthenticationFactorDisable
	if _value != nil {
            value = _value.(BACnetAccessAuthenticationFactorDisable)
	}

	// Virtual field
	_isProprietary := bool((value) == (BACnetAccessAuthenticationFactorDisable_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)
	_ = isProprietary

	// Manual Field (proprietaryValue)
	_proprietaryValue, _proprietaryValueErr := ReadProprietaryEnumGeneric(ctx, readBuffer, header.GetActualLength(), isProprietary)
	if _proprietaryValueErr != nil {
		return nil, errors.Wrap(_proprietaryValueErr, "Error parsing 'proprietaryValue' field of BACnetAccessAuthenticationFactorDisableTagged")
	}
	var proprietaryValue uint32
	if _proprietaryValue != nil {
            proprietaryValue = _proprietaryValue.(uint32)
	}

	if closeErr := readBuffer.CloseContext("BACnetAccessAuthenticationFactorDisableTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAccessAuthenticationFactorDisableTagged")
	}

	// Create the instance
	return &_BACnetAccessAuthenticationFactorDisableTagged{
            TagNumber: tagNumber,
            TagClass: tagClass,
			Header: header,
			Value: value,
			ProprietaryValue: proprietaryValue,
		}, nil
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("BACnetAccessAuthenticationFactorDisableTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAccessAuthenticationFactorDisableTagged")
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
	if _isProprietaryErr := writeBuffer.WriteVirtual(ctx, "isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	// Manual Field (proprietaryValue)
	_proprietaryValueErr := WriteProprietaryEnumGeneric(ctx, writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	if _proprietaryValueErr != nil {
		return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAccessAuthenticationFactorDisableTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAccessAuthenticationFactorDisableTagged")
	}
	return nil
}


////
// Arguments Getter

func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetAccessAuthenticationFactorDisableTagged) GetTagClass() TagClass {
	return m.TagClass
}
//
////

func (m *_BACnetAccessAuthenticationFactorDisableTagged) isBACnetAccessAuthenticationFactorDisableTagged() bool {
	return true
}

func (m *_BACnetAccessAuthenticationFactorDisableTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



