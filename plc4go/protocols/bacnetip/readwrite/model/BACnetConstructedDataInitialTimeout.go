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

// BACnetConstructedDataInitialTimeout is the data-structure of this message
type BACnetConstructedDataInitialTimeout struct {
	*BACnetConstructedData
	InitialTimeout *BACnetApplicationTagUnsignedInteger

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataInitialTimeout is the corresponding interface of BACnetConstructedDataInitialTimeout
type IBACnetConstructedDataInitialTimeout interface {
	IBACnetConstructedData
	// GetInitialTimeout returns InitialTimeout (property field)
	GetInitialTimeout() *BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagUnsignedInteger
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

func (m *BACnetConstructedDataInitialTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataInitialTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INITIAL_TIMEOUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataInitialTimeout) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataInitialTimeout) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataInitialTimeout) GetInitialTimeout() *BACnetApplicationTagUnsignedInteger {
	return m.InitialTimeout
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataInitialTimeout) GetActualValue() *BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetInitialTimeout())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataInitialTimeout factory function for BACnetConstructedDataInitialTimeout
func NewBACnetConstructedDataInitialTimeout(initialTimeout *BACnetApplicationTagUnsignedInteger, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataInitialTimeout {
	_result := &BACnetConstructedDataInitialTimeout{
		InitialTimeout:        initialTimeout,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataInitialTimeout(structType interface{}) *BACnetConstructedDataInitialTimeout {
	if casted, ok := structType.(BACnetConstructedDataInitialTimeout); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInitialTimeout); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataInitialTimeout(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataInitialTimeout(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataInitialTimeout) GetTypeName() string {
	return "BACnetConstructedDataInitialTimeout"
}

func (m *BACnetConstructedDataInitialTimeout) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataInitialTimeout) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (initialTimeout)
	lengthInBits += m.InitialTimeout.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataInitialTimeout) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataInitialTimeoutParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataInitialTimeout, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInitialTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInitialTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (initialTimeout)
	if pullErr := readBuffer.PullContext("initialTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for initialTimeout")
	}
	_initialTimeout, _initialTimeoutErr := BACnetApplicationTagParse(readBuffer)
	if _initialTimeoutErr != nil {
		return nil, errors.Wrap(_initialTimeoutErr, "Error parsing 'initialTimeout' field")
	}
	initialTimeout := CastBACnetApplicationTagUnsignedInteger(_initialTimeout)
	if closeErr := readBuffer.CloseContext("initialTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for initialTimeout")
	}

	// Virtual field
	_actualValue := initialTimeout
	actualValue := CastBACnetApplicationTagUnsignedInteger(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInitialTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInitialTimeout")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataInitialTimeout{
		InitialTimeout:        CastBACnetApplicationTagUnsignedInteger(initialTimeout),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataInitialTimeout) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInitialTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInitialTimeout")
		}

		// Simple Field (initialTimeout)
		if pushErr := writeBuffer.PushContext("initialTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for initialTimeout")
		}
		_initialTimeoutErr := writeBuffer.WriteSerializable(m.InitialTimeout)
		if popErr := writeBuffer.PopContext("initialTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for initialTimeout")
		}
		if _initialTimeoutErr != nil {
			return errors.Wrap(_initialTimeoutErr, "Error serializing 'initialTimeout' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInitialTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInitialTimeout")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataInitialTimeout) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
