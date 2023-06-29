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

// BACnetConstructedDataEntryPoints is the corresponding interface of BACnetConstructedDataEntryPoints
type BACnetConstructedDataEntryPoints interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetEntryPoints returns EntryPoints (property field)
	GetEntryPoints() []BACnetDeviceObjectReference
}

// BACnetConstructedDataEntryPointsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEntryPoints.
// This is useful for switch cases.
type BACnetConstructedDataEntryPointsExactly interface {
	BACnetConstructedDataEntryPoints
	isBACnetConstructedDataEntryPoints() bool
}

// _BACnetConstructedDataEntryPoints is the data-structure of this message
type _BACnetConstructedDataEntryPoints struct {
	*_BACnetConstructedData
	EntryPoints []BACnetDeviceObjectReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEntryPoints) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEntryPoints) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ENTRY_POINTS
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEntryPoints) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataEntryPoints) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEntryPoints) GetEntryPoints() []BACnetDeviceObjectReference {
	return m.EntryPoints
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEntryPoints factory function for _BACnetConstructedDataEntryPoints
func NewBACnetConstructedDataEntryPoints(entryPoints []BACnetDeviceObjectReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEntryPoints {
	_result := &_BACnetConstructedDataEntryPoints{
		EntryPoints:            entryPoints,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEntryPoints(structType any) BACnetConstructedDataEntryPoints {
	if casted, ok := structType.(BACnetConstructedDataEntryPoints); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEntryPoints); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEntryPoints) GetTypeName() string {
	return "BACnetConstructedDataEntryPoints"
}

func (m *_BACnetConstructedDataEntryPoints) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.EntryPoints) > 0 {
		for _, element := range m.EntryPoints {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataEntryPoints) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataEntryPointsParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEntryPoints, error) {
	return BACnetConstructedDataEntryPointsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataEntryPointsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEntryPoints, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEntryPoints"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEntryPoints")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (entryPoints)
	if pullErr := readBuffer.PullContext("entryPoints", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for entryPoints")
	}
	// Terminated array
	var entryPoints []BACnetDeviceObjectReference
	{
		for !bool(IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber)) {
			_item, _err := BACnetDeviceObjectReferenceParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'entryPoints' field of BACnetConstructedDataEntryPoints")
			}
			entryPoints = append(entryPoints, _item.(BACnetDeviceObjectReference))
		}
	}
	if closeErr := readBuffer.CloseContext("entryPoints", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for entryPoints")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEntryPoints"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEntryPoints")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataEntryPoints{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		EntryPoints: entryPoints,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataEntryPoints) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEntryPoints) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEntryPoints"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEntryPoints")
		}

		// Array Field (entryPoints)
		if pushErr := writeBuffer.PushContext("entryPoints", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for entryPoints")
		}
		for _curItem, _element := range m.GetEntryPoints() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetEntryPoints()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'entryPoints' field")
			}
		}
		if popErr := writeBuffer.PopContext("entryPoints", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for entryPoints")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEntryPoints"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEntryPoints")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEntryPoints) isBACnetConstructedDataEntryPoints() bool {
	return true
}

func (m *_BACnetConstructedDataEntryPoints) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
