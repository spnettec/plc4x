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

// SecurityDataMainsFailure is the corresponding interface of SecurityDataMainsFailure
type SecurityDataMainsFailure interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataMainsFailureExactly can be used when we want exactly this type and not a type which fulfills SecurityDataMainsFailure.
// This is useful for switch cases.
type SecurityDataMainsFailureExactly interface {
	SecurityDataMainsFailure
	isSecurityDataMainsFailure() bool
}

// _SecurityDataMainsFailure is the data-structure of this message
type _SecurityDataMainsFailure struct {
	*_SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataMainsFailure) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataMainsFailure) GetParent() SecurityData {
	return m._SecurityData
}

// NewSecurityDataMainsFailure factory function for _SecurityDataMainsFailure
func NewSecurityDataMainsFailure(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataMainsFailure {
	_result := &_SecurityDataMainsFailure{
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataMainsFailure(structType interface{}) SecurityDataMainsFailure {
	if casted, ok := structType.(SecurityDataMainsFailure); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataMainsFailure); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataMainsFailure) GetTypeName() string {
	return "SecurityDataMainsFailure"
}

func (m *_SecurityDataMainsFailure) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataMainsFailure) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_SecurityDataMainsFailure) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataMainsFailureParse(readBuffer utils.ReadBuffer) (SecurityDataMainsFailure, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataMainsFailure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataMainsFailure")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataMainsFailure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataMainsFailure")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataMainsFailure{
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataMainsFailure) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataMainsFailure"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataMainsFailure")
		}

		if popErr := writeBuffer.PopContext("SecurityDataMainsFailure"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataMainsFailure")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SecurityDataMainsFailure) isSecurityDataMainsFailure() bool {
	return true
}

func (m *_SecurityDataMainsFailure) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
