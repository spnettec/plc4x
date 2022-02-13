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

// The data-structure of this message
type KnxNetRemoteConfigurationAndDiagnosis struct {
	*ServiceId
	Version uint8
}

// The corresponding interface
type IKnxNetRemoteConfigurationAndDiagnosis interface {
	// GetVersion returns Version
	GetVersion() uint8
	// LengthInBytes returns the length in bytes
	LengthInBytes() uint16
	// LengthInBits returns the length in bits
	LengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
// Accessors for discriminator values.
///////////////////////////////////////////////////////////
func (m *KnxNetRemoteConfigurationAndDiagnosis) ServiceType() uint8 {
	return 0x07
}

func (m *KnxNetRemoteConfigurationAndDiagnosis) GetServiceType() uint8 {
	return 0x07
}

func (m *KnxNetRemoteConfigurationAndDiagnosis) InitializeParent(parent *ServiceId) {}

///////////////////////////////////////////////////////////
// Accessors for property fields.
///////////////////////////////////////////////////////////
func (m *KnxNetRemoteConfigurationAndDiagnosis) GetVersion() uint8 {
	return m.Version
}

///////////////////////////////////////////////////////////
// Accessors for virtual fields.
///////////////////////////////////////////////////////////

func NewKnxNetRemoteConfigurationAndDiagnosis(version uint8) *ServiceId {
	child := &KnxNetRemoteConfigurationAndDiagnosis{
		Version:   version,
		ServiceId: NewServiceId(),
	}
	child.Child = child
	return child.ServiceId
}

func CastKnxNetRemoteConfigurationAndDiagnosis(structType interface{}) *KnxNetRemoteConfigurationAndDiagnosis {
	castFunc := func(typ interface{}) *KnxNetRemoteConfigurationAndDiagnosis {
		if casted, ok := typ.(KnxNetRemoteConfigurationAndDiagnosis); ok {
			return &casted
		}
		if casted, ok := typ.(*KnxNetRemoteConfigurationAndDiagnosis); ok {
			return casted
		}
		if casted, ok := typ.(ServiceId); ok {
			return CastKnxNetRemoteConfigurationAndDiagnosis(casted.Child)
		}
		if casted, ok := typ.(*ServiceId); ok {
			return CastKnxNetRemoteConfigurationAndDiagnosis(casted.Child)
		}
		return nil
	}
	return castFunc(structType)
}

func (m *KnxNetRemoteConfigurationAndDiagnosis) GetTypeName() string {
	return "KnxNetRemoteConfigurationAndDiagnosis"
}

func (m *KnxNetRemoteConfigurationAndDiagnosis) LengthInBits() uint16 {
	return m.LengthInBitsConditional(false)
}

func (m *KnxNetRemoteConfigurationAndDiagnosis) LengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.ParentLengthInBits())

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *KnxNetRemoteConfigurationAndDiagnosis) LengthInBytes() uint16 {
	return m.LengthInBits() / 8
}

func KnxNetRemoteConfigurationAndDiagnosisParse(readBuffer utils.ReadBuffer) (*ServiceId, error) {
	if pullErr := readBuffer.PullContext("KnxNetRemoteConfigurationAndDiagnosis"); pullErr != nil {
		return nil, pullErr
	}

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field")
	}
	version := _version

	if closeErr := readBuffer.CloseContext("KnxNetRemoteConfigurationAndDiagnosis"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &KnxNetRemoteConfigurationAndDiagnosis{
		Version:   version,
		ServiceId: &ServiceId{},
	}
	_child.ServiceId.Child = _child
	return _child.ServiceId, nil
}

func (m *KnxNetRemoteConfigurationAndDiagnosis) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetRemoteConfigurationAndDiagnosis"); pushErr != nil {
			return pushErr
		}

		// Simple Field (version)
		version := uint8(m.Version)
		_versionErr := writeBuffer.WriteUint8("version", 8, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetRemoteConfigurationAndDiagnosis"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *KnxNetRemoteConfigurationAndDiagnosis) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	m.Serialize(buffer)
	return buffer.GetBox().String()
}
