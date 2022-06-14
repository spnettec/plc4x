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

// BACnetConfirmedServiceRequestCreateObjectObjectSpecifier is the data-structure of this message
type BACnetConfirmedServiceRequestCreateObjectObjectSpecifier struct {
	OpeningTag       *BACnetOpeningTag
	RawObjectType    *BACnetContextTagEnumerated
	ObjectIdentifier *BACnetContextTagObjectIdentifier
	ClosingTag       *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetConfirmedServiceRequestCreateObjectObjectSpecifier is the corresponding interface of BACnetConfirmedServiceRequestCreateObjectObjectSpecifier
type IBACnetConfirmedServiceRequestCreateObjectObjectSpecifier interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetRawObjectType returns RawObjectType (property field)
	GetRawObjectType() *BACnetContextTagEnumerated
	// GetObjectIdentifier returns ObjectIdentifier (property field)
	GetObjectIdentifier() *BACnetContextTagObjectIdentifier
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
	// GetIsObjectType returns IsObjectType (virtual field)
	GetIsObjectType() bool
	// GetObjectType returns ObjectType (virtual field)
	GetObjectType() BACnetObjectType
	// GetIsObjectIdentifier returns IsObjectIdentifier (virtual field)
	GetIsObjectIdentifier() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetRawObjectType() *BACnetContextTagEnumerated {
	return m.RawObjectType
}

func (m *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetObjectIdentifier() *BACnetContextTagObjectIdentifier {
	return m.ObjectIdentifier
}

func (m *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetIsObjectType() bool {
	rawObjectType := m.RawObjectType
	_ = rawObjectType
	objectIdentifier := m.ObjectIdentifier
	_ = objectIdentifier
	return bool(bool((m.GetRawObjectType()) != (nil)))
}

func (m *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetObjectType() BACnetObjectType {
	rawObjectType := m.RawObjectType
	_ = rawObjectType
	objectIdentifier := m.ObjectIdentifier
	_ = objectIdentifier
	return CastBACnetObjectType(MapBACnetObjectType((*m.GetRawObjectType())))
}

func (m *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetIsObjectIdentifier() bool {
	rawObjectType := m.RawObjectType
	_ = rawObjectType
	objectIdentifier := m.ObjectIdentifier
	_ = objectIdentifier
	return bool(bool((m.GetObjectIdentifier()) != (nil)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestCreateObjectObjectSpecifier factory function for BACnetConfirmedServiceRequestCreateObjectObjectSpecifier
func NewBACnetConfirmedServiceRequestCreateObjectObjectSpecifier(openingTag *BACnetOpeningTag, rawObjectType *BACnetContextTagEnumerated, objectIdentifier *BACnetContextTagObjectIdentifier, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier {
	return &BACnetConfirmedServiceRequestCreateObjectObjectSpecifier{OpeningTag: openingTag, RawObjectType: rawObjectType, ObjectIdentifier: objectIdentifier, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetConfirmedServiceRequestCreateObjectObjectSpecifier(structType interface{}) *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier {
	if casted, ok := structType.(BACnetConfirmedServiceRequestCreateObjectObjectSpecifier); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestCreateObjectObjectSpecifier); ok {
		return casted
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetTypeName() string {
	return "BACnetConfirmedServiceRequestCreateObjectObjectSpecifier"
}

func (m *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Optional Field (rawObjectType)
	if m.RawObjectType != nil {
		lengthInBits += (*m.RawObjectType).GetLengthInBits()
	}

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Optional Field (objectIdentifier)
	if m.ObjectIdentifier != nil {
		lengthInBits += (*m.ObjectIdentifier).GetLengthInBits()
	}

	// A virtual field doesn't have any in- or output.

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestCreateObjectObjectSpecifierParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetConfirmedServiceRequestCreateObjectObjectSpecifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestCreateObjectObjectSpecifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestCreateObjectObjectSpecifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
	_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer, uint8(tagNumber))
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field")
	}
	openingTag := CastBACnetOpeningTag(_openingTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Optional Field (rawObjectType) (Can be skipped, if a given expression evaluates to false)
	var rawObjectType *BACnetContextTagEnumerated = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("rawObjectType"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for rawObjectType")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(0), BACnetDataType_ENUMERATED)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'rawObjectType' field")
		default:
			rawObjectType = CastBACnetContextTagEnumerated(_val)
			if closeErr := readBuffer.CloseContext("rawObjectType"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for rawObjectType")
			}
		}
	}

	// Virtual field
	_isObjectType := bool((rawObjectType) != (nil))
	isObjectType := bool(_isObjectType)
	_ = isObjectType

	// Virtual field
	_objectType := MapBACnetObjectType((*rawObjectType))
	objectType := CastBACnetObjectType(_objectType)
	_ = objectType

	// Optional Field (objectIdentifier) (Can be skipped, if a given expression evaluates to false)
	var objectIdentifier *BACnetContextTagObjectIdentifier = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("objectIdentifier"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for objectIdentifier")
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_BACNET_OBJECT_IDENTIFIER)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'objectIdentifier' field")
		default:
			objectIdentifier = CastBACnetContextTagObjectIdentifier(_val)
			if closeErr := readBuffer.CloseContext("objectIdentifier"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for objectIdentifier")
			}
		}
	}

	// Virtual field
	_isObjectIdentifier := bool((objectIdentifier) != (nil))
	isObjectIdentifier := bool(_isObjectIdentifier)
	_ = isObjectIdentifier

	// Validation
	if !(bool(isObjectType) || bool(isObjectIdentifier)) {
		return nil, errors.WithStack(utils.ParseValidationError{"either we need a objectType or a objectIdentifier"})
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
	_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer, uint8(tagNumber))
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field")
	}
	closingTag := CastBACnetClosingTag(_closingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestCreateObjectObjectSpecifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestCreateObjectObjectSpecifier")
	}

	// Create the instance
	return NewBACnetConfirmedServiceRequestCreateObjectObjectSpecifier(openingTag, rawObjectType, objectIdentifier, closingTag, tagNumber), nil
}

func (m *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestCreateObjectObjectSpecifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestCreateObjectObjectSpecifier")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.OpeningTag)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Optional Field (rawObjectType) (Can be skipped, if the value is null)
	var rawObjectType *BACnetContextTagEnumerated = nil
	if m.RawObjectType != nil {
		if pushErr := writeBuffer.PushContext("rawObjectType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for rawObjectType")
		}
		rawObjectType = m.RawObjectType
		_rawObjectTypeErr := writeBuffer.WriteSerializable(rawObjectType)
		if popErr := writeBuffer.PopContext("rawObjectType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for rawObjectType")
		}
		if _rawObjectTypeErr != nil {
			return errors.Wrap(_rawObjectTypeErr, "Error serializing 'rawObjectType' field")
		}
	}
	// Virtual field
	if _isObjectTypeErr := writeBuffer.WriteVirtual("isObjectType", m.GetIsObjectType()); _isObjectTypeErr != nil {
		return errors.Wrap(_isObjectTypeErr, "Error serializing 'isObjectType' field")
	}
	// Virtual field
	if _objectTypeErr := writeBuffer.WriteVirtual("objectType", m.GetObjectType()); _objectTypeErr != nil {
		return errors.Wrap(_objectTypeErr, "Error serializing 'objectType' field")
	}

	// Optional Field (objectIdentifier) (Can be skipped, if the value is null)
	var objectIdentifier *BACnetContextTagObjectIdentifier = nil
	if m.ObjectIdentifier != nil {
		if pushErr := writeBuffer.PushContext("objectIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for objectIdentifier")
		}
		objectIdentifier = m.ObjectIdentifier
		_objectIdentifierErr := writeBuffer.WriteSerializable(objectIdentifier)
		if popErr := writeBuffer.PopContext("objectIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for objectIdentifier")
		}
		if _objectIdentifierErr != nil {
			return errors.Wrap(_objectIdentifierErr, "Error serializing 'objectIdentifier' field")
		}
	}
	// Virtual field
	if _isObjectIdentifierErr := writeBuffer.WriteVirtual("isObjectIdentifier", m.GetIsObjectIdentifier()); _isObjectIdentifierErr != nil {
		return errors.Wrap(_isObjectIdentifierErr, "Error serializing 'isObjectIdentifier' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.ClosingTag)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestCreateObjectObjectSpecifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestCreateObjectObjectSpecifier")
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestCreateObjectObjectSpecifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
