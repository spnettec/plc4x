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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type APDUUnknown struct {
	*APDU
	UnknownBytes []byte

	// Arguments.
	ApduLength uint16
}

// The corresponding interface
type IAPDUUnknown interface {
	// GetUnknownBytes returns UnknownBytes
	GetUnknownBytes() []byte
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *APDUUnknown) ApduType() uint8 {
	return 0
}

func (m *APDUUnknown) GetApduType() uint8 {
	return 0
}

func (m *APDUUnknown) InitializeParent(parent *APDU) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *APDUUnknown) GetUnknownBytes() []byte {
	return m.UnknownBytes
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewAPDUUnknown factory function for APDUUnknown
func NewAPDUUnknown(unknownBytes []byte, apduLength uint16) *APDU {
	child := &APDUUnknown{
		UnknownBytes: unknownBytes,
		APDU:         NewAPDU(apduLength),
	}
	child.Child = child
	return child.APDU
}

func CastAPDUUnknown(structType interface{}) *APDUUnknown {
	castFunc := func(typ interface{}) *APDUUnknown {
		if casted, ok := typ.(APDUUnknown); ok {
			return &casted
		}
		if casted, ok := typ.(*APDUUnknown); ok {
			return casted
		}
		if casted, ok := typ.(APDU); ok {
			return CastAPDUUnknown(casted.Child)
		}
		if casted, ok := typ.(*APDU); ok {
			return CastAPDUUnknown(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *APDUUnknown) GetTypeName() string {
	return "APDUUnknown"
}

func (m *APDUUnknown) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *APDUUnknown) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.UnknownBytes) > 0 {
		lengthInBits += 8 * uint16(len(m.UnknownBytes))
	}

	return lengthInBits
}

func (m *APDUUnknown) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func APDUUnknownParse(readBuffer utils.ReadBuffer, apduLength uint16) (*APDU, error) {
	if pullErr := readBuffer.PullContext("APDUUnknown"); pullErr != nil {
		return nil, pullErr
	}
	// Byte Array field (unknownBytes)
	numberOfBytesunknownBytes := int(utils.InlineIf(bool(bool((apduLength) > (0))), func() interface{} { return uint16(uint16(uint16(apduLength) - uint16(uint16(1)))) }, func() interface{} { return uint16(uint16(0)) }).(uint16))
	unknownBytes, _readArrayErr := readBuffer.ReadByteArray("unknownBytes", numberOfBytesunknownBytes)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'unknownBytes' field")
	}

	if closeErr := readBuffer.CloseContext("APDUUnknown"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &APDUUnknown{
		UnknownBytes: unknownBytes,
		APDU:         &APDU{},
	}
	_child.APDU.Child = _child
	return _child.APDU, nil
}

func (m *APDUUnknown) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUUnknown"); pushErr != nil {
			return pushErr
		}

		// Array Field (unknownBytes)
		if m.UnknownBytes != nil {
			// Byte Array field (unknownBytes)
			_writeArrayErr := writeBuffer.WriteByteArray("unknownBytes", m.UnknownBytes)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'unknownBytes' field")
			}
		}

		if popErr := writeBuffer.PopContext("APDUUnknown"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *APDUUnknown) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
