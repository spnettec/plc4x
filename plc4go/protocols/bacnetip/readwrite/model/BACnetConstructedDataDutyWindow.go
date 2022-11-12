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


// BACnetConstructedDataDutyWindow is the corresponding interface of BACnetConstructedDataDutyWindow
type BACnetConstructedDataDutyWindow interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDutyWindow returns DutyWindow (property field)
	GetDutyWindow() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataDutyWindowExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDutyWindow.
// This is useful for switch cases.
type BACnetConstructedDataDutyWindowExactly interface {
	BACnetConstructedDataDutyWindow
	isBACnetConstructedDataDutyWindow() bool
}

// _BACnetConstructedDataDutyWindow is the data-structure of this message
type _BACnetConstructedDataDutyWindow struct {
	*_BACnetConstructedData
        DutyWindow BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDutyWindow)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataDutyWindow)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_DUTY_WINDOW}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDutyWindow) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDutyWindow)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDutyWindow) GetDutyWindow() BACnetApplicationTagUnsignedInteger {
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

func (m *_BACnetConstructedDataDutyWindow) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetDutyWindow())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataDutyWindow factory function for _BACnetConstructedDataDutyWindow
func NewBACnetConstructedDataDutyWindow( dutyWindow BACnetApplicationTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataDutyWindow {
	_result := &_BACnetConstructedDataDutyWindow{
		DutyWindow: dutyWindow,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDutyWindow(structType interface{}) BACnetConstructedDataDutyWindow {
    if casted, ok := structType.(BACnetConstructedDataDutyWindow); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDutyWindow); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDutyWindow) GetTypeName() string {
	return "BACnetConstructedDataDutyWindow"
}

func (m *_BACnetConstructedDataDutyWindow) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDutyWindow) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (dutyWindow)
	lengthInBits += m.DutyWindow.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataDutyWindow) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDutyWindowParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDutyWindow, error) {
	return BACnetConstructedDataDutyWindowParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataDutyWindowParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDutyWindow, error) {
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
_dutyWindow, _dutyWindowErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _dutyWindowErr != nil {
		return nil, errors.Wrap(_dutyWindowErr, "Error parsing 'dutyWindow' field of BACnetConstructedDataDutyWindow")
	}
	dutyWindow := _dutyWindow.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("dutyWindow"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dutyWindow")
	}

	// Virtual field
	_actualValue := dutyWindow
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDutyWindow"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDutyWindow")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDutyWindow{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DutyWindow: dutyWindow,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDutyWindow) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDutyWindow) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
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
	_dutyWindowErr := writeBuffer.WriteSerializable(m.GetDutyWindow())
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


func (m *_BACnetConstructedDataDutyWindow) isBACnetConstructedDataDutyWindow() bool {
	return true
}

func (m *_BACnetConstructedDataDutyWindow) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



