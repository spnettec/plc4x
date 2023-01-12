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


// AccessControlDataAccessPointLeftOpen is the corresponding interface of AccessControlDataAccessPointLeftOpen
type AccessControlDataAccessPointLeftOpen interface {
	utils.LengthAware
	utils.Serializable
	AccessControlData
}

// AccessControlDataAccessPointLeftOpenExactly can be used when we want exactly this type and not a type which fulfills AccessControlDataAccessPointLeftOpen.
// This is useful for switch cases.
type AccessControlDataAccessPointLeftOpenExactly interface {
	AccessControlDataAccessPointLeftOpen
	isAccessControlDataAccessPointLeftOpen() bool
}

// _AccessControlDataAccessPointLeftOpen is the data-structure of this message
type _AccessControlDataAccessPointLeftOpen struct {
	*_AccessControlData
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_AccessControlDataAccessPointLeftOpen) InitializeParent(parent AccessControlData , commandTypeContainer AccessControlCommandTypeContainer , networkId byte , accessPointId byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.NetworkId = networkId
	m.AccessPointId = accessPointId
}

func (m *_AccessControlDataAccessPointLeftOpen)  GetParent() AccessControlData {
	return m._AccessControlData
}


// NewAccessControlDataAccessPointLeftOpen factory function for _AccessControlDataAccessPointLeftOpen
func NewAccessControlDataAccessPointLeftOpen( commandTypeContainer AccessControlCommandTypeContainer , networkId byte , accessPointId byte ) *_AccessControlDataAccessPointLeftOpen {
	_result := &_AccessControlDataAccessPointLeftOpen{
    	_AccessControlData: NewAccessControlData(commandTypeContainer, networkId, accessPointId),
	}
	_result._AccessControlData._AccessControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAccessControlDataAccessPointLeftOpen(structType interface{}) AccessControlDataAccessPointLeftOpen {
    if casted, ok := structType.(AccessControlDataAccessPointLeftOpen); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlDataAccessPointLeftOpen); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlDataAccessPointLeftOpen) GetTypeName() string {
	return "AccessControlDataAccessPointLeftOpen"
}

func (m *_AccessControlDataAccessPointLeftOpen) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AccessControlDataAccessPointLeftOpen) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}


func (m *_AccessControlDataAccessPointLeftOpen) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AccessControlDataAccessPointLeftOpenParse(theBytes []byte) (AccessControlDataAccessPointLeftOpen, error) {
	return AccessControlDataAccessPointLeftOpenParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func AccessControlDataAccessPointLeftOpenParseWithBuffer(readBuffer utils.ReadBuffer) (AccessControlDataAccessPointLeftOpen, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AccessControlDataAccessPointLeftOpen"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlDataAccessPointLeftOpen")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AccessControlDataAccessPointLeftOpen"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlDataAccessPointLeftOpen")
	}

	// Create a partially initialized instance
	_child := &_AccessControlDataAccessPointLeftOpen{
		_AccessControlData: &_AccessControlData{
		},
	}
	_child._AccessControlData._AccessControlDataChildRequirements = _child
	return _child, nil
}

func (m *_AccessControlDataAccessPointLeftOpen) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_AccessControlDataAccessPointLeftOpen) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AccessControlDataAccessPointLeftOpen"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AccessControlDataAccessPointLeftOpen")
		}

		if popErr := writeBuffer.PopContext("AccessControlDataAccessPointLeftOpen"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AccessControlDataAccessPointLeftOpen")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_AccessControlDataAccessPointLeftOpen) isAccessControlDataAccessPointLeftOpen() bool {
	return true
}

func (m *_AccessControlDataAccessPointLeftOpen) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



