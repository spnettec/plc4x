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

// BACnetConstructedDataGroupMode is the corresponding interface of BACnetConstructedDataGroupMode
type BACnetConstructedDataGroupMode interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	utils.Copyable
	BACnetConstructedData
	// GetGroupMode returns GroupMode (property field)
	GetGroupMode() BACnetLiftGroupModeTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetLiftGroupModeTagged
	// IsBACnetConstructedDataGroupMode is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataGroupMode()
	// CreateBuilder creates a BACnetConstructedDataGroupModeBuilder
	CreateBACnetConstructedDataGroupModeBuilder() BACnetConstructedDataGroupModeBuilder
}

// _BACnetConstructedDataGroupMode is the data-structure of this message
type _BACnetConstructedDataGroupMode struct {
	BACnetConstructedDataContract
	GroupMode BACnetLiftGroupModeTagged
}

var _ BACnetConstructedDataGroupMode = (*_BACnetConstructedDataGroupMode)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataGroupMode)(nil)

// NewBACnetConstructedDataGroupMode factory function for _BACnetConstructedDataGroupMode
func NewBACnetConstructedDataGroupMode(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, groupMode BACnetLiftGroupModeTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataGroupMode {
	if groupMode == nil {
		panic("groupMode of type BACnetLiftGroupModeTagged for BACnetConstructedDataGroupMode must not be nil")
	}
	_result := &_BACnetConstructedDataGroupMode{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		GroupMode:                     groupMode,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Builder
///////////////////////

// BACnetConstructedDataGroupModeBuilder is a builder for BACnetConstructedDataGroupMode
type BACnetConstructedDataGroupModeBuilder interface {
	utils.Copyable
	// WithMandatoryFields adds all mandatory fields (convenience for using multiple builder calls)
	WithMandatoryFields(groupMode BACnetLiftGroupModeTagged) BACnetConstructedDataGroupModeBuilder
	// WithGroupMode adds GroupMode (property field)
	WithGroupMode(BACnetLiftGroupModeTagged) BACnetConstructedDataGroupModeBuilder
	// WithGroupModeBuilder adds GroupMode (property field) which is build by the builder
	WithGroupModeBuilder(func(BACnetLiftGroupModeTaggedBuilder) BACnetLiftGroupModeTaggedBuilder) BACnetConstructedDataGroupModeBuilder
	// Build builds the BACnetConstructedDataGroupMode or returns an error if something is wrong
	Build() (BACnetConstructedDataGroupMode, error)
	// MustBuild does the same as Build but panics on error
	MustBuild() BACnetConstructedDataGroupMode
}

// NewBACnetConstructedDataGroupModeBuilder() creates a BACnetConstructedDataGroupModeBuilder
func NewBACnetConstructedDataGroupModeBuilder() BACnetConstructedDataGroupModeBuilder {
	return &_BACnetConstructedDataGroupModeBuilder{_BACnetConstructedDataGroupMode: new(_BACnetConstructedDataGroupMode)}
}

type _BACnetConstructedDataGroupModeBuilder struct {
	*_BACnetConstructedDataGroupMode

	parentBuilder *_BACnetConstructedDataBuilder

	err *utils.MultiError
}

var _ (BACnetConstructedDataGroupModeBuilder) = (*_BACnetConstructedDataGroupModeBuilder)(nil)

func (b *_BACnetConstructedDataGroupModeBuilder) setParent(contract BACnetConstructedDataContract) {
	b.BACnetConstructedDataContract = contract
}

func (b *_BACnetConstructedDataGroupModeBuilder) WithMandatoryFields(groupMode BACnetLiftGroupModeTagged) BACnetConstructedDataGroupModeBuilder {
	return b.WithGroupMode(groupMode)
}

func (b *_BACnetConstructedDataGroupModeBuilder) WithGroupMode(groupMode BACnetLiftGroupModeTagged) BACnetConstructedDataGroupModeBuilder {
	b.GroupMode = groupMode
	return b
}

func (b *_BACnetConstructedDataGroupModeBuilder) WithGroupModeBuilder(builderSupplier func(BACnetLiftGroupModeTaggedBuilder) BACnetLiftGroupModeTaggedBuilder) BACnetConstructedDataGroupModeBuilder {
	builder := builderSupplier(b.GroupMode.CreateBACnetLiftGroupModeTaggedBuilder())
	var err error
	b.GroupMode, err = builder.Build()
	if err != nil {
		if b.err == nil {
			b.err = &utils.MultiError{MainError: errors.New("sub builder failed")}
		}
		b.err.Append(errors.Wrap(err, "BACnetLiftGroupModeTaggedBuilder failed"))
	}
	return b
}

func (b *_BACnetConstructedDataGroupModeBuilder) Build() (BACnetConstructedDataGroupMode, error) {
	if b.GroupMode == nil {
		if b.err == nil {
			b.err = new(utils.MultiError)
		}
		b.err.Append(errors.New("mandatory field 'groupMode' not set"))
	}
	if b.err != nil {
		return nil, errors.Wrap(b.err, "error occurred during build")
	}
	return b._BACnetConstructedDataGroupMode.deepCopy(), nil
}

func (b *_BACnetConstructedDataGroupModeBuilder) MustBuild() BACnetConstructedDataGroupMode {
	build, err := b.Build()
	if err != nil {
		panic(err)
	}
	return build
}

// Done is used to finish work on this child and return to the parent builder
func (b *_BACnetConstructedDataGroupModeBuilder) Done() BACnetConstructedDataBuilder {
	return b.parentBuilder
}

func (b *_BACnetConstructedDataGroupModeBuilder) buildForBACnetConstructedData() (BACnetConstructedData, error) {
	return b.Build()
}

func (b *_BACnetConstructedDataGroupModeBuilder) DeepCopy() any {
	_copy := b.CreateBACnetConstructedDataGroupModeBuilder().(*_BACnetConstructedDataGroupModeBuilder)
	if b.err != nil {
		_copy.err = b.err.DeepCopy().(*utils.MultiError)
	}
	return _copy
}

// CreateBACnetConstructedDataGroupModeBuilder creates a BACnetConstructedDataGroupModeBuilder
func (b *_BACnetConstructedDataGroupMode) CreateBACnetConstructedDataGroupModeBuilder() BACnetConstructedDataGroupModeBuilder {
	if b == nil {
		return NewBACnetConstructedDataGroupModeBuilder()
	}
	return &_BACnetConstructedDataGroupModeBuilder{_BACnetConstructedDataGroupMode: b.deepCopy()}
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataGroupMode) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataGroupMode) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_GROUP_MODE
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataGroupMode) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataGroupMode) GetGroupMode() BACnetLiftGroupModeTagged {
	return m.GroupMode
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataGroupMode) GetActualValue() BACnetLiftGroupModeTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetLiftGroupModeTagged(m.GetGroupMode())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataGroupMode(structType any) BACnetConstructedDataGroupMode {
	if casted, ok := structType.(BACnetConstructedDataGroupMode); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataGroupMode); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataGroupMode) GetTypeName() string {
	return "BACnetConstructedDataGroupMode"
}

func (m *_BACnetConstructedDataGroupMode) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (groupMode)
	lengthInBits += m.GroupMode.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataGroupMode) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataGroupMode) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataGroupMode BACnetConstructedDataGroupMode, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataGroupMode"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataGroupMode")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	groupMode, err := ReadSimpleField[BACnetLiftGroupModeTagged](ctx, "groupMode", ReadComplex[BACnetLiftGroupModeTagged](BACnetLiftGroupModeTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'groupMode' field"))
	}
	m.GroupMode = groupMode

	actualValue, err := ReadVirtualField[BACnetLiftGroupModeTagged](ctx, "actualValue", (*BACnetLiftGroupModeTagged)(nil), groupMode)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataGroupMode"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataGroupMode")
	}

	return m, nil
}

func (m *_BACnetConstructedDataGroupMode) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataGroupMode) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataGroupMode"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataGroupMode")
		}

		if err := WriteSimpleField[BACnetLiftGroupModeTagged](ctx, "groupMode", m.GetGroupMode(), WriteComplex[BACnetLiftGroupModeTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'groupMode' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataGroupMode"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataGroupMode")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataGroupMode) IsBACnetConstructedDataGroupMode() {}

func (m *_BACnetConstructedDataGroupMode) DeepCopy() any {
	return m.deepCopy()
}

func (m *_BACnetConstructedDataGroupMode) deepCopy() *_BACnetConstructedDataGroupMode {
	if m == nil {
		return nil
	}
	_BACnetConstructedDataGroupModeCopy := &_BACnetConstructedDataGroupMode{
		m.BACnetConstructedDataContract.(*_BACnetConstructedData).deepCopy(),
		m.GroupMode.DeepCopy().(BACnetLiftGroupModeTagged),
	}
	m.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = m
	return _BACnetConstructedDataGroupModeCopy
}

func (m *_BACnetConstructedDataGroupMode) String() string {
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
