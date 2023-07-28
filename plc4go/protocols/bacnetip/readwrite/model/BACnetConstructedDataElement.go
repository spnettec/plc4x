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


// BACnetConstructedDataElement is the corresponding interface of BACnetConstructedDataElement
type BACnetConstructedDataElement interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetApplicationTag returns ApplicationTag (property field)
	GetApplicationTag() BACnetApplicationTag
	// GetContextTag returns ContextTag (property field)
	GetContextTag() BACnetContextTag
	// GetConstructedData returns ConstructedData (property field)
	GetConstructedData() BACnetConstructedData
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
	// GetIsApplicationTag returns IsApplicationTag (virtual field)
	GetIsApplicationTag() bool
	// GetIsConstructedData returns IsConstructedData (virtual field)
	GetIsConstructedData() bool
	// GetIsContextTag returns IsContextTag (virtual field)
	GetIsContextTag() bool
}

// BACnetConstructedDataElementExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataElement.
// This is useful for switch cases.
type BACnetConstructedDataElementExactly interface {
	BACnetConstructedDataElement
	isBACnetConstructedDataElement() bool
}

// _BACnetConstructedDataElement is the data-structure of this message
type _BACnetConstructedDataElement struct {
        PeekedTagHeader BACnetTagHeader
        ApplicationTag BACnetApplicationTag
        ContextTag BACnetContextTag
        ConstructedData BACnetConstructedData

	// Arguments.
	ObjectTypeArgument BACnetObjectType
	PropertyIdentifierArgument BACnetPropertyIdentifier
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataElement) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetConstructedDataElement) GetApplicationTag() BACnetApplicationTag {
	return m.ApplicationTag
}

func (m *_BACnetConstructedDataElement) GetContextTag() BACnetContextTag {
	return m.ContextTag
}

func (m *_BACnetConstructedDataElement) GetConstructedData() BACnetConstructedData {
	return m.ConstructedData
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataElement) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	applicationTag := m.ApplicationTag
	_ = applicationTag
	contextTag := m.ContextTag
	_ = contextTag
	constructedData := m.ConstructedData
	_ = constructedData
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

func (m *_BACnetConstructedDataElement) GetIsApplicationTag() bool {
	ctx := context.Background()
	_ = ctx
	applicationTag := m.ApplicationTag
	_ = applicationTag
	contextTag := m.ContextTag
	_ = contextTag
	constructedData := m.ConstructedData
	_ = constructedData
	return bool(bool((m.GetPeekedTagHeader().GetTagClass()) == (TagClass_APPLICATION_TAGS)))
}

func (m *_BACnetConstructedDataElement) GetIsConstructedData() bool {
	ctx := context.Background()
	_ = ctx
	applicationTag := m.ApplicationTag
	_ = applicationTag
	contextTag := m.ContextTag
	_ = contextTag
	constructedData := m.ConstructedData
	_ = constructedData
	return bool(bool(!(m.GetIsApplicationTag())) && bool(bool((m.GetPeekedTagHeader().GetLengthValueType()) == (0x6))))
}

func (m *_BACnetConstructedDataElement) GetIsContextTag() bool {
	ctx := context.Background()
	_ = ctx
	applicationTag := m.ApplicationTag
	_ = applicationTag
	contextTag := m.ContextTag
	_ = contextTag
	constructedData := m.ConstructedData
	_ = constructedData
	return bool(bool(!(m.GetIsConstructedData())) && bool(!(m.GetIsApplicationTag())))
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConstructedDataElement factory function for _BACnetConstructedDataElement
func NewBACnetConstructedDataElement( peekedTagHeader BACnetTagHeader , applicationTag BACnetApplicationTag , contextTag BACnetContextTag , constructedData BACnetConstructedData , objectTypeArgument BACnetObjectType , propertyIdentifierArgument BACnetPropertyIdentifier , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetConstructedDataElement {
return &_BACnetConstructedDataElement{ PeekedTagHeader: peekedTagHeader , ApplicationTag: applicationTag , ContextTag: contextTag , ConstructedData: constructedData , ObjectTypeArgument: objectTypeArgument , PropertyIdentifierArgument: propertyIdentifierArgument , ArrayIndexArgument: arrayIndexArgument }
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataElement(structType any) BACnetConstructedDataElement {
    if casted, ok := structType.(BACnetConstructedDataElement); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataElement); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataElement) GetTypeName() string {
	return "BACnetConstructedDataElement"
}

func (m *_BACnetConstructedDataElement) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// Optional Field (applicationTag)
	if m.ApplicationTag != nil {
		lengthInBits += m.ApplicationTag.GetLengthInBits(ctx)
	}

	// Optional Field (contextTag)
	if m.ContextTag != nil {
		lengthInBits += m.ContextTag.GetLengthInBits(ctx)
	}

	// Optional Field (constructedData)
	if m.ConstructedData != nil {
		lengthInBits += m.ConstructedData.GetLengthInBits(ctx)
	}

	return lengthInBits
}


func (m *_BACnetConstructedDataElement) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataElementParse(ctx context.Context, theBytes []byte, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataElement, error) {
	return BACnetConstructedDataElementParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataElementParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataElement, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetConstructedDataElement"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataElement")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

				// Peek Field (peekedTagHeader)
				currentPos = positionAware.GetPos()
				if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
					return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
				}
peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
				readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Virtual field
	_isApplicationTag := bool((peekedTagHeader.GetTagClass()) == (TagClass_APPLICATION_TAGS))
	isApplicationTag := bool(_isApplicationTag)
	_ = isApplicationTag

	// Virtual field
	_isConstructedData := bool(!(isApplicationTag)) && bool(bool((peekedTagHeader.GetLengthValueType()) == (0x6)))
	isConstructedData := bool(_isConstructedData)
	_ = isConstructedData

	// Virtual field
	_isContextTag := bool(!(isConstructedData)) && bool(!(isApplicationTag))
	isContextTag := bool(_isContextTag)
	_ = isContextTag

	// Validation
	if (!(bool(!(isContextTag)) || bool((bool(isContextTag) && bool(bool((peekedTagHeader.GetLengthValueType()) != (0x7))))))) {
		return nil, errors.WithStack(utils.ParseValidationError{"unexpected closing tag"})
	}

	// Optional Field (applicationTag) (Can be skipped, if a given expression evaluates to false)
	var applicationTag BACnetApplicationTag = nil
	if isApplicationTag {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("applicationTag"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for applicationTag")
		}
_val, _err := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'applicationTag' field of BACnetConstructedDataElement")
		default:
			applicationTag = _val.(BACnetApplicationTag)
			if closeErr := readBuffer.CloseContext("applicationTag"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for applicationTag")
			}
		}
	}

	// Optional Field (contextTag) (Can be skipped, if a given expression evaluates to false)
	var contextTag BACnetContextTag = nil
	if isContextTag {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("contextTag"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for contextTag")
		}
_val, _err := BACnetContextTagParseWithBuffer(ctx, readBuffer , peekedTagNumber , BACnetDataType_UNKNOWN )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'contextTag' field of BACnetConstructedDataElement")
		default:
			contextTag = _val.(BACnetContextTag)
			if closeErr := readBuffer.CloseContext("contextTag"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for contextTag")
			}
		}
	}

	// Optional Field (constructedData) (Can be skipped, if a given expression evaluates to false)
	var constructedData BACnetConstructedData = nil
	if isConstructedData {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("constructedData"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for constructedData")
		}
_val, _err := BACnetConstructedDataParseWithBuffer(ctx, readBuffer , peekedTagNumber , objectTypeArgument , propertyIdentifierArgument , arrayIndexArgument )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'constructedData' field of BACnetConstructedDataElement")
		default:
			constructedData = _val.(BACnetConstructedData)
			if closeErr := readBuffer.CloseContext("constructedData"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for constructedData")
			}
		}
	}

	// Validation
	if (!(bool(bool((bool(isApplicationTag) && bool(bool(((applicationTag)) != (nil))))) || bool((bool(isContextTag) && bool(bool(((contextTag)) != (nil)))))) || bool((bool(isConstructedData) && bool(bool(((constructedData)) != (nil))))))) {
		return nil, errors.WithStack(utils.ParseValidationError{"BACnetConstructedDataElement could not parse anything"})
	}

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataElement"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataElement")
	}

	// Create the instance
	return &_BACnetConstructedDataElement{
            ObjectTypeArgument: objectTypeArgument,
            PropertyIdentifierArgument: propertyIdentifierArgument,
            ArrayIndexArgument: arrayIndexArgument,
			PeekedTagHeader: peekedTagHeader,
			ApplicationTag: applicationTag,
			ContextTag: contextTag,
			ConstructedData: constructedData,
		}, nil
}

func (m *_BACnetConstructedDataElement) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataElement) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pushErr :=writeBuffer.PushContext("BACnetConstructedDataElement"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataElement")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}
	// Virtual field
	if _isApplicationTagErr := writeBuffer.WriteVirtual(ctx, "isApplicationTag", m.GetIsApplicationTag()); _isApplicationTagErr != nil {
		return errors.Wrap(_isApplicationTagErr, "Error serializing 'isApplicationTag' field")
	}
	// Virtual field
	if _isConstructedDataErr := writeBuffer.WriteVirtual(ctx, "isConstructedData", m.GetIsConstructedData()); _isConstructedDataErr != nil {
		return errors.Wrap(_isConstructedDataErr, "Error serializing 'isConstructedData' field")
	}
	// Virtual field
	if _isContextTagErr := writeBuffer.WriteVirtual(ctx, "isContextTag", m.GetIsContextTag()); _isContextTagErr != nil {
		return errors.Wrap(_isContextTagErr, "Error serializing 'isContextTag' field")
	}

	// Optional Field (applicationTag) (Can be skipped, if the value is null)
	var applicationTag BACnetApplicationTag = nil
	if m.GetApplicationTag() != nil {
		if pushErr := writeBuffer.PushContext("applicationTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for applicationTag")
		}
		applicationTag = m.GetApplicationTag()
		_applicationTagErr := writeBuffer.WriteSerializable(ctx, applicationTag)
		if popErr := writeBuffer.PopContext("applicationTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for applicationTag")
		}
		if _applicationTagErr != nil {
			return errors.Wrap(_applicationTagErr, "Error serializing 'applicationTag' field")
		}
	}

	// Optional Field (contextTag) (Can be skipped, if the value is null)
	var contextTag BACnetContextTag = nil
	if m.GetContextTag() != nil {
		if pushErr := writeBuffer.PushContext("contextTag"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for contextTag")
		}
		contextTag = m.GetContextTag()
		_contextTagErr := writeBuffer.WriteSerializable(ctx, contextTag)
		if popErr := writeBuffer.PopContext("contextTag"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for contextTag")
		}
		if _contextTagErr != nil {
			return errors.Wrap(_contextTagErr, "Error serializing 'contextTag' field")
		}
	}

	// Optional Field (constructedData) (Can be skipped, if the value is null)
	var constructedData BACnetConstructedData = nil
	if m.GetConstructedData() != nil {
		if pushErr := writeBuffer.PushContext("constructedData"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for constructedData")
		}
		constructedData = m.GetConstructedData()
		_constructedDataErr := writeBuffer.WriteSerializable(ctx, constructedData)
		if popErr := writeBuffer.PopContext("constructedData"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for constructedData")
		}
		if _constructedDataErr != nil {
			return errors.Wrap(_constructedDataErr, "Error serializing 'constructedData' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetConstructedDataElement"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetConstructedDataElement")
	}
	return nil
}


////
// Arguments Getter

func (m *_BACnetConstructedDataElement) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}
func (m *_BACnetConstructedDataElement) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return m.PropertyIdentifierArgument
}
func (m *_BACnetConstructedDataElement) GetArrayIndexArgument() BACnetTagPayloadUnsignedInteger {
	return m.ArrayIndexArgument
}
//
////

func (m *_BACnetConstructedDataElement) isBACnetConstructedDataElement() bool {
	return true
}

func (m *_BACnetConstructedDataElement) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



