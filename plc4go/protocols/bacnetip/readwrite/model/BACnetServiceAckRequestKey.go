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

// BACnetServiceAckRequestKey is the corresponding interface of BACnetServiceAckRequestKey
type BACnetServiceAckRequestKey interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetServiceAck
	// GetBytesOfRemovedService returns BytesOfRemovedService (property field)
	GetBytesOfRemovedService() []byte
}

// BACnetServiceAckRequestKeyExactly can be used when we want exactly this type and not a type which fulfills BACnetServiceAckRequestKey.
// This is useful for switch cases.
type BACnetServiceAckRequestKeyExactly interface {
	BACnetServiceAckRequestKey
	isBACnetServiceAckRequestKey() bool
}

// _BACnetServiceAckRequestKey is the data-structure of this message
type _BACnetServiceAckRequestKey struct {
	*_BACnetServiceAck
	BytesOfRemovedService []byte

	// Arguments.
	ServiceAckPayloadLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetServiceAckRequestKey) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_REQUEST_KEY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetServiceAckRequestKey) InitializeParent(parent BACnetServiceAck) {}

func (m *_BACnetServiceAckRequestKey) GetParent() BACnetServiceAck {
	return m._BACnetServiceAck
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetServiceAckRequestKey) GetBytesOfRemovedService() []byte {
	return m.BytesOfRemovedService
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetServiceAckRequestKey factory function for _BACnetServiceAckRequestKey
func NewBACnetServiceAckRequestKey(bytesOfRemovedService []byte, serviceAckPayloadLength uint32, serviceAckLength uint32) *_BACnetServiceAckRequestKey {
	_result := &_BACnetServiceAckRequestKey{
		BytesOfRemovedService: bytesOfRemovedService,
		_BACnetServiceAck:     NewBACnetServiceAck(serviceAckLength),
	}
	_result._BACnetServiceAck._BACnetServiceAckChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetServiceAckRequestKey(structType any) BACnetServiceAckRequestKey {
	if casted, ok := structType.(BACnetServiceAckRequestKey); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetServiceAckRequestKey); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetServiceAckRequestKey) GetTypeName() string {
	return "BACnetServiceAckRequestKey"
}

func (m *_BACnetServiceAckRequestKey) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.BytesOfRemovedService) > 0 {
		lengthInBits += 8 * uint16(len(m.BytesOfRemovedService))
	}

	return lengthInBits
}

func (m *_BACnetServiceAckRequestKey) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetServiceAckRequestKeyParse(ctx context.Context, theBytes []byte, serviceAckPayloadLength uint32, serviceAckLength uint32) (BACnetServiceAckRequestKey, error) {
	return BACnetServiceAckRequestKeyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceAckPayloadLength, serviceAckLength)
}

func BACnetServiceAckRequestKeyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceAckPayloadLength uint32, serviceAckLength uint32) (BACnetServiceAckRequestKey, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetServiceAckRequestKey"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetServiceAckRequestKey")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (bytesOfRemovedService)
	numberOfBytesbytesOfRemovedService := int(serviceAckPayloadLength)
	bytesOfRemovedService, _readArrayErr := readBuffer.ReadByteArray("bytesOfRemovedService", numberOfBytesbytesOfRemovedService)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'bytesOfRemovedService' field of BACnetServiceAckRequestKey")
	}

	if closeErr := readBuffer.CloseContext("BACnetServiceAckRequestKey"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetServiceAckRequestKey")
	}

	// Create a partially initialized instance
	_child := &_BACnetServiceAckRequestKey{
		_BACnetServiceAck: &_BACnetServiceAck{
			ServiceAckLength: serviceAckLength,
		},
		BytesOfRemovedService: bytesOfRemovedService,
	}
	_child._BACnetServiceAck._BACnetServiceAckChildRequirements = _child
	return _child, nil
}

func (m *_BACnetServiceAckRequestKey) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetServiceAckRequestKey) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetServiceAckRequestKey"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetServiceAckRequestKey")
		}

		// Array Field (bytesOfRemovedService)
		// Byte Array field (bytesOfRemovedService)
		if err := writeBuffer.WriteByteArray("bytesOfRemovedService", m.GetBytesOfRemovedService()); err != nil {
			return errors.Wrap(err, "Error serializing 'bytesOfRemovedService' field")
		}

		if popErr := writeBuffer.PopContext("BACnetServiceAckRequestKey"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetServiceAckRequestKey")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetServiceAckRequestKey) GetServiceAckPayloadLength() uint32 {
	return m.ServiceAckPayloadLength
}

//
////

func (m *_BACnetServiceAckRequestKey) isBACnetServiceAckRequestKey() bool {
	return true
}

func (m *_BACnetServiceAckRequestKey) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
