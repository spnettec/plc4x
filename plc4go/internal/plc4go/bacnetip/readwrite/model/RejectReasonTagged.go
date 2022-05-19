/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// RejectReasonTagged is the data-structure of this message
type RejectReasonTagged struct {
	Value            RejectReason
	ProprietaryValue uint32

	// Arguments.
	ActualLength uint32
}

// IRejectReasonTagged is the corresponding interface of RejectReasonTagged
type IRejectReasonTagged interface {
	// GetValue returns Value (property field)
	GetValue() RejectReason
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *RejectReasonTagged) GetValue() RejectReason {
	return m.Value
}

func (m *RejectReasonTagged) GetProprietaryValue() uint32 {
	return m.ProprietaryValue
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *RejectReasonTagged) GetIsProprietary() bool {
	return bool(bool((m.GetValue()) == (RejectReason_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewRejectReasonTagged factory function for RejectReasonTagged
func NewRejectReasonTagged(value RejectReason, proprietaryValue uint32, actualLength uint32) *RejectReasonTagged {
	return &RejectReasonTagged{Value: value, ProprietaryValue: proprietaryValue, ActualLength: actualLength}
}

func CastRejectReasonTagged(structType interface{}) *RejectReasonTagged {
	if casted, ok := structType.(RejectReasonTagged); ok {
		return &casted
	}
	if casted, ok := structType.(*RejectReasonTagged); ok {
		return casted
	}
	return nil
}

func (m *RejectReasonTagged) GetTypeName() string {
	return "RejectReasonTagged"
}

func (m *RejectReasonTagged) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *RejectReasonTagged) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Manual Field (value)
	lengthInBits += uint16(int32(m.ActualLength) * int32(int32(8)))

	// A virtual field doesn't have any in- or output.

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(utils.InlineIf(m.GetIsProprietary(), func() interface{} { return int32(int32(m.ActualLength) * int32(int32(8))) }, func() interface{} { return int32(int32(0)) }).(int32))

	return lengthInBits
}

func (m *RejectReasonTagged) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func RejectReasonTaggedParse(readBuffer utils.ReadBuffer, actualLength uint32) (*RejectReasonTagged, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RejectReasonTagged"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Manual Field (value)
	_value, _valueErr := ReadEnumGeneric(readBuffer, actualLength, RejectReason_VENDOR_PROPRIETARY_VALUE)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field")
	}
	value := _value.(RejectReason)

	// Virtual field
	_isProprietary := bool((value) == (RejectReason_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)
	_ = isProprietary

	// Manual Field (proprietaryValue)
	_proprietaryValue, _proprietaryValueErr := ReadProprietaryEnumGeneric(readBuffer, actualLength, isProprietary)
	if _proprietaryValueErr != nil {
		return nil, errors.Wrap(_proprietaryValueErr, "Error parsing 'proprietaryValue' field")
	}
	proprietaryValue := _proprietaryValue.(uint32)

	if closeErr := readBuffer.CloseContext("RejectReasonTagged"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewRejectReasonTagged(value, proprietaryValue, actualLength), nil
}

func (m *RejectReasonTagged) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("RejectReasonTagged"); pushErr != nil {
		return pushErr
	}

	// Manual Field (value)
	_valueErr := WriteEnumGeneric(writeBuffer, m.GetValue())
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}
	// Virtual field
	if _isProprietaryErr := writeBuffer.WriteVirtual("isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	// Manual Field (proprietaryValue)
	_proprietaryValueErr := WriteProprietaryEnumGeneric(writeBuffer, m.GetProprietaryValue(), m.GetIsProprietary())
	if _proprietaryValueErr != nil {
		return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
	}

	if popErr := writeBuffer.PopContext("RejectReasonTagged"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *RejectReasonTagged) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
