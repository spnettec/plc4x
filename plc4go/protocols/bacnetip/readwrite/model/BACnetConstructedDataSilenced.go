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

// BACnetConstructedDataSilenced is the corresponding interface of BACnetConstructedDataSilenced
type BACnetConstructedDataSilenced interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSilenced returns Silenced (property field)
	GetSilenced() BACnetSilencedStateTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetSilencedStateTagged
}

// BACnetConstructedDataSilencedExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSilenced.
// This is useful for switch cases.
type BACnetConstructedDataSilencedExactly interface {
	BACnetConstructedDataSilenced
	isBACnetConstructedDataSilenced() bool
}

// _BACnetConstructedDataSilenced is the data-structure of this message
type _BACnetConstructedDataSilenced struct {
	*_BACnetConstructedData
	Silenced BACnetSilencedStateTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSilenced) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataSilenced) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_SILENCED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSilenced) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSilenced) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSilenced) GetSilenced() BACnetSilencedStateTagged {
	return m.Silenced
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSilenced) GetActualValue() BACnetSilencedStateTagged {
	return CastBACnetSilencedStateTagged(m.GetSilenced())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataSilenced factory function for _BACnetConstructedDataSilenced
func NewBACnetConstructedDataSilenced(silenced BACnetSilencedStateTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataSilenced {
	_result := &_BACnetConstructedDataSilenced{
		Silenced:               silenced,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSilenced(structType interface{}) BACnetConstructedDataSilenced {
	if casted, ok := structType.(BACnetConstructedDataSilenced); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSilenced); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSilenced) GetTypeName() string {
	return "BACnetConstructedDataSilenced"
}

func (m *_BACnetConstructedDataSilenced) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataSilenced) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (silenced)
	lengthInBits += m.Silenced.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataSilenced) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSilencedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSilenced, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSilenced"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSilenced")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (silenced)
	if pullErr := readBuffer.PullContext("silenced"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for silenced")
	}
	_silenced, _silencedErr := BACnetSilencedStateTaggedParse(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _silencedErr != nil {
		return nil, errors.Wrap(_silencedErr, "Error parsing 'silenced' field of BACnetConstructedDataSilenced")
	}
	silenced := _silenced.(BACnetSilencedStateTagged)
	if closeErr := readBuffer.CloseContext("silenced"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for silenced")
	}

	// Virtual field
	_actualValue := silenced
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSilenced"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSilenced")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSilenced{
		Silenced: silenced,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSilenced) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSilenced"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSilenced")
		}

		// Simple Field (silenced)
		if pushErr := writeBuffer.PushContext("silenced"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for silenced")
		}
		_silencedErr := writeBuffer.WriteSerializable(m.GetSilenced())
		if popErr := writeBuffer.PopContext("silenced"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for silenced")
		}
		if _silencedErr != nil {
			return errors.Wrap(_silencedErr, "Error serializing 'silenced' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSilenced"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSilenced")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataSilenced) isBACnetConstructedDataSilenced() bool {
	return true
}

func (m *_BACnetConstructedDataSilenced) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
