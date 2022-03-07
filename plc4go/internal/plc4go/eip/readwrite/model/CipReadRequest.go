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

// The data-structure of this message
type CipReadRequest struct {
	*CipService
	RequestPathSize int8
	Tag             []byte
	ElementNb       uint16

	// Arguments.
	ServiceLen uint16
}

// The corresponding interface
type ICipReadRequest interface {
	ICipService
	// GetRequestPathSize returns RequestPathSize (property field)
	GetRequestPathSize() int8
	// GetTag returns Tag (property field)
	GetTag() []byte
	// GetElementNb returns ElementNb (property field)
	GetElementNb() uint16
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *CipReadRequest) Service() uint8 {
	return 0x4C
}

func (m *CipReadRequest) GetService() uint8 {
	return 0x4C
}

func (m *CipReadRequest) InitializeParent(parent *CipService) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *CipReadRequest) GetRequestPathSize() int8 {
	return m.RequestPathSize
}

func (m *CipReadRequest) GetTag() []byte {
	return m.Tag
}

func (m *CipReadRequest) GetElementNb() uint16 {
	return m.ElementNb
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCipReadRequest factory function for CipReadRequest
func NewCipReadRequest(requestPathSize int8, tag []byte, elementNb uint16, serviceLen uint16) *CipService {
	child := &CipReadRequest{
		RequestPathSize: requestPathSize,
		Tag:             tag,
		ElementNb:       elementNb,
		CipService:      NewCipService(serviceLen),
	}
	child.Child = child
	return child.CipService
}

func CastCipReadRequest(structType interface{}) *CipReadRequest {
	if casted, ok := structType.(CipReadRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*CipReadRequest); ok {
		return casted
	}
	if casted, ok := structType.(CipService); ok {
		return CastCipReadRequest(casted.Child)
	}
	if casted, ok := structType.(*CipService); ok {
		return CastCipReadRequest(casted.Child)
	}
	return nil
}

func (m *CipReadRequest) GetTypeName() string {
	return "CipReadRequest"
}

func (m *CipReadRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CipReadRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (requestPathSize)
	lengthInBits += 8

	// Array field
	if len(m.Tag) > 0 {
		lengthInBits += 8 * uint16(len(m.Tag))
	}

	// Simple field (elementNb)
	lengthInBits += 16

	return lengthInBits
}

func (m *CipReadRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CipReadRequestParse(readBuffer utils.ReadBuffer, serviceLen uint16) (*CipService, error) {
	if pullErr := readBuffer.PullContext("CipReadRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (requestPathSize)
	_requestPathSize, _requestPathSizeErr := readBuffer.ReadInt8("requestPathSize", 8)
	if _requestPathSizeErr != nil {
		return nil, errors.Wrap(_requestPathSizeErr, "Error parsing 'requestPathSize' field")
	}
	requestPathSize := _requestPathSize
	// Byte Array field (tag)
	numberOfBytestag := int(uint16(uint16(requestPathSize) * uint16(uint16(2))))
	tag, _readArrayErr := readBuffer.ReadByteArray("tag", numberOfBytestag)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'tag' field")
	}

	// Simple Field (elementNb)
	_elementNb, _elementNbErr := readBuffer.ReadUint16("elementNb", 16)
	if _elementNbErr != nil {
		return nil, errors.Wrap(_elementNbErr, "Error parsing 'elementNb' field")
	}
	elementNb := _elementNb

	if closeErr := readBuffer.CloseContext("CipReadRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &CipReadRequest{
		RequestPathSize: requestPathSize,
		Tag:             tag,
		ElementNb:       elementNb,
		CipService:      &CipService{},
	}
	_child.CipService.Child = _child
	return _child.CipService, nil
}

func (m *CipReadRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipReadRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (requestPathSize)
		requestPathSize := int8(m.RequestPathSize)
		_requestPathSizeErr := writeBuffer.WriteInt8("requestPathSize", 8, (requestPathSize))
		if _requestPathSizeErr != nil {
			return errors.Wrap(_requestPathSizeErr, "Error serializing 'requestPathSize' field")
		}

		// Array Field (tag)
		if m.Tag != nil {
			// Byte Array field (tag)
			_writeArrayErr := writeBuffer.WriteByteArray("tag", m.Tag)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'tag' field")
			}
		}

		// Simple Field (elementNb)
		elementNb := uint16(m.ElementNb)
		_elementNbErr := writeBuffer.WriteUint16("elementNb", 16, (elementNb))
		if _elementNbErr != nil {
			return errors.Wrap(_elementNbErr, "Error serializing 'elementNb' field")
		}

		if popErr := writeBuffer.PopContext("CipReadRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *CipReadRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
