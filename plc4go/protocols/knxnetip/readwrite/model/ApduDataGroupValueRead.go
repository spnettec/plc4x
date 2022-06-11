/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
)

// Code generated by code-generation. DO NOT EDIT.

// ApduDataGroupValueRead is the data-structure of this message
type ApduDataGroupValueRead struct {
	*ApduData

	// Arguments.
	DataLength uint8
}

// IApduDataGroupValueRead is the corresponding interface of ApduDataGroupValueRead
type IApduDataGroupValueRead interface {
	IApduData
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *ApduDataGroupValueRead) GetApciType() uint8 {
	return 0x0
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *ApduDataGroupValueRead) InitializeParent(parent *ApduData) {}

func (m *ApduDataGroupValueRead) GetParent() *ApduData {
	return m.ApduData
}

// NewApduDataGroupValueRead factory function for ApduDataGroupValueRead
func NewApduDataGroupValueRead(dataLength uint8) *ApduDataGroupValueRead {
	_result := &ApduDataGroupValueRead{
		ApduData: NewApduData(dataLength),
	}
	_result.Child = _result
	return _result
}

func CastApduDataGroupValueRead(structType interface{}) *ApduDataGroupValueRead {
	if casted, ok := structType.(ApduDataGroupValueRead); ok {
		return &casted
	}
	if casted, ok := structType.(*ApduDataGroupValueRead); ok {
		return casted
	}
	if casted, ok := structType.(ApduData); ok {
		return CastApduDataGroupValueRead(casted.Child)
	}
	if casted, ok := structType.(*ApduData); ok {
		return CastApduDataGroupValueRead(casted.Child)
	}
	return nil
}

func (m *ApduDataGroupValueRead) GetTypeName() string {
	return "ApduDataGroupValueRead"
}

func (m *ApduDataGroupValueRead) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *ApduDataGroupValueRead) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Reserved Field (reserved)
	lengthInBits += 6

	return lengthInBits
}

func (m *ApduDataGroupValueRead) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataGroupValueReadParse(readBuffer utils.ReadBuffer, dataLength uint8) (*ApduDataGroupValueRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataGroupValueRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataGroupValueRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 6)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	if closeErr := readBuffer.CloseContext("ApduDataGroupValueRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataGroupValueRead")
	}

	// Create a partially initialized instance
	_child := &ApduDataGroupValueRead{
		ApduData: &ApduData{},
	}
	_child.ApduData.Child = _child
	return _child, nil
}

func (m *ApduDataGroupValueRead) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataGroupValueRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataGroupValueRead")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 6, uint8(0x00))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		if popErr := writeBuffer.PopContext("ApduDataGroupValueRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataGroupValueRead")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *ApduDataGroupValueRead) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
