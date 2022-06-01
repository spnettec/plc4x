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

// BACnetConstructedDataListOfObjectPropertyReferences is the data-structure of this message
type BACnetConstructedDataListOfObjectPropertyReferences struct {
	*BACnetConstructedData
	References []*BACnetDeviceObjectPropertyReference

	// Arguments.
	TagNumber uint8
}

// IBACnetConstructedDataListOfObjectPropertyReferences is the corresponding interface of BACnetConstructedDataListOfObjectPropertyReferences
type IBACnetConstructedDataListOfObjectPropertyReferences interface {
	IBACnetConstructedData
	// GetReferences returns References (property field)
	GetReferences() []*BACnetDeviceObjectPropertyReference
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

func (m *BACnetConstructedDataListOfObjectPropertyReferences) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataListOfObjectPropertyReferences) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIST_OF_OBJECT_PROPERTY_REFERENCES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataListOfObjectPropertyReferences) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataListOfObjectPropertyReferences) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataListOfObjectPropertyReferences) GetReferences() []*BACnetDeviceObjectPropertyReference {
	return m.References
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataListOfObjectPropertyReferences factory function for BACnetConstructedDataListOfObjectPropertyReferences
func NewBACnetConstructedDataListOfObjectPropertyReferences(references []*BACnetDeviceObjectPropertyReference, openingTag *BACnetOpeningTag, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConstructedDataListOfObjectPropertyReferences {
	_result := &BACnetConstructedDataListOfObjectPropertyReferences{
		References:            references,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataListOfObjectPropertyReferences(structType interface{}) *BACnetConstructedDataListOfObjectPropertyReferences {
	if casted, ok := structType.(BACnetConstructedDataListOfObjectPropertyReferences); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataListOfObjectPropertyReferences); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataListOfObjectPropertyReferences(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataListOfObjectPropertyReferences(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataListOfObjectPropertyReferences) GetTypeName() string {
	return "BACnetConstructedDataListOfObjectPropertyReferences"
}

func (m *BACnetConstructedDataListOfObjectPropertyReferences) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataListOfObjectPropertyReferences) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.References) > 0 {
		for _, element := range m.References {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataListOfObjectPropertyReferences) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataListOfObjectPropertyReferencesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier) (*BACnetConstructedDataListOfObjectPropertyReferences, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataListOfObjectPropertyReferences"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (references)
	if pullErr := readBuffer.PullContext("references", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	references := make([]*BACnetDeviceObjectPropertyReference, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDeviceObjectPropertyReferenceParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'references' field")
			}
			references = append(references, CastBACnetDeviceObjectPropertyReference(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("references", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataListOfObjectPropertyReferences"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataListOfObjectPropertyReferences{
		References:            references,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataListOfObjectPropertyReferences) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataListOfObjectPropertyReferences"); pushErr != nil {
			return pushErr
		}

		// Array Field (references)
		if m.References != nil {
			if pushErr := writeBuffer.PushContext("references", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.References {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'references' field")
				}
			}
			if popErr := writeBuffer.PopContext("references", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataListOfObjectPropertyReferences"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataListOfObjectPropertyReferences) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
