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

// StatusResult is the corresponding interface of StatusResult
type StatusResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetDiagnosticInfo returns DiagnosticInfo (property field)
	GetDiagnosticInfo() DiagnosticInfo
}

// StatusResultExactly can be used when we want exactly this type and not a type which fulfills StatusResult.
// This is useful for switch cases.
type StatusResultExactly interface {
	StatusResult
	isStatusResult() bool
}

// _StatusResult is the data-structure of this message
type _StatusResult struct {
	*_ExtensionObjectDefinition
	StatusCode     StatusCode
	DiagnosticInfo DiagnosticInfo
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_StatusResult) GetIdentifier() string {
	return "301"
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_StatusResult) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_StatusResult) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StatusResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_StatusResult) GetDiagnosticInfo() DiagnosticInfo {
	return m.DiagnosticInfo
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewStatusResult factory function for _StatusResult
func NewStatusResult(statusCode StatusCode, diagnosticInfo DiagnosticInfo) *_StatusResult {
	_result := &_StatusResult{
		StatusCode:                 statusCode,
		DiagnosticInfo:             diagnosticInfo,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastStatusResult(structType any) StatusResult {
	if casted, ok := structType.(StatusResult); ok {
		return casted
	}
	if casted, ok := structType.(*StatusResult); ok {
		return *casted
	}
	return nil
}

func (m *_StatusResult) GetTypeName() string {
	return "StatusResult"
}

func (m *_StatusResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (diagnosticInfo)
	lengthInBits += m.DiagnosticInfo.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_StatusResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func StatusResultParse(ctx context.Context, theBytes []byte, identifier string) (StatusResult, error) {
	return StatusResultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func StatusResultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (StatusResult, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("StatusResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StatusResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (statusCode)
	if pullErr := readBuffer.PullContext("statusCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusCode")
	}
	_statusCode, _statusCodeErr := StatusCodeParseWithBuffer(ctx, readBuffer)
	if _statusCodeErr != nil {
		return nil, errors.Wrap(_statusCodeErr, "Error parsing 'statusCode' field of StatusResult")
	}
	statusCode := _statusCode.(StatusCode)
	if closeErr := readBuffer.CloseContext("statusCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusCode")
	}

	// Simple Field (diagnosticInfo)
	if pullErr := readBuffer.PullContext("diagnosticInfo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for diagnosticInfo")
	}
	_diagnosticInfo, _diagnosticInfoErr := DiagnosticInfoParseWithBuffer(ctx, readBuffer)
	if _diagnosticInfoErr != nil {
		return nil, errors.Wrap(_diagnosticInfoErr, "Error parsing 'diagnosticInfo' field of StatusResult")
	}
	diagnosticInfo := _diagnosticInfo.(DiagnosticInfo)
	if closeErr := readBuffer.CloseContext("diagnosticInfo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for diagnosticInfo")
	}

	if closeErr := readBuffer.CloseContext("StatusResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StatusResult")
	}

	// Create a partially initialized instance
	_child := &_StatusResult{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		StatusCode:                 statusCode,
		DiagnosticInfo:             diagnosticInfo,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_StatusResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_StatusResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("StatusResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for StatusResult")
		}

		// Simple Field (statusCode)
		if pushErr := writeBuffer.PushContext("statusCode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusCode")
		}
		_statusCodeErr := writeBuffer.WriteSerializable(ctx, m.GetStatusCode())
		if popErr := writeBuffer.PopContext("statusCode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusCode")
		}
		if _statusCodeErr != nil {
			return errors.Wrap(_statusCodeErr, "Error serializing 'statusCode' field")
		}

		// Simple Field (diagnosticInfo)
		if pushErr := writeBuffer.PushContext("diagnosticInfo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for diagnosticInfo")
		}
		_diagnosticInfoErr := writeBuffer.WriteSerializable(ctx, m.GetDiagnosticInfo())
		if popErr := writeBuffer.PopContext("diagnosticInfo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for diagnosticInfo")
		}
		if _diagnosticInfoErr != nil {
			return errors.Wrap(_diagnosticInfoErr, "Error serializing 'diagnosticInfo' field")
		}

		if popErr := writeBuffer.PopContext("StatusResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for StatusResult")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_StatusResult) isStatusResult() bool {
	return true
}

func (m *_StatusResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
