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


// MediaTransportControlDataTotalTracks is the corresponding interface of MediaTransportControlDataTotalTracks
type MediaTransportControlDataTotalTracks interface {
	utils.LengthAware
	utils.Serializable
	MediaTransportControlData
	// GetTotalTracksMSB returns TotalTracksMSB (property field)
	GetTotalTracksMSB() byte
	// GetTotalTracksMMSB returns TotalTracksMMSB (property field)
	GetTotalTracksMMSB() byte
	// GetTotalTracksMLSB returns TotalTracksMLSB (property field)
	GetTotalTracksMLSB() byte
	// GetTotalTracksLSB returns TotalTracksLSB (property field)
	GetTotalTracksLSB() byte
}

// MediaTransportControlDataTotalTracksExactly can be used when we want exactly this type and not a type which fulfills MediaTransportControlDataTotalTracks.
// This is useful for switch cases.
type MediaTransportControlDataTotalTracksExactly interface {
	MediaTransportControlDataTotalTracks
	isMediaTransportControlDataTotalTracks() bool
}

// _MediaTransportControlDataTotalTracks is the data-structure of this message
type _MediaTransportControlDataTotalTracks struct {
	*_MediaTransportControlData
        TotalTracksMSB byte
        TotalTracksMMSB byte
        TotalTracksMLSB byte
        TotalTracksLSB byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MediaTransportControlDataTotalTracks) InitializeParent(parent MediaTransportControlData , commandTypeContainer MediaTransportControlCommandTypeContainer , mediaLinkGroup byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.MediaLinkGroup = mediaLinkGroup
}

func (m *_MediaTransportControlDataTotalTracks)  GetParent() MediaTransportControlData {
	return m._MediaTransportControlData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MediaTransportControlDataTotalTracks) GetTotalTracksMSB() byte {
	return m.TotalTracksMSB
}

func (m *_MediaTransportControlDataTotalTracks) GetTotalTracksMMSB() byte {
	return m.TotalTracksMMSB
}

func (m *_MediaTransportControlDataTotalTracks) GetTotalTracksMLSB() byte {
	return m.TotalTracksMLSB
}

func (m *_MediaTransportControlDataTotalTracks) GetTotalTracksLSB() byte {
	return m.TotalTracksLSB
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewMediaTransportControlDataTotalTracks factory function for _MediaTransportControlDataTotalTracks
func NewMediaTransportControlDataTotalTracks( totalTracksMSB byte , totalTracksMMSB byte , totalTracksMLSB byte , totalTracksLSB byte , commandTypeContainer MediaTransportControlCommandTypeContainer , mediaLinkGroup byte ) *_MediaTransportControlDataTotalTracks {
	_result := &_MediaTransportControlDataTotalTracks{
		TotalTracksMSB: totalTracksMSB,
		TotalTracksMMSB: totalTracksMMSB,
		TotalTracksMLSB: totalTracksMLSB,
		TotalTracksLSB: totalTracksLSB,
    	_MediaTransportControlData: NewMediaTransportControlData(commandTypeContainer, mediaLinkGroup),
	}
	_result._MediaTransportControlData._MediaTransportControlDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMediaTransportControlDataTotalTracks(structType interface{}) MediaTransportControlDataTotalTracks {
    if casted, ok := structType.(MediaTransportControlDataTotalTracks); ok {
		return casted
	}
	if casted, ok := structType.(*MediaTransportControlDataTotalTracks); ok {
		return *casted
	}
	return nil
}

func (m *_MediaTransportControlDataTotalTracks) GetTypeName() string {
	return "MediaTransportControlDataTotalTracks"
}

func (m *_MediaTransportControlDataTotalTracks) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_MediaTransportControlDataTotalTracks) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (totalTracksMSB)
	lengthInBits += 8;

	// Simple field (totalTracksMMSB)
	lengthInBits += 8;

	// Simple field (totalTracksMLSB)
	lengthInBits += 8;

	// Simple field (totalTracksLSB)
	lengthInBits += 8;

	return lengthInBits
}


func (m *_MediaTransportControlDataTotalTracks) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func MediaTransportControlDataTotalTracksParse(readBuffer utils.ReadBuffer) (MediaTransportControlDataTotalTracks, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MediaTransportControlDataTotalTracks"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MediaTransportControlDataTotalTracks")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (totalTracksMSB)
_totalTracksMSB, _totalTracksMSBErr := readBuffer.ReadByte("totalTracksMSB")
	if _totalTracksMSBErr != nil {
		return nil, errors.Wrap(_totalTracksMSBErr, "Error parsing 'totalTracksMSB' field of MediaTransportControlDataTotalTracks")
	}
	totalTracksMSB := _totalTracksMSB

	// Simple Field (totalTracksMMSB)
_totalTracksMMSB, _totalTracksMMSBErr := readBuffer.ReadByte("totalTracksMMSB")
	if _totalTracksMMSBErr != nil {
		return nil, errors.Wrap(_totalTracksMMSBErr, "Error parsing 'totalTracksMMSB' field of MediaTransportControlDataTotalTracks")
	}
	totalTracksMMSB := _totalTracksMMSB

	// Simple Field (totalTracksMLSB)
_totalTracksMLSB, _totalTracksMLSBErr := readBuffer.ReadByte("totalTracksMLSB")
	if _totalTracksMLSBErr != nil {
		return nil, errors.Wrap(_totalTracksMLSBErr, "Error parsing 'totalTracksMLSB' field of MediaTransportControlDataTotalTracks")
	}
	totalTracksMLSB := _totalTracksMLSB

	// Simple Field (totalTracksLSB)
_totalTracksLSB, _totalTracksLSBErr := readBuffer.ReadByte("totalTracksLSB")
	if _totalTracksLSBErr != nil {
		return nil, errors.Wrap(_totalTracksLSBErr, "Error parsing 'totalTracksLSB' field of MediaTransportControlDataTotalTracks")
	}
	totalTracksLSB := _totalTracksLSB

	if closeErr := readBuffer.CloseContext("MediaTransportControlDataTotalTracks"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MediaTransportControlDataTotalTracks")
	}

	// Create a partially initialized instance
	_child := &_MediaTransportControlDataTotalTracks{
		_MediaTransportControlData: &_MediaTransportControlData{
		},
		TotalTracksMSB: totalTracksMSB,
		TotalTracksMMSB: totalTracksMMSB,
		TotalTracksMLSB: totalTracksMLSB,
		TotalTracksLSB: totalTracksLSB,
	}
	_child._MediaTransportControlData._MediaTransportControlDataChildRequirements = _child
	return _child, nil
}

func (m *_MediaTransportControlDataTotalTracks) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MediaTransportControlDataTotalTracks"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MediaTransportControlDataTotalTracks")
		}

	// Simple Field (totalTracksMSB)
	totalTracksMSB := byte(m.GetTotalTracksMSB())
	_totalTracksMSBErr := writeBuffer.WriteByte("totalTracksMSB", (totalTracksMSB))
	if _totalTracksMSBErr != nil {
		return errors.Wrap(_totalTracksMSBErr, "Error serializing 'totalTracksMSB' field")
	}

	// Simple Field (totalTracksMMSB)
	totalTracksMMSB := byte(m.GetTotalTracksMMSB())
	_totalTracksMMSBErr := writeBuffer.WriteByte("totalTracksMMSB", (totalTracksMMSB))
	if _totalTracksMMSBErr != nil {
		return errors.Wrap(_totalTracksMMSBErr, "Error serializing 'totalTracksMMSB' field")
	}

	// Simple Field (totalTracksMLSB)
	totalTracksMLSB := byte(m.GetTotalTracksMLSB())
	_totalTracksMLSBErr := writeBuffer.WriteByte("totalTracksMLSB", (totalTracksMLSB))
	if _totalTracksMLSBErr != nil {
		return errors.Wrap(_totalTracksMLSBErr, "Error serializing 'totalTracksMLSB' field")
	}

	// Simple Field (totalTracksLSB)
	totalTracksLSB := byte(m.GetTotalTracksLSB())
	_totalTracksLSBErr := writeBuffer.WriteByte("totalTracksLSB", (totalTracksLSB))
	if _totalTracksLSBErr != nil {
		return errors.Wrap(_totalTracksLSBErr, "Error serializing 'totalTracksLSB' field")
	}

		if popErr := writeBuffer.PopContext("MediaTransportControlDataTotalTracks"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MediaTransportControlDataTotalTracks")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_MediaTransportControlDataTotalTracks) isMediaTransportControlDataTotalTracks() bool {
	return true
}

func (m *_MediaTransportControlDataTotalTracks) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



