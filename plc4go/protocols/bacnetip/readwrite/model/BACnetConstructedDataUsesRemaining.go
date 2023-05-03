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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataUsesRemaining is the corresponding interface of BACnetConstructedDataUsesRemaining
type BACnetConstructedDataUsesRemaining interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetUsesRemaining returns UsesRemaining (property field)
	GetUsesRemaining() BACnetApplicationTagSignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagSignedInteger
}

// BACnetConstructedDataUsesRemainingExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataUsesRemaining.
// This is useful for switch cases.
type BACnetConstructedDataUsesRemainingExactly interface {
	BACnetConstructedDataUsesRemaining
	isBACnetConstructedDataUsesRemaining() bool
}

// _BACnetConstructedDataUsesRemaining is the data-structure of this message
type _BACnetConstructedDataUsesRemaining struct {
	*_BACnetConstructedData
	UsesRemaining BACnetApplicationTagSignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUsesRemaining) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataUsesRemaining) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_USES_REMAINING
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUsesRemaining) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataUsesRemaining) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUsesRemaining) GetUsesRemaining() BACnetApplicationTagSignedInteger {
	return m.UsesRemaining
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataUsesRemaining) GetActualValue() BACnetApplicationTagSignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagSignedInteger(m.GetUsesRemaining())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataUsesRemaining factory function for _BACnetConstructedDataUsesRemaining
func NewBACnetConstructedDataUsesRemaining(usesRemaining BACnetApplicationTagSignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataUsesRemaining {
	_result := &_BACnetConstructedDataUsesRemaining{
		UsesRemaining:          usesRemaining,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUsesRemaining(structType any) BACnetConstructedDataUsesRemaining {
	if casted, ok := structType.(BACnetConstructedDataUsesRemaining); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUsesRemaining); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUsesRemaining) GetTypeName() string {
	return "BACnetConstructedDataUsesRemaining"
}

func (m *_BACnetConstructedDataUsesRemaining) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (usesRemaining)
	lengthInBits += m.UsesRemaining.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataUsesRemaining) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataUsesRemainingParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataUsesRemaining, error) {
	return BACnetConstructedDataUsesRemainingParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataUsesRemainingParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataUsesRemaining, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUsesRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUsesRemaining")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (usesRemaining)
	if pullErr := readBuffer.PullContext("usesRemaining"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for usesRemaining")
	}
	_usesRemaining, _usesRemainingErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _usesRemainingErr != nil {
		return nil, errors.Wrap(_usesRemainingErr, "Error parsing 'usesRemaining' field of BACnetConstructedDataUsesRemaining")
	}
	usesRemaining := _usesRemaining.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("usesRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for usesRemaining")
	}

	// Virtual field
	_actualValue := usesRemaining
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUsesRemaining"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUsesRemaining")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataUsesRemaining{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		UsesRemaining: usesRemaining,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataUsesRemaining) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataUsesRemaining) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUsesRemaining"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUsesRemaining")
		}

		// Simple Field (usesRemaining)
		if pushErr := writeBuffer.PushContext("usesRemaining"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for usesRemaining")
		}
		_usesRemainingErr := writeBuffer.WriteSerializable(ctx, m.GetUsesRemaining())
		if popErr := writeBuffer.PopContext("usesRemaining"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for usesRemaining")
		}
		if _usesRemainingErr != nil {
			return errors.Wrap(_usesRemainingErr, "Error serializing 'usesRemaining' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUsesRemaining"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUsesRemaining")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataUsesRemaining) isBACnetConstructedDataUsesRemaining() bool {
	return true
}

func (m *_BACnetConstructedDataUsesRemaining) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
