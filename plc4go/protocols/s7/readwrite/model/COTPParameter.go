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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// COTPParameter is the corresponding interface of COTPParameter
type COTPParameter interface {
	utils.LengthAware
	utils.Serializable
	// GetParameterType returns ParameterType (discriminator field)
	GetParameterType() uint8
}

// COTPParameterExactly can be used when we want exactly this type and not a type which fulfills COTPParameter.
// This is useful for switch cases.
type COTPParameterExactly interface {
	COTPParameter
	isCOTPParameter() bool
}

// _COTPParameter is the data-structure of this message
type _COTPParameter struct {
	_COTPParameterChildRequirements

	// Arguments.
	Rest uint8
}

type _COTPParameterChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetParameterType() uint8
}


type COTPParameterParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child COTPParameter, serializeChildFunction func() error) error
	GetTypeName() string
}

type COTPParameterChild interface {
	utils.Serializable
InitializeParent(parent COTPParameter )
	GetParent() *COTPParameter

	GetTypeName() string
	COTPParameter
}


// NewCOTPParameter factory function for _COTPParameter
func NewCOTPParameter( rest uint8 ) *_COTPParameter {
return &_COTPParameter{ Rest: rest }
}

// Deprecated: use the interface for direct cast
func CastCOTPParameter(structType interface{}) COTPParameter {
    if casted, ok := structType.(COTPParameter); ok {
		return casted
	}
	if casted, ok := structType.(*COTPParameter); ok {
		return *casted
	}
	return nil
}

func (m *_COTPParameter) GetTypeName() string {
	return "COTPParameter"
}


func (m *_COTPParameter) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (parameterType)
	lengthInBits += 8;

	// Implicit Field (parameterLength)
	lengthInBits += 8

	return lengthInBits
}

func (m *_COTPParameter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func COTPParameterParse(theBytes []byte, rest uint8) (COTPParameter, error) {
	return COTPParameterParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), rest)
}

func COTPParameterParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, rest uint8) (COTPParameter, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPParameter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPParameter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (parameterType) (Used as input to a switch field)
	parameterType, _parameterTypeErr := readBuffer.ReadUint8("parameterType", 8)
	if _parameterTypeErr != nil {
		return nil, errors.Wrap(_parameterTypeErr, "Error parsing 'parameterType' field of COTPParameter")
	}

	// Implicit Field (parameterLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	parameterLength, _parameterLengthErr := readBuffer.ReadUint8("parameterLength", 8)
	_ = parameterLength
	if _parameterLengthErr != nil {
		return nil, errors.Wrap(_parameterLengthErr, "Error parsing 'parameterLength' field of COTPParameter")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type COTPParameterChildSerializeRequirement interface {
		COTPParameter
		InitializeParent(COTPParameter )
		GetParent() COTPParameter
	}
	var _childTemp interface{}
	var _child COTPParameterChildSerializeRequirement
	var typeSwitchError error
	switch {
case parameterType == 0xC0 : // COTPParameterTpduSize
		_childTemp, typeSwitchError = COTPParameterTpduSizeParseWithBuffer(ctx, readBuffer, rest)
case parameterType == 0xC1 : // COTPParameterCallingTsap
		_childTemp, typeSwitchError = COTPParameterCallingTsapParseWithBuffer(ctx, readBuffer, rest)
case parameterType == 0xC2 : // COTPParameterCalledTsap
		_childTemp, typeSwitchError = COTPParameterCalledTsapParseWithBuffer(ctx, readBuffer, rest)
case parameterType == 0xC3 : // COTPParameterChecksum
		_childTemp, typeSwitchError = COTPParameterChecksumParseWithBuffer(ctx, readBuffer, rest)
case parameterType == 0xE0 : // COTPParameterDisconnectAdditionalInformation
		_childTemp, typeSwitchError = COTPParameterDisconnectAdditionalInformationParseWithBuffer(ctx, readBuffer, rest)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [parameterType=%v]", parameterType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of COTPParameter")
	}
	_child = _childTemp.(COTPParameterChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("COTPParameter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPParameter")
	}

	// Finish initializing
_child.InitializeParent(_child )
	return _child, nil
}

func (pm *_COTPParameter) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child COTPParameter, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("COTPParameter"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for COTPParameter")
	}

	// Discriminator Field (parameterType) (Used as input to a switch field)
	parameterType := uint8(child.GetParameterType())
	_parameterTypeErr := writeBuffer.WriteUint8("parameterType", 8, (parameterType))

	if _parameterTypeErr != nil {
		return errors.Wrap(_parameterTypeErr, "Error serializing 'parameterType' field")
	}

	// Implicit Field (parameterLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	parameterLength := uint8(uint8(uint8(m.GetLengthInBytes(ctx))) - uint8(uint8(2)))
	_parameterLengthErr := writeBuffer.WriteUint8("parameterLength", 8, (parameterLength))
	if _parameterLengthErr != nil {
		return errors.Wrap(_parameterLengthErr, "Error serializing 'parameterLength' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("COTPParameter"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for COTPParameter")
	}
	return nil
}


////
// Arguments Getter

func (m *_COTPParameter) GetRest() uint8 {
	return m.Rest
}
//
////

func (m *_COTPParameter) isCOTPParameter() bool {
	return true
}

func (m *_COTPParameter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



