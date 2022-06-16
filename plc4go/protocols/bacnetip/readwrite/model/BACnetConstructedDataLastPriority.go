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

// BACnetConstructedDataLastPriority is the corresponding interface of BACnetConstructedDataLastPriority
type BACnetConstructedDataLastPriority interface {
	BACnetConstructedData
	// GetLastPriority returns LastPriority (property field)
	GetLastPriority() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetConstructedDataLastPriority is the data-structure of this message
type _BACnetConstructedDataLastPriority struct {
	*_BACnetConstructedData
	LastPriority BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastPriority) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLastPriority) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LAST_PRIORITY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastPriority) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLastPriority) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastPriority) GetLastPriority() BACnetApplicationTagUnsignedInteger {
	return m.LastPriority
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastPriority) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetLastPriority())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLastPriority factory function for _BACnetConstructedDataLastPriority
func NewBACnetConstructedDataLastPriority(lastPriority BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLastPriority {
	_result := &_BACnetConstructedDataLastPriority{
		LastPriority:           lastPriority,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastPriority(structType interface{}) BACnetConstructedDataLastPriority {
	if casted, ok := structType.(BACnetConstructedDataLastPriority); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastPriority); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastPriority) GetTypeName() string {
	return "BACnetConstructedDataLastPriority"
}

func (m *_BACnetConstructedDataLastPriority) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLastPriority) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lastPriority)
	lengthInBits += m.LastPriority.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLastPriority) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLastPriorityParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastPriority, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastPriority")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lastPriority)
	if pullErr := readBuffer.PullContext("lastPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lastPriority")
	}
	_lastPriority, _lastPriorityErr := BACnetApplicationTagParse(readBuffer)
	if _lastPriorityErr != nil {
		return nil, errors.Wrap(_lastPriorityErr, "Error parsing 'lastPriority' field")
	}
	lastPriority := _lastPriority.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("lastPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lastPriority")
	}

	// Virtual field
	_actualValue := lastPriority
	actualValue := _actualValue.(BACnetApplicationTagUnsignedInteger)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastPriority")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLastPriority{
		LastPriority:           lastPriority,
		_BACnetConstructedData: &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLastPriority) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastPriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastPriority")
		}

		// Simple Field (lastPriority)
		if pushErr := writeBuffer.PushContext("lastPriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lastPriority")
		}
		_lastPriorityErr := writeBuffer.WriteSerializable(m.GetLastPriority())
		if popErr := writeBuffer.PopContext("lastPriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lastPriority")
		}
		if _lastPriorityErr != nil {
			return errors.Wrap(_lastPriorityErr, "Error serializing 'lastPriority' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastPriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastPriority")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLastPriority) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
