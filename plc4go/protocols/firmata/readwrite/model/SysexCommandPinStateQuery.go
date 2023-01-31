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


// SysexCommandPinStateQuery is the corresponding interface of SysexCommandPinStateQuery
type SysexCommandPinStateQuery interface {
	utils.LengthAware
	utils.Serializable
	SysexCommand
	// GetPin returns Pin (property field)
	GetPin() uint8
}

// SysexCommandPinStateQueryExactly can be used when we want exactly this type and not a type which fulfills SysexCommandPinStateQuery.
// This is useful for switch cases.
type SysexCommandPinStateQueryExactly interface {
	SysexCommandPinStateQuery
	isSysexCommandPinStateQuery() bool
}

// _SysexCommandPinStateQuery is the data-structure of this message
type _SysexCommandPinStateQuery struct {
	*_SysexCommand
        Pin uint8
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandPinStateQuery)  GetCommandType() uint8 {
return 0x6D}

func (m *_SysexCommandPinStateQuery)  GetResponse() bool {
return false}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandPinStateQuery) InitializeParent(parent SysexCommand ) {}

func (m *_SysexCommandPinStateQuery)  GetParent() SysexCommand {
	return m._SysexCommand
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SysexCommandPinStateQuery) GetPin() uint8 {
	return m.Pin
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewSysexCommandPinStateQuery factory function for _SysexCommandPinStateQuery
func NewSysexCommandPinStateQuery( pin uint8 ) *_SysexCommandPinStateQuery {
	_result := &_SysexCommandPinStateQuery{
		Pin: pin,
    	_SysexCommand: NewSysexCommand(),
	}
	_result._SysexCommand._SysexCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandPinStateQuery(structType interface{}) SysexCommandPinStateQuery {
    if casted, ok := structType.(SysexCommandPinStateQuery); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandPinStateQuery); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandPinStateQuery) GetTypeName() string {
	return "SysexCommandPinStateQuery"
}

func (m *_SysexCommandPinStateQuery) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SysexCommandPinStateQuery) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (pin)
	lengthInBits += 8;

	return lengthInBits
}


func (m *_SysexCommandPinStateQuery) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandPinStateQueryParse(theBytes []byte, response bool) (SysexCommandPinStateQuery, error) {
	return SysexCommandPinStateQueryParseWithBuffer(utils.NewReadBufferByteBased(theBytes), response)
}

func SysexCommandPinStateQueryParseWithBuffer(readBuffer utils.ReadBuffer, response bool) (SysexCommandPinStateQuery, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandPinStateQuery"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandPinStateQuery")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (pin)
_pin, _pinErr := readBuffer.ReadUint8("pin", 8)
	if _pinErr != nil {
		return nil, errors.Wrap(_pinErr, "Error parsing 'pin' field of SysexCommandPinStateQuery")
	}
	pin := _pin

	if closeErr := readBuffer.CloseContext("SysexCommandPinStateQuery"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandPinStateQuery")
	}

	// Create a partially initialized instance
	_child := &_SysexCommandPinStateQuery{
		_SysexCommand: &_SysexCommand{
		},
		Pin: pin,
	}
	_child._SysexCommand._SysexCommandChildRequirements = _child
	return _child, nil
}

func (m *_SysexCommandPinStateQuery) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandPinStateQuery) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandPinStateQuery"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandPinStateQuery")
		}

	// Simple Field (pin)
	pin := uint8(m.GetPin())
	_pinErr := writeBuffer.WriteUint8("pin", 8, (pin))
	if _pinErr != nil {
		return errors.Wrap(_pinErr, "Error serializing 'pin' field")
	}

		if popErr := writeBuffer.PopContext("SysexCommandPinStateQuery"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandPinStateQuery")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_SysexCommandPinStateQuery) isSysexCommandPinStateQuery() bool {
	return true
}

func (m *_SysexCommandPinStateQuery) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



