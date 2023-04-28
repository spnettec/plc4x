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

// BACnetConstructedDataTrendLogLogBuffer is the corresponding interface of BACnetConstructedDataTrendLogLogBuffer
type BACnetConstructedDataTrendLogLogBuffer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFloorText returns FloorText (property field)
	GetFloorText() []BACnetLogRecord
}

// BACnetConstructedDataTrendLogLogBufferExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTrendLogLogBuffer.
// This is useful for switch cases.
type BACnetConstructedDataTrendLogLogBufferExactly interface {
	BACnetConstructedDataTrendLogLogBuffer
	isBACnetConstructedDataTrendLogLogBuffer() bool
}

// _BACnetConstructedDataTrendLogLogBuffer is the data-structure of this message
type _BACnetConstructedDataTrendLogLogBuffer struct {
	*_BACnetConstructedData
	FloorText []BACnetLogRecord
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTrendLogLogBuffer) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TREND_LOG
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOG_BUFFER
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTrendLogLogBuffer) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTrendLogLogBuffer) GetFloorText() []BACnetLogRecord {
	return m.FloorText
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTrendLogLogBuffer factory function for _BACnetConstructedDataTrendLogLogBuffer
func NewBACnetConstructedDataTrendLogLogBuffer(floorText []BACnetLogRecord, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTrendLogLogBuffer {
	_result := &_BACnetConstructedDataTrendLogLogBuffer{
		FloorText:              floorText,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTrendLogLogBuffer(structType any) BACnetConstructedDataTrendLogLogBuffer {
	if casted, ok := structType.(BACnetConstructedDataTrendLogLogBuffer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTrendLogLogBuffer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) GetTypeName() string {
	return "BACnetConstructedDataTrendLogLogBuffer"
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.FloorText) > 0 {
		for _, element := range m.FloorText {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataTrendLogLogBufferParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTrendLogLogBuffer, error) {
	return BACnetConstructedDataTrendLogLogBufferParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataTrendLogLogBufferParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTrendLogLogBuffer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTrendLogLogBuffer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTrendLogLogBuffer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (floorText)
	if pullErr := readBuffer.PullContext("floorText", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for floorText")
	}
	// Terminated array
	var floorText []BACnetLogRecord
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetLogRecordParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'floorText' field of BACnetConstructedDataTrendLogLogBuffer")
			}
			floorText = append(floorText, _item.(BACnetLogRecord))
		}
	}
	if closeErr := readBuffer.CloseContext("floorText", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for floorText")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTrendLogLogBuffer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTrendLogLogBuffer")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTrendLogLogBuffer{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FloorText: floorText,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTrendLogLogBuffer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTrendLogLogBuffer")
		}

		// Array Field (floorText)
		if pushErr := writeBuffer.PushContext("floorText", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for floorText")
		}
		for _curItem, _element := range m.GetFloorText() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetFloorText()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'floorText' field")
			}
		}
		if popErr := writeBuffer.PopContext("floorText", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for floorText")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTrendLogLogBuffer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTrendLogLogBuffer")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) isBACnetConstructedDataTrendLogLogBuffer() bool {
	return true
}

func (m *_BACnetConstructedDataTrendLogLogBuffer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
