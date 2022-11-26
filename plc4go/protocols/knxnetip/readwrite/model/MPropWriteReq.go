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

// MPropWriteReq is the corresponding interface of MPropWriteReq
type MPropWriteReq interface {
	utils.LengthAware
	utils.Serializable
	CEMI
}

// MPropWriteReqExactly can be used when we want exactly this type and not a type which fulfills MPropWriteReq.
// This is useful for switch cases.
type MPropWriteReqExactly interface {
	MPropWriteReq
	isMPropWriteReq() bool
}

// _MPropWriteReq is the data-structure of this message
type _MPropWriteReq struct {
	*_CEMI
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MPropWriteReq) GetMessageCode() uint8 {
	return 0xF6
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MPropWriteReq) InitializeParent(parent CEMI) {}

func (m *_MPropWriteReq) GetParent() CEMI {
	return m._CEMI
}

// NewMPropWriteReq factory function for _MPropWriteReq
func NewMPropWriteReq(size uint16) *_MPropWriteReq {
	_result := &_MPropWriteReq{
		_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMPropWriteReq(structType interface{}) MPropWriteReq {
	if casted, ok := structType.(MPropWriteReq); ok {
		return casted
	}
	if casted, ok := structType.(*MPropWriteReq); ok {
		return *casted
	}
	return nil
}

func (m *_MPropWriteReq) GetTypeName() string {
	return "MPropWriteReq"
}

func (m *_MPropWriteReq) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_MPropWriteReq) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_MPropWriteReq) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MPropWriteReqParse(theBytes []byte, size uint16) (MPropWriteReq, error) {
	return MPropWriteReqParseWithBuffer(utils.NewReadBufferByteBased(theBytes), size)
}

func MPropWriteReqParseWithBuffer(readBuffer utils.ReadBuffer, size uint16) (MPropWriteReq, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MPropWriteReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MPropWriteReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MPropWriteReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MPropWriteReq")
	}

	// Create a partially initialized instance
	_child := &_MPropWriteReq{
		_CEMI: &_CEMI{
			Size: size,
		},
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_MPropWriteReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MPropWriteReq) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MPropWriteReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MPropWriteReq")
		}

		if popErr := writeBuffer.PopContext("MPropWriteReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MPropWriteReq")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_MPropWriteReq) isMPropWriteReq() bool {
	return true
}

func (m *_MPropWriteReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
