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

// SALDataMediaTransport is the corresponding interface of SALDataMediaTransport
type SALDataMediaTransport interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	SALData
	// GetMediaTransportControlData returns MediaTransportControlData (property field)
	GetMediaTransportControlData() MediaTransportControlData
}

// SALDataMediaTransportExactly can be used when we want exactly this type and not a type which fulfills SALDataMediaTransport.
// This is useful for switch cases.
type SALDataMediaTransportExactly interface {
	SALDataMediaTransport
	isSALDataMediaTransport() bool
}

// _SALDataMediaTransport is the data-structure of this message
type _SALDataMediaTransport struct {
	*_SALData
	MediaTransportControlData MediaTransportControlData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataMediaTransport) GetApplicationId() ApplicationId {
	return ApplicationId_MEDIA_TRANSPORT_CONTROL
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataMediaTransport) InitializeParent(parent SALData, salData SALData) {
	m.SalData = salData
}

func (m *_SALDataMediaTransport) GetParent() SALData {
	return m._SALData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataMediaTransport) GetMediaTransportControlData() MediaTransportControlData {
	return m.MediaTransportControlData
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewSALDataMediaTransport factory function for _SALDataMediaTransport
func NewSALDataMediaTransport(mediaTransportControlData MediaTransportControlData, salData SALData) *_SALDataMediaTransport {
	_result := &_SALDataMediaTransport{
		MediaTransportControlData: mediaTransportControlData,
		_SALData:                  NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataMediaTransport(structType any) SALDataMediaTransport {
	if casted, ok := structType.(SALDataMediaTransport); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataMediaTransport); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataMediaTransport) GetTypeName() string {
	return "SALDataMediaTransport"
}

func (m *_SALDataMediaTransport) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (mediaTransportControlData)
	lengthInBits += m.MediaTransportControlData.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_SALDataMediaTransport) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func SALDataMediaTransportParse(ctx context.Context, theBytes []byte, applicationId ApplicationId) (SALDataMediaTransport, error) {
	return SALDataMediaTransportParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), applicationId)
}

func SALDataMediaTransportParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataMediaTransport, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("SALDataMediaTransport"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataMediaTransport")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (mediaTransportControlData)
	if pullErr := readBuffer.PullContext("mediaTransportControlData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for mediaTransportControlData")
	}
	_mediaTransportControlData, _mediaTransportControlDataErr := MediaTransportControlDataParseWithBuffer(ctx, readBuffer)
	if _mediaTransportControlDataErr != nil {
		return nil, errors.Wrap(_mediaTransportControlDataErr, "Error parsing 'mediaTransportControlData' field of SALDataMediaTransport")
	}
	mediaTransportControlData := _mediaTransportControlData.(MediaTransportControlData)
	if closeErr := readBuffer.CloseContext("mediaTransportControlData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for mediaTransportControlData")
	}

	if closeErr := readBuffer.CloseContext("SALDataMediaTransport"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataMediaTransport")
	}

	// Create a partially initialized instance
	_child := &_SALDataMediaTransport{
		_SALData:                  &_SALData{},
		MediaTransportControlData: mediaTransportControlData,
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataMediaTransport) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_SALDataMediaTransport) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataMediaTransport"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataMediaTransport")
		}

		// Simple Field (mediaTransportControlData)
		if pushErr := writeBuffer.PushContext("mediaTransportControlData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for mediaTransportControlData")
		}
		_mediaTransportControlDataErr := writeBuffer.WriteSerializable(ctx, m.GetMediaTransportControlData())
		if popErr := writeBuffer.PopContext("mediaTransportControlData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for mediaTransportControlData")
		}
		if _mediaTransportControlDataErr != nil {
			return errors.Wrap(_mediaTransportControlDataErr, "Error serializing 'mediaTransportControlData' field")
		}

		if popErr := writeBuffer.PopContext("SALDataMediaTransport"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataMediaTransport")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_SALDataMediaTransport) isSALDataMediaTransport() bool {
	return true
}

func (m *_SALDataMediaTransport) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
