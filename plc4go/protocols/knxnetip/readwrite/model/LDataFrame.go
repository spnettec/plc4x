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

// LDataFrame is the data-structure of this message
type LDataFrame struct {
	FrameType            bool
	NotRepeated          bool
	Priority             CEMIPriority
	AcknowledgeRequested bool
	ErrorFlag            bool
	Child                ILDataFrameChild
}

// ILDataFrame is the corresponding interface of LDataFrame
type ILDataFrame interface {
	// GetNotAckFrame returns NotAckFrame (discriminator field)
	GetNotAckFrame() bool
	// GetPolling returns Polling (discriminator field)
	GetPolling() bool
	// GetFrameType returns FrameType (property field)
	GetFrameType() bool
	// GetNotRepeated returns NotRepeated (property field)
	GetNotRepeated() bool
	// GetPriority returns Priority (property field)
	GetPriority() CEMIPriority
	// GetAcknowledgeRequested returns AcknowledgeRequested (property field)
	GetAcknowledgeRequested() bool
	// GetErrorFlag returns ErrorFlag (property field)
	GetErrorFlag() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type ILDataFrameParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child ILDataFrame, serializeChildFunction func() error) error
	GetTypeName() string
}

type ILDataFrameChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *LDataFrame, frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool)
	GetParent() *LDataFrame

	GetTypeName() string
	ILDataFrame
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *LDataFrame) GetFrameType() bool {
	return m.FrameType
}

func (m *LDataFrame) GetNotRepeated() bool {
	return m.NotRepeated
}

func (m *LDataFrame) GetPriority() CEMIPriority {
	return m.Priority
}

func (m *LDataFrame) GetAcknowledgeRequested() bool {
	return m.AcknowledgeRequested
}

func (m *LDataFrame) GetErrorFlag() bool {
	return m.ErrorFlag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewLDataFrame factory function for LDataFrame
func NewLDataFrame(frameType bool, notRepeated bool, priority CEMIPriority, acknowledgeRequested bool, errorFlag bool) *LDataFrame {
	return &LDataFrame{FrameType: frameType, NotRepeated: notRepeated, Priority: priority, AcknowledgeRequested: acknowledgeRequested, ErrorFlag: errorFlag}
}

func CastLDataFrame(structType interface{}) *LDataFrame {
	if casted, ok := structType.(LDataFrame); ok {
		return &casted
	}
	if casted, ok := structType.(*LDataFrame); ok {
		return casted
	}
	if casted, ok := structType.(ILDataFrameChild); ok {
		return casted.GetParent()
	}
	return nil
}

func (m *LDataFrame) GetTypeName() string {
	return "LDataFrame"
}

func (m *LDataFrame) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *LDataFrame) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *LDataFrame) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// Simple field (frameType)
	lengthInBits += 1
	// Discriminator Field (polling)
	lengthInBits += 1

	// Simple field (notRepeated)
	lengthInBits += 1
	// Discriminator Field (notAckFrame)
	lengthInBits += 1

	// Simple field (priority)
	lengthInBits += 2

	// Simple field (acknowledgeRequested)
	lengthInBits += 1

	// Simple field (errorFlag)
	lengthInBits += 1

	return lengthInBits
}

func (m *LDataFrame) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LDataFrameParse(readBuffer utils.ReadBuffer) (*LDataFrame, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LDataFrame"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LDataFrame")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (frameType)
	_frameType, _frameTypeErr := readBuffer.ReadBit("frameType")
	if _frameTypeErr != nil {
		return nil, errors.Wrap(_frameTypeErr, "Error parsing 'frameType' field")
	}
	frameType := _frameType

	// Discriminator Field (polling) (Used as input to a switch field)
	polling, _pollingErr := readBuffer.ReadBit("polling")
	if _pollingErr != nil {
		return nil, errors.Wrap(_pollingErr, "Error parsing 'polling' field")
	}

	// Simple Field (notRepeated)
	_notRepeated, _notRepeatedErr := readBuffer.ReadBit("notRepeated")
	if _notRepeatedErr != nil {
		return nil, errors.Wrap(_notRepeatedErr, "Error parsing 'notRepeated' field")
	}
	notRepeated := _notRepeated

	// Discriminator Field (notAckFrame) (Used as input to a switch field)
	notAckFrame, _notAckFrameErr := readBuffer.ReadBit("notAckFrame")
	if _notAckFrameErr != nil {
		return nil, errors.Wrap(_notAckFrameErr, "Error parsing 'notAckFrame' field")
	}

	// Simple Field (priority)
	if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for priority")
	}
	_priority, _priorityErr := CEMIPriorityParse(readBuffer)
	if _priorityErr != nil {
		return nil, errors.Wrap(_priorityErr, "Error parsing 'priority' field")
	}
	priority := _priority
	if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for priority")
	}

	// Simple Field (acknowledgeRequested)
	_acknowledgeRequested, _acknowledgeRequestedErr := readBuffer.ReadBit("acknowledgeRequested")
	if _acknowledgeRequestedErr != nil {
		return nil, errors.Wrap(_acknowledgeRequestedErr, "Error parsing 'acknowledgeRequested' field")
	}
	acknowledgeRequested := _acknowledgeRequested

	// Simple Field (errorFlag)
	_errorFlag, _errorFlagErr := readBuffer.ReadBit("errorFlag")
	if _errorFlagErr != nil {
		return nil, errors.Wrap(_errorFlagErr, "Error parsing 'errorFlag' field")
	}
	errorFlag := _errorFlag

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type LDataFrameChild interface {
		InitializeParent(*LDataFrame, bool, bool, CEMIPriority, bool, bool)
		GetParent() *LDataFrame
	}
	var _child LDataFrameChild
	var typeSwitchError error
	switch {
	case notAckFrame == bool(true) && polling == bool(false): // LDataExtended
		_child, typeSwitchError = LDataExtendedParse(readBuffer)
	case notAckFrame == bool(true) && polling == bool(true): // LPollData
		_child, typeSwitchError = LPollDataParse(readBuffer)
	case notAckFrame == bool(false): // LDataFrameACK
		_child, typeSwitchError = LDataFrameACKParse(readBuffer)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("LDataFrame"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LDataFrame")
	}

	// Finish initializing
	_child.InitializeParent(_child.GetParent(), frameType, notRepeated, priority, acknowledgeRequested, errorFlag)
	return _child.GetParent(), nil
}

func (m *LDataFrame) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *LDataFrame) SerializeParent(writeBuffer utils.WriteBuffer, child ILDataFrame, serializeChildFunction func() error) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("LDataFrame"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for LDataFrame")
	}

	// Simple Field (frameType)
	frameType := bool(m.FrameType)
	_frameTypeErr := writeBuffer.WriteBit("frameType", (frameType))
	if _frameTypeErr != nil {
		return errors.Wrap(_frameTypeErr, "Error serializing 'frameType' field")
	}

	// Discriminator Field (polling) (Used as input to a switch field)
	polling := bool(child.GetPolling())
	_pollingErr := writeBuffer.WriteBit("polling", (polling))

	if _pollingErr != nil {
		return errors.Wrap(_pollingErr, "Error serializing 'polling' field")
	}

	// Simple Field (notRepeated)
	notRepeated := bool(m.NotRepeated)
	_notRepeatedErr := writeBuffer.WriteBit("notRepeated", (notRepeated))
	if _notRepeatedErr != nil {
		return errors.Wrap(_notRepeatedErr, "Error serializing 'notRepeated' field")
	}

	// Discriminator Field (notAckFrame) (Used as input to a switch field)
	notAckFrame := bool(child.GetNotAckFrame())
	_notAckFrameErr := writeBuffer.WriteBit("notAckFrame", (notAckFrame))

	if _notAckFrameErr != nil {
		return errors.Wrap(_notAckFrameErr, "Error serializing 'notAckFrame' field")
	}

	// Simple Field (priority)
	if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for priority")
	}
	_priorityErr := writeBuffer.WriteSerializable(m.Priority)
	if popErr := writeBuffer.PopContext("priority"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for priority")
	}
	if _priorityErr != nil {
		return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
	}

	// Simple Field (acknowledgeRequested)
	acknowledgeRequested := bool(m.AcknowledgeRequested)
	_acknowledgeRequestedErr := writeBuffer.WriteBit("acknowledgeRequested", (acknowledgeRequested))
	if _acknowledgeRequestedErr != nil {
		return errors.Wrap(_acknowledgeRequestedErr, "Error serializing 'acknowledgeRequested' field")
	}

	// Simple Field (errorFlag)
	errorFlag := bool(m.ErrorFlag)
	_errorFlagErr := writeBuffer.WriteBit("errorFlag", (errorFlag))
	if _errorFlagErr != nil {
		return errors.Wrap(_errorFlagErr, "Error serializing 'errorFlag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("LDataFrame"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for LDataFrame")
	}
	return nil
}

func (m *LDataFrame) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
