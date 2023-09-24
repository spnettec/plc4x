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


// BACnetConstructedDataRelinquishDefault is the corresponding interface of BACnetConstructedDataRelinquishDefault
type BACnetConstructedDataRelinquishDefault interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetRelinquishDefault returns RelinquishDefault (property field)
	GetRelinquishDefault() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataRelinquishDefaultExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataRelinquishDefault.
// This is useful for switch cases.
type BACnetConstructedDataRelinquishDefaultExactly interface {
	BACnetConstructedDataRelinquishDefault
	isBACnetConstructedDataRelinquishDefault() bool
}

// _BACnetConstructedDataRelinquishDefault is the data-structure of this message
type _BACnetConstructedDataRelinquishDefault struct {
	*_BACnetConstructedData
        RelinquishDefault BACnetApplicationTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataRelinquishDefault)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataRelinquishDefault)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_RELINQUISH_DEFAULT}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataRelinquishDefault) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataRelinquishDefault)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataRelinquishDefault) GetRelinquishDefault() BACnetApplicationTagUnsignedInteger {
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

func (m *_BACnetConstructedDataRelinquishDefault) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetRelinquishDefault())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataRelinquishDefault factory function for _BACnetConstructedDataRelinquishDefault
func NewBACnetConstructedDataRelinquishDefault( relinquishDefault BACnetApplicationTagUnsignedInteger , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataRelinquishDefault {
	_result := &_BACnetConstructedDataRelinquishDefault{
		RelinquishDefault: relinquishDefault,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataRelinquishDefault(structType any) BACnetConstructedDataRelinquishDefault {
    if casted, ok := structType.(BACnetConstructedDataRelinquishDefault); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataRelinquishDefault); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataRelinquishDefault) GetTypeName() string {
	return "BACnetConstructedDataRelinquishDefault"
}

func (m *_BACnetConstructedDataRelinquishDefault) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (relinquishDefault)
	lengthInBits += m.RelinquishDefault.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataRelinquishDefault) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataRelinquishDefaultParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataRelinquishDefault, error) {
	return BACnetConstructedDataRelinquishDefaultParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataRelinquishDefaultParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataRelinquishDefault, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataRelinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataRelinquishDefault")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (relinquishDefault)
	if pullErr := readBuffer.PullContext("relinquishDefault"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for relinquishDefault")
	}
_relinquishDefault, _relinquishDefaultErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _relinquishDefaultErr != nil {
		return nil, errors.Wrap(_relinquishDefaultErr, "Error parsing 'relinquishDefault' field of BACnetConstructedDataRelinquishDefault")
	}
	relinquishDefault := _relinquishDefault.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("relinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for relinquishDefault")
	}

	// Virtual field
	_actualValue := relinquishDefault
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataRelinquishDefault"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataRelinquishDefault")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataRelinquishDefault{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		RelinquishDefault: relinquishDefault,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataRelinquishDefault) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataRelinquishDefault) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataRelinquishDefault"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataRelinquishDefault")
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

		if popErr := writeBuffer.PopContext("BACnetConstructedDataRelinquishDefault"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataRelinquishDefault")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataRelinquishDefault) isBACnetConstructedDataRelinquishDefault() bool {
	return true
}

func (m *_BACnetConstructedDataRelinquishDefault) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



