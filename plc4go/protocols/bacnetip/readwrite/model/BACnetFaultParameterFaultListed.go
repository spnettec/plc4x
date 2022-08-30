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


// BACnetFaultParameterFaultListed is the corresponding interface of BACnetFaultParameterFaultListed
type BACnetFaultParameterFaultListed interface {
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameter
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetFaultListReference returns FaultListReference (property field)
	GetFaultListReference() BACnetDeviceObjectPropertyReferenceEnclosed
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// BACnetFaultParameterFaultListedExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultListed.
// This is useful for switch cases.
type BACnetFaultParameterFaultListedExactly interface {
	BACnetFaultParameterFaultListed
	isBACnetFaultParameterFaultListed() bool
}

// _BACnetFaultParameterFaultListed is the data-structure of this message
type _BACnetFaultParameterFaultListed struct {
	*_BACnetFaultParameter
        OpeningTag BACnetOpeningTag
        FaultListReference BACnetDeviceObjectPropertyReferenceEnclosed
        ClosingTag BACnetClosingTag
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultListed) InitializeParent(parent BACnetFaultParameter , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetFaultParameterFaultListed)  GetParent() BACnetFaultParameter {
	return m._BACnetFaultParameter
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultListed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_BACnetFaultParameterFaultListed) GetFaultListReference() BACnetDeviceObjectPropertyReferenceEnclosed {
	return m.FaultListReference
}

func (m *_BACnetFaultParameterFaultListed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetFaultParameterFaultListed factory function for _BACnetFaultParameterFaultListed
func NewBACnetFaultParameterFaultListed( openingTag BACnetOpeningTag , faultListReference BACnetDeviceObjectPropertyReferenceEnclosed , closingTag BACnetClosingTag , peekedTagHeader BACnetTagHeader ) *_BACnetFaultParameterFaultListed {
	_result := &_BACnetFaultParameterFaultListed{
		OpeningTag: openingTag,
		FaultListReference: faultListReference,
		ClosingTag: closingTag,
    	_BACnetFaultParameter: NewBACnetFaultParameter(peekedTagHeader),
	}
	_result._BACnetFaultParameter._BACnetFaultParameterChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultListed(structType interface{}) BACnetFaultParameterFaultListed {
    if casted, ok := structType.(BACnetFaultParameterFaultListed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultListed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultListed) GetTypeName() string {
	return "BACnetFaultParameterFaultListed"
}

func (m *_BACnetFaultParameterFaultListed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetFaultParameterFaultListed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (faultListReference)
	lengthInBits += m.FaultListReference.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetFaultParameterFaultListed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetFaultParameterFaultListedParse(readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultListed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultListed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultListed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
_openingTag, _openingTagErr := BACnetOpeningTagParse(readBuffer , uint8( uint8(7) ) )
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of BACnetFaultParameterFaultListed")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (faultListReference)
	if pullErr := readBuffer.PullContext("faultListReference"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for faultListReference")
	}
_faultListReference, _faultListReferenceErr := BACnetDeviceObjectPropertyReferenceEnclosedParse(readBuffer , uint8( uint8(0) ) )
	if _faultListReferenceErr != nil {
		return nil, errors.Wrap(_faultListReferenceErr, "Error parsing 'faultListReference' field of BACnetFaultParameterFaultListed")
	}
	faultListReference := _faultListReference.(BACnetDeviceObjectPropertyReferenceEnclosed)
	if closeErr := readBuffer.CloseContext("faultListReference"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for faultListReference")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
_closingTag, _closingTagErr := BACnetClosingTagParse(readBuffer , uint8( uint8(7) ) )
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of BACnetFaultParameterFaultListed")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultListed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultListed")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultListed{
		_BACnetFaultParameter: &_BACnetFaultParameter{
		},
		OpeningTag: openingTag,
		FaultListReference: faultListReference,
		ClosingTag: closingTag,
	}
	_child._BACnetFaultParameter._BACnetFaultParameterChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultListed) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultListed"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultListed")
		}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (faultListReference)
	if pushErr := writeBuffer.PushContext("faultListReference"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for faultListReference")
	}
	_faultListReferenceErr := writeBuffer.WriteSerializable(m.GetFaultListReference())
	if popErr := writeBuffer.PopContext("faultListReference"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for faultListReference")
	}
	if _faultListReferenceErr != nil {
		return errors.Wrap(_faultListReferenceErr, "Error serializing 'faultListReference' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultListed"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultListed")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetFaultParameterFaultListed) isBACnetFaultParameterFaultListed() bool {
	return true
}

func (m *_BACnetFaultParameterFaultListed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



