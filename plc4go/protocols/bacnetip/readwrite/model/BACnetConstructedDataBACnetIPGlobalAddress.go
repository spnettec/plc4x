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

// BACnetConstructedDataBACnetIPGlobalAddress is the corresponding interface of BACnetConstructedDataBACnetIPGlobalAddress
type BACnetConstructedDataBACnetIPGlobalAddress interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBacnetIpGlobalAddress returns BacnetIpGlobalAddress (property field)
	GetBacnetIpGlobalAddress() BACnetHostNPort
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetHostNPort
}

// BACnetConstructedDataBACnetIPGlobalAddressExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBACnetIPGlobalAddress.
// This is useful for switch cases.
type BACnetConstructedDataBACnetIPGlobalAddressExactly interface {
	BACnetConstructedDataBACnetIPGlobalAddress
	isBACnetConstructedDataBACnetIPGlobalAddress() bool
}

// _BACnetConstructedDataBACnetIPGlobalAddress is the data-structure of this message
type _BACnetConstructedDataBACnetIPGlobalAddress struct {
	*_BACnetConstructedData
	BacnetIpGlobalAddress BACnetHostNPort
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPGlobalAddress) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBACnetIPGlobalAddress) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACNET_IP_GLOBAL_ADDRESS
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBACnetIPGlobalAddress) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBACnetIPGlobalAddress) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPGlobalAddress) GetBacnetIpGlobalAddress() BACnetHostNPort {
	return m.BacnetIpGlobalAddress
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBACnetIPGlobalAddress) GetActualValue() BACnetHostNPort {
	ctx := context.Background()
	_ = ctx
	return CastBACnetHostNPort(m.GetBacnetIpGlobalAddress())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBACnetIPGlobalAddress factory function for _BACnetConstructedDataBACnetIPGlobalAddress
func NewBACnetConstructedDataBACnetIPGlobalAddress(bacnetIpGlobalAddress BACnetHostNPort, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBACnetIPGlobalAddress {
	_result := &_BACnetConstructedDataBACnetIPGlobalAddress{
		BacnetIpGlobalAddress:  bacnetIpGlobalAddress,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBACnetIPGlobalAddress(structType any) BACnetConstructedDataBACnetIPGlobalAddress {
	if casted, ok := structType.(BACnetConstructedDataBACnetIPGlobalAddress); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBACnetIPGlobalAddress); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBACnetIPGlobalAddress) GetTypeName() string {
	return "BACnetConstructedDataBACnetIPGlobalAddress"
}

func (m *_BACnetConstructedDataBACnetIPGlobalAddress) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (bacnetIpGlobalAddress)
	lengthInBits += m.BacnetIpGlobalAddress.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBACnetIPGlobalAddress) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBACnetIPGlobalAddressParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBACnetIPGlobalAddress, error) {
	return BACnetConstructedDataBACnetIPGlobalAddressParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBACnetIPGlobalAddressParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBACnetIPGlobalAddress, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBACnetIPGlobalAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBACnetIPGlobalAddress")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bacnetIpGlobalAddress)
	if pullErr := readBuffer.PullContext("bacnetIpGlobalAddress"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bacnetIpGlobalAddress")
	}
	_bacnetIpGlobalAddress, _bacnetIpGlobalAddressErr := BACnetHostNPortParseWithBuffer(ctx, readBuffer)
	if _bacnetIpGlobalAddressErr != nil {
		return nil, errors.Wrap(_bacnetIpGlobalAddressErr, "Error parsing 'bacnetIpGlobalAddress' field of BACnetConstructedDataBACnetIPGlobalAddress")
	}
	bacnetIpGlobalAddress := _bacnetIpGlobalAddress.(BACnetHostNPort)
	if closeErr := readBuffer.CloseContext("bacnetIpGlobalAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bacnetIpGlobalAddress")
	}

	// Virtual field
	_actualValue := bacnetIpGlobalAddress
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBACnetIPGlobalAddress"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBACnetIPGlobalAddress")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBACnetIPGlobalAddress{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		BacnetIpGlobalAddress: bacnetIpGlobalAddress,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBACnetIPGlobalAddress) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBACnetIPGlobalAddress) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBACnetIPGlobalAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBACnetIPGlobalAddress")
		}

		// Simple Field (bacnetIpGlobalAddress)
		if pushErr := writeBuffer.PushContext("bacnetIpGlobalAddress"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bacnetIpGlobalAddress")
		}
		_bacnetIpGlobalAddressErr := writeBuffer.WriteSerializable(ctx, m.GetBacnetIpGlobalAddress())
		if popErr := writeBuffer.PopContext("bacnetIpGlobalAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bacnetIpGlobalAddress")
		}
		if _bacnetIpGlobalAddressErr != nil {
			return errors.Wrap(_bacnetIpGlobalAddressErr, "Error serializing 'bacnetIpGlobalAddress' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBACnetIPGlobalAddress"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBACnetIPGlobalAddress")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBACnetIPGlobalAddress) isBACnetConstructedDataBACnetIPGlobalAddress() bool {
	return true
}

func (m *_BACnetConstructedDataBACnetIPGlobalAddress) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
