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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetConstructedDataOccupancyUpperLimitEnforced is the corresponding interface of BACnetConstructedDataOccupancyUpperLimitEnforced
type BACnetConstructedDataOccupancyUpperLimitEnforced interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetOccupancyUpperLimitEnforced returns OccupancyUpperLimitEnforced (property field)
	GetOccupancyUpperLimitEnforced() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataOccupancyUpperLimitEnforcedExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataOccupancyUpperLimitEnforced.
// This is useful for switch cases.
type BACnetConstructedDataOccupancyUpperLimitEnforcedExactly interface {
	BACnetConstructedDataOccupancyUpperLimitEnforced
	isBACnetConstructedDataOccupancyUpperLimitEnforced() bool
}

// _BACnetConstructedDataOccupancyUpperLimitEnforced is the data-structure of this message
type _BACnetConstructedDataOccupancyUpperLimitEnforced struct {
	*_BACnetConstructedData
        OccupancyUpperLimitEnforced BACnetApplicationTagBoolean
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_OCCUPANCY_UPPER_LIMIT_ENFORCED}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) GetOccupancyUpperLimitEnforced() BACnetApplicationTagBoolean {
	return m.OccupancyUpperLimitEnforced
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetOccupancyUpperLimitEnforced())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataOccupancyUpperLimitEnforced factory function for _BACnetConstructedDataOccupancyUpperLimitEnforced
func NewBACnetConstructedDataOccupancyUpperLimitEnforced( occupancyUpperLimitEnforced BACnetApplicationTagBoolean , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataOccupancyUpperLimitEnforced {
	_result := &_BACnetConstructedDataOccupancyUpperLimitEnforced{
		OccupancyUpperLimitEnforced: occupancyUpperLimitEnforced,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataOccupancyUpperLimitEnforced(structType interface{}) BACnetConstructedDataOccupancyUpperLimitEnforced {
    if casted, ok := structType.(BACnetConstructedDataOccupancyUpperLimitEnforced); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataOccupancyUpperLimitEnforced); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) GetTypeName() string {
	return "BACnetConstructedDataOccupancyUpperLimitEnforced"
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (occupancyUpperLimitEnforced)
	lengthInBits += m.OccupancyUpperLimitEnforced.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataOccupancyUpperLimitEnforcedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataOccupancyUpperLimitEnforced, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataOccupancyUpperLimitEnforced"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataOccupancyUpperLimitEnforced")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (occupancyUpperLimitEnforced)
	if pullErr := readBuffer.PullContext("occupancyUpperLimitEnforced"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for occupancyUpperLimitEnforced")
	}
_occupancyUpperLimitEnforced, _occupancyUpperLimitEnforcedErr := BACnetApplicationTagParse(readBuffer)
	if _occupancyUpperLimitEnforcedErr != nil {
		return nil, errors.Wrap(_occupancyUpperLimitEnforcedErr, "Error parsing 'occupancyUpperLimitEnforced' field of BACnetConstructedDataOccupancyUpperLimitEnforced")
	}
	occupancyUpperLimitEnforced := _occupancyUpperLimitEnforced.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("occupancyUpperLimitEnforced"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for occupancyUpperLimitEnforced")
	}

	// Virtual field
	_actualValue := occupancyUpperLimitEnforced
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataOccupancyUpperLimitEnforced"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataOccupancyUpperLimitEnforced")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataOccupancyUpperLimitEnforced{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		OccupancyUpperLimitEnforced: occupancyUpperLimitEnforced,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataOccupancyUpperLimitEnforced"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataOccupancyUpperLimitEnforced")
		}

	// Simple Field (occupancyUpperLimitEnforced)
	if pushErr := writeBuffer.PushContext("occupancyUpperLimitEnforced"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for occupancyUpperLimitEnforced")
	}
	_occupancyUpperLimitEnforcedErr := writeBuffer.WriteSerializable(m.GetOccupancyUpperLimitEnforced())
	if popErr := writeBuffer.PopContext("occupancyUpperLimitEnforced"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for occupancyUpperLimitEnforced")
	}
	if _occupancyUpperLimitEnforcedErr != nil {
		return errors.Wrap(_occupancyUpperLimitEnforcedErr, "Error serializing 'occupancyUpperLimitEnforced' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataOccupancyUpperLimitEnforced"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataOccupancyUpperLimitEnforced")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) isBACnetConstructedDataOccupancyUpperLimitEnforced() bool {
	return true
}

func (m *_BACnetConstructedDataOccupancyUpperLimitEnforced) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



