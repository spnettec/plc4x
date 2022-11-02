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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetConstructedDataModelName is the corresponding interface of BACnetConstructedDataModelName
type BACnetConstructedDataModelName interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetModelName returns ModelName (property field)
	GetModelName() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataModelNameExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataModelName.
// This is useful for switch cases.
type BACnetConstructedDataModelNameExactly interface {
	BACnetConstructedDataModelName
	isBACnetConstructedDataModelName() bool
}

// _BACnetConstructedDataModelName is the data-structure of this message
type _BACnetConstructedDataModelName struct {
	*_BACnetConstructedData
        ModelName BACnetApplicationTagCharacterString
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataModelName)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataModelName)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_MODEL_NAME}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataModelName) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataModelName)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataModelName) GetModelName() BACnetApplicationTagCharacterString {
	return m.ModelName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataModelName) GetActualValue() BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetModelName())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataModelName factory function for _BACnetConstructedDataModelName
func NewBACnetConstructedDataModelName( modelName BACnetApplicationTagCharacterString , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataModelName {
	_result := &_BACnetConstructedDataModelName{
		ModelName: modelName,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataModelName(structType interface{}) BACnetConstructedDataModelName {
    if casted, ok := structType.(BACnetConstructedDataModelName); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataModelName); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataModelName) GetTypeName() string {
	return "BACnetConstructedDataModelName"
}

func (m *_BACnetConstructedDataModelName) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataModelName) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (modelName)
	lengthInBits += m.ModelName.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataModelName) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataModelNameParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataModelName, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataModelName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataModelName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (modelName)
	if pullErr := readBuffer.PullContext("modelName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for modelName")
	}
_modelName, _modelNameErr := BACnetApplicationTagParse(readBuffer)
	if _modelNameErr != nil {
		return nil, errors.Wrap(_modelNameErr, "Error parsing 'modelName' field of BACnetConstructedDataModelName")
	}
	modelName := _modelName.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("modelName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for modelName")
	}

	// Virtual field
	_actualValue := modelName
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataModelName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataModelName")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataModelName{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ModelName: modelName,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataModelName) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataModelName) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataModelName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataModelName")
		}

	// Simple Field (modelName)
	if pushErr := writeBuffer.PushContext("modelName"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for modelName")
	}
	_modelNameErr := writeBuffer.WriteSerializable(m.GetModelName())
	if popErr := writeBuffer.PopContext("modelName"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for modelName")
	}
	if _modelNameErr != nil {
		return errors.Wrap(_modelNameErr, "Error serializing 'modelName' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataModelName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataModelName")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataModelName) isBACnetConstructedDataModelName() bool {
	return true
}

func (m *_BACnetConstructedDataModelName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



