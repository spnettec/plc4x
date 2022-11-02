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


// BACnetConstructedDataLowDiffLimit is the corresponding interface of BACnetConstructedDataLowDiffLimit
type BACnetConstructedDataLowDiffLimit interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLowDiffLimit returns LowDiffLimit (property field)
	GetLowDiffLimit() BACnetOptionalREAL
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetOptionalREAL
}

// BACnetConstructedDataLowDiffLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLowDiffLimit.
// This is useful for switch cases.
type BACnetConstructedDataLowDiffLimitExactly interface {
	BACnetConstructedDataLowDiffLimit
	isBACnetConstructedDataLowDiffLimit() bool
}

// _BACnetConstructedDataLowDiffLimit is the data-structure of this message
type _BACnetConstructedDataLowDiffLimit struct {
	*_BACnetConstructedData
        LowDiffLimit BACnetOptionalREAL
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLowDiffLimit)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataLowDiffLimit)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_LOW_DIFF_LIMIT}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLowDiffLimit) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLowDiffLimit)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLowDiffLimit) GetLowDiffLimit() BACnetOptionalREAL {
	return m.LowDiffLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLowDiffLimit) GetActualValue() BACnetOptionalREAL {
	return CastBACnetOptionalREAL(m.GetLowDiffLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataLowDiffLimit factory function for _BACnetConstructedDataLowDiffLimit
func NewBACnetConstructedDataLowDiffLimit( lowDiffLimit BACnetOptionalREAL , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataLowDiffLimit {
	_result := &_BACnetConstructedDataLowDiffLimit{
		LowDiffLimit: lowDiffLimit,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLowDiffLimit(structType interface{}) BACnetConstructedDataLowDiffLimit {
    if casted, ok := structType.(BACnetConstructedDataLowDiffLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLowDiffLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLowDiffLimit) GetTypeName() string {
	return "BACnetConstructedDataLowDiffLimit"
}

func (m *_BACnetConstructedDataLowDiffLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLowDiffLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lowDiffLimit)
	lengthInBits += m.LowDiffLimit.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataLowDiffLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLowDiffLimitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLowDiffLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLowDiffLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLowDiffLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lowDiffLimit)
	if pullErr := readBuffer.PullContext("lowDiffLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lowDiffLimit")
	}
_lowDiffLimit, _lowDiffLimitErr := BACnetOptionalREALParse(readBuffer)
	if _lowDiffLimitErr != nil {
		return nil, errors.Wrap(_lowDiffLimitErr, "Error parsing 'lowDiffLimit' field of BACnetConstructedDataLowDiffLimit")
	}
	lowDiffLimit := _lowDiffLimit.(BACnetOptionalREAL)
	if closeErr := readBuffer.CloseContext("lowDiffLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lowDiffLimit")
	}

	// Virtual field
	_actualValue := lowDiffLimit
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLowDiffLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLowDiffLimit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLowDiffLimit{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LowDiffLimit: lowDiffLimit,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLowDiffLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLowDiffLimit) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLowDiffLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLowDiffLimit")
		}

	// Simple Field (lowDiffLimit)
	if pushErr := writeBuffer.PushContext("lowDiffLimit"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for lowDiffLimit")
	}
	_lowDiffLimitErr := writeBuffer.WriteSerializable(m.GetLowDiffLimit())
	if popErr := writeBuffer.PopContext("lowDiffLimit"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for lowDiffLimit")
	}
	if _lowDiffLimitErr != nil {
		return errors.Wrap(_lowDiffLimitErr, "Error serializing 'lowDiffLimit' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLowDiffLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLowDiffLimit")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataLowDiffLimit) isBACnetConstructedDataLowDiffLimit() bool {
	return true
}

func (m *_BACnetConstructedDataLowDiffLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



