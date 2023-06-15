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


// MediaTransportControlDataSetCategory is the corresponding interface of MediaTransportControlDataSetCategory
type MediaTransportControlDataSetCategory interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetCategoryNumber returns CategoryNumber (property field)
	GetCategoryNumber() uint8
}

// MediaTransportControlDataSetCategoryExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataSetCategory.
// This is useful for switch cases.
type MediaTransportControlDataSetCategoryExactly interface {
	MediaTransportControlDataSetCategory
	isMediaTransportControlDataSetCategory() bool
}

// _MediaTransportControlDataSetCategory is the data-structure of this message
type _MediaTransportControlDataSetCategory struct {
	*_MediaTransportControlData
        CategoryNumber uint8
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataSetCategory) InitializeParent(parent MediaTransportControlData , commandTypeContainer MediaTransportControlCommandTypeContainer , mediaLinkGroup byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.MediaLinkGroup = mediaLinkGroup
}

func (m *_MediaTransportControlDataSetCategory)  GetParent() MediaTransportControlData {
	return m._MediaTransportControlData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataSetCategory) GetCategoryNumber() uint8 {
	return m.CategoryNumber
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewMediaTransportControlDataSetCategory factory function for _MediaTransportControlDataSetCategory
func NewMediaTransportControlDataSetCategory( categoryNumber uint8 , commandTypeContainer MediaTransportControlCommandTypeContainer , mediaLinkGroup byte ) *_MediaTransportControlDataSetCategory {
	_result := &_MediaTransportControlDataSetCategory{
		CategoryNumber: categoryNumber,
    	_MediaTransportControlData: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result._MediaTransportControlData._MediaTransportControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataSetCategory(structType any) MediaTransportControlDataSetCategory {
    if casted, ok := structType.(MediaTransportControlDataSetCategory); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataSetCategory); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataSetCategory) GetTypeName() string {
	return "MediaTransportControlDataSetCategory"
}

func (m *_MediaTransportControlDataSetCategory) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (categoryNumber)
	lengthInBits += 8;

	return lengthInBits
}


func (m *_MediaTransportControlDataSetCategory) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MediaTransportControlDataSetCategoryParse(ctx context.Context, theBytes []byte) (MediaTransportControlDataSetCategory, error) {
	return MediaTransportControlDataSetCategoryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func MediaTransportControlDataSetCategoryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (MediaTransportControlDataSetCategory, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MediaTransportControlDataSetCategory"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataSetCategory")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (categoryNumber)
_categoryNumber, _categoryNumberErr := readBuffer.ReadUint8("categoryNumber", 8)
	if _categoryNumberErr != nil {
		return nil, errors.Wrap(_categoryNumberErr, "Error parsing 'categoryNumber' field of MediaTransportControlDataSetCategory")
	}
	categoryNumber := _categoryNumber

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataSetCategory"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataSetCategory")
	}

	// Create a partially initialized instance
	_child := &_MediaTransportControlDataSetCategory{
		_MediaTransportControlData: &_MediaTransportControlData{
		},
		CategoryNumber: categoryNumber,
	}
	_child._MediaTransportControlData._MediaTransportControlDataChildRequirements = _child
	return _child, nil
}

func (m *_MediaTransportControlDataSetCategory) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataSetCategory) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataSetCategory"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataSetCategory")
		}

	// Simple Field (categoryNumber)
	categoryNumber := uint8(m.GetCategoryNumber())
	_categoryNumberErr := writeBuffer.WriteUint8("categoryNumber", 8, (categoryNumber))
	if _categoryNumberErr != nil {
		return errors.Wrap(_categoryNumberErr, "Error serializing 'categoryNumber' field")
	}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataSetCategory"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataSetCategory")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_MediaTransportControlDataSetCategory) isMediaTransportControlDataSetCategory() bool {
	return true
}

func (m *_MediaTransportControlDataSetCategory) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



