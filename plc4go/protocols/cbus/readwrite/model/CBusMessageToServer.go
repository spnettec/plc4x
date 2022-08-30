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


// CBusMessageToServer is the corresponding interface of CBusMessageToServer
type CBusMessageToServer interface {
	utils.LengthAware
	utils.Serializable
	CBusMessage
	// GetRequest returns Request (property field)
	GetRequest() Request
}

// CBusMessageToServerExactly can be used when we want exactly this type and not a type which fulfills CBusMessageToServer.
// This is useful for switch cases.
type CBusMessageToServerExactly interface {
	CBusMessageToServer
	isCBusMessageToServer() bool
}

// _CBusMessageToServer is the data-structure of this message
type _CBusMessageToServer struct {
	*_CBusMessage
        Request Request
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_CBusMessageToServer)  GetIsResponse() bool {
return bool(false)}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CBusMessageToServer) InitializeParent(parent CBusMessage ) {}

func (m *_CBusMessageToServer)  GetParent() CBusMessage {
	return m._CBusMessage
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusMessageToServer) GetRequest() Request {
	return m.Request
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewCBusMessageToServer factory function for _CBusMessageToServer
func NewCBusMessageToServer( request Request , requestContext RequestContext , cBusOptions CBusOptions ) *_CBusMessageToServer {
	_result := &_CBusMessageToServer{
		Request: request,
    	_CBusMessage: NewCBusMessage(requestContext, cBusOptions),
	}
	_result._CBusMessage._CBusMessageChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCBusMessageToServer(structType interface{}) CBusMessageToServer {
    if casted, ok := structType.(CBusMessageToServer); ok {
		return casted
	}
	if casted, ok := structType.(*CBusMessageToServer); ok {
		return *casted
	}
	return nil
}

func (m *_CBusMessageToServer) GetTypeName() string {
	return "CBusMessageToServer"
}

func (m *_CBusMessageToServer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CBusMessageToServer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (request)
	lengthInBits += m.Request.GetLengthInBits()

	return lengthInBits
}


func (m *_CBusMessageToServer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CBusMessageToServerParse(readBuffer utils.ReadBuffer, isResponse bool, requestContext RequestContext, cBusOptions CBusOptions) (CBusMessageToServer, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CBusMessageToServer"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusMessageToServer")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (request)
	if pullErr := readBuffer.PullContext("request"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for request")
	}
_request, _requestErr := RequestParse(readBuffer , cBusOptions )
	if _requestErr != nil {
		return nil, errors.Wrap(_requestErr, "Error parsing 'request' field of CBusMessageToServer")
	}
	request := _request.(Request)
	if closeErr := readBuffer.CloseContext("request"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for request")
	}

	if closeErr := readBuffer.CloseContext("CBusMessageToServer"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusMessageToServer")
	}

	// Create a partially initialized instance
	_child := &_CBusMessageToServer{
		_CBusMessage: &_CBusMessage{
			RequestContext: requestContext,
			CBusOptions: cBusOptions,
		},
		Request: request,
	}
	_child._CBusMessage._CBusMessageChildRequirements = _child
	return _child, nil
}

func (m *_CBusMessageToServer) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusMessageToServer"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusMessageToServer")
		}

	// Simple Field (request)
	if pushErr := writeBuffer.PushContext("request"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for request")
	}
	_requestErr := writeBuffer.WriteSerializable(m.GetRequest())
	if popErr := writeBuffer.PopContext("request"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for request")
	}
	if _requestErr != nil {
		return errors.Wrap(_requestErr, "Error serializing 'request' field")
	}

		if popErr := writeBuffer.PopContext("CBusMessageToServer"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusMessageToServer")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_CBusMessageToServer) isCBusMessageToServer() bool {
	return true
}

func (m *_CBusMessageToServer) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



