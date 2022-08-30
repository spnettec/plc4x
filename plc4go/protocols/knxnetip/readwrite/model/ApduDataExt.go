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


// ApduDataExt is the corresponding interface of ApduDataExt
type ApduDataExt interface {
	utils.LengthAware
	utils.Serializable
	// GetExtApciType returns ExtApciType (discriminator field)
	GetExtApciType() uint8
}

// ApduDataExtExactly can be used when we want exactly this type and not a type which fulfills ApduDataExt.
// This is useful for switch cases.
type ApduDataExtExactly interface {
	ApduDataExt
	isApduDataExt() bool
}

// _ApduDataExt is the data-structure of this message
type _ApduDataExt struct {
	_ApduDataExtChildRequirements

	// Arguments.
	Length uint8
}

type _ApduDataExtChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
	GetExtApciType() uint8
}


type ApduDataExtParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child ApduDataExt, serializeChildFunction func() error) error
	GetTypeName() string
}

type ApduDataExtChild interface {
	utils.Serializable
InitializeParent(parent ApduDataExt )
	GetParent() *ApduDataExt

	GetTypeName() string
	ApduDataExt
}


// NewApduDataExt factory function for _ApduDataExt
func NewApduDataExt( length uint8 ) *_ApduDataExt {
return &_ApduDataExt{ Length: length }
}

// Deprecated: use the interface for direct cast
func CastApduDataExt(structType interface{}) ApduDataExt {
    if casted, ok := structType.(ApduDataExt); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExt); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExt) GetTypeName() string {
	return "ApduDataExt"
}



func (m *_ApduDataExt) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (extApciType)
	lengthInBits += 6;

	return lengthInBits
}

func (m *_ApduDataExt) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ApduDataExtParse(readBuffer utils.ReadBuffer, length uint8) (ApduDataExt, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ApduDataExt"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExt")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Discriminator Field (extApciType) (Used as input to a switch field)
	extApciType, _extApciTypeErr := readBuffer.ReadUint8("extApciType", 6)
	if _extApciTypeErr != nil {
		return nil, errors.Wrap(_extApciTypeErr, "Error parsing 'extApciType' field of ApduDataExt")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type ApduDataExtChildSerializeRequirement interface {
		ApduDataExt
		InitializeParent(ApduDataExt )
		GetParent() ApduDataExt
	}
	var _childTemp interface{}
	var _child ApduDataExtChildSerializeRequirement
	var typeSwitchError error
	switch {
case extApciType == 0x00 : // ApduDataExtOpenRoutingTableRequest
		_childTemp, typeSwitchError = ApduDataExtOpenRoutingTableRequestParse(readBuffer, length)
case extApciType == 0x01 : // ApduDataExtReadRoutingTableRequest
		_childTemp, typeSwitchError = ApduDataExtReadRoutingTableRequestParse(readBuffer, length)
case extApciType == 0x02 : // ApduDataExtReadRoutingTableResponse
		_childTemp, typeSwitchError = ApduDataExtReadRoutingTableResponseParse(readBuffer, length)
case extApciType == 0x03 : // ApduDataExtWriteRoutingTableRequest
		_childTemp, typeSwitchError = ApduDataExtWriteRoutingTableRequestParse(readBuffer, length)
case extApciType == 0x08 : // ApduDataExtReadRouterMemoryRequest
		_childTemp, typeSwitchError = ApduDataExtReadRouterMemoryRequestParse(readBuffer, length)
case extApciType == 0x09 : // ApduDataExtReadRouterMemoryResponse
		_childTemp, typeSwitchError = ApduDataExtReadRouterMemoryResponseParse(readBuffer, length)
case extApciType == 0x0A : // ApduDataExtWriteRouterMemoryRequest
		_childTemp, typeSwitchError = ApduDataExtWriteRouterMemoryRequestParse(readBuffer, length)
case extApciType == 0x0D : // ApduDataExtReadRouterStatusRequest
		_childTemp, typeSwitchError = ApduDataExtReadRouterStatusRequestParse(readBuffer, length)
case extApciType == 0x0E : // ApduDataExtReadRouterStatusResponse
		_childTemp, typeSwitchError = ApduDataExtReadRouterStatusResponseParse(readBuffer, length)
case extApciType == 0x0F : // ApduDataExtWriteRouterStatusRequest
		_childTemp, typeSwitchError = ApduDataExtWriteRouterStatusRequestParse(readBuffer, length)
case extApciType == 0x10 : // ApduDataExtMemoryBitWrite
		_childTemp, typeSwitchError = ApduDataExtMemoryBitWriteParse(readBuffer, length)
case extApciType == 0x11 : // ApduDataExtAuthorizeRequest
		_childTemp, typeSwitchError = ApduDataExtAuthorizeRequestParse(readBuffer, length)
case extApciType == 0x12 : // ApduDataExtAuthorizeResponse
		_childTemp, typeSwitchError = ApduDataExtAuthorizeResponseParse(readBuffer, length)
case extApciType == 0x13 : // ApduDataExtKeyWrite
		_childTemp, typeSwitchError = ApduDataExtKeyWriteParse(readBuffer, length)
case extApciType == 0x14 : // ApduDataExtKeyResponse
		_childTemp, typeSwitchError = ApduDataExtKeyResponseParse(readBuffer, length)
case extApciType == 0x15 : // ApduDataExtPropertyValueRead
		_childTemp, typeSwitchError = ApduDataExtPropertyValueReadParse(readBuffer, length)
case extApciType == 0x16 : // ApduDataExtPropertyValueResponse
		_childTemp, typeSwitchError = ApduDataExtPropertyValueResponseParse(readBuffer, length)
case extApciType == 0x17 : // ApduDataExtPropertyValueWrite
		_childTemp, typeSwitchError = ApduDataExtPropertyValueWriteParse(readBuffer, length)
case extApciType == 0x18 : // ApduDataExtPropertyDescriptionRead
		_childTemp, typeSwitchError = ApduDataExtPropertyDescriptionReadParse(readBuffer, length)
case extApciType == 0x19 : // ApduDataExtPropertyDescriptionResponse
		_childTemp, typeSwitchError = ApduDataExtPropertyDescriptionResponseParse(readBuffer, length)
case extApciType == 0x1A : // ApduDataExtNetworkParameterRead
		_childTemp, typeSwitchError = ApduDataExtNetworkParameterReadParse(readBuffer, length)
case extApciType == 0x1B : // ApduDataExtNetworkParameterResponse
		_childTemp, typeSwitchError = ApduDataExtNetworkParameterResponseParse(readBuffer, length)
case extApciType == 0x1C : // ApduDataExtIndividualAddressSerialNumberRead
		_childTemp, typeSwitchError = ApduDataExtIndividualAddressSerialNumberReadParse(readBuffer, length)
case extApciType == 0x1D : // ApduDataExtIndividualAddressSerialNumberResponse
		_childTemp, typeSwitchError = ApduDataExtIndividualAddressSerialNumberResponseParse(readBuffer, length)
case extApciType == 0x1E : // ApduDataExtIndividualAddressSerialNumberWrite
		_childTemp, typeSwitchError = ApduDataExtIndividualAddressSerialNumberWriteParse(readBuffer, length)
case extApciType == 0x20 : // ApduDataExtDomainAddressWrite
		_childTemp, typeSwitchError = ApduDataExtDomainAddressWriteParse(readBuffer, length)
case extApciType == 0x21 : // ApduDataExtDomainAddressRead
		_childTemp, typeSwitchError = ApduDataExtDomainAddressReadParse(readBuffer, length)
case extApciType == 0x22 : // ApduDataExtDomainAddressResponse
		_childTemp, typeSwitchError = ApduDataExtDomainAddressResponseParse(readBuffer, length)
case extApciType == 0x23 : // ApduDataExtDomainAddressSelectiveRead
		_childTemp, typeSwitchError = ApduDataExtDomainAddressSelectiveReadParse(readBuffer, length)
case extApciType == 0x24 : // ApduDataExtNetworkParameterWrite
		_childTemp, typeSwitchError = ApduDataExtNetworkParameterWriteParse(readBuffer, length)
case extApciType == 0x25 : // ApduDataExtLinkRead
		_childTemp, typeSwitchError = ApduDataExtLinkReadParse(readBuffer, length)
case extApciType == 0x26 : // ApduDataExtLinkResponse
		_childTemp, typeSwitchError = ApduDataExtLinkResponseParse(readBuffer, length)
case extApciType == 0x27 : // ApduDataExtLinkWrite
		_childTemp, typeSwitchError = ApduDataExtLinkWriteParse(readBuffer, length)
case extApciType == 0x28 : // ApduDataExtGroupPropertyValueRead
		_childTemp, typeSwitchError = ApduDataExtGroupPropertyValueReadParse(readBuffer, length)
case extApciType == 0x29 : // ApduDataExtGroupPropertyValueResponse
		_childTemp, typeSwitchError = ApduDataExtGroupPropertyValueResponseParse(readBuffer, length)
case extApciType == 0x2A : // ApduDataExtGroupPropertyValueWrite
		_childTemp, typeSwitchError = ApduDataExtGroupPropertyValueWriteParse(readBuffer, length)
case extApciType == 0x2B : // ApduDataExtGroupPropertyValueInfoReport
		_childTemp, typeSwitchError = ApduDataExtGroupPropertyValueInfoReportParse(readBuffer, length)
case extApciType == 0x2C : // ApduDataExtDomainAddressSerialNumberRead
		_childTemp, typeSwitchError = ApduDataExtDomainAddressSerialNumberReadParse(readBuffer, length)
case extApciType == 0x2D : // ApduDataExtDomainAddressSerialNumberResponse
		_childTemp, typeSwitchError = ApduDataExtDomainAddressSerialNumberResponseParse(readBuffer, length)
case extApciType == 0x2E : // ApduDataExtDomainAddressSerialNumberWrite
		_childTemp, typeSwitchError = ApduDataExtDomainAddressSerialNumberWriteParse(readBuffer, length)
case extApciType == 0x30 : // ApduDataExtFileStreamInfoReport
		_childTemp, typeSwitchError = ApduDataExtFileStreamInfoReportParse(readBuffer, length)
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [extApciType=%v]", extApciType)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of ApduDataExt")
	}
	_child = _childTemp.(ApduDataExtChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("ApduDataExt"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExt")
	}

	// Finish initializing
_child.InitializeParent(_child )
	return _child, nil
}

func (pm *_ApduDataExt) SerializeParent(writeBuffer utils.WriteBuffer, child ApduDataExt, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("ApduDataExt"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ApduDataExt")
	}

	// Discriminator Field (extApciType) (Used as input to a switch field)
	extApciType := uint8(child.GetExtApciType())
	_extApciTypeErr := writeBuffer.WriteUint8("extApciType", 6, (extApciType))

	if _extApciTypeErr != nil {
		return errors.Wrap(_extApciTypeErr, "Error serializing 'extApciType' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("ApduDataExt"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ApduDataExt")
	}
	return nil
}


////
// Arguments Getter

func (m *_ApduDataExt) GetLength() uint8 {
	return m.Length
}
//
////

func (m *_ApduDataExt) isApduDataExt() bool {
	return true
}

func (m *_ApduDataExt) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



