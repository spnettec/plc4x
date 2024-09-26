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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyStatesLiftCarDoorCommand is the corresponding interface of BACnetPropertyStatesLiftCarDoorCommand
type BACnetPropertyStatesLiftCarDoorCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLiftCarDoorCommand returns LiftCarDoorCommand (property field)
	GetLiftCarDoorCommand() BACnetLiftCarDoorCommandTagged
	// IsBACnetPropertyStatesLiftCarDoorCommand is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesLiftCarDoorCommand()
}

// _BACnetPropertyStatesLiftCarDoorCommand is the data-structure of this message
type _BACnetPropertyStatesLiftCarDoorCommand struct {
	BACnetPropertyStatesContract
	LiftCarDoorCommand BACnetLiftCarDoorCommandTagged
}

var _ BACnetPropertyStatesLiftCarDoorCommand = (*_BACnetPropertyStatesLiftCarDoorCommand)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesLiftCarDoorCommand)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLiftCarDoorCommand) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLiftCarDoorCommand) GetLiftCarDoorCommand() BACnetLiftCarDoorCommandTagged {
	return m.LiftCarDoorCommand
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesLiftCarDoorCommand factory function for _BACnetPropertyStatesLiftCarDoorCommand
func NewBACnetPropertyStatesLiftCarDoorCommand(peekedTagHeader BACnetTagHeader, liftCarDoorCommand BACnetLiftCarDoorCommandTagged) *_BACnetPropertyStatesLiftCarDoorCommand {
	if liftCarDoorCommand == nil {
		panic("liftCarDoorCommand of type BACnetLiftCarDoorCommandTagged for BACnetPropertyStatesLiftCarDoorCommand must not be nil")
	}
	_result := &_BACnetPropertyStatesLiftCarDoorCommand{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		LiftCarDoorCommand:           liftCarDoorCommand,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLiftCarDoorCommand(structType any) BACnetPropertyStatesLiftCarDoorCommand {
	if casted, ok := structType.(BACnetPropertyStatesLiftCarDoorCommand); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLiftCarDoorCommand); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLiftCarDoorCommand) GetTypeName() string {
	return "BACnetPropertyStatesLiftCarDoorCommand"
}

func (m *_BACnetPropertyStatesLiftCarDoorCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (liftCarDoorCommand)
	lengthInBits += m.LiftCarDoorCommand.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesLiftCarDoorCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesLiftCarDoorCommand) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesLiftCarDoorCommand BACnetPropertyStatesLiftCarDoorCommand, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLiftCarDoorCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLiftCarDoorCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	liftCarDoorCommand, err := ReadSimpleField[BACnetLiftCarDoorCommandTagged](ctx, "liftCarDoorCommand", ReadComplex[BACnetLiftCarDoorCommandTagged](BACnetLiftCarDoorCommandTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'liftCarDoorCommand' field"))
	}
	m.LiftCarDoorCommand = liftCarDoorCommand

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLiftCarDoorCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLiftCarDoorCommand")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesLiftCarDoorCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLiftCarDoorCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLiftCarDoorCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLiftCarDoorCommand")
		}

		if err := WriteSimpleField[BACnetLiftCarDoorCommandTagged](ctx, "liftCarDoorCommand", m.GetLiftCarDoorCommand(), WriteComplex[BACnetLiftCarDoorCommandTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'liftCarDoorCommand' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLiftCarDoorCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLiftCarDoorCommand")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesLiftCarDoorCommand) IsBACnetPropertyStatesLiftCarDoorCommand() {}

func (m *_BACnetPropertyStatesLiftCarDoorCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
