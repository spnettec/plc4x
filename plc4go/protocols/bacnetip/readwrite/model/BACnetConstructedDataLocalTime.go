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


// BACnetConstructedDataLocalTime is the corresponding interface of BACnetConstructedDataLocalTime
type BACnetConstructedDataLocalTime interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLocalTime returns LocalTime (property field)
	GetLocalTime() BACnetApplicationTagTime
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagTime
}

// BACnetConstructedDataLocalTimeExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLocalTime.
// This is useful for switch cases.
type BACnetConstructedDataLocalTimeExactly interface {
	BACnetConstructedDataLocalTime
	isBACnetConstructedDataLocalTime() bool
}

// _BACnetConstructedDataLocalTime is the data-structure of this message
type _BACnetConstructedDataLocalTime struct {
	*_BACnetConstructedData
        LocalTime BACnetApplicationTagTime
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLocalTime)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataLocalTime)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_LOCAL_TIME}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLocalTime) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLocalTime)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLocalTime) GetLocalTime() BACnetApplicationTagTime {
	return m.LocalTime
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLocalTime) GetActualValue() BACnetApplicationTagTime {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagTime(m.GetLocalTime())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataLocalTime factory function for _BACnetConstructedDataLocalTime
func NewBACnetConstructedDataLocalTime( localTime BACnetApplicationTagTime , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataLocalTime {
	_result := &_BACnetConstructedDataLocalTime{
		LocalTime: localTime,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLocalTime(structType any) BACnetConstructedDataLocalTime {
    if casted, ok := structType.(BACnetConstructedDataLocalTime); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLocalTime); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLocalTime) GetTypeName() string {
	return "BACnetConstructedDataLocalTime"
}

func (m *_BACnetConstructedDataLocalTime) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (localTime)
	lengthInBits += m.LocalTime.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataLocalTime) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLocalTimeParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLocalTime, error) {
	return BACnetConstructedDataLocalTimeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLocalTimeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLocalTime, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLocalTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLocalTime")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (localTime)
	if pullErr := readBuffer.PullContext("localTime"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for localTime")
	}
_localTime, _localTimeErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _localTimeErr != nil {
		return nil, errors.Wrap(_localTimeErr, "Error parsing 'localTime' field of BACnetConstructedDataLocalTime")
	}
	localTime := _localTime.(BACnetApplicationTagTime)
	if closeErr := readBuffer.CloseContext("localTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for localTime")
	}

	// Virtual field
	_actualValue := localTime
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLocalTime"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLocalTime")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLocalTime{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LocalTime: localTime,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLocalTime) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLocalTime) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLocalTime"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLocalTime")
		}

	// Simple Field (localTime)
	if pushErr := writeBuffer.PushContext("localTime"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for localTime")
	}
	_localTimeErr := writeBuffer.WriteSerializable(ctx, m.GetLocalTime())
	if popErr := writeBuffer.PopContext("localTime"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for localTime")
	}
	if _localTimeErr != nil {
		return errors.Wrap(_localTimeErr, "Error serializing 'localTime' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLocalTime"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLocalTime")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataLocalTime) isBACnetConstructedDataLocalTime() bool {
	return true
}

func (m *_BACnetConstructedDataLocalTime) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



