/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetEventParameterOutOfRange is the data-structure of this message
type BACnetEventParameterOutOfRange struct {
	*BACnetEventParameter
	OpeningTag    *BACnetOpeningTag
	TimeDelay     *BACnetContextTagUnsignedInteger
	LowDiffLimit  *BACnetContextTagReal
	HighDiffLimit *BACnetContextTagReal
	Deadband      *BACnetContextTagReal
	ClosingTag    *BACnetClosingTag
}

// IBACnetEventParameterOutOfRange is the corresponding interface of BACnetEventParameterOutOfRange
type IBACnetEventParameterOutOfRange interface {
	IBACnetEventParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetTimeDelay returns TimeDelay (property field)
	GetTimeDelay() *BACnetContextTagUnsignedInteger
	// GetLowDiffLimit returns LowDiffLimit (property field)
	GetLowDiffLimit() *BACnetContextTagReal
	// GetHighDiffLimit returns HighDiffLimit (property field)
	GetHighDiffLimit() *BACnetContextTagReal
	// GetDeadband returns Deadband (property field)
	GetDeadband() *BACnetContextTagReal
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetEventParameterOutOfRange) InitializeParent(parent *BACnetEventParameter, peekedTagHeader *BACnetTagHeader) {
	m.BACnetEventParameter.PeekedTagHeader = peekedTagHeader
}

func (m *BACnetEventParameterOutOfRange) GetParent() *BACnetEventParameter {
	return m.BACnetEventParameter
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetEventParameterOutOfRange) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetEventParameterOutOfRange) GetTimeDelay() *BACnetContextTagUnsignedInteger {
	return m.TimeDelay
}

func (m *BACnetEventParameterOutOfRange) GetLowDiffLimit() *BACnetContextTagReal {
	return m.LowDiffLimit
}

func (m *BACnetEventParameterOutOfRange) GetHighDiffLimit() *BACnetContextTagReal {
	return m.HighDiffLimit
}

func (m *BACnetEventParameterOutOfRange) GetDeadband() *BACnetContextTagReal {
	return m.Deadband
}

func (m *BACnetEventParameterOutOfRange) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterOutOfRange factory function for BACnetEventParameterOutOfRange
func NewBACnetEventParameterOutOfRange(openingTag *BACnetOpeningTag, timeDelay *BACnetContextTagUnsignedInteger, lowDiffLimit *BACnetContextTagReal, highDiffLimit *BACnetContextTagReal, deadband *BACnetContextTagReal, closingTag *BACnetClosingTag, peekedTagHeader *BACnetTagHeader) *BACnetEventParameterOutOfRange {
	_result := &BACnetEventParameterOutOfRange{
		OpeningTag:           openingTag,
		TimeDelay:            timeDelay,
		LowDiffLimit:         lowDiffLimit,
		HighDiffLimit:        highDiffLimit,
		Deadband:             deadband,
		ClosingTag:           closingTag,
		BACnetEventParameter: NewBACnetEventParameter(peekedTagHeader),
	}
	_result.Child = _result
	return _result
}

func CastBACnetEventParameterOutOfRange(structType interface{}) *BACnetEventParameterOutOfRange {
	if casted, ok := structType.(BACnetEventParameterOutOfRange); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventParameterOutOfRange); ok {
		return casted
	}
	if casted, ok := structType.(BACnetEventParameter); ok {
		return CastBACnetEventParameterOutOfRange(casted.Child)
	}
	if casted, ok := structType.(*BACnetEventParameter); ok {
		return CastBACnetEventParameterOutOfRange(casted.Child)
	}
	return nil
}

func (m *BACnetEventParameterOutOfRange) GetTypeName() string {
	return "BACnetEventParameterOutOfRange"
}

func (m *BACnetEventParameterOutOfRange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventParameterOutOfRange) GetLengthInBitsConditional(lastItem bool) uint16 {
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

func (m *BACnetEventParameterOutOfRange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterOutOfRangeParse(readBuffer utils.ReadBuffer) (*BACnetEventParameterOutOfRange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterOutOfRange"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(5)))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (timeDelay)
	if pullErr := readBuffer.PullContext("timeDelay"); pullErr != nil {
		return nil, pullErr
	}
	_timeDelay, _timeDelayErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _timeDelayErr != nil {
		return nil, errors.Wrap(_timeDelayErr, "Error parsing 'timeDelay' field")
	}
	timeDelay := CastBACnetContextTagUnsignedInteger(_timeDelay)
	if closeErr := readBuffer.CloseContext("timeDelay"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (lowDiffLimit)
	if pullErr := readBuffer.PullContext("lowDiffLimit"); pullErr != nil {
		return nil, pullErr
	}
	_lowDiffLimit, _lowDiffLimitErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_REAL))
	if _lowDiffLimitErr != nil {
		return nil, errors.Wrap(_lowDiffLimitErr, "Error parsing 'lowDiffLimit' field")
	}
	lowDiffLimit := CastBACnetContextTagReal(_lowDiffLimit)
	if closeErr := readBuffer.CloseContext("lowDiffLimit"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (highDiffLimit)
	if pullErr := readBuffer.PullContext("highDiffLimit"); pullErr != nil {
		return nil, pullErr
	}
	_highDiffLimit, _highDiffLimitErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_REAL))
	if _highDiffLimitErr != nil {
		return nil, errors.Wrap(_highDiffLimitErr, "Error parsing 'highDiffLimit' field")
	}
	highDiffLimit := CastBACnetContextTagReal(_highDiffLimit)
	if closeErr := readBuffer.CloseContext("highDiffLimit"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (deadband)
	if pullErr := readBuffer.PullContext("deadband"); pullErr != nil {
		return nil, pullErr
	}
	_deadband, _deadbandErr := BACnetContextTagParse(readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_REAL))
	if _deadbandErr != nil {
		return nil, errors.Wrap(_deadbandErr, "Error parsing 'deadband' field")
	}
	deadband := CastBACnetContextTagReal(_deadband)
	if closeErr := readBuffer.CloseContext("deadband"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(uint8(5)))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterOutOfRange"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetEventParameterOutOfRange{
		OpeningTag:           CastBACnetOpeningTag(openingTag),
		TimeDelay:            CastBACnetContextTagUnsignedInteger(timeDelay),
		LowDiffLimit:         CastBACnetContextTagReal(lowDiffLimit),
		HighDiffLimit:        CastBACnetContextTagReal(highDiffLimit),
		Deadband:             CastBACnetContextTagReal(deadband),
		ClosingTag:           CastBACnetClosingTag(closingTag),
		BACnetEventParameter: &BACnetEventParameter{},
	}
	_child.BACnetEventParameter.Child = _child
	return _child, nil
}

func (m *BACnetEventParameterOutOfRange) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterOutOfRange"); pushErr != nil {
			return pushErr
		}

		// Simple Field (openingTag)
		if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
			return pushErr
		}
		_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
			return popErr
		}
		if _openingTagErr != nil {
			return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
		}

		// Simple Field (timeDelay)
		if pushErr := writeBuffer.PushContext("timeDelay"); pushErr != nil {
			return pushErr
		}
		_timeDelayErr := m.TimeDelay.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("timeDelay"); popErr != nil {
			return popErr
		}
		if _timeDelayErr != nil {
			return errors.Wrap(_timeDelayErr, "Error serializing 'timeDelay' field")
		}

		// Simple Field (lowDiffLimit)
		if pushErr := writeBuffer.PushContext("lowDiffLimit"); pushErr != nil {
			return pushErr
		}
		_lowDiffLimitErr := m.LowDiffLimit.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("lowDiffLimit"); popErr != nil {
			return popErr
		}
		if _lowDiffLimitErr != nil {
			return errors.Wrap(_lowDiffLimitErr, "Error serializing 'lowDiffLimit' field")
		}

		// Simple Field (highDiffLimit)
		if pushErr := writeBuffer.PushContext("highDiffLimit"); pushErr != nil {
			return pushErr
		}
		_highDiffLimitErr := m.HighDiffLimit.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("highDiffLimit"); popErr != nil {
			return popErr
		}
		if _highDiffLimitErr != nil {
			return errors.Wrap(_highDiffLimitErr, "Error serializing 'highDiffLimit' field")
		}

		// Simple Field (deadband)
		if pushErr := writeBuffer.PushContext("deadband"); pushErr != nil {
			return pushErr
		}
		_deadbandErr := m.Deadband.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("deadband"); popErr != nil {
			return popErr
		}
		if _deadbandErr != nil {
			return errors.Wrap(_deadbandErr, "Error serializing 'deadband' field")
		}

		// Simple Field (closingTag)
		if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
			return pushErr
		}
		_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
			return popErr
		}
		if _closingTagErr != nil {
			return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventParameterOutOfRange"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetEventParameterOutOfRange) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
