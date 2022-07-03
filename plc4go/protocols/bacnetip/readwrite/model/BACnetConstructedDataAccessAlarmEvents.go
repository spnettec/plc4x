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

// BACnetConstructedDataAccessAlarmEvents is the corresponding interface of BACnetConstructedDataAccessAlarmEvents
type BACnetConstructedDataAccessAlarmEvents interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAccessAlarmEvents returns AccessAlarmEvents (property field)
	GetAccessAlarmEvents() []BACnetAccessEventTagged
}

// BACnetConstructedDataAccessAlarmEventsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAccessAlarmEvents.
// This is useful for switch cases.
type BACnetConstructedDataAccessAlarmEventsExactly interface {
	BACnetConstructedDataAccessAlarmEvents
	isBACnetConstructedDataAccessAlarmEvents() bool
}

// _BACnetConstructedDataAccessAlarmEvents is the data-structure of this message
type _BACnetConstructedDataAccessAlarmEvents struct {
	*_BACnetConstructedData
	AccessAlarmEvents []BACnetAccessEventTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAccessAlarmEvents) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAccessAlarmEvents) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCESS_ALARM_EVENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAccessAlarmEvents) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAccessAlarmEvents) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAccessAlarmEvents) GetAccessAlarmEvents() []BACnetAccessEventTagged {
	return m.AccessAlarmEvents
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessAlarmEvents factory function for _BACnetConstructedDataAccessAlarmEvents
func NewBACnetConstructedDataAccessAlarmEvents(accessAlarmEvents []BACnetAccessEventTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAccessAlarmEvents {
	_result := &_BACnetConstructedDataAccessAlarmEvents{
		AccessAlarmEvents:      accessAlarmEvents,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAccessAlarmEvents(structType interface{}) BACnetConstructedDataAccessAlarmEvents {
	if casted, ok := structType.(BACnetConstructedDataAccessAlarmEvents); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessAlarmEvents); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAccessAlarmEvents) GetTypeName() string {
	return "BACnetConstructedDataAccessAlarmEvents"
}

func (m *_BACnetConstructedDataAccessAlarmEvents) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAccessAlarmEvents) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.AccessAlarmEvents) > 0 {
		for _, element := range m.AccessAlarmEvents {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAccessAlarmEvents) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAccessAlarmEventsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAccessAlarmEvents, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessAlarmEvents"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessAlarmEvents")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (accessAlarmEvents)
	if pullErr := readBuffer.PullContext("accessAlarmEvents", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessAlarmEvents")
	}
	// Terminated array
	var accessAlarmEvents []BACnetAccessEventTagged
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetAccessEventTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'accessAlarmEvents' field of BACnetConstructedDataAccessAlarmEvents")
			}
			accessAlarmEvents = append(accessAlarmEvents, _item.(BACnetAccessEventTagged))

		}
	}
	if closeErr := readBuffer.CloseContext("accessAlarmEvents", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessAlarmEvents")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessAlarmEvents"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessAlarmEvents")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAccessAlarmEvents{
		AccessAlarmEvents: accessAlarmEvents,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAccessAlarmEvents) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessAlarmEvents"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessAlarmEvents")
		}

		// Array Field (accessAlarmEvents)
		if pushErr := writeBuffer.PushContext("accessAlarmEvents", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for accessAlarmEvents")
		}
		for _, _element := range m.GetAccessAlarmEvents() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'accessAlarmEvents' field")
			}
		}
		if popErr := writeBuffer.PopContext("accessAlarmEvents", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for accessAlarmEvents")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessAlarmEvents"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessAlarmEvents")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAccessAlarmEvents) isBACnetConstructedDataAccessAlarmEvents() bool {
	return true
}

func (m *_BACnetConstructedDataAccessAlarmEvents) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
