/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer is the data-structure of this message
type BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer struct {
	*BACnetUnconfirmedServiceRequest
	VendorId          *BACnetContextTagUnsignedInteger
	ServiceNumber     *BACnetContextTagUnsignedInteger
	ServiceParameters *BACnetPropertyValues

	// Arguments.
	Len uint16
}

// IBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer is the corresponding interface of BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
type IBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer interface {
	IBACnetUnconfirmedServiceRequest
	// GetVendorId returns VendorId (property field)
	GetVendorId() *BACnetContextTagUnsignedInteger
	// GetServiceNumber returns ServiceNumber (property field)
	GetServiceNumber() *BACnetContextTagUnsignedInteger
	// GetServiceParameters returns ServiceParameters (property field)
	GetServiceParameters() *BACnetPropertyValues
	// GetLengthInBytes returns the length in bytes
	GetLengthInBytes() uint16
	// GetLengthInBits returns the length in bits
	GetLengthInBits() uint16
	// Serialize serializes this type
	Serialize(writeBuffer utils.WriteBuffer) error
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetServiceChoice() uint8 {
	return 0x04
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) InitializeParent(parent *BACnetUnconfirmedServiceRequest) {
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetParent() *BACnetUnconfirmedServiceRequest {
	return m.BACnetUnconfirmedServiceRequest
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetVendorId() *BACnetContextTagUnsignedInteger {
	return m.VendorId
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetServiceNumber() *BACnetContextTagUnsignedInteger {
	return m.ServiceNumber
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetServiceParameters() *BACnetPropertyValues {
	return m.ServiceParameters
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer factory function for BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer
func NewBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(vendorId *BACnetContextTagUnsignedInteger, serviceNumber *BACnetContextTagUnsignedInteger, serviceParameters *BACnetPropertyValues, len uint16) *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer {
	_result := &BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer{
		VendorId:                        vendorId,
		ServiceNumber:                   serviceNumber,
		ServiceParameters:               serviceParameters,
		BACnetUnconfirmedServiceRequest: NewBACnetUnconfirmedServiceRequest(len),
	}
	_result.Child = _result
	return _result
}

func CastBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(structType interface{}) *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer {
	if casted, ok := structType.(BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer); ok {
		return &casted
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer); ok {
		return casted
	}
	if casted, ok := structType.(BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(casted.Child)
	}
	if casted, ok := structType.(*BACnetUnconfirmedServiceRequest); ok {
		return CastBACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer(casted.Child)
	}
	return nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetTypeName() string {
	return "BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (vendorId)
	lengthInBits += m.VendorId.GetLengthInBits()

	// Simple field (serviceNumber)
	lengthInBits += m.ServiceNumber.GetLengthInBits()

	// Optional Field (serviceParameters)
	if m.ServiceParameters != nil {
		lengthInBits += (*m.ServiceParameters).GetLengthInBits()
	}

	return lengthInBits
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransferParse(readBuffer utils.ReadBuffer, len uint16) (*BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer, error) {
	if pullErr := readBuffer.PullContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"); pullErr != nil {
		return nil, pullErr
	}
	currentPos := readBuffer.GetPos()
	_ = currentPos

	// Simple Field (vendorId)
	if pullErr := readBuffer.PullContext("vendorId"); pullErr != nil {
		return nil, pullErr
	}
	_vendorId, _vendorIdErr := BACnetContextTagParse(readBuffer, uint8(uint8(1)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _vendorIdErr != nil {
		return nil, errors.Wrap(_vendorIdErr, "Error parsing 'vendorId' field")
	}
	vendorId := CastBACnetContextTagUnsignedInteger(_vendorId)
	if closeErr := readBuffer.CloseContext("vendorId"); closeErr != nil {
		return nil, closeErr
	}

	// Simple Field (serviceNumber)
	if pullErr := readBuffer.PullContext("serviceNumber"); pullErr != nil {
		return nil, pullErr
	}
	_serviceNumber, _serviceNumberErr := BACnetContextTagParse(readBuffer, uint8(uint8(2)), BACnetDataType(BACnetDataType_UNSIGNED_INTEGER))
	if _serviceNumberErr != nil {
		return nil, errors.Wrap(_serviceNumberErr, "Error parsing 'serviceNumber' field")
	}
	serviceNumber := CastBACnetContextTagUnsignedInteger(_serviceNumber)
	if closeErr := readBuffer.CloseContext("serviceNumber"); closeErr != nil {
		return nil, closeErr
	}

	// Optional Field (serviceParameters) (Can be skipped, if a given expression evaluates to false)
	var serviceParameters *BACnetPropertyValues = nil
	{
		currentPos = readBuffer.GetPos()
		if pullErr := readBuffer.PullContext("serviceParameters"); pullErr != nil {
			return nil, pullErr
		}
		_val, _err := BACnetPropertyValuesParse(readBuffer, uint8(2), BACnetObjectType_VENDOR_PROPRIETARY_VALUE)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'serviceParameters' field")
		default:
			serviceParameters = CastBACnetPropertyValues(_val)
			if closeErr := readBuffer.CloseContext("serviceParameters"); closeErr != nil {
				return nil, closeErr
			}
		}
	}

	if closeErr := readBuffer.CloseContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"); closeErr != nil {
		return nil, closeErr
	}

	// Create a partially initialized instance
	_child := &BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer{
		VendorId:                        CastBACnetContextTagUnsignedInteger(vendorId),
		ServiceNumber:                   CastBACnetContextTagUnsignedInteger(serviceNumber),
		ServiceParameters:               CastBACnetPropertyValues(serviceParameters),
		BACnetUnconfirmedServiceRequest: &BACnetUnconfirmedServiceRequest{},
	}
	_child.BACnetUnconfirmedServiceRequest.Child = _child
	return _child, nil
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) Serialize(writeBuffer utils.WriteBuffer) error {
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"); pushErr != nil {
			return pushErr
		}

		// Simple Field (vendorId)
		if pushErr := writeBuffer.PushContext("vendorId"); pushErr != nil {
			return pushErr
		}
		_vendorIdErr := m.VendorId.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("vendorId"); popErr != nil {
			return popErr
		}
		if _vendorIdErr != nil {
			return errors.Wrap(_vendorIdErr, "Error serializing 'vendorId' field")
		}

		// Simple Field (serviceNumber)
		if pushErr := writeBuffer.PushContext("serviceNumber"); pushErr != nil {
			return pushErr
		}
		_serviceNumberErr := m.ServiceNumber.Serialize(writeBuffer)
		if popErr := writeBuffer.PopContext("serviceNumber"); popErr != nil {
			return popErr
		}
		if _serviceNumberErr != nil {
			return errors.Wrap(_serviceNumberErr, "Error serializing 'serviceNumber' field")
		}

		// Optional Field (serviceParameters) (Can be skipped, if the value is null)
		var serviceParameters *BACnetPropertyValues = nil
		if m.ServiceParameters != nil {
			if pushErr := writeBuffer.PushContext("serviceParameters"); pushErr != nil {
				return pushErr
			}
			serviceParameters = m.ServiceParameters
			_serviceParametersErr := serviceParameters.Serialize(writeBuffer)
			if popErr := writeBuffer.PopContext("serviceParameters"); popErr != nil {
				return popErr
			}
			if _serviceParametersErr != nil {
				return errors.Wrap(_serviceParametersErr, "Error serializing 'serviceParameters' field")
			}
		}

		if popErr := writeBuffer.PopContext("BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer"); popErr != nil {
			return popErr
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}

func (m *BACnetUnconfirmedServiceRequestUnconfirmedPrivateTransfer) String() string {
	if m == nil {
		return "<nil>"
	}
	buffer := utils.NewBoxedWriteBufferWithOptions(true, true)
	if err := m.Serialize(buffer); err != nil {
		return err.Error()
	}
	return buffer.GetBox().String()
}
