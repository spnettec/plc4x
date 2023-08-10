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

// BACnetConstructedDataTransition is the corresponding interface of BACnetConstructedDataTransition
type BACnetConstructedDataTransition interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTransition returns Transition (property field)
	GetTransition() BACnetLightingTransitionTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLightingTransitionTagged
}

// BACnetConstructedDataTransitionExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTransition.
// This is useful for switch cases.
type BACnetConstructedDataTransitionExactly interface {
	BACnetConstructedDataTransition
	isBACnetConstructedDataTransition() bool
}

// _BACnetConstructedDataTransition is the data-structure of this message
type _BACnetConstructedDataTransition struct {
	*_BACnetConstructedData
	Transition BACnetLightingTransitionTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTransition) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataTransition) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_TRANSITION
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTransition) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTransition) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTransition) GetTransition() BACnetLightingTransitionTagged {
	return m.Transition
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTransition) GetActualValue() BACnetLightingTransitionTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLightingTransitionTagged(m.GetTransition())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTransition factory function for _BACnetConstructedDataTransition
func NewBACnetConstructedDataTransition(transition BACnetLightingTransitionTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataTransition {
	_result := &_BACnetConstructedDataTransition{
		Transition:             transition,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTransition(structType any) BACnetConstructedDataTransition {
	if casted, ok := structType.(BACnetConstructedDataTransition); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTransition); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTransition) GetTypeName() string {
	return "BACnetConstructedDataTransition"
}

func (m *_BACnetConstructedDataTransition) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (transition)
	lengthInBits += m.Transition.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataTransition) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataTransitionParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTransition, error) {
	return BACnetConstructedDataTransitionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataTransitionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTransition, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTransition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTransition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (transition)
	if pullErr := readBuffer.PullContext("transition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transition")
	}
	_transition, _transitionErr := BACnetLightingTransitionTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _transitionErr != nil {
		return nil, errors.Wrap(_transitionErr, "Error parsing 'transition' field of BACnetConstructedDataTransition")
	}
	transition := _transition.(BACnetLightingTransitionTagged)
	if closeErr := readBuffer.CloseContext("transition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transition")
	}

	// Virtual field
	_actualValue := transition
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTransition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTransition")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTransition{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Transition: transition,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTransition) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTransition) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTransition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTransition")
		}

		// Simple Field (transition)
		if pushErr := writeBuffer.PushContext("transition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for transition")
		}
		_transitionErr := writeBuffer.WriteSerializable(ctx, m.GetTransition())
		if popErr := writeBuffer.PopContext("transition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for transition")
		}
		if _transitionErr != nil {
			return errors.Wrap(_transitionErr, "Error serializing 'transition' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTransition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTransition")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataTransition) isBACnetConstructedDataTransition() bool {
	return true
}

func (m *_BACnetConstructedDataTransition) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
