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


// ModbusPDUReadWriteMultipleHoldingRegistersResponse is the corresponding interface of ModbusPDUReadWriteMultipleHoldingRegistersResponse
type ModbusPDUReadWriteMultipleHoldingRegistersResponse interface {
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetValue returns Value (property field)
	GetValue() []byte
}

// ModbusPDUReadWriteMultipleHoldingRegistersResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReadWriteMultipleHoldingRegistersResponse.
// This is useful for switch cases.
type ModbusPDUReadWriteMultipleHoldingRegistersResponseExactly interface {
	ModbusPDUReadWriteMultipleHoldingRegistersResponse
	isModbusPDUReadWriteMultipleHoldingRegistersResponse() bool
}

// _ModbusPDUReadWriteMultipleHoldingRegistersResponse is the data-structure of this message
type _ModbusPDUReadWriteMultipleHoldingRegistersResponse struct {
	*_ModbusPDU
        Value []byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse)  GetErrorFlag() bool {
return bool(false)}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse)  GetFunctionFlag() uint8 {
return 0x17}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse)  GetResponse() bool {
return bool(true)}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) InitializeParent(parent ModbusPDU ) {}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse)  GetParent() ModbusPDU {
	return m._ModbusPDU
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewModbusPDUReadWriteMultipleHoldingRegistersResponse factory function for _ModbusPDUReadWriteMultipleHoldingRegistersResponse
func NewModbusPDUReadWriteMultipleHoldingRegistersResponse( value []byte ) *_ModbusPDUReadWriteMultipleHoldingRegistersResponse {
	_result := &_ModbusPDUReadWriteMultipleHoldingRegistersResponse{
		Value: value,
    	_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReadWriteMultipleHoldingRegistersResponse(structType interface{}) ModbusPDUReadWriteMultipleHoldingRegistersResponse {
    if casted, ok := structType.(ModbusPDUReadWriteMultipleHoldingRegistersResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReadWriteMultipleHoldingRegistersResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetTypeName() string {
	return "ModbusPDUReadWriteMultipleHoldingRegistersResponse"
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}


func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUReadWriteMultipleHoldingRegistersResponseParse(readBuffer utils.ReadBuffer, response bool) (ModbusPDUReadWriteMultipleHoldingRegistersResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUReadWriteMultipleHoldingRegistersResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReadWriteMultipleHoldingRegistersResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field of ModbusPDUReadWriteMultipleHoldingRegistersResponse")
	}
	// Byte Array field (value)
	numberOfBytesvalue := int(byteCount)
	value, _readArrayErr := readBuffer.ReadByteArray("value", numberOfBytesvalue)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'value' field of ModbusPDUReadWriteMultipleHoldingRegistersResponse")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReadWriteMultipleHoldingRegistersResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReadWriteMultipleHoldingRegistersResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReadWriteMultipleHoldingRegistersResponse{
		_ModbusPDU: &_ModbusPDU{
		},
		Value: value,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReadWriteMultipleHoldingRegistersResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReadWriteMultipleHoldingRegistersResponse")
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

		if popErr := writeBuffer.PopContext("ModbusPDUReadWriteMultipleHoldingRegistersResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReadWriteMultipleHoldingRegistersResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) isModbusPDUReadWriteMultipleHoldingRegistersResponse() bool {
	return true
}

func (m *_ModbusPDUReadWriteMultipleHoldingRegistersResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



