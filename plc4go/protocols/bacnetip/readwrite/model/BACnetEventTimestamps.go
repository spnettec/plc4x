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


// BACnetEventTimestamps is the corresponding interface of BACnetEventTimestamps
type BACnetEventTimestamps interface {
	utils.LengthAware
	utils.Serializable
	// GetToOffnormal returns ToOffnormal (property field)
	GetToOffnormal() BACnetTimeStamp
	// GetToFault returns ToFault (property field)
	GetToFault() BACnetTimeStamp
	// GetToNormal returns ToNormal (property field)
	GetToNormal() BACnetTimeStamp
}

// BACnetEventTimestampsExactly can be used when we want exactly this type and not a type which fulfills BACnetEventTimestamps.
// This is useful for switch cases.
type BACnetEventTimestampsExactly interface {
	BACnetEventTimestamps
	isBACnetEventTimestamps() bool
}

// _BACnetEventTimestamps is the data-structure of this message
type _BACnetEventTimestamps struct {
        ToOffnormal BACnetTimeStamp
        ToFault BACnetTimeStamp
        ToNormal BACnetTimeStamp
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventTimestamps) GetToOffnormal() BACnetTimeStamp {
	return m.ToOffnormal
}

func (m *_BACnetEventTimestamps) GetToFault() BACnetTimeStamp {
	return m.ToFault
}

func (m *_BACnetEventTimestamps) GetToNormal() BACnetTimeStamp {
	return m.ToNormal
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetEventTimestamps factory function for _BACnetEventTimestamps
func NewBACnetEventTimestamps( toOffnormal BACnetTimeStamp , toFault BACnetTimeStamp , toNormal BACnetTimeStamp ) *_BACnetEventTimestamps {
return &_BACnetEventTimestamps{ ToOffnormal: toOffnormal , ToFault: toFault , ToNormal: toNormal }
}

// Deprecated: use the interface for direct cast
func CastBACnetEventTimestamps(structType interface{}) BACnetEventTimestamps {
    if casted, ok := structType.(BACnetEventTimestamps); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventTimestamps); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventTimestamps) GetTypeName() string {
	return "BACnetEventTimestamps"
}

func (m *_BACnetEventTimestamps) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetEventTimestamps) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (toOffnormal)
	lengthInBits += m.ToOffnormal.GetLengthInBits()

	// Simple field (toFault)
	lengthInBits += m.ToFault.GetLengthInBits()

	// Simple field (toNormal)
	lengthInBits += m.ToNormal.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetEventTimestamps) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetEventTimestampsParse(theBytes []byte) (BACnetEventTimestamps, error) {
	return BACnetEventTimestampsParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventTimestampsParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetEventTimestamps, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventTimestamps"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventTimestamps")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (toOffnormal)
	if pullErr := readBuffer.PullContext("toOffnormal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for toOffnormal")
	}
_toOffnormal, _toOffnormalErr := BACnetTimeStampParseWithBuffer(readBuffer)
	if _toOffnormalErr != nil {
		return nil, errors.Wrap(_toOffnormalErr, "Error parsing 'toOffnormal' field of BACnetEventTimestamps")
	}
	toOffnormal := _toOffnormal.(BACnetTimeStamp)
	if closeErr := readBuffer.CloseContext("toOffnormal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for toOffnormal")
	}

	// Simple Field (toFault)
	if pullErr := readBuffer.PullContext("toFault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for toFault")
	}
_toFault, _toFaultErr := BACnetTimeStampParseWithBuffer(readBuffer)
	if _toFaultErr != nil {
		return nil, errors.Wrap(_toFaultErr, "Error parsing 'toFault' field of BACnetEventTimestamps")
	}
	toFault := _toFault.(BACnetTimeStamp)
	if closeErr := readBuffer.CloseContext("toFault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for toFault")
	}

	// Simple Field (toNormal)
	if pullErr := readBuffer.PullContext("toNormal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for toNormal")
	}
_toNormal, _toNormalErr := BACnetTimeStampParseWithBuffer(readBuffer)
	if _toNormalErr != nil {
		return nil, errors.Wrap(_toNormalErr, "Error parsing 'toNormal' field of BACnetEventTimestamps")
	}
	toNormal := _toNormal.(BACnetTimeStamp)
	if closeErr := readBuffer.CloseContext("toNormal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for toNormal")
	}

	if closeErr := readBuffer.CloseContext("BACnetEventTimestamps"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventTimestamps")
	}

	// Create the instance
	return &_BACnetEventTimestamps{
			ToOffnormal: toOffnormal,
			ToFault: toFault,
			ToNormal: toNormal,
		}, nil
}

func (m *_BACnetEventTimestamps) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventTimestamps) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetEventTimestamps"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventTimestamps")
	}

	// Simple Field (toOffnormal)
	if pushErr := writeBuffer.PushContext("toOffnormal"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for toOffnormal")
	}
	_toOffnormalErr := writeBuffer.WriteSerializable(m.GetToOffnormal())
	if popErr := writeBuffer.PopContext("toOffnormal"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for toOffnormal")
	}
	if _toOffnormalErr != nil {
		return errors.Wrap(_toOffnormalErr, "Error serializing 'toOffnormal' field")
	}

	// Simple Field (toFault)
	if pushErr := writeBuffer.PushContext("toFault"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for toFault")
	}
	_toFaultErr := writeBuffer.WriteSerializable(m.GetToFault())
	if popErr := writeBuffer.PopContext("toFault"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for toFault")
	}
	if _toFaultErr != nil {
		return errors.Wrap(_toFaultErr, "Error serializing 'toFault' field")
	}

	// Simple Field (toNormal)
	if pushErr := writeBuffer.PushContext("toNormal"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for toNormal")
	}
	_toNormalErr := writeBuffer.WriteSerializable(m.GetToNormal())
	if popErr := writeBuffer.PopContext("toNormal"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for toNormal")
	}
	if _toNormalErr != nil {
		return errors.Wrap(_toNormalErr, "Error serializing 'toNormal' field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventTimestamps"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventTimestamps")
	}
	return nil
}


func (m *_BACnetEventTimestamps) isBACnetEventTimestamps() bool {
	return true
}

func (m *_BACnetEventTimestamps) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



