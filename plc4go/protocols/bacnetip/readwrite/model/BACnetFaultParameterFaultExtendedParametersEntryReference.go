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

// BACnetFaultParameterFaultExtendedParametersEntryReference is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryReference
type BACnetFaultParameterFaultExtendedParametersEntryReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetReference returns Reference (property field)
	GetReference() BACnetDeviceObjectPropertyReferenceEnclosed
}

// BACnetFaultParameterFaultExtendedParametersEntryReferenceExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultExtendedParametersEntryReference.
// This is useful for switch cases.
type BACnetFaultParameterFaultExtendedParametersEntryReferenceExactly interface {
	BACnetFaultParameterFaultExtendedParametersEntryReference
	isBACnetFaultParameterFaultExtendedParametersEntryReference() bool
}

// _BACnetFaultParameterFaultExtendedParametersEntryReference is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryReference struct {
	*_BACnetFaultParameterFaultExtendedParametersEntry
	Reference BACnetDeviceObjectPropertyReferenceEnclosed
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) InitializeParent(parent BACnetFaultParameterFaultExtendedParametersEntry, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) GetParent() BACnetFaultParameterFaultExtendedParametersEntry {
	return m._BACnetFaultParameterFaultExtendedParametersEntry
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) GetReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.Reference
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetFaultParameterFaultExtendedParametersEntryReference factory function for _BACnetFaultParameterFaultExtendedParametersEntryReference
func NewBACnetFaultParameterFaultExtendedParametersEntryReference(reference BACnetDeviceObjectPropertyReferenceEnclosed, peekedTagHeader BACnetTagHeader) *_BACnetFaultParameterFaultExtendedParametersEntryReference {
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryReference{
		Reference: reference,
		_BACnetFaultParameterFaultExtendedParametersEntry: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
	}
	_result._BACnetFaultParameterFaultExtendedParametersEntry._BACnetFaultParameterFaultExtendedParametersEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryReference(structType any) BACnetFaultParameterFaultExtendedParametersEntryReference {
	if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryReference"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (reference)
	lengthInBits += m.Reference.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultExtendedParametersEntryReferenceParse(ctx context.Context, theBytes []byte) (BACnetFaultParameterFaultExtendedParametersEntryReference, error) {
	return BACnetFaultParameterFaultExtendedParametersEntryReferenceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetFaultParameterFaultExtendedParametersEntryReferenceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultExtendedParametersEntryReference, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (reference)
	if pullErr := readBuffer.PullContext("reference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for reference")
	}
	_reference, _referenceErr := BACnetDeviceObjectPropertyReferenceEnclosedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)))
	if _referenceErr != nil {
		return nil, errors.Wrap(_referenceErr, "Error parsing 'reference' field of BACnetFaultParameterFaultExtendedParametersEntryReference")
	}
	reference := _reference.(BACnetDeviceObjectPropertyReferenceEnclosed)
	if closeErr := readBuffer.CloseContext("reference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for reference")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryReference")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultExtendedParametersEntryReference{
		_BACnetFaultParameterFaultExtendedParametersEntry: &_BACnetFaultParameterFaultExtendedParametersEntry{},
		Reference: reference,
	}
	_child._BACnetFaultParameterFaultExtendedParametersEntry._BACnetFaultParameterFaultExtendedParametersEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryReference")
		}

		// Simple Field (reference)
		if pushErr := writeBuffer.PushContext("reference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for reference")
		}
		_referenceErr := writeBuffer.WriteSerializable(ctx, m.GetReference())
		if popErr := writeBuffer.PopContext("reference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for reference")
		}
		if _referenceErr != nil {
			return errors.Wrap(_referenceErr, "Error serializing 'reference' field")
		}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryReference")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) isBACnetFaultParameterFaultExtendedParametersEntryReference() bool {
	return true
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
