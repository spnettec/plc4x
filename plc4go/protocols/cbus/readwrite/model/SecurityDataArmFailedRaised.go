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

// SecurityDataArmFailedRaised is the corresponding interface of SecurityDataArmFailedRaised
type SecurityDataArmFailedRaised interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataArmFailedRaisedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataArmFailedRaised.
// This is useful for switch cases.
type SecurityDataArmFailedRaisedExactly interface {
	SecurityDataArmFailedRaised
	isSecurityDataArmFailedRaised() bool
}

// _SecurityDataArmFailedRaised is the data-structure of this message
type _SecurityDataArmFailedRaised struct {
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

func (m *_SecurityDataArmFailedRaised) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataArmFailedRaised) GetParent() SecurityData {
	return m._SecurityData
}

// NewSecurityDataArmFailedRaised factory function for _SecurityDataArmFailedRaised
func NewSecurityDataArmFailedRaised(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataArmFailedRaised {
	_result := &_SecurityDataArmFailedRaised{
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataArmFailedRaised(structType interface{}) SecurityDataArmFailedRaised {
	if casted, ok := structType.(SecurityDataArmFailedRaised); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataArmFailedRaised); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataArmFailedRaised) GetTypeName() string {
	return "SecurityDataArmFailedRaised"
}

func (m *_SecurityDataArmFailedRaised) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataArmFailedRaised) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_SecurityDataArmFailedRaised) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataArmFailedRaisedParse(theBytes []byte) (SecurityDataArmFailedRaised, error) {
	return SecurityDataArmFailedRaisedParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataArmFailedRaisedParseWithBuffer(readBuffer utils.ReadBuffer) (SecurityDataArmFailedRaised, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataArmFailedRaised"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataArmFailedRaised")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataArmFailedRaised"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataArmFailedRaised")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataArmFailedRaised{
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataArmFailedRaised) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataArmFailedRaised) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataArmFailedRaised"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataArmFailedRaised")
		}

		if popErr := writeBuffer.PopContext("SecurityDataArmFailedRaised"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataArmFailedRaised")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SecurityDataArmFailedRaised) isSecurityDataArmFailedRaised() bool {
	return true
}

func (m *_SecurityDataArmFailedRaised) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
