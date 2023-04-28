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

// BACnetCredentialAuthenticationFactor is the corresponding interface of BACnetCredentialAuthenticationFactor
type BACnetCredentialAuthenticationFactor interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetDisable returns Disable (property field)
	GetDisable() BACnetAccessAuthenticationFactorDisableTagged
	// GetAuthenticationFactor returns AuthenticationFactor (property field)
	GetAuthenticationFactor() BACnetAuthenticationFactorEnclosed
}

// BACnetCredentialAuthenticationFactorExactly can be used when we want exactly this type and not a type which fulfills BACnetCredentialAuthenticationFactor.
// This is useful for switch cases.
type BACnetCredentialAuthenticationFactorExactly interface {
	BACnetCredentialAuthenticationFactor
	isBACnetCredentialAuthenticationFactor() bool
}

// _BACnetCredentialAuthenticationFactor is the data-structure of this message
type _BACnetCredentialAuthenticationFactor struct {
	Disable              BACnetAccessAuthenticationFactorDisableTagged
	AuthenticationFactor BACnetAuthenticationFactorEnclosed
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetCredentialAuthenticationFactor) GetDisable() BACnetAccessAuthenticationFactorDisableTagged {
	return m.Disable
}

func (m *_BACnetCredentialAuthenticationFactor) GetAuthenticationFactor() BACnetAuthenticationFactorEnclosed {
	return m.AuthenticationFactor
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetCredentialAuthenticationFactor factory function for _BACnetCredentialAuthenticationFactor
func NewBACnetCredentialAuthenticationFactor(disable BACnetAccessAuthenticationFactorDisableTagged, authenticationFactor BACnetAuthenticationFactorEnclosed) *_BACnetCredentialAuthenticationFactor {
	return &_BACnetCredentialAuthenticationFactor{Disable: disable, AuthenticationFactor: authenticationFactor}
}

// Deprecated: use the interface for direct cast
func CastBACnetCredentialAuthenticationFactor(structType any) BACnetCredentialAuthenticationFactor {
	if casted, ok := structType.(BACnetCredentialAuthenticationFactor); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetCredentialAuthenticationFactor); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetCredentialAuthenticationFactor) GetTypeName() string {
	return "BACnetCredentialAuthenticationFactor"
}

func (m *_BACnetCredentialAuthenticationFactor) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (disable)
	lengthInBits += m.Disable.GetLengthInBits(ctx)

	// Simple field (authenticationFactor)
	lengthInBits += m.AuthenticationFactor.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetCredentialAuthenticationFactor) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetCredentialAuthenticationFactorParse(theBytes []byte) (BACnetCredentialAuthenticationFactor, error) {
	return BACnetCredentialAuthenticationFactorParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetCredentialAuthenticationFactorParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetCredentialAuthenticationFactor, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetCredentialAuthenticationFactor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetCredentialAuthenticationFactor")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (disable)
	if pullErr := readBuffer.PullContext("disable"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for disable")
	}
	_disable, _disableErr := BACnetAccessAuthenticationFactorDisableTaggedParseWithBuffer(ctx, readBuffer, uint8(uint8(0)), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _disableErr != nil {
		return nil, errors.Wrap(_disableErr, "Error parsing 'disable' field of BACnetCredentialAuthenticationFactor")
	}
	disable := _disable.(BACnetAccessAuthenticationFactorDisableTagged)
	if closeErr := readBuffer.CloseContext("disable"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for disable")
	}

	// Simple Field (authenticationFactor)
	if pullErr := readBuffer.PullContext("authenticationFactor"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for authenticationFactor")
	}
	_authenticationFactor, _authenticationFactorErr := BACnetAuthenticationFactorEnclosedParseWithBuffer(ctx, readBuffer, uint8(uint8(1)))
	if _authenticationFactorErr != nil {
		return nil, errors.Wrap(_authenticationFactorErr, "Error parsing 'authenticationFactor' field of BACnetCredentialAuthenticationFactor")
	}
	authenticationFactor := _authenticationFactor.(BACnetAuthenticationFactorEnclosed)
	if closeErr := readBuffer.CloseContext("authenticationFactor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for authenticationFactor")
	}

	if closeErr := readBuffer.CloseContext("BACnetCredentialAuthenticationFactor"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetCredentialAuthenticationFactor")
	}

	// Create the instance
	return &_BACnetCredentialAuthenticationFactor{
		Disable:              disable,
		AuthenticationFactor: authenticationFactor,
	}, nil
}

func (m *_BACnetCredentialAuthenticationFactor) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetCredentialAuthenticationFactor) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetCredentialAuthenticationFactor"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetCredentialAuthenticationFactor")
	}

	// Simple Field (disable)
	if pushErr := writeBuffer.PushContext("disable"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for disable")
	}
	_disableErr := writeBuffer.WriteSerializable(ctx, m.GetDisable())
	if popErr := writeBuffer.PopContext("disable"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for disable")
	}
	if _disableErr != nil {
		return errors.Wrap(_disableErr, "Error serializing 'disable' field")
	}

	// Simple Field (authenticationFactor)
	if pushErr := writeBuffer.PushContext("authenticationFactor"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for authenticationFactor")
	}
	_authenticationFactorErr := writeBuffer.WriteSerializable(ctx, m.GetAuthenticationFactor())
	if popErr := writeBuffer.PopContext("authenticationFactor"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for authenticationFactor")
	}
	if _authenticationFactorErr != nil {
		return errors.Wrap(_authenticationFactorErr, "Error serializing 'authenticationFactor' field")
	}

	if popErr := writeBuffer.PopContext("BACnetCredentialAuthenticationFactor"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetCredentialAuthenticationFactor")
	}
	return nil
}

func (m *_BACnetCredentialAuthenticationFactor) isBACnetCredentialAuthenticationFactor() bool {
	return true
}

func (m *_BACnetCredentialAuthenticationFactor) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
