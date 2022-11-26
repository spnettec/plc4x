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
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const RequestTermination_CR byte = 0x0D

// RequestTermination is the corresponding interface of RequestTermination
type RequestTermination interface {
	utils.LengthAware
	utils.Serializable
}

// RequestTerminationExactly can be used when we want exactly this type and not a type which fulfills RequestTermination.
// This is useful for switch cases.
type RequestTerminationExactly interface {
	RequestTermination
	isRequestTermination() bool
}

// _RequestTermination is the data-structure of this message
type _RequestTermination struct {
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_RequestTermination) GetCr() byte {
	return RequestTermination_CR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRequestTermination factory function for _RequestTermination
func NewRequestTermination() *_RequestTermination {
	return &_RequestTermination{}
}

// Deprecated: use the interface for direct cast
func CastRequestTermination(structType interface{}) RequestTermination {
	if casted, ok := structType.(RequestTermination); ok {
		return casted
	}
	if casted, ok := structType.(*RequestTermination); ok {
		return *casted
	}
	return nil
}

func (m *_RequestTermination) GetTypeName() string {
	return "RequestTermination"
}

func (m *_RequestTermination) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_RequestTermination) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Const Field (cr)
	lengthInBits += 8

	return lengthInBits
}

func (m *_RequestTermination) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func RequestTerminationParse(theBytes []byte) (RequestTermination, error) {
	return RequestTerminationParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func RequestTerminationParseWithBuffer(readBuffer utils.ReadBuffer) (RequestTermination, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestTermination"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestTermination")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (cr)
	cr, _crErr := readBuffer.ReadByte("cr")
	if _crErr != nil {
		return nil, errors.Wrap(_crErr, "Error parsing 'cr' field of RequestTermination")
	}
	if cr != RequestTermination_CR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", RequestTermination_CR) + " but got " + fmt.Sprintf("%d", cr))
	}

	if closeErr := readBuffer.CloseContext("RequestTermination"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestTermination")
	}

	// Create the instance
	return &_RequestTermination{}, nil
}

func (m *_RequestTermination) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RequestTermination) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("RequestTermination"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for RequestTermination")
	}

	// Const Field (cr)
	_crErr := writeBuffer.WriteByte("cr", 0x0D)
	if _crErr != nil {
		return errors.Wrap(_crErr, "Error serializing 'cr' field")
	}

	if popErr := writeBuffer.PopContext("RequestTermination"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for RequestTermination")
	}
	return nil
}

func (m *_RequestTermination) isRequestTermination() bool {
	return true
}

func (m *_RequestTermination) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
