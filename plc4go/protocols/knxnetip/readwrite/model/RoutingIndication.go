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


// RoutingIndication is the corresponding interface of RoutingIndication
type RoutingIndication interface {
	utils.LengthAware
	utils.Serializable
	KnxNetIpMessage
}

// RoutingIndicationExactly can be used when we want exactly this type and not a type which fulfills RoutingIndication.
// This is useful for switch cases.
type RoutingIndicationExactly interface {
	RoutingIndication
	isRoutingIndication() bool
}

// _RoutingIndication is the data-structure of this message
type _RoutingIndication struct {
	*_KnxNetIpMessage
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RoutingIndication)  GetMsgType() uint16 {
return 0x0530}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RoutingIndication) InitializeParent(parent KnxNetIpMessage ) {}

func (m *_RoutingIndication)  GetParent() KnxNetIpMessage {
	return m._KnxNetIpMessage
}


// NewRoutingIndication factory function for _RoutingIndication
func NewRoutingIndication( ) *_RoutingIndication {
	_result := &_RoutingIndication{
    	_KnxNetIpMessage: NewKnxNetIpMessage(),
	}
	_result._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRoutingIndication(structType interface{}) RoutingIndication {
    if casted, ok := structType.(RoutingIndication); ok {
		return casted
	}
	if casted, ok := structType.(*RoutingIndication); ok {
		return *casted
	}
	return nil
}

func (m *_RoutingIndication) GetTypeName() string {
	return "RoutingIndication"
}

func (m *_RoutingIndication) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_RoutingIndication) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}


func (m *_RoutingIndication) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func RoutingIndicationParse(theBytes []byte) (RoutingIndication, error) {
	return RoutingIndicationParseWithBuffer(utils.NewReadBufferByteBased(theBytes, utils.WithByteOrderForReadBufferByteBased(binary.BigEndian)))
}

func RoutingIndicationParseWithBuffer(readBuffer utils.ReadBuffer) (RoutingIndication, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RoutingIndication"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RoutingIndication")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("RoutingIndication"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RoutingIndication")
	}

	// Create a partially initialized instance
	_child := &_RoutingIndication{
		_KnxNetIpMessage: &_KnxNetIpMessage{
		},
	}
	_child._KnxNetIpMessage._KnxNetIpMessageChildRequirements = _child
	return _child, nil
}

func (m *_RoutingIndication) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())), utils.WithByteOrderForByteBasedBuffer(binary.BigEndian))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RoutingIndication) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RoutingIndication"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RoutingIndication")
		}

		if popErr := writeBuffer.PopContext("RoutingIndication"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RoutingIndication")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_RoutingIndication) isRoutingIndication() bool {
	return true
}

func (m *_RoutingIndication) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



