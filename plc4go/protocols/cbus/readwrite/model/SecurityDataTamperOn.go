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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// SecurityDataTamperOn is the corresponding interface of SecurityDataTamperOn
type SecurityDataTamperOn interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataTamperOnExactly can be used when we want exactly this type and not a type which fulfills SecurityDataTamperOn.
// This is useful for switch cases.
type SecurityDataTamperOnExactly interface {
	SecurityDataTamperOn
	isSecurityDataTamperOn() bool
}

// _SecurityDataTamperOn is the data-structure of this message
type _SecurityDataTamperOn struct {
	*_SecurityData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataTamperOn) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataTamperOn) GetParent() SecurityData {
	return m._SecurityData
}

// NewSecurityDataTamperOn factory function for _SecurityDataTamperOn
func NewSecurityDataTamperOn(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataTamperOn {
	_result := &_SecurityDataTamperOn{
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataTamperOn(structType any) SecurityDataTamperOn {
	if casted, ok := structType.(SecurityDataTamperOn); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataTamperOn); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataTamperOn) GetTypeName() string {
	return "SecurityDataTamperOn"
}

func (m *_SecurityDataTamperOn) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataTamperOn) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataTamperOnParse(theBytes []byte) (SecurityDataTamperOn, error) {
	return SecurityDataTamperOnParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataTamperOnParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataTamperOn, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataTamperOn"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataTamperOn")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataTamperOn"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataTamperOn")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataTamperOn{
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataTamperOn) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataTamperOn) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataTamperOn"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataTamperOn")
		}

		if popErr := writeBuffer.PopContext("SecurityDataTamperOn"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataTamperOn")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataTamperOn) isSecurityDataTamperOn() bool {
	return true
}

func (m *_SecurityDataTamperOn) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
