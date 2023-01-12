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


// BACnetPropertyStatesLightningTransition is the corresponding interface of BACnetPropertyStatesLightningTransition
type BACnetPropertyStatesLightningTransition interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetLightningTransition returns LightningTransition (property field)
	GetLightningTransition() BACnetLightingTransitionTagged
}

// BACnetPropertyStatesLightningTransitionExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesLightningTransition.
// This is useful for switch cases.
type BACnetPropertyStatesLightningTransitionExactly interface {
	BACnetPropertyStatesLightningTransition
	isBACnetPropertyStatesLightningTransition() bool
}

// _BACnetPropertyStatesLightningTransition is the data-structure of this message
type _BACnetPropertyStatesLightningTransition struct {
	*_BACnetPropertyStates
        LightningTransition BACnetLightingTransitionTagged
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesLightningTransition) InitializeParent(parent BACnetPropertyStates , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesLightningTransition)  GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesLightningTransition) GetLightningTransition() BACnetLightingTransitionTagged {
	return m.LightningTransition
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetPropertyStatesLightningTransition factory function for _BACnetPropertyStatesLightningTransition
func NewBACnetPropertyStatesLightningTransition( lightningTransition BACnetLightingTransitionTagged , peekedTagHeader BACnetTagHeader ) *_BACnetPropertyStatesLightningTransition {
	_result := &_BACnetPropertyStatesLightningTransition{
		LightningTransition: lightningTransition,
    	_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesLightningTransition(structType interface{}) BACnetPropertyStatesLightningTransition {
    if casted, ok := structType.(BACnetPropertyStatesLightningTransition); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesLightningTransition); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesLightningTransition) GetTypeName() string {
	return "BACnetPropertyStatesLightningTransition"
}

func (m *_BACnetPropertyStatesLightningTransition) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesLightningTransition) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (lightningTransition)
	lengthInBits += m.LightningTransition.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetPropertyStatesLightningTransition) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesLightningTransitionParse(theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesLightningTransition, error) {
	return BACnetPropertyStatesLightningTransitionParseWithBuffer(utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesLightningTransitionParseWithBuffer(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesLightningTransition, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesLightningTransition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesLightningTransition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (lightningTransition)
	if pullErr := readBuffer.PullContext("lightningTransition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for lightningTransition")
	}
_lightningTransition, _lightningTransitionErr := BACnetLightingTransitionTaggedParseWithBuffer(readBuffer , uint8( peekedTagNumber ) , TagClass( TagClass_CONTEXT_SPECIFIC_TAGS ) )
	if _lightningTransitionErr != nil {
		return nil, errors.Wrap(_lightningTransitionErr, "Error parsing 'lightningTransition' field of BACnetPropertyStatesLightningTransition")
	}
	lightningTransition := _lightningTransition.(BACnetLightingTransitionTagged)
	if closeErr := readBuffer.CloseContext("lightningTransition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for lightningTransition")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesLightningTransition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesLightningTransition")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesLightningTransition{
		_BACnetPropertyStates: &_BACnetPropertyStates{
		},
		LightningTransition: lightningTransition,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesLightningTransition) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesLightningTransition) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesLightningTransition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesLightningTransition")
		}

	// Simple Field (lightningTransition)
	if pushErr := writeBuffer.PushContext("lightningTransition"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for lightningTransition")
	}
	_lightningTransitionErr := writeBuffer.WriteSerializable(m.GetLightningTransition())
	if popErr := writeBuffer.PopContext("lightningTransition"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for lightningTransition")
	}
	if _lightningTransitionErr != nil {
		return errors.Wrap(_lightningTransitionErr, "Error serializing 'lightningTransition' field")
	}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesLightningTransition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesLightningTransition")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetPropertyStatesLightningTransition) isBACnetPropertyStatesLightningTransition() bool {
	return true
}

func (m *_BACnetPropertyStatesLightningTransition) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



