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


// BACnetConstructedDataGlobalIdentifier is the corresponding interface of BACnetConstructedDataGlobalIdentifier
type BACnetConstructedDataGlobalIdentifier interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetGlobalIdentifier returns GlobalIdentifier (property field)
	GetGlobalIdentifier() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataGlobalIdentifierExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataGlobalIdentifier.
// This is useful for switch cases.
type BACnetConstructedDataGlobalIdentifierExactly interface {
	BACnetConstructedDataGlobalIdentifier
	isBACnetConstructedDataGlobalIdentifier() bool
}

// _BACnetConstructedDataGlobalIdentifier is the data-structure of this message
type _BACnetConstructedDataGlobalIdentifier struct {
	*_BACnetConstructedData
        GlobalIdentifier BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataGlobalIdentifier)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataGlobalIdentifier)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_GLOBAL_IDENTIFIER}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataGlobalIdentifier) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataGlobalIdentifier)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataGlobalIdentifier) GetGlobalIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.GlobalIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataGlobalIdentifier) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetGlobalIdentifier())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataGlobalIdentifier factory function for _BACnetConstructedDataGlobalIdentifier
func NewBACnetConstructedDataGlobalIdentifier( globalIdentifier BACnetApplicationTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataGlobalIdentifier {
	_result := &_BACnetConstructedDataGlobalIdentifier{
		GlobalIdentifier: globalIdentifier,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataGlobalIdentifier(structType interface{}) BACnetConstructedDataGlobalIdentifier {
    if casted, ok := structType.(BACnetConstructedDataGlobalIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGlobalIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataGlobalIdentifier) GetTypeName() string {
	return "BACnetConstructedDataGlobalIdentifier"
}

func (m *_BACnetConstructedDataGlobalIdentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataGlobalIdentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (globalIdentifier)
	lengthInBits += m.GlobalIdentifier.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataGlobalIdentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataGlobalIdentifierParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataGlobalIdentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGlobalIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataGlobalIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (globalIdentifier)
	if pullErr := readBuffer.PullContext("globalIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for globalIdentifier")
	}
_globalIdentifier, _globalIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _globalIdentifierErr != nil {
		return nil, errors.Wrap(_globalIdentifierErr, "Error parsing 'globalIdentifier' field of BACnetConstructedDataGlobalIdentifier")
	}
	globalIdentifier := _globalIdentifier.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("globalIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for globalIdentifier")
	}

	// Virtual field
	_actualValue := globalIdentifier
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGlobalIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataGlobalIdentifier")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataGlobalIdentifier{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		GlobalIdentifier: globalIdentifier,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataGlobalIdentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataGlobalIdentifier) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGlobalIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataGlobalIdentifier")
		}

	// Simple Field (globalIdentifier)
	if pushErr := writeBuffer.PushContext("globalIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for globalIdentifier")
	}
	_globalIdentifierErr := writeBuffer.WriteSerializable(m.GetGlobalIdentifier())
	if popErr := writeBuffer.PopContext("globalIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for globalIdentifier")
	}
	if _globalIdentifierErr != nil {
		return errors.Wrap(_globalIdentifierErr, "Error serializing 'globalIdentifier' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGlobalIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataGlobalIdentifier")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataGlobalIdentifier) isBACnetConstructedDataGlobalIdentifier() bool {
	return true
}

func (m *_BACnetConstructedDataGlobalIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



