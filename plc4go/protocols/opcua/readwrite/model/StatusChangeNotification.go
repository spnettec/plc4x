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


// StatusChangeNotification is the corresponding interface of StatusChangeNotification
type StatusChangeNotification interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStatus returns Status (property field)
	GetStatus() StatusCode
	// GetDiagnosticInfo returns DiagnosticInfo (property field)
	GetDiagnosticInfo() DiagnosticInfo
}

// StatusChangeNotificationExactly can be used when we want exactly this type and not a type which fulfills StatusChangeNotification.
// This is useful for switch cases.
type StatusChangeNotificationExactly interface {
	StatusChangeNotification
	isStatusChangeNotification() bool
}

// _StatusChangeNotification is the data-structure of this message
type _StatusChangeNotification struct {
	*_ExtensionObjectDefinition
        Status StatusCode
        DiagnosticInfo DiagnosticInfo
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_StatusChangeNotification)  GetIdentifier() string {
return "820"}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_StatusChangeNotification) InitializeParent(parent ExtensionObjectDefinition ) {}

func (m *_StatusChangeNotification)  GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_StatusChangeNotification) GetStatus() StatusCode {
	return m.Status
}

func (m *_StatusChangeNotification) GetDiagnosticInfo() DiagnosticInfo {
	return m.DiagnosticInfo
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewStatusChangeNotification factory function for _StatusChangeNotification
func NewStatusChangeNotification( status StatusCode , diagnosticInfo DiagnosticInfo ) *_StatusChangeNotification {
	_result := &_StatusChangeNotification{
		Status: status,
		DiagnosticInfo: diagnosticInfo,
    	_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastStatusChangeNotification(structType any) StatusChangeNotification {
    if casted, ok := structType.(StatusChangeNotification); ok {
		return casted
	}
	if casted, ok := structType.(*StatusChangeNotification); ok {
		return *casted
	}
	return nil
}

func (m *_StatusChangeNotification) GetTypeName() string {
	return "StatusChangeNotification"
}

func (m *_StatusChangeNotification) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Implicit Field (notificationLength)
	lengthInBits += 32

	// Simple field (status)
	lengthInBits += m.Status.GetLengthInBits(ctx)

	// Simple field (diagnosticInfo)
	lengthInBits += m.DiagnosticInfo.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_StatusChangeNotification) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func StatusChangeNotificationParse(ctx context.Context, theBytes []byte, identifier string) (StatusChangeNotification, error) {
	return StatusChangeNotificationParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func StatusChangeNotificationParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (StatusChangeNotification, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("StatusChangeNotification"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for StatusChangeNotification")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Implicit Field (notificationLength) (Used for parsing, but its value is not stored as it's implicitly given by the objects content)
	notificationLength, _notificationLengthErr := readBuffer.ReadInt32("notificationLength", 32)
	_ = notificationLength
	if _notificationLengthErr != nil {
		return nil, errors.Wrap(_notificationLengthErr, "Error parsing 'notificationLength' field of StatusChangeNotification")
	}

	// Simple Field (status)
	if pullErr := readBuffer.PullContext("status"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for status")
	}
_status, _statusErr := StatusCodeParseWithBuffer(ctx, readBuffer)
	if _statusErr != nil {
		return nil, errors.Wrap(_statusErr, "Error parsing 'status' field of StatusChangeNotification")
	}
	status := _status.(StatusCode)
	if closeErr := readBuffer.CloseContext("status"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for status")
	}

	// Simple Field (diagnosticInfo)
	if pullErr := readBuffer.PullContext("diagnosticInfo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for diagnosticInfo")
	}
_diagnosticInfo, _diagnosticInfoErr := DiagnosticInfoParseWithBuffer(ctx, readBuffer)
	if _diagnosticInfoErr != nil {
		return nil, errors.Wrap(_diagnosticInfoErr, "Error parsing 'diagnosticInfo' field of StatusChangeNotification")
	}
	diagnosticInfo := _diagnosticInfo.(DiagnosticInfo)
	if closeErr := readBuffer.CloseContext("diagnosticInfo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for diagnosticInfo")
	}

	if closeErr := readBuffer.CloseContext("StatusChangeNotification"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for StatusChangeNotification")
	}

	// Create a partially initialized instance
	_child := &_StatusChangeNotification{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{
		},
		Status: status,
		DiagnosticInfo: diagnosticInfo,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_StatusChangeNotification) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_StatusChangeNotification) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("StatusChangeNotification"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for StatusChangeNotification")
		}

	// Implicit Field (notificationLength) (Used for parsing, but it's value is not stored as it's implicitly given by the objects content)
	notificationLength := int32(int32(m.GetLengthInBytes(ctx)))
	_notificationLengthErr := writeBuffer.WriteInt32("notificationLength", 32, (notificationLength))
	if _notificationLengthErr != nil {
		return errors.Wrap(_notificationLengthErr, "Error serializing 'notificationLength' field")
	}

	// Simple Field (status)
	if pushErr := writeBuffer.PushContext("status"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for status")
	}
	_statusErr := writeBuffer.WriteSerializable(ctx, m.GetStatus())
	if popErr := writeBuffer.PopContext("status"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for status")
	}
	if _statusErr != nil {
		return errors.Wrap(_statusErr, "Error serializing 'status' field")
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

		if popErr := writeBuffer.PopContext("StatusChangeNotification"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for StatusChangeNotification")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_StatusChangeNotification) isStatusChangeNotification() bool {
	return true
}

func (m *_StatusChangeNotification) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



