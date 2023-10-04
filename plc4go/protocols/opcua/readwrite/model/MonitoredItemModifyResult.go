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

// MonitoredItemModifyResult is the corresponding interface of MonitoredItemModifyResult
type MonitoredItemModifyResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetRevisedSamplingInterval returns RevisedSamplingInterval (property field)
	GetRevisedSamplingInterval() float64
	// GetRevisedQueueSize returns RevisedQueueSize (property field)
	GetRevisedQueueSize() uint32
	// GetFilterResult returns FilterResult (property field)
	GetFilterResult() ExtensionObject
}

// MonitoredItemModifyResultExactly can be used when we want exactly this type and not a type which fulfills MonitoredItemModifyResult.
// This is useful for switch cases.
type MonitoredItemModifyResultExactly interface {
	MonitoredItemModifyResult
	isMonitoredItemModifyResult() bool
}

// _MonitoredItemModifyResult is the data-structure of this message
type _MonitoredItemModifyResult struct {
	*_ExtensionObjectDefinition
	StatusCode              StatusCode
	RevisedSamplingInterval float64
	RevisedQueueSize        uint32
	FilterResult            ExtensionObject
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MonitoredItemModifyResult) GetIdentifier() string {
	return "760"
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredItemModifyResult) InitializeParent(parent ExtensionObjectDefinition) {}

func (m *_MonitoredItemModifyResult) GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredItemModifyResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_MonitoredItemModifyResult) GetRevisedSamplingInterval() float64 {
	return m.RevisedSamplingInterval
}

func (m *_MonitoredItemModifyResult) GetRevisedQueueSize() uint32 {
	return m.RevisedQueueSize
}

func (m *_MonitoredItemModifyResult) GetFilterResult() ExtensionObject {
	return m.FilterResult
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewMonitoredItemModifyResult factory function for _MonitoredItemModifyResult
func NewMonitoredItemModifyResult(statusCode StatusCode, revisedSamplingInterval float64, revisedQueueSize uint32, filterResult ExtensionObject) *_MonitoredItemModifyResult {
	_result := &_MonitoredItemModifyResult{
		StatusCode:                 statusCode,
		RevisedSamplingInterval:    revisedSamplingInterval,
		RevisedQueueSize:           revisedQueueSize,
		FilterResult:               filterResult,
		_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMonitoredItemModifyResult(structType any) MonitoredItemModifyResult {
	if casted, ok := structType.(MonitoredItemModifyResult); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredItemModifyResult); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredItemModifyResult) GetTypeName() string {
	return "MonitoredItemModifyResult"
}

func (m *_MonitoredItemModifyResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (revisedSamplingInterval)
	lengthInBits += 64

	// Simple field (revisedQueueSize)
	lengthInBits += 32

	// Simple field (filterResult)
	lengthInBits += m.FilterResult.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_MonitoredItemModifyResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MonitoredItemModifyResultParse(ctx context.Context, theBytes []byte, identifier string) (MonitoredItemModifyResult, error) {
	return MonitoredItemModifyResultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func MonitoredItemModifyResultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (MonitoredItemModifyResult, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MonitoredItemModifyResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredItemModifyResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (statusCode)
	if pullErr := readBuffer.PullContext("statusCode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for statusCode")
	}
	_statusCode, _statusCodeErr := StatusCodeParseWithBuffer(ctx, readBuffer)
	if _statusCodeErr != nil {
		return nil, errors.Wrap(_statusCodeErr, "Error parsing 'statusCode' field of MonitoredItemModifyResult")
	}
	statusCode := _statusCode.(StatusCode)
	if closeErr := readBuffer.CloseContext("statusCode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for statusCode")
	}

	// Simple Field (revisedSamplingInterval)
	_revisedSamplingInterval, _revisedSamplingIntervalErr := readBuffer.ReadFloat64("revisedSamplingInterval", 64)
	if _revisedSamplingIntervalErr != nil {
		return nil, errors.Wrap(_revisedSamplingIntervalErr, "Error parsing 'revisedSamplingInterval' field of MonitoredItemModifyResult")
	}
	revisedSamplingInterval := _revisedSamplingInterval

	// Simple Field (revisedQueueSize)
	_revisedQueueSize, _revisedQueueSizeErr := readBuffer.ReadUint32("revisedQueueSize", 32)
	if _revisedQueueSizeErr != nil {
		return nil, errors.Wrap(_revisedQueueSizeErr, "Error parsing 'revisedQueueSize' field of MonitoredItemModifyResult")
	}
	revisedQueueSize := _revisedQueueSize

	// Simple Field (filterResult)
	if pullErr := readBuffer.PullContext("filterResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for filterResult")
	}
	_filterResult, _filterResultErr := ExtensionObjectParseWithBuffer(ctx, readBuffer, bool(bool(true)))
	if _filterResultErr != nil {
		return nil, errors.Wrap(_filterResultErr, "Error parsing 'filterResult' field of MonitoredItemModifyResult")
	}
	filterResult := _filterResult.(ExtensionObject)
	if closeErr := readBuffer.CloseContext("filterResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for filterResult")
	}

	if closeErr := readBuffer.CloseContext("MonitoredItemModifyResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredItemModifyResult")
	}

	// Create a partially initialized instance
	_child := &_MonitoredItemModifyResult{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{},
		StatusCode:                 statusCode,
		RevisedSamplingInterval:    revisedSamplingInterval,
		RevisedQueueSize:           revisedQueueSize,
		FilterResult:               filterResult,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_MonitoredItemModifyResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredItemModifyResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredItemModifyResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredItemModifyResult")
		}

		// Simple Field (statusCode)
		if pushErr := writeBuffer.PushContext("statusCode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for statusCode")
		}
		_statusCodeErr := writeBuffer.WriteSerializable(ctx, m.GetStatusCode())
		if popErr := writeBuffer.PopContext("statusCode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for statusCode")
		}
		if _statusCodeErr != nil {
			return errors.Wrap(_statusCodeErr, "Error serializing 'statusCode' field")
		}

		// Simple Field (revisedSamplingInterval)
		revisedSamplingInterval := float64(m.GetRevisedSamplingInterval())
		_revisedSamplingIntervalErr := writeBuffer.WriteFloat64("revisedSamplingInterval", 64, (revisedSamplingInterval))
		if _revisedSamplingIntervalErr != nil {
			return errors.Wrap(_revisedSamplingIntervalErr, "Error serializing 'revisedSamplingInterval' field")
		}

		// Simple Field (revisedQueueSize)
		revisedQueueSize := uint32(m.GetRevisedQueueSize())
		_revisedQueueSizeErr := writeBuffer.WriteUint32("revisedQueueSize", 32, (revisedQueueSize))
		if _revisedQueueSizeErr != nil {
			return errors.Wrap(_revisedQueueSizeErr, "Error serializing 'revisedQueueSize' field")
		}

		// Simple Field (filterResult)
		if pushErr := writeBuffer.PushContext("filterResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for filterResult")
		}
		_filterResultErr := writeBuffer.WriteSerializable(ctx, m.GetFilterResult())
		if popErr := writeBuffer.PopContext("filterResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for filterResult")
		}
		if _filterResultErr != nil {
			return errors.Wrap(_filterResultErr, "Error serializing 'filterResult' field")
		}

		if popErr := writeBuffer.PopContext("MonitoredItemModifyResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredItemModifyResult")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredItemModifyResult) isMonitoredItemModifyResult() bool {
	return true
}

func (m *_MonitoredItemModifyResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
