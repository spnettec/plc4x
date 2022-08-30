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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetEventParameterOutOfRange is the corresponding interface of BACnetEventParameterOutOfRange
type BACnetEventParameterOutOfRange interface {
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() BACnetContextTagUnsignedInteger
	// GetLowDiffLimit returns LowDiffLimit (property field)
	GetLowDiffLimit() BACnetContextTagReal
	// GetHighDiffLimit returns HighDiffLimit (property field)
	GetHighDiffLimit() BACnetContextTagReal
	// GetDeadband returns Deadband (property field)
	GetDeadband() BACnetContextTagReal
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterOutOfRangeExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterOutOfRange.
// This is useful for switch cases.
type BACnetEventParameterOutOfRangeExactly interface {
	BACnetEventParameterOutOfRange
	isBACnetEventParameterOutOfRange() bool
}

// _BACnetEventParameterOutOfRange is the data-structure of this message
type _BACnetEventParameterOutOfRange struct {
	*_BACnetEventParameter
        OpeningTag BACnetOpeningTag
        TimeDelay BACnetContextTagUnsignedInteger
        LowDiffLimit BACnetContextTagReal
        HighDiffLimit BACnetContextTagReal
        Deadband BACnetContextTagReal
        ClosingTag BACnetClosingTag
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterOutOfRange) InitializeParent(parent BACnetEventParameter , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetEventParameterOutOfRange)  GetParent() BACnetEventParameter {
	return m._BACnetEventParameter
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterOutOfRange) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterOutOfRange) GetTimeDelay() BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *_BACnetEventParameterOutOfRange) GetLowDiffLimit() BACnetContextTagReal {
	return m.LowDiffLimit
}

func (m *_BACnetEventParameterOutOfRange) GetHighDiffLimit() BACnetContextTagReal {
	return m.HighDiffLimit
}

func (m *_BACnetEventParameterOutOfRange) GetDeadband() BACnetContextTagReal {
	return m.Deadband
}

func (m *_BACnetEventParameterOutOfRange) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetEventParameterOutOfRange factory function for _BACnetEventParameterOutOfRange
func NewBACnetEventParameterOutOfRange( openingTag BACnetOpeningTag , timeDelay BACnetContextTagUnsignedInteger , lowDiffLimit BACnetContextTagReal , highDiffLimit BACnetContextTagReal , deadband BACnetContextTagReal , closingTag BACnetClosingTag , peekedTagHeader BACnetTagHeader ) *_BACnetEventParameterOutOfRange {
	_result := &_BACnetEventParameterOutOfRange{
		OpeningTag: openingTag,
		TimeDelay: timeDelay,
		LowDiffLimit: lowDiffLimit,
		HighDiffLimit: highDiffLimit,
		Deadband: deadband,
		ClosingTag: closingTag,
    	_BACnetEventParameter: NewBACnetEventParameter(peekedTagHeader),
	}
	_result._BACnetEventParameter._BACnetEventParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterOutOfRange(structType interface{}) BACnetEventParameterOutOfRange {
    if casted, ok := structType.(BACnetEventParameterOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterOutOfRange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterOutOfRange) GetTypeName() string {
	return "BACnetEventParameterOutOfRange"
}

func (m *_BACnetEventParameterOutOfRange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventParameterOutOfRange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (timeDelay)
	lengthInBits += m.TimeDelay.GetLengthInBits()

	// Simple field (lowDiffLimit)
	lengthInBits += m.LowDiffLimit.GetLengthInBits()

	// Simple field (highDiffLimit)
	lengthInBits += m.HighDiffLimit.GetLengthInBits()

	// Simple field (deadband)
	lengthInBits += m.Deadband.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetEventParameterOutOfRange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterOutOfRangeParse(readBuffer utils.ReadBuffer) (BACnetEventParameterOutOfRange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterOutOfRange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterOutOfRange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer , uint8( uint8(5) ) )
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventParameterOutOfRange")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (timeDelay)
	if pullErr := readBuffer.PullContext("timeDelay"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeDelay")
	}
_timeDelay, _timeDelayErr := BACnetContextTagParse(readBuffer , uint8( uint8(0) ) , BACnetDataType( BACnetDataType_UNSIGNED_INTEGER ) )
	if _timeDelayErr != nil {
		return nil, errors.Wrap(_timeDelayErr, "Error parsing 'timeDelay' field of BACnetEventParameterOutOfRange")
	}
	timeDelay := _timeDelay.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("timeDelay"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeDelay")
	}

	// Simple Field (lowDiffLimit)
	if pullErr := readBuffer.PullContext("lowDiffLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lowDiffLimit")
	}
_lowDiffLimit, _lowDiffLimitErr := BACnetContextTagParse(readBuffer , uint8( uint8(1) ) , BACnetDataType( BACnetDataType_REAL ) )
	if _lowDiffLimitErr != nil {
		return nil, errors.Wrap(_lowDiffLimitErr, "Error parsing 'lowDiffLimit' field of BACnetEventParameterOutOfRange")
	}
	lowDiffLimit := _lowDiffLimit.(BACnetContextTagReal)
	if closeErr := readBuffer.CloseContext("lowDiffLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lowDiffLimit")
	}

	// Simple Field (highDiffLimit)
	if pullErr := readBuffer.PullContext("highDiffLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for highDiffLimit")
	}
_highDiffLimit, _highDiffLimitErr := BACnetContextTagParse(readBuffer , uint8( uint8(2) ) , BACnetDataType( BACnetDataType_REAL ) )
	if _highDiffLimitErr != nil {
		return nil, errors.Wrap(_highDiffLimitErr, "Error parsing 'highDiffLimit' field of BACnetEventParameterOutOfRange")
	}
	highDiffLimit := _highDiffLimit.(BACnetContextTagReal)
	if closeErr := readBuffer.CloseContext("highDiffLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for highDiffLimit")
	}

	// Simple Field (deadband)
	if pullErr := readBuffer.PullContext("deadband"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for deadband")
	}
_deadband, _deadbandErr := BACnetContextTagParse(readBuffer , uint8( uint8(3) ) , BACnetDataType( BACnetDataType_REAL ) )
	if _deadbandErr != nil {
		return nil, errors.Wrap(_deadbandErr, "Error parsing 'deadband' field of BACnetEventParameterOutOfRange")
	}
	deadband := _deadband.(BACnetContextTagReal)
	if closeErr := readBuffer.CloseContext("deadband"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for deadband")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer , uint8( uint8(5) ) )
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventParameterOutOfRange")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterOutOfRange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterOutOfRange")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterOutOfRange{
		_BACnetEventParameter: &_BACnetEventParameter{
		},
		OpeningTag: openingTag,
		TimeDelay: timeDelay,
		LowDiffLimit: lowDiffLimit,
		HighDiffLimit: highDiffLimit,
		Deadband: deadband,
		ClosingTag: closingTag,
	}
	_child._BACnetEventParameter._BACnetEventParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterOutOfRange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterOutOfRange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterOutOfRange")
		}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (timeDelay)
	if pushErr := writeBuffer.PushContext("timeDelay"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timeDelay")
	}
	_timeDelayErr := writeBuffer.WriteSerializable(m.GetTimeDelay())
	if popErr := writeBuffer.PopContext("timeDelay"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timeDelay")
	}
	if _timeDelayErr != nil {
		return errors.Wrap(_timeDelayErr, "Error serializing 'timeDelay' field")
	}

	// Simple Field (lowDiffLimit)
	if pushErr := writeBuffer.PushContext("lowDiffLimit"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for lowDiffLimit")
	}
	_lowDiffLimitErr := writeBuffer.WriteSerializable(m.GetLowDiffLimit())
	if popErr := writeBuffer.PopContext("lowDiffLimit"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for lowDiffLimit")
	}
	if _lowDiffLimitErr != nil {
		return errors.Wrap(_lowDiffLimitErr, "Error serializing 'lowDiffLimit' field")
	}

	// Simple Field (highDiffLimit)
	if pushErr := writeBuffer.PushContext("highDiffLimit"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for highDiffLimit")
	}
	_highDiffLimitErr := writeBuffer.WriteSerializable(m.GetHighDiffLimit())
	if popErr := writeBuffer.PopContext("highDiffLimit"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for highDiffLimit")
	}
	if _highDiffLimitErr != nil {
		return errors.Wrap(_highDiffLimitErr, "Error serializing 'highDiffLimit' field")
	}

	// Simple Field (deadband)
	if pushErr := writeBuffer.PushContext("deadband"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for deadband")
	}
	_deadbandErr := writeBuffer.WriteSerializable(m.GetDeadband())
	if popErr := writeBuffer.PopContext("deadband"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for deadband")
	}
	if _deadbandErr != nil {
		return errors.Wrap(_deadbandErr, "Error serializing 'deadband' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

		if popErr := writeBuffer.PopContext("BACnetEventParameterOutOfRange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterOutOfRange")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetEventParameterOutOfRange) isBACnetEventParameterOutOfRange() bool {
	return true
}

func (m *_BACnetEventParameterOutOfRange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



