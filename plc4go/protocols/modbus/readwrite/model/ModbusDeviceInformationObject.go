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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// ModbusDeviceInformationObject is the corresponding interface of ModbusDeviceInformationObject
type ModbusDeviceInformationObject interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetObjectId returns ObjectId (property field)
	GetObjectId() uint8
	// GetData returns Data (property field)
	GetData() []byte
}

// ModbusDeviceInformationObjectExactly can be used when we want exactly this type and not a type which fulfills ModbusDeviceInformationObject.
// This is useful for switch cases.
type ModbusDeviceInformationObjectExactly interface {
	ModbusDeviceInformationObject
	isModbusDeviceInformationObject() bool
}

// _ModbusDeviceInformationObject is the data-structure of this message
type _ModbusDeviceInformationObject struct {
	ObjectId uint8
	Data     []byte
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModbusDeviceInformationObject) GetObjectId() uint8 {
	return m.ObjectId
}

func (m *_ModbusDeviceInformationObject) GetData() []byte {
	return m.Data
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewModbusDeviceInformationObject factory function for _ModbusDeviceInformationObject
func NewModbusDeviceInformationObject(objectId uint8, data []byte) *_ModbusDeviceInformationObject {
	return &_ModbusDeviceInformationObject{ObjectId: objectId, Data: data}
}

// Deprecated: use the interface for direct cast
func CastModbusDeviceInformationObject(structType any) ModbusDeviceInformationObject {
	if casted, ok := structType.(ModbusDeviceInformationObject); ok {
		return casted
	}
	if casted, ok := structType.(*ModbusDeviceInformationObject); ok {
		return *casted
	}
	return nil
}

func (m *_ModbusDeviceInformationObject) GetTypeName() string {
	return "ModbusDeviceInformationObject"
}

func (m *_ModbusDeviceInformationObject) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectId)
	lengthInBits += 8

	// Implicit Field (objectLength)
	lengthInBits += 8

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_ModbusDeviceInformationObject) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModbusDeviceInformationObjectParse(theBytes []byte) (ModbusDeviceInformationObject, error) {
	return ModbusDeviceInformationObjectParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func ModbusDeviceInformationObjectParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (ModbusDeviceInformationObject, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ModbusDeviceInformationObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModbusDeviceInformationObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectId)
	_objectId, _objectIdErr := readBuffer.ReadUint8("objectId", 8)
	if _objectIdErr != nil {
		return nil, errors.Wrap(_objectIdErr, "Error parsing 'objectId' field of ModbusDeviceInformationObject")
	}
	objectId := _objectId

	// Implicit Field (objectLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	objectLength, _objectLengthErr := readBuffer.ReadUint8("objectLength", 8)
	_ = objectLength
	if _objectLengthErr != nil {
		return nil, errors.Wrap(_objectLengthErr, "Error parsing 'objectLength' field of ModbusDeviceInformationObject")
	}
	// Byte Array field (data)
	numberOfBytesdata := int(objectLength)
	data, _readArrayErr := readBuffer.ReadByteArray("data", numberOfBytesdata)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'data' field of ModbusDeviceInformationObject")
	}

	if closeErr := readBuffer.CloseContext("ModbusDeviceInformationObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModbusDeviceInformationObject")
	}

	// Create the instance
	return &_ModbusDeviceInformationObject{
		ObjectId: objectId,
		Data:     data,
	}, nil
}

func (m *_ModbusDeviceInformationObject) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModbusDeviceInformationObject) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("ModbusDeviceInformationObject"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ModbusDeviceInformationObject")
	}

	// Simple Field (objectId)
	objectId := uint8(m.GetObjectId())
	_objectIdErr := writeBuffer.WriteUint8("objectId", 8, (objectId))
	if _objectIdErr != nil {
		return errors.Wrap(_objectIdErr, "Error serializing 'objectId' field")
	}

	// Implicit Field (objectLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	objectLength := uint8(uint8(len(m.GetData())))
	_objectLengthErr := writeBuffer.WriteUint8("objectLength", 8, (objectLength))
	if _objectLengthErr != nil {
		return errors.Wrap(_objectLengthErr, "Error serializing 'objectLength' field")
	}

	// Array Field (data)
	// Byte Array field (data)
	if err := writeBuffer.WriteByteArray("data", m.GetData()); err != nil {
		return errors.Wrap(err, "Error serializing 'data' field")
	}

	if popErr := writeBuffer.PopContext("ModbusDeviceInformationObject"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ModbusDeviceInformationObject")
	}
	return nil
}

func (m *_ModbusDeviceInformationObject) isModbusDeviceInformationObject() bool {
	return true
}

func (m *_ModbusDeviceInformationObject) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
