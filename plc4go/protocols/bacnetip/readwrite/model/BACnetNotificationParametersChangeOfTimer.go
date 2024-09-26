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

// BACnetNotificationParametersChangeOfTimer is the corresponding interface of BACnetNotificationParametersChangeOfTimer
type BACnetNotificationParametersChangeOfTimer interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetNewValue returns NewValue (property field)
	GetNewValue() BACnetTimerStateTagged
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetUpdateTime returns UpdateTime (property field)
	GetUpdateTime() BACnetDateTimeEnclosed
	// GetLastStateChange returns LastStateChange (property field)
	GetLastStateChange() BACnetTimerTransitionTagged
	// GetInitialTimeout returns InitialTimeout (property field)
	GetInitialTimeout() BACnetContextTagUnsignedInteger
	// GetExpirationTime returns ExpirationTime (property field)
	GetExpirationTime() BACnetDateTimeEnclosed
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
	// IsBACnetNotificationParametersChangeOfTimer is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetNotificationParametersChangeOfTimer()
}

// _BACnetNotificationParametersChangeOfTimer is the data-structure of this message
type _BACnetNotificationParametersChangeOfTimer struct {
	BACnetNotificationParametersContract
	InnerOpeningTag BACnetOpeningTag
	NewValue        BACnetTimerStateTagged
	StatusFlags     BACnetStatusFlagsTagged
	UpdateTime      BACnetDateTimeEnclosed
	LastStateChange BACnetTimerTransitionTagged
	InitialTimeout  BACnetContextTagUnsignedInteger
	ExpirationTime  BACnetDateTimeEnclosed
	InnerClosingTag BACnetClosingTag
}

var _ BACnetNotificationParametersChangeOfTimer = (*_BACnetNotificationParametersChangeOfTimer)(nil)
var _ BACnetNotificationParametersRequirements = (*_BACnetNotificationParametersChangeOfTimer)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersChangeOfTimer) GetParent() BACnetNotificationParametersContract {
	return m.BACnetNotificationParametersContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersChangeOfTimer) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetNewValue() BACnetTimerStateTagged {
	return m.NewValue
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetUpdateTime() BACnetDateTimeEnclosed {
	return m.UpdateTime
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetLastStateChange() BACnetTimerTransitionTagged {
	return m.LastStateChange
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetInitialTimeout() BACnetContextTagUnsignedInteger {
	return m.InitialTimeout
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetExpirationTime() BACnetDateTimeEnclosed {
	return m.ExpirationTime
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersChangeOfTimer factory function for _BACnetNotificationParametersChangeOfTimer
func NewBACnetNotificationParametersChangeOfTimer(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, innerOpeningTag BACnetOpeningTag, newValue BACnetTimerStateTagged, statusFlags BACnetStatusFlagsTagged, updateTime BACnetDateTimeEnclosed, lastStateChange BACnetTimerTransitionTagged, initialTimeout BACnetContextTagUnsignedInteger, expirationTime BACnetDateTimeEnclosed, innerClosingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersChangeOfTimer {
	if innerOpeningTag == nil {
		panic("innerOpeningTag of type BACnetOpeningTag for BACnetNotificationParametersChangeOfTimer must not be nil")
	}
	if newValue == nil {
		panic("newValue of type BACnetTimerStateTagged for BACnetNotificationParametersChangeOfTimer must not be nil")
	}
	if statusFlags == nil {
		panic("statusFlags of type BACnetStatusFlagsTagged for BACnetNotificationParametersChangeOfTimer must not be nil")
	}
	if updateTime == nil {
		panic("updateTime of type BACnetDateTimeEnclosed for BACnetNotificationParametersChangeOfTimer must not be nil")
	}
	if innerClosingTag == nil {
		panic("innerClosingTag of type BACnetClosingTag for BACnetNotificationParametersChangeOfTimer must not be nil")
	}
	_result := &_BACnetNotificationParametersChangeOfTimer{
		BACnetNotificationParametersContract: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
		InnerOpeningTag:                      innerOpeningTag,
		NewValue:                             newValue,
		StatusFlags:                          statusFlags,
		UpdateTime:                           updateTime,
		LastStateChange:                      lastStateChange,
		InitialTimeout:                       initialTimeout,
		ExpirationTime:                       expirationTime,
		InnerClosingTag:                      innerClosingTag,
	}
	_result.BACnetNotificationParametersContract.(*_BACnetNotificationParameters)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersChangeOfTimer(structType any) BACnetNotificationParametersChangeOfTimer {
	if casted, ok := structType.(BACnetNotificationParametersChangeOfTimer); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersChangeOfTimer); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetTypeName() string {
	return "BACnetNotificationParametersChangeOfTimer"
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).getLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (newValue)
	lengthInBits += m.NewValue.GetLengthInBits(ctx)

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// Simple field (updateTime)
	lengthInBits += m.UpdateTime.GetLengthInBits(ctx)

	// Optional Field (lastStateChange)
	if m.LastStateChange != nil {
		lengthInBits += m.LastStateChange.GetLengthInBits(ctx)
	}

	// Optional Field (initialTimeout)
	if m.InitialTimeout != nil {
		lengthInBits += m.InitialTimeout.GetLengthInBits(ctx)
	}

	// Optional Field (expirationTime)
	if m.ExpirationTime != nil {
		lengthInBits += m.ExpirationTime.GetLengthInBits(ctx)
	}

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersChangeOfTimer) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetNotificationParametersChangeOfTimer) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetNotificationParameters, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (__bACnetNotificationParametersChangeOfTimer BACnetNotificationParametersChangeOfTimer, err error) {
	m.BACnetNotificationParametersContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersChangeOfTimer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersChangeOfTimer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	innerOpeningTag, err := ReadSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", ReadComplex[BACnetOpeningTag](BACnetOpeningTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerOpeningTag' field"))
	}
	m.InnerOpeningTag = innerOpeningTag

	newValue, err := ReadSimpleField[BACnetTimerStateTagged](ctx, "newValue", ReadComplex[BACnetTimerStateTagged](BACnetTimerStateTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'newValue' field"))
	}
	m.NewValue = newValue

	statusFlags, err := ReadSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", ReadComplex[BACnetStatusFlagsTagged](BACnetStatusFlagsTaggedParseWithBufferProducer((uint8)(uint8(1)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusFlags' field"))
	}
	m.StatusFlags = statusFlags

	updateTime, err := ReadSimpleField[BACnetDateTimeEnclosed](ctx, "updateTime", ReadComplex[BACnetDateTimeEnclosed](BACnetDateTimeEnclosedParseWithBufferProducer((uint8)(uint8(2))), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'updateTime' field"))
	}
	m.UpdateTime = updateTime

	var lastStateChange BACnetTimerTransitionTagged
	_lastStateChange, err := ReadOptionalField[BACnetTimerTransitionTagged](ctx, "lastStateChange", ReadComplex[BACnetTimerTransitionTagged](BACnetTimerTransitionTaggedParseWithBufferProducer((uint8)(uint8(3)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'lastStateChange' field"))
	}
	if _lastStateChange != nil {
		lastStateChange = *_lastStateChange
		m.LastStateChange = lastStateChange
	}

	var initialTimeout BACnetContextTagUnsignedInteger
	_initialTimeout, err := ReadOptionalField[BACnetContextTagUnsignedInteger](ctx, "initialTimeout", ReadComplex[BACnetContextTagUnsignedInteger](BACnetContextTagParseWithBufferProducer[BACnetContextTagUnsignedInteger]((uint8)(uint8(4)), (BACnetDataType)(BACnetDataType_UNSIGNED_INTEGER)), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'initialTimeout' field"))
	}
	if _initialTimeout != nil {
		initialTimeout = *_initialTimeout
		m.InitialTimeout = initialTimeout
	}

	var expirationTime BACnetDateTimeEnclosed
	_expirationTime, err := ReadOptionalField[BACnetDateTimeEnclosed](ctx, "expirationTime", ReadComplex[BACnetDateTimeEnclosed](BACnetDateTimeEnclosedParseWithBufferProducer((uint8)(uint8(5))), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'expirationTime' field"))
	}
	if _expirationTime != nil {
		expirationTime = *_expirationTime
		m.ExpirationTime = expirationTime
	}

	innerClosingTag, err := ReadSimpleField[BACnetClosingTag](ctx, "innerClosingTag", ReadComplex[BACnetClosingTag](BACnetClosingTagParseWithBufferProducer((uint8)(peekedTagNumber)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'innerClosingTag' field"))
	}
	m.InnerClosingTag = innerClosingTag

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersChangeOfTimer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersChangeOfTimer")
	}

	return m, nil
}

func (m *_BACnetNotificationParametersChangeOfTimer) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersChangeOfTimer) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersChangeOfTimer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersChangeOfTimer")
		}

		if err := WriteSimpleField[BACnetOpeningTag](ctx, "innerOpeningTag", m.GetInnerOpeningTag(), WriteComplex[BACnetOpeningTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerOpeningTag' field")
		}

		if err := WriteSimpleField[BACnetTimerStateTagged](ctx, "newValue", m.GetNewValue(), WriteComplex[BACnetTimerStateTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'newValue' field")
		}

		if err := WriteSimpleField[BACnetStatusFlagsTagged](ctx, "statusFlags", m.GetStatusFlags(), WriteComplex[BACnetStatusFlagsTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusFlags' field")
		}

		if err := WriteSimpleField[BACnetDateTimeEnclosed](ctx, "updateTime", m.GetUpdateTime(), WriteComplex[BACnetDateTimeEnclosed](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'updateTime' field")
		}

		if err := WriteOptionalField[BACnetTimerTransitionTagged](ctx, "lastStateChange", GetRef(m.GetLastStateChange()), WriteComplex[BACnetTimerTransitionTagged](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'lastStateChange' field")
		}

		if err := WriteOptionalField[BACnetContextTagUnsignedInteger](ctx, "initialTimeout", GetRef(m.GetInitialTimeout()), WriteComplex[BACnetContextTagUnsignedInteger](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'initialTimeout' field")
		}

		if err := WriteOptionalField[BACnetDateTimeEnclosed](ctx, "expirationTime", GetRef(m.GetExpirationTime()), WriteComplex[BACnetDateTimeEnclosed](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'expirationTime' field")
		}

		if err := WriteSimpleField[BACnetClosingTag](ctx, "innerClosingTag", m.GetInnerClosingTag(), WriteComplex[BACnetClosingTag](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersChangeOfTimer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersChangeOfTimer")
		}
		return nil
	}
	return m.BACnetNotificationParametersContract.(*_BACnetNotificationParameters).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersChangeOfTimer) IsBACnetNotificationParametersChangeOfTimer() {}

func (m *_BACnetNotificationParametersChangeOfTimer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
