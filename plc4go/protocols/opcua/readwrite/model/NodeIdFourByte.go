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

// NodeIdFourByte is the corresponding interface of NodeIdFourByte
type NodeIdFourByte interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	NodeIdTypeDefinition
	// GetNamespaceIndex returns NamespaceIndex (property field)
	GetNamespaceIndex() uint8
	// GetId returns Id (property field)
	GetId() uint16
	// GetIdentifier returns Identifier (virtual field)
	GetIdentifier() string
}

// NodeIdFourByteExactly can be used when we want exactly this type and not a type which fulfills NodeIdFourByte.
// This is useful for switch cases.
type NodeIdFourByteExactly interface {
	NodeIdFourByte
	isNodeIdFourByte() bool
}

// _NodeIdFourByte is the data-structure of this message
type _NodeIdFourByte struct {
	*_NodeIdTypeDefinition
	NamespaceIndex uint8
	Id             uint16
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeIdFourByte) GetNodeType() NodeIdType {
	return NodeIdType_nodeIdTypeFourByte
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeIdFourByte) InitializeParent(parent NodeIdTypeDefinition) {}

func (m *_NodeIdFourByte) GetParent() NodeIdTypeDefinition {
	return m._NodeIdTypeDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeIdFourByte) GetNamespaceIndex() uint8 {
	return m.NamespaceIndex
}

func (m *_NodeIdFourByte) GetId() uint16 {
	return m.Id
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_NodeIdFourByte) GetIdentifier() string {
	ctx := context.Background()
	_ = ctx
	return fmt.Sprintf("%v", m.GetId())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewNodeIdFourByte factory function for _NodeIdFourByte
func NewNodeIdFourByte(namespaceIndex uint8, id uint16) *_NodeIdFourByte {
	_result := &_NodeIdFourByte{
		NamespaceIndex:        namespaceIndex,
		Id:                    id,
		_NodeIdTypeDefinition: NewNodeIdTypeDefinition(),
	}
	_result._NodeIdTypeDefinition._NodeIdTypeDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNodeIdFourByte(structType any) NodeIdFourByte {
	if casted, ok := structType.(NodeIdFourByte); ok {
		return casted
	}
	if casted, ok := structType.(*NodeIdFourByte); ok {
		return *casted
	}
	return nil
}

func (m *_NodeIdFourByte) GetTypeName() string {
	return "NodeIdFourByte"
}

func (m *_NodeIdFourByte) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (namespaceIndex)
	lengthInBits += 8

	// Simple field (id)
	lengthInBits += 16

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_NodeIdFourByte) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NodeIdFourByteParse(ctx context.Context, theBytes []byte) (NodeIdFourByte, error) {
	return NodeIdFourByteParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NodeIdFourByteParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NodeIdFourByte, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NodeIdFourByte"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeIdFourByte")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (namespaceIndex)
	_namespaceIndex, _namespaceIndexErr := readBuffer.ReadUint8("namespaceIndex", 8)
	if _namespaceIndexErr != nil {
		return nil, errors.Wrap(_namespaceIndexErr, "Error parsing 'namespaceIndex' field of NodeIdFourByte")
	}
	namespaceIndex := _namespaceIndex

	// Simple Field (id)
	_id, _idErr := readBuffer.ReadUint16("id", 16)
	if _idErr != nil {
		return nil, errors.Wrap(_idErr, "Error parsing 'id' field of NodeIdFourByte")
	}
	id := _id

	// Virtual field
	_identifier := id
	identifier := fmt.Sprintf("%v", _identifier)
	_ = identifier

	if closeErr := readBuffer.CloseContext("NodeIdFourByte"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeIdFourByte")
	}

	// Create a partially initialized instance
	_child := &_NodeIdFourByte{
		_NodeIdTypeDefinition: &_NodeIdTypeDefinition{},
		NamespaceIndex:        namespaceIndex,
		Id:                    id,
	}
	_child._NodeIdTypeDefinition._NodeIdTypeDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_NodeIdFourByte) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeIdFourByte) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeIdFourByte"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeIdFourByte")
		}

		// Simple Field (namespaceIndex)
		namespaceIndex := uint8(m.GetNamespaceIndex())
		_namespaceIndexErr := writeBuffer.WriteUint8("namespaceIndex", 8, (namespaceIndex))
		if _namespaceIndexErr != nil {
			return errors.Wrap(_namespaceIndexErr, "Error serializing 'namespaceIndex' field")
		}

		// Simple Field (id)
		id := uint16(m.GetId())
		_idErr := writeBuffer.WriteUint16("id", 16, (id))
		if _idErr != nil {
			return errors.Wrap(_idErr, "Error serializing 'id' field")
		}
		// Virtual field
		identifier := m.GetIdentifier()
		_ = identifier
		if _identifierErr := writeBuffer.WriteVirtual(ctx, "identifier", m.GetIdentifier()); _identifierErr != nil {
			return errors.Wrap(_identifierErr, "Error serializing 'identifier' field")
		}

		if popErr := writeBuffer.PopContext("NodeIdFourByte"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeIdFourByte")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NodeIdFourByte) isNodeIdFourByte() bool {
	return true
}

func (m *_NodeIdFourByte) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
