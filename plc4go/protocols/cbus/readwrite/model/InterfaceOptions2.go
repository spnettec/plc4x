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


// InterfaceOptions2 is the corresponding interface of InterfaceOptions2
type InterfaceOptions2 interface {
	utils.LengthAware
	utils.Serializable
	// GetBurden returns Burden (property field)
	GetBurden() bool
	// GetClockGen returns ClockGen (property field)
	GetClockGen() bool
}

// InterfaceOptions2Exactly can be used when we want exactly this type and not a type which fulfills InterfaceOptions2.
// This is useful for switch cases.
type InterfaceOptions2Exactly interface {
	InterfaceOptions2
	isInterfaceOptions2() bool
}

// _InterfaceOptions2 is the data-structure of this message
type _InterfaceOptions2 struct {
        Burden bool
        ClockGen bool
	// Reserved Fields
	reservedField0 *bool
	reservedField1 *bool
	reservedField2 *bool
	reservedField3 *bool
	reservedField4 *bool
	reservedField5 *bool
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_InterfaceOptions2) GetBurden() bool {
	return m.Burden
}

func (m *_InterfaceOptions2) GetClockGen() bool {
	return m.ClockGen
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewInterfaceOptions2 factory function for _InterfaceOptions2
func NewInterfaceOptions2( burden bool , clockGen bool ) *_InterfaceOptions2 {
return &_InterfaceOptions2{ Burden: burden , ClockGen: clockGen }
}

// Deprecated: use the interface for direct cast
func CastInterfaceOptions2(structType interface{}) InterfaceOptions2 {
    if casted, ok := structType.(InterfaceOptions2); ok {
		return casted
	}
	if casted, ok := structType.(*InterfaceOptions2); ok {
		return *casted
	}
	return nil
}

func (m *_InterfaceOptions2) GetTypeName() string {
	return "InterfaceOptions2"
}

func (m *_InterfaceOptions2) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_InterfaceOptions2) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (burden)
	lengthInBits += 1;

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (clockGen)
	lengthInBits += 1;

	return lengthInBits
}


func (m *_InterfaceOptions2) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func InterfaceOptions2Parse(theBytes []byte) (InterfaceOptions2, error) {
	return InterfaceOptions2ParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func InterfaceOptions2ParseWithBuffer(readBuffer utils.ReadBuffer) (InterfaceOptions2, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("InterfaceOptions2"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for InterfaceOptions2")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of InterfaceOptions2")
		}
		if reserved != bool(false) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value": reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (burden)
_burden, _burdenErr := readBuffer.ReadBit("burden")
	if _burdenErr != nil {
		return nil, errors.Wrap(_burdenErr, "Error parsing 'burden' field of InterfaceOptions2")
	}
	burden := _burden

	var reservedField1 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of InterfaceOptions2")
		}
		if reserved != bool(false) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value": reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField1 = &reserved
		}
	}

	var reservedField2 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of InterfaceOptions2")
		}
		if reserved != bool(false) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value": reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField2 = &reserved
		}
	}

	var reservedField3 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of InterfaceOptions2")
		}
		if reserved != bool(false) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value": reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField3 = &reserved
		}
	}

	var reservedField4 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of InterfaceOptions2")
		}
		if reserved != bool(false) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value": reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField4 = &reserved
		}
	}

	var reservedField5 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of InterfaceOptions2")
		}
		if reserved != bool(false) {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value": reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField5 = &reserved
		}
	}

	// Simple Field (clockGen)
_clockGen, _clockGenErr := readBuffer.ReadBit("clockGen")
	if _clockGenErr != nil {
		return nil, errors.Wrap(_clockGenErr, "Error parsing 'clockGen' field of InterfaceOptions2")
	}
	clockGen := _clockGen

	if closeErr := readBuffer.CloseContext("InterfaceOptions2"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for InterfaceOptions2")
	}

	// Create the instance
	return &_InterfaceOptions2{
			Burden: burden,
			ClockGen: clockGen,
			reservedField0: reservedField0,
			reservedField1: reservedField1,
			reservedField2: reservedField2,
			reservedField3: reservedField3,
			reservedField4: reservedField4,
			reservedField5: reservedField5,
		}, nil
}

func (m *_InterfaceOptions2) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_InterfaceOptions2) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("InterfaceOptions2"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for InterfaceOptions2")
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField0 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value": reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (burden)
	burden := bool(m.GetBurden())
	_burdenErr := writeBuffer.WriteBit("burden", (burden))
	if _burdenErr != nil {
		return errors.Wrap(_burdenErr, "Error serializing 'burden' field")
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField1 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value": reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField1
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField2 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value": reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField2
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField3 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value": reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField3
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField4 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value": reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField4
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Reserved Field (reserved)
	{
		var reserved bool = bool(false)
		if m.reservedField5 != nil {
			Plc4xModelLog.Info().Fields(map[string]interface{}{
				"expected value": bool(false),
				"got value": reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField5
		}
		_err := writeBuffer.WriteBit("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (clockGen)
	clockGen := bool(m.GetClockGen())
	_clockGenErr := writeBuffer.WriteBit("clockGen", (clockGen))
	if _clockGenErr != nil {
		return errors.Wrap(_clockGenErr, "Error serializing 'clockGen' field")
	}

	if popErr := writeBuffer.PopContext("InterfaceOptions2"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for InterfaceOptions2")
	}
	return nil
}


func (m *_InterfaceOptions2) isInterfaceOptions2() bool {
	return true
}

func (m *_InterfaceOptions2) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



