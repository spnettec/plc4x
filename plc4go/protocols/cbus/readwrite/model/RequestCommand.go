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
	"io"
	"github.com/rs/zerolog"
)

	// Code generated by code-generation. DO NOT EDIT.


// Constant values.
const RequestCommand_INITIATOR byte = 0x5C

// RequestCommand is the corresponding interface of RequestCommand
type RequestCommand interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Request
	// GetCbusCommand returns CbusCommand (property field)
	GetCbusCommand() CBusCommand
	// GetChksum returns Chksum (property field)
	GetChksum() Checksum
	// GetAlpha returns Alpha (property field)
	GetAlpha() Alpha
	// GetCbusCommandDecoded returns CbusCommandDecoded (virtual field)
	GetCbusCommandDecoded() CBusCommand
	// GetChksumDecoded returns ChksumDecoded (virtual field)
	GetChksumDecoded() Checksum
}

// RequestCommandExactly can be used when we want exactly this type and not a type which fulfills RequestCommand.
// This is useful for switch cases.
type RequestCommandExactly interface {
	RequestCommand
	isRequestCommand() bool
}

// _RequestCommand is the data-structure of this message
type _RequestCommand struct {
	*_Request
        CbusCommand CBusCommand
        Chksum Checksum
        Alpha Alpha
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RequestCommand) InitializeParent(parent Request , peekedByte RequestType , startingCR * RequestType , resetMode * RequestType , secondPeek RequestType , termination RequestTermination ) {	m.PeekedByte = peekedByte
	m.StartingCR = startingCR
	m.ResetMode = resetMode
	m.SecondPeek = secondPeek
	m.Termination = termination
}

func (m *_RequestCommand)  GetParent() Request {
	return m._Request
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RequestCommand) GetCbusCommand() CBusCommand {
	return m.CbusCommand
}

func (m *_RequestCommand) GetChksum() Checksum {
	return m.Chksum
}

func (m *_RequestCommand) GetAlpha() Alpha {
	return m.Alpha
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_RequestCommand) GetCbusCommandDecoded() CBusCommand {
	ctx := context.Background()
	_ = ctx
	alpha := m.Alpha
	_ = alpha
	return CastCBusCommand(m.GetCbusCommand())
}

func (m *_RequestCommand) GetChksumDecoded() Checksum {
	ctx := context.Background()
	_ = ctx
	alpha := m.Alpha
	_ = alpha
	return CastChecksum(m.GetChksum())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_RequestCommand) GetInitiator() byte {
	return RequestCommand_INITIATOR
}

///////////////////////-4
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewRequestCommand factory function for _RequestCommand
func NewRequestCommand( cbusCommand CBusCommand , chksum Checksum , alpha Alpha , peekedByte RequestType , startingCR *RequestType , resetMode *RequestType , secondPeek RequestType , termination RequestTermination , cBusOptions CBusOptions ) *_RequestCommand {
	_result := &_RequestCommand{
		CbusCommand: cbusCommand,
		Chksum: chksum,
		Alpha: alpha,
    	_Request: NewRequest(peekedByte, startingCR, resetMode, secondPeek, termination, cBusOptions),
	}
	_result._Request._RequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRequestCommand(structType any) RequestCommand {
    if casted, ok := structType.(RequestCommand); ok {
		return casted
	}
	if casted, ok := structType.(*RequestCommand); ok {
		return *casted
	}
	return nil
}

func (m *_RequestCommand) GetTypeName() string {
	return "RequestCommand"
}

func (m *_RequestCommand) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Const Field (initiator)
	lengthInBits += 8

	// Manual Field (cbusCommand)
	lengthInBits += uint16(int32((int32(m.GetCbusCommand().GetLengthInBytes(ctx)) * int32(int32(2)))) * int32(int32(8)))

	// A virtual field doesn't have any in- or output.

	// Manual Field (chksum)
	lengthInBits += uint16(utils.InlineIf((m.CBusOptions.GetSrchk()), func() any {return int32((int32(16)))}, func() any {return int32((int32(0)))}).(int32))

	// A virtual field doesn't have any in- or output.

	// Optional Field (alpha)
	if m.Alpha != nil {
		lengthInBits += m.Alpha.GetLengthInBits(ctx)
	}

	return lengthInBits
}


func (m *_RequestCommand) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RequestCommandParse(ctx context.Context, theBytes []byte, cBusOptions CBusOptions) (RequestCommand, error) {
	return RequestCommandParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func RequestCommandParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (RequestCommand, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("RequestCommand"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestCommand")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (initiator)
	initiator, _initiatorErr := readBuffer.ReadByte("initiator")
	if _initiatorErr != nil {
		return nil, errors.Wrap(_initiatorErr, "Error parsing 'initiator' field of RequestCommand")
	}
	if initiator != RequestCommand_INITIATOR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", RequestCommand_INITIATOR) + " but got " + fmt.Sprintf("%d", initiator))
	}

	// Manual Field (cbusCommand)
	_cbusCommand, _cbusCommandErr := ReadCBusCommand(ctx, readBuffer, cBusOptions, cBusOptions.GetSrchk())
	if _cbusCommandErr != nil {
		return nil, errors.Wrap(_cbusCommandErr, "Error parsing 'cbusCommand' field of RequestCommand")
	}
	var cbusCommand CBusCommand
	if _cbusCommand != nil {
            cbusCommand = _cbusCommand.(CBusCommand)
	}

	// Virtual field
	_cbusCommandDecoded := cbusCommand
	cbusCommandDecoded := _cbusCommandDecoded
	_ = cbusCommandDecoded

	// Manual Field (chksum)
	_chksum, _chksumErr := ReadAndValidateChecksum(ctx, readBuffer, cbusCommand, cBusOptions.GetSrchk())
	if _chksumErr != nil {
		return nil, errors.Wrap(_chksumErr, "Error parsing 'chksum' field of RequestCommand")
	}
	var chksum Checksum
	if _chksum != nil {
            chksum = _chksum.(Checksum)
	}

	// Virtual field
	_chksumDecoded := chksum
	chksumDecoded := _chksumDecoded
	_ = chksumDecoded

	// Optional Field (alpha) (Can be skipped, if a given expression evaluates to false)
	var alpha Alpha = nil
{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("alpha"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for alpha")
		}
_val, _err := AlphaParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'alpha' field of RequestCommand")
		default:
			alpha = _val.(Alpha)
			if closeErr := readBuffer.CloseContext("alpha"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for alpha")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("RequestCommand"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestCommand")
	}

	// Create a partially initialized instance
	_child := &_RequestCommand{
		_Request: &_Request{
			CBusOptions: cBusOptions,
		},
		CbusCommand: cbusCommand,
		Chksum: chksum,
		Alpha: alpha,
	}
	_child._Request._RequestChildRequirements = _child
	return _child, nil
}

func (m *_RequestCommand) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RequestCommand) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RequestCommand"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RequestCommand")
		}

	// Const Field (initiator)
	_initiatorErr := writeBuffer.WriteByte("initiator", 0x5C)
	if _initiatorErr != nil {
		return errors.Wrap(_initiatorErr, "Error serializing 'initiator' field")
	}

	// Manual Field (cbusCommand)
	_cbusCommandErr := WriteCBusCommand(ctx, writeBuffer, m.GetCbusCommand())
	if _cbusCommandErr != nil {
		return errors.Wrap(_cbusCommandErr, "Error serializing 'cbusCommand' field")
	}
	// Virtual field
	cbusCommandDecoded := m.GetCbusCommandDecoded()
	_ =	cbusCommandDecoded
	if _cbusCommandDecodedErr := writeBuffer.WriteVirtual(ctx, "cbusCommandDecoded", m.GetCbusCommandDecoded()); _cbusCommandDecodedErr != nil {
		return errors.Wrap(_cbusCommandDecodedErr, "Error serializing 'cbusCommandDecoded' field")
	}

	// Manual Field (chksum)
	_chksumErr := CalculateChecksum(ctx, writeBuffer, m.GetCbusCommand(), m.CBusOptions.GetSrchk())
	if _chksumErr != nil {
		return errors.Wrap(_chksumErr, "Error serializing 'chksum' field")
	}
	// Virtual field
	chksumDecoded := m.GetChksumDecoded()
	_ =	chksumDecoded
	if _chksumDecodedErr := writeBuffer.WriteVirtual(ctx, "chksumDecoded", m.GetChksumDecoded()); _chksumDecodedErr != nil {
		return errors.Wrap(_chksumDecodedErr, "Error serializing 'chksumDecoded' field")
	}

	// Optional Field (alpha) (Can be skipped, if the value is null)
	var alpha Alpha = nil
	if m.GetAlpha() != nil {
		if pushErr := writeBuffer.PushContext("alpha"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for alpha")
		}
		alpha = m.GetAlpha()
		_alphaErr := writeBuffer.WriteSerializable(ctx, alpha)
		if popErr := writeBuffer.PopContext("alpha"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for alpha")
		}
		if _alphaErr != nil {
			return errors.Wrap(_alphaErr, "Error serializing 'alpha' field")
		}
	}

		if popErr := writeBuffer.PopContext("RequestCommand"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RequestCommand")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_RequestCommand) isRequestCommand() bool {
	return true
}

func (m *_RequestCommand) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



