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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// GroupObjectDescriptorRealisationType6 is the data-structure of this message
type GroupObjectDescriptorRealisationType6 struct {
}

// IGroupObjectDescriptorRealisationType6 is the corresponding interface of GroupObjectDescriptorRealisationType6
type IGroupObjectDescriptorRealisationType6 interface {
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// NewGroupObjectDescriptorRealisationType6 factory function for GroupObjectDescriptorRealisationType6
func NewGroupObjectDescriptorRealisationType6() *GroupObjectDescriptorRealisationType6 {
	return &GroupObjectDescriptorRealisationType6{}
}

func CastGroupObjectDescriptorRealisationType6(structType interface{}) *GroupObjectDescriptorRealisationType6 {
	if casted, ok := structType.(GroupObjectDescriptorRealisationType6); ok {
		return &casted
	}
	if casted, ok := structType.(*GroupObjectDescriptorRealisationType6); ok {
		return casted
	}
	return nil
}

func (m *GroupObjectDescriptorRealisationType6) GetTypeName() string {
	return "GroupObjectDescriptorRealisationType6"
}

func (m *GroupObjectDescriptorRealisationType6) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *GroupObjectDescriptorRealisationType6) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *GroupObjectDescriptorRealisationType6) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func GroupObjectDescriptorRealisationType6Parse(readBuffer utils.ReadBuffer) (*GroupObjectDescriptorRealisationType6, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GroupObjectDescriptorRealisationType6"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GroupObjectDescriptorRealisationType6")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("GroupObjectDescriptorRealisationType6"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GroupObjectDescriptorRealisationType6")
	}

	// Create the instance
	return NewGroupObjectDescriptorRealisationType6(), nil
}

func (m *GroupObjectDescriptorRealisationType6) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("GroupObjectDescriptorRealisationType6"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for GroupObjectDescriptorRealisationType6")
	}

	if popErr := writeBuffer.PopContext("GroupObjectDescriptorRealisationType6"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for GroupObjectDescriptorRealisationType6")
	}
	return nil
}

func (m *GroupObjectDescriptorRealisationType6) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
