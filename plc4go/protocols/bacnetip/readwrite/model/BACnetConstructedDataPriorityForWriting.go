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

// BACnetConstructedDataPriorityForWriting is the corresponding interface of BACnetConstructedDataPriorityForWriting
type BACnetConstructedDataPriorityForWriting interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPriorityForWriting returns PriorityForWriting (property field)
	GetPriorityForWriting() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataPriorityForWritingExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataPriorityForWriting.
// This is useful for switch cases.
type BACnetConstructedDataPriorityForWritingExactly interface {
	BACnetConstructedDataPriorityForWriting
	isBACnetConstructedDataPriorityForWriting() bool
}

// _BACnetConstructedDataPriorityForWriting is the data-structure of this message
type _BACnetConstructedDataPriorityForWriting struct {
	*_BACnetConstructedData
	PriorityForWriting BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPriorityForWriting) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataPriorityForWriting) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRIORITY_FOR_WRITING
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPriorityForWriting) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataPriorityForWriting) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPriorityForWriting) GetPriorityForWriting() BACnetApplicationTagUnsignedInteger {
	return m.PriorityForWriting
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPriorityForWriting) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetPriorityForWriting())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPriorityForWriting factory function for _BACnetConstructedDataPriorityForWriting
func NewBACnetConstructedDataPriorityForWriting(priorityForWriting BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPriorityForWriting {
	_result := &_BACnetConstructedDataPriorityForWriting{
		PriorityForWriting:     priorityForWriting,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPriorityForWriting(structType any) BACnetConstructedDataPriorityForWriting {
	if casted, ok := structType.(BACnetConstructedDataPriorityForWriting); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPriorityForWriting); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPriorityForWriting) GetTypeName() string {
	return "BACnetConstructedDataPriorityForWriting"
}

func (m *_BACnetConstructedDataPriorityForWriting) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (priorityForWriting)
	lengthInBits += m.PriorityForWriting.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPriorityForWriting) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataPriorityForWritingParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPriorityForWriting, error) {
	return BACnetConstructedDataPriorityForWritingParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataPriorityForWritingParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPriorityForWriting, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPriorityForWriting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPriorityForWriting")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (priorityForWriting)
	if pullErr := readBuffer.PullContext("priorityForWriting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for priorityForWriting")
	}
	_priorityForWriting, _priorityForWritingErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _priorityForWritingErr != nil {
		return nil, errors.Wrap(_priorityForWritingErr, "Error parsing 'priorityForWriting' field of BACnetConstructedDataPriorityForWriting")
	}
	priorityForWriting := _priorityForWriting.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("priorityForWriting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for priorityForWriting")
	}

	// Virtual field
	_actualValue := priorityForWriting
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPriorityForWriting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPriorityForWriting")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataPriorityForWriting{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		PriorityForWriting: priorityForWriting,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataPriorityForWriting) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPriorityForWriting) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPriorityForWriting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPriorityForWriting")
		}

		// Simple Field (priorityForWriting)
		if pushErr := writeBuffer.PushContext("priorityForWriting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for priorityForWriting")
		}
		_priorityForWritingErr := writeBuffer.WriteSerializable(ctx, m.GetPriorityForWriting())
		if popErr := writeBuffer.PopContext("priorityForWriting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for priorityForWriting")
		}
		if _priorityForWritingErr != nil {
			return errors.Wrap(_priorityForWritingErr, "Error serializing 'priorityForWriting' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPriorityForWriting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPriorityForWriting")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPriorityForWriting) isBACnetConstructedDataPriorityForWriting() bool {
	return true
}

func (m *_BACnetConstructedDataPriorityForWriting) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
