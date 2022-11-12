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


// CALDataStatus is the corresponding interface of CALDataStatus
type CALDataStatus interface {
	utils.LengthAware
	utils.Serializable
	CALData
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetBlockStart returns BlockStart (property field)
	GetBlockStart() uint8
	// GetStatusBytes returns StatusBytes (property field)
	GetStatusBytes() []StatusByte
}

// CALDataStatusExactly can be used when we want exactly this type and not a type which fulfills CALDataStatus.
// This is useful for switch cases.
type CALDataStatusExactly interface {
	CALDataStatus
	isCALDataStatus() bool
}

// _CALDataStatus is the data-structure of this message
type _CALDataStatus struct {
	*_CALData
        Application ApplicationIdContainer
        BlockStart uint8
        StatusBytes []StatusByte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CALDataStatus) InitializeParent(parent CALData , commandTypeContainer CALCommandTypeContainer , additionalData CALData ) {	m.CommandTypeContainer = commandTypeContainer
	m.AdditionalData = additionalData
}

func (m *_CALDataStatus)  GetParent() CALData {
	return m._CALData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CALDataStatus) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_CALDataStatus) GetBlockStart() uint8 {
	return m.BlockStart
}

func (m *_CALDataStatus) GetStatusBytes() []StatusByte {
	return m.StatusBytes
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewCALDataStatus factory function for _CALDataStatus
func NewCALDataStatus( application ApplicationIdContainer , blockStart uint8 , statusBytes []StatusByte , commandTypeContainer CALCommandTypeContainer , additionalData CALData , requestContext RequestContext ) *_CALDataStatus {
	_result := &_CALDataStatus{
		Application: application,
		BlockStart: blockStart,
		StatusBytes: statusBytes,
    	_CALData: NewCALData(commandTypeContainer, additionalData, requestContext),
	}
	_result._CALData._CALDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCALDataStatus(structType interface{}) CALDataStatus {
    if casted, ok := structType.(CALDataStatus); ok {
		return casted
	}
	if casted, ok := structType.(*CALDataStatus); ok {
		return *casted
	}
	return nil
}

func (m *_CALDataStatus) GetTypeName() string {
	return "CALDataStatus"
}

func (m *_CALDataStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_CALDataStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (application)
	lengthInBits += 8

	// Simple field (blockStart)
	lengthInBits += 8;

	// Array field
	if len(m.StatusBytes) > 0 {
		for i, element := range m.StatusBytes {
			last := i == len(m.StatusBytes) -1
			lengthInBits += element.(interface{GetLengthInBitsConditional(bool) uint16}).GetLengthInBitsConditional(last)
		}
	}

	return lengthInBits
}


func (m *_CALDataStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CALDataStatusParse(theBytes []byte, requestContext RequestContext, commandTypeContainer CALCommandTypeContainer) (CALDataStatus, error) {
	return CALDataStatusParseWithBuffer(utils.NewReadBufferByteBased(theBytes), requestContext, commandTypeContainer)
}

func CALDataStatusParseWithBuffer(readBuffer utils.ReadBuffer, requestContext RequestContext, commandTypeContainer CALCommandTypeContainer) (CALDataStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("CALDataStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CALDataStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (application)
	if pullErr := readBuffer.PullContext("application"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for application")
	}
_application, _applicationErr := ApplicationIdContainerParseWithBuffer(readBuffer)
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field of CALDataStatus")
	}
	application := _application
	if closeErr := readBuffer.CloseContext("application"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for application")
	}

	// Simple Field (blockStart)
_blockStart, _blockStartErr := readBuffer.ReadUint8("blockStart", 8)
	if _blockStartErr != nil {
		return nil, errors.Wrap(_blockStartErr, "Error parsing 'blockStart' field of CALDataStatus")
	}
	blockStart := _blockStart

	// Array field (statusBytes)
	if pullErr := readBuffer.PullContext("statusBytes", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusBytes")
	}
	// Count array
	statusBytes := make([]StatusByte, uint16(commandTypeContainer.NumBytes()) - uint16(uint16(2)))
	// This happens when the size is set conditional to 0
	if len(statusBytes) == 0 {
		statusBytes = nil
	}
	{
		for curItem := uint16(0); curItem < uint16(uint16(commandTypeContainer.NumBytes()) - uint16(uint16(2))); curItem++ {
_item, _err := StatusByteParseWithBuffer(readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'statusBytes' field of CALDataStatus")
			}
			statusBytes[curItem] = _item.(StatusByte)
		}
	}
	if closeErr := readBuffer.CloseContext("statusBytes", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusBytes")
	}

	if closeErr := readBuffer.CloseContext("CALDataStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CALDataStatus")
	}

	// Create a partially initialized instance
	_child := &_CALDataStatus{
		_CALData: &_CALData{
			RequestContext: requestContext,
		},
		Application: application,
		BlockStart: blockStart,
		StatusBytes: statusBytes,
	}
	_child._CALData._CALDataChildRequirements = _child
	return _child, nil
}

func (m *_CALDataStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CALDataStatus) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CALDataStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CALDataStatus")
		}

	// Simple Field (application)
	if pushErr := writeBuffer.PushContext("application"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for application")
	}
	_applicationErr := writeBuffer.WriteSerializable(m.GetApplication())
	if popErr := writeBuffer.PopContext("application"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for application")
	}
	if _applicationErr != nil {
		return errors.Wrap(_applicationErr, "Error serializing 'application' field")
	}

	// Simple Field (blockStart)
	blockStart := uint8(m.GetBlockStart())
	_blockStartErr := writeBuffer.WriteUint8("blockStart", 8, (blockStart))
	if _blockStartErr != nil {
		return errors.Wrap(_blockStartErr, "Error serializing 'blockStart' field")
	}

	// Array Field (statusBytes)
	if pushErr := writeBuffer.PushContext("statusBytes", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for statusBytes")
	}
	for _, _element := range m.GetStatusBytes() {
		_elementErr := writeBuffer.WriteSerializable(_element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'statusBytes' field")
		}
	}
	if popErr := writeBuffer.PopContext("statusBytes", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for statusBytes")
	}

		if popErr := writeBuffer.PopContext("CALDataStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CALDataStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_CALDataStatus) isCALDataStatus() bool {
	return true
}

func (m *_CALDataStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



