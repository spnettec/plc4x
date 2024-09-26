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

// SecurityDataOff is the corresponding interface of SecurityDataOff
type SecurityDataOff interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetData returns Data (property field)
	GetData() []byte
	// IsSecurityDataOff is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsSecurityDataOff()
}

// _SecurityDataOff is the data-structure of this message
type _SecurityDataOff struct {
	SecurityDataContract
	Data []byte
}

var _ SecurityDataOff = (*_SecurityDataOff)(nil)
var _ SecurityDataRequirements = (*_SecurityDataOff)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataOff) GetParent() SecurityDataContract {
	return m.SecurityDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataOff) GetData() []byte {
	return m.Data
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSecurityDataOff factory function for _SecurityDataOff
func NewSecurityDataOff(commandTypeContainer SecurityCommandTypeContainer, argument byte, data []byte) *_SecurityDataOff {
	_result := &_SecurityDataOff{
		SecurityDataContract: NewSecurityData(commandTypeContainer, argument),
		Data:                 data,
	}
	_result.SecurityDataContract.(*_SecurityData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataOff(structType any) SecurityDataOff {
	if casted, ok := structType.(SecurityDataOff); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataOff); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataOff) GetTypeName() string {
	return "SecurityDataOff"
}

func (m *_SecurityDataOff) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.SecurityDataContract.(*_SecurityData).getLengthInBits(ctx))

	// Array field
	if len(m.Data) > 0 {
		lengthInBits += 8 * uint16(len(m.Data))
	}

	return lengthInBits
}

func (m *_SecurityDataOff) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_SecurityDataOff) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_SecurityData, commandTypeContainer SecurityCommandTypeContainer) (__securityDataOff SecurityDataOff, err error) {
	m.SecurityDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataOff"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataOff")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	data, err := readBuffer.ReadByteArray("data", int(int32(commandTypeContainer.NumBytes())-int32(int32(1))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'data' field"))
	}
	m.Data = data

	if closeErr := readBuffer.CloseContext("SecurityDataOff"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataOff")
	}

	return m, nil
}

func (m *_SecurityDataOff) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SecurityDataOff) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataOff"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataOff")
		}

		if err := WriteByteArrayField(ctx, "data", m.GetData(), WriteByteArray(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'data' field")
		}

		if popErr := writeBuffer.PopContext("SecurityDataOff"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataOff")
		}
		return nil
	}
	return m.SecurityDataContract.(*_SecurityData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SecurityDataOff) IsSecurityDataOff() {}

func (m *_SecurityDataOff) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
