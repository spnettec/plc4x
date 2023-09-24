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


// DeviceConfigurationAckDataBlock is the corresponding interface of DeviceConfigurationAckDataBlock
type DeviceConfigurationAckDataBlock interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetSequenceCounter returns SequenceCounter (property field)
	GetSequenceCounter() uint8
	// GetStatus returns Status (property field)
	GetStatus() Status
}

// DeviceConfigurationAckDataBlockExactly can be used when we want exactly this type and not a type which fulfills DeviceConfigurationAckDataBlock.
// This is useful for switch cases.
type DeviceConfigurationAckDataBlockExactly interface {
	DeviceConfigurationAckDataBlock
	isDeviceConfigurationAckDataBlock() bool
}

// _DeviceConfigurationAckDataBlock is the data-structure of this message
type _DeviceConfigurationAckDataBlock struct {
        CommunicationChannelId uint8
        SequenceCounter uint8
        Status Status
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeviceConfigurationAckDataBlock) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_DeviceConfigurationAckDataBlock) GetSequenceCounter() uint8 {
	return m.SequenceCounter
}

func (m *_DeviceConfigurationAckDataBlock) GetStatus() Status {
	return m.Status
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewDeviceConfigurationAckDataBlock factory function for _DeviceConfigurationAckDataBlock
func NewDeviceConfigurationAckDataBlock( communicationChannelId uint8 , sequenceCounter uint8 , status Status ) *_DeviceConfigurationAckDataBlock {
return &_DeviceConfigurationAckDataBlock{ CommunicationChannelId: communicationChannelId , SequenceCounter: sequenceCounter , Status: status }
}

// Deprecated: use the interface for direct cast
func CastDeviceConfigurationAckDataBlock(structType any) DeviceConfigurationAckDataBlock {
    if casted, ok := structType.(DeviceConfigurationAckDataBlock); ok {
		return casted
	}
	if casted, ok := structType.(*DeviceConfigurationAckDataBlock); ok {
		return *casted
	}
	return nil
}

func (m *_DeviceConfigurationAckDataBlock) GetTypeName() string {
	return "DeviceConfigurationAckDataBlock"
}

func (m *_DeviceConfigurationAckDataBlock) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (communicationChannelId)
	lengthInBits += 8;

	// Simple field (sequenceCounter)
	lengthInBits += 8;

	// Simple field (status)
	lengthInBits += 8

	return lengthInBits
}


func (m *_DeviceConfigurationAckDataBlock) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DeviceConfigurationAckDataBlockParse(ctx context.Context, theBytes []byte) (DeviceConfigurationAckDataBlock, error) {
	return DeviceConfigurationAckDataBlockParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func DeviceConfigurationAckDataBlockParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (DeviceConfigurationAckDataBlock, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("DeviceConfigurationAckDataBlock"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeviceConfigurationAckDataBlock")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (structureLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := readBuffer.ReadUint8("structureLength", 8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field of DeviceConfigurationAckDataBlock")
	}

	// Simple Field (communicationChannelId)
_communicationChannelId, _communicationChannelIdErr := readBuffer.ReadUint8("communicationChannelId", 8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field of DeviceConfigurationAckDataBlock")
	}
	communicationChannelId := _communicationChannelId

	// Simple Field (sequenceCounter)
_sequenceCounter, _sequenceCounterErr := readBuffer.ReadUint8("sequenceCounter", 8)
	if _sequenceCounterErr != nil {
		return nil, errors.Wrap(_sequenceCounterErr, "Error parsing 'sequenceCounter' field of DeviceConfigurationAckDataBlock")
	}
	sequenceCounter := _sequenceCounter

	// Simple Field (status)
	if pullErr := readBuffer.PullContext("status"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for status")
	}
_status, _statusErr := StatusParseWithBuffer(ctx, readBuffer)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field of DeviceConfigurationAckDataBlock")
	}
	status := _status
	if closeErr := readBuffer.CloseContext("status"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for status")
	}

	if closeErr := readBuffer.CloseContext("DeviceConfigurationAckDataBlock"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeviceConfigurationAckDataBlock")
	}

	// Create the instance
	return &_DeviceConfigurationAckDataBlock{
			CommunicationChannelId: communicationChannelId,
			SequenceCounter: sequenceCounter,
			Status: status,
		}, nil
}

func (m *_DeviceConfigurationAckDataBlock) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeviceConfigurationAckDataBlock) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("DeviceConfigurationAckDataBlock"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for DeviceConfigurationAckDataBlock")
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.GetLengthInBytes(ctx)))
	_structureLengthErr := writeBuffer.WriteUint8("structureLength", 8, (structureLength))
	if _structureLengthErr != nil {
		return errors.Wrap(_structureLengthErr, "Error serializing 'structureLength' field")
	}

	// Simple Field (communicationChannelId)
	communicationChannelId := uint8(m.GetCommunicationChannelId())
	_communicationChannelIdErr := writeBuffer.WriteUint8("communicationChannelId", 8, (communicationChannelId))
	if _communicationChannelIdErr != nil {
		return errors.Wrap(_communicationChannelIdErr, "Error serializing 'communicationChannelId' field")
	}

	// Simple Field (sequenceCounter)
	sequenceCounter := uint8(m.GetSequenceCounter())
	_sequenceCounterErr := writeBuffer.WriteUint8("sequenceCounter", 8, (sequenceCounter))
	if _sequenceCounterErr != nil {
		return errors.Wrap(_sequenceCounterErr, "Error serializing 'sequenceCounter' field")
	}

	// Simple Field (status)
	if pushErr := writeBuffer.PushContext("status"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for status")
	}
	_statusErr := writeBuffer.WriteSerializable(ctx, m.GetStatus())
	if popErr := writeBuffer.PopContext("status"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for status")
	}
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
	}

	if popErr := writeBuffer.PopContext("DeviceConfigurationAckDataBlock"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for DeviceConfigurationAckDataBlock")
	}
	return nil
}


func (m *_DeviceConfigurationAckDataBlock) isDeviceConfigurationAckDataBlock() bool {
	return true
}

func (m *_DeviceConfigurationAckDataBlock) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



