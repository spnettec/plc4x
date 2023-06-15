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


// BACnetConstructedDataCountBeforeChange is the corresponding interface of BACnetConstructedDataCountBeforeChange
type BACnetConstructedDataCountBeforeChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCountBeforeChange returns CountBeforeChange (property field)
	GetCountBeforeChange() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataCountBeforeChangeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCountBeforeChange.
// This is useful for switch cases.
type BACnetConstructedDataCountBeforeChangeExactly interface {
	BACnetConstructedDataCountBeforeChange
	isBACnetConstructedDataCountBeforeChange() bool
}

// _BACnetConstructedDataCountBeforeChange is the data-structure of this message
type _BACnetConstructedDataCountBeforeChange struct {
	*_BACnetConstructedData
        CountBeforeChange BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCountBeforeChange)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataCountBeforeChange)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_COUNT_BEFORE_CHANGE}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCountBeforeChange) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCountBeforeChange)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCountBeforeChange) GetCountBeforeChange() BACnetApplicationTagUnsignedInteger {
	return m.CountBeforeChange
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCountBeforeChange) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetCountBeforeChange())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataCountBeforeChange factory function for _BACnetConstructedDataCountBeforeChange
func NewBACnetConstructedDataCountBeforeChange( countBeforeChange BACnetApplicationTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataCountBeforeChange {
	_result := &_BACnetConstructedDataCountBeforeChange{
		CountBeforeChange: countBeforeChange,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCountBeforeChange(structType any) BACnetConstructedDataCountBeforeChange {
    if casted, ok := structType.(BACnetConstructedDataCountBeforeChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCountBeforeChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCountBeforeChange) GetTypeName() string {
	return "BACnetConstructedDataCountBeforeChange"
}

func (m *_BACnetConstructedDataCountBeforeChange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (countBeforeChange)
	lengthInBits += m.CountBeforeChange.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataCountBeforeChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataCountBeforeChangeParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCountBeforeChange, error) {
	return BACnetConstructedDataCountBeforeChangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataCountBeforeChangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCountBeforeChange, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCountBeforeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCountBeforeChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (countBeforeChange)
	if pullErr := readBuffer.PullContext("countBeforeChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for countBeforeChange")
	}
_countBeforeChange, _countBeforeChangeErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _countBeforeChangeErr != nil {
		return nil, errors.Wrap(_countBeforeChangeErr, "Error parsing 'countBeforeChange' field of BACnetConstructedDataCountBeforeChange")
	}
	countBeforeChange := _countBeforeChange.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("countBeforeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for countBeforeChange")
	}

	// Virtual field
	_actualValue := countBeforeChange
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCountBeforeChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCountBeforeChange")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCountBeforeChange{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		CountBeforeChange: countBeforeChange,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCountBeforeChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCountBeforeChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCountBeforeChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCountBeforeChange")
		}

	// Simple Field (countBeforeChange)
	if pushErr := writeBuffer.PushContext("countBeforeChange"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for countBeforeChange")
	}
	_countBeforeChangeErr := writeBuffer.WriteSerializable(ctx, m.GetCountBeforeChange())
	if popErr := writeBuffer.PopContext("countBeforeChange"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for countBeforeChange")
	}
	if _countBeforeChangeErr != nil {
		return errors.Wrap(_countBeforeChangeErr, "Error serializing 'countBeforeChange' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCountBeforeChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCountBeforeChange")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataCountBeforeChange) isBACnetConstructedDataCountBeforeChange() bool {
	return true
}

func (m *_BACnetConstructedDataCountBeforeChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



