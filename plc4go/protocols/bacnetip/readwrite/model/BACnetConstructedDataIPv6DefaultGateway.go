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

// BACnetConstructedDataIPv6DefaultGateway is the corresponding interface of BACnetConstructedDataIPv6DefaultGateway
type BACnetConstructedDataIPv6DefaultGateway interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpv6DefaultGateway returns Ipv6DefaultGateway (property field)
	GetIpv6DefaultGateway() BACnetApplicationTagOctetString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagOctetString
}

// BACnetConstructedDataIPv6DefaultGatewayExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIPv6DefaultGateway.
// This is useful for switch cases.
type BACnetConstructedDataIPv6DefaultGatewayExactly interface {
	BACnetConstructedDataIPv6DefaultGateway
	isBACnetConstructedDataIPv6DefaultGateway() bool
}

// _BACnetConstructedDataIPv6DefaultGateway is the data-structure of this message
type _BACnetConstructedDataIPv6DefaultGateway struct {
	*_BACnetConstructedData
	Ipv6DefaultGateway BACnetApplicationTagOctetString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPv6DefaultGateway) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IPV6_DEFAULT_GATEWAY
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPv6DefaultGateway) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DefaultGateway) GetIpv6DefaultGateway() BACnetApplicationTagOctetString {
	return m.Ipv6DefaultGateway
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPv6DefaultGateway) GetActualValue() BACnetApplicationTagOctetString {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagOctetString(m.GetIpv6DefaultGateway())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPv6DefaultGateway factory function for _BACnetConstructedDataIPv6DefaultGateway
func NewBACnetConstructedDataIPv6DefaultGateway(ipv6DefaultGateway BACnetApplicationTagOctetString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPv6DefaultGateway {
	_result := &_BACnetConstructedDataIPv6DefaultGateway{
		Ipv6DefaultGateway:     ipv6DefaultGateway,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPv6DefaultGateway(structType any) BACnetConstructedDataIPv6DefaultGateway {
	if casted, ok := structType.(BACnetConstructedDataIPv6DefaultGateway); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPv6DefaultGateway); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) GetTypeName() string {
	return "BACnetConstructedDataIPv6DefaultGateway"
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (ipv6DefaultGateway)
	lengthInBits += m.Ipv6DefaultGateway.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataIPv6DefaultGatewayParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPv6DefaultGateway, error) {
	return BACnetConstructedDataIPv6DefaultGatewayParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataIPv6DefaultGatewayParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPv6DefaultGateway, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPv6DefaultGateway"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPv6DefaultGateway")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipv6DefaultGateway)
	if pullErr := readBuffer.PullContext("ipv6DefaultGateway"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipv6DefaultGateway")
	}
	_ipv6DefaultGateway, _ipv6DefaultGatewayErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _ipv6DefaultGatewayErr != nil {
		return nil, errors.Wrap(_ipv6DefaultGatewayErr, "Error parsing 'ipv6DefaultGateway' field of BACnetConstructedDataIPv6DefaultGateway")
	}
	ipv6DefaultGateway := _ipv6DefaultGateway.(BACnetApplicationTagOctetString)
	if closeErr := readBuffer.CloseContext("ipv6DefaultGateway"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipv6DefaultGateway")
	}

	// Virtual field
	_actualValue := ipv6DefaultGateway
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPv6DefaultGateway"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPv6DefaultGateway")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIPv6DefaultGateway{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Ipv6DefaultGateway: ipv6DefaultGateway,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPv6DefaultGateway"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPv6DefaultGateway")
		}

		// Simple Field (ipv6DefaultGateway)
		if pushErr := writeBuffer.PushContext("ipv6DefaultGateway"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipv6DefaultGateway")
		}
		_ipv6DefaultGatewayErr := writeBuffer.WriteSerializable(ctx, m.GetIpv6DefaultGateway())
		if popErr := writeBuffer.PopContext("ipv6DefaultGateway"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipv6DefaultGateway")
		}
		if _ipv6DefaultGatewayErr != nil {
			return errors.Wrap(_ipv6DefaultGatewayErr, "Error serializing 'ipv6DefaultGateway' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPv6DefaultGateway"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPv6DefaultGateway")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) isBACnetConstructedDataIPv6DefaultGateway() bool {
	return true
}

func (m *_BACnetConstructedDataIPv6DefaultGateway) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
