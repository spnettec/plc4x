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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataMemberOf is the corresponding interface of BACnetConstructedDataMemberOf
type BACnetConstructedDataMemberOf interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetZones returns Zones (property field)
	GetZones() []BACnetDeviceObjectReference
}

// BACnetConstructedDataMemberOfExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMemberOf.
// This is useful for switch cases.
type BACnetConstructedDataMemberOfExactly interface {
	BACnetConstructedDataMemberOf
	isBACnetConstructedDataMemberOf() bool
}

// _BACnetConstructedDataMemberOf is the data-structure of this message
type _BACnetConstructedDataMemberOf struct {
	*_BACnetConstructedData
	Zones []BACnetDeviceObjectReference
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMemberOf) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMemberOf) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MEMBER_OF
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMemberOf) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMemberOf) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMemberOf) GetZones() []BACnetDeviceObjectReference {
	return m.Zones
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMemberOf factory function for _BACnetConstructedDataMemberOf
func NewBACnetConstructedDataMemberOf(zones []BACnetDeviceObjectReference, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMemberOf {
	_result := &_BACnetConstructedDataMemberOf{
		Zones:                  zones,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMemberOf(structType interface{}) BACnetConstructedDataMemberOf {
	if casted, ok := structType.(BACnetConstructedDataMemberOf); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMemberOf); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMemberOf) GetTypeName() string {
	return "BACnetConstructedDataMemberOf"
}

func (m *_BACnetConstructedDataMemberOf) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataMemberOf) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.Zones) > 0 {
		for _, element := range m.Zones {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataMemberOf) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMemberOfParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMemberOf, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMemberOf"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMemberOf")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (zones)
	if pullErr := readBuffer.PullContext("zones", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zones")
	}
	// Terminated array
	var zones []BACnetDeviceObjectReference
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetDeviceObjectReferenceParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'zones' field of BACnetConstructedDataMemberOf")
			}
			zones = append(zones, _item.(BACnetDeviceObjectReference))

		}
	}
	if closeErr := readBuffer.CloseContext("zones", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zones")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMemberOf"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMemberOf")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMemberOf{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Zones: zones,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMemberOf) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMemberOf"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMemberOf")
		}

		// Array Field (zones)
		if pushErr := writeBuffer.PushContext("zones", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for zones")
		}
		for _, _element := range m.GetZones() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'zones' field")
			}
		}
		if popErr := writeBuffer.PopContext("zones", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for zones")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMemberOf"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMemberOf")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMemberOf) isBACnetConstructedDataMemberOf() bool {
	return true
}

func (m *_BACnetConstructedDataMemberOf) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
