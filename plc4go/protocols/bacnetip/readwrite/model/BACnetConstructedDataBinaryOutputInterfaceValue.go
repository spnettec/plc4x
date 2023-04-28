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

// BACnetConstructedDataBinaryOutputInterfaceValue is the corresponding interface of BACnetConstructedDataBinaryOutputInterfaceValue
type BACnetConstructedDataBinaryOutputInterfaceValue interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetInterfaceValue returns InterfaceValue (property field)
	GetInterfaceValue() BACnetOptionalBinaryPV
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetOptionalBinaryPV
}

// BACnetConstructedDataBinaryOutputInterfaceValueExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBinaryOutputInterfaceValue.
// This is useful for switch cases.
type BACnetConstructedDataBinaryOutputInterfaceValueExactly interface {
	BACnetConstructedDataBinaryOutputInterfaceValue
	isBACnetConstructedDataBinaryOutputInterfaceValue() bool
}

// _BACnetConstructedDataBinaryOutputInterfaceValue is the data-structure of this message
type _BACnetConstructedDataBinaryOutputInterfaceValue struct {
	*_BACnetConstructedData
	InterfaceValue BACnetOptionalBinaryPV
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_BINARY_OUTPUT
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INTERFACE_VALUE
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetInterfaceValue() BACnetOptionalBinaryPV {
	return m.InterfaceValue
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetActualValue() BACnetOptionalBinaryPV {
	ctx := context.Background()
	_ = ctx
	return CastBACnetOptionalBinaryPV(m.GetInterfaceValue())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBinaryOutputInterfaceValue factory function for _BACnetConstructedDataBinaryOutputInterfaceValue
func NewBACnetConstructedDataBinaryOutputInterfaceValue(interfaceValue BACnetOptionalBinaryPV, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBinaryOutputInterfaceValue {
	_result := &_BACnetConstructedDataBinaryOutputInterfaceValue{
		InterfaceValue:         interfaceValue,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBinaryOutputInterfaceValue(structType any) BACnetConstructedDataBinaryOutputInterfaceValue {
	if casted, ok := structType.(BACnetConstructedDataBinaryOutputInterfaceValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBinaryOutputInterfaceValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetTypeName() string {
	return "BACnetConstructedDataBinaryOutputInterfaceValue"
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (interfaceValue)
	lengthInBits += m.InterfaceValue.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBinaryOutputInterfaceValueParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBinaryOutputInterfaceValue, error) {
	return BACnetConstructedDataBinaryOutputInterfaceValueParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBinaryOutputInterfaceValueParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBinaryOutputInterfaceValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBinaryOutputInterfaceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBinaryOutputInterfaceValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (interfaceValue)
	if pullErr := readBuffer.PullContext("interfaceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for interfaceValue")
	}
	_interfaceValue, _interfaceValueErr := BACnetOptionalBinaryPVParseWithBuffer(ctx, readBuffer)
	if _interfaceValueErr != nil {
		return nil, errors.Wrap(_interfaceValueErr, "Error parsing 'interfaceValue' field of BACnetConstructedDataBinaryOutputInterfaceValue")
	}
	interfaceValue := _interfaceValue.(BACnetOptionalBinaryPV)
	if closeErr := readBuffer.CloseContext("interfaceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for interfaceValue")
	}

	// Virtual field
	_actualValue := interfaceValue
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBinaryOutputInterfaceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBinaryOutputInterfaceValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBinaryOutputInterfaceValue{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		InterfaceValue: interfaceValue,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBinaryOutputInterfaceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBinaryOutputInterfaceValue")
		}

		// Simple Field (interfaceValue)
		if pushErr := writeBuffer.PushContext("interfaceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for interfaceValue")
		}
		_interfaceValueErr := writeBuffer.WriteSerializable(ctx, m.GetInterfaceValue())
		if popErr := writeBuffer.PopContext("interfaceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for interfaceValue")
		}
		if _interfaceValueErr != nil {
			return errors.Wrap(_interfaceValueErr, "Error serializing 'interfaceValue' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBinaryOutputInterfaceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBinaryOutputInterfaceValue")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) isBACnetConstructedDataBinaryOutputInterfaceValue() bool {
	return true
}

func (m *_BACnetConstructedDataBinaryOutputInterfaceValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
