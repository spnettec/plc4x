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

// SessionSecurityDiagnosticsDataType is the corresponding interface of SessionSecurityDiagnosticsDataType
type SessionSecurityDiagnosticsDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetSessionId returns SessionId (property field)
	GetSessionId() NodeId
	// GetClientUserIdOfSession returns ClientUserIdOfSession (property field)
	GetClientUserIdOfSession() PascalString
	// GetNoOfClientUserIdHistory returns NoOfClientUserIdHistory (property field)
	GetNoOfClientUserIdHistory() int32
	// GetClientUserIdHistory returns ClientUserIdHistory (property field)
	GetClientUserIdHistory() []PascalString
	// GetAuthenticationMechanism returns AuthenticationMechanism (property field)
	GetAuthenticationMechanism() PascalString
	// GetEncoding returns Encoding (property field)
	GetEncoding() PascalString
	// GetTransportProtocol returns TransportProtocol (property field)
	GetTransportProtocol() PascalString
	// GetSecurityMode returns SecurityMode (property field)
	GetSecurityMode() MessageSecurityMode
	// GetSecurityPolicyUri returns SecurityPolicyUri (property field)
	GetSecurityPolicyUri() PascalString
	// GetClientCertificate returns ClientCertificate (property field)
	GetClientCertificate() PascalByteString
}

// SessionSecurityDiagnosticsDataTypeExactly can be used when we want exactly this type and not a type which fulfills SessionSecurityDiagnosticsDataType.
// This is useful for switch cases.
type SessionSecurityDiagnosticsDataTypeExactly interface {
	SessionSecurityDiagnosticsDataType
	isSessionSecurityDiagnosticsDataType() bool
}

// _SessionSecurityDiagnosticsDataType is the data-structure of this message
type _SessionSecurityDiagnosticsDataType struct {
	*_ExtensionObjectDefinition
	SessionId               NodeId
	ClientUserIdOfSession   PascalString
	NoOfClientUserIdHistory int32
	ClientUserIdHistory     []PascalString
	AuthenticationMechanism PascalString
	Encoding                PascalString
	TransportProtocol       PascalString
	SecurityMode            MessageSecurityMode
	SecurityPolicyUri       PascalString
	ClientCertificate       PascalByteString
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SessionSecurityDiagnosticsDataType) GetIdentifier() string {
	return "870"
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SessionSecurityDiagnosticsDataType) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_SessionSecurityDiagnosticsDataType) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SessionSecurityDiagnosticsDataType) GetSessionId() NodeId {
	return m.SessionId
}

func (m *_SessionSecurityDiagnosticsDataType) GetClientUserIdOfSession() PascalString {
	return m.ClientUserIdOfSession
}

func (m *_SessionSecurityDiagnosticsDataType) GetNoOfClientUserIdHistory() int32 {
	return m.NoOfClientUserIdHistory
}

func (m *_SessionSecurityDiagnosticsDataType) GetClientUserIdHistory() []PascalString {
	return m.ClientUserIdHistory
}

func (m *_SessionSecurityDiagnosticsDataType) GetAuthenticationMechanism() PascalString {
	return m.AuthenticationMechanism
}

func (m *_SessionSecurityDiagnosticsDataType) GetEncoding() PascalString {
	return m.Encoding
}

func (m *_SessionSecurityDiagnosticsDataType) GetTransportProtocol() PascalString {
	return m.TransportProtocol
}

func (m *_SessionSecurityDiagnosticsDataType) GetSecurityMode() MessageSecurityMode {
	return m.SecurityMode
}

func (m *_SessionSecurityDiagnosticsDataType) GetSecurityPolicyUri() PascalString {
	return m.SecurityPolicyUri
}

func (m *_SessionSecurityDiagnosticsDataType) GetClientCertificate() PascalByteString {
	return m.ClientCertificate
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSessionSecurityDiagnosticsDataType factory function for _SessionSecurityDiagnosticsDataType
func NewSessionSecurityDiagnosticsDataType(sessionId NodeId, clientUserIdOfSession PascalString, noOfClientUserIdHistory int32, clientUserIdHistory []PascalString, authenticationMechanism PascalString, encoding PascalString, transportProtocol PascalString, securityMode MessageSecurityMode, securityPolicyUri PascalString, clientCertificate PascalByteString) *_SessionSecurityDiagnosticsDataType {
	_result := &_SessionSecurityDiagnosticsDataType{
		SessionId:                  sessionId,
		ClientUserIdOfSession:      clientUserIdOfSession,
		NoOfClientUserIdHistory:    noOfClientUserIdHistory,
		ClientUserIdHistory:        clientUserIdHistory,
		AuthenticationMechanism:    authenticationMechanism,
		Encoding:                   encoding,
		TransportProtocol:          transportProtocol,
		SecurityMode:               securityMode,
		SecurityPolicyUri:          securityPolicyUri,
		ClientCertificate:          clientCertificate,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSessionSecurityDiagnosticsDataType(structType any) SessionSecurityDiagnosticsDataType {
	if casted, ok := structType.(SessionSecurityDiagnosticsDataType); ok {
		return casted
	}
	if casted, ok := structType.(*SessionSecurityDiagnosticsDataType); ok {
		return *casted
	}
	return nil
}

func (m *_SessionSecurityDiagnosticsDataType) GetTypeName() string {
	return "SessionSecurityDiagnosticsDataType"
}

func (m *_SessionSecurityDiagnosticsDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (sessionId)
	lengthInBits += m.SessionId.GetLengthInBits(ctx)

	// Simple field (clientUserIdOfSession)
	lengthInBits += m.ClientUserIdOfSession.GetLengthInBits(ctx)

	// Simple field (noOfClientUserIdHistory)
	lengthInBits += 32

	// Array field
	if len(m.ClientUserIdHistory) > 0 {
		for _curItem, element := range m.ClientUserIdHistory {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.ClientUserIdHistory), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (authenticationMechanism)
	lengthInBits += m.AuthenticationMechanism.GetLengthInBits(ctx)

	// Simple field (encoding)
	lengthInBits += m.Encoding.GetLengthInBits(ctx)

	// Simple field (transportProtocol)
	lengthInBits += m.TransportProtocol.GetLengthInBits(ctx)

	// Simple field (securityMode)
	lengthInBits += 32

	// Simple field (securityPolicyUri)
	lengthInBits += m.SecurityPolicyUri.GetLengthInBits(ctx)

	// Simple field (clientCertificate)
	lengthInBits += m.ClientCertificate.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SessionSecurityDiagnosticsDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SessionSecurityDiagnosticsDataTypeParse(ctx context.Context, theBytes []byte, identifier string) (SessionSecurityDiagnosticsDataType, error) {
	return SessionSecurityDiagnosticsDataTypeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func SessionSecurityDiagnosticsDataTypeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (SessionSecurityDiagnosticsDataType, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SessionSecurityDiagnosticsDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SessionSecurityDiagnosticsDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (sessionId)
	if pullErr := readBuffer.PullContext("sessionId"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for sessionId")
	}
	_sessionId, _sessionIdErr := NodeIdParseWithBuffer(ctx, readBuffer)
	if _sessionIdErr != nil {
		return nil, errors.Wrap(_sessionIdErr, "Error parsing 'sessionId' field of SessionSecurityDiagnosticsDataType")
	}
	sessionId := _sessionId.(NodeId)
	if closeErr := readBuffer.CloseContext("sessionId"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for sessionId")
	}

	// Simple Field (clientUserIdOfSession)
	if pullErr := readBuffer.PullContext("clientUserIdOfSession"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for clientUserIdOfSession")
	}
	_clientUserIdOfSession, _clientUserIdOfSessionErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _clientUserIdOfSessionErr != nil {
		return nil, errors.Wrap(_clientUserIdOfSessionErr, "Error parsing 'clientUserIdOfSession' field of SessionSecurityDiagnosticsDataType")
	}
	clientUserIdOfSession := _clientUserIdOfSession.(PascalString)
	if closeErr := readBuffer.CloseContext("clientUserIdOfSession"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for clientUserIdOfSession")
	}

	// Simple Field (noOfClientUserIdHistory)
	_noOfClientUserIdHistory, _noOfClientUserIdHistoryErr := readBuffer.ReadInt32("noOfClientUserIdHistory", 32)
	if _noOfClientUserIdHistoryErr != nil {
		return nil, errors.Wrap(_noOfClientUserIdHistoryErr, "Error parsing 'noOfClientUserIdHistory' field of SessionSecurityDiagnosticsDataType")
	}
	noOfClientUserIdHistory := _noOfClientUserIdHistory

	// Array field (clientUserIdHistory)
	if pullErr := readBuffer.PullContext("clientUserIdHistory", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for clientUserIdHistory")
	}
	// Count array
	clientUserIdHistory := make([]PascalString, utils.Max(noOfClientUserIdHistory, 0))
	// This happens when the size is set conditional to 0
	if len(clientUserIdHistory) == 0 {
		clientUserIdHistory = nil
	}
	{
		_numItems := uint16(utils.Max(noOfClientUserIdHistory, 0))
		for _curItem := uint16(0); _curItem < _numItems; _curItem++ {
			arrayCtx := utils.CreateArrayContext(ctx, int(_numItems), int(_curItem))
			_ = arrayCtx
			_ = _curItem
			_item, _err := PascalStringParseWithBuffer(arrayCtx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'clientUserIdHistory' field of SessionSecurityDiagnosticsDataType")
			}
			clientUserIdHistory[_curItem] = _item.(PascalString)
		}
	}
	if closeErr := readBuffer.CloseContext("clientUserIdHistory", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for clientUserIdHistory")
	}

	// Simple Field (authenticationMechanism)
	if pullErr := readBuffer.PullContext("authenticationMechanism"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for authenticationMechanism")
	}
	_authenticationMechanism, _authenticationMechanismErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _authenticationMechanismErr != nil {
		return nil, errors.Wrap(_authenticationMechanismErr, "Error parsing 'authenticationMechanism' field of SessionSecurityDiagnosticsDataType")
	}
	authenticationMechanism := _authenticationMechanism.(PascalString)
	if closeErr := readBuffer.CloseContext("authenticationMechanism"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for authenticationMechanism")
	}

	// Simple Field (encoding)
	if pullErr := readBuffer.PullContext("encoding"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for encoding")
	}
	_encoding, _encodingErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _encodingErr != nil {
		return nil, errors.Wrap(_encodingErr, "Error parsing 'encoding' field of SessionSecurityDiagnosticsDataType")
	}
	encoding := _encoding.(PascalString)
	if closeErr := readBuffer.CloseContext("encoding"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for encoding")
	}

	// Simple Field (transportProtocol)
	if pullErr := readBuffer.PullContext("transportProtocol"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for transportProtocol")
	}
	_transportProtocol, _transportProtocolErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _transportProtocolErr != nil {
		return nil, errors.Wrap(_transportProtocolErr, "Error parsing 'transportProtocol' field of SessionSecurityDiagnosticsDataType")
	}
	transportProtocol := _transportProtocol.(PascalString)
	if closeErr := readBuffer.CloseContext("transportProtocol"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for transportProtocol")
	}

	// Simple Field (securityMode)
	if pullErr := readBuffer.PullContext("securityMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for securityMode")
	}
	_securityMode, _securityModeErr := MessageSecurityModeParseWithBuffer(ctx, readBuffer)
	if _securityModeErr != nil {
		return nil, errors.Wrap(_securityModeErr, "Error parsing 'securityMode' field of SessionSecurityDiagnosticsDataType")
	}
	securityMode := _securityMode
	if closeErr := readBuffer.CloseContext("securityMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for securityMode")
	}

	// Simple Field (securityPolicyUri)
	if pullErr := readBuffer.PullContext("securityPolicyUri"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for securityPolicyUri")
	}
	_securityPolicyUri, _securityPolicyUriErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _securityPolicyUriErr != nil {
		return nil, errors.Wrap(_securityPolicyUriErr, "Error parsing 'securityPolicyUri' field of SessionSecurityDiagnosticsDataType")
	}
	securityPolicyUri := _securityPolicyUri.(PascalString)
	if closeErr := readBuffer.CloseContext("securityPolicyUri"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for securityPolicyUri")
	}

	// Simple Field (clientCertificate)
	if pullErr := readBuffer.PullContext("clientCertificate"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for clientCertificate")
	}
	_clientCertificate, _clientCertificateErr := PascalByteStringParseWithBuffer(ctx, readBuffer)
	if _clientCertificateErr != nil {
		return nil, errors.Wrap(_clientCertificateErr, "Error parsing 'clientCertificate' field of SessionSecurityDiagnosticsDataType")
	}
	clientCertificate := _clientCertificate.(PascalByteString)
	if closeErr := readBuffer.CloseContext("clientCertificate"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for clientCertificate")
	}

	if closeErr := readBuffer.CloseContext("SessionSecurityDiagnosticsDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SessionSecurityDiagnosticsDataType")
	}

	// Create a partially initialized instance
	_child := &_SessionSecurityDiagnosticsDataType{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		SessionId:                  sessionId,
		ClientUserIdOfSession:      clientUserIdOfSession,
		NoOfClientUserIdHistory:    noOfClientUserIdHistory,
		ClientUserIdHistory:        clientUserIdHistory,
		AuthenticationMechanism:    authenticationMechanism,
		Encoding:                   encoding,
		TransportProtocol:          transportProtocol,
		SecurityMode:               securityMode,
		SecurityPolicyUri:          securityPolicyUri,
		ClientCertificate:          clientCertificate,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_SessionSecurityDiagnosticsDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SessionSecurityDiagnosticsDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SessionSecurityDiagnosticsDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SessionSecurityDiagnosticsDataType")
		}

		// Simple Field (sessionId)
		if pushErr := writeBuffer.PushContext("sessionId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for sessionId")
		}
		_sessionIdErr := writeBuffer.WriteSerializable(ctx, m.GetSessionId())
		if popErr := writeBuffer.PopContext("sessionId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for sessionId")
		}
		if _sessionIdErr != nil {
			return errors.Wrap(_sessionIdErr, "Error serializing 'sessionId' field")
		}

		// Simple Field (clientUserIdOfSession)
		if pushErr := writeBuffer.PushContext("clientUserIdOfSession"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for clientUserIdOfSession")
		}
		_clientUserIdOfSessionErr := writeBuffer.WriteSerializable(ctx, m.GetClientUserIdOfSession())
		if popErr := writeBuffer.PopContext("clientUserIdOfSession"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for clientUserIdOfSession")
		}
		if _clientUserIdOfSessionErr != nil {
			return errors.Wrap(_clientUserIdOfSessionErr, "Error serializing 'clientUserIdOfSession' field")
		}

		// Simple Field (noOfClientUserIdHistory)
		noOfClientUserIdHistory := int32(m.GetNoOfClientUserIdHistory())
		_noOfClientUserIdHistoryErr := writeBuffer.WriteInt32("noOfClientUserIdHistory", 32, (noOfClientUserIdHistory))
		if _noOfClientUserIdHistoryErr != nil {
			return errors.Wrap(_noOfClientUserIdHistoryErr, "Error serializing 'noOfClientUserIdHistory' field")
		}

		// Array Field (clientUserIdHistory)
		if pushErr := writeBuffer.PushContext("clientUserIdHistory", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for clientUserIdHistory")
		}
		for _curItem, _element := range m.GetClientUserIdHistory() {
			_ = _curItem
			arrayCtx := utils.CreateArrayContext(ctx, len(m.GetClientUserIdHistory()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'clientUserIdHistory' field")
			}
		}
		if popErr := writeBuffer.PopContext("clientUserIdHistory", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for clientUserIdHistory")
		}

		// Simple Field (authenticationMechanism)
		if pushErr := writeBuffer.PushContext("authenticationMechanism"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for authenticationMechanism")
		}
		_authenticationMechanismErr := writeBuffer.WriteSerializable(ctx, m.GetAuthenticationMechanism())
		if popErr := writeBuffer.PopContext("authenticationMechanism"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for authenticationMechanism")
		}
		if _authenticationMechanismErr != nil {
			return errors.Wrap(_authenticationMechanismErr, "Error serializing 'authenticationMechanism' field")
		}

		// Simple Field (encoding)
		if pushErr := writeBuffer.PushContext("encoding"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for encoding")
		}
		_encodingErr := writeBuffer.WriteSerializable(ctx, m.GetEncoding())
		if popErr := writeBuffer.PopContext("encoding"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for encoding")
		}
		if _encodingErr != nil {
			return errors.Wrap(_encodingErr, "Error serializing 'encoding' field")
		}

		// Simple Field (transportProtocol)
		if pushErr := writeBuffer.PushContext("transportProtocol"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for transportProtocol")
		}
		_transportProtocolErr := writeBuffer.WriteSerializable(ctx, m.GetTransportProtocol())
		if popErr := writeBuffer.PopContext("transportProtocol"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for transportProtocol")
		}
		if _transportProtocolErr != nil {
			return errors.Wrap(_transportProtocolErr, "Error serializing 'transportProtocol' field")
		}

		// Simple Field (securityMode)
		if pushErr := writeBuffer.PushContext("securityMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securityMode")
		}
		_securityModeErr := writeBuffer.WriteSerializable(ctx, m.GetSecurityMode())
		if popErr := writeBuffer.PopContext("securityMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securityMode")
		}
		if _securityModeErr != nil {
			return errors.Wrap(_securityModeErr, "Error serializing 'securityMode' field")
		}

		// Simple Field (securityPolicyUri)
		if pushErr := writeBuffer.PushContext("securityPolicyUri"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for securityPolicyUri")
		}
		_securityPolicyUriErr := writeBuffer.WriteSerializable(ctx, m.GetSecurityPolicyUri())
		if popErr := writeBuffer.PopContext("securityPolicyUri"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for securityPolicyUri")
		}
		if _securityPolicyUriErr != nil {
			return errors.Wrap(_securityPolicyUriErr, "Error serializing 'securityPolicyUri' field")
		}

		// Simple Field (clientCertificate)
		if pushErr := writeBuffer.PushContext("clientCertificate"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for clientCertificate")
		}
		_clientCertificateErr := writeBuffer.WriteSerializable(ctx, m.GetClientCertificate())
		if popErr := writeBuffer.PopContext("clientCertificate"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for clientCertificate")
		}
		if _clientCertificateErr != nil {
			return errors.Wrap(_clientCertificateErr, "Error serializing 'clientCertificate' field")
		}

		if popErr := writeBuffer.PopContext("SessionSecurityDiagnosticsDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SessionSecurityDiagnosticsDataType")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SessionSecurityDiagnosticsDataType) isSessionSecurityDiagnosticsDataType() bool {
	return true
}

func (m *_SessionSecurityDiagnosticsDataType) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
