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


// BACnetEventParameterNone is the corresponding interface of BACnetEventParameterNone
type BACnetEventParameterNone interface {
	utils.LengthAware
	utils.Serializable
	BACnetEventParameter
	// GetNone returns None (property field)
	GetNone() BACnetContextTagNull
}

// BACnetEventParameterNoneExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameterNone.
// This is useful for switch cases.
type BACnetEventParameterNoneExactly interface {
	BACnetEventParameterNone
	isBACnetEventParameterNone() bool
}

// _BACnetEventParameterNone is the data-structure of this message
type _BACnetEventParameterNone struct {
	*_BACnetEventParameter
        None BACnetContextTagNull
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetEventParameterNone) InitializeParent(parent BACnetEventParameter , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetEventParameterNone)  GetParent() BACnetEventParameter {
	return m._BACnetEventParameter
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameterNone) GetNone() BACnetContextTagNull {
	return m.None
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetEventParameterNone factory function for _BACnetEventParameterNone
func NewBACnetEventParameterNone( none BACnetContextTagNull , peekedTagHeader BACnetTagHeader ) *_BACnetEventParameterNone {
	_result := &_BACnetEventParameterNone{
		None: none,
    	_BACnetEventParameter: NewBACnetEventParameter(peekedTagHeader),
	}
	_result._BACnetEventParameter._BACnetEventParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameterNone(structType interface{}) BACnetEventParameterNone {
    if casted, ok := structType.(BACnetEventParameterNone); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameterNone); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameterNone) GetTypeName() string {
	return "BACnetEventParameterNone"
}

func (m *_BACnetEventParameterNone) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventParameterNone) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (none)
	lengthInBits += m.None.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetEventParameterNone) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventParameterNoneParse(theBytes []byte) (BACnetEventParameterNone, error) {
	return BACnetEventParameterNoneParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventParameterNoneParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetEventParameterNone, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameterNone"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameterNone")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (none)
	if pullErr := readBuffer.PullContext("none"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for none")
	}
_none, _noneErr := BACnetContextTagParseWithBuffer(readBuffer , uint8( uint8(20) ) , BACnetDataType( BACnetDataType_NULL ) )
	if _noneErr != nil {
		return nil, errors.Wrap(_noneErr, "Error parsing 'none' field of BACnetEventParameterNone")
	}
	none := _none.(BACnetContextTagNull)
	if closeErr := readBuffer.CloseContext("none"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for none")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventParameterNone"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameterNone")
	}

	// Create a partially initialized instance
	_child := &_BACnetEventParameterNone{
		_BACnetEventParameter: &_BACnetEventParameter{
		},
		None: none,
	}
	_child._BACnetEventParameter._BACnetEventParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetEventParameterNone) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventParameterNone) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventParameterNone"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventParameterNone")
		}

	// Simple Field (none)
	if pushErr := writeBuffer.PushContext("none"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for none")
	}
	_noneErr := writeBuffer.WriteSerializable(m.GetNone())
	if popErr := writeBuffer.PopContext("none"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for none")
	}
	if _noneErr != nil {
		return errors.Wrap(_noneErr, "Error serializing 'none' field")
	}

		if popErr := writeBuffer.PopContext("BACnetEventParameterNone"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventParameterNone")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetEventParameterNone) isBACnetEventParameterNone() bool {
	return true
}

func (m *_BACnetEventParameterNone) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



