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


// VariantExpandedNodeId is the corresponding interface of VariantExpandedNodeId
type VariantExpandedNodeId interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	Variant
	// GetArrayLength returns ArrayLength (property field)
	GetArrayLength() *int32
	// GetValue returns Value (property field)
	GetValue() []ExpandedNodeId
}

// VariantExpandedNodeIdExactly can be used when we want exactly this type and not a type which fulfills VariantExpandedNodeId.
// This is useful for switch cases.
type VariantExpandedNodeIdExactly interface {
	VariantExpandedNodeId
	isVariantExpandedNodeId() bool
}

// _VariantExpandedNodeId is the data-structure of this message
type _VariantExpandedNodeId struct {
	*_Variant
        ArrayLength *int32
        Value []ExpandedNodeId
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_VariantExpandedNodeId)  GetVariantType() uint8 {
return uint8(18)}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_VariantExpandedNodeId) InitializeParent(parent Variant , arrayLengthSpecified bool , arrayDimensionsSpecified bool , noOfArrayDimensions * int32 , arrayDimensions []bool ) {	m.ArrayLengthSpecified = arrayLengthSpecified
	m.ArrayDimensionsSpecified = arrayDimensionsSpecified
	m.NoOfArrayDimensions = noOfArrayDimensions
	m.ArrayDimensions = arrayDimensions
}

func (m *_VariantExpandedNodeId)  GetParent() Variant {
	return m._Variant
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_VariantExpandedNodeId) GetArrayLength() *int32 {
	return m.ArrayLength
}

func (m *_VariantExpandedNodeId) GetValue() []ExpandedNodeId {
	return m.Value
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewVariantExpandedNodeId factory function for _VariantExpandedNodeId
func NewVariantExpandedNodeId( arrayLength *int32 , value []ExpandedNodeId , arrayLengthSpecified bool , arrayDimensionsSpecified bool , noOfArrayDimensions *int32 , arrayDimensions []bool ) *_VariantExpandedNodeId {
	_result := &_VariantExpandedNodeId{
		ArrayLength: arrayLength,
		Value: value,
    	_Variant: NewVariant(arrayLengthSpecified, arrayDimensionsSpecified, noOfArrayDimensions, arrayDimensions),
	}
	_result._Variant._VariantChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastVariantExpandedNodeId(structType any) VariantExpandedNodeId {
    if casted, ok := structType.(VariantExpandedNodeId); ok {
		return casted
	}
	if casted, ok := structType.(*VariantExpandedNodeId); ok {
		return *casted
	}
	return nil
}

func (m *_VariantExpandedNodeId) GetTypeName() string {
	return "VariantExpandedNodeId"
}

func (m *_VariantExpandedNodeId) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Optional Field (arrayLength)
	if m.ArrayLength != nil {
		lengthInBits += 32
	}

	// Array field
	if len(m.Value) > 0 {
		for _curItem, element := range m.Value {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.Value), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{GetLengthInBits(context.Context) uint16}).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}


func (m *_VariantExpandedNodeId) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func VariantExpandedNodeIdParse(ctx context.Context, theBytes []byte, arrayLengthSpecified bool) (VariantExpandedNodeId, error) {
	return VariantExpandedNodeIdParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), arrayLengthSpecified)
}

func VariantExpandedNodeIdParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, arrayLengthSpecified bool) (VariantExpandedNodeId, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("VariantExpandedNodeId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for VariantExpandedNodeId")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Optional Field (arrayLength) (Can be skipped, if a given expression evaluates to false)
	var arrayLength *int32 = nil
	if arrayLengthSpecified {
		_val, _err := readBuffer.ReadInt32("arrayLength", 32)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'arrayLength' field of VariantExpandedNodeId")
		}
		arrayLength = &_val
	}

	// Array field (value)
	if pullErr := readBuffer.PullContext("value", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
	// Count array
	value := make([]ExpandedNodeId, utils.Max(utils.InlineIf(bool(((arrayLength)) == (nil)), func() any {return uint16(uint16(1))}, func() any {return uint16((*arrayLength))}).(uint16), 0))
	// This happens when the size is set conditional to 0
	if len(value) == 0 {
		value = nil
	}
	{
		_numItems := uint16(utils.Max(utils.InlineIf(bool(((arrayLength)) == (nil)), func() any {return uint16(uint16(1))}, func() any {return uint16((*arrayLength))}).(uint16), 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := ExpandedNodeIdParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'value' field of VariantExpandedNodeId")
			}
			value[_curItem] = _item.(ExpandedNodeId)
		}
	}
	if closeErr := readBuffer.CloseContext("value", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}

	if closeErr := readBuffer.CloseContext("VariantExpandedNodeId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for VariantExpandedNodeId")
	}

	// Create a partially initialized instance
	_child := &_VariantExpandedNodeId{
		_Variant: &_Variant{
		},
		ArrayLength: arrayLength,
		Value: value,
	}
	_child._Variant._VariantChildRequirements = _child
	return _child, nil
}

func (m *_VariantExpandedNodeId) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_VariantExpandedNodeId) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("VariantExpandedNodeId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for VariantExpandedNodeId")
		}

	// Optional Field (arrayLength) (Can be skipped, if the value is null)
	var arrayLength *int32 = nil
	if m.GetArrayLength() != nil {
		arrayLength = m.GetArrayLength()
		_arrayLengthErr := writeBuffer.WriteInt32("arrayLength", 32, *(arrayLength))
		if _arrayLengthErr != nil {
			return errors.Wrap(_arrayLengthErr, "Error serializing 'arrayLength' field")
		}
	}

	// Array Field (value)
	if pushErr := writeBuffer.PushContext("value", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for value")
	}
	for _curItem, _element := range m.GetValue() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetValue()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'value' field")
		}
	}
	if popErr := writeBuffer.PopContext("value", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for value")
	}

		if popErr := writeBuffer.PopContext("VariantExpandedNodeId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for VariantExpandedNodeId")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_VariantExpandedNodeId) isVariantExpandedNodeId() bool {
	return true
}

func (m *_VariantExpandedNodeId) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



