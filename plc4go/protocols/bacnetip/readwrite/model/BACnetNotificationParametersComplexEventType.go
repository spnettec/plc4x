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


// BACnetNotificationParametersComplexEventType is the corresponding interface of BACnetNotificationParametersComplexEventType
type BACnetNotificationParametersComplexEventType interface {
	utils.LengthAware
	utils.Serializable
	BACnetNotificationParameters
	// GetListOfValues returns ListOfValues (property field)
	GetListOfValues() BACnetPropertyValues
}

// BACnetNotificationParametersComplexEventTypeExactly can be used when we want exactly this type and not a type which fulfills BACnetNotificationParametersComplexEventType.
// This is useful for switch cases.
type BACnetNotificationParametersComplexEventTypeExactly interface {
	BACnetNotificationParametersComplexEventType
	isBACnetNotificationParametersComplexEventType() bool
}

// _BACnetNotificationParametersComplexEventType is the data-structure of this message
type _BACnetNotificationParametersComplexEventType struct {
	*_BACnetNotificationParameters
        ListOfValues BACnetPropertyValues
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetNotificationParametersComplexEventType) InitializeParent(parent BACnetNotificationParameters , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetNotificationParametersComplexEventType)  GetParent() BACnetNotificationParameters {
	return m._BACnetNotificationParameters
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetNotificationParametersComplexEventType) GetListOfValues() BACnetPropertyValues {
	return m.ListOfValues
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetNotificationParametersComplexEventType factory function for _BACnetNotificationParametersComplexEventType
func NewBACnetNotificationParametersComplexEventType( listOfValues BACnetPropertyValues , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , objectTypeArgument BACnetObjectType ) *_BACnetNotificationParametersComplexEventType {
	_result := &_BACnetNotificationParametersComplexEventType{
		ListOfValues: listOfValues,
    	_BACnetNotificationParameters: NewBACnetNotificationParameters(openingTag, peekedTagHeader, closingTag, tagNumber, objectTypeArgument),
	}
	_result._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetNotificationParametersComplexEventType(structType interface{}) BACnetNotificationParametersComplexEventType {
    if casted, ok := structType.(BACnetNotificationParametersComplexEventType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetNotificationParametersComplexEventType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetNotificationParametersComplexEventType) GetTypeName() string {
	return "BACnetNotificationParametersComplexEventType"
}

func (m *_BACnetNotificationParametersComplexEventType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetNotificationParametersComplexEventType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (listOfValues)
	lengthInBits += m.ListOfValues.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetNotificationParametersComplexEventType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetNotificationParametersComplexEventTypeParse(theBytes []byte, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParametersComplexEventType, error) {
	return BACnetNotificationParametersComplexEventTypeParseWithBuffer(utils.NewReadBufferByteBased(theBytes), peekedTagNumber, tagNumber, objectTypeArgument)
}

func BACnetNotificationParametersComplexEventTypeParseWithBuffer(readBuffer utils.ReadBuffer, peekedTagNumber uint8, tagNumber uint8, objectTypeArgument BACnetObjectType) (BACnetNotificationParametersComplexEventType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetNotificationParametersComplexEventType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetNotificationParametersComplexEventType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (listOfValues)
	if pullErr := readBuffer.PullContext("listOfValues"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for listOfValues")
	}
_listOfValues, _listOfValuesErr := BACnetPropertyValuesParseWithBuffer(readBuffer , uint8( peekedTagNumber ) , BACnetObjectType( objectTypeArgument ) )
	if _listOfValuesErr != nil {
		return nil, errors.Wrap(_listOfValuesErr, "Error parsing 'listOfValues' field of BACnetNotificationParametersComplexEventType")
	}
	listOfValues := _listOfValues.(BACnetPropertyValues)
	if closeErr := readBuffer.CloseContext("listOfValues"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for listOfValues")
	}

	if closeErr := readBuffer.CloseContext("BACnetNotificationParametersComplexEventType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetNotificationParametersComplexEventType")
	}

	// Create a partially initialized instance
	_child := &_BACnetNotificationParametersComplexEventType{
		_BACnetNotificationParameters: &_BACnetNotificationParameters{
			TagNumber: tagNumber,
			ObjectTypeArgument: objectTypeArgument,
		},
		ListOfValues: listOfValues,
	}
	_child._BACnetNotificationParameters._BACnetNotificationParametersChildRequirements = _child
	return _child, nil
}

func (m *_BACnetNotificationParametersComplexEventType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetNotificationParametersComplexEventType) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetNotificationParametersComplexEventType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetNotificationParametersComplexEventType")
		}

	// Simple Field (listOfValues)
	if pushErr := writeBuffer.PushContext("listOfValues"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for listOfValues")
	}
	_listOfValuesErr := writeBuffer.WriteSerializable(m.GetListOfValues())
	if popErr := writeBuffer.PopContext("listOfValues"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for listOfValues")
	}
	if _listOfValuesErr != nil {
		return errors.Wrap(_listOfValuesErr, "Error serializing 'listOfValues' field")
	}

		if popErr := writeBuffer.PopContext("BACnetNotificationParametersComplexEventType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetNotificationParametersComplexEventType")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetNotificationParametersComplexEventType) isBACnetNotificationParametersComplexEventType() bool {
	return true
}

func (m *_BACnetNotificationParametersComplexEventType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



