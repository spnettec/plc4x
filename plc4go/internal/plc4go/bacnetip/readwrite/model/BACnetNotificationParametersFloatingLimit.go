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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// The data-structure of this message
type BACnetNotificationParametersFloatingLimit struct {
	*BACnetNotificationParameters
	InnerOpeningTag *BACnetOpeningTag
	ReferenceValue  *BACnetContextTagReal
	StatusFlags     *BACnetStatusFlags
	SetPointValue   *BACnetContextTagReal
	ErrorLimit      *BACnetContextTagReal
	InnerClosingTag *BACnetClosingTag

	// Arguments.
	TagNumber  uint8
	ObjectType BACnetObjectType
}

// The corresponding interface
type IBACnetNotificationParametersFloatingLimit interface {
	// GetInnerOpeningTag returns InnerOpeningTag
	GetInnerOpeningTag() *BACnetOpeningTag
	// GetReferenceValue returns ReferenceValue
	GetReferenceValue() *BACnetContextTagReal
	// GetStatusFlags returns StatusFlags
	GetStatusFlags() *BACnetStatusFlags
	// GetSetPointValue returns SetPointValue
	GetSetPointValue() *BACnetContextTagReal
	// GetErrorLimit returns ErrorLimit
	GetErrorLimit() *BACnetContextTagReal
	// GetInnerClosingTag returns InnerClosingTag
	GetInnerClosingTag() *BACnetClosingTag
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *BACnetNotificationParametersFloatingLimit) PeekedTagNumber() uint8 {
	return uint8(4)
}

func (m *BACnetNotificationParametersFloatingLimit) GetPeekedTagNumber() uint8 {
	return uint8(4)
}

func (m *BACnetNotificationParametersFloatingLimit) InitializeParent(parent *BACnetNotificationParameters, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetNotificationParameters.OpeningTag = openingTag
	m.BACnetNotificationParameters.PeekedTagHeader = peekedTagHeader
	m.BACnetNotificationParameters.ClosingTag = closingTag
}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *BACnetNotificationParametersFloatingLimit) GetInnerOpeningTag() *BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *BACnetNotificationParametersFloatingLimit) GetReferenceValue() *BACnetContextTagReal {
	return m.ReferenceValue
}

func (m *BACnetNotificationParametersFloatingLimit) GetStatusFlags() *BACnetStatusFlags {
	return m.StatusFlags
}

func (m *BACnetNotificationParametersFloatingLimit) GetSetPointValue() *BACnetContextTagReal {
	return m.SetPointValue
}

func (m *BACnetNotificationParametersFloatingLimit) GetErrorLimit() *BACnetContextTagReal {
	return m.ErrorLimit
}

func (m *BACnetNotificationParametersFloatingLimit) GetInnerClosingTag() *BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersFloatingLimit factory function for BACnetNotificationParametersFloatingLimit
func NewBACnetNotificationParametersFloatingLimit(innerOpeningTag *BACnetOpeningTag, referenceValue *BACnetContextTagReal, statusFlags *BACnetStatusFlags, setPointValue *BACnetContextTagReal, errorLimit *BACnetContextTagReal, innerClosingTag *BACnetClosingTag, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, objectType BACnetObjectType) *BACnetNotificationParameters {
	child := &BACnetNotificationParametersFloatingLimit{
		InnerOpeningTag:              innerOpeningTag,
		ReferenceValue:               referenceValue,
		StatusFlags:                  statusFlags,
		SetPointValue:                setPointValue,
		ErrorLimit:                   errorLimit,
		InnerClosingTag:              innerClosingTag,
		BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectType),
	}
	child.Child = child
	return child.BACnetNotificationParameters
}

func CastBACnetNotificationParametersFloatingLimit(structType interface{}) *BACnetNotificationParametersFloatingLimit {
	castFunc := func(typ interface{}) *BACnetNotificationParametersFloatingLimit {
		if casted, ok := typ.(BACnetNotificationParametersFloatingLimit); ok {
			return &casted
		}
		if casted, ok := typ.(*BACnetNotificationParametersFloatingLimit); ok {
			return casted
		}
		if casted, ok := typ.(BACnetNotificationParameters); ok {
			return CastBACnetNotificationParametersFloatingLimit(casted.Child)
		}
		if casted, ok := typ.(*BACnetNotificationParameters); ok {
			return CastBACnetNotificationParametersFloatingLimit(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *BACnetNotificationParametersFloatingLimit) GetTypeName() string {
	return "BACnetNotificationParametersFloatingLimit"
}

func (m *BACnetNotificationParametersFloatingLimit) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetNotificationParametersFloatingLimit) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Simple field (referenceValue)
	lengthInBits += m.ReferenceValue.GetLengthInBits()

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits()

	// Simple field (setPointValue)
	lengthInBits += m.SetPointValue.GetLengthInBits()

	// Simple field (errorLimit)
	lengthInBits += m.ErrorLimit.GetLengthInBits()

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetNotificationParametersFloatingLimit) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersFloatingLimitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectType BACnetObjectType, peekedTagNumber uint8) (*BACnetNotificationParameters, error) {
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersFloatingLimit"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, pullErr
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetContextTagParse(readBuffer, uint8(peekedTagNumber), BACnetDataType(BACnetDataType_OPENING_TAG))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field")
	}
	innerOpeningTag := CastBACnetOpeningTag(_innerOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (referenceValue)
	if pullErr := readBuffer.PullContext("referenceValue"); pullErr != nil {
		return nil, pullErr
	}
	_referenceValue, _referenceValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_REAL))
	if _referenceValueErr != nil {
		return nil, errors.Wrap(_referenceValueErr, "Error parsing 'referenceValue' field")
	}
	referenceValue := CastBACnetContextTagReal(_referenceValue)
	if closeErr := readBuffer.CloseContext("referenceValue"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (statusFlags)
	if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
		return nil, pullErr
	}
	_statusFlags, _statusFlagsErr := BACnetStatusFlagsParse(readBuffer, uint8(uint8(1)))
	if _statusFlagsErr != nil {
		return nil, errors.Wrap(_statusFlagsErr, "Error parsing 'statusFlags' field")
	}
	statusFlags := CastBACnetStatusFlags(_statusFlags)
	if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (setPointValue)
	if pullErr := readBuffer.PullContext("setPointValue"); pullErr != nil {
		return nil, pullErr
	}
	_setPointValue, _setPointValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_REAL))
	if _setPointValueErr != nil {
		return nil, errors.Wrap(_setPointValueErr, "Error parsing 'setPointValue' field")
	}
	setPointValue := CastBACnetContextTagReal(_setPointValue)
	if closeErr := readBuffer.CloseContext("setPointValue"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (errorLimit)
	if pullErr := readBuffer.PullContext("errorLimit"); pullErr != nil {
		return nil, pullErr
	}
	_errorLimit, _errorLimitErr := BACnetContextTagParse(readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_REAL))
	if _errorLimitErr != nil {
		return nil, errors.Wrap(_errorLimitErr, "Error parsing 'errorLimit' field")
	}
	errorLimit := CastBACnetContextTagReal(_errorLimit)
	if closeErr := readBuffer.CloseContext("errorLimit"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, pullErr
	}
	_innerClosingTag, _innerClosingTagErr := BACnetContextTagParse(readBuffer, uint8(peekedTagNumber), BACnetDataType(BACnetDataType_CLOSING_TAG))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field")
	}
	innerClosingTag := CastBACnetClosingTag(_innerClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersFloatingLimit"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersFloatingLimit{
		InnerOpeningTag:              CastBACnetOpeningTag(innerOpeningTag),
		ReferenceValue:               CastBACnetContextTagReal(referenceValue),
		StatusFlags:                  CastBACnetStatusFlags(statusFlags),
		SetPointValue:                CastBACnetContextTagReal(setPointValue),
		ErrorLimit:                   CastBACnetContextTagReal(errorLimit),
		InnerClosingTag:              CastBACnetClosingTag(innerClosingTag),
		BACnetNotificationParameters: &BACnetNotificationParameters{},
	}
	_child.BACnetNotificationParameters.Child = _child
	return _child.BACnetNotificationParameters, nil
}

func (m *BACnetNotificationParametersFloatingLimit) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersFloatingLimit"); pushErr != nil {
			return pushErr
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return pushErr
		}
		_innerOpeningTagErr := m.InnerOpeningTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return popErr
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (referenceValue)
		if pushErr := writeBuffer.PushContext("referenceValue"); pushErr != nil {
			return pushErr
		}
		_referenceValueErr := m.ReferenceValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("referenceValue"); popErr != nil {
			return popErr
		}
		if _referenceValueErr != nil {
			return errors.Wrap(_referenceValueErr, "Error serializing 'referenceValue' field")
		}

		// Simple Field (statusFlags)
		if pushErr := writeBuffer.PushContext("statusFlags"); pushErr != nil {
			return pushErr
		}
		_statusFlagsErr := m.StatusFlags.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("statusFlags"); popErr != nil {
			return popErr
		}
		if _statusFlagsErr != nil {
			return errors.Wrap(_statusFlagsErr, "Error serializing 'statusFlags' field")
		}

		// Simple Field (setPointValue)
		if pushErr := writeBuffer.PushContext("setPointValue"); pushErr != nil {
			return pushErr
		}
		_setPointValueErr := m.SetPointValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("setPointValue"); popErr != nil {
			return popErr
		}
		if _setPointValueErr != nil {
			return errors.Wrap(_setPointValueErr, "Error serializing 'setPointValue' field")
		}

		// Simple Field (errorLimit)
		if pushErr := writeBuffer.PushContext("errorLimit"); pushErr != nil {
			return pushErr
		}
		_errorLimitErr := m.ErrorLimit.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("errorLimit"); popErr != nil {
			return popErr
		}
		if _errorLimitErr != nil {
			return errors.Wrap(_errorLimitErr, "Error serializing 'errorLimit' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return pushErr
		}
		_innerClosingTagErr := m.InnerClosingTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return popErr
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersFloatingLimit"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetNotificationParametersFloatingLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
