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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetNameValueCollection is the data-structure of this message
type BACnetNameValueCollection struct {
	OpeningTag *BACnetOpeningTag
	Members    []*BACnetNameValue
	ClosingTag *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetNameValueCollection is the corresponding interface of BACnetNameValueCollection
type IBACnetNameValueCollection interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetMembers returns Members (property field)
	GetMembers() []*BACnetNameValue
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

func (m *BACnetNameValueCollection) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetNameValueCollection) GetMembers() []*BACnetNameValue {
	return m.Members
}

func (m *BACnetNameValueCollection) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNameValueCollection factory function for BACnetNameValueCollection
func NewBACnetNameValueCollection(openingTag *BACnetOpeningTag, members []*BACnetNameValue, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetNameValueCollection {
	return &BACnetNameValueCollection{OpeningTag: openingTag, Members: members, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetNameValueCollection(structType interface{}) *BACnetNameValueCollection {
	if casted, ok := structType.(BACnetNameValueCollection); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNameValueCollection); ok {
		return casted
	}
	return nil
}

func (m *BACnetNameValueCollection) GetTypeName() string {
	return "BACnetNameValueCollection"
}

func (m *BACnetNameValueCollection) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNameValueCollection) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.Members) > 0 {
		for _, element := range m.Members {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNameValueCollection) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNameValueCollectionParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetNameValueCollection, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNameValueCollection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNameValueCollection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (members)
	if pullErr := readBuffer.PullContext("members", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for members")
	}
	// Terminated array
	members := make([]*BACnetNameValue, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetNameValueParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'members' field")
			}
			members = append(members, CastBACnetNameValue(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("members", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for members")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetNameValueCollection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNameValueCollection")
	}

	// Create the instance
	return NewBACnetNameValueCollection(openingTag, members, closingTag, tagNumber), nil
}

func (m *BACnetNameValueCollection) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetNameValueCollection"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetNameValueCollection")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.OpeningTag)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (members)
	if m.Members != nil {
		if pushErr := writeBuffer.PushContext("members", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for members")
		}
		for _, _element := range m.Members {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'members' field")
			}
		}
		if popErr := writeBuffer.PopContext("members", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for members")
		}
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.ClosingTag)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetNameValueCollection"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetNameValueCollection")
	}
	return nil
}

func (m *BACnetNameValueCollection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
