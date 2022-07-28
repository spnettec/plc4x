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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// Constant values.
const RequestNull_NULLINDICATOR uint32 = 0x6E756C6C

// RequestNull is the corresponding interface of RequestNull
type RequestNull interface {
	utils.LengthAware
	utils.Serializable
	Request
}

// RequestNullExactly can be used when we want exactly this type and not a type which fulfills RequestNull.
// This is useful for switch cases.
type RequestNullExactly interface {
	RequestNull
	isRequestNull() bool
}

// _RequestNull is the data-structure of this message
type _RequestNull struct {
	*_Request
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RequestNull) InitializeParent(parent Request, peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination) {
	m.PeekedByte = peekedByte
	m.StartingCR = startingCR
	m.ResetMode = resetMode
	m.SecondPeek = secondPeek
	m.Termination = termination
}

func (m *_RequestNull) GetParent() Request {
	return m._Request
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_RequestNull) GetNullIndicator() uint32 {
	return RequestNull_NULLINDICATOR
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRequestNull factory function for _RequestNull
func NewRequestNull(peekedByte RequestType, startingCR *RequestType, resetMode *RequestType, secondPeek RequestType, termination RequestTermination, cBusOptions CBusOptions) *_RequestNull {
	_result := &_RequestNull{
		_Request: NewRequest(peekedByte, startingCR, resetMode, secondPeek, termination, cBusOptions),
	}
	_result._Request._RequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRequestNull(structType interface{}) RequestNull {
	if casted, ok := structType.(RequestNull); ok {
		return casted
	}
	if casted, ok := structType.(*RequestNull); ok {
		return *casted
	}
	return nil
}

func (m *_RequestNull) GetTypeName() string {
	return "RequestNull"
}

func (m *_RequestNull) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_RequestNull) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Const Field (nullIndicator)
	lengthInBits += 32

	return lengthInBits
}

func (m *_RequestNull) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func RequestNullParse(readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (RequestNull, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestNull"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestNull")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (nullIndicator)
	nullIndicator, _nullIndicatorErr := readBuffer.ReadUint32("nullIndicator", 32)
	if _nullIndicatorErr != nil {
		return nil, errors.Wrap(_nullIndicatorErr, "Error parsing 'nullIndicator' field of RequestNull")
	}
	if nullIndicator != RequestNull_NULLINDICATOR {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", RequestNull_NULLINDICATOR) + " but got " + fmt.Sprintf("%d", nullIndicator))
	}

	if closeErr := readBuffer.CloseContext("RequestNull"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestNull")
	}

	// Create a partially initialized instance
	_child := &_RequestNull{
		_Request: &_Request{
			CBusOptions: cBusOptions,
		},
	}
	_child._Request._RequestChildRequirements = _child
	return _child, nil
}

func (m *_RequestNull) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RequestNull"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RequestNull")
		}

		// Const Field (nullIndicator)
		_nullIndicatorErr := writeBuffer.WriteUint32("nullIndicator", 32, 0x6E756C6C)
		if _nullIndicatorErr != nil {
			return errors.Wrap(_nullIndicatorErr, "Error serializing 'nullIndicator' field")
		}

		if popErr := writeBuffer.PopContext("RequestNull"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RequestNull")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_RequestNull) isRequestNull() bool {
	return true
}

func (m *_RequestNull) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
