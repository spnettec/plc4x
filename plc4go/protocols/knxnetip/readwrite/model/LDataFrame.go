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


// LDataFrame is the corresponding interface of LDataFrame
type LDataFrame interface {
	utils.LengthAware
	utils.Serializable
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
}

// LDataFrameExactly can be used when we want exactly this type and not a type which fulfills LDataFrame.
// This is useful for switch cases.
type LDataFrameExactly interface {
	LDataFrame
	isLDataFrame() bool
}

// _LDataFrame is the data-structure of this message
type _LDataFrame struct {
	_LDataFrameChildRequirements
        FrameType bool
        NotRepeated bool
        Priority CEMIPriority
        AcknowledgeRequested bool
        ErrorFlag bool
}

type _LDataFrameChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetNotAckFrame() bool
	GetPolling() bool
}


type LDataFrameParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child LDataFrame, serializeChildFunction func() error) error
	GetTypeName() string
}

type LDataFrameChild interface {
	utils.Serializable
InitializeParent(parent LDataFrame , frameType bool , notRepeated bool , priority CEMIPriority , acknowledgeRequested bool , errorFlag bool )
	GetParent() *LDataFrame

	GetTypeName() string
	LDataFrame
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LDataFrame) GetFrameType() bool {
	return m.FrameType
}

func (m *_LDataFrame) GetNotRepeated() bool {
	return m.NotRepeated
}

func (m *_LDataFrame) GetPriority() CEMIPriority {
	return m.Priority
}

func (m *_LDataFrame) GetAcknowledgeRequested() bool {
	return m.AcknowledgeRequested
}

func (m *_LDataFrame) GetErrorFlag() bool {
	return m.ErrorFlag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewLDataFrame factory function for _LDataFrame
func NewLDataFrame( frameType bool , notRepeated bool , priority CEMIPriority , acknowledgeRequested bool , errorFlag bool ) *_LDataFrame {
return &_LDataFrame{ FrameType: frameType , NotRepeated: notRepeated , Priority: priority , AcknowledgeRequested: acknowledgeRequested , ErrorFlag: errorFlag }
}

// Deprecated: use the interface for direct cast
func CastLDataFrame(structType interface{}) LDataFrame {
    if casted, ok := structType.(LDataFrame); ok {
		return casted
	}
	if casted, ok := structType.(*LDataFrame); ok {
		return *casted
	}
	return nil
}

func (m *_LDataFrame) GetTypeName() string {
	return "LDataFrame"
}


func (m *_LDataFrame) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (frameType)
	lengthInBits += 1;
	// Discriminator Field (polling)
	lengthInBits += 1;

	// Simple field (notRepeated)
	lengthInBits += 1;
	// Discriminator Field (notAckFrame)
	lengthInBits += 1;

	// Simple field (priority)
	lengthInBits += 2

	// Simple field (acknowledgeRequested)
	lengthInBits += 1;

	// Simple field (errorFlag)
	lengthInBits += 1;

	return lengthInBits
}

func (m *_LDataFrame) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LDataFrameParse(theBytes []byte) (LDataFrame, error) {
	return LDataFrameParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func LDataFrameParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LDataFrame, error) {
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
		return nil, errors.Wrap(_frameTypeErr, "Error parsing 'frameType' field of LDataFrame")
	}
	frameType := _frameType

	// Discriminator Field (polling) (Used as input to a switch field)
	polling, _pollingErr := readBuffer.ReadBit("polling")
	if _pollingErr != nil {
		return nil, errors.Wrap(_pollingErr, "Error parsing 'polling' field of LDataFrame")
	}

	// Simple Field (notRepeated)
_notRepeated, _notRepeatedErr := readBuffer.ReadBit("notRepeated")
	if _notRepeatedErr != nil {
		return nil, errors.Wrap(_notRepeatedErr, "Error parsing 'notRepeated' field of LDataFrame")
	}
	notRepeated := _notRepeated

	// Discriminator Field (notAckFrame) (Used as input to a switch field)
	notAckFrame, _notAckFrameErr := readBuffer.ReadBit("notAckFrame")
	if _notAckFrameErr != nil {
		return nil, errors.Wrap(_notAckFrameErr, "Error parsing 'notAckFrame' field of LDataFrame")
	}

	// Simple Field (priority)
	if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for priority")
	}
_priority, _priorityErr := CEMIPriorityParseWithBuffer(ctx, readBuffer)
	if _priorityErr != nil {
		return nil, errors.Wrap(_priorityErr, "Error parsing 'priority' field of LDataFrame")
	}
	priority := _priority
	if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for priority")
	}

	// Simple Field (acknowledgeRequested)
_acknowledgeRequested, _acknowledgeRequestedErr := readBuffer.ReadBit("acknowledgeRequested")
	if _acknowledgeRequestedErr != nil {
		return nil, errors.Wrap(_acknowledgeRequestedErr, "Error parsing 'acknowledgeRequested' field of LDataFrame")
	}
	acknowledgeRequested := _acknowledgeRequested

	// Simple Field (errorFlag)
_errorFlag, _errorFlagErr := readBuffer.ReadBit("errorFlag")
	if _errorFlagErr != nil {
		return nil, errors.Wrap(_errorFlagErr, "Error parsing 'errorFlag' field of LDataFrame")
	}
	errorFlag := _errorFlag

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type LDataFrameChildSerializeRequirement interface {
		LDataFrame
		InitializeParent(LDataFrame,  bool, bool, CEMIPriority, bool, bool)
		GetParent() LDataFrame
	}
	var _childTemp interface{}
	var _child LDataFrameChildSerializeRequirement
	var typeSwitchError error
	switch {
case notAckFrame == bool(true) && polling == bool(false) : // LDataExtended
		_childTemp, typeSwitchError = LDataExtendedParseWithBuffer(ctx, readBuffer, )
case notAckFrame == bool(true) && polling == bool(true) : // LPollData
		_childTemp, typeSwitchError = LPollDataParseWithBuffer(ctx, readBuffer, )
case notAckFrame == bool(false) : // LDataFrameACK
		_childTemp, typeSwitchError = LDataFrameACKParseWithBuffer(ctx, readBuffer, )
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [notAckFrame=%v, polling=%v]", notAckFrame, polling)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of LDataFrame")
	}
	_child = _childTemp.(LDataFrameChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("LDataFrame"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LDataFrame")
	}

	// Finish initializing
_child.InitializeParent(_child , frameType , notRepeated , priority , acknowledgeRequested , errorFlag )
	return _child, nil
}

func (pm *_LDataFrame) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child LDataFrame, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("LDataFrame"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for LDataFrame")
	}

	// Simple Field (frameType)
	frameType := bool(m.GetFrameType())
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
	notRepeated := bool(m.GetNotRepeated())
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
	_priorityErr := writeBuffer.WriteSerializable(ctx, m.GetPriority())
	if popErr := writeBuffer.PopContext("priority"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for priority")
	}
	if _priorityErr != nil {
		return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
	}

	// Simple Field (acknowledgeRequested)
	acknowledgeRequested := bool(m.GetAcknowledgeRequested())
	_acknowledgeRequestedErr := writeBuffer.WriteBit("acknowledgeRequested", (acknowledgeRequested))
	if _acknowledgeRequestedErr != nil {
		return errors.Wrap(_acknowledgeRequestedErr, "Error serializing 'acknowledgeRequested' field")
	}

	// Simple Field (errorFlag)
	errorFlag := bool(m.GetErrorFlag())
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


func (m *_LDataFrame) isLDataFrame() bool {
	return true
}

func (m *_LDataFrame) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



