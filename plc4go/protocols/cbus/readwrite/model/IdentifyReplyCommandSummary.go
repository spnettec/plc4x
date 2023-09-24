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


// IdentifyReplyCommandSummary is the corresponding interface of IdentifyReplyCommandSummary
type IdentifyReplyCommandSummary interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetPartName returns PartName (property field)
	GetPartName() string
	// GetUnitServiceType returns UnitServiceType (property field)
	GetUnitServiceType() byte
	// GetVersion returns Version (property field)
	GetVersion() string
}

// IdentifyReplyCommandSummaryExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandSummary.
// This is useful for switch cases.
type IdentifyReplyCommandSummaryExactly interface {
	IdentifyReplyCommandSummary
	isIdentifyReplyCommandSummary() bool
}

// _IdentifyReplyCommandSummary is the data-structure of this message
type _IdentifyReplyCommandSummary struct {
	*_IdentifyReplyCommand
        PartName string
        UnitServiceType byte
        Version string
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandSummary)  GetAttribute() Attribute {
return Attribute_Summary}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandSummary) InitializeParent(parent IdentifyReplyCommand ) {}

func (m *_IdentifyReplyCommandSummary)  GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandSummary) GetPartName() string {
	return m.PartName
}

func (m *_IdentifyReplyCommandSummary) GetUnitServiceType() byte {
	return m.UnitServiceType
}

func (m *_IdentifyReplyCommandSummary) GetVersion() string {
	return m.Version
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewIdentifyReplyCommandSummary factory function for _IdentifyReplyCommandSummary
func NewIdentifyReplyCommandSummary( partName string , unitServiceType byte , version string , numBytes uint8 ) *_IdentifyReplyCommandSummary {
	_result := &_IdentifyReplyCommandSummary{
		PartName: partName,
		UnitServiceType: unitServiceType,
		Version: version,
    	_IdentifyReplyCommand: NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandSummary(structType any) IdentifyReplyCommandSummary {
    if casted, ok := structType.(IdentifyReplyCommandSummary); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandSummary); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandSummary) GetTypeName() string {
	return "IdentifyReplyCommandSummary"
}

func (m *_IdentifyReplyCommandSummary) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (partName)
	lengthInBits += 48;

	// Simple field (unitServiceType)
	lengthInBits += 8;

	// Simple field (version)
	lengthInBits += 32;

	return lengthInBits
}


func (m *_IdentifyReplyCommandSummary) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IdentifyReplyCommandSummaryParse(ctx context.Context, theBytes []byte, attribute Attribute, numBytes uint8) (IdentifyReplyCommandSummary, error) {
	return IdentifyReplyCommandSummaryParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), attribute, numBytes)
}

func IdentifyReplyCommandSummaryParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandSummary, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandSummary"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandSummary")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (partName)
_partName, _partNameErr := readBuffer.ReadString("partName", uint32(48), "UTF-8")
	if _partNameErr != nil {
		return nil, errors.Wrap(_partNameErr, "Error parsing 'partName' field of IdentifyReplyCommandSummary")
	}
	partName := _partName

	// Simple Field (unitServiceType)
_unitServiceType, _unitServiceTypeErr := readBuffer.ReadByte("unitServiceType")
	if _unitServiceTypeErr != nil {
		return nil, errors.Wrap(_unitServiceTypeErr, "Error parsing 'unitServiceType' field of IdentifyReplyCommandSummary")
	}
	unitServiceType := _unitServiceType

	// Simple Field (version)
_version, _versionErr := readBuffer.ReadString("version", uint32(32), "UTF-8")
	if _versionErr != nil {
		return nil, errors.Wrap(_versionErr, "Error parsing 'version' field of IdentifyReplyCommandSummary")
	}
	version := _version

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandSummary"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandSummary")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandSummary{
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
		PartName: partName,
		UnitServiceType: unitServiceType,
		Version: version,
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandSummary) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandSummary) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandSummary"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandSummary")
		}

	// Simple Field (partName)
	partName := string(m.GetPartName())
	_partNameErr := writeBuffer.WriteString("partName", uint32(48), "UTF-8", (partName))
	if _partNameErr != nil {
		return errors.Wrap(_partNameErr, "Error serializing 'partName' field")
	}

	// Simple Field (unitServiceType)
	unitServiceType := byte(m.GetUnitServiceType())
	_unitServiceTypeErr := writeBuffer.WriteByte("unitServiceType", (unitServiceType))
	if _unitServiceTypeErr != nil {
		return errors.Wrap(_unitServiceTypeErr, "Error serializing 'unitServiceType' field")
	}

	// Simple Field (version)
	version := string(m.GetVersion())
	_versionErr := writeBuffer.WriteString("version", uint32(32), "UTF-8", (version))
	if _versionErr != nil {
		return errors.Wrap(_versionErr, "Error serializing 'version' field")
	}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandSummary"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandSummary")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_IdentifyReplyCommandSummary) isIdentifyReplyCommandSummary() bool {
	return true
}

func (m *_IdentifyReplyCommandSummary) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



