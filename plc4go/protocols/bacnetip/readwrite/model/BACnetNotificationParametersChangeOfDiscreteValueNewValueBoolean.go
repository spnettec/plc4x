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

// BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean is the corresponding interface of BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
type BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParametersChangeOfDiscreteValueNewValue
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() BACnetApplicationTagBoolean
}

// BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean.
// This is useful for switch cases.
type BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanExactly interface {
	BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
	isBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean() bool
}

// _BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean is the data-structure of this message
type _BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean struct {
	*_BACnetNotificationParametersChangeOfDiscreteValueNewValue
	BooleanValue BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) InitializeParent(parent BACnetNotificationParametersChangeOfDiscreteValueNewValue, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetParent() BACnetNotificationParametersChangeOfDiscreteValueNewValue {
	return m._BACnetNotificationParametersChangeOfDiscreteValueNewValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetBooleanValue() BACnetApplicationTagBoolean {
	return m.BooleanValue
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean factory function for _BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean
func NewBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean(booleanValue BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean {
	_result := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean{
		BooleanValue: booleanValue,
		_BACnetNotificationParametersChangeOfDiscreteValueNewValue: NewBACnetNotificationParametersChangeOfDiscreteValueNewValue(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetNotificationParametersChangeOfDiscreteValueNewValue._BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean(structType any) BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean, error) {
	return BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetNotificationParametersChangeOfDiscreteValueNewValueBooleanParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (booleanValue)
	if pullErr := readBuffer.PullContext("booleanValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for booleanValue")
	}
	_booleanValue, _booleanValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _booleanValueErr != nil {
		return nil, errors.Wrap(_booleanValueErr, "Error parsing 'booleanValue' field of BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
	}
	booleanValue := _booleanValue.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("booleanValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for booleanValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean{
		_BACnetNotificationParametersChangeOfDiscreteValueNewValue: &_BACnetNotificationParametersChangeOfDiscreteValueNewValue{
			TagNumber: tagNumber,
		},
		BooleanValue: booleanValue,
	}
	_child._BACnetNotificationParametersChangeOfDiscreteValueNewValue._BACnetNotificationParametersChangeOfDiscreteValueNewValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
		}

		// Simple Field (booleanValue)
		if pushErr := writeBuffer.PushContext("booleanValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for booleanValue")
		}
		_booleanValueErr := writeBuffer.WriteSerializable(ctx, m.GetBooleanValue())
		if popErr := writeBuffer.PopContext("booleanValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for booleanValue")
		}
		if _booleanValueErr != nil {
			return errors.Wrap(_booleanValueErr, "Error serializing 'booleanValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) isBACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean() bool {
	return true
}

func (m *_BACnetNotificationParametersChangeOfDiscreteValueNewValueBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
