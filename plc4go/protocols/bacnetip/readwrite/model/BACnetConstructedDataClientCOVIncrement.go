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


// BACnetConstructedDataClientCOVIncrement is the corresponding interface of BACnetConstructedDataClientCOVIncrement
type BACnetConstructedDataClientCOVIncrement interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() BACnetClientCOV
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetClientCOV
}

// BACnetConstructedDataClientCOVIncrementExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataClientCOVIncrement.
// This is useful for switch cases.
type BACnetConstructedDataClientCOVIncrementExactly interface {
	BACnetConstructedDataClientCOVIncrement
	isBACnetConstructedDataClientCOVIncrement() bool
}

// _BACnetConstructedDataClientCOVIncrement is the data-structure of this message
type _BACnetConstructedDataClientCOVIncrement struct {
	*_BACnetConstructedData
        CovIncrement BACnetClientCOV
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataClientCOVIncrement)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataClientCOVIncrement)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_CLIENT_COV_INCREMENT}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataClientCOVIncrement) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataClientCOVIncrement)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataClientCOVIncrement) GetCovIncrement() BACnetClientCOV {
	return m.CovIncrement
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataClientCOVIncrement) GetActualValue() BACnetClientCOV {
	return CastBACnetClientCOV(m.GetCovIncrement())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataClientCOVIncrement factory function for _BACnetConstructedDataClientCOVIncrement
func NewBACnetConstructedDataClientCOVIncrement( covIncrement BACnetClientCOV , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataClientCOVIncrement {
	_result := &_BACnetConstructedDataClientCOVIncrement{
		CovIncrement: covIncrement,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataClientCOVIncrement(structType interface{}) BACnetConstructedDataClientCOVIncrement {
    if casted, ok := structType.(BACnetConstructedDataClientCOVIncrement); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataClientCOVIncrement); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataClientCOVIncrement) GetTypeName() string {
	return "BACnetConstructedDataClientCOVIncrement"
}

func (m *_BACnetConstructedDataClientCOVIncrement) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataClientCOVIncrement) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (covIncrement)
	lengthInBits += m.CovIncrement.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataClientCOVIncrement) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataClientCOVIncrementParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataClientCOVIncrement, error) {
	return BACnetConstructedDataClientCOVIncrementParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataClientCOVIncrementParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataClientCOVIncrement, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataClientCOVIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataClientCOVIncrement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (covIncrement)
	if pullErr := readBuffer.PullContext("covIncrement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for covIncrement")
	}
_covIncrement, _covIncrementErr := BACnetClientCOVParseWithBuffer(readBuffer)
	if _covIncrementErr != nil {
		return nil, errors.Wrap(_covIncrementErr, "Error parsing 'covIncrement' field of BACnetConstructedDataClientCOVIncrement")
	}
	covIncrement := _covIncrement.(BACnetClientCOV)
	if closeErr := readBuffer.CloseContext("covIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for covIncrement")
	}

	// Virtual field
	_actualValue := covIncrement
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataClientCOVIncrement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataClientCOVIncrement")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataClientCOVIncrement{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		CovIncrement: covIncrement,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataClientCOVIncrement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataClientCOVIncrement) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataClientCOVIncrement"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataClientCOVIncrement")
		}

	// Simple Field (covIncrement)
	if pushErr := writeBuffer.PushContext("covIncrement"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for covIncrement")
	}
	_covIncrementErr := writeBuffer.WriteSerializable(m.GetCovIncrement())
	if popErr := writeBuffer.PopContext("covIncrement"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for covIncrement")
	}
	if _covIncrementErr != nil {
		return errors.Wrap(_covIncrementErr, "Error serializing 'covIncrement' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataClientCOVIncrement"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataClientCOVIncrement")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataClientCOVIncrement) isBACnetConstructedDataClientCOVIncrement() bool {
	return true
}

func (m *_BACnetConstructedDataClientCOVIncrement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



