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

	"github.com/pkg/errors"
	"github.com/rs/zerolog"

	. "github.com/apache/plc4x/plc4go/spi/codegen/fields"
	. "github.com/apache/plc4x/plc4go/spi/codegen/io"
	"github.com/apache/plc4x/plc4go/spi/utils"
)

// Code generated by code-generation. DO NOT EDIT.

// MonitoredSALLongFormSmartMode is the corresponding interface of MonitoredSALLongFormSmartMode
type MonitoredSALLongFormSmartMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	MonitoredSAL
	// GetTerminatingByte returns TerminatingByte (property field)
	GetTerminatingByte() uint32
	// GetUnitAddress returns UnitAddress (property field)
	GetUnitAddress() UnitAddress
	// GetBridgeAddress returns BridgeAddress (property field)
	GetBridgeAddress() BridgeAddress
	// GetApplication returns Application (property field)
	GetApplication() ApplicationIdContainer
	// GetReservedByte returns ReservedByte (property field)
	GetReservedByte() *byte
	// GetReplyNetwork returns ReplyNetwork (property field)
	GetReplyNetwork() ReplyNetwork
	// GetSalData returns SalData (property field)
	GetSalData() SALData
	// GetIsUnitAddress returns IsUnitAddress (virtual field)
	GetIsUnitAddress() bool
	// IsMonitoredSALLongFormSmartMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsMonitoredSALLongFormSmartMode()
	// CreateBuilder creates a MonitoredSALLongFormSmartModeBuilder
	CreateMonitoredSALLongFormSmartModeBuilder() MonitoredSALLongFormSmartModeBuilder
}

// _MonitoredSALLongFormSmartMode is the data-structure of this message
type _MonitoredSALLongFormSmartMode struct {
	MonitoredSALContract
	TerminatingByte uint32
	UnitAddress     UnitAddress
	BridgeAddress   BridgeAddress
	Application     ApplicationIdContainer
	ReservedByte    *byte
	ReplyNetwork    ReplyNetwork
	SalData         SALData
	// Reserved Fields
	reservedField0 *byte
}

var _ MonitoredSALLongFormSmartMode = (*_MonitoredSALLongFormSmartMode)(nil)
var _ MonitoredSALRequirements = (*_MonitoredSALLongFormSmartMode)(nil)

// NewMonitoredSALLongFormSmartMode factory function for _MonitoredSALLongFormSmartMode
func NewMonitoredSALLongFormSmartMode(salType byte, terminatingByte uint32, unitAddress UnitAddress, bridgeAddress BridgeAddress, application ApplicationIdContainer, reservedByte *byte, replyNetwork ReplyNetwork, salData SALData, cBusOptions CBusOptions) *_MonitoredSALLongFormSmartMode {
	_result := &_MonitoredSALLongFormSmartMode{
		MonitoredSALContract: NewMonitoredSAL(salType, cBusOptions),
		TerminatingByte:      terminatingByte,
		UnitAddress:          unitAddress,
		BridgeAddress:        bridgeAddress,
		Application:          application,
		ReservedByte:         reservedByte,
		ReplyNetwork:         replyNetwork,
		SalData:              salData,
	}
	_result.MonitoredSALContract.(*_MonitoredSAL)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// MonitoredSALLongFormSmartModeBuilder is a builder for MonitoredSALLongFormSmartMode
type MonitoredSALLongFormSmartModeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(terminatingByte uint32, application ApplicationIdContainer) MonitoredSALLongFormSmartModeBuilder
	// WithTerminatingByte adds TerminatingByte (property field)
	WithTerminatingByte(uint32) MonitoredSALLongFormSmartModeBuilder
	// WithUnitAddress adds UnitAddress (property field)
	WithOptionalUnitAddress(UnitAddress) MonitoredSALLongFormSmartModeBuilder
	// WithOptionalUnitAddressBuilder adds UnitAddress (property field) which is build by the builder
	WithOptionalUnitAddressBuilder(func(UnitAddressBuilder) UnitAddressBuilder) MonitoredSALLongFormSmartModeBuilder
	// WithBridgeAddress adds BridgeAddress (property field)
	WithOptionalBridgeAddress(BridgeAddress) MonitoredSALLongFormSmartModeBuilder
	// WithOptionalBridgeAddressBuilder adds BridgeAddress (property field) which is build by the builder
	WithOptionalBridgeAddressBuilder(func(BridgeAddressBuilder) BridgeAddressBuilder) MonitoredSALLongFormSmartModeBuilder
	// WithApplication adds Application (property field)
	WithApplication(ApplicationIdContainer) MonitoredSALLongFormSmartModeBuilder
	// WithReservedByte adds ReservedByte (property field)
	WithOptionalReservedByte(byte) MonitoredSALLongFormSmartModeBuilder
	// WithReplyNetwork adds ReplyNetwork (property field)
	WithOptionalReplyNetwork(ReplyNetwork) MonitoredSALLongFormSmartModeBuilder
	// WithOptionalReplyNetworkBuilder adds ReplyNetwork (property field) which is build by the builder
	WithOptionalReplyNetworkBuilder(func(ReplyNetworkBuilder) ReplyNetworkBuilder) MonitoredSALLongFormSmartModeBuilder
	// WithSalData adds SalData (property field)
	WithOptionalSalData(SALData) MonitoredSALLongFormSmartModeBuilder
	// WithOptionalSalDataBuilder adds SalData (property field) which is build by the builder
	WithOptionalSalDataBuilder(func(SALDataBuilder) SALDataBuilder) MonitoredSALLongFormSmartModeBuilder
	// Build builds the MonitoredSALLongFormSmartMode or returns an error if something is wrong
	Build() (MonitoredSALLongFormSmartMode, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() MonitoredSALLongFormSmartMode
}

// NewMonitoredSALLongFormSmartModeBuilder() creates a MonitoredSALLongFormSmartModeBuilder
func NewMonitoredSALLongFormSmartModeBuilder() MonitoredSALLongFormSmartModeBuilder {
	return &_MonitoredSALLongFormSmartModeBuilder{_MonitoredSALLongFormSmartMode: new(_MonitoredSALLongFormSmartMode)}
}

type _MonitoredSALLongFormSmartModeBuilder struct {
	*_MonitoredSALLongFormSmartMode

	parentBuilder *_MonitoredSALBuilder

	err *utils.MultiError
}

var _ (MonitoredSALLongFormSmartModeBuilder) = (*_MonitoredSALLongFormSmartModeBuilder)(nil)

func (b *_MonitoredSALLongFormSmartModeBuilder) setParent(contract MonitoredSALContract) {
	b.MonitoredSALContract = contract
}

func (b *_MonitoredSALLongFormSmartModeBuilder) WithMandatoryFields(terminatingByte uint32, application ApplicationIdContainer) MonitoredSALLongFormSmartModeBuilder {
	return b.WithTerminatingByte(terminatingByte).WithApplication(application)
}

func (b *_MonitoredSALLongFormSmartModeBuilder) WithTerminatingByte(terminatingByte uint32) MonitoredSALLongFormSmartModeBuilder {
	b.TerminatingByte = terminatingByte
	return b
}

func (b *_MonitoredSALLongFormSmartModeBuilder) WithOptionalUnitAddress(unitAddress UnitAddress) MonitoredSALLongFormSmartModeBuilder {
	b.UnitAddress = unitAddress
	return b
}

func (b *_MonitoredSALLongFormSmartModeBuilder) WithOptionalUnitAddressBuilder(builderSupplier func(UnitAddressBuilder) UnitAddressBuilder) MonitoredSALLongFormSmartModeBuilder {
	builder := builderSupplier(b.UnitAddress.CreateUnitAddressBuilder())
	var err error
	b.UnitAddress, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "UnitAddressBuilder failed"))
	}
	return b
}

func (b *_MonitoredSALLongFormSmartModeBuilder) WithOptionalBridgeAddress(bridgeAddress BridgeAddress) MonitoredSALLongFormSmartModeBuilder {
	b.BridgeAddress = bridgeAddress
	return b
}

func (b *_MonitoredSALLongFormSmartModeBuilder) WithOptionalBridgeAddressBuilder(builderSupplier func(BridgeAddressBuilder) BridgeAddressBuilder) MonitoredSALLongFormSmartModeBuilder {
	builder := builderSupplier(b.BridgeAddress.CreateBridgeAddressBuilder())
	var err error
	b.BridgeAddress, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BridgeAddressBuilder failed"))
	}
	return b
}

func (b *_MonitoredSALLongFormSmartModeBuilder) WithApplication(application ApplicationIdContainer) MonitoredSALLongFormSmartModeBuilder {
	b.Application = application
	return b
}

func (b *_MonitoredSALLongFormSmartModeBuilder) WithOptionalReservedByte(reservedByte byte) MonitoredSALLongFormSmartModeBuilder {
	b.ReservedByte = &reservedByte
	return b
}

func (b *_MonitoredSALLongFormSmartModeBuilder) WithOptionalReplyNetwork(replyNetwork ReplyNetwork) MonitoredSALLongFormSmartModeBuilder {
	b.ReplyNetwork = replyNetwork
	return b
}

func (b *_MonitoredSALLongFormSmartModeBuilder) WithOptionalReplyNetworkBuilder(builderSupplier func(ReplyNetworkBuilder) ReplyNetworkBuilder) MonitoredSALLongFormSmartModeBuilder {
	builder := builderSupplier(b.ReplyNetwork.CreateReplyNetworkBuilder())
	var err error
	b.ReplyNetwork, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "ReplyNetworkBuilder failed"))
	}
	return b
}

func (b *_MonitoredSALLongFormSmartModeBuilder) WithOptionalSalData(salData SALData) MonitoredSALLongFormSmartModeBuilder {
	b.SalData = salData
	return b
}

func (b *_MonitoredSALLongFormSmartModeBuilder) WithOptionalSalDataBuilder(builderSupplier func(SALDataBuilder) SALDataBuilder) MonitoredSALLongFormSmartModeBuilder {
	builder := builderSupplier(b.SalData.CreateSALDataBuilder())
	var err error
	b.SalData, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "SALDataBuilder failed"))
	}
	return b
}

func (b *_MonitoredSALLongFormSmartModeBuilder) Build() (MonitoredSALLongFormSmartMode, error) {
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._MonitoredSALLongFormSmartMode.deepCopy(), nil
}

func (b *_MonitoredSALLongFormSmartModeBuilder) MustBuild() MonitoredSALLongFormSmartMode {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_MonitoredSALLongFormSmartModeBuilder) Done() MonitoredSALBuilder {
	return b.parentBuilder
}

func (b *_MonitoredSALLongFormSmartModeBuilder) buildForMonitoredSAL() (MonitoredSAL, error) {
	return b.Build()
}

func (b *_MonitoredSALLongFormSmartModeBuilder) DeepCopy() any {
	_copy := b.CreateMonitoredSALLongFormSmartModeBuilder().(*_MonitoredSALLongFormSmartModeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateMonitoredSALLongFormSmartModeBuilder creates a MonitoredSALLongFormSmartModeBuilder
func (b *_MonitoredSALLongFormSmartMode) CreateMonitoredSALLongFormSmartModeBuilder() MonitoredSALLongFormSmartModeBuilder {
	if b == nil {
		return NewMonitoredSALLongFormSmartModeBuilder()
	}
	return &_MonitoredSALLongFormSmartModeBuilder{_MonitoredSALLongFormSmartMode: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MonitoredSALLongFormSmartMode) GetParent() MonitoredSALContract {
	return m.MonitoredSALContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MonitoredSALLongFormSmartMode) GetTerminatingByte() uint32 {
	return m.TerminatingByte
}

func (m *_MonitoredSALLongFormSmartMode) GetUnitAddress() UnitAddress {
	return m.UnitAddress
}

func (m *_MonitoredSALLongFormSmartMode) GetBridgeAddress() BridgeAddress {
	return m.BridgeAddress
}

func (m *_MonitoredSALLongFormSmartMode) GetApplication() ApplicationIdContainer {
	return m.Application
}

func (m *_MonitoredSALLongFormSmartMode) GetReservedByte() *byte {
	return m.ReservedByte
}

func (m *_MonitoredSALLongFormSmartMode) GetReplyNetwork() ReplyNetwork {
	return m.ReplyNetwork
}

func (m *_MonitoredSALLongFormSmartMode) GetSalData() SALData {
	return m.SalData
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_MonitoredSALLongFormSmartMode) GetIsUnitAddress() bool {
	ctx := context.Background()
	_ = ctx
	unitAddress := m.GetUnitAddress()
	_ = unitAddress
	bridgeAddress := m.GetBridgeAddress()
	_ = bridgeAddress
	reservedByte := m.GetReservedByte()
	_ = reservedByte
	replyNetwork := m.GetReplyNetwork()
	_ = replyNetwork
	salData := m.GetSalData()
	_ = salData
	return bool(bool((m.GetTerminatingByte() & 0xff) == (0x00)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastMonitoredSALLongFormSmartMode(structType any) MonitoredSALLongFormSmartMode {
	if casted, ok := structType.(MonitoredSALLongFormSmartMode); ok {
		return casted
	}
	if casted, ok := structType.(*MonitoredSALLongFormSmartMode); ok {
		return *casted
	}
	return nil
}

func (m *_MonitoredSALLongFormSmartMode) GetTypeName() string {
	return "MonitoredSALLongFormSmartMode"
}

func (m *_MonitoredSALLongFormSmartMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.MonitoredSALContract.(*_MonitoredSAL).getLengthInBits(ctx))

	// Reserved Field (reserved)
	lengthInBits += 8

	// A virtual field doesn't have any in- or output.

	// Optional Field (unitAddress)
	if m.UnitAddress != nil {
		lengthInBits += m.UnitAddress.GetLengthInBits(ctx)
	}

	// Optional Field (bridgeAddress)
	if m.BridgeAddress != nil {
		lengthInBits += m.BridgeAddress.GetLengthInBits(ctx)
	}

	// Simple field (application)
	lengthInBits += 8

	// Optional Field (reservedByte)
	if m.ReservedByte != nil {
		lengthInBits += 8
	}

	// Optional Field (replyNetwork)
	if m.ReplyNetwork != nil {
		lengthInBits += m.ReplyNetwork.GetLengthInBits(ctx)
	}

	// Optional Field (salData)
	if m.SalData != nil {
		lengthInBits += m.SalData.GetLengthInBits(ctx)
	}

	return lengthInBits
}

func (m *_MonitoredSALLongFormSmartMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_MonitoredSALLongFormSmartMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_MonitoredSAL, cBusOptions CBusOptions) (__monitoredSALLongFormSmartMode MonitoredSALLongFormSmartMode, err error) {
	m.MonitoredSALContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("MonitoredSALLongFormSmartMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MonitoredSALLongFormSmartMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	reservedField0, err := ReadReservedField(ctx, "reserved", ReadByte(readBuffer, 8), byte(0x05))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing reserved field"))
	}
	m.reservedField0 = reservedField0

	terminatingByte, err := ReadPeekField[uint32](ctx, "terminatingByte", ReadUnsignedInt(readBuffer, uint8(24)), 0)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'terminatingByte' field"))
	}
	m.TerminatingByte = terminatingByte

	isUnitAddress, err := ReadVirtualField[bool](ctx, "isUnitAddress", (*bool)(nil), bool((terminatingByte&0xff) == (0x00)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'isUnitAddress' field"))
	}
	_ = isUnitAddress

	var unitAddress UnitAddress
	_unitAddress, err := ReadOptionalField[UnitAddress](ctx, "unitAddress", ReadComplex[UnitAddress](UnitAddressParseWithBuffer, readBuffer), isUnitAddress)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'unitAddress' field"))
	}
	if _unitAddress != nil {
		unitAddress = *_unitAddress
		m.UnitAddress = unitAddress
	}

	var bridgeAddress BridgeAddress
	_bridgeAddress, err := ReadOptionalField[BridgeAddress](ctx, "bridgeAddress", ReadComplex[BridgeAddress](BridgeAddressParseWithBuffer, readBuffer), !(isUnitAddress))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'bridgeAddress' field"))
	}
	if _bridgeAddress != nil {
		bridgeAddress = *_bridgeAddress
		m.BridgeAddress = bridgeAddress
	}

	application, err := ReadEnumField[ApplicationIdContainer](ctx, "application", "ApplicationIdContainer", ReadEnum(ApplicationIdContainerByValue, ReadUnsignedByte(readBuffer, uint8(8))))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'application' field"))
	}
	m.Application = application

	var reservedByte *byte
	reservedByte, err = ReadOptionalField[byte](ctx, "reservedByte", ReadByte(readBuffer, 8), isUnitAddress)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'reservedByte' field"))
	}
	m.ReservedByte = reservedByte

	// Validation
	if !(bool(bool(isUnitAddress) && bool(bool((*reservedByte) == (0x00)))) || bool(!(isUnitAddress))) {
		return nil, errors.WithStack(utils.ParseValidationError{Message: "invalid unit address"})
	}

	var replyNetwork ReplyNetwork
	_replyNetwork, err := ReadOptionalField[ReplyNetwork](ctx, "replyNetwork", ReadComplex[ReplyNetwork](ReplyNetworkParseWithBuffer, readBuffer), !(isUnitAddress))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'replyNetwork' field"))
	}
	if _replyNetwork != nil {
		replyNetwork = *_replyNetwork
		m.ReplyNetwork = replyNetwork
	}

	var salData SALData
	_salData, err := ReadOptionalField[SALData](ctx, "salData", ReadComplex[SALData](SALDataParseWithBufferProducer[SALData]((ApplicationId)(application.ApplicationId())), readBuffer), true)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'salData' field"))
	}
	if _salData != nil {
		salData = *_salData
		m.SalData = salData
	}

	if closeErr := readBuffer.CloseContext("MonitoredSALLongFormSmartMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MonitoredSALLongFormSmartMode")
	}

	return m, nil
}

func (m *_MonitoredSALLongFormSmartMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MonitoredSALLongFormSmartMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MonitoredSALLongFormSmartMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MonitoredSALLongFormSmartMode")
		}

		if err := WriteReservedField[byte](ctx, "reserved", byte(0x05), WriteByte(writeBuffer, 8)); err != nil {
			return errors.Wrap(err, "Error serializing 'reserved' field number 1")
		}
		// Virtual field
		isUnitAddress := m.GetIsUnitAddress()
		_ = isUnitAddress
		if _isUnitAddressErr := writeBuffer.WriteVirtual(ctx, "isUnitAddress", m.GetIsUnitAddress()); _isUnitAddressErr != nil {
			return errors.Wrap(_isUnitAddressErr, "Error serializing 'isUnitAddress' field")
		}

		if err := WriteOptionalField[UnitAddress](ctx, "unitAddress", GetRef(m.GetUnitAddress()), WriteComplex[UnitAddress](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'unitAddress' field")
		}

		if err := WriteOptionalField[BridgeAddress](ctx, "bridgeAddress", GetRef(m.GetBridgeAddress()), WriteComplex[BridgeAddress](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'bridgeAddress' field")
		}

		if err := WriteSimpleEnumField[ApplicationIdContainer](ctx, "application", "ApplicationIdContainer", m.GetApplication(), WriteEnum[ApplicationIdContainer, uint8](ApplicationIdContainer.GetValue, ApplicationIdContainer.PLC4XEnumName, WriteUnsignedByte(writeBuffer, 8))); err != nil {
			return errors.Wrap(err, "Error serializing 'application' field")
		}

		if err := WriteOptionalField[byte](ctx, "reservedByte", m.GetReservedByte(), WriteByte(writeBuffer, 8), true); err != nil {
			return errors.Wrap(err, "Error serializing 'reservedByte' field")
		}

		if err := WriteOptionalField[ReplyNetwork](ctx, "replyNetwork", GetRef(m.GetReplyNetwork()), WriteComplex[ReplyNetwork](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'replyNetwork' field")
		}

		if err := WriteOptionalField[SALData](ctx, "salData", GetRef(m.GetSalData()), WriteComplex[SALData](writeBuffer), true); err != nil {
			return errors.Wrap(err, "Error serializing 'salData' field")
		}

		if popErr := writeBuffer.PopContext("MonitoredSALLongFormSmartMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MonitoredSALLongFormSmartMode")
		}
		return nil
	}
	return m.MonitoredSALContract.(*_MonitoredSAL).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_MonitoredSALLongFormSmartMode) IsMonitoredSALLongFormSmartMode() {}

func (m *_MonitoredSALLongFormSmartMode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_MonitoredSALLongFormSmartMode) deepCopy() *_MonitoredSALLongFormSmartMode {
	if m == nil {
		return nil
	}
	_MonitoredSALLongFormSmartModeCopy := &_MonitoredSALLongFormSmartMode{
		m.MonitoredSALContract.(*_MonitoredSAL).deepCopy(),
		m.TerminatingByte,
		m.UnitAddress.DeepCopy().(UnitAddress),
		m.BridgeAddress.DeepCopy().(BridgeAddress),
		m.Application,
		utils.CopyPtr[byte](m.ReservedByte),
		m.ReplyNetwork.DeepCopy().(ReplyNetwork),
		m.SalData.DeepCopy().(SALData),
		m.reservedField0,
	}
	m.MonitoredSALContract.(*_MonitoredSAL)._SubType = m
	return _MonitoredSALLongFormSmartModeCopy
}

func (m *_MonitoredSALLongFormSmartMode) String() string {
	if m == nil {
		return "<nil>"
	}
	wb := utils.NewWriteBufferBoxBased(
		utils.WithWriteBufferBoxBasedMergeSingleBoxes(),
		utils.WithWriteBufferBoxBasedOmitEmptyBoxes(),
		utils.WithWriteBufferBoxBasedPrintPosLengthFooter(),
	)
	if err := wb.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return wb.GetBox().String()
}
