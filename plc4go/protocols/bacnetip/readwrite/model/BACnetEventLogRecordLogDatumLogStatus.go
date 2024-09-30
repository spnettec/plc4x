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

// BACnetEventLogRecordLogDatumLogStatus is the corresponding interface of BACnetEventLogRecordLogDatumLogStatus
type BACnetEventLogRecordLogDatumLogStatus interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetEventLogRecordLogDatum
	// GetLogStatus returns LogStatus (property field)
	GetLogStatus() BACnetLogStatusTagged
	// IsBACnetEventLogRecordLogDatumLogStatus is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetEventLogRecordLogDatumLogStatus()
	// CreateBuilder creates a BACnetEventLogRecordLogDatumLogStatusBuilder
	CreateBACnetEventLogRecordLogDatumLogStatusBuilder() BACnetEventLogRecordLogDatumLogStatusBuilder
}

// _BACnetEventLogRecordLogDatumLogStatus is the data-structure of this message
type _BACnetEventLogRecordLogDatumLogStatus struct {
	BACnetEventLogRecordLogDatumContract
	LogStatus BACnetLogStatusTagged
}

var _ BACnetEventLogRecordLogDatumLogStatus = (*_BACnetEventLogRecordLogDatumLogStatus)(nil)
var _ BACnetEventLogRecordLogDatumRequirements = (*_BACnetEventLogRecordLogDatumLogStatus)(nil)

// NewBACnetEventLogRecordLogDatumLogStatus factory function for _BACnetEventLogRecordLogDatumLogStatus
func NewBACnetEventLogRecordLogDatumLogStatus(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, logStatus BACnetLogStatusTagged, tagNumber uint8) *_BACnetEventLogRecordLogDatumLogStatus {
	if logStatus == nil {
		panic("logStatus of type BACnetLogStatusTagged for BACnetEventLogRecordLogDatumLogStatus must not be nil")
	}
	_result := &_BACnetEventLogRecordLogDatumLogStatus{
		BACnetEventLogRecordLogDatumContract: NewBACnetEventLogRecordLogDatum(openingTag, peekedTagHeader, closingTag, tagNumber),
		LogStatus:                            logStatus,
	}
	_result.BACnetEventLogRecordLogDatumContract.(*_BACnetEventLogRecordLogDatum)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetEventLogRecordLogDatumLogStatusBuilder is a builder for BACnetEventLogRecordLogDatumLogStatus
type BACnetEventLogRecordLogDatumLogStatusBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(logStatus BACnetLogStatusTagged) BACnetEventLogRecordLogDatumLogStatusBuilder
	// WithLogStatus adds LogStatus (property field)
	WithLogStatus(BACnetLogStatusTagged) BACnetEventLogRecordLogDatumLogStatusBuilder
	// WithLogStatusBuilder adds LogStatus (property field) which is build by the builder
	WithLogStatusBuilder(func(BACnetLogStatusTaggedBuilder) BACnetLogStatusTaggedBuilder) BACnetEventLogRecordLogDatumLogStatusBuilder
	// Build builds the BACnetEventLogRecordLogDatumLogStatus or returns an error if something is wrong
	Build() (BACnetEventLogRecordLogDatumLogStatus, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetEventLogRecordLogDatumLogStatus
}

// NewBACnetEventLogRecordLogDatumLogStatusBuilder() creates a BACnetEventLogRecordLogDatumLogStatusBuilder
func NewBACnetEventLogRecordLogDatumLogStatusBuilder() BACnetEventLogRecordLogDatumLogStatusBuilder {
	return &_BACnetEventLogRecordLogDatumLogStatusBuilder{_BACnetEventLogRecordLogDatumLogStatus: new(_BACnetEventLogRecordLogDatumLogStatus)}
}

type _BACnetEventLogRecordLogDatumLogStatusBuilder struct {
	*_BACnetEventLogRecordLogDatumLogStatus

	parentBuilder *_BACnetEventLogRecordLogDatumBuilder

	err *utils.MultiError
}

var _ (BACnetEventLogRecordLogDatumLogStatusBuilder) = (*_BACnetEventLogRecordLogDatumLogStatusBuilder)(nil)

func (b *_BACnetEventLogRecordLogDatumLogStatusBuilder) setParent(contract BACnetEventLogRecordLogDatumContract) {
	b.BACnetEventLogRecordLogDatumContract = contract
}

func (b *_BACnetEventLogRecordLogDatumLogStatusBuilder) WithMandatoryFields(logStatus BACnetLogStatusTagged) BACnetEventLogRecordLogDatumLogStatusBuilder {
	return b.WithLogStatus(logStatus)
}

func (b *_BACnetEventLogRecordLogDatumLogStatusBuilder) WithLogStatus(logStatus BACnetLogStatusTagged) BACnetEventLogRecordLogDatumLogStatusBuilder {
	b.LogStatus = logStatus
	return b
}

func (b *_BACnetEventLogRecordLogDatumLogStatusBuilder) WithLogStatusBuilder(builderSupplier func(BACnetLogStatusTaggedBuilder) BACnetLogStatusTaggedBuilder) BACnetEventLogRecordLogDatumLogStatusBuilder {
	builder := builderSupplier(b.LogStatus.CreateBACnetLogStatusTaggedBuilder())
	var err error
	b.LogStatus, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLogStatusTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetEventLogRecordLogDatumLogStatusBuilder) Build() (BACnetEventLogRecordLogDatumLogStatus, error) {
	if b.LogStatus == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'logStatus' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetEventLogRecordLogDatumLogStatus.deepCopy(), nil
}

func (b *_BACnetEventLogRecordLogDatumLogStatusBuilder) MustBuild() BACnetEventLogRecordLogDatumLogStatus {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetEventLogRecordLogDatumLogStatusBuilder) Done() BACnetEventLogRecordLogDatumBuilder {
	return b.parentBuilder
}

func (b *_BACnetEventLogRecordLogDatumLogStatusBuilder) buildForBACnetEventLogRecordLogDatum() (BACnetEventLogRecordLogDatum, error) {
	return b.Build()
}

func (b *_BACnetEventLogRecordLogDatumLogStatusBuilder) DeepCopy() any {
	_copy := b.CreateBACnetEventLogRecordLogDatumLogStatusBuilder().(*_BACnetEventLogRecordLogDatumLogStatusBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetEventLogRecordLogDatumLogStatusBuilder creates a BACnetEventLogRecordLogDatumLogStatusBuilder
func (b *_BACnetEventLogRecordLogDatumLogStatus) CreateBACnetEventLogRecordLogDatumLogStatusBuilder() BACnetEventLogRecordLogDatumLogStatusBuilder {
	if b == nil {
		return NewBACnetEventLogRecordLogDatumLogStatusBuilder()
	}
	return &_BACnetEventLogRecordLogDatumLogStatusBuilder{_BACnetEventLogRecordLogDatumLogStatus: b.deepCopy()}
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

func (m *_BACnetEventLogRecordLogDatumLogStatus) GetParent() BACnetEventLogRecordLogDatumContract {
	return m.BACnetEventLogRecordLogDatumContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventLogRecordLogDatumLogStatus) GetLogStatus() BACnetLogStatusTagged {
	return m.LogStatus
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetEventLogRecordLogDatumLogStatus(structType any) BACnetEventLogRecordLogDatumLogStatus {
	if casted, ok := structType.(BACnetEventLogRecordLogDatumLogStatus); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventLogRecordLogDatumLogStatus); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) GetTypeName() string {
	return "BACnetEventLogRecordLogDatumLogStatus"
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetEventLogRecordLogDatumContract.(*_BACnetEventLogRecordLogDatum).getLengthInBits(ctx))

	// Simple field (logStatus)
	lengthInBits += m.LogStatus.GetLengthInBits(ctx)

	return lengthInBits
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetEventLogRecordLogDatum, tagNumber uint8) (__bACnetEventLogRecordLogDatumLogStatus BACnetEventLogRecordLogDatumLogStatus, err error) {
	m.BACnetEventLogRecordLogDatumContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventLogRecordLogDatumLogStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventLogRecordLogDatumLogStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	logStatus, err := ReadSimpleField[BACnetLogStatusTagged](ctx, "logStatus", ReadComplex[BACnetLogStatusTagged](BACnetLogStatusTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_CONTEXT_SPECIFIC_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'logStatus' field"))
	}
	m.LogStatus = logStatus

	if closeErr := readBuffer.CloseContext("BACnetEventLogRecordLogDatumLogStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventLogRecordLogDatumLogStatus")
	}

	return m, nil
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetEventLogRecordLogDatumLogStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetEventLogRecordLogDatumLogStatus")
		}

		if err := WriteSimpleField[BACnetLogStatusTagged](ctx, "logStatus", m.GetLogStatus(), WriteComplex[BACnetLogStatusTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'logStatus' field")
		}

		if popErr := writeBuffer.PopContext("BACnetEventLogRecordLogDatumLogStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetEventLogRecordLogDatumLogStatus")
		}
		return nil
	}
	return m.BACnetEventLogRecordLogDatumContract.(*_BACnetEventLogRecordLogDatum).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) IsBACnetEventLogRecordLogDatumLogStatus() {}

func (m *_BACnetEventLogRecordLogDatumLogStatus) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) deepCopy() *_BACnetEventLogRecordLogDatumLogStatus {
	if m == nil {
		return nil
	}
	_BACnetEventLogRecordLogDatumLogStatusCopy := &_BACnetEventLogRecordLogDatumLogStatus{
		m.BACnetEventLogRecordLogDatumContract.(*_BACnetEventLogRecordLogDatum).deepCopy(),
		m.LogStatus.DeepCopy().(BACnetLogStatusTagged),
	}
	m.BACnetEventLogRecordLogDatumContract.(*_BACnetEventLogRecordLogDatum)._SubType = m
	return _BACnetEventLogRecordLogDatumLogStatusCopy
}

func (m *_BACnetEventLogRecordLogDatumLogStatus) String() string {
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
