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

// BACnetConstructedDataIPDHCPLeaseTime is the corresponding interface of BACnetConstructedDataIPDHCPLeaseTime
type BACnetConstructedDataIPDHCPLeaseTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpDhcpLeaseTime returns IpDhcpLeaseTime (property field)
	GetIpDhcpLeaseTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataIPDHCPLeaseTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIPDHCPLeaseTime.
// This is useful for switch cases.
type BACnetConstructedDataIPDHCPLeaseTimeExactly interface {
	BACnetConstructedDataIPDHCPLeaseTime
	isBACnetConstructedDataIPDHCPLeaseTime() bool
}

// _BACnetConstructedDataIPDHCPLeaseTime is the data-structure of this message
type _BACnetConstructedDataIPDHCPLeaseTime struct {
	*_BACnetConstructedData
	IpDhcpLeaseTime BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPLeaseTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IP_DHCP_LEASE_TIME
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPDHCPLeaseTime) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPLeaseTime) GetIpDhcpLeaseTime() BACnetApplicationTagUnsignedInteger {
	return m.IpDhcpLeaseTime
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPLeaseTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetIpDhcpLeaseTime())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPDHCPLeaseTime factory function for _BACnetConstructedDataIPDHCPLeaseTime
func NewBACnetConstructedDataIPDHCPLeaseTime(ipDhcpLeaseTime BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPDHCPLeaseTime {
	_result := &_BACnetConstructedDataIPDHCPLeaseTime{
		IpDhcpLeaseTime:        ipDhcpLeaseTime,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPDHCPLeaseTime(structType any) BACnetConstructedDataIPDHCPLeaseTime {
	if casted, ok := structType.(BACnetConstructedDataIPDHCPLeaseTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPDHCPLeaseTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) GetTypeName() string {
	return "BACnetConstructedDataIPDHCPLeaseTime"
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (ipDhcpLeaseTime)
	lengthInBits += m.IpDhcpLeaseTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataIPDHCPLeaseTimeParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPDHCPLeaseTime, error) {
	return BACnetConstructedDataIPDHCPLeaseTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataIPDHCPLeaseTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPDHCPLeaseTime, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPDHCPLeaseTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPDHCPLeaseTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipDhcpLeaseTime)
	if pullErr := readBuffer.PullContext("ipDhcpLeaseTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipDhcpLeaseTime")
	}
	_ipDhcpLeaseTime, _ipDhcpLeaseTimeErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _ipDhcpLeaseTimeErr != nil {
		return nil, errors.Wrap(_ipDhcpLeaseTimeErr, "Error parsing 'ipDhcpLeaseTime' field of BACnetConstructedDataIPDHCPLeaseTime")
	}
	ipDhcpLeaseTime := _ipDhcpLeaseTime.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("ipDhcpLeaseTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipDhcpLeaseTime")
	}

	// Virtual field
	_actualValue := ipDhcpLeaseTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPDHCPLeaseTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPDHCPLeaseTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIPDHCPLeaseTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		IpDhcpLeaseTime: ipDhcpLeaseTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPDHCPLeaseTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPDHCPLeaseTime")
		}

		// Simple Field (ipDhcpLeaseTime)
		if pushErr := writeBuffer.PushContext("ipDhcpLeaseTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipDhcpLeaseTime")
		}
		_ipDhcpLeaseTimeErr := writeBuffer.WriteSerializable(ctx, m.GetIpDhcpLeaseTime())
		if popErr := writeBuffer.PopContext("ipDhcpLeaseTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipDhcpLeaseTime")
		}
		if _ipDhcpLeaseTimeErr != nil {
			return errors.Wrap(_ipDhcpLeaseTimeErr, "Error serializing 'ipDhcpLeaseTime' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPDHCPLeaseTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPDHCPLeaseTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) isBACnetConstructedDataIPDHCPLeaseTime() bool {
	return true
}

func (m *_BACnetConstructedDataIPDHCPLeaseTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
