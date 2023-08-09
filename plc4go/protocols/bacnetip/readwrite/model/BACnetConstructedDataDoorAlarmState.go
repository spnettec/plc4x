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

// BACnetConstructedDataDoorAlarmState is the corresponding interface of BACnetConstructedDataDoorAlarmState
type BACnetConstructedDataDoorAlarmState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDoorAlarmState returns DoorAlarmState (property field)
	GetDoorAlarmState() BACnetDoorAlarmStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDoorAlarmStateTagged
}

// BACnetConstructedDataDoorAlarmStateExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDoorAlarmState.
// This is useful for switch cases.
type BACnetConstructedDataDoorAlarmStateExactly interface {
	BACnetConstructedDataDoorAlarmState
	isBACnetConstructedDataDoorAlarmState() bool
}

// _BACnetConstructedDataDoorAlarmState is the data-structure of this message
type _BACnetConstructedDataDoorAlarmState struct {
	*_BACnetConstructedData
	DoorAlarmState BACnetDoorAlarmStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDoorAlarmState) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataDoorAlarmState) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DOOR_ALARM_STATE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDoorAlarmState) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDoorAlarmState) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDoorAlarmState) GetDoorAlarmState() BACnetDoorAlarmStateTagged {
	return m.DoorAlarmState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDoorAlarmState) GetActualValue() BACnetDoorAlarmStateTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDoorAlarmStateTagged(m.GetDoorAlarmState())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDoorAlarmState factory function for _BACnetConstructedDataDoorAlarmState
func NewBACnetConstructedDataDoorAlarmState(doorAlarmState BACnetDoorAlarmStateTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataDoorAlarmState {
	_result := &_BACnetConstructedDataDoorAlarmState{
		DoorAlarmState:         doorAlarmState,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDoorAlarmState(structType any) BACnetConstructedDataDoorAlarmState {
	if casted, ok := structType.(BACnetConstructedDataDoorAlarmState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDoorAlarmState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDoorAlarmState) GetTypeName() string {
	return "BACnetConstructedDataDoorAlarmState"
}

func (m *_BACnetConstructedDataDoorAlarmState) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (doorAlarmState)
	lengthInBits += m.DoorAlarmState.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataDoorAlarmState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataDoorAlarmStateParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDoorAlarmState, error) {
	return BACnetConstructedDataDoorAlarmStateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataDoorAlarmStateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDoorAlarmState, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDoorAlarmState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDoorAlarmState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (doorAlarmState)
	if pullErr := readBuffer.PullContext("doorAlarmState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for doorAlarmState")
	}
	_doorAlarmState, _doorAlarmStateErr := BACnetDoorAlarmStateTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _doorAlarmStateErr != nil {
		return nil, errors.Wrap(_doorAlarmStateErr, "Error parsing 'doorAlarmState' field of BACnetConstructedDataDoorAlarmState")
	}
	doorAlarmState := _doorAlarmState.(BACnetDoorAlarmStateTagged)
	if closeErr := readBuffer.CloseContext("doorAlarmState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for doorAlarmState")
	}

	// Virtual field
	_actualValue := doorAlarmState
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDoorAlarmState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDoorAlarmState")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDoorAlarmState{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DoorAlarmState: doorAlarmState,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDoorAlarmState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDoorAlarmState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDoorAlarmState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDoorAlarmState")
		}

		// Simple Field (doorAlarmState)
		if pushErr := writeBuffer.PushContext("doorAlarmState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for doorAlarmState")
		}
		_doorAlarmStateErr := writeBuffer.WriteSerializable(ctx, m.GetDoorAlarmState())
		if popErr := writeBuffer.PopContext("doorAlarmState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for doorAlarmState")
		}
		if _doorAlarmStateErr != nil {
			return errors.Wrap(_doorAlarmStateErr, "Error serializing 'doorAlarmState' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDoorAlarmState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDoorAlarmState")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataDoorAlarmState) isBACnetConstructedDataDoorAlarmState() bool {
	return true
}

func (m *_BACnetConstructedDataDoorAlarmState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
