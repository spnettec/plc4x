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


// QueryDataSet is the corresponding interface of QueryDataSet
type QueryDataSet interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNodeId returns NodeId (property field)
	GetNodeId() ExpandedNodeId
	// GetTypeDefinitionNode returns TypeDefinitionNode (property field)
	GetTypeDefinitionNode() ExpandedNodeId
	// GetNoOfValues returns NoOfValues (property field)
	GetNoOfValues() int32
	// GetValues returns Values (property field)
	GetValues() []Variant
}

// QueryDataSetExactly can be used when we want exactly this type and not a type which fulfills QueryDataSet.
// This is useful for switch cases.
type QueryDataSetExactly interface {
	QueryDataSet
	isQueryDataSet() bool
}

// _QueryDataSet is the data-structure of this message
type _QueryDataSet struct {
	*_ExtensionObjectDefinition
        NodeId ExpandedNodeId
        TypeDefinitionNode ExpandedNodeId
        NoOfValues int32
        Values []Variant
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_QueryDataSet)  GetIdentifier() string {
return "579"}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_QueryDataSet) InitializeParent(parent ExtensionObjectDefinition ) {}

func (m *_QueryDataSet)  GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_QueryDataSet) GetNodeId() ExpandedNodeId {
	return m.NodeId
}

func (m *_QueryDataSet) GetTypeDefinitionNode() ExpandedNodeId {
	return m.TypeDefinitionNode
}

func (m *_QueryDataSet) GetNoOfValues() int32 {
	return m.NoOfValues
}

func (m *_QueryDataSet) GetValues() []Variant {
	return m.Values
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewQueryDataSet factory function for _QueryDataSet
func NewQueryDataSet( nodeId ExpandedNodeId , typeDefinitionNode ExpandedNodeId , noOfValues int32 , values []Variant ) *_QueryDataSet {
	_result := &_QueryDataSet{
		NodeId: nodeId,
		TypeDefinitionNode: typeDefinitionNode,
		NoOfValues: noOfValues,
		Values: values,
    	_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastQueryDataSet(structType any) QueryDataSet {
    if casted, ok := structType.(QueryDataSet); ok {
		return casted
	}
	if casted, ok := structType.(*QueryDataSet); ok {
		return *casted
	}
	return nil
}

func (m *_QueryDataSet) GetTypeName() string {
	return "QueryDataSet"
}

func (m *_QueryDataSet) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (nodeId)
	lengthInBits += m.NodeId.GetLengthInBits(ctx)

	// Simple field (typeDefinitionNode)
	lengthInBits += m.TypeDefinitionNode.GetLengthInBits(ctx)

	// Simple field (noOfValues)
	lengthInBits += 32;

	// Array field
	if len(m.Values) > 0 {
		for _curItem, element := range m.Values {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Values), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{GetLengthInBits(context.Context) uint16}).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}


func (m *_QueryDataSet) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func QueryDataSetParse(ctx context.Context, theBytes []byte, identifier string) (QueryDataSet, error) {
	return QueryDataSetParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func QueryDataSetParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (QueryDataSet, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("QueryDataSet"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for QueryDataSet")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (nodeId)
	if pullErr := readBuffer.PullContext("nodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodeId")
	}
_nodeId, _nodeIdErr := ExpandedNodeIdParseWithBuffer(ctx, readBuffer)
	if _nodeIdErr != nil {
		return nil, errors.Wrap(_nodeIdErr, "Error parsing 'nodeId' field of QueryDataSet")
	}
	nodeId := _nodeId.(ExpandedNodeId)
	if closeErr := readBuffer.CloseContext("nodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodeId")
	}

	// Simple Field (typeDefinitionNode)
	if pullErr := readBuffer.PullContext("typeDefinitionNode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for typeDefinitionNode")
	}
_typeDefinitionNode, _typeDefinitionNodeErr := ExpandedNodeIdParseWithBuffer(ctx, readBuffer)
	if _typeDefinitionNodeErr != nil {
		return nil, errors.Wrap(_typeDefinitionNodeErr, "Error parsing 'typeDefinitionNode' field of QueryDataSet")
	}
	typeDefinitionNode := _typeDefinitionNode.(ExpandedNodeId)
	if closeErr := readBuffer.CloseContext("typeDefinitionNode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for typeDefinitionNode")
	}

	// Simple Field (noOfValues)
_noOfValues, _noOfValuesErr := readBuffer.ReadInt32("noOfValues", 32)
	if _noOfValuesErr != nil {
		return nil, errors.Wrap(_noOfValuesErr, "Error parsing 'noOfValues' field of QueryDataSet")
	}
	noOfValues := _noOfValues

	// Array field (values)
	if pullErr := readBuffer.PullContext("values", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for values")
	}
	// Count array
	values := make([]Variant, utils.Max(noOfValues, 0))
	// This happens when the size is set conditional to 0
	if len(values) == 0 {
		values = nil
	}
	{
		_numItems := uint16(utils.Max(noOfValues, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := VariantParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'values' field of QueryDataSet")
			}
			values[_curItem] = _item.(Variant)
		}
	}
	if closeErr := readBuffer.CloseContext("values", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for values")
	}

	if closeErr := readBuffer.CloseContext("QueryDataSet"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for QueryDataSet")
	}

	// Create a partially initialized instance
	_child := &_QueryDataSet{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{
		},
		NodeId: nodeId,
		TypeDefinitionNode: typeDefinitionNode,
		NoOfValues: noOfValues,
		Values: values,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_QueryDataSet) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_QueryDataSet) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("QueryDataSet"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for QueryDataSet")
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

	// Simple Field (typeDefinitionNode)
	if pushErr := writeBuffer.PushContext("typeDefinitionNode"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for typeDefinitionNode")
	}
	_typeDefinitionNodeErr := writeBuffer.WriteSerializable(ctx, m.GetTypeDefinitionNode())
	if popErr := writeBuffer.PopContext("typeDefinitionNode"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for typeDefinitionNode")
	}
	if _typeDefinitionNodeErr != nil {
		return errors.Wrap(_typeDefinitionNodeErr, "Error serializing 'typeDefinitionNode' field")
	}

	// Simple Field (noOfValues)
	noOfValues := int32(m.GetNoOfValues())
	_noOfValuesErr := writeBuffer.WriteInt32("noOfValues", 32, (noOfValues))
	if _noOfValuesErr != nil {
		return errors.Wrap(_noOfValuesErr, "Error serializing 'noOfValues' field")
	}

	// Array Field (values)
	if pushErr := writeBuffer.PushContext("values", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for values")
	}
	for _curItem, _element := range m.GetValues() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetValues()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'values' field")
		}
	}
	if popErr := writeBuffer.PopContext("values", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for values")
	}

		if popErr := writeBuffer.PopContext("QueryDataSet"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for QueryDataSet")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_QueryDataSet) isQueryDataSet() bool {
	return true
}

func (m *_QueryDataSet) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



