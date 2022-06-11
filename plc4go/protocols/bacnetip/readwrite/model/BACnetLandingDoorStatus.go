/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetLandingDoorStatus is the data-structure of this message
type BACnetLandingDoorStatus struct {
	LandingDoors *BACnetLandingDoorStatusLandingDoorsList
}

// IBACnetLandingDoorStatus is the corresponding interface of BACnetLandingDoorStatus
type IBACnetLandingDoorStatus interface {
	// GetLandingDoors returns LandingDoors (property field)
	GetLandingDoors() *BACnetLandingDoorStatusLandingDoorsList
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetLandingDoorStatus) GetLandingDoors() *BACnetLandingDoorStatusLandingDoorsList {
	return m.LandingDoors
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLandingDoorStatus factory function for BACnetLandingDoorStatus
func NewBACnetLandingDoorStatus(landingDoors *BACnetLandingDoorStatusLandingDoorsList) *BACnetLandingDoorStatus {
	return &BACnetLandingDoorStatus{LandingDoors: landingDoors}
}

func CastBACnetLandingDoorStatus(structType interface{}) *BACnetLandingDoorStatus {
	if casted, ok := structType.(BACnetLandingDoorStatus); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetLandingDoorStatus); ok {
		return casted
	}
	return nil
}

func (m *BACnetLandingDoorStatus) GetTypeName() string {
	return "BACnetLandingDoorStatus"
}

func (m *BACnetLandingDoorStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetLandingDoorStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (landingDoors)
	lengthInBits += m.LandingDoors.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetLandingDoorStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLandingDoorStatusParse(readBuffer utils.ReadBuffer) (*BACnetLandingDoorStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLandingDoorStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLandingDoorStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (landingDoors)
	if pullErr := readBuffer.PullContext("landingDoors"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for landingDoors")
	}
	_landingDoors, _landingDoorsErr := BACnetLandingDoorStatusLandingDoorsListParse(readBuffer, uint8(uint8(0)))
	if _landingDoorsErr != nil {
		return nil, errors.Wrap(_landingDoorsErr, "Error parsing 'landingDoors' field")
	}
	landingDoors := CastBACnetLandingDoorStatusLandingDoorsList(_landingDoors)
	if closeErr := readBuffer.CloseContext("landingDoors"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for landingDoors")
	}

	if closeErr := readBuffer.CloseContext("BACnetLandingDoorStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLandingDoorStatus")
	}

	// Create the instance
	return NewBACnetLandingDoorStatus(landingDoors), nil
}

func (m *BACnetLandingDoorStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLandingDoorStatus"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLandingDoorStatus")
	}

	// Simple Field (landingDoors)
	if pushErr := writeBuffer.PushContext("landingDoors"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for landingDoors")
	}
	_landingDoorsErr := m.LandingDoors.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("landingDoors"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for landingDoors")
	}
	if _landingDoorsErr != nil {
		return errors.Wrap(_landingDoorsErr, "Error serializing 'landingDoors' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLandingDoorStatus"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLandingDoorStatus")
	}
	return nil
}

func (m *BACnetLandingDoorStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
