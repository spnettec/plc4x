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


// BACnetConstructedDataMultiStateInputAll is the corresponding interface of BACnetConstructedDataMultiStateInputAll
type BACnetConstructedDataMultiStateInputAll interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
}

// BACnetConstructedDataMultiStateInputAllExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataMultiStateInputAll.
// This is useful for switch cases.
type BACnetConstructedDataMultiStateInputAllExactly interface {
	BACnetConstructedDataMultiStateInputAll
	isBACnetConstructedDataMultiStateInputAll() bool
}

// _BACnetConstructedDataMultiStateInputAll is the data-structure of this message
type _BACnetConstructedDataMultiStateInputAll struct {
	*_BACnetConstructedData
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMultiStateInputAll)  GetObjectTypeArgument() BACnetObjectType {
return BACnetObjectType_MULTI_STATE_INPUT}

func (m *_BACnetConstructedDataMultiStateInputAll)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_ALL}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMultiStateInputAll) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataMultiStateInputAll)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}


// NewBACnetConstructedDataMultiStateInputAll factory function for _BACnetConstructedDataMultiStateInputAll
func NewBACnetConstructedDataMultiStateInputAll( openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataMultiStateInputAll {
	_result := &_BACnetConstructedDataMultiStateInputAll{
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMultiStateInputAll(structType interface{}) BACnetConstructedDataMultiStateInputAll {
    if casted, ok := structType.(BACnetConstructedDataMultiStateInputAll); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMultiStateInputAll); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMultiStateInputAll) GetTypeName() string {
	return "BACnetConstructedDataMultiStateInputAll"
}

func (m *_BACnetConstructedDataMultiStateInputAll) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataMultiStateInputAll) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}


func (m *_BACnetConstructedDataMultiStateInputAll) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataMultiStateInputAllParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataMultiStateInputAll, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMultiStateInputAll"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMultiStateInputAll")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if (!(bool(((1)) == ((2))))) {
		return nil, errors.WithStack(utils.ParseValidationError{"All should never occur in context of constructed data. If it does please report"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMultiStateInputAll"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMultiStateInputAll")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataMultiStateInputAll{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataMultiStateInputAll) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMultiStateInputAll"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMultiStateInputAll")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMultiStateInputAll"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMultiStateInputAll")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataMultiStateInputAll) isBACnetConstructedDataMultiStateInputAll() bool {
	return true
}

func (m *_BACnetConstructedDataMultiStateInputAll) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



