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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const BacnetConstants_BACNETUDPDEFAULTPORT uint16 = uint16(47808)

// BacnetConstants is the corresponding interface of BacnetConstants
type BacnetConstants interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// BacnetConstantsExactly can be used when we want exactly this type and not a type which fulfills BacnetConstants.
// This is useful for switch cases.
type BacnetConstantsExactly interface {
	BacnetConstants
	isBacnetConstants() bool
}

// _BacnetConstants is the data-structure of this message
type _BacnetConstants struct {
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_BacnetConstants) GetBacnetUdpDefaultPort() uint16 {
	return BacnetConstants_BACNETUDPDEFAULTPORT
}

///////////////////////-4
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBacnetConstants factory function for _BacnetConstants
func NewBacnetConstants() *_BacnetConstants {
	return &_BacnetConstants{}
}

// Deprecated: use the interface for direct cast
func CastBacnetConstants(structType any) BacnetConstants {
	if casted, ok := structType.(BacnetConstants); ok {
		return casted
	}
	if casted, ok := structType.(*BacnetConstants); ok {
		return *casted
	}
	return nil
}

func (m *_BacnetConstants) GetTypeName() string {
	return "BacnetConstants"
}

func (m *_BacnetConstants) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (bacnetUdpDefaultPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *_BacnetConstants) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BacnetConstantsParse(ctx context.Context, theBytes []byte) (BacnetConstants, error) {
	return BacnetConstantsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BacnetConstantsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BacnetConstants, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BacnetConstants"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BacnetConstants")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (bacnetUdpDefaultPort)
	bacnetUdpDefaultPort, _bacnetUdpDefaultPortErr := readBuffer.ReadUint16("bacnetUdpDefaultPort", 16)
	if _bacnetUdpDefaultPortErr != nil {
		return nil, errors.Wrap(_bacnetUdpDefaultPortErr, "Error parsing 'bacnetUdpDefaultPort' field of BacnetConstants")
	}
	if bacnetUdpDefaultPort != BacnetConstants_BACNETUDPDEFAULTPORT {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", BacnetConstants_BACNETUDPDEFAULTPORT) + " but got " + fmt.Sprintf("%d", bacnetUdpDefaultPort))
	}

	if closeErr := readBuffer.CloseContext("BacnetConstants"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BacnetConstants")
	}

	// Create the instance
	return &_BacnetConstants{}, nil
}

func (m *_BacnetConstants) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BacnetConstants) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BacnetConstants"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BacnetConstants")
	}

	// Const Field (bacnetUdpDefaultPort)
	_bacnetUdpDefaultPortErr := writeBuffer.WriteUint16("bacnetUdpDefaultPort", 16, 47808)
	if _bacnetUdpDefaultPortErr != nil {
		return errors.Wrap(_bacnetUdpDefaultPortErr, "Error serializing 'bacnetUdpDefaultPort' field")
	}

	if popErr := writeBuffer.PopContext("BacnetConstants"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BacnetConstants")
	}
	return nil
}

func (m *_BacnetConstants) isBacnetConstants() bool {
	return true
}

func (m *_BacnetConstants) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
