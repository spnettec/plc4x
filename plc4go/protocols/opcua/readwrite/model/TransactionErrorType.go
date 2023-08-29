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


// TransactionErrorType is the corresponding interface of TransactionErrorType
type TransactionErrorType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetTargetId returns TargetId (property field)
	GetTargetId() NodeId
	// GetError returns Error (property field)
	GetError() StatusCode
	// GetMessage returns Message (property field)
	GetMessage() LocalizedText
}

// TransactionErrorTypeExactly can be used when we want exactly this type and not a type which fulfills TransactionErrorType.
// This is useful for switch cases.
type TransactionErrorTypeExactly interface {
	TransactionErrorType
	isTransactionErrorType() bool
}

// _TransactionErrorType is the data-structure of this message
type _TransactionErrorType struct {
	*_ExtensionObjectDefinition
        TargetId NodeId
        Error StatusCode
        Message LocalizedText
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_TransactionErrorType)  GetIdentifier() string {
return "32287"}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_TransactionErrorType) InitializeParent(parent ExtensionObjectDefinition ) {}

func (m *_TransactionErrorType)  GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_TransactionErrorType) GetTargetId() NodeId {
	return m.TargetId
}

func (m *_TransactionErrorType) GetError() StatusCode {
	return m.Error
}

func (m *_TransactionErrorType) GetMessage() LocalizedText {
	return m.Message
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewTransactionErrorType factory function for _TransactionErrorType
func NewTransactionErrorType( targetId NodeId , error StatusCode , message LocalizedText ) *_TransactionErrorType {
	_result := &_TransactionErrorType{
		TargetId: targetId,
		Error: error,
		Message: message,
    	_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastTransactionErrorType(structType any) TransactionErrorType {
    if casted, ok := structType.(TransactionErrorType); ok {
		return casted
	}
	if casted, ok := structType.(*TransactionErrorType); ok {
		return *casted
	}
	return nil
}

func (m *_TransactionErrorType) GetTypeName() string {
	return "TransactionErrorType"
}

func (m *_TransactionErrorType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (targetId)
	lengthInBits += m.TargetId.GetLengthInBits(ctx)

	// Simple field (error)
	lengthInBits += m.Error.GetLengthInBits(ctx)

	// Simple field (message)
	lengthInBits += m.Message.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_TransactionErrorType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func TransactionErrorTypeParse(ctx context.Context, theBytes []byte, identifier string) (TransactionErrorType, error) {
	return TransactionErrorTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func TransactionErrorTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (TransactionErrorType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("TransactionErrorType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for TransactionErrorType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (targetId)
	if pullErr := readBuffer.PullContext("targetId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for targetId")
	}
_targetId, _targetIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _targetIdErr != nil {
		return nil, errors.Wrap(_targetIdErr, "Error parsing 'targetId' field of TransactionErrorType")
	}
	targetId := _targetId.(NodeId)
	if closeErr := readBuffer.CloseContext("targetId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for targetId")
	}

	// Simple Field (error)
	if pullErr := readBuffer.PullContext("error"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for error")
	}
_error, _errorErr := StatusCodeParseWithBuffer(ctx, readBuffer)
	if _errorErr != nil {
		return nil, errors.Wrap(_errorErr, "Error parsing 'error' field of TransactionErrorType")
	}
	error := _error.(StatusCode)
	if closeErr := readBuffer.CloseContext("error"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for error")
	}

	// Simple Field (message)
	if pullErr := readBuffer.PullContext("message"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for message")
	}
_message, _messageErr := LocalizedTextParseWithBuffer(ctx, readBuffer)
	if _messageErr != nil {
		return nil, errors.Wrap(_messageErr, "Error parsing 'message' field of TransactionErrorType")
	}
	message := _message.(LocalizedText)
	if closeErr := readBuffer.CloseContext("message"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for message")
	}

	if closeErr := readBuffer.CloseContext("TransactionErrorType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for TransactionErrorType")
	}

	// Create a partially initialized instance
	_child := &_TransactionErrorType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{
		},
		TargetId: targetId,
		Error: error,
		Message: message,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_TransactionErrorType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_TransactionErrorType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("TransactionErrorType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for TransactionErrorType")
		}

	// Simple Field (targetId)
	if pushErr := writeBuffer.PushContext("targetId"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for targetId")
	}
	_targetIdErr := writeBuffer.WriteSerializable(ctx, m.GetTargetId())
	if popErr := writeBuffer.PopContext("targetId"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for targetId")
	}
	if _targetIdErr != nil {
		return errors.Wrap(_targetIdErr, "Error serializing 'targetId' field")
	}

	// Simple Field (error)
	if pushErr := writeBuffer.PushContext("error"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for error")
	}
	_errorErr := writeBuffer.WriteSerializable(ctx, m.GetError())
	if popErr := writeBuffer.PopContext("error"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for error")
	}
	if _errorErr != nil {
		return errors.Wrap(_errorErr, "Error serializing 'error' field")
	}

	// Simple Field (message)
	if pushErr := writeBuffer.PushContext("message"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for message")
	}
	_messageErr := writeBuffer.WriteSerializable(ctx, m.GetMessage())
	if popErr := writeBuffer.PopContext("message"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for message")
	}
	if _messageErr != nil {
		return errors.Wrap(_messageErr, "Error serializing 'message' field")
	}

		if popErr := writeBuffer.PopContext("TransactionErrorType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for TransactionErrorType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_TransactionErrorType) isTransactionErrorType() bool {
	return true
}

func (m *_TransactionErrorType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



