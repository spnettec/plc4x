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

// LogicalSegmentType is the corresponding interface of LogicalSegmentType
type LogicalSegmentType interface {
	utils.LengthAware
	utils.Serializable
	// GetLogicalSegmentType returns LogicalSegmentType (discriminator field)
	GetLogicalSegmentType() uint8
}

// LogicalSegmentTypeExactly can be used when we want exactly this type and not a type which fulfills LogicalSegmentType.
// This is useful for switch cases.
type LogicalSegmentTypeExactly interface {
	LogicalSegmentType
	isLogicalSegmentType() bool
}

// _LogicalSegmentType is the data-structure of this message
type _LogicalSegmentType struct {
	_LogicalSegmentTypeChildRequirements
}

type _LogicalSegmentTypeChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetLogicalSegmentType() uint8
}

type LogicalSegmentTypeParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child LogicalSegmentType, serializeChildFunction func() error) error
	GetTypeName() string
}

type LogicalSegmentTypeChild interface {
	utils.Serializable
	InitializeParent(parent LogicalSegmentType)
	GetParent() *LogicalSegmentType

	GetTypeName() string
	LogicalSegmentType
}

// NewLogicalSegmentType factory function for _LogicalSegmentType
func NewLogicalSegmentType() *_LogicalSegmentType {
	return &_LogicalSegmentType{}
}

// Deprecated: use the interface for direct cast
func CastLogicalSegmentType(structType interface{}) LogicalSegmentType {
	if casted, ok := structType.(LogicalSegmentType); ok {
		return casted
	}
	if casted, ok := structType.(*LogicalSegmentType); ok {
		return *casted
	}
	return nil
}

func (m *_LogicalSegmentType) GetTypeName() string {
	return "LogicalSegmentType"
}

func (m *_LogicalSegmentType) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (logicalSegmentType)
	lengthInBits += 3

	return lengthInBits
}

func (m *_LogicalSegmentType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func LogicalSegmentTypeParse(theBytes []byte) (LogicalSegmentType, error) {
	return LogicalSegmentTypeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func LogicalSegmentTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (LogicalSegmentType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LogicalSegmentType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LogicalSegmentType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (logicalSegmentType) (Used as input to a switch field)
	logicalSegmentType, _logicalSegmentTypeErr := readBuffer.ReadUint8("logicalSegmentType", 3)
	if _logicalSegmentTypeErr != nil {
		return nil, errors.Wrap(_logicalSegmentTypeErr, "Error parsing 'logicalSegmentType' field of LogicalSegmentType")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type LogicalSegmentTypeChildSerializeRequirement interface {
		LogicalSegmentType
		InitializeParent(LogicalSegmentType)
		GetParent() LogicalSegmentType
	}
	var _childTemp interface{}
	var _child LogicalSegmentTypeChildSerializeRequirement
	var typeSwitchError error
	switch {
	case logicalSegmentType == 0x00: // ClassID
		_childTemp, typeSwitchError = ClassIDParseWithBuffer(ctx, readBuffer)
	case logicalSegmentType == 0x01: // InstanceID
		_childTemp, typeSwitchError = InstanceIDParseWithBuffer(ctx, readBuffer)
	case logicalSegmentType == 0x02: // MemberID
		_childTemp, typeSwitchError = MemberIDParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [logicalSegmentType=%v]", logicalSegmentType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of LogicalSegmentType")
	}
	_child = _childTemp.(LogicalSegmentTypeChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("LogicalSegmentType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LogicalSegmentType")
	}

	// Finish initializing
	_child.InitializeParent(_child)
	return _child, nil
}

func (pm *_LogicalSegmentType) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child LogicalSegmentType, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("LogicalSegmentType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for LogicalSegmentType")
	}

	// Discriminator Field (logicalSegmentType) (Used as input to a switch field)
	logicalSegmentType := uint8(child.GetLogicalSegmentType())
	_logicalSegmentTypeErr := writeBuffer.WriteUint8("logicalSegmentType", 3, (logicalSegmentType))

	if _logicalSegmentTypeErr != nil {
		return errors.Wrap(_logicalSegmentTypeErr, "Error serializing 'logicalSegmentType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("LogicalSegmentType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for LogicalSegmentType")
	}
	return nil
}

func (m *_LogicalSegmentType) isLogicalSegmentType() bool {
	return true
}

func (m *_LogicalSegmentType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
