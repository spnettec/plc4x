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
	"github.com/rs/zerolog/log"
	"io"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetConstructedDataConfigurationFiles is the corresponding interface of BACnetConstructedDataConfigurationFiles
type BACnetConstructedDataConfigurationFiles interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetConfigurationFiles returns ConfigurationFiles (property field)
	GetConfigurationFiles() []BACnetApplicationTagObjectIdentifier
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataConfigurationFilesExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataConfigurationFiles.
// This is useful for switch cases.
type BACnetConstructedDataConfigurationFilesExactly interface {
	BACnetConstructedDataConfigurationFiles
	isBACnetConstructedDataConfigurationFiles() bool
}

// _BACnetConstructedDataConfigurationFiles is the data-structure of this message
type _BACnetConstructedDataConfigurationFiles struct {
	*_BACnetConstructedData
        NumberOfDataElements BACnetApplicationTagUnsignedInteger
        ConfigurationFiles []BACnetApplicationTagObjectIdentifier
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataConfigurationFiles)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataConfigurationFiles)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_CONFIGURATION_FILES}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataConfigurationFiles) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataConfigurationFiles)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataConfigurationFiles) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataConfigurationFiles) GetConfigurationFiles() []BACnetApplicationTagObjectIdentifier {
	return m.ConfigurationFiles
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataConfigurationFiles) GetZero() uint64 {
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataConfigurationFiles factory function for _BACnetConstructedDataConfigurationFiles
func NewBACnetConstructedDataConfigurationFiles( numberOfDataElements BACnetApplicationTagUnsignedInteger , configurationFiles []BACnetApplicationTagObjectIdentifier , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataConfigurationFiles {
	_result := &_BACnetConstructedDataConfigurationFiles{
		NumberOfDataElements: numberOfDataElements,
		ConfigurationFiles: configurationFiles,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataConfigurationFiles(structType interface{}) BACnetConstructedDataConfigurationFiles {
    if casted, ok := structType.(BACnetConstructedDataConfigurationFiles); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataConfigurationFiles); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataConfigurationFiles) GetTypeName() string {
	return "BACnetConstructedDataConfigurationFiles"
}

func (m *_BACnetConstructedDataConfigurationFiles) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConstructedDataConfigurationFiles) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits()
	}

	// Array field
	if len(m.ConfigurationFiles) > 0 {
		for _, element := range m.ConfigurationFiles {
			lengthInBits += element.GetLengthInBits()
		}
	}

	return lengthInBits
}


func (m *_BACnetConstructedDataConfigurationFiles) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConstructedDataConfigurationFilesParse(readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataConfigurationFiles, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataConfigurationFiles"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataConfigurationFiles")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
_val, _err := BACnetApplicationTagParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field of BACnetConstructedDataConfigurationFiles")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (configurationFiles)
	if pullErr := readBuffer.PullContext("configurationFiles", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for configurationFiles")
	}
	// Terminated array
	var configurationFiles []BACnetApplicationTagObjectIdentifier
	{
		for ;!bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)); {
_item, _err := BACnetApplicationTagParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'configurationFiles' field of BACnetConstructedDataConfigurationFiles")
			}
			configurationFiles = append(configurationFiles, _item.(BACnetApplicationTagObjectIdentifier))

		}
	}
	if closeErr := readBuffer.CloseContext("configurationFiles", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for configurationFiles")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataConfigurationFiles"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataConfigurationFiles")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataConfigurationFiles{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NumberOfDataElements: numberOfDataElements,
		ConfigurationFiles: configurationFiles,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataConfigurationFiles) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataConfigurationFiles"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataConfigurationFiles")
		}
	// Virtual field
	if _zeroErr := writeBuffer.WriteVirtual("zero", m.GetZero()); _zeroErr != nil {
		return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
	}

	// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if m.GetNumberOfDataElements() != nil {
		if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
		}
		numberOfDataElements = m.GetNumberOfDataElements()
		_numberOfDataElementsErr := writeBuffer.WriteSerializable(numberOfDataElements)
		if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for numberOfDataElements")
		}
		if _numberOfDataElementsErr != nil {
			return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
		}
	}

	// Array Field (configurationFiles)
	if pushErr := writeBuffer.PushContext("configurationFiles", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for configurationFiles")
	}
	for _, _element := range m.GetConfigurationFiles() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'configurationFiles' field")
		}
	}
	if popErr := writeBuffer.PopContext("configurationFiles", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for configurationFiles")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataConfigurationFiles"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataConfigurationFiles")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataConfigurationFiles) isBACnetConstructedDataConfigurationFiles() bool {
	return true
}

func (m *_BACnetConstructedDataConfigurationFiles) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



