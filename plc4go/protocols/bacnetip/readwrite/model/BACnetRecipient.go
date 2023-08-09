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


// BACnetRecipient is the corresponding interface of BACnetRecipient
type BACnetRecipient interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetRecipientExactly can be used when we want exactly this type and not a type which fulfills BACnetRecipient.
// This is useful for switch cases.
type BACnetRecipientExactly interface {
	BACnetRecipient
	isBACnetRecipient() bool
}

// _BACnetRecipient is the data-structure of this message
type _BACnetRecipient struct {
	_BACnetRecipientChildRequirements
        PeekedTagHeader BACnetTagHeader
}

type _BACnetRecipientChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}


type BACnetRecipientParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetRecipient, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetRecipientChild interface {
	utils.Serializable
InitializeParent(parent BACnetRecipient , peekedTagHeader BACnetTagHeader )
	GetParent() *BACnetRecipient

	GetTypeName() string
	BACnetRecipient
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRecipient) GetPeekedTagHeader() BACnetTagHeader {
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

func (m *_BACnetRecipient) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetRecipient factory function for _BACnetRecipient
func NewBACnetRecipient( peekedTagHeader BACnetTagHeader ) *_BACnetRecipient {
return &_BACnetRecipient{ PeekedTagHeader: peekedTagHeader }
}

// Deprecated: use the interface for direct cast
func CastBACnetRecipient(structType any) BACnetRecipient {
    if casted, ok := structType.(BACnetRecipient); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRecipient); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRecipient) GetTypeName() string {
	return "BACnetRecipient"
}


func (m *_BACnetRecipient) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetRecipient) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRecipientParse(ctx context.Context, theBytes []byte) (BACnetRecipient, error) {
	return BACnetRecipientParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetRecipientParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRecipient, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetRecipient"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRecipient")
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
	type BACnetRecipientChildSerializeRequirement interface {
		BACnetRecipient
		InitializeParent(BACnetRecipient,  BACnetTagHeader)
		GetParent() BACnetRecipient
	}
	var _childTemp any
	var _child BACnetRecipientChildSerializeRequirement
	var typeSwitchError error
	switch {
case peekedTagNumber == uint8(0) : // BACnetRecipientDevice
		_childTemp, typeSwitchError = BACnetRecipientDeviceParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(1) : // BACnetRecipientAddress
		_childTemp, typeSwitchError = BACnetRecipientAddressParseWithBuffer(ctx, readBuffer, )
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetRecipient")
	}
	_child = _childTemp.(BACnetRecipientChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetRecipient"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRecipient")
	}

	// Finish initializing
_child.InitializeParent(_child , peekedTagHeader )
	return _child, nil
}

func (pm *_BACnetRecipient) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetRecipient, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("BACnetRecipient"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRecipient")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ =	peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetRecipient"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRecipient")
	}
	return nil
}


func (m *_BACnetRecipient) isBACnetRecipient() bool {
	return true
}

func (m *_BACnetRecipient) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



