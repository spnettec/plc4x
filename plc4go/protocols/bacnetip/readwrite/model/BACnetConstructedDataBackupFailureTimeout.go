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

// BACnetConstructedDataBackupFailureTimeout is the corresponding interface of BACnetConstructedDataBackupFailureTimeout
type BACnetConstructedDataBackupFailureTimeout interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBackupFailureTimeout returns BackupFailureTimeout (property field)
	GetBackupFailureTimeout() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataBackupFailureTimeoutExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBackupFailureTimeout.
// This is useful for switch cases.
type BACnetConstructedDataBackupFailureTimeoutExactly interface {
	BACnetConstructedDataBackupFailureTimeout
	isBACnetConstructedDataBackupFailureTimeout() bool
}

// _BACnetConstructedDataBackupFailureTimeout is the data-structure of this message
type _BACnetConstructedDataBackupFailureTimeout struct {
	*_BACnetConstructedData
	BackupFailureTimeout BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBackupFailureTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataBackupFailureTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_BACKUP_FAILURE_TIMEOUT
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBackupFailureTimeout) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBackupFailureTimeout) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBackupFailureTimeout) GetBackupFailureTimeout() BACnetApplicationTagUnsignedInteger {
	return m.BackupFailureTimeout
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBackupFailureTimeout) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetBackupFailureTimeout())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataBackupFailureTimeout factory function for _BACnetConstructedDataBackupFailureTimeout
func NewBACnetConstructedDataBackupFailureTimeout(backupFailureTimeout BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataBackupFailureTimeout {
	_result := &_BACnetConstructedDataBackupFailureTimeout{
		BackupFailureTimeout:   backupFailureTimeout,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBackupFailureTimeout(structType any) BACnetConstructedDataBackupFailureTimeout {
	if casted, ok := structType.(BACnetConstructedDataBackupFailureTimeout); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBackupFailureTimeout); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBackupFailureTimeout) GetTypeName() string {
	return "BACnetConstructedDataBackupFailureTimeout"
}

func (m *_BACnetConstructedDataBackupFailureTimeout) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (backupFailureTimeout)
	lengthInBits += m.BackupFailureTimeout.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataBackupFailureTimeout) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBackupFailureTimeoutParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBackupFailureTimeout, error) {
	return BACnetConstructedDataBackupFailureTimeoutParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBackupFailureTimeoutParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBackupFailureTimeout, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBackupFailureTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBackupFailureTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (backupFailureTimeout)
	if pullErr := readBuffer.PullContext("backupFailureTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for backupFailureTimeout")
	}
	_backupFailureTimeout, _backupFailureTimeoutErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _backupFailureTimeoutErr != nil {
		return nil, errors.Wrap(_backupFailureTimeoutErr, "Error parsing 'backupFailureTimeout' field of BACnetConstructedDataBackupFailureTimeout")
	}
	backupFailureTimeout := _backupFailureTimeout.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("backupFailureTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for backupFailureTimeout")
	}

	// Virtual field
	_actualValue := backupFailureTimeout
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBackupFailureTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBackupFailureTimeout")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBackupFailureTimeout{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		BackupFailureTimeout: backupFailureTimeout,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBackupFailureTimeout) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBackupFailureTimeout) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBackupFailureTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBackupFailureTimeout")
		}

		// Simple Field (backupFailureTimeout)
		if pushErr := writeBuffer.PushContext("backupFailureTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for backupFailureTimeout")
		}
		_backupFailureTimeoutErr := writeBuffer.WriteSerializable(ctx, m.GetBackupFailureTimeout())
		if popErr := writeBuffer.PopContext("backupFailureTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for backupFailureTimeout")
		}
		if _backupFailureTimeoutErr != nil {
			return errors.Wrap(_backupFailureTimeoutErr, "Error serializing 'backupFailureTimeout' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBackupFailureTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBackupFailureTimeout")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataBackupFailureTimeout) isBACnetConstructedDataBackupFailureTimeout() bool {
	return true
}

func (m *_BACnetConstructedDataBackupFailureTimeout) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
