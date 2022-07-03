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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// CBusCommandPointToPointToMultiPointNormal is the corresponding interface of CBusCommandPointToPointToMultiPointNormal
type CBusCommandPointToPointToMultiPointNormal interface {
	utils.LengthAware
	utils.Serializable
	CBusPointToPointToMultipointCommand
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetSalData returns SalData (property field)
	GetSalData() SALData
	// GetCrc returns Crc (property field)
	GetCrc() Checksum
	// GetAlpha returns Alpha (property field)
	GetAlpha() Alpha
}

// CBusCommandPointToPointToMultiPointNormalExactly can be used when we want exactly this type and not a type which fulfills CBusCommandPointToPointToMultiPointNormal.
// This is useful for switch cases.
type CBusCommandPointToPointToMultiPointNormalExactly interface {
	CBusCommandPointToPointToMultiPointNormal
	isCBusCommandPointToPointToMultiPointNormal() bool
}

// _CBusCommandPointToPointToMultiPointNormal is the data-structure of this message
type _CBusCommandPointToPointToMultiPointNormal struct {
	*_CBusPointToPointToMultipointCommand
	Application ApplicationIdContainer
	SalData     SALData
	Crc         Checksum
	Alpha       Alpha
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CBusCommandPointToPointToMultiPointNormal) InitializeParent(parent CBusPointToPointToMultipointCommand, bridgeAddress BridgeAddress, networkRoute NetworkRoute, peekedApplication byte, termination RequestTermination) {
	m.BridgeAddress = bridgeAddress
	m.NetworkRoute = networkRoute
	m.PeekedApplication = peekedApplication
	m.Termination = termination
}

func (m *_CBusCommandPointToPointToMultiPointNormal) GetParent() CBusPointToPointToMultipointCommand {
	return m._CBusPointToPointToMultipointCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusCommandPointToPointToMultiPointNormal) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_CBusCommandPointToPointToMultiPointNormal) GetSalData() SALData {
	return m.SalData
}

func (m *_CBusCommandPointToPointToMultiPointNormal) GetCrc() Checksum {
	return m.Crc
}

func (m *_CBusCommandPointToPointToMultiPointNormal) GetAlpha() Alpha {
	return m.Alpha
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusCommandPointToPointToMultiPointNormal factory function for _CBusCommandPointToPointToMultiPointNormal
func NewCBusCommandPointToPointToMultiPointNormal(application ApplicationIdContainer, salData SALData, crc Checksum, alpha Alpha, bridgeAddress BridgeAddress, networkRoute NetworkRoute, peekedApplication byte, termination RequestTermination, srchk bool) *_CBusCommandPointToPointToMultiPointNormal {
	_result := &_CBusCommandPointToPointToMultiPointNormal{
		Application:                          application,
		SalData:                              salData,
		Crc:                                  crc,
		Alpha:                                alpha,
		_CBusPointToPointToMultipointCommand: NewCBusPointToPointToMultipointCommand(bridgeAddress, networkRoute, peekedApplication, termination, srchk),
	}
	_result._CBusPointToPointToMultipointCommand._CBusPointToPointToMultipointCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCBusCommandPointToPointToMultiPointNormal(structType interface{}) CBusCommandPointToPointToMultiPointNormal {
	if casted, ok := structType.(CBusCommandPointToPointToMultiPointNormal); ok {
		return casted
	}
	if casted, ok := structType.(*CBusCommandPointToPointToMultiPointNormal); ok {
		return *casted
	}
	return nil
}

func (m *_CBusCommandPointToPointToMultiPointNormal) GetTypeName() string {
	return "CBusCommandPointToPointToMultiPointNormal"
}

func (m *_CBusCommandPointToPointToMultiPointNormal) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CBusCommandPointToPointToMultiPointNormal) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (application)
	lengthInBits += 8

	// Simple field (salData)
	lengthInBits += m.SalData.GetLengthInBits()

	// Optional Field (crc)
	if m.Crc != nil {
		lengthInBits += m.Crc.GetLengthInBits()
	}

	// Optional Field (alpha)
	if m.Alpha != nil {
		lengthInBits += m.Alpha.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_CBusCommandPointToPointToMultiPointNormal) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusCommandPointToPointToMultiPointNormalParse(readBuffer utils.ReadBuffer, srchk bool) (CBusCommandPointToPointToMultiPointNormal, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusCommandPointToPointToMultiPointNormal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusCommandPointToPointToMultiPointNormal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (application)
	if pullErr := readBuffer.PullContext("application"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for application")
	}
	_application, _applicationErr := ApplicationIdContainerParse(readBuffer)
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field of CBusCommandPointToPointToMultiPointNormal")
	}
	application := _application
	if closeErr := readBuffer.CloseContext("application"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for application")
	}

	// Simple Field (salData)
	if pullErr := readBuffer.PullContext("salData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for salData")
	}
	_salData, _salDataErr := SALDataParse(readBuffer)
	if _salDataErr != nil {
		return nil, errors.Wrap(_salDataErr, "Error parsing 'salData' field of CBusCommandPointToPointToMultiPointNormal")
	}
	salData := _salData.(SALData)
	if closeErr := readBuffer.CloseContext("salData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for salData")
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
			return nil, errors.Wrap(_err, "Error parsing 'crc' field of CBusCommandPointToPointToMultiPointNormal")
		default:
			crc = _val.(Checksum)
			if closeErr := readBuffer.CloseContext("crc"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for crc")
			}
		}
	}

	// Optional Field (alpha) (Can be skipped, if a given expression evaluates to false)
	var alpha Alpha = nil
	{
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
			return nil, errors.Wrap(_err, "Error parsing 'alpha' field of CBusCommandPointToPointToMultiPointNormal")
		default:
			alpha = _val.(Alpha)
			if closeErr := readBuffer.CloseContext("alpha"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for alpha")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("CBusCommandPointToPointToMultiPointNormal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusCommandPointToPointToMultiPointNormal")
	}

	// Create a partially initialized instance
	_child := &_CBusCommandPointToPointToMultiPointNormal{
		Application: application,
		SalData:     salData,
		Crc:         crc,
		Alpha:       alpha,
		_CBusPointToPointToMultipointCommand: &_CBusPointToPointToMultipointCommand{
			Srchk: srchk,
		},
	}
	_child._CBusPointToPointToMultipointCommand._CBusPointToPointToMultipointCommandChildRequirements = _child
	return _child, nil
}

func (m *_CBusCommandPointToPointToMultiPointNormal) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusCommandPointToPointToMultiPointNormal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusCommandPointToPointToMultiPointNormal")
		}

		// Simple Field (application)
		if pushErr := writeBuffer.PushContext("application"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for application")
		}
		_applicationErr := writeBuffer.WriteSerializable(m.GetApplication())
		if popErr := writeBuffer.PopContext("application"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for application")
		}
		if _applicationErr != nil {
			return errors.Wrap(_applicationErr, "Error serializing 'application' field")
		}

		// Simple Field (salData)
		if pushErr := writeBuffer.PushContext("salData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for salData")
		}
		_salDataErr := writeBuffer.WriteSerializable(m.GetSalData())
		if popErr := writeBuffer.PopContext("salData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for salData")
		}
		if _salDataErr != nil {
			return errors.Wrap(_salDataErr, "Error serializing 'salData' field")
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

		if popErr := writeBuffer.PopContext("CBusCommandPointToPointToMultiPointNormal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusCommandPointToPointToMultiPointNormal")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CBusCommandPointToPointToMultiPointNormal) isCBusCommandPointToPointToMultiPointNormal() bool {
	return true
}

func (m *_CBusCommandPointToPointToMultiPointNormal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
