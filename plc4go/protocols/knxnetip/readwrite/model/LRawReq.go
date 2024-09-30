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

// LRawReq is the corresponding interface of LRawReq
type LRawReq interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	CEMI
	// IsLRawReq is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsLRawReq()
	// CreateBuilder creates a LRawReqBuilder
	CreateLRawReqBuilder() LRawReqBuilder
}

// _LRawReq is the data-structure of this message
type _LRawReq struct {
	CEMIContract
}

var _ LRawReq = (*_LRawReq)(nil)
var _ CEMIRequirements = (*_LRawReq)(nil)

// NewLRawReq factory function for _LRawReq
func NewLRawReq(size uint16) *_LRawReq {
	_result := &_LRawReq{
		CEMIContract: NewCEMI(size),
	}
	_result.CEMIContract.(*_CEMI)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// LRawReqBuilder is a builder for LRawReq
type LRawReqBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields() LRawReqBuilder
	// Build builds the LRawReq or returns an error if something is wrong
	Build() (LRawReq, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() LRawReq
}

// NewLRawReqBuilder() creates a LRawReqBuilder
func NewLRawReqBuilder() LRawReqBuilder {
	return &_LRawReqBuilder{_LRawReq: new(_LRawReq)}
}

type _LRawReqBuilder struct {
	*_LRawReq

	parentBuilder *_CEMIBuilder

	err *utils.MultiError
}

var _ (LRawReqBuilder) = (*_LRawReqBuilder)(nil)

func (b *_LRawReqBuilder) setParent(contract CEMIContract) {
	b.CEMIContract = contract
}

func (b *_LRawReqBuilder) WithMandatoryFields() LRawReqBuilder {
	return b
}

func (b *_LRawReqBuilder) Build() (LRawReq, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._LRawReq.deepCopy(), nil
}

func (b *_LRawReqBuilder) MustBuild() LRawReq {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_LRawReqBuilder) Done() CEMIBuilder {
	return b.parentBuilder
}

func (b *_LRawReqBuilder) buildForCEMI() (CEMI, error) {
	return b.Build()
}

func (b *_LRawReqBuilder) DeepCopy() any {
	_copy := b.CreateLRawReqBuilder().(*_LRawReqBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateLRawReqBuilder creates a LRawReqBuilder
func (b *_LRawReq) CreateLRawReqBuilder() LRawReqBuilder {
	if b == nil {
		return NewLRawReqBuilder()
	}
	return &_LRawReqBuilder{_LRawReq: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LRawReq) GetMessageCode() uint8 {
	return 0x10
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LRawReq) GetParent() CEMIContract {
	return m.CEMIContract
}

// Deprecated: use the interface for direct cast
func CastLRawReq(structType any) LRawReq {
	if casted, ok := structType.(LRawReq); ok {
		return casted
	}
	if casted, ok := structType.(*LRawReq); ok {
		return *casted
	}
	return nil
}

func (m *_LRawReq) GetTypeName() string {
	return "LRawReq"
}

func (m *_LRawReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.CEMIContract.(*_CEMI).getLengthInBits(ctx))

	return lengthInBits
}

func (m *_LRawReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_LRawReq) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_CEMI, size uint16) (__lRawReq LRawReq, err error) {
	m.CEMIContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LRawReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LRawReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("LRawReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LRawReq")
	}

	return m, nil
}

func (m *_LRawReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_LRawReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LRawReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LRawReq")
		}

		if popErr := writeBuffer.PopContext("LRawReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LRawReq")
		}
		return nil
	}
	return m.CEMIContract.(*_CEMI).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_LRawReq) IsLRawReq() {}

func (m *_LRawReq) DeepCopy() any {
	return m.deepCopy()
}

func (m *_LRawReq) deepCopy() *_LRawReq {
	if m == nil {
		return nil
	}
	_LRawReqCopy := &_LRawReq{
		m.CEMIContract.(*_CEMI).deepCopy(),
	}
	m.CEMIContract.(*_CEMI)._SubType = m
	return _LRawReqCopy
}

func (m *_LRawReq) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
