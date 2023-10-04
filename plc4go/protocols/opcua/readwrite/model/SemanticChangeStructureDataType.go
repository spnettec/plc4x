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

// SemanticChangeStructureDataType is the corresponding interface of SemanticChangeStructureDataType
type SemanticChangeStructureDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetAffected returns Affected (property field)
	GetAffected() NodeId
	// GetAffectedType returns AffectedType (property field)
	GetAffectedType() NodeId
}

// SemanticChangeStructureDataTypeExactly can be used when we want exactly this type and not a type which fulfills SemanticChangeStructureDataType.
// This is useful for switch cases.
type SemanticChangeStructureDataTypeExactly interface {
	SemanticChangeStructureDataType
	isSemanticChangeStructureDataType() bool
}

// _SemanticChangeStructureDataType is the data-structure of this message
type _SemanticChangeStructureDataType struct {
	*_ExtensionObjectDefinition
	Affected     NodeId
	AffectedType NodeId
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SemanticChangeStructureDataType) GetIdentifier() string {
	return "899"
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SemanticChangeStructureDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_SemanticChangeStructureDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SemanticChangeStructureDataType) GetAffected() NodeId {
	return m.Affected
}

func (m *_SemanticChangeStructureDataType) GetAffectedType() NodeId {
	return m.AffectedType
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSemanticChangeStructureDataType factory function for _SemanticChangeStructureDataType
func NewSemanticChangeStructureDataType(affected NodeId, affectedType NodeId) *_SemanticChangeStructureDataType {
	_result := &_SemanticChangeStructureDataType{
		Affected:                   affected,
		AffectedType:               affectedType,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSemanticChangeStructureDataType(structType any) SemanticChangeStructureDataType {
	if casted, ok := structType.(SemanticChangeStructureDataType); ok {
		return casted
	}
	if casted, ok := structType.(*SemanticChangeStructureDataType); ok {
		return *casted
	}
	return nil
}

func (m *_SemanticChangeStructureDataType) GetTypeName() string {
	return "SemanticChangeStructureDataType"
}

func (m *_SemanticChangeStructureDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (affected)
	lengthInBits += m.Affected.GetLengthInBits(ctx)

	// Simple field (affectedType)
	lengthInBits += m.AffectedType.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SemanticChangeStructureDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SemanticChangeStructureDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (SemanticChangeStructureDataType, error) {
	return SemanticChangeStructureDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func SemanticChangeStructureDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (SemanticChangeStructureDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SemanticChangeStructureDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SemanticChangeStructureDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (affected)
	if pullErr := readBuffer.PullContext("affected"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for affected")
	}
	_affected, _affectedErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _affectedErr != nil {
		return nil, errors.Wrap(_affectedErr, "Error parsing 'affected' field of SemanticChangeStructureDataType")
	}
	affected := _affected.(NodeId)
	if closeErr := readBuffer.CloseContext("affected"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for affected")
	}

	// Simple Field (affectedType)
	if pullErr := readBuffer.PullContext("affectedType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for affectedType")
	}
	_affectedType, _affectedTypeErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _affectedTypeErr != nil {
		return nil, errors.Wrap(_affectedTypeErr, "Error parsing 'affectedType' field of SemanticChangeStructureDataType")
	}
	affectedType := _affectedType.(NodeId)
	if closeErr := readBuffer.CloseContext("affectedType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for affectedType")
	}

	if closeErr := readBuffer.CloseContext("SemanticChangeStructureDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SemanticChangeStructureDataType")
	}

	// Create a partially initialized instance
	_child := &_SemanticChangeStructureDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		Affected:                   affected,
		AffectedType:               affectedType,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_SemanticChangeStructureDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SemanticChangeStructureDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SemanticChangeStructureDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SemanticChangeStructureDataType")
		}

		// Simple Field (affected)
		if pushErr := writeBuffer.PushContext("affected"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for affected")
		}
		_affectedErr := writeBuffer.WriteSerializable(ctx, m.GetAffected())
		if popErr := writeBuffer.PopContext("affected"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for affected")
		}
		if _affectedErr != nil {
			return errors.Wrap(_affectedErr, "Error serializing 'affected' field")
		}

		// Simple Field (affectedType)
		if pushErr := writeBuffer.PushContext("affectedType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for affectedType")
		}
		_affectedTypeErr := writeBuffer.WriteSerializable(ctx, m.GetAffectedType())
		if popErr := writeBuffer.PopContext("affectedType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for affectedType")
		}
		if _affectedTypeErr != nil {
			return errors.Wrap(_affectedTypeErr, "Error serializing 'affectedType' field")
		}

		if popErr := writeBuffer.PopContext("SemanticChangeStructureDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SemanticChangeStructureDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SemanticChangeStructureDataType) isSemanticChangeStructureDataType() bool {
	return true
}

func (m *_SemanticChangeStructureDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
