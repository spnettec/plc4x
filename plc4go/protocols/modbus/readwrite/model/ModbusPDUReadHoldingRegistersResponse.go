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


// ModbusPDUReadHoldingRegistersResponse is the corresponding interface of ModbusPDUReadHoldingRegistersResponse
type ModbusPDUReadHoldingRegistersResponse interface {
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetValue returns Value (property field)
	GetValue() []byte
}

// ModbusPDUReadHoldingRegistersResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadHoldingRegistersResponse.
// This is useful for switch cases.
type ModbusPDUReadHoldingRegistersResponseExactly interface {
	ModbusPDUReadHoldingRegistersResponse
	isModbusPDUReadHoldingRegistersResponse() bool
}

// _ModbusPDUReadHoldingRegistersResponse is the data-structure of this message
type _ModbusPDUReadHoldingRegistersResponse struct {
	*_ModbusPDU
        Value []byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadHoldingRegistersResponse)  GetErrorFlag() bool {
return bool(false)}

func (m *_ModbusPDUReadHoldingRegistersResponse)  GetFunctionFlag() uint8 {
return 0x03}

func (m *_ModbusPDUReadHoldingRegistersResponse)  GetResponse() bool {
return bool(true)}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadHoldingRegistersResponse) InitializeParent(parent ModbusPDU ) {}

func (m *_ModbusPDUReadHoldingRegistersResponse)  GetParent() ModbusPDU {
	return m._ModbusPDU
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadHoldingRegistersResponse) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewModbusPDUReadHoldingRegistersResponse factory function for _ModbusPDUReadHoldingRegistersResponse
func NewModbusPDUReadHoldingRegistersResponse( value []byte ) *_ModbusPDUReadHoldingRegistersResponse {
	_result := &_ModbusPDUReadHoldingRegistersResponse{
		Value: value,
    	_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadHoldingRegistersResponse(structType interface{}) ModbusPDUReadHoldingRegistersResponse {
    if casted, ok := structType.(ModbusPDUReadHoldingRegistersResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadHoldingRegistersResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadHoldingRegistersResponse) GetTypeName() string {
	return "ModbusPDUReadHoldingRegistersResponse"
}

func (m *_ModbusPDUReadHoldingRegistersResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ModbusPDUReadHoldingRegistersResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}


func (m *_ModbusPDUReadHoldingRegistersResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadHoldingRegistersResponseParse(readBuffer utils.ReadBuffer, response bool) (ModbusPDUReadHoldingRegistersResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadHoldingRegistersResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadHoldingRegistersResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field of ModbusPDUReadHoldingRegistersResponse")
	}
	// Byte Array field (value)
	numberOfBytesvalue := int(byteCount)
	value, _readArrayErr := readBuffer.ReadByteArray("value", numberOfBytesvalue)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'value' field of ModbusPDUReadHoldingRegistersResponse")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadHoldingRegistersResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadHoldingRegistersResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReadHoldingRegistersResponse{
		_ModbusPDU: &_ModbusPDU{
		},
		Value: value,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReadHoldingRegistersResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReadHoldingRegistersResponse) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadHoldingRegistersResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadHoldingRegistersResponse")
		}

	// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	byteCount := uint8(uint8(len(m.GetValue())))
	_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
	if _byteCountErr != nil {
		return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
	}

	// Array Field (value)
	// Byte Array field (value)
	if err := writeBuffer.WriteByteArray("value", m.GetValue()); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

		if popErr := writeBuffer.PopContext("ModbusPDUReadHoldingRegistersResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadHoldingRegistersResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_ModbusPDUReadHoldingRegistersResponse) isModbusPDUReadHoldingRegistersResponse() bool {
	return true
}

func (m *_ModbusPDUReadHoldingRegistersResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



