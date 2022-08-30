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


// BACnetPropertyStatesZoneOccupanyState is the corresponding interface of BACnetPropertyStatesZoneOccupanyState
type BACnetPropertyStatesZoneOccupanyState interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetZoneOccupanyState returns ZoneOccupanyState (property field)
	GetZoneOccupanyState() BACnetAccessZoneOccupancyStateTagged
}

// BACnetPropertyStatesZoneOccupanyStateExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesZoneOccupanyState.
// This is useful for switch cases.
type BACnetPropertyStatesZoneOccupanyStateExactly interface {
	BACnetPropertyStatesZoneOccupanyState
	isBACnetPropertyStatesZoneOccupanyState() bool
}

// _BACnetPropertyStatesZoneOccupanyState is the data-structure of this message
type _BACnetPropertyStatesZoneOccupanyState struct {
	*_BACnetPropertyStates
        ZoneOccupanyState BACnetAccessZoneOccupancyStateTagged
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesZoneOccupanyState) InitializeParent(parent BACnetPropertyStates , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesZoneOccupanyState)  GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesZoneOccupanyState) GetZoneOccupanyState() BACnetAccessZoneOccupancyStateTagged {
	return m.ZoneOccupanyState
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetPropertyStatesZoneOccupanyState factory function for _BACnetPropertyStatesZoneOccupanyState
func NewBACnetPropertyStatesZoneOccupanyState( zoneOccupanyState BACnetAccessZoneOccupancyStateTagged , peekedTagHeader BACnetTagHeader ) *_BACnetPropertyStatesZoneOccupanyState {
	_result := &_BACnetPropertyStatesZoneOccupanyState{
		ZoneOccupanyState: zoneOccupanyState,
    	_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesZoneOccupanyState(structType interface{}) BACnetPropertyStatesZoneOccupanyState {
    if casted, ok := structType.(BACnetPropertyStatesZoneOccupanyState); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesZoneOccupanyState); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesZoneOccupanyState) GetTypeName() string {
	return "BACnetPropertyStatesZoneOccupanyState"
}

func (m *_BACnetPropertyStatesZoneOccupanyState) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesZoneOccupanyState) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (zoneOccupanyState)
	lengthInBits += m.ZoneOccupanyState.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetPropertyStatesZoneOccupanyState) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesZoneOccupanyStateParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesZoneOccupanyState, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesZoneOccupanyState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesZoneOccupanyState")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneOccupanyState)
	if pullErr := readBuffer.PullContext("zoneOccupanyState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for zoneOccupanyState")
	}
_zoneOccupanyState, _zoneOccupanyStateErr := BACnetAccessZoneOccupancyStateTaggedParse(readBuffer , uint8( peekedTagNumber ) , TagClass( TagClass_CONTEXT_SPECIFIC_TAGS ) )
	if _zoneOccupanyStateErr != nil {
		return nil, errors.Wrap(_zoneOccupanyStateErr, "Error parsing 'zoneOccupanyState' field of BACnetPropertyStatesZoneOccupanyState")
	}
	zoneOccupanyState := _zoneOccupanyState.(BACnetAccessZoneOccupancyStateTagged)
	if closeErr := readBuffer.CloseContext("zoneOccupanyState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for zoneOccupanyState")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesZoneOccupanyState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesZoneOccupanyState")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesZoneOccupanyState{
		_BACnetPropertyStates: &_BACnetPropertyStates{
		},
		ZoneOccupanyState: zoneOccupanyState,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesZoneOccupanyState) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesZoneOccupanyState"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesZoneOccupanyState")
		}

	// Simple Field (zoneOccupanyState)
	if pushErr := writeBuffer.PushContext("zoneOccupanyState"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for zoneOccupanyState")
	}
	_zoneOccupanyStateErr := writeBuffer.WriteSerializable(m.GetZoneOccupanyState())
	if popErr := writeBuffer.PopContext("zoneOccupanyState"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for zoneOccupanyState")
	}
	if _zoneOccupanyStateErr != nil {
		return errors.Wrap(_zoneOccupanyStateErr, "Error serializing 'zoneOccupanyState' field")
	}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesZoneOccupanyState"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesZoneOccupanyState")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetPropertyStatesZoneOccupanyState) isBACnetPropertyStatesZoneOccupanyState() bool {
	return true
}

func (m *_BACnetPropertyStatesZoneOccupanyState) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



