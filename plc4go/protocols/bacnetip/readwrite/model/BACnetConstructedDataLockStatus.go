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

// BACnetConstructedDataLockStatus is the corresponding interface of BACnetConstructedDataLockStatus
type BACnetConstructedDataLockStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLockStatus returns LockStatus (property field)
	GetLockStatus() BACnetLockStatusTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLockStatusTagged
}

// BACnetConstructedDataLockStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLockStatus.
// This is useful for switch cases.
type BACnetConstructedDataLockStatusExactly interface {
	BACnetConstructedDataLockStatus
	isBACnetConstructedDataLockStatus() bool
}

// _BACnetConstructedDataLockStatus is the data-structure of this message
type _BACnetConstructedDataLockStatus struct {
	*_BACnetConstructedData
	LockStatus BACnetLockStatusTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLockStatus) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataLockStatus) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOCK_STATUS
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLockStatus) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLockStatus) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLockStatus) GetLockStatus() BACnetLockStatusTagged {
	return m.LockStatus
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLockStatus) GetActualValue() BACnetLockStatusTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLockStatusTagged(m.GetLockStatus())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataLockStatus factory function for _BACnetConstructedDataLockStatus
func NewBACnetConstructedDataLockStatus(lockStatus BACnetLockStatusTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataLockStatus {
	_result := &_BACnetConstructedDataLockStatus{
		LockStatus:             lockStatus,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLockStatus(structType any) BACnetConstructedDataLockStatus {
	if casted, ok := structType.(BACnetConstructedDataLockStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLockStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLockStatus) GetTypeName() string {
	return "BACnetConstructedDataLockStatus"
}

func (m *_BACnetConstructedDataLockStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (lockStatus)
	lengthInBits += m.LockStatus.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataLockStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLockStatusParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLockStatus, error) {
	return BACnetConstructedDataLockStatusParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLockStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLockStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLockStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLockStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lockStatus)
	if pullErr := readBuffer.PullContext("lockStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lockStatus")
	}
	_lockStatus, _lockStatusErr := BACnetLockStatusTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _lockStatusErr != nil {
		return nil, errors.Wrap(_lockStatusErr, "Error parsing 'lockStatus' field of BACnetConstructedDataLockStatus")
	}
	lockStatus := _lockStatus.(BACnetLockStatusTagged)
	if closeErr := readBuffer.CloseContext("lockStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lockStatus")
	}

	// Virtual field
	_actualValue := lockStatus
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLockStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLockStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLockStatus{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LockStatus: lockStatus,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLockStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLockStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLockStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLockStatus")
		}

		// Simple Field (lockStatus)
		if pushErr := writeBuffer.PushContext("lockStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for lockStatus")
		}
		_lockStatusErr := writeBuffer.WriteSerializable(ctx, m.GetLockStatus())
		if popErr := writeBuffer.PopContext("lockStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for lockStatus")
		}
		if _lockStatusErr != nil {
			return errors.Wrap(_lockStatusErr, "Error serializing 'lockStatus' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLockStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLockStatus")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataLockStatus) isBACnetConstructedDataLockStatus() bool {
	return true
}

func (m *_BACnetConstructedDataLockStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
