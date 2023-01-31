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


// State is the corresponding interface of State
type State interface {
	utils.LengthAware
	utils.Serializable
	// GetSIG_8 returns SIG_8 (property field)
	GetSIG_8() bool
	// GetSIG_7 returns SIG_7 (property field)
	GetSIG_7() bool
	// GetSIG_6 returns SIG_6 (property field)
	GetSIG_6() bool
	// GetSIG_5 returns SIG_5 (property field)
	GetSIG_5() bool
	// GetSIG_4 returns SIG_4 (property field)
	GetSIG_4() bool
	// GetSIG_3 returns SIG_3 (property field)
	GetSIG_3() bool
	// GetSIG_2 returns SIG_2 (property field)
	GetSIG_2() bool
	// GetSIG_1 returns SIG_1 (property field)
	GetSIG_1() bool
}

// StateExactly can be used when we want exactly this type and not a type which fulfills State.
// This is useful for switch cases.
type StateExactly interface {
	State
	isState() bool
}

// _State is the data-structure of this message
type _State struct {
        SIG_8 bool
        SIG_7 bool
        SIG_6 bool
        SIG_5 bool
        SIG_4 bool
        SIG_3 bool
        SIG_2 bool
        SIG_1 bool
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_State) GetSIG_8() bool {
	return m.SIG_8
}

func (m *_State) GetSIG_7() bool {
	return m.SIG_7
}

func (m *_State) GetSIG_6() bool {
	return m.SIG_6
}

func (m *_State) GetSIG_5() bool {
	return m.SIG_5
}

func (m *_State) GetSIG_4() bool {
	return m.SIG_4
}

func (m *_State) GetSIG_3() bool {
	return m.SIG_3
}

func (m *_State) GetSIG_2() bool {
	return m.SIG_2
}

func (m *_State) GetSIG_1() bool {
	return m.SIG_1
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewState factory function for _State
func NewState( SIG_8 bool , SIG_7 bool , SIG_6 bool , SIG_5 bool , SIG_4 bool , SIG_3 bool , SIG_2 bool , SIG_1 bool ) *_State {
return &_State{ SIG_8: SIG_8 , SIG_7: SIG_7 , SIG_6: SIG_6 , SIG_5: SIG_5 , SIG_4: SIG_4 , SIG_3: SIG_3 , SIG_2: SIG_2 , SIG_1: SIG_1 }
}

// Deprecated: use the interface for direct cast
func CastState(structType interface{}) State {
    if casted, ok := structType.(State); ok {
		return casted
	}
	if casted, ok := structType.(*State); ok {
		return *casted
	}
	return nil
}

func (m *_State) GetTypeName() string {
	return "State"
}

func (m *_State) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_State) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (SIG_8)
	lengthInBits += 1;

	// Simple field (SIG_7)
	lengthInBits += 1;

	// Simple field (SIG_6)
	lengthInBits += 1;

	// Simple field (SIG_5)
	lengthInBits += 1;

	// Simple field (SIG_4)
	lengthInBits += 1;

	// Simple field (SIG_3)
	lengthInBits += 1;

	// Simple field (SIG_2)
	lengthInBits += 1;

	// Simple field (SIG_1)
	lengthInBits += 1;

	return lengthInBits
}


func (m *_State) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func StateParse(theBytes []byte) (State, error) {
	return StateParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func StateParseWithBuffer(readBuffer utils.ReadBuffer) (State, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("State"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for State")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (SIG_8)
_SIG_8, _SIG_8Err := readBuffer.ReadBit("SIG_8")
	if _SIG_8Err != nil {
		return nil, errors.Wrap(_SIG_8Err, "Error parsing 'SIG_8' field of State")
	}
	SIG_8 := _SIG_8

	// Simple Field (SIG_7)
_SIG_7, _SIG_7Err := readBuffer.ReadBit("SIG_7")
	if _SIG_7Err != nil {
		return nil, errors.Wrap(_SIG_7Err, "Error parsing 'SIG_7' field of State")
	}
	SIG_7 := _SIG_7

	// Simple Field (SIG_6)
_SIG_6, _SIG_6Err := readBuffer.ReadBit("SIG_6")
	if _SIG_6Err != nil {
		return nil, errors.Wrap(_SIG_6Err, "Error parsing 'SIG_6' field of State")
	}
	SIG_6 := _SIG_6

	// Simple Field (SIG_5)
_SIG_5, _SIG_5Err := readBuffer.ReadBit("SIG_5")
	if _SIG_5Err != nil {
		return nil, errors.Wrap(_SIG_5Err, "Error parsing 'SIG_5' field of State")
	}
	SIG_5 := _SIG_5

	// Simple Field (SIG_4)
_SIG_4, _SIG_4Err := readBuffer.ReadBit("SIG_4")
	if _SIG_4Err != nil {
		return nil, errors.Wrap(_SIG_4Err, "Error parsing 'SIG_4' field of State")
	}
	SIG_4 := _SIG_4

	// Simple Field (SIG_3)
_SIG_3, _SIG_3Err := readBuffer.ReadBit("SIG_3")
	if _SIG_3Err != nil {
		return nil, errors.Wrap(_SIG_3Err, "Error parsing 'SIG_3' field of State")
	}
	SIG_3 := _SIG_3

	// Simple Field (SIG_2)
_SIG_2, _SIG_2Err := readBuffer.ReadBit("SIG_2")
	if _SIG_2Err != nil {
		return nil, errors.Wrap(_SIG_2Err, "Error parsing 'SIG_2' field of State")
	}
	SIG_2 := _SIG_2

	// Simple Field (SIG_1)
_SIG_1, _SIG_1Err := readBuffer.ReadBit("SIG_1")
	if _SIG_1Err != nil {
		return nil, errors.Wrap(_SIG_1Err, "Error parsing 'SIG_1' field of State")
	}
	SIG_1 := _SIG_1

	if closeErr := readBuffer.CloseContext("State"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for State")
	}

	// Create the instance
	return &_State{
			SIG_8: SIG_8,
			SIG_7: SIG_7,
			SIG_6: SIG_6,
			SIG_5: SIG_5,
			SIG_4: SIG_4,
			SIG_3: SIG_3,
			SIG_2: SIG_2,
			SIG_1: SIG_1,
		}, nil
}

func (m *_State) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_State) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("State"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for State")
	}

	// Simple Field (SIG_8)
	SIG_8 := bool(m.GetSIG_8())
	_SIG_8Err := writeBuffer.WriteBit("SIG_8", (SIG_8))
	if _SIG_8Err != nil {
		return errors.Wrap(_SIG_8Err, "Error serializing 'SIG_8' field")
	}

	// Simple Field (SIG_7)
	SIG_7 := bool(m.GetSIG_7())
	_SIG_7Err := writeBuffer.WriteBit("SIG_7", (SIG_7))
	if _SIG_7Err != nil {
		return errors.Wrap(_SIG_7Err, "Error serializing 'SIG_7' field")
	}

	// Simple Field (SIG_6)
	SIG_6 := bool(m.GetSIG_6())
	_SIG_6Err := writeBuffer.WriteBit("SIG_6", (SIG_6))
	if _SIG_6Err != nil {
		return errors.Wrap(_SIG_6Err, "Error serializing 'SIG_6' field")
	}

	// Simple Field (SIG_5)
	SIG_5 := bool(m.GetSIG_5())
	_SIG_5Err := writeBuffer.WriteBit("SIG_5", (SIG_5))
	if _SIG_5Err != nil {
		return errors.Wrap(_SIG_5Err, "Error serializing 'SIG_5' field")
	}

	// Simple Field (SIG_4)
	SIG_4 := bool(m.GetSIG_4())
	_SIG_4Err := writeBuffer.WriteBit("SIG_4", (SIG_4))
	if _SIG_4Err != nil {
		return errors.Wrap(_SIG_4Err, "Error serializing 'SIG_4' field")
	}

	// Simple Field (SIG_3)
	SIG_3 := bool(m.GetSIG_3())
	_SIG_3Err := writeBuffer.WriteBit("SIG_3", (SIG_3))
	if _SIG_3Err != nil {
		return errors.Wrap(_SIG_3Err, "Error serializing 'SIG_3' field")
	}

	// Simple Field (SIG_2)
	SIG_2 := bool(m.GetSIG_2())
	_SIG_2Err := writeBuffer.WriteBit("SIG_2", (SIG_2))
	if _SIG_2Err != nil {
		return errors.Wrap(_SIG_2Err, "Error serializing 'SIG_2' field")
	}

	// Simple Field (SIG_1)
	SIG_1 := bool(m.GetSIG_1())
	_SIG_1Err := writeBuffer.WriteBit("SIG_1", (SIG_1))
	if _SIG_1Err != nil {
		return errors.Wrap(_SIG_1Err, "Error serializing 'SIG_1' field")
	}

	if popErr := writeBuffer.PopContext("State"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for State")
	}
	return nil
}


func (m *_State) isState() bool {
	return true
}

func (m *_State) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



