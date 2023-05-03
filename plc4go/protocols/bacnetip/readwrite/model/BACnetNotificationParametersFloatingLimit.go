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

// BACnetNotificationParametersFloatingLimit is the corresponding interface of BACnetNotificationParametersFloatingLimit
type BACnetNotificationParametersFloatingLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetReferenceValue returns ReferenceValue (property field)
	GetReferenceValue() BACnetContextTagReal
	// GetStatusFlags returns StatusFlags (property field)
	GetStatusFlags() BACnetStatusFlagsTagged
	// GetSetPointValue returns SetPointValue (property field)
	GetSetPointValue() BACnetContextTagReal
	// GetErrorLimit returns ErrorLimit (property field)
	GetErrorLimit() BACnetContextTagReal
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// BACnetNotificationParametersFloatingLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersFloatingLimit.
// This is useful for switch cases.
type BACnetNotificationParametersFloatingLimitExactly interface {
	BACnetNotificationParametersFloatingLimit
	isBACnetNotificationParametersFloatingLimit() bool
}

// _BACnetNotificationParametersFloatingLimit is the data-structure of this message
type _BACnetNotificationParametersFloatingLimit struct {
	*_BACnetNotificationParameters
	InnerOpeningTag BACnetOpeningTag
	ReferenceValue  BACnetContextTagReal
	StatusFlags     BACnetStatusFlagsTagged
	SetPointValue   BACnetContextTagReal
	ErrorLimit      BACnetContextTagReal
	InnerClosingTag BACnetClosingTag
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersFloatingLimit) InitializeParent(parent BACnetNotificationParameters, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersFloatingLimit) GetParent() BACnetNotificationParameters {
	return m._BACnetNotificationParameters
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersFloatingLimit) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetNotificationParametersFloatingLimit) GetReferenceValue() BACnetContextTagReal {
	return m.ReferenceValue
}

func (m *_BACnetNotificationParametersFloatingLimit) GetStatusFlags() BACnetStatusFlagsTagged {
	return m.StatusFlags
}

func (m *_BACnetNotificationParametersFloatingLimit) GetSetPointValue() BACnetContextTagReal {
	return m.SetPointValue
}

func (m *_BACnetNotificationParametersFloatingLimit) GetErrorLimit() BACnetContextTagReal {
	return m.ErrorLimit
}

func (m *_BACnetNotificationParametersFloatingLimit) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetNotificationParametersFloatingLimit factory function for _BACnetNotificationParametersFloatingLimit
func NewBACnetNotificationParametersFloatingLimit(innerOpeningTag BACnetOpeningTag, referenceValue BACnetContextTagReal, statusFlags BACnetStatusFlagsTagged, setPointValue BACnetContextTagReal, errorLimit BACnetContextTagReal, innerClosingTag BACnetClosingTag, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, objectTypeArgument BACnetObjectType) *_BACnetNotificationParametersFloatingLimit {
	_result := &_BACnetNotificationParametersFloatingLimit{
		InnerOpeningTag:               innerOpeningTag,
		ReferenceValue:                referenceValue,
		StatusFlags:                   statusFlags,
		SetPointValue:                 setPointValue,
		ErrorLimit:                    errorLimit,
		InnerClosingTag:               innerClosingTag,
		_BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
	}
	_result._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersFloatingLimit(structType any) BACnetNotificationParametersFloatingLimit {
	if casted, ok := structType.(BACnetNotificationParametersFloatingLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersFloatingLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersFloatingLimit) GetTypeName() string {
	return "BACnetNotificationParametersFloatingLimit"
}

func (m *_BACnetNotificationParametersFloatingLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits(ctx)

	// Simple field (referenceValue)
	lengthInBits += m.ReferenceValue.GetLengthInBits(ctx)

	// Simple field (statusFlags)
	lengthInBits += m.StatusFlags.GetLengthInBits(ctx)

	// Simple field (setPointValue)
	lengthInBits += m.SetPointValue.GetLengthInBits(ctx)

	// Simple field (errorLimit)
	lengthInBits += m.ErrorLimit.GetLengthInBits(ctx)

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetNotificationParametersFloatingLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetNotificationParametersFloatingLimitParse(theBytes []byte, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParametersFloatingLimit, error) {
	return BACnetNotificationParametersFloatingLimitParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), peekedTagNumber, tagNumber, objectTypeArgument)
}

func BACnetNotificationParametersFloatingLimitParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParametersFloatingLimit, error) {
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
	_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber))
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field of BACnetNotificationParametersFloatingLimit")
	}
	innerOpeningTag := _innerOpeningTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerOpeningTag")
	}

	// Simple Field (referenceValue)
	if pullErr := readBuffer.PullContext("referenceValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for referenceValue")
	}
	_referenceValue, _referenceValueErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), BACnetDataType(BACnetDataType_REAL))
	if _referenceValueErr != nil {
		return nil, errors.Wrap(_referenceValueErr, "Error parsing 'referenceValue' field of BACnetNotificationParametersFloatingLimit")
	}
	referenceValue := _referenceValue.(BACnetContextTagReal)
	if closeErr := readBuffer.CloseContext("referenceValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for referenceValue")
	}

	// Simple Field (statusFlags)
	if pullErr := readBuffer.PullContext("statusFlags"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusFlags")
	}
	_statusFlags, _statusFlagsErr := BACnetStatusFlagsTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(1)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _statusFlagsErr != nil {
		return nil, errors.Wrap(_statusFlagsErr, "Error parsing 'statusFlags' field of BACnetNotificationParametersFloatingLimit")
	}
	statusFlags := _statusFlags.(BACnetStatusFlagsTagged)
	if closeErr := readBuffer.CloseContext("statusFlags"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusFlags")
	}

	// Simple Field (setPointValue)
	if pullErr := readBuffer.PullContext("setPointValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for setPointValue")
	}
	_setPointValue, _setPointValueErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_REAL))
	if _setPointValueErr != nil {
		return nil, errors.Wrap(_setPointValueErr, "Error parsing 'setPointValue' field of BACnetNotificationParametersFloatingLimit")
	}
	setPointValue := _setPointValue.(BACnetContextTagReal)
	if closeErr := readBuffer.CloseContext("setPointValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for setPointValue")
	}

	// Simple Field (errorLimit)
	if pullErr := readBuffer.PullContext("errorLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for errorLimit")
	}
	_errorLimit, _errorLimitErr := BACnetContextTagParseWithBuffer(ctx, readBuffer, uint8(uint8(3)), BACnetDataType(BACnetDataType_REAL))
	if _errorLimitErr != nil {
		return nil, errors.Wrap(_errorLimitErr, "Error parsing 'errorLimit' field of BACnetNotificationParametersFloatingLimit")
	}
	errorLimit := _errorLimit.(BACnetContextTagReal)
	if closeErr := readBuffer.CloseContext("errorLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for errorLimit")
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerClosingTag")
	}
	_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParseWithBuffer(ctx, readBuffer, uint8(peekedTagNumber))
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field of BACnetNotificationParametersFloatingLimit")
	}
	innerClosingTag := _innerClosingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerClosingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersFloatingLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersFloatingLimit")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersFloatingLimit{
		_BACnetNotificationParameters: &_BACnetNotificationParameters{
			TagNumber:          tagNumber,
			ObjectTypeArgument: objectTypeArgument,
		},
		InnerOpeningTag: innerOpeningTag,
		ReferenceValue:  referenceValue,
		StatusFlags:     statusFlags,
		SetPointValue:   setPointValue,
		ErrorLimit:      errorLimit,
		InnerClosingTag: innerClosingTag,
	}
	_child._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersFloatingLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersFloatingLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
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
		_innerOpeningTagErr := writeBuffer.WriteSerializable(ctx, m.GetInnerOpeningTag())
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
		_referenceValueErr := writeBuffer.WriteSerializable(ctx, m.GetReferenceValue())
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
		_statusFlagsErr := writeBuffer.WriteSerializable(ctx, m.GetStatusFlags())
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
		_setPointValueErr := writeBuffer.WriteSerializable(ctx, m.GetSetPointValue())
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
		_errorLimitErr := writeBuffer.WriteSerializable(ctx, m.GetErrorLimit())
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
		_innerClosingTagErr := writeBuffer.WriteSerializable(ctx, m.GetInnerClosingTag())
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
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetNotificationParametersFloatingLimit) isBACnetNotificationParametersFloatingLimit() bool {
	return true
}

func (m *_BACnetNotificationParametersFloatingLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
