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


// MPropReadReq is the corresponding interface of MPropReadReq
type MPropReadReq interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	CEMI
	// GetInterfaceObjectType returns InterfaceObjectType (property field)
	GetInterfaceObjectType() uint16
	// GetObjectInstance returns ObjectInstance (property field)
	GetObjectInstance() uint8
	// GetPropertyId returns PropertyId (property field)
	GetPropertyId() uint8
	// GetNumberOfElements returns NumberOfElements (property field)
	GetNumberOfElements() uint8
	// GetStartIndex returns StartIndex (property field)
	GetStartIndex() uint16
}

// MPropReadReqExactly can be used when we want exactly this type and not a type which fulfills MPropReadReq.
// This is useful for switch cases.
type MPropReadReqExactly interface {
	MPropReadReq
	isMPropReadReq() bool
}

// _MPropReadReq is the data-structure of this message
type _MPropReadReq struct {
	*_CEMI
        InterfaceObjectType uint16
        ObjectInstance uint8
        PropertyId uint8
        NumberOfElements uint8
        StartIndex uint16
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_MPropReadReq)  GetMessageCode() uint8 {
return 0xFC}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_MPropReadReq) InitializeParent(parent CEMI ) {}

func (m *_MPropReadReq)  GetParent() CEMI {
	return m._CEMI
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_MPropReadReq) GetInterfaceObjectType() uint16 {
	return m.InterfaceObjectType
}

func (m *_MPropReadReq) GetObjectInstance() uint8 {
	return m.ObjectInstance
}

func (m *_MPropReadReq) GetPropertyId() uint8 {
	return m.PropertyId
}

func (m *_MPropReadReq) GetNumberOfElements() uint8 {
	return m.NumberOfElements
}

func (m *_MPropReadReq) GetStartIndex() uint16 {
	return m.StartIndex
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewMPropReadReq factory function for _MPropReadReq
func NewMPropReadReq( interfaceObjectType uint16 , objectInstance uint8 , propertyId uint8 , numberOfElements uint8 , startIndex uint16 , size uint16 ) *_MPropReadReq {
	_result := &_MPropReadReq{
		InterfaceObjectType: interfaceObjectType,
		ObjectInstance: objectInstance,
		PropertyId: propertyId,
		NumberOfElements: numberOfElements,
		StartIndex: startIndex,
    	_CEMI: NewCEMI(size),
	}
	_result._CEMI._CEMIChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastMPropReadReq(structType any) MPropReadReq {
    if casted, ok := structType.(MPropReadReq); ok {
		return casted
	}
	if casted, ok := structType.(*MPropReadReq); ok {
		return *casted
	}
	return nil
}

func (m *_MPropReadReq) GetTypeName() string {
	return "MPropReadReq"
}

func (m *_MPropReadReq) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (interfaceObjectType)
	lengthInBits += 16;

	// Simple field (objectInstance)
	lengthInBits += 8;

	// Simple field (propertyId)
	lengthInBits += 8;

	// Simple field (numberOfElements)
	lengthInBits += 4;

	// Simple field (startIndex)
	lengthInBits += 12;

	return lengthInBits
}


func (m *_MPropReadReq) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func MPropReadReqParse(ctx context.Context, theBytes []byte, size uint16) (MPropReadReq, error) {
	return MPropReadReqParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), size)
}

func MPropReadReqParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, size uint16) (MPropReadReq, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("MPropReadReq"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for MPropReadReq")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (interfaceObjectType)
_interfaceObjectType, _interfaceObjectTypeErr := readBuffer.ReadUint16("interfaceObjectType", 16)
	if _interfaceObjectTypeErr != nil {
		return nil, errors.Wrap(_interfaceObjectTypeErr, "Error parsing 'interfaceObjectType' field of MPropReadReq")
	}
	interfaceObjectType := _interfaceObjectType

	// Simple Field (objectInstance)
_objectInstance, _objectInstanceErr := readBuffer.ReadUint8("objectInstance", 8)
	if _objectInstanceErr != nil {
		return nil, errors.Wrap(_objectInstanceErr, "Error parsing 'objectInstance' field of MPropReadReq")
	}
	objectInstance := _objectInstance

	// Simple Field (propertyId)
_propertyId, _propertyIdErr := readBuffer.ReadUint8("propertyId", 8)
	if _propertyIdErr != nil {
		return nil, errors.Wrap(_propertyIdErr, "Error parsing 'propertyId' field of MPropReadReq")
	}
	propertyId := _propertyId

	// Simple Field (numberOfElements)
_numberOfElements, _numberOfElementsErr := readBuffer.ReadUint8("numberOfElements", 4)
	if _numberOfElementsErr != nil {
		return nil, errors.Wrap(_numberOfElementsErr, "Error parsing 'numberOfElements' field of MPropReadReq")
	}
	numberOfElements := _numberOfElements

	// Simple Field (startIndex)
_startIndex, _startIndexErr := readBuffer.ReadUint16("startIndex", 12)
	if _startIndexErr != nil {
		return nil, errors.Wrap(_startIndexErr, "Error parsing 'startIndex' field of MPropReadReq")
	}
	startIndex := _startIndex

	if closeErr := readBuffer.CloseContext("MPropReadReq"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for MPropReadReq")
	}

	// Create a partially initialized instance
	_child := &_MPropReadReq{
		_CEMI: &_CEMI{
			Size: size,
		},
		InterfaceObjectType: interfaceObjectType,
		ObjectInstance: objectInstance,
		PropertyId: propertyId,
		NumberOfElements: numberOfElements,
		StartIndex: startIndex,
	}
	_child._CEMI._CEMIChildRequirements = _child
	return _child, nil
}

func (m *_MPropReadReq) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_MPropReadReq) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("MPropReadReq"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for MPropReadReq")
		}

	// Simple Field (interfaceObjectType)
	interfaceObjectType := uint16(m.GetInterfaceObjectType())
	_interfaceObjectTypeErr := writeBuffer.WriteUint16("interfaceObjectType", 16, (interfaceObjectType))
	if _interfaceObjectTypeErr != nil {
		return errors.Wrap(_interfaceObjectTypeErr, "Error serializing 'interfaceObjectType' field")
	}

	// Simple Field (objectInstance)
	objectInstance := uint8(m.GetObjectInstance())
	_objectInstanceErr := writeBuffer.WriteUint8("objectInstance", 8, (objectInstance))
	if _objectInstanceErr != nil {
		return errors.Wrap(_objectInstanceErr, "Error serializing 'objectInstance' field")
	}

	// Simple Field (propertyId)
	propertyId := uint8(m.GetPropertyId())
	_propertyIdErr := writeBuffer.WriteUint8("propertyId", 8, (propertyId))
	if _propertyIdErr != nil {
		return errors.Wrap(_propertyIdErr, "Error serializing 'propertyId' field")
	}

	// Simple Field (numberOfElements)
	numberOfElements := uint8(m.GetNumberOfElements())
	_numberOfElementsErr := writeBuffer.WriteUint8("numberOfElements", 4, (numberOfElements))
	if _numberOfElementsErr != nil {
		return errors.Wrap(_numberOfElementsErr, "Error serializing 'numberOfElements' field")
	}

	// Simple Field (startIndex)
	startIndex := uint16(m.GetStartIndex())
	_startIndexErr := writeBuffer.WriteUint16("startIndex", 12, (startIndex))
	if _startIndexErr != nil {
		return errors.Wrap(_startIndexErr, "Error serializing 'startIndex' field")
	}

		if popErr := writeBuffer.PopContext("MPropReadReq"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for MPropReadReq")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_MPropReadReq) isMPropReadReq() bool {
	return true
}

func (m *_MPropReadReq) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



