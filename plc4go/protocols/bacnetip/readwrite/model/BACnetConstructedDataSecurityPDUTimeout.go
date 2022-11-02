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


// BACnetConstructedDataSecurityPDUTimeout is the corresponding interface of BACnetConstructedDataSecurityPDUTimeout
type BACnetConstructedDataSecurityPDUTimeout interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSecurityPduTimeout returns SecurityPduTimeout (property field)
	GetSecurityPduTimeout() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataSecurityPDUTimeoutExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSecurityPDUTimeout.
// This is useful for switch cases.
type BACnetConstructedDataSecurityPDUTimeoutExactly interface {
	BACnetConstructedDataSecurityPDUTimeout
	isBACnetConstructedDataSecurityPDUTimeout() bool
}

// _BACnetConstructedDataSecurityPDUTimeout is the data-structure of this message
type _BACnetConstructedDataSecurityPDUTimeout struct {
	*_BACnetConstructedData
        SecurityPduTimeout BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSecurityPDUTimeout)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataSecurityPDUTimeout)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_SECURITY_PDU_TIMEOUT}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSecurityPDUTimeout) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSecurityPDUTimeout)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSecurityPDUTimeout) GetSecurityPduTimeout() BACnetApplicationTagUnsignedInteger {
	return m.SecurityPduTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSecurityPDUTimeout) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetSecurityPduTimeout())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataSecurityPDUTimeout factory function for _BACnetConstructedDataSecurityPDUTimeout
func NewBACnetConstructedDataSecurityPDUTimeout( securityPduTimeout BACnetApplicationTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataSecurityPDUTimeout {
	_result := &_BACnetConstructedDataSecurityPDUTimeout{
		SecurityPduTimeout: securityPduTimeout,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSecurityPDUTimeout(structType interface{}) BACnetConstructedDataSecurityPDUTimeout {
    if casted, ok := structType.(BACnetConstructedDataSecurityPDUTimeout); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSecurityPDUTimeout); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSecurityPDUTimeout) GetTypeName() string {
	return "BACnetConstructedDataSecurityPDUTimeout"
}

func (m *_BACnetConstructedDataSecurityPDUTimeout) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataSecurityPDUTimeout) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (securityPduTimeout)
	lengthInBits += m.SecurityPduTimeout.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataSecurityPDUTimeout) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSecurityPDUTimeoutParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSecurityPDUTimeout, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSecurityPDUTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSecurityPDUTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (securityPduTimeout)
	if pullErr := readBuffer.PullContext("securityPduTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for securityPduTimeout")
	}
_securityPduTimeout, _securityPduTimeoutErr := BACnetApplicationTagParse(readBuffer)
	if _securityPduTimeoutErr != nil {
		return nil, errors.Wrap(_securityPduTimeoutErr, "Error parsing 'securityPduTimeout' field of BACnetConstructedDataSecurityPDUTimeout")
	}
	securityPduTimeout := _securityPduTimeout.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("securityPduTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for securityPduTimeout")
	}

	// Virtual field
	_actualValue := securityPduTimeout
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSecurityPDUTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSecurityPDUTimeout")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSecurityPDUTimeout{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		SecurityPduTimeout: securityPduTimeout,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSecurityPDUTimeout) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataSecurityPDUTimeout) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSecurityPDUTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSecurityPDUTimeout")
		}

	// Simple Field (securityPduTimeout)
	if pushErr := writeBuffer.PushContext("securityPduTimeout"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for securityPduTimeout")
	}
	_securityPduTimeoutErr := writeBuffer.WriteSerializable(m.GetSecurityPduTimeout())
	if popErr := writeBuffer.PopContext("securityPduTimeout"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for securityPduTimeout")
	}
	if _securityPduTimeoutErr != nil {
		return errors.Wrap(_securityPduTimeoutErr, "Error serializing 'securityPduTimeout' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSecurityPDUTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSecurityPDUTimeout")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataSecurityPDUTimeout) isBACnetConstructedDataSecurityPDUTimeout() bool {
	return true
}

func (m *_BACnetConstructedDataSecurityPDUTimeout) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



