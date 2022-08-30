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


// BACnetConstructedDataProcessIdentifier is the corresponding interface of BACnetConstructedDataProcessIdentifier
type BACnetConstructedDataProcessIdentifier interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetProcessIdentifier returns ProcessIdentifier (property field)
	GetProcessIdentifier() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataProcessIdentifierExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataProcessIdentifier.
// This is useful for switch cases.
type BACnetConstructedDataProcessIdentifierExactly interface {
	BACnetConstructedDataProcessIdentifier
	isBACnetConstructedDataProcessIdentifier() bool
}

// _BACnetConstructedDataProcessIdentifier is the data-structure of this message
type _BACnetConstructedDataProcessIdentifier struct {
	*_BACnetConstructedData
        ProcessIdentifier BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataProcessIdentifier)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataProcessIdentifier)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_PROCESS_IDENTIFIER}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataProcessIdentifier) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataProcessIdentifier)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataProcessIdentifier) GetProcessIdentifier() BACnetApplicationTagUnsignedInteger {
	return m.ProcessIdentifier
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataProcessIdentifier) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetProcessIdentifier())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataProcessIdentifier factory function for _BACnetConstructedDataProcessIdentifier
func NewBACnetConstructedDataProcessIdentifier( processIdentifier BACnetApplicationTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataProcessIdentifier {
	_result := &_BACnetConstructedDataProcessIdentifier{
		ProcessIdentifier: processIdentifier,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataProcessIdentifier(structType interface{}) BACnetConstructedDataProcessIdentifier {
    if casted, ok := structType.(BACnetConstructedDataProcessIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataProcessIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataProcessIdentifier) GetTypeName() string {
	return "BACnetConstructedDataProcessIdentifier"
}

func (m *_BACnetConstructedDataProcessIdentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataProcessIdentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (processIdentifier)
	lengthInBits += m.ProcessIdentifier.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataProcessIdentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataProcessIdentifierParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataProcessIdentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataProcessIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataProcessIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (processIdentifier)
	if pullErr := readBuffer.PullContext("processIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for processIdentifier")
	}
_processIdentifier, _processIdentifierErr := BACnetApplicationTagParse(readBuffer)
	if _processIdentifierErr != nil {
		return nil, errors.Wrap(_processIdentifierErr, "Error parsing 'processIdentifier' field of BACnetConstructedDataProcessIdentifier")
	}
	processIdentifier := _processIdentifier.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("processIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for processIdentifier")
	}

	// Virtual field
	_actualValue := processIdentifier
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataProcessIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataProcessIdentifier")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataProcessIdentifier{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ProcessIdentifier: processIdentifier,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataProcessIdentifier) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataProcessIdentifier"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataProcessIdentifier")
		}

	// Simple Field (processIdentifier)
	if pushErr := writeBuffer.PushContext("processIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for processIdentifier")
	}
	_processIdentifierErr := writeBuffer.WriteSerializable(m.GetProcessIdentifier())
	if popErr := writeBuffer.PopContext("processIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for processIdentifier")
	}
	if _processIdentifierErr != nil {
		return errors.Wrap(_processIdentifierErr, "Error serializing 'processIdentifier' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataProcessIdentifier"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataProcessIdentifier")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataProcessIdentifier) isBACnetConstructedDataProcessIdentifier() bool {
	return true
}

func (m *_BACnetConstructedDataProcessIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



