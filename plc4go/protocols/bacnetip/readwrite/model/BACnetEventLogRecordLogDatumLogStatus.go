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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetEventLogRecordLogDatumLogStatus is the corresponding interface of BACnetEventLogRecordLogDatumLogStatus
type BACnetEventLogRecordLogDatumLogStatus interface {
	utils.LengthAware
	utils.Serializable
	BACnetEventLogRecordLogDatum
	// GetLogStatus returns LogStatus (property field)
	GetLogStatus() BACnetLogStatusTagged
}

// BACnetEventLogRecordLogDatumLogStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetEventLogRecordLogDatumLogStatus.
// This is useful for switch cases.
type BACnetEventLogRecordLogDatumLogStatusExactly interface {
	BACnetEventLogRecordLogDatumLogStatus
	isBACnetEventLogRecordLogDatumLogStatus() bool
}

// _BACnetEventLogRecordLogDatumLogStatus is the data-structure of this message
type _BACnetEventLogRecordLogDatumLogStatus struct {
	*_BACnetEventLogRecordLogDatum
	LogStatus BACnetLogStatusTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventLogRecordLogDatumLogStatus) InitializeParent(parent BACnetEventLogRecordLogDatum, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) GetParent() BACnetEventLogRecordLogDatum {
	return m._BACnetEventLogRecordLogDatum
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventLogRecordLogDatumLogStatus) GetLogStatus() BACnetLogStatusTagged {
	return m.LogStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetEventLogRecordLogDatumLogStatus factory function for _BACnetEventLogRecordLogDatumLogStatus
func NewBACnetEventLogRecordLogDatumLogStatus(logStatus BACnetLogStatusTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8) *_BACnetEventLogRecordLogDatumLogStatus {
	_result := &_BACnetEventLogRecordLogDatumLogStatus{
		LogStatus:                     logStatus,
		_BACnetEventLogRecordLogDatum: NewBACnetEventLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetEventLogRecordLogDatum._BACnetEventLogRecordLogDatumChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventLogRecordLogDatumLogStatus(structType interface{}) BACnetEventLogRecordLogDatumLogStatus {
	if casted, ok := structType.(BACnetEventLogRecordLogDatumLogStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventLogRecordLogDatumLogStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) GetTypeName() string {
	return "BACnetEventLogRecordLogDatumLogStatus"
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (logStatus)
	lengthInBits += m.LogStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventLogRecordLogDatumLogStatusParse(theBytes []byte, tagNumber uint8) (BACnetEventLogRecordLogDatumLogStatus, error) {
	return BACnetEventLogRecordLogDatumLogStatusParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetEventLogRecordLogDatumLogStatusParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetEventLogRecordLogDatumLogStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventLogRecordLogDatumLogStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventLogRecordLogDatumLogStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (logStatus)
	if pullErr := readBuffer.PullContext("logStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for logStatus")
	}
	_logStatus, _logStatusErr := BACnetLogStatusTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _logStatusErr != nil {
		return nil, errors.Wrap(_logStatusErr, "Error parsing 'logStatus' field of BACnetEventLogRecordLogDatumLogStatus")
	}
	logStatus := _logStatus.(BACnetLogStatusTagged)
	if closeErr := readBuffer.CloseContext("logStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for logStatus")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventLogRecordLogDatumLogStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventLogRecordLogDatumLogStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventLogRecordLogDatumLogStatus{
		_BACnetEventLogRecordLogDatum: &_BACnetEventLogRecordLogDatum{
			TagNumber: tagNumber,
		},
		LogStatus: logStatus,
	}
	_child._BACnetEventLogRecordLogDatum._BACnetEventLogRecordLogDatumChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventLogRecordLogDatumLogStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventLogRecordLogDatumLogStatus")
		}

		// Simple Field (logStatus)
		if pushErr := writeBuffer.PushContext("logStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for logStatus")
		}
		_logStatusErr := writeBuffer.WriteSerializable(ctx, m.GetLogStatus())
		if popErr := writeBuffer.PopContext("logStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for logStatus")
		}
		if _logStatusErr != nil {
			return errors.Wrap(_logStatusErr, "Error serializing 'logStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventLogRecordLogDatumLogStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventLogRecordLogDatumLogStatus")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) isBACnetEventLogRecordLogDatumLogStatus() bool {
	return true
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
