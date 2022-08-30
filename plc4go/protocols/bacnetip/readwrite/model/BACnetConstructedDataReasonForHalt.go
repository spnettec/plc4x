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


// BACnetConstructedDataReasonForHalt is the corresponding interface of BACnetConstructedDataReasonForHalt
type BACnetConstructedDataReasonForHalt interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetProgramError returns ProgramError (property field)
	GetProgramError() BACnetProgramErrorTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetProgramErrorTagged
}

// BACnetConstructedDataReasonForHaltExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataReasonForHalt.
// This is useful for switch cases.
type BACnetConstructedDataReasonForHaltExactly interface {
	BACnetConstructedDataReasonForHalt
	isBACnetConstructedDataReasonForHalt() bool
}

// _BACnetConstructedDataReasonForHalt is the data-structure of this message
type _BACnetConstructedDataReasonForHalt struct {
	*_BACnetConstructedData
        ProgramError BACnetProgramErrorTagged
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataReasonForHalt)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataReasonForHalt)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_REASON_FOR_HALT}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataReasonForHalt) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataReasonForHalt)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataReasonForHalt) GetProgramError() BACnetProgramErrorTagged {
	return m.ProgramError
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataReasonForHalt) GetActualValue() BACnetProgramErrorTagged {
	return CastBACnetProgramErrorTagged(m.GetProgramError())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataReasonForHalt factory function for _BACnetConstructedDataReasonForHalt
func NewBACnetConstructedDataReasonForHalt( programError BACnetProgramErrorTagged , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataReasonForHalt {
	_result := &_BACnetConstructedDataReasonForHalt{
		ProgramError: programError,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataReasonForHalt(structType interface{}) BACnetConstructedDataReasonForHalt {
    if casted, ok := structType.(BACnetConstructedDataReasonForHalt); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataReasonForHalt); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataReasonForHalt) GetTypeName() string {
	return "BACnetConstructedDataReasonForHalt"
}

func (m *_BACnetConstructedDataReasonForHalt) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataReasonForHalt) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (programError)
	lengthInBits += m.ProgramError.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataReasonForHalt) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataReasonForHaltParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataReasonForHalt, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataReasonForHalt"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataReasonForHalt")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (programError)
	if pullErr := readBuffer.PullContext("programError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for programError")
	}
_programError, _programErrorErr := BACnetProgramErrorTaggedParse(readBuffer , uint8( uint8(0) ) , TagClass( TagClass_APPLICATION_TAGS ) )
	if _programErrorErr != nil {
		return nil, errors.Wrap(_programErrorErr, "Error parsing 'programError' field of BACnetConstructedDataReasonForHalt")
	}
	programError := _programError.(BACnetProgramErrorTagged)
	if closeErr := readBuffer.CloseContext("programError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for programError")
	}

	// Virtual field
	_actualValue := programError
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataReasonForHalt"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataReasonForHalt")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataReasonForHalt{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ProgramError: programError,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataReasonForHalt) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataReasonForHalt"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataReasonForHalt")
		}

	// Simple Field (programError)
	if pushErr := writeBuffer.PushContext("programError"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for programError")
	}
	_programErrorErr := writeBuffer.WriteSerializable(m.GetProgramError())
	if popErr := writeBuffer.PopContext("programError"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for programError")
	}
	if _programErrorErr != nil {
		return errors.Wrap(_programErrorErr, "Error serializing 'programError' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataReasonForHalt"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataReasonForHalt")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataReasonForHalt) isBACnetConstructedDataReasonForHalt() bool {
	return true
}

func (m *_BACnetConstructedDataReasonForHalt) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



