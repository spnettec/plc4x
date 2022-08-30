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


// AccessControlDataAccessPointClosed is the corresponding interface of AccessControlDataAccessPointClosed
type AccessControlDataAccessPointClosed interface {
	utils.LengthAware
	utils.Serializable
	AccessControlData
}

// AccessControlDataAccessPointClosedExactly can be used when we want exactly this type and not a type which fulfills AccessControlDataAccessPointClosed.
// This is useful for switch cases.
type AccessControlDataAccessPointClosedExactly interface {
	AccessControlDataAccessPointClosed
	isAccessControlDataAccessPointClosed() bool
}

// _AccessControlDataAccessPointClosed is the data-structure of this message
type _AccessControlDataAccessPointClosed struct {
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

func (m *_AccessControlDataAccessPointClosed) InitializeParent(parent AccessControlData , commandTypeContainer AccessControlCommandTypeContainer , networkId byte , accessPointId byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.NetworkId = networkId
	m.AccessPointId = accessPointId
}

func (m *_AccessControlDataAccessPointClosed)  GetParent() AccessControlData {
	return m._AccessControlData
}


// NewAccessControlDataAccessPointClosed factory function for _AccessControlDataAccessPointClosed
func NewAccessControlDataAccessPointClosed( commandTypeContainer AccessControlCommandTypeContainer , networkId byte , accessPointId byte ) *_AccessControlDataAccessPointClosed {
	_result := &_AccessControlDataAccessPointClosed{
    	_AccessControlData: NewAccessControlData(commandTypeContainer, networkId, accessPointId),
	}
	_result._AccessControlData._AccessControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastAccessControlDataAccessPointClosed(structType interface{}) AccessControlDataAccessPointClosed {
    if casted, ok := structType.(AccessControlDataAccessPointClosed); ok {
		return casted
	}
	if casted, ok := structType.(*AccessControlDataAccessPointClosed); ok {
		return *casted
	}
	return nil
}

func (m *_AccessControlDataAccessPointClosed) GetTypeName() string {
	return "AccessControlDataAccessPointClosed"
}

func (m *_AccessControlDataAccessPointClosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_AccessControlDataAccessPointClosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	return lengthInBits
}


func (m *_AccessControlDataAccessPointClosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func AccessControlDataAccessPointClosedParse(readBuffer utils.ReadBuffer) (AccessControlDataAccessPointClosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("AccessControlDataAccessPointClosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for AccessControlDataAccessPointClosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	if closeErr := readBuffer.CloseContext("AccessControlDataAccessPointClosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for AccessControlDataAccessPointClosed")
	}

	// Create a partially initialized instance
	_child := &_AccessControlDataAccessPointClosed{
		_AccessControlData: &_AccessControlData{
		},
	}
	_child._AccessControlData._AccessControlDataChildRequirements = _child
	return _child, nil
}

func (m *_AccessControlDataAccessPointClosed) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("AccessControlDataAccessPointClosed"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for AccessControlDataAccessPointClosed")
		}

		if popErr := writeBuffer.PopContext("AccessControlDataAccessPointClosed"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for AccessControlDataAccessPointClosed")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_AccessControlDataAccessPointClosed) isAccessControlDataAccessPointClosed() bool {
	return true
}

func (m *_AccessControlDataAccessPointClosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



