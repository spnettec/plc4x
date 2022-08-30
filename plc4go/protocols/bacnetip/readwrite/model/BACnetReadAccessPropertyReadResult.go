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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"github.com/rs/zerolog/log"
	"io"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetReadAccessPropertyReadResult is the corresponding interface of BACnetReadAccessPropertyReadResult
type BACnetReadAccessPropertyReadResult interface {
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPropertyValue returns PropertyValue (property field)
	GetPropertyValue() BACnetConstructedData
	// GetPropertyAccessError returns PropertyAccessError (property field)
	GetPropertyAccessError() ErrorEnclosed
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetReadAccessPropertyReadResultExactly can be used when we want exactly this type and not a type which fulfills BACnetReadAccessPropertyReadResult.
// This is useful for switch cases.
type BACnetReadAccessPropertyReadResultExactly interface {
	BACnetReadAccessPropertyReadResult
	isBACnetReadAccessPropertyReadResult() bool
}

// _BACnetReadAccessPropertyReadResult is the data-structure of this message
type _BACnetReadAccessPropertyReadResult struct {
        PeekedTagHeader BACnetTagHeader
        PropertyValue BACnetConstructedData
        PropertyAccessError ErrorEnclosed

	// Arguments.
	ObjectTypeArgument BACnetObjectType
	PropertyIdentifierArgument BACnetPropertyIdentifier
	ArrayIndexArgument BACnetTagPayloadUnsignedInteger
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetReadAccessPropertyReadResult) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

func (m *_BACnetReadAccessPropertyReadResult) GetPropertyValue() BACnetConstructedData {
	return m.PropertyValue
}

func (m *_BACnetReadAccessPropertyReadResult) GetPropertyAccessError() ErrorEnclosed {
	return m.PropertyAccessError
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetReadAccessPropertyReadResult) GetPeekedTagNumber() uint8 {
	propertyValue := m.PropertyValue
	_ = propertyValue
	propertyAccessError := m.PropertyAccessError
	_ = propertyAccessError
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetReadAccessPropertyReadResult factory function for _BACnetReadAccessPropertyReadResult
func NewBACnetReadAccessPropertyReadResult( peekedTagHeader BACnetTagHeader , propertyValue BACnetConstructedData , propertyAccessError ErrorEnclosed , objectTypeArgument BACnetObjectType , propertyIdentifierArgument BACnetPropertyIdentifier , arrayIndexArgument BACnetTagPayloadUnsignedInteger ) *_BACnetReadAccessPropertyReadResult {
return &_BACnetReadAccessPropertyReadResult{ PeekedTagHeader: peekedTagHeader , PropertyValue: propertyValue , PropertyAccessError: propertyAccessError , ObjectTypeArgument: objectTypeArgument , PropertyIdentifierArgument: propertyIdentifierArgument , ArrayIndexArgument: arrayIndexArgument }
}

// Deprecated: use the interface for direct cast
func CastBACnetReadAccessPropertyReadResult(structType interface{}) BACnetReadAccessPropertyReadResult {
    if casted, ok := structType.(BACnetReadAccessPropertyReadResult); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetReadAccessPropertyReadResult); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetReadAccessPropertyReadResult) GetTypeName() string {
	return "BACnetReadAccessPropertyReadResult"
}

func (m *_BACnetReadAccessPropertyReadResult) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetReadAccessPropertyReadResult) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	// Optional Field (propertyValue)
	if m.PropertyValue != nil {
		lengthInBits += m.PropertyValue.GetLengthInBits()
	}

	// Optional Field (propertyAccessError)
	if m.PropertyAccessError != nil {
		lengthInBits += m.PropertyAccessError.GetLengthInBits()
	}

	return lengthInBits
}


func (m *_BACnetReadAccessPropertyReadResult) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetReadAccessPropertyReadResultParse(readBuffer utils.ReadBuffer, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetReadAccessPropertyReadResult, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetReadAccessPropertyReadResult"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetReadAccessPropertyReadResult")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

				// Peek Field (peekedTagHeader)
				currentPos = positionAware.GetPos()
				if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
					return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
				}
peekedTagHeader, _ := BACnetTagHeaderParse(readBuffer)
				readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Optional Field (propertyValue) (Can be skipped, if a given expression evaluates to false)
	var propertyValue BACnetConstructedData = nil
	if bool((peekedTagNumber) == ((4))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("propertyValue"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for propertyValue")
		}
_val, _err := BACnetConstructedDataParse(readBuffer , uint8(4) , objectTypeArgument , propertyIdentifierArgument , arrayIndexArgument )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyValue' field of BACnetReadAccessPropertyReadResult")
		default:
			propertyValue = _val.(BACnetConstructedData)
			if closeErr := readBuffer.CloseContext("propertyValue"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for propertyValue")
			}
		}
	}

	// Validation
	if (!(bool((bool(bool((peekedTagNumber) == ((4)))) && bool(bool(((propertyValue)) != (nil))))) || bool(bool((peekedTagNumber) != ((4)))))) {
		return nil, errors.WithStack(utils.ParseValidationError{"failure parsing field 4"})
	}

	// Optional Field (propertyAccessError) (Can be skipped, if a given expression evaluates to false)
	var propertyAccessError ErrorEnclosed = nil
	if bool((peekedTagNumber) == ((5))) {
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("propertyAccessError"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for propertyAccessError")
		}
_val, _err := ErrorEnclosedParse(readBuffer , uint8(5) )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			log.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'propertyAccessError' field of BACnetReadAccessPropertyReadResult")
		default:
			propertyAccessError = _val.(ErrorEnclosed)
			if closeErr := readBuffer.CloseContext("propertyAccessError"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for propertyAccessError")
			}
		}
	}

	// Validation
	if (!(bool((bool(bool((peekedTagNumber) == ((5)))) && bool(bool(((propertyAccessError)) != (nil))))) || bool(bool((peekedTagNumber) != ((5)))))) {
		return nil, errors.WithStack(utils.ParseValidationError{"failure parsing field 5"})
	}

	// Validation
	if (!(bool(bool((peekedTagNumber) == ((4)))) || bool(bool((peekedTagNumber) == ((5)))))) {
		return nil, errors.WithStack(utils.ParseAssertError{"should be either 4 or 5"})
	}

	if closeErr := readBuffer.CloseContext("BACnetReadAccessPropertyReadResult"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetReadAccessPropertyReadResult")
	}

	// Create the instance
	return &_BACnetReadAccessPropertyReadResult{
            ObjectTypeArgument: objectTypeArgument,
            PropertyIdentifierArgument: propertyIdentifierArgument,
            ArrayIndexArgument: arrayIndexArgument,
			PeekedTagHeader: peekedTagHeader,
			PropertyValue: propertyValue,
			PropertyAccessError: propertyAccessError,
		}, nil
}

func (m *_BACnetReadAccessPropertyReadResult) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetReadAccessPropertyReadResult"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetReadAccessPropertyReadResult")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Optional Field (propertyValue) (Can be skipped, if the value is null)
	var propertyValue BACnetConstructedData = nil
	if m.GetPropertyValue() != nil {
		if pushErr := writeBuffer.PushContext("propertyValue"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyValue")
		}
		propertyValue = m.GetPropertyValue()
		_propertyValueErr := writeBuffer.WriteSerializable(propertyValue)
		if popErr := writeBuffer.PopContext("propertyValue"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyValue")
		}
		if _propertyValueErr != nil {
			return errors.Wrap(_propertyValueErr, "Error serializing 'propertyValue' field")
		}
	}

	// Optional Field (propertyAccessError) (Can be skipped, if the value is null)
	var propertyAccessError ErrorEnclosed = nil
	if m.GetPropertyAccessError() != nil {
		if pushErr := writeBuffer.PushContext("propertyAccessError"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for propertyAccessError")
		}
		propertyAccessError = m.GetPropertyAccessError()
		_propertyAccessErrorErr := writeBuffer.WriteSerializable(propertyAccessError)
		if popErr := writeBuffer.PopContext("propertyAccessError"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for propertyAccessError")
		}
		if _propertyAccessErrorErr != nil {
			return errors.Wrap(_propertyAccessErrorErr, "Error serializing 'propertyAccessError' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetReadAccessPropertyReadResult"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetReadAccessPropertyReadResult")
	}
	return nil
}


////
// Arguments Getter

func (m *_BACnetReadAccessPropertyReadResult) GetObjectTypeArgument() BACnetObjectType {
	return m.ObjectTypeArgument
}
func (m *_BACnetReadAccessPropertyReadResult) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return m.PropertyIdentifierArgument
}
func (m *_BACnetReadAccessPropertyReadResult) GetArrayIndexArgument() BACnetTagPayloadUnsignedInteger {
	return m.ArrayIndexArgument
}
//
////

func (m *_BACnetReadAccessPropertyReadResult) isBACnetReadAccessPropertyReadResult() bool {
	return true
}

func (m *_BACnetReadAccessPropertyReadResult) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



