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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

	// Code generated by code-generation. DO NOT EDIT.


// SALDataAirConditioning is the corresponding interface of SALDataAirConditioning
type SALDataAirConditioning interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetAirConditioningData returns AirConditioningData (property field)
	GetAirConditioningData() AirConditioningData
}

// SALDataAirConditioningExactly can be used when we want exactly this type and not a type which fulfills SALDataAirConditioning.
// This is useful for switch cases.
type SALDataAirConditioningExactly interface {
	SALDataAirConditioning
	isSALDataAirConditioning() bool
}

// _SALDataAirConditioning is the data-structure of this message
type _SALDataAirConditioning struct {
	*_SALData
        AirConditioningData AirConditioningData
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataAirConditioning)  GetApplicationId() ApplicationId {
return ApplicationId_AIR_CONDITIONING}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataAirConditioning) InitializeParent(parent SALData , salData SALData ) {	m.SalData = salData
}

func (m *_SALDataAirConditioning)  GetParent() SALData {
	return m._SALData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataAirConditioning) GetAirConditioningData() AirConditioningData {
	return m.AirConditioningData
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewSALDataAirConditioning factory function for _SALDataAirConditioning
func NewSALDataAirConditioning( airConditioningData AirConditioningData , salData SALData ) *_SALDataAirConditioning {
	_result := &_SALDataAirConditioning{
		AirConditioningData: airConditioningData,
    	_SALData: NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataAirConditioning(structType any) SALDataAirConditioning {
    if casted, ok := structType.(SALDataAirConditioning); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataAirConditioning); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataAirConditioning) GetTypeName() string {
	return "SALDataAirConditioning"
}

func (m *_SALDataAirConditioning) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (airConditioningData)
	lengthInBits += m.AirConditioningData.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_SALDataAirConditioning) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SALDataAirConditioningParse(ctx context.Context, theBytes []byte, applicationId ApplicationId) (SALDataAirConditioning, error) {
	return SALDataAirConditioningParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), applicationId)
}

func SALDataAirConditioningParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataAirConditioning, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SALDataAirConditioning"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataAirConditioning")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (airConditioningData)
	if pullErr := readBuffer.PullContext("airConditioningData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for airConditioningData")
	}
_airConditioningData, _airConditioningDataErr := AirConditioningDataParseWithBuffer(ctx, readBuffer)
	if _airConditioningDataErr != nil {
		return nil, errors.Wrap(_airConditioningDataErr, "Error parsing 'airConditioningData' field of SALDataAirConditioning")
	}
	airConditioningData := _airConditioningData.(AirConditioningData)
	if closeErr := readBuffer.CloseContext("airConditioningData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for airConditioningData")
	}

	if closeErr := readBuffer.CloseContext("SALDataAirConditioning"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataAirConditioning")
	}

	// Create a partially initialized instance
	_child := &_SALDataAirConditioning{
		_SALData: &_SALData{
		},
		AirConditioningData: airConditioningData,
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataAirConditioning) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataAirConditioning) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataAirConditioning"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataAirConditioning")
		}

	// Simple Field (airConditioningData)
	if pushErr := writeBuffer.PushContext("airConditioningData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for airConditioningData")
	}
	_airConditioningDataErr := writeBuffer.WriteSerializable(ctx, m.GetAirConditioningData())
	if popErr := writeBuffer.PopContext("airConditioningData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for airConditioningData")
	}
	if _airConditioningDataErr != nil {
		return errors.Wrap(_airConditioningDataErr, "Error serializing 'airConditioningData' field")
	}

		if popErr := writeBuffer.PopContext("SALDataAirConditioning"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataAirConditioning")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_SALDataAirConditioning) isSALDataAirConditioning() bool {
	return true
}

func (m *_SALDataAirConditioning) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



