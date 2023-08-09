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


// BACnetConstructedDataDateValueRelinquishDefault is the corresponding interface of BACnetConstructedDataDateValueRelinquishDefault
type BACnetConstructedDataDateValueRelinquishDefault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetRelinquishDefault returns RelinquishDefault (property field)
	GetRelinquishDefault() BACnetApplicationTagDate
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagDate
}

// BACnetConstructedDataDateValueRelinquishDefaultExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataDateValueRelinquishDefault.
// This is useful for switch cases.
type BACnetConstructedDataDateValueRelinquishDefaultExactly interface {
	BACnetConstructedDataDateValueRelinquishDefault
	isBACnetConstructedDataDateValueRelinquishDefault() bool
}

// _BACnetConstructedDataDateValueRelinquishDefault is the data-structure of this message
type _BACnetConstructedDataDateValueRelinquishDefault struct {
	*_BACnetConstructedData
        RelinquishDefault BACnetApplicationTagDate
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataDateValueRelinquishDefault)  GetObjectTypeArgument() BACnetObjectType {
return BACnetObjectType_DATE_VALUE}

func (m *_BACnetConstructedDataDateValueRelinquishDefault)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_RELINQUISH_DEFAULT}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataDateValueRelinquishDefault) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetRelinquishDefault() BACnetApplicationTagDate {
	return m.RelinquishDefault
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetActualValue() BACnetApplicationTagDate {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagDate(m.GetRelinquishDefault())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataDateValueRelinquishDefault factory function for _BACnetConstructedDataDateValueRelinquishDefault
func NewBACnetConstructedDataDateValueRelinquishDefault( relinquishDefault BACnetApplicationTagDate , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataDateValueRelinquishDefault {
	_result := &_BACnetConstructedDataDateValueRelinquishDefault{
		RelinquishDefault: relinquishDefault,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataDateValueRelinquishDefault(structType any) BACnetConstructedDataDateValueRelinquishDefault {
    if casted, ok := structType.(BACnetConstructedDataDateValueRelinquishDefault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataDateValueRelinquishDefault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetTypeName() string {
	return "BACnetConstructedDataDateValueRelinquishDefault"
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (relinquishDefault)
	lengthInBits += m.RelinquishDefault.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataDateValueRelinquishDefault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataDateValueRelinquishDefaultParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDateValueRelinquishDefault, error) {
	return BACnetConstructedDataDateValueRelinquishDefaultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataDateValueRelinquishDefaultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataDateValueRelinquishDefault, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataDateValueRelinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataDateValueRelinquishDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (relinquishDefault)
	if pullErr := readBuffer.PullContext("relinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for relinquishDefault")
	}
_relinquishDefault, _relinquishDefaultErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _relinquishDefaultErr != nil {
		return nil, errors.Wrap(_relinquishDefaultErr, "Error parsing 'relinquishDefault' field of BACnetConstructedDataDateValueRelinquishDefault")
	}
	relinquishDefault := _relinquishDefault.(BACnetApplicationTagDate)
	if closeErr := readBuffer.CloseContext("relinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for relinquishDefault")
	}

	// Virtual field
	_actualValue := relinquishDefault
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataDateValueRelinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataDateValueRelinquishDefault")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataDateValueRelinquishDefault{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		RelinquishDefault: relinquishDefault,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataDateValueRelinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataDateValueRelinquishDefault")
		}

	// Simple Field (relinquishDefault)
	if pushErr := writeBuffer.PushContext("relinquishDefault"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for relinquishDefault")
	}
	_relinquishDefaultErr := writeBuffer.WriteSerializable(ctx, m.GetRelinquishDefault())
	if popErr := writeBuffer.PopContext("relinquishDefault"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for relinquishDefault")
	}
	if _relinquishDefaultErr != nil {
		return errors.Wrap(_relinquishDefaultErr, "Error serializing 'relinquishDefault' field")
	}
	// Virtual field
	actualValue := m.GetActualValue()
	_ =	actualValue
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataDateValueRelinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataDateValueRelinquishDefault")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataDateValueRelinquishDefault) isBACnetConstructedDataDateValueRelinquishDefault() bool {
	return true
}

func (m *_BACnetConstructedDataDateValueRelinquishDefault) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



