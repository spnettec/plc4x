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
	"github.com/rs/zerolog"
)

	// Code generated by code-generation. DO NOT EDIT.


// IntegerId is the corresponding interface of IntegerId
type IntegerId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// IntegerIdExactly can be used when we want exactly this type and not a type which fulfills IntegerId.
// This is useful for switch cases.
type IntegerIdExactly interface {
	IntegerId
	isIntegerId() bool
}

// _IntegerId is the data-structure of this message
type _IntegerId struct {
}




// NewIntegerId factory function for _IntegerId
func NewIntegerId( ) *_IntegerId {
return &_IntegerId{ }
}

// Deprecated: use the interface for direct cast
func CastIntegerId(structType any) IntegerId {
    if casted, ok := structType.(IntegerId); ok {
		return casted
	}
	if casted, ok := structType.(*IntegerId); ok {
		return *casted
	}
	return nil
}

func (m *_IntegerId) GetTypeName() string {
	return "IntegerId"
}

func (m *_IntegerId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}


func (m *_IntegerId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IntegerIdParse(ctx context.Context, theBytes []byte) (IntegerId, error) {
	return IntegerIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func IntegerIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (IntegerId, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("IntegerId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IntegerId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("IntegerId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IntegerId")
	}

	// Create the instance
	return &_IntegerId{
		}, nil
}

func (m *_IntegerId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IntegerId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("IntegerId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for IntegerId")
	}

	if popErr := writeBuffer.PopContext("IntegerId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for IntegerId")
	}
	return nil
}


func (m *_IntegerId) isIntegerId() bool {
	return true
}

func (m *_IntegerId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



