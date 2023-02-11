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
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetApplicationTagReal is the corresponding interface of BACnetApplicationTagReal
type BACnetApplicationTagReal interface {
	utils.LengthAware
	utils.Serializable
	BACnetApplicationTag
	// GetPayload returns Payload (property field)
	GetPayload() BACnetTagPayloadReal
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() float32
}

// BACnetApplicationTagRealExactly can be used when we want exactly this type and not a type which fulfills BACnetApplicationTagReal.
// This is useful for switch cases.
type BACnetApplicationTagRealExactly interface {
	BACnetApplicationTagReal
	isBACnetApplicationTagReal() bool
}

// _BACnetApplicationTagReal is the data-structure of this message
type _BACnetApplicationTagReal struct {
	*_BACnetApplicationTag
        Payload BACnetTagPayloadReal
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetApplicationTagReal) InitializeParent(parent BACnetApplicationTag , header BACnetTagHeader ) {	m.Header = header
}

func (m *_BACnetApplicationTagReal)  GetParent() BACnetApplicationTag {
	return m._BACnetApplicationTag
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetApplicationTagReal) GetPayload() BACnetTagPayloadReal {
	return m.Payload
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetApplicationTagReal) GetActualValue() float32 {
	ctx := context.Background()
	_ = ctx
	return float32(m.GetPayload().GetValue())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetApplicationTagReal factory function for _BACnetApplicationTagReal
func NewBACnetApplicationTagReal( payload BACnetTagPayloadReal , header BACnetTagHeader ) *_BACnetApplicationTagReal {
	_result := &_BACnetApplicationTagReal{
		Payload: payload,
    	_BACnetApplicationTag: NewBACnetApplicationTag(header),
	}
	_result._BACnetApplicationTag._BACnetApplicationTagChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetApplicationTagReal(structType interface{}) BACnetApplicationTagReal {
    if casted, ok := structType.(BACnetApplicationTagReal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetApplicationTagReal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetApplicationTagReal) GetTypeName() string {
	return "BACnetApplicationTagReal"
}

func (m *_BACnetApplicationTagReal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (payload)
	lengthInBits += m.Payload.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_BACnetApplicationTagReal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetApplicationTagRealParse(theBytes []byte) (BACnetApplicationTagReal, error) {
	return BACnetApplicationTagRealParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetApplicationTagRealParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetApplicationTagReal, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetApplicationTagReal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetApplicationTagReal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (payload)
	if pullErr := readBuffer.PullContext("payload"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for payload")
	}
_payload, _payloadErr := BACnetTagPayloadRealParseWithBuffer(ctx, readBuffer)
	if _payloadErr != nil {
		return nil, errors.Wrap(_payloadErr, "Error parsing 'payload' field of BACnetApplicationTagReal")
	}
	payload := _payload.(BACnetTagPayloadReal)
	if closeErr := readBuffer.CloseContext("payload"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for payload")
	}

	// Virtual field
	_actualValue := payload.GetValue()
	actualValue := float32(_actualValue)
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetApplicationTagReal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetApplicationTagReal")
	}

	// Create a partially initialized instance
	_child := &_BACnetApplicationTagReal{
		_BACnetApplicationTag: &_BACnetApplicationTag{
		},
		Payload: payload,
	}
	_child._BACnetApplicationTag._BACnetApplicationTagChildRequirements = _child
	return _child, nil
}

func (m *_BACnetApplicationTagReal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetApplicationTagReal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetApplicationTagReal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetApplicationTagReal")
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

		if popErr := writeBuffer.PopContext("BACnetApplicationTagReal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetApplicationTagReal")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetApplicationTagReal) isBACnetApplicationTagReal() bool {
	return true
}

func (m *_BACnetApplicationTagReal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



