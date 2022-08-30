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


// BACnetConstructedDataDerivativeConstant is the corresponding interface of BACnetConstructedDataDerivativeConstant
type BACnetConstructedDataDerivativeConstant interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetDerivativeConstant returns DerivativeConstant (property field)
	GetDerivativeConstant() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataDerivativeConstantExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDerivativeConstant.
// This is useful for switch cases.
type BACnetConstructedDataDerivativeConstantExactly interface {
	BACnetConstructedDataDerivativeConstant
	isBACnetConstructedDataDerivativeConstant() bool
}

// _BACnetConstructedDataDerivativeConstant is the data-structure of this message
type _BACnetConstructedDataDerivativeConstant struct {
	*_BACnetConstructedData
        DerivativeConstant BACnetApplicationTagReal
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDerivativeConstant)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataDerivativeConstant)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_DERIVATIVE_CONSTANT}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDerivativeConstant) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDerivativeConstant)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDerivativeConstant) GetDerivativeConstant() BACnetApplicationTagReal {
	return m.DerivativeConstant
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDerivativeConstant) GetActualValue() BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetDerivativeConstant())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataDerivativeConstant factory function for _BACnetConstructedDataDerivativeConstant
func NewBACnetConstructedDataDerivativeConstant( derivativeConstant BACnetApplicationTagReal , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataDerivativeConstant {
	_result := &_BACnetConstructedDataDerivativeConstant{
		DerivativeConstant: derivativeConstant,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDerivativeConstant(structType interface{}) BACnetConstructedDataDerivativeConstant {
    if casted, ok := structType.(BACnetConstructedDataDerivativeConstant); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDerivativeConstant); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDerivativeConstant) GetTypeName() string {
	return "BACnetConstructedDataDerivativeConstant"
}

func (m *_BACnetConstructedDataDerivativeConstant) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDerivativeConstant) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (derivativeConstant)
	lengthInBits += m.DerivativeConstant.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataDerivativeConstant) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDerivativeConstantParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDerivativeConstant, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDerivativeConstant"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDerivativeConstant")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (derivativeConstant)
	if pullErr := readBuffer.PullContext("derivativeConstant"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for derivativeConstant")
	}
_derivativeConstant, _derivativeConstantErr := BACnetApplicationTagParse(readBuffer)
	if _derivativeConstantErr != nil {
		return nil, errors.Wrap(_derivativeConstantErr, "Error parsing 'derivativeConstant' field of BACnetConstructedDataDerivativeConstant")
	}
	derivativeConstant := _derivativeConstant.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("derivativeConstant"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for derivativeConstant")
	}

	// Virtual field
	_actualValue := derivativeConstant
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDerivativeConstant"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDerivativeConstant")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDerivativeConstant{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		DerivativeConstant: derivativeConstant,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDerivativeConstant) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDerivativeConstant"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDerivativeConstant")
		}

	// Simple Field (derivativeConstant)
	if pushErr := writeBuffer.PushContext("derivativeConstant"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for derivativeConstant")
	}
	_derivativeConstantErr := writeBuffer.WriteSerializable(m.GetDerivativeConstant())
	if popErr := writeBuffer.PopContext("derivativeConstant"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for derivativeConstant")
	}
	if _derivativeConstantErr != nil {
		return errors.Wrap(_derivativeConstantErr, "Error serializing 'derivativeConstant' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDerivativeConstant"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDerivativeConstant")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataDerivativeConstant) isBACnetConstructedDataDerivativeConstant() bool {
	return true
}

func (m *_BACnetConstructedDataDerivativeConstant) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



