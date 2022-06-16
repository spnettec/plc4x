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
	"fmt"
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const CBusCommandPointToPointToMultiPointStatus_CR byte = 0xD

// CBusCommandPointToPointToMultiPointStatus is the corresponding interface of CBusCommandPointToPointToMultiPointStatus
type CBusCommandPointToPointToMultiPointStatus interface {
	CBusPointToPointToMultipointCommand
	// GetStatusRequest returns StatusRequest (property field)
	GetStatusRequest() StatusRequest
	// GetCrc returns Crc (property field)
	GetCrc() Checksum
	// GetPeekAlpha returns PeekAlpha (property field)
	GetPeekAlpha() byte
	// GetAlpha returns Alpha (property field)
	GetAlpha() Alpha
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _CBusCommandPointToPointToMultiPointStatus is the data-structure of this message
type _CBusCommandPointToPointToMultiPointStatus struct {
	*_CBusPointToPointToMultipointCommand
	StatusRequest StatusRequest
	Crc           Checksum
	PeekAlpha     byte
	Alpha         Alpha

	// Arguments.
	Srchk bool
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CBusCommandPointToPointToMultiPointStatus) InitializeParent(parent CBusPointToPointToMultipointCommand, bridgeAddress BridgeAddress, networkRoute NetworkRoute, peekedApplication byte) {
	m.BridgeAddress = bridgeAddress
	m.NetworkRoute = networkRoute
	m.PeekedApplication = peekedApplication
}

func (m *_CBusCommandPointToPointToMultiPointStatus) GetParent() CBusPointToPointToMultipointCommand {
	return m._CBusPointToPointToMultipointCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusCommandPointToPointToMultiPointStatus) GetStatusRequest() StatusRequest {
	return m.StatusRequest
}

func (m *_CBusCommandPointToPointToMultiPointStatus) GetCrc() Checksum {
	return m.Crc
}

func (m *_CBusCommandPointToPointToMultiPointStatus) GetPeekAlpha() byte {
	return m.PeekAlpha
}

func (m *_CBusCommandPointToPointToMultiPointStatus) GetAlpha() Alpha {
	return m.Alpha
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_CBusCommandPointToPointToMultiPointStatus) GetCr() byte {
	return CBusCommandPointToPointToMultiPointStatus_CR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusCommandPointToPointToMultiPointStatus factory function for _CBusCommandPointToPointToMultiPointStatus
func NewCBusCommandPointToPointToMultiPointStatus(statusRequest StatusRequest, crc Checksum, peekAlpha byte, alpha Alpha, bridgeAddress BridgeAddress, networkRoute NetworkRoute, peekedApplication byte, srchk bool) *_CBusCommandPointToPointToMultiPointStatus {
	_result := &_CBusCommandPointToPointToMultiPointStatus{
		StatusRequest:                        statusRequest,
		Crc:                                  crc,
		PeekAlpha:                            peekAlpha,
		Alpha:                                alpha,
		_CBusPointToPointToMultipointCommand: NewCBusPointToPointToMultipointCommand(bridgeAddress, networkRoute, peekedApplication, srchk),
	}
	_result._CBusPointToPointToMultipointCommand._CBusPointToPointToMultipointCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCBusCommandPointToPointToMultiPointStatus(structType interface{}) CBusCommandPointToPointToMultiPointStatus {
	if casted, ok := structType.(CBusCommandPointToPointToMultiPointStatus); ok {
		return casted
	}
	if casted, ok := structType.(*CBusCommandPointToPointToMultiPointStatus); ok {
		return *casted
	}
	return nil
}

func (m *_CBusCommandPointToPointToMultiPointStatus) GetTypeName() string {
	return "CBusCommandPointToPointToMultiPointStatus"
}

func (m *_CBusCommandPointToPointToMultiPointStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CBusCommandPointToPointToMultiPointStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (statusRequest)
	lengthInBits += m.StatusRequest.GetLengthInBits()

	// Optional Field (crc)
	if m.Crc != nil {
		lengthInBits += m.Crc.GetLengthInBits()
	}

	// Optional Field (alpha)
	if m.Alpha != nil {
		lengthInBits += m.Alpha.GetLengthInBits()
	}

	// Const Field (cr)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CBusCommandPointToPointToMultiPointStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusCommandPointToPointToMultiPointStatusParse(readBuffer utils.ReadBuffer, srchk bool) (CBusCommandPointToPointToMultiPointStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusCommandPointToPointToMultiPointStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusCommandPointToPointToMultiPointStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != byte(0xFF) {
			log.Info().Fields(map[string]interface{}{
				"expected value": byte(0xFF),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (statusRequest)
	if pullErr := readBuffer.PullContext("statusRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusRequest")
	}
	_statusRequest, _statusRequestErr := StatusRequestParse(readBuffer)
	if _statusRequestErr != nil {
		return nil, errors.Wrap(_statusRequestErr, "Error parsing 'statusRequest' field")
	}
	statusRequest := _statusRequest.(StatusRequest)
	if closeErr := readBuffer.CloseContext("statusRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusRequest")
	}

	// Optional Field (crc) (Can be skipped, if a given expression evaluates to false)
	var crc Checksum = nil
	if srchk {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("crc"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for crc")
		}
		_val, _err := ChecksumParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'crc' field")
		default:
			crc = _val.(Checksum)
			if closeErr := readBuffer.CloseContext("crc"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for crc")
			}
		}
	}

	// Peek Field (peekAlpha)
	currentPos = positionAware.GetPos()
	peekAlpha, _err := readBuffer.ReadByte("peekAlpha")
	if _err != nil {
		return nil, errors.Wrap(_err, "Error parsing 'peekAlpha' field")
	}

	readBuffer.Reset(currentPos)

	// Optional Field (alpha) (Can be skipped, if a given expression evaluates to false)
	var alpha Alpha = nil
	if bool(bool(bool((peekAlpha) >= (0x67)))) && bool(bool(bool((peekAlpha) <= (0x7A)))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("alpha"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for alpha")
		}
		_val, _err := AlphaParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'alpha' field")
		default:
			alpha = _val.(Alpha)
			if closeErr := readBuffer.CloseContext("alpha"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for alpha")
			}
		}
	}

	// Const Field (cr)
	cr, _crErr := readBuffer.ReadByte("cr")
	if _crErr != nil {
		return nil, errors.Wrap(_crErr, "Error parsing 'cr' field")
	}
	if cr != CBusCommandPointToPointToMultiPointStatus_CR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", CBusCommandPointToPointToMultiPointStatus_CR) + " but got " + fmt.Sprintf("%d", cr))
	}

	if closeErr := readBuffer.CloseContext("CBusCommandPointToPointToMultiPointStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusCommandPointToPointToMultiPointStatus")
	}

	// Create a partially initialized instance
	_child := &_CBusCommandPointToPointToMultiPointStatus{
		StatusRequest:                        statusRequest,
		Crc:                                  crc,
		PeekAlpha:                            peekAlpha,
		Alpha:                                alpha,
		_CBusPointToPointToMultipointCommand: &_CBusPointToPointToMultipointCommand{},
	}
	_child._CBusPointToPointToMultipointCommand._CBusPointToPointToMultipointCommandChildRequirements = _child
	return _child, nil
}

func (m *_CBusCommandPointToPointToMultiPointStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusCommandPointToPointToMultiPointStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusCommandPointToPointToMultiPointStatus")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteByte("reserved", byte(0xFF))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (statusRequest)
		if pushErr := writeBuffer.PushContext("statusRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusRequest")
		}
		_statusRequestErr := writeBuffer.WriteSerializable(m.GetStatusRequest())
		if popErr := writeBuffer.PopContext("statusRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusRequest")
		}
		if _statusRequestErr != nil {
			return errors.Wrap(_statusRequestErr, "Error serializing 'statusRequest' field")
		}

		// Optional Field (crc) (Can be skipped, if the value is null)
		var crc Checksum = nil
		if m.GetCrc() != nil {
			if pushErr := writeBuffer.PushContext("crc"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for crc")
			}
			crc = m.GetCrc()
			_crcErr := writeBuffer.WriteSerializable(crc)
			if popErr := writeBuffer.PopContext("crc"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for crc")
			}
			if _crcErr != nil {
				return errors.Wrap(_crcErr, "Error serializing 'crc' field")
			}
		}

		// Optional Field (alpha) (Can be skipped, if the value is null)
		var alpha Alpha = nil
		if m.GetAlpha() != nil {
			if pushErr := writeBuffer.PushContext("alpha"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for alpha")
			}
			alpha = m.GetAlpha()
			_alphaErr := writeBuffer.WriteSerializable(alpha)
			if popErr := writeBuffer.PopContext("alpha"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for alpha")
			}
			if _alphaErr != nil {
				return errors.Wrap(_alphaErr, "Error serializing 'alpha' field")
			}
		}

		// Const Field (cr)
		_crErr := writeBuffer.WriteByte("cr", 0xD)
		if _crErr != nil {
			return errors.Wrap(_crErr, "Error serializing 'cr' field")
		}

		if popErr := writeBuffer.PopContext("CBusCommandPointToPointToMultiPointStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusCommandPointToPointToMultiPointStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CBusCommandPointToPointToMultiPointStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
