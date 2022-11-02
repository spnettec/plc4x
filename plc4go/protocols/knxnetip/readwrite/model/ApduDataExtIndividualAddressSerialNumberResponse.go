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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// ApduDataExtIndividualAddressSerialNumberResponse is the corresponding interface of ApduDataExtIndividualAddressSerialNumberResponse
type ApduDataExtIndividualAddressSerialNumberResponse interface {
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtIndividualAddressSerialNumberResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtIndividualAddressSerialNumberResponse.
// This is useful for switch cases.
type ApduDataExtIndividualAddressSerialNumberResponseExactly interface {
	ApduDataExtIndividualAddressSerialNumberResponse
	isApduDataExtIndividualAddressSerialNumberResponse() bool
}

// _ApduDataExtIndividualAddressSerialNumberResponse is the data-structure of this message
type _ApduDataExtIndividualAddressSerialNumberResponse struct {
	*_ApduDataExt
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtIndividualAddressSerialNumberResponse)  GetExtApciType() uint8 {
return 0x1D}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) InitializeParent(parent ApduDataExt ) {}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse)  GetParent() ApduDataExt {
	return m._ApduDataExt
}


// NewApduDataExtIndividualAddressSerialNumberResponse factory function for _ApduDataExtIndividualAddressSerialNumberResponse
func NewApduDataExtIndividualAddressSerialNumberResponse( length uint8 ) *_ApduDataExtIndividualAddressSerialNumberResponse {
	_result := &_ApduDataExtIndividualAddressSerialNumberResponse{
    	_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtIndividualAddressSerialNumberResponse(structType interface{}) ApduDataExtIndividualAddressSerialNumberResponse {
    if casted, ok := structType.(ApduDataExtIndividualAddressSerialNumberResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtIndividualAddressSerialNumberResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) GetTypeName() string {
	return "ApduDataExtIndividualAddressSerialNumberResponse"
}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}


func (m *_ApduDataExtIndividualAddressSerialNumberResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtIndividualAddressSerialNumberResponseParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExtIndividualAddressSerialNumberResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExtIndividualAddressSerialNumberResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtIndividualAddressSerialNumberResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtIndividualAddressSerialNumberResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtIndividualAddressSerialNumberResponse")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtIndividualAddressSerialNumberResponse{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtIndividualAddressSerialNumberResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtIndividualAddressSerialNumberResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtIndividualAddressSerialNumberResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtIndividualAddressSerialNumberResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_ApduDataExtIndividualAddressSerialNumberResponse) isApduDataExtIndividualAddressSerialNumberResponse() bool {
	return true
}

func (m *_ApduDataExtIndividualAddressSerialNumberResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



