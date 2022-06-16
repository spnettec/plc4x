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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// StatusRequest is the corresponding interface of StatusRequest
type StatusRequest interface {
	// GetStatusType returns StatusType (property field)
	GetStatusType() byte
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _StatusRequest is the data-structure of this message
type _StatusRequest struct {
	_StatusRequestChildRequirements
	StatusType byte
}

type _StatusRequestChildRequirements interface {
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

type StatusRequestParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child StatusRequest, serializeChildFunction func() error) error
	GetTypeName() string
}

type StatusRequestChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent StatusRequest, statusType byte)
	GetParent() *StatusRequest

	GetTypeName() string
	StatusRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StatusRequest) GetStatusType() byte {
	return m.StatusType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewStatusRequest factory function for _StatusRequest
func NewStatusRequest(statusType byte) *_StatusRequest {
	return &_StatusRequest{StatusType: statusType}
}

// Deprecated: use the interface for direct cast
func CastStatusRequest(structType interface{}) StatusRequest {
	if casted, ok := structType.(StatusRequest); ok {
		return casted
	}
	if casted, ok := structType.(*StatusRequest); ok {
		return *casted
	}
	return nil
}

func (m *_StatusRequest) GetTypeName() string {
	return "StatusRequest"
}

func (m *_StatusRequest) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_StatusRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func StatusRequestParse(readBuffer utils.ReadBuffer) (StatusRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("StatusRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StatusRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (statusType)
	currentPos = positionAware.GetPos()
	statusType, _err := readBuffer.ReadByte("statusType")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'statusType' field")
	}

	readBuffer.Reset(currentPos)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type StatusRequestChildSerializeRequirement interface {
		StatusRequest
		InitializeParent(StatusRequest, byte)
		GetParent() StatusRequest
	}
	var _childTemp interface{}
	var _child StatusRequestChildSerializeRequirement
	var typeSwitchError error
	switch {
	case statusType == 0x7A: // StatusRequestBinaryState
		_childTemp, typeSwitchError = StatusRequestBinaryStateParse(readBuffer)
		_child = _childTemp.(StatusRequestChildSerializeRequirement)
	case statusType == 0x73: // StatusRequestLevel
		_childTemp, typeSwitchError = StatusRequestLevelParse(readBuffer)
		_child = _childTemp.(StatusRequestChildSerializeRequirement)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("StatusRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StatusRequest")
	}

	// Finish initializing
	_child.InitializeParent(_child, statusType)
	return _child, nil
}

func (pm *_StatusRequest) SerializeParent(writeBuffer utils.WriteBuffer, child StatusRequest, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("StatusRequest"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for StatusRequest")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("StatusRequest"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for StatusRequest")
	}
	return nil
}

func (m *_StatusRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
