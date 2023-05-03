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

// BACnetConstructedDataIPDHCPEnable is the corresponding interface of BACnetConstructedDataIPDHCPEnable
type BACnetConstructedDataIPDHCPEnable interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetIpDhcpEnable returns IpDhcpEnable (property field)
	GetIpDhcpEnable() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataIPDHCPEnableExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataIPDHCPEnable.
// This is useful for switch cases.
type BACnetConstructedDataIPDHCPEnableExactly interface {
	BACnetConstructedDataIPDHCPEnable
	isBACnetConstructedDataIPDHCPEnable() bool
}

// _BACnetConstructedDataIPDHCPEnable is the data-structure of this message
type _BACnetConstructedDataIPDHCPEnable struct {
	*_BACnetConstructedData
	IpDhcpEnable BACnetApplicationTagBoolean
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPEnable) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataIPDHCPEnable) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_IP_DHCP_ENABLE
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataIPDHCPEnable) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataIPDHCPEnable) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPEnable) GetIpDhcpEnable() BACnetApplicationTagBoolean {
	return m.IpDhcpEnable
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataIPDHCPEnable) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetIpDhcpEnable())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataIPDHCPEnable factory function for _BACnetConstructedDataIPDHCPEnable
func NewBACnetConstructedDataIPDHCPEnable(ipDhcpEnable BACnetApplicationTagBoolean, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataIPDHCPEnable {
	_result := &_BACnetConstructedDataIPDHCPEnable{
		IpDhcpEnable:           ipDhcpEnable,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataIPDHCPEnable(structType any) BACnetConstructedDataIPDHCPEnable {
	if casted, ok := structType.(BACnetConstructedDataIPDHCPEnable); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataIPDHCPEnable); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataIPDHCPEnable) GetTypeName() string {
	return "BACnetConstructedDataIPDHCPEnable"
}

func (m *_BACnetConstructedDataIPDHCPEnable) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (ipDhcpEnable)
	lengthInBits += m.IpDhcpEnable.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataIPDHCPEnable) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataIPDHCPEnableParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPDHCPEnable, error) {
	return BACnetConstructedDataIPDHCPEnableParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataIPDHCPEnableParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataIPDHCPEnable, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataIPDHCPEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataIPDHCPEnable")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ipDhcpEnable)
	if pullErr := readBuffer.PullContext("ipDhcpEnable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ipDhcpEnable")
	}
	_ipDhcpEnable, _ipDhcpEnableErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _ipDhcpEnableErr != nil {
		return nil, errors.Wrap(_ipDhcpEnableErr, "Error parsing 'ipDhcpEnable' field of BACnetConstructedDataIPDHCPEnable")
	}
	ipDhcpEnable := _ipDhcpEnable.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("ipDhcpEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ipDhcpEnable")
	}

	// Virtual field
	_actualValue := ipDhcpEnable
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataIPDHCPEnable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataIPDHCPEnable")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataIPDHCPEnable{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		IpDhcpEnable: ipDhcpEnable,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataIPDHCPEnable) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataIPDHCPEnable) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataIPDHCPEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataIPDHCPEnable")
		}

		// Simple Field (ipDhcpEnable)
		if pushErr := writeBuffer.PushContext("ipDhcpEnable"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ipDhcpEnable")
		}
		_ipDhcpEnableErr := writeBuffer.WriteSerializable(ctx, m.GetIpDhcpEnable())
		if popErr := writeBuffer.PopContext("ipDhcpEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ipDhcpEnable")
		}
		if _ipDhcpEnableErr != nil {
			return errors.Wrap(_ipDhcpEnableErr, "Error serializing 'ipDhcpEnable' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataIPDHCPEnable"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataIPDHCPEnable")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataIPDHCPEnable) isBACnetConstructedDataIPDHCPEnable() bool {
	return true
}

func (m *_BACnetConstructedDataIPDHCPEnable) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
