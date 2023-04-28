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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataEnergyMeter is the corresponding interface of BACnetConstructedDataEnergyMeter
type BACnetConstructedDataEnergyMeter interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetEnergyMeter returns EnergyMeter (property field)
	GetEnergyMeter() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataEnergyMeterExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEnergyMeter.
// This is useful for switch cases.
type BACnetConstructedDataEnergyMeterExactly interface {
	BACnetConstructedDataEnergyMeter
	isBACnetConstructedDataEnergyMeter() bool
}

// _BACnetConstructedDataEnergyMeter is the data-structure of this message
type _BACnetConstructedDataEnergyMeter struct {
	*_BACnetConstructedData
	EnergyMeter BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEnergyMeter) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataEnergyMeter) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ENERGY_METER
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEnergyMeter) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataEnergyMeter) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataEnergyMeter) GetEnergyMeter() BACnetApplicationTagReal {
	return m.EnergyMeter
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataEnergyMeter) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetEnergyMeter())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEnergyMeter factory function for _BACnetConstructedDataEnergyMeter
func NewBACnetConstructedDataEnergyMeter(energyMeter BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataEnergyMeter {
	_result := &_BACnetConstructedDataEnergyMeter{
		EnergyMeter:            energyMeter,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEnergyMeter(structType any) BACnetConstructedDataEnergyMeter {
	if casted, ok := structType.(BACnetConstructedDataEnergyMeter); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEnergyMeter); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEnergyMeter) GetTypeName() string {
	return "BACnetConstructedDataEnergyMeter"
}

func (m *_BACnetConstructedDataEnergyMeter) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (energyMeter)
	lengthInBits += m.EnergyMeter.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataEnergyMeter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataEnergyMeterParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEnergyMeter, error) {
	return BACnetConstructedDataEnergyMeterParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataEnergyMeterParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEnergyMeter, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEnergyMeter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEnergyMeter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (energyMeter)
	if pullErr := readBuffer.PullContext("energyMeter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for energyMeter")
	}
	_energyMeter, _energyMeterErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _energyMeterErr != nil {
		return nil, errors.Wrap(_energyMeterErr, "Error parsing 'energyMeter' field of BACnetConstructedDataEnergyMeter")
	}
	energyMeter := _energyMeter.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("energyMeter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for energyMeter")
	}

	// Virtual field
	_actualValue := energyMeter
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEnergyMeter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEnergyMeter")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataEnergyMeter{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		EnergyMeter: energyMeter,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataEnergyMeter) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataEnergyMeter) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEnergyMeter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEnergyMeter")
		}

		// Simple Field (energyMeter)
		if pushErr := writeBuffer.PushContext("energyMeter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for energyMeter")
		}
		_energyMeterErr := writeBuffer.WriteSerializable(ctx, m.GetEnergyMeter())
		if popErr := writeBuffer.PopContext("energyMeter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for energyMeter")
		}
		if _energyMeterErr != nil {
			return errors.Wrap(_energyMeterErr, "Error serializing 'energyMeter' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEnergyMeter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEnergyMeter")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataEnergyMeter) isBACnetConstructedDataEnergyMeter() bool {
	return true
}

func (m *_BACnetConstructedDataEnergyMeter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
