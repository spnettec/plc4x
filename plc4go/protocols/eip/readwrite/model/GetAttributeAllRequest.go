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

// GetAttributeAllRequest is the corresponding interface of GetAttributeAllRequest
type GetAttributeAllRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetClassSegment returns ClassSegment (property field)
	GetClassSegment() PathSegment
	// GetInstanceSegment returns InstanceSegment (property field)
	GetInstanceSegment() PathSegment
}

// GetAttributeAllRequestExactly can be used when we want exactly this type and not a type which fulfills GetAttributeAllRequest.
// This is useful for switch cases.
type GetAttributeAllRequestExactly interface {
	GetAttributeAllRequest
	isGetAttributeAllRequest() bool
}

// _GetAttributeAllRequest is the data-structure of this message
type _GetAttributeAllRequest struct {
	*_CipService
	ClassSegment    PathSegment
	InstanceSegment PathSegment
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GetAttributeAllRequest) GetService() uint8 {
	return 0x01
}

func (m *_GetAttributeAllRequest) GetResponse() bool {
	return bool(false)
}

func (m *_GetAttributeAllRequest) GetConnected() bool {
	return false
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GetAttributeAllRequest) InitializeParent(parent CipService) {}

func (m *_GetAttributeAllRequest) GetParent() CipService {
	return m._CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GetAttributeAllRequest) GetClassSegment() PathSegment {
	return m.ClassSegment
}

func (m *_GetAttributeAllRequest) GetInstanceSegment() PathSegment {
	return m.InstanceSegment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewGetAttributeAllRequest factory function for _GetAttributeAllRequest
func NewGetAttributeAllRequest(classSegment PathSegment, instanceSegment PathSegment, serviceLen uint16) *_GetAttributeAllRequest {
	_result := &_GetAttributeAllRequest{
		ClassSegment:    classSegment,
		InstanceSegment: instanceSegment,
		_CipService:     NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastGetAttributeAllRequest(structType any) GetAttributeAllRequest {
	if casted, ok := structType.(GetAttributeAllRequest); ok {
		return casted
	}
	if casted, ok := structType.(*GetAttributeAllRequest); ok {
		return *casted
	}
	return nil
}

func (m *_GetAttributeAllRequest) GetTypeName() string {
	return "GetAttributeAllRequest"
}

func (m *_GetAttributeAllRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (requestPathSize)
	lengthInBits += 8

	// Simple field (classSegment)
	lengthInBits += m.ClassSegment.GetLengthInBits(ctx)

	// Simple field (instanceSegment)
	lengthInBits += m.InstanceSegment.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_GetAttributeAllRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func GetAttributeAllRequestParse(ctx context.Context, theBytes []byte, connected bool, serviceLen uint16) (GetAttributeAllRequest, error) {
	return GetAttributeAllRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func GetAttributeAllRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (GetAttributeAllRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("GetAttributeAllRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GetAttributeAllRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (requestPathSize) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	requestPathSize, _requestPathSizeErr := readBuffer.ReadUint8("requestPathSize", 8)
	_ = requestPathSize
	if _requestPathSizeErr != nil {
		return nil, errors.Wrap(_requestPathSizeErr, "Error parsing 'requestPathSize' field of GetAttributeAllRequest")
	}

	// Simple Field (classSegment)
	if pullErr := readBuffer.PullContext("classSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for classSegment")
	}
	_classSegment, _classSegmentErr := PathSegmentParseWithBuffer(ctx, readBuffer)
	if _classSegmentErr != nil {
		return nil, errors.Wrap(_classSegmentErr, "Error parsing 'classSegment' field of GetAttributeAllRequest")
	}
	classSegment := _classSegment.(PathSegment)
	if closeErr := readBuffer.CloseContext("classSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for classSegment")
	}

	// Simple Field (instanceSegment)
	if pullErr := readBuffer.PullContext("instanceSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for instanceSegment")
	}
	_instanceSegment, _instanceSegmentErr := PathSegmentParseWithBuffer(ctx, readBuffer)
	if _instanceSegmentErr != nil {
		return nil, errors.Wrap(_instanceSegmentErr, "Error parsing 'instanceSegment' field of GetAttributeAllRequest")
	}
	instanceSegment := _instanceSegment.(PathSegment)
	if closeErr := readBuffer.CloseContext("instanceSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for instanceSegment")
	}

	if closeErr := readBuffer.CloseContext("GetAttributeAllRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GetAttributeAllRequest")
	}

	// Create a partially initialized instance
	_child := &_GetAttributeAllRequest{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
		ClassSegment:    classSegment,
		InstanceSegment: instanceSegment,
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_GetAttributeAllRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GetAttributeAllRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GetAttributeAllRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GetAttributeAllRequest")
		}

		// Implicit Field (requestPathSize) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		requestPathSize := uint8(uint8((uint8(m.GetClassSegment().GetLengthInBytes(ctx)) + uint8(m.GetInstanceSegment().GetLengthInBytes(ctx)))) / uint8(uint8(2)))
		_requestPathSizeErr := writeBuffer.WriteUint8("requestPathSize", 8, uint8((requestPathSize)))
		if _requestPathSizeErr != nil {
			return errors.Wrap(_requestPathSizeErr, "Error serializing 'requestPathSize' field")
		}

		// Simple Field (classSegment)
		if pushErr := writeBuffer.PushContext("classSegment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for classSegment")
		}
		_classSegmentErr := writeBuffer.WriteSerializable(ctx, m.GetClassSegment())
		if popErr := writeBuffer.PopContext("classSegment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for classSegment")
		}
		if _classSegmentErr != nil {
			return errors.Wrap(_classSegmentErr, "Error serializing 'classSegment' field")
		}

		// Simple Field (instanceSegment)
		if pushErr := writeBuffer.PushContext("instanceSegment"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for instanceSegment")
		}
		_instanceSegmentErr := writeBuffer.WriteSerializable(ctx, m.GetInstanceSegment())
		if popErr := writeBuffer.PopContext("instanceSegment"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for instanceSegment")
		}
		if _instanceSegmentErr != nil {
			return errors.Wrap(_instanceSegmentErr, "Error serializing 'instanceSegment' field")
		}

		if popErr := writeBuffer.PopContext("GetAttributeAllRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GetAttributeAllRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GetAttributeAllRequest) isGetAttributeAllRequest() bool {
	return true
}

func (m *_GetAttributeAllRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
