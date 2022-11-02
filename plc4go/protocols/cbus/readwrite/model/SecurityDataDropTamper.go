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


// SecurityDataDropTamper is the corresponding interface of SecurityDataDropTamper
type SecurityDataDropTamper interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataDropTamperExactly can be used when we want exactly this type and not a type which fulfills SecurityDataDropTamper.
// This is useful for switch cases.
type SecurityDataDropTamperExactly interface {
	SecurityDataDropTamper
	isSecurityDataDropTamper() bool
}

// _SecurityDataDropTamper is the data-structure of this message
type _SecurityDataDropTamper struct {
	*_SecurityData
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataDropTamper) InitializeParent(parent SecurityData , commandTypeContainer SecurityCommandTypeContainer , argument byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataDropTamper)  GetParent() SecurityData {
	return m._SecurityData
}


// NewSecurityDataDropTamper factory function for _SecurityDataDropTamper
func NewSecurityDataDropTamper( commandTypeContainer SecurityCommandTypeContainer , argument byte ) *_SecurityDataDropTamper {
	_result := &_SecurityDataDropTamper{
    	_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataDropTamper(structType interface{}) SecurityDataDropTamper {
    if casted, ok := structType.(SecurityDataDropTamper); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataDropTamper); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataDropTamper) GetTypeName() string {
	return "SecurityDataDropTamper"
}

func (m *_SecurityDataDropTamper) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataDropTamper) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}


func (m *_SecurityDataDropTamper) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataDropTamperParse(readBuffer utils.ReadBuffer) (SecurityDataDropTamper, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataDropTamper"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataDropTamper")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataDropTamper"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataDropTamper")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataDropTamper{
		_SecurityData: &_SecurityData{
		},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataDropTamper) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataDropTamper) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataDropTamper"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataDropTamper")
		}

		if popErr := writeBuffer.PopContext("SecurityDataDropTamper"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataDropTamper")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_SecurityDataDropTamper) isSecurityDataDropTamper() bool {
	return true
}

func (m *_SecurityDataDropTamper) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



