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

// BACnetFaultParameter is the corresponding interface of BACnetFaultParameter
type BACnetFaultParameter interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetFaultParameterExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameter.
// This is useful for switch cases.
type BACnetFaultParameterExactly interface {
	BACnetFaultParameter
	isBACnetFaultParameter() bool
}

// _BACnetFaultParameter is the data-structure of this message
type _BACnetFaultParameter struct {
	_BACnetFaultParameterChildRequirements
	PeekedTagHeader BACnetTagHeader
}

type _BACnetFaultParameterChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}

type BACnetFaultParameterParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetFaultParameter, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetFaultParameterChild interface {
	utils.Serializable
	InitializeParent(parent BACnetFaultParameter, peekedTagHeader BACnetTagHeader)
	GetParent() *BACnetFaultParameter

	GetTypeName() string
	BACnetFaultParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameter) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetFaultParameter) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameter factory function for _BACnetFaultParameter
func NewBACnetFaultParameter(peekedTagHeader BACnetTagHeader) *_BACnetFaultParameter {
	return &_BACnetFaultParameter{PeekedTagHeader: peekedTagHeader}
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameter(structType any) BACnetFaultParameter {
	if casted, ok := structType.(BACnetFaultParameter); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameter); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameter) GetTypeName() string {
	return "BACnetFaultParameter"
}

func (m *_BACnetFaultParameter) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetFaultParameter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterParse(ctx context.Context, theBytes []byte) (BACnetFaultParameter, error) {
	return BACnetFaultParameterParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetFaultParameterParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameter, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetFaultParameter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Peek Field (peekedTagHeader)
	currentPos = positionAware.GetPos()
	if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
	}
	peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetFaultParameterChildSerializeRequirement interface {
		BACnetFaultParameter
		InitializeParent(BACnetFaultParameter, BACnetTagHeader)
		GetParent() BACnetFaultParameter
	}
	var _childTemp any
	var _child BACnetFaultParameterChildSerializeRequirement
	var typeSwitchError error
	switch {
	case peekedTagNumber == uint8(0): // BACnetFaultParameterNone
		_childTemp, typeSwitchError = BACnetFaultParameterNoneParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(1): // BACnetFaultParameterFaultCharacterString
		_childTemp, typeSwitchError = BACnetFaultParameterFaultCharacterStringParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(2): // BACnetFaultParameterFaultExtended
		_childTemp, typeSwitchError = BACnetFaultParameterFaultExtendedParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(3): // BACnetFaultParameterFaultLifeSafety
		_childTemp, typeSwitchError = BACnetFaultParameterFaultLifeSafetyParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(4): // BACnetFaultParameterFaultState
		_childTemp, typeSwitchError = BACnetFaultParameterFaultStateParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(5): // BACnetFaultParameterFaultStatusFlags
		_childTemp, typeSwitchError = BACnetFaultParameterFaultStatusFlagsParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(6): // BACnetFaultParameterFaultOutOfRange
		_childTemp, typeSwitchError = BACnetFaultParameterFaultOutOfRangeParseWithBuffer(ctx, readBuffer)
	case peekedTagNumber == uint8(7): // BACnetFaultParameterFaultListed
		_childTemp, typeSwitchError = BACnetFaultParameterFaultListedParseWithBuffer(ctx, readBuffer)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetFaultParameter")
	}
	_child = _childTemp.(BACnetFaultParameterChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetFaultParameter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameter")
	}

	// Finish initializing
	_child.InitializeParent(_child, peekedTagHeader)
	return _child, nil
}

func (pm *_BACnetFaultParameter) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetFaultParameter, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetFaultParameter"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameter")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ = peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetFaultParameter"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultParameter")
	}
	return nil
}

func (m *_BACnetFaultParameter) isBACnetFaultParameter() bool {
	return true
}

func (m *_BACnetFaultParameter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
