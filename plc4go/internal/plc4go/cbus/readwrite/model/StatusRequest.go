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
type StatusRequest struct {
	StatusType byte
	Child      IStatusRequestChild
}

// The corresponding interface
type IStatusRequest interface {
	// GetStatusType returns StatusType
	GetStatusType() byte
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type IStatusRequestParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child IStatusRequest, serializeChildFunction func() error) error
	GetTypeName() string
}

type IStatusRequestChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *StatusRequest, statusType byte)
	GetTypeName() string
	IStatusRequest
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *StatusRequest) GetStatusType() byte {
	return m.StatusType
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewStatusRequest factory function for StatusRequest
func NewStatusRequest(statusType byte) *StatusRequest {
	return &StatusRequest{StatusType: statusType}
}

func CastStatusRequest(structType interface{}) *StatusRequest {
	castFunc := func(typ interface{}) *StatusRequest {
		if casted, ok := typ.(StatusRequest); ok {
			return &casted
		}
		if casted, ok := typ.(*StatusRequest); ok {
			return casted
		}
		return nil
	}
	return castFunc(structType)
}

func (m *StatusRequest) GetTypeName() string {
	return "StatusRequest"
}

func (m *StatusRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *StatusRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *StatusRequest) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *StatusRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func StatusRequestParse(readBuffer utils.ReadBuffer) (*StatusRequest, error) {
	if pullErr := readBuffer.PullContext("StatusRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Peek Field (statusType)
	currentPos = readBuffer.GetPos()
	statusType, _err := readBuffer.ReadByte("statusType")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'statusType' field")
	}

	readBuffer.Reset(currentPos)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *StatusRequest
	var typeSwitchError error
	switch {
	case statusType == 0x7A: // StatusRequestBinaryState
		_parent, typeSwitchError = StatusRequestBinaryStateParse(readBuffer)
	case statusType == 0x73: // StatusRequestLevel
		_parent, typeSwitchError = StatusRequestLevelParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("StatusRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent, statusType)
	return _parent, nil
}

func (m *StatusRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *StatusRequest) SerializeParent(writeBuffer utils.WriteBuffer, child IStatusRequest, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("StatusRequest"); pushErr != nil {
		return pushErr
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("StatusRequest"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *StatusRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
