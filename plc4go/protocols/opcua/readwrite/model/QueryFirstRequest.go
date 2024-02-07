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

// QueryFirstRequest is the corresponding interface of QueryFirstRequest
type QueryFirstRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetView returns View (property field)
	GetView() ExtensionObjectDefinition
	// GetNoOfNodeTypes returns NoOfNodeTypes (property field)
	GetNoOfNodeTypes() int32
	// GetNodeTypes returns NodeTypes (property field)
	GetNodeTypes() []ExtensionObjectDefinition
	// GetFilter returns Filter (property field)
	GetFilter() ExtensionObjectDefinition
	// GetMaxDataSetsToReturn returns MaxDataSetsToReturn (property field)
	GetMaxDataSetsToReturn() uint32
	// GetMaxReferencesToReturn returns MaxReferencesToReturn (property field)
	GetMaxReferencesToReturn() uint32
}

// QueryFirstRequestExactly can be used when we want exactly this type and not a type which fulfills QueryFirstRequest.
// This is useful for switch cases.
type QueryFirstRequestExactly interface {
	QueryFirstRequest
	isQueryFirstRequest() bool
}

// _QueryFirstRequest is the data-structure of this message
type _QueryFirstRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader         ExtensionObjectDefinition
	View                  ExtensionObjectDefinition
	NoOfNodeTypes         int32
	NodeTypes             []ExtensionObjectDefinition
	Filter                ExtensionObjectDefinition
	MaxDataSetsToReturn   uint32
	MaxReferencesToReturn uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_QueryFirstRequest) GetIdentifier() string {
	return "615"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_QueryFirstRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_QueryFirstRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_QueryFirstRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_QueryFirstRequest) GetView() ExtensionObjectDefinition {
	return m.View
}

func (m *_QueryFirstRequest) GetNoOfNodeTypes() int32 {
	return m.NoOfNodeTypes
}

func (m *_QueryFirstRequest) GetNodeTypes() []ExtensionObjectDefinition {
	return m.NodeTypes
}

func (m *_QueryFirstRequest) GetFilter() ExtensionObjectDefinition {
	return m.Filter
}

func (m *_QueryFirstRequest) GetMaxDataSetsToReturn() uint32 {
	return m.MaxDataSetsToReturn
}

func (m *_QueryFirstRequest) GetMaxReferencesToReturn() uint32 {
	return m.MaxReferencesToReturn
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewQueryFirstRequest factory function for _QueryFirstRequest
func NewQueryFirstRequest(requestHeader ExtensionObjectDefinition, view ExtensionObjectDefinition, noOfNodeTypes int32, nodeTypes []ExtensionObjectDefinition, filter ExtensionObjectDefinition, maxDataSetsToReturn uint32, maxReferencesToReturn uint32) *_QueryFirstRequest {
	_result := &_QueryFirstRequest{
		RequestHeader:              requestHeader,
		View:                       view,
		NoOfNodeTypes:              noOfNodeTypes,
		NodeTypes:                  nodeTypes,
		Filter:                     filter,
		MaxDataSetsToReturn:        maxDataSetsToReturn,
		MaxReferencesToReturn:      maxReferencesToReturn,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastQueryFirstRequest(structType any) QueryFirstRequest {
	if casted, ok := structType.(QueryFirstRequest); ok {
		return casted
	}
	if casted, ok := structType.(*QueryFirstRequest); ok {
		return *casted
	}
	return nil
}

func (m *_QueryFirstRequest) GetTypeName() string {
	return "QueryFirstRequest"
}

func (m *_QueryFirstRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (view)
	lengthInBits += m.View.GetLengthInBits(ctx)

	// Simple field (noOfNodeTypes)
	lengthInBits += 32

	// Array field
	if len(m.NodeTypes) > 0 {
		for _curItem, element := range m.NodeTypes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NodeTypes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (filter)
	lengthInBits += m.Filter.GetLengthInBits(ctx)

	// Simple field (maxDataSetsToReturn)
	lengthInBits += 32

	// Simple field (maxReferencesToReturn)
	lengthInBits += 32

	return lengthInBits
}

func (m *_QueryFirstRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func QueryFirstRequestParse(ctx context.Context, theBytes []byte, identifier string) (QueryFirstRequest, error) {
	return QueryFirstRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func QueryFirstRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (QueryFirstRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("QueryFirstRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for QueryFirstRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of QueryFirstRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (view)
	if pullErr := readBuffer.PullContext("view"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for view")
	}
	_view, _viewErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("513"))
	if _viewErr != nil {
		return nil, errors.Wrap(_viewErr, "Error parsing 'view' field of QueryFirstRequest")
	}
	view := _view.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("view"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for view")
	}

	// Simple Field (noOfNodeTypes)
	_noOfNodeTypes, _noOfNodeTypesErr := readBuffer.ReadInt32("noOfNodeTypes", 32)
	if _noOfNodeTypesErr != nil {
		return nil, errors.Wrap(_noOfNodeTypesErr, "Error parsing 'noOfNodeTypes' field of QueryFirstRequest")
	}
	noOfNodeTypes := _noOfNodeTypes

	// Array field (nodeTypes)
	if pullErr := readBuffer.PullContext("nodeTypes", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for nodeTypes")
	}
	// Count array
	nodeTypes := make([]ExtensionObjectDefinition, utils.Max(noOfNodeTypes, 0))
	// This happens when the size is set conditional to 0
	if len(nodeTypes) == 0 {
		nodeTypes = nil
	}
	{
		_numItems := uint16(utils.Max(noOfNodeTypes, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "575")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'nodeTypes' field of QueryFirstRequest")
			}
			nodeTypes[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("nodeTypes", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for nodeTypes")
	}

	// Simple Field (filter)
	if pullErr := readBuffer.PullContext("filter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for filter")
	}
	_filter, _filterErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("588"))
	if _filterErr != nil {
		return nil, errors.Wrap(_filterErr, "Error parsing 'filter' field of QueryFirstRequest")
	}
	filter := _filter.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("filter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for filter")
	}

	// Simple Field (maxDataSetsToReturn)
	_maxDataSetsToReturn, _maxDataSetsToReturnErr := readBuffer.ReadUint32("maxDataSetsToReturn", 32)
	if _maxDataSetsToReturnErr != nil {
		return nil, errors.Wrap(_maxDataSetsToReturnErr, "Error parsing 'maxDataSetsToReturn' field of QueryFirstRequest")
	}
	maxDataSetsToReturn := _maxDataSetsToReturn

	// Simple Field (maxReferencesToReturn)
	_maxReferencesToReturn, _maxReferencesToReturnErr := readBuffer.ReadUint32("maxReferencesToReturn", 32)
	if _maxReferencesToReturnErr != nil {
		return nil, errors.Wrap(_maxReferencesToReturnErr, "Error parsing 'maxReferencesToReturn' field of QueryFirstRequest")
	}
	maxReferencesToReturn := _maxReferencesToReturn

	if closeErr := readBuffer.CloseContext("QueryFirstRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for QueryFirstRequest")
	}

	// Create a partially initialized instance
	_child := &_QueryFirstRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		View:                       view,
		NoOfNodeTypes:              noOfNodeTypes,
		NodeTypes:                  nodeTypes,
		Filter:                     filter,
		MaxDataSetsToReturn:        maxDataSetsToReturn,
		MaxReferencesToReturn:      maxReferencesToReturn,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_QueryFirstRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_QueryFirstRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("QueryFirstRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for QueryFirstRequest")
		}

		// Simple Field (requestHeader)
		if pushErr := writeBuffer.PushContext("requestHeader"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for requestHeader")
		}
		_requestHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetRequestHeader())
		if popErr := writeBuffer.PopContext("requestHeader"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for requestHeader")
		}
		if _requestHeaderErr != nil {
			return errors.Wrap(_requestHeaderErr, "Error serializing 'requestHeader' field")
		}

		// Simple Field (view)
		if pushErr := writeBuffer.PushContext("view"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for view")
		}
		_viewErr := writeBuffer.WriteSerializable(ctx, m.GetView())
		if popErr := writeBuffer.PopContext("view"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for view")
		}
		if _viewErr != nil {
			return errors.Wrap(_viewErr, "Error serializing 'view' field")
		}

		// Simple Field (noOfNodeTypes)
		noOfNodeTypes := int32(m.GetNoOfNodeTypes())
		_noOfNodeTypesErr := writeBuffer.WriteInt32("noOfNodeTypes", 32, int32((noOfNodeTypes)))
		if _noOfNodeTypesErr != nil {
			return errors.Wrap(_noOfNodeTypesErr, "Error serializing 'noOfNodeTypes' field")
		}

		// Array Field (nodeTypes)
		if pushErr := writeBuffer.PushContext("nodeTypes", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for nodeTypes")
		}
		for _curItem, _element := range m.GetNodeTypes() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetNodeTypes()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'nodeTypes' field")
			}
		}
		if popErr := writeBuffer.PopContext("nodeTypes", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for nodeTypes")
		}

		// Simple Field (filter)
		if pushErr := writeBuffer.PushContext("filter"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for filter")
		}
		_filterErr := writeBuffer.WriteSerializable(ctx, m.GetFilter())
		if popErr := writeBuffer.PopContext("filter"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for filter")
		}
		if _filterErr != nil {
			return errors.Wrap(_filterErr, "Error serializing 'filter' field")
		}

		// Simple Field (maxDataSetsToReturn)
		maxDataSetsToReturn := uint32(m.GetMaxDataSetsToReturn())
		_maxDataSetsToReturnErr := writeBuffer.WriteUint32("maxDataSetsToReturn", 32, uint32((maxDataSetsToReturn)))
		if _maxDataSetsToReturnErr != nil {
			return errors.Wrap(_maxDataSetsToReturnErr, "Error serializing 'maxDataSetsToReturn' field")
		}

		// Simple Field (maxReferencesToReturn)
		maxReferencesToReturn := uint32(m.GetMaxReferencesToReturn())
		_maxReferencesToReturnErr := writeBuffer.WriteUint32("maxReferencesToReturn", 32, uint32((maxReferencesToReturn)))
		if _maxReferencesToReturnErr != nil {
			return errors.Wrap(_maxReferencesToReturnErr, "Error serializing 'maxReferencesToReturn' field")
		}

		if popErr := writeBuffer.PopContext("QueryFirstRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for QueryFirstRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_QueryFirstRequest) isQueryFirstRequest() bool {
	return true
}

func (m *_QueryFirstRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
