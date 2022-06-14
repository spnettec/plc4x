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

// BACnetConstructedDataDutyWindow is the data-structure of this message
type BACnetConstructedDataDutyWindow struct {
	*BACnetConstructedData
	DutyWindow *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataDutyWindow is the corresponding interface of BACnetConstructedDataDutyWindow
type IBACnetConstructedDataDutyWindow interface {
	IBACnetConstructedData
	// GetDutyWindow returns DutyWindow (property field)
	GetDutyWindow() *BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataDutyWindow) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataDutyWindow) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DUTY_WINDOW
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataDutyWindow) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataDutyWindow) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataDutyWindow) GetDutyWindow() *BACnetApplicationTagUnsignedInteger {
	return m.DutyWindow
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataDutyWindow) GetActualValue() *BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetDutyWindow())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDutyWindow factory function for BACnetConstructedDataDutyWindow
func NewBACnetConstructedDataDutyWindow(dutyWindow *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataDutyWindow {
	_result := &BACnetConstructedDataDutyWindow{
		DutyWindow:            dutyWindow,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataDutyWindow(structType interface{}) *BACnetConstructedDataDutyWindow {
	if casted, ok := structType.(BACnetConstructedDataDutyWindow); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDutyWindow); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataDutyWindow(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataDutyWindow(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataDutyWindow) GetTypeName() string {
	return "BACnetConstructedDataDutyWindow"
}

func (m *BACnetConstructedDataDutyWindow) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataDutyWindow) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dutyWindow)
	lengthInBits += m.DutyWindow.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataDutyWindow) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDutyWindowParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataDutyWindow, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDutyWindow"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDutyWindow")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dutyWindow)
	if pullErr := readBuffer.PullContext("dutyWindow"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dutyWindow")
	}
	_dutyWindow, _dutyWindowErr := BACnetApplicationTagParse(readBuffer)
	if _dutyWindowErr != nil {
		return nil, errors.Wrap(_dutyWindowErr, "Error parsing 'dutyWindow' field")
	}
	dutyWindow := CastBACnetApplicationTagUnsignedInteger(_dutyWindow)
	if closeErr := readBuffer.CloseContext("dutyWindow"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dutyWindow")
	}

	// Virtual field
	_actualValue := dutyWindow
	actualValue := CastBACnetApplicationTagUnsignedInteger(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDutyWindow"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDutyWindow")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataDutyWindow{
		DutyWindow:            CastBACnetApplicationTagUnsignedInteger(dutyWindow),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataDutyWindow) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDutyWindow"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDutyWindow")
		}

		// Simple Field (dutyWindow)
		if pushErr := writeBuffer.PushContext("dutyWindow"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for dutyWindow")
		}
		_dutyWindowErr := writeBuffer.WriteSerializable(m.DutyWindow)
		if popErr := writeBuffer.PopContext("dutyWindow"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for dutyWindow")
		}
		if _dutyWindowErr != nil {
			return errors.Wrap(_dutyWindowErr, "Error serializing 'dutyWindow' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDutyWindow"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDutyWindow")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataDutyWindow) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
