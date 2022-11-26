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

// TDataConnectedReq is the corresponding interface of TDataConnectedReq
type TDataConnectedReq interface {
	utils.LengthAware
	utils.Serializable
	CEMI
}

// TDataConnectedReqExactly can be used when we want exactly this type and not a type which fulfills TDataConnectedReq.
// This is useful for switch cases.
type TDataConnectedReqExactly interface {
	TDataConnectedReq
	isTDataConnectedReq() bool
}

// _TDataConnectedReq is the data-structure of this message
type _TDataConnectedReq struct {
	*_CEMI
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TDataConnectedReq) GetMessageCode() uint8 {
	return 0x41
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TDataConnectedReq) InitializeParent(parent CEMI) {}

func (m *_TDataConnectedReq) GetParent() CEMI {
	return m._CEMI
}

// NewTDataConnectedReq factory function for _TDataConnectedReq
func NewTDataConnectedReq(size uint16) *_TDataConnectedReq {
	_result := &_TDataConnectedReq{
		_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTDataConnectedReq(structType interface{}) TDataConnectedReq {
	if casted, ok := structType.(TDataConnectedReq); ok {
		return casted
	}
	if casted, ok := structType.(*TDataConnectedReq); ok {
		return *casted
	}
	return nil
}

func (m *_TDataConnectedReq) GetTypeName() string {
	return "TDataConnectedReq"
}

func (m *_TDataConnectedReq) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_TDataConnectedReq) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_TDataConnectedReq) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TDataConnectedReqParse(theBytes []byte, size uint16) (TDataConnectedReq, error) {
	return TDataConnectedReqParseWithBuffer(utils.NewReadBufferByteBased(theBytes), size)
}

func TDataConnectedReqParseWithBuffer(readBuffer utils.ReadBuffer, size uint16) (TDataConnectedReq, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TDataConnectedReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TDataConnectedReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TDataConnectedReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TDataConnectedReq")
	}

	// Create a partially initialized instance
	_child := &_TDataConnectedReq{
		_CEMI: &_CEMI{
			Size: size,
		},
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_TDataConnectedReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TDataConnectedReq) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TDataConnectedReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TDataConnectedReq")
		}

		if popErr := writeBuffer.PopContext("TDataConnectedReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TDataConnectedReq")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_TDataConnectedReq) isTDataConnectedReq() bool {
	return true
}

func (m *_TDataConnectedReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
