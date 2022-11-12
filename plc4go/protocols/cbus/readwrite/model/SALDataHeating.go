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


// SALDataHeating is the corresponding interface of SALDataHeating
type SALDataHeating interface {
	utils.LengthAware
	utils.Serializable
	SALData
	// GetHeatingData returns HeatingData (property field)
	GetHeatingData() LightingData
}

// SALDataHeatingExactly can be used when we want exactly this type and not a type which fulfills SALDataHeating.
// This is useful for switch cases.
type SALDataHeatingExactly interface {
	SALDataHeating
	isSALDataHeating() bool
}

// _SALDataHeating is the data-structure of this message
type _SALDataHeating struct {
	*_SALData
        HeatingData LightingData
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataHeating)  GetApplicationId() ApplicationId {
return ApplicationId_HEATING}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataHeating) InitializeParent(parent SALData , salData SALData ) {	m.SalData = salData
}

func (m *_SALDataHeating)  GetParent() SALData {
	return m._SALData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataHeating) GetHeatingData() LightingData {
	return m.HeatingData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewSALDataHeating factory function for _SALDataHeating
func NewSALDataHeating( heatingData LightingData , salData SALData ) *_SALDataHeating {
	_result := &_SALDataHeating{
		HeatingData: heatingData,
    	_SALData: NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataHeating(structType interface{}) SALDataHeating {
    if casted, ok := structType.(SALDataHeating); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataHeating); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataHeating) GetTypeName() string {
	return "SALDataHeating"
}

func (m *_SALDataHeating) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SALDataHeating) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (heatingData)
	lengthInBits += m.HeatingData.GetLengthInBits()

	return lengthInBits
}


func (m *_SALDataHeating) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALDataHeatingParse(theBytes []byte, applicationId ApplicationId) (SALDataHeating, error) {
	return SALDataHeatingParseWithBuffer(utils.NewReadBufferByteBased(theBytes), applicationId)
}

func SALDataHeatingParseWithBuffer(readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataHeating, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataHeating"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataHeating")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (heatingData)
	if pullErr := readBuffer.PullContext("heatingData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for heatingData")
	}
_heatingData, _heatingDataErr := LightingDataParseWithBuffer(readBuffer)
	if _heatingDataErr != nil {
		return nil, errors.Wrap(_heatingDataErr, "Error parsing 'heatingData' field of SALDataHeating")
	}
	heatingData := _heatingData.(LightingData)
	if closeErr := readBuffer.CloseContext("heatingData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for heatingData")
	}

	if closeErr := readBuffer.CloseContext("SALDataHeating"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataHeating")
	}

	// Create a partially initialized instance
	_child := &_SALDataHeating{
		_SALData: &_SALData{
		},
		HeatingData: heatingData,
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataHeating) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataHeating) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataHeating"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataHeating")
		}

	// Simple Field (heatingData)
	if pushErr := writeBuffer.PushContext("heatingData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for heatingData")
	}
	_heatingDataErr := writeBuffer.WriteSerializable(m.GetHeatingData())
	if popErr := writeBuffer.PopContext("heatingData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for heatingData")
	}
	if _heatingDataErr != nil {
		return errors.Wrap(_heatingDataErr, "Error serializing 'heatingData' field")
	}

		if popErr := writeBuffer.PopContext("SALDataHeating"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataHeating")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_SALDataHeating) isSALDataHeating() bool {
	return true
}

func (m *_SALDataHeating) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



