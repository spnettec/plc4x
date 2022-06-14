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
	"github.com/apache/plc4x/plc4go/internal/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataDeployedProfileLocation is the data-structure of this message
type BACnetConstructedDataDeployedProfileLocation struct {
	*BACnetConstructedData
	DeployedProfileLocation *BACnetApplicationTagCharacterString

	// Arguments.
	TagNumber          uint8
	ArrayIndexArgument *BACnetTagPayloadUnsignedInteger
}

// IBACnetConstructedDataDeployedProfileLocation is the corresponding interface of BACnetConstructedDataDeployedProfileLocation
type IBACnetConstructedDataDeployedProfileLocation interface {
	IBACnetConstructedData
	// GetDeployedProfileLocation returns DeployedProfileLocation (property field)
	GetDeployedProfileLocation() *BACnetApplicationTagCharacterString
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetConstructedDataDeployedProfileLocation) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *BACnetConstructedDataDeployedProfileLocation) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_DEPLOYED_PROFILE_LOCATION
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetConstructedDataDeployedProfileLocation) InitializeParent(parent *BACnetConstructedData, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag) {
	m.BACnetConstructedData.OpeningTag = openingTag
	m.BACnetConstructedData.PeekedTagHeader = peekedTagHeader
	m.BACnetConstructedData.ClosingTag = closingTag
}

func (m *BACnetConstructedDataDeployedProfileLocation) GetParent() *BACnetConstructedData {
	return m.BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetConstructedDataDeployedProfileLocation) GetDeployedProfileLocation() *BACnetApplicationTagCharacterString {
	return m.DeployedProfileLocation
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataDeployedProfileLocation factory function for BACnetConstructedDataDeployedProfileLocation
func NewBACnetConstructedDataDeployedProfileLocation(deployedProfileLocation *BACnetApplicationTagCharacterString, openingTag *BACnetOpeningTag, peekedTagHeader *BACnetTagHeader, closingTag *BACnetClosingTag, tagNumber uint8, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) *BACnetConstructedDataDeployedProfileLocation {
	_result := &BACnetConstructedDataDeployedProfileLocation{
		DeployedProfileLocation: deployedProfileLocation,
		BACnetConstructedData:   NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result.Child = _result
	return _result
}

func CastBACnetConstructedDataDeployedProfileLocation(structType interface{}) *BACnetConstructedDataDeployedProfileLocation {
	if casted, ok := structType.(BACnetConstructedDataDeployedProfileLocation); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDeployedProfileLocation); ok {
		return casted
	}
	if casted, ok := structType.(BACnetConstructedData); ok {
		return CastBACnetConstructedDataDeployedProfileLocation(casted.Child)
	}
	if casted, ok := structType.(*BACnetConstructedData); ok {
		return CastBACnetConstructedDataDeployedProfileLocation(casted.Child)
	}
	return nil
}

func (m *BACnetConstructedDataDeployedProfileLocation) GetTypeName() string {
	return "BACnetConstructedDataDeployedProfileLocation"
}

func (m *BACnetConstructedDataDeployedProfileLocation) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConstructedDataDeployedProfileLocation) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (deployedProfileLocation)
	lengthInBits += m.DeployedProfileLocation.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConstructedDataDeployedProfileLocation) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataDeployedProfileLocationParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument *BACnetTagPayloadUnsignedInteger) (*BACnetConstructedDataDeployedProfileLocation, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDeployedProfileLocation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDeployedProfileLocation")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (deployedProfileLocation)
	if pullErr := readBuffer.PullContext("deployedProfileLocation"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for deployedProfileLocation")
	}
	_deployedProfileLocation, _deployedProfileLocationErr := BACnetApplicationTagParse(readBuffer)
	if _deployedProfileLocationErr != nil {
		return nil, errors.Wrap(_deployedProfileLocationErr, "Error parsing 'deployedProfileLocation' field")
	}
	deployedProfileLocation := CastBACnetApplicationTagCharacterString(_deployedProfileLocation)
	if closeErr := readBuffer.CloseContext("deployedProfileLocation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for deployedProfileLocation")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDeployedProfileLocation"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDeployedProfileLocation")
	}

	// Create a partially initialized instance
	_child := &BACnetConstructedDataDeployedProfileLocation{
		DeployedProfileLocation: CastBACnetApplicationTagCharacterString(deployedProfileLocation),
		BACnetConstructedData:   &BACnetConstructedData{},
	}
	_child.BACnetConstructedData.Child = _child
	return _child, nil
}

func (m *BACnetConstructedDataDeployedProfileLocation) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDeployedProfileLocation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDeployedProfileLocation")
		}

		// Simple Field (deployedProfileLocation)
		if pushErr := writeBuffer.PushContext("deployedProfileLocation"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for deployedProfileLocation")
		}
		_deployedProfileLocationErr := writeBuffer.WriteSerializable(m.DeployedProfileLocation)
		if popErr := writeBuffer.PopContext("deployedProfileLocation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for deployedProfileLocation")
		}
		if _deployedProfileLocationErr != nil {
			return errors.Wrap(_deployedProfileLocationErr, "Error serializing 'deployedProfileLocation' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDeployedProfileLocation"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDeployedProfileLocation")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetConstructedDataDeployedProfileLocation) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
