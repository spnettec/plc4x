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

// MFuncPropCon is the corresponding interface of MFuncPropCon
type MFuncPropCon interface {
	utils.LengthAware
	utils.Serializable
	CEMI
}

// MFuncPropConExactly can be used when we want exactly this type and not a type which fulfills MFuncPropCon.
// This is useful for switch cases.
type MFuncPropConExactly interface {
	MFuncPropCon
	isMFuncPropCon() bool
}

// _MFuncPropCon is the data-structure of this message
type _MFuncPropCon struct {
	*_CEMI
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MFuncPropCon) GetMessageCode() uint8 {
	return 0xFA
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MFuncPropCon) InitializeParent(parent CEMI) {}

func (m *_MFuncPropCon) GetParent() CEMI {
	return m._CEMI
}

// NewMFuncPropCon factory function for _MFuncPropCon
func NewMFuncPropCon(size uint16) *_MFuncPropCon {
	_result := &_MFuncPropCon{
		_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMFuncPropCon(structType interface{}) MFuncPropCon {
	if casted, ok := structType.(MFuncPropCon); ok {
		return casted
	}
	if casted, ok := structType.(*MFuncPropCon); ok {
		return *casted
	}
	return nil
}

func (m *_MFuncPropCon) GetTypeName() string {
	return "MFuncPropCon"
}

func (m *_MFuncPropCon) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_MFuncPropCon) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_MFuncPropCon) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MFuncPropConParse(readBuffer utils.ReadBuffer, size uint16) (MFuncPropCon, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MFuncPropCon"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MFuncPropCon")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("MFuncPropCon"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MFuncPropCon")
	}

	// Create a partially initialized instance
	_child := &_MFuncPropCon{
		_CEMI: &_CEMI{
			Size: size,
		},
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_MFuncPropCon) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MFuncPropCon"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MFuncPropCon")
		}

		if popErr := writeBuffer.PopContext("MFuncPropCon"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MFuncPropCon")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_MFuncPropCon) isMFuncPropCon() bool {
	return true
}

func (m *_MFuncPropCon) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
