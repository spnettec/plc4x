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

// BACnetFaultParameterFaultExtendedParametersEntryBitString is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryBitString
type BACnetFaultParameterFaultExtendedParametersEntryBitString interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetBitStringValue returns BitStringValue (property field)
	GetBitStringValue() BACnetApplicationTagBitString
}

// BACnetFaultParameterFaultExtendedParametersEntryBitStringExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultExtendedParametersEntryBitString.
// This is useful for switch cases.
type BACnetFaultParameterFaultExtendedParametersEntryBitStringExactly interface {
	BACnetFaultParameterFaultExtendedParametersEntryBitString
	isBACnetFaultParameterFaultExtendedParametersEntryBitString() bool
}

// _BACnetFaultParameterFaultExtendedParametersEntryBitString is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryBitString struct {
	*_BACnetFaultParameterFaultExtendedParametersEntry
	BitStringValue BACnetApplicationTagBitString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) InitializeParent(parent BACnetFaultParameterFaultExtendedParametersEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) GetParent() BACnetFaultParameterFaultExtendedParametersEntry {
	return m._BACnetFaultParameterFaultExtendedParametersEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) GetBitStringValue() BACnetApplicationTagBitString {
	return m.BitStringValue
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtendedParametersEntryBitString factory function for _BACnetFaultParameterFaultExtendedParametersEntryBitString
func NewBACnetFaultParameterFaultExtendedParametersEntryBitString(bitStringValue BACnetApplicationTagBitString, peekedTagHeader BACnetTagHeader) *_BACnetFaultParameterFaultExtendedParametersEntryBitString {
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryBitString{
		BitStringValue: bitStringValue,
		_BACnetFaultParameterFaultExtendedParametersEntry: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
	}
	_result._BACnetFaultParameterFaultExtendedParametersEntry._BACnetFaultParameterFaultExtendedParametersEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryBitString(structType any) BACnetFaultParameterFaultExtendedParametersEntryBitString {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryBitString); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryBitString); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryBitString"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (bitStringValue)
	lengthInBits += m.BitStringValue.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultExtendedParametersEntryBitStringParse(ctx context.Context, theBytes []byte) (BACnetFaultParameterFaultExtendedParametersEntryBitString, error) {
	return BACnetFaultParameterFaultExtendedParametersEntryBitStringParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetFaultParameterFaultExtendedParametersEntryBitStringParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultExtendedParametersEntryBitString, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryBitString"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryBitString")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bitStringValue)
	if pullErr := readBuffer.PullContext("bitStringValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bitStringValue")
	}
	_bitStringValue, _bitStringValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _bitStringValueErr != nil {
		return nil, errors.Wrap(_bitStringValueErr, "Error parsing 'bitStringValue' field of BACnetFaultParameterFaultExtendedParametersEntryBitString")
	}
	bitStringValue := _bitStringValue.(BACnetApplicationTagBitString)
	if closeErr := readBuffer.CloseContext("bitStringValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bitStringValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryBitString"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryBitString")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultExtendedParametersEntryBitString{
		_BACnetFaultParameterFaultExtendedParametersEntry: &_BACnetFaultParameterFaultExtendedParametersEntry{},
		BitStringValue: bitStringValue,
	}
	_child._BACnetFaultParameterFaultExtendedParametersEntry._BACnetFaultParameterFaultExtendedParametersEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryBitString"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryBitString")
		}

		// Simple Field (bitStringValue)
		if pushErr := writeBuffer.PushContext("bitStringValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bitStringValue")
		}
		_bitStringValueErr := writeBuffer.WriteSerializable(ctx, m.GetBitStringValue())
		if popErr := writeBuffer.PopContext("bitStringValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bitStringValue")
		}
		if _bitStringValueErr != nil {
			return errors.Wrap(_bitStringValueErr, "Error serializing 'bitStringValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryBitString"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryBitString")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) isBACnetFaultParameterFaultExtendedParametersEntryBitString() bool {
	return true
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryBitString) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
