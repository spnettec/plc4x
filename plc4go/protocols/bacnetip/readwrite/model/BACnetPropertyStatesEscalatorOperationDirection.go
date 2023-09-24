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


// BACnetPropertyStatesEscalatorOperationDirection is the corresponding interface of BACnetPropertyStatesEscalatorOperationDirection
type BACnetPropertyStatesEscalatorOperationDirection interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetEscalatorOperationDirection returns EscalatorOperationDirection (property field)
	GetEscalatorOperationDirection() BACnetEscalatorOperationDirectionTagged
}

// BACnetPropertyStatesEscalatorOperationDirectionExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesEscalatorOperationDirection.
// This is useful for switch cases.
type BACnetPropertyStatesEscalatorOperationDirectionExactly interface {
	BACnetPropertyStatesEscalatorOperationDirection
	isBACnetPropertyStatesEscalatorOperationDirection() bool
}

// _BACnetPropertyStatesEscalatorOperationDirection is the data-structure of this message
type _BACnetPropertyStatesEscalatorOperationDirection struct {
	*_BACnetPropertyStates
        EscalatorOperationDirection BACnetEscalatorOperationDirectionTagged
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesEscalatorOperationDirection) InitializeParent(parent BACnetPropertyStates , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesEscalatorOperationDirection)  GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesEscalatorOperationDirection) GetEscalatorOperationDirection() BACnetEscalatorOperationDirectionTagged {
	return m.EscalatorOperationDirection
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetPropertyStatesEscalatorOperationDirection factory function for _BACnetPropertyStatesEscalatorOperationDirection
func NewBACnetPropertyStatesEscalatorOperationDirection( escalatorOperationDirection BACnetEscalatorOperationDirectionTagged , peekedTagHeader BACnetTagHeader ) *_BACnetPropertyStatesEscalatorOperationDirection {
	_result := &_BACnetPropertyStatesEscalatorOperationDirection{
		EscalatorOperationDirection: escalatorOperationDirection,
    	_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesEscalatorOperationDirection(structType any) BACnetPropertyStatesEscalatorOperationDirection {
    if casted, ok := structType.(BACnetPropertyStatesEscalatorOperationDirection); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesEscalatorOperationDirection); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesEscalatorOperationDirection) GetTypeName() string {
	return "BACnetPropertyStatesEscalatorOperationDirection"
}

func (m *_BACnetPropertyStatesEscalatorOperationDirection) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (escalatorOperationDirection)
	lengthInBits += m.EscalatorOperationDirection.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetPropertyStatesEscalatorOperationDirection) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesEscalatorOperationDirectionParse(ctx context.Context, theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesEscalatorOperationDirection, error) {
	return BACnetPropertyStatesEscalatorOperationDirectionParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesEscalatorOperationDirectionParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesEscalatorOperationDirection, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesEscalatorOperationDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesEscalatorOperationDirection")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (escalatorOperationDirection)
	if pullErr := readBuffer.PullContext("escalatorOperationDirection"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for escalatorOperationDirection")
	}
_escalatorOperationDirection, _escalatorOperationDirectionErr := BACnetEscalatorOperationDirectionTaggedParseWithBuffer(ctx, readBuffer , uint8( peekedTagNumber ) , TagClass( TagClass_CONTEXT_SPECIFIC_TAGS ) )
	if _escalatorOperationDirectionErr != nil {
		return nil, errors.Wrap(_escalatorOperationDirectionErr, "Error parsing 'escalatorOperationDirection' field of BACnetPropertyStatesEscalatorOperationDirection")
	}
	escalatorOperationDirection := _escalatorOperationDirection.(BACnetEscalatorOperationDirectionTagged)
	if closeErr := readBuffer.CloseContext("escalatorOperationDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for escalatorOperationDirection")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesEscalatorOperationDirection"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesEscalatorOperationDirection")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesEscalatorOperationDirection{
		_BACnetPropertyStates: &_BACnetPropertyStates{
		},
		EscalatorOperationDirection: escalatorOperationDirection,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesEscalatorOperationDirection) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesEscalatorOperationDirection) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesEscalatorOperationDirection"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesEscalatorOperationDirection")
		}

	// Simple Field (escalatorOperationDirection)
	if pushErr := writeBuffer.PushContext("escalatorOperationDirection"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for escalatorOperationDirection")
	}
	_escalatorOperationDirectionErr := writeBuffer.WriteSerializable(ctx, m.GetEscalatorOperationDirection())
	if popErr := writeBuffer.PopContext("escalatorOperationDirection"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for escalatorOperationDirection")
	}
	if _escalatorOperationDirectionErr != nil {
		return errors.Wrap(_escalatorOperationDirectionErr, "Error serializing 'escalatorOperationDirection' field")
	}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesEscalatorOperationDirection"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesEscalatorOperationDirection")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetPropertyStatesEscalatorOperationDirection) isBACnetPropertyStatesEscalatorOperationDirection() bool {
	return true
}

func (m *_BACnetPropertyStatesEscalatorOperationDirection) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



