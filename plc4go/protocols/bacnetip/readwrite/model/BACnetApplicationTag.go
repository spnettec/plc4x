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


// BACnetApplicationTag is the corresponding interface of BACnetApplicationTag
type BACnetApplicationTag interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetHeader returns Header (property field)
	GetHeader() BACnetTagHeader
	// GetActualTagNumber returns ActualTagNumber (virtual field)
	GetActualTagNumber() uint8
	// GetActualLength returns ActualLength (virtual field)
	GetActualLength() uint32
}

// BACnetApplicationTagExactly can be used when we want exactly this type and not a type which fulfills BACnetApplicationTag.
// This is useful for switch cases.
type BACnetApplicationTagExactly interface {
	BACnetApplicationTag
	isBACnetApplicationTag() bool
}

// _BACnetApplicationTag is the data-structure of this message
type _BACnetApplicationTag struct {
	_BACnetApplicationTagChildRequirements
        Header BACnetTagHeader
}

type _BACnetApplicationTagChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetActualTagNumber() uint8
}


type BACnetApplicationTagParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetApplicationTag, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetApplicationTagChild interface {
	utils.Serializable
InitializeParent(parent BACnetApplicationTag , header BACnetTagHeader )
	GetParent() *BACnetApplicationTag

	GetTypeName() string
	BACnetApplicationTag
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTag) GetHeader() BACnetTagHeader {
	return m.Header
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetApplicationTag) GetActualTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetHeader().GetActualTagNumber())
}

func (m *_BACnetApplicationTag) GetActualLength() uint32 {
	ctx := context.Background()
	_ = ctx
	return uint32(m.GetHeader().GetActualLength())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetApplicationTag factory function for _BACnetApplicationTag
func NewBACnetApplicationTag( header BACnetTagHeader ) *_BACnetApplicationTag {
return &_BACnetApplicationTag{ Header: header }
}

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTag(structType any) BACnetApplicationTag {
    if casted, ok := structType.(BACnetApplicationTag); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTag); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTag) GetTypeName() string {
	return "BACnetApplicationTag"
}


func (m *_BACnetApplicationTag) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (header)
	lengthInBits += m.Header.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetApplicationTag) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetApplicationTagParse(ctx context.Context, theBytes []byte) (BACnetApplicationTag, error) {
	return BACnetApplicationTagParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetApplicationTagParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetApplicationTag, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetApplicationTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTag")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (header)
	if pullErr := readBuffer.PullContext("header"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for header")
	}
_header, _headerErr := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
	if _headerErr != nil {
		return nil, errors.Wrap(_headerErr, "Error parsing 'header' field of BACnetApplicationTag")
	}
	header := _header.(BACnetTagHeader)
	if closeErr := readBuffer.CloseContext("header"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for header")
	}

	// Validation
	if (!(bool((header.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) {
		return nil, errors.WithStack(utils.ParseValidationError{"should be a application tag"})
	}

	// Virtual field
	_actualTagNumber := header.GetActualTagNumber()
	actualTagNumber := uint8(_actualTagNumber)
	_ = actualTagNumber

	// Virtual field
	_actualLength := header.GetActualLength()
	actualLength := uint32(_actualLength)
	_ = actualLength

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetApplicationTagChildSerializeRequirement interface {
		BACnetApplicationTag
		InitializeParent(BACnetApplicationTag,  BACnetTagHeader)
		GetParent() BACnetApplicationTag
	}
	var _childTemp any
	var _child BACnetApplicationTagChildSerializeRequirement
	var typeSwitchError error
	switch {
case actualTagNumber == 0x0 : // BACnetApplicationTagNull
		_childTemp, typeSwitchError = BACnetApplicationTagNullParseWithBuffer(ctx, readBuffer, )
case actualTagNumber == 0x1 : // BACnetApplicationTagBoolean
		_childTemp, typeSwitchError = BACnetApplicationTagBooleanParseWithBuffer(ctx, readBuffer, header)
case actualTagNumber == 0x2 : // BACnetApplicationTagUnsignedInteger
		_childTemp, typeSwitchError = BACnetApplicationTagUnsignedIntegerParseWithBuffer(ctx, readBuffer, header)
case actualTagNumber == 0x3 : // BACnetApplicationTagSignedInteger
		_childTemp, typeSwitchError = BACnetApplicationTagSignedIntegerParseWithBuffer(ctx, readBuffer, header)
case actualTagNumber == 0x4 : // BACnetApplicationTagReal
		_childTemp, typeSwitchError = BACnetApplicationTagRealParseWithBuffer(ctx, readBuffer, )
case actualTagNumber == 0x5 : // BACnetApplicationTagDouble
		_childTemp, typeSwitchError = BACnetApplicationTagDoubleParseWithBuffer(ctx, readBuffer, )
case actualTagNumber == 0x6 : // BACnetApplicationTagOctetString
		_childTemp, typeSwitchError = BACnetApplicationTagOctetStringParseWithBuffer(ctx, readBuffer, header)
case actualTagNumber == 0x7 : // BACnetApplicationTagCharacterString
		_childTemp, typeSwitchError = BACnetApplicationTagCharacterStringParseWithBuffer(ctx, readBuffer, header)
case actualTagNumber == 0x8 : // BACnetApplicationTagBitString
		_childTemp, typeSwitchError = BACnetApplicationTagBitStringParseWithBuffer(ctx, readBuffer, header)
case actualTagNumber == 0x9 : // BACnetApplicationTagEnumerated
		_childTemp, typeSwitchError = BACnetApplicationTagEnumeratedParseWithBuffer(ctx, readBuffer, header)
case actualTagNumber == 0xA : // BACnetApplicationTagDate
		_childTemp, typeSwitchError = BACnetApplicationTagDateParseWithBuffer(ctx, readBuffer, )
case actualTagNumber == 0xB : // BACnetApplicationTagTime
		_childTemp, typeSwitchError = BACnetApplicationTagTimeParseWithBuffer(ctx, readBuffer, )
case actualTagNumber == 0xC : // BACnetApplicationTagObjectIdentifier
		_childTemp, typeSwitchError = BACnetApplicationTagObjectIdentifierParseWithBuffer(ctx, readBuffer, )
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [actualTagNumber=%v]", actualTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetApplicationTag")
	}
	_child = _childTemp.(BACnetApplicationTagChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetApplicationTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTag")
	}

	// Finish initializing
_child.InitializeParent(_child , header )
	return _child, nil
}

func (pm *_BACnetApplicationTag) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetApplicationTag, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("BACnetApplicationTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTag")
	}

	// Simple Field (header)
	if pushErr := writeBuffer.PushContext("header"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for header")
	}
	_headerErr := writeBuffer.WriteSerializable(ctx, m.GetHeader())
	if popErr := writeBuffer.PopContext("header"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for header")
	}
	if _headerErr != nil {
		return errors.Wrap(_headerErr, "Error serializing 'header' field")
	}
	// Virtual field
	actualTagNumber := m.GetActualTagNumber()
	_ = actualTagNumber
	if _actualTagNumberErr := writeBuffer.WriteVirtual(ctx, "actualTagNumber", m.GetActualTagNumber()); _actualTagNumberErr != nil {
		return errors.Wrap(_actualTagNumberErr, "Error serializing 'actualTagNumber' field")
	}
	// Virtual field
	actualLength := m.GetActualLength()
	_ = actualLength
	if _actualLengthErr := writeBuffer.WriteVirtual(ctx, "actualLength", m.GetActualLength()); _actualLengthErr != nil {
		return errors.Wrap(_actualLengthErr, "Error serializing 'actualLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetApplicationTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetApplicationTag")
	}
	return nil
}


func (m *_BACnetApplicationTag) isBACnetApplicationTag() bool {
	return true
}

func (m *_BACnetApplicationTag) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



