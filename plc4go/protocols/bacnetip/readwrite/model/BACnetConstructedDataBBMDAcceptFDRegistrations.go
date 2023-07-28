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


// BACnetConstructedDataBBMDAcceptFDRegistrations is the corresponding interface of BACnetConstructedDataBBMDAcceptFDRegistrations
type BACnetConstructedDataBBMDAcceptFDRegistrations interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetBbmdAcceptFDRegistrations returns BbmdAcceptFDRegistrations (property field)
	GetBbmdAcceptFDRegistrations() BACnetApplicationTagBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagBoolean
}

// BACnetConstructedDataBBMDAcceptFDRegistrationsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataBBMDAcceptFDRegistrations.
// This is useful for switch cases.
type BACnetConstructedDataBBMDAcceptFDRegistrationsExactly interface {
	BACnetConstructedDataBBMDAcceptFDRegistrations
	isBACnetConstructedDataBBMDAcceptFDRegistrations() bool
}

// _BACnetConstructedDataBBMDAcceptFDRegistrations is the data-structure of this message
type _BACnetConstructedDataBBMDAcceptFDRegistrations struct {
	*_BACnetConstructedData
        BbmdAcceptFDRegistrations BACnetApplicationTagBoolean
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_BBMD_ACCEPT_FD_REGISTRATIONS}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) GetBbmdAcceptFDRegistrations() BACnetApplicationTagBoolean {
	return m.BbmdAcceptFDRegistrations
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) GetActualValue() BACnetApplicationTagBoolean {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagBoolean(m.GetBbmdAcceptFDRegistrations())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataBBMDAcceptFDRegistrations factory function for _BACnetConstructedDataBBMDAcceptFDRegistrations
func NewBACnetConstructedDataBBMDAcceptFDRegistrations( bbmdAcceptFDRegistrations BACnetApplicationTagBoolean , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataBBMDAcceptFDRegistrations {
	_result := &_BACnetConstructedDataBBMDAcceptFDRegistrations{
		BbmdAcceptFDRegistrations: bbmdAcceptFDRegistrations,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataBBMDAcceptFDRegistrations(structType any) BACnetConstructedDataBBMDAcceptFDRegistrations {
    if casted, ok := structType.(BACnetConstructedDataBBMDAcceptFDRegistrations); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataBBMDAcceptFDRegistrations); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) GetTypeName() string {
	return "BACnetConstructedDataBBMDAcceptFDRegistrations"
}

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (bbmdAcceptFDRegistrations)
	lengthInBits += m.BbmdAcceptFDRegistrations.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataBBMDAcceptFDRegistrationsParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBBMDAcceptFDRegistrations, error) {
	return BACnetConstructedDataBBMDAcceptFDRegistrationsParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataBBMDAcceptFDRegistrationsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataBBMDAcceptFDRegistrations, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataBBMDAcceptFDRegistrations"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataBBMDAcceptFDRegistrations")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (bbmdAcceptFDRegistrations)
	if pullErr := readBuffer.PullContext("bbmdAcceptFDRegistrations"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for bbmdAcceptFDRegistrations")
	}
_bbmdAcceptFDRegistrations, _bbmdAcceptFDRegistrationsErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _bbmdAcceptFDRegistrationsErr != nil {
		return nil, errors.Wrap(_bbmdAcceptFDRegistrationsErr, "Error parsing 'bbmdAcceptFDRegistrations' field of BACnetConstructedDataBBMDAcceptFDRegistrations")
	}
	bbmdAcceptFDRegistrations := _bbmdAcceptFDRegistrations.(BACnetApplicationTagBoolean)
	if closeErr := readBuffer.CloseContext("bbmdAcceptFDRegistrations"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for bbmdAcceptFDRegistrations")
	}

	// Virtual field
	_actualValue := bbmdAcceptFDRegistrations
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataBBMDAcceptFDRegistrations"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataBBMDAcceptFDRegistrations")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataBBMDAcceptFDRegistrations{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		BbmdAcceptFDRegistrations: bbmdAcceptFDRegistrations,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataBBMDAcceptFDRegistrations"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataBBMDAcceptFDRegistrations")
		}

	// Simple Field (bbmdAcceptFDRegistrations)
	if pushErr := writeBuffer.PushContext("bbmdAcceptFDRegistrations"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for bbmdAcceptFDRegistrations")
	}
	_bbmdAcceptFDRegistrationsErr := writeBuffer.WriteSerializable(ctx, m.GetBbmdAcceptFDRegistrations())
	if popErr := writeBuffer.PopContext("bbmdAcceptFDRegistrations"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for bbmdAcceptFDRegistrations")
	}
	if _bbmdAcceptFDRegistrationsErr != nil {
		return errors.Wrap(_bbmdAcceptFDRegistrationsErr, "Error serializing 'bbmdAcceptFDRegistrations' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataBBMDAcceptFDRegistrations"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataBBMDAcceptFDRegistrations")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) isBACnetConstructedDataBBMDAcceptFDRegistrations() bool {
	return true
}

func (m *_BACnetConstructedDataBBMDAcceptFDRegistrations) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



