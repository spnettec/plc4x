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

// BACnetConstructedDataEgressActive is the data-structure of this message
type BACnetConstructedDataEgressActive struct {
	*BACnetConstructedData
	EgressActive *BACnetApplicationTagBoolean

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataEgressActive is the corresponding interface of BACnetConstructedDataEgressActive
type IBACnetConstructedDataEgressActive interface {
	IBACnetConstructedData
	// GetEgressActive returns EgressActive (property field)
	GetEgressActive() *BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() *BACnetApplicationTagBoolean
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

func (m *BACnetConstructedDataEgressActive) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataEgressActive) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_EGRESS_ACTIVE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataEgressActive) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataEgressActive) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataEgressActive) GetEgressActive() *BACnetApplicationTagBoolean {
	return m.EgressActive
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConstructedDataEgressActive) GetActualValue() *BACnetApplicationTagBoolean {
	return CastBACnetApplicationTagBoolean(m.GetEgressActive())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataEgressActive factory function for BACnetConstructedDataEgressActive
func NewBACnetConstructedDataEgressActive(egressActive *BACnetApplicationTagBoolean, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataEgressActive {
	_result := &BACnetConstructedDataEgressActive{
		EgressActive:          egressActive,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataEgressActive(structType interface{}) *BACnetConstructedDataEgressActive {
	if casted, ok := structType.(BACnetConstructedDataEgressActive); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEgressActive); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataEgressActive(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataEgressActive(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataEgressActive) GetTypeName() string {
	return "BACnetConstructedDataEgressActive"
}

func (m *BACnetConstructedDataEgressActive) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataEgressActive) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (egressActive)
	lengthInBits += m.EgressActive.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *BACnetConstructedDataEgressActive) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEgressActiveParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataEgressActive, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEgressActive"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEgressActive")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (egressActive)
	if pullErr := readBuffer.PullContext("egressActive"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for egressActive")
	}
	_egressActive, _egressActiveErr := BACnetApplicationTagParse(readBuffer)
	if _egressActiveErr != nil {
		return nil, errors.Wrap(_egressActiveErr, "Error parsing 'egressActive' field")
	}
	egressActive := CastBACnetApplicationTagBoolean(_egressActive)
	if closeErr := readBuffer.CloseContext("egressActive"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for egressActive")
	}

	// Virtual field
	_actualValue := egressActive
	actualValue := CastBACnetApplicationTagBoolean(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEgressActive"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEgressActive")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataEgressActive{
		EgressActive:          CastBACnetApplicationTagBoolean(egressActive),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataEgressActive) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEgressActive"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEgressActive")
		}

		// Simple Field (egressActive)
		if pushErr := writeBuffer.PushContext("egressActive"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for egressActive")
		}
		_egressActiveErr := writeBuffer.WriteSerializable(m.EgressActive)
		if popErr := writeBuffer.PopContext("egressActive"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for egressActive")
		}
		if _egressActiveErr != nil {
			return errors.Wrap(_egressActiveErr, "Error serializing 'egressActive' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEgressActive"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEgressActive")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataEgressActive) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
