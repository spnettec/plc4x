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


// ApduControlAck is the corresponding interface of ApduControlAck
type ApduControlAck interface {
	utils.LengthAware
	utils.Serializable
	ApduControl
}

// ApduControlAckExactly can be used when we want exactly this type and not a type which fulfills ApduControlAck.
// This is useful for switch cases.
type ApduControlAckExactly interface {
	ApduControlAck
	isApduControlAck() bool
}

// _ApduControlAck is the data-structure of this message
type _ApduControlAck struct {
	*_ApduControl
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduControlAck)  GetControlType() uint8 {
return 0x2}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduControlAck) InitializeParent(parent ApduControl ) {}

func (m *_ApduControlAck)  GetParent() ApduControl {
	return m._ApduControl
}


// NewApduControlAck factory function for _ApduControlAck
func NewApduControlAck( ) *_ApduControlAck {
	_result := &_ApduControlAck{
    	_ApduControl: NewApduControl(),
	}
	_result._ApduControl._ApduControlChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduControlAck(structType interface{}) ApduControlAck {
    if casted, ok := structType.(ApduControlAck); ok {
		return casted
	}
	if casted, ok := structType.(*ApduControlAck); ok {
		return *casted
	}
	return nil
}

func (m *_ApduControlAck) GetTypeName() string {
	return "ApduControlAck"
}

func (m *_ApduControlAck) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduControlAck) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}


func (m *_ApduControlAck) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduControlAckParse(theBytes []byte) (ApduControlAck, error) {
	return ApduControlAckParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func ApduControlAckParseWithBuffer(readBuffer utils.ReadBuffer) (ApduControlAck, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduControlAck"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduControlAck")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduControlAck"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduControlAck")
	}

	// Create a partially initialized instance
	_child := &_ApduControlAck{
		_ApduControl: &_ApduControl{
		},
	}
	_child._ApduControl._ApduControlChildRequirements = _child
	return _child, nil
}

func (m *_ApduControlAck) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduControlAck) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduControlAck"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduControlAck")
		}

		if popErr := writeBuffer.PopContext("ApduControlAck"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduControlAck")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_ApduControlAck) isApduControlAck() bool {
	return true
}

func (m *_ApduControlAck) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



