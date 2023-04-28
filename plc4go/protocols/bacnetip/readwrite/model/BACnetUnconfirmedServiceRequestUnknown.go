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

// BACnetUnconfirmedServiceRequestUnknown is the corresponding interface of BACnetUnconfirmedServiceRequestUnknown
type BACnetUnconfirmedServiceRequestUnknown interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetUnconfirmedServiceRequest
	// GetUnknownBytes returns UnknownBytes (property field)
	GetUnknownBytes() []byte
}

// BACnetUnconfirmedServiceRequestUnknownExactly can be used when we want exactly this type and not a type which fulfills BACnetUnconfirmedServiceRequestUnknown.
// This is useful for switch cases.
type BACnetUnconfirmedServiceRequestUnknownExactly interface {
	BACnetUnconfirmedServiceRequestUnknown
	isBACnetUnconfirmedServiceRequestUnknown() bool
}

// _BACnetUnconfirmedServiceRequestUnknown is the data-structure of this message
type _BACnetUnconfirmedServiceRequestUnknown struct {
	*_BACnetUnconfirmedServiceRequest
	UnknownBytes []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnknown) GetServiceChoice() BACnetUnconfirmedServiceChoice {
	return 0
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnknown) InitializeParent(parent BACnetUnconfirmedServiceRequest) {
}

func (m *_BACnetUnconfirmedServiceRequestUnknown) GetParent() BACnetUnconfirmedServiceRequest {
	return m._BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetUnconfirmedServiceRequestUnknown) GetUnknownBytes() []byte {
	return m.UnknownBytes
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestUnknown factory function for _BACnetUnconfirmedServiceRequestUnknown
func NewBACnetUnconfirmedServiceRequestUnknown(unknownBytes []byte, serviceRequestLength uint16) *_BACnetUnconfirmedServiceRequestUnknown {
	_result := &_BACnetUnconfirmedServiceRequestUnknown{
		UnknownBytes:                     unknownBytes,
		_BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetUnconfirmedServiceRequestUnknown(structType any) BACnetUnconfirmedServiceRequestUnknown {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUnknown); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUnknown); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetUnconfirmedServiceRequestUnknown) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnknown"
}

func (m *_BACnetUnconfirmedServiceRequestUnknown) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.UnknownBytes) > 0 {
		lengthInBits += 8 * uint16(len(m.UnknownBytes))
	}

	return lengthInBits
}

func (m *_BACnetUnconfirmedServiceRequestUnknown) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetUnconfirmedServiceRequestUnknownParse(theBytes []byte, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestUnknown, error) {
	return BACnetUnconfirmedServiceRequestUnknownParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetUnconfirmedServiceRequestUnknownParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetUnconfirmedServiceRequestUnknown, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnknown"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetUnconfirmedServiceRequestUnknown")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (unknownBytes)
	numberOfBytesunknownBytes := int(utils.InlineIf((bool((serviceRequestLength) > (0))), func() any { return uint16((uint16(serviceRequestLength) - uint16(uint16(1)))) }, func() any { return uint16(uint16(0)) }).(uint16))
	unknownBytes, _readArrayErr := readBuffer.ReadByteArray("unknownBytes", numberOfBytesunknownBytes)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'unknownBytes' field of BACnetUnconfirmedServiceRequestUnknown")
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnknown"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetUnconfirmedServiceRequestUnknown")
	}

	// Create a partially initialized instance
	_child := &_BACnetUnconfirmedServiceRequestUnknown{
		_BACnetUnconfirmedServiceRequest: &_BACnetUnconfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		UnknownBytes: unknownBytes,
	}
	_child._BACnetUnconfirmedServiceRequest._BACnetUnconfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetUnconfirmedServiceRequestUnknown) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetUnconfirmedServiceRequestUnknown) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnknown"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetUnconfirmedServiceRequestUnknown")
		}

		// Array Field (unknownBytes)
		// Byte Array field (unknownBytes)
		if err := writeBuffer.WriteByteArray("unknownBytes", m.GetUnknownBytes()); err != nil {
			return errors.Wrap(err, "Error serializing 'unknownBytes' field")
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnknown"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetUnconfirmedServiceRequestUnknown")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetUnconfirmedServiceRequestUnknown) isBACnetUnconfirmedServiceRequestUnknown() bool {
	return true
}

func (m *_BACnetUnconfirmedServiceRequestUnknown) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
