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

// BACnetFaultParameterFaultOutOfRange is the corresponding interface of BACnetFaultParameterFaultOutOfRange
type BACnetFaultParameterFaultOutOfRange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetMinNormalValue returns MinNormalValue (property field)
	GetMinNormalValue() BACnetFaultParameterFaultOutOfRangeMinNormalValue
	// GetMaxNormalValue returns MaxNormalValue (property field)
	GetMaxNormalValue() BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetFaultParameterFaultOutOfRangeExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultOutOfRange.
// This is useful for switch cases.
type BACnetFaultParameterFaultOutOfRangeExactly interface {
	BACnetFaultParameterFaultOutOfRange
	isBACnetFaultParameterFaultOutOfRange() bool
}

// _BACnetFaultParameterFaultOutOfRange is the data-structure of this message
type _BACnetFaultParameterFaultOutOfRange struct {
	*_BACnetFaultParameter
	OpeningTag     BACnetOpeningTag
	MinNormalValue BACnetFaultParameterFaultOutOfRangeMinNormalValue
	MaxNormalValue BACnetFaultParameterFaultOutOfRangeMaxNormalValue
	ClosingTag     BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultOutOfRange) InitializeParent(parent BACnetFaultParameter, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetFaultParameterFaultOutOfRange) GetParent() BACnetFaultParameter {
	return m._BACnetFaultParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultOutOfRange) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultOutOfRange) GetMinNormalValue() BACnetFaultParameterFaultOutOfRangeMinNormalValue {
	return m.MinNormalValue
}

func (m *_BACnetFaultParameterFaultOutOfRange) GetMaxNormalValue() BACnetFaultParameterFaultOutOfRangeMaxNormalValue {
	return m.MaxNormalValue
}

func (m *_BACnetFaultParameterFaultOutOfRange) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultOutOfRange factory function for _BACnetFaultParameterFaultOutOfRange
func NewBACnetFaultParameterFaultOutOfRange(openingTag BACnetOpeningTag, minNormalValue BACnetFaultParameterFaultOutOfRangeMinNormalValue, maxNormalValue BACnetFaultParameterFaultOutOfRangeMaxNormalValue, closingTag BACnetClosingTag, peekedTagHeader BACnetTagHeader) *_BACnetFaultParameterFaultOutOfRange {
	_result := &_BACnetFaultParameterFaultOutOfRange{
		OpeningTag:            openingTag,
		MinNormalValue:        minNormalValue,
		MaxNormalValue:        maxNormalValue,
		ClosingTag:            closingTag,
		_BACnetFaultParameter: NewBACnetFaultParameter(peekedTagHeader),
	}
	_result._BACnetFaultParameter._BACnetFaultParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultOutOfRange(structType any) BACnetFaultParameterFaultOutOfRange {
	if casted, ok := structType.(BACnetFaultParameterFaultOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultOutOfRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultOutOfRange) GetTypeName() string {
	return "BACnetFaultParameterFaultOutOfRange"
}

func (m *_BACnetFaultParameterFaultOutOfRange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (minNormalValue)
	lengthInBits += m.MinNormalValue.GetLengthInBits(ctx)

	// Simple field (maxNormalValue)
	lengthInBits += m.MaxNormalValue.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultOutOfRange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultOutOfRangeParse(ctx context.Context, theBytes []byte) (BACnetFaultParameterFaultOutOfRange, error) {
	return BACnetFaultParameterFaultOutOfRangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetFaultParameterFaultOutOfRangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultOutOfRange, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultOutOfRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultOutOfRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(uint8(6)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetFaultParameterFaultOutOfRange")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (minNormalValue)
	if pullErr := readBuffer.PullContext("minNormalValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for minNormalValue")
	}
	_minNormalValue, _minNormalValueErr := BACnetFaultParameterFaultOutOfRangeMinNormalValueParseWithBuffer(ctx, readBuffer, uint8(uint8(0)))
	if _minNormalValueErr != nil {
		return nil, errors.Wrap(_minNormalValueErr, "Error parsing 'minNormalValue' field of BACnetFaultParameterFaultOutOfRange")
	}
	minNormalValue := _minNormalValue.(BACnetFaultParameterFaultOutOfRangeMinNormalValue)
	if closeErr := readBuffer.CloseContext("minNormalValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for minNormalValue")
	}

	// Simple Field (maxNormalValue)
	if pullErr := readBuffer.PullContext("maxNormalValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxNormalValue")
	}
	_maxNormalValue, _maxNormalValueErr := BACnetFaultParameterFaultOutOfRangeMaxNormalValueParseWithBuffer(ctx, readBuffer, uint8(uint8(0)))
	if _maxNormalValueErr != nil {
		return nil, errors.Wrap(_maxNormalValueErr, "Error parsing 'maxNormalValue' field of BACnetFaultParameterFaultOutOfRange")
	}
	maxNormalValue := _maxNormalValue.(BACnetFaultParameterFaultOutOfRangeMaxNormalValue)
	if closeErr := readBuffer.CloseContext("maxNormalValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxNormalValue")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(uint8(6)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetFaultParameterFaultOutOfRange")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultOutOfRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultOutOfRange")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultOutOfRange{
		_BACnetFaultParameter: &_BACnetFaultParameter{},
		OpeningTag:            openingTag,
		MinNormalValue:        minNormalValue,
		MaxNormalValue:        maxNormalValue,
		ClosingTag:            closingTag,
	}
	_child._BACnetFaultParameter._BACnetFaultParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultOutOfRange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultOutOfRange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultOutOfRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultOutOfRange")
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

		// Simple Field (minNormalValue)
		if pushErr := writeBuffer.PushContext("minNormalValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for minNormalValue")
		}
		_minNormalValueErr := writeBuffer.WriteSerializable(ctx, m.GetMinNormalValue())
		if popErr := writeBuffer.PopContext("minNormalValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for minNormalValue")
		}
		if _minNormalValueErr != nil {
			return errors.Wrap(_minNormalValueErr, "Error serializing 'minNormalValue' field")
		}

		// Simple Field (maxNormalValue)
		if pushErr := writeBuffer.PushContext("maxNormalValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maxNormalValue")
		}
		_maxNormalValueErr := writeBuffer.WriteSerializable(ctx, m.GetMaxNormalValue())
		if popErr := writeBuffer.PopContext("maxNormalValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maxNormalValue")
		}
		if _maxNormalValueErr != nil {
			return errors.Wrap(_maxNormalValueErr, "Error serializing 'maxNormalValue' field")
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

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultOutOfRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultOutOfRange")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultOutOfRange) isBACnetFaultParameterFaultOutOfRange() bool {
	return true
}

func (m *_BACnetFaultParameterFaultOutOfRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
