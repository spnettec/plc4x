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


// BACnetContextTagBoolean is the corresponding interface of BACnetContextTagBoolean
type BACnetContextTagBoolean interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetContextTag
	// GetValue returns Value (property field)
	GetValue() uint8
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadBoolean
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() bool
}

// BACnetContextTagBooleanExactly can be used when we want exactly this type and not a type which fulfills BACnetContextTagBoolean.
// This is useful for switch cases.
type BACnetContextTagBooleanExactly interface {
	BACnetContextTagBoolean
	isBACnetContextTagBoolean() bool
}

// _BACnetContextTagBoolean is the data-structure of this message
type _BACnetContextTagBoolean struct {
	*_BACnetContextTag
        Value uint8
        Payload BACnetTagPayloadBoolean
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetContextTagBoolean)  GetDataType() BACnetDataType {
return BACnetDataType_BOOLEAN}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetContextTagBoolean) InitializeParent(parent BACnetContextTag , header BACnetTagHeader ) {	m.Header = header
}

func (m *_BACnetContextTagBoolean)  GetParent() BACnetContextTag {
	return m._BACnetContextTag
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetContextTagBoolean) GetValue() uint8 {
	return m.Value
}

func (m *_BACnetContextTagBoolean) GetPayload() BACnetTagPayloadBoolean {
	return m.Payload
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetContextTagBoolean) GetActualValue() bool {
	ctx := context.Background()
	_ = ctx
	return bool(m.GetPayload().GetValue())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetContextTagBoolean factory function for _BACnetContextTagBoolean
func NewBACnetContextTagBoolean( value uint8 , payload BACnetTagPayloadBoolean , header BACnetTagHeader , tagNumberArgument uint8 ) *_BACnetContextTagBoolean {
	_result := &_BACnetContextTagBoolean{
		Value: value,
		Payload: payload,
    	_BACnetContextTag: NewBACnetContextTag(header, tagNumberArgument),
	}
	_result._BACnetContextTag._BACnetContextTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetContextTagBoolean(structType any) BACnetContextTagBoolean {
    if casted, ok := structType.(BACnetContextTagBoolean); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetContextTagBoolean); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetContextTagBoolean) GetTypeName() string {
	return "BACnetContextTagBoolean"
}

func (m *_BACnetContextTagBoolean) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (value)
	lengthInBits += 8;

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetContextTagBoolean) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetContextTagBooleanParse(ctx context.Context, theBytes []byte, header BACnetTagHeader, tagNumberArgument uint8, dataType BACnetDataType) (BACnetContextTagBoolean, error) {
	return BACnetContextTagBooleanParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), header, tagNumberArgument, dataType)
}

func BACnetContextTagBooleanParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, header BACnetTagHeader, tagNumberArgument uint8, dataType BACnetDataType) (BACnetContextTagBoolean, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetContextTagBoolean"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetContextTagBoolean")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Validation
	if (!(bool((header.GetActualLength()) == ((1))))) {
		return nil, errors.WithStack(utils.ParseValidationError{"length field should be 1"})
	}

	// Simple Field (value)
_value, _valueErr := readBuffer.ReadUint8("value", 8)
	if _valueErr != nil {
		return nil, errors.Wrap(_valueErr, "Error parsing 'value' field of BACnetContextTagBoolean")
	}
	value := _value

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
_payload, _payloadErr := BACnetTagPayloadBooleanParseWithBuffer(ctx, readBuffer , uint32( value ) )
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field of BACnetContextTagBoolean")
	}
	payload := _payload.(BACnetTagPayloadBoolean)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_actualValue := payload.GetValue()
	actualValue := bool(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetContextTagBoolean"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetContextTagBoolean")
	}

	// Create a partially initialized instance
	_child := &_BACnetContextTagBoolean{
		_BACnetContextTag: &_BACnetContextTag{
			TagNumberArgument: tagNumberArgument,
		},
		Value: value,
		Payload: payload,
	}
	_child._BACnetContextTag._BACnetContextTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetContextTagBoolean) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetContextTagBoolean) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetContextTagBoolean"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetContextTagBoolean")
		}

	// Simple Field (value)
	value := uint8(m.GetValue())
	_valueErr := writeBuffer.WriteUint8("value", 8, (value))
	if _valueErr != nil {
		return errors.Wrap(_valueErr, "Error serializing 'value' field")
	}

	// Simple Field (payload)
	if pushErr := writeBuffer.PushContext("payload"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for payload")
	}
	_payloadErr := writeBuffer.WriteSerializable(ctx, m.GetPayload())
	if popErr := writeBuffer.PopContext("payload"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for payload")
	}
	if _payloadErr != nil {
		return errors.Wrap(_payloadErr, "Error serializing 'payload' field")
	}
	// Virtual field
	if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
		return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetContextTagBoolean"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetContextTagBoolean")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetContextTagBoolean) isBACnetContextTagBoolean() bool {
	return true
}

func (m *_BACnetContextTagBoolean) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



