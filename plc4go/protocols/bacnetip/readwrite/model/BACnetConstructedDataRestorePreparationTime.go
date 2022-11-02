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


// BACnetConstructedDataRestorePreparationTime is the corresponding interface of BACnetConstructedDataRestorePreparationTime
type BACnetConstructedDataRestorePreparationTime interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetRestorePreparationTime returns RestorePreparationTime (property field)
	GetRestorePreparationTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataRestorePreparationTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataRestorePreparationTime.
// This is useful for switch cases.
type BACnetConstructedDataRestorePreparationTimeExactly interface {
	BACnetConstructedDataRestorePreparationTime
	isBACnetConstructedDataRestorePreparationTime() bool
}

// _BACnetConstructedDataRestorePreparationTime is the data-structure of this message
type _BACnetConstructedDataRestorePreparationTime struct {
	*_BACnetConstructedData
        RestorePreparationTime BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRestorePreparationTime)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataRestorePreparationTime)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_RESTORE_PREPARATION_TIME}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRestorePreparationTime) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataRestorePreparationTime)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRestorePreparationTime) GetRestorePreparationTime() BACnetApplicationTagUnsignedInteger {
	return m.RestorePreparationTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataRestorePreparationTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetRestorePreparationTime())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataRestorePreparationTime factory function for _BACnetConstructedDataRestorePreparationTime
func NewBACnetConstructedDataRestorePreparationTime( restorePreparationTime BACnetApplicationTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataRestorePreparationTime {
	_result := &_BACnetConstructedDataRestorePreparationTime{
		RestorePreparationTime: restorePreparationTime,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRestorePreparationTime(structType interface{}) BACnetConstructedDataRestorePreparationTime {
    if casted, ok := structType.(BACnetConstructedDataRestorePreparationTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRestorePreparationTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRestorePreparationTime) GetTypeName() string {
	return "BACnetConstructedDataRestorePreparationTime"
}

func (m *_BACnetConstructedDataRestorePreparationTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataRestorePreparationTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (restorePreparationTime)
	lengthInBits += m.RestorePreparationTime.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataRestorePreparationTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataRestorePreparationTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataRestorePreparationTime, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRestorePreparationTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRestorePreparationTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (restorePreparationTime)
	if pullErr := readBuffer.PullContext("restorePreparationTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for restorePreparationTime")
	}
_restorePreparationTime, _restorePreparationTimeErr := BACnetApplicationTagParse(readBuffer)
	if _restorePreparationTimeErr != nil {
		return nil, errors.Wrap(_restorePreparationTimeErr, "Error parsing 'restorePreparationTime' field of BACnetConstructedDataRestorePreparationTime")
	}
	restorePreparationTime := _restorePreparationTime.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("restorePreparationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for restorePreparationTime")
	}

	// Virtual field
	_actualValue := restorePreparationTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRestorePreparationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRestorePreparationTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataRestorePreparationTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		RestorePreparationTime: restorePreparationTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataRestorePreparationTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataRestorePreparationTime) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRestorePreparationTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRestorePreparationTime")
		}

	// Simple Field (restorePreparationTime)
	if pushErr := writeBuffer.PushContext("restorePreparationTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for restorePreparationTime")
	}
	_restorePreparationTimeErr := writeBuffer.WriteSerializable(m.GetRestorePreparationTime())
	if popErr := writeBuffer.PopContext("restorePreparationTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for restorePreparationTime")
	}
	if _restorePreparationTimeErr != nil {
		return errors.Wrap(_restorePreparationTimeErr, "Error serializing 'restorePreparationTime' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRestorePreparationTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRestorePreparationTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataRestorePreparationTime) isBACnetConstructedDataRestorePreparationTime() bool {
	return true
}

func (m *_BACnetConstructedDataRestorePreparationTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



