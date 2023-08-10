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

// RelativePath is the corresponding interface of RelativePath
type RelativePath interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetNoOfElements returns NoOfElements (property field)
	GetNoOfElements() int32
	// GetElements returns Elements (property field)
	GetElements() []ExtensionObjectDefinition
}

// RelativePathExactly can be used when we want exactly this type and not a type which fulfills RelativePath.
// This is useful for switch cases.
type RelativePathExactly interface {
	RelativePath
	isRelativePath() bool
}

// _RelativePath is the data-structure of this message
type _RelativePath struct {
	*_ExtensionObjectDefinition
	NoOfElements int32
	Elements     []ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_RelativePath) GetIdentifier() string {
	return "542"
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RelativePath) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_RelativePath) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RelativePath) GetNoOfElements() int32 {
	return m.NoOfElements
}

func (m *_RelativePath) GetElements() []ExtensionObjectDefinition {
	return m.Elements
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRelativePath factory function for _RelativePath
func NewRelativePath(noOfElements int32, elements []ExtensionObjectDefinition) *_RelativePath {
	_result := &_RelativePath{
		NoOfElements:               noOfElements,
		Elements:                   elements,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRelativePath(structType any) RelativePath {
	if casted, ok := structType.(RelativePath); ok {
		return casted
	}
	if casted, ok := structType.(*RelativePath); ok {
		return *casted
	}
	return nil
}

func (m *_RelativePath) GetTypeName() string {
	return "RelativePath"
}

func (m *_RelativePath) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (noOfElements)
	lengthInBits += 32

	// Array field
	if len(m.Elements) > 0 {
		for _curItem, element := range m.Elements {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Elements), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_RelativePath) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RelativePathParse(ctx context.Context, theBytes []byte, identifier string) (RelativePath, error) {
	return RelativePathParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func RelativePathParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (RelativePath, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("RelativePath"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RelativePath")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (noOfElements)
	_noOfElements, _noOfElementsErr := readBuffer.ReadInt32("noOfElements", 32)
	if _noOfElementsErr != nil {
		return nil, errors.Wrap(_noOfElementsErr, "Error parsing 'noOfElements' field of RelativePath")
	}
	noOfElements := _noOfElements

	// Array field (elements)
	if pullErr := readBuffer.PullContext("elements", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for elements")
	}
	// Count array
	elements := make([]ExtensionObjectDefinition, utils.Max(noOfElements, 0))
	// This happens when the size is set conditional to 0
	if len(elements) == 0 {
		elements = nil
	}
	{
		_numItems := uint16(utils.Max(noOfElements, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := ExtensionObjectDefinitionParseWithBuffer(arrayCtx, readBuffer, "539")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'elements' field of RelativePath")
			}
			elements[_curItem] = _item.(ExtensionObjectDefinition)
		}
	}
	if closeErr := readBuffer.CloseContext("elements", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for elements")
	}

	if closeErr := readBuffer.CloseContext("RelativePath"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RelativePath")
	}

	// Create a partially initialized instance
	_child := &_RelativePath{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		NoOfElements:               noOfElements,
		Elements:                   elements,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_RelativePath) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RelativePath) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RelativePath"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RelativePath")
		}

		// Simple Field (noOfElements)
		noOfElements := int32(m.GetNoOfElements())
		_noOfElementsErr := writeBuffer.WriteInt32("noOfElements", 32, (noOfElements))
		if _noOfElementsErr != nil {
			return errors.Wrap(_noOfElementsErr, "Error serializing 'noOfElements' field")
		}

		// Array Field (elements)
		if pushErr := writeBuffer.PushContext("elements", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for elements")
		}
		for _curItem, _element := range m.GetElements() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetElements()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'elements' field")
			}
		}
		if popErr := writeBuffer.PopContext("elements", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for elements")
		}

		if popErr := writeBuffer.PopContext("RelativePath"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RelativePath")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_RelativePath) isRelativePath() bool {
	return true
}

func (m *_RelativePath) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
