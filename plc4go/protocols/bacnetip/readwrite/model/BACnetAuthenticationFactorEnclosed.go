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

// BACnetAuthenticationFactorEnclosed is the data-structure of this message
type BACnetAuthenticationFactorEnclosed struct {
	OpeningTag           *BACnetOpeningTag
	AuthenticationFactor *BACnetAuthenticationFactor
	ClosingTag           *BACnetClosingTag

	// Arguments.
	TagNumber uint8
}

// IBACnetAuthenticationFactorEnclosed is the corresponding interface of BACnetAuthenticationFactorEnclosed
type IBACnetAuthenticationFactorEnclosed interface {
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() *BACnetOpeningTag
	// GetAuthenticationFactor returns AuthenticationFactor (property field)
	GetAuthenticationFactor() *BACnetAuthenticationFactor
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() *BACnetClosingTag
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

func (m *BACnetAuthenticationFactorEnclosed) GetOpeningTag() *BACnetOpeningTag {
	return m.OpeningTag
}

func (m *BACnetAuthenticationFactorEnclosed) GetAuthenticationFactor() *BACnetAuthenticationFactor {
	return m.AuthenticationFactor
}

func (m *BACnetAuthenticationFactorEnclosed) GetClosingTag() *BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetAuthenticationFactorEnclosed factory function for BACnetAuthenticationFactorEnclosed
func NewBACnetAuthenticationFactorEnclosed(openingTag *BACnetOpeningTag, authenticationFactor *BACnetAuthenticationFactor, closingTag *BACnetClosingTag, tagNumber uint8) *BACnetAuthenticationFactorEnclosed {
	return &BACnetAuthenticationFactorEnclosed{OpeningTag: openingTag, AuthenticationFactor: authenticationFactor, ClosingTag: closingTag, TagNumber: tagNumber}
}

func CastBACnetAuthenticationFactorEnclosed(structType interface{}) *BACnetAuthenticationFactorEnclosed {
	if casted, ok := structType.(BACnetAuthenticationFactorEnclosed); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetAuthenticationFactorEnclosed); ok {
		return casted
	}
	return nil
}

func (m *BACnetAuthenticationFactorEnclosed) GetTypeName() string {
	return "BACnetAuthenticationFactorEnclosed"
}

func (m *BACnetAuthenticationFactorEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetAuthenticationFactorEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (authenticationFactor)
	lengthInBits += m.AuthenticationFactor.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetAuthenticationFactorEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAuthenticationFactorEnclosedParse(readBuffer utils.ReadBuffer, tagNumber uint8) (*BACnetAuthenticationFactorEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthenticationFactorEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthenticationFactorEnclosed")
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

	// Simple Field (authenticationFactor)
	if pullErr := readBuffer.PullContext("authenticationFactor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for authenticationFactor")
	}
	_authenticationFactor, _authenticationFactorErr := BACnetAuthenticationFactorParse(readBuffer)
	if _authenticationFactorErr != nil {
		return nil, errors.Wrap(_authenticationFactorErr, "Error parsing 'authenticationFactor' field")
	}
	authenticationFactor := CastBACnetAuthenticationFactor(_authenticationFactor)
	if closeErr := readBuffer.CloseContext("authenticationFactor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for authenticationFactor")
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

	if closeErr := readBuffer.CloseContext("BACnetAuthenticationFactorEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthenticationFactorEnclosed")
	}

	// Create the instance
	return NewBACnetAuthenticationFactorEnclosed(openingTag, authenticationFactor, closingTag, tagNumber), nil
}

func (m *BACnetAuthenticationFactorEnclosed) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetAuthenticationFactorEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthenticationFactorEnclosed")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := m.OpeningTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (authenticationFactor)
	if pushErr := writeBuffer.PushContext("authenticationFactor"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for authenticationFactor")
	}
	_authenticationFactorErr := m.AuthenticationFactor.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("authenticationFactor"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for authenticationFactor")
	}
	if _authenticationFactorErr != nil {
		return errors.Wrap(_authenticationFactorErr, "Error serializing 'authenticationFactor' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := m.ClosingTag.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("BACnetAuthenticationFactorEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAuthenticationFactorEnclosed")
	}
	return nil
}

func (m *BACnetAuthenticationFactorEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
