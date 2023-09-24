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


// CBusPointToMultiPointCommandNormal is the corresponding interface of CBusPointToMultiPointCommandNormal
type CBusPointToMultiPointCommandNormal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CBusPointToMultiPointCommand
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetSalData returns SalData (property field)
	GetSalData() SALData
}

// CBusPointToMultiPointCommandNormalExactly can be used when we want exactly this type and not a type which fulfills CBusPointToMultiPointCommandNormal.
// This is useful for switch cases.
type CBusPointToMultiPointCommandNormalExactly interface {
	CBusPointToMultiPointCommandNormal
	isCBusPointToMultiPointCommandNormal() bool
}

// _CBusPointToMultiPointCommandNormal is the data-structure of this message
type _CBusPointToMultiPointCommandNormal struct {
	*_CBusPointToMultiPointCommand
        Application ApplicationIdContainer
        SalData SALData
	// Reserved Fields
	reservedField0 *byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_CBusPointToMultiPointCommandNormal) InitializeParent(parent CBusPointToMultiPointCommand , peekedApplication byte ) {	m.PeekedApplication = peekedApplication
}

func (m *_CBusPointToMultiPointCommandNormal)  GetParent() CBusPointToMultiPointCommand {
	return m._CBusPointToMultiPointCommand
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_CBusPointToMultiPointCommandNormal) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_CBusPointToMultiPointCommandNormal) GetSalData() SALData {
	return m.SalData
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewCBusPointToMultiPointCommandNormal factory function for _CBusPointToMultiPointCommandNormal
func NewCBusPointToMultiPointCommandNormal( application ApplicationIdContainer , salData SALData , peekedApplication byte , cBusOptions CBusOptions ) *_CBusPointToMultiPointCommandNormal {
	_result := &_CBusPointToMultiPointCommandNormal{
		Application: application,
		SalData: salData,
    	_CBusPointToMultiPointCommand: NewCBusPointToMultiPointCommand(peekedApplication, cBusOptions),
	}
	_result._CBusPointToMultiPointCommand._CBusPointToMultiPointCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastCBusPointToMultiPointCommandNormal(structType any) CBusPointToMultiPointCommandNormal {
    if casted, ok := structType.(CBusPointToMultiPointCommandNormal); ok {
		return casted
	}
	if casted, ok := structType.(*CBusPointToMultiPointCommandNormal); ok {
		return *casted
	}
	return nil
}

func (m *_CBusPointToMultiPointCommandNormal) GetTypeName() string {
	return "CBusPointToMultiPointCommandNormal"
}

func (m *_CBusPointToMultiPointCommandNormal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (application)
	lengthInBits += 8

	// Reserved Field (reserved)
	lengthInBits += 8

	// Simple field (salData)
	lengthInBits += m.SalData.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_CBusPointToMultiPointCommandNormal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func CBusPointToMultiPointCommandNormalParse(ctx context.Context, theBytes []byte, cBusOptions CBusOptions) (CBusPointToMultiPointCommandNormal, error) {
	return CBusPointToMultiPointCommandNormalParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), cBusOptions)
}

func CBusPointToMultiPointCommandNormalParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (CBusPointToMultiPointCommandNormal, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("CBusPointToMultiPointCommandNormal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for CBusPointToMultiPointCommandNormal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (application)
	if pullErr := readBuffer.PullContext("application"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for application")
	}
_application, _applicationErr := ApplicationIdContainerParseWithBuffer(ctx, readBuffer)
	if _applicationErr != nil {
		return nil, errors.Wrap(_applicationErr, "Error parsing 'application' field of CBusPointToMultiPointCommandNormal")
	}
	application := _application
	if closeErr := readBuffer.CloseContext("application"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for application")
	}

	var reservedField0 *byte
	// Reserved Field (Compartmentalized so the "reserved" variable can't leak)
	{
		reserved, _err := readBuffer.ReadByte("reserved")
		if _err != nil {
			return nil, errors.Wrap(_err, "Error parsing 'reserved' field of CBusPointToMultiPointCommandNormal")
		}
		if reserved != byte(0x00) {
			log.Info().Fields(map[string]any{
				"expected value": byte(0x00),
				"got value": reserved,
			}).Msg("Got unexpected response for reserved field.")
			// We save the value, so it can be re-serialized
			reservedField0 = &reserved
		}
	}

	// Simple Field (salData)
	if pullErr := readBuffer.PullContext("salData"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for salData")
	}
_salData, _salDataErr := SALDataParseWithBuffer(ctx, readBuffer , ApplicationId( application.ApplicationId() ) )
	if _salDataErr != nil {
		return nil, errors.Wrap(_salDataErr, "Error parsing 'salData' field of CBusPointToMultiPointCommandNormal")
	}
	salData := _salData.(SALData)
	if closeErr := readBuffer.CloseContext("salData"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for salData")
	}

	if closeErr := readBuffer.CloseContext("CBusPointToMultiPointCommandNormal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for CBusPointToMultiPointCommandNormal")
	}

	// Create a partially initialized instance
	_child := &_CBusPointToMultiPointCommandNormal{
		_CBusPointToMultiPointCommand: &_CBusPointToMultiPointCommand{
			CBusOptions: cBusOptions,
		},
		Application: application,
		SalData: salData,
		reservedField0: reservedField0,
	}
	_child._CBusPointToMultiPointCommand._CBusPointToMultiPointCommandChildRequirements = _child
	return _child, nil
}

func (m *_CBusPointToMultiPointCommandNormal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_CBusPointToMultiPointCommandNormal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("CBusPointToMultiPointCommandNormal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for CBusPointToMultiPointCommandNormal")
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

	// Reserved Field (reserved)
	{
		var reserved byte = byte(0x00)
		if m.reservedField0 != nil {
			log.Info().Fields(map[string]any{
				"expected value": byte(0x00),
				"got value": reserved,
			}).Msg("Overriding reserved field with unexpected value.")
			reserved = *m.reservedField0
		}
		_err := writeBuffer.WriteByte("reserved", reserved)
		if _err != nil {
			return errors.Wrap(_err, "Error serializing 'reserved' field")
		}
	}

	// Simple Field (salData)
	if pushErr := writeBuffer.PushContext("salData"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for salData")
	}
	_salDataErr := writeBuffer.WriteSerializable(ctx, m.GetSalData())
	if popErr := writeBuffer.PopContext("salData"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for salData")
	}
	if _salDataErr != nil {
		return errors.Wrap(_salDataErr, "Error serializing 'salData' field")
	}

		if popErr := writeBuffer.PopContext("CBusPointToMultiPointCommandNormal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for CBusPointToMultiPointCommandNormal")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_CBusPointToMultiPointCommandNormal) isCBusPointToMultiPointCommandNormal() bool {
	return true
}

func (m *_CBusPointToMultiPointCommandNormal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



