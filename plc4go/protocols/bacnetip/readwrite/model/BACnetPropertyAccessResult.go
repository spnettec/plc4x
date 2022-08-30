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
	"github.com/rs/zerolog/log"
	"io"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetPropertyAccessResult is the corresponding interface of BACnetPropertyAccessResult
type BACnetPropertyAccessResult interface {
	utils.LengthAware
	utils.Serializable
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetPropertyArrayIndex returns PropertyArrayIndex (property field)
	GetPropertyArrayIndex() BACnetContextTagUnsignedInteger
	// GetDeviceIdentifier returns DeviceIdentifier (property field)
	GetDeviceIdentifier() BACnetContextTagObjectIdentifier
	// GetAccessResult returns AccessResult (property field)
	GetAccessResult() BACnetPropertyAccessResultAccessResult
}

// BACnetPropertyAccessResultExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyAccessResult.
// This is useful for switch cases.
type BACnetPropertyAccessResultExactly interface {
	BACnetPropertyAccessResult
	isBACnetPropertyAccessResult() bool
}

// _BACnetPropertyAccessResult is the data-structure of this message
type _BACnetPropertyAccessResult struct {
        ObjectIdentifier BACnetContextTagObjectIdentifier
        PropertyIdentifier BACnetPropertyIdentifierTagged
        PropertyArrayIndex BACnetContextTagUnsignedInteger
        DeviceIdentifier BACnetContextTagObjectIdentifier
        AccessResult BACnetPropertyAccessResultAccessResult
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyAccessResult) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetPropertyAccessResult) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetPropertyAccessResult) GetPropertyArrayIndex() BACnetContextTagUnsignedInteger {
	return m.PropertyArrayIndex
}

func (m *_BACnetPropertyAccessResult) GetDeviceIdentifier() BACnetContextTagObjectIdentifier {
	return m.DeviceIdentifier
}

func (m *_BACnetPropertyAccessResult) GetAccessResult() BACnetPropertyAccessResultAccessResult {
	return m.AccessResult
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetPropertyAccessResult factory function for _BACnetPropertyAccessResult
func NewBACnetPropertyAccessResult( objectIdentifier BACnetContextTagObjectIdentifier , propertyIdentifier BACnetPropertyIdentifierTagged , propertyArrayIndex BACnetContextTagUnsignedInteger , deviceIdentifier BACnetContextTagObjectIdentifier , accessResult BACnetPropertyAccessResultAccessResult ) *_BACnetPropertyAccessResult {
return &_BACnetPropertyAccessResult{ ObjectIdentifier: objectIdentifier , PropertyIdentifier: propertyIdentifier , PropertyArrayIndex: propertyArrayIndex , DeviceIdentifier: deviceIdentifier , AccessResult: accessResult }
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyAccessResult(structType interface{}) BACnetPropertyAccessResult {
    if casted, ok := structType.(BACnetPropertyAccessResult); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyAccessResult); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyAccessResult) GetTypeName() string {
	return "BACnetPropertyAccessResult"
}

func (m *_BACnetPropertyAccessResult) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyAccessResult) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits()

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits()

	// Optional Field (propertyArrayIndex)
	if m.PropertyArrayIndex != nil {
		lengthInBits += m.PropertyArrayIndex.GetLengthInBits()
	}

	// Optional Field (deviceIdentifier)
	if m.DeviceIdentifier != nil {
		lengthInBits += m.DeviceIdentifier.GetLengthInBits()
	}

	// Simple field (accessResult)
	lengthInBits += m.AccessResult.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetPropertyAccessResult) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyAccessResultParse(readBuffer utils.ReadBuffer) (BACnetPropertyAccessResult, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyAccessResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyAccessResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
_objectIdentifier, _objectIdentifierErr := BACnetContextTagParse(readBuffer , uint8( uint8(0) ) , BACnetDataType( BACnetDataType_BACNET_OBJECT_IDENTIFIER ) )
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field of BACnetPropertyAccessResult")
	}
	objectIdentifier := _objectIdentifier.(BACnetContextTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
	}

	// Simple Field (propertyIdentifier)
	if pullErr := readBuffer.PullContext("propertyIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyIdentifier")
	}
_propertyIdentifier, _propertyIdentifierErr := BACnetPropertyIdentifierTaggedParse(readBuffer , uint8( uint8(1) ) , TagClass( TagClass_CONTEXT_SPECIFIC_TAGS ) )
	if _propertyIdentifierErr != nil {
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field of BACnetPropertyAccessResult")
	}
	propertyIdentifier := _propertyIdentifier.(BACnetPropertyIdentifierTagged)
	if closeErr := readBuffer.CloseContext("propertyIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyIdentifier")
	}

	// Optional Field (propertyArrayIndex) (Can be skipped, if a given expression evaluates to false)
	var propertyArrayIndex BACnetContextTagUnsignedInteger = nil
{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("propertyArrayIndex"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for propertyArrayIndex")
		}
_val, _err := BACnetContextTagParse(readBuffer , uint8(2) , BACnetDataType_UNSIGNED_INTEGER )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyArrayIndex' field of BACnetPropertyAccessResult")
		default:
			propertyArrayIndex = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("propertyArrayIndex"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for propertyArrayIndex")
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
_val, _err := BACnetContextTagParse(readBuffer , uint8(3) , BACnetDataType_BACNET_OBJECT_IDENTIFIER )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'deviceIdentifier' field of BACnetPropertyAccessResult")
		default:
			deviceIdentifier = _val.(BACnetContextTagObjectIdentifier)
			if closeErr := readBuffer.CloseContext("deviceIdentifier"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for deviceIdentifier")
			}
		}
	}

	// Simple Field (accessResult)
	if pullErr := readBuffer.PullContext("accessResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for accessResult")
	}
_accessResult, _accessResultErr := BACnetPropertyAccessResultAccessResultParse(readBuffer , BACnetObjectType( objectIdentifier.GetObjectType() ) , BACnetPropertyIdentifier( propertyIdentifier.GetValue() ) , (CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool(((propertyArrayIndex)) != (nil)), func() interface{} {return CastBACnetTagPayloadUnsignedInteger((propertyArrayIndex).GetPayload())}, func() interface{} {return CastBACnetTagPayloadUnsignedInteger(nil)}))) )
	if _accessResultErr != nil {
		return nil, errors.Wrap(_accessResultErr, "Error parsing 'accessResult' field of BACnetPropertyAccessResult")
	}
	accessResult := _accessResult.(BACnetPropertyAccessResultAccessResult)
	if closeErr := readBuffer.CloseContext("accessResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for accessResult")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyAccessResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyAccessResult")
	}

	// Create the instance
	return &_BACnetPropertyAccessResult{
			ObjectIdentifier: objectIdentifier,
			PropertyIdentifier: propertyIdentifier,
			PropertyArrayIndex: propertyArrayIndex,
			DeviceIdentifier: deviceIdentifier,
			AccessResult: accessResult,
		}, nil
}

func (m *_BACnetPropertyAccessResult) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetPropertyAccessResult"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetPropertyAccessResult")
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

	// Simple Field (propertyIdentifier)
	if pushErr := writeBuffer.PushContext("propertyIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for propertyIdentifier")
	}
	_propertyIdentifierErr := writeBuffer.WriteSerializable(m.GetPropertyIdentifier())
	if popErr := writeBuffer.PopContext("propertyIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for propertyIdentifier")
	}
	if _propertyIdentifierErr != nil {
		return errors.Wrap(_propertyIdentifierErr, "Error serializing 'propertyIdentifier' field")
	}

	// Optional Field (propertyArrayIndex) (Can be skipped, if the value is null)
	var propertyArrayIndex BACnetContextTagUnsignedInteger = nil
	if m.GetPropertyArrayIndex() != nil {
		if pushErr := writeBuffer.PushContext("propertyArrayIndex"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyArrayIndex")
		}
		propertyArrayIndex = m.GetPropertyArrayIndex()
		_propertyArrayIndexErr := writeBuffer.WriteSerializable(propertyArrayIndex)
		if popErr := writeBuffer.PopContext("propertyArrayIndex"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyArrayIndex")
		}
		if _propertyArrayIndexErr != nil {
			return errors.Wrap(_propertyArrayIndexErr, "Error serializing 'propertyArrayIndex' field")
		}
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

	// Simple Field (accessResult)
	if pushErr := writeBuffer.PushContext("accessResult"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for accessResult")
	}
	_accessResultErr := writeBuffer.WriteSerializable(m.GetAccessResult())
	if popErr := writeBuffer.PopContext("accessResult"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for accessResult")
	}
	if _accessResultErr != nil {
		return errors.Wrap(_accessResultErr, "Error serializing 'accessResult' field")
	}

	if popErr := writeBuffer.PopContext("BACnetPropertyAccessResult"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetPropertyAccessResult")
	}
	return nil
}


func (m *_BACnetPropertyAccessResult) isBACnetPropertyAccessResult() bool {
	return true
}

func (m *_BACnetPropertyAccessResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



