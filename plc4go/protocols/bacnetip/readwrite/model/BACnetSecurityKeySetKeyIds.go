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

// BACnetSecurityKeySetKeyIds is the corresponding interface of BACnetSecurityKeySetKeyIds
type BACnetSecurityKeySetKeyIds interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetKeyIds returns KeyIds (property field)
	GetKeyIds() []BACnetKeyIdentifier
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetSecurityKeySetKeyIdsExactly can be used when we want exactly this type and not a type which fulfills BACnetSecurityKeySetKeyIds.
// This is useful for switch cases.
type BACnetSecurityKeySetKeyIdsExactly interface {
	BACnetSecurityKeySetKeyIds
	isBACnetSecurityKeySetKeyIds() bool
}

// _BACnetSecurityKeySetKeyIds is the data-structure of this message
type _BACnetSecurityKeySetKeyIds struct {
	OpeningTag BACnetOpeningTag
	KeyIds     []BACnetKeyIdentifier
	ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetSecurityKeySetKeyIds) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetSecurityKeySetKeyIds) GetKeyIds() []BACnetKeyIdentifier {
	return m.KeyIds
}

func (m *_BACnetSecurityKeySetKeyIds) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetSecurityKeySetKeyIds factory function for _BACnetSecurityKeySetKeyIds
func NewBACnetSecurityKeySetKeyIds(openingTag BACnetOpeningTag, keyIds []BACnetKeyIdentifier, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetSecurityKeySetKeyIds {
	return &_BACnetSecurityKeySetKeyIds{OpeningTag: openingTag, KeyIds: keyIds, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetSecurityKeySetKeyIds(structType any) BACnetSecurityKeySetKeyIds {
	if casted, ok := structType.(BACnetSecurityKeySetKeyIds); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetSecurityKeySetKeyIds); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetSecurityKeySetKeyIds) GetTypeName() string {
	return "BACnetSecurityKeySetKeyIds"
}

func (m *_BACnetSecurityKeySetKeyIds) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Array field
	if len(m.KeyIds) > 0 {
		for _, element := range m.KeyIds {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetSecurityKeySetKeyIds) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetSecurityKeySetKeyIdsParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetSecurityKeySetKeyIds, error) {
	return BACnetSecurityKeySetKeyIdsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetSecurityKeySetKeyIdsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetSecurityKeySetKeyIds, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetSecurityKeySetKeyIds"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetSecurityKeySetKeyIds")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetSecurityKeySetKeyIds")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (keyIds)
	if pullErr := readBuffer.PullContext("keyIds", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for keyIds")
	}
	// Terminated array
	var keyIds []BACnetKeyIdentifier
	{
		for !bool(IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber)) {
			_item, _err := BACnetKeyIdentifierParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'keyIds' field of BACnetSecurityKeySetKeyIds")
			}
			keyIds = append(keyIds, _item.(BACnetKeyIdentifier))
		}
	}
	if closeErr := readBuffer.CloseContext("keyIds", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for keyIds")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetSecurityKeySetKeyIds")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetSecurityKeySetKeyIds"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetSecurityKeySetKeyIds")
	}

	// Create the instance
	return &_BACnetSecurityKeySetKeyIds{
		TagNumber:  tagNumber,
		OpeningTag: openingTag,
		KeyIds:     keyIds,
		ClosingTag: closingTag,
	}, nil
}

func (m *_BACnetSecurityKeySetKeyIds) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetSecurityKeySetKeyIds) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetSecurityKeySetKeyIds"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetSecurityKeySetKeyIds")
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

	// Array Field (keyIds)
	if pushErr := writeBuffer.PushContext("keyIds", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for keyIds")
	}
	for _curItem, _element := range m.GetKeyIds() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetKeyIds()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'keyIds' field")
		}
	}
	if popErr := writeBuffer.PopContext("keyIds", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for keyIds")
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

	if popErr := writeBuffer.PopContext("BACnetSecurityKeySetKeyIds"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetSecurityKeySetKeyIds")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetSecurityKeySetKeyIds) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetSecurityKeySetKeyIds) isBACnetSecurityKeySetKeyIds() bool {
	return true
}

func (m *_BACnetSecurityKeySetKeyIds) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
