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

// BACnetConstructedDataSecuredStatus is the corresponding interface of BACnetConstructedDataSecuredStatus
type BACnetConstructedDataSecuredStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSecuredStatus returns SecuredStatus (property field)
	GetSecuredStatus() BACnetDoorSecuredStatusTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDoorSecuredStatusTagged
}

// BACnetConstructedDataSecuredStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSecuredStatus.
// This is useful for switch cases.
type BACnetConstructedDataSecuredStatusExactly interface {
	BACnetConstructedDataSecuredStatus
	isBACnetConstructedDataSecuredStatus() bool
}

// _BACnetConstructedDataSecuredStatus is the data-structure of this message
type _BACnetConstructedDataSecuredStatus struct {
	*_BACnetConstructedData
	SecuredStatus BACnetDoorSecuredStatusTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSecuredStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSecuredStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SECURED_STATUS
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSecuredStatus) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSecuredStatus) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSecuredStatus) GetSecuredStatus() BACnetDoorSecuredStatusTagged {
	return m.SecuredStatus
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSecuredStatus) GetActualValue() BACnetDoorSecuredStatusTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetDoorSecuredStatusTagged(m.GetSecuredStatus())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSecuredStatus factory function for _BACnetConstructedDataSecuredStatus
func NewBACnetConstructedDataSecuredStatus(securedStatus BACnetDoorSecuredStatusTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSecuredStatus {
	_result := &_BACnetConstructedDataSecuredStatus{
		SecuredStatus:          securedStatus,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSecuredStatus(structType any) BACnetConstructedDataSecuredStatus {
	if casted, ok := structType.(BACnetConstructedDataSecuredStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSecuredStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSecuredStatus) GetTypeName() string {
	return "BACnetConstructedDataSecuredStatus"
}

func (m *_BACnetConstructedDataSecuredStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (securedStatus)
	lengthInBits += m.SecuredStatus.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSecuredStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataSecuredStatusParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSecuredStatus, error) {
	return BACnetConstructedDataSecuredStatusParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataSecuredStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSecuredStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSecuredStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSecuredStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (securedStatus)
	if pullErr := readBuffer.PullContext("securedStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for securedStatus")
	}
	_securedStatus, _securedStatusErr := BACnetDoorSecuredStatusTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _securedStatusErr != nil {
		return nil, errors.Wrap(_securedStatusErr, "Error parsing 'securedStatus' field of BACnetConstructedDataSecuredStatus")
	}
	securedStatus := _securedStatus.(BACnetDoorSecuredStatusTagged)
	if closeErr := readBuffer.CloseContext("securedStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for securedStatus")
	}

	// Virtual field
	_actualValue := securedStatus
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSecuredStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSecuredStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSecuredStatus{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		SecuredStatus: securedStatus,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSecuredStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSecuredStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSecuredStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSecuredStatus")
		}

		// Simple Field (securedStatus)
		if pushErr := writeBuffer.PushContext("securedStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securedStatus")
		}
		_securedStatusErr := writeBuffer.WriteSerializable(ctx, m.GetSecuredStatus())
		if popErr := writeBuffer.PopContext("securedStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securedStatus")
		}
		if _securedStatusErr != nil {
			return errors.Wrap(_securedStatusErr, "Error serializing 'securedStatus' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSecuredStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSecuredStatus")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSecuredStatus) isBACnetConstructedDataSecuredStatus() bool {
	return true
}

func (m *_BACnetConstructedDataSecuredStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
