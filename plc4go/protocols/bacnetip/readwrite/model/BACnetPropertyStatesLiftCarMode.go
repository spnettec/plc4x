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


// BACnetPropertyStatesLiftCarMode is the corresponding interface of BACnetPropertyStatesLiftCarMode
type BACnetPropertyStatesLiftCarMode interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLiftCarMode returns LiftCarMode (property field)
	GetLiftCarMode() BACnetLiftCarModeTagged
}

// BACnetPropertyStatesLiftCarModeExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesLiftCarMode.
// This is useful for switch cases.
type BACnetPropertyStatesLiftCarModeExactly interface {
	BACnetPropertyStatesLiftCarMode
	isBACnetPropertyStatesLiftCarMode() bool
}

// _BACnetPropertyStatesLiftCarMode is the data-structure of this message
type _BACnetPropertyStatesLiftCarMode struct {
	*_BACnetPropertyStates
        LiftCarMode BACnetLiftCarModeTagged
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLiftCarMode) InitializeParent(parent BACnetPropertyStates , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesLiftCarMode)  GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLiftCarMode) GetLiftCarMode() BACnetLiftCarModeTagged {
	return m.LiftCarMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetPropertyStatesLiftCarMode factory function for _BACnetPropertyStatesLiftCarMode
func NewBACnetPropertyStatesLiftCarMode( liftCarMode BACnetLiftCarModeTagged , peekedTagHeader BACnetTagHeader ) *_BACnetPropertyStatesLiftCarMode {
	_result := &_BACnetPropertyStatesLiftCarMode{
		LiftCarMode: liftCarMode,
    	_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLiftCarMode(structType interface{}) BACnetPropertyStatesLiftCarMode {
    if casted, ok := structType.(BACnetPropertyStatesLiftCarMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLiftCarMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLiftCarMode) GetTypeName() string {
	return "BACnetPropertyStatesLiftCarMode"
}

func (m *_BACnetPropertyStatesLiftCarMode) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesLiftCarMode) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (liftCarMode)
	lengthInBits += m.LiftCarMode.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetPropertyStatesLiftCarMode) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesLiftCarModeParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesLiftCarMode, error) {
	return BACnetPropertyStatesLiftCarModeParseWithBuffer(utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesLiftCarModeParseWithBuffer(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesLiftCarMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLiftCarMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLiftCarMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (liftCarMode)
	if pullErr := readBuffer.PullContext("liftCarMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for liftCarMode")
	}
_liftCarMode, _liftCarModeErr := BACnetLiftCarModeTaggedParseWithBuffer(readBuffer , uint8( peekedTagNumber ) , TagClass( TagClass_CONTEXT_SPECIFIC_TAGS ) )
	if _liftCarModeErr != nil {
		return nil, errors.Wrap(_liftCarModeErr, "Error parsing 'liftCarMode' field of BACnetPropertyStatesLiftCarMode")
	}
	liftCarMode := _liftCarMode.(BACnetLiftCarModeTagged)
	if closeErr := readBuffer.CloseContext("liftCarMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for liftCarMode")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLiftCarMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLiftCarMode")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesLiftCarMode{
		_BACnetPropertyStates: &_BACnetPropertyStates{
		},
		LiftCarMode: liftCarMode,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesLiftCarMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLiftCarMode) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLiftCarMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLiftCarMode")
		}

	// Simple Field (liftCarMode)
	if pushErr := writeBuffer.PushContext("liftCarMode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for liftCarMode")
	}
	_liftCarModeErr := writeBuffer.WriteSerializable(m.GetLiftCarMode())
	if popErr := writeBuffer.PopContext("liftCarMode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for liftCarMode")
	}
	if _liftCarModeErr != nil {
		return errors.Wrap(_liftCarModeErr, "Error serializing 'liftCarMode' field")
	}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLiftCarMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLiftCarMode")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetPropertyStatesLiftCarMode) isBACnetPropertyStatesLiftCarMode() bool {
	return true
}

func (m *_BACnetPropertyStatesLiftCarMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



