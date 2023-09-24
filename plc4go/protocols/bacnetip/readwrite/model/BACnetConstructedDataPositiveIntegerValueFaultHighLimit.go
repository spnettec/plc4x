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


// BACnetConstructedDataPositiveIntegerValueFaultHighLimit is the corresponding interface of BACnetConstructedDataPositiveIntegerValueFaultHighLimit
type BACnetConstructedDataPositiveIntegerValueFaultHighLimit interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFaultHighLimit returns FaultHighLimit (property field)
	GetFaultHighLimit() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataPositiveIntegerValueFaultHighLimitExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataPositiveIntegerValueFaultHighLimit.
// This is useful for switch cases.
type BACnetConstructedDataPositiveIntegerValueFaultHighLimitExactly interface {
	BACnetConstructedDataPositiveIntegerValueFaultHighLimit
	isBACnetConstructedDataPositiveIntegerValueFaultHighLimit() bool
}

// _BACnetConstructedDataPositiveIntegerValueFaultHighLimit is the data-structure of this message
type _BACnetConstructedDataPositiveIntegerValueFaultHighLimit struct {
	*_BACnetConstructedData
        FaultHighLimit BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueFaultHighLimit)  GetObjectTypeArgument() BACnetObjectType {
return BACnetObjectType_POSITIVE_INTEGER_VALUE}

func (m *_BACnetConstructedDataPositiveIntegerValueFaultHighLimit)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_FAULT_HIGH_LIMIT}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueFaultHighLimit) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataPositiveIntegerValueFaultHighLimit)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueFaultHighLimit) GetFaultHighLimit() BACnetApplicationTagUnsignedInteger {
	return m.FaultHighLimit
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataPositiveIntegerValueFaultHighLimit) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetFaultHighLimit())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataPositiveIntegerValueFaultHighLimit factory function for _BACnetConstructedDataPositiveIntegerValueFaultHighLimit
func NewBACnetConstructedDataPositiveIntegerValueFaultHighLimit( faultHighLimit BACnetApplicationTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataPositiveIntegerValueFaultHighLimit {
	_result := &_BACnetConstructedDataPositiveIntegerValueFaultHighLimit{
		FaultHighLimit: faultHighLimit,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataPositiveIntegerValueFaultHighLimit(structType any) BACnetConstructedDataPositiveIntegerValueFaultHighLimit {
    if casted, ok := structType.(BACnetConstructedDataPositiveIntegerValueFaultHighLimit); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataPositiveIntegerValueFaultHighLimit); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueFaultHighLimit) GetTypeName() string {
	return "BACnetConstructedDataPositiveIntegerValueFaultHighLimit"
}

func (m *_BACnetConstructedDataPositiveIntegerValueFaultHighLimit) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (faultHighLimit)
	lengthInBits += m.FaultHighLimit.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataPositiveIntegerValueFaultHighLimit) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataPositiveIntegerValueFaultHighLimitParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPositiveIntegerValueFaultHighLimit, error) {
	return BACnetConstructedDataPositiveIntegerValueFaultHighLimitParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataPositiveIntegerValueFaultHighLimitParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataPositiveIntegerValueFaultHighLimit, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataPositiveIntegerValueFaultHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataPositiveIntegerValueFaultHighLimit")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (faultHighLimit)
	if pullErr := readBuffer.PullContext("faultHighLimit"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for faultHighLimit")
	}
_faultHighLimit, _faultHighLimitErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _faultHighLimitErr != nil {
		return nil, errors.Wrap(_faultHighLimitErr, "Error parsing 'faultHighLimit' field of BACnetConstructedDataPositiveIntegerValueFaultHighLimit")
	}
	faultHighLimit := _faultHighLimit.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("faultHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for faultHighLimit")
	}

	// Virtual field
	_actualValue := faultHighLimit
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataPositiveIntegerValueFaultHighLimit"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataPositiveIntegerValueFaultHighLimit")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataPositiveIntegerValueFaultHighLimit{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FaultHighLimit: faultHighLimit,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueFaultHighLimit) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataPositiveIntegerValueFaultHighLimit) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataPositiveIntegerValueFaultHighLimit"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataPositiveIntegerValueFaultHighLimit")
		}

	// Simple Field (faultHighLimit)
	if pushErr := writeBuffer.PushContext("faultHighLimit"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for faultHighLimit")
	}
	_faultHighLimitErr := writeBuffer.WriteSerializable(ctx, m.GetFaultHighLimit())
	if popErr := writeBuffer.PopContext("faultHighLimit"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for faultHighLimit")
	}
	if _faultHighLimitErr != nil {
		return errors.Wrap(_faultHighLimitErr, "Error serializing 'faultHighLimit' field")
	}
	// Virtual field
	actualValue := m.GetActualValue()
	_ =	actualValue
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataPositiveIntegerValueFaultHighLimit"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataPositiveIntegerValueFaultHighLimit")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataPositiveIntegerValueFaultHighLimit) isBACnetConstructedDataPositiveIntegerValueFaultHighLimit() bool {
	return true
}

func (m *_BACnetConstructedDataPositiveIntegerValueFaultHighLimit) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



