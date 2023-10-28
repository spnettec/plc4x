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

// BACnetValueSourceNone is the corresponding interface of BACnetValueSourceNone
type BACnetValueSourceNone interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetValueSource
	// GetNone returns None (property field)
	GetNone() BACnetContextTagNull
}

// BACnetValueSourceNoneExactly can be used when we want exactly this type and not a type which fulfills BACnetValueSourceNone.
// This is useful for switch cases.
type BACnetValueSourceNoneExactly interface {
	BACnetValueSourceNone
	isBACnetValueSourceNone() bool
}

// _BACnetValueSourceNone is the data-structure of this message
type _BACnetValueSourceNone struct {
	*_BACnetValueSource
	None BACnetContextTagNull
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetValueSourceNone) InitializeParent(parent BACnetValueSource, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetValueSourceNone) GetParent() BACnetValueSource {
	return m._BACnetValueSource
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetValueSourceNone) GetNone() BACnetContextTagNull {
	return m.None
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetValueSourceNone factory function for _BACnetValueSourceNone
func NewBACnetValueSourceNone(none BACnetContextTagNull, peekedTagHeader BACnetTagHeader) *_BACnetValueSourceNone {
	_result := &_BACnetValueSourceNone{
		None:               none,
		_BACnetValueSource: NewBACnetValueSource(peekedTagHeader),
	}
	_result._BACnetValueSource._BACnetValueSourceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetValueSourceNone(structType any) BACnetValueSourceNone {
	if casted, ok := structType.(BACnetValueSourceNone); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetValueSourceNone); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetValueSourceNone) GetTypeName() string {
	return "BACnetValueSourceNone"
}

func (m *_BACnetValueSourceNone) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (none)
	lengthInBits += m.None.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetValueSourceNone) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetValueSourceNoneParse(ctx context.Context, theBytes []byte) (BACnetValueSourceNone, error) {
	return BACnetValueSourceNoneParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetValueSourceNoneParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetValueSourceNone, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetValueSourceNone"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetValueSourceNone")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (none)
	if pullErr := readBuffer.PullContext("none"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for none")
	}
	_none, _noneErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_NULL))
	if _noneErr != nil {
		return nil, errors.Wrap(_noneErr, "Error parsing 'none' field of BACnetValueSourceNone")
	}
	none := _none.(BACnetContextTagNull)
	if closeErr := readBuffer.CloseContext("none"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for none")
	}

	if closeErr := readBuffer.CloseContext("BACnetValueSourceNone"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetValueSourceNone")
	}

	// Create a partially initialized instance
	_child := &_BACnetValueSourceNone{
		_BACnetValueSource: &_BACnetValueSource{},
		None:               none,
	}
	_child._BACnetValueSource._BACnetValueSourceChildRequirements = _child
	return _child, nil
}

func (m *_BACnetValueSourceNone) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetValueSourceNone) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetValueSourceNone"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetValueSourceNone")
		}

		// Simple Field (none)
		if pushErr := writeBuffer.PushContext("none"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for none")
		}
		_noneErr := writeBuffer.WriteSerializable(ctx, m.GetNone())
		if popErr := writeBuffer.PopContext("none"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for none")
		}
		if _noneErr != nil {
			return errors.Wrap(_noneErr, "Error serializing 'none' field")
		}

		if popErr := writeBuffer.PopContext("BACnetValueSourceNone"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetValueSourceNone")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetValueSourceNone) isBACnetValueSourceNone() bool {
	return true
}

func (m *_BACnetValueSourceNone) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
