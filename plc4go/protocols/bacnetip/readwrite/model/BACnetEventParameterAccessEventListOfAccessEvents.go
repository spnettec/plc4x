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

// BACnetEventParameterAccessEventListOfAccessEvents is the corresponding interface of BACnetEventParameterAccessEventListOfAccessEvents
type BACnetEventParameterAccessEventListOfAccessEvents interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfAccessEvents returns ListOfAccessEvents (property field)
	GetListOfAccessEvents() []BACnetDeviceObjectPropertyReference
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetEventParameterAccessEventListOfAccessEventsExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterAccessEventListOfAccessEvents.
// This is useful for switch cases.
type BACnetEventParameterAccessEventListOfAccessEventsExactly interface {
	BACnetEventParameterAccessEventListOfAccessEvents
	isBACnetEventParameterAccessEventListOfAccessEvents() bool
}

// _BACnetEventParameterAccessEventListOfAccessEvents is the data-structure of this message
type _BACnetEventParameterAccessEventListOfAccessEvents struct {
	OpeningTag         BACnetOpeningTag
	ListOfAccessEvents []BACnetDeviceObjectPropertyReference
	ClosingTag         BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) GetListOfAccessEvents() []BACnetDeviceObjectPropertyReference {
	return m.ListOfAccessEvents
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventParameterAccessEventListOfAccessEvents factory function for _BACnetEventParameterAccessEventListOfAccessEvents
func NewBACnetEventParameterAccessEventListOfAccessEvents(openingTag BACnetOpeningTag, listOfAccessEvents []BACnetDeviceObjectPropertyReference, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventParameterAccessEventListOfAccessEvents {
	return &_BACnetEventParameterAccessEventListOfAccessEvents{OpeningTag: openingTag, ListOfAccessEvents: listOfAccessEvents, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterAccessEventListOfAccessEvents(structType interface{}) BACnetEventParameterAccessEventListOfAccessEvents {
	if casted, ok := structType.(BACnetEventParameterAccessEventListOfAccessEvents); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterAccessEventListOfAccessEvents); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) GetTypeName() string {
	return "BACnetEventParameterAccessEventListOfAccessEvents"
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.ListOfAccessEvents) > 0 {
		for _, element := range m.ListOfAccessEvents {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterAccessEventListOfAccessEventsParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventParameterAccessEventListOfAccessEvents, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterAccessEventListOfAccessEvents"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterAccessEventListOfAccessEvents")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetEventParameterAccessEventListOfAccessEvents")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (listOfAccessEvents)
	if pullErr := readBuffer.PullContext("listOfAccessEvents", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfAccessEvents")
	}
	// Terminated array
	var listOfAccessEvents []BACnetDeviceObjectPropertyReference
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDeviceObjectPropertyReferenceParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfAccessEvents' field of BACnetEventParameterAccessEventListOfAccessEvents")
			}
			listOfAccessEvents = append(listOfAccessEvents, _item.(BACnetDeviceObjectPropertyReference))

		}
	}
	if closeErr := readBuffer.CloseContext("listOfAccessEvents", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfAccessEvents")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetEventParameterAccessEventListOfAccessEvents")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterAccessEventListOfAccessEvents"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterAccessEventListOfAccessEvents")
	}

	// Create the instance
	return &_BACnetEventParameterAccessEventListOfAccessEvents{
		TagNumber:          tagNumber,
		OpeningTag:         openingTag,
		ListOfAccessEvents: listOfAccessEvents,
		ClosingTag:         closingTag,
	}, nil
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetEventParameterAccessEventListOfAccessEvents"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterAccessEventListOfAccessEvents")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Array Field (listOfAccessEvents)
	if pushErr := writeBuffer.PushContext("listOfAccessEvents", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfAccessEvents")
	}
	for _, _element := range m.GetListOfAccessEvents() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listOfAccessEvents' field")
		}
	}
	if popErr := writeBuffer.PopContext("listOfAccessEvents", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfAccessEvents")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameterAccessEventListOfAccessEvents"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameterAccessEventListOfAccessEvents")
	}
	return nil
}

////
// Arguments Getter

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) GetTagNumber() uint8 {
	return m.TagNumber
}

//
////

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) isBACnetEventParameterAccessEventListOfAccessEvents() bool {
	return true
}

func (m *_BACnetEventParameterAccessEventListOfAccessEvents) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
