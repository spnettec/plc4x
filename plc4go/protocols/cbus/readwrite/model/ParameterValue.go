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


// ParameterValue is the corresponding interface of ParameterValue
type ParameterValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetParameterType returns ParameterType (discriminator field)
	GetParameterType() ParameterType
}

// ParameterValueExactly can be used when we want exactly this type and not a type which fulfills ParameterValue.
// This is useful for switch cases.
type ParameterValueExactly interface {
	ParameterValue
	isParameterValue() bool
}

// _ParameterValue is the data-structure of this message
type _ParameterValue struct {
	_ParameterValueChildRequirements

	// Arguments.
	NumBytes uint8
}

type _ParameterValueChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetParameterType() ParameterType
}


type ParameterValueParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ParameterValue, serializeChildFunction func() error) error
	GetTypeName() string
}

type ParameterValueChild interface {
	utils.Serializable
InitializeParent(parent ParameterValue )
	GetParent() *ParameterValue

	GetTypeName() string
	ParameterValue
}


// NewParameterValue factory function for _ParameterValue
func NewParameterValue( numBytes uint8 ) *_ParameterValue {
return &_ParameterValue{ NumBytes: numBytes }
}

// Deprecated: use the interface for direct cast
func CastParameterValue(structType any) ParameterValue {
    if casted, ok := structType.(ParameterValue); ok {
		return casted
	}
	if casted, ok := structType.(*ParameterValue); ok {
		return *casted
	}
	return nil
}

func (m *_ParameterValue) GetTypeName() string {
	return "ParameterValue"
}


func (m *_ParameterValue) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_ParameterValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ParameterValueParse(ctx context.Context, theBytes []byte, parameterType ParameterType, numBytes uint8) (ParameterValue, error) {
	return ParameterValueParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), parameterType, numBytes)
}

func ParameterValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, parameterType ParameterType, numBytes uint8) (ParameterValue, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ParameterValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ParameterValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ParameterValueChildSerializeRequirement interface {
		ParameterValue
		InitializeParent(ParameterValue )
		GetParent() ParameterValue
	}
	var _childTemp any
	var _child ParameterValueChildSerializeRequirement
	var typeSwitchError error
	switch {
case parameterType == ParameterType_APPLICATION_ADDRESS_1 : // ParameterValueApplicationAddress1
		_childTemp, typeSwitchError = ParameterValueApplicationAddress1ParseWithBuffer(ctx, readBuffer, parameterType, numBytes)
case parameterType == ParameterType_APPLICATION_ADDRESS_2 : // ParameterValueApplicationAddress2
		_childTemp, typeSwitchError = ParameterValueApplicationAddress2ParseWithBuffer(ctx, readBuffer, parameterType, numBytes)
case parameterType == ParameterType_INTERFACE_OPTIONS_1 : // ParameterValueInterfaceOptions1
		_childTemp, typeSwitchError = ParameterValueInterfaceOptions1ParseWithBuffer(ctx, readBuffer, parameterType, numBytes)
case parameterType == ParameterType_BAUD_RATE_SELECTOR : // ParameterValueBaudRateSelector
		_childTemp, typeSwitchError = ParameterValueBaudRateSelectorParseWithBuffer(ctx, readBuffer, parameterType, numBytes)
case parameterType == ParameterType_INTERFACE_OPTIONS_2 : // ParameterValueInterfaceOptions2
		_childTemp, typeSwitchError = ParameterValueInterfaceOptions2ParseWithBuffer(ctx, readBuffer, parameterType, numBytes)
case parameterType == ParameterType_INTERFACE_OPTIONS_1_POWER_UP_SETTINGS : // ParameterValueInterfaceOptions1PowerUpSettings
		_childTemp, typeSwitchError = ParameterValueInterfaceOptions1PowerUpSettingsParseWithBuffer(ctx, readBuffer, parameterType, numBytes)
case parameterType == ParameterType_INTERFACE_OPTIONS_3 : // ParameterValueInterfaceOptions3
		_childTemp, typeSwitchError = ParameterValueInterfaceOptions3ParseWithBuffer(ctx, readBuffer, parameterType, numBytes)
case parameterType == ParameterType_CUSTOM_MANUFACTURER : // ParameterValueCustomManufacturer
		_childTemp, typeSwitchError = ParameterValueCustomManufacturerParseWithBuffer(ctx, readBuffer, parameterType, numBytes)
case parameterType == ParameterType_SERIAL_NUMBER : // ParameterValueSerialNumber
		_childTemp, typeSwitchError = ParameterValueSerialNumberParseWithBuffer(ctx, readBuffer, parameterType, numBytes)
case parameterType == ParameterType_CUSTOM_TYPE : // ParameterValueCustomTypes
		_childTemp, typeSwitchError = ParameterValueCustomTypesParseWithBuffer(ctx, readBuffer, parameterType, numBytes)
case 0==0 : // ParameterValueRaw
		_childTemp, typeSwitchError = ParameterValueRawParseWithBuffer(ctx, readBuffer, parameterType, numBytes)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [parameterType=%v]", parameterType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of ParameterValue")
	}
	_child = _childTemp.(ParameterValueChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("ParameterValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ParameterValue")
	}

	// Finish initializing
_child.InitializeParent(_child )
	return _child, nil
}

func (pm *_ParameterValue) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child ParameterValue, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("ParameterValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ParameterValue")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ParameterValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ParameterValue")
	}
	return nil
}


////
// Arguments Getter

func (m *_ParameterValue) GetNumBytes() uint8 {
	return m.NumBytes
}
//
////

func (m *_ParameterValue) isParameterValue() bool {
	return true
}

func (m *_ParameterValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



