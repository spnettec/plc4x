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


// BACnetConstructedDataValueSource is the corresponding interface of BACnetConstructedDataValueSource
type BACnetConstructedDataValueSource interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetValueSource returns ValueSource (property field)
	GetValueSource() BACnetValueSource
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetValueSource
}

// BACnetConstructedDataValueSourceExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataValueSource.
// This is useful for switch cases.
type BACnetConstructedDataValueSourceExactly interface {
	BACnetConstructedDataValueSource
	isBACnetConstructedDataValueSource() bool
}

// _BACnetConstructedDataValueSource is the data-structure of this message
type _BACnetConstructedDataValueSource struct {
	*_BACnetConstructedData
        ValueSource BACnetValueSource
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataValueSource)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataValueSource)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_VALUE_SOURCE}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataValueSource) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataValueSource)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataValueSource) GetValueSource() BACnetValueSource {
	return m.ValueSource
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataValueSource) GetActualValue() BACnetValueSource {
	return CastBACnetValueSource(m.GetValueSource())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataValueSource factory function for _BACnetConstructedDataValueSource
func NewBACnetConstructedDataValueSource( valueSource BACnetValueSource , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataValueSource {
	_result := &_BACnetConstructedDataValueSource{
		ValueSource: valueSource,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataValueSource(structType interface{}) BACnetConstructedDataValueSource {
    if casted, ok := structType.(BACnetConstructedDataValueSource); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataValueSource); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataValueSource) GetTypeName() string {
	return "BACnetConstructedDataValueSource"
}

func (m *_BACnetConstructedDataValueSource) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataValueSource) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (valueSource)
	lengthInBits += m.ValueSource.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataValueSource) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataValueSourceParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataValueSource, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataValueSource"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataValueSource")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (valueSource)
	if pullErr := readBuffer.PullContext("valueSource"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for valueSource")
	}
_valueSource, _valueSourceErr := BACnetValueSourceParse(readBuffer)
	if _valueSourceErr != nil {
		return nil, errors.Wrap(_valueSourceErr, "Error parsing 'valueSource' field of BACnetConstructedDataValueSource")
	}
	valueSource := _valueSource.(BACnetValueSource)
	if closeErr := readBuffer.CloseContext("valueSource"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for valueSource")
	}

	// Virtual field
	_actualValue := valueSource
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataValueSource"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataValueSource")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataValueSource{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ValueSource: valueSource,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataValueSource) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataValueSource"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataValueSource")
		}

	// Simple Field (valueSource)
	if pushErr := writeBuffer.PushContext("valueSource"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for valueSource")
	}
	_valueSourceErr := writeBuffer.WriteSerializable(m.GetValueSource())
	if popErr := writeBuffer.PopContext("valueSource"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for valueSource")
	}
	if _valueSourceErr != nil {
		return errors.Wrap(_valueSourceErr, "Error serializing 'valueSource' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataValueSource"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataValueSource")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataValueSource) isBACnetConstructedDataValueSource() bool {
	return true
}

func (m *_BACnetConstructedDataValueSource) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



