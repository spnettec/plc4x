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
	"github.com/rs/zerolog"
)

	// Code generated by code-generation. DO NOT EDIT.


// ModbusPDUReportServerIdResponse is the corresponding interface of ModbusPDUReportServerIdResponse
type ModbusPDUReportServerIdResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetValue returns Value (property field)
	GetValue() []byte
}

// ModbusPDUReportServerIdResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUReportServerIdResponse.
// This is useful for switch cases.
type ModbusPDUReportServerIdResponseExactly interface {
	ModbusPDUReportServerIdResponse
	isModbusPDUReportServerIdResponse() bool
}

// _ModbusPDUReportServerIdResponse is the data-structure of this message
type _ModbusPDUReportServerIdResponse struct {
	*_ModbusPDU
        Value []byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUReportServerIdResponse)  GetErrorFlag() bool {
return bool(false)}

func (m *_ModbusPDUReportServerIdResponse)  GetFunctionFlag() uint8 {
return 0x11}

func (m *_ModbusPDUReportServerIdResponse)  GetResponse() bool {
return bool(true)}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUReportServerIdResponse) InitializeParent(parent ModbusPDU ) {}

func (m *_ModbusPDUReportServerIdResponse)  GetParent() ModbusPDU {
	return m._ModbusPDU
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUReportServerIdResponse) GetValue() []byte {
	return m.Value
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewModbusPDUReportServerIdResponse factory function for _ModbusPDUReportServerIdResponse
func NewModbusPDUReportServerIdResponse( value []byte ) *_ModbusPDUReportServerIdResponse {
	_result := &_ModbusPDUReportServerIdResponse{
		Value: value,
    	_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUReportServerIdResponse(structType any) ModbusPDUReportServerIdResponse {
    if casted, ok := structType.(ModbusPDUReportServerIdResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUReportServerIdResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUReportServerIdResponse) GetTypeName() string {
	return "ModbusPDUReportServerIdResponse"
}

func (m *_ModbusPDUReportServerIdResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (byteCount)
	lengthInBits += 8

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}


func (m *_ModbusPDUReportServerIdResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUReportServerIdResponseParse(ctx context.Context, theBytes []byte, response bool) (ModbusPDUReportServerIdResponse, error) {
	return ModbusPDUReportServerIdResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUReportServerIdResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUReportServerIdResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ModbusPDUReportServerIdResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUReportServerIdResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (byteCount) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	byteCount, _byteCountErr := readBuffer.ReadUint8("byteCount", 8)
	_ = byteCount
	if _byteCountErr != nil {
		return nil, errors.Wrap(_byteCountErr, "Error parsing 'byteCount' field of ModbusPDUReportServerIdResponse")
	}
	// Byte Array field (value)
	numberOfBytesvalue := int(byteCount)
	value, _readArrayErr := readBuffer.ReadByteArray("value", numberOfBytesvalue)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'value' field of ModbusPDUReportServerIdResponse")
	}

	if closeErr := readBuffer.CloseContext("ModbusPDUReportServerIdResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUReportServerIdResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUReportServerIdResponse{
		_ModbusPDU: &_ModbusPDU{
		},
		Value: value,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUReportServerIdResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUReportServerIdResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUReportServerIdResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUReportServerIdResponse")
		}

	// Implicit Field (byteCount) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	byteCount := uint8(uint8(len(m.GetValue())))
	_byteCountErr := writeBuffer.WriteUint8("byteCount", 8, (byteCount))
	if _byteCountErr != nil {
		return errors.Wrap(_byteCountErr, "Error serializing 'byteCount' field")
	}

	// Array Field (value)
	// Byte Array field (value)
	if err := writeBuffer.WriteByteArray("value", m.GetValue()); err != nil {
		return errors.Wrap(err, "Error serializing 'value' field")
	}

		if popErr := writeBuffer.PopContext("ModbusPDUReportServerIdResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUReportServerIdResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_ModbusPDUReportServerIdResponse) isModbusPDUReportServerIdResponse() bool {
	return true
}

func (m *_ModbusPDUReportServerIdResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



