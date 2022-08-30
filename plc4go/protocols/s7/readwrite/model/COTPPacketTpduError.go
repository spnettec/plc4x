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


// COTPPacketTpduError is the corresponding interface of COTPPacketTpduError
type COTPPacketTpduError interface {
	utils.LengthAware
	utils.Serializable
	COTPPacket
	// GetDestinationReference returns DestinationReference (property field)
	GetDestinationReference() uint16
	// GetRejectCause returns RejectCause (property field)
	GetRejectCause() uint8
}

// COTPPacketTpduErrorExactly can be used when we want exactly this type and not a type which fulfills COTPPacketTpduError.
// This is useful for switch cases.
type COTPPacketTpduErrorExactly interface {
	COTPPacketTpduError
	isCOTPPacketTpduError() bool
}

// _COTPPacketTpduError is the data-structure of this message
type _COTPPacketTpduError struct {
	*_COTPPacket
        DestinationReference uint16
        RejectCause uint8
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_COTPPacketTpduError)  GetTpduCode() uint8 {
return 0x70}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_COTPPacketTpduError) InitializeParent(parent COTPPacket , parameters []COTPParameter , payload S7Message ) {	m.Parameters = parameters
	m.Payload = payload
}

func (m *_COTPPacketTpduError)  GetParent() COTPPacket {
	return m._COTPPacket
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_COTPPacketTpduError) GetDestinationReference() uint16 {
	return m.DestinationReference
}

func (m *_COTPPacketTpduError) GetRejectCause() uint8 {
	return m.RejectCause
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewCOTPPacketTpduError factory function for _COTPPacketTpduError
func NewCOTPPacketTpduError( destinationReference uint16 , rejectCause uint8 , parameters []COTPParameter , payload S7Message , cotpLen uint16 ) *_COTPPacketTpduError {
	_result := &_COTPPacketTpduError{
		DestinationReference: destinationReference,
		RejectCause: rejectCause,
    	_COTPPacket: NewCOTPPacket(parameters, payload, cotpLen),
	}
	_result._COTPPacket._COTPPacketChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCOTPPacketTpduError(structType interface{}) COTPPacketTpduError {
    if casted, ok := structType.(COTPPacketTpduError); ok {
		return casted
	}
	if casted, ok := structType.(*COTPPacketTpduError); ok {
		return *casted
	}
	return nil
}

func (m *_COTPPacketTpduError) GetTypeName() string {
	return "COTPPacketTpduError"
}

func (m *_COTPPacketTpduError) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_COTPPacketTpduError) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (destinationReference)
	lengthInBits += 16;

	// Simple field (rejectCause)
	lengthInBits += 8;

	return lengthInBits
}


func (m *_COTPPacketTpduError) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func COTPPacketTpduErrorParse(readBuffer utils.ReadBuffer, cotpLen uint16) (COTPPacketTpduError, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("COTPPacketTpduError"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for COTPPacketTpduError")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (destinationReference)
_destinationReference, _destinationReferenceErr := readBuffer.ReadUint16("destinationReference", 16)
	if _destinationReferenceErr != nil {
		return nil, errors.Wrap(_destinationReferenceErr, "Error parsing 'destinationReference' field of COTPPacketTpduError")
	}
	destinationReference := _destinationReference

	// Simple Field (rejectCause)
_rejectCause, _rejectCauseErr := readBuffer.ReadUint8("rejectCause", 8)
	if _rejectCauseErr != nil {
		return nil, errors.Wrap(_rejectCauseErr, "Error parsing 'rejectCause' field of COTPPacketTpduError")
	}
	rejectCause := _rejectCause

	if closeErr := readBuffer.CloseContext("COTPPacketTpduError"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for COTPPacketTpduError")
	}

	// Create a partially initialized instance
	_child := &_COTPPacketTpduError{
		_COTPPacket: &_COTPPacket{
			CotpLen: cotpLen,
		},
		DestinationReference: destinationReference,
		RejectCause: rejectCause,
	}
	_child._COTPPacket._COTPPacketChildRequirements = _child
	return _child, nil
}

func (m *_COTPPacketTpduError) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("COTPPacketTpduError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for COTPPacketTpduError")
		}

	// Simple Field (destinationReference)
	destinationReference := uint16(m.GetDestinationReference())
	_destinationReferenceErr := writeBuffer.WriteUint16("destinationReference", 16, (destinationReference))
	if _destinationReferenceErr != nil {
		return errors.Wrap(_destinationReferenceErr, "Error serializing 'destinationReference' field")
	}

	// Simple Field (rejectCause)
	rejectCause := uint8(m.GetRejectCause())
	_rejectCauseErr := writeBuffer.WriteUint8("rejectCause", 8, (rejectCause))
	if _rejectCauseErr != nil {
		return errors.Wrap(_rejectCauseErr, "Error serializing 'rejectCause' field")
	}

		if popErr := writeBuffer.PopContext("COTPPacketTpduError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for COTPPacketTpduError")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_COTPPacketTpduError) isCOTPPacketTpduError() bool {
	return true
}

func (m *_COTPPacketTpduError) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



