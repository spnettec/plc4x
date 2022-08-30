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


// BACnetConstructedDataNetworkPortMaxMaster is the corresponding interface of BACnetConstructedDataNetworkPortMaxMaster
type BACnetConstructedDataNetworkPortMaxMaster interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaxMaster returns MaxMaster (property field)
	GetMaxMaster() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataNetworkPortMaxMasterExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataNetworkPortMaxMaster.
// This is useful for switch cases.
type BACnetConstructedDataNetworkPortMaxMasterExactly interface {
	BACnetConstructedDataNetworkPortMaxMaster
	isBACnetConstructedDataNetworkPortMaxMaster() bool
}

// _BACnetConstructedDataNetworkPortMaxMaster is the data-structure of this message
type _BACnetConstructedDataNetworkPortMaxMaster struct {
	*_BACnetConstructedData
        MaxMaster BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataNetworkPortMaxMaster)  GetObjectTypeArgument() BACnetObjectType {
return BACnetObjectType_NETWORK_PORT}

func (m *_BACnetConstructedDataNetworkPortMaxMaster)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_MAX_MASTER}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataNetworkPortMaxMaster) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataNetworkPortMaxMaster)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkPortMaxMaster) GetMaxMaster() BACnetApplicationTagUnsignedInteger {
	return m.MaxMaster
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataNetworkPortMaxMaster) GetActualValue() BACnetApplicationTagUnsignedInteger {
	return CastBACnetApplicationTagUnsignedInteger(m.GetMaxMaster())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataNetworkPortMaxMaster factory function for _BACnetConstructedDataNetworkPortMaxMaster
func NewBACnetConstructedDataNetworkPortMaxMaster( maxMaster BACnetApplicationTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataNetworkPortMaxMaster {
	_result := &_BACnetConstructedDataNetworkPortMaxMaster{
		MaxMaster: maxMaster,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataNetworkPortMaxMaster(structType interface{}) BACnetConstructedDataNetworkPortMaxMaster {
    if casted, ok := structType.(BACnetConstructedDataNetworkPortMaxMaster); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataNetworkPortMaxMaster); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataNetworkPortMaxMaster) GetTypeName() string {
	return "BACnetConstructedDataNetworkPortMaxMaster"
}

func (m *_BACnetConstructedDataNetworkPortMaxMaster) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataNetworkPortMaxMaster) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (maxMaster)
	lengthInBits += m.MaxMaster.GetLengthInBits()

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataNetworkPortMaxMaster) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataNetworkPortMaxMasterParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataNetworkPortMaxMaster, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataNetworkPortMaxMaster"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataNetworkPortMaxMaster")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (maxMaster)
	if pullErr := readBuffer.PullContext("maxMaster"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for maxMaster")
	}
_maxMaster, _maxMasterErr := BACnetApplicationTagParse(readBuffer)
	if _maxMasterErr != nil {
		return nil, errors.Wrap(_maxMasterErr, "Error parsing 'maxMaster' field of BACnetConstructedDataNetworkPortMaxMaster")
	}
	maxMaster := _maxMaster.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("maxMaster"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for maxMaster")
	}

	// Virtual field
	_actualValue := maxMaster
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataNetworkPortMaxMaster"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataNetworkPortMaxMaster")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataNetworkPortMaxMaster{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		MaxMaster: maxMaster,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataNetworkPortMaxMaster) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataNetworkPortMaxMaster"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataNetworkPortMaxMaster")
		}

	// Simple Field (maxMaster)
	if pushErr := writeBuffer.PushContext("maxMaster"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for maxMaster")
	}
	_maxMasterErr := writeBuffer.WriteSerializable(m.GetMaxMaster())
	if popErr := writeBuffer.PopContext("maxMaster"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for maxMaster")
	}
	if _maxMasterErr != nil {
		return errors.Wrap(_maxMasterErr, "Error serializing 'maxMaster' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual("actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataNetworkPortMaxMaster"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataNetworkPortMaxMaster")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataNetworkPortMaxMaster) isBACnetConstructedDataNetworkPortMaxMaster() bool {
	return true
}

func (m *_BACnetConstructedDataNetworkPortMaxMaster) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



