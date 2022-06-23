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

// BACnetConstructedDataAcceptedModes is the corresponding interface of BACnetConstructedDataAcceptedModes
type BACnetConstructedDataAcceptedModes interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAcceptedModes returns AcceptedModes (property field)
	GetAcceptedModes() []BACnetLifeSafetyModeTagged
}

// BACnetConstructedDataAcceptedModesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAcceptedModes.
// This is useful for switch cases.
type BACnetConstructedDataAcceptedModesExactly interface {
	BACnetConstructedDataAcceptedModes
	isBACnetConstructedDataAcceptedModes() bool
}

// _BACnetConstructedDataAcceptedModes is the data-structure of this message
type _BACnetConstructedDataAcceptedModes struct {
	*_BACnetConstructedData
	AcceptedModes []BACnetLifeSafetyModeTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAcceptedModes) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAcceptedModes) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACCEPTED_MODES
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAcceptedModes) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAcceptedModes) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAcceptedModes) GetAcceptedModes() []BACnetLifeSafetyModeTagged {
	return m.AcceptedModes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAcceptedModes factory function for _BACnetConstructedDataAcceptedModes
func NewBACnetConstructedDataAcceptedModes(acceptedModes []BACnetLifeSafetyModeTagged, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAcceptedModes {
	_result := &_BACnetConstructedDataAcceptedModes{
		AcceptedModes:          acceptedModes,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAcceptedModes(structType interface{}) BACnetConstructedDataAcceptedModes {
	if casted, ok := structType.(BACnetConstructedDataAcceptedModes); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAcceptedModes); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAcceptedModes) GetTypeName() string {
	return "BACnetConstructedDataAcceptedModes"
}

func (m *_BACnetConstructedDataAcceptedModes) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataAcceptedModes) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Array field
	if len(m.AcceptedModes) > 0 {
		for _, element := range m.AcceptedModes {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataAcceptedModes) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataAcceptedModesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAcceptedModes, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAcceptedModes"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAcceptedModes")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (acceptedModes)
	if pullErr := readBuffer.PullContext("acceptedModes", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for acceptedModes")
	}
	// Terminated array
	var acceptedModes []BACnetLifeSafetyModeTagged
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetLifeSafetyModeTaggedParse(readBuffer, uint8(0), TagClass_APPLICATION_TAGS)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'acceptedModes' field")
			}
			acceptedModes = append(acceptedModes, _item.(BACnetLifeSafetyModeTagged))

		}
	}
	if closeErr := readBuffer.CloseContext("acceptedModes", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for acceptedModes")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAcceptedModes"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAcceptedModes")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAcceptedModes{
		AcceptedModes: acceptedModes,
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAcceptedModes) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAcceptedModes"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAcceptedModes")
		}

		// Array Field (acceptedModes)
		if pushErr := writeBuffer.PushContext("acceptedModes", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for acceptedModes")
		}
		for _, _element := range m.GetAcceptedModes() {
			_elementErr := writeBuffer.WriteSerializable(_element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'acceptedModes' field")
			}
		}
		if popErr := writeBuffer.PopContext("acceptedModes", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for acceptedModes")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAcceptedModes"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAcceptedModes")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAcceptedModes) isBACnetConstructedDataAcceptedModes() bool {
	return true
}

func (m *_BACnetConstructedDataAcceptedModes) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
