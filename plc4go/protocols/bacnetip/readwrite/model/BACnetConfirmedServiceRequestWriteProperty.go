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

// BACnetConfirmedServiceRequestWriteProperty is the corresponding interface of BACnetConfirmedServiceRequestWriteProperty
type BACnetConfirmedServiceRequestWriteProperty interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequest
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() BACnetContextTagObjectIdentifier
	// GetPropertyIdentifier returns PropertyIdentifier (property field)
	GetPropertyIdentifier() BACnetPropertyIdentifierTagged
	// GetArrayIndex returns ArrayIndex (property field)
	GetArrayIndex() BACnetContextTagUnsignedInteger
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() BACnetConstructedData
	// GetPriority returns Priority (property field)
	GetPriority() BACnetContextTagUnsignedInteger
}

// BACnetConfirmedServiceRequestWritePropertyExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestWriteProperty.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestWritePropertyExactly interface {
	BACnetConfirmedServiceRequestWriteProperty
	isBACnetConfirmedServiceRequestWriteProperty() bool
}

// _BACnetConfirmedServiceRequestWriteProperty is the data-structure of this message
type _BACnetConfirmedServiceRequestWriteProperty struct {
	*_BACnetConfirmedServiceRequest
	ObjectIdentifier   BACnetContextTagObjectIdentifier
	PropertyIdentifier BACnetPropertyIdentifierTagged
	ArrayIndex         BACnetContextTagUnsignedInteger
	PropertyValue      BACnetConstructedData
	Priority           BACnetContextTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetServiceChoice() BACnetConfirmedServiceChoice {
	return BACnetConfirmedServiceChoice_WRITE_PROPERTY
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestWriteProperty) InitializeParent(parent BACnetConfirmedServiceRequest) {
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetParent() BACnetConfirmedServiceRequest {
	return m._BACnetConfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetObjectIdentifier() BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetPropertyIdentifier() BACnetPropertyIdentifierTagged {
	return m.PropertyIdentifier
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetArrayIndex() BACnetContextTagUnsignedInteger {
	return m.ArrayIndex
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetPropertyValue() BACnetConstructedData {
	return m.PropertyValue
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetPriority() BACnetContextTagUnsignedInteger {
	return m.Priority
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestWriteProperty factory function for _BACnetConfirmedServiceRequestWriteProperty
func NewBACnetConfirmedServiceRequestWriteProperty(objectIdentifier BACnetContextTagObjectIdentifier, propertyIdentifier BACnetPropertyIdentifierTagged, arrayIndex BACnetContextTagUnsignedInteger, propertyValue BACnetConstructedData, priority BACnetContextTagUnsignedInteger, serviceRequestLength uint32) *_BACnetConfirmedServiceRequestWriteProperty {
	_result := &_BACnetConfirmedServiceRequestWriteProperty{
		ObjectIdentifier:               objectIdentifier,
		PropertyIdentifier:             propertyIdentifier,
		ArrayIndex:                     arrayIndex,
		PropertyValue:                  propertyValue,
		Priority:                       priority,
		_BACnetConfirmedServiceRequest: NewBACnetConfirmedServiceRequest(serviceRequestLength),
	}
	_result._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestWriteProperty(structType any) BACnetConfirmedServiceRequestWriteProperty {
	if casted, ok := structType.(BACnetConfirmedServiceRequestWriteProperty); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestWriteProperty); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetTypeName() string {
	return "BACnetConfirmedServiceRequestWriteProperty"
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (objectIdentifier)
	lengthInBits += m.ObjectIdentifier.GetLengthInBits(ctx)

	// Simple field (propertyIdentifier)
	lengthInBits += m.PropertyIdentifier.GetLengthInBits(ctx)

	// Optional Field (arrayIndex)
	if m.ArrayIndex != nil {
		lengthInBits += m.ArrayIndex.GetLengthInBits(ctx)
	}

	// Simple field (propertyValue)
	lengthInBits += m.PropertyValue.GetLengthInBits(ctx)

	// Optional Field (priority)
	if m.Priority != nil {
		lengthInBits += m.Priority.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConfirmedServiceRequestWritePropertyParse(ctx context.Context, theBytes []byte, serviceRequestLength uint32) (BACnetConfirmedServiceRequestWriteProperty, error) {
	return BACnetConfirmedServiceRequestWritePropertyParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), serviceRequestLength)
}

func BACnetConfirmedServiceRequestWritePropertyParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, serviceRequestLength uint32) (BACnetConfirmedServiceRequestWriteProperty, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestWriteProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestWriteProperty")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (objectIdentifier)
	if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
	}
	_objectIdentifier, _objectIdentifierErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_BACNET_OBJECT_IDENTIFIER))
	if _objectIdentifierErr != nil {
		return nil, errors.Wrap(_objectIdentifierErr, "Error parsing 'objectIdentifier' field of BACnetConfirmedServiceRequestWriteProperty")
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
		return nil, errors.Wrap(_propertyIdentifierErr, "Error parsing 'propertyIdentifier' field of BACnetConfirmedServiceRequestWriteProperty")
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
			return nil, errors.Wrap(_err, "Error parsing 'arrayIndex' field of BACnetConfirmedServiceRequestWriteProperty")
		default:
			arrayIndex = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("arrayIndex"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for arrayIndex")
			}
		}
	}

	// Simple Field (propertyValue)
	if pullErr := readBuffer.PullContext("propertyValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for propertyValue")
	}
	_propertyValue, _propertyValueErr := BACnetConstructedDataParseWithBuffer(ctx, readBuffer, uint8(uint8(3)), BACnetObjectType(objectIdentifier.GetObjectType()), BACnetPropertyIdentifier(propertyIdentifier.GetValue()), (CastBACnetTagPayloadUnsignedInteger(utils.InlineIf(bool((arrayIndex) != (nil)), func() any { return CastBACnetTagPayloadUnsignedInteger((arrayIndex).GetPayload()) }, func() any { return CastBACnetTagPayloadUnsignedInteger(nil) }))))
	if _propertyValueErr != nil {
		return nil, errors.Wrap(_propertyValueErr, "Error parsing 'propertyValue' field of BACnetConfirmedServiceRequestWriteProperty")
	}
	propertyValue := _propertyValue.(BACnetConstructedData)
	if closeErr := readBuffer.CloseContext("propertyValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for propertyValue")
	}

	// Optional Field (priority) (Can be skipped, if a given expression evaluates to false)
	var priority BACnetContextTagUnsignedInteger = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("priority"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for priority")
		}
		_val, _err := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(4), BACnetDataType_UNSIGNED_INTEGER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'priority' field of BACnetConfirmedServiceRequestWriteProperty")
		default:
			priority = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("priority"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for priority")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestWriteProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestWriteProperty")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestWriteProperty{
		_BACnetConfirmedServiceRequest: &_BACnetConfirmedServiceRequest{
			ServiceRequestLength: serviceRequestLength,
		},
		ObjectIdentifier:   objectIdentifier,
		PropertyIdentifier: propertyIdentifier,
		ArrayIndex:         arrayIndex,
		PropertyValue:      propertyValue,
		Priority:           priority,
	}
	_child._BACnetConfirmedServiceRequest._BACnetConfirmedServiceRequestChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestWriteProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestWriteProperty")
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

		// Simple Field (propertyValue)
		if pushErr := writeBuffer.PushContext("propertyValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyValue")
		}
		_propertyValueErr := writeBuffer.WriteSerializable(ctx, m.GetPropertyValue())
		if popErr := writeBuffer.PopContext("propertyValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyValue")
		}
		if _propertyValueErr != nil {
			return errors.Wrap(_propertyValueErr, "Error serializing 'propertyValue' field")
		}

		// Optional Field (priority) (Can be skipped, if the value is null)
		var priority BACnetContextTagUnsignedInteger = nil
		if m.GetPriority() != nil {
			if pushErr := writeBuffer.PushContext("priority"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for priority")
			}
			priority = m.GetPriority()
			_priorityErr := writeBuffer.WriteSerializable(ctx, priority)
			if popErr := writeBuffer.PopContext("priority"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for priority")
			}
			if _priorityErr != nil {
				return errors.Wrap(_priorityErr, "Error serializing 'priority' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestWriteProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestWriteProperty")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) isBACnetConfirmedServiceRequestWriteProperty() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestWriteProperty) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
