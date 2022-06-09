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

// BACnetConstructedDataRecordsSinceNotification is the data-structure of this message
type BACnetConstructedDataRecordsSinceNotification struct {
	*BACnetConstructedData
	RecordsSinceNotifications *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataRecordsSinceNotification is the corresponding interface of BACnetConstructedDataRecordsSinceNotification
type IBACnetConstructedDataRecordsSinceNotification interface {
	IBACnetConstructedData
	// GetRecordsSinceNotifications returns RecordsSinceNotifications (property field)
	GetRecordsSinceNotifications() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataRecordsSinceNotification) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataRecordsSinceNotification) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RECORDS_SINCE_NOTIFICATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataRecordsSinceNotification) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataRecordsSinceNotification) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataRecordsSinceNotification) GetRecordsSinceNotifications() *BACnetApplicationTagUnsignedInteger {
	return m.RecordsSinceNotifications
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataRecordsSinceNotification factory function for BACnetConstructedDataRecordsSinceNotification
func NewBACnetConstructedDataRecordsSinceNotification(recordsSinceNotifications *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataRecordsSinceNotification {
	_result := &BACnetConstructedDataRecordsSinceNotification{
		RecordsSinceNotifications: recordsSinceNotifications,
		BACnetConstructedData:     NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataRecordsSinceNotification(structType interface{}) *BACnetConstructedDataRecordsSinceNotification {
	if casted, ok := structType.(BACnetConstructedDataRecordsSinceNotification); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRecordsSinceNotification); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataRecordsSinceNotification(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataRecordsSinceNotification(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataRecordsSinceNotification) GetTypeName() string {
	return "BACnetConstructedDataRecordsSinceNotification"
}

func (m *BACnetConstructedDataRecordsSinceNotification) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataRecordsSinceNotification) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (recordsSinceNotifications)
	lengthInBits += m.RecordsSinceNotifications.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataRecordsSinceNotification) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataRecordsSinceNotificationParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataRecordsSinceNotification, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRecordsSinceNotification"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (recordsSinceNotifications)
	if pullErr := readBuffer.PullContext("recordsSinceNotifications"); pullErr != nil {
		return nil, pullErr
	}
	_recordsSinceNotifications, _recordsSinceNotificationsErr := BACnetApplicationTagParse(readBuffer)
	if _recordsSinceNotificationsErr != nil {
		return nil, errors.Wrap(_recordsSinceNotificationsErr, "Error parsing 'recordsSinceNotifications' field")
	}
	recordsSinceNotifications := CastBACnetApplicationTagUnsignedInteger(_recordsSinceNotifications)
	if closeErr := readBuffer.CloseContext("recordsSinceNotifications"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRecordsSinceNotification"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataRecordsSinceNotification{
		RecordsSinceNotifications: CastBACnetApplicationTagUnsignedInteger(recordsSinceNotifications),
		BACnetConstructedData:     &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataRecordsSinceNotification) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRecordsSinceNotification"); pushErr != nil {
			return pushErr
		}

		// Simple Field (recordsSinceNotifications)
		if pushErr := writeBuffer.PushContext("recordsSinceNotifications"); pushErr != nil {
			return pushErr
		}
		_recordsSinceNotificationsErr := m.RecordsSinceNotifications.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("recordsSinceNotifications"); popErr != nil {
			return popErr
		}
		if _recordsSinceNotificationsErr != nil {
			return errors.Wrap(_recordsSinceNotificationsErr, "Error serializing 'recordsSinceNotifications' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRecordsSinceNotification"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataRecordsSinceNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
