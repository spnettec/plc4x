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
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetTagPayloadObjectIdentifier is the corresponding interface of BACnetTagPayloadObjectIdentifier
type BACnetTagPayloadObjectIdentifier interface {
	utils.LengthAware
	utils.Serializable
	// GetObjectType returns ObjectType (property field)
	GetObjectType() BACnetObjectType
	// GetProprietaryValue returns ProprietaryValue (property field)
	GetProprietaryValue() uint16
	// GetInstanceNumber returns InstanceNumber (property field)
	GetInstanceNumber() uint32
	// GetIsProprietary returns IsProprietary (virtual field)
	GetIsProprietary() bool
}

// BACnetTagPayloadObjectIdentifierExactly can be used when we want exactly this type and not a type which fulfills BACnetTagPayloadObjectIdentifier.
// This is useful for switch cases.
type BACnetTagPayloadObjectIdentifierExactly interface {
	BACnetTagPayloadObjectIdentifier
	isBACnetTagPayloadObjectIdentifier() bool
}

// _BACnetTagPayloadObjectIdentifier is the data-structure of this message
type _BACnetTagPayloadObjectIdentifier struct {
        ObjectType BACnetObjectType
        ProprietaryValue uint16
        InstanceNumber uint32
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetTagPayloadObjectIdentifier) GetObjectType() BACnetObjectType {
	return m.ObjectType
}

func (m *_BACnetTagPayloadObjectIdentifier) GetProprietaryValue() uint16 {
	return m.ProprietaryValue
}

func (m *_BACnetTagPayloadObjectIdentifier) GetInstanceNumber() uint32 {
	return m.InstanceNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetTagPayloadObjectIdentifier) GetIsProprietary() bool {
	return bool(bool((m.GetObjectType()) == (BACnetObjectType_VENDOR_PROPRIETARY_VALUE)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetTagPayloadObjectIdentifier factory function for _BACnetTagPayloadObjectIdentifier
func NewBACnetTagPayloadObjectIdentifier( objectType BACnetObjectType , proprietaryValue uint16 , instanceNumber uint32 ) *_BACnetTagPayloadObjectIdentifier {
return &_BACnetTagPayloadObjectIdentifier{ ObjectType: objectType , ProprietaryValue: proprietaryValue , InstanceNumber: instanceNumber }
}

// Deprecated: use the interface for direct cast
func CastBACnetTagPayloadObjectIdentifier(structType interface{}) BACnetTagPayloadObjectIdentifier {
    if casted, ok := structType.(BACnetTagPayloadObjectIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetTagPayloadObjectIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetTagPayloadObjectIdentifier) GetTypeName() string {
	return "BACnetTagPayloadObjectIdentifier"
}

func (m *_BACnetTagPayloadObjectIdentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetTagPayloadObjectIdentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Manual Field (objectType)
	lengthInBits += uint16(int32(10))

	// Manual Field (proprietaryValue)
	lengthInBits += uint16(int32(0))

	// A virtual field doesn't have any in- or output.

	// Simple field (instanceNumber)
	lengthInBits += 22;

	return lengthInBits
}


func (m *_BACnetTagPayloadObjectIdentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetTagPayloadObjectIdentifierParse(theBytes []byte) (BACnetTagPayloadObjectIdentifier, error) {
	return BACnetTagPayloadObjectIdentifierParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetTagPayloadObjectIdentifierParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetTagPayloadObjectIdentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetTagPayloadObjectIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetTagPayloadObjectIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Manual Field (objectType)
	_objectType, _objectTypeErr := ReadObjectType(readBuffer)
	if _objectTypeErr != nil {
		return nil, errors.Wrap(_objectTypeErr, "Error parsing 'objectType' field of BACnetTagPayloadObjectIdentifier")
	}
	var objectType BACnetObjectType
	if _objectType != nil {
            objectType = _objectType.(BACnetObjectType)
	}

	// Manual Field (proprietaryValue)
	_proprietaryValue, _proprietaryValueErr := ReadProprietaryObjectType(readBuffer, objectType)
	if _proprietaryValueErr != nil {
		return nil, errors.Wrap(_proprietaryValueErr, "Error parsing 'proprietaryValue' field of BACnetTagPayloadObjectIdentifier")
	}
	var proprietaryValue uint16
	if _proprietaryValue != nil {
            proprietaryValue = _proprietaryValue.(uint16)
	}

	// Virtual field
	_isProprietary := bool((objectType) == (BACnetObjectType_VENDOR_PROPRIETARY_VALUE))
	isProprietary := bool(_isProprietary)
	_ = isProprietary

	// Simple Field (instanceNumber)
_instanceNumber, _instanceNumberErr := readBuffer.ReadUint32("instanceNumber", 22)
	if _instanceNumberErr != nil {
		return nil, errors.Wrap(_instanceNumberErr, "Error parsing 'instanceNumber' field of BACnetTagPayloadObjectIdentifier")
	}
	instanceNumber := _instanceNumber

	if closeErr := readBuffer.CloseContext("BACnetTagPayloadObjectIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetTagPayloadObjectIdentifier")
	}

	// Create the instance
	return &_BACnetTagPayloadObjectIdentifier{
			ObjectType: objectType,
			ProprietaryValue: proprietaryValue,
			InstanceNumber: instanceNumber,
		}, nil
}

func (m *_BACnetTagPayloadObjectIdentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetTagPayloadObjectIdentifier) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetTagPayloadObjectIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetTagPayloadObjectIdentifier")
	}

	// Manual Field (objectType)
	_objectTypeErr := WriteObjectType(writeBuffer, m.GetObjectType())
	if _objectTypeErr != nil {
		return errors.Wrap(_objectTypeErr, "Error serializing 'objectType' field")
	}

	// Manual Field (proprietaryValue)
	_proprietaryValueErr := WriteProprietaryObjectType(writeBuffer, m.GetObjectType(), m.GetProprietaryValue())
	if _proprietaryValueErr != nil {
		return errors.Wrap(_proprietaryValueErr, "Error serializing 'proprietaryValue' field")
	}
	// Virtual field
	if _isProprietaryErr := writeBuffer.WriteVirtual("isProprietary", m.GetIsProprietary()); _isProprietaryErr != nil {
		return errors.Wrap(_isProprietaryErr, "Error serializing 'isProprietary' field")
	}

	// Simple Field (instanceNumber)
	instanceNumber := uint32(m.GetInstanceNumber())
	_instanceNumberErr := writeBuffer.WriteUint32("instanceNumber", 22, (instanceNumber))
	if _instanceNumberErr != nil {
		return errors.Wrap(_instanceNumberErr, "Error serializing 'instanceNumber' field")
	}

	if popErr := writeBuffer.PopContext("BACnetTagPayloadObjectIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetTagPayloadObjectIdentifier")
	}
	return nil
}


func (m *_BACnetTagPayloadObjectIdentifier) isBACnetTagPayloadObjectIdentifier() bool {
	return true
}

func (m *_BACnetTagPayloadObjectIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



