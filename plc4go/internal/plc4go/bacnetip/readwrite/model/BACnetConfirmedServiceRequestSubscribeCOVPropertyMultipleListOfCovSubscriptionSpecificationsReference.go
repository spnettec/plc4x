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
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference is the data-structure of this message
type BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference struct {
	MonitoredProperty *BACnetPropertyReferenceEnclosed
	CovIncrement      *BACnetContextTagReal
	Timestamped       *BACnetContextTagBoolean
}

// IBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference is the corresponding interface of BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference
type IBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference interface {
	// GetMonitoredProperty returns MonitoredProperty (property field)
	GetMonitoredProperty() *BACnetPropertyReferenceEnclosed
	// GetCovIncrement returns CovIncrement (property field)
	GetCovIncrement() *BACnetContextTagReal
	// GetTimestamped returns Timestamped (property field)
	GetTimestamped() *BACnetContextTagBoolean
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

func (m *BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) GetMonitoredProperty() *BACnetPropertyReferenceEnclosed {
	return m.MonitoredProperty
}

func (m *BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) GetCovIncrement() *BACnetContextTagReal {
	return m.CovIncrement
}

func (m *BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) GetTimestamped() *BACnetContextTagBoolean {
	return m.Timestamped
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference factory function for BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference
func NewBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference(monitoredProperty *BACnetPropertyReferenceEnclosed, covIncrement *BACnetContextTagReal, timestamped *BACnetContextTagBoolean) *BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference {
	return &BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference{MonitoredProperty: monitoredProperty, CovIncrement: covIncrement, Timestamped: timestamped}
}

func CastBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference(structType interface{}) *BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference {
	if casted, ok := structType.(BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference); ok {
		return casted
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) GetTypeName() string {
	return "BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference"
}

func (m *BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (monitoredProperty)
	lengthInBits += m.MonitoredProperty.GetLengthInBits()

	// Optional Field (covIncrement)
	if m.CovIncrement != nil {
		lengthInBits += (*m.CovIncrement).GetLengthInBits()
	}

	// Simple field (timestamped)
	lengthInBits += m.Timestamped.GetLengthInBits()

	return lengthInBits
}

func (m *BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReferenceParse(readBuffer utils.ReadBuffer) (*BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (monitoredProperty)
	if pullErr := readBuffer.PullContext("monitoredProperty"); pullErr != nil {
		return nil, pullErr
	}
	_monitoredProperty, _monitoredPropertyErr := BACnetPropertyReferenceEnclosedParse(readBuffer, uint8(uint8(1)))
	if _monitoredPropertyErr != nil {
		return nil, errors.Wrap(_monitoredPropertyErr, "Error parsing 'monitoredProperty' field")
	}
	monitoredProperty := CastBACnetPropertyReferenceEnclosed(_monitoredProperty)
	if closeErr := readBuffer.CloseContext("monitoredProperty"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (covIncrement) (Can be skipped, if a given expression evaluates to false)
	var covIncrement *BACnetContextTagReal = nil
	{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("covIncrement"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetContextTagParse(readBuffer, uint8(1), BACnetDataType_REAL)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'covIncrement' field")
		default:
			covIncrement = CastBACnetContextTagReal(_val)
			if closeErr := readBuffer.CloseContext("covIncrement"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	// Simple Field (timestamped)
	if pullErr := readBuffer.PullContext("timestamped"); pullErr != nil {
		return nil, pullErr
	}
	_timestamped, _timestampedErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_BOOLEAN))
	if _timestampedErr != nil {
		return nil, errors.Wrap(_timestampedErr, "Error parsing 'timestamped' field")
	}
	timestamped := CastBACnetContextTagBoolean(_timestamped)
	if closeErr := readBuffer.CloseContext("timestamped"); closeErr != nil {
		return nil, closeErr
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference"); closeErr != nil {
		return nil, closeErr
	}

	// Create the instance
	return NewBACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference(monitoredProperty, covIncrement, timestamped), nil
}

func (m *BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference"); pushErr != nil {
		return pushErr
	}

	// Simple Field (monitoredProperty)
	if pushErr := writeBuffer.PushContext("monitoredProperty"); pushErr != nil {
		return pushErr
	}
	_monitoredPropertyErr := m.MonitoredProperty.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("monitoredProperty"); popErr != nil {
		return popErr
	}
	if _monitoredPropertyErr != nil {
		return errors.Wrap(_monitoredPropertyErr, "Error serializing 'monitoredProperty' field")
	}

	// Optional Field (covIncrement) (Can be skipped, if the value is null)
	var covIncrement *BACnetContextTagReal = nil
	if m.CovIncrement != nil {
		if pushErr := writeBuffer.PushContext("covIncrement"); pushErr != nil {
			return pushErr
		}
		covIncrement = m.CovIncrement
		_covIncrementErr := covIncrement.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("covIncrement"); popErr != nil {
			return popErr
		}
		if _covIncrementErr != nil {
			return errors.Wrap(_covIncrementErr, "Error serializing 'covIncrement' field")
		}
	}

	// Simple Field (timestamped)
	if pushErr := writeBuffer.PushContext("timestamped"); pushErr != nil {
		return pushErr
	}
	_timestampedErr := m.Timestamped.Serialize(writeBuffer)
	if popErr := writeBuffer.PopContext("timestamped"); popErr != nil {
		return popErr
	}
	if _timestampedErr != nil {
		return errors.Wrap(_timestampedErr, "Error serializing 'timestamped' field")
	}

	if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference"); popErr != nil {
		return popErr
	}
	return nil
}

func (m *BACnetConfirmedServiceRequestSubscribeCOVPropertyMultipleListOfCovSubscriptionSpecificationsReference) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
