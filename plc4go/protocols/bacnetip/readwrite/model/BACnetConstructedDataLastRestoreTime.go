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


// BACnetConstructedDataLastRestoreTime is the corresponding interface of BACnetConstructedDataLastRestoreTime
type BACnetConstructedDataLastRestoreTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLastRestoreTime returns LastRestoreTime (property field)
	GetLastRestoreTime() BACnetTimeStamp
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetTimeStamp
}

// BACnetConstructedDataLastRestoreTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLastRestoreTime.
// This is useful for switch cases.
type BACnetConstructedDataLastRestoreTimeExactly interface {
	BACnetConstructedDataLastRestoreTime
	isBACnetConstructedDataLastRestoreTime() bool
}

// _BACnetConstructedDataLastRestoreTime is the data-structure of this message
type _BACnetConstructedDataLastRestoreTime struct {
	*_BACnetConstructedData
        LastRestoreTime BACnetTimeStamp
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLastRestoreTime)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataLastRestoreTime)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_LAST_RESTORE_TIME}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLastRestoreTime) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLastRestoreTime)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLastRestoreTime) GetLastRestoreTime() BACnetTimeStamp {
	return m.LastRestoreTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLastRestoreTime) GetActualValue() BACnetTimeStamp {
	return CastBACnetTimeStamp(m.GetLastRestoreTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataLastRestoreTime factory function for _BACnetConstructedDataLastRestoreTime
func NewBACnetConstructedDataLastRestoreTime( lastRestoreTime BACnetTimeStamp , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataLastRestoreTime {
	_result := &_BACnetConstructedDataLastRestoreTime{
		LastRestoreTime: lastRestoreTime,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLastRestoreTime(structType interface{}) BACnetConstructedDataLastRestoreTime {
    if casted, ok := structType.(BACnetConstructedDataLastRestoreTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLastRestoreTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLastRestoreTime) GetTypeName() string {
	return "BACnetConstructedDataLastRestoreTime"
}

func (m *_BACnetConstructedDataLastRestoreTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLastRestoreTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lastRestoreTime)
	lengthInBits += m.LastRestoreTime.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataLastRestoreTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLastRestoreTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLastRestoreTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLastRestoreTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLastRestoreTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lastRestoreTime)
	if pullErr := readBuffer.PullContext("lastRestoreTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lastRestoreTime")
	}
_lastRestoreTime, _lastRestoreTimeErr := BACnetTimeStampParse(readBuffer)
	if _lastRestoreTimeErr != nil {
		return nil, errors.Wrap(_lastRestoreTimeErr, "Error parsing 'lastRestoreTime' field of BACnetConstructedDataLastRestoreTime")
	}
	lastRestoreTime := _lastRestoreTime.(BACnetTimeStamp)
	if closeErr := readBuffer.CloseContext("lastRestoreTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lastRestoreTime")
	}

	// Virtual field
	_actualValue := lastRestoreTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLastRestoreTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLastRestoreTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLastRestoreTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LastRestoreTime: lastRestoreTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLastRestoreTime) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLastRestoreTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLastRestoreTime")
		}

	// Simple Field (lastRestoreTime)
	if pushErr := writeBuffer.PushContext("lastRestoreTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for lastRestoreTime")
	}
	_lastRestoreTimeErr := writeBuffer.WriteSerializable(m.GetLastRestoreTime())
	if popErr := writeBuffer.PopContext("lastRestoreTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for lastRestoreTime")
	}
	if _lastRestoreTimeErr != nil {
		return errors.Wrap(_lastRestoreTimeErr, "Error serializing 'lastRestoreTime' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLastRestoreTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLastRestoreTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataLastRestoreTime) isBACnetConstructedDataLastRestoreTime() bool {
	return true
}

func (m *_BACnetConstructedDataLastRestoreTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



