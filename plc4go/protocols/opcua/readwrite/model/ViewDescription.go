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

// ViewDescription is the corresponding interface of ViewDescription
type ViewDescription interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetViewId returns ViewId (property field)
	GetViewId() NodeId
	// GetTimestamp returns Timestamp (property field)
	GetTimestamp() int64
	// GetViewVersion returns ViewVersion (property field)
	GetViewVersion() uint32
}

// ViewDescriptionExactly can be used when we want exactly this type and not a type which fulfills ViewDescription.
// This is useful for switch cases.
type ViewDescriptionExactly interface {
	ViewDescription
	isViewDescription() bool
}

// _ViewDescription is the data-structure of this message
type _ViewDescription struct {
	*_ExtensionObjectDefinition
	ViewId      NodeId
	Timestamp   int64
	ViewVersion uint32
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ViewDescription) GetIdentifier() string {
	return "513"
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ViewDescription) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_ViewDescription) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ViewDescription) GetViewId() NodeId {
	return m.ViewId
}

func (m *_ViewDescription) GetTimestamp() int64 {
	return m.Timestamp
}

func (m *_ViewDescription) GetViewVersion() uint32 {
	return m.ViewVersion
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewViewDescription factory function for _ViewDescription
func NewViewDescription(viewId NodeId, timestamp int64, viewVersion uint32) *_ViewDescription {
	_result := &_ViewDescription{
		ViewId:                     viewId,
		Timestamp:                  timestamp,
		ViewVersion:                viewVersion,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastViewDescription(structType any) ViewDescription {
	if casted, ok := structType.(ViewDescription); ok {
		return casted
	}
	if casted, ok := structType.(*ViewDescription); ok {
		return *casted
	}
	return nil
}

func (m *_ViewDescription) GetTypeName() string {
	return "ViewDescription"
}

func (m *_ViewDescription) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (viewId)
	lengthInBits += m.ViewId.GetLengthInBits(ctx)

	// Simple field (timestamp)
	lengthInBits += 64

	// Simple field (viewVersion)
	lengthInBits += 32

	return lengthInBits
}

func (m *_ViewDescription) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func ViewDescriptionParse(ctx context.Context, theBytes []byte, identifier string) (ViewDescription, error) {
	return ViewDescriptionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func ViewDescriptionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (ViewDescription, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("ViewDescription"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ViewDescription")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (viewId)
	if pullErr := readBuffer.PullContext("viewId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for viewId")
	}
	_viewId, _viewIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _viewIdErr != nil {
		return nil, errors.Wrap(_viewIdErr, "Error parsing 'viewId' field of ViewDescription")
	}
	viewId := _viewId.(NodeId)
	if closeErr := readBuffer.CloseContext("viewId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for viewId")
	}

	// Simple Field (timestamp)
	_timestamp, _timestampErr := readBuffer.ReadInt64("timestamp", 64)
	if _timestampErr != nil {
		return nil, errors.Wrap(_timestampErr, "Error parsing 'timestamp' field of ViewDescription")
	}
	timestamp := _timestamp

	// Simple Field (viewVersion)
	_viewVersion, _viewVersionErr := readBuffer.ReadUint32("viewVersion", 32)
	if _viewVersionErr != nil {
		return nil, errors.Wrap(_viewVersionErr, "Error parsing 'viewVersion' field of ViewDescription")
	}
	viewVersion := _viewVersion

	if closeErr := readBuffer.CloseContext("ViewDescription"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ViewDescription")
	}

	// Create a partially initialized instance
	_child := &_ViewDescription{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		ViewId:                     viewId,
		Timestamp:                  timestamp,
		ViewVersion:                viewVersion,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_ViewDescription) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ViewDescription) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ViewDescription"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ViewDescription")
		}

		// Simple Field (viewId)
		if pushErr := writeBuffer.PushContext("viewId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for viewId")
		}
		_viewIdErr := writeBuffer.WriteSerializable(ctx, m.GetViewId())
		if popErr := writeBuffer.PopContext("viewId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for viewId")
		}
		if _viewIdErr != nil {
			return errors.Wrap(_viewIdErr, "Error serializing 'viewId' field")
		}

		// Simple Field (timestamp)
		timestamp := int64(m.GetTimestamp())
		_timestampErr := writeBuffer.WriteInt64("timestamp", 64, (timestamp))
		if _timestampErr != nil {
			return errors.Wrap(_timestampErr, "Error serializing 'timestamp' field")
		}

		// Simple Field (viewVersion)
		viewVersion := uint32(m.GetViewVersion())
		_viewVersionErr := writeBuffer.WriteUint32("viewVersion", 32, (viewVersion))
		if _viewVersionErr != nil {
			return errors.Wrap(_viewVersionErr, "Error serializing 'viewVersion' field")
		}

		if popErr := writeBuffer.PopContext("ViewDescription"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ViewDescription")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ViewDescription) isViewDescription() bool {
	return true
}

func (m *_ViewDescription) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
