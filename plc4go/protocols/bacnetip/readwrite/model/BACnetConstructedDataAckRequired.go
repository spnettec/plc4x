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


// BACnetConstructedDataAckRequired is the corresponding interface of BACnetConstructedDataAckRequired
type BACnetConstructedDataAckRequired interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAckRequired returns AckRequired (property field)
	GetAckRequired() BACnetEventTransitionBitsTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetEventTransitionBitsTagged
}

// BACnetConstructedDataAckRequiredExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAckRequired.
// This is useful for switch cases.
type BACnetConstructedDataAckRequiredExactly interface {
	BACnetConstructedDataAckRequired
	isBACnetConstructedDataAckRequired() bool
}

// _BACnetConstructedDataAckRequired is the data-structure of this message
type _BACnetConstructedDataAckRequired struct {
	*_BACnetConstructedData
        AckRequired BACnetEventTransitionBitsTagged
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAckRequired)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataAckRequired)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_ACK_REQUIRED}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAckRequired) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAckRequired)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAckRequired) GetAckRequired() BACnetEventTransitionBitsTagged {
	return m.AckRequired
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAckRequired) GetActualValue() BACnetEventTransitionBitsTagged {
	return CastBACnetEventTransitionBitsTagged(m.GetAckRequired())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataAckRequired factory function for _BACnetConstructedDataAckRequired
func NewBACnetConstructedDataAckRequired( ackRequired BACnetEventTransitionBitsTagged , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataAckRequired {
	_result := &_BACnetConstructedDataAckRequired{
		AckRequired: ackRequired,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAckRequired(structType interface{}) BACnetConstructedDataAckRequired {
    if casted, ok := structType.(BACnetConstructedDataAckRequired); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAckRequired); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAckRequired) GetTypeName() string {
	return "BACnetConstructedDataAckRequired"
}

func (m *_BACnetConstructedDataAckRequired) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAckRequired) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (ackRequired)
	lengthInBits += m.AckRequired.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataAckRequired) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAckRequiredParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAckRequired, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAckRequired"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAckRequired")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (ackRequired)
	if pullErr := readBuffer.PullContext("ackRequired"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ackRequired")
	}
_ackRequired, _ackRequiredErr := BACnetEventTransitionBitsTaggedParse(readBuffer , uint8( uint8(0) ) , TagClass( TagClass_APPLICATION_TAGS ) )
	if _ackRequiredErr != nil {
		return nil, errors.Wrap(_ackRequiredErr, "Error parsing 'ackRequired' field of BACnetConstructedDataAckRequired")
	}
	ackRequired := _ackRequired.(BACnetEventTransitionBitsTagged)
	if closeErr := readBuffer.CloseContext("ackRequired"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ackRequired")
	}

	// Virtual field
	_actualValue := ackRequired
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAckRequired"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAckRequired")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAckRequired{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		AckRequired: ackRequired,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAckRequired) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAckRequired) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAckRequired"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAckRequired")
		}

	// Simple Field (ackRequired)
	if pushErr := writeBuffer.PushContext("ackRequired"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ackRequired")
	}
	_ackRequiredErr := writeBuffer.WriteSerializable(m.GetAckRequired())
	if popErr := writeBuffer.PopContext("ackRequired"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ackRequired")
	}
	if _ackRequiredErr != nil {
		return errors.Wrap(_ackRequiredErr, "Error serializing 'ackRequired' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAckRequired"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAckRequired")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataAckRequired) isBACnetConstructedDataAckRequired() bool {
	return true
}

func (m *_BACnetConstructedDataAckRequired) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



