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


// BACnetConstructedDataGroupMembers is the corresponding interface of BACnetConstructedDataGroupMembers
type BACnetConstructedDataGroupMembers interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetGroupMembers returns GroupMembers (property field)
	GetGroupMembers() []BACnetApplicationTagObjectIdentifier
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataGroupMembersExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataGroupMembers.
// This is useful for switch cases.
type BACnetConstructedDataGroupMembersExactly interface {
	BACnetConstructedDataGroupMembers
	isBACnetConstructedDataGroupMembers() bool
}

// _BACnetConstructedDataGroupMembers is the data-structure of this message
type _BACnetConstructedDataGroupMembers struct {
	*_BACnetConstructedData
        NumberOfDataElements BACnetApplicationTagUnsignedInteger
        GroupMembers []BACnetApplicationTagObjectIdentifier
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataGroupMembers)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataGroupMembers)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_GROUP_MEMBERS}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataGroupMembers) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataGroupMembers)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataGroupMembers) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataGroupMembers) GetGroupMembers() []BACnetApplicationTagObjectIdentifier {
	return m.GroupMembers
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataGroupMembers) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataGroupMembers factory function for _BACnetConstructedDataGroupMembers
func NewBACnetConstructedDataGroupMembers( numberOfDataElements BACnetApplicationTagUnsignedInteger , groupMembers []BACnetApplicationTagObjectIdentifier , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataGroupMembers {
	_result := &_BACnetConstructedDataGroupMembers{
		NumberOfDataElements: numberOfDataElements,
		GroupMembers: groupMembers,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataGroupMembers(structType any) BACnetConstructedDataGroupMembers {
    if casted, ok := structType.(BACnetConstructedDataGroupMembers); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGroupMembers); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataGroupMembers) GetTypeName() string {
	return "BACnetConstructedDataGroupMembers"
}

func (m *_BACnetConstructedDataGroupMembers) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.GroupMembers) > 0 {
		for _, element := range m.GroupMembers {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}


func (m *_BACnetConstructedDataGroupMembers) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataGroupMembersParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataGroupMembers, error) {
	return BACnetConstructedDataGroupMembersParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataGroupMembersParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataGroupMembers, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGroupMembers"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataGroupMembers")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
_val, _err := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field of BACnetConstructedDataGroupMembers")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (groupMembers)
	if pullErr := readBuffer.PullContext("groupMembers", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for groupMembers")
	}
	// Terminated array
	var groupMembers []BACnetApplicationTagObjectIdentifier
	{
		for ;!bool(IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber)); {
_item, _err := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'groupMembers' field of BACnetConstructedDataGroupMembers")
			}
			groupMembers = append(groupMembers, _item.(BACnetApplicationTagObjectIdentifier))
		}
	}
	if closeErr := readBuffer.CloseContext("groupMembers", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for groupMembers")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGroupMembers"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataGroupMembers")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataGroupMembers{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NumberOfDataElements: numberOfDataElements,
		GroupMembers: groupMembers,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataGroupMembers) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataGroupMembers) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGroupMembers"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataGroupMembers")
		}
	// Virtual field
	zero := m.GetZero()
	_ =	zero
	if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
		return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
	}

	// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if m.GetNumberOfDataElements() != nil {
		if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
		}
		numberOfDataElements = m.GetNumberOfDataElements()
		_numberOfDataElementsErr := writeBuffer.WriteSerializable(ctx, numberOfDataElements)
		if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for numberOfDataElements")
		}
		if _numberOfDataElementsErr != nil {
			return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
		}
	}

	// Array Field (groupMembers)
	if pushErr := writeBuffer.PushContext("groupMembers", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for groupMembers")
	}
	for _curItem, _element := range m.GetGroupMembers() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetGroupMembers()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'groupMembers' field")
		}
	}
	if popErr := writeBuffer.PopContext("groupMembers", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for groupMembers")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGroupMembers"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataGroupMembers")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataGroupMembers) isBACnetConstructedDataGroupMembers() bool {
	return true
}

func (m *_BACnetConstructedDataGroupMembers) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



