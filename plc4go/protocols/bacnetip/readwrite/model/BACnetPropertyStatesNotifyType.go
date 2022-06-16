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

// BACnetPropertyStatesNotifyType is the corresponding interface of BACnetPropertyStatesNotifyType
type BACnetPropertyStatesNotifyType interface {
	BACnetPropertyStates
	// GetNotifyType returns NotifyType (property field)
	GetNotifyType() BACnetNotifyTypeTagged
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

// _BACnetPropertyStatesNotifyType is the data-structure of this message
type _BACnetPropertyStatesNotifyType struct {
	*_BACnetPropertyStates
	NotifyType BACnetNotifyTypeTagged
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesNotifyType) InitializeParent(parent BACnetPropertyStates, peekedTagHeader BACnetTagHeader) {
	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesNotifyType) GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesNotifyType) GetNotifyType() BACnetNotifyTypeTagged {
	return m.NotifyType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetPropertyStatesNotifyType factory function for _BACnetPropertyStatesNotifyType
func NewBACnetPropertyStatesNotifyType(notifyType BACnetNotifyTypeTagged, peekedTagHeader BACnetTagHeader) *_BACnetPropertyStatesNotifyType {
	_result := &_BACnetPropertyStatesNotifyType{
		NotifyType:            notifyType,
		_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesNotifyType(structType interface{}) BACnetPropertyStatesNotifyType {
	if casted, ok := structType.(BACnetPropertyStatesNotifyType); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesNotifyType); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesNotifyType) GetTypeName() string {
	return "BACnetPropertyStatesNotifyType"
}

func (m *_BACnetPropertyStatesNotifyType) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetPropertyStatesNotifyType) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (notifyType)
	lengthInBits += m.NotifyType.GetLengthInBits()

	return lengthInBits
}

func (m *_BACnetPropertyStatesNotifyType) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetPropertyStatesNotifyTypeParse(readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesNotifyType, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesNotifyType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesNotifyType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (notifyType)
	if pullErr := readBuffer.PullContext("notifyType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for notifyType")
	}
	_notifyType, _notifyTypeErr := BACnetNotifyTypeTaggedParse(readBuffer, uint8(peekedTagNumber), TagClass(TagClass_CONTEXT_SPECIFIC_TAGS))
	if _notifyTypeErr != nil {
		return nil, errors.Wrap(_notifyTypeErr, "Error parsing 'notifyType' field")
	}
	notifyType := _notifyType.(BACnetNotifyTypeTagged)
	if closeErr := readBuffer.CloseContext("notifyType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for notifyType")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesNotifyType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesNotifyType")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesNotifyType{
		NotifyType:            notifyType,
		_BACnetPropertyStates: &_BACnetPropertyStates{},
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesNotifyType) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesNotifyType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesNotifyType")
		}

		// Simple Field (notifyType)
		if pushErr := writeBuffer.PushContext("notifyType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for notifyType")
		}
		_notifyTypeErr := writeBuffer.WriteSerializable(m.GetNotifyType())
		if popErr := writeBuffer.PopContext("notifyType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for notifyType")
		}
		if _notifyTypeErr != nil {
			return errors.Wrap(_notifyTypeErr, "Error serializing 'notifyType' field")
		}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesNotifyType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesNotifyType")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *_BACnetPropertyStatesNotifyType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
