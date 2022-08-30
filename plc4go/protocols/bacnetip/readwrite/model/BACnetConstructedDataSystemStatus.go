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


// BACnetConstructedDataSystemStatus is the corresponding interface of BACnetConstructedDataSystemStatus
type BACnetConstructedDataSystemStatus interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetSystemStatus returns SystemStatus (property field)
	GetSystemStatus() BACnetDeviceStatusTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetDeviceStatusTagged
}

// BACnetConstructedDataSystemStatusExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataSystemStatus.
// This is useful for switch cases.
type BACnetConstructedDataSystemStatusExactly interface {
	BACnetConstructedDataSystemStatus
	isBACnetConstructedDataSystemStatus() bool
}

// _BACnetConstructedDataSystemStatus is the data-structure of this message
type _BACnetConstructedDataSystemStatus struct {
	*_BACnetConstructedData
        SystemStatus BACnetDeviceStatusTagged
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataSystemStatus)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataSystemStatus)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_SYSTEM_STATUS}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataSystemStatus) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataSystemStatus)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataSystemStatus) GetSystemStatus() BACnetDeviceStatusTagged {
	return m.SystemStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataSystemStatus) GetActualValue() BACnetDeviceStatusTagged {
	return CastBACnetDeviceStatusTagged(m.GetSystemStatus())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataSystemStatus factory function for _BACnetConstructedDataSystemStatus
func NewBACnetConstructedDataSystemStatus( systemStatus BACnetDeviceStatusTagged , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataSystemStatus {
	_result := &_BACnetConstructedDataSystemStatus{
		SystemStatus: systemStatus,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataSystemStatus(structType interface{}) BACnetConstructedDataSystemStatus {
    if casted, ok := structType.(BACnetConstructedDataSystemStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataSystemStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataSystemStatus) GetTypeName() string {
	return "BACnetConstructedDataSystemStatus"
}

func (m *_BACnetConstructedDataSystemStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataSystemStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (systemStatus)
	lengthInBits += m.SystemStatus.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataSystemStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataSystemStatusParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataSystemStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataSystemStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataSystemStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (systemStatus)
	if pullErr := readBuffer.PullContext("systemStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for systemStatus")
	}
_systemStatus, _systemStatusErr := BACnetDeviceStatusTaggedParse(readBuffer , uint8( uint8(0) ) , TagClass( TagClass_APPLICATION_TAGS ) )
	if _systemStatusErr != nil {
		return nil, errors.Wrap(_systemStatusErr, "Error parsing 'systemStatus' field of BACnetConstructedDataSystemStatus")
	}
	systemStatus := _systemStatus.(BACnetDeviceStatusTagged)
	if closeErr := readBuffer.CloseContext("systemStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for systemStatus")
	}

	// Virtual field
	_actualValue := systemStatus
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataSystemStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataSystemStatus")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataSystemStatus{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		SystemStatus: systemStatus,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataSystemStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataSystemStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataSystemStatus")
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
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataSystemStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataSystemStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataSystemStatus) isBACnetConstructedDataSystemStatus() bool {
	return true
}

func (m *_BACnetConstructedDataSystemStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



