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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyStatesReasonForHalt is the corresponding interface of BACnetPropertyStatesReasonForHalt
type BACnetPropertyStatesReasonForHalt interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetReasonForHalt returns ReasonForHalt (property field)
	GetReasonForHalt() BACnetProgramErrorTagged
	// IsBACnetPropertyStatesReasonForHalt is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesReasonForHalt()
}

// _BACnetPropertyStatesReasonForHalt is the data-structure of this message
type _BACnetPropertyStatesReasonForHalt struct {
	BACnetPropertyStatesContract
	ReasonForHalt BACnetProgramErrorTagged
}

var _ BACnetPropertyStatesReasonForHalt = (*_BACnetPropertyStatesReasonForHalt)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesReasonForHalt)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesReasonForHalt) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesReasonForHalt) GetReasonForHalt() BACnetProgramErrorTagged {
	return m.ReasonForHalt
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesReasonForHalt factory function for _BACnetPropertyStatesReasonForHalt
func NewBACnetPropertyStatesReasonForHalt(peekedTagHeader BACnetTagHeader, reasonForHalt BACnetProgramErrorTagged) *_BACnetPropertyStatesReasonForHalt {
	if reasonForHalt == nil {
		panic("reasonForHalt of type BACnetProgramErrorTagged for BACnetPropertyStatesReasonForHalt must not be nil")
	}
	_result := &_BACnetPropertyStatesReasonForHalt{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		ReasonForHalt:                reasonForHalt,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesReasonForHalt(structType any) BACnetPropertyStatesReasonForHalt {
	if casted, ok := structType.(BACnetPropertyStatesReasonForHalt); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesReasonForHalt); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesReasonForHalt) GetTypeName() string {
	return "BACnetPropertyStatesReasonForHalt"
}

func (m *_BACnetPropertyStatesReasonForHalt) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (reasonForHalt)
	lengthInBits += m.ReasonForHalt.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesReasonForHalt) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesReasonForHalt) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesReasonForHalt BACnetPropertyStatesReasonForHalt, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesReasonForHalt"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesReasonForHalt")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reasonForHalt, err := ReadSimpleField[BACnetProgramErrorTagged](ctx, "reasonForHalt", ReadComplex[BACnetProgramErrorTagged](BACnetProgramErrorTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reasonForHalt' field"))
	}
	m.ReasonForHalt = reasonForHalt

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesReasonForHalt"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesReasonForHalt")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesReasonForHalt) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesReasonForHalt) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesReasonForHalt"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesReasonForHalt")
		}

		if err := WriteSimpleField[BACnetProgramErrorTagged](ctx, "reasonForHalt", m.GetReasonForHalt(), WriteComplex[BACnetProgramErrorTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'reasonForHalt' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesReasonForHalt"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesReasonForHalt")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesReasonForHalt) IsBACnetPropertyStatesReasonForHalt() {}

func (m *_BACnetPropertyStatesReasonForHalt) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
