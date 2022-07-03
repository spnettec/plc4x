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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataBinaryInputInterfaceValue is the corresponding interface of BACnetConstructedDataBinaryInputInterfaceValue
type BACnetConstructedDataBinaryInputInterfaceValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetInterfaceValue returns InterfaceValue (property field)
	GetInterfaceValue() BACnetOptionalBinaryPV
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetOptionalBinaryPV
}

// BACnetConstructedDataBinaryInputInterfaceValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBinaryInputInterfaceValue.
// This is useful for switch cases.
type BACnetConstructedDataBinaryInputInterfaceValueExactly interface {
	BACnetConstructedDataBinaryInputInterfaceValue
	isBACnetConstructedDataBinaryInputInterfaceValue() bool
}

// _BACnetConstructedDataBinaryInputInterfaceValue is the data-structure of this message
type _BACnetConstructedDataBinaryInputInterfaceValue struct {
	*_BACnetConstructedData
	InterfaceValue BACnetOptionalBinaryPV
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBinaryInputInterfaceValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BINARY_INPUT
}

func (m *_BACnetConstructedDataBinaryInputInterfaceValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INTERFACE_VALUE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBinaryInputInterfaceValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBinaryInputInterfaceValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryInputInterfaceValue) GetInterfaceValue() BACnetOptionalBinaryPV {
	return m.InterfaceValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryInputInterfaceValue) GetActualValue() BACnetOptionalBinaryPV {
	return CastBACnetOptionalBinaryPV(m.GetInterfaceValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBinaryInputInterfaceValue factory function for _BACnetConstructedDataBinaryInputInterfaceValue
func NewBACnetConstructedDataBinaryInputInterfaceValue(interfaceValue BACnetOptionalBinaryPV, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBinaryInputInterfaceValue {
	_result := &_BACnetConstructedDataBinaryInputInterfaceValue{
		InterfaceValue:         interfaceValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBinaryInputInterfaceValue(structType interface{}) BACnetConstructedDataBinaryInputInterfaceValue {
	if casted, ok := structType.(BACnetConstructedDataBinaryInputInterfaceValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBinaryInputInterfaceValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBinaryInputInterfaceValue) GetTypeName() string {
	return "BACnetConstructedDataBinaryInputInterfaceValue"
}

func (m *_BACnetConstructedDataBinaryInputInterfaceValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataBinaryInputInterfaceValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (interfaceValue)
	lengthInBits += m.InterfaceValue.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBinaryInputInterfaceValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataBinaryInputInterfaceValueParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBinaryInputInterfaceValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBinaryInputInterfaceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBinaryInputInterfaceValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (interfaceValue)
	if pullErr := readBuffer.PullContext("interfaceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for interfaceValue")
	}
	_interfaceValue, _interfaceValueErr := BACnetOptionalBinaryPVParse(readBuffer)
	if _interfaceValueErr != nil {
		return nil, errors.Wrap(_interfaceValueErr, "Error parsing 'interfaceValue' field of BACnetConstructedDataBinaryInputInterfaceValue")
	}
	interfaceValue := _interfaceValue.(BACnetOptionalBinaryPV)
	if closeErr := readBuffer.CloseContext("interfaceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for interfaceValue")
	}

	// Virtual field
	_actualValue := interfaceValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBinaryInputInterfaceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBinaryInputInterfaceValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBinaryInputInterfaceValue{
		InterfaceValue: interfaceValue,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBinaryInputInterfaceValue) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBinaryInputInterfaceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBinaryInputInterfaceValue")
		}

		// Simple Field (interfaceValue)
		if pushErr := writeBuffer.PushContext("interfaceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for interfaceValue")
		}
		_interfaceValueErr := writeBuffer.WriteSerializable(m.GetInterfaceValue())
		if popErr := writeBuffer.PopContext("interfaceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for interfaceValue")
		}
		if _interfaceValueErr != nil {
			return errors.Wrap(_interfaceValueErr, "Error serializing 'interfaceValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBinaryInputInterfaceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBinaryInputInterfaceValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBinaryInputInterfaceValue) isBACnetConstructedDataBinaryInputInterfaceValue() bool {
	return true
}

func (m *_BACnetConstructedDataBinaryInputInterfaceValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
