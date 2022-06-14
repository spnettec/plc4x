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

// BACnetConstructedDataAccessTransactionEvents is the data-structure of this message
type BACnetConstructedDataAccessTransactionEvents struct {
	*BACnetConstructedData
	AccessTransactionEvents []*BACnetAccessEventTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataAccessTransactionEvents is the corresponding interface of BACnetConstructedDataAccessTransactionEvents
type IBACnetConstructedDataAccessTransactionEvents interface {
	IBACnetConstructedData
	// GetAccessTransactionEvents returns AccessTransactionEvents (property field)
	GetAccessTransactionEvents() []*BACnetAccessEventTagged
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

func (m *BACnetConstructedDataAccessTransactionEvents) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataAccessTransactionEvents) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCESS_TRANSACTION_EVENTS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataAccessTransactionEvents) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataAccessTransactionEvents) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataAccessTransactionEvents) GetAccessTransactionEvents() []*BACnetAccessEventTagged {
	return m.AccessTransactionEvents
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAccessTransactionEvents factory function for BACnetConstructedDataAccessTransactionEvents
func NewBACnetConstructedDataAccessTransactionEvents(accessTransactionEvents []*BACnetAccessEventTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataAccessTransactionEvents {
	_result := &BACnetConstructedDataAccessTransactionEvents{
		AccessTransactionEvents: accessTransactionEvents,
		BACnetConstructedData:   NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataAccessTransactionEvents(structType interface{}) *BACnetConstructedDataAccessTransactionEvents {
	if casted, ok := structType.(BACnetConstructedDataAccessTransactionEvents); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAccessTransactionEvents); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataAccessTransactionEvents(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataAccessTransactionEvents(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataAccessTransactionEvents) GetTypeName() string {
	return "BACnetConstructedDataAccessTransactionEvents"
}

func (m *BACnetConstructedDataAccessTransactionEvents) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataAccessTransactionEvents) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.AccessTransactionEvents) > 0 {
		for _, element := range m.AccessTransactionEvents {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataAccessTransactionEvents) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAccessTransactionEventsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataAccessTransactionEvents, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAccessTransactionEvents"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAccessTransactionEvents")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (accessTransactionEvents)
	if pullErr := readBuffer.PullContext("accessTransactionEvents", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessTransactionEvents")
	}
	// Terminated array
	accessTransactionEvents := make([]*BACnetAccessEventTagged, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetAccessEventTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'accessTransactionEvents' field")
			}
			accessTransactionEvents = append(accessTransactionEvents, CastBACnetAccessEventTagged(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("accessTransactionEvents", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessTransactionEvents")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAccessTransactionEvents"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAccessTransactionEvents")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataAccessTransactionEvents{
		AccessTransactionEvents: accessTransactionEvents,
		BACnetConstructedData:   &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataAccessTransactionEvents) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAccessTransactionEvents"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAccessTransactionEvents")
		}

		// Array Field (accessTransactionEvents)
		if m.AccessTransactionEvents != nil {
			if pushErr := writeBuffer.PushContext("accessTransactionEvents", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for accessTransactionEvents")
			}
			for _, _element := range m.AccessTransactionEvents {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'accessTransactionEvents' field")
				}
			}
			if popErr := writeBuffer.PopContext("accessTransactionEvents", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for accessTransactionEvents")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAccessTransactionEvents"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAccessTransactionEvents")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataAccessTransactionEvents) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
