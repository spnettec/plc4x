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


// BACnetConstructedDataAnalogOutputMaxPresValue is the corresponding interface of BACnetConstructedDataAnalogOutputMaxPresValue
type BACnetConstructedDataAnalogOutputMaxPresValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaxPresValue returns MaxPresValue (property field)
	GetMaxPresValue() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataAnalogOutputMaxPresValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAnalogOutputMaxPresValue.
// This is useful for switch cases.
type BACnetConstructedDataAnalogOutputMaxPresValueExactly interface {
	BACnetConstructedDataAnalogOutputMaxPresValue
	isBACnetConstructedDataAnalogOutputMaxPresValue() bool
}

// _BACnetConstructedDataAnalogOutputMaxPresValue is the data-structure of this message
type _BACnetConstructedDataAnalogOutputMaxPresValue struct {
	*_BACnetConstructedData
        MaxPresValue BACnetApplicationTagReal
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue)  GetObjectTypeArgument() BACnetObjectType {
return BACnetObjectType_ANALOG_OUTPUT}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_MAX_PRES_VALUE}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) GetMaxPresValue() BACnetApplicationTagReal {
	return m.MaxPresValue
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetMaxPresValue())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataAnalogOutputMaxPresValue factory function for _BACnetConstructedDataAnalogOutputMaxPresValue
func NewBACnetConstructedDataAnalogOutputMaxPresValue( maxPresValue BACnetApplicationTagReal , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataAnalogOutputMaxPresValue {
	_result := &_BACnetConstructedDataAnalogOutputMaxPresValue{
		MaxPresValue: maxPresValue,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAnalogOutputMaxPresValue(structType any) BACnetConstructedDataAnalogOutputMaxPresValue {
    if casted, ok := structType.(BACnetConstructedDataAnalogOutputMaxPresValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAnalogOutputMaxPresValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) GetTypeName() string {
	return "BACnetConstructedDataAnalogOutputMaxPresValue"
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (maxPresValue)
	lengthInBits += m.MaxPresValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAnalogOutputMaxPresValueParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAnalogOutputMaxPresValue, error) {
	return BACnetConstructedDataAnalogOutputMaxPresValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAnalogOutputMaxPresValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAnalogOutputMaxPresValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAnalogOutputMaxPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAnalogOutputMaxPresValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maxPresValue)
	if pullErr := readBuffer.PullContext("maxPresValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxPresValue")
	}
_maxPresValue, _maxPresValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _maxPresValueErr != nil {
		return nil, errors.Wrap(_maxPresValueErr, "Error parsing 'maxPresValue' field of BACnetConstructedDataAnalogOutputMaxPresValue")
	}
	maxPresValue := _maxPresValue.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("maxPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxPresValue")
	}

	// Virtual field
	_actualValue := maxPresValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAnalogOutputMaxPresValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAnalogOutputMaxPresValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAnalogOutputMaxPresValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MaxPresValue: maxPresValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAnalogOutputMaxPresValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAnalogOutputMaxPresValue")
		}

	// Simple Field (maxPresValue)
	if pushErr := writeBuffer.PushContext("maxPresValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for maxPresValue")
	}
	_maxPresValueErr := writeBuffer.WriteSerializable(ctx, m.GetMaxPresValue())
	if popErr := writeBuffer.PopContext("maxPresValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for maxPresValue")
	}
	if _maxPresValueErr != nil {
		return errors.Wrap(_maxPresValueErr, "Error serializing 'maxPresValue' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAnalogOutputMaxPresValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAnalogOutputMaxPresValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) isBACnetConstructedDataAnalogOutputMaxPresValue() bool {
	return true
}

func (m *_BACnetConstructedDataAnalogOutputMaxPresValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



