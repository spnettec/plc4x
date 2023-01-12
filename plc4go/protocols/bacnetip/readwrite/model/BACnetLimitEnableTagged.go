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


// BACnetLimitEnableTagged is the corresponding interface of BACnetLimitEnableTagged
type BACnetLimitEnableTagged interface {
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBitString
	// GetLowLimitEnable returns LowLimitEnable (virtual field)
	GetLowLimitEnable() bool
	// GetHighLimitEnable returns HighLimitEnable (virtual field)
	GetHighLimitEnable() bool
}

// BACnetLimitEnableTaggedExactly can be used when we want exactly this type and not a type which fulfills BACnetLimitEnableTagged.
// This is useful for switch cases.
type BACnetLimitEnableTaggedExactly interface {
	BACnetLimitEnableTagged
	isBACnetLimitEnableTagged() bool
}

// _BACnetLimitEnableTagged is the data-structure of this message
type _BACnetLimitEnableTagged struct {
        Header BACnetTagHeader
        Payload BACnetTagPayloadBitString

	// Arguments.
	TagNumber uint8
	TagClass TagClass
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLimitEnableTagged) GetHeader() BACnetTagHeader {
	return m.Header
}

func (m *_BACnetLimitEnableTagged) GetPayload() BACnetTagPayloadBitString {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetLimitEnableTagged) GetLowLimitEnable() bool {
	return bool(utils.InlineIf((bool(((len(m.GetPayload().GetData()))) > ((0)))), func() interface{} {return bool(m.GetPayload().GetData()[0])}, func() interface{} {return bool(bool(false))}).(bool))
}

func (m *_BACnetLimitEnableTagged) GetHighLimitEnable() bool {
	return bool(utils.InlineIf((bool(((len(m.GetPayload().GetData()))) > ((1)))), func() interface{} {return bool(m.GetPayload().GetData()[1])}, func() interface{} {return bool(bool(false))}).(bool))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetLimitEnableTagged factory function for _BACnetLimitEnableTagged
func NewBACnetLimitEnableTagged( header BACnetTagHeader , payload BACnetTagPayloadBitString , tagNumber uint8 , tagClass TagClass ) *_BACnetLimitEnableTagged {
return &_BACnetLimitEnableTagged{ Header: header , Payload: payload , TagNumber: tagNumber , TagClass: tagClass }
}

// Deprecated: use the interface for direct cast
func CastBACnetLimitEnableTagged(structType interface{}) BACnetLimitEnableTagged {
    if casted, ok := structType.(BACnetLimitEnableTagged); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLimitEnableTagged); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLimitEnableTagged) GetTypeName() string {
	return "BACnetLimitEnableTagged"
}

func (m *_BACnetLimitEnableTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLimitEnableTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits()

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetLimitEnableTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLimitEnableTaggedParse(theBytes []byte, tagNumber uint8, tagClass TagClass) (BACnetLimitEnableTagged, error) {
	return BACnetLimitEnableTaggedParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, tagClass)
}

func BACnetLimitEnableTaggedParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, tagClass TagClass) (BACnetLimitEnableTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLimitEnableTagged"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLimitEnableTagged")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
_header, _headerErr := BACnetTagHeaderParseWithBuffer(readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetLimitEnableTagged")
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

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
_payload, _payloadErr := BACnetTagPayloadBitStringParseWithBuffer(readBuffer , uint32( header.GetActualLength() ) )
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field of BACnetLimitEnableTagged")
	}
	payload := _payload.(BACnetTagPayloadBitString)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_lowLimitEnable := utils.InlineIf((bool(((len(payload.GetData()))) > ((0)))), func() interface{} {return bool(payload.GetData()[0])}, func() interface{} {return bool(bool(false))}).(bool)
	lowLimitEnable := bool(_lowLimitEnable)
	_ = lowLimitEnable

	// Virtual field
	_highLimitEnable := utils.InlineIf((bool(((len(payload.GetData()))) > ((1)))), func() interface{} {return bool(payload.GetData()[1])}, func() interface{} {return bool(bool(false))}).(bool)
	highLimitEnable := bool(_highLimitEnable)
	_ = highLimitEnable

	if closeErr := readBuffer.CloseContext("BACnetLimitEnableTagged"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLimitEnableTagged")
	}

	// Create the instance
	return &_BACnetLimitEnableTagged{
            TagNumber: tagNumber,
            TagClass: tagClass,
			Header: header,
			Payload: payload,
		}, nil
}

func (m *_BACnetLimitEnableTagged) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLimitEnableTagged) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetLimitEnableTagged"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLimitEnableTagged")
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

	// Simple Field (payload)
	if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for payload")
	}
	_payloadErr := writeBuffer.WriteSerializable(m.GetPayload())
	if popErr := writeBuffer.PopContext("payload"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for payload")
	}
	if _payloadErr != nil {
		return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
	}
	// Virtual field
	if _lowLimitEnableErr := writeBuffer.WriteVirtual("lowLimitEnable", m.GetLowLimitEnable()); _lowLimitEnableErr != nil {
		return errors.Wrap(_lowLimitEnableErr, "Error serializing 'lowLimitEnable' field")
	}
	// Virtual field
	if _highLimitEnableErr := writeBuffer.WriteVirtual("highLimitEnable", m.GetHighLimitEnable()); _highLimitEnableErr != nil {
		return errors.Wrap(_highLimitEnableErr, "Error serializing 'highLimitEnable' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLimitEnableTagged"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLimitEnableTagged")
	}
	return nil
}


////
// Arguments Getter

func (m *_BACnetLimitEnableTagged) GetTagNumber() uint8 {
	return m.TagNumber
}
func (m *_BACnetLimitEnableTagged) GetTagClass() TagClass {
	return m.TagClass
}
//
////

func (m *_BACnetLimitEnableTagged) isBACnetLimitEnableTagged() bool {
	return true
}

func (m *_BACnetLimitEnableTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



