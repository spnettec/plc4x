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

// SALDataTesting is the corresponding interface of SALDataTesting
type SALDataTesting interface {
	utils.LengthAware
	utils.Serializable
	SALData
}

// SALDataTestingExactly can be used when we want exactly this type and not a type which fulfills SALDataTesting.
// This is useful for switch cases.
type SALDataTestingExactly interface {
	SALDataTesting
	isSALDataTesting() bool
}

// _SALDataTesting is the data-structure of this message
type _SALDataTesting struct {
	*_SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataTesting) GetApplicationId() ApplicationId {
	return ApplicationId_TESTING
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataTesting) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataTesting) GetParent() SALData {
	return m._SALData
}

// NewSALDataTesting factory function for _SALDataTesting
func NewSALDataTesting(salData SALData) *_SALDataTesting {
	_result := &_SALDataTesting{
		_SALData: NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataTesting(structType interface{}) SALDataTesting {
	if casted, ok := structType.(SALDataTesting); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataTesting); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataTesting) GetTypeName() string {
	return "SALDataTesting"
}

func (m *_SALDataTesting) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SALDataTesting) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}

func (m *_SALDataTesting) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALDataTestingParse(theBytes []byte, applicationId ApplicationId) (SALDataTesting, error) {
	return SALDataTestingParseWithBuffer(utils.NewReadBufferByteBased(theBytes), applicationId)
}

func SALDataTestingParseWithBuffer(readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataTesting, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataTesting"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataTesting")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if !(bool((1) == (2))) {
		return nil, errors.WithStack(utils.ParseValidationError{"TESTING Not yet implemented"})
	}

	if closeErr := readBuffer.CloseContext("SALDataTesting"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataTesting")
	}

	// Create a partially initialized instance
	_child := &_SALDataTesting{
		_SALData: &_SALData{},
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataTesting) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataTesting) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataTesting"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataTesting")
		}

		if popErr := writeBuffer.PopContext("SALDataTesting"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataTesting")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_SALDataTesting) isSALDataTesting() bool {
	return true
}

func (m *_SALDataTesting) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
