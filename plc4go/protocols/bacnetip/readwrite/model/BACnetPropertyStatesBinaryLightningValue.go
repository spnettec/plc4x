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


// BACnetPropertyStatesBinaryLightningValue is the corresponding interface of BACnetPropertyStatesBinaryLightningValue
type BACnetPropertyStatesBinaryLightningValue interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetBinaryLightningValue returns BinaryLightningValue (property field)
	GetBinaryLightningValue() BACnetBinaryLightingPVTagged
}

// BACnetPropertyStatesBinaryLightningValueExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesBinaryLightningValue.
// This is useful for switch cases.
type BACnetPropertyStatesBinaryLightningValueExactly interface {
	BACnetPropertyStatesBinaryLightningValue
	isBACnetPropertyStatesBinaryLightningValue() bool
}

// _BACnetPropertyStatesBinaryLightningValue is the data-structure of this message
type _BACnetPropertyStatesBinaryLightningValue struct {
	*_BACnetPropertyStates
        BinaryLightningValue BACnetBinaryLightingPVTagged
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesBinaryLightningValue) InitializeParent(parent BACnetPropertyStates , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesBinaryLightningValue)  GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesBinaryLightningValue) GetBinaryLightningValue() BACnetBinaryLightingPVTagged {
	return m.BinaryLightningValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetPropertyStatesBinaryLightningValue factory function for _BACnetPropertyStatesBinaryLightningValue
func NewBACnetPropertyStatesBinaryLightningValue( binaryLightningValue BACnetBinaryLightingPVTagged , peekedTagHeader BACnetTagHeader ) *_BACnetPropertyStatesBinaryLightningValue {
	_result := &_BACnetPropertyStatesBinaryLightningValue{
		BinaryLightningValue: binaryLightningValue,
    	_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesBinaryLightningValue(structType interface{}) BACnetPropertyStatesBinaryLightningValue {
    if casted, ok := structType.(BACnetPropertyStatesBinaryLightningValue); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesBinaryLightningValue); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesBinaryLightningValue) GetTypeName() string {
	return "BACnetPropertyStatesBinaryLightningValue"
}

func (m *_BACnetPropertyStatesBinaryLightningValue) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesBinaryLightningValue) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (binaryLightningValue)
	lengthInBits += m.BinaryLightningValue.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetPropertyStatesBinaryLightningValue) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesBinaryLightningValueParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesBinaryLightningValue, error) {
	return BACnetPropertyStatesBinaryLightningValueParseWithBuffer(utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesBinaryLightningValueParseWithBuffer(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesBinaryLightningValue, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesBinaryLightningValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesBinaryLightningValue")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (binaryLightningValue)
	if pullErr := readBuffer.PullContext("binaryLightningValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for binaryLightningValue")
	}
_binaryLightningValue, _binaryLightningValueErr := BACnetBinaryLightingPVTaggedParseWithBuffer(readBuffer , uint8( peekedTagNumber ) , TagClass( TagClass_CONTEXT_SPECIFIC_TAGS ) )
	if _binaryLightningValueErr != nil {
		return nil, errors.Wrap(_binaryLightningValueErr, "Error parsing 'binaryLightningValue' field of BACnetPropertyStatesBinaryLightningValue")
	}
	binaryLightningValue := _binaryLightningValue.(BACnetBinaryLightingPVTagged)
	if closeErr := readBuffer.CloseContext("binaryLightningValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for binaryLightningValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesBinaryLightningValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesBinaryLightningValue")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesBinaryLightningValue{
		_BACnetPropertyStates: &_BACnetPropertyStates{
		},
		BinaryLightningValue: binaryLightningValue,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesBinaryLightningValue) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesBinaryLightningValue) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesBinaryLightningValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesBinaryLightningValue")
		}

	// Simple Field (binaryLightningValue)
	if pushErr := writeBuffer.PushContext("binaryLightningValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for binaryLightningValue")
	}
	_binaryLightningValueErr := writeBuffer.WriteSerializable(m.GetBinaryLightningValue())
	if popErr := writeBuffer.PopContext("binaryLightningValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for binaryLightningValue")
	}
	if _binaryLightningValueErr != nil {
		return errors.Wrap(_binaryLightningValueErr, "Error serializing 'binaryLightningValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesBinaryLightningValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesBinaryLightningValue")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetPropertyStatesBinaryLightningValue) isBACnetPropertyStatesBinaryLightningValue() bool {
	return true
}

func (m *_BACnetPropertyStatesBinaryLightningValue) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



