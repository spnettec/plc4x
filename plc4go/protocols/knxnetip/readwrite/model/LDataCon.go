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


// LDataCon is the corresponding interface of LDataCon
type LDataCon interface {
	utils.LengthAware
	utils.Serializable
	CEMI
	// GetAdditionalInformationLength returns AdditionalInformationLength (property field)
	GetAdditionalInformationLength() uint8
	// GetAdditionalInformation returns AdditionalInformation (property field)
	GetAdditionalInformation() []CEMIAdditionalInformation
	// GetDataFrame returns DataFrame (property field)
	GetDataFrame() LDataFrame
}

// LDataConExactly can be used when we want exactly this type and not a type which fulfills LDataCon.
// This is useful for switch cases.
type LDataConExactly interface {
	LDataCon
	isLDataCon() bool
}

// _LDataCon is the data-structure of this message
type _LDataCon struct {
	*_CEMI
        AdditionalInformationLength uint8
        AdditionalInformation []CEMIAdditionalInformation
        DataFrame LDataFrame
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_LDataCon)  GetMessageCode() uint8 {
return 0x2E}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_LDataCon) InitializeParent(parent CEMI ) {}

func (m *_LDataCon)  GetParent() CEMI {
	return m._CEMI
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_LDataCon) GetAdditionalInformationLength() uint8 {
	return m.AdditionalInformationLength
}

func (m *_LDataCon) GetAdditionalInformation() []CEMIAdditionalInformation {
	return m.AdditionalInformation
}

func (m *_LDataCon) GetDataFrame() LDataFrame {
	return m.DataFrame
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewLDataCon factory function for _LDataCon
func NewLDataCon( additionalInformationLength uint8 , additionalInformation []CEMIAdditionalInformation , dataFrame LDataFrame , size uint16 ) *_LDataCon {
	_result := &_LDataCon{
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation: additionalInformation,
		DataFrame: dataFrame,
    	_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastLDataCon(structType interface{}) LDataCon {
    if casted, ok := structType.(LDataCon); ok {
		return casted
	}
	if casted, ok := structType.(*LDataCon); ok {
		return *casted
	}
	return nil
}

func (m *_LDataCon) GetTypeName() string {
	return "LDataCon"
}

func (m *_LDataCon) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_LDataCon) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (additionalInformationLength)
	lengthInBits += 8;

	// Array field
	if len(m.AdditionalInformation) > 0 {
		for _, element := range m.AdditionalInformation {
			lengthInBits += element.GetLengthInBits()
		}
	}

	// Simple field (dataFrame)
	lengthInBits += m.DataFrame.GetLengthInBits()

	return lengthInBits
}


func (m *_LDataCon) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func LDataConParse(readBuffer utils.ReadBuffer, size uint16) (LDataCon, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("LDataCon"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for LDataCon")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (additionalInformationLength)
_additionalInformationLength, _additionalInformationLengthErr := readBuffer.ReadUint8("additionalInformationLength", 8)
	if _additionalInformationLengthErr != nil {
		return nil, errors.Wrap(_additionalInformationLengthErr, "Error parsing 'additionalInformationLength' field of LDataCon")
	}
	additionalInformationLength := _additionalInformationLength

	// Array field (additionalInformation)
	if pullErr := readBuffer.PullContext("additionalInformation", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for additionalInformation")
	}
	// Length array
	var additionalInformation []CEMIAdditionalInformation
	{
		_additionalInformationLength := additionalInformationLength
		_additionalInformationEndPos := positionAware.GetPos() + uint16(_additionalInformationLength)
		for ;positionAware.GetPos() < _additionalInformationEndPos; {
_item, _err := CEMIAdditionalInformationParse(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'additionalInformation' field of LDataCon")
			}
			additionalInformation = append(additionalInformation, _item.(CEMIAdditionalInformation))
		}
	}
	if closeErr := readBuffer.CloseContext("additionalInformation", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for additionalInformation")
	}

	// Simple Field (dataFrame)
	if pullErr := readBuffer.PullContext("dataFrame"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for dataFrame")
	}
_dataFrame, _dataFrameErr := LDataFrameParse(readBuffer)
	if _dataFrameErr != nil {
		return nil, errors.Wrap(_dataFrameErr, "Error parsing 'dataFrame' field of LDataCon")
	}
	dataFrame := _dataFrame.(LDataFrame)
	if closeErr := readBuffer.CloseContext("dataFrame"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for dataFrame")
	}

	if closeErr := readBuffer.CloseContext("LDataCon"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for LDataCon")
	}

	// Create a partially initialized instance
	_child := &_LDataCon{
		_CEMI: &_CEMI{
			Size: size,
		},
		AdditionalInformationLength: additionalInformationLength,
		AdditionalInformation: additionalInformation,
		DataFrame: dataFrame,
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_LDataCon) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("LDataCon"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for LDataCon")
		}

	// Simple Field (additionalInformationLength)
	additionalInformationLength := uint8(m.GetAdditionalInformationLength())
	_additionalInformationLengthErr := writeBuffer.WriteUint8("additionalInformationLength", 8, (additionalInformationLength))
	if _additionalInformationLengthErr != nil {
		return errors.Wrap(_additionalInformationLengthErr, "Error serializing 'additionalInformationLength' field")
	}

	// Array Field (additionalInformation)
	if pushErr := writeBuffer.PushContext("additionalInformation", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for additionalInformation")
	}
	for _, _element := range m.GetAdditionalInformation() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'additionalInformation' field")
		}
	}
	if popErr := writeBuffer.PopContext("additionalInformation", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for additionalInformation")
	}

	// Simple Field (dataFrame)
	if pushErr := writeBuffer.PushContext("dataFrame"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for dataFrame")
	}
	_dataFrameErr := writeBuffer.WriteSerializable(m.GetDataFrame())
	if popErr := writeBuffer.PopContext("dataFrame"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for dataFrame")
	}
	if _dataFrameErr != nil {
		return errors.Wrap(_dataFrameErr, "Error serializing 'dataFrame' field")
	}

		if popErr := writeBuffer.PopContext("LDataCon"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for LDataCon")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_LDataCon) isLDataCon() bool {
	return true
}

func (m *_LDataCon) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



