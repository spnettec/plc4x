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


// BACnetConstructedDataPowerMode is the corresponding interface of BACnetConstructedDataPowerMode
type BACnetConstructedDataPowerMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPowerMode returns PowerMode (property field)
	GetPowerMode() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataPowerModeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataPowerMode.
// This is useful for switch cases.
type BACnetConstructedDataPowerModeExactly interface {
	BACnetConstructedDataPowerMode
	isBACnetConstructedDataPowerMode() bool
}

// _BACnetConstructedDataPowerMode is the data-structure of this message
type _BACnetConstructedDataPowerMode struct {
	*_BACnetConstructedData
        PowerMode BACnetApplicationTagBoolean
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPowerMode)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataPowerMode)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_POWER_MODE}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPowerMode) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataPowerMode)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPowerMode) GetPowerMode() BACnetApplicationTagBoolean {
	return m.PowerMode
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPowerMode) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetPowerMode())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataPowerMode factory function for _BACnetConstructedDataPowerMode
func NewBACnetConstructedDataPowerMode( powerMode BACnetApplicationTagBoolean , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataPowerMode {
	_result := &_BACnetConstructedDataPowerMode{
		PowerMode: powerMode,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPowerMode(structType any) BACnetConstructedDataPowerMode {
    if casted, ok := structType.(BACnetConstructedDataPowerMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPowerMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPowerMode) GetTypeName() string {
	return "BACnetConstructedDataPowerMode"
}

func (m *_BACnetConstructedDataPowerMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (powerMode)
	lengthInBits += m.PowerMode.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataPowerMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataPowerModeParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPowerMode, error) {
	return BACnetConstructedDataPowerModeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataPowerModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPowerMode, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPowerMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPowerMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (powerMode)
	if pullErr := readBuffer.PullContext("powerMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for powerMode")
	}
_powerMode, _powerModeErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _powerModeErr != nil {
		return nil, errors.Wrap(_powerModeErr, "Error parsing 'powerMode' field of BACnetConstructedDataPowerMode")
	}
	powerMode := _powerMode.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("powerMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for powerMode")
	}

	// Virtual field
	_actualValue := powerMode
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPowerMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPowerMode")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataPowerMode{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		PowerMode: powerMode,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataPowerMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPowerMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPowerMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPowerMode")
		}

	// Simple Field (powerMode)
	if pushErr := writeBuffer.PushContext("powerMode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for powerMode")
	}
	_powerModeErr := writeBuffer.WriteSerializable(ctx, m.GetPowerMode())
	if popErr := writeBuffer.PopContext("powerMode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for powerMode")
	}
	if _powerModeErr != nil {
		return errors.Wrap(_powerModeErr, "Error serializing 'powerMode' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPowerMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPowerMode")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataPowerMode) isBACnetConstructedDataPowerMode() bool {
	return true
}

func (m *_BACnetConstructedDataPowerMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



