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

// CIPData is the corresponding interface of CIPData
type CIPData interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetDataType returns DataType (property field)
	GetDataType() CIPDataTypeCode
	// GetData returns Data (property field)
	GetData() []byte
}

// CIPDataExactly can be used when we want exactly this type and not a type which fulfills CIPData.
// This is useful for switch cases.
type CIPDataExactly interface {
	CIPData
	isCIPData() bool
}

// _CIPData is the data-structure of this message
type _CIPData struct {
	DataType CIPDataTypeCode
	Data     []byte

	// Arguments.
	PacketLength uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CIPData) GetDataType() CIPDataTypeCode {
	return m.DataType
}

func (m *_CIPData) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCIPData factory function for _CIPData
func NewCIPData(dataType CIPDataTypeCode, data []byte, packetLength uint16) *_CIPData {
	return &_CIPData{DataType: dataType, Data: data, PacketLength: packetLength}
}

// Deprecated: use the interface for direct cast
func CastCIPData(structType any) CIPData {
	if casted, ok := structType.(CIPData); ok {
		return casted
	}
	if casted, ok := structType.(*CIPData); ok {
		return *casted
	}
	return nil
}

func (m *_CIPData) GetTypeName() string {
	return "CIPData"
}

func (m *_CIPData) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (dataType)
	lengthInBits += 16

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_CIPData) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CIPDataParse(ctx context.Context, theBytes []byte, packetLength uint16) (CIPData, error) {
	return CIPDataParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), packetLength)
}

func CIPDataParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, packetLength uint16) (CIPData, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CIPData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CIPData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dataType)
	if pullErr := readBuffer.PullContext("dataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dataType")
	}
	_dataType, _dataTypeErr := CIPDataTypeCodeParseWithBuffer(ctx, readBuffer)
	if _dataTypeErr != nil {
		return nil, errors.Wrap(_dataTypeErr, "Error parsing 'dataType' field of CIPData")
	}
	dataType := _dataType
	if closeErr := readBuffer.CloseContext("dataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dataType")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(uint16(packetLength) - uint16(uint16(2)))
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of CIPData")
	}

	if closeErr := readBuffer.CloseContext("CIPData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CIPData")
	}

	// Create the instance
	return &_CIPData{
		PacketLength: packetLength,
		DataType:     dataType,
		Data:         data,
	}, nil
}

func (m *_CIPData) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CIPData) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("CIPData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CIPData")
	}

	// Simple Field (dataType)
	if pushErr := writeBuffer.PushContext("dataType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for dataType")
	}
	_dataTypeErr := writeBuffer.WriteSerializable(ctx, m.GetDataType())
	if popErr := writeBuffer.PopContext("dataType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for dataType")
	}
	if _dataTypeErr != nil {
		return errors.Wrap(_dataTypeErr, "Error serializing 'dataType' field")
	}

	// Array Field (data)
	// Byte Array field (data)
	if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

	if popErr := writeBuffer.PopContext("CIPData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CIPData")
	}
	return nil
}

////
// Arguments Getter

func (m *_CIPData) GetPacketLength() uint16 {
	return m.PacketLength
}

//
////

func (m *_CIPData) isCIPData() bool {
	return true
}

func (m *_CIPData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
