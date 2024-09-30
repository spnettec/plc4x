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

// NetworkGroupDataType is the corresponding interface of NetworkGroupDataType
type NetworkGroupDataType interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetServerUri returns ServerUri (property field)
	GetServerUri() PascalString
	// GetNoOfNetworkPaths returns NoOfNetworkPaths (property field)
	GetNoOfNetworkPaths() int32
	// GetNetworkPaths returns NetworkPaths (property field)
	GetNetworkPaths() []ExtensionObjectDefinition
	// IsNetworkGroupDataType is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsNetworkGroupDataType()
	// CreateBuilder creates a NetworkGroupDataTypeBuilder
	CreateNetworkGroupDataTypeBuilder() NetworkGroupDataTypeBuilder
}

// _NetworkGroupDataType is the data-structure of this message
type _NetworkGroupDataType struct {
	ExtensionObjectDefinitionContract
	ServerUri        PascalString
	NoOfNetworkPaths int32
	NetworkPaths     []ExtensionObjectDefinition
}

var _ NetworkGroupDataType = (*_NetworkGroupDataType)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_NetworkGroupDataType)(nil)

// NewNetworkGroupDataType factory function for _NetworkGroupDataType
func NewNetworkGroupDataType(serverUri PascalString, noOfNetworkPaths int32, networkPaths []ExtensionObjectDefinition) *_NetworkGroupDataType {
	if serverUri == nil {
		panic("serverUri of type PascalString for NetworkGroupDataType must not be nil")
	}
	_result := &_NetworkGroupDataType{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		ServerUri:                         serverUri,
		NoOfNetworkPaths:                  noOfNetworkPaths,
		NetworkPaths:                      networkPaths,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// NetworkGroupDataTypeBuilder is a builder for NetworkGroupDataType
type NetworkGroupDataTypeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(serverUri PascalString, noOfNetworkPaths int32, networkPaths []ExtensionObjectDefinition) NetworkGroupDataTypeBuilder
	// WithServerUri adds ServerUri (property field)
	WithServerUri(PascalString) NetworkGroupDataTypeBuilder
	// WithServerUriBuilder adds ServerUri (property field) which is build by the builder
	WithServerUriBuilder(func(PascalStringBuilder) PascalStringBuilder) NetworkGroupDataTypeBuilder
	// WithNoOfNetworkPaths adds NoOfNetworkPaths (property field)
	WithNoOfNetworkPaths(int32) NetworkGroupDataTypeBuilder
	// WithNetworkPaths adds NetworkPaths (property field)
	WithNetworkPaths(...ExtensionObjectDefinition) NetworkGroupDataTypeBuilder
	// Build builds the NetworkGroupDataType or returns an error if something is wrong
	Build() (NetworkGroupDataType, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() NetworkGroupDataType
}

// NewNetworkGroupDataTypeBuilder() creates a NetworkGroupDataTypeBuilder
func NewNetworkGroupDataTypeBuilder() NetworkGroupDataTypeBuilder {
	return &_NetworkGroupDataTypeBuilder{_NetworkGroupDataType: new(_NetworkGroupDataType)}
}

type _NetworkGroupDataTypeBuilder struct {
	*_NetworkGroupDataType

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (NetworkGroupDataTypeBuilder) = (*_NetworkGroupDataTypeBuilder)(nil)

func (b *_NetworkGroupDataTypeBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_NetworkGroupDataTypeBuilder) WithMandatoryFields(serverUri PascalString, noOfNetworkPaths int32, networkPaths []ExtensionObjectDefinition) NetworkGroupDataTypeBuilder {
	return b.WithServerUri(serverUri).WithNoOfNetworkPaths(noOfNetworkPaths).WithNetworkPaths(networkPaths...)
}

func (b *_NetworkGroupDataTypeBuilder) WithServerUri(serverUri PascalString) NetworkGroupDataTypeBuilder {
	b.ServerUri = serverUri
	return b
}

func (b *_NetworkGroupDataTypeBuilder) WithServerUriBuilder(builderSupplier func(PascalStringBuilder) PascalStringBuilder) NetworkGroupDataTypeBuilder {
	builder := builderSupplier(b.ServerUri.CreatePascalStringBuilder())
	var err error
	b.ServerUri, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "PascalStringBuilder failed"))
	}
	return b
}

func (b *_NetworkGroupDataTypeBuilder) WithNoOfNetworkPaths(noOfNetworkPaths int32) NetworkGroupDataTypeBuilder {
	b.NoOfNetworkPaths = noOfNetworkPaths
	return b
}

func (b *_NetworkGroupDataTypeBuilder) WithNetworkPaths(networkPaths ...ExtensionObjectDefinition) NetworkGroupDataTypeBuilder {
	b.NetworkPaths = networkPaths
	return b
}

func (b *_NetworkGroupDataTypeBuilder) Build() (NetworkGroupDataType, error) {
	if b.ServerUri == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'serverUri' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._NetworkGroupDataType.deepCopy(), nil
}

func (b *_NetworkGroupDataTypeBuilder) MustBuild() NetworkGroupDataType {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_NetworkGroupDataTypeBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_NetworkGroupDataTypeBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_NetworkGroupDataTypeBuilder) DeepCopy() any {
	_copy := b.CreateNetworkGroupDataTypeBuilder().(*_NetworkGroupDataTypeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateNetworkGroupDataTypeBuilder creates a NetworkGroupDataTypeBuilder
func (b *_NetworkGroupDataType) CreateNetworkGroupDataTypeBuilder() NetworkGroupDataTypeBuilder {
	if b == nil {
		return NewNetworkGroupDataTypeBuilder()
	}
	return &_NetworkGroupDataTypeBuilder{_NetworkGroupDataType: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_NetworkGroupDataType) GetIdentifier() string {
	return "11946"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_NetworkGroupDataType) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_NetworkGroupDataType) GetServerUri() PascalString {
	return m.ServerUri
}

func (m *_NetworkGroupDataType) GetNoOfNetworkPaths() int32 {
	return m.NoOfNetworkPaths
}

func (m *_NetworkGroupDataType) GetNetworkPaths() []ExtensionObjectDefinition {
	return m.NetworkPaths
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastNetworkGroupDataType(structType any) NetworkGroupDataType {
	if casted, ok := structType.(NetworkGroupDataType); ok {
		return casted
	}
	if casted, ok := structType.(*NetworkGroupDataType); ok {
		return *casted
	}
	return nil
}

func (m *_NetworkGroupDataType) GetTypeName() string {
	return "NetworkGroupDataType"
}

func (m *_NetworkGroupDataType) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (serverUri)
	lengthInBits += m.ServerUri.GetLengthInBits(ctx)

	// Simple field (noOfNetworkPaths)
	lengthInBits += 32

	// Array field
	if len(m.NetworkPaths) > 0 {
		for _curItem, element := range m.NetworkPaths {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.NetworkPaths), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_NetworkGroupDataType) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_NetworkGroupDataType) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__networkGroupDataType NetworkGroupDataType, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("NetworkGroupDataType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for NetworkGroupDataType")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	serverUri, err := ReadSimpleField[PascalString](ctx, "serverUri", ReadComplex[PascalString](PascalStringParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'serverUri' field"))
	}
	m.ServerUri = serverUri

	noOfNetworkPaths, err := ReadSimpleField(ctx, "noOfNetworkPaths", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfNetworkPaths' field"))
	}
	m.NoOfNetworkPaths = noOfNetworkPaths

	networkPaths, err := ReadCountArrayField[ExtensionObjectDefinition](ctx, "networkPaths", ReadComplex[ExtensionObjectDefinition](ExtensionObjectDefinitionParseWithBufferProducer[ExtensionObjectDefinition]((string)("11945")), readBuffer), uint64(noOfNetworkPaths))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'networkPaths' field"))
	}
	m.NetworkPaths = networkPaths

	if closeErr := readBuffer.CloseContext("NetworkGroupDataType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for NetworkGroupDataType")
	}

	return m, nil
}

func (m *_NetworkGroupDataType) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_NetworkGroupDataType) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("NetworkGroupDataType"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for NetworkGroupDataType")
		}

		if err := WriteSimpleField[PascalString](ctx, "serverUri", m.GetServerUri(), WriteComplex[PascalString](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'serverUri' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfNetworkPaths", m.GetNoOfNetworkPaths(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfNetworkPaths' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "networkPaths", m.GetNetworkPaths(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'networkPaths' field")
		}

		if popErr := writeBuffer.PopContext("NetworkGroupDataType"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for NetworkGroupDataType")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_NetworkGroupDataType) IsNetworkGroupDataType() {}

func (m *_NetworkGroupDataType) DeepCopy() any {
	return m.deepCopy()
}

func (m *_NetworkGroupDataType) deepCopy() *_NetworkGroupDataType {
	if m == nil {
		return nil
	}
	_NetworkGroupDataTypeCopy := &_NetworkGroupDataType{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.ServerUri.DeepCopy().(PascalString),
		m.NoOfNetworkPaths,
		utils.DeepCopySlice[ExtensionObjectDefinition, ExtensionObjectDefinition](m.NetworkPaths),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _NetworkGroupDataTypeCopy
}

func (m *_NetworkGroupDataType) String() string {
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
