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

// BACnetConstructedDataProgramChange is the corresponding interface of BACnetConstructedDataProgramChange
type BACnetConstructedDataProgramChange interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetProgramChange returns ProgramChange (property field)
	GetProgramChange() BACnetProgramRequestTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetProgramRequestTagged
}

// BACnetConstructedDataProgramChangeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataProgramChange.
// This is useful for switch cases.
type BACnetConstructedDataProgramChangeExactly interface {
	BACnetConstructedDataProgramChange
	isBACnetConstructedDataProgramChange() bool
}

// _BACnetConstructedDataProgramChange is the data-structure of this message
type _BACnetConstructedDataProgramChange struct {
	*_BACnetConstructedData
	ProgramChange BACnetProgramRequestTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProgramChange) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataProgramChange) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_PROGRAM_CHANGE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProgramChange) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataProgramChange) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProgramChange) GetProgramChange() BACnetProgramRequestTagged {
	return m.ProgramChange
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProgramChange) GetActualValue() BACnetProgramRequestTagged {
	return CastBACnetProgramRequestTagged(m.GetProgramChange())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataProgramChange factory function for _BACnetConstructedDataProgramChange
func NewBACnetConstructedDataProgramChange(programChange BACnetProgramRequestTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataProgramChange {
	_result := &_BACnetConstructedDataProgramChange{
		ProgramChange:          programChange,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProgramChange(structType interface{}) BACnetConstructedDataProgramChange {
	if casted, ok := structType.(BACnetConstructedDataProgramChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProgramChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProgramChange) GetTypeName() string {
	return "BACnetConstructedDataProgramChange"
}

func (m *_BACnetConstructedDataProgramChange) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataProgramChange) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (programChange)
	lengthInBits += m.ProgramChange.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataProgramChange) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataProgramChangeParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProgramChange, error) {
	return BACnetConstructedDataProgramChangeParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataProgramChangeParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProgramChange, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProgramChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProgramChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (programChange)
	if pullErr := readBuffer.PullContext("programChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for programChange")
	}
	_programChange, _programChangeErr := BACnetProgramRequestTaggedParseWithBuffer(readBuffer, uint8(uint8(0)), TagClass(TagClass_APPLICATION_TAGS))
	if _programChangeErr != nil {
		return nil, errors.Wrap(_programChangeErr, "Error parsing 'programChange' field of BACnetConstructedDataProgramChange")
	}
	programChange := _programChange.(BACnetProgramRequestTagged)
	if closeErr := readBuffer.CloseContext("programChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for programChange")
	}

	// Virtual field
	_actualValue := programChange
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProgramChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProgramChange")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataProgramChange{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ProgramChange: programChange,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataProgramChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataProgramChange) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProgramChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProgramChange")
		}

		// Simple Field (programChange)
		if pushErr := writeBuffer.PushContext("programChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for programChange")
		}
		_programChangeErr := writeBuffer.WriteSerializable(m.GetProgramChange())
		if popErr := writeBuffer.PopContext("programChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for programChange")
		}
		if _programChangeErr != nil {
			return errors.Wrap(_programChangeErr, "Error serializing 'programChange' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProgramChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProgramChange")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataProgramChange) isBACnetConstructedDataProgramChange() bool {
	return true
}

func (m *_BACnetConstructedDataProgramChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
