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


// BACnetConstructedDataCurrentCommandPriority is the corresponding interface of BACnetConstructedDataCurrentCommandPriority
type BACnetConstructedDataCurrentCommandPriority interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCurrentCommandPriority returns CurrentCommandPriority (property field)
	GetCurrentCommandPriority() BACnetOptionalUnsigned
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetOptionalUnsigned
}

// BACnetConstructedDataCurrentCommandPriorityExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCurrentCommandPriority.
// This is useful for switch cases.
type BACnetConstructedDataCurrentCommandPriorityExactly interface {
	BACnetConstructedDataCurrentCommandPriority
	isBACnetConstructedDataCurrentCommandPriority() bool
}

// _BACnetConstructedDataCurrentCommandPriority is the data-structure of this message
type _BACnetConstructedDataCurrentCommandPriority struct {
	*_BACnetConstructedData
        CurrentCommandPriority BACnetOptionalUnsigned
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCurrentCommandPriority)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataCurrentCommandPriority)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_CURRENT_COMMAND_PRIORITY}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCurrentCommandPriority) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCurrentCommandPriority)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCurrentCommandPriority) GetCurrentCommandPriority() BACnetOptionalUnsigned {
	return m.CurrentCommandPriority
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCurrentCommandPriority) GetActualValue() BACnetOptionalUnsigned {
	ctx := context.Background()
	_ = ctx
	return CastBACnetOptionalUnsigned(m.GetCurrentCommandPriority())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataCurrentCommandPriority factory function for _BACnetConstructedDataCurrentCommandPriority
func NewBACnetConstructedDataCurrentCommandPriority( currentCommandPriority BACnetOptionalUnsigned , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataCurrentCommandPriority {
	_result := &_BACnetConstructedDataCurrentCommandPriority{
		CurrentCommandPriority: currentCommandPriority,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCurrentCommandPriority(structType any) BACnetConstructedDataCurrentCommandPriority {
    if casted, ok := structType.(BACnetConstructedDataCurrentCommandPriority); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCurrentCommandPriority); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCurrentCommandPriority) GetTypeName() string {
	return "BACnetConstructedDataCurrentCommandPriority"
}

func (m *_BACnetConstructedDataCurrentCommandPriority) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (currentCommandPriority)
	lengthInBits += m.CurrentCommandPriority.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataCurrentCommandPriority) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataCurrentCommandPriorityParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCurrentCommandPriority, error) {
	return BACnetConstructedDataCurrentCommandPriorityParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataCurrentCommandPriorityParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCurrentCommandPriority, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCurrentCommandPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCurrentCommandPriority")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (currentCommandPriority)
	if pullErr := readBuffer.PullContext("currentCommandPriority"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for currentCommandPriority")
	}
_currentCommandPriority, _currentCommandPriorityErr := BACnetOptionalUnsignedParseWithBuffer(ctx, readBuffer)
	if _currentCommandPriorityErr != nil {
		return nil, errors.Wrap(_currentCommandPriorityErr, "Error parsing 'currentCommandPriority' field of BACnetConstructedDataCurrentCommandPriority")
	}
	currentCommandPriority := _currentCommandPriority.(BACnetOptionalUnsigned)
	if closeErr := readBuffer.CloseContext("currentCommandPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for currentCommandPriority")
	}

	// Virtual field
	_actualValue := currentCommandPriority
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCurrentCommandPriority"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCurrentCommandPriority")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCurrentCommandPriority{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		CurrentCommandPriority: currentCommandPriority,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCurrentCommandPriority) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCurrentCommandPriority) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCurrentCommandPriority"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCurrentCommandPriority")
		}

	// Simple Field (currentCommandPriority)
	if pushErr := writeBuffer.PushContext("currentCommandPriority"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for currentCommandPriority")
	}
	_currentCommandPriorityErr := writeBuffer.WriteSerializable(ctx, m.GetCurrentCommandPriority())
	if popErr := writeBuffer.PopContext("currentCommandPriority"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for currentCommandPriority")
	}
	if _currentCommandPriorityErr != nil {
		return errors.Wrap(_currentCommandPriorityErr, "Error serializing 'currentCommandPriority' field")
	}
	// Virtual field
	actualValue := m.GetActualValue()
	_ =	actualValue
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCurrentCommandPriority"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCurrentCommandPriority")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataCurrentCommandPriority) isBACnetConstructedDataCurrentCommandPriority() bool {
	return true
}

func (m *_BACnetConstructedDataCurrentCommandPriority) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



