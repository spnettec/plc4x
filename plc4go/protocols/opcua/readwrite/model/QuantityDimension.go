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

// QuantityDimension is the corresponding interface of QuantityDimension
type QuantityDimension interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetMassExponent returns MassExponent (property field)
	GetMassExponent() int8
	// GetLengthExponent returns LengthExponent (property field)
	GetLengthExponent() int8
	// GetTimeExponent returns TimeExponent (property field)
	GetTimeExponent() int8
	// GetElectricCurrentExponent returns ElectricCurrentExponent (property field)
	GetElectricCurrentExponent() int8
	// GetAmountOfSubstanceExponent returns AmountOfSubstanceExponent (property field)
	GetAmountOfSubstanceExponent() int8
	// GetLuminousIntensityExponent returns LuminousIntensityExponent (property field)
	GetLuminousIntensityExponent() int8
	// GetAbsoluteTemperatureExponent returns AbsoluteTemperatureExponent (property field)
	GetAbsoluteTemperatureExponent() int8
	// GetDimensionlessExponent returns DimensionlessExponent (property field)
	GetDimensionlessExponent() int8
}

// QuantityDimensionExactly can be used when we want exactly this type and not a type which fulfills QuantityDimension.
// This is useful for switch cases.
type QuantityDimensionExactly interface {
	QuantityDimension
	isQuantityDimension() bool
}

// _QuantityDimension is the data-structure of this message
type _QuantityDimension struct {
	*_ExtensionObjectDefinition
	MassExponent                int8
	LengthExponent              int8
	TimeExponent                int8
	ElectricCurrentExponent     int8
	AmountOfSubstanceExponent   int8
	LuminousIntensityExponent   int8
	AbsoluteTemperatureExponent int8
	DimensionlessExponent       int8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_QuantityDimension) GetIdentifier() string {
	return "32440"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_QuantityDimension) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_QuantityDimension) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_QuantityDimension) GetMassExponent() int8 {
	return m.MassExponent
}

func (m *_QuantityDimension) GetLengthExponent() int8 {
	return m.LengthExponent
}

func (m *_QuantityDimension) GetTimeExponent() int8 {
	return m.TimeExponent
}

func (m *_QuantityDimension) GetElectricCurrentExponent() int8 {
	return m.ElectricCurrentExponent
}

func (m *_QuantityDimension) GetAmountOfSubstanceExponent() int8 {
	return m.AmountOfSubstanceExponent
}

func (m *_QuantityDimension) GetLuminousIntensityExponent() int8 {
	return m.LuminousIntensityExponent
}

func (m *_QuantityDimension) GetAbsoluteTemperatureExponent() int8 {
	return m.AbsoluteTemperatureExponent
}

func (m *_QuantityDimension) GetDimensionlessExponent() int8 {
	return m.DimensionlessExponent
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewQuantityDimension factory function for _QuantityDimension
func NewQuantityDimension(massExponent int8, lengthExponent int8, timeExponent int8, electricCurrentExponent int8, amountOfSubstanceExponent int8, luminousIntensityExponent int8, absoluteTemperatureExponent int8, dimensionlessExponent int8) *_QuantityDimension {
	_result := &_QuantityDimension{
		MassExponent:                massExponent,
		LengthExponent:              lengthExponent,
		TimeExponent:                timeExponent,
		ElectricCurrentExponent:     electricCurrentExponent,
		AmountOfSubstanceExponent:   amountOfSubstanceExponent,
		LuminousIntensityExponent:   luminousIntensityExponent,
		AbsoluteTemperatureExponent: absoluteTemperatureExponent,
		DimensionlessExponent:       dimensionlessExponent,
		_ExtensionObjectDefinition:  NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastQuantityDimension(structType any) QuantityDimension {
	if casted, ok := structType.(QuantityDimension); ok {
		return casted
	}
	if casted, ok := structType.(*QuantityDimension); ok {
		return *casted
	}
	return nil
}

func (m *_QuantityDimension) GetTypeName() string {
	return "QuantityDimension"
}

func (m *_QuantityDimension) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (massExponent)
	lengthInBits += 8

	// Simple field (lengthExponent)
	lengthInBits += 8

	// Simple field (timeExponent)
	lengthInBits += 8

	// Simple field (electricCurrentExponent)
	lengthInBits += 8

	// Simple field (amountOfSubstanceExponent)
	lengthInBits += 8

	// Simple field (luminousIntensityExponent)
	lengthInBits += 8

	// Simple field (absoluteTemperatureExponent)
	lengthInBits += 8

	// Simple field (dimensionlessExponent)
	lengthInBits += 8

	return lengthInBits
}

func (m *_QuantityDimension) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func QuantityDimensionParse(ctx context.Context, theBytes []byte, identifier string) (QuantityDimension, error) {
	return QuantityDimensionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func QuantityDimensionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (QuantityDimension, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("QuantityDimension"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for QuantityDimension")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (massExponent)
	_massExponent, _massExponentErr := readBuffer.ReadInt8("massExponent", 8)
	if _massExponentErr != nil {
		return nil, errors.Wrap(_massExponentErr, "Error parsing 'massExponent' field of QuantityDimension")
	}
	massExponent := _massExponent

	// Simple Field (lengthExponent)
	_lengthExponent, _lengthExponentErr := readBuffer.ReadInt8("lengthExponent", 8)
	if _lengthExponentErr != nil {
		return nil, errors.Wrap(_lengthExponentErr, "Error parsing 'lengthExponent' field of QuantityDimension")
	}
	lengthExponent := _lengthExponent

	// Simple Field (timeExponent)
	_timeExponent, _timeExponentErr := readBuffer.ReadInt8("timeExponent", 8)
	if _timeExponentErr != nil {
		return nil, errors.Wrap(_timeExponentErr, "Error parsing 'timeExponent' field of QuantityDimension")
	}
	timeExponent := _timeExponent

	// Simple Field (electricCurrentExponent)
	_electricCurrentExponent, _electricCurrentExponentErr := readBuffer.ReadInt8("electricCurrentExponent", 8)
	if _electricCurrentExponentErr != nil {
		return nil, errors.Wrap(_electricCurrentExponentErr, "Error parsing 'electricCurrentExponent' field of QuantityDimension")
	}
	electricCurrentExponent := _electricCurrentExponent

	// Simple Field (amountOfSubstanceExponent)
	_amountOfSubstanceExponent, _amountOfSubstanceExponentErr := readBuffer.ReadInt8("amountOfSubstanceExponent", 8)
	if _amountOfSubstanceExponentErr != nil {
		return nil, errors.Wrap(_amountOfSubstanceExponentErr, "Error parsing 'amountOfSubstanceExponent' field of QuantityDimension")
	}
	amountOfSubstanceExponent := _amountOfSubstanceExponent

	// Simple Field (luminousIntensityExponent)
	_luminousIntensityExponent, _luminousIntensityExponentErr := readBuffer.ReadInt8("luminousIntensityExponent", 8)
	if _luminousIntensityExponentErr != nil {
		return nil, errors.Wrap(_luminousIntensityExponentErr, "Error parsing 'luminousIntensityExponent' field of QuantityDimension")
	}
	luminousIntensityExponent := _luminousIntensityExponent

	// Simple Field (absoluteTemperatureExponent)
	_absoluteTemperatureExponent, _absoluteTemperatureExponentErr := readBuffer.ReadInt8("absoluteTemperatureExponent", 8)
	if _absoluteTemperatureExponentErr != nil {
		return nil, errors.Wrap(_absoluteTemperatureExponentErr, "Error parsing 'absoluteTemperatureExponent' field of QuantityDimension")
	}
	absoluteTemperatureExponent := _absoluteTemperatureExponent

	// Simple Field (dimensionlessExponent)
	_dimensionlessExponent, _dimensionlessExponentErr := readBuffer.ReadInt8("dimensionlessExponent", 8)
	if _dimensionlessExponentErr != nil {
		return nil, errors.Wrap(_dimensionlessExponentErr, "Error parsing 'dimensionlessExponent' field of QuantityDimension")
	}
	dimensionlessExponent := _dimensionlessExponent

	if closeErr := readBuffer.CloseContext("QuantityDimension"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for QuantityDimension")
	}

	// Create a partially initialized instance
	_child := &_QuantityDimension{
		_ExtensionObjectDefinition:  &_ExtensionObjectDefinition{},
		MassExponent:                massExponent,
		LengthExponent:              lengthExponent,
		TimeExponent:                timeExponent,
		ElectricCurrentExponent:     electricCurrentExponent,
		AmountOfSubstanceExponent:   amountOfSubstanceExponent,
		LuminousIntensityExponent:   luminousIntensityExponent,
		AbsoluteTemperatureExponent: absoluteTemperatureExponent,
		DimensionlessExponent:       dimensionlessExponent,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_QuantityDimension) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_QuantityDimension) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("QuantityDimension"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for QuantityDimension")
		}

		// Simple Field (massExponent)
		massExponent := int8(m.GetMassExponent())
		_massExponentErr := writeBuffer.WriteInt8("massExponent", 8, int8((massExponent)))
		if _massExponentErr != nil {
			return errors.Wrap(_massExponentErr, "Error serializing 'massExponent' field")
		}

		// Simple Field (lengthExponent)
		lengthExponent := int8(m.GetLengthExponent())
		_lengthExponentErr := writeBuffer.WriteInt8("lengthExponent", 8, int8((lengthExponent)))
		if _lengthExponentErr != nil {
			return errors.Wrap(_lengthExponentErr, "Error serializing 'lengthExponent' field")
		}

		// Simple Field (timeExponent)
		timeExponent := int8(m.GetTimeExponent())
		_timeExponentErr := writeBuffer.WriteInt8("timeExponent", 8, int8((timeExponent)))
		if _timeExponentErr != nil {
			return errors.Wrap(_timeExponentErr, "Error serializing 'timeExponent' field")
		}

		// Simple Field (electricCurrentExponent)
		electricCurrentExponent := int8(m.GetElectricCurrentExponent())
		_electricCurrentExponentErr := writeBuffer.WriteInt8("electricCurrentExponent", 8, int8((electricCurrentExponent)))
		if _electricCurrentExponentErr != nil {
			return errors.Wrap(_electricCurrentExponentErr, "Error serializing 'electricCurrentExponent' field")
		}

		// Simple Field (amountOfSubstanceExponent)
		amountOfSubstanceExponent := int8(m.GetAmountOfSubstanceExponent())
		_amountOfSubstanceExponentErr := writeBuffer.WriteInt8("amountOfSubstanceExponent", 8, int8((amountOfSubstanceExponent)))
		if _amountOfSubstanceExponentErr != nil {
			return errors.Wrap(_amountOfSubstanceExponentErr, "Error serializing 'amountOfSubstanceExponent' field")
		}

		// Simple Field (luminousIntensityExponent)
		luminousIntensityExponent := int8(m.GetLuminousIntensityExponent())
		_luminousIntensityExponentErr := writeBuffer.WriteInt8("luminousIntensityExponent", 8, int8((luminousIntensityExponent)))
		if _luminousIntensityExponentErr != nil {
			return errors.Wrap(_luminousIntensityExponentErr, "Error serializing 'luminousIntensityExponent' field")
		}

		// Simple Field (absoluteTemperatureExponent)
		absoluteTemperatureExponent := int8(m.GetAbsoluteTemperatureExponent())
		_absoluteTemperatureExponentErr := writeBuffer.WriteInt8("absoluteTemperatureExponent", 8, int8((absoluteTemperatureExponent)))
		if _absoluteTemperatureExponentErr != nil {
			return errors.Wrap(_absoluteTemperatureExponentErr, "Error serializing 'absoluteTemperatureExponent' field")
		}

		// Simple Field (dimensionlessExponent)
		dimensionlessExponent := int8(m.GetDimensionlessExponent())
		_dimensionlessExponentErr := writeBuffer.WriteInt8("dimensionlessExponent", 8, int8((dimensionlessExponent)))
		if _dimensionlessExponentErr != nil {
			return errors.Wrap(_dimensionlessExponentErr, "Error serializing 'dimensionlessExponent' field")
		}

		if popErr := writeBuffer.PopContext("QuantityDimension"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for QuantityDimension")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_QuantityDimension) isQuantityDimension() bool {
	return true
}

func (m *_QuantityDimension) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
