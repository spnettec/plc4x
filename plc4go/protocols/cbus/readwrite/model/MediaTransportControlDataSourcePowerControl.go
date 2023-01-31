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


// MediaTransportControlDataSourcePowerControl is the corresponding interface of MediaTransportControlDataSourcePowerControl
type MediaTransportControlDataSourcePowerControl interface {
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetState returns State (property field)
	GetState() byte
	// GetIsShouldPowerOn returns IsShouldPowerOn (virtual field)
	GetIsShouldPowerOn() bool
	// GetIsShouldPowerOff returns IsShouldPowerOff (virtual field)
	GetIsShouldPowerOff() bool
}

// MediaTransportControlDataSourcePowerControlExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataSourcePowerControl.
// This is useful for switch cases.
type MediaTransportControlDataSourcePowerControlExactly interface {
	MediaTransportControlDataSourcePowerControl
	isMediaTransportControlDataSourcePowerControl() bool
}

// _MediaTransportControlDataSourcePowerControl is the data-structure of this message
type _MediaTransportControlDataSourcePowerControl struct {
	*_MediaTransportControlData
        State byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataSourcePowerControl) InitializeParent(parent MediaTransportControlData , commandTypeContainer MediaTransportControlCommandTypeContainer , mediaLinkGroup byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.MediaLinkGroup = mediaLinkGroup
}

func (m *_MediaTransportControlDataSourcePowerControl)  GetParent() MediaTransportControlData {
	return m._MediaTransportControlData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataSourcePowerControl) GetState() byte {
	return m.State
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MediaTransportControlDataSourcePowerControl) GetIsShouldPowerOn() bool {
	return bool(bool((m.GetState()) == (0x00)))
}

func (m *_MediaTransportControlDataSourcePowerControl) GetIsShouldPowerOff() bool {
	return bool(bool((m.GetState()) != (0x00)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewMediaTransportControlDataSourcePowerControl factory function for _MediaTransportControlDataSourcePowerControl
func NewMediaTransportControlDataSourcePowerControl( state byte , commandTypeContainer MediaTransportControlCommandTypeContainer , mediaLinkGroup byte ) *_MediaTransportControlDataSourcePowerControl {
	_result := &_MediaTransportControlDataSourcePowerControl{
		State: state,
    	_MediaTransportControlData: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result._MediaTransportControlData._MediaTransportControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataSourcePowerControl(structType interface{}) MediaTransportControlDataSourcePowerControl {
    if casted, ok := structType.(MediaTransportControlDataSourcePowerControl); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataSourcePowerControl); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataSourcePowerControl) GetTypeName() string {
	return "MediaTransportControlDataSourcePowerControl"
}

func (m *_MediaTransportControlDataSourcePowerControl) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_MediaTransportControlDataSourcePowerControl) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (state)
	lengthInBits += 8;

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_MediaTransportControlDataSourcePowerControl) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MediaTransportControlDataSourcePowerControlParse(theBytes []byte) (MediaTransportControlDataSourcePowerControl, error) {
	return MediaTransportControlDataSourcePowerControlParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func MediaTransportControlDataSourcePowerControlParseWithBuffer(readBuffer utils.ReadBuffer) (MediaTransportControlDataSourcePowerControl, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataSourcePowerControl"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataSourcePowerControl")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (state)
_state, _stateErr := readBuffer.ReadByte("state")
	if _stateErr != nil {
		return nil, errors.Wrap(_stateErr, "Error parsing 'state' field of MediaTransportControlDataSourcePowerControl")
	}
	state := _state

	// Virtual field
	_isShouldPowerOn := bool((state) == (0x00))
	isShouldPowerOn := bool(_isShouldPowerOn)
	_ = isShouldPowerOn

	// Virtual field
	_isShouldPowerOff := bool((state) != (0x00))
	isShouldPowerOff := bool(_isShouldPowerOff)
	_ = isShouldPowerOff

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataSourcePowerControl"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataSourcePowerControl")
	}

	// Create a partially initialized instance
	_child := &_MediaTransportControlDataSourcePowerControl{
		_MediaTransportControlData: &_MediaTransportControlData{
		},
		State: state,
	}
	_child._MediaTransportControlData._MediaTransportControlDataChildRequirements = _child
	return _child, nil
}

func (m *_MediaTransportControlDataSourcePowerControl) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MediaTransportControlDataSourcePowerControl) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataSourcePowerControl"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataSourcePowerControl")
		}

	// Simple Field (state)
	state := byte(m.GetState())
	_stateErr := writeBuffer.WriteByte("state", (state))
	if _stateErr != nil {
		return errors.Wrap(_stateErr, "Error serializing 'state' field")
	}
	// Virtual field
	if _isShouldPowerOnErr := writeBuffer.WriteVirtual("isShouldPowerOn", m.GetIsShouldPowerOn()); _isShouldPowerOnErr != nil {
		return errors.Wrap(_isShouldPowerOnErr, "Error serializing 'isShouldPowerOn' field")
	}
	// Virtual field
	if _isShouldPowerOffErr := writeBuffer.WriteVirtual("isShouldPowerOff", m.GetIsShouldPowerOff()); _isShouldPowerOffErr != nil {
		return errors.Wrap(_isShouldPowerOffErr, "Error serializing 'isShouldPowerOff' field")
	}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataSourcePowerControl"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataSourcePowerControl")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_MediaTransportControlDataSourcePowerControl) isMediaTransportControlDataSourcePowerControl() bool {
	return true
}

func (m *_MediaTransportControlDataSourcePowerControl) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



