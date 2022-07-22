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

// SecurityDataGasAlarmRaised is the corresponding interface of SecurityDataGasAlarmRaised
type SecurityDataGasAlarmRaised interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataGasAlarmRaisedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataGasAlarmRaised.
// This is useful for switch cases.
type SecurityDataGasAlarmRaisedExactly interface {
	SecurityDataGasAlarmRaised
	isSecurityDataGasAlarmRaised() bool
}

// _SecurityDataGasAlarmRaised is the data-structure of this message
type _SecurityDataGasAlarmRaised struct {
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

func (m *_SecurityDataGasAlarmRaised) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataGasAlarmRaised) GetParent() SecurityData {
	return m._SecurityData
}

// NewSecurityDataGasAlarmRaised factory function for _SecurityDataGasAlarmRaised
func NewSecurityDataGasAlarmRaised(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataGasAlarmRaised {
	_result := &_SecurityDataGasAlarmRaised{
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataGasAlarmRaised(structType interface{}) SecurityDataGasAlarmRaised {
	if casted, ok := structType.(SecurityDataGasAlarmRaised); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataGasAlarmRaised); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataGasAlarmRaised) GetTypeName() string {
	return "SecurityDataGasAlarmRaised"
}

func (m *_SecurityDataGasAlarmRaised) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataGasAlarmRaised) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_SecurityDataGasAlarmRaised) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataGasAlarmRaisedParse(readBuffer utils.ReadBuffer) (SecurityDataGasAlarmRaised, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataGasAlarmRaised"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataGasAlarmRaised")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataGasAlarmRaised"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataGasAlarmRaised")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataGasAlarmRaised{
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataGasAlarmRaised) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataGasAlarmRaised"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataGasAlarmRaised")
		}

		if popErr := writeBuffer.PopContext("SecurityDataGasAlarmRaised"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataGasAlarmRaised")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SecurityDataGasAlarmRaised) isSecurityDataGasAlarmRaised() bool {
	return true
}

func (m *_SecurityDataGasAlarmRaised) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
