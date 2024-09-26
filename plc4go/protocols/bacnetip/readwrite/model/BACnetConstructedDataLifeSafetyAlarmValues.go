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

// BACnetConstructedDataLifeSafetyAlarmValues is the corresponding interface of BACnetConstructedDataLifeSafetyAlarmValues
type BACnetConstructedDataLifeSafetyAlarmValues interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAlarmValues returns AlarmValues (property field)
	GetAlarmValues() []BACnetLifeSafetyStateTagged
	// IsBACnetConstructedDataLifeSafetyAlarmValues is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataLifeSafetyAlarmValues()
}

// _BACnetConstructedDataLifeSafetyAlarmValues is the data-structure of this message
type _BACnetConstructedDataLifeSafetyAlarmValues struct {
	BACnetConstructedDataContract
	AlarmValues []BACnetLifeSafetyStateTagged
}

var _ BACnetConstructedDataLifeSafetyAlarmValues = (*_BACnetConstructedDataLifeSafetyAlarmValues)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataLifeSafetyAlarmValues)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIFE_SAFETY_ALARM_VALUES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) GetAlarmValues() []BACnetLifeSafetyStateTagged {
	return m.AlarmValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLifeSafetyAlarmValues factory function for _BACnetConstructedDataLifeSafetyAlarmValues
func NewBACnetConstructedDataLifeSafetyAlarmValues(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, alarmValues []BACnetLifeSafetyStateTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLifeSafetyAlarmValues {
	_result := &_BACnetConstructedDataLifeSafetyAlarmValues{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AlarmValues:                   alarmValues,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLifeSafetyAlarmValues(structType any) BACnetConstructedDataLifeSafetyAlarmValues {
	if casted, ok := structType.(BACnetConstructedDataLifeSafetyAlarmValues); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLifeSafetyAlarmValues); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) GetTypeName() string {
	return "BACnetConstructedDataLifeSafetyAlarmValues"
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Array field
	if len(m.AlarmValues) > 0 {
		for _, element := range m.AlarmValues {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataLifeSafetyAlarmValues BACnetConstructedDataLifeSafetyAlarmValues, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLifeSafetyAlarmValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLifeSafetyAlarmValues")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	alarmValues, err := ReadTerminatedArrayField[BACnetLifeSafetyStateTagged](ctx, "alarmValues", ReadComplex[BACnetLifeSafetyStateTagged](BACnetLifeSafetyStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer), IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'alarmValues' field"))
	}
	m.AlarmValues = alarmValues

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLifeSafetyAlarmValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLifeSafetyAlarmValues")
	}

	return m, nil
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLifeSafetyAlarmValues"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLifeSafetyAlarmValues")
		}

		if err := WriteComplexTypeArrayField(ctx, "alarmValues", m.GetAlarmValues(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'alarmValues' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLifeSafetyAlarmValues"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLifeSafetyAlarmValues")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) IsBACnetConstructedDataLifeSafetyAlarmValues() {
}

func (m *_BACnetConstructedDataLifeSafetyAlarmValues) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
