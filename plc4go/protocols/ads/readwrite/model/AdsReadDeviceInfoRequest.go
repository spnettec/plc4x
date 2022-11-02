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


// AdsReadDeviceInfoRequest is the corresponding interface of AdsReadDeviceInfoRequest
type AdsReadDeviceInfoRequest interface {
	utils.LengthAware
	utils.Serializable
	AmsPacket
}

// AdsReadDeviceInfoRequestExactly can be used when we want exactly this type and not a type which fulfills AdsReadDeviceInfoRequest.
// This is useful for switch cases.
type AdsReadDeviceInfoRequestExactly interface {
	AdsReadDeviceInfoRequest
	isAdsReadDeviceInfoRequest() bool
}

// _AdsReadDeviceInfoRequest is the data-structure of this message
type _AdsReadDeviceInfoRequest struct {
	*_AmsPacket
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_AdsReadDeviceInfoRequest)  GetCommandId() CommandId {
return CommandId_ADS_READ_DEVICE_INFO}

func (m *_AdsReadDeviceInfoRequest)  GetResponse() bool {
return bool(false)}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AdsReadDeviceInfoRequest) InitializeParent(parent AmsPacket , targetAmsNetId AmsNetId , targetAmsPort uint16 , sourceAmsNetId AmsNetId , sourceAmsPort uint16 , errorCode uint32 , invokeId uint32 ) {	m.TargetAmsNetId = targetAmsNetId
	m.TargetAmsPort = targetAmsPort
	m.SourceAmsNetId = sourceAmsNetId
	m.SourceAmsPort = sourceAmsPort
	m.ErrorCode = errorCode
	m.InvokeId = invokeId
}

func (m *_AdsReadDeviceInfoRequest)  GetParent() AmsPacket {
	return m._AmsPacket
}


// NewAdsReadDeviceInfoRequest factory function for _AdsReadDeviceInfoRequest
func NewAdsReadDeviceInfoRequest( targetAmsNetId AmsNetId , targetAmsPort uint16 , sourceAmsNetId AmsNetId , sourceAmsPort uint16 , errorCode uint32 , invokeId uint32 ) *_AdsReadDeviceInfoRequest {
	_result := &_AdsReadDeviceInfoRequest{
    	_AmsPacket: NewAmsPacket(targetAmsNetId, targetAmsPort, sourceAmsNetId, sourceAmsPort, errorCode, invokeId),
	}
	_result._AmsPacket._AmsPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAdsReadDeviceInfoRequest(structType interface{}) AdsReadDeviceInfoRequest {
    if casted, ok := structType.(AdsReadDeviceInfoRequest); ok {
		return casted
	}
	if casted, ok := structType.(*AdsReadDeviceInfoRequest); ok {
		return *casted
	}
	return nil
}

func (m *_AdsReadDeviceInfoRequest) GetTypeName() string {
	return "AdsReadDeviceInfoRequest"
}

func (m *_AdsReadDeviceInfoRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AdsReadDeviceInfoRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}


func (m *_AdsReadDeviceInfoRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AdsReadDeviceInfoRequestParse(readBuffer utils.ReadBuffer) (AdsReadDeviceInfoRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AdsReadDeviceInfoRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AdsReadDeviceInfoRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AdsReadDeviceInfoRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AdsReadDeviceInfoRequest")
	}

	// Create a partially initialized instance
	_child := &_AdsReadDeviceInfoRequest{
		_AmsPacket: &_AmsPacket{
		},
	}
	_child._AmsPacket._AmsPacketChildRequirements = _child
	return _child, nil
}

func (m *_AdsReadDeviceInfoRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AdsReadDeviceInfoRequest) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AdsReadDeviceInfoRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AdsReadDeviceInfoRequest")
		}

		if popErr := writeBuffer.PopContext("AdsReadDeviceInfoRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AdsReadDeviceInfoRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_AdsReadDeviceInfoRequest) isAdsReadDeviceInfoRequest() bool {
	return true
}

func (m *_AdsReadDeviceInfoRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



