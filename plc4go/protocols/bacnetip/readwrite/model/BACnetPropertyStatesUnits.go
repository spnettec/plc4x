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


// BACnetPropertyStatesUnits is the corresponding interface of BACnetPropertyStatesUnits
type BACnetPropertyStatesUnits interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetUnits returns Units (property field)
	GetUnits() BACnetEngineeringUnitsTagged
}

// BACnetPropertyStatesUnitsExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesUnits.
// This is useful for switch cases.
type BACnetPropertyStatesUnitsExactly interface {
	BACnetPropertyStatesUnits
	isBACnetPropertyStatesUnits() bool
}

// _BACnetPropertyStatesUnits is the data-structure of this message
type _BACnetPropertyStatesUnits struct {
	*_BACnetPropertyStates
        Units BACnetEngineeringUnitsTagged
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesUnits) InitializeParent(parent BACnetPropertyStates , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesUnits)  GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesUnits) GetUnits() BACnetEngineeringUnitsTagged {
	return m.Units
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetPropertyStatesUnits factory function for _BACnetPropertyStatesUnits
func NewBACnetPropertyStatesUnits( units BACnetEngineeringUnitsTagged , peekedTagHeader BACnetTagHeader ) *_BACnetPropertyStatesUnits {
	_result := &_BACnetPropertyStatesUnits{
		Units: units,
    	_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesUnits(structType interface{}) BACnetPropertyStatesUnits {
    if casted, ok := structType.(BACnetPropertyStatesUnits); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesUnits); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesUnits) GetTypeName() string {
	return "BACnetPropertyStatesUnits"
}

func (m *_BACnetPropertyStatesUnits) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesUnits) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (units)
	lengthInBits += m.Units.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetPropertyStatesUnits) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesUnitsParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesUnits, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesUnits"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesUnits")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (units)
	if pullErr := readBuffer.PullContext("units"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for units")
	}
_units, _unitsErr := BACnetEngineeringUnitsTaggedParse(readBuffer , uint8( peekedTagNumber ) , TagClass( TagClass_CONTEXT_SPECIFIC_TAGS ) )
	if _unitsErr != nil {
		return nil, errors.Wrap(_unitsErr, "Error parsing 'units' field of BACnetPropertyStatesUnits")
	}
	units := _units.(BACnetEngineeringUnitsTagged)
	if closeErr := readBuffer.CloseContext("units"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for units")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesUnits"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesUnits")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesUnits{
		_BACnetPropertyStates: &_BACnetPropertyStates{
		},
		Units: units,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesUnits) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesUnits"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesUnits")
		}

	// Simple Field (units)
	if pushErr := writeBuffer.PushContext("units"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for units")
	}
	_unitsErr := writeBuffer.WriteSerializable(m.GetUnits())
	if popErr := writeBuffer.PopContext("units"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for units")
	}
	if _unitsErr != nil {
		return errors.Wrap(_unitsErr, "Error serializing 'units' field")
	}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesUnits"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesUnits")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetPropertyStatesUnits) isBACnetPropertyStatesUnits() bool {
	return true
}

func (m *_BACnetPropertyStatesUnits) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



