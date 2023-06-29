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

// CipSecurityInformation is the corresponding interface of CipSecurityInformation
type CipSecurityInformation interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CommandSpecificDataItem
	// GetTodoImplement returns TodoImplement (property field)
	GetTodoImplement() []uint8
}

// CipSecurityInformationExactly can be used when we want exactly this type and not a type which fulfills CipSecurityInformation.
// This is useful for switch cases.
type CipSecurityInformationExactly interface {
	CipSecurityInformation
	isCipSecurityInformation() bool
}

// _CipSecurityInformation is the data-structure of this message
type _CipSecurityInformation struct {
	*_CommandSpecificDataItem
	TodoImplement []uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CipSecurityInformation) GetItemType() uint16 {
	return 0x0086
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CipSecurityInformation) InitializeParent(parent CommandSpecificDataItem) {}

func (m *_CipSecurityInformation) GetParent() CommandSpecificDataItem {
	return m._CommandSpecificDataItem
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CipSecurityInformation) GetTodoImplement() []uint8 {
	return m.TodoImplement
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCipSecurityInformation factory function for _CipSecurityInformation
func NewCipSecurityInformation(todoImplement []uint8) *_CipSecurityInformation {
	_result := &_CipSecurityInformation{
		TodoImplement:            todoImplement,
		_CommandSpecificDataItem: NewCommandSpecificDataItem(),
	}
	_result._CommandSpecificDataItem._CommandSpecificDataItemChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCipSecurityInformation(structType any) CipSecurityInformation {
	if casted, ok := structType.(CipSecurityInformation); ok {
		return casted
	}
	if casted, ok := structType.(*CipSecurityInformation); ok {
		return *casted
	}
	return nil
}

func (m *_CipSecurityInformation) GetTypeName() string {
	return "CipSecurityInformation"
}

func (m *_CipSecurityInformation) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (itemLength)
	lengthInBits += 16

	// Array field
	if len(m.TodoImplement) > 0 {
		lengthInBits += 8 * uint16(len(m.TodoImplement))
	}

	return lengthInBits
}

func (m *_CipSecurityInformation) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CipSecurityInformationParse(ctx context.Context, theBytes []byte) (CipSecurityInformation, error) {
	return CipSecurityInformationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func CipSecurityInformationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (CipSecurityInformation, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CipSecurityInformation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CipSecurityInformation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (itemLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	itemLength, _itemLengthErr := readBuffer.ReadUint16("itemLength", 16)
	_ = itemLength
	if _itemLengthErr != nil {
		return nil, errors.Wrap(_itemLengthErr, "Error parsing 'itemLength' field of CipSecurityInformation")
	}

	// Array field (todoImplement)
	if pullErr := readBuffer.PullContext("todoImplement", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for todoImplement")
	}
	// Count array
	todoImplement := make([]uint8, itemLength)
	// This happens when the size is set conditional to 0
	if len(todoImplement) == 0 {
		todoImplement = nil
	}
	{
		_numItems := uint16(itemLength)
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := readBuffer.ReadUint8("", 8)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'todoImplement' field of CipSecurityInformation")
			}
			todoImplement[_curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("todoImplement", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for todoImplement")
	}

	if closeErr := readBuffer.CloseContext("CipSecurityInformation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CipSecurityInformation")
	}

	// Create a partially initialized instance
	_child := &_CipSecurityInformation{
		_CommandSpecificDataItem: &_CommandSpecificDataItem{},
		TodoImplement:            todoImplement,
	}
	_child._CommandSpecificDataItem._CommandSpecificDataItemChildRequirements = _child
	return _child, nil
}

func (m *_CipSecurityInformation) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CipSecurityInformation) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CipSecurityInformation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CipSecurityInformation")
		}

		// Implicit Field (itemLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
		itemLength := uint16(uint16(len(m.GetTodoImplement())))
		_itemLengthErr := writeBuffer.WriteUint16("itemLength", 16, (itemLength))
		if _itemLengthErr != nil {
			return errors.Wrap(_itemLengthErr, "Error serializing 'itemLength' field")
		}

		// Array Field (todoImplement)
		if pushErr := writeBuffer.PushContext("todoImplement", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for todoImplement")
		}
		for _curItem, _element := range m.GetTodoImplement() {
			_ = _curItem
			_elementErr := writeBuffer.WriteUint8("", 8, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'todoImplement' field")
			}
		}
		if popErr := writeBuffer.PopContext("todoImplement", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for todoImplement")
		}

		if popErr := writeBuffer.PopContext("CipSecurityInformation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CipSecurityInformation")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_CipSecurityInformation) isCipSecurityInformation() bool {
	return true
}

func (m *_CipSecurityInformation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
