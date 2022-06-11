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

// BACnetConstructedDataInputReference is the data-structure of this message
type BACnetConstructedDataInputReference struct {
	*BACnetConstructedData
	InputReference *BACnetObjectPropertyReference

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataInputReference is the corresponding interface of BACnetConstructedDataInputReference
type IBACnetConstructedDataInputReference interface {
	IBACnetConstructedData
	// GetInputReference returns InputReference (property field)
	GetInputReference() *BACnetObjectPropertyReference
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

func (m *BACnetConstructedDataInputReference) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataInputReference) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INPUT_REFERENCE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataInputReference) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataInputReference) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataInputReference) GetInputReference() *BACnetObjectPropertyReference {
	return m.InputReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataInputReference factory function for BACnetConstructedDataInputReference
func NewBACnetConstructedDataInputReference(inputReference *BACnetObjectPropertyReference, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataInputReference {
	_result := &BACnetConstructedDataInputReference{
		InputReference:        inputReference,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataInputReference(structType interface{}) *BACnetConstructedDataInputReference {
	if casted, ok := structType.(BACnetConstructedDataInputReference); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInputReference); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataInputReference(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataInputReference(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataInputReference) GetTypeName() string {
	return "BACnetConstructedDataInputReference"
}

func (m *BACnetConstructedDataInputReference) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataInputReference) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (inputReference)
	lengthInBits += m.InputReference.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataInputReference) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataInputReferenceParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataInputReference, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInputReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInputReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (inputReference)
	if pullErr := readBuffer.PullContext("inputReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for inputReference")
	}
	_inputReference, _inputReferenceErr := BACnetObjectPropertyReferenceParse(readBuffer)
	if _inputReferenceErr != nil {
		return nil, errors.Wrap(_inputReferenceErr, "Error parsing 'inputReference' field")
	}
	inputReference := CastBACnetObjectPropertyReference(_inputReference)
	if closeErr := readBuffer.CloseContext("inputReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for inputReference")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInputReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInputReference")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataInputReference{
		InputReference:        CastBACnetObjectPropertyReference(inputReference),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataInputReference) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInputReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInputReference")
		}

		// Simple Field (inputReference)
		if pushErr := writeBuffer.PushContext("inputReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for inputReference")
		}
		_inputReferenceErr := m.InputReference.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("inputReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for inputReference")
		}
		if _inputReferenceErr != nil {
			return errors.Wrap(_inputReferenceErr, "Error serializing 'inputReference' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInputReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInputReference")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataInputReference) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
