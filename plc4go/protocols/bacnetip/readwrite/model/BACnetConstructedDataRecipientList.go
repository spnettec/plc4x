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


// BACnetConstructedDataRecipientList is the corresponding interface of BACnetConstructedDataRecipientList
type BACnetConstructedDataRecipientList interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetRecipientList returns RecipientList (property field)
	GetRecipientList() []BACnetDestination
}

// BACnetConstructedDataRecipientListExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataRecipientList.
// This is useful for switch cases.
type BACnetConstructedDataRecipientListExactly interface {
	BACnetConstructedDataRecipientList
	isBACnetConstructedDataRecipientList() bool
}

// _BACnetConstructedDataRecipientList is the data-structure of this message
type _BACnetConstructedDataRecipientList struct {
	*_BACnetConstructedData
        RecipientList []BACnetDestination
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRecipientList)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataRecipientList)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_RECIPIENT_LIST}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRecipientList) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataRecipientList)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRecipientList) GetRecipientList() []BACnetDestination {
	return m.RecipientList
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataRecipientList factory function for _BACnetConstructedDataRecipientList
func NewBACnetConstructedDataRecipientList( recipientList []BACnetDestination , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataRecipientList {
	_result := &_BACnetConstructedDataRecipientList{
		RecipientList: recipientList,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRecipientList(structType interface{}) BACnetConstructedDataRecipientList {
    if casted, ok := structType.(BACnetConstructedDataRecipientList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRecipientList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRecipientList) GetTypeName() string {
	return "BACnetConstructedDataRecipientList"
}

func (m *_BACnetConstructedDataRecipientList) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataRecipientList) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.RecipientList) > 0 {
		for _, element := range m.RecipientList {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}


func (m *_BACnetConstructedDataRecipientList) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataRecipientListParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataRecipientList, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRecipientList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRecipientList")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (recipientList)
	if pullErr := readBuffer.PullContext("recipientList", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for recipientList")
	}
	// Terminated array
	var recipientList []BACnetDestination
	{
		for ;!bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)); {
_item, _err := BACnetDestinationParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'recipientList' field of BACnetConstructedDataRecipientList")
			}
			recipientList = append(recipientList, _item.(BACnetDestination))

		}
	}
	if closeErr := readBuffer.CloseContext("recipientList", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for recipientList")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRecipientList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRecipientList")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataRecipientList{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		RecipientList: recipientList,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataRecipientList) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRecipientList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRecipientList")
		}

	// Array Field (recipientList)
	if pushErr := writeBuffer.PushContext("recipientList", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for recipientList")
	}
	for _, _element := range m.GetRecipientList() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'recipientList' field")
		}
	}
	if popErr := writeBuffer.PopContext("recipientList", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for recipientList")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRecipientList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRecipientList")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataRecipientList) isBACnetConstructedDataRecipientList() bool {
	return true
}

func (m *_BACnetConstructedDataRecipientList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



