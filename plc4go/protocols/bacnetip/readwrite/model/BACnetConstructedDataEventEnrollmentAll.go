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


// BACnetConstructedDataEventEnrollmentAll is the corresponding interface of BACnetConstructedDataEventEnrollmentAll
type BACnetConstructedDataEventEnrollmentAll interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
}

// BACnetConstructedDataEventEnrollmentAllExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataEventEnrollmentAll.
// This is useful for switch cases.
type BACnetConstructedDataEventEnrollmentAllExactly interface {
	BACnetConstructedDataEventEnrollmentAll
	isBACnetConstructedDataEventEnrollmentAll() bool
}

// _BACnetConstructedDataEventEnrollmentAll is the data-structure of this message
type _BACnetConstructedDataEventEnrollmentAll struct {
	*_BACnetConstructedData
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataEventEnrollmentAll)  GetObjectTypeArgument() BACnetObjectType {
return BACnetObjectType_EVENT_ENROLLMENT}

func (m *_BACnetConstructedDataEventEnrollmentAll)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_ALL}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataEventEnrollmentAll) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataEventEnrollmentAll)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}


// NewBACnetConstructedDataEventEnrollmentAll factory function for _BACnetConstructedDataEventEnrollmentAll
func NewBACnetConstructedDataEventEnrollmentAll( openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataEventEnrollmentAll {
	_result := &_BACnetConstructedDataEventEnrollmentAll{
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataEventEnrollmentAll(structType interface{}) BACnetConstructedDataEventEnrollmentAll {
    if casted, ok := structType.(BACnetConstructedDataEventEnrollmentAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataEventEnrollmentAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataEventEnrollmentAll) GetTypeName() string {
	return "BACnetConstructedDataEventEnrollmentAll"
}

func (m *_BACnetConstructedDataEventEnrollmentAll) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataEventEnrollmentAll) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}


func (m *_BACnetConstructedDataEventEnrollmentAll) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataEventEnrollmentAllParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataEventEnrollmentAll, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataEventEnrollmentAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataEventEnrollmentAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if (!(bool(((1)) == ((2))))) {
		return nil, errors.WithStack(utils.ParseValidationError{"All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataEventEnrollmentAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataEventEnrollmentAll")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataEventEnrollmentAll{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataEventEnrollmentAll) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataEventEnrollmentAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataEventEnrollmentAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataEventEnrollmentAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataEventEnrollmentAll")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataEventEnrollmentAll) isBACnetConstructedDataEventEnrollmentAll() bool {
	return true
}

func (m *_BACnetConstructedDataEventEnrollmentAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



