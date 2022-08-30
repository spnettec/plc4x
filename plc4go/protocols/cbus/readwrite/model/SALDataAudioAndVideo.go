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


// SALDataAudioAndVideo is the corresponding interface of SALDataAudioAndVideo
type SALDataAudioAndVideo interface {
	utils.LengthAware
	utils.Serializable
	SALData
	// GetAudioVideoData returns AudioVideoData (property field)
	GetAudioVideoData() LightingData
}

// SALDataAudioAndVideoExactly can be used when we want exactly this type and not a type which fulfills SALDataAudioAndVideo.
// This is useful for switch cases.
type SALDataAudioAndVideoExactly interface {
	SALDataAudioAndVideo
	isSALDataAudioAndVideo() bool
}

// _SALDataAudioAndVideo is the data-structure of this message
type _SALDataAudioAndVideo struct {
	*_SALData
        AudioVideoData LightingData
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_SALDataAudioAndVideo)  GetApplicationId() ApplicationId {
return ApplicationId_AUDIO_AND_VIDEO}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SALDataAudioAndVideo) InitializeParent(parent SALData , salData SALData ) {	m.SalData = salData
}

func (m *_SALDataAudioAndVideo)  GetParent() SALData {
	return m._SALData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SALDataAudioAndVideo) GetAudioVideoData() LightingData {
	return m.AudioVideoData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewSALDataAudioAndVideo factory function for _SALDataAudioAndVideo
func NewSALDataAudioAndVideo( audioVideoData LightingData , salData SALData ) *_SALDataAudioAndVideo {
	_result := &_SALDataAudioAndVideo{
		AudioVideoData: audioVideoData,
    	_SALData: NewSALData(salData),
	}
	_result._SALData._SALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSALDataAudioAndVideo(structType interface{}) SALDataAudioAndVideo {
    if casted, ok := structType.(SALDataAudioAndVideo); ok {
		return casted
	}
	if casted, ok := structType.(*SALDataAudioAndVideo); ok {
		return *casted
	}
	return nil
}

func (m *_SALDataAudioAndVideo) GetTypeName() string {
	return "SALDataAudioAndVideo"
}

func (m *_SALDataAudioAndVideo) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SALDataAudioAndVideo) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (audioVideoData)
	lengthInBits += m.AudioVideoData.GetLengthInBits()

	return lengthInBits
}


func (m *_SALDataAudioAndVideo) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SALDataAudioAndVideoParse(readBuffer utils.ReadBuffer, applicationId ApplicationId) (SALDataAudioAndVideo, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SALDataAudioAndVideo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SALDataAudioAndVideo")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (audioVideoData)
	if pullErr := readBuffer.PullContext("audioVideoData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for audioVideoData")
	}
_audioVideoData, _audioVideoDataErr := LightingDataParse(readBuffer)
	if _audioVideoDataErr != nil {
		return nil, errors.Wrap(_audioVideoDataErr, "Error parsing 'audioVideoData' field of SALDataAudioAndVideo")
	}
	audioVideoData := _audioVideoData.(LightingData)
	if closeErr := readBuffer.CloseContext("audioVideoData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for audioVideoData")
	}

	if closeErr := readBuffer.CloseContext("SALDataAudioAndVideo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SALDataAudioAndVideo")
	}

	// Create a partially initialized instance
	_child := &_SALDataAudioAndVideo{
		_SALData: &_SALData{
		},
		AudioVideoData: audioVideoData,
	}
	_child._SALData._SALDataChildRequirements = _child
	return _child, nil
}

func (m *_SALDataAudioAndVideo) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SALDataAudioAndVideo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SALDataAudioAndVideo")
		}

	// Simple Field (audioVideoData)
	if pushErr := writeBuffer.PushContext("audioVideoData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for audioVideoData")
	}
	_audioVideoDataErr := writeBuffer.WriteSerializable(m.GetAudioVideoData())
	if popErr := writeBuffer.PopContext("audioVideoData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for audioVideoData")
	}
	if _audioVideoDataErr != nil {
		return errors.Wrap(_audioVideoDataErr, "Error serializing 'audioVideoData' field")
	}

		if popErr := writeBuffer.PopContext("SALDataAudioAndVideo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SALDataAudioAndVideo")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_SALDataAudioAndVideo) isSALDataAudioAndVideo() bool {
	return true
}

func (m *_SALDataAudioAndVideo) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



