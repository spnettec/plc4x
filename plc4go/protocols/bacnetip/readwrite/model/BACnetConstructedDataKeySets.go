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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
	spiContext "github.com/apache/plc4x/plc4go/spi/context"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetConstructedDataKeySets is the corresponding interface of BACnetConstructedDataKeySets
type BACnetConstructedDataKeySets interface {
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetKeySets returns KeySets (property field)
	GetKeySets() []BACnetSecurityKeySet
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataKeySetsExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataKeySets.
// This is useful for switch cases.
type BACnetConstructedDataKeySetsExactly interface {
	BACnetConstructedDataKeySets
	isBACnetConstructedDataKeySets() bool
}

// _BACnetConstructedDataKeySets is the data-structure of this message
type _BACnetConstructedDataKeySets struct {
	*_BACnetConstructedData
        NumberOfDataElements BACnetApplicationTagUnsignedInteger
        KeySets []BACnetSecurityKeySet
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataKeySets)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataKeySets)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_KEY_SETS}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataKeySets) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataKeySets)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataKeySets) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataKeySets) GetKeySets() []BACnetSecurityKeySet {
	return m.KeySets
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataKeySets) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataKeySets factory function for _BACnetConstructedDataKeySets
func NewBACnetConstructedDataKeySets( numberOfDataElements BACnetApplicationTagUnsignedInteger , keySets []BACnetSecurityKeySet , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataKeySets {
	_result := &_BACnetConstructedDataKeySets{
		NumberOfDataElements: numberOfDataElements,
		KeySets: keySets,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataKeySets(structType interface{}) BACnetConstructedDataKeySets {
    if casted, ok := structType.(BACnetConstructedDataKeySets); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataKeySets); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataKeySets) GetTypeName() string {
	return "BACnetConstructedDataKeySets"
}

func (m *_BACnetConstructedDataKeySets) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.KeySets) > 0 {
		for _, element := range m.KeySets {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}


func (m *_BACnetConstructedDataKeySets) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataKeySetsParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataKeySets, error) {
	return BACnetConstructedDataKeySetsParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataKeySetsParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataKeySets, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataKeySets"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataKeySets")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Virtual field
	_zero := uint64(0)
	zero := uint64(_zero)
	_ = zero

	// Optional Field (numberOfDataElements) (Can be skipped, if a given expression evaluates to false)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if bool(bool((arrayIndexArgument) != (nil))) && bool(bool((arrayIndexArgument.GetActualValue()) == (zero))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("numberOfDataElements"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for numberOfDataElements")
		}
_val, _err := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field of BACnetConstructedDataKeySets")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (keySets)
	if pullErr := readBuffer.PullContext("keySets", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for keySets")
	}
	// Terminated array
	var keySets []BACnetSecurityKeySet
	{
		for ;!bool(IsBACnetConstructedDataClosingTag(readBuffer, false, tagNumber)); {
_item, _err := BACnetSecurityKeySetParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'keySets' field of BACnetConstructedDataKeySets")
			}
			keySets = append(keySets, _item.(BACnetSecurityKeySet))
		}
	}
	if closeErr := readBuffer.CloseContext("keySets", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for keySets")
	}

	// Validation
	if (!(bool(bool((arrayIndexArgument) != (nil))) || bool(bool(((len(keySets))) == ((2)))))) {
		return nil, errors.WithStack(utils.ParseValidationError{"keySets should have exactly 2 values"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataKeySets"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataKeySets")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataKeySets{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NumberOfDataElements: numberOfDataElements,
		KeySets: keySets,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataKeySets) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataKeySets) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataKeySets"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataKeySets")
		}
	// Virtual field
	if _zeroErr := writeBuffer.WriteVirtual(ctx, "zero", m.GetZero()); _zeroErr != nil {
		return errors.Wrap(_zeroErr, "Error serializing 'zero' field")
	}

	// Optional Field (numberOfDataElements) (Can be skipped, if the value is null)
	var numberOfDataElements BACnetApplicationTagUnsignedInteger = nil
	if m.GetNumberOfDataElements() != nil {
		if pushErr := writeBuffer.PushContext("numberOfDataElements"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for numberOfDataElements")
		}
		numberOfDataElements = m.GetNumberOfDataElements()
		_numberOfDataElementsErr := writeBuffer.WriteSerializable(ctx, numberOfDataElements)
		if popErr := writeBuffer.PopContext("numberOfDataElements"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for numberOfDataElements")
		}
		if _numberOfDataElementsErr != nil {
			return errors.Wrap(_numberOfDataElementsErr, "Error serializing 'numberOfDataElements' field")
		}
	}

	// Array Field (keySets)
	if pushErr := writeBuffer.PushContext("keySets", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for keySets")
	}
	for _curItem, _element := range m.GetKeySets() {
		_ = _curItem
		arrayCtx := spiContext.CreateArrayContext(ctx, len(m.GetKeySets()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'keySets' field")
		}
	}
	if popErr := writeBuffer.PopContext("keySets", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for keySets")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataKeySets"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataKeySets")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataKeySets) isBACnetConstructedDataKeySets() bool {
	return true
}

func (m *_BACnetConstructedDataKeySets) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



