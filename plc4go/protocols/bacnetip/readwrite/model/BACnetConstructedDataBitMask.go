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


// BACnetConstructedDataBitMask is the corresponding interface of BACnetConstructedDataBitMask
type BACnetConstructedDataBitMask interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBitString returns BitString (property field)
	GetBitString() BACnetApplicationTagBitString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBitString
}

// BACnetConstructedDataBitMaskExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBitMask.
// This is useful for switch cases.
type BACnetConstructedDataBitMaskExactly interface {
	BACnetConstructedDataBitMask
	isBACnetConstructedDataBitMask() bool
}

// _BACnetConstructedDataBitMask is the data-structure of this message
type _BACnetConstructedDataBitMask struct {
	*_BACnetConstructedData
        BitString BACnetApplicationTagBitString
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBitMask)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataBitMask)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_BIT_MASK}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBitMask) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBitMask)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBitMask) GetBitString() BACnetApplicationTagBitString {
	return m.BitString
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBitMask) GetActualValue() BACnetApplicationTagBitString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBitString(m.GetBitString())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataBitMask factory function for _BACnetConstructedDataBitMask
func NewBACnetConstructedDataBitMask( bitString BACnetApplicationTagBitString , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataBitMask {
	_result := &_BACnetConstructedDataBitMask{
		BitString: bitString,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBitMask(structType any) BACnetConstructedDataBitMask {
    if casted, ok := structType.(BACnetConstructedDataBitMask); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBitMask); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBitMask) GetTypeName() string {
	return "BACnetConstructedDataBitMask"
}

func (m *_BACnetConstructedDataBitMask) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (bitString)
	lengthInBits += m.BitString.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataBitMask) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBitMaskParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBitMask, error) {
	return BACnetConstructedDataBitMaskParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBitMaskParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBitMask, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBitMask"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBitMask")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bitString)
	if pullErr := readBuffer.PullContext("bitString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bitString")
	}
_bitString, _bitStringErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _bitStringErr != nil {
		return nil, errors.Wrap(_bitStringErr, "Error parsing 'bitString' field of BACnetConstructedDataBitMask")
	}
	bitString := _bitString.(BACnetApplicationTagBitString)
	if closeErr := readBuffer.CloseContext("bitString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bitString")
	}

	// Virtual field
	_actualValue := bitString
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBitMask"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBitMask")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBitMask{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		BitString: bitString,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBitMask) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBitMask) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBitMask"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBitMask")
		}

	// Simple Field (bitString)
	if pushErr := writeBuffer.PushContext("bitString"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for bitString")
	}
	_bitStringErr := writeBuffer.WriteSerializable(ctx, m.GetBitString())
	if popErr := writeBuffer.PopContext("bitString"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for bitString")
	}
	if _bitStringErr != nil {
		return errors.Wrap(_bitStringErr, "Error serializing 'bitString' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBitMask"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBitMask")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataBitMask) isBACnetConstructedDataBitMask() bool {
	return true
}

func (m *_BACnetConstructedDataBitMask) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



