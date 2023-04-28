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

// SetAttributeListResponse is the corresponding interface of SetAttributeListResponse
type SetAttributeListResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
}

// SetAttributeListResponseExactly can be used when we want exactly this type and not a type which fulfills SetAttributeListResponse.
// This is useful for switch cases.
type SetAttributeListResponseExactly interface {
	SetAttributeListResponse
	isSetAttributeListResponse() bool
}

// _SetAttributeListResponse is the data-structure of this message
type _SetAttributeListResponse struct {
	*_CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SetAttributeListResponse) GetService() uint8 {
	return 0x04
}

func (m *_SetAttributeListResponse) GetResponse() bool {
	return bool(true)
}

func (m *_SetAttributeListResponse) GetConnected() bool {
	return false
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SetAttributeListResponse) InitializeParent(parent CipService) {}

func (m *_SetAttributeListResponse) GetParent() CipService {
	return m._CipService
}

// NewSetAttributeListResponse factory function for _SetAttributeListResponse
func NewSetAttributeListResponse(serviceLen uint16) *_SetAttributeListResponse {
	_result := &_SetAttributeListResponse{
		_CipService: NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSetAttributeListResponse(structType any) SetAttributeListResponse {
	if casted, ok := structType.(SetAttributeListResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SetAttributeListResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SetAttributeListResponse) GetTypeName() string {
	return "SetAttributeListResponse"
}

func (m *_SetAttributeListResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}

func (m *_SetAttributeListResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SetAttributeListResponseParse(theBytes []byte, connected bool, serviceLen uint16) (SetAttributeListResponse, error) {
	return SetAttributeListResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func SetAttributeListResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (SetAttributeListResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SetAttributeListResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SetAttributeListResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SetAttributeListResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SetAttributeListResponse")
	}

	// Create a partially initialized instance
	_child := &_SetAttributeListResponse{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_SetAttributeListResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SetAttributeListResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SetAttributeListResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SetAttributeListResponse")
		}

		if popErr := writeBuffer.PopContext("SetAttributeListResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SetAttributeListResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SetAttributeListResponse) isSetAttributeListResponse() bool {
	return true
}

func (m *_SetAttributeListResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
