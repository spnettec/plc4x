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


// BACnetConstructedDataListOfObjectPropertyReferences is the corresponding interface of BACnetConstructedDataListOfObjectPropertyReferences
type BACnetConstructedDataListOfObjectPropertyReferences interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetReferences returns References (property field)
	GetReferences() []BACnetDeviceObjectPropertyReference
}

// BACnetConstructedDataListOfObjectPropertyReferencesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataListOfObjectPropertyReferences.
// This is useful for switch cases.
type BACnetConstructedDataListOfObjectPropertyReferencesExactly interface {
	BACnetConstructedDataListOfObjectPropertyReferences
	isBACnetConstructedDataListOfObjectPropertyReferences() bool
}

// _BACnetConstructedDataListOfObjectPropertyReferences is the data-structure of this message
type _BACnetConstructedDataListOfObjectPropertyReferences struct {
	*_BACnetConstructedData
        References []BACnetDeviceObjectPropertyReference
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataListOfObjectPropertyReferences)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataListOfObjectPropertyReferences)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_LIST_OF_OBJECT_PROPERTY_REFERENCES}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataListOfObjectPropertyReferences) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataListOfObjectPropertyReferences)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataListOfObjectPropertyReferences) GetReferences() []BACnetDeviceObjectPropertyReference {
	return m.References
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataListOfObjectPropertyReferences factory function for _BACnetConstructedDataListOfObjectPropertyReferences
func NewBACnetConstructedDataListOfObjectPropertyReferences( references []BACnetDeviceObjectPropertyReference , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataListOfObjectPropertyReferences {
	_result := &_BACnetConstructedDataListOfObjectPropertyReferences{
		References: references,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataListOfObjectPropertyReferences(structType interface{}) BACnetConstructedDataListOfObjectPropertyReferences {
    if casted, ok := structType.(BACnetConstructedDataListOfObjectPropertyReferences); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataListOfObjectPropertyReferences); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataListOfObjectPropertyReferences) GetTypeName() string {
	return "BACnetConstructedDataListOfObjectPropertyReferences"
}

func (m *_BACnetConstructedDataListOfObjectPropertyReferences) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataListOfObjectPropertyReferences) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.References) > 0 {
		for _, element := range m.References {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}


func (m *_BACnetConstructedDataListOfObjectPropertyReferences) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataListOfObjectPropertyReferencesParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataListOfObjectPropertyReferences, error) {
	return BACnetConstructedDataListOfObjectPropertyReferencesParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataListOfObjectPropertyReferencesParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataListOfObjectPropertyReferences, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataListOfObjectPropertyReferences"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataListOfObjectPropertyReferences")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (references)
	if pullErr := readBuffer.PullContext("references", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for references")
	}
	// Terminated array
	var references []BACnetDeviceObjectPropertyReference
	{
		for ;!bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)); {
_item, _err := BACnetDeviceObjectPropertyReferenceParseWithBuffer(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'references' field of BACnetConstructedDataListOfObjectPropertyReferences")
			}
			references = append(references, _item.(BACnetDeviceObjectPropertyReference))
		}
	}
	if closeErr := readBuffer.CloseContext("references", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for references")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataListOfObjectPropertyReferences"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataListOfObjectPropertyReferences")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataListOfObjectPropertyReferences{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		References: references,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataListOfObjectPropertyReferences) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataListOfObjectPropertyReferences) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataListOfObjectPropertyReferences"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataListOfObjectPropertyReferences")
		}

	// Array Field (references)
	if pushErr := writeBuffer.PushContext("references", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for references")
	}
	for _, _element := range m.GetReferences() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'references' field")
		}
	}
	if popErr := writeBuffer.PopContext("references", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for references")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataListOfObjectPropertyReferences"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataListOfObjectPropertyReferences")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataListOfObjectPropertyReferences) isBACnetConstructedDataListOfObjectPropertyReferences() bool {
	return true
}

func (m *_BACnetConstructedDataListOfObjectPropertyReferences) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



