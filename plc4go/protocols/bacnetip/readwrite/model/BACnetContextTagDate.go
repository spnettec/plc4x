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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetContextTagDate is the corresponding interface of BACnetContextTagDate
type BACnetContextTagDate interface {
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadDate
}

// BACnetContextTagDateExactly can be used when we want exactly this type and not a type which fulfills BACnetContextTagDate.
// This is useful for switch cases.
type BACnetContextTagDateExactly interface {
	BACnetContextTagDate
	isBACnetContextTagDate() bool
}

// _BACnetContextTagDate is the data-structure of this message
type _BACnetContextTagDate struct {
	*_BACnetContextTag
        Payload BACnetTagPayloadDate
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagDate)  GetDataType() BACnetDataType {
return BACnetDataType_DATE}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagDate) InitializeParent(parent BACnetContextTag , header BACnetTagHeader ) {	m.Header = header
}

func (m *_BACnetContextTagDate)  GetParent() BACnetContextTag {
	return m._BACnetContextTag
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagDate) GetPayload() BACnetTagPayloadDate {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetContextTagDate factory function for _BACnetContextTagDate
func NewBACnetContextTagDate( payload BACnetTagPayloadDate , header BACnetTagHeader , tagNumberArgument uint8 ) *_BACnetContextTagDate {
	_result := &_BACnetContextTagDate{
		Payload: payload,
    	_BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	_result._BACnetContextTag._BACnetContextTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagDate(structType interface{}) BACnetContextTagDate {
    if casted, ok := structType.(BACnetContextTagDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagDate) GetTypeName() string {
	return "BACnetContextTagDate"
}

func (m *_BACnetContextTagDate) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetContextTagDate) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetContextTagDate) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetContextTagDateParse(readBuffer utils.ReadBuffer, tagNumberArgument uint8, dataType BACnetDataType) (BACnetContextTagDate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetContextTagDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
_payload, _payloadErr := BACnetTagPayloadDateParse(readBuffer)
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field of BACnetContextTagDate")
	}
	payload := _payload.(BACnetTagPayloadDate)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTagDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagDate")
	}

	// Create a partially initialized instance
	_child := &_BACnetContextTagDate{
		_BACnetContextTag: &_BACnetContextTag{
			TagNumberArgument: tagNumberArgument,
		},
		Payload: payload,
	}
	_child._BACnetContextTag._BACnetContextTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetContextTagDate) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagDate) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagDate")
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

		if popErr := writeBuffer.PopContext("BACnetContextTagDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagDate")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetContextTagDate) isBACnetContextTagDate() bool {
	return true
}

func (m *_BACnetContextTagDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



