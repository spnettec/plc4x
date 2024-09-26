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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
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
	// IsMediaTransportControlDataSetCategory is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMediaTransportControlDataSetCategory()
}

// _MediaTransportControlDataSetCategory is the data-structure of this message
type _MediaTransportControlDataSetCategory struct {
	MediaTransportControlDataContract
	CategoryNumber uint8
}

var _ MediaTransportControlDataSetCategory = (*_MediaTransportControlDataSetCategory)(nil)
var _ MediaTransportControlDataRequirements = (*_MediaTransportControlDataSetCategory)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataSetCategory) GetParent() MediaTransportControlDataContract {
	return m.MediaTransportControlDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataSetCategory) GetCategoryNumber() uint8 {
	return m.CategoryNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMediaTransportControlDataSetCategory factory function for _MediaTransportControlDataSetCategory
func NewMediaTransportControlDataSetCategory(commandTypeContainer MediaTransportControlCommandTypeContainer, mediaLinkGroup byte, categoryNumber uint8) *_MediaTransportControlDataSetCategory {
	_result := &_MediaTransportControlDataSetCategory{
		MediaTransportControlDataContract: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
		CategoryNumber:                    categoryNumber,
	}
	_result.MediaTransportControlDataContract.(*_MediaTransportControlData)._SubType = _result
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
	lengthInBits := uint16(m.MediaTransportControlDataContract.(*_MediaTransportControlData).getLengthInBits(ctx))

	// Simple field (categoryNumber)
	lengthInBits += 8

	return lengthInBits
}

func (m *_MediaTransportControlDataSetCategory) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MediaTransportControlDataSetCategory) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MediaTransportControlData) (__mediaTransportControlDataSetCategory MediaTransportControlDataSetCategory, err error) {
	m.MediaTransportControlDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataSetCategory"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataSetCategory")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	categoryNumber, err := ReadSimpleField(ctx, "categoryNumber", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'categoryNumber' field"))
	}
	m.CategoryNumber = categoryNumber

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataSetCategory"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataSetCategory")
	}

	return m, nil
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

		if err := WriteSimpleField[uint8](ctx, "categoryNumber", m.GetCategoryNumber(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'categoryNumber' field")
		}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataSetCategory"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataSetCategory")
		}
		return nil
	}
	return m.MediaTransportControlDataContract.(*_MediaTransportControlData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MediaTransportControlDataSetCategory) IsMediaTransportControlDataSetCategory() {}

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
