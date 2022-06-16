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
)

// Code generated by code-generation. DO NOT EDIT.

// TDataIndividualReq is the corresponding interface of TDataIndividualReq
type TDataIndividualReq interface {
	CEMI
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _TDataIndividualReq is the data-structure of this message
type _TDataIndividualReq struct {
	*_CEMI

	// Arguments.
	Size uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TDataIndividualReq) GetMessageCode() uint8 {
	return 0x4A
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TDataIndividualReq) InitializeParent(parent CEMI) {}

func (m *_TDataIndividualReq) GetParent() CEMI {
	return m._CEMI
}

// NewTDataIndividualReq factory function for _TDataIndividualReq
func NewTDataIndividualReq(size uint16) *_TDataIndividualReq {
	_result := &_TDataIndividualReq{
		_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTDataIndividualReq(structType interface{}) TDataIndividualReq {
	if casted, ok := structType.(TDataIndividualReq); ok {
		return casted
	}
	if casted, ok := structType.(*TDataIndividualReq); ok {
		return *casted
	}
	return nil
}

func (m *_TDataIndividualReq) GetTypeName() string {
	return "TDataIndividualReq"
}

func (m *_TDataIndividualReq) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_TDataIndividualReq) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_TDataIndividualReq) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func TDataIndividualReqParse(readBuffer utils.ReadBuffer, size uint16) (TDataIndividualReq, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("TDataIndividualReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TDataIndividualReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TDataIndividualReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TDataIndividualReq")
	}

	// Create a partially initialized instance
	_child := &_TDataIndividualReq{
		_CEMI: &_CEMI{},
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_TDataIndividualReq) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TDataIndividualReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TDataIndividualReq")
		}

		if popErr := writeBuffer.PopContext("TDataIndividualReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TDataIndividualReq")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_TDataIndividualReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
