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
	"io"
)

	// Code generated by code-generation. DO NOT EDIT.


// MonitoredSALShortFormBasicMode is the corresponding interface of MonitoredSALShortFormBasicMode
type MonitoredSALShortFormBasicMode interface {
	utils.LengthAware
	utils.Serializable
	MonitoredSAL
	// GetCounts returns Counts (property field)
	GetCounts() byte
	// GetBridgeCount returns BridgeCount (property field)
	GetBridgeCount() *uint8
	// GetNetworkNumber returns NetworkNumber (property field)
	GetNetworkNumber() *uint8
	// GetNoCounts returns NoCounts (property field)
	GetNoCounts() *byte
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetSalData returns SalData (property field)
	GetSalData() SALData
}

// MonitoredSALShortFormBasicModeExactly can be used when we want exactly this type and not a type which fulfills MonitoredSALShortFormBasicMode.
// This is useful for switch cases.
type MonitoredSALShortFormBasicModeExactly interface {
	MonitoredSALShortFormBasicMode
	isMonitoredSALShortFormBasicMode() bool
}

// _MonitoredSALShortFormBasicMode is the data-structure of this message
type _MonitoredSALShortFormBasicMode struct {
	*_MonitoredSAL
        Counts byte
        BridgeCount *uint8
        NetworkNumber *uint8
        NoCounts *byte
        Application ApplicationIdContainer
        SalData SALData
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredSALShortFormBasicMode) InitializeParent(parent MonitoredSAL , salType byte ) {	m.SalType = salType
}

func (m *_MonitoredSALShortFormBasicMode)  GetParent() MonitoredSAL {
	return m._MonitoredSAL
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredSALShortFormBasicMode) GetCounts() byte {
	return m.Counts
}

func (m *_MonitoredSALShortFormBasicMode) GetBridgeCount() *uint8 {
	return m.BridgeCount
}

func (m *_MonitoredSALShortFormBasicMode) GetNetworkNumber() *uint8 {
	return m.NetworkNumber
}

func (m *_MonitoredSALShortFormBasicMode) GetNoCounts() *byte {
	return m.NoCounts
}

func (m *_MonitoredSALShortFormBasicMode) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_MonitoredSALShortFormBasicMode) GetSalData() SALData {
	return m.SalData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewMonitoredSALShortFormBasicMode factory function for _MonitoredSALShortFormBasicMode
func NewMonitoredSALShortFormBasicMode( counts byte , bridgeCount *uint8 , networkNumber *uint8 , noCounts *byte , application ApplicationIdContainer , salData SALData , salType byte , cBusOptions CBusOptions ) *_MonitoredSALShortFormBasicMode {
	_result := &_MonitoredSALShortFormBasicMode{
		Counts: counts,
		BridgeCount: bridgeCount,
		NetworkNumber: networkNumber,
		NoCounts: noCounts,
		Application: application,
		SalData: salData,
    	_MonitoredSAL: NewMonitoredSAL(salType, cBusOptions),
	}
	_result._MonitoredSAL._MonitoredSALChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoredSALShortFormBasicMode(structType interface{}) MonitoredSALShortFormBasicMode {
    if casted, ok := structType.(MonitoredSALShortFormBasicMode); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredSALShortFormBasicMode); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredSALShortFormBasicMode) GetTypeName() string {
	return "MonitoredSALShortFormBasicMode"
}

func (m *_MonitoredSALShortFormBasicMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Optional Field (bridgeCount)
	if m.BridgeCount != nil {
		lengthInBits += 8
	}

	// Optional Field (networkNumber)
	if m.NetworkNumber != nil {
		lengthInBits += 8
	}

	// Optional Field (noCounts)
	if m.NoCounts != nil {
		lengthInBits += 8
	}

	// Simple field (application)
	lengthInBits += 8

	// Optional Field (salData)
	if m.SalData != nil {
		lengthInBits += m.SalData.GetLengthInBits(ctx)
	}

	return lengthInBits
}


func (m *_MonitoredSALShortFormBasicMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MonitoredSALShortFormBasicModeParse(theBytes []byte, cBusOptions CBusOptions) (MonitoredSALShortFormBasicMode, error) {
	return MonitoredSALShortFormBasicModeParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func MonitoredSALShortFormBasicModeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (MonitoredSALShortFormBasicMode, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MonitoredSALShortFormBasicMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredSALShortFormBasicMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

				// Peek Field (counts)
				currentPos = positionAware.GetPos()
				counts, _err := readBuffer.ReadByte("counts")
				if _err != nil {
					return nil, errors.Wrap(_err, "Error parsing 'counts' field of MonitoredSALShortFormBasicMode")
				}

				readBuffer.Reset(currentPos)

	// Optional Field (bridgeCount) (Can be skipped, if a given expression evaluates to false)
	var bridgeCount *uint8 = nil
	if bool((counts) != (0x00)) {
		_val, _err := readBuffer.ReadUint8("bridgeCount", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'bridgeCount' field of MonitoredSALShortFormBasicMode")
		}
		bridgeCount = &_val
	}

	// Optional Field (networkNumber) (Can be skipped, if a given expression evaluates to false)
	var networkNumber *uint8 = nil
	if bool((counts) != (0x00)) {
		_val, _err := readBuffer.ReadUint8("networkNumber", 8)
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'networkNumber' field of MonitoredSALShortFormBasicMode")
		}
		networkNumber = &_val
	}

	// Optional Field (noCounts) (Can be skipped, if a given expression evaluates to false)
	var noCounts *byte = nil
	if bool((counts) == (0x00)) {
		_val, _err := readBuffer.ReadByte("noCounts")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'noCounts' field of MonitoredSALShortFormBasicMode")
		}
		noCounts = &_val
	}

	// Simple Field (application)
	if pullErr := readBuffer.PullContext("application"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for application")
	}
_application, _applicationErr := ApplicationIdContainerParseWithBuffer(ctx, readBuffer)
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field of MonitoredSALShortFormBasicMode")
	}
	application := _application
	if closeErr := readBuffer.CloseContext("application"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for application")
	}

	// Optional Field (salData) (Can be skipped, if a given expression evaluates to false)
	var salData SALData = nil
{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("salData"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for salData")
		}
_val, _err := SALDataParseWithBuffer(ctx, readBuffer , application.ApplicationId() )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'salData' field of MonitoredSALShortFormBasicMode")
		default:
			salData = _val.(SALData)
			if closeErr := readBuffer.CloseContext("salData"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for salData")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("MonitoredSALShortFormBasicMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredSALShortFormBasicMode")
	}

	// Create a partially initialized instance
	_child := &_MonitoredSALShortFormBasicMode{
		_MonitoredSAL: &_MonitoredSAL{
			CBusOptions: cBusOptions,
		},
		Counts: counts,
		BridgeCount: bridgeCount,
		NetworkNumber: networkNumber,
		NoCounts: noCounts,
		Application: application,
		SalData: salData,
	}
	_child._MonitoredSAL._MonitoredSALChildRequirements = _child
	return _child, nil
}

func (m *_MonitoredSALShortFormBasicMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredSALShortFormBasicMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredSALShortFormBasicMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredSALShortFormBasicMode")
		}

	// Optional Field (bridgeCount) (Can be skipped, if the value is null)
	var bridgeCount *uint8 = nil
	if m.GetBridgeCount() != nil {
		bridgeCount = m.GetBridgeCount()
		_bridgeCountErr := writeBuffer.WriteUint8("bridgeCount", 8, *(bridgeCount))
		if _bridgeCountErr != nil {
			return errors.Wrap(_bridgeCountErr, "Error serializing 'bridgeCount' field")
		}
	}

	// Optional Field (networkNumber) (Can be skipped, if the value is null)
	var networkNumber *uint8 = nil
	if m.GetNetworkNumber() != nil {
		networkNumber = m.GetNetworkNumber()
		_networkNumberErr := writeBuffer.WriteUint8("networkNumber", 8, *(networkNumber))
		if _networkNumberErr != nil {
			return errors.Wrap(_networkNumberErr, "Error serializing 'networkNumber' field")
		}
	}

	// Optional Field (noCounts) (Can be skipped, if the value is null)
	var noCounts *byte = nil
	if m.GetNoCounts() != nil {
		noCounts = m.GetNoCounts()
		_noCountsErr := writeBuffer.WriteByte("noCounts", *(noCounts))
		if _noCountsErr != nil {
			return errors.Wrap(_noCountsErr, "Error serializing 'noCounts' field")
		}
	}

	// Simple Field (application)
	if pushErr := writeBuffer.PushContext("application"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for application")
	}
	_applicationErr := writeBuffer.WriteSerializable(ctx, m.GetApplication())
	if popErr := writeBuffer.PopContext("application"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for application")
	}
	if _applicationErr != nil {
		return errors.Wrap(_applicationErr, "Error serializing 'application' field")
	}

	// Optional Field (salData) (Can be skipped, if the value is null)
	var salData SALData = nil
	if m.GetSalData() != nil {
		if pushErr := writeBuffer.PushContext("salData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for salData")
		}
		salData = m.GetSalData()
		_salDataErr := writeBuffer.WriteSerializable(ctx, salData)
		if popErr := writeBuffer.PopContext("salData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for salData")
		}
		if _salDataErr != nil {
			return errors.Wrap(_salDataErr, "Error serializing 'salData' field")
		}
	}

		if popErr := writeBuffer.PopContext("MonitoredSALShortFormBasicMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredSALShortFormBasicMode")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_MonitoredSALShortFormBasicMode) isMonitoredSALShortFormBasicMode() bool {
	return true
}

func (m *_MonitoredSALShortFormBasicMode) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



