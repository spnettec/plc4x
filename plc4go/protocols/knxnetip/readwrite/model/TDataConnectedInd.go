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


// TDataConnectedInd is the corresponding interface of TDataConnectedInd
type TDataConnectedInd interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMI
}

// TDataConnectedIndExactly can be used when we want exactly this type and not a type which fulfills TDataConnectedInd.
// This is useful for switch cases.
type TDataConnectedIndExactly interface {
	TDataConnectedInd
	isTDataConnectedInd() bool
}

// _TDataConnectedInd is the data-structure of this message
type _TDataConnectedInd struct {
	*_CEMI
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TDataConnectedInd)  GetMessageCode() uint8 {
return 0x89}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TDataConnectedInd) InitializeParent(parent CEMI ) {}

func (m *_TDataConnectedInd)  GetParent() CEMI {
	return m._CEMI
}


// NewTDataConnectedInd factory function for _TDataConnectedInd
func NewTDataConnectedInd( size uint16 ) *_TDataConnectedInd {
	_result := &_TDataConnectedInd{
    	_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTDataConnectedInd(structType any) TDataConnectedInd {
    if casted, ok := structType.(TDataConnectedInd); ok {
		return casted
	}
	if casted, ok := structType.(*TDataConnectedInd); ok {
		return *casted
	}
	return nil
}

func (m *_TDataConnectedInd) GetTypeName() string {
	return "TDataConnectedInd"
}

func (m *_TDataConnectedInd) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	return lengthInBits
}


func (m *_TDataConnectedInd) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TDataConnectedIndParse(ctx context.Context, theBytes []byte, size uint16) (TDataConnectedInd, error) {
	return TDataConnectedIndParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), size)
}

func TDataConnectedIndParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, size uint16) (TDataConnectedInd, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("TDataConnectedInd"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TDataConnectedInd")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("TDataConnectedInd"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TDataConnectedInd")
	}

	// Create a partially initialized instance
	_child := &_TDataConnectedInd{
		_CEMI: &_CEMI{
			Size: size,
		},
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_TDataConnectedInd) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TDataConnectedInd) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TDataConnectedInd"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TDataConnectedInd")
		}

		if popErr := writeBuffer.PopContext("TDataConnectedInd"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TDataConnectedInd")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_TDataConnectedInd) isTDataConnectedInd() bool {
	return true
}

func (m *_TDataConnectedInd) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



