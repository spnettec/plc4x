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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetNotificationParametersBufferReady is the corresponding interface of BACnetNotificationParametersBufferReady
type BACnetNotificationParametersBufferReady interface {
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetBufferProperty returns BufferProperty (property field)
	GetBufferProperty() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetPreviousNotification returns PreviousNotification (property field)
	GetPreviousNotification() BACnetContextTagUnsignedInteger
	// GetCurrentNotification returns CurrentNotification (property field)
	GetCurrentNotification() BACnetContextTagUnsignedInteger
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// BACnetNotificationParametersBufferReadyExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersBufferReady.
// This is useful for switch cases.
type BACnetNotificationParametersBufferReadyExactly interface {
	BACnetNotificationParametersBufferReady
	isBACnetNotificationParametersBufferReady() bool
}

// _BACnetNotificationParametersBufferReady is the data-structure of this message
type _BACnetNotificationParametersBufferReady struct {
	*_BACnetNotificationParameters
	InnerOpeningTag      BACnetOpeningTag
	BufferProperty       BACnetDeviceObjectPropertyReferenceEnclosed
	PreviousNotification BACnetContextTagUnsignedInteger
	CurrentNotification  BACnetContextTagUnsignedInteger
	InnerClosingTag      BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersBufferReady) InitializeParent(parent BACnetNotificationParameters, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersBufferReady) GetParent() BACnetNotificationParameters {
	return m._BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersBufferReady) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersBufferReady) GetBufferProperty() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.BufferProperty
}

func (m *_BACnetNotificationParametersBufferReady) GetPreviousNotification() BACnetContextTagUnsignedInteger {
	return m.PreviousNotification
}

func (m *_BACnetNotificationParametersBufferReady) GetCurrentNotification() BACnetContextTagUnsignedInteger {
	return m.CurrentNotification
}

func (m *_BACnetNotificationParametersBufferReady) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersBufferReady factory function for _BACnetNotificationParametersBufferReady
func NewBACnetNotificationParametersBufferReady(innerOpeningTag BACnetOpeningTag, bufferProperty BACnetDeviceObjectPropertyReferenceEnclosed, previousNotification BACnetContextTagUnsignedInteger, currentNotification BACnetContextTagUnsignedInteger, innerClosingTag BACnetClosingTag, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersBufferReady {
	_result := &_BACnetNotificationParametersBufferReady{
		InnerOpeningTag:               innerOpeningTag,
		BufferProperty:                bufferProperty,
		PreviousNotification:          previousNotification,
		CurrentNotification:           currentNotification,
		InnerClosingTag:               innerClosingTag,
		_BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
	}
	_result._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersBufferReady(structType interface{}) BACnetNotificationParametersBufferReady {
	if casted, ok := structType.(BACnetNotificationParametersBufferReady); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersBufferReady); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersBufferReady) GetTypeName() string {
	return "BACnetNotificationParametersBufferReady"
}

func (m *_BACnetNotificationParametersBufferReady) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetNotificationParametersBufferReady) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Simple field (bufferProperty)
	lengthInBits += m.BufferProperty.GetLengthInBits()

	// Simple field (previousNotification)
	lengthInBits += m.PreviousNotification.GetLengthInBits()

	// Simple field (currentNotification)
	lengthInBits += m.CurrentNotification.GetLengthInBits()

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetNotificationParametersBufferReady) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersBufferReadyParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, peekedTagNumber uint8) (BACnetNotificationParametersBufferReady, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersBufferReady"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersBufferReady")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerOpeningTag")
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParse(readBuffer, uint8(peekedTagNumber))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field of BACnetNotificationParametersBufferReady")
	}
	innerOpeningTag := _innerOpeningTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerOpeningTag")
	}

	// Simple Field (bufferProperty)
	if pullErr := readBuffer.PullContext("bufferProperty"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bufferProperty")
	}
	_bufferProperty, _bufferPropertyErr := BACnetDeviceObjectPropertyReferenceEnclosedParse(readBuffer, uint8(uint8(0)))
	if _bufferPropertyErr != nil {
		return nil, errors.Wrap(_bufferPropertyErr, "Error parsing 'bufferProperty' field of BACnetNotificationParametersBufferReady")
	}
	bufferProperty := _bufferProperty.(BACnetDeviceObjectPropertyReferenceEnclosed)
	if closeErr := readBuffer.CloseContext("bufferProperty"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bufferProperty")
	}

	// Simple Field (previousNotification)
	if pullErr := readBuffer.PullContext("previousNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for previousNotification")
	}
	_previousNotification, _previousNotificationErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _previousNotificationErr != nil {
		return nil, errors.Wrap(_previousNotificationErr, "Error parsing 'previousNotification' field of BACnetNotificationParametersBufferReady")
	}
	previousNotification := _previousNotification.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("previousNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for previousNotification")
	}

	// Simple Field (currentNotification)
	if pullErr := readBuffer.PullContext("currentNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for currentNotification")
	}
	_currentNotification, _currentNotificationErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _currentNotificationErr != nil {
		return nil, errors.Wrap(_currentNotificationErr, "Error parsing 'currentNotification' field of BACnetNotificationParametersBufferReady")
	}
	currentNotification := _currentNotification.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("currentNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for currentNotification")
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerClosingTag")
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParse(readBuffer, uint8(peekedTagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field of BACnetNotificationParametersBufferReady")
	}
	innerClosingTag := _innerClosingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerClosingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersBufferReady"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersBufferReady")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersBufferReady{
		_BACnetNotificationParameters: &_BACnetNotificationParameters{
			TagNumber:          tagNumber,
			ObjectTypeArgument: objectTypeArgument,
		},
		InnerOpeningTag:      innerOpeningTag,
		BufferProperty:       bufferProperty,
		PreviousNotification: previousNotification,
		CurrentNotification:  currentNotification,
		InnerClosingTag:      innerClosingTag,
	}
	_child._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersBufferReady) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersBufferReady"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersBufferReady")
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerOpeningTag")
		}
		_innerOpeningTagErr := writeBuffer.WriteSerializable(m.GetInnerOpeningTag())
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerOpeningTag")
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (bufferProperty)
		if pushErr := writeBuffer.PushContext("bufferProperty"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for bufferProperty")
		}
		_bufferPropertyErr := writeBuffer.WriteSerializable(m.GetBufferProperty())
		if popErr := writeBuffer.PopContext("bufferProperty"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for bufferProperty")
		}
		if _bufferPropertyErr != nil {
			return errors.Wrap(_bufferPropertyErr, "Error serializing 'bufferProperty' field")
		}

		// Simple Field (previousNotification)
		if pushErr := writeBuffer.PushContext("previousNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for previousNotification")
		}
		_previousNotificationErr := writeBuffer.WriteSerializable(m.GetPreviousNotification())
		if popErr := writeBuffer.PopContext("previousNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for previousNotification")
		}
		if _previousNotificationErr != nil {
			return errors.Wrap(_previousNotificationErr, "Error serializing 'previousNotification' field")
		}

		// Simple Field (currentNotification)
		if pushErr := writeBuffer.PushContext("currentNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for currentNotification")
		}
		_currentNotificationErr := writeBuffer.WriteSerializable(m.GetCurrentNotification())
		if popErr := writeBuffer.PopContext("currentNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for currentNotification")
		}
		if _currentNotificationErr != nil {
			return errors.Wrap(_currentNotificationErr, "Error serializing 'currentNotification' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerClosingTag")
		}
		_innerClosingTagErr := writeBuffer.WriteSerializable(m.GetInnerClosingTag())
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerClosingTag")
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersBufferReady"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersBufferReady")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersBufferReady) isBACnetNotificationParametersBufferReady() bool {
	return true
}

func (m *_BACnetNotificationParametersBufferReady) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
