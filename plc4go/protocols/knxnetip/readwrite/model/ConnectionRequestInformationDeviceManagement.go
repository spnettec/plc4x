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

// ConnectionRequestInformationDeviceManagement is the corresponding interface of ConnectionRequestInformationDeviceManagement
type ConnectionRequestInformationDeviceManagement interface {
	utils.LengthAware
	utils.Serializable
	ConnectionRequestInformation
}

// ConnectionRequestInformationDeviceManagementExactly can be used when we want exactly this type and not a type which fulfills ConnectionRequestInformationDeviceManagement.
// This is useful for switch cases.
type ConnectionRequestInformationDeviceManagementExactly interface {
	ConnectionRequestInformationDeviceManagement
	isConnectionRequestInformationDeviceManagement() bool
}

// _ConnectionRequestInformationDeviceManagement is the data-structure of this message
type _ConnectionRequestInformationDeviceManagement struct {
	*_ConnectionRequestInformation
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ConnectionRequestInformationDeviceManagement) GetConnectionType() uint8 {
	return 0x03
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ConnectionRequestInformationDeviceManagement) InitializeParent(parent ConnectionRequestInformation) {
}

func (m *_ConnectionRequestInformationDeviceManagement) GetParent() ConnectionRequestInformation {
	return m._ConnectionRequestInformation
}

// NewConnectionRequestInformationDeviceManagement factory function for _ConnectionRequestInformationDeviceManagement
func NewConnectionRequestInformationDeviceManagement() *_ConnectionRequestInformationDeviceManagement {
	_result := &_ConnectionRequestInformationDeviceManagement{
		_ConnectionRequestInformation: NewConnectionRequestInformation(),
	}
	_result._ConnectionRequestInformation._ConnectionRequestInformationChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastConnectionRequestInformationDeviceManagement(structType interface{}) ConnectionRequestInformationDeviceManagement {
	if casted, ok := structType.(ConnectionRequestInformationDeviceManagement); ok {
		return casted
	}
	if casted, ok := structType.(*ConnectionRequestInformationDeviceManagement); ok {
		return *casted
	}
	return nil
}

func (m *_ConnectionRequestInformationDeviceManagement) GetTypeName() string {
	return "ConnectionRequestInformationDeviceManagement"
}

func (m *_ConnectionRequestInformationDeviceManagement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_ConnectionRequestInformationDeviceManagement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ConnectionRequestInformationDeviceManagementParse(theBytes []byte) (ConnectionRequestInformationDeviceManagement, error) {
	return ConnectionRequestInformationDeviceManagementParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func ConnectionRequestInformationDeviceManagementParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ConnectionRequestInformationDeviceManagement, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ConnectionRequestInformationDeviceManagement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ConnectionRequestInformationDeviceManagement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ConnectionRequestInformationDeviceManagement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ConnectionRequestInformationDeviceManagement")
	}

	// Create a partially initialized instance
	_child := &_ConnectionRequestInformationDeviceManagement{
		_ConnectionRequestInformation: &_ConnectionRequestInformation{},
	}
	_child._ConnectionRequestInformation._ConnectionRequestInformationChildRequirements = _child
	return _child, nil
}

func (m *_ConnectionRequestInformationDeviceManagement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ConnectionRequestInformationDeviceManagement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ConnectionRequestInformationDeviceManagement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ConnectionRequestInformationDeviceManagement")
		}

		if popErr := writeBuffer.PopContext("ConnectionRequestInformationDeviceManagement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ConnectionRequestInformationDeviceManagement")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ConnectionRequestInformationDeviceManagement) isConnectionRequestInformationDeviceManagement() bool {
	return true
}

func (m *_ConnectionRequestInformationDeviceManagement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
