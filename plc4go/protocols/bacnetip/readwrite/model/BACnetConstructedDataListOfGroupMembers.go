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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataListOfGroupMembers is the corresponding interface of BACnetConstructedDataListOfGroupMembers
type BACnetConstructedDataListOfGroupMembers interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetListOfGroupMembers returns ListOfGroupMembers (property field)
	GetListOfGroupMembers() []BACnetReadAccessSpecification
}

// BACnetConstructedDataListOfGroupMembersExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataListOfGroupMembers.
// This is useful for switch cases.
type BACnetConstructedDataListOfGroupMembersExactly interface {
	BACnetConstructedDataListOfGroupMembers
	isBACnetConstructedDataListOfGroupMembers() bool
}

// _BACnetConstructedDataListOfGroupMembers is the data-structure of this message
type _BACnetConstructedDataListOfGroupMembers struct {
	*_BACnetConstructedData
	ListOfGroupMembers []BACnetReadAccessSpecification
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataListOfGroupMembers) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataListOfGroupMembers) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LIST_OF_GROUP_MEMBERS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataListOfGroupMembers) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataListOfGroupMembers) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataListOfGroupMembers) GetListOfGroupMembers() []BACnetReadAccessSpecification {
	return m.ListOfGroupMembers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataListOfGroupMembers factory function for _BACnetConstructedDataListOfGroupMembers
func NewBACnetConstructedDataListOfGroupMembers(listOfGroupMembers []BACnetReadAccessSpecification, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataListOfGroupMembers {
	_result := &_BACnetConstructedDataListOfGroupMembers{
		ListOfGroupMembers:     listOfGroupMembers,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataListOfGroupMembers(structType interface{}) BACnetConstructedDataListOfGroupMembers {
	if casted, ok := structType.(BACnetConstructedDataListOfGroupMembers); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataListOfGroupMembers); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataListOfGroupMembers) GetTypeName() string {
	return "BACnetConstructedDataListOfGroupMembers"
}

func (m *_BACnetConstructedDataListOfGroupMembers) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataListOfGroupMembers) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.ListOfGroupMembers) > 0 {
		for _, element := range m.ListOfGroupMembers {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataListOfGroupMembers) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataListOfGroupMembersParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataListOfGroupMembers, error) {
	return BACnetConstructedDataListOfGroupMembersParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataListOfGroupMembersParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataListOfGroupMembers, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataListOfGroupMembers"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataListOfGroupMembers")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (listOfGroupMembers)
	if pullErr := readBuffer.PullContext("listOfGroupMembers", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfGroupMembers")
	}
	// Terminated array
	var listOfGroupMembers []BACnetReadAccessSpecification
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetReadAccessSpecificationParseWithBuffer(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'listOfGroupMembers' field of BACnetConstructedDataListOfGroupMembers")
			}
			listOfGroupMembers = append(listOfGroupMembers, _item.(BACnetReadAccessSpecification))
		}
	}
	if closeErr := readBuffer.CloseContext("listOfGroupMembers", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfGroupMembers")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataListOfGroupMembers"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataListOfGroupMembers")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataListOfGroupMembers{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ListOfGroupMembers: listOfGroupMembers,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataListOfGroupMembers) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataListOfGroupMembers) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataListOfGroupMembers"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataListOfGroupMembers")
		}

		// Array Field (listOfGroupMembers)
		if pushErr := writeBuffer.PushContext("listOfGroupMembers", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for listOfGroupMembers")
		}
		for _, _element := range m.GetListOfGroupMembers() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'listOfGroupMembers' field")
			}
		}
		if popErr := writeBuffer.PopContext("listOfGroupMembers", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for listOfGroupMembers")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataListOfGroupMembers"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataListOfGroupMembers")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataListOfGroupMembers) isBACnetConstructedDataListOfGroupMembers() bool {
	return true
}

func (m *_BACnetConstructedDataListOfGroupMembers) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
