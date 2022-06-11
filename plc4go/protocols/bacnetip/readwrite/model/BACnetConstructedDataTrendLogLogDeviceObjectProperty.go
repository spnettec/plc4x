/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataTrendLogLogDeviceObjectProperty is the data-structure of this message
type BACnetConstructedDataTrendLogLogDeviceObjectProperty struct {
	*BACnetConstructedData
	LogDeviceObjectProperty *BACnetDeviceObjectPropertyReference

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataTrendLogLogDeviceObjectProperty is the corresponding interface of BACnetConstructedDataTrendLogLogDeviceObjectProperty
type IBACnetConstructedDataTrendLogLogDeviceObjectProperty interface {
	IBACnetConstructedData
	// GetLogDeviceObjectProperty returns LogDeviceObjectProperty (property field)
	GetLogDeviceObjectProperty() *BACnetDeviceObjectPropertyReference
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetObjectTypeArgument() BACnetObjectType {
	return BACnetObjectType_TREND_LOG
}

func (m *BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_LOG_DEVICE_OBJECT_PROPERTY
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataTrendLogLogDeviceObjectProperty) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetLogDeviceObjectProperty() *BACnetDeviceObjectPropertyReference {
	return m.LogDeviceObjectProperty
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataTrendLogLogDeviceObjectProperty factory function for BACnetConstructedDataTrendLogLogDeviceObjectProperty
func NewBACnetConstructedDataTrendLogLogDeviceObjectProperty(logDeviceObjectProperty *BACnetDeviceObjectPropertyReference, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataTrendLogLogDeviceObjectProperty {
	_result := &BACnetConstructedDataTrendLogLogDeviceObjectProperty{
		LogDeviceObjectProperty: logDeviceObjectProperty,
		BACnetConstructedData:   NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataTrendLogLogDeviceObjectProperty(structType interface{}) *BACnetConstructedDataTrendLogLogDeviceObjectProperty {
	if casted, ok := structType.(BACnetConstructedDataTrendLogLogDeviceObjectProperty); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTrendLogLogDeviceObjectProperty); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataTrendLogLogDeviceObjectProperty(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataTrendLogLogDeviceObjectProperty(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetTypeName() string {
	return "BACnetConstructedDataTrendLogLogDeviceObjectProperty"
}

func (m *BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (logDeviceObjectProperty)
	lengthInBits += m.LogDeviceObjectProperty.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataTrendLogLogDeviceObjectProperty) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataTrendLogLogDeviceObjectPropertyParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataTrendLogLogDeviceObjectProperty, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTrendLogLogDeviceObjectProperty")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (logDeviceObjectProperty)
	if pullErr := readBuffer.PullContext("logDeviceObjectProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for logDeviceObjectProperty")
	}
	_logDeviceObjectProperty, _logDeviceObjectPropertyErr := BACnetDeviceObjectPropertyReferenceParse(readBuffer)
	if _logDeviceObjectPropertyErr != nil {
		return nil, errors.Wrap(_logDeviceObjectPropertyErr, "Error parsing 'logDeviceObjectProperty' field")
	}
	logDeviceObjectProperty := CastBACnetDeviceObjectPropertyReference(_logDeviceObjectProperty)
	if closeErr := readBuffer.CloseContext("logDeviceObjectProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for logDeviceObjectProperty")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTrendLogLogDeviceObjectProperty")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataTrendLogLogDeviceObjectProperty{
		LogDeviceObjectProperty: CastBACnetDeviceObjectPropertyReference(logDeviceObjectProperty),
		BACnetConstructedData:   &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataTrendLogLogDeviceObjectProperty) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTrendLogLogDeviceObjectProperty")
		}

		// Simple Field (logDeviceObjectProperty)
		if pushErr := writeBuffer.PushContext("logDeviceObjectProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for logDeviceObjectProperty")
		}
		_logDeviceObjectPropertyErr := m.LogDeviceObjectProperty.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("logDeviceObjectProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for logDeviceObjectProperty")
		}
		if _logDeviceObjectPropertyErr != nil {
			return errors.Wrap(_logDeviceObjectPropertyErr, "Error serializing 'logDeviceObjectProperty' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTrendLogLogDeviceObjectProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTrendLogLogDeviceObjectProperty")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataTrendLogLogDeviceObjectProperty) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
