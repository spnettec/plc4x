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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type CEMI struct {

	// Arguments.
	Size  uint16
	Child ICEMIChild
}

// The corresponding interface
type ICEMI interface {
	// GetMessageCode returns MessageCode (discriminator field)
	GetMessageCode() uint8
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

type ICEMIParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child ICEMI, serializeChildFunction func() error) error
	GetTypeName() string
}

type ICEMIChild interface {
	Serialize(writeBuffer utils.WriteBuffer) error
	InitializeParent(parent *CEMI)
	GetTypeName() string
	ICEMI
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewCEMI factory function for CEMI
func NewCEMI(size uint16) *CEMI {
	return &CEMI{Size: size}
}

func CastCEMI(structType interface{}) *CEMI {
	if casted, ok := structType.(CEMI); ok {
		return &casted
	}
	if casted, ok := structType.(*CEMI); ok {
		return casted
	}
	return nil
}

func (m *CEMI) GetTypeName() string {
	return "CEMI"
}

func (m *CEMI) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *CEMI) GetLengthInBitsConditional(lastItem bool) uint16 {
	return m.Child.GetLengthInBits()
}

func (m *CEMI) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)
	// Discriminator Field (messageCode)
	lengthInBits += 8

	return lengthInBits
}

func (m *CEMI) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CEMIParse(readBuffer utils.ReadBuffer, size uint16) (*CEMI, error) {
	if pullErr := readBuffer.PullContext("CEMI"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Discriminator Field (messageCode) (Used as input to a switch field)
	messageCode, _messageCodeErr := readBuffer.ReadUint8("messageCode", 8)
	if _messageCodeErr != nil {
		return nil, errors.Wrap(_messageCodeErr, "Error parsing 'messageCode' field")
	}

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	var _parent *CEMI
	var typeSwitchError error
	switch {
	case messageCode == 0x2B: // LBusmonInd
		_parent, typeSwitchError = LBusmonIndParse(readBuffer, size)
	case messageCode == 0x11: // LDataReq
		_parent, typeSwitchError = LDataReqParse(readBuffer, size)
	case messageCode == 0x29: // LDataInd
		_parent, typeSwitchError = LDataIndParse(readBuffer, size)
	case messageCode == 0x2E: // LDataCon
		_parent, typeSwitchError = LDataConParse(readBuffer, size)
	case messageCode == 0x10: // LRawReq
		_parent, typeSwitchError = LRawReqParse(readBuffer, size)
	case messageCode == 0x2D: // LRawInd
		_parent, typeSwitchError = LRawIndParse(readBuffer, size)
	case messageCode == 0x2F: // LRawCon
		_parent, typeSwitchError = LRawConParse(readBuffer, size)
	case messageCode == 0x13: // LPollDataReq
		_parent, typeSwitchError = LPollDataReqParse(readBuffer, size)
	case messageCode == 0x25: // LPollDataCon
		_parent, typeSwitchError = LPollDataConParse(readBuffer, size)
	case messageCode == 0x41: // TDataConnectedReq
		_parent, typeSwitchError = TDataConnectedReqParse(readBuffer, size)
	case messageCode == 0x89: // TDataConnectedInd
		_parent, typeSwitchError = TDataConnectedIndParse(readBuffer, size)
	case messageCode == 0x4A: // TDataIndividualReq
		_parent, typeSwitchError = TDataIndividualReqParse(readBuffer, size)
	case messageCode == 0x94: // TDataIndividualInd
		_parent, typeSwitchError = TDataIndividualIndParse(readBuffer, size)
	case messageCode == 0xFC: // MPropReadReq
		_parent, typeSwitchError = MPropReadReqParse(readBuffer, size)
	case messageCode == 0xFB: // MPropReadCon
		_parent, typeSwitchError = MPropReadConParse(readBuffer, size)
	case messageCode == 0xF6: // MPropWriteReq
		_parent, typeSwitchError = MPropWriteReqParse(readBuffer, size)
	case messageCode == 0xF5: // MPropWriteCon
		_parent, typeSwitchError = MPropWriteConParse(readBuffer, size)
	case messageCode == 0xF7: // MPropInfoInd
		_parent, typeSwitchError = MPropInfoIndParse(readBuffer, size)
	case messageCode == 0xF8: // MFuncPropCommandReq
		_parent, typeSwitchError = MFuncPropCommandReqParse(readBuffer, size)
	case messageCode == 0xF9: // MFuncPropStateReadReq
		_parent, typeSwitchError = MFuncPropStateReadReqParse(readBuffer, size)
	case messageCode == 0xFA: // MFuncPropCon
		_parent, typeSwitchError = MFuncPropConParse(readBuffer, size)
	case messageCode == 0xF1: // MResetReq
		_parent, typeSwitchError = MResetReqParse(readBuffer, size)
	case messageCode == 0xF0: // MResetInd
		_parent, typeSwitchError = MResetIndParse(readBuffer, size)
	default:
		// TODO: return actual type
		typeSwitchError = errors.New("Unmapped type")
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch.")
	}

	if closeErr := readBuffer.CloseContext("CEMI"); closeErr != nil {
		return nil, closeErr
	}

	// Finish initializing
	_parent.Child.InitializeParent(_parent)
	return _parent, nil
}

func (m *CEMI) Serialize(writeBuffer utils.WriteBuffer) error {
	return m.Child.Serialize(writeBuffer)
}

func (m *CEMI) SerializeParent(writeBuffer utils.WriteBuffer, child ICEMI, serializeChildFunction func() error) error {
	if pushErr := writeBuffer.PushContext("CEMI"); pushErr != nil {
		return pushErr
	}

	// Discriminator Field (messageCode) (Used as input to a switch field)
	messageCode := uint8(child.GetMessageCode())
	_messageCodeErr := writeBuffer.WriteUint8("messageCode", 8, (messageCode))

	if _messageCodeErr != nil {
		return errors.Wrap(_messageCodeErr, "Error serializing 'messageCode' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("CEMI"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *CEMI) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
