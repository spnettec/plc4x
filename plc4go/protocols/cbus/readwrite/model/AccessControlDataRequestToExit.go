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

// AccessControlDataRequestToExit is the corresponding interface of AccessControlDataRequestToExit
type AccessControlDataRequestToExit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	AccessControlData
}

// AccessControlDataRequestToExitExactly can be used when we want exactly this type and not a type which fulfills AccessControlDataRequestToExit.
// This is useful for switch cases.
type AccessControlDataRequestToExitExactly interface {
	AccessControlDataRequestToExit
	isAccessControlDataRequestToExit() bool
}

// _AccessControlDataRequestToExit is the data-structure of this message
type _AccessControlDataRequestToExit struct {
	*_AccessControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AccessControlDataRequestToExit) InitializeParent(parent AccessControlData, commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.NetworkId = networkId
	m.AccessPointId = accessPointId
}

func (m *_AccessControlDataRequestToExit) GetParent() AccessControlData {
	return m._AccessControlData
}

// NewAccessControlDataRequestToExit factory function for _AccessControlDataRequestToExit
func NewAccessControlDataRequestToExit(commandTypeContainer AccessControlCommandTypeContainer, networkId byte, accessPointId byte) *_AccessControlDataRequestToExit {
	_result := &_AccessControlDataRequestToExit{
		_AccessControlData: NewAccessControlData(commandTypeContainer, networkId, accessPointId),
	}
	_result._AccessControlData._AccessControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAccessControlDataRequestToExit(structType any) AccessControlDataRequestToExit {
	if casted, ok := structType.(AccessControlDataRequestToExit); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlDataRequestToExit); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlDataRequestToExit) GetTypeName() string {
	return "AccessControlDataRequestToExit"
}

func (m *_AccessControlDataRequestToExit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_AccessControlDataRequestToExit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AccessControlDataRequestToExitParse(theBytes []byte) (AccessControlDataRequestToExit, error) {
	return AccessControlDataRequestToExitParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func AccessControlDataRequestToExitParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AccessControlDataRequestToExit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AccessControlDataRequestToExit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlDataRequestToExit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AccessControlDataRequestToExit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlDataRequestToExit")
	}

	// Create a partially initialized instance
	_child := &_AccessControlDataRequestToExit{
		_AccessControlData: &_AccessControlData{},
	}
	_child._AccessControlData._AccessControlDataChildRequirements = _child
	return _child, nil
}

func (m *_AccessControlDataRequestToExit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AccessControlDataRequestToExit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AccessControlDataRequestToExit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AccessControlDataRequestToExit")
		}

		if popErr := writeBuffer.PopContext("AccessControlDataRequestToExit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AccessControlDataRequestToExit")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_AccessControlDataRequestToExit) isAccessControlDataRequestToExit() bool {
	return true
}

func (m *_AccessControlDataRequestToExit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
