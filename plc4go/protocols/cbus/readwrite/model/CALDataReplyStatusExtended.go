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
)

// Code generated by code-generation. DO NOT EDIT.

// CALDataReplyStatusExtended is the corresponding interface of CALDataReplyStatusExtended
type CALDataReplyStatusExtended interface {
	CALData
	// GetEncoding returns Encoding (property field)
	GetEncoding() uint8
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetBlockStart returns BlockStart (property field)
	GetBlockStart() uint8
	// GetData returns Data (property field)
	GetData() []byte
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _CALDataReplyStatusExtended is the data-structure of this message
type _CALDataReplyStatusExtended struct {
	*_CALData
	Encoding    uint8
	Application ApplicationIdContainer
	BlockStart  uint8
	Data        []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataReplyStatusExtended) InitializeParent(parent CALData, commandTypeContainer CALCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_CALDataReplyStatusExtended) GetParent() CALData {
	return m._CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataReplyStatusExtended) GetEncoding() uint8 {
	return m.Encoding
}

func (m *_CALDataReplyStatusExtended) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_CALDataReplyStatusExtended) GetBlockStart() uint8 {
	return m.BlockStart
}

func (m *_CALDataReplyStatusExtended) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataReplyStatusExtended factory function for _CALDataReplyStatusExtended
func NewCALDataReplyStatusExtended(encoding uint8, application ApplicationIdContainer, blockStart uint8, data []byte, commandTypeContainer CALCommandTypeContainer) *_CALDataReplyStatusExtended {
	_result := &_CALDataReplyStatusExtended{
		Encoding:    encoding,
		Application: application,
		BlockStart:  blockStart,
		Data:        data,
		_CALData:    NewCALData(commandTypeContainer),
	}
	_result._CALData._CALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataReplyStatusExtended(structType interface{}) CALDataReplyStatusExtended {
	if casted, ok := structType.(CALDataReplyStatusExtended); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataReplyStatusExtended); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataReplyStatusExtended) GetTypeName() string {
	return "CALDataReplyStatusExtended"
}

func (m *_CALDataReplyStatusExtended) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CALDataReplyStatusExtended) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (encoding)
	lengthInBits += 8

	// Simple field (application)
	lengthInBits += 8

	// Simple field (blockStart)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_CALDataReplyStatusExtended) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataReplyStatusExtendedParse(readBuffer utils.ReadBuffer, commandTypeContainer CALCommandTypeContainer) (CALDataReplyStatusExtended, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataReplyStatusExtended"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataReplyStatusExtended")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (encoding)
	_encoding, _encodingErr := readBuffer.ReadUint8("encoding", 8)
	if _encodingErr != nil {
		return nil, errors.Wrap(_encodingErr, "Error parsing 'encoding' field")
	}
	encoding := _encoding

	// Simple Field (application)
	if pullErr := readBuffer.PullContext("application"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for application")
	}
	_application, _applicationErr := ApplicationIdContainerParse(readBuffer)
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field")
	}
	application := _application
	if closeErr := readBuffer.CloseContext("application"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for application")
	}

	// Simple Field (blockStart)
	_blockStart, _blockStartErr := readBuffer.ReadUint8("blockStart", 8)
	if _blockStartErr != nil {
		return nil, errors.Wrap(_blockStartErr, "Error parsing 'blockStart' field")
	}
	blockStart := _blockStart
	// Byte Array field (data)
	numberOfBytesdata := int(commandTypeContainer.NumBytes())
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field")
	}

	if closeErr := readBuffer.CloseContext("CALDataReplyStatusExtended"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataReplyStatusExtended")
	}

	// Create a partially initialized instance
	_child := &_CALDataReplyStatusExtended{
		Encoding:    encoding,
		Application: application,
		BlockStart:  blockStart,
		Data:        data,
		_CALData:    &_CALData{},
	}
	_child._CALData._CALDataChildRequirements = _child
	return _child, nil
}

func (m *_CALDataReplyStatusExtended) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataReplyStatusExtended"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataReplyStatusExtended")
		}

		// Simple Field (encoding)
		encoding := uint8(m.GetEncoding())
		_encodingErr := writeBuffer.WriteUint8("encoding", 8, (encoding))
		if _encodingErr != nil {
			return errors.Wrap(_encodingErr, "Error serializing 'encoding' field")
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

		// Simple Field (blockStart)
		blockStart := uint8(m.GetBlockStart())
		_blockStartErr := writeBuffer.WriteUint8("blockStart", 8, (blockStart))
		if _blockStartErr != nil {
			return errors.Wrap(_blockStartErr, "Error serializing 'blockStart' field")
		}

		// Array Field (data)
		if m.GetData() != nil {
			// Byte Array field (data)
			_writeArrayErr := writeBuffer.WriteByteArray("data", m.GetData())
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'data' field")
			}
		}

		if popErr := writeBuffer.PopContext("CALDataReplyStatusExtended"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataReplyStatusExtended")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CALDataReplyStatusExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
