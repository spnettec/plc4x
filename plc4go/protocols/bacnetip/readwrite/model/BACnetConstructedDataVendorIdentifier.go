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


// BACnetConstructedDataVendorIdentifier is the corresponding interface of BACnetConstructedDataVendorIdentifier
type BACnetConstructedDataVendorIdentifier interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetVendorIdentifier returns VendorIdentifier (property field)
	GetVendorIdentifier() BACnetVendorIdTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetVendorIdTagged
}

// BACnetConstructedDataVendorIdentifierExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataVendorIdentifier.
// This is useful for switch cases.
type BACnetConstructedDataVendorIdentifierExactly interface {
	BACnetConstructedDataVendorIdentifier
	isBACnetConstructedDataVendorIdentifier() bool
}

// _BACnetConstructedDataVendorIdentifier is the data-structure of this message
type _BACnetConstructedDataVendorIdentifier struct {
	*_BACnetConstructedData
        VendorIdentifier BACnetVendorIdTagged
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataVendorIdentifier)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataVendorIdentifier)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_VENDOR_IDENTIFIER}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataVendorIdentifier) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataVendorIdentifier)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataVendorIdentifier) GetVendorIdentifier() BACnetVendorIdTagged {
	return m.VendorIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataVendorIdentifier) GetActualValue() BACnetVendorIdTagged {
	return CastBACnetVendorIdTagged(m.GetVendorIdentifier())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataVendorIdentifier factory function for _BACnetConstructedDataVendorIdentifier
func NewBACnetConstructedDataVendorIdentifier( vendorIdentifier BACnetVendorIdTagged , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataVendorIdentifier {
	_result := &_BACnetConstructedDataVendorIdentifier{
		VendorIdentifier: vendorIdentifier,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataVendorIdentifier(structType interface{}) BACnetConstructedDataVendorIdentifier {
    if casted, ok := structType.(BACnetConstructedDataVendorIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVendorIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataVendorIdentifier) GetTypeName() string {
	return "BACnetConstructedDataVendorIdentifier"
}

func (m *_BACnetConstructedDataVendorIdentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataVendorIdentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (vendorIdentifier)
	lengthInBits += m.VendorIdentifier.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataVendorIdentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataVendorIdentifierParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataVendorIdentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVendorIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataVendorIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (vendorIdentifier)
	if pullErr := readBuffer.PullContext("vendorIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for vendorIdentifier")
	}
_vendorIdentifier, _vendorIdentifierErr := BACnetVendorIdTaggedParse(readBuffer , uint8( uint8(0) ) , TagClass( TagClass_APPLICATION_TAGS ) )
	if _vendorIdentifierErr != nil {
		return nil, errors.Wrap(_vendorIdentifierErr, "Error parsing 'vendorIdentifier' field of BACnetConstructedDataVendorIdentifier")
	}
	vendorIdentifier := _vendorIdentifier.(BACnetVendorIdTagged)
	if closeErr := readBuffer.CloseContext("vendorIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for vendorIdentifier")
	}

	// Virtual field
	_actualValue := vendorIdentifier
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVendorIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataVendorIdentifier")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataVendorIdentifier{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		VendorIdentifier: vendorIdentifier,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataVendorIdentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataVendorIdentifier) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVendorIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataVendorIdentifier")
		}

	// Simple Field (vendorIdentifier)
	if pushErr := writeBuffer.PushContext("vendorIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for vendorIdentifier")
	}
	_vendorIdentifierErr := writeBuffer.WriteSerializable(m.GetVendorIdentifier())
	if popErr := writeBuffer.PopContext("vendorIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for vendorIdentifier")
	}
	if _vendorIdentifierErr != nil {
		return errors.Wrap(_vendorIdentifierErr, "Error serializing 'vendorIdentifier' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVendorIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataVendorIdentifier")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataVendorIdentifier) isBACnetConstructedDataVendorIdentifier() bool {
	return true
}

func (m *_BACnetConstructedDataVendorIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



