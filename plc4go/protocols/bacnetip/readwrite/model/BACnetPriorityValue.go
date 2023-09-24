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


// BACnetPriorityValue is the corresponding interface of BACnetPriorityValue
type BACnetPriorityValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetPeekedIsContextTag returns PeekedIsContextTag (virtual field)
	GetPeekedIsContextTag() bool
}

// BACnetPriorityValueExactly can be used when we want exactly this type and not a type which fulfills BACnetPriorityValue.
// This is useful for switch cases.
type BACnetPriorityValueExactly interface {
	BACnetPriorityValue
	isBACnetPriorityValue() bool
}

// _BACnetPriorityValue is the data-structure of this message
type _BACnetPriorityValue struct {
	_BACnetPriorityValueChildRequirements
        PeekedTagHeader BACnetTagHeader

	// Arguments.
	ObjectTypeArgument BACnetObjectType
}

type _BACnetPriorityValueChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
	GetPeekedIsContextTag() bool
}


type BACnetPriorityValueParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetPriorityValue, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetPriorityValueChild interface {
	utils.Serializable
InitializeParent(parent BACnetPriorityValue , peekedTagHeader BACnetTagHeader )
	GetParent() *BACnetPriorityValue

	GetTypeName() string
	BACnetPriorityValue
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValue) GetPeekedTagHeader() BACnetTagHeader {
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

func (m *_BACnetPriorityValue) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (m *_BACnetPriorityValue) GetPeekedIsContextTag() bool {
	ctx := context.Background()
	_ = ctx
	return bool(bool((m.GetPeekedTagHeader().GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS)))
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetPriorityValue factory function for _BACnetPriorityValue
func NewBACnetPriorityValue( peekedTagHeader BACnetTagHeader , objectTypeArgument BACnetObjectType ) *_BACnetPriorityValue {
return &_BACnetPriorityValue{ PeekedTagHeader: peekedTagHeader , ObjectTypeArgument: objectTypeArgument }
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValue(structType any) BACnetPriorityValue {
    if casted, ok := structType.(BACnetPriorityValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValue) GetTypeName() string {
	return "BACnetPriorityValue"
}


func (m *_BACnetPriorityValue) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetPriorityValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPriorityValueParse(ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType) (BACnetPriorityValue, error) {
	return BACnetPriorityValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument)
}

func BACnetPriorityValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetPriorityValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetPriorityValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValue")
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

	// Virtual field
	_peekedIsContextTag := bool((peekedTagHeader.GetTagClass()) == (TagClass_CONTEXT_SPECIFIC_TAGS))
	peekedIsContextTag := bool(_peekedIsContextTag)
	_ = peekedIsContextTag

	// Validation
	if (!(bool((!(peekedIsContextTag))) || bool((bool(bool(peekedIsContextTag) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x6)))) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x7))))))) {
		return nil, errors.WithStack(utils.ParseValidationError{"unexpected opening or closing tag"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetPriorityValueChildSerializeRequirement interface {
		BACnetPriorityValue
		InitializeParent(BACnetPriorityValue,  BACnetTagHeader)
		GetParent() BACnetPriorityValue
	}
	var _childTemp any
	var _child BACnetPriorityValueChildSerializeRequirement
	var typeSwitchError error
	switch {
case peekedTagNumber == 0x0 && peekedIsContextTag == bool(false) : // BACnetPriorityValueNull
		_childTemp, typeSwitchError = BACnetPriorityValueNullParseWithBuffer(ctx, readBuffer, objectTypeArgument)
case peekedTagNumber == 0x4 && peekedIsContextTag == bool(false) : // BACnetPriorityValueReal
		_childTemp, typeSwitchError = BACnetPriorityValueRealParseWithBuffer(ctx, readBuffer, objectTypeArgument)
case peekedTagNumber == 0x9 && peekedIsContextTag == bool(false) : // BACnetPriorityValueEnumerated
		_childTemp, typeSwitchError = BACnetPriorityValueEnumeratedParseWithBuffer(ctx, readBuffer, objectTypeArgument)
case peekedTagNumber == 0x2 && peekedIsContextTag == bool(false) : // BACnetPriorityValueUnsigned
		_childTemp, typeSwitchError = BACnetPriorityValueUnsignedParseWithBuffer(ctx, readBuffer, objectTypeArgument)
case peekedTagNumber == 0x1 && peekedIsContextTag == bool(false) : // BACnetPriorityValueBoolean
		_childTemp, typeSwitchError = BACnetPriorityValueBooleanParseWithBuffer(ctx, readBuffer, objectTypeArgument)
case peekedTagNumber == 0x3 && peekedIsContextTag == bool(false) : // BACnetPriorityValueInteger
		_childTemp, typeSwitchError = BACnetPriorityValueIntegerParseWithBuffer(ctx, readBuffer, objectTypeArgument)
case peekedTagNumber == 0x5 && peekedIsContextTag == bool(false) : // BACnetPriorityValueDouble
		_childTemp, typeSwitchError = BACnetPriorityValueDoubleParseWithBuffer(ctx, readBuffer, objectTypeArgument)
case peekedTagNumber == 0xB && peekedIsContextTag == bool(false) : // BACnetPriorityValueTime
		_childTemp, typeSwitchError = BACnetPriorityValueTimeParseWithBuffer(ctx, readBuffer, objectTypeArgument)
case peekedTagNumber == 0x7 && peekedIsContextTag == bool(false) : // BACnetPriorityValueCharacterString
		_childTemp, typeSwitchError = BACnetPriorityValueCharacterStringParseWithBuffer(ctx, readBuffer, objectTypeArgument)
case peekedTagNumber == 0x6 && peekedIsContextTag == bool(false) : // BACnetPriorityValueOctetString
		_childTemp, typeSwitchError = BACnetPriorityValueOctetStringParseWithBuffer(ctx, readBuffer, objectTypeArgument)
case peekedTagNumber == 0x8 && peekedIsContextTag == bool(false) : // BACnetPriorityValueBitString
		_childTemp, typeSwitchError = BACnetPriorityValueBitStringParseWithBuffer(ctx, readBuffer, objectTypeArgument)
case peekedTagNumber == 0xA && peekedIsContextTag == bool(false) : // BACnetPriorityValueDate
		_childTemp, typeSwitchError = BACnetPriorityValueDateParseWithBuffer(ctx, readBuffer, objectTypeArgument)
case peekedTagNumber == 0xC && peekedIsContextTag == bool(false) : // BACnetPriorityValueObjectidentifier
		_childTemp, typeSwitchError = BACnetPriorityValueObjectidentifierParseWithBuffer(ctx, readBuffer, objectTypeArgument)
case peekedTagNumber == uint8(0) && peekedIsContextTag == bool(true) : // BACnetPriorityValueConstructedValue
		_childTemp, typeSwitchError = BACnetPriorityValueConstructedValueParseWithBuffer(ctx, readBuffer, objectTypeArgument)
case peekedTagNumber == uint8(1) && peekedIsContextTag == bool(true) : // BACnetPriorityValueDateTime
		_childTemp, typeSwitchError = BACnetPriorityValueDateTimeParseWithBuffer(ctx, readBuffer, objectTypeArgument)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v, peekedIsContextTag=%v]", peekedTagNumber, peekedIsContextTag)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetPriorityValue")
	}
	_child = _childTemp.(BACnetPriorityValueChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetPriorityValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValue")
	}

	// Finish initializing
_child.InitializeParent(_child , peekedTagHeader )
	return _child, nil
}

func (pm *_BACnetPriorityValue) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetPriorityValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("BACnetPriorityValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValue")
	}
	// Virtual field
	peekedTagNumber := m.GetPeekedTagNumber()
	_ =	peekedTagNumber
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}
	// Virtual field
	peekedIsContextTag := m.GetPeekedIsContextTag()
	_ =	peekedIsContextTag
	if _peekedIsContextTagErr := writeBuffer.WriteVirtual(ctx, "peekedIsContextTag", m.GetPeekedIsContextTag()); _peekedIsContextTagErr != nil {
		return errors.Wrap(_peekedIsContextTagErr, "Error serializing 'peekedIsContextTag' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetPriorityValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPriorityValue")
	}
	return nil
}


////
// Arguments Getter

func (m *_BACnetPriorityValue) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}
//
////

func (m *_BACnetPriorityValue) isBACnetPriorityValue() bool {
	return true
}

func (m *_BACnetPriorityValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



