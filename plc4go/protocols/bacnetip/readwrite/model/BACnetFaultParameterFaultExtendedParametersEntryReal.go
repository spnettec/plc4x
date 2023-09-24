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


// BACnetFaultParameterFaultExtendedParametersEntryReal is the corresponding interface of BACnetFaultParameterFaultExtendedParametersEntryReal
type BACnetFaultParameterFaultExtendedParametersEntryReal interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetFaultParameterFaultExtendedParametersEntry
	// GetRealValue returns RealValue (property field)
	GetRealValue() BACnetApplicationTagReal
}

// BACnetFaultParameterFaultExtendedParametersEntryRealExactly can be used when we want exactly this type and not a type which fulfills BACnetFaultParameterFaultExtendedParametersEntryReal.
// This is useful for switch cases.
type BACnetFaultParameterFaultExtendedParametersEntryRealExactly interface {
	BACnetFaultParameterFaultExtendedParametersEntryReal
	isBACnetFaultParameterFaultExtendedParametersEntryReal() bool
}

// _BACnetFaultParameterFaultExtendedParametersEntryReal is the data-structure of this message
type _BACnetFaultParameterFaultExtendedParametersEntryReal struct {
	*_BACnetFaultParameterFaultExtendedParametersEntry
        RealValue BACnetApplicationTagReal
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) InitializeParent(parent BACnetFaultParameterFaultExtendedParametersEntry , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal)  GetParent() BACnetFaultParameterFaultExtendedParametersEntry {
	return m._BACnetFaultParameterFaultExtendedParametersEntry
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) GetRealValue() BACnetApplicationTagReal {
	return m.RealValue
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetFaultParameterFaultExtendedParametersEntryReal factory function for _BACnetFaultParameterFaultExtendedParametersEntryReal
func NewBACnetFaultParameterFaultExtendedParametersEntryReal( realValue BACnetApplicationTagReal , peekedTagHeader BACnetTagHeader ) *_BACnetFaultParameterFaultExtendedParametersEntryReal {
	_result := &_BACnetFaultParameterFaultExtendedParametersEntryReal{
		RealValue: realValue,
    	_BACnetFaultParameterFaultExtendedParametersEntry: NewBACnetFaultParameterFaultExtendedParametersEntry(peekedTagHeader),
	}
	_result._BACnetFaultParameterFaultExtendedParametersEntry._BACnetFaultParameterFaultExtendedParametersEntryChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetFaultParameterFaultExtendedParametersEntryReal(structType any) BACnetFaultParameterFaultExtendedParametersEntryReal {
    if casted, ok := structType.(BACnetFaultParameterFaultExtendedParametersEntryReal); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetFaultParameterFaultExtendedParametersEntryReal); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) GetTypeName() string {
	return "BACnetFaultParameterFaultExtendedParametersEntryReal"
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (realValue)
	lengthInBits += m.RealValue.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetFaultParameterFaultExtendedParametersEntryRealParse(ctx context.Context, theBytes []byte) (BACnetFaultParameterFaultExtendedParametersEntryReal, error) {
	return BACnetFaultParameterFaultExtendedParametersEntryRealParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetFaultParameterFaultExtendedParametersEntryRealParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetFaultParameterFaultExtendedParametersEntryReal, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetFaultParameterFaultExtendedParametersEntryReal"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetFaultParameterFaultExtendedParametersEntryReal")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (realValue)
	if pullErr := readBuffer.PullContext("realValue"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for realValue")
	}
_realValue, _realValueErr := BACnetApplicationTagParseWithBuffer(ctx, readBuffer)
	if _realValueErr != nil {
		return nil, errors.Wrap(_realValueErr, "Error parsing 'realValue' field of BACnetFaultParameterFaultExtendedParametersEntryReal")
	}
	realValue := _realValue.(BACnetApplicationTagReal)
	if closeErr := readBuffer.CloseContext("realValue"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for realValue")
	}

	if closeErr := readBuffer.CloseContext("BACnetFaultParameterFaultExtendedParametersEntryReal"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetFaultParameterFaultExtendedParametersEntryReal")
	}

	// Create a partially initialized instance
	_child := &_BACnetFaultParameterFaultExtendedParametersEntryReal{
		_BACnetFaultParameterFaultExtendedParametersEntry: &_BACnetFaultParameterFaultExtendedParametersEntry{
		},
		RealValue: realValue,
	}
	_child._BACnetFaultParameterFaultExtendedParametersEntry._BACnetFaultParameterFaultExtendedParametersEntryChildRequirements = _child
	return _child, nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetFaultParameterFaultExtendedParametersEntryReal"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetFaultParameterFaultExtendedParametersEntryReal")
		}

	// Simple Field (realValue)
	if pushErr := writeBuffer.PushContext("realValue"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for realValue")
	}
	_realValueErr := writeBuffer.WriteSerializable(ctx, m.GetRealValue())
	if popErr := writeBuffer.PopContext("realValue"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for realValue")
	}
	if _realValueErr != nil {
		return errors.Wrap(_realValueErr, "Error serializing 'realValue' field")
	}

		if popErr := writeBuffer.PopContext("BACnetFaultParameterFaultExtendedParametersEntryReal"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetFaultParameterFaultExtendedParametersEntryReal")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) isBACnetFaultParameterFaultExtendedParametersEntryReal() bool {
	return true
}

func (m *_BACnetFaultParameterFaultExtendedParametersEntryReal) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



