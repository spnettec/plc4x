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

// BACnetConstructedDataAllowGroupDelayInhibit is the corresponding interface of BACnetConstructedDataAllowGroupDelayInhibit
type BACnetConstructedDataAllowGroupDelayInhibit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetAllowGroupDelayInhibit returns AllowGroupDelayInhibit (property field)
	GetAllowGroupDelayInhibit() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
	// IsBACnetConstructedDataAllowGroupDelayInhibit is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataAllowGroupDelayInhibit()
}

// _BACnetConstructedDataAllowGroupDelayInhibit is the data-structure of this message
type _BACnetConstructedDataAllowGroupDelayInhibit struct {
	BACnetConstructedDataContract
	AllowGroupDelayInhibit BACnetApplicationTagBoolean
}

var _ BACnetConstructedDataAllowGroupDelayInhibit = (*_BACnetConstructedDataAllowGroupDelayInhibit)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataAllowGroupDelayInhibit)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAllowGroupDelayInhibit) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataAllowGroupDelayInhibit) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_ALLOW_GROUP_DELAY_INHIBIT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAllowGroupDelayInhibit) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAllowGroupDelayInhibit) GetAllowGroupDelayInhibit() BACnetApplicationTagBoolean {
	return m.AllowGroupDelayInhibit
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAllowGroupDelayInhibit) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetAllowGroupDelayInhibit())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataAllowGroupDelayInhibit factory function for _BACnetConstructedDataAllowGroupDelayInhibit
func NewBACnetConstructedDataAllowGroupDelayInhibit(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, allowGroupDelayInhibit BACnetApplicationTagBoolean, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataAllowGroupDelayInhibit {
	if allowGroupDelayInhibit == nil {
		panic("allowGroupDelayInhibit of type BACnetApplicationTagBoolean for BACnetConstructedDataAllowGroupDelayInhibit must not be nil")
	}
	_result := &_BACnetConstructedDataAllowGroupDelayInhibit{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		AllowGroupDelayInhibit:        allowGroupDelayInhibit,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAllowGroupDelayInhibit(structType any) BACnetConstructedDataAllowGroupDelayInhibit {
	if casted, ok := structType.(BACnetConstructedDataAllowGroupDelayInhibit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAllowGroupDelayInhibit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAllowGroupDelayInhibit) GetTypeName() string {
	return "BACnetConstructedDataAllowGroupDelayInhibit"
}

func (m *_BACnetConstructedDataAllowGroupDelayInhibit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (allowGroupDelayInhibit)
	lengthInBits += m.AllowGroupDelayInhibit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataAllowGroupDelayInhibit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataAllowGroupDelayInhibit) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataAllowGroupDelayInhibit BACnetConstructedDataAllowGroupDelayInhibit, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAllowGroupDelayInhibit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAllowGroupDelayInhibit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	allowGroupDelayInhibit, err := ReadSimpleField[BACnetApplicationTagBoolean](ctx, "allowGroupDelayInhibit", ReadComplex[BACnetApplicationTagBoolean](BACnetApplicationTagParseWithBufferProducer[BACnetApplicationTagBoolean](), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'allowGroupDelayInhibit' field"))
	}
	m.AllowGroupDelayInhibit = allowGroupDelayInhibit

	actualValue, err := ReadVirtualField[BACnetApplicationTagBoolean](ctx, "actualValue", (*BACnetApplicationTagBoolean)(nil), allowGroupDelayInhibit)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAllowGroupDelayInhibit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAllowGroupDelayInhibit")
	}

	return m, nil
}

func (m *_BACnetConstructedDataAllowGroupDelayInhibit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAllowGroupDelayInhibit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAllowGroupDelayInhibit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAllowGroupDelayInhibit")
		}

		if err := WriteSimpleField[BACnetApplicationTagBoolean](ctx, "allowGroupDelayInhibit", m.GetAllowGroupDelayInhibit(), WriteComplex[BACnetApplicationTagBoolean](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'allowGroupDelayInhibit' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAllowGroupDelayInhibit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAllowGroupDelayInhibit")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataAllowGroupDelayInhibit) IsBACnetConstructedDataAllowGroupDelayInhibit() {
}

func (m *_BACnetConstructedDataAllowGroupDelayInhibit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
