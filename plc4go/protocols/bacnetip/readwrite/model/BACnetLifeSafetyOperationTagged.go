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


// BACnetLifeSafetyOperationTagged is the corresponding interface of BACnetLifeSafetyOperationTagged
type BACnetLifeSafetyOperationTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetValue returns Value (property field)
	GetValue() BACnetLifeSafetyOperation
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
}

// BACnetLifeSafetyOperationTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetLifeSafetyOperationTagged.
// This is useful for switch cases.
type BACnetLifeSafetyOperationTaggedExactly interface {
	BACnetLifeSafetyOperationTagged
	isBACnetLifeSafetyOperationTagged() bool
}

// _BACnetLifeSafetyOperationTagged is the data-structure of this message
type _BACnetLifeSafetyOperationTagged struct {
        Header BACnetTagHeader
        Value BACnetLifeSafetyOperation
        ProprietaryValue uint32

	// Arguments.
	TagNumber uint8
	TagClass TagClass
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLifeSafetyOperationTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetLifeSafetyOperationTagged) GetValue() BACnetLifeSafetyOperation {
	return m.Value
}

func (m *_BACnetLifeSafetyOperationTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetLifeSafetyOperationTagged) GetIsProprietary() bool {
	return bool(bool((m.GetValue()) == (BACnetLifeSafetyOperation_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetLifeSafetyOperationTagged factory function for _BACnetLifeSafetyOperationTagged
func NewBACnetLifeSafetyOperationTagged( header BACnetTagHeader , value BACnetLifeSafetyOperation , proprietaryValue uint32 , tagNumber uint8 , tagClass TagClass ) *_BACnetLifeSafetyOperationTagged {
return &_BACnetLifeSafetyOperationTagged{ Header: header , Value: value , ProprietaryValue: proprietaryValue , TagNumber: tagNumber , TagClass: tagClass }
}

// Deprecated: use the interface for direct cast
func CastBACnetLifeSafetyOperationTagged(structType interface{}) BACnetLifeSafetyOperationTagged {
    if casted, ok := structType.(BACnetLifeSafetyOperationTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLifeSafetyOperationTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLifeSafetyOperationTagged) GetTypeName() string {
	return "BACnetLifeSafetyOperationTagged"
}

func (m *_BACnetLifeSafetyOperationTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLifeSafetyOperationTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Manual Field (value)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() interface{} {return int32(int32(0))}, func() interface{} {return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8))))}).(int32))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() interface{} {return int32((int32(m.GetHeader().GetActualLength()) * int32(int32(8))))}, func() interface{} {return int32(int32(0))}).(int32))

	return lengthInBits
}


func (m *_BACnetLifeSafetyOperationTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLifeSafetyOperationTaggedParse(theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetLifeSafetyOperationTagged, error) {
	return BACnetLifeSafetyOperationTaggedParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetLifeSafetyOperationTaggedParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetLifeSafetyOperationTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLifeSafetyOperationTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLifeSafetyOperationTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
_header, _headerErr := BACnetTagHeaderParseWithBuffer(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetLifeSafetyOperationTagged")
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
	_value, _valueErr := ReadEnumGeneric(readBuffer, header.GetActualLength(), BACnetLifeSafetyOperation_VENDOR_PROPRIETARY_VALUE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetLifeSafetyOperationTagged")
	}
	var value BACnetLifeSafetyOperation
	if _value != nil {
            value = _value.(BACnetLifeSafetyOperation)
	}

	// Virtual field
	_isProprietary := bool((value) == (BACnetLifeSafetyOperation_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)
	_ = isProprietary

	// Manual Field (proprietaryValue)
	_proprietaryValue, _proprietaryValueErr := ReadProprietaryEnumGeneric(readBuffer, header.GetActualLength(), isProprietary)
	if _proprietaryValueErr != nil {
		return nil, errors.Wrap(_proprietaryValueErr, "Error parsing 'proprietaryValue' field of BACnetLifeSafetyOperationTagged")
	}
	var proprietaryValue uint32
	if _proprietaryValue != nil {
            proprietaryValue = _proprietaryValue.(uint32)
	}

	if closeErr := readBuffer.CloseContext("BACnetLifeSafetyOperationTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLifeSafetyOperationTagged")
	}

	// Create the instance
	return &_BACnetLifeSafetyOperationTagged{
            TagNumber: tagNumber,
            TagClass: tagClass,
			Header: header,
			Value: value,
			ProprietaryValue: proprietaryValue,
		}, nil
}

func (m *_BACnetLifeSafetyOperationTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLifeSafetyOperationTagged) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetLifeSafetyOperationTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLifeSafetyOperationTagged")
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
	// Virtual field
	if _isProprietaryErr := writeBuffer.WriteVirtual("isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	// Manual Field (proprietaryValue)
	_proprietaryValueErr := WriteProprietaryEnumGeneric(writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	if _proprietaryValueErr != nil {
		return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLifeSafetyOperationTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLifeSafetyOperationTagged")
	}
	return nil
}


////
// Arguments Getter

func (m *_BACnetLifeSafetyOperationTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetLifeSafetyOperationTagged) GetTagClass() TagClass {
	return m.TagClass
}
//
////

func (m *_BACnetLifeSafetyOperationTagged) isBACnetLifeSafetyOperationTagged() bool {
	return true
}

func (m *_BACnetLifeSafetyOperationTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



