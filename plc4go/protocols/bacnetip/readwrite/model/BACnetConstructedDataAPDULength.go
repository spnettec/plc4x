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


// BACnetConstructedDataAPDULength is the corresponding interface of BACnetConstructedDataAPDULength
type BACnetConstructedDataAPDULength interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetApduLength returns ApduLength (property field)
	GetApduLength() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataAPDULengthExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAPDULength.
// This is useful for switch cases.
type BACnetConstructedDataAPDULengthExactly interface {
	BACnetConstructedDataAPDULength
	isBACnetConstructedDataAPDULength() bool
}

// _BACnetConstructedDataAPDULength is the data-structure of this message
type _BACnetConstructedDataAPDULength struct {
	*_BACnetConstructedData
        ApduLength BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAPDULength)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataAPDULength)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_APDU_LENGTH}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAPDULength) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAPDULength)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAPDULength) GetApduLength() BACnetApplicationTagUnsignedInteger {
	return m.ApduLength
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAPDULength) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetApduLength())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataAPDULength factory function for _BACnetConstructedDataAPDULength
func NewBACnetConstructedDataAPDULength( apduLength BACnetApplicationTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataAPDULength {
	_result := &_BACnetConstructedDataAPDULength{
		ApduLength: apduLength,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAPDULength(structType interface{}) BACnetConstructedDataAPDULength {
    if casted, ok := structType.(BACnetConstructedDataAPDULength); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAPDULength); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAPDULength) GetTypeName() string {
	return "BACnetConstructedDataAPDULength"
}

func (m *_BACnetConstructedDataAPDULength) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAPDULength) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (apduLength)
	lengthInBits += m.ApduLength.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataAPDULength) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAPDULengthParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAPDULength, error) {
	return BACnetConstructedDataAPDULengthParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAPDULengthParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAPDULength, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAPDULength"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAPDULength")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (apduLength)
	if pullErr := readBuffer.PullContext("apduLength"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for apduLength")
	}
_apduLength, _apduLengthErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _apduLengthErr != nil {
		return nil, errors.Wrap(_apduLengthErr, "Error parsing 'apduLength' field of BACnetConstructedDataAPDULength")
	}
	apduLength := _apduLength.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("apduLength"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for apduLength")
	}

	// Virtual field
	_actualValue := apduLength
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAPDULength"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAPDULength")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAPDULength{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ApduLength: apduLength,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAPDULength) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAPDULength) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAPDULength"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAPDULength")
		}

	// Simple Field (apduLength)
	if pushErr := writeBuffer.PushContext("apduLength"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for apduLength")
	}
	_apduLengthErr := writeBuffer.WriteSerializable(m.GetApduLength())
	if popErr := writeBuffer.PopContext("apduLength"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for apduLength")
	}
	if _apduLengthErr != nil {
		return errors.Wrap(_apduLengthErr, "Error serializing 'apduLength' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAPDULength"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAPDULength")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataAPDULength) isBACnetConstructedDataAPDULength() bool {
	return true
}

func (m *_BACnetConstructedDataAPDULength) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



