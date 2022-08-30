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


// BACnetConstructedDataUpdateTime is the corresponding interface of BACnetConstructedDataUpdateTime
type BACnetConstructedDataUpdateTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetUpdateTime returns UpdateTime (property field)
	GetUpdateTime() BACnetDateTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDateTime
}

// BACnetConstructedDataUpdateTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataUpdateTime.
// This is useful for switch cases.
type BACnetConstructedDataUpdateTimeExactly interface {
	BACnetConstructedDataUpdateTime
	isBACnetConstructedDataUpdateTime() bool
}

// _BACnetConstructedDataUpdateTime is the data-structure of this message
type _BACnetConstructedDataUpdateTime struct {
	*_BACnetConstructedData
        UpdateTime BACnetDateTime
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataUpdateTime)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataUpdateTime)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_UPDATE_TIME}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataUpdateTime) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataUpdateTime)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataUpdateTime) GetUpdateTime() BACnetDateTime {
	return m.UpdateTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataUpdateTime) GetActualValue() BACnetDateTime {
	return CastBACnetDateTime(m.GetUpdateTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataUpdateTime factory function for _BACnetConstructedDataUpdateTime
func NewBACnetConstructedDataUpdateTime( updateTime BACnetDateTime , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataUpdateTime {
	_result := &_BACnetConstructedDataUpdateTime{
		UpdateTime: updateTime,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataUpdateTime(structType interface{}) BACnetConstructedDataUpdateTime {
    if casted, ok := structType.(BACnetConstructedDataUpdateTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataUpdateTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataUpdateTime) GetTypeName() string {
	return "BACnetConstructedDataUpdateTime"
}

func (m *_BACnetConstructedDataUpdateTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataUpdateTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (updateTime)
	lengthInBits += m.UpdateTime.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataUpdateTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataUpdateTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataUpdateTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataUpdateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataUpdateTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (updateTime)
	if pullErr := readBuffer.PullContext("updateTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for updateTime")
	}
_updateTime, _updateTimeErr := BACnetDateTimeParse(readBuffer)
	if _updateTimeErr != nil {
		return nil, errors.Wrap(_updateTimeErr, "Error parsing 'updateTime' field of BACnetConstructedDataUpdateTime")
	}
	updateTime := _updateTime.(BACnetDateTime)
	if closeErr := readBuffer.CloseContext("updateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for updateTime")
	}

	// Virtual field
	_actualValue := updateTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataUpdateTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataUpdateTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataUpdateTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		UpdateTime: updateTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataUpdateTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataUpdateTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataUpdateTime")
		}

	// Simple Field (updateTime)
	if pushErr := writeBuffer.PushContext("updateTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for updateTime")
	}
	_updateTimeErr := writeBuffer.WriteSerializable(m.GetUpdateTime())
	if popErr := writeBuffer.PopContext("updateTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for updateTime")
	}
	if _updateTimeErr != nil {
		return errors.Wrap(_updateTimeErr, "Error serializing 'updateTime' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataUpdateTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataUpdateTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataUpdateTime) isBACnetConstructedDataUpdateTime() bool {
	return true
}

func (m *_BACnetConstructedDataUpdateTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



