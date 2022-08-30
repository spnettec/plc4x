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


// BACnetPropertyStatesSystemStatus is the corresponding interface of BACnetPropertyStatesSystemStatus
type BACnetPropertyStatesSystemStatus interface {
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetSystemStatus returns SystemStatus (property field)
	GetSystemStatus() BACnetDeviceStatusTagged
}

// BACnetPropertyStatesSystemStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesSystemStatus.
// This is useful for switch cases.
type BACnetPropertyStatesSystemStatusExactly interface {
	BACnetPropertyStatesSystemStatus
	isBACnetPropertyStatesSystemStatus() bool
}

// _BACnetPropertyStatesSystemStatus is the data-structure of this message
type _BACnetPropertyStatesSystemStatus struct {
	*_BACnetPropertyStates
        SystemStatus BACnetDeviceStatusTagged
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesSystemStatus) InitializeParent(parent BACnetPropertyStates , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesSystemStatus)  GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesSystemStatus) GetSystemStatus() BACnetDeviceStatusTagged {
	return m.SystemStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetPropertyStatesSystemStatus factory function for _BACnetPropertyStatesSystemStatus
func NewBACnetPropertyStatesSystemStatus( systemStatus BACnetDeviceStatusTagged , peekedTagHeader BACnetTagHeader ) *_BACnetPropertyStatesSystemStatus {
	_result := &_BACnetPropertyStatesSystemStatus{
		SystemStatus: systemStatus,
    	_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesSystemStatus(structType interface{}) BACnetPropertyStatesSystemStatus {
    if casted, ok := structType.(BACnetPropertyStatesSystemStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesSystemStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesSystemStatus) GetTypeName() string {
	return "BACnetPropertyStatesSystemStatus"
}

func (m *_BACnetPropertyStatesSystemStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesSystemStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (systemStatus)
	lengthInBits += m.SystemStatus.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetPropertyStatesSystemStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesSystemStatusParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesSystemStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesSystemStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesSystemStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (systemStatus)
	if pullErr := readBuffer.PullContext("systemStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for systemStatus")
	}
_systemStatus, _systemStatusErr := BACnetDeviceStatusTaggedParse(readBuffer , uint8( peekedTagNumber ) , TagClass( TagClass_CONTEXT_SPECIFIC_TAGS ) )
	if _systemStatusErr != nil {
		return nil, errors.Wrap(_systemStatusErr, "Error parsing 'systemStatus' field of BACnetPropertyStatesSystemStatus")
	}
	systemStatus := _systemStatus.(BACnetDeviceStatusTagged)
	if closeErr := readBuffer.CloseContext("systemStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for systemStatus")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesSystemStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesSystemStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesSystemStatus{
		_BACnetPropertyStates: &_BACnetPropertyStates{
		},
		SystemStatus: systemStatus,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesSystemStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesSystemStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesSystemStatus")
		}

	// Simple Field (systemStatus)
	if pushErr := writeBuffer.PushContext("systemStatus"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for systemStatus")
	}
	_systemStatusErr := writeBuffer.WriteSerializable(m.GetSystemStatus())
	if popErr := writeBuffer.PopContext("systemStatus"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for systemStatus")
	}
	if _systemStatusErr != nil {
		return errors.Wrap(_systemStatusErr, "Error serializing 'systemStatus' field")
	}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesSystemStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesSystemStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetPropertyStatesSystemStatus) isBACnetPropertyStatesSystemStatus() bool {
	return true
}

func (m *_BACnetPropertyStatesSystemStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



