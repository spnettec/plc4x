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


// BACnetConstructedDataZoneFrom is the corresponding interface of BACnetConstructedDataZoneFrom
type BACnetConstructedDataZoneFrom interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetZoneFrom returns ZoneFrom (property field)
	GetZoneFrom() BACnetDeviceObjectReference
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceObjectReference
}

// BACnetConstructedDataZoneFromExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataZoneFrom.
// This is useful for switch cases.
type BACnetConstructedDataZoneFromExactly interface {
	BACnetConstructedDataZoneFrom
	isBACnetConstructedDataZoneFrom() bool
}

// _BACnetConstructedDataZoneFrom is the data-structure of this message
type _BACnetConstructedDataZoneFrom struct {
	*_BACnetConstructedData
        ZoneFrom BACnetDeviceObjectReference
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataZoneFrom)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataZoneFrom)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_ZONE_FROM}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataZoneFrom) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataZoneFrom)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataZoneFrom) GetZoneFrom() BACnetDeviceObjectReference {
	return m.ZoneFrom
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataZoneFrom) GetActualValue() BACnetDeviceObjectReference {
	return CastBACnetDeviceObjectReference(m.GetZoneFrom())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataZoneFrom factory function for _BACnetConstructedDataZoneFrom
func NewBACnetConstructedDataZoneFrom( zoneFrom BACnetDeviceObjectReference , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataZoneFrom {
	_result := &_BACnetConstructedDataZoneFrom{
		ZoneFrom: zoneFrom,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataZoneFrom(structType interface{}) BACnetConstructedDataZoneFrom {
    if casted, ok := structType.(BACnetConstructedDataZoneFrom); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataZoneFrom); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataZoneFrom) GetTypeName() string {
	return "BACnetConstructedDataZoneFrom"
}

func (m *_BACnetConstructedDataZoneFrom) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataZoneFrom) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (zoneFrom)
	lengthInBits += m.ZoneFrom.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataZoneFrom) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataZoneFromParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataZoneFrom, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataZoneFrom"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataZoneFrom")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneFrom)
	if pullErr := readBuffer.PullContext("zoneFrom"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneFrom")
	}
_zoneFrom, _zoneFromErr := BACnetDeviceObjectReferenceParse(readBuffer)
	if _zoneFromErr != nil {
		return nil, errors.Wrap(_zoneFromErr, "Error parsing 'zoneFrom' field of BACnetConstructedDataZoneFrom")
	}
	zoneFrom := _zoneFrom.(BACnetDeviceObjectReference)
	if closeErr := readBuffer.CloseContext("zoneFrom"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneFrom")
	}

	// Virtual field
	_actualValue := zoneFrom
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataZoneFrom"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataZoneFrom")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataZoneFrom{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ZoneFrom: zoneFrom,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataZoneFrom) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataZoneFrom"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataZoneFrom")
		}

	// Simple Field (zoneFrom)
	if pushErr := writeBuffer.PushContext("zoneFrom"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for zoneFrom")
	}
	_zoneFromErr := writeBuffer.WriteSerializable(m.GetZoneFrom())
	if popErr := writeBuffer.PopContext("zoneFrom"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for zoneFrom")
	}
	if _zoneFromErr != nil {
		return errors.Wrap(_zoneFromErr, "Error serializing 'zoneFrom' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataZoneFrom"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataZoneFrom")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataZoneFrom) isBACnetConstructedDataZoneFrom() bool {
	return true
}

func (m *_BACnetConstructedDataZoneFrom) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



