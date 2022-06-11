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

// BACnetConstructedDataMinimumOutput is the data-structure of this message
type BACnetConstructedDataMinimumOutput struct {
	*BACnetConstructedData
	MinimumOutput *BACnetApplicationTagReal

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataMinimumOutput is the corresponding interface of BACnetConstructedDataMinimumOutput
type IBACnetConstructedDataMinimumOutput interface {
	IBACnetConstructedData
	// GetMinimumOutput returns MinimumOutput (property field)
	GetMinimumOutput() *BACnetApplicationTagReal
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

func (m *BACnetConstructedDataMinimumOutput) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataMinimumOutput) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MINIMUM_OUTPUT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataMinimumOutput) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataMinimumOutput) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataMinimumOutput) GetMinimumOutput() *BACnetApplicationTagReal {
	return m.MinimumOutput
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMinimumOutput factory function for BACnetConstructedDataMinimumOutput
func NewBACnetConstructedDataMinimumOutput(minimumOutput *BACnetApplicationTagReal, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataMinimumOutput {
	_result := &BACnetConstructedDataMinimumOutput{
		MinimumOutput:         minimumOutput,
		BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataMinimumOutput(structType interface{}) *BACnetConstructedDataMinimumOutput {
	if casted, ok := structType.(BACnetConstructedDataMinimumOutput); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMinimumOutput); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataMinimumOutput(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataMinimumOutput(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataMinimumOutput) GetTypeName() string {
	return "BACnetConstructedDataMinimumOutput"
}

func (m *BACnetConstructedDataMinimumOutput) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataMinimumOutput) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (minimumOutput)
	lengthInBits += m.MinimumOutput.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataMinimumOutput) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMinimumOutputParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataMinimumOutput, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMinimumOutput"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMinimumOutput")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (minimumOutput)
	if pullErr := readBuffer.PullContext("minimumOutput"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for minimumOutput")
	}
	_minimumOutput, _minimumOutputErr := BACnetApplicationTagParse(readBuffer)
	if _minimumOutputErr != nil {
		return nil, errors.Wrap(_minimumOutputErr, "Error parsing 'minimumOutput' field")
	}
	minimumOutput := CastBACnetApplicationTagReal(_minimumOutput)
	if closeErr := readBuffer.CloseContext("minimumOutput"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for minimumOutput")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMinimumOutput"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMinimumOutput")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataMinimumOutput{
		MinimumOutput:         CastBACnetApplicationTagReal(minimumOutput),
		BACnetConstructedData: &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataMinimumOutput) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMinimumOutput"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMinimumOutput")
		}

		// Simple Field (minimumOutput)
		if pushErr := writeBuffer.PushContext("minimumOutput"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for minimumOutput")
		}
		_minimumOutputErr := m.MinimumOutput.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("minimumOutput"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for minimumOutput")
		}
		if _minimumOutputErr != nil {
			return errors.Wrap(_minimumOutputErr, "Error serializing 'minimumOutput' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMinimumOutput"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMinimumOutput")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataMinimumOutput) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
