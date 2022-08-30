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


// IdentifyReplyCommandNetworkTerminalLevels is the corresponding interface of IdentifyReplyCommandNetworkTerminalLevels
type IdentifyReplyCommandNetworkTerminalLevels interface {
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetNetworkTerminalLevels returns NetworkTerminalLevels (property field)
	GetNetworkTerminalLevels() []byte
}

// IdentifyReplyCommandNetworkTerminalLevelsExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandNetworkTerminalLevels.
// This is useful for switch cases.
type IdentifyReplyCommandNetworkTerminalLevelsExactly interface {
	IdentifyReplyCommandNetworkTerminalLevels
	isIdentifyReplyCommandNetworkTerminalLevels() bool
}

// _IdentifyReplyCommandNetworkTerminalLevels is the data-structure of this message
type _IdentifyReplyCommandNetworkTerminalLevels struct {
	*_IdentifyReplyCommand
        NetworkTerminalLevels []byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandNetworkTerminalLevels)  GetAttribute() Attribute {
return Attribute_NetworkTerminalLevels}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandNetworkTerminalLevels) InitializeParent(parent IdentifyReplyCommand ) {}

func (m *_IdentifyReplyCommandNetworkTerminalLevels)  GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandNetworkTerminalLevels) GetNetworkTerminalLevels() []byte {
	return m.NetworkTerminalLevels
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewIdentifyReplyCommandNetworkTerminalLevels factory function for _IdentifyReplyCommandNetworkTerminalLevels
func NewIdentifyReplyCommandNetworkTerminalLevels( networkTerminalLevels []byte , numBytes uint8 ) *_IdentifyReplyCommandNetworkTerminalLevels {
	_result := &_IdentifyReplyCommandNetworkTerminalLevels{
		NetworkTerminalLevels: networkTerminalLevels,
    	_IdentifyReplyCommand: NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandNetworkTerminalLevels(structType interface{}) IdentifyReplyCommandNetworkTerminalLevels {
    if casted, ok := structType.(IdentifyReplyCommandNetworkTerminalLevels); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandNetworkTerminalLevels); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandNetworkTerminalLevels) GetTypeName() string {
	return "IdentifyReplyCommandNetworkTerminalLevels"
}

func (m *_IdentifyReplyCommandNetworkTerminalLevels) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_IdentifyReplyCommandNetworkTerminalLevels) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.NetworkTerminalLevels) > 0 {
		lengthInBits += 8 * uint16(len(m.NetworkTerminalLevels))
	}

	return lengthInBits
}


func (m *_IdentifyReplyCommandNetworkTerminalLevels) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func IdentifyReplyCommandNetworkTerminalLevelsParse(readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandNetworkTerminalLevels, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandNetworkTerminalLevels"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandNetworkTerminalLevels")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (networkTerminalLevels)
	numberOfBytesnetworkTerminalLevels := int(numBytes)
	networkTerminalLevels, _readArrayErr := readBuffer.ReadByteArray("networkTerminalLevels", numberOfBytesnetworkTerminalLevels)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'networkTerminalLevels' field of IdentifyReplyCommandNetworkTerminalLevels")
	}

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandNetworkTerminalLevels"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandNetworkTerminalLevels")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandNetworkTerminalLevels{
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
		NetworkTerminalLevels: networkTerminalLevels,
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandNetworkTerminalLevels) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandNetworkTerminalLevels"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandNetworkTerminalLevels")
		}

	// Array Field (networkTerminalLevels)
	// Byte Array field (networkTerminalLevels)
	if err := writeBuffer.WriteByteArray("networkTerminalLevels", m.GetNetworkTerminalLevels()); err != nil {
		return errors.Wrap(err, "Error serializing 'networkTerminalLevels' field")
	}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandNetworkTerminalLevels"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandNetworkTerminalLevels")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_IdentifyReplyCommandNetworkTerminalLevels) isIdentifyReplyCommandNetworkTerminalLevels() bool {
	return true
}

func (m *_IdentifyReplyCommandNetworkTerminalLevels) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



