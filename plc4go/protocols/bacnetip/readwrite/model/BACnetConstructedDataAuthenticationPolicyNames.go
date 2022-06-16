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

// BACnetConstructedDataAuthenticationPolicyNames is the corresponding interface of BACnetConstructedDataAuthenticationPolicyNames
type BACnetConstructedDataAuthenticationPolicyNames interface {
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetAuthenticationPolicyNames returns AuthenticationPolicyNames (property field)
	GetAuthenticationPolicyNames() []BACnetApplicationTagCharacterString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetConstructedDataAuthenticationPolicyNames is the data-structure of this message
type _BACnetConstructedDataAuthenticationPolicyNames struct {
	*_BACnetConstructedData
	NumberOfDataElements      BACnetApplicationTagUnsignedInteger
	AuthenticationPolicyNames []BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_AUTHENTICATION_POLICY_NAMES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyNames) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetAuthenticationPolicyNames() []BACnetApplicationTagCharacterString {
	return m.AuthenticationPolicyNames
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetZero() uint64 {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAuthenticationPolicyNames factory function for _BACnetConstructedDataAuthenticationPolicyNames
func NewBACnetConstructedDataAuthenticationPolicyNames(numberOfDataElements BACnetApplicationTagUnsignedInteger, authenticationPolicyNames []BACnetApplicationTagCharacterString, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAuthenticationPolicyNames {
	_result := &_BACnetConstructedDataAuthenticationPolicyNames{
		NumberOfDataElements:      numberOfDataElements,
		AuthenticationPolicyNames: authenticationPolicyNames,
		_BACnetConstructedData:    NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAuthenticationPolicyNames(structType interface{}) BACnetConstructedDataAuthenticationPolicyNames {
	if casted, ok := structType.(BACnetConstructedDataAuthenticationPolicyNames); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAuthenticationPolicyNames); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetTypeName() string {
	return "BACnetConstructedDataAuthenticationPolicyNames"
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits()
	}

	// Array field
	if len(m.AuthenticationPolicyNames) > 0 {
		for _, element := range m.AuthenticationPolicyNames {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAuthenticationPolicyNamesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAuthenticationPolicyNames, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAuthenticationPolicyNames"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAuthenticationPolicyNames")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
		_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (authenticationPolicyNames)
	if pullErr := readBuffer.PullContext("authenticationPolicyNames", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for authenticationPolicyNames")
	}
	// Terminated array
	authenticationPolicyNames := make([]BACnetApplicationTagCharacterString, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetApplicationTagParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'authenticationPolicyNames' field")
			}
			authenticationPolicyNames = append(authenticationPolicyNames, _item.(BACnetApplicationTagCharacterString))

		}
	}
	if closeErr := readBuffer.CloseContext("authenticationPolicyNames", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for authenticationPolicyNames")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAuthenticationPolicyNames"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAuthenticationPolicyNames")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAuthenticationPolicyNames{
		NumberOfDataElements:      numberOfDataElements,
		AuthenticationPolicyNames: authenticationPolicyNames,
		_BACnetConstructedData:    &_BACnetConstructedData{},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAuthenticationPolicyNames"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAuthenticationPolicyNames")
		}
		// Virtual field
		if _zeroErr := writeBuffer.WriteVirtual("zero", m.GetZero()); _zeroErr != nil {
			return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
		}

		// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
		var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
		if m.GetNumberOfDataElements() != nil {
			if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
			}
			numberOfDataElements = m.GetNumberOfDataElements()
			_numberOfDataElementsErr := writeBuffer.WriteSerializable(numberOfDataElements)
			if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
				return errors.Wrap(popErr, "Error popping for numberOfDataElements")
			}
			if _numberOfDataElementsErr != nil {
				return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
			}
		}

		// Array Field (authenticationPolicyNames)
		if m.GetAuthenticationPolicyNames() != nil {
			if pushErr := writeBuffer.PushContext("authenticationPolicyNames", utils.WithRenderAsList(true)); pushErr != nil {
				return errors.Wrap(pushErr, "Error pushing for authenticationPolicyNames")
			}
			for _, _element := range m.GetAuthenticationPolicyNames() {
				_elementErr := writeBuffer.WriteSerializable(_element)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'authenticationPolicyNames' field")
				}
			}
			if popErr := writeBuffer.PopContext("authenticationPolicyNames", utils.WithRenderAsList(true)); popErr != nil {
				return errors.Wrap(popErr, "Error popping for authenticationPolicyNames")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAuthenticationPolicyNames"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAuthenticationPolicyNames")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAuthenticationPolicyNames) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
