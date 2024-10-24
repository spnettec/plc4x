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

// ContentFilterElementResult is the corresponding interface of ContentFilterElementResult
type ContentFilterElementResult interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	ExtensionObjectDefinition
	// GetStatusCode returns StatusCode (property field)
	GetStatusCode() StatusCode
	// GetNoOfOperandStatusCodes returns NoOfOperandStatusCodes (property field)
	GetNoOfOperandStatusCodes() int32
	// GetOperandStatusCodes returns OperandStatusCodes (property field)
	GetOperandStatusCodes() []StatusCode
	// GetNoOfOperandDiagnosticInfos returns NoOfOperandDiagnosticInfos (property field)
	GetNoOfOperandDiagnosticInfos() int32
	// GetOperandDiagnosticInfos returns OperandDiagnosticInfos (property field)
	GetOperandDiagnosticInfos() []DiagnosticInfo
	// IsContentFilterElementResult is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsContentFilterElementResult()
	// CreateBuilder creates a ContentFilterElementResultBuilder
	CreateContentFilterElementResultBuilder() ContentFilterElementResultBuilder
}

// _ContentFilterElementResult is the data-structure of this message
type _ContentFilterElementResult struct {
	ExtensionObjectDefinitionContract
	StatusCode                 StatusCode
	NoOfOperandStatusCodes     int32
	OperandStatusCodes         []StatusCode
	NoOfOperandDiagnosticInfos int32
	OperandDiagnosticInfos     []DiagnosticInfo
}

var _ ContentFilterElementResult = (*_ContentFilterElementResult)(nil)
var _ ExtensionObjectDefinitionRequirements = (*_ContentFilterElementResult)(nil)

// NewContentFilterElementResult factory function for _ContentFilterElementResult
func NewContentFilterElementResult(statusCode StatusCode, noOfOperandStatusCodes int32, operandStatusCodes []StatusCode, noOfOperandDiagnosticInfos int32, operandDiagnosticInfos []DiagnosticInfo) *_ContentFilterElementResult {
	if statusCode == nil {
		panic("statusCode of type StatusCode for ContentFilterElementResult must not be nil")
	}
	_result := &_ContentFilterElementResult{
		ExtensionObjectDefinitionContract: NewExtensionObjectDefinition(),
		StatusCode:                        statusCode,
		NoOfOperandStatusCodes:            noOfOperandStatusCodes,
		OperandStatusCodes:                operandStatusCodes,
		NoOfOperandDiagnosticInfos:        noOfOperandDiagnosticInfos,
		OperandDiagnosticInfos:            operandDiagnosticInfos,
	}
	_result.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// ContentFilterElementResultBuilder is a builder for ContentFilterElementResult
type ContentFilterElementResultBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(statusCode StatusCode, noOfOperandStatusCodes int32, operandStatusCodes []StatusCode, noOfOperandDiagnosticInfos int32, operandDiagnosticInfos []DiagnosticInfo) ContentFilterElementResultBuilder
	// WithStatusCode adds StatusCode (property field)
	WithStatusCode(StatusCode) ContentFilterElementResultBuilder
	// WithStatusCodeBuilder adds StatusCode (property field) which is build by the builder
	WithStatusCodeBuilder(func(StatusCodeBuilder) StatusCodeBuilder) ContentFilterElementResultBuilder
	// WithNoOfOperandStatusCodes adds NoOfOperandStatusCodes (property field)
	WithNoOfOperandStatusCodes(int32) ContentFilterElementResultBuilder
	// WithOperandStatusCodes adds OperandStatusCodes (property field)
	WithOperandStatusCodes(...StatusCode) ContentFilterElementResultBuilder
	// WithNoOfOperandDiagnosticInfos adds NoOfOperandDiagnosticInfos (property field)
	WithNoOfOperandDiagnosticInfos(int32) ContentFilterElementResultBuilder
	// WithOperandDiagnosticInfos adds OperandDiagnosticInfos (property field)
	WithOperandDiagnosticInfos(...DiagnosticInfo) ContentFilterElementResultBuilder
	// Build builds the ContentFilterElementResult or returns an error if something is wrong
	Build() (ContentFilterElementResult, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() ContentFilterElementResult
}

// NewContentFilterElementResultBuilder() creates a ContentFilterElementResultBuilder
func NewContentFilterElementResultBuilder() ContentFilterElementResultBuilder {
	return &_ContentFilterElementResultBuilder{_ContentFilterElementResult: new(_ContentFilterElementResult)}
}

type _ContentFilterElementResultBuilder struct {
	*_ContentFilterElementResult

	parentBuilder *_ExtensionObjectDefinitionBuilder

	err *utils.MultiError
}

var _ (ContentFilterElementResultBuilder) = (*_ContentFilterElementResultBuilder)(nil)

func (b *_ContentFilterElementResultBuilder) setParent(contract ExtensionObjectDefinitionContract) {
	b.ExtensionObjectDefinitionContract = contract
}

func (b *_ContentFilterElementResultBuilder) WithMandatoryFields(statusCode StatusCode, noOfOperandStatusCodes int32, operandStatusCodes []StatusCode, noOfOperandDiagnosticInfos int32, operandDiagnosticInfos []DiagnosticInfo) ContentFilterElementResultBuilder {
	return b.WithStatusCode(statusCode).WithNoOfOperandStatusCodes(noOfOperandStatusCodes).WithOperandStatusCodes(operandStatusCodes...).WithNoOfOperandDiagnosticInfos(noOfOperandDiagnosticInfos).WithOperandDiagnosticInfos(operandDiagnosticInfos...)
}

func (b *_ContentFilterElementResultBuilder) WithStatusCode(statusCode StatusCode) ContentFilterElementResultBuilder {
	b.StatusCode = statusCode
	return b
}

func (b *_ContentFilterElementResultBuilder) WithStatusCodeBuilder(builderSupplier func(StatusCodeBuilder) StatusCodeBuilder) ContentFilterElementResultBuilder {
	builder := builderSupplier(b.StatusCode.CreateStatusCodeBuilder())
	var err error
	b.StatusCode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "StatusCodeBuilder failed"))
	}
	return b
}

func (b *_ContentFilterElementResultBuilder) WithNoOfOperandStatusCodes(noOfOperandStatusCodes int32) ContentFilterElementResultBuilder {
	b.NoOfOperandStatusCodes = noOfOperandStatusCodes
	return b
}

func (b *_ContentFilterElementResultBuilder) WithOperandStatusCodes(operandStatusCodes ...StatusCode) ContentFilterElementResultBuilder {
	b.OperandStatusCodes = operandStatusCodes
	return b
}

func (b *_ContentFilterElementResultBuilder) WithNoOfOperandDiagnosticInfos(noOfOperandDiagnosticInfos int32) ContentFilterElementResultBuilder {
	b.NoOfOperandDiagnosticInfos = noOfOperandDiagnosticInfos
	return b
}

func (b *_ContentFilterElementResultBuilder) WithOperandDiagnosticInfos(operandDiagnosticInfos ...DiagnosticInfo) ContentFilterElementResultBuilder {
	b.OperandDiagnosticInfos = operandDiagnosticInfos
	return b
}

func (b *_ContentFilterElementResultBuilder) Build() (ContentFilterElementResult, error) {
	if b.StatusCode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'statusCode' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._ContentFilterElementResult.deepCopy(), nil
}

func (b *_ContentFilterElementResultBuilder) MustBuild() ContentFilterElementResult {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_ContentFilterElementResultBuilder) Done() ExtensionObjectDefinitionBuilder {
	return b.parentBuilder
}

func (b *_ContentFilterElementResultBuilder) buildForExtensionObjectDefinition() (ExtensionObjectDefinition, error) {
	return b.Build()
}

func (b *_ContentFilterElementResultBuilder) DeepCopy() any {
	_copy := b.CreateContentFilterElementResultBuilder().(*_ContentFilterElementResultBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateContentFilterElementResultBuilder creates a ContentFilterElementResultBuilder
func (b *_ContentFilterElementResult) CreateContentFilterElementResultBuilder() ContentFilterElementResultBuilder {
	if b == nil {
		return NewContentFilterElementResultBuilder()
	}
	return &_ContentFilterElementResultBuilder{_ContentFilterElementResult: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ContentFilterElementResult) GetIdentifier() string {
	return "606"
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ContentFilterElementResult) GetParent() ExtensionObjectDefinitionContract {
	return m.ExtensionObjectDefinitionContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ContentFilterElementResult) GetStatusCode() StatusCode {
	return m.StatusCode
}

func (m *_ContentFilterElementResult) GetNoOfOperandStatusCodes() int32 {
	return m.NoOfOperandStatusCodes
}

func (m *_ContentFilterElementResult) GetOperandStatusCodes() []StatusCode {
	return m.OperandStatusCodes
}

func (m *_ContentFilterElementResult) GetNoOfOperandDiagnosticInfos() int32 {
	return m.NoOfOperandDiagnosticInfos
}

func (m *_ContentFilterElementResult) GetOperandDiagnosticInfos() []DiagnosticInfo {
	return m.OperandDiagnosticInfos
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastContentFilterElementResult(structType any) ContentFilterElementResult {
	if casted, ok := structType.(ContentFilterElementResult); ok {
		return casted
	}
	if casted, ok := structType.(*ContentFilterElementResult); ok {
		return *casted
	}
	return nil
}

func (m *_ContentFilterElementResult) GetTypeName() string {
	return "ContentFilterElementResult"
}

func (m *_ContentFilterElementResult) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).getLengthInBits(ctx))

	// Simple field (statusCode)
	lengthInBits += m.StatusCode.GetLengthInBits(ctx)

	// Simple field (noOfOperandStatusCodes)
	lengthInBits += 32

	// Array field
	if len(m.OperandStatusCodes) > 0 {
		for _curItem, element := range m.OperandStatusCodes {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.OperandStatusCodes), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	// Simple field (noOfOperandDiagnosticInfos)
	lengthInBits += 32

	// Array field
	if len(m.OperandDiagnosticInfos) > 0 {
		for _curItem, element := range m.OperandDiagnosticInfos {
			arrayCtx := utils.CreateArrayContext(ctx, len(m.OperandDiagnosticInfos), _curItem)
			_ = arrayCtx
			_ = _curItem
			lengthInBits += element.(interface{ GetLengthInBits(context.Context) uint16 }).GetLengthInBits(arrayCtx)
		}
	}

	return lengthInBits
}

func (m *_ContentFilterElementResult) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_ContentFilterElementResult) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_ExtensionObjectDefinition, identifier string) (__contentFilterElementResult ContentFilterElementResult, err error) {
	m.ExtensionObjectDefinitionContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ContentFilterElementResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ContentFilterElementResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	statusCode, err := ReadSimpleField[StatusCode](ctx, "statusCode", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'statusCode' field"))
	}
	m.StatusCode = statusCode

	noOfOperandStatusCodes, err := ReadSimpleField(ctx, "noOfOperandStatusCodes", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfOperandStatusCodes' field"))
	}
	m.NoOfOperandStatusCodes = noOfOperandStatusCodes

	operandStatusCodes, err := ReadCountArrayField[StatusCode](ctx, "operandStatusCodes", ReadComplex[StatusCode](StatusCodeParseWithBuffer, readBuffer), uint64(noOfOperandStatusCodes))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operandStatusCodes' field"))
	}
	m.OperandStatusCodes = operandStatusCodes

	noOfOperandDiagnosticInfos, err := ReadSimpleField(ctx, "noOfOperandDiagnosticInfos", ReadSignedInt(readBuffer, uint8(32)))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'noOfOperandDiagnosticInfos' field"))
	}
	m.NoOfOperandDiagnosticInfos = noOfOperandDiagnosticInfos

	operandDiagnosticInfos, err := ReadCountArrayField[DiagnosticInfo](ctx, "operandDiagnosticInfos", ReadComplex[DiagnosticInfo](DiagnosticInfoParseWithBuffer, readBuffer), uint64(noOfOperandDiagnosticInfos))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'operandDiagnosticInfos' field"))
	}
	m.OperandDiagnosticInfos = operandDiagnosticInfos

	if closeErr := readBuffer.CloseContext("ContentFilterElementResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ContentFilterElementResult")
	}

	return m, nil
}

func (m *_ContentFilterElementResult) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ContentFilterElementResult) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ContentFilterElementResult"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ContentFilterElementResult")
		}

		if err := WriteSimpleField[StatusCode](ctx, "statusCode", m.GetStatusCode(), WriteComplex[StatusCode](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'statusCode' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfOperandStatusCodes", m.GetNoOfOperandStatusCodes(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfOperandStatusCodes' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "operandStatusCodes", m.GetOperandStatusCodes(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'operandStatusCodes' field")
		}

		if err := WriteSimpleField[int32](ctx, "noOfOperandDiagnosticInfos", m.GetNoOfOperandDiagnosticInfos(), WriteSignedInt(writeBuffer, 32)); err != nil {
			return errors.Wrap(err, "Error serializing 'noOfOperandDiagnosticInfos' field")
		}

		if err := WriteComplexTypeArrayField(ctx, "operandDiagnosticInfos", m.GetOperandDiagnosticInfos(), writeBuffer); err != nil {
			return errors.Wrap(err, "Error serializing 'operandDiagnosticInfos' field")
		}

		if popErr := writeBuffer.PopContext("ContentFilterElementResult"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ContentFilterElementResult")
		}
		return nil
	}
	return m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_ContentFilterElementResult) IsContentFilterElementResult() {}

func (m *_ContentFilterElementResult) DeepCopy() any {
	return m.deepCopy()
}

func (m *_ContentFilterElementResult) deepCopy() *_ContentFilterElementResult {
	if m == nil {
		return nil
	}
	_ContentFilterElementResultCopy := &_ContentFilterElementResult{
		m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition).deepCopy(),
		m.StatusCode.DeepCopy().(StatusCode),
		m.NoOfOperandStatusCodes,
		utils.DeepCopySlice[StatusCode, StatusCode](m.OperandStatusCodes),
		m.NoOfOperandDiagnosticInfos,
		utils.DeepCopySlice[DiagnosticInfo, DiagnosticInfo](m.OperandDiagnosticInfos),
	}
	m.ExtensionObjectDefinitionContract.(*_ExtensionObjectDefinition)._SubType = m
	return _ContentFilterElementResultCopy
}

func (m *_ContentFilterElementResult) String() string {
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
