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

// BACnetConstructedDataThreatLevel is the corresponding interface of BACnetConstructedDataThreatLevel
type BACnetConstructedDataThreatLevel interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetThreatLevel returns ThreatLevel (property field)
	GetThreatLevel() BACnetAccessThreatLevel
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetAccessThreatLevel
	// IsBACnetConstructedDataThreatLevel is a marker method to prevent unintentional type checks (interfaces of same signature)
	IsBACnetConstructedDataThreatLevel()
}

// _BACnetConstructedDataThreatLevel is the data-structure of this message
type _BACnetConstructedDataThreatLevel struct {
	BACnetConstructedDataContract
	ThreatLevel BACnetAccessThreatLevel
}

var _ BACnetConstructedDataThreatLevel = (*_BACnetConstructedDataThreatLevel)(nil)
var _ BACnetConstructedDataRequirements = (*_BACnetConstructedDataThreatLevel)(nil)

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataThreatLevel) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataThreatLevel) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_THREAT_LEVEL
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataThreatLevel) GetParent() BACnetConstructedDataContract {
	return m.BACnetConstructedDataContract
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataThreatLevel) GetThreatLevel() BACnetAccessThreatLevel {
	return m.ThreatLevel
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataThreatLevel) GetActualValue() BACnetAccessThreatLevel {
	ctx := context.Background()
	_ = ctx
	return CastBACnetAccessThreatLevel(m.GetThreatLevel())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataThreatLevel factory function for _BACnetConstructedDataThreatLevel
func NewBACnetConstructedDataThreatLevel(openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, threatLevel BACnetAccessThreatLevel, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataThreatLevel {
	if threatLevel == nil {
		panic("threatLevel of type BACnetAccessThreatLevel for BACnetConstructedDataThreatLevel must not be nil")
	}
	_result := &_BACnetConstructedDataThreatLevel{
		BACnetConstructedDataContract: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
		ThreatLevel:                   threatLevel,
	}
	_result.BACnetConstructedDataContract.(*_BACnetConstructedData)._SubType = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataThreatLevel(structType any) BACnetConstructedDataThreatLevel {
	if casted, ok := structType.(BACnetConstructedDataThreatLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataThreatLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataThreatLevel) GetTypeName() string {
	return "BACnetConstructedDataThreatLevel"
}

func (m *_BACnetConstructedDataThreatLevel) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.BACnetConstructedDataContract.(*_BACnetConstructedData).getLengthInBits(ctx))

	// Simple field (threatLevel)
	lengthInBits += m.ThreatLevel.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataThreatLevel) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func (m *_BACnetConstructedDataThreatLevel) parse(ctx context.Context, readBuffer utils.ReadBuffer, parent *_BACnetConstructedData, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (__bACnetConstructedDataThreatLevel BACnetConstructedDataThreatLevel, err error) {
	m.BACnetConstructedDataContract = parent
	parent._SubType = m
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataThreatLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataThreatLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	threatLevel, err := ReadSimpleField[BACnetAccessThreatLevel](ctx, "threatLevel", ReadComplex[BACnetAccessThreatLevel](BACnetAccessThreatLevelParseWithBuffer, readBuffer))
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'threatLevel' field"))
	}
	m.ThreatLevel = threatLevel

	actualValue, err := ReadVirtualField[BACnetAccessThreatLevel](ctx, "actualValue", (*BACnetAccessThreatLevel)(nil), threatLevel)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("Error parsing 'actualValue' field"))
	}
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataThreatLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataThreatLevel")
	}

	return m, nil
}

func (m *_BACnetConstructedDataThreatLevel) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataThreatLevel) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataThreatLevel"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataThreatLevel")
		}

		if err := WriteSimpleField[BACnetAccessThreatLevel](ctx, "threatLevel", m.GetThreatLevel(), WriteComplex[BACnetAccessThreatLevel](writeBuffer)); err != nil {
			return errors.Wrap(err, "Error serializing 'threatLevel' field")
		}
		// Virtual field
		actualValue := m.GetActualValue()
		_ = actualValue
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataThreatLevel"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataThreatLevel")
		}
		return nil
	}
	return m.BACnetConstructedDataContract.(*_BACnetConstructedData).serializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataThreatLevel) IsBACnetConstructedDataThreatLevel() {}

func (m *_BACnetConstructedDataThreatLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
