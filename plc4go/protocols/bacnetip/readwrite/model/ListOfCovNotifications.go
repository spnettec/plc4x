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


// ListOfCovNotifications is the corresponding interface of ListOfCovNotifications
type ListOfCovNotifications interface {
	utils.LengthAware
	utils.Serializable
	// GetMonitoredObjectIdentifier returns MonitoredObjectIdentifier (property field)
	GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetListOfValues returns ListOfValues (property field)
	GetListOfValues() []ListOfCovNotificationsValue
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// ListOfCovNotificationsExactly can be used when we want exactly this type and not a type which fulfills ListOfCovNotifications.
// This is useful for switch cases.
type ListOfCovNotificationsExactly interface {
	ListOfCovNotifications
	isListOfCovNotifications() bool
}

// _ListOfCovNotifications is the data-structure of this message
type _ListOfCovNotifications struct {
        MonitoredObjectIdentifier BACnetContextTagObjectIdentifier
        OpeningTag BACnetOpeningTag
        ListOfValues []ListOfCovNotificationsValue
        ClosingTag BACnetClosingTag
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ListOfCovNotifications) GetMonitoredObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.MonitoredObjectIdentifier
}

func (m *_ListOfCovNotifications) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_ListOfCovNotifications) GetListOfValues() []ListOfCovNotificationsValue {
	return m.ListOfValues
}

func (m *_ListOfCovNotifications) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewListOfCovNotifications factory function for _ListOfCovNotifications
func NewListOfCovNotifications( monitoredObjectIdentifier BACnetContextTagObjectIdentifier , openingTag BACnetOpeningTag , listOfValues []ListOfCovNotificationsValue , closingTag BACnetClosingTag ) *_ListOfCovNotifications {
return &_ListOfCovNotifications{ MonitoredObjectIdentifier: monitoredObjectIdentifier , OpeningTag: openingTag , ListOfValues: listOfValues , ClosingTag: closingTag }
}

// Deprecated: use the interface for direct cast
func CastListOfCovNotifications(structType interface{}) ListOfCovNotifications {
    if casted, ok := structType.(ListOfCovNotifications); ok {
		return casted
	}
	if casted, ok := structType.(*ListOfCovNotifications); ok {
		return *casted
	}
	return nil
}

func (m *_ListOfCovNotifications) GetTypeName() string {
	return "ListOfCovNotifications"
}

func (m *_ListOfCovNotifications) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ListOfCovNotifications) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (monitoredObjectIdentifier)
	lengthInBits += m.MonitoredObjectIdentifier.GetLengthInBits()

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Array field
	if len(m.ListOfValues) > 0 {
		for _, element := range m.ListOfValues {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}


func (m *_ListOfCovNotifications) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ListOfCovNotificationsParse(readBuffer utils.ReadBuffer) (ListOfCovNotifications, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ListOfCovNotifications"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ListOfCovNotifications")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (monitoredObjectIdentifier)
	if pullErr := readBuffer.PullContext("monitoredObjectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for monitoredObjectIdentifier")
	}
_monitoredObjectIdentifier, _monitoredObjectIdentifierErr := BACnetContextTagParse(readBuffer , uint8( uint8(0) ) , BACnetDataType( BACnetDataType_BACNET_OBJECT_IDENTIFIER ) )
	if _monitoredObjectIdentifierErr != nil {
		return nil, errors.Wrap(_monitoredObjectIdentifierErr, "Error parsing 'monitoredObjectIdentifier' field of ListOfCovNotifications")
	}
	monitoredObjectIdentifier := _monitoredObjectIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("monitoredObjectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for monitoredObjectIdentifier")
	}

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer , uint8( uint8(1) ) )
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of ListOfCovNotifications")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Array field (listOfValues)
	if pullErr := readBuffer.PullContext("listOfValues", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfValues")
	}
	// Terminated array
	var listOfValues []ListOfCovNotificationsValue
	{
		for ;!bool(IsBACnetConstructedDataClosingTag(readBuffer, false, 1)); {
_item, _err := ListOfCovNotificationsValueParse(readBuffer , monitoredObjectIdentifier.GetObjectType() )
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfValues' field of ListOfCovNotifications")
			}
			listOfValues = append(listOfValues, _item.(ListOfCovNotificationsValue))

		}
	}
	if closeErr := readBuffer.CloseContext("listOfValues", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfValues")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer , uint8( uint8(1) ) )
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of ListOfCovNotifications")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("ListOfCovNotifications"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ListOfCovNotifications")
	}

	// Create the instance
	return &_ListOfCovNotifications{
			MonitoredObjectIdentifier: monitoredObjectIdentifier,
			OpeningTag: openingTag,
			ListOfValues: listOfValues,
			ClosingTag: closingTag,
		}, nil
}

func (m *_ListOfCovNotifications) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ListOfCovNotifications) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("ListOfCovNotifications"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ListOfCovNotifications")
	}

	// Simple Field (monitoredObjectIdentifier)
	if pushErr := writeBuffer.PushContext("monitoredObjectIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for monitoredObjectIdentifier")
	}
	_monitoredObjectIdentifierErr := writeBuffer.WriteSerializable(m.GetMonitoredObjectIdentifier())
	if popErr := writeBuffer.PopContext("monitoredObjectIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for monitoredObjectIdentifier")
	}
	if _monitoredObjectIdentifierErr != nil {
		return errors.Wrap(_monitoredObjectIdentifierErr, "Error serializing 'monitoredObjectIdentifier' field")
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

	// Array Field (listOfValues)
	if pushErr := writeBuffer.PushContext("listOfValues", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfValues")
	}
	for _, _element := range m.GetListOfValues() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'listOfValues' field")
		}
	}
	if popErr := writeBuffer.PopContext("listOfValues", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfValues")
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

	if popErr := writeBuffer.PopContext("ListOfCovNotifications"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ListOfCovNotifications")
	}
	return nil
}


func (m *_ListOfCovNotifications) isListOfCovNotifications() bool {
	return true
}

func (m *_ListOfCovNotifications) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



