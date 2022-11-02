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


// APDUSimpleAck is the corresponding interface of APDUSimpleAck
type APDUSimpleAck interface {
	utils.LengthAware
	utils.Serializable
	APDU
	// GetOriginalInvokeId returns OriginalInvokeId (property field)
	GetOriginalInvokeId() uint8
	// GetServiceChoice returns ServiceChoice (property field)
	GetServiceChoice() uint8
}

// APDUSimpleAckExactly can be used when we want exactly this type and not a type which fulfills APDUSimpleAck.
// This is useful for switch cases.
type APDUSimpleAckExactly interface {
	APDUSimpleAck
	isAPDUSimpleAck() bool
}

// _APDUSimpleAck is the data-structure of this message
type _APDUSimpleAck struct {
	*_APDU
        OriginalInvokeId uint8
        ServiceChoice uint8
	// Reserved Fields
	reservedField0 *uint8
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_APDUSimpleAck)  GetApduType() ApduType {
return ApduType_SIMPLE_ACK_PDU}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_APDUSimpleAck) InitializeParent(parent APDU ) {}

func (m *_APDUSimpleAck)  GetParent() APDU {
	return m._APDU
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_APDUSimpleAck) GetOriginalInvokeId() uint8 {
	return m.OriginalInvokeId
}

func (m *_APDUSimpleAck) GetServiceChoice() uint8 {
	return m.ServiceChoice
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewAPDUSimpleAck factory function for _APDUSimpleAck
func NewAPDUSimpleAck( originalInvokeId uint8 , serviceChoice uint8 , apduLength uint16 ) *_APDUSimpleAck {
	_result := &_APDUSimpleAck{
		OriginalInvokeId: originalInvokeId,
		ServiceChoice: serviceChoice,
    	_APDU: NewAPDU(apduLength),
	}
	_result._APDU._APDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAPDUSimpleAck(structType interface{}) APDUSimpleAck {
    if casted, ok := structType.(APDUSimpleAck); ok {
		return casted
	}
	if casted, ok := structType.(*APDUSimpleAck); ok {
		return *casted
	}
	return nil
}

func (m *_APDUSimpleAck) GetTypeName() string {
	return "APDUSimpleAck"
}

func (m *_APDUSimpleAck) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_APDUSimpleAck) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 4

	// Simple field (originalInvokeId)
	lengthInBits += 8;

	// Simple field (serviceChoice)
	lengthInBits += 8;

	return lengthInBits
}


func (m *_APDUSimpleAck) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func APDUSimpleAckParse(readBuffer utils.ReadBuffer, apduLength uint16) (APDUSimpleAck, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUSimpleAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUSimpleAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 4)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of APDUSimpleAck")
		}
		if reserved != uint8(0) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value": reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (originalInvokeId)
_originalInvokeId, _originalInvokeIdErr := readBuffer.ReadUint8("originalInvokeId", 8)
	if _originalInvokeIdErr != nil {
		return nil, errors.Wrap(_originalInvokeIdErr, "Error parsing 'originalInvokeId' field of APDUSimpleAck")
	}
	originalInvokeId := _originalInvokeId

	// Simple Field (serviceChoice)
_serviceChoice, _serviceChoiceErr := readBuffer.ReadUint8("serviceChoice", 8)
	if _serviceChoiceErr != nil {
		return nil, errors.Wrap(_serviceChoiceErr, "Error parsing 'serviceChoice' field of APDUSimpleAck")
	}
	serviceChoice := _serviceChoice

	if closeErr := readBuffer.CloseContext("APDUSimpleAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUSimpleAck")
	}

	// Create a partially initialized instance
	_child := &_APDUSimpleAck{
		_APDU: &_APDU{
			ApduLength: apduLength,
		},
		OriginalInvokeId: originalInvokeId,
		ServiceChoice: serviceChoice,
		reservedField0: reservedField0,
	}
	_child._APDU._APDUChildRequirements = _child
	return _child, nil
}

func (m *_APDUSimpleAck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_APDUSimpleAck) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUSimpleAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUSimpleAck")
		}

	// Reserved Field (reserved)
	{
		var reserved uint8 = uint8(0)
		if m.reservedField0 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value": reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteUint8("reserved", 4, reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (originalInvokeId)
	originalInvokeId := uint8(m.GetOriginalInvokeId())
	_originalInvokeIdErr := writeBuffer.WriteUint8("originalInvokeId", 8, (originalInvokeId))
	if _originalInvokeIdErr != nil {
		return errors.Wrap(_originalInvokeIdErr, "Error serializing 'originalInvokeId' field")
	}

	// Simple Field (serviceChoice)
	serviceChoice := uint8(m.GetServiceChoice())
	_serviceChoiceErr := writeBuffer.WriteUint8("serviceChoice", 8, (serviceChoice))
	if _serviceChoiceErr != nil {
		return errors.Wrap(_serviceChoiceErr, "Error serializing 'serviceChoice' field")
	}

		if popErr := writeBuffer.PopContext("APDUSimpleAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUSimpleAck")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_APDUSimpleAck) isAPDUSimpleAck() bool {
	return true
}

func (m *_APDUSimpleAck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



