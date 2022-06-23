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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// SALDataOn is the corresponding interface of SALDataOn
type SALDataOn interface {
	utils.LengthAware
	utils.Serializable
	SALData
	// GetGroup returns Group (property field)
	GetGroup() byte
}

// SALDataOnExactly can be used when we want exactly this type and not a type which fulfills SALDataOn.
// This is useful for switch cases.
type SALDataOnExactly interface {
	SALDataOn
	isSALDataOn() bool
}

// _SALDataOn is the data-structure of this message
type _SALDataOn struct {
	*_SALData
	Group byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataOn) InitializeParent(parent SALData, commandTypeContainer SALCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_SALDataOn) GetParent() SALData {
	return m._SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataOn) GetGroup() byte {
	return m.Group
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataOn factory function for _SALDataOn
func NewSALDataOn(group byte, commandTypeContainer SALCommandTypeContainer) *_SALDataOn {
	_result := &_SALDataOn{
		Group:    group,
		_SALData: NewSALData(commandTypeContainer),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataOn(structType interface{}) SALDataOn {
	if casted, ok := structType.(SALDataOn); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataOn); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataOn) GetTypeName() string {
	return "SALDataOn"
}

func (m *_SALDataOn) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SALDataOn) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (group)
	lengthInBits += 8

	return lengthInBits
}

func (m *_SALDataOn) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALDataOnParse(readBuffer utils.ReadBuffer) (SALDataOn, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataOn"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataOn")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (group)
	_group, _groupErr := readBuffer.ReadByte("group")
	if _groupErr != nil {
		return nil, errors.Wrap(_groupErr, "Error parsing 'group' field")
	}
	group := _group

	if closeErr := readBuffer.CloseContext("SALDataOn"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataOn")
	}

	// Create a partially initialized instance
	_child := &_SALDataOn{
		Group:    group,
		_SALData: &_SALData{},
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataOn) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataOn"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataOn")
		}

		// Simple Field (group)
		group := byte(m.GetGroup())
		_groupErr := writeBuffer.WriteByte("group", (group))
		if _groupErr != nil {
			return errors.Wrap(_groupErr, "Error serializing 'group' field")
		}

		if popErr := writeBuffer.PopContext("SALDataOn"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataOn")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SALDataOn) isSALDataOn() bool {
	return true
}

func (m *_SALDataOn) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
