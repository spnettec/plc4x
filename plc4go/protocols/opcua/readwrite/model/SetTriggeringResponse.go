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


// SetTriggeringResponse is the corresponding interface of SetTriggeringResponse
type SetTriggeringResponse interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetResponseHeader returns ResponseHeader (property field)
	GetResponseHeader() ExtensionObjectDefinition
	// GetNoOfAddResults returns NoOfAddResults (property field)
	GetNoOfAddResults() int32
	// GetAddResults returns AddResults (property field)
	GetAddResults() []StatusCode
	// GetNoOfAddDiagnosticInfos returns NoOfAddDiagnosticInfos (property field)
	GetNoOfAddDiagnosticInfos() int32
	// GetAddDiagnosticInfos returns AddDiagnosticInfos (property field)
	GetAddDiagnosticInfos() []DiagnosticInfo
	// GetNoOfRemoveResults returns NoOfRemoveResults (property field)
	GetNoOfRemoveResults() int32
	// GetRemoveResults returns RemoveResults (property field)
	GetRemoveResults() []StatusCode
	// GetNoOfRemoveDiagnosticInfos returns NoOfRemoveDiagnosticInfos (property field)
	GetNoOfRemoveDiagnosticInfos() int32
	// GetRemoveDiagnosticInfos returns RemoveDiagnosticInfos (property field)
	GetRemoveDiagnosticInfos() []DiagnosticInfo
}

// SetTriggeringResponseExactly can be used when we want exactly this type and not a type which fulfills SetTriggeringResponse.
// This is useful for switch cases.
type SetTriggeringResponseExactly interface {
	SetTriggeringResponse
	isSetTriggeringResponse() bool
}

// _SetTriggeringResponse is the data-structure of this message
type _SetTriggeringResponse struct {
	*_ExtensionObjectDefinition
        ResponseHeader ExtensionObjectDefinition
        NoOfAddResults int32
        AddResults []StatusCode
        NoOfAddDiagnosticInfos int32
        AddDiagnosticInfos []DiagnosticInfo
        NoOfRemoveResults int32
        RemoveResults []StatusCode
        NoOfRemoveDiagnosticInfos int32
        RemoveDiagnosticInfos []DiagnosticInfo
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SetTriggeringResponse)  GetIdentifier() string {
return "778"}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SetTriggeringResponse) InitializeParent(parent ExtensionObjectDefinition ) {}

func (m *_SetTriggeringResponse)  GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SetTriggeringResponse) GetResponseHeader() ExtensionObjectDefinition {
	return m.ResponseHeader
}

func (m *_SetTriggeringResponse) GetNoOfAddResults() int32 {
	return m.NoOfAddResults
}

func (m *_SetTriggeringResponse) GetAddResults() []StatusCode {
	return m.AddResults
}

func (m *_SetTriggeringResponse) GetNoOfAddDiagnosticInfos() int32 {
	return m.NoOfAddDiagnosticInfos
}

func (m *_SetTriggeringResponse) GetAddDiagnosticInfos() []DiagnosticInfo {
	return m.AddDiagnosticInfos
}

func (m *_SetTriggeringResponse) GetNoOfRemoveResults() int32 {
	return m.NoOfRemoveResults
}

func (m *_SetTriggeringResponse) GetRemoveResults() []StatusCode {
	return m.RemoveResults
}

func (m *_SetTriggeringResponse) GetNoOfRemoveDiagnosticInfos() int32 {
	return m.NoOfRemoveDiagnosticInfos
}

func (m *_SetTriggeringResponse) GetRemoveDiagnosticInfos() []DiagnosticInfo {
	return m.RemoveDiagnosticInfos
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewSetTriggeringResponse factory function for _SetTriggeringResponse
func NewSetTriggeringResponse( responseHeader ExtensionObjectDefinition , noOfAddResults int32 , addResults []StatusCode , noOfAddDiagnosticInfos int32 , addDiagnosticInfos []DiagnosticInfo , noOfRemoveResults int32 , removeResults []StatusCode , noOfRemoveDiagnosticInfos int32 , removeDiagnosticInfos []DiagnosticInfo ) *_SetTriggeringResponse {
	_result := &_SetTriggeringResponse{
		ResponseHeader: responseHeader,
		NoOfAddResults: noOfAddResults,
		AddResults: addResults,
		NoOfAddDiagnosticInfos: noOfAddDiagnosticInfos,
		AddDiagnosticInfos: addDiagnosticInfos,
		NoOfRemoveResults: noOfRemoveResults,
		RemoveResults: removeResults,
		NoOfRemoveDiagnosticInfos: noOfRemoveDiagnosticInfos,
		RemoveDiagnosticInfos: removeDiagnosticInfos,
    	_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSetTriggeringResponse(structType any) SetTriggeringResponse {
    if casted, ok := structType.(SetTriggeringResponse); ok {
		return casted
	}
	if casted, ok := structType.(*SetTriggeringResponse); ok {
		return *casted
	}
	return nil
}

func (m *_SetTriggeringResponse) GetTypeName() string {
	return "SetTriggeringResponse"
}

func (m *_SetTriggeringResponse) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (responseHeader)
	lengthInBits += m.ResponseHeader.GetLengthInBits(ctx)

	// Simple field (noOfAddResults)
	lengthInBits += 32;

	// Array field
	if len(m.AddResults) > 0 {
		for _curItem, element := range m.AddResults {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.AddResults), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{GetLengthInBits(context.Context) uint16}).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfAddDiagnosticInfos)
	lengthInBits += 32;

	// Array field
	if len(m.AddDiagnosticInfos) > 0 {
		for _curItem, element := range m.AddDiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.AddDiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{GetLengthInBits(context.Context) uint16}).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfRemoveResults)
	lengthInBits += 32;

	// Array field
	if len(m.RemoveResults) > 0 {
		for _curItem, element := range m.RemoveResults {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.RemoveResults), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{GetLengthInBits(context.Context) uint16}).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfRemoveDiagnosticInfos)
	lengthInBits += 32;

	// Array field
	if len(m.RemoveDiagnosticInfos) > 0 {
		for _curItem, element := range m.RemoveDiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.RemoveDiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{GetLengthInBits(context.Context) uint16}).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}


func (m *_SetTriggeringResponse) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SetTriggeringResponseParse(ctx context.Context, theBytes []byte, identifier string) (SetTriggeringResponse, error) {
	return SetTriggeringResponseParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func SetTriggeringResponseParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (SetTriggeringResponse, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SetTriggeringResponse"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SetTriggeringResponse")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (responseHeader)
	if pullErr := readBuffer.PullContext("responseHeader"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for responseHeader")
	}
_responseHeader, _responseHeaderErr := ExtensionObjectDefinitionParseWithBuffer(ctx, readBuffer , string( "394" ) )
	if _responseHeaderErr != nil {
		return nil, errors.Wrap(_responseHeaderErr, "Error parsing 'responseHeader' field of SetTriggeringResponse")
	}
	responseHeader := _responseHeader.(ExtensionObjectDefinition)
	if closeErr := readBuffer.CloseContext("responseHeader"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for responseHeader")
	}

	// Simple Field (noOfAddResults)
_noOfAddResults, _noOfAddResultsErr := readBuffer.ReadInt32("noOfAddResults", 32)
	if _noOfAddResultsErr != nil {
		return nil, errors.Wrap(_noOfAddResultsErr, "Error parsing 'noOfAddResults' field of SetTriggeringResponse")
	}
	noOfAddResults := _noOfAddResults

	// Array field (addResults)
	if pullErr := readBuffer.PullContext("addResults", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for addResults")
	}
	// Count array
	addResults := make([]StatusCode, noOfAddResults)
	// This happens when the size is set conditional to 0
	if len(addResults) == 0 {
		addResults = nil
	}
	{
		_numItems := uint16(noOfAddResults)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := StatusCodeParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'addResults' field of SetTriggeringResponse")
			}
			addResults[_curItem] = _item.(StatusCode)
		}
	}
	if closeErr := readBuffer.CloseContext("addResults", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for addResults")
	}

	// Simple Field (noOfAddDiagnosticInfos)
_noOfAddDiagnosticInfos, _noOfAddDiagnosticInfosErr := readBuffer.ReadInt32("noOfAddDiagnosticInfos", 32)
	if _noOfAddDiagnosticInfosErr != nil {
		return nil, errors.Wrap(_noOfAddDiagnosticInfosErr, "Error parsing 'noOfAddDiagnosticInfos' field of SetTriggeringResponse")
	}
	noOfAddDiagnosticInfos := _noOfAddDiagnosticInfos

	// Array field (addDiagnosticInfos)
	if pullErr := readBuffer.PullContext("addDiagnosticInfos", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for addDiagnosticInfos")
	}
	// Count array
	addDiagnosticInfos := make([]DiagnosticInfo, noOfAddDiagnosticInfos)
	// This happens when the size is set conditional to 0
	if len(addDiagnosticInfos) == 0 {
		addDiagnosticInfos = nil
	}
	{
		_numItems := uint16(noOfAddDiagnosticInfos)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := DiagnosticInfoParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'addDiagnosticInfos' field of SetTriggeringResponse")
			}
			addDiagnosticInfos[_curItem] = _item.(DiagnosticInfo)
		}
	}
	if closeErr := readBuffer.CloseContext("addDiagnosticInfos", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for addDiagnosticInfos")
	}

	// Simple Field (noOfRemoveResults)
_noOfRemoveResults, _noOfRemoveResultsErr := readBuffer.ReadInt32("noOfRemoveResults", 32)
	if _noOfRemoveResultsErr != nil {
		return nil, errors.Wrap(_noOfRemoveResultsErr, "Error parsing 'noOfRemoveResults' field of SetTriggeringResponse")
	}
	noOfRemoveResults := _noOfRemoveResults

	// Array field (removeResults)
	if pullErr := readBuffer.PullContext("removeResults", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for removeResults")
	}
	// Count array
	removeResults := make([]StatusCode, noOfRemoveResults)
	// This happens when the size is set conditional to 0
	if len(removeResults) == 0 {
		removeResults = nil
	}
	{
		_numItems := uint16(noOfRemoveResults)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := StatusCodeParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'removeResults' field of SetTriggeringResponse")
			}
			removeResults[_curItem] = _item.(StatusCode)
		}
	}
	if closeErr := readBuffer.CloseContext("removeResults", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for removeResults")
	}

	// Simple Field (noOfRemoveDiagnosticInfos)
_noOfRemoveDiagnosticInfos, _noOfRemoveDiagnosticInfosErr := readBuffer.ReadInt32("noOfRemoveDiagnosticInfos", 32)
	if _noOfRemoveDiagnosticInfosErr != nil {
		return nil, errors.Wrap(_noOfRemoveDiagnosticInfosErr, "Error parsing 'noOfRemoveDiagnosticInfos' field of SetTriggeringResponse")
	}
	noOfRemoveDiagnosticInfos := _noOfRemoveDiagnosticInfos

	// Array field (removeDiagnosticInfos)
	if pullErr := readBuffer.PullContext("removeDiagnosticInfos", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for removeDiagnosticInfos")
	}
	// Count array
	removeDiagnosticInfos := make([]DiagnosticInfo, noOfRemoveDiagnosticInfos)
	// This happens when the size is set conditional to 0
	if len(removeDiagnosticInfos) == 0 {
		removeDiagnosticInfos = nil
	}
	{
		_numItems := uint16(noOfRemoveDiagnosticInfos)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := DiagnosticInfoParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'removeDiagnosticInfos' field of SetTriggeringResponse")
			}
			removeDiagnosticInfos[_curItem] = _item.(DiagnosticInfo)
		}
	}
	if closeErr := readBuffer.CloseContext("removeDiagnosticInfos", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for removeDiagnosticInfos")
	}

	if closeErr := readBuffer.CloseContext("SetTriggeringResponse"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SetTriggeringResponse")
	}

	// Create a partially initialized instance
	_child := &_SetTriggeringResponse{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{
		},
		ResponseHeader: responseHeader,
		NoOfAddResults: noOfAddResults,
		AddResults: addResults,
		NoOfAddDiagnosticInfos: noOfAddDiagnosticInfos,
		AddDiagnosticInfos: addDiagnosticInfos,
		NoOfRemoveResults: noOfRemoveResults,
		RemoveResults: removeResults,
		NoOfRemoveDiagnosticInfos: noOfRemoveDiagnosticInfos,
		RemoveDiagnosticInfos: removeDiagnosticInfos,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_SetTriggeringResponse) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SetTriggeringResponse) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SetTriggeringResponse"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SetTriggeringResponse")
		}

	// Simple Field (responseHeader)
	if pushErr := writeBuffer.PushContext("responseHeader"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for responseHeader")
	}
	_responseHeaderErr := writeBuffer.WriteSerializable(ctx, m.GetResponseHeader())
	if popErr := writeBuffer.PopContext("responseHeader"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for responseHeader")
	}
	if _responseHeaderErr != nil {
		return errors.Wrap(_responseHeaderErr, "Error serializing 'responseHeader' field")
	}

	// Simple Field (noOfAddResults)
	noOfAddResults := int32(m.GetNoOfAddResults())
	_noOfAddResultsErr := writeBuffer.WriteInt32("noOfAddResults", 32, (noOfAddResults))
	if _noOfAddResultsErr != nil {
		return errors.Wrap(_noOfAddResultsErr, "Error serializing 'noOfAddResults' field")
	}

	// Array Field (addResults)
	if pushErr := writeBuffer.PushContext("addResults", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for addResults")
	}
	for _curItem, _element := range m.GetAddResults() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetAddResults()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'addResults' field")
		}
	}
	if popErr := writeBuffer.PopContext("addResults", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for addResults")
	}

	// Simple Field (noOfAddDiagnosticInfos)
	noOfAddDiagnosticInfos := int32(m.GetNoOfAddDiagnosticInfos())
	_noOfAddDiagnosticInfosErr := writeBuffer.WriteInt32("noOfAddDiagnosticInfos", 32, (noOfAddDiagnosticInfos))
	if _noOfAddDiagnosticInfosErr != nil {
		return errors.Wrap(_noOfAddDiagnosticInfosErr, "Error serializing 'noOfAddDiagnosticInfos' field")
	}

	// Array Field (addDiagnosticInfos)
	if pushErr := writeBuffer.PushContext("addDiagnosticInfos", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for addDiagnosticInfos")
	}
	for _curItem, _element := range m.GetAddDiagnosticInfos() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetAddDiagnosticInfos()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'addDiagnosticInfos' field")
		}
	}
	if popErr := writeBuffer.PopContext("addDiagnosticInfos", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for addDiagnosticInfos")
	}

	// Simple Field (noOfRemoveResults)
	noOfRemoveResults := int32(m.GetNoOfRemoveResults())
	_noOfRemoveResultsErr := writeBuffer.WriteInt32("noOfRemoveResults", 32, (noOfRemoveResults))
	if _noOfRemoveResultsErr != nil {
		return errors.Wrap(_noOfRemoveResultsErr, "Error serializing 'noOfRemoveResults' field")
	}

	// Array Field (removeResults)
	if pushErr := writeBuffer.PushContext("removeResults", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for removeResults")
	}
	for _curItem, _element := range m.GetRemoveResults() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetRemoveResults()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'removeResults' field")
		}
	}
	if popErr := writeBuffer.PopContext("removeResults", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for removeResults")
	}

	// Simple Field (noOfRemoveDiagnosticInfos)
	noOfRemoveDiagnosticInfos := int32(m.GetNoOfRemoveDiagnosticInfos())
	_noOfRemoveDiagnosticInfosErr := writeBuffer.WriteInt32("noOfRemoveDiagnosticInfos", 32, (noOfRemoveDiagnosticInfos))
	if _noOfRemoveDiagnosticInfosErr != nil {
		return errors.Wrap(_noOfRemoveDiagnosticInfosErr, "Error serializing 'noOfRemoveDiagnosticInfos' field")
	}

	// Array Field (removeDiagnosticInfos)
	if pushErr := writeBuffer.PushContext("removeDiagnosticInfos", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for removeDiagnosticInfos")
	}
	for _curItem, _element := range m.GetRemoveDiagnosticInfos() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetRemoveDiagnosticInfos()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'removeDiagnosticInfos' field")
		}
	}
	if popErr := writeBuffer.PopContext("removeDiagnosticInfos", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for removeDiagnosticInfos")
	}

		if popErr := writeBuffer.PopContext("SetTriggeringResponse"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SetTriggeringResponse")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_SetTriggeringResponse) isSetTriggeringResponse() bool {
	return true
}

func (m *_SetTriggeringResponse) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



