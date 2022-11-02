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


// IdentifyReplyCommandGAVValuesCurrent is the corresponding interface of IdentifyReplyCommandGAVValuesCurrent
type IdentifyReplyCommandGAVValuesCurrent interface {
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetValues returns Values (property field)
	GetValues() []byte
}

// IdentifyReplyCommandGAVValuesCurrentExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandGAVValuesCurrent.
// This is useful for switch cases.
type IdentifyReplyCommandGAVValuesCurrentExactly interface {
	IdentifyReplyCommandGAVValuesCurrent
	isIdentifyReplyCommandGAVValuesCurrent() bool
}

// _IdentifyReplyCommandGAVValuesCurrent is the data-structure of this message
type _IdentifyReplyCommandGAVValuesCurrent struct {
	*_IdentifyReplyCommand
        Values []byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandGAVValuesCurrent)  GetAttribute() Attribute {
return Attribute_GAVValuesCurrent}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandGAVValuesCurrent) InitializeParent(parent IdentifyReplyCommand ) {}

func (m *_IdentifyReplyCommandGAVValuesCurrent)  GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandGAVValuesCurrent) GetValues() []byte {
	return m.Values
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewIdentifyReplyCommandGAVValuesCurrent factory function for _IdentifyReplyCommandGAVValuesCurrent
func NewIdentifyReplyCommandGAVValuesCurrent( values []byte , numBytes uint8 ) *_IdentifyReplyCommandGAVValuesCurrent {
	_result := &_IdentifyReplyCommandGAVValuesCurrent{
		Values: values,
    	_IdentifyReplyCommand: NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandGAVValuesCurrent(structType interface{}) IdentifyReplyCommandGAVValuesCurrent {
    if casted, ok := structType.(IdentifyReplyCommandGAVValuesCurrent); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandGAVValuesCurrent); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandGAVValuesCurrent) GetTypeName() string {
	return "IdentifyReplyCommandGAVValuesCurrent"
}

func (m *_IdentifyReplyCommandGAVValuesCurrent) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_IdentifyReplyCommandGAVValuesCurrent) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Values) > 0 {
		lengthInBits += 8 * uint16(len(m.Values))
	}

	return lengthInBits
}


func (m *_IdentifyReplyCommandGAVValuesCurrent) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandGAVValuesCurrentParse(readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandGAVValuesCurrent, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandGAVValuesCurrent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandGAVValuesCurrent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (values)
	numberOfBytesvalues := int(numBytes)
	values, _readArrayErr := readBuffer.ReadByteArray("values", numberOfBytesvalues)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'values' field of IdentifyReplyCommandGAVValuesCurrent")
	}

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandGAVValuesCurrent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandGAVValuesCurrent")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandGAVValuesCurrent{
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
		Values: values,
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandGAVValuesCurrent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandGAVValuesCurrent) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandGAVValuesCurrent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandGAVValuesCurrent")
		}

	// Array Field (values)
	// Byte Array field (values)
	if err := writeBuffer.WriteByteArray("values", m.GetValues()); err != nil {
		return errors.Wrap(err, "Error serializing 'values' field")
	}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandGAVValuesCurrent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandGAVValuesCurrent")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_IdentifyReplyCommandGAVValuesCurrent) isIdentifyReplyCommandGAVValuesCurrent() bool {
	return true
}

func (m *_IdentifyReplyCommandGAVValuesCurrent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



