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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataControlledVariableReference is the corresponding interface of BACnetConstructedDataControlledVariableReference
type BACnetConstructedDataControlledVariableReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetControlledVariableReference returns ControlledVariableReference (property field)
	GetControlledVariableReference() BACnetObjectPropertyReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetObjectPropertyReference
	// IsBACnetConstructedDataControlledVariableReference is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataControlledVariableReference()
}

// _BACnetConstructedDataControlledVariableReference is the data-structure of this message
type _BACnetConstructedDataControlledVariableReference struct {
	BACnetConstructedDataContract
	ControlledVariableReference BACnetObjectPropertyReference
}

var _ BACnetConstructedDataControlledVariableReference = (*_BACnetConstructedDataControlledVariableReference)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataControlledVariableReference)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableReference) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataControlledVariableReference) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CONTROLLED_VARIABLE_REFERENCE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataControlledVariableReference) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableReference) GetControlledVariableReference() BACnetObjectPropertyReference {
	return m.ControlledVariableReference
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataControlledVariableReference) GetActualValue() BACnetObjectPropertyReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetObjectPropertyReference(m.GetControlledVariableReference())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataControlledVariableReference factory function for _BACnetConstructedDataControlledVariableReference
func NewBACnetConstructedDataControlledVariableReference(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, controlledVariableReference BACnetObjectPropertyReference, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataControlledVariableReference {
	if controlledVariableReference == nil {
		panic("controlledVariableReference of type BACnetObjectPropertyReference for BACnetConstructedDataControlledVariableReference must not be nil")
	}
	_result := &_BACnetConstructedDataControlledVariableReference{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ControlledVariableReference:   controlledVariableReference,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataControlledVariableReference(structType any) BACnetConstructedDataControlledVariableReference {
	if casted, ok := structType.(BACnetConstructedDataControlledVariableReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataControlledVariableReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataControlledVariableReference) GetTypeName() string {
	return "BACnetConstructedDataControlledVariableReference"
}

func (m *_BACnetConstructedDataControlledVariableReference) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (controlledVariableReference)
	lengthInBits += m.ControlledVariableReference.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataControlledVariableReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataControlledVariableReference) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataControlledVariableReference BACnetConstructedDataControlledVariableReference, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataControlledVariableReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataControlledVariableReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	controlledVariableReference, err := ReadSimpleField[BACnetObjectPropertyReference](ctx, "controlledVariableReference", ReadComplex[BACnetObjectPropertyReference](BACnetObjectPropertyReferenceParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'controlledVariableReference' field"))
	}
	m.ControlledVariableReference = controlledVariableReference

	actualValue, err := ReadVirtualField[BACnetObjectPropertyReference](ctx, "actualValue", (*BACnetObjectPropertyReference)(nil), controlledVariableReference)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataControlledVariableReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataControlledVariableReference")
	}

	return m, nil
}

func (m *_BACnetConstructedDataControlledVariableReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataControlledVariableReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataControlledVariableReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataControlledVariableReference")
		}

		if err := WriteSimpleField[BACnetObjectPropertyReference](ctx, "controlledVariableReference", m.GetControlledVariableReference(), WriteComplex[BACnetObjectPropertyReference](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'controlledVariableReference' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataControlledVariableReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataControlledVariableReference")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataControlledVariableReference) IsBACnetConstructedDataControlledVariableReference() {
}

func (m *_BACnetConstructedDataControlledVariableReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
