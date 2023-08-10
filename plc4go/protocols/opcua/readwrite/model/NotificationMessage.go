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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// NotificationMessage is the corresponding interface of NotificationMessage
type NotificationMessage interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() uint32
	// GetPublishTime returns PublishTime (property field)
	GetPublishTime() int64
	// GetNoOfNotificationData returns NoOfNotificationData (property field)
	GetNoOfNotificationData() int32
	// GetNotificationData returns NotificationData (property field)
	GetNotificationData() []ExtensionObject
}

// NotificationMessageExactly can be used when we want exactly this type and not a type which fulfills NotificationMessage.
// This is useful for switch cases.
type NotificationMessageExactly interface {
	NotificationMessage
	isNotificationMessage() bool
}

// _NotificationMessage is the data-structure of this message
type _NotificationMessage struct {
	*_ExtensionObjectDefinition
	SequenceNumber       uint32
	PublishTime          int64
	NoOfNotificationData int32
	NotificationData     []ExtensionObject
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NotificationMessage) GetIdentifier() string {
	return "805"
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NotificationMessage) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_NotificationMessage) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NotificationMessage) GetSequenceNumber() uint32 {
	return m.SequenceNumber
}

func (m *_NotificationMessage) GetPublishTime() int64 {
	return m.PublishTime
}

func (m *_NotificationMessage) GetNoOfNotificationData() int32 {
	return m.NoOfNotificationData
}

func (m *_NotificationMessage) GetNotificationData() []ExtensionObject {
	return m.NotificationData
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNotificationMessage factory function for _NotificationMessage
func NewNotificationMessage(sequenceNumber uint32, publishTime int64, noOfNotificationData int32, notificationData []ExtensionObject) *_NotificationMessage {
	_result := &_NotificationMessage{
		SequenceNumber:             sequenceNumber,
		PublishTime:                publishTime,
		NoOfNotificationData:       noOfNotificationData,
		NotificationData:           notificationData,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNotificationMessage(structType any) NotificationMessage {
	if casted, ok := structType.(NotificationMessage); ok {
		return casted
	}
	if casted, ok := structType.(*NotificationMessage); ok {
		return *casted
	}
	return nil
}

func (m *_NotificationMessage) GetTypeName() string {
	return "NotificationMessage"
}

func (m *_NotificationMessage) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (sequenceNumber)
	lengthInBits += 32

	// Simple field (publishTime)
	lengthInBits += 64

	// Simple field (noOfNotificationData)
	lengthInBits += 32

	// Array field
	if len(m.NotificationData) > 0 {
		for _curItem, element := range m.NotificationData {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NotificationData), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_NotificationMessage) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NotificationMessageParse(ctx context.Context, theBytes []byte, identifier string) (NotificationMessage, error) {
	return NotificationMessageParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func NotificationMessageParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (NotificationMessage, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NotificationMessage"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NotificationMessage")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (sequenceNumber)
	_sequenceNumber, _sequenceNumberErr := readBuffer.ReadUint32("sequenceNumber", 32)
	if _sequenceNumberErr != nil {
		return nil, errors.Wrap(_sequenceNumberErr, "Error parsing 'sequenceNumber' field of NotificationMessage")
	}
	sequenceNumber := _sequenceNumber

	// Simple Field (publishTime)
	_publishTime, _publishTimeErr := readBuffer.ReadInt64("publishTime", 64)
	if _publishTimeErr != nil {
		return nil, errors.Wrap(_publishTimeErr, "Error parsing 'publishTime' field of NotificationMessage")
	}
	publishTime := _publishTime

	// Simple Field (noOfNotificationData)
	_noOfNotificationData, _noOfNotificationDataErr := readBuffer.ReadInt32("noOfNotificationData", 32)
	if _noOfNotificationDataErr != nil {
		return nil, errors.Wrap(_noOfNotificationDataErr, "Error parsing 'noOfNotificationData' field of NotificationMessage")
	}
	noOfNotificationData := _noOfNotificationData

	// Array field (notificationData)
	if pullErr := readBuffer.PullContext("notificationData", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for notificationData")
	}
	// Count array
	notificationData := make([]ExtensionObject, utils.Max(noOfNotificationData, 0))
	// This happens when the size is set conditional to 0
	if len(notificationData) == 0 {
		notificationData = nil
	}
	{
		_numItems := uint16(utils.Max(noOfNotificationData, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectParseWithBuffer(arrayCtx, readBuffer, bool(true))
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'notificationData' field of NotificationMessage")
			}
			notificationData[_curItem] = _item.(ExtensionObject)
		}
	}
	if closeErr := readBuffer.CloseContext("notificationData", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for notificationData")
	}

	if closeErr := readBuffer.CloseContext("NotificationMessage"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NotificationMessage")
	}

	// Create a partially initialized instance
	_child := &_NotificationMessage{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		SequenceNumber:             sequenceNumber,
		PublishTime:                publishTime,
		NoOfNotificationData:       noOfNotificationData,
		NotificationData:           notificationData,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_NotificationMessage) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NotificationMessage) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NotificationMessage"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NotificationMessage")
		}

		// Simple Field (sequenceNumber)
		sequenceNumber := uint32(m.GetSequenceNumber())
		_sequenceNumberErr := writeBuffer.WriteUint32("sequenceNumber", 32, (sequenceNumber))
		if _sequenceNumberErr != nil {
			return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
		}

		// Simple Field (publishTime)
		publishTime := int64(m.GetPublishTime())
		_publishTimeErr := writeBuffer.WriteInt64("publishTime", 64, (publishTime))
		if _publishTimeErr != nil {
			return errors.Wrap(_publishTimeErr, "Error serializing 'publishTime' field")
		}

		// Simple Field (noOfNotificationData)
		noOfNotificationData := int32(m.GetNoOfNotificationData())
		_noOfNotificationDataErr := writeBuffer.WriteInt32("noOfNotificationData", 32, (noOfNotificationData))
		if _noOfNotificationDataErr != nil {
			return errors.Wrap(_noOfNotificationDataErr, "Error serializing 'noOfNotificationData' field")
		}

		// Array Field (notificationData)
		if pushErr := writeBuffer.PushContext("notificationData", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for notificationData")
		}
		for _curItem, _element := range m.GetNotificationData() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetNotificationData()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'notificationData' field")
			}
		}
		if popErr := writeBuffer.PopContext("notificationData", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for notificationData")
		}

		if popErr := writeBuffer.PopContext("NotificationMessage"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NotificationMessage")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NotificationMessage) isNotificationMessage() bool {
	return true
}

func (m *_NotificationMessage) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
