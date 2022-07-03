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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestLifeSafetyOperation is the corresponding interface of BACnetConfirmedServiceRequestLifeSafetyOperation
type BACnetConfirmedServiceRequestLifeSafetyOperation interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetRequestingProcessIdentifier returns RequestingProcessIdentifier (property field)
	GetRequestingProcessIdentifier() BACnetContextTagUnsignedInteger
	// GetRequestingSource returns RequestingSource (property field)
	GetRequestingSource() BACnetContextTagCharacterString
	// GetRequest returns Request (property field)
	GetRequest() BACnetLifeSafetyOperationTagged
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
}

// BACnetConfirmedServiceRequestLifeSafetyOperationExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestLifeSafetyOperation.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestLifeSafetyOperationExactly interface {
	BACnetConfirmedServiceRequestLifeSafetyOperation
	isBACnetConfirmedServiceRequestLifeSafetyOperation() bool
}

// _BACnetConfirmedServiceRequestLifeSafetyOperation is the data-structure of this message
type _BACnetConfirmedServiceRequestLifeSafetyOperation struct {
	*_BACnetConfirmedServiceRequest
	RequestingProcessIdentifier BACnetContextTagUnsignedInteger
	RequestingSource            BACnetContextTagCharacterString
	Request                     BACnetLifeSafetyOperationTagged
	ObjectIdentifier            BACnetContextTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_LIFE_SAFETY_OPERATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetRequestingProcessIdentifier() BACnetContextTagUnsignedInteger {
	return m.RequestingProcessIdentifier
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetRequestingSource() BACnetContextTagCharacterString {
	return m.RequestingSource
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetRequest() BACnetLifeSafetyOperationTagged {
	return m.Request
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestLifeSafetyOperation factory function for _BACnetConfirmedServiceRequestLifeSafetyOperation
func NewBACnetConfirmedServiceRequestLifeSafetyOperation(requestingProcessIdentifier BACnetContextTagUnsignedInteger, requestingSource BACnetContextTagCharacterString, request BACnetLifeSafetyOperationTagged, objectIdentifier BACnetContextTagObjectIdentifier, serviceRequestLength uint16) *_BACnetConfirmedServiceRequestLifeSafetyOperation {
	_result := &_BACnetConfirmedServiceRequestLifeSafetyOperation{
		RequestingProcessIdentifier:    requestingProcessIdentifier,
		RequestingSource:               requestingSource,
		Request:                        request,
		ObjectIdentifier:               objectIdentifier,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestLifeSafetyOperation(structType interface{}) BACnetConfirmedServiceRequestLifeSafetyOperation {
	if casted, ok := structType.(BACnetConfirmedServiceRequestLifeSafetyOperation); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestLifeSafetyOperation); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetTypeName() string {
	return "BACnetConfirmedServiceRequestLifeSafetyOperation"
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (requestingProcessIdentifier)
	lengthInBits += m.RequestingProcessIdentifier.GetLengthInBits()

	// Simple field (requestingSource)
	lengthInBits += m.RequestingSource.GetLengthInBits()

	// Simple field (request)
	lengthInBits += m.Request.GetLengthInBits()

	// Optional Field (objectIdentifier)
	if m.ObjectIdentifier != nil {
		lengthInBits += m.ObjectIdentifier.GetLengthInBits()
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestLifeSafetyOperationParse(readBuffer utils.ReadBuffer, serviceRequestLength uint16) (BACnetConfirmedServiceRequestLifeSafetyOperation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestLifeSafetyOperation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestLifeSafetyOperation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestingProcessIdentifier)
	if pullErr := readBuffer.PullContext("requestingProcessIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestingProcessIdentifier")
	}
	_requestingProcessIdentifier, _requestingProcessIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _requestingProcessIdentifierErr != nil {
		return nil, errors.Wrap(_requestingProcessIdentifierErr, "Error parsing 'requestingProcessIdentifier' field of BACnetConfirmedServiceRequestLifeSafetyOperation")
	}
	requestingProcessIdentifier := _requestingProcessIdentifier.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("requestingProcessIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestingProcessIdentifier")
	}

	// Simple Field (requestingSource)
	if pullErr := readBuffer.PullContext("requestingSource"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestingSource")
	}
	_requestingSource, _requestingSourceErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_CHARACTER_STRING))
	if _requestingSourceErr != nil {
		return nil, errors.Wrap(_requestingSourceErr, "Error parsing 'requestingSource' field of BACnetConfirmedServiceRequestLifeSafetyOperation")
	}
	requestingSource := _requestingSource.(BACnetContextTagCharacterString)
	if closeErr := readBuffer.CloseContext("requestingSource"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestingSource")
	}

	// Simple Field (request)
	if pullErr := readBuffer.PullContext("request"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for request")
	}
	_request, _requestErr := BACnetLifeSafetyOperationTaggedParse(readBuffer, uint8(uint8(2)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _requestErr != nil {
		return nil, errors.Wrap(_requestErr, "Error parsing 'request' field of BACnetConfirmedServiceRequestLifeSafetyOperation")
	}
	request := _request.(BACnetLifeSafetyOperationTagged)
	if closeErr := readBuffer.CloseContext("request"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for request")
	}

	// Optional Field (objectIdentifier) (Can be skipped, if a given expression evaluates to false)
	var objectIdentifier BACnetContextTagObjectIdentifier = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(3), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'objectIdentifier' field of BACnetConfirmedServiceRequestLifeSafetyOperation")
		default:
			objectIdentifier = _val.(BACnetContextTagObjectIdentifier)
			if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestLifeSafetyOperation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestLifeSafetyOperation")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestLifeSafetyOperation{
		RequestingProcessIdentifier: requestingProcessIdentifier,
		RequestingSource:            requestingSource,
		Request:                     request,
		ObjectIdentifier:            objectIdentifier,
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestLifeSafetyOperation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestLifeSafetyOperation")
		}

		// Simple Field (requestingProcessIdentifier)
		if pushErr := writeBuffer.PushContext("requestingProcessIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestingProcessIdentifier")
		}
		_requestingProcessIdentifierErr := writeBuffer.WriteSerializable(m.GetRequestingProcessIdentifier())
		if popErr := writeBuffer.PopContext("requestingProcessIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestingProcessIdentifier")
		}
		if _requestingProcessIdentifierErr != nil {
			return errors.Wrap(_requestingProcessIdentifierErr, "Error serializing 'requestingProcessIdentifier' field")
		}

		// Simple Field (requestingSource)
		if pushErr := writeBuffer.PushContext("requestingSource"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestingSource")
		}
		_requestingSourceErr := writeBuffer.WriteSerializable(m.GetRequestingSource())
		if popErr := writeBuffer.PopContext("requestingSource"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestingSource")
		}
		if _requestingSourceErr != nil {
			return errors.Wrap(_requestingSourceErr, "Error serializing 'requestingSource' field")
		}

		// Simple Field (request)
		if pushErr := writeBuffer.PushContext("request"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for request")
		}
		_requestErr := writeBuffer.WriteSerializable(m.GetRequest())
		if popErr := writeBuffer.PopContext("request"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for request")
		}
		if _requestErr != nil {
			return errors.Wrap(_requestErr, "Error serializing 'request' field")
		}

		// Optional Field (objectIdentifier) (Can be skipped, if the value is null)
		var objectIdentifier BACnetContextTagObjectIdentifier = nil
		if m.GetObjectIdentifier() != nil {
			if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
			}
			objectIdentifier = m.GetObjectIdentifier()
			_objectIdentifierErr := writeBuffer.WriteSerializable(objectIdentifier)
			if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for objectIdentifier")
			}
			if _objectIdentifierErr != nil {
				return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestLifeSafetyOperation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestLifeSafetyOperation")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) isBACnetConfirmedServiceRequestLifeSafetyOperation() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestLifeSafetyOperation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
