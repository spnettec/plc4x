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

// TranslateBrowsePathsToNodeIdsRequest is the corresponding interface of TranslateBrowsePathsToNodeIdsRequest
type TranslateBrowsePathsToNodeIdsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfBrowsePaths returns NoOfBrowsePaths (property field)
	GetNoOfBrowsePaths() int32
	// GetBrowsePaths returns BrowsePaths (property field)
	GetBrowsePaths() []ExtensionObjectDefinition
}

// TranslateBrowsePathsToNodeIdsRequestExactly can be used when we want exactly this type and not a type which fulfills TranslateBrowsePathsToNodeIdsRequest.
// This is useful for switch cases.
type TranslateBrowsePathsToNodeIdsRequestExactly interface {
	TranslateBrowsePathsToNodeIdsRequest
	isTranslateBrowsePathsToNodeIdsRequest() bool
}

// _TranslateBrowsePathsToNodeIdsRequest is the data-structure of this message
type _TranslateBrowsePathsToNodeIdsRequest struct {
	*_ExtensionObjectDefinition
	RequestHeader   ExtensionObjectDefinition
	NoOfBrowsePaths int32
	BrowsePaths     []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TranslateBrowsePathsToNodeIdsRequest) GetIdentifier() string {
	return "554"
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TranslateBrowsePathsToNodeIdsRequest) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_TranslateBrowsePathsToNodeIdsRequest) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TranslateBrowsePathsToNodeIdsRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) GetNoOfBrowsePaths() int32 {
	return m.NoOfBrowsePaths
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) GetBrowsePaths() []ExtensionObjectDefinition {
	return m.BrowsePaths
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewTranslateBrowsePathsToNodeIdsRequest factory function for _TranslateBrowsePathsToNodeIdsRequest
func NewTranslateBrowsePathsToNodeIdsRequest(requestHeader ExtensionObjectDefinition, noOfBrowsePaths int32, browsePaths []ExtensionObjectDefinition) *_TranslateBrowsePathsToNodeIdsRequest {
	_result := &_TranslateBrowsePathsToNodeIdsRequest{
		RequestHeader:              requestHeader,
		NoOfBrowsePaths:            noOfBrowsePaths,
		BrowsePaths:                browsePaths,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTranslateBrowsePathsToNodeIdsRequest(structType any) TranslateBrowsePathsToNodeIdsRequest {
	if casted, ok := structType.(TranslateBrowsePathsToNodeIdsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*TranslateBrowsePathsToNodeIdsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) GetTypeName() string {
	return "TranslateBrowsePathsToNodeIdsRequest"
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfBrowsePaths)
	lengthInBits += 32

	// Array field
	if len(m.BrowsePaths) > 0 {
		for _curItem, element := range m.BrowsePaths {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.BrowsePaths), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TranslateBrowsePathsToNodeIdsRequestParse(ctx context.Context, theBytes []byte, identifier string) (TranslateBrowsePathsToNodeIdsRequest, error) {
	return TranslateBrowsePathsToNodeIdsRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func TranslateBrowsePathsToNodeIdsRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (TranslateBrowsePathsToNodeIdsRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("TranslateBrowsePathsToNodeIdsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TranslateBrowsePathsToNodeIdsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
	_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer, string("391"))
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of TranslateBrowsePathsToNodeIdsRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (noOfBrowsePaths)
	_noOfBrowsePaths, _noOfBrowsePathsErr := readBuffer.ReadInt32("noOfBrowsePaths", 32)
	if _noOfBrowsePathsErr != nil {
		return nil, errors.Wrap(_noOfBrowsePathsErr, "Error parsing 'noOfBrowsePaths' field of TranslateBrowsePathsToNodeIdsRequest")
	}
	noOfBrowsePaths := _noOfBrowsePaths

	// Array field (browsePaths)
	if pullErr := readBuffer.PullContext("browsePaths", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for browsePaths")
	}
	// Count array
	browsePaths := make([]ExtensionObjectDefinition, utils.Max(noOfBrowsePaths, 0))
	// This happens when the size is set conditional to 0
	if len(browsePaths) == 0 {
		browsePaths = nil
	}
	{
		_numItems := uint16(utils.Max(noOfBrowsePaths, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "545")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'browsePaths' field of TranslateBrowsePathsToNodeIdsRequest")
			}
			browsePaths[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("browsePaths", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for browsePaths")
	}

	if closeErr := readBuffer.CloseContext("TranslateBrowsePathsToNodeIdsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TranslateBrowsePathsToNodeIdsRequest")
	}

	// Create a partially initialized instance
	_child := &_TranslateBrowsePathsToNodeIdsRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		RequestHeader:              requestHeader,
		NoOfBrowsePaths:            noOfBrowsePaths,
		BrowsePaths:                browsePaths,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TranslateBrowsePathsToNodeIdsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TranslateBrowsePathsToNodeIdsRequest")
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

		// Simple Field (noOfBrowsePaths)
		noOfBrowsePaths := int32(m.GetNoOfBrowsePaths())
		_noOfBrowsePathsErr := writeBuffer.WriteInt32("noOfBrowsePaths", 32, (noOfBrowsePaths))
		if _noOfBrowsePathsErr != nil {
			return errors.Wrap(_noOfBrowsePathsErr, "Error serializing 'noOfBrowsePaths' field")
		}

		// Array Field (browsePaths)
		if pushErr := writeBuffer.PushContext("browsePaths", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for browsePaths")
		}
		for _curItem, _element := range m.GetBrowsePaths() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetBrowsePaths()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'browsePaths' field")
			}
		}
		if popErr := writeBuffer.PopContext("browsePaths", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for browsePaths")
		}

		if popErr := writeBuffer.PopContext("TranslateBrowsePathsToNodeIdsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TranslateBrowsePathsToNodeIdsRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) isTranslateBrowsePathsToNodeIdsRequest() bool {
	return true
}

func (m *_TranslateBrowsePathsToNodeIdsRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
