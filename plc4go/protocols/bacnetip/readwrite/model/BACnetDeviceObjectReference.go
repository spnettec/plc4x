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
	"github.com/rs/zerolog/log"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetDeviceObjectReference is the corresponding interface of BACnetDeviceObjectReference
type BACnetDeviceObjectReference interface {
	utils.LengthAware
	utils.Serializable
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
}

// BACnetDeviceObjectReferenceExactly can be used when we want exactly this type and not a type which fulfills BACnetDeviceObjectReference.
// This is useful for switch cases.
type BACnetDeviceObjectReferenceExactly interface {
	BACnetDeviceObjectReference
	isBACnetDeviceObjectReference() bool
}

// _BACnetDeviceObjectReference is the data-structure of this message
type _BACnetDeviceObjectReference struct {
	DeviceIdentifier BACnetContextTagObjectIdentifier
	ObjectIdentifier BACnetContextTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDeviceObjectReference) GetDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.DeviceIdentifier
}

func (m *_BACnetDeviceObjectReference) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDeviceObjectReference factory function for _BACnetDeviceObjectReference
func NewBACnetDeviceObjectReference(deviceIdentifier BACnetContextTagObjectIdentifier, objectIdentifier BACnetContextTagObjectIdentifier) *_BACnetDeviceObjectReference {
	return &_BACnetDeviceObjectReference{DeviceIdentifier: deviceIdentifier, ObjectIdentifier: objectIdentifier}
}

// Deprecated: use the interface for direct cast
func CastBACnetDeviceObjectReference(structType interface{}) BACnetDeviceObjectReference {
	if casted, ok := structType.(BACnetDeviceObjectReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDeviceObjectReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDeviceObjectReference) GetTypeName() string {
	return "BACnetDeviceObjectReference"
}

func (m *_BACnetDeviceObjectReference) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetDeviceObjectReference) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Optional Field (deviceIdentifier)
	if m.DeviceIdentifier != nil {
		lengthInBits += m.DeviceIdentifier.GetLengthInBits()
	}

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetDeviceObjectReference) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetDeviceObjectReferenceParse(readBuffer utils.ReadBuffer) (BACnetDeviceObjectReference, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetDeviceObjectReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDeviceObjectReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Optional Field (deviceIdentifier) (Can be skipped, if a given expression evaluates to false)
	var deviceIdentifier BACnetContextTagObjectIdentifier = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("deviceIdentifier"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for deviceIdentifier")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(0), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'deviceIdentifier' field of BACnetDeviceObjectReference")
		default:
			deviceIdentifier = _val.(BACnetContextTagObjectIdentifier)
			if closeErr := readBuffer.CloseContext("deviceIdentifier"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for deviceIdentifier")
			}
		}
	}

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field of BACnetDeviceObjectReference")
	}
	objectIdentifier := _objectIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	if closeErr := readBuffer.CloseContext("BACnetDeviceObjectReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDeviceObjectReference")
	}

	// Create the instance
	return NewBACnetDeviceObjectReference(deviceIdentifier, objectIdentifier), nil
}

func (m *_BACnetDeviceObjectReference) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetDeviceObjectReference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDeviceObjectReference")
	}

	// Optional Field (deviceIdentifier) (Can be skipped, if the value is null)
	var deviceIdentifier BACnetContextTagObjectIdentifier = nil
	if m.GetDeviceIdentifier() != nil {
		if pushErr := writeBuffer.PushContext("deviceIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deviceIdentifier")
		}
		deviceIdentifier = m.GetDeviceIdentifier()
		_deviceIdentifierErr := writeBuffer.WriteSerializable(deviceIdentifier)
		if popErr := writeBuffer.PopContext("deviceIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deviceIdentifier")
		}
		if _deviceIdentifierErr != nil {
			return errors.Wrap(_deviceIdentifierErr, "Error serializing 'deviceIdentifier' field")
		}
	}

	// Simple Field (objectIdentifier)
	if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
	}
	_objectIdentifierErr := writeBuffer.WriteSerializable(m.GetObjectIdentifier())
	if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for objectIdentifier")
	}
	if _objectIdentifierErr != nil {
		return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
	}

	if popErr := writeBuffer.PopContext("BACnetDeviceObjectReference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDeviceObjectReference")
	}
	return nil
}

func (m *_BACnetDeviceObjectReference) isBACnetDeviceObjectReference() bool {
	return true
}

func (m *_BACnetDeviceObjectReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
