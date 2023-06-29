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
	"github.com/rs/zerolog"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetDeviceObjectPropertyReference is the corresponding interface of BACnetDeviceObjectPropertyReference
type BACnetDeviceObjectPropertyReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() BACnetContextTagObjectIdentifier
}

// BACnetDeviceObjectPropertyReferenceExactly can be used when we want exactly this type and not a type which fulfills BACnetDeviceObjectPropertyReference.
// This is useful for switch cases.
type BACnetDeviceObjectPropertyReferenceExactly interface {
	BACnetDeviceObjectPropertyReference
	isBACnetDeviceObjectPropertyReference() bool
}

// _BACnetDeviceObjectPropertyReference is the data-structure of this message
type _BACnetDeviceObjectPropertyReference struct {
	ObjectIdentifier   BACnetContextTagObjectIdentifier
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
	DeviceIdentifier   BACnetContextTagObjectIdentifier
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetDeviceObjectPropertyReference) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetDeviceObjectPropertyReference) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetDeviceObjectPropertyReference) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *_BACnetDeviceObjectPropertyReference) GetDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.DeviceIdentifier
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetDeviceObjectPropertyReference factory function for _BACnetDeviceObjectPropertyReference
func NewBACnetDeviceObjectPropertyReference(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger, deviceIdentifier BACnetContextTagObjectIdentifier) *_BACnetDeviceObjectPropertyReference {
	return &_BACnetDeviceObjectPropertyReference{ObjectIdentifier: objectIdentifier, PropertyIdentifier: propertyIdentifier, ArrayIndex: arrayIndex, DeviceIdentifier: deviceIdentifier}
}

// Deprecated: use the interface for direct cast
func CastBACnetDeviceObjectPropertyReference(structType any) BACnetDeviceObjectPropertyReference {
	if casted, ok := structType.(BACnetDeviceObjectPropertyReference); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetDeviceObjectPropertyReference); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetDeviceObjectPropertyReference) GetTypeName() string {
	return "BACnetDeviceObjectPropertyReference"
}

func (m *_BACnetDeviceObjectPropertyReference) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits(ctx)
	}

	// Optional Field (deviceIdentifier)
	if m.DeviceIdentifier != nil {
		lengthInBits += m.DeviceIdentifier.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetDeviceObjectPropertyReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetDeviceObjectPropertyReferenceParse(ctx context.Context, theBytes []byte) (BACnetDeviceObjectPropertyReference, error) {
	return BACnetDeviceObjectPropertyReferenceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetDeviceObjectPropertyReferenceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetDeviceObjectPropertyReference, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetDeviceObjectPropertyReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetDeviceObjectPropertyReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field of BACnetDeviceObjectPropertyReference")
	}
	objectIdentifier := _objectIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyIdentifier")
	}
	_propertyIdentifier, _propertyIdentifierErr := BACnetPropertyIdentifierTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field of BACnetDeviceObjectPropertyReference")
	}
	propertyIdentifier := _propertyIdentifier.(BACnetPropertyIdentifierTagged)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyIdentifier")
	}

	// Optional Field (arrayIndex) (Can be skipped, if a given expression evaluates to false)
	var arrayIndex BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("arrayIndex"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for arrayIndex")
		}
		_val, _err := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(2), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'arrayIndex' field of BACnetDeviceObjectPropertyReference")
		default:
			arrayIndex = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("arrayIndex"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for arrayIndex")
			}
		}
	}

	// Optional Field (deviceIdentifier) (Can be skipped, if a given expression evaluates to false)
	var deviceIdentifier BACnetContextTagObjectIdentifier = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("deviceIdentifier"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for deviceIdentifier")
		}
		_val, _err := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(3), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'deviceIdentifier' field of BACnetDeviceObjectPropertyReference")
		default:
			deviceIdentifier = _val.(BACnetContextTagObjectIdentifier)
			if closeErr := readBuffer.CloseContext("deviceIdentifier"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for deviceIdentifier")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetDeviceObjectPropertyReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetDeviceObjectPropertyReference")
	}

	// Create the instance
	return &_BACnetDeviceObjectPropertyReference{
		ObjectIdentifier:   objectIdentifier,
		PropertyIdentifier: propertyIdentifier,
		ArrayIndex:         arrayIndex,
		DeviceIdentifier:   deviceIdentifier,
	}, nil
}

func (m *_BACnetDeviceObjectPropertyReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetDeviceObjectPropertyReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := writeBuffer.PushContext("BACnetDeviceObjectPropertyReference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetDeviceObjectPropertyReference")
	}

	// Simple Field (objectIdentifier)
	if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
	}
	_objectIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetObjectIdentifier())
	if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for objectIdentifier")
	}
	if _objectIdentifierErr != nil {
		return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
	}

	// Simple Field (propertyIdentifier)
	if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for propertyIdentifier")
	}
	_propertyIdentifierErr := writeBuffer.WriteSerializable(ctx, m.GetPropertyIdentifier())
	if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for propertyIdentifier")
	}
	if _propertyIdentifierErr != nil {
		return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
	}

	// Optional Field (arrayIndex) (Can be skipped, if the value is null)
	var arrayIndex BACnetContextTagUnsignedInteger = nil
	if m.GetArrayIndex() != nil {
		if pushErr := writeBuffer.PushContext("arrayIndex"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for arrayIndex")
		}
		arrayIndex = m.GetArrayIndex()
		_arrayIndexErr := writeBuffer.WriteSerializable(ctx, arrayIndex)
		if popErr := writeBuffer.PopContext("arrayIndex"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for arrayIndex")
		}
		if _arrayIndexErr != nil {
			return errors.Wrap(_arrayIndexErr, "Error serializing 'arrayIndex' field")
		}
	}

	// Optional Field (deviceIdentifier) (Can be skipped, if the value is null)
	var deviceIdentifier BACnetContextTagObjectIdentifier = nil
	if m.GetDeviceIdentifier() != nil {
		if pushErr := writeBuffer.PushContext("deviceIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deviceIdentifier")
		}
		deviceIdentifier = m.GetDeviceIdentifier()
		_deviceIdentifierErr := writeBuffer.WriteSerializable(ctx, deviceIdentifier)
		if popErr := writeBuffer.PopContext("deviceIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deviceIdentifier")
		}
		if _deviceIdentifierErr != nil {
			return errors.Wrap(_deviceIdentifierErr, "Error serializing 'deviceIdentifier' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetDeviceObjectPropertyReference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetDeviceObjectPropertyReference")
	}
	return nil
}

func (m *_BACnetDeviceObjectPropertyReference) isBACnetDeviceObjectPropertyReference() bool {
	return true
}

func (m *_BACnetDeviceObjectPropertyReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
