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
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

	// Code generated by code-generation. DO NOT EDIT.


// Dummy is the corresponding interface of Dummy
type Dummy interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetDummy returns Dummy (property field)
	GetDummy() uint16
}

// DummyExactly can be used when we want exactly this type and not a type which fulfills Dummy.
// This is useful for switch cases.
type DummyExactly interface {
	Dummy
	isDummy() bool
}

// _Dummy is the data-structure of this message
type _Dummy struct {
        Dummy uint16
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_Dummy) GetDummy() uint16 {
	return m.Dummy
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewDummy factory function for _Dummy
func NewDummy( dummy uint16 ) *_Dummy {
return &_Dummy{ Dummy: dummy }
}

// Deprecated: use the interface for direct cast
func CastDummy(structType any) Dummy {
    if casted, ok := structType.(Dummy); ok {
		return casted
	}
	if casted, ok := structType.(*Dummy); ok {
		return *casted
	}
	return nil
}

func (m *_Dummy) GetTypeName() string {
	return "Dummy"
}

func (m *_Dummy) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (dummy)
	lengthInBits += 16;

	return lengthInBits
}


func (m *_Dummy) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func DummyParse(ctx context.Context, theBytes []byte) (Dummy, error) {
	return DummyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func DummyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Dummy, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("Dummy"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Dummy")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (dummy)
_dummy, _dummyErr := readBuffer.ReadUint16("dummy", 16)
	if _dummyErr != nil {
		return nil, errors.Wrap(_dummyErr, "Error parsing 'dummy' field of Dummy")
	}
	dummy := _dummy

	if closeErr := readBuffer.CloseContext("Dummy"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Dummy")
	}

	// Create the instance
	return &_Dummy{
			Dummy: dummy,
		}, nil
}

func (m *_Dummy) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Dummy) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("Dummy"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Dummy")
	}

	// Simple Field (dummy)
	dummy := uint16(m.GetDummy())
	_dummyErr := writeBuffer.WriteUint16("dummy", 16, (dummy))
	if _dummyErr != nil {
		return errors.Wrap(_dummyErr, "Error serializing 'dummy' field")
	}

	if popErr := writeBuffer.PopContext("Dummy"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Dummy")
	}
	return nil
}


func (m *_Dummy) isDummy() bool {
	return true
}

func (m *_Dummy) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



