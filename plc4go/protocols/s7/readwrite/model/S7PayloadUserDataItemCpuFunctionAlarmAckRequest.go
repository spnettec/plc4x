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


// Constant values.
const S7PayloadUserDataItemCpuFunctionAlarmAckRequest_FUNCTIONID uint8 = 0x09

// S7PayloadUserDataItemCpuFunctionAlarmAckRequest is the corresponding interface of S7PayloadUserDataItemCpuFunctionAlarmAckRequest
type S7PayloadUserDataItemCpuFunctionAlarmAckRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	S7PayloadUserDataItem
	// GetMessageObjects returns MessageObjects (property field)
	GetMessageObjects() []AlarmMessageObjectAckType
}

// S7PayloadUserDataItemCpuFunctionAlarmAckRequestExactly can be used when we want exactly this type and not a type which fulfills S7PayloadUserDataItemCpuFunctionAlarmAckRequest.
// This is useful for switch cases.
type S7PayloadUserDataItemCpuFunctionAlarmAckRequestExactly interface {
	S7PayloadUserDataItemCpuFunctionAlarmAckRequest
	isS7PayloadUserDataItemCpuFunctionAlarmAckRequest() bool
}

// _S7PayloadUserDataItemCpuFunctionAlarmAckRequest is the data-structure of this message
type _S7PayloadUserDataItemCpuFunctionAlarmAckRequest struct {
	*_S7PayloadUserDataItem
        MessageObjects []AlarmMessageObjectAckType
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckRequest)  GetCpuFunctionGroup() uint8 {
return 0x04}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckRequest)  GetCpuFunctionType() uint8 {
return 0x04}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckRequest)  GetCpuSubfunction() uint8 {
return 0x0b}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckRequest) InitializeParent(parent S7PayloadUserDataItem , returnCode DataTransportErrorCode , transportSize DataTransportSize , dataLength uint16 ) {	m.ReturnCode = returnCode
	m.TransportSize = transportSize
	m.DataLength = dataLength
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckRequest)  GetParent() S7PayloadUserDataItem {
	return m._S7PayloadUserDataItem
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckRequest) GetMessageObjects() []AlarmMessageObjectAckType {
	return m.MessageObjects
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckRequest) GetFunctionId() uint8 {
	return S7PayloadUserDataItemCpuFunctionAlarmAckRequest_FUNCTIONID
}

///////////////////////-4
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewS7PayloadUserDataItemCpuFunctionAlarmAckRequest factory function for _S7PayloadUserDataItemCpuFunctionAlarmAckRequest
func NewS7PayloadUserDataItemCpuFunctionAlarmAckRequest( messageObjects []AlarmMessageObjectAckType , returnCode DataTransportErrorCode , transportSize DataTransportSize , dataLength uint16 ) *_S7PayloadUserDataItemCpuFunctionAlarmAckRequest {
	_result := &_S7PayloadUserDataItemCpuFunctionAlarmAckRequest{
		MessageObjects: messageObjects,
    	_S7PayloadUserDataItem: NewS7PayloadUserDataItem(returnCode, transportSize, dataLength),
	}
	_result._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastS7PayloadUserDataItemCpuFunctionAlarmAckRequest(structType any) S7PayloadUserDataItemCpuFunctionAlarmAckRequest {
    if casted, ok := structType.(S7PayloadUserDataItemCpuFunctionAlarmAckRequest); ok {
		return casted
	}
	if casted, ok := structType.(*S7PayloadUserDataItemCpuFunctionAlarmAckRequest); ok {
		return *casted
	}
	return nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckRequest) GetTypeName() string {
	return "S7PayloadUserDataItemCpuFunctionAlarmAckRequest"
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Const Field (functionId)
	lengthInBits += 8

	// Implicit Field (numberOfObjects)
	lengthInBits += 8

	// Array field
	if len(m.MessageObjects) > 0 {
		for _curItem, element := range m.MessageObjects {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.MessageObjects), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{GetLengthInBits(context.Context) uint16}).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}


func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func S7PayloadUserDataItemCpuFunctionAlarmAckRequestParse(ctx context.Context, theBytes []byte, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionAlarmAckRequest, error) {
	return S7PayloadUserDataItemCpuFunctionAlarmAckRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cpuFunctionGroup, cpuFunctionType, cpuSubfunction)
}

func S7PayloadUserDataItemCpuFunctionAlarmAckRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cpuFunctionGroup uint8, cpuFunctionType uint8, cpuSubfunction uint8) (S7PayloadUserDataItemCpuFunctionAlarmAckRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("S7PayloadUserDataItemCpuFunctionAlarmAckRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for S7PayloadUserDataItemCpuFunctionAlarmAckRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (functionId)
	functionId, _functionIdErr := readBuffer.ReadUint8("functionId", 8)
	if _functionIdErr != nil {
		return nil, errors.Wrap(_functionIdErr, "Error parsing 'functionId' field of S7PayloadUserDataItemCpuFunctionAlarmAckRequest")
	}
	if functionId != S7PayloadUserDataItemCpuFunctionAlarmAckRequest_FUNCTIONID {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", S7PayloadUserDataItemCpuFunctionAlarmAckRequest_FUNCTIONID) + " but got " + fmt.Sprintf("%d", functionId))
	}

	// Implicit Field (numberOfObjects) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	numberOfObjects, _numberOfObjectsErr := readBuffer.ReadUint8("numberOfObjects", 8)
	_ = numberOfObjects
	if _numberOfObjectsErr != nil {
		return nil, errors.Wrap(_numberOfObjectsErr, "Error parsing 'numberOfObjects' field of S7PayloadUserDataItemCpuFunctionAlarmAckRequest")
	}

	// Array field (messageObjects)
	if pullErr := readBuffer.PullContext("messageObjects", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for messageObjects")
	}
	// Count array
	messageObjects := make([]AlarmMessageObjectAckType, numberOfObjects)
	// This happens when the size is set conditional to 0
	if len(messageObjects) == 0 {
		messageObjects = nil
	}
	{
		_numItems := uint16(numberOfObjects)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := AlarmMessageObjectAckTypeParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'messageObjects' field of S7PayloadUserDataItemCpuFunctionAlarmAckRequest")
			}
			messageObjects[_curItem] = _item.(AlarmMessageObjectAckType)
		}
	}
	if closeErr := readBuffer.CloseContext("messageObjects", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for messageObjects")
	}

	if closeErr := readBuffer.CloseContext("S7PayloadUserDataItemCpuFunctionAlarmAckRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for S7PayloadUserDataItemCpuFunctionAlarmAckRequest")
	}

	// Create a partially initialized instance
	_child := &_S7PayloadUserDataItemCpuFunctionAlarmAckRequest{
		_S7PayloadUserDataItem: &_S7PayloadUserDataItem{
		},
		MessageObjects: messageObjects,
	}
	_child._S7PayloadUserDataItem._S7PayloadUserDataItemChildRequirements = _child
	return _child, nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("S7PayloadUserDataItemCpuFunctionAlarmAckRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for S7PayloadUserDataItemCpuFunctionAlarmAckRequest")
		}

	// Const Field (functionId)
	_functionIdErr := writeBuffer.WriteUint8("functionId", 8, 0x09)
	if _functionIdErr != nil {
		return errors.Wrap(_functionIdErr, "Error serializing 'functionId' field")
	}

	// Implicit Field (numberOfObjects) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	numberOfObjects := uint8(uint8(len(m.GetMessageObjects())))
	_numberOfObjectsErr := writeBuffer.WriteUint8("numberOfObjects", 8, (numberOfObjects))
	if _numberOfObjectsErr != nil {
		return errors.Wrap(_numberOfObjectsErr, "Error serializing 'numberOfObjects' field")
	}

	// Array Field (messageObjects)
	if pushErr := writeBuffer.PushContext("messageObjects", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for messageObjects")
	}
	for _curItem, _element := range m.GetMessageObjects() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetMessageObjects()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'messageObjects' field")
		}
	}
	if popErr := writeBuffer.PopContext("messageObjects", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for messageObjects")
	}

		if popErr := writeBuffer.PopContext("S7PayloadUserDataItemCpuFunctionAlarmAckRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for S7PayloadUserDataItemCpuFunctionAlarmAckRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckRequest) isS7PayloadUserDataItemCpuFunctionAlarmAckRequest() bool {
	return true
}

func (m *_S7PayloadUserDataItemCpuFunctionAlarmAckRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



