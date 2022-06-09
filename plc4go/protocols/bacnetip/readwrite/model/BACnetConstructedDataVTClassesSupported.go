/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataVTClassesSupported is the data-structure of this message
type BACnetConstructedDataVTClassesSupported struct {
	*BACnetConstructedData
	VtClassesSupported []*BACnetVTClassTagged

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataVTClassesSupported is the corresponding interface of BACnetConstructedDataVTClassesSupported
type IBACnetConstructedDataVTClassesSupported interface {
	IBACnetConstructedData
	// GetVtClassesSupported returns VtClassesSupported (property field)
	GetVtClassesSupported() []*BACnetVTClassTagged
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataVTClassesSupported) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataVTClassesSupported) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_VT_CLASSES_SUPPORTED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataVTClassesSupported) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataVTClassesSupported) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataVTClassesSupported) GetVtClassesSupported() []*BACnetVTClassTagged {
	return m.VtClassesSupported
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataVTClassesSupported factory function for BACnetConstructedDataVTClassesSupported
func NewBACnetConstructedDataVTClassesSupported(vtClassesSupported []*BACnetVTClassTagged, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataVTClassesSupported {
	_result := &BACnetConstructedDataVTClassesSupported{
		VtClassesSupported:    vtClassesSupported,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataVTClassesSupported(structType interface{}) *BACnetConstructedDataVTClassesSupported {
	if casted, ok := structType.(BACnetConstructedDataVTClassesSupported); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataVTClassesSupported); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataVTClassesSupported(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataVTClassesSupported(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataVTClassesSupported) GetTypeName() string {
	return "BACnetConstructedDataVTClassesSupported"
}

func (m *BACnetConstructedDataVTClassesSupported) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataVTClassesSupported) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.VtClassesSupported) > 0 {
		for _, element := range m.VtClassesSupported {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *BACnetConstructedDataVTClassesSupported) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataVTClassesSupportedParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataVTClassesSupported, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataVTClassesSupported"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (vtClassesSupported)
	if pullErr := readBuffer.PullContext("vtClassesSupported", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, pullErr
	}
	// Terminated array
	vtClassesSupported := make([]*BACnetVTClassTagged, 0)
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetVTClassTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'vtClassesSupported' field")
			}
			vtClassesSupported = append(vtClassesSupported, CastBACnetVTClassTagged(_item))

		}
	}
	if closeErr := readBuffer.CloseContext("vtClassesSupported", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataVTClassesSupported"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataVTClassesSupported{
		VtClassesSupported:    vtClassesSupported,
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataVTClassesSupported) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataVTClassesSupported"); pushErr != nil {
			return pushErr
		}

		// Array Field (vtClassesSupported)
		if m.VtClassesSupported != nil {
			if pushErr := writeBuffer.PushContext("vtClassesSupported", utils.WithRenderAsList(true)); pushErr != nil {
				return pushErr
			}
			for _, _element := range m.VtClassesSupported {
				_elementErr := _element.Serialize(writeBuffer)
				if _elementErr != nil {
					return errors.Wrap(_elementErr, "Error serializing 'vtClassesSupported' field")
				}
			}
			if popErr := writeBuffer.PopContext("vtClassesSupported", utils.WithRenderAsList(true)); popErr != nil {
				return popErr
			}
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataVTClassesSupported"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataVTClassesSupported) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
