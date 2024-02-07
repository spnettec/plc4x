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

// BACnetConstructedDataObjectPropertyReference is the corresponding interface of BACnetConstructedDataObjectPropertyReference
type BACnetConstructedDataObjectPropertyReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPropertyReference returns PropertyReference (property field)
	GetPropertyReference() BACnetDeviceObjectPropertyReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectPropertyReference
}

// BACnetConstructedDataObjectPropertyReferenceExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataObjectPropertyReference.
// This is useful for switch cases.
type BACnetConstructedDataObjectPropertyReferenceExactly interface {
	BACnetConstructedDataObjectPropertyReference
	isBACnetConstructedDataObjectPropertyReference() bool
}

// _BACnetConstructedDataObjectPropertyReference is the data-structure of this message
type _BACnetConstructedDataObjectPropertyReference struct {
	*_BACnetConstructedData
	PropertyReference BACnetDeviceObjectPropertyReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataObjectPropertyReference) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataObjectPropertyReference) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_OBJECT_PROPERTY_REFERENCE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataObjectPropertyReference) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataObjectPropertyReference) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataObjectPropertyReference) GetPropertyReference() BACnetDeviceObjectPropertyReference {
	return m.PropertyReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataObjectPropertyReference) GetActualValue() BACnetDeviceObjectPropertyReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDeviceObjectPropertyReference(m.GetPropertyReference())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataObjectPropertyReference factory function for _BACnetConstructedDataObjectPropertyReference
func NewBACnetConstructedDataObjectPropertyReference(propertyReference BACnetDeviceObjectPropertyReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataObjectPropertyReference {
	_result := &_BACnetConstructedDataObjectPropertyReference{
		PropertyReference:      propertyReference,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataObjectPropertyReference(structType any) BACnetConstructedDataObjectPropertyReference {
	if casted, ok := structType.(BACnetConstructedDataObjectPropertyReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataObjectPropertyReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataObjectPropertyReference) GetTypeName() string {
	return "BACnetConstructedDataObjectPropertyReference"
}

func (m *_BACnetConstructedDataObjectPropertyReference) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (propertyReference)
	lengthInBits += m.PropertyReference.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataObjectPropertyReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataObjectPropertyReferenceParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataObjectPropertyReference, error) {
	return BACnetConstructedDataObjectPropertyReferenceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataObjectPropertyReferenceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataObjectPropertyReference, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataObjectPropertyReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataObjectPropertyReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (propertyReference)
	if pullErr := readBuffer.PullContext("propertyReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyReference")
	}
	_propertyReference, _propertyReferenceErr := BACnetDeviceObjectPropertyReferenceParseWithBuffer(ctx, readBuffer)
	if _propertyReferenceErr != nil {
		return nil, errors.Wrap(_propertyReferenceErr, "Error parsing 'propertyReference' field of BACnetConstructedDataObjectPropertyReference")
	}
	propertyReference := _propertyReference.(BACnetDeviceObjectPropertyReference)
	if closeErr := readBuffer.CloseContext("propertyReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyReference")
	}

	// Virtual field
	_actualValue := propertyReference
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataObjectPropertyReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataObjectPropertyReference")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataObjectPropertyReference{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		PropertyReference: propertyReference,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataObjectPropertyReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataObjectPropertyReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataObjectPropertyReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataObjectPropertyReference")
		}

		// Simple Field (propertyReference)
		if pushErr := writeBuffer.PushContext("propertyReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyReference")
		}
		_propertyReferenceErr := writeBuffer.WriteSerializable(ctx, m.GetPropertyReference())
		if popErr := writeBuffer.PopContext("propertyReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyReference")
		}
		if _propertyReferenceErr != nil {
			return errors.Wrap(_propertyReferenceErr, "Error serializing 'propertyReference' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataObjectPropertyReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataObjectPropertyReference")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataObjectPropertyReference) isBACnetConstructedDataObjectPropertyReference() bool {
	return true
}

func (m *_BACnetConstructedDataObjectPropertyReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
