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

// NodeIdTwoByte is the corresponding interface of NodeIdTwoByte
type NodeIdTwoByte interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	NodeIdTypeDefinition
	// GetId returns Id (property field)
	GetId() uint8
	// GetIdentifier returns Identifier (virtual field)
	GetIdentifier() string
	// IsNodeIdTwoByte is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNodeIdTwoByte()
	// CreateBuilder creates a NodeIdTwoByteBuilder
	CreateNodeIdTwoByteBuilder() NodeIdTwoByteBuilder
}

// _NodeIdTwoByte is the data-structure of this message
type _NodeIdTwoByte struct {
	NodeIdTypeDefinitionContract
	Id uint8
}

var _ NodeIdTwoByte = (*_NodeIdTwoByte)(nil)
var _ NodeIdTypeDefinitionRequirements = (*_NodeIdTwoByte)(nil)

// NewNodeIdTwoByte factory function for _NodeIdTwoByte
func NewNodeIdTwoByte(id uint8) *_NodeIdTwoByte {
	_result := &_NodeIdTwoByte{
		NodeIdTypeDefinitionContract: NewNodeIdTypeDefinition(),
		Id:                           id,
	}
	_result.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NodeIdTwoByteBuilder is a builder for NodeIdTwoByte
type NodeIdTwoByteBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(id uint8) NodeIdTwoByteBuilder
	// WithId adds Id (property field)
	WithId(uint8) NodeIdTwoByteBuilder
	// Build builds the NodeIdTwoByte or returns an error if something is wrong
	Build() (NodeIdTwoByte, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NodeIdTwoByte
}

// NewNodeIdTwoByteBuilder() creates a NodeIdTwoByteBuilder
func NewNodeIdTwoByteBuilder() NodeIdTwoByteBuilder {
	return &_NodeIdTwoByteBuilder{_NodeIdTwoByte: new(_NodeIdTwoByte)}
}

type _NodeIdTwoByteBuilder struct {
	*_NodeIdTwoByte

	parentBuilder *_NodeIdTypeDefinitionBuilder

	err *utils.MultiError
}

var _ (NodeIdTwoByteBuilder) = (*_NodeIdTwoByteBuilder)(nil)

func (b *_NodeIdTwoByteBuilder) setParent(contract NodeIdTypeDefinitionContract) {
	b.NodeIdTypeDefinitionContract = contract
}

func (b *_NodeIdTwoByteBuilder) WithMandatoryFields(id uint8) NodeIdTwoByteBuilder {
	return b.WithId(id)
}

func (b *_NodeIdTwoByteBuilder) WithId(id uint8) NodeIdTwoByteBuilder {
	b.Id = id
	return b
}

func (b *_NodeIdTwoByteBuilder) Build() (NodeIdTwoByte, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NodeIdTwoByte.deepCopy(), nil
}

func (b *_NodeIdTwoByteBuilder) MustBuild() NodeIdTwoByte {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_NodeIdTwoByteBuilder) Done() NodeIdTypeDefinitionBuilder {
	return b.parentBuilder
}

func (b *_NodeIdTwoByteBuilder) buildForNodeIdTypeDefinition() (NodeIdTypeDefinition, error) {
	return b.Build()
}

func (b *_NodeIdTwoByteBuilder) DeepCopy() any {
	_copy := b.CreateNodeIdTwoByteBuilder().(*_NodeIdTwoByteBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNodeIdTwoByteBuilder creates a NodeIdTwoByteBuilder
func (b *_NodeIdTwoByte) CreateNodeIdTwoByteBuilder() NodeIdTwoByteBuilder {
	if b == nil {
		return NewNodeIdTwoByteBuilder()
	}
	return &_NodeIdTwoByteBuilder{_NodeIdTwoByte: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeIdTwoByte) GetNodeType() NodeIdType {
	return NodeIdType_nodeIdTypeTwoByte
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeIdTwoByte) GetParent() NodeIdTypeDefinitionContract {
	return m.NodeIdTypeDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeIdTwoByte) GetId() uint8 {
	return m.Id
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_NodeIdTwoByte) GetIdentifier() string {
	ctx := context.Background()
	_ = ctx
	return fmt.Sprintf("%v", m.GetId())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNodeIdTwoByte(structType any) NodeIdTwoByte {
	if casted, ok := structType.(NodeIdTwoByte); ok {
		return casted
	}
	if casted, ok := structType.(*NodeIdTwoByte); ok {
		return *casted
	}
	return nil
}

func (m *_NodeIdTwoByte) GetTypeName() string {
	return "NodeIdTwoByte"
}

func (m *_NodeIdTwoByte) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).getLengthInBits(ctx))

	// Simple field (id)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_NodeIdTwoByte) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NodeIdTwoByte) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_NodeIdTypeDefinition) (__nodeIdTwoByte NodeIdTwoByte, err error) {
	m.NodeIdTypeDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NodeIdTwoByte"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeIdTwoByte")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	id, err := ReadSimpleField(ctx, "id", ReadUnsignedByte(readBuffer, uint8(8)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'id' field"))
	}
	m.Id = id

	identifier, err := ReadVirtualField[string](ctx, "identifier", (*string)(nil), id)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'identifier' field"))
	}
	_ = identifier

	if closeErr := readBuffer.CloseContext("NodeIdTwoByte"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeIdTwoByte")
	}

	return m, nil
}

func (m *_NodeIdTwoByte) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeIdTwoByte) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeIdTwoByte"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeIdTwoByte")
		}

		if err := WriteSimpleField[uint8](ctx, "id", m.GetId(), WriteUnsignedByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'id' field")
		}
		// Virtual field
		identifier := m.GetIdentifier()
		_ = identifier
		if _identifierErr := writeBuffer.WriteVirtual(ctx, "identifier", m.GetIdentifier()); _identifierErr != nil {
			return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
		}

		if popErr := writeBuffer.PopContext("NodeIdTwoByte"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeIdTwoByte")
		}
		return nil
	}
	return m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeIdTwoByte) IsNodeIdTwoByte() {}

func (m *_NodeIdTwoByte) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NodeIdTwoByte) deepCopy() *_NodeIdTwoByte {
	if m == nil {
		return nil
	}
	_NodeIdTwoByteCopy := &_NodeIdTwoByte{
		m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition).deepCopy(),
		m.Id,
	}
	m.NodeIdTypeDefinitionContract.(*_NodeIdTypeDefinition)._SubType = m
	return _NodeIdTwoByteCopy
}

func (m *_NodeIdTwoByte) String() string {
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
