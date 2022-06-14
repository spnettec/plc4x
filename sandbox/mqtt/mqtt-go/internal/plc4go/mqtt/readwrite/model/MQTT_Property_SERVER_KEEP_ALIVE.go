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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// The data-structure of this message
type MQTT_Property_SERVER_KEEP_ALIVE struct {
	*MQTT_Property
	Value uint16
}

// The corresponding interface
type IMQTT_Property_SERVER_KEEP_ALIVE interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////

func (m *MQTT_Property_SERVER_KEEP_ALIVE) InitializeParent(parent *MQTT_Property, propertyType MQTT_PropertyType) {
	m.PropertyType = propertyType
}

func NewMQTT_Property_SERVER_KEEP_ALIVE(value uint16, propertyType MQTT_PropertyType) *MQTT_Property {
	child := &MQTT_Property_SERVER_KEEP_ALIVE{
		Value: value,
    	MQTT_Property: NewMQTT_Property(propertyType),
	}
	child.Child = child
	return child.MQTT_Property
}

func CastMQTT_Property_SERVER_KEEP_ALIVE(structType interface{}) *MQTT_Property_SERVER_KEEP_ALIVE {
	castFunc := func(typ interface{}) *MQTT_Property_SERVER_KEEP_ALIVE {
		if casted, ok := typ.(MQTT_Property_SERVER_KEEP_ALIVE); ok {
			return &casted
		}
		if casted, ok := typ.(*MQTT_Property_SERVER_KEEP_ALIVE); ok {
			return casted
		}
		if casted, ok := typ.(MQTT_Property); ok {
			return CastMQTT_Property_SERVER_KEEP_ALIVE(casted.Child)
		}
		if casted, ok := typ.(*MQTT_Property); ok {
			return CastMQTT_Property_SERVER_KEEP_ALIVE(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *MQTT_Property_SERVER_KEEP_ALIVE) GetTypeName() string {
	return "MQTT_Property_SERVER_KEEP_ALIVE"
}

func (m *MQTT_Property_SERVER_KEEP_ALIVE) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *MQTT_Property_SERVER_KEEP_ALIVE) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (value)
	lengthInBits += 16;

	return lengthInBits
}


func (m *MQTT_Property_SERVER_KEEP_ALIVE) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func MQTT_Property_SERVER_KEEP_ALIVEParse(readBuffer utils.ReadBuffer) (*MQTT_Property, error) {
	if pullErr := readBuffer.PullContext("MQTT_Property_SERVER_KEEP_ALIVE"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (value)
_value, _valueErr := readBuffer.ReadUint16("value", 16)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value

	if closeErr := readBuffer.CloseContext("MQTT_Property_SERVER_KEEP_ALIVE"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &MQTT_Property_SERVER_KEEP_ALIVE{
		Value: value,
        MQTT_Property: &MQTT_Property{},
	}
	_child.MQTT_Property.Child = _child
	return _child.MQTT_Property, nil
}

func (m *MQTT_Property_SERVER_KEEP_ALIVE) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MQTT_Property_SERVER_KEEP_ALIVE"); pushErr != nil {
			return pushErr
		}

	// Simple Field (value)
	value := uint16(m.Value)
	_valueErr := writeBuffer.WriteUint16("value", 16, (value))
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

		if popErr := writeBuffer.PopContext("MQTT_Property_SERVER_KEEP_ALIVE"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *MQTT_Property_SERVER_KEEP_ALIVE) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}



