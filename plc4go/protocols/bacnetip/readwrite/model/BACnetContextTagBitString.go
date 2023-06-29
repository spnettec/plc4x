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

// BACnetContextTagBitString is the corresponding interface of BACnetContextTagBitString
type BACnetContextTagBitString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBitString
}

// BACnetContextTagBitStringExactly can be used when we want exactly this type and not a type which fulfills BACnetContextTagBitString.
// This is useful for switch cases.
type BACnetContextTagBitStringExactly interface {
	BACnetContextTagBitString
	isBACnetContextTagBitString() bool
}

// _BACnetContextTagBitString is the data-structure of this message
type _BACnetContextTagBitString struct {
	*_BACnetContextTag
	Payload BACnetTagPayloadBitString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagBitString) GetDataType() BACnetDataType {
	return BACnetDataType_BIT_STRING
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagBitString) InitializeParent(parent BACnetContextTag, header BACnetTagHeader) {
	m.Header = header
}

func (m *_BACnetContextTagBitString) GetParent() BACnetContextTag {
	return m._BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagBitString) GetPayload() BACnetTagPayloadBitString {
	return m.Payload
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagBitString factory function for _BACnetContextTagBitString
func NewBACnetContextTagBitString(payload BACnetTagPayloadBitString, header BACnetTagHeader, tagNumberArgument uint8) *_BACnetContextTagBitString {
	_result := &_BACnetContextTagBitString{
		Payload:           payload,
		_BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	_result._BACnetContextTag._BACnetContextTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagBitString(structType any) BACnetContextTagBitString {
	if casted, ok := structType.(BACnetContextTagBitString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagBitString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagBitString) GetTypeName() string {
	return "BACnetContextTagBitString"
}

func (m *_BACnetContextTagBitString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetContextTagBitString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetContextTagBitStringParse(ctx context.Context, theBytes []byte, header BACnetTagHeader, tagNumberArgument uint8, dataType BACnetDataType) (BACnetContextTagBitString, error) {
	return BACnetContextTagBitStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), header, tagNumberArgument, dataType)
}

func BACnetContextTagBitStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, header BACnetTagHeader, tagNumberArgument uint8, dataType BACnetDataType) (BACnetContextTagBitString, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetContextTagBitString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagBitString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
	_payload, _payloadErr := BACnetTagPayloadBitStringParseWithBuffer(ctx, readBuffer, uint32(header.GetActualLength()))
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field of BACnetContextTagBitString")
	}
	payload := _payload.(BACnetTagPayloadBitString)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTagBitString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagBitString")
	}

	// Create a partially initialized instance
	_child := &_BACnetContextTagBitString{
		_BACnetContextTag: &_BACnetContextTag{
			TagNumberArgument: tagNumberArgument,
		},
		Payload: payload,
	}
	_child._BACnetContextTag._BACnetContextTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetContextTagBitString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagBitString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagBitString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagBitString")
		}

		// Simple Field (payload)
		if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for payload")
		}
		_payloadErr := writeBuffer.WriteSerializable(ctx, m.GetPayload())
		if popErr := writeBuffer.PopContext("payload"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for payload")
		}
		if _payloadErr != nil {
			return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagBitString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagBitString")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetContextTagBitString) isBACnetContextTagBitString() bool {
	return true
}

func (m *_BACnetContextTagBitString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
