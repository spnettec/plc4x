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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// RequestContext is the corresponding interface of RequestContext
type RequestContext interface {
	utils.LengthAware
	utils.Serializable
	// GetSendIdentifyRequestBefore returns SendIdentifyRequestBefore (property field)
	GetSendIdentifyRequestBefore() bool
}

// RequestContextExactly can be used when we want exactly this type and not a type which fulfills RequestContext.
// This is useful for switch cases.
type RequestContextExactly interface {
	RequestContext
	isRequestContext() bool
}

// _RequestContext is the data-structure of this message
type _RequestContext struct {
        SendIdentifyRequestBefore bool
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RequestContext) GetSendIdentifyRequestBefore() bool {
	return m.SendIdentifyRequestBefore
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewRequestContext factory function for _RequestContext
func NewRequestContext( sendIdentifyRequestBefore bool ) *_RequestContext {
return &_RequestContext{ SendIdentifyRequestBefore: sendIdentifyRequestBefore }
}

// Deprecated: use the interface for direct cast
func CastRequestContext(structType interface{}) RequestContext {
    if casted, ok := structType.(RequestContext); ok {
		return casted
	}
	if casted, ok := structType.(*RequestContext); ok {
		return *casted
	}
	return nil
}

func (m *_RequestContext) GetTypeName() string {
	return "RequestContext"
}

func (m *_RequestContext) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// Simple field (sendIdentifyRequestBefore)
	lengthInBits += 1;

	return lengthInBits
}


func (m *_RequestContext) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func RequestContextParse(theBytes []byte) (RequestContext, error) {
	return RequestContextParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func RequestContextParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (RequestContext, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestContext"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestContext")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (sendIdentifyRequestBefore)
_sendIdentifyRequestBefore, _sendIdentifyRequestBeforeErr := readBuffer.ReadBit("sendIdentifyRequestBefore")
	if _sendIdentifyRequestBeforeErr != nil {
		return nil, errors.Wrap(_sendIdentifyRequestBeforeErr, "Error parsing 'sendIdentifyRequestBefore' field of RequestContext")
	}
	sendIdentifyRequestBefore := _sendIdentifyRequestBefore

	if closeErr := readBuffer.CloseContext("RequestContext"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestContext")
	}

	// Create the instance
	return &_RequestContext{
			SendIdentifyRequestBefore: sendIdentifyRequestBefore,
		}, nil
}

func (m *_RequestContext) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RequestContext) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("RequestContext"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for RequestContext")
	}

	// Simple Field (sendIdentifyRequestBefore)
	sendIdentifyRequestBefore := bool(m.GetSendIdentifyRequestBefore())
	_sendIdentifyRequestBeforeErr := writeBuffer.WriteBit("sendIdentifyRequestBefore", (sendIdentifyRequestBefore))
	if _sendIdentifyRequestBeforeErr != nil {
		return errors.Wrap(_sendIdentifyRequestBeforeErr, "Error serializing 'sendIdentifyRequestBefore' field")
	}

	if popErr := writeBuffer.PopContext("RequestContext"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for RequestContext")
	}
	return nil
}


func (m *_RequestContext) isRequestContext() bool {
	return true
}

func (m *_RequestContext) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



