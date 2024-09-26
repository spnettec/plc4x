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

// BACnetConstructedDataMaintenanceRequired is the corresponding interface of BACnetConstructedDataMaintenanceRequired
type BACnetConstructedDataMaintenanceRequired interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetMaintenanceRequired returns MaintenanceRequired (property field)
	GetMaintenanceRequired() BACnetMaintenanceTagged
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetMaintenanceTagged
	// IsBACnetConstructedDataMaintenanceRequired is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataMaintenanceRequired()
}

// _BACnetConstructedDataMaintenanceRequired is the data-structure of this message
type _BACnetConstructedDataMaintenanceRequired struct {
	BACnetConstructedDataContract
	MaintenanceRequired BACnetMaintenanceTagged
}

var _ BACnetConstructedDataMaintenanceRequired = (*_BACnetConstructedDataMaintenanceRequired)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataMaintenanceRequired)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataMaintenanceRequired) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataMaintenanceRequired) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_MAINTENANCE_REQUIRED
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataMaintenanceRequired) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataMaintenanceRequired) GetMaintenanceRequired() BACnetMaintenanceTagged {
	return m.MaintenanceRequired
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataMaintenanceRequired) GetActualValue() BACnetMaintenanceTagged {
	ctx := context.Background()
	_ = ctx
	return CastBACnetMaintenanceTagged(m.GetMaintenanceRequired())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataMaintenanceRequired factory function for _BACnetConstructedDataMaintenanceRequired
func NewBACnetConstructedDataMaintenanceRequired(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, maintenanceRequired BACnetMaintenanceTagged, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataMaintenanceRequired {
	if maintenanceRequired == nil {
		panic("maintenanceRequired of type BACnetMaintenanceTagged for BACnetConstructedDataMaintenanceRequired must not be nil")
	}
	_result := &_BACnetConstructedDataMaintenanceRequired{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		MaintenanceRequired:           maintenanceRequired,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataMaintenanceRequired(structType any) BACnetConstructedDataMaintenanceRequired {
	if casted, ok := structType.(BACnetConstructedDataMaintenanceRequired); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataMaintenanceRequired); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataMaintenanceRequired) GetTypeName() string {
	return "BACnetConstructedDataMaintenanceRequired"
}

func (m *_BACnetConstructedDataMaintenanceRequired) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (maintenanceRequired)
	lengthInBits += m.MaintenanceRequired.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataMaintenanceRequired) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataMaintenanceRequired) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataMaintenanceRequired BACnetConstructedDataMaintenanceRequired, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataMaintenanceRequired"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataMaintenanceRequired")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	maintenanceRequired, err := ReadSimpleField[BACnetMaintenanceTagged](ctx, "maintenanceRequired", ReadComplex[BACnetMaintenanceTagged](BACnetMaintenanceTaggedParseWithBufferProducer((uint8)(uint8(0)), (TagClass)(TagClass_APPLICATION_TAGS)), readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'maintenanceRequired' field"))
	}
	m.MaintenanceRequired = maintenanceRequired

	actualValue, err := ReadVirtualField[BACnetMaintenanceTagged](ctx, "actualValue", (*BACnetMaintenanceTagged)(nil), maintenanceRequired)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataMaintenanceRequired"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataMaintenanceRequired")
	}

	return m, nil
}

func (m *_BACnetConstructedDataMaintenanceRequired) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataMaintenanceRequired) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataMaintenanceRequired"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataMaintenanceRequired")
		}

		if err := WriteSimpleField[BACnetMaintenanceTagged](ctx, "maintenanceRequired", m.GetMaintenanceRequired(), WriteComplex[BACnetMaintenanceTagged](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'maintenanceRequired' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataMaintenanceRequired"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataMaintenanceRequired")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataMaintenanceRequired) IsBACnetConstructedDataMaintenanceRequired() {}

func (m *_BACnetConstructedDataMaintenanceRequired) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
