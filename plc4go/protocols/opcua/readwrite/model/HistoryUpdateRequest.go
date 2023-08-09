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


// HistoryUpdateRequest is the corresponding interface of HistoryUpdateRequest
type HistoryUpdateRequest interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetRequestHeader returns RequestHeader (property field)
	GetRequestHeader() ExtensionObjectDefinition
	// GetNoOfHistoryUpdateDetails returns NoOfHistoryUpdateDetails (property field)
	GetNoOfHistoryUpdateDetails() int32
	// GetHistoryUpdateDetails returns HistoryUpdateDetails (property field)
	GetHistoryUpdateDetails() []ExtensionObject
}

// HistoryUpdateRequestExactly can be used when we want exactly this type and not a type which fulfills HistoryUpdateRequest.
// This is useful for switch cases.
type HistoryUpdateRequestExactly interface {
	HistoryUpdateRequest
	isHistoryUpdateRequest() bool
}

// _HistoryUpdateRequest is the data-structure of this message
type _HistoryUpdateRequest struct {
	*_ExtensionObjectDefinition
        RequestHeader ExtensionObjectDefinition
        NoOfHistoryUpdateDetails int32
        HistoryUpdateDetails []ExtensionObject
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_HistoryUpdateRequest)  GetIdentifier() string {
return "700"}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_HistoryUpdateRequest) InitializeParent(parent ExtensionObjectDefinition ) {}

func (m *_HistoryUpdateRequest)  GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_HistoryUpdateRequest) GetRequestHeader() ExtensionObjectDefinition {
	return m.RequestHeader
}

func (m *_HistoryUpdateRequest) GetNoOfHistoryUpdateDetails() int32 {
	return m.NoOfHistoryUpdateDetails
}

func (m *_HistoryUpdateRequest) GetHistoryUpdateDetails() []ExtensionObject {
	return m.HistoryUpdateDetails
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewHistoryUpdateRequest factory function for _HistoryUpdateRequest
func NewHistoryUpdateRequest( requestHeader ExtensionObjectDefinition , noOfHistoryUpdateDetails int32 , historyUpdateDetails []ExtensionObject ) *_HistoryUpdateRequest {
	_result := &_HistoryUpdateRequest{
		RequestHeader: requestHeader,
		NoOfHistoryUpdateDetails: noOfHistoryUpdateDetails,
		HistoryUpdateDetails: historyUpdateDetails,
    	_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastHistoryUpdateRequest(structType any) HistoryUpdateRequest {
    if casted, ok := structType.(HistoryUpdateRequest); ok {
		return casted
	}
	if casted, ok := structType.(*HistoryUpdateRequest); ok {
		return *casted
	}
	return nil
}

func (m *_HistoryUpdateRequest) GetTypeName() string {
	return "HistoryUpdateRequest"
}

func (m *_HistoryUpdateRequest) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (requestHeader)
	lengthInBits += m.RequestHeader.GetLengthInBits(ctx)

	// Simple field (noOfHistoryUpdateDetails)
	lengthInBits += 32;

	// Array field
	if len(m.HistoryUpdateDetails) > 0 {
		for _curItem, element := range m.HistoryUpdateDetails {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.HistoryUpdateDetails), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{GetLengthInBits(context.Context) uint16}).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}


func (m *_HistoryUpdateRequest) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func HistoryUpdateRequestParse(ctx context.Context, theBytes []byte, identifier string) (HistoryUpdateRequest, error) {
	return HistoryUpdateRequestParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func HistoryUpdateRequestParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (HistoryUpdateRequest, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("HistoryUpdateRequest"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for HistoryUpdateRequest")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (requestHeader)
	if pullErr := readBuffer.PullContext("requestHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for requestHeader")
	}
_requestHeader, _requestHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer , string( "391" ) )
	if _requestHeaderErr != nil {
		return nil, errors.Wrap(_requestHeaderErr, "Error parsing 'requestHeader' field of HistoryUpdateRequest")
	}
	requestHeader := _requestHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("requestHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for requestHeader")
	}

	// Simple Field (noOfHistoryUpdateDetails)
_noOfHistoryUpdateDetails, _noOfHistoryUpdateDetailsErr := readBuffer.ReadInt32("noOfHistoryUpdateDetails", 32)
	if _noOfHistoryUpdateDetailsErr != nil {
		return nil, errors.Wrap(_noOfHistoryUpdateDetailsErr, "Error parsing 'noOfHistoryUpdateDetails' field of HistoryUpdateRequest")
	}
	noOfHistoryUpdateDetails := _noOfHistoryUpdateDetails

	// Array field (historyUpdateDetails)
	if pullErr := readBuffer.PullContext("historyUpdateDetails", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for historyUpdateDetails")
	}
	// Count array
	historyUpdateDetails := make([]ExtensionObject, utils.Max(noOfHistoryUpdateDetails, 0))
	// This happens when the size is set conditional to 0
	if len(historyUpdateDetails) == 0 {
		historyUpdateDetails = nil
	}
	{
		_numItems := uint16(utils.Max(noOfHistoryUpdateDetails, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := ExtensionObjectParseWithBuffer(arrayCtx, readBuffer , bool(true) )
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'historyUpdateDetails' field of HistoryUpdateRequest")
			}
			historyUpdateDetails[_curItem] = _item.(ExtensionObject)
		}
	}
	if closeErr := readBuffer.CloseContext("historyUpdateDetails", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for historyUpdateDetails")
	}

	if closeErr := readBuffer.CloseContext("HistoryUpdateRequest"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for HistoryUpdateRequest")
	}

	// Create a partially initialized instance
	_child := &_HistoryUpdateRequest{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{
		},
		RequestHeader: requestHeader,
		NoOfHistoryUpdateDetails: noOfHistoryUpdateDetails,
		HistoryUpdateDetails: historyUpdateDetails,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_HistoryUpdateRequest) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_HistoryUpdateRequest) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("HistoryUpdateRequest"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for HistoryUpdateRequest")
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

	// Simple Field (noOfHistoryUpdateDetails)
	noOfHistoryUpdateDetails := int32(m.GetNoOfHistoryUpdateDetails())
	_noOfHistoryUpdateDetailsErr := writeBuffer.WriteInt32("noOfHistoryUpdateDetails", 32, (noOfHistoryUpdateDetails))
	if _noOfHistoryUpdateDetailsErr != nil {
		return errors.Wrap(_noOfHistoryUpdateDetailsErr, "Error serializing 'noOfHistoryUpdateDetails' field")
	}

	// Array Field (historyUpdateDetails)
	if pushErr := writeBuffer.PushContext("historyUpdateDetails", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for historyUpdateDetails")
	}
	for _curItem, _element := range m.GetHistoryUpdateDetails() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetHistoryUpdateDetails()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'historyUpdateDetails' field")
		}
	}
	if popErr := writeBuffer.PopContext("historyUpdateDetails", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for historyUpdateDetails")
	}

		if popErr := writeBuffer.PopContext("HistoryUpdateRequest"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for HistoryUpdateRequest")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_HistoryUpdateRequest) isHistoryUpdateRequest() bool {
	return true
}

func (m *_HistoryUpdateRequest) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



