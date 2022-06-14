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

// BACnetConstructedDataRestorePreparationTime is the data-structure of this message
type BACnetConstructedDataRestorePreparationTime struct {
	*BACnetConstructedData
	RestorePreparationTime *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataRestorePreparationTime is the corresponding interface of BACnetConstructedDataRestorePreparationTime
type IBACnetConstructedDataRestorePreparationTime interface {
	IBACnetConstructedData
	// GetRestorePreparationTime returns RestorePreparationTime (property field)
	GetRestorePreparationTime() *BACnetApplicationTagUnsignedInteger
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataRestorePreparationTime) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataRestorePreparationTime) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_RESTORE_PREPARATION_TIME
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataRestorePreparationTime) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataRestorePreparationTime) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataRestorePreparationTime) GetRestorePreparationTime() *BACnetApplicationTagUnsignedInteger {
	return m.RestorePreparationTime
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataRestorePreparationTime factory function for BACnetConstructedDataRestorePreparationTime
func NewBACnetConstructedDataRestorePreparationTime(restorePreparationTime *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataRestorePreparationTime {
	_result := &BACnetConstructedDataRestorePreparationTime{
		RestorePreparationTime: restorePreparationTime,
		BACnetConstructedData:  NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataRestorePreparationTime(structType interface{}) *BACnetConstructedDataRestorePreparationTime {
	if casted, ok := structType.(BACnetConstructedDataRestorePreparationTime); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRestorePreparationTime); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataRestorePreparationTime(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataRestorePreparationTime(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataRestorePreparationTime) GetTypeName() string {
	return "BACnetConstructedDataRestorePreparationTime"
}

func (m *BACnetConstructedDataRestorePreparationTime) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataRestorePreparationTime) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (restorePreparationTime)
	lengthInBits += m.RestorePreparationTime.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataRestorePreparationTime) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataRestorePreparationTimeParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataRestorePreparationTime, error) {
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
		return nil, errors.Wrap(_restorePreparationTimeErr, "Error parsing 'restorePreparationTime' field")
	}
	restorePreparationTime := CastBACnetApplicationTagUnsignedInteger(_restorePreparationTime)
	if closeErr := readBuffer.CloseContext("restorePreparationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for restorePreparationTime")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRestorePreparationTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRestorePreparationTime")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataRestorePreparationTime{
		RestorePreparationTime: CastBACnetApplicationTagUnsignedInteger(restorePreparationTime),
		BACnetConstructedData:  &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataRestorePreparationTime) Serialize(writeBuffer utils.WriteBuffer) error {
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
		_restorePreparationTimeErr := writeBuffer.WriteSerializable(m.RestorePreparationTime)
		if popErr := writeBuffer.PopContext("restorePreparationTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for restorePreparationTime")
		}
		if _restorePreparationTimeErr != nil {
			return errors.Wrap(_restorePreparationTimeErr, "Error serializing 'restorePreparationTime' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRestorePreparationTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRestorePreparationTime")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataRestorePreparationTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
