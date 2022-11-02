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


// BACnetConstructedDataAlignIntervals is the corresponding interface of BACnetConstructedDataAlignIntervals
type BACnetConstructedDataAlignIntervals interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAlignIntervals returns AlignIntervals (property field)
	GetAlignIntervals() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataAlignIntervalsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAlignIntervals.
// This is useful for switch cases.
type BACnetConstructedDataAlignIntervalsExactly interface {
	BACnetConstructedDataAlignIntervals
	isBACnetConstructedDataAlignIntervals() bool
}

// _BACnetConstructedDataAlignIntervals is the data-structure of this message
type _BACnetConstructedDataAlignIntervals struct {
	*_BACnetConstructedData
        AlignIntervals BACnetApplicationTagBoolean
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAlignIntervals)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataAlignIntervals)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_ALIGN_INTERVALS}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAlignIntervals) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAlignIntervals)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAlignIntervals) GetAlignIntervals() BACnetApplicationTagBoolean {
	return m.AlignIntervals
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAlignIntervals) GetActualValue() BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetAlignIntervals())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataAlignIntervals factory function for _BACnetConstructedDataAlignIntervals
func NewBACnetConstructedDataAlignIntervals( alignIntervals BACnetApplicationTagBoolean , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataAlignIntervals {
	_result := &_BACnetConstructedDataAlignIntervals{
		AlignIntervals: alignIntervals,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAlignIntervals(structType interface{}) BACnetConstructedDataAlignIntervals {
    if casted, ok := structType.(BACnetConstructedDataAlignIntervals); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAlignIntervals); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAlignIntervals) GetTypeName() string {
	return "BACnetConstructedDataAlignIntervals"
}

func (m *_BACnetConstructedDataAlignIntervals) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAlignIntervals) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (alignIntervals)
	lengthInBits += m.AlignIntervals.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataAlignIntervals) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAlignIntervalsParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAlignIntervals, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAlignIntervals"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAlignIntervals")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (alignIntervals)
	if pullErr := readBuffer.PullContext("alignIntervals"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for alignIntervals")
	}
_alignIntervals, _alignIntervalsErr := BACnetApplicationTagParse(readBuffer)
	if _alignIntervalsErr != nil {
		return nil, errors.Wrap(_alignIntervalsErr, "Error parsing 'alignIntervals' field of BACnetConstructedDataAlignIntervals")
	}
	alignIntervals := _alignIntervals.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("alignIntervals"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for alignIntervals")
	}

	// Virtual field
	_actualValue := alignIntervals
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAlignIntervals"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAlignIntervals")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAlignIntervals{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AlignIntervals: alignIntervals,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAlignIntervals) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAlignIntervals) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAlignIntervals"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAlignIntervals")
		}

	// Simple Field (alignIntervals)
	if pushErr := writeBuffer.PushContext("alignIntervals"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for alignIntervals")
	}
	_alignIntervalsErr := writeBuffer.WriteSerializable(m.GetAlignIntervals())
	if popErr := writeBuffer.PopContext("alignIntervals"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for alignIntervals")
	}
	if _alignIntervalsErr != nil {
		return errors.Wrap(_alignIntervalsErr, "Error serializing 'alignIntervals' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAlignIntervals"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAlignIntervals")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataAlignIntervals) isBACnetConstructedDataAlignIntervals() bool {
	return true
}

func (m *_BACnetConstructedDataAlignIntervals) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



