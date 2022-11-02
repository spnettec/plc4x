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


// BACnetConstructedDataMaxSegmentsAccepted is the corresponding interface of BACnetConstructedDataMaxSegmentsAccepted
type BACnetConstructedDataMaxSegmentsAccepted interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaxSegmentsAccepted returns MaxSegmentsAccepted (property field)
	GetMaxSegmentsAccepted() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataMaxSegmentsAcceptedExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMaxSegmentsAccepted.
// This is useful for switch cases.
type BACnetConstructedDataMaxSegmentsAcceptedExactly interface {
	BACnetConstructedDataMaxSegmentsAccepted
	isBACnetConstructedDataMaxSegmentsAccepted() bool
}

// _BACnetConstructedDataMaxSegmentsAccepted is the data-structure of this message
type _BACnetConstructedDataMaxSegmentsAccepted struct {
	*_BACnetConstructedData
        MaxSegmentsAccepted BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaxSegmentsAccepted)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataMaxSegmentsAccepted)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_MAX_SEGMENTS_ACCEPTED}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaxSegmentsAccepted) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaxSegmentsAccepted) GetMaxSegmentsAccepted() BACnetApplicationTagUnsignedInteger {
	return m.MaxSegmentsAccepted
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMaxSegmentsAccepted) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetMaxSegmentsAccepted())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataMaxSegmentsAccepted factory function for _BACnetConstructedDataMaxSegmentsAccepted
func NewBACnetConstructedDataMaxSegmentsAccepted( maxSegmentsAccepted BACnetApplicationTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataMaxSegmentsAccepted {
	_result := &_BACnetConstructedDataMaxSegmentsAccepted{
		MaxSegmentsAccepted: maxSegmentsAccepted,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaxSegmentsAccepted(structType interface{}) BACnetConstructedDataMaxSegmentsAccepted {
    if casted, ok := structType.(BACnetConstructedDataMaxSegmentsAccepted); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaxSegmentsAccepted); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) GetTypeName() string {
	return "BACnetConstructedDataMaxSegmentsAccepted"
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maxSegmentsAccepted)
	lengthInBits += m.MaxSegmentsAccepted.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataMaxSegmentsAccepted) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMaxSegmentsAcceptedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMaxSegmentsAccepted, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaxSegmentsAccepted"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaxSegmentsAccepted")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maxSegmentsAccepted)
	if pullErr := readBuffer.PullContext("maxSegmentsAccepted"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxSegmentsAccepted")
	}
_maxSegmentsAccepted, _maxSegmentsAcceptedErr := BACnetApplicationTagParse(readBuffer)
	if _maxSegmentsAcceptedErr != nil {
		return nil, errors.Wrap(_maxSegmentsAcceptedErr, "Error parsing 'maxSegmentsAccepted' field of BACnetConstructedDataMaxSegmentsAccepted")
	}
	maxSegmentsAccepted := _maxSegmentsAccepted.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("maxSegmentsAccepted"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxSegmentsAccepted")
	}

	// Virtual field
	_actualValue := maxSegmentsAccepted
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaxSegmentsAccepted"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaxSegmentsAccepted")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMaxSegmentsAccepted{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MaxSegmentsAccepted: maxSegmentsAccepted,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaxSegmentsAccepted"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaxSegmentsAccepted")
		}

	// Simple Field (maxSegmentsAccepted)
	if pushErr := writeBuffer.PushContext("maxSegmentsAccepted"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for maxSegmentsAccepted")
	}
	_maxSegmentsAcceptedErr := writeBuffer.WriteSerializable(m.GetMaxSegmentsAccepted())
	if popErr := writeBuffer.PopContext("maxSegmentsAccepted"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for maxSegmentsAccepted")
	}
	if _maxSegmentsAcceptedErr != nil {
		return errors.Wrap(_maxSegmentsAcceptedErr, "Error serializing 'maxSegmentsAccepted' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaxSegmentsAccepted"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaxSegmentsAccepted")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataMaxSegmentsAccepted) isBACnetConstructedDataMaxSegmentsAccepted() bool {
	return true
}

func (m *_BACnetConstructedDataMaxSegmentsAccepted) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



