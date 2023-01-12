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


// BACnetLiftCarCallList is the corresponding interface of BACnetLiftCarCallList
type BACnetLiftCarCallList interface {
	utils.LengthAware
	utils.Serializable
	// GetFloorNumbers returns FloorNumbers (property field)
	GetFloorNumbers() BACnetLiftCarCallListFloorList
}

// BACnetLiftCarCallListExactly can be used when we want exactly this type and not a type which fulfills BACnetLiftCarCallList.
// This is useful for switch cases.
type BACnetLiftCarCallListExactly interface {
	BACnetLiftCarCallList
	isBACnetLiftCarCallList() bool
}

// _BACnetLiftCarCallList is the data-structure of this message
type _BACnetLiftCarCallList struct {
        FloorNumbers BACnetLiftCarCallListFloorList
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLiftCarCallList) GetFloorNumbers() BACnetLiftCarCallListFloorList {
	return m.FloorNumbers
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetLiftCarCallList factory function for _BACnetLiftCarCallList
func NewBACnetLiftCarCallList( floorNumbers BACnetLiftCarCallListFloorList ) *_BACnetLiftCarCallList {
return &_BACnetLiftCarCallList{ FloorNumbers: floorNumbers }
}

// Deprecated: use the interface for direct cast
func CastBACnetLiftCarCallList(structType interface{}) BACnetLiftCarCallList {
    if casted, ok := structType.(BACnetLiftCarCallList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLiftCarCallList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLiftCarCallList) GetTypeName() string {
	return "BACnetLiftCarCallList"
}

func (m *_BACnetLiftCarCallList) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLiftCarCallList) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (floorNumbers)
	lengthInBits += m.FloorNumbers.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetLiftCarCallList) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLiftCarCallListParse(theBytes []byte) (BACnetLiftCarCallList, error) {
	return BACnetLiftCarCallListParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetLiftCarCallListParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetLiftCarCallList, error) {
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
_floorNumbers, _floorNumbersErr := BACnetLiftCarCallListFloorListParseWithBuffer(readBuffer , uint8( uint8(0) ) )
	if _floorNumbersErr != nil {
		return nil, errors.Wrap(_floorNumbersErr, "Error parsing 'floorNumbers' field of BACnetLiftCarCallList")
	}
	floorNumbers := _floorNumbers.(BACnetLiftCarCallListFloorList)
	if closeErr := readBuffer.CloseContext("floorNumbers"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for floorNumbers")
	}

	if closeErr := readBuffer.CloseContext("BACnetLiftCarCallList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLiftCarCallList")
	}

	// Create the instance
	return &_BACnetLiftCarCallList{
			FloorNumbers: floorNumbers,
		}, nil
}

func (m *_BACnetLiftCarCallList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetLiftCarCallList) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetLiftCarCallList"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetLiftCarCallList")
	}

	// Simple Field (floorNumbers)
	if pushErr := writeBuffer.PushContext("floorNumbers"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for floorNumbers")
	}
	_floorNumbersErr := writeBuffer.WriteSerializable(m.GetFloorNumbers())
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


func (m *_BACnetLiftCarCallList) isBACnetLiftCarCallList() bool {
	return true
}

func (m *_BACnetLiftCarCallList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



