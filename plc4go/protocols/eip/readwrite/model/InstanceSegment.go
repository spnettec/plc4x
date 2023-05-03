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

// InstanceSegment is the corresponding interface of InstanceSegment
type InstanceSegment interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPathSegmentType returns PathSegmentType (property field)
	GetPathSegmentType() uint8
	// GetLogicalSegmentType returns LogicalSegmentType (property field)
	GetLogicalSegmentType() uint8
	// GetLogicalSegmentFormat returns LogicalSegmentFormat (property field)
	GetLogicalSegmentFormat() uint8
	// GetInstance returns Instance (property field)
	GetInstance() uint8
}

// InstanceSegmentExactly can be used when we want exactly this type and not a type which fulfills InstanceSegment.
// This is useful for switch cases.
type InstanceSegmentExactly interface {
	InstanceSegment
	isInstanceSegment() bool
}

// _InstanceSegment is the data-structure of this message
type _InstanceSegment struct {
	PathSegmentType      uint8
	LogicalSegmentType   uint8
	LogicalSegmentFormat uint8
	Instance             uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_InstanceSegment) GetPathSegmentType() uint8 {
	return m.PathSegmentType
}

func (m *_InstanceSegment) GetLogicalSegmentType() uint8 {
	return m.LogicalSegmentType
}

func (m *_InstanceSegment) GetLogicalSegmentFormat() uint8 {
	return m.LogicalSegmentFormat
}

func (m *_InstanceSegment) GetInstance() uint8 {
	return m.Instance
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewInstanceSegment factory function for _InstanceSegment
func NewInstanceSegment(pathSegmentType uint8, logicalSegmentType uint8, logicalSegmentFormat uint8, instance uint8) *_InstanceSegment {
	return &_InstanceSegment{PathSegmentType: pathSegmentType, LogicalSegmentType: logicalSegmentType, LogicalSegmentFormat: logicalSegmentFormat, Instance: instance}
}

// Deprecated: use the interface for direct cast
func CastInstanceSegment(structType any) InstanceSegment {
	if casted, ok := structType.(InstanceSegment); ok {
		return casted
	}
	if casted, ok := structType.(*InstanceSegment); ok {
		return *casted
	}
	return nil
}

func (m *_InstanceSegment) GetTypeName() string {
	return "InstanceSegment"
}

func (m *_InstanceSegment) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (pathSegmentType)
	lengthInBits += 3

	// Simple field (logicalSegmentType)
	lengthInBits += 3

	// Simple field (logicalSegmentFormat)
	lengthInBits += 2

	// Simple field (instance)
	lengthInBits += 8

	return lengthInBits
}

func (m *_InstanceSegment) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func InstanceSegmentParse(theBytes []byte) (InstanceSegment, error) {
	return InstanceSegmentParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func InstanceSegmentParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (InstanceSegment, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("InstanceSegment"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for InstanceSegment")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (pathSegmentType)
	_pathSegmentType, _pathSegmentTypeErr := readBuffer.ReadUint8("pathSegmentType", 3)
	if _pathSegmentTypeErr != nil {
		return nil, errors.Wrap(_pathSegmentTypeErr, "Error parsing 'pathSegmentType' field of InstanceSegment")
	}
	pathSegmentType := _pathSegmentType

	// Simple Field (logicalSegmentType)
	_logicalSegmentType, _logicalSegmentTypeErr := readBuffer.ReadUint8("logicalSegmentType", 3)
	if _logicalSegmentTypeErr != nil {
		return nil, errors.Wrap(_logicalSegmentTypeErr, "Error parsing 'logicalSegmentType' field of InstanceSegment")
	}
	logicalSegmentType := _logicalSegmentType

	// Simple Field (logicalSegmentFormat)
	_logicalSegmentFormat, _logicalSegmentFormatErr := readBuffer.ReadUint8("logicalSegmentFormat", 2)
	if _logicalSegmentFormatErr != nil {
		return nil, errors.Wrap(_logicalSegmentFormatErr, "Error parsing 'logicalSegmentFormat' field of InstanceSegment")
	}
	logicalSegmentFormat := _logicalSegmentFormat

	// Simple Field (instance)
	_instance, _instanceErr := readBuffer.ReadUint8("instance", 8)
	if _instanceErr != nil {
		return nil, errors.Wrap(_instanceErr, "Error parsing 'instance' field of InstanceSegment")
	}
	instance := _instance

	if closeErr := readBuffer.CloseContext("InstanceSegment"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for InstanceSegment")
	}

	// Create the instance
	return &_InstanceSegment{
		PathSegmentType:      pathSegmentType,
		LogicalSegmentType:   logicalSegmentType,
		LogicalSegmentFormat: logicalSegmentFormat,
		Instance:             instance,
	}, nil
}

func (m *_InstanceSegment) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_InstanceSegment) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("InstanceSegment"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for InstanceSegment")
	}

	// Simple Field (pathSegmentType)
	pathSegmentType := uint8(m.GetPathSegmentType())
	_pathSegmentTypeErr := writeBuffer.WriteUint8("pathSegmentType", 3, (pathSegmentType))
	if _pathSegmentTypeErr != nil {
		return errors.Wrap(_pathSegmentTypeErr, "Error serializing 'pathSegmentType' field")
	}

	// Simple Field (logicalSegmentType)
	logicalSegmentType := uint8(m.GetLogicalSegmentType())
	_logicalSegmentTypeErr := writeBuffer.WriteUint8("logicalSegmentType", 3, (logicalSegmentType))
	if _logicalSegmentTypeErr != nil {
		return errors.Wrap(_logicalSegmentTypeErr, "Error serializing 'logicalSegmentType' field")
	}

	// Simple Field (logicalSegmentFormat)
	logicalSegmentFormat := uint8(m.GetLogicalSegmentFormat())
	_logicalSegmentFormatErr := writeBuffer.WriteUint8("logicalSegmentFormat", 2, (logicalSegmentFormat))
	if _logicalSegmentFormatErr != nil {
		return errors.Wrap(_logicalSegmentFormatErr, "Error serializing 'logicalSegmentFormat' field")
	}

	// Simple Field (instance)
	instance := uint8(m.GetInstance())
	_instanceErr := writeBuffer.WriteUint8("instance", 8, (instance))
	if _instanceErr != nil {
		return errors.Wrap(_instanceErr, "Error serializing 'instance' field")
	}

	if popErr := writeBuffer.PopContext("InstanceSegment"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for InstanceSegment")
	}
	return nil
}

func (m *_InstanceSegment) isInstanceSegment() bool {
	return true
}

func (m *_InstanceSegment) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
