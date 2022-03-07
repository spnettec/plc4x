/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type ConnectionResponseDataBlock struct {
	Child IConnectionResponseDataBlockChild
}

// The corresponding interface
type IConnectionResponseDataBlock interface {
	// GetConnectionType returns ConnectionType (discriminator field)
	GetConnectionType() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IConnectionResponseDataBlockParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IConnectionResponseDataBlock, serializeChildFunction func() error) error
	GetTypeName() string
}

type IConnectionResponseDataBlockChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *ConnectionResponseDataBlock)
	GetTypeName() string
	IConnectionResponseDataBlock
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewConnectionResponseDataBlock factory function for ConnectionResponseDataBlock
func NewConnectionResponseDataBlock() *ConnectionResponseDataBlock {
	return &ConnectionResponseDataBlock{}
}

func CastConnectionResponseDataBlock(structType interface{}) *ConnectionResponseDataBlock {
	if casted, ok := structType.(ConnectionResponseDataBlock); ok {
		return &casted
	}
	if casted, ok := structType.(*ConnectionResponseDataBlock); ok {
		return casted
	}
	return nil
}

func (m *ConnectionResponseDataBlock) GetTypeName() string {
	return "ConnectionResponseDataBlock"
}

func (m *ConnectionResponseDataBlock) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ConnectionResponseDataBlock) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *ConnectionResponseDataBlock) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8
	// Discriminator Field (connectionType)
	lengthInBits += 8

	return lengthInBits
}

func (m *ConnectionResponseDataBlock) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ConnectionResponseDataBlockParse(readBuffer utils.ReadBuffer) (*ConnectionResponseDataBlock, error) {
	if pullErr := readBuffer.PullContext("ConnectionResponseDataBlock"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Implicit Field (structureLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := readBuffer.ReadUint8("structureLength", 8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field")
	}

	// Discriminator Field (connectionType) (Used as input to a switch field)
	connectionType, _connectionTypeErr := readBuffer.ReadUint8("connectionType", 8)
	if _connectionTypeErr != nil {
		return nil, errors.Wrap(_connectionTypeErr, "Error parsing 'connectionType' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *ConnectionResponseDataBlock
	var typeSwitchError error
	switch {
	case connectionType == 0x03: // ConnectionResponseDataBlockDeviceManagement
		_parent, typeSwitchError = ConnectionResponseDataBlockDeviceManagementParse(readBuffer)
	case connectionType == 0x04: // ConnectionResponseDataBlockTunnelConnection
		_parent, typeSwitchError = ConnectionResponseDataBlockTunnelConnectionParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("ConnectionResponseDataBlock"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *ConnectionResponseDataBlock) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *ConnectionResponseDataBlock) SerializeParent(writeBuffer utils.WriteBuffer, child IConnectionResponseDataBlock, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("ConnectionResponseDataBlock"); pushErr != nil {
		return pushErr
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.GetLengthInBytes()))
	_structureLengthErr := writeBuffer.WriteUint8("structureLength", 8, (structureLength))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Discriminator Field (connectionType) (Used as input to a switch field)
	connectionType := uint8(child.GetConnectionType())
	_connectionTypeErr := writeBuffer.WriteUint8("connectionType", 8, (connectionType))

	if _connectionTypeErr != nil {
		return errors.Wrap(_connectionTypeErr, "Error serializing 'connectionType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ConnectionResponseDataBlock"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *ConnectionResponseDataBlock) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
