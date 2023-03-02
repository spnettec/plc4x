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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// SetAttributeAllRequest is the corresponding interface of SetAttributeAllRequest
type SetAttributeAllRequest interface {
	utils.LengthAware
	utils.Serializable
	CipService
}

// SetAttributeAllRequestExactly can be used when we want exactly this type and not a type which fulfills SetAttributeAllRequest.
// This is useful for switch cases.
type SetAttributeAllRequestExactly interface {
	SetAttributeAllRequest
	isSetAttributeAllRequest() bool
}

// _SetAttributeAllRequest is the data-structure of this message
type _SetAttributeAllRequest struct {
	*_CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SetAttributeAllRequest) GetService() uint8 {
	return 0x02
}

func (m *_SetAttributeAllRequest) GetResponse() bool {
	return bool(false)
}

func (m *_SetAttributeAllRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SetAttributeAllRequest) InitializeParent(parent CipService) {}

func (m *_SetAttributeAllRequest) GetParent() CipService {
	return m._CipService
}

// NewSetAttributeAllRequest factory function for _SetAttributeAllRequest
func NewSetAttributeAllRequest(serviceLen uint16) *_SetAttributeAllRequest {
	_result := &_SetAttributeAllRequest{
		_CipService: NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSetAttributeAllRequest(structType interface{}) SetAttributeAllRequest {
	if casted, ok := structType.(SetAttributeAllRequest); ok {
		return casted
	}
	if casted, ok := structType.(*SetAttributeAllRequest); ok {
		return *casted
	}
	return nil
}

func (m *_SetAttributeAllRequest) GetTypeName() string {
	return "SetAttributeAllRequest"
}

func (m *_SetAttributeAllRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SetAttributeAllRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SetAttributeAllRequestParse(theBytes []byte, connected bool, serviceLen uint16) (SetAttributeAllRequest, error) {
	return SetAttributeAllRequestParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func SetAttributeAllRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (SetAttributeAllRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SetAttributeAllRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SetAttributeAllRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SetAttributeAllRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SetAttributeAllRequest")
	}

	// Create a partially initialized instance
	_child := &_SetAttributeAllRequest{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_SetAttributeAllRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SetAttributeAllRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SetAttributeAllRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SetAttributeAllRequest")
		}

		if popErr := writeBuffer.PopContext("SetAttributeAllRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SetAttributeAllRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SetAttributeAllRequest) isSetAttributeAllRequest() bool {
	return true
}

func (m *_SetAttributeAllRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
