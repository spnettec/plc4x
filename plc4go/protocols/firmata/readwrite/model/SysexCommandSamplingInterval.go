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


// SysexCommandSamplingInterval is the corresponding interface of SysexCommandSamplingInterval
type SysexCommandSamplingInterval interface {
	utils.LengthAware
	utils.Serializable
	SysexCommand
}

// SysexCommandSamplingIntervalExactly can be used when we want exactly this type and not a type which fulfills SysexCommandSamplingInterval.
// This is useful for switch cases.
type SysexCommandSamplingIntervalExactly interface {
	SysexCommandSamplingInterval
	isSysexCommandSamplingInterval() bool
}

// _SysexCommandSamplingInterval is the data-structure of this message
type _SysexCommandSamplingInterval struct {
	*_SysexCommand
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SysexCommandSamplingInterval)  GetCommandType() uint8 {
return 0x7A}

func (m *_SysexCommandSamplingInterval)  GetResponse() bool {
return false}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SysexCommandSamplingInterval) InitializeParent(parent SysexCommand ) {}

func (m *_SysexCommandSamplingInterval)  GetParent() SysexCommand {
	return m._SysexCommand
}


// NewSysexCommandSamplingInterval factory function for _SysexCommandSamplingInterval
func NewSysexCommandSamplingInterval( ) *_SysexCommandSamplingInterval {
	_result := &_SysexCommandSamplingInterval{
    	_SysexCommand: NewSysexCommand(),
	}
	_result._SysexCommand._SysexCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSysexCommandSamplingInterval(structType interface{}) SysexCommandSamplingInterval {
    if casted, ok := structType.(SysexCommandSamplingInterval); ok {
		return casted
	}
	if casted, ok := structType.(*SysexCommandSamplingInterval); ok {
		return *casted
	}
	return nil
}

func (m *_SysexCommandSamplingInterval) GetTypeName() string {
	return "SysexCommandSamplingInterval"
}

func (m *_SysexCommandSamplingInterval) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SysexCommandSamplingInterval) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}


func (m *_SysexCommandSamplingInterval) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SysexCommandSamplingIntervalParse(readBuffer utils.ReadBuffer, response bool) (SysexCommandSamplingInterval, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SysexCommandSamplingInterval"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SysexCommandSamplingInterval")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SysexCommandSamplingInterval"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SysexCommandSamplingInterval")
	}

	// Create a partially initialized instance
	_child := &_SysexCommandSamplingInterval{
		_SysexCommand: &_SysexCommand{
		},
	}
	_child._SysexCommand._SysexCommandChildRequirements = _child
	return _child, nil
}

func (m *_SysexCommandSamplingInterval) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SysexCommandSamplingInterval) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SysexCommandSamplingInterval"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SysexCommandSamplingInterval")
		}

		if popErr := writeBuffer.PopContext("SysexCommandSamplingInterval"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SysexCommandSamplingInterval")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_SysexCommandSamplingInterval) isSysexCommandSamplingInterval() bool {
	return true
}

func (m *_SysexCommandSamplingInterval) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



