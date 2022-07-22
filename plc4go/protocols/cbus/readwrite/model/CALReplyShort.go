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

// CALReplyShort is the corresponding interface of CALReplyShort
type CALReplyShort interface {
	utils.LengthAware
	utils.Serializable
	CALReply
}

// CALReplyShortExactly can be used when we want exactly this type and not a type which fulfills CALReplyShort.
// This is useful for switch cases.
type CALReplyShortExactly interface {
	CALReplyShort
	isCALReplyShort() bool
}

// _CALReplyShort is the data-structure of this message
type _CALReplyShort struct {
	*_CALReply
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALReplyShort) InitializeParent(parent CALReply, calType byte, calData CALData) {
	m.CalType = calType
	m.CalData = calData
}

func (m *_CALReplyShort) GetParent() CALReply {
	return m._CALReply
}

// NewCALReplyShort factory function for _CALReplyShort
func NewCALReplyShort(calType byte, calData CALData, cBusOptions CBusOptions, requestContext RequestContext) *_CALReplyShort {
	_result := &_CALReplyShort{
		_CALReply: NewCALReply(calType, calData, cBusOptions, requestContext),
	}
	_result._CALReply._CALReplyChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALReplyShort(structType interface{}) CALReplyShort {
	if casted, ok := structType.(CALReplyShort); ok {
		return casted
	}
	if casted, ok := structType.(*CALReplyShort); ok {
		return *casted
	}
	return nil
}

func (m *_CALReplyShort) GetTypeName() string {
	return "CALReplyShort"
}

func (m *_CALReplyShort) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CALReplyShort) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_CALReplyShort) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALReplyShortParse(readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (CALReplyShort, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALReplyShort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALReplyShort")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("CALReplyShort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALReplyShort")
	}

	// Create a partially initialized instance
	_child := &_CALReplyShort{
		_CALReply: &_CALReply{
			CBusOptions:    cBusOptions,
			RequestContext: requestContext,
		},
	}
	_child._CALReply._CALReplyChildRequirements = _child
	return _child, nil
}

func (m *_CALReplyShort) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALReplyShort"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALReplyShort")
		}

		if popErr := writeBuffer.PopContext("CALReplyShort"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALReplyShort")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CALReplyShort) isCALReplyShort() bool {
	return true
}

func (m *_CALReplyShort) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
