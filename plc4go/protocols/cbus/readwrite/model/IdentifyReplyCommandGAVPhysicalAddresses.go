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


// IdentifyReplyCommandGAVPhysicalAddresses is the corresponding interface of IdentifyReplyCommandGAVPhysicalAddresses
type IdentifyReplyCommandGAVPhysicalAddresses interface {
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetValues returns Values (property field)
	GetValues() []byte
}

// IdentifyReplyCommandGAVPhysicalAddressesExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandGAVPhysicalAddresses.
// This is useful for switch cases.
type IdentifyReplyCommandGAVPhysicalAddressesExactly interface {
	IdentifyReplyCommandGAVPhysicalAddresses
	isIdentifyReplyCommandGAVPhysicalAddresses() bool
}

// _IdentifyReplyCommandGAVPhysicalAddresses is the data-structure of this message
type _IdentifyReplyCommandGAVPhysicalAddresses struct {
	*_IdentifyReplyCommand
        Values []byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandGAVPhysicalAddresses)  GetAttribute() Attribute {
return Attribute_GAVPhysicalAddresses}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandGAVPhysicalAddresses) InitializeParent(parent IdentifyReplyCommand ) {}

func (m *_IdentifyReplyCommandGAVPhysicalAddresses)  GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandGAVPhysicalAddresses) GetValues() []byte {
	return m.Values
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewIdentifyReplyCommandGAVPhysicalAddresses factory function for _IdentifyReplyCommandGAVPhysicalAddresses
func NewIdentifyReplyCommandGAVPhysicalAddresses( values []byte , numBytes uint8 ) *_IdentifyReplyCommandGAVPhysicalAddresses {
	_result := &_IdentifyReplyCommandGAVPhysicalAddresses{
		Values: values,
    	_IdentifyReplyCommand: NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandGAVPhysicalAddresses(structType interface{}) IdentifyReplyCommandGAVPhysicalAddresses {
    if casted, ok := structType.(IdentifyReplyCommandGAVPhysicalAddresses); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandGAVPhysicalAddresses); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandGAVPhysicalAddresses) GetTypeName() string {
	return "IdentifyReplyCommandGAVPhysicalAddresses"
}

func (m *_IdentifyReplyCommandGAVPhysicalAddresses) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_IdentifyReplyCommandGAVPhysicalAddresses) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Values) > 0 {
		lengthInBits += 8 * uint16(len(m.Values))
	}

	return lengthInBits
}


func (m *_IdentifyReplyCommandGAVPhysicalAddresses) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandGAVPhysicalAddressesParse(readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandGAVPhysicalAddresses, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandGAVPhysicalAddresses"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandGAVPhysicalAddresses")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (values)
	numberOfBytesvalues := int(numBytes)
	values, _readArrayErr := readBuffer.ReadByteArray("values", numberOfBytesvalues)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'values' field of IdentifyReplyCommandGAVPhysicalAddresses")
	}

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandGAVPhysicalAddresses"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandGAVPhysicalAddresses")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandGAVPhysicalAddresses{
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
		Values: values,
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandGAVPhysicalAddresses) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandGAVPhysicalAddresses) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandGAVPhysicalAddresses"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandGAVPhysicalAddresses")
		}

	// Array Field (values)
	// Byte Array field (values)
	if err := writeBuffer.WriteByteArray("values", m.GetValues()); err != nil {
		return errors.Wrap(err, "Error serializing 'values' field")
	}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandGAVPhysicalAddresses"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandGAVPhysicalAddresses")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_IdentifyReplyCommandGAVPhysicalAddresses) isIdentifyReplyCommandGAVPhysicalAddresses() bool {
	return true
}

func (m *_IdentifyReplyCommandGAVPhysicalAddresses) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



