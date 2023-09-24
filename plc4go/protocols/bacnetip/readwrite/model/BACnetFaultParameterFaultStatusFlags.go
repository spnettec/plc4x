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


// BACnetFaultParameterFaultStatusFlags is the corresponding interface of BACnetFaultParameterFaultStatusFlags
type BACnetFaultParameterFaultStatusFlags interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetStatusFlagsReference returns StatusFlagsReference (property field)
	GetStatusFlagsReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetFaultParameterFaultStatusFlagsExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultStatusFlags.
// This is useful for switch cases.
type BACnetFaultParameterFaultStatusFlagsExactly interface {
	BACnetFaultParameterFaultStatusFlags
	isBACnetFaultParameterFaultStatusFlags() bool
}

// _BACnetFaultParameterFaultStatusFlags is the data-structure of this message
type _BACnetFaultParameterFaultStatusFlags struct {
	*_BACnetFaultParameter
        OpeningTag BACnetOpeningTag
        StatusFlagsReference BACnetDeviceObjectPropertyReferenceEnclosed
        ClosingTag BACnetClosingTag
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultStatusFlags) InitializeParent(parent BACnetFaultParameter , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetFaultParameterFaultStatusFlags)  GetParent() BACnetFaultParameter {
	return m._BACnetFaultParameter
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultStatusFlags) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultStatusFlags) GetStatusFlagsReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.StatusFlagsReference
}

func (m *_BACnetFaultParameterFaultStatusFlags) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetFaultParameterFaultStatusFlags factory function for _BACnetFaultParameterFaultStatusFlags
func NewBACnetFaultParameterFaultStatusFlags( openingTag BACnetOpeningTag , statusFlagsReference BACnetDeviceObjectPropertyReferenceEnclosed , closingTag BACnetClosingTag , peekedTagHeader BACnetTagHeader ) *_BACnetFaultParameterFaultStatusFlags {
	_result := &_BACnetFaultParameterFaultStatusFlags{
		OpeningTag: openingTag,
		StatusFlagsReference: statusFlagsReference,
		ClosingTag: closingTag,
    	_BACnetFaultParameter: NewBACnetFaultParameter(peekedTagHeader),
	}
	_result._BACnetFaultParameter._BACnetFaultParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultStatusFlags(structType any) BACnetFaultParameterFaultStatusFlags {
    if casted, ok := structType.(BACnetFaultParameterFaultStatusFlags); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultStatusFlags); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultStatusFlags) GetTypeName() string {
	return "BACnetFaultParameterFaultStatusFlags"
}

func (m *_BACnetFaultParameterFaultStatusFlags) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits(ctx)

	// Simple field (statusFlagsReference)
	lengthInBits += m.StatusFlagsReference.GetLengthInBits(ctx)

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetFaultParameterFaultStatusFlags) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultStatusFlagsParse(ctx context.Context, theBytes []byte) (BACnetFaultParameterFaultStatusFlags, error) {
	return BACnetFaultParameterFaultStatusFlagsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetFaultParameterFaultStatusFlagsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultStatusFlags, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultStatusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultStatusFlags")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer , uint8( uint8(5) ) )
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetFaultParameterFaultStatusFlags")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (statusFlagsReference)
	if pullErr := readBuffer.PullContext("statusFlagsReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusFlagsReference")
	}
_statusFlagsReference, _statusFlagsReferenceErr := BACnetDeviceObjectPropertyReferenceEnclosedParseWithBuffer(ctx, readBuffer , uint8( uint8(1) ) )
	if _statusFlagsReferenceErr != nil {
		return nil, errors.Wrap(_statusFlagsReferenceErr, "Error parsing 'statusFlagsReference' field of BACnetFaultParameterFaultStatusFlags")
	}
	statusFlagsReference := _statusFlagsReference.(BACnetDeviceObjectPropertyReferenceEnclosed)
	if closeErr := readBuffer.CloseContext("statusFlagsReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusFlagsReference")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer , uint8( uint8(5) ) )
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetFaultParameterFaultStatusFlags")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultStatusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultStatusFlags")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultStatusFlags{
		_BACnetFaultParameter: &_BACnetFaultParameter{
		},
		OpeningTag: openingTag,
		StatusFlagsReference: statusFlagsReference,
		ClosingTag: closingTag,
	}
	_child._BACnetFaultParameter._BACnetFaultParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultStatusFlags) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultStatusFlags) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultStatusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultStatusFlags")
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

	// Simple Field (statusFlagsReference)
	if pushErr := writeBuffer.PushContext("statusFlagsReference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for statusFlagsReference")
	}
	_statusFlagsReferenceErr := writeBuffer.WriteSerializable(ctx, m.GetStatusFlagsReference())
	if popErr := writeBuffer.PopContext("statusFlagsReference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for statusFlagsReference")
	}
	if _statusFlagsReferenceErr != nil {
		return errors.Wrap(_statusFlagsReferenceErr, "Error serializing 'statusFlagsReference' field")
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

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultStatusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultStatusFlags")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetFaultParameterFaultStatusFlags) isBACnetFaultParameterFaultStatusFlags() bool {
	return true
}

func (m *_BACnetFaultParameterFaultStatusFlags) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



