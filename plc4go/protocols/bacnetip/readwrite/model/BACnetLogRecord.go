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
	"io"
	"github.com/rs/zerolog"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetLogRecord is the corresponding interface of BACnetLogRecord
type BACnetLogRecord interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() BACnetDateTimeEnclosed
	// GetLogDatum returns LogDatum (property field)
	GetLogDatum() BACnetLogRecordLogDatum
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
}

// BACnetLogRecordExactly can be used when we want exactly this type and not a type which fulfills BACnetLogRecord.
// This is useful for switch cases.
type BACnetLogRecordExactly interface {
	BACnetLogRecord
	isBACnetLogRecord() bool
}

// _BACnetLogRecord is the data-structure of this message
type _BACnetLogRecord struct {
        Timestamp BACnetDateTimeEnclosed
        LogDatum BACnetLogRecordLogDatum
        StatusFlags BACnetStatusFlagsTagged
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogRecord) GetTimestamp() BACnetDateTimeEnclosed {
	return m.Timestamp
}

func (m *_BACnetLogRecord) GetLogDatum() BACnetLogRecordLogDatum {
	return m.LogDatum
}

func (m *_BACnetLogRecord) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetLogRecord factory function for _BACnetLogRecord
func NewBACnetLogRecord( timestamp BACnetDateTimeEnclosed , logDatum BACnetLogRecordLogDatum , statusFlags BACnetStatusFlagsTagged ) *_BACnetLogRecord {
return &_BACnetLogRecord{ Timestamp: timestamp , LogDatum: logDatum , StatusFlags: statusFlags }
}

// Deprecated: use the interface for direct cast
func CastBACnetLogRecord(structType any) BACnetLogRecord {
    if casted, ok := structType.(BACnetLogRecord); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogRecord); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogRecord) GetTypeName() string {
	return "BACnetLogRecord"
}

func (m *_BACnetLogRecord) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (timestamp)
	lengthInBits += m.Timestamp.GetLengthInBits(ctx)

	// Simple field (logDatum)
	lengthInBits += m.LogDatum.GetLengthInBits(ctx)

	// Optional Field (statusFlags)
	if m.StatusFlags != nil {
		lengthInBits += m.StatusFlags.GetLengthInBits(ctx)
	}

	return lengthInBits
}


func (m *_BACnetLogRecord) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetLogRecordParse(ctx context.Context, theBytes []byte) (BACnetLogRecord, error) {
	return BACnetLogRecordParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetLogRecordParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetLogRecord, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetLogRecord"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogRecord")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timestamp)
	if pullErr := readBuffer.PullContext("timestamp"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timestamp")
	}
_timestamp, _timestampErr := BACnetDateTimeEnclosedParseWithBuffer(ctx, readBuffer , uint8( uint8(0) ) )
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field of BACnetLogRecord")
	}
	timestamp := _timestamp.(BACnetDateTimeEnclosed)
	if closeErr := readBuffer.CloseContext("timestamp"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timestamp")
	}

	// Simple Field (logDatum)
	if pullErr := readBuffer.PullContext("logDatum"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for logDatum")
	}
_logDatum, _logDatumErr := BACnetLogRecordLogDatumParseWithBuffer(ctx, readBuffer , uint8( uint8(1) ) )
	if _logDatumErr != nil {
		return nil, errors.Wrap(_logDatumErr, "Error parsing 'logDatum' field of BACnetLogRecord")
	}
	logDatum := _logDatum.(BACnetLogRecordLogDatum)
	if closeErr := readBuffer.CloseContext("logDatum"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for logDatum")
	}

	// Optional Field (statusFlags) (Can be skipped, if a given expression evaluates to false)
	var statusFlags BACnetStatusFlagsTagged = nil
{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for statusFlags")
		}
_val, _err := BACnetStatusFlagsTaggedParseWithBuffer(ctx, readBuffer , uint8(2) , TagClass_CONTEXT_SPECIFIC_TAGS )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'statusFlags' field of BACnetLogRecord")
		default:
			statusFlags = _val.(BACnetStatusFlagsTagged)
			if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for statusFlags")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetLogRecord"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogRecord")
	}

	// Create the instance
	return &_BACnetLogRecord{
			Timestamp: timestamp,
			LogDatum: logDatum,
			StatusFlags: statusFlags,
		}, nil
}

func (m *_BACnetLogRecord) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLogRecord) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("BACnetLogRecord"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLogRecord")
	}

	// Simple Field (timestamp)
	if pushErr := writeBuffer.PushContext("timestamp"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timestamp")
	}
	_timestampErr := writeBuffer.WriteSerializable(ctx, m.GetTimestamp())
	if popErr := writeBuffer.PopContext("timestamp"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timestamp")
	}
	if _timestampErr != nil {
		return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
	}

	// Simple Field (logDatum)
	if pushErr := writeBuffer.PushContext("logDatum"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for logDatum")
	}
	_logDatumErr := writeBuffer.WriteSerializable(ctx, m.GetLogDatum())
	if popErr := writeBuffer.PopContext("logDatum"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for logDatum")
	}
	if _logDatumErr != nil {
		return errors.Wrap(_logDatumErr, "Error serializing 'logDatum' field")
	}

	// Optional Field (statusFlags) (Can be skipped, if the value is null)
	var statusFlags BACnetStatusFlagsTagged = nil
	if m.GetStatusFlags() != nil {
		if pushErr := writeBuffer.PushContext("statusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusFlags")
		}
		statusFlags = m.GetStatusFlags()
		_statusFlagsErr := writeBuffer.WriteSerializable(ctx, statusFlags)
		if popErr := writeBuffer.PopContext("statusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusFlags")
		}
		if _statusFlagsErr != nil {
			return errors.Wrap(_statusFlagsErr, "Error serializing 'statusFlags' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetLogRecord"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLogRecord")
	}
	return nil
}


func (m *_BACnetLogRecord) isBACnetLogRecord() bool {
	return true
}

func (m *_BACnetLogRecord) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



