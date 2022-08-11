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

// BACnetConstructedDataCarLoad is the corresponding interface of BACnetConstructedDataCarLoad
type BACnetConstructedDataCarLoad interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCarLoad returns CarLoad (property field)
	GetCarLoad() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataCarLoadExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataCarLoad.
// This is useful for switch cases.
type BACnetConstructedDataCarLoadExactly interface {
	BACnetConstructedDataCarLoad
	isBACnetConstructedDataCarLoad() bool
}

// _BACnetConstructedDataCarLoad is the data-structure of this message
type _BACnetConstructedDataCarLoad struct {
	*_BACnetConstructedData
	CarLoad BACnetApplicationTagReal
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataCarLoad) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataCarLoad) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_CAR_LOAD
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataCarLoad) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataCarLoad) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataCarLoad) GetCarLoad() BACnetApplicationTagReal {
	return m.CarLoad
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataCarLoad) GetActualValue() BACnetApplicationTagReal {
	return CastBACnetApplicationTagReal(m.GetCarLoad())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataCarLoad factory function for _BACnetConstructedDataCarLoad
func NewBACnetConstructedDataCarLoad(carLoad BACnetApplicationTagReal, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataCarLoad {
	_result := &_BACnetConstructedDataCarLoad{
		CarLoad:                carLoad,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataCarLoad(structType interface{}) BACnetConstructedDataCarLoad {
	if casted, ok := structType.(BACnetConstructedDataCarLoad); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataCarLoad); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataCarLoad) GetTypeName() string {
	return "BACnetConstructedDataCarLoad"
}

func (m *_BACnetConstructedDataCarLoad) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataCarLoad) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (carLoad)
	lengthInBits += m.CarLoad.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataCarLoad) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataCarLoadParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataCarLoad, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataCarLoad"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataCarLoad")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (carLoad)
	if pullErr := readBuffer.PullContext("carLoad"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for carLoad")
	}
	_carLoad, _carLoadErr := BACnetApplicationTagParse(readBuffer)
	if _carLoadErr != nil {
		return nil, errors.Wrap(_carLoadErr, "Error parsing 'carLoad' field of BACnetConstructedDataCarLoad")
	}
	carLoad := _carLoad.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("carLoad"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for carLoad")
	}

	// Virtual field
	_actualValue := carLoad
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataCarLoad"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataCarLoad")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataCarLoad{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		CarLoad: carLoad,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataCarLoad) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataCarLoad"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataCarLoad")
		}

		// Simple Field (carLoad)
		if pushErr := writeBuffer.PushContext("carLoad"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for carLoad")
		}
		_carLoadErr := writeBuffer.WriteSerializable(m.GetCarLoad())
		if popErr := writeBuffer.PopContext("carLoad"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for carLoad")
		}
		if _carLoadErr != nil {
			return errors.Wrap(_carLoadErr, "Error serializing 'carLoad' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataCarLoad"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataCarLoad")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataCarLoad) isBACnetConstructedDataCarLoad() bool {
	return true
}

func (m *_BACnetConstructedDataCarLoad) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
