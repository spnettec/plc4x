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


// ModifyMonitoredItemsRequest is the corresponding interface of ModifyMonitoredItemsRequest
type ModifyMonitoredItemsRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetSubscriptionId returns SubscriptionId (property field)
	GetSubscriptionId() uint32
	// GetTimestampsToReturn returns TimestampsToReturn (property field)
	GetTimestampsToReturn() TimestampsToReturn
	// GetNoOfItemsToModify returns NoOfItemsToModify (property field)
	GetNoOfItemsToModify() int32
	// GetItemsToModify returns ItemsToModify (property field)
	GetItemsToModify() []ExtensionObjectDefinition
}

// ModifyMonitoredItemsRequestExactly can be used when we want exactly this type and not a type which fulfills ModifyMonitoredItemsRequest.
// This is useful for switch cases.
type ModifyMonitoredItemsRequestExactly interface {
	ModifyMonitoredItemsRequest
	isModifyMonitoredItemsRequest() bool
}

// _ModifyMonitoredItemsRequest is the data-structure of this message
type _ModifyMonitoredItemsRequest struct {
	*_ExtensionObjectDefinition
        RequestHeader ExtensionObjectDefinition
        SubscriptionId uint32
        TimestampsToReturn TimestampsToReturn
        NoOfItemsToModify int32
        ItemsToModify []ExtensionObjectDefinition
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ModifyMonitoredItemsRequest)  GetIdentifier() string {
return "763"}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ModifyMonitoredItemsRequest) InitializeParent(parent ExtensionObjectDefinition ) {}

func (m *_ModifyMonitoredItemsRequest)  GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ModifyMonitoredItemsRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_ModifyMonitoredItemsRequest) GetSubscriptionId() uint32 {
	return m.SubscriptionId
}

func (m *_ModifyMonitoredItemsRequest) GetTimestampsToReturn() TimestampsToReturn {
	return m.TimestampsToReturn
}

func (m *_ModifyMonitoredItemsRequest) GetNoOfItemsToModify() int32 {
	return m.NoOfItemsToModify
}

func (m *_ModifyMonitoredItemsRequest) GetItemsToModify() []ExtensionObjectDefinition {
	return m.ItemsToModify
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewModifyMonitoredItemsRequest factory function for _ModifyMonitoredItemsRequest
func NewModifyMonitoredItemsRequest( requestHeader ExtensionObjectDefinition , subscriptionId uint32 , timestampsToReturn TimestampsToReturn , noOfItemsToModify int32 , itemsToModify []ExtensionObjectDefinition ) *_ModifyMonitoredItemsRequest {
	_result := &_ModifyMonitoredItemsRequest{
		RequestHeader: requestHeader,
		SubscriptionId: subscriptionId,
		TimestampsToReturn: timestampsToReturn,
		NoOfItemsToModify: noOfItemsToModify,
		ItemsToModify: itemsToModify,
    	_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastModifyMonitoredItemsRequest(structType any) ModifyMonitoredItemsRequest {
    if casted, ok := structType.(ModifyMonitoredItemsRequest); ok {
		return casted
	}
	if casted, ok := structType.(*ModifyMonitoredItemsRequest); ok {
		return *casted
	}
	return nil
}

func (m *_ModifyMonitoredItemsRequest) GetTypeName() string {
	return "ModifyMonitoredItemsRequest"
}

func (m *_ModifyMonitoredItemsRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (subscriptionId)
	lengthInBits += 32;

	// Simple field (timestampsToReturn)
	lengthInBits += 32

	// Simple field (noOfItemsToModify)
	lengthInBits += 32;

	// Array field
	if len(m.ItemsToModify) > 0 {
		for _curItem, element := range m.ItemsToModify {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ItemsToModify), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{GetLengthInBits(context.Context) uint16}).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}


func (m *_ModifyMonitoredItemsRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ModifyMonitoredItemsRequestParse(ctx context.Context, theBytes []byte, identifier string) (ModifyMonitoredItemsRequest, error) {
	return ModifyMonitoredItemsRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ModifyMonitoredItemsRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ModifyMonitoredItemsRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ModifyMonitoredItemsRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ModifyMonitoredItemsRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer , string( "391" ) )
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of ModifyMonitoredItemsRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (subscriptionId)
_subscriptionId, _subscriptionIdErr := readBuffer.ReadUint32("subscriptionId", 32)
	if _subscriptionIdErr != nil {
		return nil, errors.Wrap(_subscriptionIdErr, "Error parsing 'subscriptionId' field of ModifyMonitoredItemsRequest")
	}
	subscriptionId := _subscriptionId

	// Simple Field (timestampsToReturn)
	if pullErr := readBuffer.PullContext("timestampsToReturn"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timestampsToReturn")
	}
_timestampsToReturn, _timestampsToReturnErr := TimestampsToReturnParseWithBuffer(ctx, readBuffer)
	if _timestampsToReturnErr != nil {
		return nil, errors.Wrap(_timestampsToReturnErr, "Error parsing 'timestampsToReturn' field of ModifyMonitoredItemsRequest")
	}
	timestampsToReturn := _timestampsToReturn
	if closeErr := readBuffer.CloseContext("timestampsToReturn"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timestampsToReturn")
	}

	// Simple Field (noOfItemsToModify)
_noOfItemsToModify, _noOfItemsToModifyErr := readBuffer.ReadInt32("noOfItemsToModify", 32)
	if _noOfItemsToModifyErr != nil {
		return nil, errors.Wrap(_noOfItemsToModifyErr, "Error parsing 'noOfItemsToModify' field of ModifyMonitoredItemsRequest")
	}
	noOfItemsToModify := _noOfItemsToModify

	// Array field (itemsToModify)
	if pullErr := readBuffer.PullContext("itemsToModify", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for itemsToModify")
	}
	// Count array
	itemsToModify := make([]ExtensionObjectDefinition, utils.Max(noOfItemsToModify, 0))
	// This happens when the size is set conditional to 0
	if len(itemsToModify) == 0 {
		itemsToModify = nil
	}
	{
		_numItems := uint16(utils.Max(noOfItemsToModify, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer , "757" )
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'itemsToModify' field of ModifyMonitoredItemsRequest")
			}
			itemsToModify[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("itemsToModify", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for itemsToModify")
	}

	if closeErr := readBuffer.CloseContext("ModifyMonitoredItemsRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ModifyMonitoredItemsRequest")
	}

	// Create a partially initialized instance
	_child := &_ModifyMonitoredItemsRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{
		},
		RequestHeader: requestHeader,
		SubscriptionId: subscriptionId,
		TimestampsToReturn: timestampsToReturn,
		NoOfItemsToModify: noOfItemsToModify,
		ItemsToModify: itemsToModify,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ModifyMonitoredItemsRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ModifyMonitoredItemsRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ModifyMonitoredItemsRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ModifyMonitoredItemsRequest")
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

	// Simple Field (subscriptionId)
	subscriptionId := uint32(m.GetSubscriptionId())
	_subscriptionIdErr := writeBuffer.WriteUint32("subscriptionId", 32, (subscriptionId))
	if _subscriptionIdErr != nil {
		return errors.Wrap(_subscriptionIdErr, "Error serializing 'subscriptionId' field")
	}

	// Simple Field (timestampsToReturn)
	if pushErr := writeBuffer.PushContext("timestampsToReturn"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timestampsToReturn")
	}
	_timestampsToReturnErr := writeBuffer.WriteSerializable(ctx, m.GetTimestampsToReturn())
	if popErr := writeBuffer.PopContext("timestampsToReturn"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timestampsToReturn")
	}
	if _timestampsToReturnErr != nil {
		return errors.Wrap(_timestampsToReturnErr, "Error serializing 'timestampsToReturn' field")
	}

	// Simple Field (noOfItemsToModify)
	noOfItemsToModify := int32(m.GetNoOfItemsToModify())
	_noOfItemsToModifyErr := writeBuffer.WriteInt32("noOfItemsToModify", 32, (noOfItemsToModify))
	if _noOfItemsToModifyErr != nil {
		return errors.Wrap(_noOfItemsToModifyErr, "Error serializing 'noOfItemsToModify' field")
	}

	// Array Field (itemsToModify)
	if pushErr := writeBuffer.PushContext("itemsToModify", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for itemsToModify")
	}
	for _curItem, _element := range m.GetItemsToModify() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetItemsToModify()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'itemsToModify' field")
		}
	}
	if popErr := writeBuffer.PopContext("itemsToModify", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for itemsToModify")
	}

		if popErr := writeBuffer.PopContext("ModifyMonitoredItemsRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ModifyMonitoredItemsRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_ModifyMonitoredItemsRequest) isModifyMonitoredItemsRequest() bool {
	return true
}

func (m *_ModifyMonitoredItemsRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



