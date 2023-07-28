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


// BACnetContextTagNull is the corresponding interface of BACnetContextTagNull
type BACnetContextTagNull interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
}

// BACnetContextTagNullExactly can be used when we want exactly this type and not a type which fulfills BACnetContextTagNull.
// This is useful for switch cases.
type BACnetContextTagNullExactly interface {
	BACnetContextTagNull
	isBACnetContextTagNull() bool
}

// _BACnetContextTagNull is the data-structure of this message
type _BACnetContextTagNull struct {
	*_BACnetContextTag
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagNull)  GetDataType() BACnetDataType {
return BACnetDataType_NULL}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagNull) InitializeParent(parent BACnetContextTag , header BACnetTagHeader ) {	m.Header = header
}

func (m *_BACnetContextTagNull)  GetParent() BACnetContextTag {
	return m._BACnetContextTag
}


// NewBACnetContextTagNull factory function for _BACnetContextTagNull
func NewBACnetContextTagNull( header BACnetTagHeader , tagNumberArgument uint8 ) *_BACnetContextTagNull {
	_result := &_BACnetContextTagNull{
    	_BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	_result._BACnetContextTag._BACnetContextTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagNull(structType any) BACnetContextTagNull {
    if casted, ok := structType.(BACnetContextTagNull); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagNull); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagNull) GetTypeName() string {
	return "BACnetContextTagNull"
}

func (m *_BACnetContextTagNull) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}


func (m *_BACnetContextTagNull) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetContextTagNullParse(ctx context.Context, theBytes []byte, header BACnetTagHeader, tagNumberArgument uint8, dataType BACnetDataType) (BACnetContextTagNull, error) {
	return BACnetContextTagNullParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), header, tagNumberArgument, dataType)
}

func BACnetContextTagNullParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, header BACnetTagHeader, tagNumberArgument uint8, dataType BACnetDataType) (BACnetContextTagNull, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetContextTagNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if (!(bool((header.GetActualLength()) == ((0))))) {
		return nil, errors.WithStack(utils.ParseValidationError{"length field should be 0"})
	}

	if closeErr := readBuffer.CloseContext("BACnetContextTagNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagNull")
	}

	// Create a partially initialized instance
	_child := &_BACnetContextTagNull{
		_BACnetContextTag: &_BACnetContextTag{
			TagNumberArgument: tagNumberArgument,
		},
	}
	_child._BACnetContextTag._BACnetContextTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetContextTagNull) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagNull) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagNull")
		}

		if popErr := writeBuffer.PopContext("BACnetContextTagNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagNull")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetContextTagNull) isBACnetContextTagNull() bool {
	return true
}

func (m *_BACnetContextTagNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



