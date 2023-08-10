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
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// SecurityDataAlarmOff is the corresponding interface of SecurityDataAlarmOff
type SecurityDataAlarmOff interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataAlarmOffExactly can be used when we want exactly this type and not a type which fulfills SecurityDataAlarmOff.
// This is useful for switch cases.
type SecurityDataAlarmOffExactly interface {
	SecurityDataAlarmOff
	isSecurityDataAlarmOff() bool
}

// _SecurityDataAlarmOff is the data-structure of this message
type _SecurityDataAlarmOff struct {
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

func (m *_SecurityDataAlarmOff) InitializeParent(parent SecurityData, commandTypeContainer SecurityCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataAlarmOff) GetParent() SecurityData {
	return m._SecurityData
}

// NewSecurityDataAlarmOff factory function for _SecurityDataAlarmOff
func NewSecurityDataAlarmOff(commandTypeContainer SecurityCommandTypeContainer, argument byte) *_SecurityDataAlarmOff {
	_result := &_SecurityDataAlarmOff{
		_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataAlarmOff(structType any) SecurityDataAlarmOff {
	if casted, ok := structType.(SecurityDataAlarmOff); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataAlarmOff); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataAlarmOff) GetTypeName() string {
	return "SecurityDataAlarmOff"
}

func (m *_SecurityDataAlarmOff) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SecurityDataAlarmOff) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataAlarmOffParse(ctx context.Context, theBytes []byte) (SecurityDataAlarmOff, error) {
	return SecurityDataAlarmOffParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataAlarmOffParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataAlarmOff, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SecurityDataAlarmOff"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataAlarmOff")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataAlarmOff"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataAlarmOff")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataAlarmOff{
		_SecurityData: &_SecurityData{},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataAlarmOff) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataAlarmOff) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataAlarmOff"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataAlarmOff")
		}

		if popErr := writeBuffer.PopContext("SecurityDataAlarmOff"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataAlarmOff")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataAlarmOff) isSecurityDataAlarmOff() bool {
	return true
}

func (m *_SecurityDataAlarmOff) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
