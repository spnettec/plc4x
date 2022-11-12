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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// TunnelingRequestDataBlock is the corresponding interface of TunnelingRequestDataBlock
type TunnelingRequestDataBlock interface {
	utils.LengthAware
	utils.Serializable
	// GetCommunicationChannelId returns CommunicationChannelId (property field)
	GetCommunicationChannelId() uint8
	// GetSequenceCounter returns SequenceCounter (property field)
	GetSequenceCounter() uint8
}

// TunnelingRequestDataBlockExactly can be used when we want exactly this type and not a type which fulfills TunnelingRequestDataBlock.
// This is useful for switch cases.
type TunnelingRequestDataBlockExactly interface {
	TunnelingRequestDataBlock
	isTunnelingRequestDataBlock() bool
}

// _TunnelingRequestDataBlock is the data-structure of this message
type _TunnelingRequestDataBlock struct {
        CommunicationChannelId uint8
        SequenceCounter uint8
	// Reserved Fields
	reservedField0 *uint8
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TunnelingRequestDataBlock) GetCommunicationChannelId() uint8 {
	return m.CommunicationChannelId
}

func (m *_TunnelingRequestDataBlock) GetSequenceCounter() uint8 {
	return m.SequenceCounter
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewTunnelingRequestDataBlock factory function for _TunnelingRequestDataBlock
func NewTunnelingRequestDataBlock( communicationChannelId uint8 , sequenceCounter uint8 ) *_TunnelingRequestDataBlock {
return &_TunnelingRequestDataBlock{ CommunicationChannelId: communicationChannelId , SequenceCounter: sequenceCounter }
}

// Deprecated: use the interface for direct cast
func CastTunnelingRequestDataBlock(structType interface{}) TunnelingRequestDataBlock {
    if casted, ok := structType.(TunnelingRequestDataBlock); ok {
		return casted
	}
	if casted, ok := structType.(*TunnelingRequestDataBlock); ok {
		return *casted
	}
	return nil
}

func (m *_TunnelingRequestDataBlock) GetTypeName() string {
	return "TunnelingRequestDataBlock"
}

func (m *_TunnelingRequestDataBlock) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_TunnelingRequestDataBlock) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Implicit Field (structureLength)
	lengthInBits += 8

	// Simple field (communicationChannelId)
	lengthInBits += 8;

	// Simple field (sequenceCounter)
	lengthInBits += 8;

	// Reserved Field (reserved)
	lengthInBits += 8

	return lengthInBits
}


func (m *_TunnelingRequestDataBlock) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TunnelingRequestDataBlockParse(theBytes []byte) (TunnelingRequestDataBlock, error) {
	return TunnelingRequestDataBlockParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func TunnelingRequestDataBlockParseWithBuffer(readBuffer utils.ReadBuffer) (TunnelingRequestDataBlock, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TunnelingRequestDataBlock"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TunnelingRequestDataBlock")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (structureLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	structureLength, _structureLengthErr := readBuffer.ReadUint8("structureLength", 8)
	_ = structureLength
	if _structureLengthErr != nil {
		return nil, errors.Wrap(_structureLengthErr, "Error parsing 'structureLength' field of TunnelingRequestDataBlock")
	}

	// Simple Field (communicationChannelId)
_communicationChannelId, _communicationChannelIdErr := readBuffer.ReadUint8("communicationChannelId", 8)
	if _communicationChannelIdErr != nil {
		return nil, errors.Wrap(_communicationChannelIdErr, "Error parsing 'communicationChannelId' field of TunnelingRequestDataBlock")
	}
	communicationChannelId := _communicationChannelId

	// Simple Field (sequenceCounter)
_sequenceCounter, _sequenceCounterErr := readBuffer.ReadUint8("sequenceCounter", 8)
	if _sequenceCounterErr != nil {
		return nil, errors.Wrap(_sequenceCounterErr, "Error parsing 'sequenceCounter' field of TunnelingRequestDataBlock")
	}
	sequenceCounter := _sequenceCounter

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of TunnelingRequestDataBlock")
		}
		if reserved != uint8(0x00) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value": reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	if closeErr := readBuffer.CloseContext("TunnelingRequestDataBlock"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TunnelingRequestDataBlock")
	}

	// Create the instance
	return &_TunnelingRequestDataBlock{
			CommunicationChannelId: communicationChannelId,
			SequenceCounter: sequenceCounter,
			reservedField0: reservedField0,
		}, nil
}

func (m *_TunnelingRequestDataBlock) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TunnelingRequestDataBlock) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("TunnelingRequestDataBlock"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for TunnelingRequestDataBlock")
	}

	// Implicit Field (structureLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	structureLength := uint8(uint8(m.GetLengthInBytes()))
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

	// Reserved Field (reserved)
	{
		var reserved uint8 = uint8(0x00)
		if m.reservedField0 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": uint8(0x00),
				"got value": reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteUint8("reserved", 8, reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	if popErr := writeBuffer.PopContext("TunnelingRequestDataBlock"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for TunnelingRequestDataBlock")
	}
	return nil
}


func (m *_TunnelingRequestDataBlock) isTunnelingRequestDataBlock() bool {
	return true
}

func (m *_TunnelingRequestDataBlock) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



