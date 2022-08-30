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


// BACnetConstructedDataIntegralConstant is the corresponding interface of BACnetConstructedDataIntegralConstant
type BACnetConstructedDataIntegralConstant interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIntegralConstant returns IntegralConstant (property field)
	GetIntegralConstant() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataIntegralConstantExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIntegralConstant.
// This is useful for switch cases.
type BACnetConstructedDataIntegralConstantExactly interface {
	BACnetConstructedDataIntegralConstant
	isBACnetConstructedDataIntegralConstant() bool
}

// _BACnetConstructedDataIntegralConstant is the data-structure of this message
type _BACnetConstructedDataIntegralConstant struct {
	*_BACnetConstructedData
        IntegralConstant BACnetApplicationTagReal
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIntegralConstant)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataIntegralConstant)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_INTEGRAL_CONSTANT}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIntegralConstant) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIntegralConstant)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIntegralConstant) GetIntegralConstant() BACnetApplicationTagReal {
	return m.IntegralConstant
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIntegralConstant) GetActualValue() BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetIntegralConstant())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataIntegralConstant factory function for _BACnetConstructedDataIntegralConstant
func NewBACnetConstructedDataIntegralConstant( integralConstant BACnetApplicationTagReal , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataIntegralConstant {
	_result := &_BACnetConstructedDataIntegralConstant{
		IntegralConstant: integralConstant,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIntegralConstant(structType interface{}) BACnetConstructedDataIntegralConstant {
    if casted, ok := structType.(BACnetConstructedDataIntegralConstant); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIntegralConstant); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIntegralConstant) GetTypeName() string {
	return "BACnetConstructedDataIntegralConstant"
}

func (m *_BACnetConstructedDataIntegralConstant) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataIntegralConstant) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (integralConstant)
	lengthInBits += m.IntegralConstant.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataIntegralConstant) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataIntegralConstantParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIntegralConstant, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIntegralConstant"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIntegralConstant")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (integralConstant)
	if pullErr := readBuffer.PullContext("integralConstant"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for integralConstant")
	}
_integralConstant, _integralConstantErr := BACnetApplicationTagParse(readBuffer)
	if _integralConstantErr != nil {
		return nil, errors.Wrap(_integralConstantErr, "Error parsing 'integralConstant' field of BACnetConstructedDataIntegralConstant")
	}
	integralConstant := _integralConstant.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("integralConstant"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for integralConstant")
	}

	// Virtual field
	_actualValue := integralConstant
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIntegralConstant"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIntegralConstant")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIntegralConstant{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		IntegralConstant: integralConstant,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIntegralConstant) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIntegralConstant"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIntegralConstant")
		}

	// Simple Field (integralConstant)
	if pushErr := writeBuffer.PushContext("integralConstant"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for integralConstant")
	}
	_integralConstantErr := writeBuffer.WriteSerializable(m.GetIntegralConstant())
	if popErr := writeBuffer.PopContext("integralConstant"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for integralConstant")
	}
	if _integralConstantErr != nil {
		return errors.Wrap(_integralConstantErr, "Error serializing 'integralConstant' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIntegralConstant"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIntegralConstant")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataIntegralConstant) isBACnetConstructedDataIntegralConstant() bool {
	return true
}

func (m *_BACnetConstructedDataIntegralConstant) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



