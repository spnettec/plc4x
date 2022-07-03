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

// CALDataRequestGetStatus is the corresponding interface of CALDataRequestGetStatus
type CALDataRequestGetStatus interface {
	utils.LengthAware
	utils.Serializable
	CALData
	// GetParamNo returns ParamNo (property field)
	GetParamNo() uint8
	// GetCount returns Count (property field)
	GetCount() uint8
}

// CALDataRequestGetStatusExactly can be used when we want exactly this type and not a type which fulfills CALDataRequestGetStatus.
// This is useful for switch cases.
type CALDataRequestGetStatusExactly interface {
	CALDataRequestGetStatus
	isCALDataRequestGetStatus() bool
}

// _CALDataRequestGetStatus is the data-structure of this message
type _CALDataRequestGetStatus struct {
	*_CALData
	ParamNo uint8
	Count   uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataRequestGetStatus) InitializeParent(parent CALData, commandTypeContainer CALCommandTypeContainer) {
	m.CommandTypeContainer = commandTypeContainer
}

func (m *_CALDataRequestGetStatus) GetParent() CALData {
	return m._CALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataRequestGetStatus) GetParamNo() uint8 {
	return m.ParamNo
}

func (m *_CALDataRequestGetStatus) GetCount() uint8 {
	return m.Count
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewCALDataRequestGetStatus factory function for _CALDataRequestGetStatus
func NewCALDataRequestGetStatus(paramNo uint8, count uint8, commandTypeContainer CALCommandTypeContainer) *_CALDataRequestGetStatus {
	_result := &_CALDataRequestGetStatus{
		ParamNo:  paramNo,
		Count:    count,
		_CALData: NewCALData(commandTypeContainer),
	}
	_result._CALData._CALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataRequestGetStatus(structType interface{}) CALDataRequestGetStatus {
	if casted, ok := structType.(CALDataRequestGetStatus); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataRequestGetStatus); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataRequestGetStatus) GetTypeName() string {
	return "CALDataRequestGetStatus"
}

func (m *_CALDataRequestGetStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CALDataRequestGetStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (paramNo)
	lengthInBits += 8

	// Simple field (count)
	lengthInBits += 8

	return lengthInBits
}

func (m *_CALDataRequestGetStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataRequestGetStatusParse(readBuffer utils.ReadBuffer) (CALDataRequestGetStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataRequestGetStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataRequestGetStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (paramNo)
	_paramNo, _paramNoErr := readBuffer.ReadUint8("paramNo", 8)
	if _paramNoErr != nil {
		return nil, errors.Wrap(_paramNoErr, "Error parsing 'paramNo' field of CALDataRequestGetStatus")
	}
	paramNo := _paramNo

	// Simple Field (count)
	_count, _countErr := readBuffer.ReadUint8("count", 8)
	if _countErr != nil {
		return nil, errors.Wrap(_countErr, "Error parsing 'count' field of CALDataRequestGetStatus")
	}
	count := _count

	if closeErr := readBuffer.CloseContext("CALDataRequestGetStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataRequestGetStatus")
	}

	// Create a partially initialized instance
	_child := &_CALDataRequestGetStatus{
		ParamNo:  paramNo,
		Count:    count,
		_CALData: &_CALData{},
	}
	_child._CALData._CALDataChildRequirements = _child
	return _child, nil
}

func (m *_CALDataRequestGetStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataRequestGetStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataRequestGetStatus")
		}

		// Simple Field (paramNo)
		paramNo := uint8(m.GetParamNo())
		_paramNoErr := writeBuffer.WriteUint8("paramNo", 8, (paramNo))
		if _paramNoErr != nil {
			return errors.Wrap(_paramNoErr, "Error serializing 'paramNo' field")
		}

		// Simple Field (count)
		count := uint8(m.GetCount())
		_countErr := writeBuffer.WriteUint8("count", 8, (count))
		if _countErr != nil {
			return errors.Wrap(_countErr, "Error serializing 'count' field")
		}

		if popErr := writeBuffer.PopContext("CALDataRequestGetStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataRequestGetStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_CALDataRequestGetStatus) isCALDataRequestGetStatus() bool {
	return true
}

func (m *_CALDataRequestGetStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
