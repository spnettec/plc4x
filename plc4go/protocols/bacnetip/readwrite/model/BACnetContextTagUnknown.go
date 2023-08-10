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

// BACnetContextTagUnknown is the corresponding interface of BACnetContextTagUnknown
type BACnetContextTagUnknown interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
	// GetUnknownData returns UnknownData (property field)
	GetUnknownData() []byte
}

// BACnetContextTagUnknownExactly can be used when we want exactly this type and not a type which fulfills BACnetContextTagUnknown.
// This is useful for switch cases.
type BACnetContextTagUnknownExactly interface {
	BACnetContextTagUnknown
	isBACnetContextTagUnknown() bool
}

// _BACnetContextTagUnknown is the data-structure of this message
type _BACnetContextTagUnknown struct {
	*_BACnetContextTag
	UnknownData []byte

	// Arguments.
	ActualLength uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagUnknown) GetDataType() BACnetDataType {
	return BACnetDataType_UNKNOWN
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagUnknown) InitializeParent(parent BACnetContextTag, header BACnetTagHeader) {
	m.Header = header
}

func (m *_BACnetContextTagUnknown) GetParent() BACnetContextTag {
	return m._BACnetContextTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagUnknown) GetUnknownData() []byte {
	return m.UnknownData
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetContextTagUnknown factory function for _BACnetContextTagUnknown
func NewBACnetContextTagUnknown(unknownData []byte, header BACnetTagHeader, actualLength uint32, tagNumberArgument uint8) *_BACnetContextTagUnknown {
	_result := &_BACnetContextTagUnknown{
		UnknownData:       unknownData,
		_BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	_result._BACnetContextTag._BACnetContextTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagUnknown(structType any) BACnetContextTagUnknown {
	if casted, ok := structType.(BACnetContextTagUnknown); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagUnknown); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagUnknown) GetTypeName() string {
	return "BACnetContextTagUnknown"
}

func (m *_BACnetContextTagUnknown) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.UnknownData) > 0 {
		lengthInBits += 8 * uint16(len(m.UnknownData))
	}

	return lengthInBits
}

func (m *_BACnetContextTagUnknown) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetContextTagUnknownParse(ctx context.Context, theBytes []byte, actualLength uint32, tagNumberArgument uint8, dataType BACnetDataType) (BACnetContextTagUnknown, error) {
	return BACnetContextTagUnknownParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), actualLength, tagNumberArgument, dataType)
}

func BACnetContextTagUnknownParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, actualLength uint32, tagNumberArgument uint8, dataType BACnetDataType) (BACnetContextTagUnknown, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetContextTagUnknown"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagUnknown")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (unknownData)
	numberOfBytesunknownData := int(actualLength)
	unknownData, _readArrayErr := readBuffer.ReadByteArray("unknownData", numberOfBytesunknownData)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'unknownData' field of BACnetContextTagUnknown")
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTagUnknown"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagUnknown")
	}

	// Create a partially initialized instance
	_child := &_BACnetContextTagUnknown{
		_BACnetContextTag: &_BACnetContextTag{
			TagNumberArgument: tagNumberArgument,
		},
		UnknownData: unknownData,
	}
	_child._BACnetContextTag._BACnetContextTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetContextTagUnknown) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagUnknown) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagUnknown"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagUnknown")
		}

		// Array Field (unknownData)
		// Byte Array field (unknownData)
		if err := writeBuffer.WriteByteArray("unknownData", m.GetUnknownData()); err != nil {
			return errors.Wrap(err, "Error serializing 'unknownData' field")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagUnknown"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagUnknown")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

////
// Arguments Getter

func (m *_BACnetContextTagUnknown) GetActualLength() uint32 {
	return m.ActualLength
}

//
////

func (m *_BACnetContextTagUnknown) isBACnetContextTagUnknown() bool {
	return true
}

func (m *_BACnetContextTagUnknown) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
