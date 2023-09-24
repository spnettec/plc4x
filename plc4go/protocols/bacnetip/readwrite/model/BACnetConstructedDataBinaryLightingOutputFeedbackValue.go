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


// BACnetConstructedDataBinaryLightingOutputFeedbackValue is the corresponding interface of BACnetConstructedDataBinaryLightingOutputFeedbackValue
type BACnetConstructedDataBinaryLightingOutputFeedbackValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFeedbackValue returns FeedbackValue (property field)
	GetFeedbackValue() BACnetBinaryLightingPVTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetBinaryLightingPVTagged
}

// BACnetConstructedDataBinaryLightingOutputFeedbackValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBinaryLightingOutputFeedbackValue.
// This is useful for switch cases.
type BACnetConstructedDataBinaryLightingOutputFeedbackValueExactly interface {
	BACnetConstructedDataBinaryLightingOutputFeedbackValue
	isBACnetConstructedDataBinaryLightingOutputFeedbackValue() bool
}

// _BACnetConstructedDataBinaryLightingOutputFeedbackValue is the data-structure of this message
type _BACnetConstructedDataBinaryLightingOutputFeedbackValue struct {
	*_BACnetConstructedData
        FeedbackValue BACnetBinaryLightingPVTagged
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue)  GetObjectTypeArgument() BACnetObjectType {
return BACnetObjectType_BINARY_LIGHTING_OUTPUT}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_FEEDBACK_VALUE}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetFeedbackValue() BACnetBinaryLightingPVTagged {
	return m.FeedbackValue
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetActualValue() BACnetBinaryLightingPVTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetBinaryLightingPVTagged(m.GetFeedbackValue())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataBinaryLightingOutputFeedbackValue factory function for _BACnetConstructedDataBinaryLightingOutputFeedbackValue
func NewBACnetConstructedDataBinaryLightingOutputFeedbackValue( feedbackValue BACnetBinaryLightingPVTagged , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataBinaryLightingOutputFeedbackValue {
	_result := &_BACnetConstructedDataBinaryLightingOutputFeedbackValue{
		FeedbackValue: feedbackValue,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBinaryLightingOutputFeedbackValue(structType any) BACnetConstructedDataBinaryLightingOutputFeedbackValue {
    if casted, ok := structType.(BACnetConstructedDataBinaryLightingOutputFeedbackValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBinaryLightingOutputFeedbackValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetTypeName() string {
	return "BACnetConstructedDataBinaryLightingOutputFeedbackValue"
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (feedbackValue)
	lengthInBits += m.FeedbackValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBinaryLightingOutputFeedbackValueParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBinaryLightingOutputFeedbackValue, error) {
	return BACnetConstructedDataBinaryLightingOutputFeedbackValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBinaryLightingOutputFeedbackValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBinaryLightingOutputFeedbackValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBinaryLightingOutputFeedbackValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBinaryLightingOutputFeedbackValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (feedbackValue)
	if pullErr := readBuffer.PullContext("feedbackValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for feedbackValue")
	}
_feedbackValue, _feedbackValueErr := BACnetBinaryLightingPVTaggedParseWithBuffer(ctx, readBuffer , uint8( uint8(0) ) , TagClass( TagClass_APPLICATION_TAGS ) )
	if _feedbackValueErr != nil {
		return nil, errors.Wrap(_feedbackValueErr, "Error parsing 'feedbackValue' field of BACnetConstructedDataBinaryLightingOutputFeedbackValue")
	}
	feedbackValue := _feedbackValue.(BACnetBinaryLightingPVTagged)
	if closeErr := readBuffer.CloseContext("feedbackValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for feedbackValue")
	}

	// Virtual field
	_actualValue := feedbackValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBinaryLightingOutputFeedbackValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBinaryLightingOutputFeedbackValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBinaryLightingOutputFeedbackValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FeedbackValue: feedbackValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBinaryLightingOutputFeedbackValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBinaryLightingOutputFeedbackValue")
		}

	// Simple Field (feedbackValue)
	if pushErr := writeBuffer.PushContext("feedbackValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for feedbackValue")
	}
	_feedbackValueErr := writeBuffer.WriteSerializable(ctx, m.GetFeedbackValue())
	if popErr := writeBuffer.PopContext("feedbackValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for feedbackValue")
	}
	if _feedbackValueErr != nil {
		return errors.Wrap(_feedbackValueErr, "Error serializing 'feedbackValue' field")
	}
	// Virtual field
	actualValue := m.GetActualValue()
	_ =	actualValue
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBinaryLightingOutputFeedbackValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBinaryLightingOutputFeedbackValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) isBACnetConstructedDataBinaryLightingOutputFeedbackValue() bool {
	return true
}

func (m *_BACnetConstructedDataBinaryLightingOutputFeedbackValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



