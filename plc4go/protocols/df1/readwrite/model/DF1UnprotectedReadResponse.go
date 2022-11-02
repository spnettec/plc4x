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


// DF1UnprotectedReadResponse is the corresponding interface of DF1UnprotectedReadResponse
type DF1UnprotectedReadResponse interface {
	utils.LengthAware
	utils.Serializable
	DF1Command
	// GetData returns Data (property field)
	GetData() []byte
}

// DF1UnprotectedReadResponseExactly can be used when we want exactly this type and not a type which fulfills DF1UnprotectedReadResponse.
// This is useful for switch cases.
type DF1UnprotectedReadResponseExactly interface {
	DF1UnprotectedReadResponse
	isDF1UnprotectedReadResponse() bool
}

// _DF1UnprotectedReadResponse is the data-structure of this message
type _DF1UnprotectedReadResponse struct {
	*_DF1Command
        Data []byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DF1UnprotectedReadResponse)  GetCommandCode() uint8 {
return 0x41}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DF1UnprotectedReadResponse) InitializeParent(parent DF1Command , status uint8 , transactionCounter uint16 ) {	m.Status = status
	m.TransactionCounter = transactionCounter
}

func (m *_DF1UnprotectedReadResponse)  GetParent() DF1Command {
	return m._DF1Command
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DF1UnprotectedReadResponse) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewDF1UnprotectedReadResponse factory function for _DF1UnprotectedReadResponse
func NewDF1UnprotectedReadResponse( data []byte , status uint8 , transactionCounter uint16 ) *_DF1UnprotectedReadResponse {
	_result := &_DF1UnprotectedReadResponse{
		Data: data,
    	_DF1Command: NewDF1Command(status, transactionCounter),
	}
	_result._DF1Command._DF1CommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastDF1UnprotectedReadResponse(structType interface{}) DF1UnprotectedReadResponse {
    if casted, ok := structType.(DF1UnprotectedReadResponse); ok {
		return casted
	}
	if casted, ok := structType.(*DF1UnprotectedReadResponse); ok {
		return *casted
	}
	return nil
}

func (m *_DF1UnprotectedReadResponse) GetTypeName() string {
	return "DF1UnprotectedReadResponse"
}

func (m *_DF1UnprotectedReadResponse) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_DF1UnprotectedReadResponse) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Manual Array Field (data)
	lengthInBits += uint16(DataLength(m.GetData()))

	return lengthInBits
}


func (m *_DF1UnprotectedReadResponse) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DF1UnprotectedReadResponseParse(readBuffer utils.ReadBuffer) (DF1UnprotectedReadResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DF1UnprotectedReadResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DF1UnprotectedReadResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	if pullErr := readBuffer.PullContext("data", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for data")
	}
	// Manual Array Field (data)
	// Terminated array
	var _dataList []byte
	{
		_values := &_dataList
		_ = _values
		for ;!((bool) (DataTerminate(readBuffer))); {
		_dataList = append(_dataList, ((byte) (ReadData(readBuffer))))

		}
	}
	data := _dataList
	if closeErr := readBuffer.CloseContext("data", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for data")
	}

	if closeErr := readBuffer.CloseContext("DF1UnprotectedReadResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DF1UnprotectedReadResponse")
	}

	// Create a partially initialized instance
	_child := &_DF1UnprotectedReadResponse{
		_DF1Command: &_DF1Command{
		},
		Data: data,
	}
	_child._DF1Command._DF1CommandChildRequirements = _child
	return _child, nil
}

func (m *_DF1UnprotectedReadResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DF1UnprotectedReadResponse) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DF1UnprotectedReadResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DF1UnprotectedReadResponse")
		}

	// Manual Array Field (data)
	if pushErr := writeBuffer.PushContext("data", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for data")
	}
	for _, m := range m.GetData() {
			WriteData(writeBuffer, m)
	}
	if popErr := writeBuffer.PopContext("data", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for data")
	}

		if popErr := writeBuffer.PopContext("DF1UnprotectedReadResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DF1UnprotectedReadResponse")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_DF1UnprotectedReadResponse) isDF1UnprotectedReadResponse() bool {
	return true
}

func (m *_DF1UnprotectedReadResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



