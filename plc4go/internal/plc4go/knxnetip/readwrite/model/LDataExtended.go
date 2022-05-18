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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// LDataExtended is the data-structure of this message
type LDataExtended struct {
	*LDataFrame
	GroupAddress        bool
	HopCount            uint8
	ExtendedFrameFormat uint8
	SourceAddress       *KnxAddress
	DestinationAddress  []byte
	Apdu                *Apdu
}

// ILDataExtended is the corresponding interface of LDataExtended
type ILDataExtended interface {
	ILDataFrame
	// GetGroupAddress returns GroupAddress (property field)
	GetGroupAddress() bool
	// GetHopCount returns HopCount (property field)
	GetHopCount() uint8
	// GetExtendedFrameFormat returns ExtendedFrameFormat (property field)
	GetExtendedFrameFormat() uint8
	// GetSourceAddress returns SourceAddress (property field)
	GetSourceAddress() *KnxAddress
	// GetDestinationAddress returns DestinationAddress (property field)
	GetDestinationAddress() []byte
	// GetApdu returns Apdu (property field)
	GetApdu() *Apdu
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

func (m *LDataExtended) GetNotAckFrame() bool {
	return bool(true)
}

func (m *LDataExtended) GetPolling() bool {
	return bool(false)
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *LDataExtended) InitializeParent(parent *LDataFrame, frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) {
	m.LDataFrame.FrameType = frameType
	m.LDataFrame.NotRepeated = notRepeated
	m.LDataFrame.Priority = priority
	m.LDataFrame.AcknowledgeRequested = acknowledgeRequested
	m.LDataFrame.ErrorFlag = errorFlag
}

func (m *LDataExtended) GetParent() *LDataFrame {
	return m.LDataFrame
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *LDataExtended) GetGroupAddress() bool {
	return m.GroupAddress
}

func (m *LDataExtended) GetHopCount() uint8 {
	return m.HopCount
}

func (m *LDataExtended) GetExtendedFrameFormat() uint8 {
	return m.ExtendedFrameFormat
}

func (m *LDataExtended) GetSourceAddress() *KnxAddress {
	return m.SourceAddress
}

func (m *LDataExtended) GetDestinationAddress() []byte {
	return m.DestinationAddress
}

func (m *LDataExtended) GetApdu() *Apdu {
	return m.Apdu
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLDataExtended factory function for LDataExtended
func NewLDataExtended(groupAddress bool, hopCount uint8, extendedFrameFormat uint8, sourceAddress *KnxAddress, destinationAddress []byte, apdu *Apdu, frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) *LDataExtended {
	_result := &LDataExtended{
		GroupAddress:        groupAddress,
		HopCount:            hopCount,
		ExtendedFrameFormat: extendedFrameFormat,
		SourceAddress:       sourceAddress,
		DestinationAddress:  destinationAddress,
		Apdu:                apdu,
		LDataFrame:          NewLDataFrame(frameType, notRepeated, priority, acknowledgeRequested, errorFlag),
	}
	_result.Child = _result
	return _result
}

func CastLDataExtended(structType interface{}) *LDataExtended {
	if casted, ok := structType.(LDataExtended); ok {
		return &casted
	}
	if casted, ok := structType.(*LDataExtended); ok {
		return casted
	}
	if casted, ok := structType.(LDataFrame); ok {
		return CastLDataExtended(casted.Child)
	}
	if casted, ok := structType.(*LDataFrame); ok {
		return CastLDataExtended(casted.Child)
	}
	return nil
}

func (m *LDataExtended) GetTypeName() string {
	return "LDataExtended"
}

func (m *LDataExtended) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *LDataExtended) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (groupAddress)
	lengthInBits += 1

	// Simple field (hopCount)
	lengthInBits += 3

	// Simple field (extendedFrameFormat)
	lengthInBits += 4

	// Simple field (sourceAddress)
	lengthInBits += m.SourceAddress.GetLengthInBits()

	// Array field
	if len(m.DestinationAddress) > 0 {
		lengthInBits += 8 * uint16(len(m.DestinationAddress))
	}

	// Implicit Field (dataLength)
	lengthInBits += 8

	// Simple field (apdu)
	lengthInBits += m.Apdu.GetLengthInBits()

	return lengthInBits
}

func (m *LDataExtended) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LDataExtendedParse(readBuffer utils.ReadBuffer) (*LDataExtended, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LDataExtended"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (groupAddress)
	_groupAddress, _groupAddressErr := readBuffer.ReadBit("groupAddress")
	if _groupAddressErr != nil {
		return nil, errors.Wrap(_groupAddressErr, "Error parsing 'groupAddress' field")
	}
	groupAddress := _groupAddress

	// Simple Field (hopCount)
	_hopCount, _hopCountErr := readBuffer.ReadUint8("hopCount", 3)
	if _hopCountErr != nil {
		return nil, errors.Wrap(_hopCountErr, "Error parsing 'hopCount' field")
	}
	hopCount := _hopCount

	// Simple Field (extendedFrameFormat)
	_extendedFrameFormat, _extendedFrameFormatErr := readBuffer.ReadUint8("extendedFrameFormat", 4)
	if _extendedFrameFormatErr != nil {
		return nil, errors.Wrap(_extendedFrameFormatErr, "Error parsing 'extendedFrameFormat' field")
	}
	extendedFrameFormat := _extendedFrameFormat

	// Simple Field (sourceAddress)
	if pullErr := readBuffer.PullContext("sourceAddress"); pullErr != nil {
		return nil, pullErr
	}
	_sourceAddress, _sourceAddressErr := KnxAddressParse(readBuffer)
	if _sourceAddressErr != nil {
		return nil, errors.Wrap(_sourceAddressErr, "Error parsing 'sourceAddress' field")
	}
	sourceAddress := CastKnxAddress(_sourceAddress)
	if closeErr := readBuffer.CloseContext("sourceAddress"); closeErr != nil {
		return nil, closeErr
	}
	// Byte Array field (destinationAddress)
	numberOfBytesdestinationAddress := int(uint16(2))
	destinationAddress, _readArrayErr := readBuffer.ReadByteArray("destinationAddress", numberOfBytesdestinationAddress)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'destinationAddress' field")
	}

	// Implicit Field (dataLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	dataLength, _dataLengthErr := readBuffer.ReadUint8("dataLength", 8)
	_ = dataLength
	if _dataLengthErr != nil {
		return nil, errors.Wrap(_dataLengthErr, "Error parsing 'dataLength' field")
	}

	// Simple Field (apdu)
	if pullErr := readBuffer.PullContext("apdu"); pullErr != nil {
		return nil, pullErr
	}
	_apdu, _apduErr := ApduParse(readBuffer, uint8(dataLength))
	if _apduErr != nil {
		return nil, errors.Wrap(_apduErr, "Error parsing 'apdu' field")
	}
	apdu := CastApdu(_apdu)
	if closeErr := readBuffer.CloseContext("apdu"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("LDataExtended"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &LDataExtended{
		GroupAddress:        groupAddress,
		HopCount:            hopCount,
		ExtendedFrameFormat: extendedFrameFormat,
		SourceAddress:       CastKnxAddress(sourceAddress),
		DestinationAddress:  destinationAddress,
		Apdu:                CastApdu(apdu),
		LDataFrame:          &LDataFrame{},
	}
	_child.LDataFrame.Child = _child
	return _child, nil
}

func (m *LDataExtended) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LDataExtended"); pushErr != nil {
			return pushErr
		}

		// Simple Field (groupAddress)
		groupAddress := bool(m.GroupAddress)
		_groupAddressErr := writeBuffer.WriteBit("groupAddress", (groupAddress))
		if _groupAddressErr != nil {
			return errors.Wrap(_groupAddressErr, "Error serializing 'groupAddress' field")
		}

		// Simple Field (hopCount)
		hopCount := uint8(m.HopCount)
		_hopCountErr := writeBuffer.WriteUint8("hopCount", 3, (hopCount))
		if _hopCountErr != nil {
			return errors.Wrap(_hopCountErr, "Error serializing 'hopCount' field")
		}

		// Simple Field (extendedFrameFormat)
		extendedFrameFormat := uint8(m.ExtendedFrameFormat)
		_extendedFrameFormatErr := writeBuffer.WriteUint8("extendedFrameFormat", 4, (extendedFrameFormat))
		if _extendedFrameFormatErr != nil {
			return errors.Wrap(_extendedFrameFormatErr, "Error serializing 'extendedFrameFormat' field")
		}

		// Simple Field (sourceAddress)
		if pushErr := writeBuffer.PushContext("sourceAddress"); pushErr != nil {
			return pushErr
		}
		_sourceAddressErr := m.SourceAddress.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("sourceAddress"); popErr != nil {
			return popErr
		}
		if _sourceAddressErr != nil {
			return errors.Wrap(_sourceAddressErr, "Error serializing 'sourceAddress' field")
		}

		// Array Field (destinationAddress)
		if m.DestinationAddress != nil {
			// Byte Array field (destinationAddress)
			_writeArrayErr := writeBuffer.WriteByteArray("destinationAddress", m.DestinationAddress)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'destinationAddress' field")
			}
		}

		// Implicit Field (dataLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		dataLength := uint8(uint8(m.GetApdu().GetLengthInBytes()) - uint8(uint8(1)))
		_dataLengthErr := writeBuffer.WriteUint8("dataLength", 8, (dataLength))
		if _dataLengthErr != nil {
			return errors.Wrap(_dataLengthErr, "Error serializing 'dataLength' field")
		}

		// Simple Field (apdu)
		if pushErr := writeBuffer.PushContext("apdu"); pushErr != nil {
			return pushErr
		}
		_apduErr := m.Apdu.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("apdu"); popErr != nil {
			return popErr
		}
		if _apduErr != nil {
			return errors.Wrap(_apduErr, "Error serializing 'apdu' field")
		}

		if popErr := writeBuffer.PopContext("LDataExtended"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *LDataExtended) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
