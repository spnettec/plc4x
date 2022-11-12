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


// SecurityDataZoneShort is the corresponding interface of SecurityDataZoneShort
type SecurityDataZoneShort interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetZoneNumber returns ZoneNumber (property field)
	GetZoneNumber() uint8
}

// SecurityDataZoneShortExactly can be used when we want exactly this type and not a type which fulfills SecurityDataZoneShort.
// This is useful for switch cases.
type SecurityDataZoneShortExactly interface {
	SecurityDataZoneShort
	isSecurityDataZoneShort() bool
}

// _SecurityDataZoneShort is the data-structure of this message
type _SecurityDataZoneShort struct {
	*_SecurityData
        ZoneNumber uint8
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataZoneShort) InitializeParent(parent SecurityData , commandTypeContainer SecurityCommandTypeContainer , argument byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataZoneShort)  GetParent() SecurityData {
	return m._SecurityData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataZoneShort) GetZoneNumber() uint8 {
	return m.ZoneNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewSecurityDataZoneShort factory function for _SecurityDataZoneShort
func NewSecurityDataZoneShort( zoneNumber uint8 , commandTypeContainer SecurityCommandTypeContainer , argument byte ) *_SecurityDataZoneShort {
	_result := &_SecurityDataZoneShort{
		ZoneNumber: zoneNumber,
    	_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataZoneShort(structType interface{}) SecurityDataZoneShort {
    if casted, ok := structType.(SecurityDataZoneShort); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataZoneShort); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataZoneShort) GetTypeName() string {
	return "SecurityDataZoneShort"
}

func (m *_SecurityDataZoneShort) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataZoneShort) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (zoneNumber)
	lengthInBits += 8;

	return lengthInBits
}


func (m *_SecurityDataZoneShort) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataZoneShortParse(theBytes []byte) (SecurityDataZoneShort, error) {
	return SecurityDataZoneShortParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataZoneShortParseWithBuffer(readBuffer utils.ReadBuffer) (SecurityDataZoneShort, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataZoneShort"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataZoneShort")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneNumber)
_zoneNumber, _zoneNumberErr := readBuffer.ReadUint8("zoneNumber", 8)
	if _zoneNumberErr != nil {
		return nil, errors.Wrap(_zoneNumberErr, "Error parsing 'zoneNumber' field of SecurityDataZoneShort")
	}
	zoneNumber := _zoneNumber

	if closeErr := readBuffer.CloseContext("SecurityDataZoneShort"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataZoneShort")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataZoneShort{
		_SecurityData: &_SecurityData{
		},
		ZoneNumber: zoneNumber,
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataZoneShort) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataZoneShort) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataZoneShort"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataZoneShort")
		}

	// Simple Field (zoneNumber)
	zoneNumber := uint8(m.GetZoneNumber())
	_zoneNumberErr := writeBuffer.WriteUint8("zoneNumber", 8, (zoneNumber))
	if _zoneNumberErr != nil {
		return errors.Wrap(_zoneNumberErr, "Error serializing 'zoneNumber' field")
	}

		if popErr := writeBuffer.PopContext("SecurityDataZoneShort"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataZoneShort")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_SecurityDataZoneShort) isSecurityDataZoneShort() bool {
	return true
}

func (m *_SecurityDataZoneShort) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



