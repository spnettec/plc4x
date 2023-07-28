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


// BACnetRecipientProcess is the corresponding interface of BACnetRecipientProcess
type BACnetRecipientProcess interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetRecipient returns Recipient (property field)
	GetRecipient() BACnetRecipientEnclosed
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() BACnetContextTagUnsignedInteger
}

// BACnetRecipientProcessExactly can be used when we want exactly this type and not a type which fulfills BACnetRecipientProcess.
// This is useful for switch cases.
type BACnetRecipientProcessExactly interface {
	BACnetRecipientProcess
	isBACnetRecipientProcess() bool
}

// _BACnetRecipientProcess is the data-structure of this message
type _BACnetRecipientProcess struct {
        Recipient BACnetRecipientEnclosed
        ProcessIdentifier BACnetContextTagUnsignedInteger
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRecipientProcess) GetRecipient() BACnetRecipientEnclosed {
	return m.Recipient
}

func (m *_BACnetRecipientProcess) GetProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.ProcessIdentifier
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetRecipientProcess factory function for _BACnetRecipientProcess
func NewBACnetRecipientProcess( recipient BACnetRecipientEnclosed , processIdentifier BACnetContextTagUnsignedInteger ) *_BACnetRecipientProcess {
return &_BACnetRecipientProcess{ Recipient: recipient , ProcessIdentifier: processIdentifier }
}

// Deprecated: use the interface for direct cast
func CastBACnetRecipientProcess(structType any) BACnetRecipientProcess {
    if casted, ok := structType.(BACnetRecipientProcess); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRecipientProcess); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRecipientProcess) GetTypeName() string {
	return "BACnetRecipientProcess"
}

func (m *_BACnetRecipientProcess) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits(ctx)

	// Optional Field (processIdentifier)
	if m.ProcessIdentifier != nil {
		lengthInBits += m.ProcessIdentifier.GetLengthInBits(ctx)
	}

	return lengthInBits
}


func (m *_BACnetRecipientProcess) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRecipientProcessParse(ctx context.Context, theBytes []byte) (BACnetRecipientProcess, error) {
	return BACnetRecipientProcessParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetRecipientProcessParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetRecipientProcess, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetRecipientProcess"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRecipientProcess")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (recipient)
	if pullErr := readBuffer.PullContext("recipient"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for recipient")
	}
_recipient, _recipientErr := BACnetRecipientEnclosedParseWithBuffer(ctx, readBuffer , uint8( uint8(0) ) )
	if _recipientErr != nil {
		return nil, errors.Wrap(_recipientErr, "Error parsing 'recipient' field of BACnetRecipientProcess")
	}
	recipient := _recipient.(BACnetRecipientEnclosed)
	if closeErr := readBuffer.CloseContext("recipient"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for recipient")
	}

	// Optional Field (processIdentifier) (Can be skipped, if a given expression evaluates to false)
	var processIdentifier BACnetContextTagUnsignedInteger = nil
{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("processIdentifier"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for processIdentifier")
		}
_val, _err := BACnetContextTagParseWithBuffer(ctx, readBuffer , uint8(1) , BACnetDataType_UNSIGNED_INTEGER )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'processIdentifier' field of BACnetRecipientProcess")
		default:
			processIdentifier = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("processIdentifier"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for processIdentifier")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetRecipientProcess"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRecipientProcess")
	}

	// Create the instance
	return &_BACnetRecipientProcess{
			Recipient: recipient,
			ProcessIdentifier: processIdentifier,
		}, nil
}

func (m *_BACnetRecipientProcess) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetRecipientProcess) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("BACnetRecipientProcess"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRecipientProcess")
	}

	// Simple Field (recipient)
	if pushErr := writeBuffer.PushContext("recipient"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for recipient")
	}
	_recipientErr := writeBuffer.WriteSerializable(ctx, m.GetRecipient())
	if popErr := writeBuffer.PopContext("recipient"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for recipient")
	}
	if _recipientErr != nil {
		return errors.Wrap(_recipientErr, "Error serializing 'recipient' field")
	}

	// Optional Field (processIdentifier) (Can be skipped, if the value is null)
	var processIdentifier BACnetContextTagUnsignedInteger = nil
	if m.GetProcessIdentifier() != nil {
		if pushErr := writeBuffer.PushContext("processIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for processIdentifier")
		}
		processIdentifier = m.GetProcessIdentifier()
		_processIdentifierErr := writeBuffer.WriteSerializable(ctx, processIdentifier)
		if popErr := writeBuffer.PopContext("processIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for processIdentifier")
		}
		if _processIdentifierErr != nil {
			return errors.Wrap(_processIdentifierErr, "Error serializing 'processIdentifier' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetRecipientProcess"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRecipientProcess")
	}
	return nil
}


func (m *_BACnetRecipientProcess) isBACnetRecipientProcess() bool {
	return true
}

func (m *_BACnetRecipientProcess) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



