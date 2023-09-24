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


// NodeReference is the corresponding interface of NodeReference
type NodeReference interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() NodeId
	// GetReferenceTypeId returns ReferenceTypeId (property field)
	GetReferenceTypeId() NodeId
	// GetIsForward returns IsForward (property field)
	GetIsForward() bool
	// GetNoOfReferencedNodeIds returns NoOfReferencedNodeIds (property field)
	GetNoOfReferencedNodeIds() int32
	// GetReferencedNodeIds returns ReferencedNodeIds (property field)
	GetReferencedNodeIds() []NodeId
}

// NodeReferenceExactly can be used when we want exactly this type and not a type which fulfills NodeReference.
// This is useful for switch cases.
type NodeReferenceExactly interface {
	NodeReference
	isNodeReference() bool
}

// _NodeReference is the data-structure of this message
type _NodeReference struct {
	*_ExtensionObjectDefinition
        NodeId NodeId
        ReferenceTypeId NodeId
        IsForward bool
        NoOfReferencedNodeIds int32
        ReferencedNodeIds []NodeId
	// Reserved Fields
	reservedField0 *uint8
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NodeReference)  GetIdentifier() string {
return "582"}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NodeReference) InitializeParent(parent ExtensionObjectDefinition ) {}

func (m *_NodeReference)  GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NodeReference) GetNodeId() NodeId {
	return m.NodeId
}

func (m *_NodeReference) GetReferenceTypeId() NodeId {
	return m.ReferenceTypeId
}

func (m *_NodeReference) GetIsForward() bool {
	return m.IsForward
}

func (m *_NodeReference) GetNoOfReferencedNodeIds() int32 {
	return m.NoOfReferencedNodeIds
}

func (m *_NodeReference) GetReferencedNodeIds() []NodeId {
	return m.ReferencedNodeIds
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewNodeReference factory function for _NodeReference
func NewNodeReference( nodeId NodeId , referenceTypeId NodeId , isForward bool , noOfReferencedNodeIds int32 , referencedNodeIds []NodeId ) *_NodeReference {
	_result := &_NodeReference{
		NodeId: nodeId,
		ReferenceTypeId: referenceTypeId,
		IsForward: isForward,
		NoOfReferencedNodeIds: noOfReferencedNodeIds,
		ReferencedNodeIds: referencedNodeIds,
    	_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastNodeReference(structType any) NodeReference {
    if casted, ok := structType.(NodeReference); ok {
		return casted
	}
	if casted, ok := structType.(*NodeReference); ok {
		return *casted
	}
	return nil
}

func (m *_NodeReference) GetTypeName() string {
	return "NodeReference"
}

func (m *_NodeReference) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Simple field (referenceTypeId)
	lengthInBits += m.ReferenceTypeId.GetLengthInBits(ctx)

	// Reserved Field (reserved)
	lengthInBits += 7

	// Simple field (isForward)
	lengthInBits += 1;

	// Simple field (noOfReferencedNodeIds)
	lengthInBits += 32;

	// Array field
	if len(m.ReferencedNodeIds) > 0 {
		for _curItem, element := range m.ReferencedNodeIds {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ReferencedNodeIds), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{GetLengthInBits(context.Context) uint16}).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}


func (m *_NodeReference) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NodeReferenceParse(ctx context.Context, theBytes []byte, identifier string) (NodeReference, error) {
	return NodeReferenceParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func NodeReferenceParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (NodeReference, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("NodeReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NodeReference")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nodeId)
	if pullErr := readBuffer.PullContext("nodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodeId")
	}
_nodeId, _nodeIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _nodeIdErr != nil {
		return nil, errors.Wrap(_nodeIdErr, "Error parsing 'nodeId' field of NodeReference")
	}
	nodeId := _nodeId.(NodeId)
	if closeErr := readBuffer.CloseContext("nodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodeId")
	}

	// Simple Field (referenceTypeId)
	if pullErr := readBuffer.PullContext("referenceTypeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for referenceTypeId")
	}
_referenceTypeId, _referenceTypeIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _referenceTypeIdErr != nil {
		return nil, errors.Wrap(_referenceTypeIdErr, "Error parsing 'referenceTypeId' field of NodeReference")
	}
	referenceTypeId := _referenceTypeId.(NodeId)
	if closeErr := readBuffer.CloseContext("referenceTypeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for referenceTypeId")
	}

	var reservedField0 *uint8
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadUint8("reserved", 7)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of NodeReference")
		}
		if reserved != uint8(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value": reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (isForward)
_isForward, _isForwardErr := readBuffer.ReadBit("isForward")
	if _isForwardErr != nil {
		return nil, errors.Wrap(_isForwardErr, "Error parsing 'isForward' field of NodeReference")
	}
	isForward := _isForward

	// Simple Field (noOfReferencedNodeIds)
_noOfReferencedNodeIds, _noOfReferencedNodeIdsErr := readBuffer.ReadInt32("noOfReferencedNodeIds", 32)
	if _noOfReferencedNodeIdsErr != nil {
		return nil, errors.Wrap(_noOfReferencedNodeIdsErr, "Error parsing 'noOfReferencedNodeIds' field of NodeReference")
	}
	noOfReferencedNodeIds := _noOfReferencedNodeIds

	// Array field (referencedNodeIds)
	if pullErr := readBuffer.PullContext("referencedNodeIds", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for referencedNodeIds")
	}
	// Count array
	referencedNodeIds := make([]NodeId, utils.Max(noOfReferencedNodeIds, 0))
	// This happens when the size is set conditional to 0
	if len(referencedNodeIds) == 0 {
		referencedNodeIds = nil
	}
	{
		_numItems := uint16(utils.Max(noOfReferencedNodeIds, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := NodeIdParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'referencedNodeIds' field of NodeReference")
			}
			referencedNodeIds[_curItem] = _item.(NodeId)
		}
	}
	if closeErr := readBuffer.CloseContext("referencedNodeIds", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for referencedNodeIds")
	}

	if closeErr := readBuffer.CloseContext("NodeReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NodeReference")
	}

	// Create a partially initialized instance
	_child := &_NodeReference{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{
		},
		NodeId: nodeId,
		ReferenceTypeId: referenceTypeId,
		IsForward: isForward,
		NoOfReferencedNodeIds: noOfReferencedNodeIds,
		ReferencedNodeIds: referencedNodeIds,
		reservedField0: reservedField0,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_NodeReference) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NodeReference) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NodeReference"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NodeReference")
		}

	// Simple Field (nodeId)
	if pushErr := writeBuffer.PushContext("nodeId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for nodeId")
	}
	_nodeIdErr := writeBuffer.WriteSerializable(ctx, m.GetNodeId())
	if popErr := writeBuffer.PopContext("nodeId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for nodeId")
	}
	if _nodeIdErr != nil {
		return errors.Wrap(_nodeIdErr, "Error serializing 'nodeId' field")
	}

	// Simple Field (referenceTypeId)
	if pushErr := writeBuffer.PushContext("referenceTypeId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for referenceTypeId")
	}
	_referenceTypeIdErr := writeBuffer.WriteSerializable(ctx, m.GetReferenceTypeId())
	if popErr := writeBuffer.PopContext("referenceTypeId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for referenceTypeId")
	}
	if _referenceTypeIdErr != nil {
		return errors.Wrap(_referenceTypeIdErr, "Error serializing 'referenceTypeId' field")
	}

	// Reserved Field (reserved)
	{
		var reserved uint8 = uint8(0x00)
		if m.reservedField0 != nil {
			log.Info().Fields(map[string]any{
				"expected value": uint8(0x00),
				"got value": reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteUint8("reserved", 7, reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (isForward)
	isForward := bool(m.GetIsForward())
	_isForwardErr := writeBuffer.WriteBit("isForward", (isForward))
	if _isForwardErr != nil {
		return errors.Wrap(_isForwardErr, "Error serializing 'isForward' field")
	}

	// Simple Field (noOfReferencedNodeIds)
	noOfReferencedNodeIds := int32(m.GetNoOfReferencedNodeIds())
	_noOfReferencedNodeIdsErr := writeBuffer.WriteInt32("noOfReferencedNodeIds", 32, (noOfReferencedNodeIds))
	if _noOfReferencedNodeIdsErr != nil {
		return errors.Wrap(_noOfReferencedNodeIdsErr, "Error serializing 'noOfReferencedNodeIds' field")
	}

	// Array Field (referencedNodeIds)
	if pushErr := writeBuffer.PushContext("referencedNodeIds", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for referencedNodeIds")
	}
	for _curItem, _element := range m.GetReferencedNodeIds() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetReferencedNodeIds()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'referencedNodeIds' field")
		}
	}
	if popErr := writeBuffer.PopContext("referencedNodeIds", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for referencedNodeIds")
	}

		if popErr := writeBuffer.PopContext("NodeReference"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NodeReference")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_NodeReference) isNodeReference() bool {
	return true
}

func (m *_NodeReference) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



