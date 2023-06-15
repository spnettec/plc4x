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


// BACnetConstructedDataCarMode is the corresponding interface of BACnetConstructedDataCarMode
type BACnetConstructedDataCarMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCarMode returns CarMode (property field)
	GetCarMode() BACnetLiftCarModeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLiftCarModeTagged
}

// BACnetConstructedDataCarModeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCarMode.
// This is useful for switch cases.
type BACnetConstructedDataCarModeExactly interface {
	BACnetConstructedDataCarMode
	isBACnetConstructedDataCarMode() bool
}

// _BACnetConstructedDataCarMode is the data-structure of this message
type _BACnetConstructedDataCarMode struct {
	*_BACnetConstructedData
        CarMode BACnetLiftCarModeTagged
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCarMode)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataCarMode)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_CAR_MODE}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCarMode) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCarMode)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCarMode) GetCarMode() BACnetLiftCarModeTagged {
	return m.CarMode
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCarMode) GetActualValue() BACnetLiftCarModeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLiftCarModeTagged(m.GetCarMode())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataCarMode factory function for _BACnetConstructedDataCarMode
func NewBACnetConstructedDataCarMode( carMode BACnetLiftCarModeTagged , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataCarMode {
	_result := &_BACnetConstructedDataCarMode{
		CarMode: carMode,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCarMode(structType any) BACnetConstructedDataCarMode {
    if casted, ok := structType.(BACnetConstructedDataCarMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCarMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCarMode) GetTypeName() string {
	return "BACnetConstructedDataCarMode"
}

func (m *_BACnetConstructedDataCarMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (carMode)
	lengthInBits += m.CarMode.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataCarMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataCarModeParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCarMode, error) {
	return BACnetConstructedDataCarModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataCarModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCarMode, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCarMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCarMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (carMode)
	if pullErr := readBuffer.PullContext("carMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for carMode")
	}
_carMode, _carModeErr := BACnetLiftCarModeTaggedParseWithBuffer(ctx, readBuffer , uint8( uint8(0) ) , TagClass( TagClass_APPLICATION_TAGS ) )
	if _carModeErr != nil {
		return nil, errors.Wrap(_carModeErr, "Error parsing 'carMode' field of BACnetConstructedDataCarMode")
	}
	carMode := _carMode.(BACnetLiftCarModeTagged)
	if closeErr := readBuffer.CloseContext("carMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for carMode")
	}

	// Virtual field
	_actualValue := carMode
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCarMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCarMode")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCarMode{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		CarMode: carMode,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCarMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataCarMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCarMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCarMode")
		}

	// Simple Field (carMode)
	if pushErr := writeBuffer.PushContext("carMode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for carMode")
	}
	_carModeErr := writeBuffer.WriteSerializable(ctx, m.GetCarMode())
	if popErr := writeBuffer.PopContext("carMode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for carMode")
	}
	if _carModeErr != nil {
		return errors.Wrap(_carModeErr, "Error serializing 'carMode' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCarMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCarMode")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataCarMode) isBACnetConstructedDataCarMode() bool {
	return true
}

func (m *_BACnetConstructedDataCarMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



