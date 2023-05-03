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
)

// Code generated by code-generation. DO NOT EDIT.

// DF1CommandResponseMessageProtectedTypedLogicalRead is the corresponding interface of DF1CommandResponseMessageProtectedTypedLogicalRead
type DF1CommandResponseMessageProtectedTypedLogicalRead interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	DF1ResponseMessage
	// GetData returns Data (property field)
	GetData() []uint8
}

// DF1CommandResponseMessageProtectedTypedLogicalReadExactly can be used when we want exactly this type and not a type which fulfills DF1CommandResponseMessageProtectedTypedLogicalRead.
// This is useful for switch cases.
type DF1CommandResponseMessageProtectedTypedLogicalReadExactly interface {
	DF1CommandResponseMessageProtectedTypedLogicalRead
	isDF1CommandResponseMessageProtectedTypedLogicalRead() bool
}

// _DF1CommandResponseMessageProtectedTypedLogicalRead is the data-structure of this message
type _DF1CommandResponseMessageProtectedTypedLogicalRead struct {
	*_DF1ResponseMessage
	Data []uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) GetCommandCode() uint8 {
	return 0x4F
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) InitializeParent(parent DF1ResponseMessage, destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16) {
	m.DestinationAddress = destinationAddress
	m.SourceAddress = sourceAddress
	m.Status = status
	m.TransactionCounter = transactionCounter
}

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) GetParent() DF1ResponseMessage {
	return m._DF1ResponseMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) GetData() []uint8 {
	return m.Data
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDF1CommandResponseMessageProtectedTypedLogicalRead factory function for _DF1CommandResponseMessageProtectedTypedLogicalRead
func NewDF1CommandResponseMessageProtectedTypedLogicalRead(data []uint8, destinationAddress uint8, sourceAddress uint8, status uint8, transactionCounter uint16, payloadLength uint16) *_DF1CommandResponseMessageProtectedTypedLogicalRead {
	_result := &_DF1CommandResponseMessageProtectedTypedLogicalRead{
		Data:                data,
		_DF1ResponseMessage: NewDF1ResponseMessage(destinationAddress, sourceAddress, status, transactionCounter, payloadLength),
	}
	_result._DF1ResponseMessage._DF1ResponseMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDF1CommandResponseMessageProtectedTypedLogicalRead(structType any) DF1CommandResponseMessageProtectedTypedLogicalRead {
	if casted, ok := structType.(DF1CommandResponseMessageProtectedTypedLogicalRead); ok {
		return casted
	}
	if casted, ok := structType.(*DF1CommandResponseMessageProtectedTypedLogicalRead); ok {
		return *casted
	}
	return nil
}

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) GetTypeName() string {
	return "DF1CommandResponseMessageProtectedTypedLogicalRead"
}

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DF1CommandResponseMessageProtectedTypedLogicalReadParse(theBytes []byte, payloadLength uint16) (DF1CommandResponseMessageProtectedTypedLogicalRead, error) {
	return DF1CommandResponseMessageProtectedTypedLogicalReadParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), payloadLength)
}

func DF1CommandResponseMessageProtectedTypedLogicalReadParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, payloadLength uint16) (DF1CommandResponseMessageProtectedTypedLogicalRead, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1CommandResponseMessageProtectedTypedLogicalRead"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1CommandResponseMessageProtectedTypedLogicalRead")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (data)
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for data")
	}
	// Length array
	var data []uint8
	{
		_dataLength := uint16(payloadLength) - uint16(uint16(8))
		_dataEndPos := positionAware.GetPos() + uint16(_dataLength)
		for positionAware.GetPos() < _dataEndPos {
			_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'data' field of DF1CommandResponseMessageProtectedTypedLogicalRead")
			}
			data = append(data, _item)
		}
	}
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for data")
	}

	if closeErr := readBuffer.CloseContext("DF1CommandResponseMessageProtectedTypedLogicalRead"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1CommandResponseMessageProtectedTypedLogicalRead")
	}

	// Create a partially initialized instance
	_child := &_DF1CommandResponseMessageProtectedTypedLogicalRead{
		_DF1ResponseMessage: &_DF1ResponseMessage{
			PayloadLength: payloadLength,
		},
		Data: data,
	}
	_child._DF1ResponseMessage._DF1ResponseMessageChildRequirements = _child
	return _child, nil
}

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1CommandResponseMessageProtectedTypedLogicalRead"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1CommandResponseMessageProtectedTypedLogicalRead")
		}

		// Array Field (data)
		if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for data")
		}
		for _curItem, _element := range m.GetData() {
			_ = _curItem
			_elementErr := writeBuffer.WriteUint8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'data' field")
			}
		}
		if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for data")
		}

		if popErr := writeBuffer.PopContext("DF1CommandResponseMessageProtectedTypedLogicalRead"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1CommandResponseMessageProtectedTypedLogicalRead")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) isDF1CommandResponseMessageProtectedTypedLogicalRead() bool {
	return true
}

func (m *_DF1CommandResponseMessageProtectedTypedLogicalRead) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
