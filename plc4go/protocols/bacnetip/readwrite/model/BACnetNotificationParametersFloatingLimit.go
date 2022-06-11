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

// BACnetNotificationParametersFloatingLimit is the data-structure of this message
type BACnetNotificationParametersFloatingLimit struct {
	*BACnetNotificationParameters
	InnerOpeningTag *BACnetOpeningTag
	ReferenceValue  *BACnetContextTagReal
	StatusFlags     *BACnetStatusFlagsTagged
	SetPointValue   *BACnetContextTagReal
	ErrorLimit      *BACnetContextTagReal
	InnerClosingTag *BACnetClosingTag

	// Arguments.
	TagNumber          uint8
	ObjectTypeArgument BACnetObjectType
}

// IBACnetNotificationParametersFloatingLimit is the corresponding interface of BACnetNotificationParametersFloatingLimit
type IBACnetNotificationParametersFloatingLimit interface {
	IBACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() *BACnetOpeningTag
	// GetReferenceValue returns ReferenceValue (property field)
	GetReferenceValue() *BACnetContextTagReal
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() *BACnetStatusFlagsTagged
	// GetSetPointValue returns SetPointValue (property field)
	GetSetPointValue() *BACnetContextTagReal
	// GetErrorLimit returns ErrorLimit (property field)
	GetErrorLimit() *BACnetContextTagReal
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() *BACnetClosingTag
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetNotificationParametersFloatingLimit) InitializeParent(parent *BACnetNotificationParameters, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetNotificationParameters.OpeningTag = openingTag
	m.BACnetNotificationParameters.PeekedTagHeader = peekedTagHeader
	m.BACnetNotificationParameters.ClosingTag = closingTag
}

func (m *BACnetNotificationParametersFloatingLimit) GetParent() *BACnetNotificationParameters {
	return m.BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetNotificationParametersFloatingLimit) GetInnerOpeningTag() *BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *BACnetNotificationParametersFloatingLimit) GetReferenceValue() *BACnetContextTagReal {
	return m.ReferenceValue
}

func (m *BACnetNotificationParametersFloatingLimit) GetStatusFlags() *BACnetStatusFlagsTagged {
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

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersFloatingLimit factory function for BACnetNotificationParametersFloatingLimit
func NewBACnetNotificationParametersFloatingLimit(innerOpeningTag *BACnetOpeningTag, referenceValue *BACnetContextTagReal, statusFlags *BACnetStatusFlagsTagged, setPointValue *BACnetContextTagReal, errorLimit *BACnetContextTagReal, innerClosingTag *BACnetClosingTag, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *BACnetNotificationParametersFloatingLimit {
	_result := &BACnetNotificationParametersFloatingLimit{
		InnerOpeningTag:              innerOpeningTag,
		ReferenceValue:               referenceValue,
		StatusFlags:                  statusFlags,
		SetPointValue:                setPointValue,
		ErrorLimit:                   errorLimit,
		InnerClosingTag:              innerClosingTag,
		BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetNotificationParametersFloatingLimit(structType interface{}) *BACnetNotificationParametersFloatingLimit {
	if casted, ok := structType.(BACnetNotificationParametersFloatingLimit); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersFloatingLimit); ok {
		return casted
	}
	if casted, ok := structType.(BACnetNotificationParameters); ok {
		return CastBACnetNotificationParametersFloatingLimit(casted.Child)
	}
	if casted, ok := structType.(*BACnetNotificationParameters); ok {
		return CastBACnetNotificationParametersFloatingLimit(casted.Child)
	}
	return nil
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

func BACnetNotificationParametersFloatingLimitParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, peekedTagNumber uint8) (*BACnetNotificationParametersFloatingLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersFloatingLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersFloatingLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerOpeningTag")
	}
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParse(readBuffer, uint8(peekedTagNumber))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field")
	}
	innerOpeningTag := CastBACnetOpeningTag(_innerOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerOpeningTag")
	}

	// Simple Field (referenceValue)
	if pullErr := readBuffer.PullContext("referenceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for referenceValue")
	}
	_referenceValue, _referenceValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_REAL))
	if _referenceValueErr != nil {
		return nil, errors.Wrap(_referenceValueErr, "Error parsing 'referenceValue' field")
	}
	referenceValue := CastBACnetContextTagReal(_referenceValue)
	if closeErr := readBuffer.CloseContext("referenceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for referenceValue")
	}

	// Simple Field (statusFlags)
	if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusFlags")
	}
	_statusFlags, _statusFlagsErr := BACnetStatusFlagsTaggedParse(readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _statusFlagsErr != nil {
		return nil, errors.Wrap(_statusFlagsErr, "Error parsing 'statusFlags' field")
	}
	statusFlags := CastBACnetStatusFlagsTagged(_statusFlags)
	if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusFlags")
	}

	// Simple Field (setPointValue)
	if pullErr := readBuffer.PullContext("setPointValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for setPointValue")
	}
	_setPointValue, _setPointValueErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_REAL))
	if _setPointValueErr != nil {
		return nil, errors.Wrap(_setPointValueErr, "Error parsing 'setPointValue' field")
	}
	setPointValue := CastBACnetContextTagReal(_setPointValue)
	if closeErr := readBuffer.CloseContext("setPointValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for setPointValue")
	}

	// Simple Field (errorLimit)
	if pullErr := readBuffer.PullContext("errorLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for errorLimit")
	}
	_errorLimit, _errorLimitErr := BACnetContextTagParse(readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_REAL))
	if _errorLimitErr != nil {
		return nil, errors.Wrap(_errorLimitErr, "Error parsing 'errorLimit' field")
	}
	errorLimit := CastBACnetContextTagReal(_errorLimit)
	if closeErr := readBuffer.CloseContext("errorLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for errorLimit")
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerClosingTag")
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParse(readBuffer, uint8(peekedTagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field")
	}
	innerClosingTag := CastBACnetClosingTag(_innerClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerClosingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersFloatingLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersFloatingLimit")
	}

	// Create a partially initialized instance
	_child := &BACnetNotificationParametersFloatingLimit{
		InnerOpeningTag:              CastBACnetOpeningTag(innerOpeningTag),
		ReferenceValue:               CastBACnetContextTagReal(referenceValue),
		StatusFlags:                  CastBACnetStatusFlagsTagged(statusFlags),
		SetPointValue:                CastBACnetContextTagReal(setPointValue),
		ErrorLimit:                   CastBACnetContextTagReal(errorLimit),
		InnerClosingTag:              CastBACnetClosingTag(innerClosingTag),
		BACnetNotificationParameters: &BACnetNotificationParameters{},
	}
	_child.BACnetNotificationParameters.Child = _child
	return _child, nil
}

func (m *BACnetNotificationParametersFloatingLimit) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersFloatingLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersFloatingLimit")
		}

		// Simple Field (innerOpeningTag)
		if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerOpeningTag")
		}
		_innerOpeningTagErr := m.InnerOpeningTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerOpeningTag")
		}
		if _innerOpeningTagErr != nil {
			return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
		}

		// Simple Field (referenceValue)
		if pushErr := writeBuffer.PushContext("referenceValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for referenceValue")
		}
		_referenceValueErr := m.ReferenceValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("referenceValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for referenceValue")
		}
		if _referenceValueErr != nil {
			return errors.Wrap(_referenceValueErr, "Error serializing 'referenceValue' field")
		}

		// Simple Field (statusFlags)
		if pushErr := writeBuffer.PushContext("statusFlags"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusFlags")
		}
		_statusFlagsErr := m.StatusFlags.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("statusFlags"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusFlags")
		}
		if _statusFlagsErr != nil {
			return errors.Wrap(_statusFlagsErr, "Error serializing 'statusFlags' field")
		}

		// Simple Field (setPointValue)
		if pushErr := writeBuffer.PushContext("setPointValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for setPointValue")
		}
		_setPointValueErr := m.SetPointValue.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("setPointValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for setPointValue")
		}
		if _setPointValueErr != nil {
			return errors.Wrap(_setPointValueErr, "Error serializing 'setPointValue' field")
		}

		// Simple Field (errorLimit)
		if pushErr := writeBuffer.PushContext("errorLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for errorLimit")
		}
		_errorLimitErr := m.ErrorLimit.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("errorLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for errorLimit")
		}
		if _errorLimitErr != nil {
			return errors.Wrap(_errorLimitErr, "Error serializing 'errorLimit' field")
		}

		// Simple Field (innerClosingTag)
		if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for innerClosingTag")
		}
		_innerClosingTagErr := m.InnerClosingTag.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for innerClosingTag")
		}
		if _innerClosingTagErr != nil {
			return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
		}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersFloatingLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersFloatingLimit")
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
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
