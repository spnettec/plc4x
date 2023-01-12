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


// BACnetConfirmedServiceRequestRequestKey is the corresponding interface of BACnetConfirmedServiceRequestRequestKey
type BACnetConfirmedServiceRequestRequestKey interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetBytesOfRemovedService returns BytesOfRemovedService (property field)
	GetBytesOfRemovedService() []byte
}

// BACnetConfirmedServiceRequestRequestKeyExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestRequestKey.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestRequestKeyExactly interface {
	BACnetConfirmedServiceRequestRequestKey
	isBACnetConfirmedServiceRequestRequestKey() bool
}

// _BACnetConfirmedServiceRequestRequestKey is the data-structure of this message
type _BACnetConfirmedServiceRequestRequestKey struct {
	*_BACnetConfirmedServiceRequest
        BytesOfRemovedService []byte

	// Arguments.
	ServiceRequestPayloadLength uint32
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestRequestKey)  GetServiceChoice() BACnetConfirmedServiceChoice {
return BACnetConfirmedServiceChoice_REQUEST_KEY}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestRequestKey) InitializeParent(parent BACnetConfirmedServiceRequest ) {}

func (m *_BACnetConfirmedServiceRequestRequestKey)  GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestRequestKey) GetBytesOfRemovedService() []byte {
	return m.BytesOfRemovedService
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConfirmedServiceRequestRequestKey factory function for _BACnetConfirmedServiceRequestRequestKey
func NewBACnetConfirmedServiceRequestRequestKey( bytesOfRemovedService []byte , serviceRequestLength uint32 , serviceRequestPayloadLength uint32 ) *_BACnetConfirmedServiceRequestRequestKey {
	_result := &_BACnetConfirmedServiceRequestRequestKey{
		BytesOfRemovedService: bytesOfRemovedService,
    	_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestRequestKey(structType interface{}) BACnetConfirmedServiceRequestRequestKey {
    if casted, ok := structType.(BACnetConfirmedServiceRequestRequestKey); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestRequestKey); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestRequestKey) GetTypeName() string {
	return "BACnetConfirmedServiceRequestRequestKey"
}

func (m *_BACnetConfirmedServiceRequestRequestKey) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestRequestKey) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.BytesOfRemovedService) > 0 {
		lengthInBits += 8 * uint16(len(m.BytesOfRemovedService))
	}

	return lengthInBits
}


func (m *_BACnetConfirmedServiceRequestRequestKey) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestRequestKeyParse(theBytes []byte, serviceRequestLength uint32, serviceRequestPayloadLength uint32) (BACnetConfirmedServiceRequestRequestKey, error) {
	return BACnetConfirmedServiceRequestRequestKeyParseWithBuffer(utils.NewReadBufferByteBased(theBytes), serviceRequestLength, serviceRequestPayloadLength)
}

func BACnetConfirmedServiceRequestRequestKeyParseWithBuffer(readBuffer utils.ReadBuffer, serviceRequestLength uint32, serviceRequestPayloadLength uint32) (BACnetConfirmedServiceRequestRequestKey, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestRequestKey"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestRequestKey")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (bytesOfRemovedService)
	numberOfBytesbytesOfRemovedService := int(serviceRequestPayloadLength)
	bytesOfRemovedService, _readArrayErr := readBuffer.ReadByteArray("bytesOfRemovedService", numberOfBytesbytesOfRemovedService)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'bytesOfRemovedService' field of BACnetConfirmedServiceRequestRequestKey")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestRequestKey"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestRequestKey")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestRequestKey{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		BytesOfRemovedService: bytesOfRemovedService,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestRequestKey) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestRequestKey) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestRequestKey"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestRequestKey")
		}

	// Array Field (bytesOfRemovedService)
	// Byte Array field (bytesOfRemovedService)
	if err := writeBuffer.WriteByteArray("bytesOfRemovedService", m.GetBytesOfRemovedService()); err != nil {
		return errors.Wrap(err, "Error serializing 'bytesOfRemovedService' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestRequestKey"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestRequestKey")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


////
// Arguments Getter

func (m *_BACnetConfirmedServiceRequestRequestKey) GetServiceRequestPayloadLength() uint32 {
	return m.ServiceRequestPayloadLength
}
//
////

func (m *_BACnetConfirmedServiceRequestRequestKey) isBACnetConfirmedServiceRequestRequestKey() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestRequestKey) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



