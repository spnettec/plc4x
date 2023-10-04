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
	"context"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

// Code generated by code-generation. DO NOT EDIT.

// KnxNetRemoteConfigurationAndDiagnosis is the corresponding interface of KnxNetRemoteConfigurationAndDiagnosis
type KnxNetRemoteConfigurationAndDiagnosis interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ServiceId
	// GetVersion returns Version (property field)
	GetVersion() uint8
}

// KnxNetRemoteConfigurationAndDiagnosisExactly can be used when we want exactly this type and not a type which fulfills KnxNetRemoteConfigurationAndDiagnosis.
// This is useful for switch cases.
type KnxNetRemoteConfigurationAndDiagnosisExactly interface {
	KnxNetRemoteConfigurationAndDiagnosis
	isKnxNetRemoteConfigurationAndDiagnosis() bool
}

// _KnxNetRemoteConfigurationAndDiagnosis is the data-structure of this message
type _KnxNetRemoteConfigurationAndDiagnosis struct {
	*_ServiceId
	Version uint8
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_KnxNetRemoteConfigurationAndDiagnosis) GetServiceType() uint8 {
	return 0x07
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_KnxNetRemoteConfigurationAndDiagnosis) InitializeParent(parent ServiceId) {}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) GetParent() ServiceId {
	return m._ServiceId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_KnxNetRemoteConfigurationAndDiagnosis) GetVersion() uint8 {
	return m.Version
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewKnxNetRemoteConfigurationAndDiagnosis factory function for _KnxNetRemoteConfigurationAndDiagnosis
func NewKnxNetRemoteConfigurationAndDiagnosis(version uint8) *_KnxNetRemoteConfigurationAndDiagnosis {
	_result := &_KnxNetRemoteConfigurationAndDiagnosis{
		Version:    version,
		_ServiceId: NewServiceId(),
	}
	_result._ServiceId._ServiceIdChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastKnxNetRemoteConfigurationAndDiagnosis(structType any) KnxNetRemoteConfigurationAndDiagnosis {
	if casted, ok := structType.(KnxNetRemoteConfigurationAndDiagnosis); ok {
		return casted
	}
	if casted, ok := structType.(*KnxNetRemoteConfigurationAndDiagnosis); ok {
		return *casted
	}
	return nil
}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) GetTypeName() string {
	return "KnxNetRemoteConfigurationAndDiagnosis"
}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (version)
	lengthInBits += 8

	return lengthInBits
}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func KnxNetRemoteConfigurationAndDiagnosisParse(ctx context.Context, theBytes []byte) (KnxNetRemoteConfigurationAndDiagnosis, error) {
	return KnxNetRemoteConfigurationAndDiagnosisParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func KnxNetRemoteConfigurationAndDiagnosisParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (KnxNetRemoteConfigurationAndDiagnosis, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("KnxNetRemoteConfigurationAndDiagnosis"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for KnxNetRemoteConfigurationAndDiagnosis")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (version)
	_version, _versionErr := readBuffer.ReadUint8("version", 8)
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field of KnxNetRemoteConfigurationAndDiagnosis")
	}
	version := _version

	if closeErr := readBuffer.CloseContext("KnxNetRemoteConfigurationAndDiagnosis"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for KnxNetRemoteConfigurationAndDiagnosis")
	}

	// Create a partially initialized instance
	_child := &_KnxNetRemoteConfigurationAndDiagnosis{
		_ServiceId: &_ServiceId{},
		Version:    version,
	}
	_child._ServiceId._ServiceIdChildRequirements = _child
	return _child, nil
}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("KnxNetRemoteConfigurationAndDiagnosis"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for KnxNetRemoteConfigurationAndDiagnosis")
		}

		// Simple Field (version)
		version := uint8(m.GetVersion())
		_versionErr := writeBuffer.WriteUint8("version", 8, (version))
		if _versionErr != nil {
			return errors.Wrap(_versionErr, "Error serializing 'version' field")
		}

		if popErr := writeBuffer.PopContext("KnxNetRemoteConfigurationAndDiagnosis"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for KnxNetRemoteConfigurationAndDiagnosis")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) isKnxNetRemoteConfigurationAndDiagnosis() bool {
	return true
}

func (m *_KnxNetRemoteConfigurationAndDiagnosis) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
