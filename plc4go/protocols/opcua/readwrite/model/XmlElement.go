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


// XmlElement is the corresponding interface of XmlElement
type XmlElement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetLength returns Length (property field)
	GetLength() int32
	// GetValue returns Value (property field)
	GetValue() []string
}

// XmlElementExactly can be used when we want exactly this type and not a type which fulfills XmlElement.
// This is useful for switch cases.
type XmlElementExactly interface {
	XmlElement
	isXmlElement() bool
}

// _XmlElement is the data-structure of this message
type _XmlElement struct {
        Length int32
        Value []string
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_XmlElement) GetLength() int32 {
	return m.Length
}

func (m *_XmlElement) GetValue() []string {
	return m.Value
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewXmlElement factory function for _XmlElement
func NewXmlElement( length int32 , value []string ) *_XmlElement {
return &_XmlElement{ Length: length , Value: value }
}

// Deprecated: use the interface for direct cast
func CastXmlElement(structType any) XmlElement {
    if casted, ok := structType.(XmlElement); ok {
		return casted
	}
	if casted, ok := structType.(*XmlElement); ok {
		return *casted
	}
	return nil
}

func (m *_XmlElement) GetTypeName() string {
	return "XmlElement"
}

func (m *_XmlElement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (length)
	lengthInBits += 32;

	// Array field
	if len(m.Value) > 0 {
		lengthInBits += 8 * uint16(len(m.Value))
	}

	return lengthInBits
}


func (m *_XmlElement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func XmlElementParse(ctx context.Context, theBytes []byte) (XmlElement, error) {
	return XmlElementParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func XmlElementParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (XmlElement, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("XmlElement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for XmlElement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (length)
_length, _lengthErr := readBuffer.ReadInt32("length", 32)
	if _lengthErr != nil {
		return nil, errors.Wrap(_lengthErr, "Error parsing 'length' field of XmlElement")
	}
	length := _length

	// Array field (value)
	if pullErr := readBuffer.PullContext("value", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for value")
	}
	// Count array
	value := make([]string, utils.Max(length, 0))
	// This happens when the size is set conditional to 0
	if len(value) == 0 {
		value = nil
	}
	{
		_numItems := uint16(utils.Max(length, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
_item, _err := readBuffer.ReadString("", uint32(8), "UTF-8")
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'value' field of XmlElement")
			}
			value[_curItem] = _item
		}
	}
	if closeErr := readBuffer.CloseContext("value", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for value")
	}

	if closeErr := readBuffer.CloseContext("XmlElement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for XmlElement")
	}

	// Create the instance
	return &_XmlElement{
			Length: length,
			Value: value,
		}, nil
}

func (m *_XmlElement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_XmlElement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("XmlElement"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for XmlElement")
	}

	// Simple Field (length)
	length := int32(m.GetLength())
	_lengthErr := writeBuffer.WriteInt32("length", 32, (length))
	if _lengthErr != nil {
		return errors.Wrap(_lengthErr, "Error serializing 'length' field")
	}

	// Array Field (value)
	if pushErr := writeBuffer.PushContext("value", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for value")
	}
	for _curItem, _element := range m.GetValue() {
		_ = _curItem
		_elementErr := writeBuffer.WriteString("", uint32(8), "UTF-8", _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'value' field")
		}
	}
	if popErr := writeBuffer.PopContext("value", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for value")
	}

	if popErr := writeBuffer.PopContext("XmlElement"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for XmlElement")
	}
	return nil
}


func (m *_XmlElement) isXmlElement() bool {
	return true
}

func (m *_XmlElement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



