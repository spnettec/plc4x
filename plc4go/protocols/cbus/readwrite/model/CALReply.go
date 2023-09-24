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


// CALReply is the corresponding interface of CALReply
type CALReply interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetCalType returns CalType (property field)
	GetCalType() byte
	// GetCalData returns CalData (property field)
	GetCalData() CALData
}

// CALReplyExactly can be used when we want exactly this type and not a type which fulfills CALReply.
// This is useful for switch cases.
type CALReplyExactly interface {
	CALReply
	isCALReply() bool
}

// _CALReply is the data-structure of this message
type _CALReply struct {
	_CALReplyChildRequirements
        CalType byte
        CalData CALData

	// Arguments.
	CBusOptions CBusOptions
	RequestContext RequestContext
}

type _CALReplyChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetCalType() byte
}


type CALReplyParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CALReply, serializeChildFunction func() error) error
	GetTypeName() string
}

type CALReplyChild interface {
	utils.Serializable
InitializeParent(parent CALReply , calType byte , calData CALData )
	GetParent() *CALReply

	GetTypeName() string
	CALReply
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALReply) GetCalType() byte {
	return m.CalType
}

func (m *_CALReply) GetCalData() CALData {
	return m.CalData
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewCALReply factory function for _CALReply
func NewCALReply( calType byte , calData CALData , cBusOptions CBusOptions , requestContext RequestContext ) *_CALReply {
return &_CALReply{ CalType: calType , CalData: calData , CBusOptions: cBusOptions , RequestContext: requestContext }
}

// Deprecated: use the interface for direct cast
func CastCALReply(structType any) CALReply {
    if casted, ok := structType.(CALReply); ok {
		return casted
	}
	if casted, ok := structType.(*CALReply); ok {
		return *casted
	}
	return nil
}

func (m *_CALReply) GetTypeName() string {
	return "CALReply"
}


func (m *_CALReply) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (calData)
	lengthInBits += m.CalData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_CALReply) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CALReplyParse(ctx context.Context, theBytes []byte, cBusOptions CBusOptions, requestContext RequestContext) (CALReply, error) {
	return CALReplyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions, requestContext)
}

func CALReplyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions, requestContext RequestContext) (CALReply, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CALReply"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALReply")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

				// Peek Field (calType)
				currentPos = positionAware.GetPos()
				calType, _err := readBuffer.ReadByte("calType")
				if _err != nil {
					return nil, errors.Wrap(_err, "Error parsing 'calType' field of CALReply")
				}

				readBuffer.Reset(currentPos)

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type CALReplyChildSerializeRequirement interface {
		CALReply
		InitializeParent(CALReply,  byte, CALData)
		GetParent() CALReply
	}
	var _childTemp any
	var _child CALReplyChildSerializeRequirement
	var typeSwitchError error
	switch {
case calType == 0x86 : // CALReplyLong
		_childTemp, typeSwitchError = CALReplyLongParseWithBuffer(ctx, readBuffer, cBusOptions, requestContext)
case true : // CALReplyShort
		_childTemp, typeSwitchError = CALReplyShortParseWithBuffer(ctx, readBuffer, cBusOptions, requestContext)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [calType=%v]", calType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of CALReply")
	}
	_child = _childTemp.(CALReplyChildSerializeRequirement)

	// Simple Field (calData)
	if pullErr := readBuffer.PullContext("calData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for calData")
	}
_calData, _calDataErr := CALDataParseWithBuffer(ctx, readBuffer , requestContext )
	if _calDataErr != nil {
		return nil, errors.Wrap(_calDataErr, "Error parsing 'calData' field of CALReply")
	}
	calData := _calData.(CALData)
	if closeErr := readBuffer.CloseContext("calData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for calData")
	}

	if closeErr := readBuffer.CloseContext("CALReply"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALReply")
	}

	// Finish initializing
_child.InitializeParent(_child , calType , calData )
	return _child, nil
}

func (pm *_CALReply) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CALReply, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("CALReply"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CALReply")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (calData)
	if pushErr := writeBuffer.PushContext("calData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for calData")
	}
	_calDataErr := writeBuffer.WriteSerializable(ctx, m.GetCalData())
	if popErr := writeBuffer.PopContext("calData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for calData")
	}
	if _calDataErr != nil {
		return errors.Wrap(_calDataErr, "Error serializing 'calData' field")
	}

	if popErr := writeBuffer.PopContext("CALReply"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CALReply")
	}
	return nil
}


////
// Arguments Getter

func (m *_CALReply) GetCBusOptions() CBusOptions {
	return m.CBusOptions
}
func (m *_CALReply) GetRequestContext() RequestContext {
	return m.RequestContext
}
//
////

func (m *_CALReply) isCALReply() bool {
	return true
}

func (m *_CALReply) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



