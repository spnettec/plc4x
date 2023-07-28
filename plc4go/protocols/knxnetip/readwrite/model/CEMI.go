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


// CEMI is the corresponding interface of CEMI
type CEMI interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetMessageCode returns MessageCode (discriminator field)
	GetMessageCode() uint8
}

// CEMIExactly can be used when we want exactly this type and not a type which fulfills CEMI.
// This is useful for switch cases.
type CEMIExactly interface {
	CEMI
	isCEMI() bool
}

// _CEMI is the data-structure of this message
type _CEMI struct {
	_CEMIChildRequirements

	// Arguments.
	Size uint16
}

type _CEMIChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetMessageCode() uint8
}


type CEMIParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CEMI, serializeChildFunction func() error) error
	GetTypeName() string
}

type CEMIChild interface {
	utils.Serializable
InitializeParent(parent CEMI )
	GetParent() *CEMI

	GetTypeName() string
	CEMI
}


// NewCEMI factory function for _CEMI
func NewCEMI( size uint16 ) *_CEMI {
return &_CEMI{ Size: size }
}

// Deprecated: use the interface for direct cast
func CastCEMI(structType any) CEMI {
    if casted, ok := structType.(CEMI); ok {
		return casted
	}
	if casted, ok := structType.(*CEMI); ok {
		return *casted
	}
	return nil
}

func (m *_CEMI) GetTypeName() string {
	return "CEMI"
}


func (m *_CEMI) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (messageCode)
	lengthInBits += 8;

	return lengthInBits
}

func (m *_CEMI) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CEMIParse(ctx context.Context, theBytes []byte, size uint16) (CEMI, error) {
	return CEMIParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), size)
}

func CEMIParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, size uint16) (CEMI, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CEMI"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CEMI")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (messageCode) (Used as input to a switch field)
	messageCode, _messageCodeErr := readBuffer.ReadUint8("messageCode", 8)
	if _messageCodeErr != nil {
		return nil, errors.Wrap(_messageCodeErr, "Error parsing 'messageCode' field of CEMI")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type CEMIChildSerializeRequirement interface {
		CEMI
		InitializeParent(CEMI )
		GetParent() CEMI
	}
	var _childTemp any
	var _child CEMIChildSerializeRequirement
	var typeSwitchError error
	switch {
case messageCode == 0x2B : // LBusmonInd
		_childTemp, typeSwitchError = LBusmonIndParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0x11 : // LDataReq
		_childTemp, typeSwitchError = LDataReqParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0x29 : // LDataInd
		_childTemp, typeSwitchError = LDataIndParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0x2E : // LDataCon
		_childTemp, typeSwitchError = LDataConParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0x10 : // LRawReq
		_childTemp, typeSwitchError = LRawReqParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0x2D : // LRawInd
		_childTemp, typeSwitchError = LRawIndParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0x2F : // LRawCon
		_childTemp, typeSwitchError = LRawConParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0x13 : // LPollDataReq
		_childTemp, typeSwitchError = LPollDataReqParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0x25 : // LPollDataCon
		_childTemp, typeSwitchError = LPollDataConParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0x41 : // TDataConnectedReq
		_childTemp, typeSwitchError = TDataConnectedReqParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0x89 : // TDataConnectedInd
		_childTemp, typeSwitchError = TDataConnectedIndParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0x4A : // TDataIndividualReq
		_childTemp, typeSwitchError = TDataIndividualReqParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0x94 : // TDataIndividualInd
		_childTemp, typeSwitchError = TDataIndividualIndParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0xFC : // MPropReadReq
		_childTemp, typeSwitchError = MPropReadReqParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0xFB : // MPropReadCon
		_childTemp, typeSwitchError = MPropReadConParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0xF6 : // MPropWriteReq
		_childTemp, typeSwitchError = MPropWriteReqParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0xF5 : // MPropWriteCon
		_childTemp, typeSwitchError = MPropWriteConParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0xF7 : // MPropInfoInd
		_childTemp, typeSwitchError = MPropInfoIndParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0xF8 : // MFuncPropCommandReq
		_childTemp, typeSwitchError = MFuncPropCommandReqParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0xF9 : // MFuncPropStateReadReq
		_childTemp, typeSwitchError = MFuncPropStateReadReqParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0xFA : // MFuncPropCon
		_childTemp, typeSwitchError = MFuncPropConParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0xF1 : // MResetReq
		_childTemp, typeSwitchError = MResetReqParseWithBuffer(ctx, readBuffer, size)
case messageCode == 0xF0 : // MResetInd
		_childTemp, typeSwitchError = MResetIndParseWithBuffer(ctx, readBuffer, size)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [messageCode=%v]", messageCode)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of CEMI")
	}
	_child = _childTemp.(CEMIChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("CEMI"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CEMI")
	}

	// Finish initializing
_child.InitializeParent(_child )
	return _child, nil
}

func (pm *_CEMI) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child CEMI, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("CEMI"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CEMI")
	}

	// Discriminator Field (messageCode) (Used as input to a switch field)
	messageCode := uint8(child.GetMessageCode())
	_messageCodeErr := writeBuffer.WriteUint8("messageCode", 8, (messageCode))

	if _messageCodeErr != nil {
		return errors.Wrap(_messageCodeErr, "Error serializing 'messageCode' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CEMI"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CEMI")
	}
	return nil
}


////
// Arguments Getter

func (m *_CEMI) GetSize() uint16 {
	return m.Size
}
//
////

func (m *_CEMI) isCEMI() bool {
	return true
}

func (m *_CEMI) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



