/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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

// BACnetConstructedDataModelName is the data-structure of this message
type BACnetConstructedDataModelName struct {
	*BACnetConstructedData
	ModelName *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataModelName is the corresponding interface of BACnetConstructedDataModelName
type IBACnetConstructedDataModelName interface {
	IBACnetConstructedData
	// GetModelName returns ModelName (property field)
	GetModelName() *BACnetApplicationTagCharacterString
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

func (m *BACnetConstructedDataModelName) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataModelName) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MODEL_NAME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataModelName) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataModelName) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataModelName) GetModelName() *BACnetApplicationTagCharacterString {
	return m.ModelName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataModelName factory function for BACnetConstructedDataModelName
func NewBACnetConstructedDataModelName(modelName *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataModelName {
	_result := &BACnetConstructedDataModelName{
		ModelName:             modelName,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataModelName(structType interface{}) *BACnetConstructedDataModelName {
	if casted, ok := structType.(BACnetConstructedDataModelName); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataModelName); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataModelName(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataModelName(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataModelName) GetTypeName() string {
	return "BACnetConstructedDataModelName"
}

func (m *BACnetConstructedDataModelName) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataModelName) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (modelName)
	lengthInBits += m.ModelName.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataModelName) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataModelNameParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataModelName, error) {
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
		return nil, errors.Wrap(_modelNameErr, "Error parsing 'modelName' field")
	}
	modelName := CastBACnetApplicationTagCharacterString(_modelName)
	if closeErr := readBuffer.CloseContext("modelName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for modelName")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataModelName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataModelName")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataModelName{
		ModelName:             CastBACnetApplicationTagCharacterString(modelName),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataModelName) Serialize(writeBuffer utils.WriteBuffer) error {
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
		_modelNameErr := m.ModelName.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("modelName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for modelName")
		}
		if _modelNameErr != nil {
			return errors.Wrap(_modelNameErr, "Error serializing 'modelName' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataModelName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataModelName")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataModelName) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
