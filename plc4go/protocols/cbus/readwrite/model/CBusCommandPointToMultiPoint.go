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

// CBusCommandPointToMultiPoint is the corresponding interface of CBusCommandPointToMultiPoint
type CBusCommandPointToMultiPoint interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CBusCommand
	// GetCommand returns Command (property field)
	GetCommand() CBusPointToMultiPointCommand
}

// CBusCommandPointToMultiPointExactly can be used when we want exactly this type and not a type which fulfills CBusCommandPointToMultiPoint.
// This is useful for switch cases.
type CBusCommandPointToMultiPointExactly interface {
	CBusCommandPointToMultiPoint
	isCBusCommandPointToMultiPoint() bool
}

// _CBusCommandPointToMultiPoint is the data-structure of this message
type _CBusCommandPointToMultiPoint struct {
	*_CBusCommand
	Command CBusPointToMultiPointCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CBusCommandPointToMultiPoint) InitializeParent(parent CBusCommand, header CBusHeader) {
	m.Header = header
}

func (m *_CBusCommandPointToMultiPoint) GetParent() CBusCommand {
	return m._CBusCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusCommandPointToMultiPoint) GetCommand() CBusPointToMultiPointCommand {
	return m.Command
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusCommandPointToMultiPoint factory function for _CBusCommandPointToMultiPoint
func NewCBusCommandPointToMultiPoint(command CBusPointToMultiPointCommand, header CBusHeader, cBusOptions CBusOptions) *_CBusCommandPointToMultiPoint {
	_result := &_CBusCommandPointToMultiPoint{
		Command:      command,
		_CBusCommand: NewCBusCommand(header, cBusOptions),
	}
	_result._CBusCommand._CBusCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCBusCommandPointToMultiPoint(structType any) CBusCommandPointToMultiPoint {
	if casted, ok := structType.(CBusCommandPointToMultiPoint); ok {
		return casted
	}
	if casted, ok := structType.(*CBusCommandPointToMultiPoint); ok {
		return *casted
	}
	return nil
}

func (m *_CBusCommandPointToMultiPoint) GetTypeName() string {
	return "CBusCommandPointToMultiPoint"
}

func (m *_CBusCommandPointToMultiPoint) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (command)
	lengthInBits += m.Command.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CBusCommandPointToMultiPoint) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CBusCommandPointToMultiPointParse(ctx context.Context, theBytes []byte, cBusOptions CBusOptions) (CBusCommandPointToMultiPoint, error) {
	return CBusCommandPointToMultiPointParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func CBusCommandPointToMultiPointParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (CBusCommandPointToMultiPoint, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CBusCommandPointToMultiPoint"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusCommandPointToMultiPoint")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (command)
	if pullErr := readBuffer.PullContext("command"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for command")
	}
	_command, _commandErr := CBusPointToMultiPointCommandParseWithBuffer(ctx, readBuffer, cBusOptions)
	if _commandErr != nil {
		return nil, errors.Wrap(_commandErr, "Error parsing 'command' field of CBusCommandPointToMultiPoint")
	}
	command := _command.(CBusPointToMultiPointCommand)
	if closeErr := readBuffer.CloseContext("command"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for command")
	}

	if closeErr := readBuffer.CloseContext("CBusCommandPointToMultiPoint"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusCommandPointToMultiPoint")
	}

	// Create a partially initialized instance
	_child := &_CBusCommandPointToMultiPoint{
		_CBusCommand: &_CBusCommand{
			CBusOptions: cBusOptions,
		},
		Command: command,
	}
	_child._CBusCommand._CBusCommandChildRequirements = _child
	return _child, nil
}

func (m *_CBusCommandPointToMultiPoint) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusCommandPointToMultiPoint) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusCommandPointToMultiPoint"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusCommandPointToMultiPoint")
		}

		// Simple Field (command)
		if pushErr := writeBuffer.PushContext("command"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for command")
		}
		_commandErr := writeBuffer.WriteSerializable(ctx, m.GetCommand())
		if popErr := writeBuffer.PopContext("command"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for command")
		}
		if _commandErr != nil {
			return errors.Wrap(_commandErr, "Error serializing 'command' field")
		}

		if popErr := writeBuffer.PopContext("CBusCommandPointToMultiPoint"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusCommandPointToMultiPoint")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CBusCommandPointToMultiPoint) isCBusCommandPointToMultiPoint() bool {
	return true
}

func (m *_CBusCommandPointToMultiPoint) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
