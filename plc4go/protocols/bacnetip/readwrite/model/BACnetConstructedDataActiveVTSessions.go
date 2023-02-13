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
	spiContext "github.com/apache/plc4x/plc4go/spi/context"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataActiveVTSessions is the corresponding interface of BACnetConstructedDataActiveVTSessions
type BACnetConstructedDataActiveVTSessions interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetActiveVTSession returns ActiveVTSession (property field)
	GetActiveVTSession() []BACnetVTSession
}

// BACnetConstructedDataActiveVTSessionsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataActiveVTSessions.
// This is useful for switch cases.
type BACnetConstructedDataActiveVTSessionsExactly interface {
	BACnetConstructedDataActiveVTSessions
	isBACnetConstructedDataActiveVTSessions() bool
}

// _BACnetConstructedDataActiveVTSessions is the data-structure of this message
type _BACnetConstructedDataActiveVTSessions struct {
	*_BACnetConstructedData
	ActiveVTSession []BACnetVTSession
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataActiveVTSessions) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataActiveVTSessions) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ACTIVE_VT_SESSIONS
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataActiveVTSessions) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataActiveVTSessions) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataActiveVTSessions) GetActiveVTSession() []BACnetVTSession {
	return m.ActiveVTSession
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataActiveVTSessions factory function for _BACnetConstructedDataActiveVTSessions
func NewBACnetConstructedDataActiveVTSessions(activeVTSession []BACnetVTSession, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataActiveVTSessions {
	_result := &_BACnetConstructedDataActiveVTSessions{
		ActiveVTSession:        activeVTSession,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataActiveVTSessions(structType interface{}) BACnetConstructedDataActiveVTSessions {
	if casted, ok := structType.(BACnetConstructedDataActiveVTSessions); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataActiveVTSessions); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataActiveVTSessions) GetTypeName() string {
	return "BACnetConstructedDataActiveVTSessions"
}

func (m *_BACnetConstructedDataActiveVTSessions) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.ActiveVTSession) > 0 {
		for _, element := range m.ActiveVTSession {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}

func (m *_BACnetConstructedDataActiveVTSessions) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataActiveVTSessionsParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataActiveVTSessions, error) {
	return BACnetConstructedDataActiveVTSessionsParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataActiveVTSessionsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataActiveVTSessions, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataActiveVTSessions"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataActiveVTSessions")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Array field (activeVTSession)
	if pullErr := readBuffer.PullContext("activeVTSession", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for activeVTSession")
	}
	// Terminated array
	var activeVTSession []BACnetVTSession
	{
		for !bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)) {
			_item, _err := BACnetVTSessionParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'activeVTSession' field of BACnetConstructedDataActiveVTSessions")
			}
			activeVTSession = append(activeVTSession, _item.(BACnetVTSession))
		}
	}
	if closeErr := readBuffer.CloseContext("activeVTSession", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for activeVTSession")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataActiveVTSessions"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataActiveVTSessions")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataActiveVTSessions{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		ActiveVTSession: activeVTSession,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataActiveVTSessions) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataActiveVTSessions) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataActiveVTSessions"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataActiveVTSessions")
		}

		// Array Field (activeVTSession)
		if pushErr := writeBuffer.PushContext("activeVTSession", utils.WithRenderAsList(true)); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for activeVTSession")
		}
		for _curItem, _element := range m.GetActiveVTSession() {
			_ = _curItem
			arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetActiveVTSession()), _curItem)
			_ = arrayCtx
			_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
			if _elementErr != nil {
				return errors.Wrap(_elementErr, "Error serializing 'activeVTSession' field")
			}
		}
		if popErr := writeBuffer.PopContext("activeVTSession", utils.WithRenderAsList(true)); popErr != nil {
			return errors.Wrap(popErr, "Error popping for activeVTSession")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataActiveVTSessions"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataActiveVTSessions")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataActiveVTSessions) isBACnetConstructedDataActiveVTSessions() bool {
	return true
}

func (m *_BACnetConstructedDataActiveVTSessions) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
