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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetLogDataLogData is the corresponding interface of BACnetLogDataLogData
type BACnetLogDataLogData interface {
	utils.LengthAware
	utils.Serializable
	BACnetLogData
	// GetInnerOpeningTag returns InnerOpeningTag (property field)
	GetInnerOpeningTag() BACnetOpeningTag
	// GetLogData returns LogData (property field)
	GetLogData() []BACnetLogDataLogDataEntry
	// GetInnerClosingTag returns InnerClosingTag (property field)
	GetInnerClosingTag() BACnetClosingTag
}

// BACnetLogDataLogDataExactly can be used when we want exactly this type and not a type which fulfills BACnetLogDataLogData.
// This is useful for switch cases.
type BACnetLogDataLogDataExactly interface {
	BACnetLogDataLogData
	isBACnetLogDataLogData() bool
}

// _BACnetLogDataLogData is the data-structure of this message
type _BACnetLogDataLogData struct {
	*_BACnetLogData
        InnerOpeningTag BACnetOpeningTag
        LogData []BACnetLogDataLogDataEntry
        InnerClosingTag BACnetClosingTag
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetLogDataLogData) InitializeParent(parent BACnetLogData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetLogDataLogData)  GetParent() BACnetLogData {
	return m._BACnetLogData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetLogDataLogData) GetInnerOpeningTag() BACnetOpeningTag {
	return m.InnerOpeningTag
}

func (m *_BACnetLogDataLogData) GetLogData() []BACnetLogDataLogDataEntry {
	return m.LogData
}

func (m *_BACnetLogDataLogData) GetInnerClosingTag() BACnetClosingTag {
	return m.InnerClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetLogDataLogData factory function for _BACnetLogDataLogData
func NewBACnetLogDataLogData( innerOpeningTag BACnetOpeningTag , logData []BACnetLogDataLogDataEntry , innerClosingTag BACnetClosingTag , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 ) *_BACnetLogDataLogData {
	_result := &_BACnetLogDataLogData{
		InnerOpeningTag: innerOpeningTag,
		LogData: logData,
		InnerClosingTag: innerClosingTag,
    	_BACnetLogData: NewBACnetLogData(openingTag, peekedTagHeader, closingTag, tagNumber),
	}
	_result._BACnetLogData._BACnetLogDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetLogDataLogData(structType interface{}) BACnetLogDataLogData {
    if casted, ok := structType.(BACnetLogDataLogData); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetLogDataLogData); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetLogDataLogData) GetTypeName() string {
	return "BACnetLogDataLogData"
}

func (m *_BACnetLogDataLogData) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetLogDataLogData) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (innerOpeningTag)
	lengthInBits += m.InnerOpeningTag.GetLengthInBits()

	// Array field
	if len(m.LogData) > 0 {
		for _, element := range m.LogData {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (innerClosingTag)
	lengthInBits += m.InnerClosingTag.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetLogDataLogData) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetLogDataLogDataParse(readBuffer utils.ReadBuffer, tagNumber uint8) (BACnetLogDataLogData, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetLogDataLogData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetLogDataLogData")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (innerOpeningTag)
	if pullErr := readBuffer.PullContext("innerOpeningTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerOpeningTag")
	}
_innerOpeningTag, _innerOpeningTagErr := BACnetOpeningTagParse(readBuffer , uint8( uint8(1) ) )
	if _innerOpeningTagErr != nil {
		return nil, errors.Wrap(_innerOpeningTagErr, "Error parsing 'innerOpeningTag' field of BACnetLogDataLogData")
	}
	innerOpeningTag := _innerOpeningTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("innerOpeningTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerOpeningTag")
	}

	// Array field (logData)
	if pullErr := readBuffer.PullContext("logData", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for logData")
	}
	// Terminated array
	var logData []BACnetLogDataLogDataEntry
	{
		for ;!bool(IsBACnetConstructedDataClosingTag(readBuffer, false, 1)); {
_item, _err := BACnetLogDataLogDataEntryParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'logData' field of BACnetLogDataLogData")
			}
			logData = append(logData, _item.(BACnetLogDataLogDataEntry))

		}
	}
	if closeErr := readBuffer.CloseContext("logData", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for logData")
	}

	// Simple Field (innerClosingTag)
	if pullErr := readBuffer.PullContext("innerClosingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for innerClosingTag")
	}
_innerClosingTag, _innerClosingTagErr := BACnetClosingTagParse(readBuffer , uint8( uint8(1) ) )
	if _innerClosingTagErr != nil {
		return nil, errors.Wrap(_innerClosingTagErr, "Error parsing 'innerClosingTag' field of BACnetLogDataLogData")
	}
	innerClosingTag := _innerClosingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("innerClosingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for innerClosingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetLogDataLogData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetLogDataLogData")
	}

	// Create a partially initialized instance
	_child := &_BACnetLogDataLogData{
		_BACnetLogData: &_BACnetLogData{
			TagNumber: tagNumber,
		},
		InnerOpeningTag: innerOpeningTag,
		LogData: logData,
		InnerClosingTag: innerClosingTag,
	}
	_child._BACnetLogData._BACnetLogDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetLogDataLogData) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetLogDataLogData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetLogDataLogData")
		}

	// Simple Field (innerOpeningTag)
	if pushErr := writeBuffer.PushContext("innerOpeningTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for innerOpeningTag")
	}
	_innerOpeningTagErr := writeBuffer.WriteSerializable(m.GetInnerOpeningTag())
	if popErr := writeBuffer.PopContext("innerOpeningTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for innerOpeningTag")
	}
	if _innerOpeningTagErr != nil {
		return errors.Wrap(_innerOpeningTagErr, "Error serializing 'innerOpeningTag' field")
	}

	// Array Field (logData)
	if pushErr := writeBuffer.PushContext("logData", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for logData")
	}
	for _, _element := range m.GetLogData() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'logData' field")
		}
	}
	if popErr := writeBuffer.PopContext("logData", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for logData")
	}

	// Simple Field (innerClosingTag)
	if pushErr := writeBuffer.PushContext("innerClosingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for innerClosingTag")
	}
	_innerClosingTagErr := writeBuffer.WriteSerializable(m.GetInnerClosingTag())
	if popErr := writeBuffer.PopContext("innerClosingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for innerClosingTag")
	}
	if _innerClosingTagErr != nil {
		return errors.Wrap(_innerClosingTagErr, "Error serializing 'innerClosingTag' field")
	}

		if popErr := writeBuffer.PopContext("BACnetLogDataLogData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetLogDataLogData")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetLogDataLogData) isBACnetLogDataLogData() bool {
	return true
}

func (m *_BACnetLogDataLogData) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



