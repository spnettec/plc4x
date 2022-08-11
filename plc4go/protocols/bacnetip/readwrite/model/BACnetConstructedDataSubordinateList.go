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

// BACnetConstructedDataSubordinateList is the corresponding interface of BACnetConstructedDataSubordinateList
type BACnetConstructedDataSubordinateList interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetSubordinateList returns SubordinateList (property field)
	GetSubordinateList() []BACnetDeviceObjectReference
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataSubordinateListExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSubordinateList.
// This is useful for switch cases.
type BACnetConstructedDataSubordinateListExactly interface {
	BACnetConstructedDataSubordinateList
	isBACnetConstructedDataSubordinateList() bool
}

// _BACnetConstructedDataSubordinateList is the data-structure of this message
type _BACnetConstructedDataSubordinateList struct {
	*_BACnetConstructedData
	NumberOfDataElements BACnetApplicationTagUnsignedInteger
	SubordinateList      []BACnetDeviceObjectReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSubordinateList) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSubordinateList) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SUBORDINATE_LIST
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSubordinateList) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSubordinateList) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSubordinateList) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataSubordinateList) GetSubordinateList() []BACnetDeviceObjectReference {
	return m.SubordinateList
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSubordinateList) GetZero() uint64 {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSubordinateList factory function for _BACnetConstructedDataSubordinateList
func NewBACnetConstructedDataSubordinateList(numberOfDataElements BACnetApplicationTagUnsignedInteger, subordinateList []BACnetDeviceObjectReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSubordinateList {
	_result := &_BACnetConstructedDataSubordinateList{
		NumberOfDataElements:   numberOfDataElements,
		SubordinateList:        subordinateList,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSubordinateList(structType interface{}) BACnetConstructedDataSubordinateList {
	if casted, ok := structType.(BACnetConstructedDataSubordinateList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSubordinateList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSubordinateList) GetTypeName() string {
	return "BACnetConstructedDataSubordinateList"
}

func (m *_BACnetConstructedDataSubordinateList) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataSubordinateList) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits()
	}

	// Array field
	if len(m.SubordinateList) > 0 {
		for _, element := range m.SubordinateList {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataSubordinateList) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSubordinateListParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSubordinateList, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSubordinateList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSubordinateList")
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
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field of BACnetConstructedDataSubordinateList")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (subordinateList)
	if pullErr := readBuffer.PullContext("subordinateList", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for subordinateList")
	}
	// Terminated array
	var subordinateList []BACnetDeviceObjectReference
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDeviceObjectReferenceParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'subordinateList' field of BACnetConstructedDataSubordinateList")
			}
			subordinateList = append(subordinateList, _item.(BACnetDeviceObjectReference))

		}
	}
	if closeErr := readBuffer.CloseContext("subordinateList", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for subordinateList")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSubordinateList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSubordinateList")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSubordinateList{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NumberOfDataElements: numberOfDataElements,
		SubordinateList:      subordinateList,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSubordinateList) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSubordinateList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSubordinateList")
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

		// Array Field (subordinateList)
		if pushErr := writeBuffer.PushContext("subordinateList", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for subordinateList")
		}
		for _, _element := range m.GetSubordinateList() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'subordinateList' field")
			}
		}
		if popErr := writeBuffer.PopContext("subordinateList", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for subordinateList")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSubordinateList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSubordinateList")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSubordinateList) isBACnetConstructedDataSubordinateList() bool {
	return true
}

func (m *_BACnetConstructedDataSubordinateList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
