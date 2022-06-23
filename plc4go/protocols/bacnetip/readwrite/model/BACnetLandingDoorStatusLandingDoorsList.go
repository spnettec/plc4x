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

// BACnetLandingDoorStatusLandingDoorsList is the corresponding interface of BACnetLandingDoorStatusLandingDoorsList
type BACnetLandingDoorStatusLandingDoorsList interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetLandingDoors returns LandingDoors (property field)
	GetLandingDoors() []BACnetLandingDoorStatusLandingDoorsListEntry
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetLandingDoorStatusLandingDoorsListExactly can be used when we want exactly this type and not a type which fulfills BACnetLandingDoorStatusLandingDoorsList.
// This is useful for switch cases.
type BACnetLandingDoorStatusLandingDoorsListExactly interface {
	BACnetLandingDoorStatusLandingDoorsList
	isBACnetLandingDoorStatusLandingDoorsList() bool
}

// _BACnetLandingDoorStatusLandingDoorsList is the data-structure of this message
type _BACnetLandingDoorStatusLandingDoorsList struct {
	OpeningTag   BACnetOpeningTag
	LandingDoors []BACnetLandingDoorStatusLandingDoorsListEntry
	ClosingTag   BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLandingDoorStatusLandingDoorsList) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetLandingDoorStatusLandingDoorsList) GetLandingDoors() []BACnetLandingDoorStatusLandingDoorsListEntry {
	return m.LandingDoors
}

func (m *_BACnetLandingDoorStatusLandingDoorsList) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLandingDoorStatusLandingDoorsList factory function for _BACnetLandingDoorStatusLandingDoorsList
func NewBACnetLandingDoorStatusLandingDoorsList(openingTag BACnetOpeningTag, landingDoors []BACnetLandingDoorStatusLandingDoorsListEntry, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetLandingDoorStatusLandingDoorsList {
	return &_BACnetLandingDoorStatusLandingDoorsList{OpeningTag: openingTag, LandingDoors: landingDoors, ClosingTag: closingTag, TagNumber: tagNumber}
}

// Deprecated: use the interface for direct cast
func CastBACnetLandingDoorStatusLandingDoorsList(structType interface{}) BACnetLandingDoorStatusLandingDoorsList {
	if casted, ok := structType.(BACnetLandingDoorStatusLandingDoorsList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLandingDoorStatusLandingDoorsList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLandingDoorStatusLandingDoorsList) GetTypeName() string {
	return "BACnetLandingDoorStatusLandingDoorsList"
}

func (m *_BACnetLandingDoorStatusLandingDoorsList) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLandingDoorStatusLandingDoorsList) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.LandingDoors) > 0 {
		for _, element := range m.LandingDoors {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetLandingDoorStatusLandingDoorsList) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLandingDoorStatusLandingDoorsListParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLandingDoorStatusLandingDoorsList, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLandingDoorStatusLandingDoorsList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLandingDoorStatusLandingDoorsList")
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
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (landingDoors)
	if pullErr := readBuffer.PullContext("landingDoors", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for landingDoors")
	}
	// Terminated array
	var landingDoors []BACnetLandingDoorStatusLandingDoorsListEntry
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetLandingDoorStatusLandingDoorsListEntryParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'landingDoors' field")
			}
			landingDoors = append(landingDoors, _item.(BACnetLandingDoorStatusLandingDoorsListEntry))

		}
	}
	if closeErr := readBuffer.CloseContext("landingDoors", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for landingDoors")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetLandingDoorStatusLandingDoorsList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLandingDoorStatusLandingDoorsList")
	}

	// Create the instance
	return NewBACnetLandingDoorStatusLandingDoorsList(openingTag, landingDoors, closingTag, tagNumber), nil
}

func (m *_BACnetLandingDoorStatusLandingDoorsList) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLandingDoorStatusLandingDoorsList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLandingDoorStatusLandingDoorsList")
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

	// Array Field (landingDoors)
	if pushErr := writeBuffer.PushContext("landingDoors", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for landingDoors")
	}
	for _, _element := range m.GetLandingDoors() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'landingDoors' field")
		}
	}
	if popErr := writeBuffer.PopContext("landingDoors", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for landingDoors")
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

	if popErr := writeBuffer.PopContext("BACnetLandingDoorStatusLandingDoorsList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLandingDoorStatusLandingDoorsList")
	}
	return nil
}

func (m *_BACnetLandingDoorStatusLandingDoorsList) isBACnetLandingDoorStatusLandingDoorsList() bool {
	return true
}

func (m *_BACnetLandingDoorStatusLandingDoorsList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
