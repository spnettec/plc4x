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


// BACnetConstructedDataLargeAnalogValueFaultLowLimit is the corresponding interface of BACnetConstructedDataLargeAnalogValueFaultLowLimit
type BACnetConstructedDataLargeAnalogValueFaultLowLimit interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultLowLimit returns FaultLowLimit (property field)
	GetFaultLowLimit() BACnetApplicationTagDouble
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDouble
}

// BACnetConstructedDataLargeAnalogValueFaultLowLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLargeAnalogValueFaultLowLimit.
// This is useful for switch cases.
type BACnetConstructedDataLargeAnalogValueFaultLowLimitExactly interface {
	BACnetConstructedDataLargeAnalogValueFaultLowLimit
	isBACnetConstructedDataLargeAnalogValueFaultLowLimit() bool
}

// _BACnetConstructedDataLargeAnalogValueFaultLowLimit is the data-structure of this message
type _BACnetConstructedDataLargeAnalogValueFaultLowLimit struct {
	*_BACnetConstructedData
        FaultLowLimit BACnetApplicationTagDouble
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit)  GetObjectTypeArgument() BACnetObjectType {
return BACnetObjectType_LARGE_ANALOG_VALUE}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_FAULT_LOW_LIMIT}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) GetFaultLowLimit() BACnetApplicationTagDouble {
	return m.FaultLowLimit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) GetActualValue() BACnetApplicationTagDouble {
	return CastBACnetApplicationTagDouble(m.GetFaultLowLimit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataLargeAnalogValueFaultLowLimit factory function for _BACnetConstructedDataLargeAnalogValueFaultLowLimit
func NewBACnetConstructedDataLargeAnalogValueFaultLowLimit( faultLowLimit BACnetApplicationTagDouble , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataLargeAnalogValueFaultLowLimit {
	_result := &_BACnetConstructedDataLargeAnalogValueFaultLowLimit{
		FaultLowLimit: faultLowLimit,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLargeAnalogValueFaultLowLimit(structType interface{}) BACnetConstructedDataLargeAnalogValueFaultLowLimit {
    if casted, ok := structType.(BACnetConstructedDataLargeAnalogValueFaultLowLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLargeAnalogValueFaultLowLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) GetTypeName() string {
	return "BACnetConstructedDataLargeAnalogValueFaultLowLimit"
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (faultLowLimit)
	lengthInBits += m.FaultLowLimit.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataLargeAnalogValueFaultLowLimitParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLargeAnalogValueFaultLowLimit, error) {
	return BACnetConstructedDataLargeAnalogValueFaultLowLimitParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLargeAnalogValueFaultLowLimitParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLargeAnalogValueFaultLowLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLargeAnalogValueFaultLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLargeAnalogValueFaultLowLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (faultLowLimit)
	if pullErr := readBuffer.PullContext("faultLowLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for faultLowLimit")
	}
_faultLowLimit, _faultLowLimitErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _faultLowLimitErr != nil {
		return nil, errors.Wrap(_faultLowLimitErr, "Error parsing 'faultLowLimit' field of BACnetConstructedDataLargeAnalogValueFaultLowLimit")
	}
	faultLowLimit := _faultLowLimit.(BACnetApplicationTagDouble)
	if closeErr := readBuffer.CloseContext("faultLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for faultLowLimit")
	}

	// Virtual field
	_actualValue := faultLowLimit
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLargeAnalogValueFaultLowLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLargeAnalogValueFaultLowLimit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLargeAnalogValueFaultLowLimit{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FaultLowLimit: faultLowLimit,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLargeAnalogValueFaultLowLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLargeAnalogValueFaultLowLimit")
		}

	// Simple Field (faultLowLimit)
	if pushErr := writeBuffer.PushContext("faultLowLimit"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for faultLowLimit")
	}
	_faultLowLimitErr := writeBuffer.WriteSerializable(m.GetFaultLowLimit())
	if popErr := writeBuffer.PopContext("faultLowLimit"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for faultLowLimit")
	}
	if _faultLowLimitErr != nil {
		return errors.Wrap(_faultLowLimitErr, "Error serializing 'faultLowLimit' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLargeAnalogValueFaultLowLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLargeAnalogValueFaultLowLimit")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) isBACnetConstructedDataLargeAnalogValueFaultLowLimit() bool {
	return true
}

func (m *_BACnetConstructedDataLargeAnalogValueFaultLowLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



