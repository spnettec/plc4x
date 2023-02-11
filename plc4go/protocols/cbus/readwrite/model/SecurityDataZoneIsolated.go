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
	"context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// SecurityDataZoneIsolated is the corresponding interface of SecurityDataZoneIsolated
type SecurityDataZoneIsolated interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetZoneNumber returns ZoneNumber (property field)
	GetZoneNumber() uint8
}

// SecurityDataZoneIsolatedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataZoneIsolated.
// This is useful for switch cases.
type SecurityDataZoneIsolatedExactly interface {
	SecurityDataZoneIsolated
	isSecurityDataZoneIsolated() bool
}

// _SecurityDataZoneIsolated is the data-structure of this message
type _SecurityDataZoneIsolated struct {
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

func (m *_SecurityDataZoneIsolated) InitializeParent(parent SecurityData , commandTypeContainer SecurityCommandTypeContainer , argument byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataZoneIsolated)  GetParent() SecurityData {
	return m._SecurityData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataZoneIsolated) GetZoneNumber() uint8 {
	return m.ZoneNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewSecurityDataZoneIsolated factory function for _SecurityDataZoneIsolated
func NewSecurityDataZoneIsolated( zoneNumber uint8 , commandTypeContainer SecurityCommandTypeContainer , argument byte ) *_SecurityDataZoneIsolated {
	_result := &_SecurityDataZoneIsolated{
		ZoneNumber: zoneNumber,
    	_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataZoneIsolated(structType interface{}) SecurityDataZoneIsolated {
    if casted, ok := structType.(SecurityDataZoneIsolated); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataZoneIsolated); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataZoneIsolated) GetTypeName() string {
	return "SecurityDataZoneIsolated"
}

func (m *_SecurityDataZoneIsolated) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (zoneNumber)
	lengthInBits += 8;

	return lengthInBits
}


func (m *_SecurityDataZoneIsolated) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SecurityDataZoneIsolatedParse(theBytes []byte) (SecurityDataZoneIsolated, error) {
	return SecurityDataZoneIsolatedParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataZoneIsolatedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (SecurityDataZoneIsolated, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataZoneIsolated"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataZoneIsolated")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (zoneNumber)
_zoneNumber, _zoneNumberErr := readBuffer.ReadUint8("zoneNumber", 8)
	if _zoneNumberErr != nil {
		return nil, errors.Wrap(_zoneNumberErr, "Error parsing 'zoneNumber' field of SecurityDataZoneIsolated")
	}
	zoneNumber := _zoneNumber

	if closeErr := readBuffer.CloseContext("SecurityDataZoneIsolated"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataZoneIsolated")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataZoneIsolated{
		_SecurityData: &_SecurityData{
		},
		ZoneNumber: zoneNumber,
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataZoneIsolated) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataZoneIsolated) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataZoneIsolated"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataZoneIsolated")
		}

	// Simple Field (zoneNumber)
	zoneNumber := uint8(m.GetZoneNumber())
	_zoneNumberErr := writeBuffer.WriteUint8("zoneNumber", 8, (zoneNumber))
	if _zoneNumberErr != nil {
		return errors.Wrap(_zoneNumberErr, "Error serializing 'zoneNumber' field")
	}

		if popErr := writeBuffer.PopContext("SecurityDataZoneIsolated"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataZoneIsolated")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_SecurityDataZoneIsolated) isSecurityDataZoneIsolated() bool {
	return true
}

func (m *_SecurityDataZoneIsolated) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



