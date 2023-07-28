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


// BACnetConstructedDataStateDescription is the corresponding interface of BACnetConstructedDataStateDescription
type BACnetConstructedDataStateDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetStateDescription returns StateDescription (property field)
	GetStateDescription() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataStateDescriptionExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataStateDescription.
// This is useful for switch cases.
type BACnetConstructedDataStateDescriptionExactly interface {
	BACnetConstructedDataStateDescription
	isBACnetConstructedDataStateDescription() bool
}

// _BACnetConstructedDataStateDescription is the data-structure of this message
type _BACnetConstructedDataStateDescription struct {
	*_BACnetConstructedData
        StateDescription BACnetApplicationTagCharacterString
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataStateDescription)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataStateDescription)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_STATE_DESCRIPTION}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataStateDescription) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataStateDescription)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataStateDescription) GetStateDescription() BACnetApplicationTagCharacterString {
	return m.StateDescription
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataStateDescription) GetActualValue() BACnetApplicationTagCharacterString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagCharacterString(m.GetStateDescription())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataStateDescription factory function for _BACnetConstructedDataStateDescription
func NewBACnetConstructedDataStateDescription( stateDescription BACnetApplicationTagCharacterString , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataStateDescription {
	_result := &_BACnetConstructedDataStateDescription{
		StateDescription: stateDescription,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataStateDescription(structType any) BACnetConstructedDataStateDescription {
    if casted, ok := structType.(BACnetConstructedDataStateDescription); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataStateDescription); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataStateDescription) GetTypeName() string {
	return "BACnetConstructedDataStateDescription"
}

func (m *_BACnetConstructedDataStateDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (stateDescription)
	lengthInBits += m.StateDescription.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataStateDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataStateDescriptionParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataStateDescription, error) {
	return BACnetConstructedDataStateDescriptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataStateDescriptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataStateDescription, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataStateDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataStateDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (stateDescription)
	if pullErr := readBuffer.PullContext("stateDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for stateDescription")
	}
_stateDescription, _stateDescriptionErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _stateDescriptionErr != nil {
		return nil, errors.Wrap(_stateDescriptionErr, "Error parsing 'stateDescription' field of BACnetConstructedDataStateDescription")
	}
	stateDescription := _stateDescription.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("stateDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for stateDescription")
	}

	// Virtual field
	_actualValue := stateDescription
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataStateDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataStateDescription")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataStateDescription{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		StateDescription: stateDescription,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataStateDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataStateDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataStateDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataStateDescription")
		}

	// Simple Field (stateDescription)
	if pushErr := writeBuffer.PushContext("stateDescription"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for stateDescription")
	}
	_stateDescriptionErr := writeBuffer.WriteSerializable(ctx, m.GetStateDescription())
	if popErr := writeBuffer.PopContext("stateDescription"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for stateDescription")
	}
	if _stateDescriptionErr != nil {
		return errors.Wrap(_stateDescriptionErr, "Error serializing 'stateDescription' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataStateDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataStateDescription")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataStateDescription) isBACnetConstructedDataStateDescription() bool {
	return true
}

func (m *_BACnetConstructedDataStateDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



