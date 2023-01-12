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


// MeteringDataOtherWaterConsumption is the corresponding interface of MeteringDataOtherWaterConsumption
type MeteringDataOtherWaterConsumption interface {
	utils.LengthAware
	utils.Serializable
	MeteringData
	// GetKL returns KL (property field)
	GetKL() uint32
}

// MeteringDataOtherWaterConsumptionExactly can be used when we want exactly this type and not a type which fulfills MeteringDataOtherWaterConsumption.
// This is useful for switch cases.
type MeteringDataOtherWaterConsumptionExactly interface {
	MeteringDataOtherWaterConsumption
	isMeteringDataOtherWaterConsumption() bool
}

// _MeteringDataOtherWaterConsumption is the data-structure of this message
type _MeteringDataOtherWaterConsumption struct {
	*_MeteringData
        KL uint32
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MeteringDataOtherWaterConsumption) InitializeParent(parent MeteringData , commandTypeContainer MeteringCommandTypeContainer , argument byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_MeteringDataOtherWaterConsumption)  GetParent() MeteringData {
	return m._MeteringData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MeteringDataOtherWaterConsumption) GetKL() uint32 {
	return m.KL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewMeteringDataOtherWaterConsumption factory function for _MeteringDataOtherWaterConsumption
func NewMeteringDataOtherWaterConsumption( kL uint32 , commandTypeContainer MeteringCommandTypeContainer , argument byte ) *_MeteringDataOtherWaterConsumption {
	_result := &_MeteringDataOtherWaterConsumption{
		KL: kL,
    	_MeteringData: NewMeteringData(commandTypeContainer, argument),
	}
	_result._MeteringData._MeteringDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMeteringDataOtherWaterConsumption(structType interface{}) MeteringDataOtherWaterConsumption {
    if casted, ok := structType.(MeteringDataOtherWaterConsumption); ok {
		return casted
	}
	if casted, ok := structType.(*MeteringDataOtherWaterConsumption); ok {
		return *casted
	}
	return nil
}

func (m *_MeteringDataOtherWaterConsumption) GetTypeName() string {
	return "MeteringDataOtherWaterConsumption"
}

func (m *_MeteringDataOtherWaterConsumption) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_MeteringDataOtherWaterConsumption) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (kL)
	lengthInBits += 32;

	return lengthInBits
}


func (m *_MeteringDataOtherWaterConsumption) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MeteringDataOtherWaterConsumptionParse(theBytes []byte) (MeteringDataOtherWaterConsumption, error) {
	return MeteringDataOtherWaterConsumptionParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func MeteringDataOtherWaterConsumptionParseWithBuffer(readBuffer utils.ReadBuffer) (MeteringDataOtherWaterConsumption, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MeteringDataOtherWaterConsumption"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MeteringDataOtherWaterConsumption")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (kL)
_kL, _kLErr := readBuffer.ReadUint32("kL", 32)
	if _kLErr != nil {
		return nil, errors.Wrap(_kLErr, "Error parsing 'kL' field of MeteringDataOtherWaterConsumption")
	}
	kL := _kL

	if closeErr := readBuffer.CloseContext("MeteringDataOtherWaterConsumption"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MeteringDataOtherWaterConsumption")
	}

	// Create a partially initialized instance
	_child := &_MeteringDataOtherWaterConsumption{
		_MeteringData: &_MeteringData{
		},
		KL: kL,
	}
	_child._MeteringData._MeteringDataChildRequirements = _child
	return _child, nil
}

func (m *_MeteringDataOtherWaterConsumption) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MeteringDataOtherWaterConsumption) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MeteringDataOtherWaterConsumption"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MeteringDataOtherWaterConsumption")
		}

	// Simple Field (kL)
	kL := uint32(m.GetKL())
	_kLErr := writeBuffer.WriteUint32("kL", 32, (kL))
	if _kLErr != nil {
		return errors.Wrap(_kLErr, "Error serializing 'kL' field")
	}

		if popErr := writeBuffer.PopContext("MeteringDataOtherWaterConsumption"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MeteringDataOtherWaterConsumption")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_MeteringDataOtherWaterConsumption) isMeteringDataOtherWaterConsumption() bool {
	return true
}

func (m *_MeteringDataOtherWaterConsumption) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



