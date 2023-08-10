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

// ListServicesResponse is the corresponding interface of ListServicesResponse
type ListServicesResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	EipPacket
	// GetTypeIds returns TypeIds (property field)
	GetTypeIds() []TypeId
}

// ListServicesResponseExactly can be used when we want exactly this type and not a type which fulfills ListServicesResponse.
// This is useful for switch cases.
type ListServicesResponseExactly interface {
	ListServicesResponse
	isListServicesResponse() bool
}

// _ListServicesResponse is the data-structure of this message
type _ListServicesResponse struct {
	*_EipPacket
	TypeIds []TypeId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ListServicesResponse) GetCommand() uint16 {
	return 0x0004
}

func (m *_ListServicesResponse) GetResponse() bool {
	return bool(true)
}

func (m *_ListServicesResponse) GetPacketLength() uint16 {
	return 0
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ListServicesResponse) InitializeParent(parent EipPacket, sessionHandle uint32, status uint32, senderContext []byte, options uint32) {
	m.SessionHandle = sessionHandle
	m.Status = status
	m.SenderContext = senderContext
	m.Options = options
}

func (m *_ListServicesResponse) GetParent() EipPacket {
	return m._EipPacket
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ListServicesResponse) GetTypeIds() []TypeId {
	return m.TypeIds
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewListServicesResponse factory function for _ListServicesResponse
func NewListServicesResponse(typeIds []TypeId, sessionHandle uint32, status uint32, senderContext []byte, options uint32) *_ListServicesResponse {
	_result := &_ListServicesResponse{
		TypeIds:    typeIds,
		_EipPacket: NewEipPacket(sessionHandle, status, senderContext, options),
	}
	_result._EipPacket._EipPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastListServicesResponse(structType any) ListServicesResponse {
	if casted, ok := structType.(ListServicesResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ListServicesResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ListServicesResponse) GetTypeName() string {
	return "ListServicesResponse"
}

func (m *_ListServicesResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (typeIdCount)
	lengthInBits += 16

	// Array field
	if len(m.TypeIds) > 0 {
		for _curItem, element := range m.TypeIds {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.TypeIds), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ListServicesResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ListServicesResponseParse(ctx context.Context, theBytes []byte, response bool) (ListServicesResponse, error) {
	return ListServicesResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func ListServicesResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ListServicesResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ListServicesResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ListServicesResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (typeIdCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	typeIdCount, _typeIdCountErr := readBuffer.ReadUint16("typeIdCount", 16)
	_ = typeIdCount
	if _typeIdCountErr != nil {
		return nil, errors.Wrap(_typeIdCountErr, "Error parsing 'typeIdCount' field of ListServicesResponse")
	}

	// Array field (typeIds)
	if pullErr := readBuffer.PullContext("typeIds", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for typeIds")
	}
	// Count array
	typeIds := make([]TypeId, utils.Max(typeIdCount, 0))
	// This happens when the size is set conditional to 0
	if len(typeIds) == 0 {
		typeIds = nil
	}
	{
		_numItems := uint16(utils.Max(typeIdCount, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := TypeIdParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'typeIds' field of ListServicesResponse")
			}
			typeIds[_curItem] = _item.(TypeId)
		}
	}
	if closeErr := readBuffer.CloseContext("typeIds", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for typeIds")
	}

	if closeErr := readBuffer.CloseContext("ListServicesResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ListServicesResponse")
	}

	// Create a partially initialized instance
	_child := &_ListServicesResponse{
		_EipPacket: &_EipPacket{},
		TypeIds:    typeIds,
	}
	_child._EipPacket._EipPacketChildRequirements = _child
	return _child, nil
}

func (m *_ListServicesResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ListServicesResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ListServicesResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ListServicesResponse")
		}

		// Implicit Field (typeIdCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		typeIdCount := uint16(uint16(len(m.GetTypeIds())))
		_typeIdCountErr := writeBuffer.WriteUint16("typeIdCount", 16, (typeIdCount))
		if _typeIdCountErr != nil {
			return errors.Wrap(_typeIdCountErr, "Error serializing 'typeIdCount' field")
		}

		// Array Field (typeIds)
		if pushErr := writeBuffer.PushContext("typeIds", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for typeIds")
		}
		for _curItem, _element := range m.GetTypeIds() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetTypeIds()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'typeIds' field")
			}
		}
		if popErr := writeBuffer.PopContext("typeIds", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for typeIds")
		}

		if popErr := writeBuffer.PopContext("ListServicesResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ListServicesResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ListServicesResponse) isListServicesResponse() bool {
	return true
}

func (m *_ListServicesResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
