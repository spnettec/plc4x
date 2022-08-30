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


// ModbusPDUWriteSingleCoilResponse is the corresponding interface of ModbusPDUWriteSingleCoilResponse
type ModbusPDUWriteSingleCoilResponse interface {
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetAddress returns Address (property field)
	GetAddress() uint16
	// GetValue returns Value (property field)
	GetValue() uint16
}

// ModbusPDUWriteSingleCoilResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUWriteSingleCoilResponse.
// This is useful for switch cases.
type ModbusPDUWriteSingleCoilResponseExactly interface {
	ModbusPDUWriteSingleCoilResponse
	isModbusPDUWriteSingleCoilResponse() bool
}

// _ModbusPDUWriteSingleCoilResponse is the data-structure of this message
type _ModbusPDUWriteSingleCoilResponse struct {
	*_ModbusPDU
        Address uint16
        Value uint16
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUWriteSingleCoilResponse)  GetErrorFlag() bool {
return bool(false)}

func (m *_ModbusPDUWriteSingleCoilResponse)  GetFunctionFlag() uint8 {
return 0x05}

func (m *_ModbusPDUWriteSingleCoilResponse)  GetResponse() bool {
return bool(true)}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUWriteSingleCoilResponse) InitializeParent(parent ModbusPDU ) {}

func (m *_ModbusPDUWriteSingleCoilResponse)  GetParent() ModbusPDU {
	return m._ModbusPDU
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteSingleCoilResponse) GetAddress() uint16 {
	return m.Address
}

func (m *_ModbusPDUWriteSingleCoilResponse) GetValue() uint16 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewModbusPDUWriteSingleCoilResponse factory function for _ModbusPDUWriteSingleCoilResponse
func NewModbusPDUWriteSingleCoilResponse( address uint16 , value uint16 ) *_ModbusPDUWriteSingleCoilResponse {
	_result := &_ModbusPDUWriteSingleCoilResponse{
		Address: address,
		Value: value,
    	_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteSingleCoilResponse(structType interface{}) ModbusPDUWriteSingleCoilResponse {
    if casted, ok := structType.(ModbusPDUWriteSingleCoilResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteSingleCoilResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteSingleCoilResponse) GetTypeName() string {
	return "ModbusPDUWriteSingleCoilResponse"
}

func (m *_ModbusPDUWriteSingleCoilResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ModbusPDUWriteSingleCoilResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (address)
	lengthInBits += 16;

	// Simple field (value)
	lengthInBits += 16;

	return lengthInBits
}


func (m *_ModbusPDUWriteSingleCoilResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUWriteSingleCoilResponseParse(readBuffer utils.ReadBuffer, response bool) (ModbusPDUWriteSingleCoilResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteSingleCoilResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteSingleCoilResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (address)
_address, _addressErr := readBuffer.ReadUint16("address", 16)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field of ModbusPDUWriteSingleCoilResponse")
	}
	address := _address

	// Simple Field (value)
_value, _valueErr := readBuffer.ReadUint16("value", 16)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of ModbusPDUWriteSingleCoilResponse")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteSingleCoilResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteSingleCoilResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUWriteSingleCoilResponse{
		_ModbusPDU: &_ModbusPDU{
		},
		Address: address,
		Value: value,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUWriteSingleCoilResponse) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteSingleCoilResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteSingleCoilResponse")
		}

	// Simple Field (address)
	address := uint16(m.GetAddress())
	_addressErr := writeBuffer.WriteUint16("address", 16, (address))
	if _addressErr != nil {
		return errors.Wrap(_addressErr, "Error serializing 'address' field")
	}

	// Simple Field (value)
	value := uint16(m.GetValue())
	_valueErr := writeBuffer.WriteUint16("value", 16, (value))
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

		if popErr := writeBuffer.PopContext("ModbusPDUWriteSingleCoilResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUWriteSingleCoilResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_ModbusPDUWriteSingleCoilResponse) isModbusPDUWriteSingleCoilResponse() bool {
	return true
}

func (m *_ModbusPDUWriteSingleCoilResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



