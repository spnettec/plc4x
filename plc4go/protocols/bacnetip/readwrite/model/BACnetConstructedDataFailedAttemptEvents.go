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

// BACnetConstructedDataFailedAttemptEvents is the corresponding interface of BACnetConstructedDataFailedAttemptEvents
type BACnetConstructedDataFailedAttemptEvents interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFailedAttemptEvents returns FailedAttemptEvents (property field)
	GetFailedAttemptEvents() []BACnetAccessEventTagged
	// IsBACnetConstructedDataFailedAttemptEvents is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataFailedAttemptEvents()
}

// _BACnetConstructedDataFailedAttemptEvents is the data-structure of this message
type _BACnetConstructedDataFailedAttemptEvents struct {
	BACnetConstructedDataContract
	FailedAttemptEvents []BACnetAccessEventTagged
}

var _ BACnetConstructedDataFailedAttemptEvents = (*_BACnetConstructedDataFailedAttemptEvents)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataFailedAttemptEvents)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFailedAttemptEvents) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataFailedAttemptEvents) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FAILED_ATTEMPT_EVENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFailedAttemptEvents) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFailedAttemptEvents) GetFailedAttemptEvents() []BACnetAccessEventTagged {
	return m.FailedAttemptEvents
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataFailedAttemptEvents factory function for _BACnetConstructedDataFailedAttemptEvents
func NewBACnetConstructedDataFailedAttemptEvents(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, failedAttemptEvents []BACnetAccessEventTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataFailedAttemptEvents {
	_result := &_BACnetConstructedDataFailedAttemptEvents{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		FailedAttemptEvents:           failedAttemptEvents,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFailedAttemptEvents(structType any) BACnetConstructedDataFailedAttemptEvents {
	if casted, ok := structType.(BACnetConstructedDataFailedAttemptEvents); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFailedAttemptEvents); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFailedAttemptEvents) GetTypeName() string {
	return "BACnetConstructedDataFailedAttemptEvents"
}

func (m *_BACnetConstructedDataFailedAttemptEvents) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.FailedAttemptEvents) > 0 {
		for _, element := range m.FailedAttemptEvents {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataFailedAttemptEvents) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataFailedAttemptEvents) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataFailedAttemptEvents BACnetConstructedDataFailedAttemptEvents, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFailedAttemptEvents"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFailedAttemptEvents")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	failedAttemptEvents, err := ReadTerminatedArrayField[BACnetAccessEventTagged](ctx, "failedAttemptEvents", ReadComplex[BACnetAccessEventTagged](BACnetAccessEventTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'failedAttemptEvents' field"))
	}
	m.FailedAttemptEvents = failedAttemptEvents

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFailedAttemptEvents"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFailedAttemptEvents")
	}

	return m, nil
}

func (m *_BACnetConstructedDataFailedAttemptEvents) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFailedAttemptEvents) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFailedAttemptEvents"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFailedAttemptEvents")
		}

		if err := WriteComplexTypeArrayField(ctx, "failedAttemptEvents", m.GetFailedAttemptEvents(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'failedAttemptEvents' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFailedAttemptEvents"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFailedAttemptEvents")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataFailedAttemptEvents) IsBACnetConstructedDataFailedAttemptEvents() {}

func (m *_BACnetConstructedDataFailedAttemptEvents) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
