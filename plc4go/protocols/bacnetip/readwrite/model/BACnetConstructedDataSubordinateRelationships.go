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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataSubordinateRelationships is the corresponding interface of BACnetConstructedDataSubordinateRelationships
type BACnetConstructedDataSubordinateRelationships interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetSubordinateRelationships returns SubordinateRelationships (property field)
	GetSubordinateRelationships() []BACnetRelationshipTagged
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataSubordinateRelationshipsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSubordinateRelationships.
// This is useful for switch cases.
type BACnetConstructedDataSubordinateRelationshipsExactly interface {
	BACnetConstructedDataSubordinateRelationships
	isBACnetConstructedDataSubordinateRelationships() bool
}

// _BACnetConstructedDataSubordinateRelationships is the data-structure of this message
type _BACnetConstructedDataSubordinateRelationships struct {
	*_BACnetConstructedData
	NumberOfDataElements     BACnetApplicationTagUnsignedInteger
	SubordinateRelationships []BACnetRelationshipTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSubordinateRelationships) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSubordinateRelationships) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUBORDINATE_RELATIONSHIPS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSubordinateRelationships) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSubordinateRelationships) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSubordinateRelationships) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataSubordinateRelationships) GetSubordinateRelationships() []BACnetRelationshipTagged {
	return m.SubordinateRelationships
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSubordinateRelationships) GetZero() uint64 {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSubordinateRelationships factory function for _BACnetConstructedDataSubordinateRelationships
func NewBACnetConstructedDataSubordinateRelationships(numberOfDataElements BACnetApplicationTagUnsignedInteger, subordinateRelationships []BACnetRelationshipTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSubordinateRelationships {
	_result := &_BACnetConstructedDataSubordinateRelationships{
		NumberOfDataElements:     numberOfDataElements,
		SubordinateRelationships: subordinateRelationships,
		_BACnetConstructedData:   NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSubordinateRelationships(structType interface{}) BACnetConstructedDataSubordinateRelationships {
	if casted, ok := structType.(BACnetConstructedDataSubordinateRelationships); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSubordinateRelationships); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSubordinateRelationships) GetTypeName() string {
	return "BACnetConstructedDataSubordinateRelationships"
}

func (m *_BACnetConstructedDataSubordinateRelationships) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataSubordinateRelationships) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits()
	}

	// Array field
	if len(m.SubordinateRelationships) > 0 {
		for _, element := range m.SubordinateRelationships {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataSubordinateRelationships) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSubordinateRelationshipsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSubordinateRelationships, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSubordinateRelationships"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSubordinateRelationships")
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
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field of BACnetConstructedDataSubordinateRelationships")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (subordinateRelationships)
	if pullErr := readBuffer.PullContext("subordinateRelationships", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for subordinateRelationships")
	}
	// Terminated array
	var subordinateRelationships []BACnetRelationshipTagged
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetRelationshipTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'subordinateRelationships' field of BACnetConstructedDataSubordinateRelationships")
			}
			subordinateRelationships = append(subordinateRelationships, _item.(BACnetRelationshipTagged))

		}
	}
	if closeErr := readBuffer.CloseContext("subordinateRelationships", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for subordinateRelationships")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSubordinateRelationships"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSubordinateRelationships")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSubordinateRelationships{
		NumberOfDataElements:     numberOfDataElements,
		SubordinateRelationships: subordinateRelationships,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSubordinateRelationships) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSubordinateRelationships"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSubordinateRelationships")
		}
		// Virtual field
		if _zeroErr := writeBuffer.WriteVirtual("zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
		if m.GetNumberOfDataElements() != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.GetNumberOfDataElements()
			_numberOfDataElementsErr := writeBuffer.WriteSerializable(numberOfDataElements)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		// Array Field (subordinateRelationships)
		if pushErr := writeBuffer.PushContext("subordinateRelationships", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for subordinateRelationships")
		}
		for _, _element := range m.GetSubordinateRelationships() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'subordinateRelationships' field")
			}
		}
		if popErr := writeBuffer.PopContext("subordinateRelationships", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for subordinateRelationships")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSubordinateRelationships"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSubordinateRelationships")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSubordinateRelationships) isBACnetConstructedDataSubordinateRelationships() bool {
	return true
}

func (m *_BACnetConstructedDataSubordinateRelationships) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
