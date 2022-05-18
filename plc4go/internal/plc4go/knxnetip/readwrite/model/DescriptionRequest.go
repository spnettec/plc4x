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

// DescriptionRequest is the data-structure of this message
type DescriptionRequest struct {
	*KnxNetIpMessage
	HpaiControlEndpoint *HPAIControlEndpoint
}

// IDescriptionRequest is the corresponding interface of DescriptionRequest
type IDescriptionRequest interface {
	IKnxNetIpMessage
	// GetHpaiControlEndpoint returns HpaiControlEndpoint (property field)
	GetHpaiControlEndpoint() *HPAIControlEndpoint
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

func (m *DescriptionRequest) GetMsgType() uint16 {
	return 0x0203
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *DescriptionRequest) InitializeParent(parent *KnxNetIpMessage) {}

func (m *DescriptionRequest) GetParent() *KnxNetIpMessage {
	return m.KnxNetIpMessage
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *DescriptionRequest) GetHpaiControlEndpoint() *HPAIControlEndpoint {
	return m.HpaiControlEndpoint
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewDescriptionRequest factory function for DescriptionRequest
func NewDescriptionRequest(hpaiControlEndpoint *HPAIControlEndpoint) *DescriptionRequest {
	_result := &DescriptionRequest{
		HpaiControlEndpoint: hpaiControlEndpoint,
		KnxNetIpMessage:     NewKnxNetIpMessage(),
	}
	_result.Child = _result
	return _result
}

func CastDescriptionRequest(structType interface{}) *DescriptionRequest {
	if casted, ok := structType.(DescriptionRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*DescriptionRequest); ok {
		return casted
	}
	if casted, ok := structType.(KnxNetIpMessage); ok {
		return CastDescriptionRequest(casted.Child)
	}
	if casted, ok := structType.(*KnxNetIpMessage); ok {
		return CastDescriptionRequest(casted.Child)
	}
	return nil
}

func (m *DescriptionRequest) GetTypeName() string {
	return "DescriptionRequest"
}

func (m *DescriptionRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *DescriptionRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (hpaiControlEndpoint)
	lengthInBits += m.HpaiControlEndpoint.GetLengthInBits()

	return lengthInBits
}

func (m *DescriptionRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DescriptionRequestParse(readBuffer utils.ReadBuffer) (*DescriptionRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DescriptionRequest"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (hpaiControlEndpoint)
	if pullErr := readBuffer.PullContext("hpaiControlEndpoint"); pullErr != nil {
		return nil, pullErr
	}
	_hpaiControlEndpoint, _hpaiControlEndpointErr := HPAIControlEndpointParse(readBuffer)
	if _hpaiControlEndpointErr != nil {
		return nil, errors.Wrap(_hpaiControlEndpointErr, "Error parsing 'hpaiControlEndpoint' field")
	}
	hpaiControlEndpoint := CastHPAIControlEndpoint(_hpaiControlEndpoint)
	if closeErr := readBuffer.CloseContext("hpaiControlEndpoint"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("DescriptionRequest"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &DescriptionRequest{
		HpaiControlEndpoint: CastHPAIControlEndpoint(hpaiControlEndpoint),
		KnxNetIpMessage:     &KnxNetIpMessage{},
	}
	_child.KnxNetIpMessage.Child = _child
	return _child, nil
}

func (m *DescriptionRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DescriptionRequest"); pushErr != nil {
			return pushErr
		}

		// Simple Field (hpaiControlEndpoint)
		if pushErr := writeBuffer.PushContext("hpaiControlEndpoint"); pushErr != nil {
			return pushErr
		}
		_hpaiControlEndpointErr := m.HpaiControlEndpoint.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("hpaiControlEndpoint"); popErr != nil {
			return popErr
		}
		if _hpaiControlEndpointErr != nil {
			return errors.Wrap(_hpaiControlEndpointErr, "Error serializing 'hpaiControlEndpoint' field")
		}

		if popErr := writeBuffer.PopContext("DescriptionRequest"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *DescriptionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
