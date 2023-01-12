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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// CEMIAdditionalInformation is the corresponding interface of CEMIAdditionalInformation
type CEMIAdditionalInformation interface {
	utils.LengthAware
	utils.Serializable
	// GetAdditionalInformationType returns AdditionalInformationType (discriminator field)
	GetAdditionalInformationType() uint8
}

// CEMIAdditionalInformationExactly can be used when we want exactly this type and not a type which fulfills CEMIAdditionalInformation.
// This is useful for switch cases.
type CEMIAdditionalInformationExactly interface {
	CEMIAdditionalInformation
	isCEMIAdditionalInformation() bool
}

// _CEMIAdditionalInformation is the data-structure of this message
type _CEMIAdditionalInformation struct {
	_CEMIAdditionalInformationChildRequirements
}

type _CEMIAdditionalInformationChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetAdditionalInformationType() uint8
}


type CEMIAdditionalInformationParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child CEMIAdditionalInformation, serializeChildFunction func() error) error
	GetTypeName() string
}

type CEMIAdditionalInformationChild interface {
	utils.Serializable
InitializeParent(parent CEMIAdditionalInformation )
	GetParent() *CEMIAdditionalInformation

	GetTypeName() string
	CEMIAdditionalInformation
}


// NewCEMIAdditionalInformation factory function for _CEMIAdditionalInformation
func NewCEMIAdditionalInformation( ) *_CEMIAdditionalInformation {
return &_CEMIAdditionalInformation{ }
}

// Deprecated: use the interface for direct cast
func CastCEMIAdditionalInformation(structType interface{}) CEMIAdditionalInformation {
    if casted, ok := structType.(CEMIAdditionalInformation); ok {
		return casted
	}
	if casted, ok := structType.(*CEMIAdditionalInformation); ok {
		return *casted
	}
	return nil
}

func (m *_CEMIAdditionalInformation) GetTypeName() string {
	return "CEMIAdditionalInformation"
}



func (m *_CEMIAdditionalInformation) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (additionalInformationType)
	lengthInBits += 8;

	return lengthInBits
}

func (m *_CEMIAdditionalInformation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CEMIAdditionalInformationParse(theBytes []byte) (CEMIAdditionalInformation, error) {
	return CEMIAdditionalInformationParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func CEMIAdditionalInformationParseWithBuffer(readBuffer utils.ReadBuffer) (CEMIAdditionalInformation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CEMIAdditionalInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CEMIAdditionalInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (additionalInformationType) (Used as input to a switch field)
	additionalInformationType, _additionalInformationTypeErr := readBuffer.ReadUint8("additionalInformationType", 8)
	if _additionalInformationTypeErr != nil {
		return nil, errors.Wrap(_additionalInformationTypeErr, "Error parsing 'additionalInformationType' field of CEMIAdditionalInformation")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type CEMIAdditionalInformationChildSerializeRequirement interface {
		CEMIAdditionalInformation
		InitializeParent(CEMIAdditionalInformation )
		GetParent() CEMIAdditionalInformation
	}
	var _childTemp interface{}
	var _child CEMIAdditionalInformationChildSerializeRequirement
	var typeSwitchError error
	switch {
case additionalInformationType == 0x03 : // CEMIAdditionalInformationBusmonitorInfo
		_childTemp, typeSwitchError = CEMIAdditionalInformationBusmonitorInfoParseWithBuffer(readBuffer, )
case additionalInformationType == 0x04 : // CEMIAdditionalInformationRelativeTimestamp
		_childTemp, typeSwitchError = CEMIAdditionalInformationRelativeTimestampParseWithBuffer(readBuffer, )
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [additionalInformationType=%v]", additionalInformationType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of CEMIAdditionalInformation")
	}
	_child = _childTemp.(CEMIAdditionalInformationChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("CEMIAdditionalInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CEMIAdditionalInformation")
	}

	// Finish initializing
_child.InitializeParent(_child )
	return _child, nil
}

func (pm *_CEMIAdditionalInformation) SerializeParent(writeBuffer utils.WriteBuffer, child CEMIAdditionalInformation, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("CEMIAdditionalInformation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for CEMIAdditionalInformation")
	}

	// Discriminator Field (additionalInformationType) (Used as input to a switch field)
	additionalInformationType := uint8(child.GetAdditionalInformationType())
	_additionalInformationTypeErr := writeBuffer.WriteUint8("additionalInformationType", 8, (additionalInformationType))

	if _additionalInformationTypeErr != nil {
		return errors.Wrap(_additionalInformationTypeErr, "Error serializing 'additionalInformationType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CEMIAdditionalInformation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for CEMIAdditionalInformation")
	}
	return nil
}


func (m *_CEMIAdditionalInformation) isCEMIAdditionalInformation() bool {
	return true
}

func (m *_CEMIAdditionalInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



