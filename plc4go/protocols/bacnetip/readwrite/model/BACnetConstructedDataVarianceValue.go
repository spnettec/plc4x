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


// BACnetConstructedDataVarianceValue is the corresponding interface of BACnetConstructedDataVarianceValue
type BACnetConstructedDataVarianceValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetVarianceValue returns VarianceValue (property field)
	GetVarianceValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataVarianceValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataVarianceValue.
// This is useful for switch cases.
type BACnetConstructedDataVarianceValueExactly interface {
	BACnetConstructedDataVarianceValue
	isBACnetConstructedDataVarianceValue() bool
}

// _BACnetConstructedDataVarianceValue is the data-structure of this message
type _BACnetConstructedDataVarianceValue struct {
	*_BACnetConstructedData
        VarianceValue BACnetApplicationTagReal
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataVarianceValue)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataVarianceValue)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_VARIANCE_VALUE}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataVarianceValue) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataVarianceValue)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataVarianceValue) GetVarianceValue() BACnetApplicationTagReal {
	return m.VarianceValue
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataVarianceValue) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetVarianceValue())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataVarianceValue factory function for _BACnetConstructedDataVarianceValue
func NewBACnetConstructedDataVarianceValue( varianceValue BACnetApplicationTagReal , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataVarianceValue {
	_result := &_BACnetConstructedDataVarianceValue{
		VarianceValue: varianceValue,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataVarianceValue(structType any) BACnetConstructedDataVarianceValue {
    if casted, ok := structType.(BACnetConstructedDataVarianceValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVarianceValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataVarianceValue) GetTypeName() string {
	return "BACnetConstructedDataVarianceValue"
}

func (m *_BACnetConstructedDataVarianceValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (varianceValue)
	lengthInBits += m.VarianceValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataVarianceValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataVarianceValueParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataVarianceValue, error) {
	return BACnetConstructedDataVarianceValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataVarianceValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataVarianceValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVarianceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataVarianceValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (varianceValue)
	if pullErr := readBuffer.PullContext("varianceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for varianceValue")
	}
_varianceValue, _varianceValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _varianceValueErr != nil {
		return nil, errors.Wrap(_varianceValueErr, "Error parsing 'varianceValue' field of BACnetConstructedDataVarianceValue")
	}
	varianceValue := _varianceValue.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("varianceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for varianceValue")
	}

	// Virtual field
	_actualValue := varianceValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVarianceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataVarianceValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataVarianceValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		VarianceValue: varianceValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataVarianceValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataVarianceValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVarianceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataVarianceValue")
		}

	// Simple Field (varianceValue)
	if pushErr := writeBuffer.PushContext("varianceValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for varianceValue")
	}
	_varianceValueErr := writeBuffer.WriteSerializable(ctx, m.GetVarianceValue())
	if popErr := writeBuffer.PopContext("varianceValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for varianceValue")
	}
	if _varianceValueErr != nil {
		return errors.Wrap(_varianceValueErr, "Error serializing 'varianceValue' field")
	}
	// Virtual field
	actualValue := m.GetActualValue()
	_ =	actualValue
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVarianceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataVarianceValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataVarianceValue) isBACnetConstructedDataVarianceValue() bool {
	return true
}

func (m *_BACnetConstructedDataVarianceValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



