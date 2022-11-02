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


// BACnetPriorityValueBoolean is the corresponding interface of BACnetPriorityValueBoolean
type BACnetPriorityValueBoolean interface {
	utils.LengthAware
	utils.Serializable
	BACnetPriorityValue
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() BACnetApplicationTagBoolean
}

// BACnetPriorityValueBooleanExactly can be used when we want exactly this type and not a type which fulfills BACnetPriorityValueBoolean.
// This is useful for switch cases.
type BACnetPriorityValueBooleanExactly interface {
	BACnetPriorityValueBoolean
	isBACnetPriorityValueBoolean() bool
}

// _BACnetPriorityValueBoolean is the data-structure of this message
type _BACnetPriorityValueBoolean struct {
	*_BACnetPriorityValue
        BooleanValue BACnetApplicationTagBoolean
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPriorityValueBoolean) InitializeParent(parent BACnetPriorityValue , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPriorityValueBoolean)  GetParent() BACnetPriorityValue {
	return m._BACnetPriorityValue
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPriorityValueBoolean) GetBooleanValue() BACnetApplicationTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetPriorityValueBoolean factory function for _BACnetPriorityValueBoolean
func NewBACnetPriorityValueBoolean( booleanValue BACnetApplicationTagBoolean , peekedTagHeader BACnetTagHeader , objectTypeArgument BACnetObjectType ) *_BACnetPriorityValueBoolean {
	_result := &_BACnetPriorityValueBoolean{
		BooleanValue: booleanValue,
    	_BACnetPriorityValue: NewBACnetPriorityValue(peekedTagHeader, objectTypeArgument),
	}
	_result._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPriorityValueBoolean(structType interface{}) BACnetPriorityValueBoolean {
    if casted, ok := structType.(BACnetPriorityValueBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPriorityValueBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPriorityValueBoolean) GetTypeName() string {
	return "BACnetPriorityValueBoolean"
}

func (m *_BACnetPriorityValueBoolean) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPriorityValueBoolean) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetPriorityValueBoolean) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPriorityValueBooleanParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType) (BACnetPriorityValueBoolean, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPriorityValueBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPriorityValueBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (booleanValue)
	if pullErr := readBuffer.PullContext("booleanValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for booleanValue")
	}
_booleanValue, _booleanValueErr := BACnetApplicationTagParse(readBuffer)
	if _booleanValueErr != nil {
		return nil, errors.Wrap(_booleanValueErr, "Error parsing 'booleanValue' field of BACnetPriorityValueBoolean")
	}
	booleanValue := _booleanValue.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("booleanValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for booleanValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetPriorityValueBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPriorityValueBoolean")
	}

	// Create a partially initialized instance
	_child := &_BACnetPriorityValueBoolean{
		_BACnetPriorityValue: &_BACnetPriorityValue{
			ObjectTypeArgument: objectTypeArgument,
		},
		BooleanValue: booleanValue,
	}
	_child._BACnetPriorityValue._BACnetPriorityValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPriorityValueBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPriorityValueBoolean) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPriorityValueBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPriorityValueBoolean")
		}

	// Simple Field (booleanValue)
	if pushErr := writeBuffer.PushContext("booleanValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for booleanValue")
	}
	_booleanValueErr := writeBuffer.WriteSerializable(m.GetBooleanValue())
	if popErr := writeBuffer.PopContext("booleanValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for booleanValue")
	}
	if _booleanValueErr != nil {
		return errors.Wrap(_booleanValueErr, "Error serializing 'booleanValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetPriorityValueBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPriorityValueBoolean")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetPriorityValueBoolean) isBACnetPriorityValueBoolean() bool {
	return true
}

func (m *_BACnetPriorityValueBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



