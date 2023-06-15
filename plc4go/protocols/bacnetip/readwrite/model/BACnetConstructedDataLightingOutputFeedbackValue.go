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


// BACnetConstructedDataLightingOutputFeedbackValue is the corresponding interface of BACnetConstructedDataLightingOutputFeedbackValue
type BACnetConstructedDataLightingOutputFeedbackValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFeedbackValue returns FeedbackValue (property field)
	GetFeedbackValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataLightingOutputFeedbackValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLightingOutputFeedbackValue.
// This is useful for switch cases.
type BACnetConstructedDataLightingOutputFeedbackValueExactly interface {
	BACnetConstructedDataLightingOutputFeedbackValue
	isBACnetConstructedDataLightingOutputFeedbackValue() bool
}

// _BACnetConstructedDataLightingOutputFeedbackValue is the data-structure of this message
type _BACnetConstructedDataLightingOutputFeedbackValue struct {
	*_BACnetConstructedData
        FeedbackValue BACnetApplicationTagReal
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLightingOutputFeedbackValue)  GetObjectTypeArgument() BACnetObjectType {
return BACnetObjectType_LIGHTING_OUTPUT}

func (m *_BACnetConstructedDataLightingOutputFeedbackValue)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_FEEDBACK_VALUE}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLightingOutputFeedbackValue) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLightingOutputFeedbackValue)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLightingOutputFeedbackValue) GetFeedbackValue() BACnetApplicationTagReal {
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

func (m *_BACnetConstructedDataLightingOutputFeedbackValue) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetFeedbackValue())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataLightingOutputFeedbackValue factory function for _BACnetConstructedDataLightingOutputFeedbackValue
func NewBACnetConstructedDataLightingOutputFeedbackValue( feedbackValue BACnetApplicationTagReal , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataLightingOutputFeedbackValue {
	_result := &_BACnetConstructedDataLightingOutputFeedbackValue{
		FeedbackValue: feedbackValue,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLightingOutputFeedbackValue(structType any) BACnetConstructedDataLightingOutputFeedbackValue {
    if casted, ok := structType.(BACnetConstructedDataLightingOutputFeedbackValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLightingOutputFeedbackValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLightingOutputFeedbackValue) GetTypeName() string {
	return "BACnetConstructedDataLightingOutputFeedbackValue"
}

func (m *_BACnetConstructedDataLightingOutputFeedbackValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (feedbackValue)
	lengthInBits += m.FeedbackValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataLightingOutputFeedbackValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLightingOutputFeedbackValueParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLightingOutputFeedbackValue, error) {
	return BACnetConstructedDataLightingOutputFeedbackValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLightingOutputFeedbackValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLightingOutputFeedbackValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLightingOutputFeedbackValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLightingOutputFeedbackValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (feedbackValue)
	if pullErr := readBuffer.PullContext("feedbackValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for feedbackValue")
	}
_feedbackValue, _feedbackValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _feedbackValueErr != nil {
		return nil, errors.Wrap(_feedbackValueErr, "Error parsing 'feedbackValue' field of BACnetConstructedDataLightingOutputFeedbackValue")
	}
	feedbackValue := _feedbackValue.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("feedbackValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for feedbackValue")
	}

	// Virtual field
	_actualValue := feedbackValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLightingOutputFeedbackValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLightingOutputFeedbackValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLightingOutputFeedbackValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FeedbackValue: feedbackValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLightingOutputFeedbackValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLightingOutputFeedbackValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLightingOutputFeedbackValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLightingOutputFeedbackValue")
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
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLightingOutputFeedbackValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLightingOutputFeedbackValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataLightingOutputFeedbackValue) isBACnetConstructedDataLightingOutputFeedbackValue() bool {
	return true
}

func (m *_BACnetConstructedDataLightingOutputFeedbackValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



