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


// SecurityDataSystemDisarmed is the corresponding interface of SecurityDataSystemDisarmed
type SecurityDataSystemDisarmed interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
}

// SecurityDataSystemDisarmedExactly can be used when we want exactly this type and not a type which fulfills SecurityDataSystemDisarmed.
// This is useful for switch cases.
type SecurityDataSystemDisarmedExactly interface {
	SecurityDataSystemDisarmed
	isSecurityDataSystemDisarmed() bool
}

// _SecurityDataSystemDisarmed is the data-structure of this message
type _SecurityDataSystemDisarmed struct {
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

func (m *_SecurityDataSystemDisarmed) InitializeParent(parent SecurityData , commandTypeContainer SecurityCommandTypeContainer , argument byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataSystemDisarmed)  GetParent() SecurityData {
	return m._SecurityData
}


// NewSecurityDataSystemDisarmed factory function for _SecurityDataSystemDisarmed
func NewSecurityDataSystemDisarmed( commandTypeContainer SecurityCommandTypeContainer , argument byte ) *_SecurityDataSystemDisarmed {
	_result := &_SecurityDataSystemDisarmed{
    	_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataSystemDisarmed(structType interface{}) SecurityDataSystemDisarmed {
    if casted, ok := structType.(SecurityDataSystemDisarmed); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataSystemDisarmed); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataSystemDisarmed) GetTypeName() string {
	return "SecurityDataSystemDisarmed"
}

func (m *_SecurityDataSystemDisarmed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataSystemDisarmed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}


func (m *_SecurityDataSystemDisarmed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataSystemDisarmedParse(theBytes []byte) (SecurityDataSystemDisarmed, error) {
	return SecurityDataSystemDisarmedParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func SecurityDataSystemDisarmedParseWithBuffer(readBuffer utils.ReadBuffer) (SecurityDataSystemDisarmed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataSystemDisarmed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataSystemDisarmed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("SecurityDataSystemDisarmed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataSystemDisarmed")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataSystemDisarmed{
		_SecurityData: &_SecurityData{
		},
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataSystemDisarmed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataSystemDisarmed) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataSystemDisarmed"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataSystemDisarmed")
		}

		if popErr := writeBuffer.PopContext("SecurityDataSystemDisarmed"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataSystemDisarmed")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_SecurityDataSystemDisarmed) isSecurityDataSystemDisarmed() bool {
	return true
}

func (m *_SecurityDataSystemDisarmed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



