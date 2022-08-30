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


// BACnetConstructedDataModificationDate is the corresponding interface of BACnetConstructedDataModificationDate
type BACnetConstructedDataModificationDate interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetModificationDate returns ModificationDate (property field)
	GetModificationDate() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// BACnetConstructedDataModificationDateExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataModificationDate.
// This is useful for switch cases.
type BACnetConstructedDataModificationDateExactly interface {
	BACnetConstructedDataModificationDate
	isBACnetConstructedDataModificationDate() bool
}

// _BACnetConstructedDataModificationDate is the data-structure of this message
type _BACnetConstructedDataModificationDate struct {
	*_BACnetConstructedData
        ModificationDate BACnetDateTime
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataModificationDate)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataModificationDate)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_MODIFICATION_DATE}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataModificationDate) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataModificationDate)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataModificationDate) GetModificationDate() BACnetDateTime {
	return m.ModificationDate
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataModificationDate) GetActualValue() BACnetDateTime {
	return CastBACnetDateTime(m.GetModificationDate())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataModificationDate factory function for _BACnetConstructedDataModificationDate
func NewBACnetConstructedDataModificationDate( modificationDate BACnetDateTime , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataModificationDate {
	_result := &_BACnetConstructedDataModificationDate{
		ModificationDate: modificationDate,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataModificationDate(structType interface{}) BACnetConstructedDataModificationDate {
    if casted, ok := structType.(BACnetConstructedDataModificationDate); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataModificationDate); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataModificationDate) GetTypeName() string {
	return "BACnetConstructedDataModificationDate"
}

func (m *_BACnetConstructedDataModificationDate) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataModificationDate) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (modificationDate)
	lengthInBits += m.ModificationDate.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataModificationDate) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataModificationDateParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataModificationDate, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataModificationDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataModificationDate")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (modificationDate)
	if pullErr := readBuffer.PullContext("modificationDate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for modificationDate")
	}
_modificationDate, _modificationDateErr := BACnetDateTimeParse(readBuffer)
	if _modificationDateErr != nil {
		return nil, errors.Wrap(_modificationDateErr, "Error parsing 'modificationDate' field of BACnetConstructedDataModificationDate")
	}
	modificationDate := _modificationDate.(BACnetDateTime)
	if closeErr := readBuffer.CloseContext("modificationDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for modificationDate")
	}

	// Virtual field
	_actualValue := modificationDate
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataModificationDate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataModificationDate")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataModificationDate{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ModificationDate: modificationDate,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataModificationDate) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataModificationDate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataModificationDate")
		}

	// Simple Field (modificationDate)
	if pushErr := writeBuffer.PushContext("modificationDate"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for modificationDate")
	}
	_modificationDateErr := writeBuffer.WriteSerializable(m.GetModificationDate())
	if popErr := writeBuffer.PopContext("modificationDate"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for modificationDate")
	}
	if _modificationDateErr != nil {
		return errors.Wrap(_modificationDateErr, "Error serializing 'modificationDate' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataModificationDate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataModificationDate")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataModificationDate) isBACnetConstructedDataModificationDate() bool {
	return true
}

func (m *_BACnetConstructedDataModificationDate) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



