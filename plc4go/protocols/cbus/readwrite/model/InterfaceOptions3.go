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


// InterfaceOptions3 is the corresponding interface of InterfaceOptions3
type InterfaceOptions3 interface {
	utils.LengthAware
	utils.Serializable
	// GetExstat returns Exstat (property field)
	GetExstat() bool
	// GetPun returns Pun (property field)
	GetPun() bool
	// GetLocalSal returns LocalSal (property field)
	GetLocalSal() bool
	// GetPcn returns Pcn (property field)
	GetPcn() bool
}

// InterfaceOptions3Exactly can be used when we want exactly this type and not a type which fulfills InterfaceOptions3.
// This is useful for switch cases.
type InterfaceOptions3Exactly interface {
	InterfaceOptions3
	isInterfaceOptions3() bool
}

// _InterfaceOptions3 is the data-structure of this message
type _InterfaceOptions3 struct {
        Exstat bool
        Pun bool
        LocalSal bool
        Pcn bool
	// Reserved Fields
	reservedField0 *bool
	reservedField1 *bool
	reservedField2 *bool
	reservedField3 *bool
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_InterfaceOptions3) GetExstat() bool {
	return m.Exstat
}

func (m *_InterfaceOptions3) GetPun() bool {
	return m.Pun
}

func (m *_InterfaceOptions3) GetLocalSal() bool {
	return m.LocalSal
}

func (m *_InterfaceOptions3) GetPcn() bool {
	return m.Pcn
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewInterfaceOptions3 factory function for _InterfaceOptions3
func NewInterfaceOptions3( exstat bool , pun bool , localSal bool , pcn bool ) *_InterfaceOptions3 {
return &_InterfaceOptions3{ Exstat: exstat , Pun: pun , LocalSal: localSal , Pcn: pcn }
}

// Deprecated: use the interface for direct cast
func CastInterfaceOptions3(structType interface{}) InterfaceOptions3 {
    if casted, ok := structType.(InterfaceOptions3); ok {
		return casted
	}
	if casted, ok := structType.(*InterfaceOptions3); ok {
		return *casted
	}
	return nil
}

func (m *_InterfaceOptions3) GetTypeName() string {
	return "InterfaceOptions3"
}

func (m *_InterfaceOptions3) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_InterfaceOptions3) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 1

	// Simple field (exstat)
	lengthInBits += 1;

	// Simple field (pun)
	lengthInBits += 1;

	// Simple field (localSal)
	lengthInBits += 1;

	// Simple field (pcn)
	lengthInBits += 1;

	return lengthInBits
}


func (m *_InterfaceOptions3) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func InterfaceOptions3Parse(theBytes []byte) (InterfaceOptions3, error) {
	return InterfaceOptions3ParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func InterfaceOptions3ParseWithBuffer(readBuffer utils.ReadBuffer) (InterfaceOptions3, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("InterfaceOptions3"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for InterfaceOptions3")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of InterfaceOptions3")
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

	var reservedField1 *bool
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadBit("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of InterfaceOptions3")
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
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of InterfaceOptions3")
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
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of InterfaceOptions3")
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

	// Simple Field (exstat)
_exstat, _exstatErr := readBuffer.ReadBit("exstat")
	if _exstatErr != nil {
		return nil, errors.Wrap(_exstatErr, "Error parsing 'exstat' field of InterfaceOptions3")
	}
	exstat := _exstat

	// Simple Field (pun)
_pun, _punErr := readBuffer.ReadBit("pun")
	if _punErr != nil {
		return nil, errors.Wrap(_punErr, "Error parsing 'pun' field of InterfaceOptions3")
	}
	pun := _pun

	// Simple Field (localSal)
_localSal, _localSalErr := readBuffer.ReadBit("localSal")
	if _localSalErr != nil {
		return nil, errors.Wrap(_localSalErr, "Error parsing 'localSal' field of InterfaceOptions3")
	}
	localSal := _localSal

	// Simple Field (pcn)
_pcn, _pcnErr := readBuffer.ReadBit("pcn")
	if _pcnErr != nil {
		return nil, errors.Wrap(_pcnErr, "Error parsing 'pcn' field of InterfaceOptions3")
	}
	pcn := _pcn

	if closeErr := readBuffer.CloseContext("InterfaceOptions3"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for InterfaceOptions3")
	}

	// Create the instance
	return &_InterfaceOptions3{
			Exstat: exstat,
			Pun: pun,
			LocalSal: localSal,
			Pcn: pcn,
			reservedField0: reservedField0,
			reservedField1: reservedField1,
			reservedField2: reservedField2,
			reservedField3: reservedField3,
		}, nil
}

func (m *_InterfaceOptions3) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_InterfaceOptions3) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("InterfaceOptions3"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for InterfaceOptions3")
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

	// Simple Field (exstat)
	exstat := bool(m.GetExstat())
	_exstatErr := writeBuffer.WriteBit("exstat", (exstat))
	if _exstatErr != nil {
		return errors.Wrap(_exstatErr, "Error serializing 'exstat' field")
	}

	// Simple Field (pun)
	pun := bool(m.GetPun())
	_punErr := writeBuffer.WriteBit("pun", (pun))
	if _punErr != nil {
		return errors.Wrap(_punErr, "Error serializing 'pun' field")
	}

	// Simple Field (localSal)
	localSal := bool(m.GetLocalSal())
	_localSalErr := writeBuffer.WriteBit("localSal", (localSal))
	if _localSalErr != nil {
		return errors.Wrap(_localSalErr, "Error serializing 'localSal' field")
	}

	// Simple Field (pcn)
	pcn := bool(m.GetPcn())
	_pcnErr := writeBuffer.WriteBit("pcn", (pcn))
	if _pcnErr != nil {
		return errors.Wrap(_pcnErr, "Error serializing 'pcn' field")
	}

	if popErr := writeBuffer.PopContext("InterfaceOptions3"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for InterfaceOptions3")
	}
	return nil
}


func (m *_InterfaceOptions3) isInterfaceOptions3() bool {
	return true
}

func (m *_InterfaceOptions3) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



