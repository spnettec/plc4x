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


// ApduDataExtGroupPropertyValueResponse is the corresponding interface of ApduDataExtGroupPropertyValueResponse
type ApduDataExtGroupPropertyValueResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ApduDataExt
}

// ApduDataExtGroupPropertyValueResponseExactly can be used when we want exactly this type and not a type which fulfills ApduDataExtGroupPropertyValueResponse.
// This is useful for switch cases.
type ApduDataExtGroupPropertyValueResponseExactly interface {
	ApduDataExtGroupPropertyValueResponse
	isApduDataExtGroupPropertyValueResponse() bool
}

// _ApduDataExtGroupPropertyValueResponse is the data-structure of this message
type _ApduDataExtGroupPropertyValueResponse struct {
	*_ApduDataExt
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ApduDataExtGroupPropertyValueResponse)  GetExtApciType() uint8 {
return 0x29}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ApduDataExtGroupPropertyValueResponse) InitializeParent(parent ApduDataExt ) {}

func (m *_ApduDataExtGroupPropertyValueResponse)  GetParent() ApduDataExt {
	return m._ApduDataExt
}


// NewApduDataExtGroupPropertyValueResponse factory function for _ApduDataExtGroupPropertyValueResponse
func NewApduDataExtGroupPropertyValueResponse( length uint8 ) *_ApduDataExtGroupPropertyValueResponse {
	_result := &_ApduDataExtGroupPropertyValueResponse{
    	_ApduDataExt: NewApduDataExt(length),
	}
	_result._ApduDataExt._ApduDataExtChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastApduDataExtGroupPropertyValueResponse(structType any) ApduDataExtGroupPropertyValueResponse {
    if casted, ok := structType.(ApduDataExtGroupPropertyValueResponse); ok {
		return casted
	}
	if casted, ok := structType.(*ApduDataExtGroupPropertyValueResponse); ok {
		return *casted
	}
	return nil
}

func (m *_ApduDataExtGroupPropertyValueResponse) GetTypeName() string {
	return "ApduDataExtGroupPropertyValueResponse"
}

func (m *_ApduDataExtGroupPropertyValueResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}


func (m *_ApduDataExtGroupPropertyValueResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ApduDataExtGroupPropertyValueResponseParse(ctx context.Context, theBytes []byte, length uint8) (ApduDataExtGroupPropertyValueResponse, error) {
	return ApduDataExtGroupPropertyValueResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), length)
}

func ApduDataExtGroupPropertyValueResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, length uint8) (ApduDataExtGroupPropertyValueResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ApduDataExtGroupPropertyValueResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ApduDataExtGroupPropertyValueResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("ApduDataExtGroupPropertyValueResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ApduDataExtGroupPropertyValueResponse")
	}

	// Create a partially initialized instance
	_child := &_ApduDataExtGroupPropertyValueResponse{
		_ApduDataExt: &_ApduDataExt{
			Length: length,
		},
	}
	_child._ApduDataExt._ApduDataExtChildRequirements = _child
	return _child, nil
}

func (m *_ApduDataExtGroupPropertyValueResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ApduDataExtGroupPropertyValueResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ApduDataExtGroupPropertyValueResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ApduDataExtGroupPropertyValueResponse")
		}

		if popErr := writeBuffer.PopContext("ApduDataExtGroupPropertyValueResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ApduDataExtGroupPropertyValueResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_ApduDataExtGroupPropertyValueResponse) isApduDataExtGroupPropertyValueResponse() bool {
	return true
}

func (m *_ApduDataExtGroupPropertyValueResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



