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

	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// Time is the corresponding interface of Time
type Time interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

// TimeExactly can be used when we want exactly this type and not a type which fulfills Time.
// This is useful for switch cases.
type TimeExactly interface {
	Time
	isTime() bool
}

// _Time is the data-structure of this message
type _Time struct {
}

// NewTime factory function for _Time
func NewTime() *_Time {
	return &_Time{}
}

// Deprecated: use the interface for direct cast
func CastTime(structType any) Time {
	if casted, ok := structType.(Time); ok {
		return casted
	}
	if casted, ok := structType.(*Time); ok {
		return *casted
	}
	return nil
}

func (m *_Time) GetTypeName() string {
	return "Time"
}

func (m *_Time) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	return lengthInBits
}

func (m *_Time) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TimeParse(ctx context.Context, theBytes []byte) (Time, error) {
	return TimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func TimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (Time, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("Time"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for Time")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("Time"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for Time")
	}

	// Create the instance
	return &_Time{}, nil
}

func (m *_Time) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_Time) SerializeWithWriteBuffer(ctx context.Context, wb utils.WriteBuffer) error {
	positionAware := wb
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr := wb.PushContext("Time"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for Time")
	}

	if popErr := wb.PopContext("Time"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for Time")
	}
	return nil
}

func (m *_Time) isTime() bool {
	return true
}

func (m *_Time) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(utils.WithWriteBufferBoxBasedOmitEmptyBoxes(), utils.WithWriteBufferBoxBasedMergeSingleBoxes())
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
