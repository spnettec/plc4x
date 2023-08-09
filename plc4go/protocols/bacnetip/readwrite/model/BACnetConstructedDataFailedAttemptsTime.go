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


// BACnetConstructedDataFailedAttemptsTime is the corresponding interface of BACnetConstructedDataFailedAttemptsTime
type BACnetConstructedDataFailedAttemptsTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetFailedAttemptsTime returns FailedAttemptsTime (property field)
	GetFailedAttemptsTime() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataFailedAttemptsTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataFailedAttemptsTime.
// This is useful for switch cases.
type BACnetConstructedDataFailedAttemptsTimeExactly interface {
	BACnetConstructedDataFailedAttemptsTime
	isBACnetConstructedDataFailedAttemptsTime() bool
}

// _BACnetConstructedDataFailedAttemptsTime is the data-structure of this message
type _BACnetConstructedDataFailedAttemptsTime struct {
	*_BACnetConstructedData
        FailedAttemptsTime BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataFailedAttemptsTime)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataFailedAttemptsTime)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_FAILED_ATTEMPTS_TIME}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataFailedAttemptsTime) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataFailedAttemptsTime)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataFailedAttemptsTime) GetFailedAttemptsTime() BACnetApplicationTagUnsignedInteger {
	return m.FailedAttemptsTime
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataFailedAttemptsTime) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetFailedAttemptsTime())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataFailedAttemptsTime factory function for _BACnetConstructedDataFailedAttemptsTime
func NewBACnetConstructedDataFailedAttemptsTime( failedAttemptsTime BACnetApplicationTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataFailedAttemptsTime {
	_result := &_BACnetConstructedDataFailedAttemptsTime{
		FailedAttemptsTime: failedAttemptsTime,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataFailedAttemptsTime(structType any) BACnetConstructedDataFailedAttemptsTime {
    if casted, ok := structType.(BACnetConstructedDataFailedAttemptsTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataFailedAttemptsTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataFailedAttemptsTime) GetTypeName() string {
	return "BACnetConstructedDataFailedAttemptsTime"
}

func (m *_BACnetConstructedDataFailedAttemptsTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (failedAttemptsTime)
	lengthInBits += m.FailedAttemptsTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataFailedAttemptsTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataFailedAttemptsTimeParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFailedAttemptsTime, error) {
	return BACnetConstructedDataFailedAttemptsTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataFailedAttemptsTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataFailedAttemptsTime, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataFailedAttemptsTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataFailedAttemptsTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (failedAttemptsTime)
	if pullErr := readBuffer.PullContext("failedAttemptsTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for failedAttemptsTime")
	}
_failedAttemptsTime, _failedAttemptsTimeErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _failedAttemptsTimeErr != nil {
		return nil, errors.Wrap(_failedAttemptsTimeErr, "Error parsing 'failedAttemptsTime' field of BACnetConstructedDataFailedAttemptsTime")
	}
	failedAttemptsTime := _failedAttemptsTime.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("failedAttemptsTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for failedAttemptsTime")
	}

	// Virtual field
	_actualValue := failedAttemptsTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataFailedAttemptsTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataFailedAttemptsTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataFailedAttemptsTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		FailedAttemptsTime: failedAttemptsTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataFailedAttemptsTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataFailedAttemptsTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataFailedAttemptsTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataFailedAttemptsTime")
		}

	// Simple Field (failedAttemptsTime)
	if pushErr := writeBuffer.PushContext("failedAttemptsTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for failedAttemptsTime")
	}
	_failedAttemptsTimeErr := writeBuffer.WriteSerializable(ctx, m.GetFailedAttemptsTime())
	if popErr := writeBuffer.PopContext("failedAttemptsTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for failedAttemptsTime")
	}
	if _failedAttemptsTimeErr != nil {
		return errors.Wrap(_failedAttemptsTimeErr, "Error serializing 'failedAttemptsTime' field")
	}
	// Virtual field
	actualValue := m.GetActualValue()
	_ =	actualValue
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataFailedAttemptsTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataFailedAttemptsTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataFailedAttemptsTime) isBACnetConstructedDataFailedAttemptsTime() bool {
	return true
}

func (m *_BACnetConstructedDataFailedAttemptsTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



