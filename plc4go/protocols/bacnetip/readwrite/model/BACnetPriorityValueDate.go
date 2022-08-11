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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPriorityValueDate is the corresponding interface of BACnetPriorityValueDate
type BACnetPriorityValueDate interface {
	utils.LengthAware
	utils.Serializable
	BACnetPriorityValue
	// GetDateValue returns DateValue (property field)
	GetDateValue() BACnetApplicationTagDate
}

// BACnetPriorityValueDateExactly can be used when we want exactly this type and not a type which fulfills BACnetPriorityValueDate.
// This is useful for switch cases.
type BACnetPriorityValueDateExactly interface {
	BACnetPriorityValueDate
	isBACnetPriorityValueDate() bool
}

// _BACnetPriorityValueDate is the data-structure of this message
type _BACnetPriorityValueDate struct {
	*_BACnetPriorityValue
	DateValue BACnetApplicationTagDate
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPriorityValueDate) InitializeParent(parent BACnetPriorityValue, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPriorityValueDate) GetParent() BACnetPriorityValue {
	return m._BACnetPriorityValue
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueDate) GetDateValue() BACnetApplicationTagDate {
	return m.DateValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPriorityValueDate factory function for _BACnetPriorityValueDate
func NewBACnetPriorityValueDate(dateValue BACnetApplicationTagDate, peekedTagHeader BACnetTagHeader, objectTypeArgument BACnetObjectType) *_BACnetPriorityValueDate {
	_result := &_BACnetPriorityValueDate{
		DateValue:            dateValue,
		_BACnetPriorityValue: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueDate(structType interface{}) BACnetPriorityValueDate {
	if casted, ok := structType.(BACnetPriorityValueDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueDate) GetTypeName() string {
	return "BACnetPriorityValueDate"
}

func (m *_BACnetPriorityValueDate) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPriorityValueDate) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dateValue)
	lengthInBits += m.DateValue.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPriorityValueDate) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPriorityValueDateParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetPriorityValueDate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dateValue)
	if pullErr := readBuffer.PullContext("dateValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dateValue")
	}
	_dateValue, _dateValueErr := BACnetApplicationTagParse(readBuffer)
	if _dateValueErr != nil {
		return nil, errors.Wrap(_dateValueErr, "Error parsing 'dateValue' field of BACnetPriorityValueDate")
	}
	dateValue := _dateValue.(BACnetApplicationTagDate)
	if closeErr := readBuffer.CloseContext("dateValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dateValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueDate")
	}

	// Create a partially initialized instance
	_child := &_BACnetPriorityValueDate{
		_BACnetPriorityValue: &_BACnetPriorityValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		DateValue: dateValue,
	}
	_child._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPriorityValueDate) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueDate")
		}

		// Simple Field (dateValue)
		if pushErr := writeBuffer.PushContext("dateValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dateValue")
		}
		_dateValueErr := writeBuffer.WriteSerializable(m.GetDateValue())
		if popErr := writeBuffer.PopContext("dateValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dateValue")
		}
		if _dateValueErr != nil {
			return errors.Wrap(_dateValueErr, "Error serializing 'dateValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueDate")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPriorityValueDate) isBACnetPriorityValueDate() bool {
	return true
}

func (m *_BACnetPriorityValueDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
