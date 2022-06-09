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

// BACnetEventLogRecordLogDatumNotification is the data-structure of this message
type BACnetEventLogRecordLogDatumNotification struct {
	*BACnetEventLogRecordLogDatum
	InnerOpeningTag *BACnetOpeningTag
	Notification    *ConfirmedEventNotificationRequest
	InnerClosingTag *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetEventLogRecordLogDatumNotification is the corresponding interface of BACnetEventLogRecordLogDatumNotification
type IBACnetEventLogRecordLogDatumNotification interface {
	IBACnetEventLogRecordLogDatum
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() *BACnetOpeningTag
	// GetNotification returns Notification (property field)
	GetNotification() *ConfirmedEventNotificationRequest
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() *BACnetClosingTag
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetEventLogRecordLogDatumNotification) InitializeParent(parent *BACnetEventLogRecordLogDatum, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetEventLogRecordLogDatum.OpeningTag = openingTag
	m.BACnetEventLogRecordLogDatum.PeekedTagHeader = peekedTagHeader
	m.BACnetEventLogRecordLogDatum.ClosingTag = closingTag
}

func (m *BACnetEventLogRecordLogDatumNotification) GetParent() *BACnetEventLogRecordLogDatum {
	return m.BACnetEventLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetEventLogRecordLogDatumNotification) GetInnerOpeningTag() *BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *BACnetEventLogRecordLogDatumNotification) GetNotification() *ConfirmedEventNotificationRequest {
	return m.Notification
}

func (m *BACnetEventLogRecordLogDatumNotification) GetInnerClosingTag() *BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventLogRecordLogDatumNotification factory function for BACnetEventLogRecordLogDatumNotification
func NewBACnetEventLogRecordLogDatumNotification(innerOpeningTag *BACnetOpeningTag, notification *ConfirmedEventNotificationRequest, innerClosingTag *BACnetClosingTag, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetEventLogRecordLogDatumNotification {
	_result := &BACnetEventLogRecordLogDatumNotification{
		InnerOpeningTag:              innerOpeningTag,
		Notification:                 notification,
		InnerClosingTag:              innerClosingTag,
		BACnetEventLogRecordLogDatum: NewBACnetEventLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result.Child = _result
	return _result
}

func CastBACnetEventLogRecordLogDatumNotification(structType interface{}) *BACnetEventLogRecordLogDatumNotification {
	if casted, ok := structType.(BACnetEventLogRecordLogDatumNotification); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetEventLogRecordLogDatumNotification); ok {
		return casted
	}
	if casted, ok := structType.(BACnetEventLogRecordLogDatum); ok {
		return CastBACnetEventLogRecordLogDatumNotification(casted.Child)
	}
	if casted, ok := structType.(*BACnetEventLogRecordLogDatum); ok {
		return CastBACnetEventLogRecordLogDatumNotification(casted.Child)
	}
	return nil
}

func (m *BACnetEventLogRecordLogDatumNotification) GetTypeName() string {
	return "BACnetEventLogRecordLogDatumNotification"
}

func (m *BACnetEventLogRecordLogDatumNotification) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetEventLogRecordLogDatumNotification) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Simple field (notification)
	lengthInBits += m.Notification.GetLengthInBits()

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetEventLogRecordLogDatumNotification) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventLogRecordLogDatumNotificationParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetEventLogRecordLogDatumNotification, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventLogRecordLogDatumNotification"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, pullErr
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParse(readBuffer, uint8(uint8(1)))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field")
	}
	innerOpeningTag := CastBACnetOpeningTag(_innerOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (notification)
	if pullErr := readBuffer.PullContext("notification"); pullErr != nil {
		return nil, pullErr
	}
	_notification, _notificationErr := ConfirmedEventNotificationRequestParse(readBuffer)
	if _notificationErr != nil {
		return nil, errors.Wrap(_notificationErr, "Error parsing 'notification' field")
	}
	notification := CastConfirmedEventNotificationRequest(_notification)
	if closeErr := readBuffer.CloseContext("notification"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, pullErr
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field")
	}
	innerClosingTag := CastBACnetClosingTag(_innerClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetEventLogRecordLogDatumNotification"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetEventLogRecordLogDatumNotification{
		InnerOpeningTag:              CastBACnetOpeningTag(innerOpeningTag),
		Notification:                 CastConfirmedEventNotificationRequest(notification),
		InnerClosingTag:              CastBACnetClosingTag(innerClosingTag),
		BACnetEventLogRecordLogDatum: &BACnetEventLogRecordLogDatum{},
	}
	_child.BACnetEventLogRecordLogDatum.Child = _child
	return _child, nil
}

func (m *BACnetEventLogRecordLogDatumNotification) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventLogRecordLogDatumNotification"); pushErr != nil {
			return pushErr
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return pushErr
		}
		_innerOpeningTagErr := m.InnerOpeningTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return popErr
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (notification)
		if pushErr := writeBuffer.PushContext("notification"); pushErr != nil {
			return pushErr
		}
		_notificationErr := m.Notification.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("notification"); popErr != nil {
			return popErr
		}
		if _notificationErr != nil {
			return errors.Wrap(_notificationErr, "Error serializing 'notification' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return pushErr
		}
		_innerClosingTagErr := m.InnerClosingTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return popErr
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventLogRecordLogDatumNotification"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetEventLogRecordLogDatumNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
