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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetPropertyStatesAuthenticationStatus is the corresponding interface of BACnetPropertyStatesAuthenticationStatus
type BACnetPropertyStatesAuthenticationStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetAuthenticationStatus returns AuthenticationStatus (property field)
	GetAuthenticationStatus() BACnetAuthenticationStatusTagged
	// IsBACnetPropertyStatesAuthenticationStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetPropertyStatesAuthenticationStatus()
}

// _BACnetPropertyStatesAuthenticationStatus is the data-structure of this message
type _BACnetPropertyStatesAuthenticationStatus struct {
	BACnetPropertyStatesContract
	AuthenticationStatus BACnetAuthenticationStatusTagged
}

var _ BACnetPropertyStatesAuthenticationStatus = (*_BACnetPropertyStatesAuthenticationStatus)(nil)
var _ BACnetPropertyStatesRequirements = (*_BACnetPropertyStatesAuthenticationStatus)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesAuthenticationStatus) GetParent() BACnetPropertyStatesContract {
	return m.BACnetPropertyStatesContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesAuthenticationStatus) GetAuthenticationStatus() BACnetAuthenticationStatusTagged {
	return m.AuthenticationStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesAuthenticationStatus factory function for _BACnetPropertyStatesAuthenticationStatus
func NewBACnetPropertyStatesAuthenticationStatus(peekedTagHeader BACnetTagHeader, authenticationStatus BACnetAuthenticationStatusTagged) *_BACnetPropertyStatesAuthenticationStatus {
	if authenticationStatus == nil {
		panic("authenticationStatus of type BACnetAuthenticationStatusTagged for BACnetPropertyStatesAuthenticationStatus must not be nil")
	}
	_result := &_BACnetPropertyStatesAuthenticationStatus{
		BACnetPropertyStatesContract: NewBACnetPropertyStates(peekedTagHeader),
		AuthenticationStatus:         authenticationStatus,
	}
	_result.BACnetPropertyStatesContract.(*_BACnetPropertyStates)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesAuthenticationStatus(structType any) BACnetPropertyStatesAuthenticationStatus {
	if casted, ok := structType.(BACnetPropertyStatesAuthenticationStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesAuthenticationStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesAuthenticationStatus) GetTypeName() string {
	return "BACnetPropertyStatesAuthenticationStatus"
}

func (m *_BACnetPropertyStatesAuthenticationStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).getLengthInBits(ctx))

	// Simple field (authenticationStatus)
	lengthInBits += m.AuthenticationStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetPropertyStatesAuthenticationStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetPropertyStatesAuthenticationStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetPropertyStates, peekedTagNumber uint8) (__bACnetPropertyStatesAuthenticationStatus BACnetPropertyStatesAuthenticationStatus, err error) {
	m.BACnetPropertyStatesContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesAuthenticationStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesAuthenticationStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	authenticationStatus, err := ReadSimpleField[BACnetAuthenticationStatusTagged](ctx, "authenticationStatus", ReadComplex[BACnetAuthenticationStatusTagged](BACnetAuthenticationStatusTaggedParseWithBufferProducer((uint8)(peekedTagNumber), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'authenticationStatus' field"))
	}
	m.AuthenticationStatus = authenticationStatus

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesAuthenticationStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesAuthenticationStatus")
	}

	return m, nil
}

func (m *_BACnetPropertyStatesAuthenticationStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesAuthenticationStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesAuthenticationStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesAuthenticationStatus")
		}

		if err := WriteSimpleField[BACnetAuthenticationStatusTagged](ctx, "authenticationStatus", m.GetAuthenticationStatus(), WriteComplex[BACnetAuthenticationStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'authenticationStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesAuthenticationStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesAuthenticationStatus")
		}
		return nil
	}
	return m.BACnetPropertyStatesContract.(*_BACnetPropertyStates).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesAuthenticationStatus) IsBACnetPropertyStatesAuthenticationStatus() {}

func (m *_BACnetPropertyStatesAuthenticationStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
