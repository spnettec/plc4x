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

// CBusPointToMultiPointCommandNormal is the corresponding interface of CBusPointToMultiPointCommandNormal
type CBusPointToMultiPointCommandNormal interface {
	utils.LengthAware
	utils.Serializable
	CBusPointToMultiPointCommand
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetSalData returns SalData (property field)
	GetSalData() SALData
	// GetCrc returns Crc (property field)
	GetCrc() Checksum
}

// CBusPointToMultiPointCommandNormalExactly can be used when we want exactly this type and not a type which fulfills CBusPointToMultiPointCommandNormal.
// This is useful for switch cases.
type CBusPointToMultiPointCommandNormalExactly interface {
	CBusPointToMultiPointCommandNormal
	isCBusPointToMultiPointCommandNormal() bool
}

// _CBusPointToMultiPointCommandNormal is the data-structure of this message
type _CBusPointToMultiPointCommandNormal struct {
	*_CBusPointToMultiPointCommand
	Application ApplicationIdContainer
	SalData     SALData
	Crc         Checksum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CBusPointToMultiPointCommandNormal) InitializeParent(parent CBusPointToMultiPointCommand, peekedApplication byte) {
	m.PeekedApplication = peekedApplication
}

func (m *_CBusPointToMultiPointCommandNormal) GetParent() CBusPointToMultiPointCommand {
	return m._CBusPointToMultiPointCommand
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToMultiPointCommandNormal) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_CBusPointToMultiPointCommandNormal) GetSalData() SALData {
	return m.SalData
}

func (m *_CBusPointToMultiPointCommandNormal) GetCrc() Checksum {
	return m.Crc
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCBusPointToMultiPointCommandNormal factory function for _CBusPointToMultiPointCommandNormal
func NewCBusPointToMultiPointCommandNormal(application ApplicationIdContainer, salData SALData, crc Checksum, peekedApplication byte, srchk bool) *_CBusPointToMultiPointCommandNormal {
	_result := &_CBusPointToMultiPointCommandNormal{
		Application:                   application,
		SalData:                       salData,
		Crc:                           crc,
		_CBusPointToMultiPointCommand: NewCBusPointToMultiPointCommand(peekedApplication, srchk),
	}
	_result._CBusPointToMultiPointCommand._CBusPointToMultiPointCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCBusPointToMultiPointCommandNormal(structType interface{}) CBusPointToMultiPointCommandNormal {
	if casted, ok := structType.(CBusPointToMultiPointCommandNormal); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToMultiPointCommandNormal); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToMultiPointCommandNormal) GetTypeName() string {
	return "CBusPointToMultiPointCommandNormal"
}

func (m *_CBusPointToMultiPointCommandNormal) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CBusPointToMultiPointCommandNormal) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (application)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (salData)
	lengthInBits += m.SalData.GetLengthInBits()

	// Optional Field (crc)
	if m.Crc != nil {
		lengthInBits += m.Crc.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_CBusPointToMultiPointCommandNormal) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusPointToMultiPointCommandNormalParse(readBuffer utils.ReadBuffer, srchk bool) (CBusPointToMultiPointCommandNormal, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusPointToMultiPointCommandNormal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToMultiPointCommandNormal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (application)
	if pullErr := readBuffer.PullContext("application"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for application")
	}
	_application, _applicationErr := ApplicationIdContainerParse(readBuffer)
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field of CBusPointToMultiPointCommandNormal")
	}
	application := _application
	if closeErr := readBuffer.CloseContext("application"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for application")
	}

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CBusPointToMultiPointCommandNormal")
		}
		if reserved != byte(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": byte(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
		}
	}

	// Simple Field (salData)
	if pullErr := readBuffer.PullContext("salData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for salData")
	}
	_salData, _salDataErr := SALDataParse(readBuffer)
	if _salDataErr != nil {
		return nil, errors.Wrap(_salDataErr, "Error parsing 'salData' field of CBusPointToMultiPointCommandNormal")
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
			return nil, errors.Wrap(_err, "Error parsing 'crc' field of CBusPointToMultiPointCommandNormal")
		default:
			crc = _val.(Checksum)
			if closeErr := readBuffer.CloseContext("crc"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for crc")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("CBusPointToMultiPointCommandNormal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToMultiPointCommandNormal")
	}

	// Create a partially initialized instance
	_child := &_CBusPointToMultiPointCommandNormal{
		Application: application,
		SalData:     salData,
		Crc:         crc,
		_CBusPointToMultiPointCommand: &_CBusPointToMultiPointCommand{
			Srchk: srchk,
		},
	}
	_child._CBusPointToMultiPointCommand._CBusPointToMultiPointCommandChildRequirements = _child
	return _child, nil
}

func (m *_CBusPointToMultiPointCommandNormal) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusPointToMultiPointCommandNormal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusPointToMultiPointCommandNormal")
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

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteByte("reserved", byte(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
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

		if popErr := writeBuffer.PopContext("CBusPointToMultiPointCommandNormal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusPointToMultiPointCommandNormal")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CBusPointToMultiPointCommandNormal) isCBusPointToMultiPointCommandNormal() bool {
	return true
}

func (m *_CBusPointToMultiPointCommandNormal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
