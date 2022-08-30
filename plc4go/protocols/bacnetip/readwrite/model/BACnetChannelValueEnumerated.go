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


// BACnetChannelValueEnumerated is the corresponding interface of BACnetChannelValueEnumerated
type BACnetChannelValueEnumerated interface {
	utils.LengthAware
	utils.Serializable
	BACnetChannelValue
	// GetEnumeratedValue returns EnumeratedValue (property field)
	GetEnumeratedValue() BACnetApplicationTagEnumerated
}

// BACnetChannelValueEnumeratedExactly can be used when we want exactly this type and not a type which fulfills BACnetChannelValueEnumerated.
// This is useful for switch cases.
type BACnetChannelValueEnumeratedExactly interface {
	BACnetChannelValueEnumerated
	isBACnetChannelValueEnumerated() bool
}

// _BACnetChannelValueEnumerated is the data-structure of this message
type _BACnetChannelValueEnumerated struct {
	*_BACnetChannelValue
        EnumeratedValue BACnetApplicationTagEnumerated
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetChannelValueEnumerated) InitializeParent(parent BACnetChannelValue , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetChannelValueEnumerated)  GetParent() BACnetChannelValue {
	return m._BACnetChannelValue
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetChannelValueEnumerated) GetEnumeratedValue() BACnetApplicationTagEnumerated {
	return m.EnumeratedValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetChannelValueEnumerated factory function for _BACnetChannelValueEnumerated
func NewBACnetChannelValueEnumerated( enumeratedValue BACnetApplicationTagEnumerated , peekedTagHeader BACnetTagHeader ) *_BACnetChannelValueEnumerated {
	_result := &_BACnetChannelValueEnumerated{
		EnumeratedValue: enumeratedValue,
    	_BACnetChannelValue: NewBACnetChannelValue(peekedTagHeader),
	}
	_result._BACnetChannelValue._BACnetChannelValueChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetChannelValueEnumerated(structType interface{}) BACnetChannelValueEnumerated {
    if casted, ok := structType.(BACnetChannelValueEnumerated); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetChannelValueEnumerated); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetChannelValueEnumerated) GetTypeName() string {
	return "BACnetChannelValueEnumerated"
}

func (m *_BACnetChannelValueEnumerated) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetChannelValueEnumerated) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (enumeratedValue)
	lengthInBits += m.EnumeratedValue.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetChannelValueEnumerated) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetChannelValueEnumeratedParse(readBuffer utils.ReadBuffer) (BACnetChannelValueEnumerated, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetChannelValueEnumerated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetChannelValueEnumerated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (enumeratedValue)
	if pullErr := readBuffer.PullContext("enumeratedValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for enumeratedValue")
	}
_enumeratedValue, _enumeratedValueErr := BACnetApplicationTagParse(readBuffer)
	if _enumeratedValueErr != nil {
		return nil, errors.Wrap(_enumeratedValueErr, "Error parsing 'enumeratedValue' field of BACnetChannelValueEnumerated")
	}
	enumeratedValue := _enumeratedValue.(BACnetApplicationTagEnumerated)
	if closeErr := readBuffer.CloseContext("enumeratedValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for enumeratedValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetChannelValueEnumerated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetChannelValueEnumerated")
	}

	// Create a partially initialized instance
	_child := &_BACnetChannelValueEnumerated{
		_BACnetChannelValue: &_BACnetChannelValue{
		},
		EnumeratedValue: enumeratedValue,
	}
	_child._BACnetChannelValue._BACnetChannelValueChildRequirements = _child
	return _child, nil
}

func (m *_BACnetChannelValueEnumerated) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetChannelValueEnumerated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetChannelValueEnumerated")
		}

	// Simple Field (enumeratedValue)
	if pushErr := writeBuffer.PushContext("enumeratedValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for enumeratedValue")
	}
	_enumeratedValueErr := writeBuffer.WriteSerializable(m.GetEnumeratedValue())
	if popErr := writeBuffer.PopContext("enumeratedValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for enumeratedValue")
	}
	if _enumeratedValueErr != nil {
		return errors.Wrap(_enumeratedValueErr, "Error serializing 'enumeratedValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetChannelValueEnumerated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetChannelValueEnumerated")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetChannelValueEnumerated) isBACnetChannelValueEnumerated() bool {
	return true
}

func (m *_BACnetChannelValueEnumerated) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



