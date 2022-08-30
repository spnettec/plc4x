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


// SecurityDataZoneName is the corresponding interface of SecurityDataZoneName
type SecurityDataZoneName interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetZoneNumber returns ZoneNumber (property field)
	GetZoneNumber() uint8
	// GetZoneName returns ZoneName (property field)
	GetZoneName() string
}

// SecurityDataZoneNameExactly can be used when we want exactly this type and not a type which fulfills SecurityDataZoneName.
// This is useful for switch cases.
type SecurityDataZoneNameExactly interface {
	SecurityDataZoneName
	isSecurityDataZoneName() bool
}

// _SecurityDataZoneName is the data-structure of this message
type _SecurityDataZoneName struct {
	*_SecurityData
        ZoneNumber uint8
        ZoneName string
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataZoneName) InitializeParent(parent SecurityData , commandTypeContainer SecurityCommandTypeContainer , argument byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataZoneName)  GetParent() SecurityData {
	return m._SecurityData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataZoneName) GetZoneNumber() uint8 {
	return m.ZoneNumber
}

func (m *_SecurityDataZoneName) GetZoneName() string {
	return m.ZoneName
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewSecurityDataZoneName factory function for _SecurityDataZoneName
func NewSecurityDataZoneName( zoneNumber uint8 , zoneName string , commandTypeContainer SecurityCommandTypeContainer , argument byte ) *_SecurityDataZoneName {
	_result := &_SecurityDataZoneName{
		ZoneNumber: zoneNumber,
		ZoneName: zoneName,
    	_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataZoneName(structType interface{}) SecurityDataZoneName {
    if casted, ok := structType.(SecurityDataZoneName); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataZoneName); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataZoneName) GetTypeName() string {
	return "SecurityDataZoneName"
}

func (m *_SecurityDataZoneName) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataZoneName) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (zoneNumber)
	lengthInBits += 8;

	// Simple field (zoneName)
	lengthInBits += 88;

	return lengthInBits
}


func (m *_SecurityDataZoneName) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataZoneNameParse(readBuffer utils.ReadBuffer) (SecurityDataZoneName, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataZoneName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataZoneName")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneNumber)
_zoneNumber, _zoneNumberErr := readBuffer.ReadUint8("zoneNumber", 8)
	if _zoneNumberErr != nil {
		return nil, errors.Wrap(_zoneNumberErr, "Error parsing 'zoneNumber' field of SecurityDataZoneName")
	}
	zoneNumber := _zoneNumber

	// Simple Field (zoneName)
_zoneName, _zoneNameErr := readBuffer.ReadString("zoneName", uint32(88))
	if _zoneNameErr != nil {
		return nil, errors.Wrap(_zoneNameErr, "Error parsing 'zoneName' field of SecurityDataZoneName")
	}
	zoneName := _zoneName

	if closeErr := readBuffer.CloseContext("SecurityDataZoneName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataZoneName")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataZoneName{
		_SecurityData: &_SecurityData{
		},
		ZoneNumber: zoneNumber,
		ZoneName: zoneName,
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataZoneName) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataZoneName"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataZoneName")
		}

	// Simple Field (zoneNumber)
	zoneNumber := uint8(m.GetZoneNumber())
	_zoneNumberErr := writeBuffer.WriteUint8("zoneNumber", 8, (zoneNumber))
	if _zoneNumberErr != nil {
		return errors.Wrap(_zoneNumberErr, "Error serializing 'zoneNumber' field")
	}

	// Simple Field (zoneName)
	zoneName := string(m.GetZoneName())
	_zoneNameErr := writeBuffer.WriteString("zoneName", uint32(88), "UTF-8", (zoneName))
	if _zoneNameErr != nil {
		return errors.Wrap(_zoneNameErr, "Error serializing 'zoneName' field")
	}

		if popErr := writeBuffer.PopContext("SecurityDataZoneName"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataZoneName")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_SecurityDataZoneName) isSecurityDataZoneName() bool {
	return true
}

func (m *_SecurityDataZoneName) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



