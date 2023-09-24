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


// BACnetConstructedDataSetpointReference is the corresponding interface of BACnetConstructedDataSetpointReference
type BACnetConstructedDataSetpointReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSetpointReference returns SetpointReference (property field)
	GetSetpointReference() BACnetSetpointReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetSetpointReference
}

// BACnetConstructedDataSetpointReferenceExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSetpointReference.
// This is useful for switch cases.
type BACnetConstructedDataSetpointReferenceExactly interface {
	BACnetConstructedDataSetpointReference
	isBACnetConstructedDataSetpointReference() bool
}

// _BACnetConstructedDataSetpointReference is the data-structure of this message
type _BACnetConstructedDataSetpointReference struct {
	*_BACnetConstructedData
        SetpointReference BACnetSetpointReference
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSetpointReference)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataSetpointReference)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_SETPOINT_REFERENCE}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSetpointReference) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSetpointReference)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSetpointReference) GetSetpointReference() BACnetSetpointReference {
	return m.SetpointReference
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSetpointReference) GetActualValue() BACnetSetpointReference {
	ctx := context.Background()
	_ = ctx
	return CastBACnetSetpointReference(m.GetSetpointReference())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataSetpointReference factory function for _BACnetConstructedDataSetpointReference
func NewBACnetConstructedDataSetpointReference( setpointReference BACnetSetpointReference , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataSetpointReference {
	_result := &_BACnetConstructedDataSetpointReference{
		SetpointReference: setpointReference,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSetpointReference(structType any) BACnetConstructedDataSetpointReference {
    if casted, ok := structType.(BACnetConstructedDataSetpointReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSetpointReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSetpointReference) GetTypeName() string {
	return "BACnetConstructedDataSetpointReference"
}

func (m *_BACnetConstructedDataSetpointReference) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (setpointReference)
	lengthInBits += m.SetpointReference.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataSetpointReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataSetpointReferenceParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSetpointReference, error) {
	return BACnetConstructedDataSetpointReferenceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataSetpointReferenceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSetpointReference, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSetpointReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSetpointReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (setpointReference)
	if pullErr := readBuffer.PullContext("setpointReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for setpointReference")
	}
_setpointReference, _setpointReferenceErr := BACnetSetpointReferenceParseWithBuffer(ctx, readBuffer)
	if _setpointReferenceErr != nil {
		return nil, errors.Wrap(_setpointReferenceErr, "Error parsing 'setpointReference' field of BACnetConstructedDataSetpointReference")
	}
	setpointReference := _setpointReference.(BACnetSetpointReference)
	if closeErr := readBuffer.CloseContext("setpointReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for setpointReference")
	}

	// Virtual field
	_actualValue := setpointReference
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSetpointReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSetpointReference")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSetpointReference{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		SetpointReference: setpointReference,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSetpointReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSetpointReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSetpointReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSetpointReference")
		}

	// Simple Field (setpointReference)
	if pushErr := writeBuffer.PushContext("setpointReference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for setpointReference")
	}
	_setpointReferenceErr := writeBuffer.WriteSerializable(ctx, m.GetSetpointReference())
	if popErr := writeBuffer.PopContext("setpointReference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for setpointReference")
	}
	if _setpointReferenceErr != nil {
		return errors.Wrap(_setpointReferenceErr, "Error serializing 'setpointReference' field")
	}
	// Virtual field
	actualValue := m.GetActualValue()
	_ =	actualValue
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSetpointReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSetpointReference")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataSetpointReference) isBACnetConstructedDataSetpointReference() bool {
	return true
}

func (m *_BACnetConstructedDataSetpointReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



