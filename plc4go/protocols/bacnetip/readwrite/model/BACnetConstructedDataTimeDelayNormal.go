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


// BACnetConstructedDataTimeDelayNormal is the corresponding interface of BACnetConstructedDataTimeDelayNormal
type BACnetConstructedDataTimeDelayNormal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetTimeDelayNormal returns TimeDelayNormal (property field)
	GetTimeDelayNormal() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataTimeDelayNormalExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataTimeDelayNormal.
// This is useful for switch cases.
type BACnetConstructedDataTimeDelayNormalExactly interface {
	BACnetConstructedDataTimeDelayNormal
	isBACnetConstructedDataTimeDelayNormal() bool
}

// _BACnetConstructedDataTimeDelayNormal is the data-structure of this message
type _BACnetConstructedDataTimeDelayNormal struct {
	*_BACnetConstructedData
        TimeDelayNormal BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataTimeDelayNormal)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataTimeDelayNormal)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_TIME_DELAY_NORMAL}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataTimeDelayNormal) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataTimeDelayNormal)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataTimeDelayNormal) GetTimeDelayNormal() BACnetApplicationTagUnsignedInteger {
	return m.TimeDelayNormal
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataTimeDelayNormal) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetTimeDelayNormal())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataTimeDelayNormal factory function for _BACnetConstructedDataTimeDelayNormal
func NewBACnetConstructedDataTimeDelayNormal( timeDelayNormal BACnetApplicationTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataTimeDelayNormal {
	_result := &_BACnetConstructedDataTimeDelayNormal{
		TimeDelayNormal: timeDelayNormal,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataTimeDelayNormal(structType any) BACnetConstructedDataTimeDelayNormal {
    if casted, ok := structType.(BACnetConstructedDataTimeDelayNormal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataTimeDelayNormal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataTimeDelayNormal) GetTypeName() string {
	return "BACnetConstructedDataTimeDelayNormal"
}

func (m *_BACnetConstructedDataTimeDelayNormal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (timeDelayNormal)
	lengthInBits += m.TimeDelayNormal.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataTimeDelayNormal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataTimeDelayNormalParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimeDelayNormal, error) {
	return BACnetConstructedDataTimeDelayNormalParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataTimeDelayNormalParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataTimeDelayNormal, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataTimeDelayNormal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataTimeDelayNormal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (timeDelayNormal)
	if pullErr := readBuffer.PullContext("timeDelayNormal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for timeDelayNormal")
	}
_timeDelayNormal, _timeDelayNormalErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _timeDelayNormalErr != nil {
		return nil, errors.Wrap(_timeDelayNormalErr, "Error parsing 'timeDelayNormal' field of BACnetConstructedDataTimeDelayNormal")
	}
	timeDelayNormal := _timeDelayNormal.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("timeDelayNormal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for timeDelayNormal")
	}

	// Virtual field
	_actualValue := timeDelayNormal
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataTimeDelayNormal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataTimeDelayNormal")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataTimeDelayNormal{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		TimeDelayNormal: timeDelayNormal,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataTimeDelayNormal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataTimeDelayNormal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataTimeDelayNormal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataTimeDelayNormal")
		}

	// Simple Field (timeDelayNormal)
	if pushErr := writeBuffer.PushContext("timeDelayNormal"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for timeDelayNormal")
	}
	_timeDelayNormalErr := writeBuffer.WriteSerializable(ctx, m.GetTimeDelayNormal())
	if popErr := writeBuffer.PopContext("timeDelayNormal"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for timeDelayNormal")
	}
	if _timeDelayNormalErr != nil {
		return errors.Wrap(_timeDelayNormalErr, "Error serializing 'timeDelayNormal' field")
	}
	// Virtual field
	actualValue := m.GetActualValue()
	_ =	actualValue
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataTimeDelayNormal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataTimeDelayNormal")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataTimeDelayNormal) isBACnetConstructedDataTimeDelayNormal() bool {
	return true
}

func (m *_BACnetConstructedDataTimeDelayNormal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



