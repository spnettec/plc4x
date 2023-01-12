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


// ModbusPDUWriteMultipleHoldingRegistersRequest is the corresponding interface of ModbusPDUWriteMultipleHoldingRegistersRequest
type ModbusPDUWriteMultipleHoldingRegistersRequest interface {
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetStartingAddress returns StartingAddress (property field)
	GetStartingAddress() uint16
	// GetQuantity returns Quantity (property field)
	GetQuantity() uint16
	// GetValue returns Value (property field)
	GetValue() []byte
}

// ModbusPDUWriteMultipleHoldingRegistersRequestExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUWriteMultipleHoldingRegistersRequest.
// This is useful for switch cases.
type ModbusPDUWriteMultipleHoldingRegistersRequestExactly interface {
	ModbusPDUWriteMultipleHoldingRegistersRequest
	isModbusPDUWriteMultipleHoldingRegistersRequest() bool
}

// _ModbusPDUWriteMultipleHoldingRegistersRequest is the data-structure of this message
type _ModbusPDUWriteMultipleHoldingRegistersRequest struct {
	*_ModbusPDU
        StartingAddress uint16
        Quantity uint16
        Value []byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUWriteMultipleHoldingRegistersRequest)  GetErrorFlag() bool {
return bool(false)}

func (m *_ModbusPDUWriteMultipleHoldingRegistersRequest)  GetFunctionFlag() uint8 {
return 0x10}

func (m *_ModbusPDUWriteMultipleHoldingRegistersRequest)  GetResponse() bool {
return bool(false)}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUWriteMultipleHoldingRegistersRequest) InitializeParent(parent ModbusPDU ) {}

func (m *_ModbusPDUWriteMultipleHoldingRegistersRequest)  GetParent() ModbusPDU {
	return m._ModbusPDU
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUWriteMultipleHoldingRegistersRequest) GetStartingAddress() uint16 {
	return m.StartingAddress
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersRequest) GetQuantity() uint16 {
	return m.Quantity
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersRequest) GetValue() []byte {
	return m.Value
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewModbusPDUWriteMultipleHoldingRegistersRequest factory function for _ModbusPDUWriteMultipleHoldingRegistersRequest
func NewModbusPDUWriteMultipleHoldingRegistersRequest( startingAddress uint16 , quantity uint16 , value []byte ) *_ModbusPDUWriteMultipleHoldingRegistersRequest {
	_result := &_ModbusPDUWriteMultipleHoldingRegistersRequest{
		StartingAddress: startingAddress,
		Quantity: quantity,
		Value: value,
    	_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUWriteMultipleHoldingRegistersRequest(structType interface{}) ModbusPDUWriteMultipleHoldingRegistersRequest {
    if casted, ok := structType.(ModbusPDUWriteMultipleHoldingRegistersRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUWriteMultipleHoldingRegistersRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersRequest) GetTypeName() string {
	return "ModbusPDUWriteMultipleHoldingRegistersRequest"
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (startingAddress)
	lengthInBits += 16;

	// Simple field (quantity)
	lengthInBits += 16;

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}


func (m *_ModbusPDUWriteMultipleHoldingRegistersRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ModbusPDUWriteMultipleHoldingRegistersRequestParse(theBytes []byte, response bool) (ModbusPDUWriteMultipleHoldingRegistersRequest, error) {
	return ModbusPDUWriteMultipleHoldingRegistersRequestParseWithBuffer(utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUWriteMultipleHoldingRegistersRequestParseWithBuffer(readBuffer utils.ReadBuffer, response bool) (ModbusPDUWriteMultipleHoldingRegistersRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusPDUWriteMultipleHoldingRegistersRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUWriteMultipleHoldingRegistersRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (startingAddress)
_startingAddress, _startingAddressErr := readBuffer.ReadUint16("startingAddress", 16)
	if _startingAddressErr != nil {
		return nil, errors.Wrap(_startingAddressErr, "Error parsing 'startingAddress' field of ModbusPDUWriteMultipleHoldingRegistersRequest")
	}
	startingAddress := _startingAddress

	// Simple Field (quantity)
_quantity, _quantityErr := readBuffer.ReadUint16("quantity", 16)
	if _quantityErr != nil {
		return nil, errors.Wrap(_quantityErr, "Error parsing 'quantity' field of ModbusPDUWriteMultipleHoldingRegistersRequest")
	}
	quantity := _quantity

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field of ModbusPDUWriteMultipleHoldingRegistersRequest")
	}
	// Byte Array field (value)
	numberOfBytesvalue := int(byteCount)
	value, _readArrayErr := readBuffer.ReadByteArray("value", numberOfBytesvalue)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'value' field of ModbusPDUWriteMultipleHoldingRegistersRequest")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUWriteMultipleHoldingRegistersRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUWriteMultipleHoldingRegistersRequest")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUWriteMultipleHoldingRegistersRequest{
		_ModbusPDU: &_ModbusPDU{
		},
		StartingAddress: startingAddress,
		Quantity: quantity,
		Value: value,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersRequest) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUWriteMultipleHoldingRegistersRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUWriteMultipleHoldingRegistersRequest")
		}

	// Simple Field (startingAddress)
	startingAddress := uint16(m.GetStartingAddress())
	_startingAddressErr := writeBuffer.WriteUint16("startingAddress", 16, (startingAddress))
	if _startingAddressErr != nil {
		return errors.Wrap(_startingAddressErr, "Error serializing 'startingAddress' field")
	}

	// Simple Field (quantity)
	quantity := uint16(m.GetQuantity())
	_quantityErr := writeBuffer.WriteUint16("quantity", 16, (quantity))
	if _quantityErr != nil {
		return errors.Wrap(_quantityErr, "Error serializing 'quantity' field")
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

		if popErr := writeBuffer.PopContext("ModbusPDUWriteMultipleHoldingRegistersRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUWriteMultipleHoldingRegistersRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_ModbusPDUWriteMultipleHoldingRegistersRequest) isModbusPDUWriteMultipleHoldingRegistersRequest() bool {
	return true
}

func (m *_ModbusPDUWriteMultipleHoldingRegistersRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



