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

// BACnetLiftCarCallList is the data-structure of this message
type BACnetLiftCarCallList struct {
	FloorNumbers *BACnetLiftCarCallListFloorList
}

// IBACnetLiftCarCallList is the corresponding interface of BACnetLiftCarCallList
type IBACnetLiftCarCallList interface {
	// GetFloorNumbers returns FloorNumbers (property field)
	GetFloorNumbers() *BACnetLiftCarCallListFloorList
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

func (m *BACnetLiftCarCallList) GetFloorNumbers() *BACnetLiftCarCallListFloorList {
	return m.FloorNumbers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetLiftCarCallList factory function for BACnetLiftCarCallList
func NewBACnetLiftCarCallList(floorNumbers *BACnetLiftCarCallListFloorList) *BACnetLiftCarCallList {
	return &BACnetLiftCarCallList{FloorNumbers: floorNumbers}
}

func CastBACnetLiftCarCallList(structType interface{}) *BACnetLiftCarCallList {
	if casted, ok := structType.(BACnetLiftCarCallList); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetLiftCarCallList); ok {
		return casted
	}
	return nil
}

func (m *BACnetLiftCarCallList) GetTypeName() string {
	return "BACnetLiftCarCallList"
}

func (m *BACnetLiftCarCallList) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetLiftCarCallList) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (floorNumbers)
	lengthInBits += m.FloorNumbers.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetLiftCarCallList) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLiftCarCallListParse(readBuffer utils.ReadBuffer) (*BACnetLiftCarCallList, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLiftCarCallList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLiftCarCallList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (floorNumbers)
	if pullErr := readBuffer.PullContext("floorNumbers"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for floorNumbers")
	}
	_floorNumbers, _floorNumbersErr := BACnetLiftCarCallListFloorListParse(readBuffer, uint8(uint8(0)))
	if _floorNumbersErr != nil {
		return nil, errors.Wrap(_floorNumbersErr, "Error parsing 'floorNumbers' field")
	}
	floorNumbers := CastBACnetLiftCarCallListFloorList(_floorNumbers)
	if closeErr := readBuffer.CloseContext("floorNumbers"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for floorNumbers")
	}

	if closeErr := readBuffer.CloseContext("BACnetLiftCarCallList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLiftCarCallList")
	}

	// Create the instance
	return NewBACnetLiftCarCallList(floorNumbers), nil
}

func (m *BACnetLiftCarCallList) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetLiftCarCallList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLiftCarCallList")
	}

	// Simple Field (floorNumbers)
	if pushErr := writeBuffer.PushContext("floorNumbers"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for floorNumbers")
	}
	_floorNumbersErr := m.FloorNumbers.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("floorNumbers"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for floorNumbers")
	}
	if _floorNumbersErr != nil {
		return errors.Wrap(_floorNumbersErr, "Error serializing 'floorNumbers' field")
	}

	if popErr := writeBuffer.PopContext("BACnetLiftCarCallList"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetLiftCarCallList")
	}
	return nil
}

func (m *BACnetLiftCarCallList) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
