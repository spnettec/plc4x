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


// ModbusPDUWriteSingleRegisterRequest is the corresponding interface of ModbusPDUWriteSingleRegisterRequest
type ModbusPDUWriteSingleRegisterRequest interface {
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetAddress returns Address (property field)
	GetAddress() uint16
	// GetValue returns Value (property field)
	GetValue() uint16
}

// ModbusPDUWriteSingleRegisterRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUWriteSingleRegisterRequest.
// This is useful for switch cases.
type ModbusPDUWriteSingleRegisterRequestExactly interface {
	ModbusPDUWriteSingleRegisterRequest
	isModbusPDUWriteSingleRegisterRequest() bool
}

// _ModbusPDUWriteSingleRegisterRequest is the data-structure of this message
type _ModbusPDUWriteSingleRegisterRequest struct {
	*_ModbusPDU
        Address uint16
        Value uint16
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUWriteSingleRegisterRequest)  GetErrorFlag() bool {
return bool(false)}

func (m *_ModbusPDUWriteSingleRegisterRequest)  GetFunctionFlag() uint8 {
return 0x06}

func (m *_ModbusPDUWriteSingleRegisterRequest)  GetResponse() bool {
return bool(false)}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUWriteSingleRegisterRequest) InitializeParent(parent ModbusPDU ) {}

func (m *_ModbusPDUWriteSingleRegisterRequest)  GetParent() ModbusPDU {
	return m._ModbusPDU
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteSingleRegisterRequest) GetAddress() uint16 {
	return m.Address
}

func (m *_ModbusPDUWriteSingleRegisterRequest) GetValue() uint16 {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewModbusPDUWriteSingleRegisterRequest factory function for _ModbusPDUWriteSingleRegisterRequest
func NewModbusPDUWriteSingleRegisterRequest( address uint16 , value uint16 ) *_ModbusPDUWriteSingleRegisterRequest {
	_result := &_ModbusPDUWriteSingleRegisterRequest{
		Address: address,
		Value: value,
    	_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteSingleRegisterRequest(structType interface{}) ModbusPDUWriteSingleRegisterRequest {
    if casted, ok := structType.(ModbusPDUWriteSingleRegisterRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteSingleRegisterRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteSingleRegisterRequest) GetTypeName() string {
	return "ModbusPDUWriteSingleRegisterRequest"
}

func (m *_ModbusPDUWriteSingleRegisterRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (address)
	lengthInBits += 16;

	// Simple field (value)
	lengthInBits += 16;

	return lengthInBits
}


func (m *_ModbusPDUWriteSingleRegisterRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUWriteSingleRegisterRequestParse(theBytes []byte, response bool) (ModbusPDUWriteSingleRegisterRequest, error) {
	return ModbusPDUWriteSingleRegisterRequestParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUWriteSingleRegisterRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUWriteSingleRegisterRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteSingleRegisterRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteSingleRegisterRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (address)
_address, _addressErr := readBuffer.ReadUint16("address", 16)
	if _addressErr != nil {
		return nil, errors.Wrap(_addressErr, "Error parsing 'address' field of ModbusPDUWriteSingleRegisterRequest")
	}
	address := _address

	// Simple Field (value)
_value, _valueErr := readBuffer.ReadUint16("value", 16)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of ModbusPDUWriteSingleRegisterRequest")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteSingleRegisterRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteSingleRegisterRequest")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUWriteSingleRegisterRequest{
		_ModbusPDU: &_ModbusPDU{
		},
		Address: address,
		Value: value,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUWriteSingleRegisterRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUWriteSingleRegisterRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteSingleRegisterRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteSingleRegisterRequest")
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

		if popErr := writeBuffer.PopContext("ModbusPDUWriteSingleRegisterRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUWriteSingleRegisterRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_ModbusPDUWriteSingleRegisterRequest) isModbusPDUWriteSingleRegisterRequest() bool {
	return true
}

func (m *_ModbusPDUWriteSingleRegisterRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



