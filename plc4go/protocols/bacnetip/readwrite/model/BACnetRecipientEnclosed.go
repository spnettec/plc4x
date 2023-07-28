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


// BACnetRecipientEnclosed is the corresponding interface of BACnetRecipientEnclosed
type BACnetRecipientEnclosed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetRecipient returns Recipient (property field)
	GetRecipient() BACnetRecipient
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetRecipientEnclosedExactly can be used when we want exactly this type and not a type which fulfills BACnetRecipientEnclosed.
// This is useful for switch cases.
type BACnetRecipientEnclosedExactly interface {
	BACnetRecipientEnclosed
	isBACnetRecipientEnclosed() bool
}

// _BACnetRecipientEnclosed is the data-structure of this message
type _BACnetRecipientEnclosed struct {
        OpeningTag BACnetOpeningTag
        Recipient BACnetRecipient
        ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetRecipientEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetRecipientEnclosed) GetRecipient() BACnetRecipient {
	return m.Recipient
}

func (m *_BACnetRecipientEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetRecipientEnclosed factory function for _BACnetRecipientEnclosed
func NewBACnetRecipientEnclosed( openingTag BACnetOpeningTag , recipient BACnetRecipient , closingTag BACnetClosingTag , tagNumber uint8 ) *_BACnetRecipientEnclosed {
return &_BACnetRecipientEnclosed{ OpeningTag: openingTag , Recipient: recipient , ClosingTag: closingTag , TagNumber: tagNumber }
}

// Deprecated: use the interface for direct cast
func CastBACnetRecipientEnclosed(structType any) BACnetRecipientEnclosed {
    if casted, ok := structType.(BACnetRecipientEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetRecipientEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetRecipientEnclosed) GetTypeName() string {
	return "BACnetRecipientEnclosed"
}

func (m *_BACnetRecipientEnclosed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (recipient)
	lengthInBits += m.Recipient.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetRecipientEnclosed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetRecipientEnclosedParse(ctx context.Context, theBytes []byte, tagNumber uint8) (BACnetRecipientEnclosed, error) {
	return BACnetRecipientEnclosedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func BACnetRecipientEnclosedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetRecipientEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetRecipientEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetRecipientEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer , uint8( tagNumber ) )
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetRecipientEnclosed")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (recipient)
	if pullErr := readBuffer.PullContext("recipient"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for recipient")
	}
_recipient, _recipientErr := BACnetRecipientParseWithBuffer(ctx, readBuffer)
	if _recipientErr != nil {
		return nil, errors.Wrap(_recipientErr, "Error parsing 'recipient' field of BACnetRecipientEnclosed")
	}
	recipient := _recipient.(BACnetRecipient)
	if closeErr := readBuffer.CloseContext("recipient"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for recipient")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer , uint8( tagNumber ) )
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetRecipientEnclosed")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetRecipientEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetRecipientEnclosed")
	}

	// Create the instance
	return &_BACnetRecipientEnclosed{
            TagNumber: tagNumber,
			OpeningTag: openingTag,
			Recipient: recipient,
			ClosingTag: closingTag,
		}, nil
}

func (m *_BACnetRecipientEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetRecipientEnclosed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("BACnetRecipientEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetRecipientEnclosed")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(ctx, m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
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

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(ctx, m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetRecipientEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetRecipientEnclosed")
	}
	return nil
}


////
// Arguments Getter

func (m *_BACnetRecipientEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}
//
////

func (m *_BACnetRecipientEnclosed) isBACnetRecipientEnclosed() bool {
	return true
}

func (m *_BACnetRecipientEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



