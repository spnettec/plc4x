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

// BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification is the data-structure of this message
type BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification struct {
	OpeningTag                              *BACnetOpeningTag
	ListOfCovSubscriptionSpecificationEntry []*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
	ClosingTag                              *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification is the corresponding interface of BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification
type IBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetListOfCovSubscriptionSpecificationEntry returns ListOfCovSubscriptionSpecificationEntry (property field)
	GetListOfCovSubscriptionSpecificationEntry() []*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetListOfCovSubscriptionSpecificationEntry() []*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry {
	return m.ListOfCovSubscriptionSpecificationEntry
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification factory function for BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification
func NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification(openingTag *BACnetOpeningTag, listOfCovSubscriptionSpecificationEntry []*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification {
	return &BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification{OpeningTag: openingTag, ListOfCovSubscriptionSpecificationEntry: listOfCovSubscriptionSpecificationEntry, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification(structType interface{}) *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification {
	if casted, ok := structType.(BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification); ok {
		return casted
	}
	return nil
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetTypeName() string {
	return "BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.ListOfCovSubscriptionSpecificationEntry) > 0 {
		for _, element := range m.ListOfCovSubscriptionSpecificationEntry {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, pullErr
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, closeErr
	}

	// Array field (listOfCovSubscriptionSpecificationEntry)
	if pullErr := readBuffer.PullContext("listOfCovSubscriptionSpecificationEntry", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	listOfCovSubscriptionSpecificationEntry := make([]*BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntryParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfCovSubscriptionSpecificationEntry' field")
			}
			listOfCovSubscriptionSpecificationEntry = append(listOfCovSubscriptionSpecificationEntry, CastBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecificationEntry(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("listOfCovSubscriptionSpecificationEntry", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, pullErr
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification(openingTag, listOfCovSubscriptionSpecificationEntry, closingTag, tagNumber), nil
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"); pushErr != nil {
		return pushErr
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return pushErr
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return popErr
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (listOfCovSubscriptionSpecificationEntry)
	if m.ListOfCovSubscriptionSpecificationEntry != nil {
		if pushErr := writeBuffer.PushContext("listOfCovSubscriptionSpecificationEntry", utils.WithRenderAsList(true)); pushErr != nil {
			return pushErr
		}
		for _, _element := range m.ListOfCovSubscriptionSpecificationEntry {
			_elementErr := _element.Serialize(writeBuffer)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'listOfCovSubscriptionSpecificationEntry' field")
			}
		}
		if popErr := writeBuffer.PopContext("listOfCovSubscriptionSpecificationEntry", utils.WithRenderAsList(true)); popErr != nil {
			return popErr
		}
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return pushErr
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return popErr
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetCOVMultipleSubscriptionListOfCovSubscriptionSpecification) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
