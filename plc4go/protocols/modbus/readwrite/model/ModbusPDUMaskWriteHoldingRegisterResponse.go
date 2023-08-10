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

// ModbusPDUMaskWriteHoldingRegisterResponse is the corresponding interface of ModbusPDUMaskWriteHoldingRegisterResponse
type ModbusPDUMaskWriteHoldingRegisterResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetReferenceAddress returns ReferenceAddress (property field)
	GetReferenceAddress() uint16
	// GetAndMask returns AndMask (property field)
	GetAndMask() uint16
	// GetOrMask returns OrMask (property field)
	GetOrMask() uint16
}

// ModbusPDUMaskWriteHoldingRegisterResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUMaskWriteHoldingRegisterResponse.
// This is useful for switch cases.
type ModbusPDUMaskWriteHoldingRegisterResponseExactly interface {
	ModbusPDUMaskWriteHoldingRegisterResponse
	isModbusPDUMaskWriteHoldingRegisterResponse() bool
}

// _ModbusPDUMaskWriteHoldingRegisterResponse is the data-structure of this message
type _ModbusPDUMaskWriteHoldingRegisterResponse struct {
	*_ModbusPDU
	ReferenceAddress uint16
	AndMask          uint16
	OrMask           uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetErrorFlag() bool {
	return bool(false)
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetFunctionFlag() uint8 {
	return 0x16
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetResponse() bool {
	return bool(true)
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) InitializeParent(parent ModbusPDU) {}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetParent() ModbusPDU {
	return m._ModbusPDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetReferenceAddress() uint16 {
	return m.ReferenceAddress
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetAndMask() uint16 {
	return m.AndMask
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetOrMask() uint16 {
	return m.OrMask
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusPDUMaskWriteHoldingRegisterResponse factory function for _ModbusPDUMaskWriteHoldingRegisterResponse
func NewModbusPDUMaskWriteHoldingRegisterResponse(referenceAddress uint16, andMask uint16, orMask uint16) *_ModbusPDUMaskWriteHoldingRegisterResponse {
	_result := &_ModbusPDUMaskWriteHoldingRegisterResponse{
		ReferenceAddress: referenceAddress,
		AndMask:          andMask,
		OrMask:           orMask,
		_ModbusPDU:       NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUMaskWriteHoldingRegisterResponse(structType any) ModbusPDUMaskWriteHoldingRegisterResponse {
	if casted, ok := structType.(ModbusPDUMaskWriteHoldingRegisterResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUMaskWriteHoldingRegisterResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetTypeName() string {
	return "ModbusPDUMaskWriteHoldingRegisterResponse"
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (referenceAddress)
	lengthInBits += 16

	// Simple field (andMask)
	lengthInBits += 16

	// Simple field (orMask)
	lengthInBits += 16

	return lengthInBits
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUMaskWriteHoldingRegisterResponseParse(ctx context.Context, theBytes []byte, response bool) (ModbusPDUMaskWriteHoldingRegisterResponse, error) {
	return ModbusPDUMaskWriteHoldingRegisterResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUMaskWriteHoldingRegisterResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUMaskWriteHoldingRegisterResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ModbusPDUMaskWriteHoldingRegisterResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUMaskWriteHoldingRegisterResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (referenceAddress)
	_referenceAddress, _referenceAddressErr := readBuffer.ReadUint16("referenceAddress", 16)
	if _referenceAddressErr != nil {
		return nil, errors.Wrap(_referenceAddressErr, "Error parsing 'referenceAddress' field of ModbusPDUMaskWriteHoldingRegisterResponse")
	}
	referenceAddress := _referenceAddress

	// Simple Field (andMask)
	_andMask, _andMaskErr := readBuffer.ReadUint16("andMask", 16)
	if _andMaskErr != nil {
		return nil, errors.Wrap(_andMaskErr, "Error parsing 'andMask' field of ModbusPDUMaskWriteHoldingRegisterResponse")
	}
	andMask := _andMask

	// Simple Field (orMask)
	_orMask, _orMaskErr := readBuffer.ReadUint16("orMask", 16)
	if _orMaskErr != nil {
		return nil, errors.Wrap(_orMaskErr, "Error parsing 'orMask' field of ModbusPDUMaskWriteHoldingRegisterResponse")
	}
	orMask := _orMask

	if closeErr := readBuffer.CloseContext("ModbusPDUMaskWriteHoldingRegisterResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUMaskWriteHoldingRegisterResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUMaskWriteHoldingRegisterResponse{
		_ModbusPDU:       &_ModbusPDU{},
		ReferenceAddress: referenceAddress,
		AndMask:          andMask,
		OrMask:           orMask,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUMaskWriteHoldingRegisterResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUMaskWriteHoldingRegisterResponse")
		}

		// Simple Field (referenceAddress)
		referenceAddress := uint16(m.GetReferenceAddress())
		_referenceAddressErr := writeBuffer.WriteUint16("referenceAddress", 16, (referenceAddress))
		if _referenceAddressErr != nil {
			return errors.Wrap(_referenceAddressErr, "Error serializing 'referenceAddress' field")
		}

		// Simple Field (andMask)
		andMask := uint16(m.GetAndMask())
		_andMaskErr := writeBuffer.WriteUint16("andMask", 16, (andMask))
		if _andMaskErr != nil {
			return errors.Wrap(_andMaskErr, "Error serializing 'andMask' field")
		}

		// Simple Field (orMask)
		orMask := uint16(m.GetOrMask())
		_orMaskErr := writeBuffer.WriteUint16("orMask", 16, (orMask))
		if _orMaskErr != nil {
			return errors.Wrap(_orMaskErr, "Error serializing 'orMask' field")
		}

		if popErr := writeBuffer.PopContext("ModbusPDUMaskWriteHoldingRegisterResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUMaskWriteHoldingRegisterResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) isModbusPDUMaskWriteHoldingRegisterResponse() bool {
	return true
}

func (m *_ModbusPDUMaskWriteHoldingRegisterResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
