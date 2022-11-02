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
	"encoding/binary"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetAuthenticationFactorFormat is the corresponding interface of BACnetAuthenticationFactorFormat
type BACnetAuthenticationFactorFormat interface {
	utils.LengthAware
	utils.Serializable
	// GetFormatType returns FormatType (property field)
	GetFormatType() BACnetAuthenticationFactorTypeTagged
	// GetVendorId returns VendorId (property field)
	GetVendorId() BACnetVendorIdTagged
	// GetVendorFormat returns VendorFormat (property field)
	GetVendorFormat() BACnetContextTagUnsignedInteger
}

// BACnetAuthenticationFactorFormatExactly can be used when we want exactly this type and not a type which fulfills BACnetAuthenticationFactorFormat.
// This is useful for switch cases.
type BACnetAuthenticationFactorFormatExactly interface {
	BACnetAuthenticationFactorFormat
	isBACnetAuthenticationFactorFormat() bool
}

// _BACnetAuthenticationFactorFormat is the data-structure of this message
type _BACnetAuthenticationFactorFormat struct {
        FormatType BACnetAuthenticationFactorTypeTagged
        VendorId BACnetVendorIdTagged
        VendorFormat BACnetContextTagUnsignedInteger
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetAuthenticationFactorFormat) GetFormatType() BACnetAuthenticationFactorTypeTagged {
	return m.FormatType
}

func (m *_BACnetAuthenticationFactorFormat) GetVendorId() BACnetVendorIdTagged {
	return m.VendorId
}

func (m *_BACnetAuthenticationFactorFormat) GetVendorFormat() BACnetContextTagUnsignedInteger {
	return m.VendorFormat
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetAuthenticationFactorFormat factory function for _BACnetAuthenticationFactorFormat
func NewBACnetAuthenticationFactorFormat( formatType BACnetAuthenticationFactorTypeTagged , vendorId BACnetVendorIdTagged , vendorFormat BACnetContextTagUnsignedInteger ) *_BACnetAuthenticationFactorFormat {
return &_BACnetAuthenticationFactorFormat{ FormatType: formatType , VendorId: vendorId , VendorFormat: vendorFormat }
}

// Deprecated: use the interface for direct cast
func CastBACnetAuthenticationFactorFormat(structType interface{}) BACnetAuthenticationFactorFormat {
    if casted, ok := structType.(BACnetAuthenticationFactorFormat); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetAuthenticationFactorFormat); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetAuthenticationFactorFormat) GetTypeName() string {
	return "BACnetAuthenticationFactorFormat"
}

func (m *_BACnetAuthenticationFactorFormat) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetAuthenticationFactorFormat) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (formatType)
	lengthInBits += m.FormatType.GetLengthInBits()

	// Optional Field (vendorId)
	if m.VendorId != nil {
		lengthInBits += m.VendorId.GetLengthInBits()
	}

	// Optional Field (vendorFormat)
	if m.VendorFormat != nil {
		lengthInBits += m.VendorFormat.GetLengthInBits()
	}

	return lengthInBits
}


func (m *_BACnetAuthenticationFactorFormat) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetAuthenticationFactorFormatParse(readBuffer utils.ReadBuffer) (BACnetAuthenticationFactorFormat, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetAuthenticationFactorFormat"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetAuthenticationFactorFormat")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (formatType)
	if pullErr := readBuffer.PullContext("formatType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for formatType")
	}
_formatType, _formatTypeErr := BACnetAuthenticationFactorTypeTaggedParse(readBuffer , uint8( uint8(0) ) , TagClass( TagClass_CONTEXT_SPECIFIC_TAGS ) )
	if _formatTypeErr != nil {
		return nil, errors.Wrap(_formatTypeErr, "Error parsing 'formatType' field of BACnetAuthenticationFactorFormat")
	}
	formatType := _formatType.(BACnetAuthenticationFactorTypeTagged)
	if closeErr := readBuffer.CloseContext("formatType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for formatType")
	}

	// Optional Field (vendorId) (Can be skipped, if a given expression evaluates to false)
	var vendorId BACnetVendorIdTagged = nil
{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("vendorId"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for vendorId")
		}
_val, _err := BACnetVendorIdTaggedParse(readBuffer , uint8(1) , TagClass_CONTEXT_SPECIFIC_TAGS )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'vendorId' field of BACnetAuthenticationFactorFormat")
		default:
			vendorId = _val.(BACnetVendorIdTagged)
			if closeErr := readBuffer.CloseContext("vendorId"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for vendorId")
			}
		}
	}

	// Optional Field (vendorFormat) (Can be skipped, if a given expression evaluates to false)
	var vendorFormat BACnetContextTagUnsignedInteger = nil
{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("vendorFormat"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for vendorFormat")
		}
_val, _err := BACnetContextTagParse(readBuffer , uint8(2) , BACnetDataType_UNSIGNED_INTEGER )
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'vendorFormat' field of BACnetAuthenticationFactorFormat")
		default:
			vendorFormat = _val.(BACnetContextTagUnsignedInteger)
			if closeErr := readBuffer.CloseContext("vendorFormat"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for vendorFormat")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetAuthenticationFactorFormat"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetAuthenticationFactorFormat")
	}

	// Create the instance
	return &_BACnetAuthenticationFactorFormat{
			FormatType: formatType,
			VendorId: vendorId,
			VendorFormat: vendorFormat,
		}, nil
}

func (m *_BACnetAuthenticationFactorFormat) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetAuthenticationFactorFormat) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetAuthenticationFactorFormat"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetAuthenticationFactorFormat")
	}

	// Simple Field (formatType)
	if pushErr := writeBuffer.PushContext("formatType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for formatType")
	}
	_formatTypeErr := writeBuffer.WriteSerializable(m.GetFormatType())
	if popErr := writeBuffer.PopContext("formatType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for formatType")
	}
	if _formatTypeErr != nil {
		return errors.Wrap(_formatTypeErr, "Error serializing 'formatType' field")
	}

	// Optional Field (vendorId) (Can be skipped, if the value is null)
	var vendorId BACnetVendorIdTagged = nil
	if m.GetVendorId() != nil {
		if pushErr := writeBuffer.PushContext("vendorId"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vendorId")
		}
		vendorId = m.GetVendorId()
		_vendorIdErr := writeBuffer.WriteSerializable(vendorId)
		if popErr := writeBuffer.PopContext("vendorId"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vendorId")
		}
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}
	}

	// Optional Field (vendorFormat) (Can be skipped, if the value is null)
	var vendorFormat BACnetContextTagUnsignedInteger = nil
	if m.GetVendorFormat() != nil {
		if pushErr := writeBuffer.PushContext("vendorFormat"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for vendorFormat")
		}
		vendorFormat = m.GetVendorFormat()
		_vendorFormatErr := writeBuffer.WriteSerializable(vendorFormat)
		if popErr := writeBuffer.PopContext("vendorFormat"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for vendorFormat")
		}
		if _vendorFormatErr != nil {
			return errors.Wrap(_vendorFormatErr, "Error serializing 'vendorFormat' field")
		}
	}

	if popErr := writeBuffer.PopContext("BACnetAuthenticationFactorFormat"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetAuthenticationFactorFormat")
	}
	return nil
}


func (m *_BACnetAuthenticationFactorFormat) isBACnetAuthenticationFactorFormat() bool {
	return true
}

func (m *_BACnetAuthenticationFactorFormat) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



