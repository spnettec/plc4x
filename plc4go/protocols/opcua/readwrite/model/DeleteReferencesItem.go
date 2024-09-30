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

// DeleteReferencesItem is the corresponding interface of DeleteReferencesItem
type DeleteReferencesItem interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetSourceNodeId returns SourceNodeId (property field)
	GetSourceNodeId() NodeId
	// GetReferenceTypeId returns ReferenceTypeId (property field)
	GetReferenceTypeId() NodeId
	// GetIsForward returns IsForward (property field)
	GetIsForward() bool
	// GetTargetNodeId returns TargetNodeId (property field)
	GetTargetNodeId() ExpandedNodeId
	// GetDeleteBidirectional returns DeleteBidirectional (property field)
	GetDeleteBidirectional() bool
	// IsDeleteReferencesItem is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsDeleteReferencesItem()
	// CreateBuilder creates a DeleteReferencesItemBuilder
	CreateDeleteReferencesItemBuilder() DeleteReferencesItemBuilder
}

// _DeleteReferencesItem is the data-structure of this message
type _DeleteReferencesItem struct {
	ExtensionObjectDefinitionContract
	SourceNodeId        NodeId
	ReferenceTypeId     NodeId
	IsForward           bool
	TargetNodeId        ExpandedNodeId
	DeleteBidirectional bool
	// Reserved Fields
	reservedField0 *uint8
	reservedField1 *uint8
}

var _ DeleteReferencesItem = (*_DeleteReferencesItem)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_DeleteReferencesItem)(nil)

// NewDeleteReferencesItem factory function for _DeleteReferencesItem
func NewDeleteReferencesItem(sourceNodeId NodeId, referenceTypeId NodeId, isForward bool, targetNodeId ExpandedNodeId, deleteBidirectional bool) *_DeleteReferencesItem {
	if sourceNodeId == nil {
		panic("sourceNodeId of type NodeId for DeleteReferencesItem must not be nil")
	}
	if referenceTypeId == nil {
		panic("referenceTypeId of type NodeId for DeleteReferencesItem must not be nil")
	}
	if targetNodeId == nil {
		panic("targetNodeId of type ExpandedNodeId for DeleteReferencesItem must not be nil")
	}
	_result := &_DeleteReferencesItem{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		SourceNodeId:                      sourceNodeId,
		ReferenceTypeId:                   referenceTypeId,
		IsForward:                         isForward,
		TargetNodeId:                      targetNodeId,
		DeleteBidirectional:               deleteBidirectional,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// DeleteReferencesItemBuilder is a builder for DeleteReferencesItem
type DeleteReferencesItemBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(sourceNodeId NodeId, referenceTypeId NodeId, isForward bool, targetNodeId ExpandedNodeId, deleteBidirectional bool) DeleteReferencesItemBuilder
	// WithSourceNodeId adds SourceNodeId (property field)
	WithSourceNodeId(NodeId) DeleteReferencesItemBuilder
	// WithSourceNodeIdBuilder adds SourceNodeId (property field) which is build by the builder
	WithSourceNodeIdBuilder(func(NodeIdBuilder) NodeIdBuilder) DeleteReferencesItemBuilder
	// WithReferenceTypeId adds ReferenceTypeId (property field)
	WithReferenceTypeId(NodeId) DeleteReferencesItemBuilder
	// WithReferenceTypeIdBuilder adds ReferenceTypeId (property field) which is build by the builder
	WithReferenceTypeIdBuilder(func(NodeIdBuilder) NodeIdBuilder) DeleteReferencesItemBuilder
	// WithIsForward adds IsForward (property field)
	WithIsForward(bool) DeleteReferencesItemBuilder
	// WithTargetNodeId adds TargetNodeId (property field)
	WithTargetNodeId(ExpandedNodeId) DeleteReferencesItemBuilder
	// WithTargetNodeIdBuilder adds TargetNodeId (property field) which is build by the builder
	WithTargetNodeIdBuilder(func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) DeleteReferencesItemBuilder
	// WithDeleteBidirectional adds DeleteBidirectional (property field)
	WithDeleteBidirectional(bool) DeleteReferencesItemBuilder
	// Build builds the DeleteReferencesItem or returns an error if something is wrong
	Build() (DeleteReferencesItem, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() DeleteReferencesItem
}

// NewDeleteReferencesItemBuilder() creates a DeleteReferencesItemBuilder
func NewDeleteReferencesItemBuilder() DeleteReferencesItemBuilder {
	return &_DeleteReferencesItemBuilder{_DeleteReferencesItem: new(_DeleteReferencesItem)}
}

type _DeleteReferencesItemBuilder struct {
	*_DeleteReferencesItem

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (DeleteReferencesItemBuilder) = (*_DeleteReferencesItemBuilder)(nil)

func (b *_DeleteReferencesItemBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_DeleteReferencesItemBuilder) WithMandatoryFields(sourceNodeId NodeId, referenceTypeId NodeId, isForward bool, targetNodeId ExpandedNodeId, deleteBidirectional bool) DeleteReferencesItemBuilder {
	return b.WithSourceNodeId(sourceNodeId).WithReferenceTypeId(referenceTypeId).WithIsForward(isForward).WithTargetNodeId(targetNodeId).WithDeleteBidirectional(deleteBidirectional)
}

func (b *_DeleteReferencesItemBuilder) WithSourceNodeId(sourceNodeId NodeId) DeleteReferencesItemBuilder {
	b.SourceNodeId = sourceNodeId
	return b
}

func (b *_DeleteReferencesItemBuilder) WithSourceNodeIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) DeleteReferencesItemBuilder {
	builder := builderSupplier(b.SourceNodeId.CreateNodeIdBuilder())
	var err error
	b.SourceNodeId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_DeleteReferencesItemBuilder) WithReferenceTypeId(referenceTypeId NodeId) DeleteReferencesItemBuilder {
	b.ReferenceTypeId = referenceTypeId
	return b
}

func (b *_DeleteReferencesItemBuilder) WithReferenceTypeIdBuilder(builderSupplier func(NodeIdBuilder) NodeIdBuilder) DeleteReferencesItemBuilder {
	builder := builderSupplier(b.ReferenceTypeId.CreateNodeIdBuilder())
	var err error
	b.ReferenceTypeId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "NodeIdBuilder failed"))
	}
	return b
}

func (b *_DeleteReferencesItemBuilder) WithIsForward(isForward bool) DeleteReferencesItemBuilder {
	b.IsForward = isForward
	return b
}

func (b *_DeleteReferencesItemBuilder) WithTargetNodeId(targetNodeId ExpandedNodeId) DeleteReferencesItemBuilder {
	b.TargetNodeId = targetNodeId
	return b
}

func (b *_DeleteReferencesItemBuilder) WithTargetNodeIdBuilder(builderSupplier func(ExpandedNodeIdBuilder) ExpandedNodeIdBuilder) DeleteReferencesItemBuilder {
	builder := builderSupplier(b.TargetNodeId.CreateExpandedNodeIdBuilder())
	var err error
	b.TargetNodeId, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ExpandedNodeIdBuilder failed"))
	}
	return b
}

func (b *_DeleteReferencesItemBuilder) WithDeleteBidirectional(deleteBidirectional bool) DeleteReferencesItemBuilder {
	b.DeleteBidirectional = deleteBidirectional
	return b
}

func (b *_DeleteReferencesItemBuilder) Build() (DeleteReferencesItem, error) {
	if b.SourceNodeId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'sourceNodeId' not set"))
	}
	if b.ReferenceTypeId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'referenceTypeId' not set"))
	}
	if b.TargetNodeId == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'targetNodeId' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._DeleteReferencesItem.deepCopy(), nil
}

func (b *_DeleteReferencesItemBuilder) MustBuild() DeleteReferencesItem {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_DeleteReferencesItemBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_DeleteReferencesItemBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_DeleteReferencesItemBuilder) DeepCopy() any {
	_copy := b.CreateDeleteReferencesItemBuilder().(*_DeleteReferencesItemBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateDeleteReferencesItemBuilder creates a DeleteReferencesItemBuilder
func (b *_DeleteReferencesItem) CreateDeleteReferencesItemBuilder() DeleteReferencesItemBuilder {
	if b == nil {
		return NewDeleteReferencesItemBuilder()
	}
	return &_DeleteReferencesItemBuilder{_DeleteReferencesItem: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_DeleteReferencesItem) GetIdentifier() string {
	return "387"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_DeleteReferencesItem) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_DeleteReferencesItem) GetSourceNodeId() NodeId {
	return m.SourceNodeId
}

func (m *_DeleteReferencesItem) GetReferenceTypeId() NodeId {
	return m.ReferenceTypeId
}

func (m *_DeleteReferencesItem) GetIsForward() bool {
	return m.IsForward
}

func (m *_DeleteReferencesItem) GetTargetNodeId() ExpandedNodeId {
	return m.TargetNodeId
}

func (m *_DeleteReferencesItem) GetDeleteBidirectional() bool {
	return m.DeleteBidirectional
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastDeleteReferencesItem(structType any) DeleteReferencesItem {
	if casted, ok := structType.(DeleteReferencesItem); ok {
		return casted
	}
	if casted, ok := structType.(*DeleteReferencesItem); ok {
		return *casted
	}
	return nil
}

func (m *_DeleteReferencesItem) GetTypeName() string {
	return "DeleteReferencesItem"
}

func (m *_DeleteReferencesItem) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (sourceNodeId)
	lengthInBits += m.SourceNodeId.GetLengthInBits(ctx)

	// Simple field (referenceTypeId)
	lengthInBits += m.ReferenceTypeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (isForward)
	lengthInBits += 1

	// Simple field (targetNodeId)
	lengthInBits += m.TargetNodeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (deleteBidirectional)
	lengthInBits += 1

	return lengthInBits
}

func (m *_DeleteReferencesItem) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_DeleteReferencesItem) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__deleteReferencesItem DeleteReferencesItem, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("DeleteReferencesItem"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for DeleteReferencesItem")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	sourceNodeId, err := ReadSimpleField[NodeId](ctx, "sourceNodeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'sourceNodeId' field"))
	}
	m.SourceNodeId = sourceNodeId

	referenceTypeId, err := ReadSimpleField[NodeId](ctx, "referenceTypeId", ReadComplex[NodeId](NodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'referenceTypeId' field"))
	}
	m.ReferenceTypeId = referenceTypeId

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	isForward, err := ReadSimpleField(ctx, "isForward", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isForward' field"))
	}
	m.IsForward = isForward

	targetNodeId, err := ReadSimpleField[ExpandedNodeId](ctx, "targetNodeId", ReadComplex[ExpandedNodeId](ExpandedNodeIdParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'targetNodeId' field"))
	}
	m.TargetNodeId = targetNodeId

	reservedField1, err := ReadReservedField(ctx, "reserved", ReadUnsignedByte(readBuffer, uint8(7)), uint8(0x00))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField1 = reservedField1

	deleteBidirectional, err := ReadSimpleField(ctx, "deleteBidirectional", ReadBoolean(readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'deleteBidirectional' field"))
	}
	m.DeleteBidirectional = deleteBidirectional

	if closeErr := readBuffer.CloseContext("DeleteReferencesItem"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for DeleteReferencesItem")
	}

	return m, nil
}

func (m *_DeleteReferencesItem) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_DeleteReferencesItem) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("DeleteReferencesItem"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for DeleteReferencesItem")
		}

		if err := WriteSimpleField[NodeId](ctx, "sourceNodeId", m.GetSourceNodeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'sourceNodeId' field")
		}

		if err := WriteSimpleField[NodeId](ctx, "referenceTypeId", m.GetReferenceTypeId(), WriteComplex[NodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'referenceTypeId' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}

		if err := WriteSimpleField[bool](ctx, "isForward", m.GetIsForward(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'isForward' field")
		}

		if err := WriteSimpleField[ExpandedNodeId](ctx, "targetNodeId", m.GetTargetNodeId(), WriteComplex[ExpandedNodeId](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'targetNodeId' field")
		}

		if err := WriteReservedField[uint8](ctx, "reserved", uint8(0x00), WriteUnsignedByte(writeBuffer, 7)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 2")
		}

		if err := WriteSimpleField[bool](ctx, "deleteBidirectional", m.GetDeleteBidirectional(), WriteBoolean(writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'deleteBidirectional' field")
		}

		if popErr := writeBuffer.PopContext("DeleteReferencesItem"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for DeleteReferencesItem")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_DeleteReferencesItem) IsDeleteReferencesItem() {}

func (m *_DeleteReferencesItem) DeepCopy() any {
	return m.deepCopy()
}

func (m *_DeleteReferencesItem) deepCopy() *_DeleteReferencesItem {
	if m == nil {
		return nil
	}
	_DeleteReferencesItemCopy := &_DeleteReferencesItem{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.SourceNodeId.DeepCopy().(NodeId),
		m.ReferenceTypeId.DeepCopy().(NodeId),
		m.IsForward,
		m.TargetNodeId.DeepCopy().(ExpandedNodeId),
		m.DeleteBidirectional,
		m.reservedField0,
		m.reservedField1,
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _DeleteReferencesItemCopy
}

func (m *_DeleteReferencesItem) String() string {
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
