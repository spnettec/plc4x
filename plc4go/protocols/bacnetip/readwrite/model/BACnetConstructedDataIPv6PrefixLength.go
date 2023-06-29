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

// BACnetConstructedDataIPv6PrefixLength is the corresponding interface of BACnetConstructedDataIPv6PrefixLength
type BACnetConstructedDataIPv6PrefixLength interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpv6PrefixLength returns Ipv6PrefixLength (property field)
	GetIpv6PrefixLength() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataIPv6PrefixLengthExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIPv6PrefixLength.
// This is useful for switch cases.
type BACnetConstructedDataIPv6PrefixLengthExactly interface {
	BACnetConstructedDataIPv6PrefixLength
	isBACnetConstructedDataIPv6PrefixLength() bool
}

// _BACnetConstructedDataIPv6PrefixLength is the data-structure of this message
type _BACnetConstructedDataIPv6PrefixLength struct {
	*_BACnetConstructedData
	Ipv6PrefixLength BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPv6PrefixLength) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPv6PrefixLength) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_PREFIX_LENGTH
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPv6PrefixLength) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIPv6PrefixLength) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6PrefixLength) GetIpv6PrefixLength() BACnetApplicationTagUnsignedInteger {
	return m.Ipv6PrefixLength
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6PrefixLength) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetIpv6PrefixLength())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPv6PrefixLength factory function for _BACnetConstructedDataIPv6PrefixLength
func NewBACnetConstructedDataIPv6PrefixLength(ipv6PrefixLength BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPv6PrefixLength {
	_result := &_BACnetConstructedDataIPv6PrefixLength{
		Ipv6PrefixLength:       ipv6PrefixLength,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPv6PrefixLength(structType any) BACnetConstructedDataIPv6PrefixLength {
	if casted, ok := structType.(BACnetConstructedDataIPv6PrefixLength); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6PrefixLength); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPv6PrefixLength) GetTypeName() string {
	return "BACnetConstructedDataIPv6PrefixLength"
}

func (m *_BACnetConstructedDataIPv6PrefixLength) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (ipv6PrefixLength)
	lengthInBits += m.Ipv6PrefixLength.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPv6PrefixLength) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataIPv6PrefixLengthParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPv6PrefixLength, error) {
	return BACnetConstructedDataIPv6PrefixLengthParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataIPv6PrefixLengthParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPv6PrefixLength, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6PrefixLength"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6PrefixLength")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipv6PrefixLength)
	if pullErr := readBuffer.PullContext("ipv6PrefixLength"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipv6PrefixLength")
	}
	_ipv6PrefixLength, _ipv6PrefixLengthErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _ipv6PrefixLengthErr != nil {
		return nil, errors.Wrap(_ipv6PrefixLengthErr, "Error parsing 'ipv6PrefixLength' field of BACnetConstructedDataIPv6PrefixLength")
	}
	ipv6PrefixLength := _ipv6PrefixLength.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("ipv6PrefixLength"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipv6PrefixLength")
	}

	// Virtual field
	_actualValue := ipv6PrefixLength
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6PrefixLength"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6PrefixLength")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIPv6PrefixLength{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Ipv6PrefixLength: ipv6PrefixLength,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIPv6PrefixLength) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPv6PrefixLength) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6PrefixLength"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6PrefixLength")
		}

		// Simple Field (ipv6PrefixLength)
		if pushErr := writeBuffer.PushContext("ipv6PrefixLength"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipv6PrefixLength")
		}
		_ipv6PrefixLengthErr := writeBuffer.WriteSerializable(ctx, m.GetIpv6PrefixLength())
		if popErr := writeBuffer.PopContext("ipv6PrefixLength"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipv6PrefixLength")
		}
		if _ipv6PrefixLengthErr != nil {
			return errors.Wrap(_ipv6PrefixLengthErr, "Error serializing 'ipv6PrefixLength' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6PrefixLength"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6PrefixLength")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPv6PrefixLength) isBACnetConstructedDataIPv6PrefixLength() bool {
	return true
}

func (m *_BACnetConstructedDataIPv6PrefixLength) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
