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

// CALDataReplyStatus is the corresponding interface of CALDataReplyStatus
type CALDataReplyStatus interface {
	CALData
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

// _CALDataReplyStatus is the data-structure of this message
type _CALDataReplyStatus struct {
	*_CALData
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

func (m *_CALDataReplyStatus) InitializeParent(parent CALData, commandTypeContainer CALCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_CALDataReplyStatus) GetParent() CALData {
	return m._CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataReplyStatus) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_CALDataReplyStatus) GetBlockStart() uint8 {
	return m.BlockStart
}

func (m *_CALDataReplyStatus) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataReplyStatus factory function for _CALDataReplyStatus
func NewCALDataReplyStatus(application ApplicationIdContainer, blockStart uint8, data []byte, commandTypeContainer CALCommandTypeContainer) *_CALDataReplyStatus {
	_result := &_CALDataReplyStatus{
		Application: application,
		BlockStart:  blockStart,
		Data:        data,
		_CALData:    NewCALData(commandTypeContainer),
	}
	_result._CALData._CALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataReplyStatus(structType interface{}) CALDataReplyStatus {
	if casted, ok := structType.(CALDataReplyStatus); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataReplyStatus); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataReplyStatus) GetTypeName() string {
	return "CALDataReplyStatus"
}

func (m *_CALDataReplyStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CALDataReplyStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

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

func (m *_CALDataReplyStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataReplyStatusParse(readBuffer utils.ReadBuffer, commandTypeContainer CALCommandTypeContainer) (CALDataReplyStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataReplyStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataReplyStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

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

	if closeErr := readBuffer.CloseContext("CALDataReplyStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataReplyStatus")
	}

	// Create a partially initialized instance
	_child := &_CALDataReplyStatus{
		Application: application,
		BlockStart:  blockStart,
		Data:        data,
		_CALData:    &_CALData{},
	}
	_child._CALData._CALDataChildRequirements = _child
	return _child, nil
}

func (m *_CALDataReplyStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataReplyStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataReplyStatus")
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

		if popErr := writeBuffer.PopContext("CALDataReplyStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataReplyStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CALDataReplyStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
