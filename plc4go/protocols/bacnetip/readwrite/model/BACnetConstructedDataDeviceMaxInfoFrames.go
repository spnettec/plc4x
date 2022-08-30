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


// BACnetConstructedDataDeviceMaxInfoFrames is the corresponding interface of BACnetConstructedDataDeviceMaxInfoFrames
type BACnetConstructedDataDeviceMaxInfoFrames interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaxInfoFrames returns MaxInfoFrames (property field)
	GetMaxInfoFrames() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataDeviceMaxInfoFramesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDeviceMaxInfoFrames.
// This is useful for switch cases.
type BACnetConstructedDataDeviceMaxInfoFramesExactly interface {
	BACnetConstructedDataDeviceMaxInfoFrames
	isBACnetConstructedDataDeviceMaxInfoFrames() bool
}

// _BACnetConstructedDataDeviceMaxInfoFrames is the data-structure of this message
type _BACnetConstructedDataDeviceMaxInfoFrames struct {
	*_BACnetConstructedData
        MaxInfoFrames BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDeviceMaxInfoFrames)  GetObjectTypeArgument() BACnetObjectType {
return BACnetObjectType_DEVICE}

func (m *_BACnetConstructedDataDeviceMaxInfoFrames)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_MAX_INFO_FRAMES}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDeviceMaxInfoFrames)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) GetMaxInfoFrames() BACnetApplicationTagUnsignedInteger {
	return m.MaxInfoFrames
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetMaxInfoFrames())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataDeviceMaxInfoFrames factory function for _BACnetConstructedDataDeviceMaxInfoFrames
func NewBACnetConstructedDataDeviceMaxInfoFrames( maxInfoFrames BACnetApplicationTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataDeviceMaxInfoFrames {
	_result := &_BACnetConstructedDataDeviceMaxInfoFrames{
		MaxInfoFrames: maxInfoFrames,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDeviceMaxInfoFrames(structType interface{}) BACnetConstructedDataDeviceMaxInfoFrames {
    if casted, ok := structType.(BACnetConstructedDataDeviceMaxInfoFrames); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDeviceMaxInfoFrames); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) GetTypeName() string {
	return "BACnetConstructedDataDeviceMaxInfoFrames"
}

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maxInfoFrames)
	lengthInBits += m.MaxInfoFrames.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataDeviceMaxInfoFrames) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDeviceMaxInfoFramesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDeviceMaxInfoFrames, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDeviceMaxInfoFrames"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDeviceMaxInfoFrames")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maxInfoFrames)
	if pullErr := readBuffer.PullContext("maxInfoFrames"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxInfoFrames")
	}
_maxInfoFrames, _maxInfoFramesErr := BACnetApplicationTagParse(readBuffer)
	if _maxInfoFramesErr != nil {
		return nil, errors.Wrap(_maxInfoFramesErr, "Error parsing 'maxInfoFrames' field of BACnetConstructedDataDeviceMaxInfoFrames")
	}
	maxInfoFrames := _maxInfoFrames.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("maxInfoFrames"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxInfoFrames")
	}

	// Virtual field
	_actualValue := maxInfoFrames
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDeviceMaxInfoFrames"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDeviceMaxInfoFrames")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDeviceMaxInfoFrames{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MaxInfoFrames: maxInfoFrames,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDeviceMaxInfoFrames"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDeviceMaxInfoFrames")
		}

	// Simple Field (maxInfoFrames)
	if pushErr := writeBuffer.PushContext("maxInfoFrames"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for maxInfoFrames")
	}
	_maxInfoFramesErr := writeBuffer.WriteSerializable(m.GetMaxInfoFrames())
	if popErr := writeBuffer.PopContext("maxInfoFrames"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for maxInfoFrames")
	}
	if _maxInfoFramesErr != nil {
		return errors.Wrap(_maxInfoFramesErr, "Error serializing 'maxInfoFrames' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDeviceMaxInfoFrames"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDeviceMaxInfoFrames")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataDeviceMaxInfoFrames) isBACnetConstructedDataDeviceMaxInfoFrames() bool {
	return true
}

func (m *_BACnetConstructedDataDeviceMaxInfoFrames) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



