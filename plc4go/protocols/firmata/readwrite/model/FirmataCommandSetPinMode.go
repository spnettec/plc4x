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


// FirmataCommandSetPinMode is the corresponding interface of FirmataCommandSetPinMode
type FirmataCommandSetPinMode interface {
	utils.LengthAware
	utils.Serializable
	FirmataCommand
	// GetPin returns Pin (property field)
	GetPin() uint8
	// GetMode returns Mode (property field)
	GetMode() PinMode
}

// FirmataCommandSetPinModeExactly can be used when we want exactly this type and not a type which fulfills FirmataCommandSetPinMode.
// This is useful for switch cases.
type FirmataCommandSetPinModeExactly interface {
	FirmataCommandSetPinMode
	isFirmataCommandSetPinMode() bool
}

// _FirmataCommandSetPinMode is the data-structure of this message
type _FirmataCommandSetPinMode struct {
	*_FirmataCommand
        Pin uint8
        Mode PinMode
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_FirmataCommandSetPinMode)  GetCommandCode() uint8 {
return 0x4}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_FirmataCommandSetPinMode) InitializeParent(parent FirmataCommand ) {}

func (m *_FirmataCommandSetPinMode)  GetParent() FirmataCommand {
	return m._FirmataCommand
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_FirmataCommandSetPinMode) GetPin() uint8 {
	return m.Pin
}

func (m *_FirmataCommandSetPinMode) GetMode() PinMode {
	return m.Mode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewFirmataCommandSetPinMode factory function for _FirmataCommandSetPinMode
func NewFirmataCommandSetPinMode( pin uint8 , mode PinMode , response bool ) *_FirmataCommandSetPinMode {
	_result := &_FirmataCommandSetPinMode{
		Pin: pin,
		Mode: mode,
    	_FirmataCommand: NewFirmataCommand(response),
	}
	_result._FirmataCommand._FirmataCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastFirmataCommandSetPinMode(structType interface{}) FirmataCommandSetPinMode {
    if casted, ok := structType.(FirmataCommandSetPinMode); ok {
		return casted
	}
	if casted, ok := structType.(*FirmataCommandSetPinMode); ok {
		return *casted
	}
	return nil
}

func (m *_FirmataCommandSetPinMode) GetTypeName() string {
	return "FirmataCommandSetPinMode"
}

func (m *_FirmataCommandSetPinMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (pin)
	lengthInBits += 8;

	// Simple field (mode)
	lengthInBits += 8

	return lengthInBits
}


func (m *_FirmataCommandSetPinMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func FirmataCommandSetPinModeParse(theBytes []byte, response bool) (FirmataCommandSetPinMode, error) {
	return FirmataCommandSetPinModeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), response)
}

func FirmataCommandSetPinModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (FirmataCommandSetPinMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("FirmataCommandSetPinMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for FirmataCommandSetPinMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (pin)
_pin, _pinErr := readBuffer.ReadUint8("pin", 8)
	if _pinErr != nil {
		return nil, errors.Wrap(_pinErr, "Error parsing 'pin' field of FirmataCommandSetPinMode")
	}
	pin := _pin

	// Simple Field (mode)
	if pullErr := readBuffer.PullContext("mode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for mode")
	}
_mode, _modeErr := PinModeParseWithBuffer(ctx, readBuffer)
	if _modeErr != nil {
		return nil, errors.Wrap(_modeErr, "Error parsing 'mode' field of FirmataCommandSetPinMode")
	}
	mode := _mode
	if closeErr := readBuffer.CloseContext("mode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for mode")
	}

	if closeErr := readBuffer.CloseContext("FirmataCommandSetPinMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for FirmataCommandSetPinMode")
	}

	// Create a partially initialized instance
	_child := &_FirmataCommandSetPinMode{
		_FirmataCommand: &_FirmataCommand{
			Response: response,
		},
		Pin: pin,
		Mode: mode,
	}
	_child._FirmataCommand._FirmataCommandChildRequirements = _child
	return _child, nil
}

func (m *_FirmataCommandSetPinMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_FirmataCommandSetPinMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("FirmataCommandSetPinMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for FirmataCommandSetPinMode")
		}

	// Simple Field (pin)
	pin := uint8(m.GetPin())
	_pinErr := writeBuffer.WriteUint8("pin", 8, (pin))
	if _pinErr != nil {
		return errors.Wrap(_pinErr, "Error serializing 'pin' field")
	}

	// Simple Field (mode)
	if pushErr := writeBuffer.PushContext("mode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for mode")
	}
	_modeErr := writeBuffer.WriteSerializable(ctx, m.GetMode())
	if popErr := writeBuffer.PopContext("mode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for mode")
	}
	if _modeErr != nil {
		return errors.Wrap(_modeErr, "Error serializing 'mode' field")
	}

		if popErr := writeBuffer.PopContext("FirmataCommandSetPinMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for FirmataCommandSetPinMode")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_FirmataCommandSetPinMode) isFirmataCommandSetPinMode() bool {
	return true
}

func (m *_FirmataCommandSetPinMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



