/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// APDUConfirmedRequest is the data-structure of this message
type APDUConfirmedRequest struct {
	*APDU
	SegmentedMessage          bool
	MoreFollows               bool
	SegmentedResponseAccepted bool
	MaxSegmentsAccepted       MaxSegmentsAccepted
	MaxApduLengthAccepted     MaxApduLengthAccepted
	InvokeId                  uint8
	SequenceNumber            *uint8
	ProposedWindowSize        *uint8
	ServiceRequest            *BACnetConfirmedServiceRequest
	SegmentServiceChoice      *uint8
	Segment                   []byte

	// Arguments.
	ApduLength uint16
}

// IAPDUConfirmedRequest is the corresponding interface of APDUConfirmedRequest
type IAPDUConfirmedRequest interface {
	IAPDU
	// GetSegmentedMessage returns SegmentedMessage (property field)
	GetSegmentedMessage() bool
	// GetMoreFollows returns MoreFollows (property field)
	GetMoreFollows() bool
	// GetSegmentedResponseAccepted returns SegmentedResponseAccepted (property field)
	GetSegmentedResponseAccepted() bool
	// GetMaxSegmentsAccepted returns MaxSegmentsAccepted (property field)
	GetMaxSegmentsAccepted() MaxSegmentsAccepted
	// GetMaxApduLengthAccepted returns MaxApduLengthAccepted (property field)
	GetMaxApduLengthAccepted() MaxApduLengthAccepted
	// GetInvokeId returns InvokeId (property field)
	GetInvokeId() uint8
	// GetSequenceNumber returns SequenceNumber (property field)
	GetSequenceNumber() *uint8
	// GetProposedWindowSize returns ProposedWindowSize (property field)
	GetProposedWindowSize() *uint8
	// GetServiceRequest returns ServiceRequest (property field)
	GetServiceRequest() *BACnetConfirmedServiceRequest
	// GetSegmentServiceChoice returns SegmentServiceChoice (property field)
	GetSegmentServiceChoice() *uint8
	// GetSegment returns Segment (property field)
	GetSegment() []byte
	// GetApduHeaderReduction returns ApduHeaderReduction (virtual field)
	GetApduHeaderReduction() uint16
	// GetSegmentReduction returns SegmentReduction (virtual field)
	GetSegmentReduction() uint16
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *APDUConfirmedRequest) GetApduType() ApduType {
	return ApduType_CONFIRMED_REQUEST_PDU
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *APDUConfirmedRequest) InitializeParent(parent *APDU) {}

func (m *APDUConfirmedRequest) GetParent() *APDU {
	return m.APDU
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *APDUConfirmedRequest) GetSegmentedMessage() bool {
	return m.SegmentedMessage
}

func (m *APDUConfirmedRequest) GetMoreFollows() bool {
	return m.MoreFollows
}

func (m *APDUConfirmedRequest) GetSegmentedResponseAccepted() bool {
	return m.SegmentedResponseAccepted
}

func (m *APDUConfirmedRequest) GetMaxSegmentsAccepted() MaxSegmentsAccepted {
	return m.MaxSegmentsAccepted
}

func (m *APDUConfirmedRequest) GetMaxApduLengthAccepted() MaxApduLengthAccepted {
	return m.MaxApduLengthAccepted
}

func (m *APDUConfirmedRequest) GetInvokeId() uint8 {
	return m.InvokeId
}

func (m *APDUConfirmedRequest) GetSequenceNumber() *uint8 {
	return m.SequenceNumber
}

func (m *APDUConfirmedRequest) GetProposedWindowSize() *uint8 {
	return m.ProposedWindowSize
}

func (m *APDUConfirmedRequest) GetServiceRequest() *BACnetConfirmedServiceRequest {
	return m.ServiceRequest
}

func (m *APDUConfirmedRequest) GetSegmentServiceChoice() *uint8 {
	return m.SegmentServiceChoice
}

func (m *APDUConfirmedRequest) GetSegment() []byte {
	return m.Segment
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *APDUConfirmedRequest) GetApduHeaderReduction() uint16 {
	sequenceNumber := m.SequenceNumber
	_ = sequenceNumber
	proposedWindowSize := m.ProposedWindowSize
	_ = proposedWindowSize
	serviceRequest := m.ServiceRequest
	_ = serviceRequest
	segmentServiceChoice := m.SegmentServiceChoice
	_ = segmentServiceChoice
	return uint16(uint16(uint16(3)) + uint16(uint16(utils.InlineIf(m.GetSegmentedMessage(), func() interface{} { return uint16(uint16(2)) }, func() interface{} { return uint16(uint16(0)) }).(uint16))))
}

func (m *APDUConfirmedRequest) GetSegmentReduction() uint16 {
	sequenceNumber := m.SequenceNumber
	_ = sequenceNumber
	proposedWindowSize := m.ProposedWindowSize
	_ = proposedWindowSize
	serviceRequest := m.ServiceRequest
	_ = serviceRequest
	segmentServiceChoice := m.SegmentServiceChoice
	_ = segmentServiceChoice
	return uint16(utils.InlineIf(bool(bool((m.GetSegmentServiceChoice()) != (nil))), func() interface{} { return uint16(uint16(uint16(m.GetApduHeaderReduction()) + uint16(uint16(1)))) }, func() interface{} { return uint16(m.GetApduHeaderReduction()) }).(uint16))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewAPDUConfirmedRequest factory function for APDUConfirmedRequest
func NewAPDUConfirmedRequest(segmentedMessage bool, moreFollows bool, segmentedResponseAccepted bool, maxSegmentsAccepted MaxSegmentsAccepted, maxApduLengthAccepted MaxApduLengthAccepted, invokeId uint8, sequenceNumber *uint8, proposedWindowSize *uint8, serviceRequest *BACnetConfirmedServiceRequest, segmentServiceChoice *uint8, segment []byte, apduLength uint16) *APDUConfirmedRequest {
	_result := &APDUConfirmedRequest{
		SegmentedMessage:          segmentedMessage,
		MoreFollows:               moreFollows,
		SegmentedResponseAccepted: segmentedResponseAccepted,
		MaxSegmentsAccepted:       maxSegmentsAccepted,
		MaxApduLengthAccepted:     maxApduLengthAccepted,
		InvokeId:                  invokeId,
		SequenceNumber:            sequenceNumber,
		ProposedWindowSize:        proposedWindowSize,
		ServiceRequest:            serviceRequest,
		SegmentServiceChoice:      segmentServiceChoice,
		Segment:                   segment,
		APDU:                      NewAPDU(apduLength),
	}
	_result.Child = _result
	return _result
}

func CastAPDUConfirmedRequest(structType interface{}) *APDUConfirmedRequest {
	if casted, ok := structType.(APDUConfirmedRequest); ok {
		return &casted
	}
	if casted, ok := structType.(*APDUConfirmedRequest); ok {
		return casted
	}
	if casted, ok := structType.(APDU); ok {
		return CastAPDUConfirmedRequest(casted.Child)
	}
	if casted, ok := structType.(*APDU); ok {
		return CastAPDUConfirmedRequest(casted.Child)
	}
	return nil
}

func (m *APDUConfirmedRequest) GetTypeName() string {
	return "APDUConfirmedRequest"
}

func (m *APDUConfirmedRequest) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *APDUConfirmedRequest) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (segmentedMessage)
	lengthInBits += 1

	// Simple field (moreFollows)
	lengthInBits += 1

	// Simple field (segmentedResponseAccepted)
	lengthInBits += 1

	// Reserved Field (reserved)
	lengthInBits += 2

	// Simple field (maxSegmentsAccepted)
	lengthInBits += 3

	// Simple field (maxApduLengthAccepted)
	lengthInBits += 4

	// Simple field (invokeId)
	lengthInBits += 8

	// Optional Field (sequenceNumber)
	if m.SequenceNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (proposedWindowSize)
	if m.ProposedWindowSize != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Optional Field (serviceRequest)
	if m.ServiceRequest != nil {
		lengthInBits += (*m.ServiceRequest).GetLengthInBits()
	}

	// Optional Field (segmentServiceChoice)
	if m.SegmentServiceChoice != nil {
		lengthInBits += 8
	}

	// A virtual field doesn't have any in- or output.

	// Array field
	if len(m.Segment) > 0 {
		lengthInBits += 8 * uint16(len(m.Segment))
	}

	return lengthInBits
}

func (m *APDUConfirmedRequest) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func APDUConfirmedRequestParse(readBuffer utils.ReadBuffer, apduLength uint16) (*APDUConfirmedRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("APDUConfirmedRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for APDUConfirmedRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (segmentedMessage)
	_segmentedMessage, _segmentedMessageErr := readBuffer.ReadBit("segmentedMessage")
	if _segmentedMessageErr != nil {
		return nil, errors.Wrap(_segmentedMessageErr, "Error parsing 'segmentedMessage' field")
	}
	segmentedMessage := _segmentedMessage

	// Simple Field (moreFollows)
	_moreFollows, _moreFollowsErr := readBuffer.ReadBit("moreFollows")
	if _moreFollowsErr != nil {
		return nil, errors.Wrap(_moreFollowsErr, "Error parsing 'moreFollows' field")
	}
	moreFollows := _moreFollows

	// Simple Field (segmentedResponseAccepted)
	_segmentedResponseAccepted, _segmentedResponseAcceptedErr := readBuffer.ReadBit("segmentedResponseAccepted")
	if _segmentedResponseAcceptedErr != nil {
		return nil, errors.Wrap(_segmentedResponseAcceptedErr, "Error parsing 'segmentedResponseAccepted' field")
	}
	segmentedResponseAccepted := _segmentedResponseAccepted

	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 2)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field")
		}
		if reserved != uint8(0) {
			log.Info().Fields(map[string]interface{}{
				"expected value": uint8(0),
				"got value":      reserved,
			}).Msg("Got unexpected response.")
		}
	}

	// Simple Field (maxSegmentsAccepted)
	if pullErr := readBuffer.PullContext("maxSegmentsAccepted"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxSegmentsAccepted")
	}
	_maxSegmentsAccepted, _maxSegmentsAcceptedErr := MaxSegmentsAcceptedParse(readBuffer)
	if _maxSegmentsAcceptedErr != nil {
		return nil, errors.Wrap(_maxSegmentsAcceptedErr, "Error parsing 'maxSegmentsAccepted' field")
	}
	maxSegmentsAccepted := _maxSegmentsAccepted
	if closeErr := readBuffer.CloseContext("maxSegmentsAccepted"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxSegmentsAccepted")
	}

	// Simple Field (maxApduLengthAccepted)
	if pullErr := readBuffer.PullContext("maxApduLengthAccepted"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxApduLengthAccepted")
	}
	_maxApduLengthAccepted, _maxApduLengthAcceptedErr := MaxApduLengthAcceptedParse(readBuffer)
	if _maxApduLengthAcceptedErr != nil {
		return nil, errors.Wrap(_maxApduLengthAcceptedErr, "Error parsing 'maxApduLengthAccepted' field")
	}
	maxApduLengthAccepted := _maxApduLengthAccepted
	if closeErr := readBuffer.CloseContext("maxApduLengthAccepted"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxApduLengthAccepted")
	}

	// Simple Field (invokeId)
	_invokeId, _invokeIdErr := readBuffer.ReadUint8("invokeId", 8)
	if _invokeIdErr != nil {
		return nil, errors.Wrap(_invokeIdErr, "Error parsing 'invokeId' field")
	}
	invokeId := _invokeId

	// Optional Field (sequenceNumber) (Can be skipped, if a given expression evaluates to false)
	var sequenceNumber *uint8 = nil
	if segmentedMessage {
		_val, _err := readBuffer.ReadUint8("sequenceNumber", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'sequenceNumber' field")
		}
		sequenceNumber = &_val
	}

	// Optional Field (proposedWindowSize) (Can be skipped, if a given expression evaluates to false)
	var proposedWindowSize *uint8 = nil
	if segmentedMessage {
		_val, _err := readBuffer.ReadUint8("proposedWindowSize", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'proposedWindowSize' field")
		}
		proposedWindowSize = &_val
	}

	// Virtual field
	_apduHeaderReduction := uint16(uint16(3)) + uint16(uint16(utils.InlineIf(segmentedMessage, func() interface{} { return uint16(uint16(2)) }, func() interface{} { return uint16(uint16(0)) }).(uint16)))
	apduHeaderReduction := uint16(_apduHeaderReduction)
	_ = apduHeaderReduction

	// Optional Field (serviceRequest) (Can be skipped, if a given expression evaluates to false)
	var serviceRequest *BACnetConfirmedServiceRequest = nil
	if !(segmentedMessage) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("serviceRequest"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for serviceRequest")
		}
		_val, _err := BACnetConfirmedServiceRequestParse(readBuffer, uint16(apduLength)-uint16(apduHeaderReduction))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'serviceRequest' field")
		default:
			serviceRequest = CastBACnetConfirmedServiceRequest(_val)
			if closeErr := readBuffer.CloseContext("serviceRequest"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for serviceRequest")
			}
		}
	}

	// Validation
	if !(bool(bool(bool(!(segmentedMessage)) && bool(bool((serviceRequest) != (nil))))) || bool(segmentedMessage)) {
		return nil, errors.WithStack(utils.ParseValidationError{"service request should be set"})
	}

	// Optional Field (segmentServiceChoice) (Can be skipped, if a given expression evaluates to false)
	var segmentServiceChoice *uint8 = nil
	if bool(segmentedMessage) && bool(bool((*sequenceNumber) != (0))) {
		_val, _err := readBuffer.ReadUint8("segmentServiceChoice", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'segmentServiceChoice' field")
		}
		segmentServiceChoice = &_val
	}

	// Virtual field
	_segmentReduction := utils.InlineIf(bool(bool((segmentServiceChoice) != (nil))), func() interface{} { return uint16(uint16(uint16(apduHeaderReduction) + uint16(uint16(1)))) }, func() interface{} { return uint16(apduHeaderReduction) }).(uint16)
	segmentReduction := uint16(_segmentReduction)
	_ = segmentReduction
	// Byte Array field (segment)
	numberOfBytessegment := int(utils.InlineIf(segmentedMessage, func() interface{} {
		return uint16(uint16(utils.InlineIf(bool(bool((apduLength) > (0))), func() interface{} { return uint16(uint16(uint16(apduLength) - uint16(segmentReduction))) }, func() interface{} { return uint16(uint16(0)) }).(uint16)))
	}, func() interface{} { return uint16(uint16(0)) }).(uint16))
	segment, _readArrayErr := readBuffer.ReadByteArray("segment", numberOfBytessegment)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'segment' field")
	}

	if closeErr := readBuffer.CloseContext("APDUConfirmedRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for APDUConfirmedRequest")
	}

	// Create a partially initialized instance
	_child := &APDUConfirmedRequest{
		SegmentedMessage:          segmentedMessage,
		MoreFollows:               moreFollows,
		SegmentedResponseAccepted: segmentedResponseAccepted,
		MaxSegmentsAccepted:       maxSegmentsAccepted,
		MaxApduLengthAccepted:     maxApduLengthAccepted,
		InvokeId:                  invokeId,
		SequenceNumber:            sequenceNumber,
		ProposedWindowSize:        proposedWindowSize,
		ServiceRequest:            CastBACnetConfirmedServiceRequest(serviceRequest),
		SegmentServiceChoice:      segmentServiceChoice,
		Segment:                   segment,
		APDU:                      &APDU{},
	}
	_child.APDU.Child = _child
	return _child, nil
}

func (m *APDUConfirmedRequest) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("APDUConfirmedRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for APDUConfirmedRequest")
		}

		// Simple Field (segmentedMessage)
		segmentedMessage := bool(m.SegmentedMessage)
		_segmentedMessageErr := writeBuffer.WriteBit("segmentedMessage", (segmentedMessage))
		if _segmentedMessageErr != nil {
			return errors.Wrap(_segmentedMessageErr, "Error serializing 'segmentedMessage' field")
		}

		// Simple Field (moreFollows)
		moreFollows := bool(m.MoreFollows)
		_moreFollowsErr := writeBuffer.WriteBit("moreFollows", (moreFollows))
		if _moreFollowsErr != nil {
			return errors.Wrap(_moreFollowsErr, "Error serializing 'moreFollows' field")
		}

		// Simple Field (segmentedResponseAccepted)
		segmentedResponseAccepted := bool(m.SegmentedResponseAccepted)
		_segmentedResponseAcceptedErr := writeBuffer.WriteBit("segmentedResponseAccepted", (segmentedResponseAccepted))
		if _segmentedResponseAcceptedErr != nil {
			return errors.Wrap(_segmentedResponseAcceptedErr, "Error serializing 'segmentedResponseAccepted' field")
		}

		// Reserved Field (reserved)
		{
			_err := writeBuffer.WriteUint8("reserved", 2, uint8(0))
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (maxSegmentsAccepted)
		if pushErr := writeBuffer.PushContext("maxSegmentsAccepted"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maxSegmentsAccepted")
		}
		_maxSegmentsAcceptedErr := m.MaxSegmentsAccepted.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("maxSegmentsAccepted"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maxSegmentsAccepted")
		}
		if _maxSegmentsAcceptedErr != nil {
			return errors.Wrap(_maxSegmentsAcceptedErr, "Error serializing 'maxSegmentsAccepted' field")
		}

		// Simple Field (maxApduLengthAccepted)
		if pushErr := writeBuffer.PushContext("maxApduLengthAccepted"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for maxApduLengthAccepted")
		}
		_maxApduLengthAcceptedErr := m.MaxApduLengthAccepted.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("maxApduLengthAccepted"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for maxApduLengthAccepted")
		}
		if _maxApduLengthAcceptedErr != nil {
			return errors.Wrap(_maxApduLengthAcceptedErr, "Error serializing 'maxApduLengthAccepted' field")
		}

		// Simple Field (invokeId)
		invokeId := uint8(m.InvokeId)
		_invokeIdErr := writeBuffer.WriteUint8("invokeId", 8, (invokeId))
		if _invokeIdErr != nil {
			return errors.Wrap(_invokeIdErr, "Error serializing 'invokeId' field")
		}

		// Optional Field (sequenceNumber) (Can be skipped, if the value is null)
		var sequenceNumber *uint8 = nil
		if m.SequenceNumber != nil {
			sequenceNumber = m.SequenceNumber
			_sequenceNumberErr := writeBuffer.WriteUint8("sequenceNumber", 8, *(sequenceNumber))
			if _sequenceNumberErr != nil {
				return errors.Wrap(_sequenceNumberErr, "Error serializing 'sequenceNumber' field")
			}
		}

		// Optional Field (proposedWindowSize) (Can be skipped, if the value is null)
		var proposedWindowSize *uint8 = nil
		if m.ProposedWindowSize != nil {
			proposedWindowSize = m.ProposedWindowSize
			_proposedWindowSizeErr := writeBuffer.WriteUint8("proposedWindowSize", 8, *(proposedWindowSize))
			if _proposedWindowSizeErr != nil {
				return errors.Wrap(_proposedWindowSizeErr, "Error serializing 'proposedWindowSize' field")
			}
		}
		// Virtual field
		if _apduHeaderReductionErr := writeBuffer.WriteVirtual("apduHeaderReduction", m.GetApduHeaderReduction()); _apduHeaderReductionErr != nil {
			return errors.Wrap(_apduHeaderReductionErr, "Error serializing 'apduHeaderReduction' field")
		}

		// Optional Field (serviceRequest) (Can be skipped, if the value is null)
		var serviceRequest *BACnetConfirmedServiceRequest = nil
		if m.ServiceRequest != nil {
			if pushErr := writeBuffer.PushContext("serviceRequest"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for serviceRequest")
			}
			serviceRequest = m.ServiceRequest
			_serviceRequestErr := serviceRequest.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("serviceRequest"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for serviceRequest")
			}
			if _serviceRequestErr != nil {
				return errors.Wrap(_serviceRequestErr, "Error serializing 'serviceRequest' field")
			}
		}

		// Optional Field (segmentServiceChoice) (Can be skipped, if the value is null)
		var segmentServiceChoice *uint8 = nil
		if m.SegmentServiceChoice != nil {
			segmentServiceChoice = m.SegmentServiceChoice
			_segmentServiceChoiceErr := writeBuffer.WriteUint8("segmentServiceChoice", 8, *(segmentServiceChoice))
			if _segmentServiceChoiceErr != nil {
				return errors.Wrap(_segmentServiceChoiceErr, "Error serializing 'segmentServiceChoice' field")
			}
		}
		// Virtual field
		if _segmentReductionErr := writeBuffer.WriteVirtual("segmentReduction", m.GetSegmentReduction()); _segmentReductionErr != nil {
			return errors.Wrap(_segmentReductionErr, "Error serializing 'segmentReduction' field")
		}

		// Array Field (segment)
		if m.Segment != nil {
			// Byte Array field (segment)
			_writeArrayErr := writeBuffer.WriteByteArray("segment", m.Segment)
			if _writeArrayErr != nil {
				return errors.Wrap(_writeArrayErr, "Error serializing 'segment' field")
			}
		}

		if popErr := writeBuffer.PopContext("APDUConfirmedRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for APDUConfirmedRequest")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *APDUConfirmedRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
