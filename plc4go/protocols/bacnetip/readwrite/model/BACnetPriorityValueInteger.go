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


// BACnetPriorityValueInteger is the corresponding interface of BACnetPriorityValueInteger
type BACnetPriorityValueInteger interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPriorityValue
	// GetIntegerValue returns IntegerValue (property field)
	GetIntegerValue() BACnetApplicationTagSignedInteger
}

// BACnetPriorityValueIntegerExactly can be used when we want exactly this type and not a type which fulfills BACnetPriorityValueInteger.
// This is useful for switch cases.
type BACnetPriorityValueIntegerExactly interface {
	BACnetPriorityValueInteger
	isBACnetPriorityValueInteger() bool
}

// _BACnetPriorityValueInteger is the data-structure of this message
type _BACnetPriorityValueInteger struct {
	*_BACnetPriorityValue
        IntegerValue BACnetApplicationTagSignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPriorityValueInteger) InitializeParent(parent BACnetPriorityValue , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPriorityValueInteger)  GetParent() BACnetPriorityValue {
	return m._BACnetPriorityValue
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueInteger) GetIntegerValue() BACnetApplicationTagSignedInteger {
	return m.IntegerValue
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetPriorityValueInteger factory function for _BACnetPriorityValueInteger
func NewBACnetPriorityValueInteger( integerValue BACnetApplicationTagSignedInteger , peekedTagHeader BACnetTagHeader , objectTypeArgument BACnetObjectType ) *_BACnetPriorityValueInteger {
	_result := &_BACnetPriorityValueInteger{
		IntegerValue: integerValue,
    	_BACnetPriorityValue: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueInteger(structType any) BACnetPriorityValueInteger {
    if casted, ok := structType.(BACnetPriorityValueInteger); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueInteger); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueInteger) GetTypeName() string {
	return "BACnetPriorityValueInteger"
}

func (m *_BACnetPriorityValueInteger) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (integerValue)
	lengthInBits += m.IntegerValue.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetPriorityValueInteger) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPriorityValueIntegerParse(ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType) (BACnetPriorityValueInteger, error) {
	return BACnetPriorityValueIntegerParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func BACnetPriorityValueIntegerParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetPriorityValueInteger, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetPriorityValueInteger"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueInteger")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (integerValue)
	if pullErr := readBuffer.PullContext("integerValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for integerValue")
	}
_integerValue, _integerValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _integerValueErr != nil {
		return nil, errors.Wrap(_integerValueErr, "Error parsing 'integerValue' field of BACnetPriorityValueInteger")
	}
	integerValue := _integerValue.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("integerValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for integerValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueInteger"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueInteger")
	}

	// Create a partially initialized instance
	_child := &_BACnetPriorityValueInteger{
		_BACnetPriorityValue: &_BACnetPriorityValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		IntegerValue: integerValue,
	}
	_child._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPriorityValueInteger) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueInteger) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueInteger"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueInteger")
		}

	// Simple Field (integerValue)
	if pushErr := writeBuffer.PushContext("integerValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for integerValue")
	}
	_integerValueErr := writeBuffer.WriteSerializable(ctx, m.GetIntegerValue())
	if popErr := writeBuffer.PopContext("integerValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for integerValue")
	}
	if _integerValueErr != nil {
		return errors.Wrap(_integerValueErr, "Error serializing 'integerValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueInteger"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueInteger")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetPriorityValueInteger) isBACnetPriorityValueInteger() bool {
	return true
}

func (m *_BACnetPriorityValueInteger) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



