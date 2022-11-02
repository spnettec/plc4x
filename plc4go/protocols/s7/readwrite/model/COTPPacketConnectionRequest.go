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


// COTPPacketConnectionRequest is the corresponding interface of COTPPacketConnectionRequest
type COTPPacketConnectionRequest interface {
	utils.LengthAware
	utils.Serializable
	COTPPacket
	// GetDestinationReference returns DestinationReference (property field)
	GetDestinationReference() uint16
	// GetSourceReference returns SourceReference (property field)
	GetSourceReference() uint16
	// GetProtocolClass returns ProtocolClass (property field)
	GetProtocolClass() COTPProtocolClass
}

// COTPPacketConnectionRequestExactly can be used when we want exactly this type and not a type which fulfills COTPPacketConnectionRequest.
// This is useful for switch cases.
type COTPPacketConnectionRequestExactly interface {
	COTPPacketConnectionRequest
	isCOTPPacketConnectionRequest() bool
}

// _COTPPacketConnectionRequest is the data-structure of this message
type _COTPPacketConnectionRequest struct {
	*_COTPPacket
        DestinationReference uint16
        SourceReference uint16
        ProtocolClass COTPProtocolClass
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPPacketConnectionRequest)  GetTpduCode() uint8 {
return 0xE0}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPPacketConnectionRequest) InitializeParent(parent COTPPacket , parameters []COTPParameter , payload S7Message ) {	m.Parameters = parameters
	m.Payload = payload
}

func (m *_COTPPacketConnectionRequest)  GetParent() COTPPacket {
	return m._COTPPacket
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPPacketConnectionRequest) GetDestinationReference() uint16 {
	return m.DestinationReference
}

func (m *_COTPPacketConnectionRequest) GetSourceReference() uint16 {
	return m.SourceReference
}

func (m *_COTPPacketConnectionRequest) GetProtocolClass() COTPProtocolClass {
	return m.ProtocolClass
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewCOTPPacketConnectionRequest factory function for _COTPPacketConnectionRequest
func NewCOTPPacketConnectionRequest( destinationReference uint16 , sourceReference uint16 , protocolClass COTPProtocolClass , parameters []COTPParameter , payload S7Message , cotpLen uint16 ) *_COTPPacketConnectionRequest {
	_result := &_COTPPacketConnectionRequest{
		DestinationReference: destinationReference,
		SourceReference: sourceReference,
		ProtocolClass: protocolClass,
    	_COTPPacket: NewCOTPPacket(parameters, payload, cotpLen),
	}
	_result._COTPPacket._COTPPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCOTPPacketConnectionRequest(structType interface{}) COTPPacketConnectionRequest {
    if casted, ok := structType.(COTPPacketConnectionRequest); ok {
		return casted
	}
	if casted, ok := structType.(*COTPPacketConnectionRequest); ok {
		return *casted
	}
	return nil
}

func (m *_COTPPacketConnectionRequest) GetTypeName() string {
	return "COTPPacketConnectionRequest"
}

func (m *_COTPPacketConnectionRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_COTPPacketConnectionRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (destinationReference)
	lengthInBits += 16;

	// Simple field (sourceReference)
	lengthInBits += 16;

	// Simple field (protocolClass)
	lengthInBits += 8

	return lengthInBits
}


func (m *_COTPPacketConnectionRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPPacketConnectionRequestParse(readBuffer utils.ReadBuffer, cotpLen uint16) (COTPPacketConnectionRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPPacketConnectionRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPPacketConnectionRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (destinationReference)
_destinationReference, _destinationReferenceErr := readBuffer.ReadUint16("destinationReference", 16)
	if _destinationReferenceErr != nil {
		return nil, errors.Wrap(_destinationReferenceErr, "Error parsing 'destinationReference' field of COTPPacketConnectionRequest")
	}
	destinationReference := _destinationReference

	// Simple Field (sourceReference)
_sourceReference, _sourceReferenceErr := readBuffer.ReadUint16("sourceReference", 16)
	if _sourceReferenceErr != nil {
		return nil, errors.Wrap(_sourceReferenceErr, "Error parsing 'sourceReference' field of COTPPacketConnectionRequest")
	}
	sourceReference := _sourceReference

	// Simple Field (protocolClass)
	if pullErr := readBuffer.PullContext("protocolClass"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for protocolClass")
	}
_protocolClass, _protocolClassErr := COTPProtocolClassParse(readBuffer)
	if _protocolClassErr != nil {
		return nil, errors.Wrap(_protocolClassErr, "Error parsing 'protocolClass' field of COTPPacketConnectionRequest")
	}
	protocolClass := _protocolClass
	if closeErr := readBuffer.CloseContext("protocolClass"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for protocolClass")
	}

	if closeErr := readBuffer.CloseContext("COTPPacketConnectionRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPPacketConnectionRequest")
	}

	// Create a partially initialized instance
	_child := &_COTPPacketConnectionRequest{
		_COTPPacket: &_COTPPacket{
			CotpLen: cotpLen,
		},
		DestinationReference: destinationReference,
		SourceReference: sourceReference,
		ProtocolClass: protocolClass,
	}
	_child._COTPPacket._COTPPacketChildRequirements = _child
	return _child, nil
}

func (m *_COTPPacketConnectionRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_COTPPacketConnectionRequest) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPPacketConnectionRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPPacketConnectionRequest")
		}

	// Simple Field (destinationReference)
	destinationReference := uint16(m.GetDestinationReference())
	_destinationReferenceErr := writeBuffer.WriteUint16("destinationReference", 16, (destinationReference))
	if _destinationReferenceErr != nil {
		return errors.Wrap(_destinationReferenceErr, "Error serializing 'destinationReference' field")
	}

	// Simple Field (sourceReference)
	sourceReference := uint16(m.GetSourceReference())
	_sourceReferenceErr := writeBuffer.WriteUint16("sourceReference", 16, (sourceReference))
	if _sourceReferenceErr != nil {
		return errors.Wrap(_sourceReferenceErr, "Error serializing 'sourceReference' field")
	}

	// Simple Field (protocolClass)
	if pushErr := writeBuffer.PushContext("protocolClass"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for protocolClass")
	}
	_protocolClassErr := writeBuffer.WriteSerializable(m.GetProtocolClass())
	if popErr := writeBuffer.PopContext("protocolClass"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for protocolClass")
	}
	if _protocolClassErr != nil {
		return errors.Wrap(_protocolClassErr, "Error serializing 'protocolClass' field")
	}

		if popErr := writeBuffer.PopContext("COTPPacketConnectionRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPPacketConnectionRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_COTPPacketConnectionRequest) isCOTPPacketConnectionRequest() bool {
	return true
}

func (m *_COTPPacketConnectionRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



