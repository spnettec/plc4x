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


// BACnetConstructedDataLinkSpeed is the corresponding interface of BACnetConstructedDataLinkSpeed
type BACnetConstructedDataLinkSpeed interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetLinkSpeed returns LinkSpeed (property field)
	GetLinkSpeed() BACnetApplicationTagReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagReal
}

// BACnetConstructedDataLinkSpeedExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataLinkSpeed.
// This is useful for switch cases.
type BACnetConstructedDataLinkSpeedExactly interface {
	BACnetConstructedDataLinkSpeed
	isBACnetConstructedDataLinkSpeed() bool
}

// _BACnetConstructedDataLinkSpeed is the data-structure of this message
type _BACnetConstructedDataLinkSpeed struct {
	*_BACnetConstructedData
        LinkSpeed BACnetApplicationTagReal
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataLinkSpeed)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataLinkSpeed)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_LINK_SPEED}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataLinkSpeed) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataLinkSpeed)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataLinkSpeed) GetLinkSpeed() BACnetApplicationTagReal {
	return m.LinkSpeed
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataLinkSpeed) GetActualValue() BACnetApplicationTagReal {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagReal(m.GetLinkSpeed())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataLinkSpeed factory function for _BACnetConstructedDataLinkSpeed
func NewBACnetConstructedDataLinkSpeed( linkSpeed BACnetApplicationTagReal , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataLinkSpeed {
	_result := &_BACnetConstructedDataLinkSpeed{
		LinkSpeed: linkSpeed,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataLinkSpeed(structType any) BACnetConstructedDataLinkSpeed {
    if casted, ok := structType.(BACnetConstructedDataLinkSpeed); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataLinkSpeed); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataLinkSpeed) GetTypeName() string {
	return "BACnetConstructedDataLinkSpeed"
}

func (m *_BACnetConstructedDataLinkSpeed) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (linkSpeed)
	lengthInBits += m.LinkSpeed.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetConstructedDataLinkSpeed) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataLinkSpeedParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLinkSpeed, error) {
	return BACnetConstructedDataLinkSpeedParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataLinkSpeedParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataLinkSpeed, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataLinkSpeed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataLinkSpeed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (linkSpeed)
	if pullErr := readBuffer.PullContext("linkSpeed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for linkSpeed")
	}
_linkSpeed, _linkSpeedErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _linkSpeedErr != nil {
		return nil, errors.Wrap(_linkSpeedErr, "Error parsing 'linkSpeed' field of BACnetConstructedDataLinkSpeed")
	}
	linkSpeed := _linkSpeed.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("linkSpeed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for linkSpeed")
	}

	// Virtual field
	_actualValue := linkSpeed
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataLinkSpeed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataLinkSpeed")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataLinkSpeed{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		LinkSpeed: linkSpeed,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataLinkSpeed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataLinkSpeed) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataLinkSpeed"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataLinkSpeed")
		}

	// Simple Field (linkSpeed)
	if pushErr := writeBuffer.PushContext("linkSpeed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for linkSpeed")
	}
	_linkSpeedErr := writeBuffer.WriteSerializable(ctx, m.GetLinkSpeed())
	if popErr := writeBuffer.PopContext("linkSpeed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for linkSpeed")
	}
	if _linkSpeedErr != nil {
		return errors.Wrap(_linkSpeedErr, "Error serializing 'linkSpeed' field")
	}
	// Virtual field
	actualValue := m.GetActualValue()
	_ =	actualValue
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataLinkSpeed"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataLinkSpeed")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataLinkSpeed) isBACnetConstructedDataLinkSpeed() bool {
	return true
}

func (m *_BACnetConstructedDataLinkSpeed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



