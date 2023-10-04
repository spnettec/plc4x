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

// BACnetConstructedDataPriorityArray is the corresponding interface of BACnetConstructedDataPriorityArray
type BACnetConstructedDataPriorityArray interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPriorityArray returns PriorityArray (property field)
	GetPriorityArray() BACnetPriorityArray
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetPriorityArray
}

// BACnetConstructedDataPriorityArrayExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataPriorityArray.
// This is useful for switch cases.
type BACnetConstructedDataPriorityArrayExactly interface {
	BACnetConstructedDataPriorityArray
	isBACnetConstructedDataPriorityArray() bool
}

// _BACnetConstructedDataPriorityArray is the data-structure of this message
type _BACnetConstructedDataPriorityArray struct {
	*_BACnetConstructedData
	PriorityArray BACnetPriorityArray
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPriorityArray) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataPriorityArray) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PRIORITY_ARRAY
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPriorityArray) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataPriorityArray) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPriorityArray) GetPriorityArray() BACnetPriorityArray {
	return m.PriorityArray
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPriorityArray) GetActualValue() BACnetPriorityArray {
	ctx := context.Background()
	_ = ctx
	return CastBACnetPriorityArray(m.GetPriorityArray())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataPriorityArray factory function for _BACnetConstructedDataPriorityArray
func NewBACnetConstructedDataPriorityArray(priorityArray BACnetPriorityArray, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataPriorityArray {
	_result := &_BACnetConstructedDataPriorityArray{
		PriorityArray:          priorityArray,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPriorityArray(structType any) BACnetConstructedDataPriorityArray {
	if casted, ok := structType.(BACnetConstructedDataPriorityArray); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPriorityArray); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPriorityArray) GetTypeName() string {
	return "BACnetConstructedDataPriorityArray"
}

func (m *_BACnetConstructedDataPriorityArray) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (priorityArray)
	lengthInBits += m.PriorityArray.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataPriorityArray) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataPriorityArrayParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPriorityArray, error) {
	return BACnetConstructedDataPriorityArrayParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataPriorityArrayParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPriorityArray, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPriorityArray"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPriorityArray")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (priorityArray)
	if pullErr := readBuffer.PullContext("priorityArray"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for priorityArray")
	}
	_priorityArray, _priorityArrayErr := BACnetPriorityArrayParseWithBuffer(ctx, readBuffer, BACnetObjectType(objectTypeArgument), uint8(tagNumber), arrayIndexArgument)
	if _priorityArrayErr != nil {
		return nil, errors.Wrap(_priorityArrayErr, "Error parsing 'priorityArray' field of BACnetConstructedDataPriorityArray")
	}
	priorityArray := _priorityArray.(BACnetPriorityArray)
	if closeErr := readBuffer.CloseContext("priorityArray"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for priorityArray")
	}

	// Virtual field
	_actualValue := priorityArray
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPriorityArray"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPriorityArray")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataPriorityArray{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		PriorityArray: priorityArray,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataPriorityArray) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPriorityArray) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPriorityArray"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPriorityArray")
		}

		// Simple Field (priorityArray)
		if pushErr := writeBuffer.PushContext("priorityArray"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for priorityArray")
		}
		_priorityArrayErr := writeBuffer.WriteSerializable(ctx, m.GetPriorityArray())
		if popErr := writeBuffer.PopContext("priorityArray"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for priorityArray")
		}
		if _priorityArrayErr != nil {
			return errors.Wrap(_priorityArrayErr, "Error serializing 'priorityArray' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPriorityArray"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPriorityArray")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataPriorityArray) isBACnetConstructedDataPriorityArray() bool {
	return true
}

func (m *_BACnetConstructedDataPriorityArray) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
