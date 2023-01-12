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


// BACnetChannelValueBoolean is the corresponding interface of BACnetChannelValueBoolean
type BACnetChannelValueBoolean interface {
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetBooleanValue returns BooleanValue (property field)
	GetBooleanValue() BACnetApplicationTagBoolean
}

// BACnetChannelValueBooleanExactly can be used when we want exactly this type and not a type which fulfills BACnetChannelValueBoolean.
// This is useful for switch cases.
type BACnetChannelValueBooleanExactly interface {
	BACnetChannelValueBoolean
	isBACnetChannelValueBoolean() bool
}

// _BACnetChannelValueBoolean is the data-structure of this message
type _BACnetChannelValueBoolean struct {
	*_BACnetChannelValue
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

func (m *_BACnetChannelValueBoolean) InitializeParent(parent BACnetChannelValue , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetChannelValueBoolean)  GetParent() BACnetChannelValue {
	return m._BACnetChannelValue
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueBoolean) GetBooleanValue() BACnetApplicationTagBoolean {
	return m.BooleanValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetChannelValueBoolean factory function for _BACnetChannelValueBoolean
func NewBACnetChannelValueBoolean( booleanValue BACnetApplicationTagBoolean , peekedTagHeader BACnetTagHeader ) *_BACnetChannelValueBoolean {
	_result := &_BACnetChannelValueBoolean{
		BooleanValue: booleanValue,
    	_BACnetChannelValue: NewBACnetChannelValue(peekedTagHeader),
	}
	_result._BACnetChannelValue._BACnetChannelValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueBoolean(structType interface{}) BACnetChannelValueBoolean {
    if casted, ok := structType.(BACnetChannelValueBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueBoolean) GetTypeName() string {
	return "BACnetChannelValueBoolean"
}

func (m *_BACnetChannelValueBoolean) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetChannelValueBoolean) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (booleanValue)
	lengthInBits += m.BooleanValue.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetChannelValueBoolean) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetChannelValueBooleanParse(theBytes []byte) (BACnetChannelValueBoolean, error) {
	return BACnetChannelValueBooleanParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetChannelValueBooleanParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetChannelValueBoolean, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (booleanValue)
	if pullErr := readBuffer.PullContext("booleanValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for booleanValue")
	}
_booleanValue, _booleanValueErr := BACnetApplicationTagParseWithBuffer(readBuffer)
	if _booleanValueErr != nil {
		return nil, errors.Wrap(_booleanValueErr, "Error parsing 'booleanValue' field of BACnetChannelValueBoolean")
	}
	booleanValue := _booleanValue.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("booleanValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for booleanValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueBoolean")
	}

	// Create a partially initialized instance
	_child := &_BACnetChannelValueBoolean{
		_BACnetChannelValue: &_BACnetChannelValue{
		},
		BooleanValue: booleanValue,
	}
	_child._BACnetChannelValue._BACnetChannelValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetChannelValueBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetChannelValueBoolean) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueBoolean")
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

		if popErr := writeBuffer.PopContext("BACnetChannelValueBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueBoolean")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetChannelValueBoolean) isBACnetChannelValueBoolean() bool {
	return true
}

func (m *_BACnetChannelValueBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



