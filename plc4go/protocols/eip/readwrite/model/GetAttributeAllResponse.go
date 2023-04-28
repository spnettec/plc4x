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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// GetAttributeAllResponse is the corresponding interface of GetAttributeAllResponse
type GetAttributeAllResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CipService
	// GetStatus returns Status (property field)
	GetStatus() uint8
	// GetExtStatus returns ExtStatus (property field)
	GetExtStatus() uint8
	// GetAttributes returns Attributes (property field)
	GetAttributes() CIPAttributes
}

// GetAttributeAllResponseExactly can be used when we want exactly this type and not a type which fulfills GetAttributeAllResponse.
// This is useful for switch cases.
type GetAttributeAllResponseExactly interface {
	GetAttributeAllResponse
	isGetAttributeAllResponse() bool
}

// _GetAttributeAllResponse is the data-structure of this message
type _GetAttributeAllResponse struct {
	*_CipService
	Status     uint8
	ExtStatus  uint8
	Attributes CIPAttributes
	// Reserved Fields
	reservedField0 *uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_GetAttributeAllResponse) GetService() uint8 {
	return 0x01
}

func (m *_GetAttributeAllResponse) GetResponse() bool {
	return bool(true)
}

func (m *_GetAttributeAllResponse) GetConnected() bool {
	return false
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_GetAttributeAllResponse) InitializeParent(parent CipService) {}

func (m *_GetAttributeAllResponse) GetParent() CipService {
	return m._CipService
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_GetAttributeAllResponse) GetStatus() uint8 {
	return m.Status
}

func (m *_GetAttributeAllResponse) GetExtStatus() uint8 {
	return m.ExtStatus
}

func (m *_GetAttributeAllResponse) GetAttributes() CIPAttributes {
	return m.Attributes
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewGetAttributeAllResponse factory function for _GetAttributeAllResponse
func NewGetAttributeAllResponse(status uint8, extStatus uint8, attributes CIPAttributes, serviceLen uint16) *_GetAttributeAllResponse {
	_result := &_GetAttributeAllResponse{
		Status:      status,
		ExtStatus:   extStatus,
		Attributes:  attributes,
		_CipService: NewCipService(serviceLen),
	}
	_result._CipService._CipServiceChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastGetAttributeAllResponse(structType any) GetAttributeAllResponse {
	if casted, ok := structType.(GetAttributeAllResponse); ok {
		return casted
	}
	if casted, ok := structType.(*GetAttributeAllResponse); ok {
		return *casted
	}
	return nil
}

func (m *_GetAttributeAllResponse) GetTypeName() string {
	return "GetAttributeAllResponse"
}

func (m *_GetAttributeAllResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (status)
	lengthInBits += 8

	// Simple field (extStatus)
	lengthInBits += 8

	// Optional Field (attributes)
	if m.Attributes != nil {
		lengthInBits += m.Attributes.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_GetAttributeAllResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func GetAttributeAllResponseParse(theBytes []byte, connected bool, serviceLen uint16) (GetAttributeAllResponse, error) {
	return GetAttributeAllResponseParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), connected, serviceLen)
}

func GetAttributeAllResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, connected bool, serviceLen uint16) (GetAttributeAllResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("GetAttributeAllResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for GetAttributeAllResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of GetAttributeAllResponse")
		}
		if reserved != uint8(0x00) {
			Plc4xModelLog.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value":      reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (status)
	_status, _statusErr := readBuffer.ReadUint8("status", 8)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field of GetAttributeAllResponse")
	}
	status := _status

	// Simple Field (extStatus)
	_extStatus, _extStatusErr := readBuffer.ReadUint8("extStatus", 8)
	if _extStatusErr != nil {
		return nil, errors.Wrap(_extStatusErr, "Error parsing 'extStatus' field of GetAttributeAllResponse")
	}
	extStatus := _extStatus

	// Optional Field (attributes) (Can be skipped, if a given expression evaluates to false)
	var attributes CIPAttributes = nil
	if bool(((serviceLen) - (4)) > (0)) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("attributes"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for attributes")
		}
		_val, _err := CIPAttributesParseWithBuffer(ctx, readBuffer, uint16(serviceLen)-uint16(uint16(4)))
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'attributes' field of GetAttributeAllResponse")
		default:
			attributes = _val.(CIPAttributes)
			if closeErr := readBuffer.CloseContext("attributes"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for attributes")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("GetAttributeAllResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for GetAttributeAllResponse")
	}

	// Create a partially initialized instance
	_child := &_GetAttributeAllResponse{
		_CipService: &_CipService{
			ServiceLen: serviceLen,
		},
		Status:         status,
		ExtStatus:      extStatus,
		Attributes:     attributes,
		reservedField0: reservedField0,
	}
	_child._CipService._CipServiceChildRequirements = _child
	return _child, nil
}

func (m *_GetAttributeAllResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_GetAttributeAllResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("GetAttributeAllResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for GetAttributeAllResponse")
		}

		// Reserved Field (reserved)
		{
			var reserved uint8 = uint8(0x00)
			if m.reservedField0 != nil {
				Plc4xModelLog.Info().Fields(map[string]any{
					"expected value": uint8(0x00),
					"got value":      reserved,
				}).Msg("Overriding reserved field with unexpected value.")
				reserved = *m.reservedField0
			}
			_err := writeBuffer.WriteUint8("reserved", 8, reserved)
			if _err != nil {
				return errors.Wrap(_err, "Error serializing 'reserved' field")
			}
		}

		// Simple Field (status)
		status := uint8(m.GetStatus())
		_statusErr := writeBuffer.WriteUint8("status", 8, (status))
		if _statusErr != nil {
			return errors.Wrap(_statusErr, "Error serializing 'status' field")
		}

		// Simple Field (extStatus)
		extStatus := uint8(m.GetExtStatus())
		_extStatusErr := writeBuffer.WriteUint8("extStatus", 8, (extStatus))
		if _extStatusErr != nil {
			return errors.Wrap(_extStatusErr, "Error serializing 'extStatus' field")
		}

		// Optional Field (attributes) (Can be skipped, if the value is null)
		var attributes CIPAttributes = nil
		if m.GetAttributes() != nil {
			if pushErr := writeBuffer.PushContext("attributes"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for attributes")
			}
			attributes = m.GetAttributes()
			_attributesErr := writeBuffer.WriteSerializable(ctx, attributes)
			if popErr := writeBuffer.PopContext("attributes"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for attributes")
			}
			if _attributesErr != nil {
				return errors.Wrap(_attributesErr, "Error serializing 'attributes' field")
			}
		}

		if popErr := writeBuffer.PopContext("GetAttributeAllResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for GetAttributeAllResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_GetAttributeAllResponse) isGetAttributeAllResponse() bool {
	return true
}

func (m *_GetAttributeAllResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
