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

// TelephonyDataClearDiversion is the corresponding interface of TelephonyDataClearDiversion
type TelephonyDataClearDiversion interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TelephonyData
}

// TelephonyDataClearDiversionExactly can be used when we want exactly this type and not a type which fulfills TelephonyDataClearDiversion.
// This is useful for switch cases.
type TelephonyDataClearDiversionExactly interface {
	TelephonyDataClearDiversion
	isTelephonyDataClearDiversion() bool
}

// _TelephonyDataClearDiversion is the data-structure of this message
type _TelephonyDataClearDiversion struct {
	*_TelephonyData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataClearDiversion) InitializeParent(parent TelephonyData, commandTypeContainer TelephonyCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_TelephonyDataClearDiversion) GetParent() TelephonyData {
	return m._TelephonyData
}

// NewTelephonyDataClearDiversion factory function for _TelephonyDataClearDiversion
func NewTelephonyDataClearDiversion(commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyDataClearDiversion {
	_result := &_TelephonyDataClearDiversion{
		_TelephonyData: NewTelephonyData(commandTypeContainer, argument),
	}
	_result._TelephonyData._TelephonyDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataClearDiversion(structType any) TelephonyDataClearDiversion {
	if casted, ok := structType.(TelephonyDataClearDiversion); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataClearDiversion); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataClearDiversion) GetTypeName() string {
	return "TelephonyDataClearDiversion"
}

func (m *_TelephonyDataClearDiversion) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_TelephonyDataClearDiversion) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TelephonyDataClearDiversionParse(ctx context.Context, theBytes []byte) (TelephonyDataClearDiversion, error) {
	return TelephonyDataClearDiversionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TelephonyDataClearDiversionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TelephonyDataClearDiversion, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("TelephonyDataClearDiversion"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataClearDiversion")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TelephonyDataClearDiversion"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataClearDiversion")
	}

	// Create a partially initialized instance
	_child := &_TelephonyDataClearDiversion{
		_TelephonyData: &_TelephonyData{},
	}
	_child._TelephonyData._TelephonyDataChildRequirements = _child
	return _child, nil
}

func (m *_TelephonyDataClearDiversion) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataClearDiversion) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataClearDiversion"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataClearDiversion")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataClearDiversion"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataClearDiversion")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TelephonyDataClearDiversion) isTelephonyDataClearDiversion() bool {
	return true
}

func (m *_TelephonyDataClearDiversion) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
