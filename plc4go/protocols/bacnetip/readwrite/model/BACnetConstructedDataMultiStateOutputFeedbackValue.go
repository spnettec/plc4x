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

// BACnetConstructedDataMultiStateOutputFeedbackValue is the data-structure of this message
type BACnetConstructedDataMultiStateOutputFeedbackValue struct {
	*BACnetConstructedData
	FeedbackValue *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataMultiStateOutputFeedbackValue is the corresponding interface of BACnetConstructedDataMultiStateOutputFeedbackValue
type IBACnetConstructedDataMultiStateOutputFeedbackValue interface {
	IBACnetConstructedData
	// GetFeedbackValue returns FeedbackValue (property field)
	GetFeedbackValue() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataMultiStateOutputFeedbackValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_MULTI_STATE_OUTPUT
}

func (m *BACnetConstructedDataMultiStateOutputFeedbackValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_FEEDBACK_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMultiStateOutputFeedbackValue) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMultiStateOutputFeedbackValue) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMultiStateOutputFeedbackValue) GetFeedbackValue() *BACnetApplicationTagUnsignedInteger {
	return m.FeedbackValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMultiStateOutputFeedbackValue factory function for BACnetConstructedDataMultiStateOutputFeedbackValue
func NewBACnetConstructedDataMultiStateOutputFeedbackValue(feedbackValue *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataMultiStateOutputFeedbackValue {
	_result := &BACnetConstructedDataMultiStateOutputFeedbackValue{
		FeedbackValue:         feedbackValue,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMultiStateOutputFeedbackValue(structType interface{}) *BACnetConstructedDataMultiStateOutputFeedbackValue {
	if casted, ok := structType.(BACnetConstructedDataMultiStateOutputFeedbackValue); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMultiStateOutputFeedbackValue); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMultiStateOutputFeedbackValue(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMultiStateOutputFeedbackValue(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMultiStateOutputFeedbackValue) GetTypeName() string {
	return "BACnetConstructedDataMultiStateOutputFeedbackValue"
}

func (m *BACnetConstructedDataMultiStateOutputFeedbackValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMultiStateOutputFeedbackValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (feedbackValue)
	lengthInBits += m.FeedbackValue.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataMultiStateOutputFeedbackValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMultiStateOutputFeedbackValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataMultiStateOutputFeedbackValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMultiStateOutputFeedbackValue"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (feedbackValue)
	if pullErr := readBuffer.PullContext("feedbackValue"); pullErr != nil {
		return nil, pullErr
	}
	_feedbackValue, _feedbackValueErr := BACnetApplicationTagParse(readBuffer)
	if _feedbackValueErr != nil {
		return nil, errors.Wrap(_feedbackValueErr, "Error parsing 'feedbackValue' field")
	}
	feedbackValue := CastBACnetApplicationTagUnsignedInteger(_feedbackValue)
	if closeErr := readBuffer.CloseContext("feedbackValue"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMultiStateOutputFeedbackValue"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMultiStateOutputFeedbackValue{
		FeedbackValue:         CastBACnetApplicationTagUnsignedInteger(feedbackValue),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMultiStateOutputFeedbackValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMultiStateOutputFeedbackValue"); pushErr != nil {
			return pushErr
		}

		// Simple Field (feedbackValue)
		if pushErr := writeBuffer.PushContext("feedbackValue"); pushErr != nil {
			return pushErr
		}
		_feedbackValueErr := m.FeedbackValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("feedbackValue"); popErr != nil {
			return popErr
		}
		if _feedbackValueErr != nil {
			return errors.Wrap(_feedbackValueErr, "Error serializing 'feedbackValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMultiStateOutputFeedbackValue"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMultiStateOutputFeedbackValue) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
