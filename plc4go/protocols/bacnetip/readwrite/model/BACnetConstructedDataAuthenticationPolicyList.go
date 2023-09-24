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
	"io"
	"github.com/rs/zerolog"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetConstructedDataAuthenticationPolicyList is the corresponding interface of BACnetConstructedDataAuthenticationPolicyList
type BACnetConstructedDataAuthenticationPolicyList interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetNumberOfDataElements returns NumberOfDataElements (property field)
	GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger
	// GetAuthenticationPolicyList returns AuthenticationPolicyList (property field)
	GetAuthenticationPolicyList() []BACnetAuthenticationPolicy
	// GetZero returns Zero (virtual field)
	GetZero() uint64
}

// BACnetConstructedDataAuthenticationPolicyListExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataAuthenticationPolicyList.
// This is useful for switch cases.
type BACnetConstructedDataAuthenticationPolicyListExactly interface {
	BACnetConstructedDataAuthenticationPolicyList
	isBACnetConstructedDataAuthenticationPolicyList() bool
}

// _BACnetConstructedDataAuthenticationPolicyList is the data-structure of this message
type _BACnetConstructedDataAuthenticationPolicyList struct {
	*_BACnetConstructedData
        NumberOfDataElements BACnetApplicationTagUnsignedInteger
        AuthenticationPolicyList []BACnetAuthenticationPolicy
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyList)  GetObjectTypeArgument() BACnetObjectType {
return 0}

func (m *_BACnetConstructedDataAuthenticationPolicyList)  GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
return BACnetPropertyIdentifier_AUTHENTICATION_POLICY_LIST}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyList) InitializeParent(parent BACnetConstructedData , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag ) {	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataAuthenticationPolicyList)  GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyList) GetNumberOfDataElements() BACnetApplicationTagUnsignedInteger {
	return m.NumberOfDataElements
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) GetAuthenticationPolicyList() []BACnetAuthenticationPolicy {
	return m.AuthenticationPolicyList
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataAuthenticationPolicyList) GetZero() uint64 {
	ctx := context.Background()
	_ = ctx
	numberOfDataElements := m.NumberOfDataElements
	_ = numberOfDataElements
	return uint64(uint64(0))
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataAuthenticationPolicyList factory function for _BACnetConstructedDataAuthenticationPolicyList
func NewBACnetConstructedDataAuthenticationPolicyList( numberOfDataElements BACnetApplicationTagUnsignedInteger , authenticationPolicyList []BACnetAuthenticationPolicy , openingTag BACnetOpeningTag , peekedTagHeader BACnetTagHeader , closingTag BACnetClosingTag , tagNumber uint8 , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataAuthenticationPolicyList {
	_result := &_BACnetConstructedDataAuthenticationPolicyList{
		NumberOfDataElements: numberOfDataElements,
		AuthenticationPolicyList: authenticationPolicyList,
    	_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataAuthenticationPolicyList(structType any) BACnetConstructedDataAuthenticationPolicyList {
    if casted, ok := structType.(BACnetConstructedDataAuthenticationPolicyList); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataAuthenticationPolicyList); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) GetTypeName() string {
	return "BACnetConstructedDataAuthenticationPolicyList"
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// A virtual field doesn't have any in- or output.

	// Optional Field (numberOfDataElements)
	if m.NumberOfDataElements != nil {
		lengthInBits += m.NumberOfDataElements.GetLengthInBits(ctx)
	}

	// Array field
	if len(m.AuthenticationPolicyList) > 0 {
		for _, element := range m.AuthenticationPolicyList {
			lengthInBits += element.GetLengthInBits(ctx)
		}
	}

	return lengthInBits
}


func (m *_BACnetConstructedDataAuthenticationPolicyList) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataAuthenticationPolicyListParse(ctx context.Context, theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAuthenticationPolicyList, error) {
	return BACnetConstructedDataAuthenticationPolicyListParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataAuthenticationPolicyListParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataAuthenticationPolicyList, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataAuthenticationPolicyList"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataAuthenticationPolicyList")
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
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'numberOfDataElements' field of BACnetConstructedDataAuthenticationPolicyList")
		default:
			numberOfDataElements = _val.(BACnetApplicationTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("numberOfDataElements"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for numberOfDataElements")
			}
		}
	}

	// Array field (authenticationPolicyList)
	if pullErr := readBuffer.PullContext("authenticationPolicyList", utils.WithRenderAsList(true)); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for authenticationPolicyList")
	}
	// Terminated array
	var authenticationPolicyList []BACnetAuthenticationPolicy
	{
		for ;!bool(IsBACnetConstructedDataClosingTag(ctx, readBuffer, false, tagNumber)); {
_item, _err := BACnetAuthenticationPolicyParseWithBuffer(ctx, readBuffer)
			if _err != nil {
				return nil, errors.Wrap(_err, "Error parsing 'authenticationPolicyList' field of BACnetConstructedDataAuthenticationPolicyList")
			}
			authenticationPolicyList = append(authenticationPolicyList, _item.(BACnetAuthenticationPolicy))
		}
	}
	if closeErr := readBuffer.CloseContext("authenticationPolicyList", utils.WithRenderAsList(true)); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for authenticationPolicyList")
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataAuthenticationPolicyList"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataAuthenticationPolicyList")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataAuthenticationPolicyList{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber: tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		NumberOfDataElements: numberOfDataElements,
		AuthenticationPolicyList: authenticationPolicyList,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataAuthenticationPolicyList"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataAuthenticationPolicyList")
		}
	// Virtual field
	zero := m.GetZero()
	_ =	zero
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

	// Array Field (authenticationPolicyList)
	if pushErr := writeBuffer.PushContext("authenticationPolicyList", utils.WithRenderAsList(true)); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for authenticationPolicyList")
	}
	for _curItem, _element := range m.GetAuthenticationPolicyList() {
		_ = _curItem
		arrayCtx := utils.CreateArrayContext(ctx, len(m.GetAuthenticationPolicyList()), _curItem)
		_ = arrayCtx
		_elementErr := writeBuffer.WriteSerializable(arrayCtx, _element)
		if _elementErr != nil {
			return errors.Wrap(_elementErr, "Error serializing 'authenticationPolicyList' field")
		}
	}
	if popErr := writeBuffer.PopContext("authenticationPolicyList", utils.WithRenderAsList(true)); popErr != nil {
		return errors.Wrap(popErr, "Error popping for authenticationPolicyList")
	}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataAuthenticationPolicyList"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataAuthenticationPolicyList")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetConstructedDataAuthenticationPolicyList) isBACnetConstructedDataAuthenticationPolicyList() bool {
	return true
}

func (m *_BACnetConstructedDataAuthenticationPolicyList) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



