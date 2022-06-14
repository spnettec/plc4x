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
type MQTT_Property_USER_PROPERTY struct {
	*MQTT_Property
	Name *MQTT_String
	Value *MQTT_String
}

// The corresponding interface
type IMQTT_Property_USER_PROPERTY interface {
	LengthInBytes() uint16
	LengthInBits() uint16
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////

func (m *MQTT_Property_USER_PROPERTY) InitializeParent(parent *MQTT_Property, propertyType MQTT_PropertyType) {
	m.PropertyType = propertyType
}

func NewMQTT_Property_USER_PROPERTY(name *MQTT_String, value *MQTT_String, propertyType MQTT_PropertyType) *MQTT_Property {
	child := &MQTT_Property_USER_PROPERTY{
		Name: name,
		Value: value,
    	MQTT_Property: NewMQTT_Property(propertyType),
	}
	child.Child = child
	return child.MQTT_Property
}

func CastMQTT_Property_USER_PROPERTY(structType interface{}) *MQTT_Property_USER_PROPERTY {
	castFunc := func(typ interface{}) *MQTT_Property_USER_PROPERTY {
		if casted, ok := typ.(MQTT_Property_USER_PROPERTY); ok {
			return &casted
		}
		if casted, ok := typ.(*MQTT_Property_USER_PROPERTY); ok {
			return casted
		}
		if casted, ok := typ.(MQTT_Property); ok {
			return CastMQTT_Property_USER_PROPERTY(casted.Child)
		}
		if casted, ok := typ.(*MQTT_Property); ok {
			return CastMQTT_Property_USER_PROPERTY(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *MQTT_Property_USER_PROPERTY) GetTypeName() string {
	return "MQTT_Property_USER_PROPERTY"
}

func (m *MQTT_Property_USER_PROPERTY) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *MQTT_Property_USER_PROPERTY) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (name)
	lengthInBits += m.Name.LengthInBits()

	// Simple field (value)
	lengthInBits += m.Value.LengthInBits()

	return lengthInBits
}


func (m *MQTT_Property_USER_PROPERTY) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func MQTT_Property_USER_PROPERTYParse(readBuffer utils.ReadBuffer) (*MQTT_Property, error) {
	if pullErr := readBuffer.PullContext("MQTT_Property_USER_PROPERTY"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (name)
	if pullErr := readBuffer.PullContext("name"); pullErr != nil {
		return nil, pullErr
	}
_name, _nameErr := MQTT_StringParse(readBuffer)
	if _nameErr != nil {
		return nil, errors.Wrap(_nameErr, "Error parsing 'name' field")
	}
	name := CastMQTT_String(_name)
	if closeErr := readBuffer.CloseContext("name"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (value)
	if pullErr := readBuffer.PullContext("value"); pullErr != nil {
		return nil, pullErr
	}
_value, _valueErr := MQTT_StringParse(readBuffer)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := CastMQTT_String(_value)
	if closeErr := readBuffer.CloseContext("value"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("MQTT_Property_USER_PROPERTY"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &MQTT_Property_USER_PROPERTY{
		Name: CastMQTT_String(name),
		Value: CastMQTT_String(value),
        MQTT_Property: &MQTT_Property{},
	}
	_child.MQTT_Property.Child = _child
	return _child.MQTT_Property, nil
}

func (m *MQTT_Property_USER_PROPERTY) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MQTT_Property_USER_PROPERTY"); pushErr != nil {
			return pushErr
		}

	// Simple Field (name)
	if pushErr := writeBuffer.PushContext("name"); pushErr != nil {
		return pushErr
	}
	_nameErr := m.Name.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("name"); popErr != nil {
		return popErr
	}
	if _nameErr != nil {
		return errors.Wrap(_nameErr, "Error serializing 'name' field")
	}

	// Simple Field (value)
	if pushErr := writeBuffer.PushContext("value"); pushErr != nil {
		return pushErr
	}
	_valueErr := m.Value.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("value"); popErr != nil {
		return popErr
	}
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

		if popErr := writeBuffer.PopContext("MQTT_Property_USER_PROPERTY"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *MQTT_Property_USER_PROPERTY) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}



