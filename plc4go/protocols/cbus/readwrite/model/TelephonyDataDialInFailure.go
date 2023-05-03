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

// TelephonyDataDialInFailure is the corresponding interface of TelephonyDataDialInFailure
type TelephonyDataDialInFailure interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	TelephonyData
	// GetReason returns Reason (property field)
	GetReason() DialInFailureReason
}

// TelephonyDataDialInFailureExactly can be used when we want exactly this type and not a type which fulfills TelephonyDataDialInFailure.
// This is useful for switch cases.
type TelephonyDataDialInFailureExactly interface {
	TelephonyDataDialInFailure
	isTelephonyDataDialInFailure() bool
}

// _TelephonyDataDialInFailure is the data-structure of this message
type _TelephonyDataDialInFailure struct {
	*_TelephonyData
	Reason DialInFailureReason
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TelephonyDataDialInFailure) InitializeParent(parent TelephonyData, commandTypeContainer TelephonyCommandTypeContainer, argument byte) {
	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_TelephonyDataDialInFailure) GetParent() TelephonyData {
	return m._TelephonyData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TelephonyDataDialInFailure) GetReason() DialInFailureReason {
	return m.Reason
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTelephonyDataDialInFailure factory function for _TelephonyDataDialInFailure
func NewTelephonyDataDialInFailure(reason DialInFailureReason, commandTypeContainer TelephonyCommandTypeContainer, argument byte) *_TelephonyDataDialInFailure {
	_result := &_TelephonyDataDialInFailure{
		Reason:         reason,
		_TelephonyData: NewTelephonyData(commandTypeContainer, argument),
	}
	_result._TelephonyData._TelephonyDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTelephonyDataDialInFailure(structType any) TelephonyDataDialInFailure {
	if casted, ok := structType.(TelephonyDataDialInFailure); ok {
		return casted
	}
	if casted, ok := structType.(*TelephonyDataDialInFailure); ok {
		return *casted
	}
	return nil
}

func (m *_TelephonyDataDialInFailure) GetTypeName() string {
	return "TelephonyDataDialInFailure"
}

func (m *_TelephonyDataDialInFailure) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (reason)
	lengthInBits += 8

	return lengthInBits
}

func (m *_TelephonyDataDialInFailure) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TelephonyDataDialInFailureParse(theBytes []byte) (TelephonyDataDialInFailure, error) {
	return TelephonyDataDialInFailureParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func TelephonyDataDialInFailureParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (TelephonyDataDialInFailure, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TelephonyDataDialInFailure"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TelephonyDataDialInFailure")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (reason)
	if pullErr := readBuffer.PullContext("reason"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for reason")
	}
	_reason, _reasonErr := DialInFailureReasonParseWithBuffer(ctx, readBuffer)
	if _reasonErr != nil {
		return nil, errors.Wrap(_reasonErr, "Error parsing 'reason' field of TelephonyDataDialInFailure")
	}
	reason := _reason
	if closeErr := readBuffer.CloseContext("reason"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for reason")
	}

	if closeErr := readBuffer.CloseContext("TelephonyDataDialInFailure"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TelephonyDataDialInFailure")
	}

	// Create a partially initialized instance
	_child := &_TelephonyDataDialInFailure{
		_TelephonyData: &_TelephonyData{},
		Reason:         reason,
	}
	_child._TelephonyData._TelephonyDataChildRequirements = _child
	return _child, nil
}

func (m *_TelephonyDataDialInFailure) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TelephonyDataDialInFailure) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TelephonyDataDialInFailure"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TelephonyDataDialInFailure")
		}

		// Simple Field (reason)
		if pushErr := writeBuffer.PushContext("reason"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for reason")
		}
		_reasonErr := writeBuffer.WriteSerializable(ctx, m.GetReason())
		if popErr := writeBuffer.PopContext("reason"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for reason")
		}
		if _reasonErr != nil {
			return errors.Wrap(_reasonErr, "Error serializing 'reason' field")
		}

		if popErr := writeBuffer.PopContext("TelephonyDataDialInFailure"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TelephonyDataDialInFailure")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TelephonyDataDialInFailure) isTelephonyDataDialInFailure() bool {
	return true
}

func (m *_TelephonyDataDialInFailure) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
