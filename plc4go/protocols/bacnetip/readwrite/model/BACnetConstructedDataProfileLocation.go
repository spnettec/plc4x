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


// BACnetConstructedDataProfileLocation is the corresponding interface of BACnetConstructedDataProfileLocation
type BACnetConstructedDataProfileLocation interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetProfileLocation returns ProfileLocation (property field)
	GetProfileLocation() BACnetApplicationTagCharacterString
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagCharacterString
}

// BACnetConstructedDataProfileLocationExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataProfileLocation.
// This is useful for switch cases.
type BACnetConstructedDataProfileLocationExactly interface {
	BACnetConstructedDataProfileLocation
	isBACnetConstructedDataProfileLocation() bool
}

// _BACnetConstructedDataProfileLocation is the data-structure of this message
type _BACnetConstructedDataProfileLocation struct {
	*_BACnetConstructedData
        ProfileLocation BACnetApplicationTagCharacterString
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProfileLocation)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataProfileLocation)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_PROFILE_LOCATION}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProfileLocation) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataProfileLocation)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProfileLocation) GetProfileLocation() BACnetApplicationTagCharacterString {
	return m.ProfileLocation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProfileLocation) GetActualValue() BACnetApplicationTagCharacterString {
	return CastBACnetApplicationTagCharacterString(m.GetProfileLocation())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataProfileLocation factory function for _BACnetConstructedDataProfileLocation
func NewBACnetConstructedDataProfileLocation( profileLocation BACnetApplicationTagCharacterString , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataProfileLocation {
	_result := &_BACnetConstructedDataProfileLocation{
		ProfileLocation: profileLocation,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProfileLocation(structType interface{}) BACnetConstructedDataProfileLocation {
    if casted, ok := structType.(BACnetConstructedDataProfileLocation); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProfileLocation); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProfileLocation) GetTypeName() string {
	return "BACnetConstructedDataProfileLocation"
}

func (m *_BACnetConstructedDataProfileLocation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataProfileLocation) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (profileLocation)
	lengthInBits += m.ProfileLocation.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataProfileLocation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataProfileLocationParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProfileLocation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProfileLocation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProfileLocation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (profileLocation)
	if pullErr := readBuffer.PullContext("profileLocation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for profileLocation")
	}
_profileLocation, _profileLocationErr := BACnetApplicationTagParse(readBuffer)
	if _profileLocationErr != nil {
		return nil, errors.Wrap(_profileLocationErr, "Error parsing 'profileLocation' field of BACnetConstructedDataProfileLocation")
	}
	profileLocation := _profileLocation.(BACnetApplicationTagCharacterString)
	if closeErr := readBuffer.CloseContext("profileLocation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for profileLocation")
	}

	// Virtual field
	_actualValue := profileLocation
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProfileLocation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProfileLocation")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataProfileLocation{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ProfileLocation: profileLocation,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataProfileLocation) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProfileLocation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProfileLocation")
		}

	// Simple Field (profileLocation)
	if pushErr := writeBuffer.PushContext("profileLocation"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for profileLocation")
	}
	_profileLocationErr := writeBuffer.WriteSerializable(m.GetProfileLocation())
	if popErr := writeBuffer.PopContext("profileLocation"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for profileLocation")
	}
	if _profileLocationErr != nil {
		return errors.Wrap(_profileLocationErr, "Error serializing 'profileLocation' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProfileLocation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProfileLocation")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataProfileLocation) isBACnetConstructedDataProfileLocation() bool {
	return true
}

func (m *_BACnetConstructedDataProfileLocation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



