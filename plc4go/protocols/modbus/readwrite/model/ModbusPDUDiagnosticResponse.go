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


// ModbusPDUDiagnosticResponse is the corresponding interface of ModbusPDUDiagnosticResponse
type ModbusPDUDiagnosticResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ModbusPDU
	// GetSubFunction returns SubFunction (property field)
	GetSubFunction() uint16
	// GetData returns Data (property field)
	GetData() uint16
}

// ModbusPDUDiagnosticResponseExactly can be used when we want exactly this type and not a type which fulfills ModbusPDUDiagnosticResponse.
// This is useful for switch cases.
type ModbusPDUDiagnosticResponseExactly interface {
	ModbusPDUDiagnosticResponse
	isModbusPDUDiagnosticResponse() bool
}

// _ModbusPDUDiagnosticResponse is the data-structure of this message
type _ModbusPDUDiagnosticResponse struct {
	*_ModbusPDU
        SubFunction uint16
        Data uint16
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModbusPDUDiagnosticResponse)  GetErrorFlag() bool {
return bool(false)}

func (m *_ModbusPDUDiagnosticResponse)  GetFunctionFlag() uint8 {
return 0x08}

func (m *_ModbusPDUDiagnosticResponse)  GetResponse() bool {
return bool(true)}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModbusPDUDiagnosticResponse) InitializeParent(parent ModbusPDU ) {}

func (m *_ModbusPDUDiagnosticResponse)  GetParent() ModbusPDU {
	return m._ModbusPDU
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusPDUDiagnosticResponse) GetSubFunction() uint16 {
	return m.SubFunction
}

func (m *_ModbusPDUDiagnosticResponse) GetData() uint16 {
	return m.Data
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewModbusPDUDiagnosticResponse factory function for _ModbusPDUDiagnosticResponse
func NewModbusPDUDiagnosticResponse( subFunction uint16 , data uint16 ) *_ModbusPDUDiagnosticResponse {
	_result := &_ModbusPDUDiagnosticResponse{
		SubFunction: subFunction,
		Data: data,
    	_ModbusPDU: NewModbusPDU(),
	}
	_result._ModbusPDU._ModbusPDUChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModbusPDUDiagnosticResponse(structType any) ModbusPDUDiagnosticResponse {
    if casted, ok := structType.(ModbusPDUDiagnosticResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusPDUDiagnosticResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusPDUDiagnosticResponse) GetTypeName() string {
	return "ModbusPDUDiagnosticResponse"
}

func (m *_ModbusPDUDiagnosticResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (subFunction)
	lengthInBits += 16;

	// Simple field (data)
	lengthInBits += 16;

	return lengthInBits
}


func (m *_ModbusPDUDiagnosticResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusPDUDiagnosticResponseParse(ctx context.Context, theBytes []byte, response bool) (ModbusPDUDiagnosticResponse, error) {
	return ModbusPDUDiagnosticResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), response)
}

func ModbusPDUDiagnosticResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, response bool) (ModbusPDUDiagnosticResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ModbusPDUDiagnosticResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusPDUDiagnosticResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (subFunction)
_subFunction, _subFunctionErr := readBuffer.ReadUint16("subFunction", 16)
	if _subFunctionErr != nil {
		return nil, errors.Wrap(_subFunctionErr, "Error parsing 'subFunction' field of ModbusPDUDiagnosticResponse")
	}
	subFunction := _subFunction

	// Simple Field (data)
_data, _dataErr := readBuffer.ReadUint16("data", 16)
	if _dataErr != nil {
		return nil, errors.Wrap(_dataErr, "Error parsing 'data' field of ModbusPDUDiagnosticResponse")
	}
	data := _data

	if closeErr := readBuffer.CloseContext("ModbusPDUDiagnosticResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusPDUDiagnosticResponse")
	}

	// Create a partially initialized instance
	_child := &_ModbusPDUDiagnosticResponse{
		_ModbusPDU: &_ModbusPDU{
		},
		SubFunction: subFunction,
		Data: data,
	}
	_child._ModbusPDU._ModbusPDUChildRequirements = _child
	return _child, nil
}

func (m *_ModbusPDUDiagnosticResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusPDUDiagnosticResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModbusPDUDiagnosticResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModbusPDUDiagnosticResponse")
		}

	// Simple Field (subFunction)
	subFunction := uint16(m.GetSubFunction())
	_subFunctionErr := writeBuffer.WriteUint16("subFunction", 16, (subFunction))
	if _subFunctionErr != nil {
		return errors.Wrap(_subFunctionErr, "Error serializing 'subFunction' field")
	}

	// Simple Field (data)
	data := uint16(m.GetData())
	_dataErr := writeBuffer.WriteUint16("data", 16, (data))
	if _dataErr != nil {
		return errors.Wrap(_dataErr, "Error serializing 'data' field")
	}

		if popErr := writeBuffer.PopContext("ModbusPDUDiagnosticResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModbusPDUDiagnosticResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_ModbusPDUDiagnosticResponse) isModbusPDUDiagnosticResponse() bool {
	return true
}

func (m *_ModbusPDUDiagnosticResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



