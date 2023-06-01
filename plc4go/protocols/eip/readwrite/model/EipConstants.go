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
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const EipConstants_EIPUDPDISCOVERYDEFAULTPORT uint16 = uint16(44818)
const EipConstants_EIPTCPDEFAULTPORT uint16 = uint16(44818)

// EipConstants is the corresponding interface of EipConstants
type EipConstants interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// EipConstantsExactly can be used when we want exactly this type and not a type which fulfills EipConstants.
// This is useful for switch cases.
type EipConstantsExactly interface {
	EipConstants
	isEipConstants() bool
}

// _EipConstants is the data-structure of this message
type _EipConstants struct {
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_EipConstants) GetEipUdpDiscoveryDefaultPort() uint16 {
	return EipConstants_EIPUDPDISCOVERYDEFAULTPORT
}

func (m *_EipConstants) GetEipTcpDefaultPort() uint16 {
	return EipConstants_EIPTCPDEFAULTPORT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewEipConstants factory function for _EipConstants
func NewEipConstants() *_EipConstants {
	return &_EipConstants{}
}

// Deprecated: use the interface for direct cast
func CastEipConstants(structType any) EipConstants {
	if casted, ok := structType.(EipConstants); ok {
		return casted
	}
	if casted, ok := structType.(*EipConstants); ok {
		return *casted
	}
	return nil
}

func (m *_EipConstants) GetTypeName() string {
	return "EipConstants"
}

func (m *_EipConstants) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Const Field (eipUdpDiscoveryDefaultPort)
	lengthInBits += 16

	// Const Field (eipTcpDefaultPort)
	lengthInBits += 16

	return lengthInBits
}

func (m *_EipConstants) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func EipConstantsParse(theBytes []byte) (EipConstants, error) {
	return EipConstantsParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func EipConstantsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (EipConstants, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("EipConstants"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for EipConstants")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (eipUdpDiscoveryDefaultPort)
	eipUdpDiscoveryDefaultPort, _eipUdpDiscoveryDefaultPortErr := readBuffer.ReadUint16("eipUdpDiscoveryDefaultPort", 16)
	if _eipUdpDiscoveryDefaultPortErr != nil {
		return nil, errors.Wrap(_eipUdpDiscoveryDefaultPortErr, "Error parsing 'eipUdpDiscoveryDefaultPort' field of EipConstants")
	}
	if eipUdpDiscoveryDefaultPort != EipConstants_EIPUDPDISCOVERYDEFAULTPORT {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", EipConstants_EIPUDPDISCOVERYDEFAULTPORT) + " but got " + fmt.Sprintf("%d", eipUdpDiscoveryDefaultPort))
	}

	// Const Field (eipTcpDefaultPort)
	eipTcpDefaultPort, _eipTcpDefaultPortErr := readBuffer.ReadUint16("eipTcpDefaultPort", 16)
	if _eipTcpDefaultPortErr != nil {
		return nil, errors.Wrap(_eipTcpDefaultPortErr, "Error parsing 'eipTcpDefaultPort' field of EipConstants")
	}
	if eipTcpDefaultPort != EipConstants_EIPTCPDEFAULTPORT {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", EipConstants_EIPTCPDEFAULTPORT) + " but got " + fmt.Sprintf("%d", eipTcpDefaultPort))
	}

	if closeErr := readBuffer.CloseContext("EipConstants"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for EipConstants")
	}

	// Create the instance
	return &_EipConstants{}, nil
}

func (m *_EipConstants) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_EipConstants) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("EipConstants"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for EipConstants")
	}

	// Const Field (eipUdpDiscoveryDefaultPort)
	_eipUdpDiscoveryDefaultPortErr := writeBuffer.WriteUint16("eipUdpDiscoveryDefaultPort", 16, 44818)
	if _eipUdpDiscoveryDefaultPortErr != nil {
		return errors.Wrap(_eipUdpDiscoveryDefaultPortErr, "Error serializing 'eipUdpDiscoveryDefaultPort' field")
	}

	// Const Field (eipTcpDefaultPort)
	_eipTcpDefaultPortErr := writeBuffer.WriteUint16("eipTcpDefaultPort", 16, 44818)
	if _eipTcpDefaultPortErr != nil {
		return errors.Wrap(_eipTcpDefaultPortErr, "Error serializing 'eipTcpDefaultPort' field")
	}

	if popErr := writeBuffer.PopContext("EipConstants"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for EipConstants")
	}
	return nil
}

func (m *_EipConstants) isEipConstants() bool {
	return true
}

func (m *_EipConstants) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
