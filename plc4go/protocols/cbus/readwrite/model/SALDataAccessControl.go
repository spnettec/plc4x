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


// SALDataAccessControl is the corresponding interface of SALDataAccessControl
type SALDataAccessControl interface {
	utils.LengthAware
	utils.Serializable
	SALData
	// GetAccessControlData returns AccessControlData (property field)
	GetAccessControlData() AccessControlData
}

// SALDataAccessControlExactly can be used when we want exactly this type and not a type which fulfills SALDataAccessControl.
// This is useful for switch cases.
type SALDataAccessControlExactly interface {
	SALDataAccessControl
	isSALDataAccessControl() bool
}

// _SALDataAccessControl is the data-structure of this message
type _SALDataAccessControl struct {
	*_SALData
        AccessControlData AccessControlData
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataAccessControl)  GetApplicationId() ApplicationId {
return ApplicationId_ACCESS_CONTROL}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataAccessControl) InitializeParent(parent SALData , salData SALData ) {	m.SalData = salData
}

func (m *_SALDataAccessControl)  GetParent() SALData {
	return m._SALData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataAccessControl) GetAccessControlData() AccessControlData {
	return m.AccessControlData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewSALDataAccessControl factory function for _SALDataAccessControl
func NewSALDataAccessControl( accessControlData AccessControlData , salData SALData ) *_SALDataAccessControl {
	_result := &_SALDataAccessControl{
		AccessControlData: accessControlData,
    	_SALData: NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataAccessControl(structType interface{}) SALDataAccessControl {
    if casted, ok := structType.(SALDataAccessControl); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataAccessControl); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataAccessControl) GetTypeName() string {
	return "SALDataAccessControl"
}

func (m *_SALDataAccessControl) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SALDataAccessControl) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (accessControlData)
	lengthInBits += m.AccessControlData.GetLengthInBits()

	return lengthInBits
}


func (m *_SALDataAccessControl) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALDataAccessControlParse(readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataAccessControl, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataAccessControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataAccessControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (accessControlData)
	if pullErr := readBuffer.PullContext("accessControlData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessControlData")
	}
_accessControlData, _accessControlDataErr := AccessControlDataParse(readBuffer)
	if _accessControlDataErr != nil {
		return nil, errors.Wrap(_accessControlDataErr, "Error parsing 'accessControlData' field of SALDataAccessControl")
	}
	accessControlData := _accessControlData.(AccessControlData)
	if closeErr := readBuffer.CloseContext("accessControlData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessControlData")
	}

	if closeErr := readBuffer.CloseContext("SALDataAccessControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataAccessControl")
	}

	// Create a partially initialized instance
	_child := &_SALDataAccessControl{
		_SALData: &_SALData{
		},
		AccessControlData: accessControlData,
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataAccessControl) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataAccessControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataAccessControl")
		}

	// Simple Field (accessControlData)
	if pushErr := writeBuffer.PushContext("accessControlData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for accessControlData")
	}
	_accessControlDataErr := writeBuffer.WriteSerializable(m.GetAccessControlData())
	if popErr := writeBuffer.PopContext("accessControlData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for accessControlData")
	}
	if _accessControlDataErr != nil {
		return errors.Wrap(_accessControlDataErr, "Error serializing 'accessControlData' field")
	}

		if popErr := writeBuffer.PopContext("SALDataAccessControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataAccessControl")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_SALDataAccessControl) isSALDataAccessControl() bool {
	return true
}

func (m *_SALDataAccessControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



