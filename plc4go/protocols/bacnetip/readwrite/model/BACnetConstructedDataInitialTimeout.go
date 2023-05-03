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
)

// Code generated by code-generation. DO NOT EDIT.

// BACnetConstructedDataInitialTimeout is the corresponding interface of BACnetConstructedDataInitialTimeout
type BACnetConstructedDataInitialTimeout interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetConstructedData
	// GetInitialTimeout returns InitialTimeout (property field)
	GetInitialTimeout() BACnetApplicationTagUnsignedInteger
	// GetActualValue returns ActualValue (virtual field)
	GetActualValue() BACnetApplicationTagUnsignedInteger
}

// BACnetConstructedDataInitialTimeoutExactly can be used when we want exactly this type and not a type which fulfills BACnetConstructedDataInitialTimeout.
// This is useful for switch cases.
type BACnetConstructedDataInitialTimeoutExactly interface {
	BACnetConstructedDataInitialTimeout
	isBACnetConstructedDataInitialTimeout() bool
}

// _BACnetConstructedDataInitialTimeout is the data-structure of this message
type _BACnetConstructedDataInitialTimeout struct {
	*_BACnetConstructedData
	InitialTimeout BACnetApplicationTagUnsignedInteger
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BACnetConstructedDataInitialTimeout) GetObjectTypeArgument() BACnetObjectType {
	return 0
}

func (m *_BACnetConstructedDataInitialTimeout) GetPropertyIdentifierArgument() BACnetPropertyIdentifier {
	return BACnetPropertyIdentifier_INITIAL_TIMEOUT
}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConstructedDataInitialTimeout) InitializeParent(parent BACnetConstructedData, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag) {
	m.OpeningTag = openingTag
	m.PeekedTagHeader = peekedTagHeader
	m.ClosingTag = closingTag
}

func (m *_BACnetConstructedDataInitialTimeout) GetParent() BACnetConstructedData {
	return m._BACnetConstructedData
}

///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConstructedDataInitialTimeout) GetInitialTimeout() BACnetApplicationTagUnsignedInteger {
	return m.InitialTimeout
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetConstructedDataInitialTimeout) GetActualValue() BACnetApplicationTagUnsignedInteger {
	ctx := context.Background()
	_ = ctx
	return CastBACnetApplicationTagUnsignedInteger(m.GetInitialTimeout())
}

///////////////////////-3
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

// NewBACnetConstructedDataInitialTimeout factory function for _BACnetConstructedDataInitialTimeout
func NewBACnetConstructedDataInitialTimeout(initialTimeout BACnetApplicationTagUnsignedInteger, openingTag BACnetOpeningTag, peekedTagHeader BACnetTagHeader, closingTag BACnetClosingTag, tagNumber uint8, arrayIndexArgument BACnetTagPayloadUnsignedInteger) *_BACnetConstructedDataInitialTimeout {
	_result := &_BACnetConstructedDataInitialTimeout{
		InitialTimeout:         initialTimeout,
		_BACnetConstructedData: NewBACnetConstructedData(openingTag, peekedTagHeader, closingTag, tagNumber, arrayIndexArgument),
	}
	_result._BACnetConstructedData._BACnetConstructedDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConstructedDataInitialTimeout(structType any) BACnetConstructedDataInitialTimeout {
	if casted, ok := structType.(BACnetConstructedDataInitialTimeout); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConstructedDataInitialTimeout); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConstructedDataInitialTimeout) GetTypeName() string {
	return "BACnetConstructedDataInitialTimeout"
}

func (m *_BACnetConstructedDataInitialTimeout) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (initialTimeout)
	lengthInBits += m.InitialTimeout.GetLengthInBits(ctx)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetConstructedDataInitialTimeout) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetConstructedDataInitialTimeoutParse(theBytes []byte, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataInitialTimeout, error) {
	return BACnetConstructedDataInitialTimeoutParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes), tagNumber, objectTypeArgument, propertyIdentifierArgument, arrayIndexArgument)
}

func BACnetConstructedDataInitialTimeoutParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, tagNumber uint8, objectTypeArgument BACnetObjectType, propertyIdentifierArgument BACnetPropertyIdentifier, arrayIndexArgument BACnetTagPayloadUnsignedInteger) (BACnetConstructedDataInitialTimeout, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConstructedDataInitialTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConstructedDataInitialTimeout")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (initialTimeout)
	if pullErr := readBuffer.PullContext("initialTimeout"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for initialTimeout")
	}
	_initialTimeout, _initialTimeoutErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _initialTimeoutErr != nil {
		return nil, errors.Wrap(_initialTimeoutErr, "Error parsing 'initialTimeout' field of BACnetConstructedDataInitialTimeout")
	}
	initialTimeout := _initialTimeout.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("initialTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for initialTimeout")
	}

	// Virtual field
	_actualValue := initialTimeout
	actualValue := _actualValue
	_ = actualValue

	if closeErr := readBuffer.CloseContext("BACnetConstructedDataInitialTimeout"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConstructedDataInitialTimeout")
	}

	// Create a partially initialized instance
	_child := &_BACnetConstructedDataInitialTimeout{
		_BACnetConstructedData: &_BACnetConstructedData{
			TagNumber:          tagNumber,
			ArrayIndexArgument: arrayIndexArgument,
		},
		InitialTimeout: initialTimeout,
	}
	_child._BACnetConstructedData._BACnetConstructedDataChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConstructedDataInitialTimeout) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetConstructedDataInitialTimeout) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConstructedDataInitialTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConstructedDataInitialTimeout")
		}

		// Simple Field (initialTimeout)
		if pushErr := writeBuffer.PushContext("initialTimeout"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for initialTimeout")
		}
		_initialTimeoutErr := writeBuffer.WriteSerializable(ctx, m.GetInitialTimeout())
		if popErr := writeBuffer.PopContext("initialTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for initialTimeout")
		}
		if _initialTimeoutErr != nil {
			return errors.Wrap(_initialTimeoutErr, "Error serializing 'initialTimeout' field")
		}
		// Virtual field
		if _actualValueErr := writeBuffer.WriteVirtual(ctx, "actualValue", m.GetActualValue()); _actualValueErr != nil {
			return errors.Wrap(_actualValueErr, "Error serializing 'actualValue' field")
		}

		if popErr := writeBuffer.PopContext("BACnetConstructedDataInitialTimeout"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConstructedDataInitialTimeout")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}

func (m *_BACnetConstructedDataInitialTimeout) isBACnetConstructedDataInitialTimeout() bool {
	return true
}

func (m *_BACnetConstructedDataInitialTimeout) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}
