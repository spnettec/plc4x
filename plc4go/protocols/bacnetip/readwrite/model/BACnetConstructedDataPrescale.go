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


// BACnetConstructedDataPrescale is the corresponding interface of BACnetConstructedDataPrescale
type BACnetConstructedDataPrescale interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetPrescale returns Prescale (property field)
	GetPrescale() BACnetPrescale
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetPrescale
}

// BACnetConstructedDataPrescaleExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataPrescale.
// This is useful for switch cases.
type BACnetConstructedDataPrescaleExactly interface {
	BACnetConstructedDataPrescale
	isBACnetConstructedDataPrescale() bool
}

// _BACnetConstructedDataPrescale is the data-structure of this message
type _BACnetConstructedDataPrescale struct {
	*_BACnetConstructedData
        Prescale BACnetPrescale
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPrescale)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataPrescale)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_PRESCALE}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPrescale) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataPrescale)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPrescale) GetPrescale() BACnetPrescale {
	return m.Prescale
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPrescale) GetActualValue() BACnetPrescale {
	ctx := context.Background()
	_ = ctx
	return CastBACnetPrescale(m.GetPrescale())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataPrescale factory function for _BACnetConstructedDataPrescale
func NewBACnetConstructedDataPrescale( prescale BACnetPrescale , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataPrescale {
	_result := &_BACnetConstructedDataPrescale{
		Prescale: prescale,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPrescale(structType any) BACnetConstructedDataPrescale {
    if casted, ok := structType.(BACnetConstructedDataPrescale); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPrescale); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPrescale) GetTypeName() string {
	return "BACnetConstructedDataPrescale"
}

func (m *_BACnetConstructedDataPrescale) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (prescale)
	lengthInBits += m.Prescale.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataPrescale) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataPrescaleParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPrescale, error) {
	return BACnetConstructedDataPrescaleParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataPrescaleParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPrescale, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPrescale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPrescale")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (prescale)
	if pullErr := readBuffer.PullContext("prescale"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for prescale")
	}
_prescale, _prescaleErr := BACnetPrescaleParseWithBuffer(ctx, readBuffer)
	if _prescaleErr != nil {
		return nil, errors.Wrap(_prescaleErr, "Error parsing 'prescale' field of BACnetConstructedDataPrescale")
	}
	prescale := _prescale.(BACnetPrescale)
	if closeErr := readBuffer.CloseContext("prescale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for prescale")
	}

	// Virtual field
	_actualValue := prescale
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPrescale"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPrescale")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataPrescale{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		Prescale: prescale,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataPrescale) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPrescale) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPrescale"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPrescale")
		}

	// Simple Field (prescale)
	if pushErr := writeBuffer.PushContext("prescale"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for prescale")
	}
	_prescaleErr := writeBuffer.WriteSerializable(ctx, m.GetPrescale())
	if popErr := writeBuffer.PopContext("prescale"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for prescale")
	}
	if _prescaleErr != nil {
		return errors.Wrap(_prescaleErr, "Error serializing 'prescale' field")
	}
	// Virtual field
	actualValue := m.GetActualValue()
	_ =	actualValue
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPrescale"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPrescale")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataPrescale) isBACnetConstructedDataPrescale() bool {
	return true
}

func (m *_BACnetConstructedDataPrescale) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



