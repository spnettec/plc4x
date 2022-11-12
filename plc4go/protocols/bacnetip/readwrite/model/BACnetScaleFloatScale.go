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


// BACnetScaleFloatScale is the corresponding interface of BACnetScaleFloatScale
type BACnetScaleFloatScale interface {
	utils.LengthAware
	utils.Serializable
	BACnetScale
	// GetFloatScale returns FloatScale (property field)
	GetFloatScale() BACnetContextTagReal
}

// BACnetScaleFloatScaleExactly can be used when we want exactly this type and not a type which fulfills BACnetScaleFloatScale.
// This is useful for switch cases.
type BACnetScaleFloatScaleExactly interface {
	BACnetScaleFloatScale
	isBACnetScaleFloatScale() bool
}

// _BACnetScaleFloatScale is the data-structure of this message
type _BACnetScaleFloatScale struct {
	*_BACnetScale
        FloatScale BACnetContextTagReal
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetScaleFloatScale) InitializeParent(parent BACnetScale , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetScaleFloatScale)  GetParent() BACnetScale {
	return m._BACnetScale
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetScaleFloatScale) GetFloatScale() BACnetContextTagReal {
	return m.FloatScale
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetScaleFloatScale factory function for _BACnetScaleFloatScale
func NewBACnetScaleFloatScale( floatScale BACnetContextTagReal , peekedTagHeader BACnetTagHeader ) *_BACnetScaleFloatScale {
	_result := &_BACnetScaleFloatScale{
		FloatScale: floatScale,
    	_BACnetScale: NewBACnetScale(peekedTagHeader),
	}
	_result._BACnetScale._BACnetScaleChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetScaleFloatScale(structType interface{}) BACnetScaleFloatScale {
    if casted, ok := structType.(BACnetScaleFloatScale); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetScaleFloatScale); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetScaleFloatScale) GetTypeName() string {
	return "BACnetScaleFloatScale"
}

func (m *_BACnetScaleFloatScale) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetScaleFloatScale) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (floatScale)
	lengthInBits += m.FloatScale.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetScaleFloatScale) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetScaleFloatScaleParse(theBytes []byte) (BACnetScaleFloatScale, error) {
	return BACnetScaleFloatScaleParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetScaleFloatScaleParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetScaleFloatScale, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetScaleFloatScale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetScaleFloatScale")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (floatScale)
	if pullErr := readBuffer.PullContext("floatScale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for floatScale")
	}
_floatScale, _floatScaleErr := BACnetContextTagParseWithBuffer(readBuffer , uint8( uint8(0) ) , BACnetDataType( BACnetDataType_REAL ) )
	if _floatScaleErr != nil {
		return nil, errors.Wrap(_floatScaleErr, "Error parsing 'floatScale' field of BACnetScaleFloatScale")
	}
	floatScale := _floatScale.(BACnetContextTagReal)
	if closeErr := readBuffer.CloseContext("floatScale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for floatScale")
	}

	if closeErr := readBuffer.CloseContext("BACnetScaleFloatScale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetScaleFloatScale")
	}

	// Create a partially initialized instance
	_child := &_BACnetScaleFloatScale{
		_BACnetScale: &_BACnetScale{
		},
		FloatScale: floatScale,
	}
	_child._BACnetScale._BACnetScaleChildRequirements = _child
	return _child, nil
}

func (m *_BACnetScaleFloatScale) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetScaleFloatScale) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetScaleFloatScale"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetScaleFloatScale")
		}

	// Simple Field (floatScale)
	if pushErr := writeBuffer.PushContext("floatScale"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for floatScale")
	}
	_floatScaleErr := writeBuffer.WriteSerializable(m.GetFloatScale())
	if popErr := writeBuffer.PopContext("floatScale"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for floatScale")
	}
	if _floatScaleErr != nil {
		return errors.Wrap(_floatScaleErr, "Error serializing 'floatScale' field")
	}

		if popErr := writeBuffer.PopContext("BACnetScaleFloatScale"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetScaleFloatScale")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetScaleFloatScale) isBACnetScaleFloatScale() bool {
	return true
}

func (m *_BACnetScaleFloatScale) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



