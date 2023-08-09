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


// BACnetConstructedDataLoggingObject is the corresponding interface of BACnetConstructedDataLoggingObject
type BACnetConstructedDataLoggingObject interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLoggingObject returns LoggingObject (property field)
	GetLoggingObject() BACnetApplicationTagObjectIdentifier
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagObjectIdentifier
}

// BACnetConstructedDataLoggingObjectExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLoggingObject.
// This is useful for switch cases.
type BACnetConstructedDataLoggingObjectExactly interface {
	BACnetConstructedDataLoggingObject
	isBACnetConstructedDataLoggingObject() bool
}

// _BACnetConstructedDataLoggingObject is the data-structure of this message
type _BACnetConstructedDataLoggingObject struct {
	*_BACnetConstructedData
        LoggingObject BACnetApplicationTagObjectIdentifier
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLoggingObject)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataLoggingObject)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_LOGGING_OBJECT}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLoggingObject) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLoggingObject)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLoggingObject) GetLoggingObject() BACnetApplicationTagObjectIdentifier {
	return m.LoggingObject
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLoggingObject) GetActualValue() BACnetApplicationTagObjectIdentifier {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagObjectIdentifier(m.GetLoggingObject())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataLoggingObject factory function for _BACnetConstructedDataLoggingObject
func NewBACnetConstructedDataLoggingObject( loggingObject BACnetApplicationTagObjectIdentifier , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataLoggingObject {
	_result := &_BACnetConstructedDataLoggingObject{
		LoggingObject: loggingObject,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLoggingObject(structType any) BACnetConstructedDataLoggingObject {
    if casted, ok := structType.(BACnetConstructedDataLoggingObject); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLoggingObject); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLoggingObject) GetTypeName() string {
	return "BACnetConstructedDataLoggingObject"
}

func (m *_BACnetConstructedDataLoggingObject) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (loggingObject)
	lengthInBits += m.LoggingObject.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataLoggingObject) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLoggingObjectParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLoggingObject, error) {
	return BACnetConstructedDataLoggingObjectParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLoggingObjectParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLoggingObject, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLoggingObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLoggingObject")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (loggingObject)
	if pullErr := readBuffer.PullContext("loggingObject"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for loggingObject")
	}
_loggingObject, _loggingObjectErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _loggingObjectErr != nil {
		return nil, errors.Wrap(_loggingObjectErr, "Error parsing 'loggingObject' field of BACnetConstructedDataLoggingObject")
	}
	loggingObject := _loggingObject.(BACnetApplicationTagObjectIdentifier)
	if closeErr := readBuffer.CloseContext("loggingObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for loggingObject")
	}

	// Virtual field
	_actualValue := loggingObject
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLoggingObject"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLoggingObject")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLoggingObject{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LoggingObject: loggingObject,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLoggingObject) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLoggingObject) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLoggingObject"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLoggingObject")
		}

	// Simple Field (loggingObject)
	if pushErr := writeBuffer.PushContext("loggingObject"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for loggingObject")
	}
	_loggingObjectErr := writeBuffer.WriteSerializable(ctx, m.GetLoggingObject())
	if popErr := writeBuffer.PopContext("loggingObject"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for loggingObject")
	}
	if _loggingObjectErr != nil {
		return errors.Wrap(_loggingObjectErr, "Error serializing 'loggingObject' field")
	}
	// Virtual field
	actualValue := m.GetActualValue()
	_ =	actualValue
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLoggingObject"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLoggingObject")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataLoggingObject) isBACnetConstructedDataLoggingObject() bool {
	return true
}

func (m *_BACnetConstructedDataLoggingObject) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



