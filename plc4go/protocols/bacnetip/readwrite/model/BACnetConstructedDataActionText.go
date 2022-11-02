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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetConstructedDataActionText is the corresponding interface of BACnetConstructedDataActionText
type BACnetConstructedDataActionText interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetActionText returns ActionText (property field)
	GetActionText() []BACnetApplicationTagCharacterString
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataActionTextExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataActionText.
// This is useful for switch cases.
type BACnetConstructedDataActionTextExactly interface {
	BACnetConstructedDataActionText
	isBACnetConstructedDataActionText() bool
}

// _BACnetConstructedDataActionText is the data-structure of this message
type _BACnetConstructedDataActionText struct {
	*_BACnetConstructedData
        NumberOfDataElements BACnetApplicationTagUnsignedInteger
        ActionText []BACnetApplicationTagCharacterString
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataActionText)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataActionText)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_ACTION_TEXT}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataActionText) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataActionText)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataActionText) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataActionText) GetActionText() []BACnetApplicationTagCharacterString {
	return m.ActionText
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataActionText) GetZero() uint64 {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataActionText factory function for _BACnetConstructedDataActionText
func NewBACnetConstructedDataActionText( numberOfDataElements BACnetApplicationTagUnsignedInteger , actionText []BACnetApplicationTagCharacterString , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataActionText {
	_result := &_BACnetConstructedDataActionText{
		NumberOfDataElements: numberOfDataElements,
		ActionText: actionText,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataActionText(structType interface{}) BACnetConstructedDataActionText {
    if casted, ok := structType.(BACnetConstructedDataActionText); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataActionText); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataActionText) GetTypeName() string {
	return "BACnetConstructedDataActionText"
}

func (m *_BACnetConstructedDataActionText) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataActionText) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits()
	}

	// Array field
	if len(m.ActionText) > 0 {
		for _, element := range m.ActionText {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}


func (m *_BACnetConstructedDataActionText) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataActionTextParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataActionText, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataActionText"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataActionText")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field of BACnetConstructedDataActionText")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (actionText)
	if pullErr := readBuffer.PullContext("actionText", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for actionText")
	}
	// Terminated array
	var actionText []BACnetApplicationTagCharacterString
	{
		for ;!bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)); {
_item, _err := BACnetApplicationTagParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'actionText' field of BACnetConstructedDataActionText")
			}
			actionText = append(actionText, _item.(BACnetApplicationTagCharacterString))

		}
	}
	if closeErr := readBuffer.CloseContext("actionText", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for actionText")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataActionText"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataActionText")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataActionText{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NumberOfDataElements: numberOfDataElements,
		ActionText: actionText,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataActionText) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataActionText) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataActionText"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataActionText")
		}
	// Virtual field
	if _zeroErr := writeBuffer.WriteVirtual("zero", m.GetZero()); _zeroErr != nil {
		return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
	}

	// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if m.GetNumberOfDataElements() != nil {
		if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
		}
		numberOfDataElements = m.GetNumberOfDataElements()
		_numberOfDataElementsErr := writeBuffer.WriteSerializable(numberOfDataElements)
		if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for numberOfDataElements")
		}
		if _numberOfDataElementsErr != nil {
			return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
		}
	}

	// Array Field (actionText)
	if pushErr := writeBuffer.PushContext("actionText", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for actionText")
	}
	for _, _element := range m.GetActionText() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'actionText' field")
		}
	}
	if popErr := writeBuffer.PopContext("actionText", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for actionText")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataActionText"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataActionText")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataActionText) isBACnetConstructedDataActionText() bool {
	return true
}

func (m *_BACnetConstructedDataActionText) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



