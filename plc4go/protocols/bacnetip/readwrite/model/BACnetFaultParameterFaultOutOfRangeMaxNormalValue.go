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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetFaultParameterFaultOutOfRangeMaxNormalValue is the corresponding interface of BACnetFaultParameterFaultOutOfRangeMaxNormalValue
type BACnetFaultParameterFaultOutOfRangeMaxNormalValue interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetFaultParameterFaultOutOfRangeMaxNormalValueExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultOutOfRangeMaxNormalValue.
// This is useful for switch cases.
type BACnetFaultParameterFaultOutOfRangeMaxNormalValueExactly interface {
	BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	isBACnetFaultParameterFaultOutOfRangeMaxNormalValue() bool
}

// _BACnetFaultParameterFaultOutOfRangeMaxNormalValue is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRangeMaxNormalValue struct {
	_BACnetFaultParameterFaultOutOfRangeMaxNormalValueChildRequirements
        OpeningTag BACnetOpeningTag
        PeekedTagHeader BACnetTagHeader
        ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

type _BACnetFaultParameterFaultOutOfRangeMaxNormalValueChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}


type BACnetFaultParameterFaultOutOfRangeMaxNormalValueParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetFaultParameterFaultOutOfRangeMaxNormalValue, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetFaultParameterFaultOutOfRangeMaxNormalValueChild interface {
	utils.Serializable
InitializeParent(parent BACnetFaultParameterFaultOutOfRangeMaxNormalValue , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag )
	GetParent() *BACnetFaultParameterFaultOutOfRangeMaxNormalValue

	GetTypeName() string
	BACnetFaultParameterFaultOutOfRangeMaxNormalValue
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetFaultParameterFaultOutOfRangeMaxNormalValue factory function for _BACnetFaultParameterFaultOutOfRangeMaxNormalValue
func NewBACnetFaultParameterFaultOutOfRangeMaxNormalValue( openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 ) *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue {
return &_BACnetFaultParameterFaultOutOfRangeMaxNormalValue{ OpeningTag: openingTag , PeekedTagHeader: peekedTagHeader , ClosingTag: closingTag , TagNumber: tagNumber }
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRangeMaxNormalValue(structType interface{}) BACnetFaultParameterFaultOutOfRangeMaxNormalValue {
    if casted, ok := structType.(BACnetFaultParameterFaultOutOfRangeMaxNormalValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRangeMaxNormalValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRangeMaxNormalValue"
}


func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultOutOfRangeMaxNormalValueParse(theBytes []byte, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMaxNormalValue, error) {
	return BACnetFaultParameterFaultOutOfRangeMaxNormalValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetFaultParameterFaultOutOfRangeMaxNormalValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetFaultParameterFaultOutOfRangeMaxNormalValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRangeMaxNormalValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer , uint8( tagNumber ) )
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetFaultParameterFaultOutOfRangeMaxNormalValue")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

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

	// Validation
	if (!(bool((peekedTagHeader.GetTagClass()) == (TagClass_APPLICATION_TAGS)))) {
		return nil, errors.WithStack(utils.ParseValidationError{"only application tags allowed"})
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetFaultParameterFaultOutOfRangeMaxNormalValueChildSerializeRequirement interface {
		BACnetFaultParameterFaultOutOfRangeMaxNormalValue
		InitializeParent(BACnetFaultParameterFaultOutOfRangeMaxNormalValue,  BACnetOpeningTag, BACnetTagHeader, BACnetClosingTag)
		GetParent() BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	}
	var _childTemp interface{}
	var _child BACnetFaultParameterFaultOutOfRangeMaxNormalValueChildSerializeRequirement
	var typeSwitchError error
	switch {
case peekedTagNumber == 0x4 : // BACnetFaultParameterFaultOutOfRangeMaxNormalValueReal
		_childTemp, typeSwitchError = BACnetFaultParameterFaultOutOfRangeMaxNormalValueRealParseWithBuffer(ctx, readBuffer, tagNumber)
case peekedTagNumber == 0x2 : // BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsigned
		_childTemp, typeSwitchError = BACnetFaultParameterFaultOutOfRangeMaxNormalValueUnsignedParseWithBuffer(ctx, readBuffer, tagNumber)
case peekedTagNumber == 0x5 : // BACnetFaultParameterFaultOutOfRangeMaxNormalValueDouble
		_childTemp, typeSwitchError = BACnetFaultParameterFaultOutOfRangeMaxNormalValueDoubleParseWithBuffer(ctx, readBuffer, tagNumber)
case peekedTagNumber == 0x3 : // BACnetFaultParameterFaultOutOfRangeMaxNormalValueInteger
		_childTemp, typeSwitchError = BACnetFaultParameterFaultOutOfRangeMaxNormalValueIntegerParseWithBuffer(ctx, readBuffer, tagNumber)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetFaultParameterFaultOutOfRangeMaxNormalValue")
	}
	_child = _childTemp.(BACnetFaultParameterFaultOutOfRangeMaxNormalValueChildSerializeRequirement)

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer , uint8( tagNumber ) )
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetFaultParameterFaultOutOfRangeMaxNormalValue")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRangeMaxNormalValue")
	}

	// Finish initializing
_child.InitializeParent(_child , openingTag , peekedTagHeader , closingTag )
	return _child, nil
}

func (pm *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetFaultParameterFaultOutOfRangeMaxNormalValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRangeMaxNormalValue")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRangeMaxNormalValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRangeMaxNormalValue")
	}
	return nil
}


////
// Arguments Getter

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue) GetTagNumber() uint8 {
	return m.TagNumber
}
//
////

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue) isBACnetFaultParameterFaultOutOfRangeMaxNormalValue() bool {
	return true
}

func (m *_BACnetFaultParameterFaultOutOfRangeMaxNormalValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



