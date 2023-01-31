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


// BACnetConstructedDataMaximumValue is the corresponding interface of BACnetConstructedDataMaximumValue
type BACnetConstructedDataMaximumValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaximumValue returns MaximumValue (property field)
	GetMaximumValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataMaximumValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMaximumValue.
// This is useful for switch cases.
type BACnetConstructedDataMaximumValueExactly interface {
	BACnetConstructedDataMaximumValue
	isBACnetConstructedDataMaximumValue() bool
}

// _BACnetConstructedDataMaximumValue is the data-structure of this message
type _BACnetConstructedDataMaximumValue struct {
	*_BACnetConstructedData
        MaximumValue BACnetApplicationTagReal
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaximumValue)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataMaximumValue)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_MAXIMUM_VALUE}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaximumValue) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMaximumValue)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaximumValue) GetMaximumValue() BACnetApplicationTagReal {
	return m.MaximumValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMaximumValue) GetActualValue() BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetMaximumValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataMaximumValue factory function for _BACnetConstructedDataMaximumValue
func NewBACnetConstructedDataMaximumValue( maximumValue BACnetApplicationTagReal , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataMaximumValue {
	_result := &_BACnetConstructedDataMaximumValue{
		MaximumValue: maximumValue,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaximumValue(structType interface{}) BACnetConstructedDataMaximumValue {
    if casted, ok := structType.(BACnetConstructedDataMaximumValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaximumValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaximumValue) GetTypeName() string {
	return "BACnetConstructedDataMaximumValue"
}

func (m *_BACnetConstructedDataMaximumValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataMaximumValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maximumValue)
	lengthInBits += m.MaximumValue.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataMaximumValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMaximumValueParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMaximumValue, error) {
	return BACnetConstructedDataMaximumValueParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataMaximumValueParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMaximumValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaximumValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaximumValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maximumValue)
	if pullErr := readBuffer.PullContext("maximumValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maximumValue")
	}
_maximumValue, _maximumValueErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _maximumValueErr != nil {
		return nil, errors.Wrap(_maximumValueErr, "Error parsing 'maximumValue' field of BACnetConstructedDataMaximumValue")
	}
	maximumValue := _maximumValue.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("maximumValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maximumValue")
	}

	// Virtual field
	_actualValue := maximumValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaximumValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaximumValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMaximumValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MaximumValue: maximumValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMaximumValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMaximumValue) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaximumValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaximumValue")
		}

	// Simple Field (maximumValue)
	if pushErr := writeBuffer.PushContext("maximumValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for maximumValue")
	}
	_maximumValueErr := writeBuffer.WriteSerializable(m.GetMaximumValue())
	if popErr := writeBuffer.PopContext("maximumValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for maximumValue")
	}
	if _maximumValueErr != nil {
		return errors.Wrap(_maximumValueErr, "Error serializing 'maximumValue' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaximumValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaximumValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataMaximumValue) isBACnetConstructedDataMaximumValue() bool {
	return true
}

func (m *_BACnetConstructedDataMaximumValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



