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


// BACnetPropertyStatesProgramChange is the corresponding interface of BACnetPropertyStatesProgramChange
type BACnetPropertyStatesProgramChange interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetPropertyStates
	// GetProgramState returns ProgramState (property field)
	GetProgramState() BACnetProgramStateTagged
}

// BACnetPropertyStatesProgramChangeExactly can be used when we want exactly this type and not a type which fulfills BACnetPropertyStatesProgramChange.
// This is useful for switch cases.
type BACnetPropertyStatesProgramChangeExactly interface {
	BACnetPropertyStatesProgramChange
	isBACnetPropertyStatesProgramChange() bool
}

// _BACnetPropertyStatesProgramChange is the data-structure of this message
type _BACnetPropertyStatesProgramChange struct {
	*_BACnetPropertyStates
        ProgramState BACnetProgramStateTagged
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetPropertyStatesProgramChange) InitializeParent(parent BACnetPropertyStates , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetPropertyStatesProgramChange)  GetParent() BACnetPropertyStates {
	return m._BACnetPropertyStates
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetPropertyStatesProgramChange) GetProgramState() BACnetProgramStateTagged {
	return m.ProgramState
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetPropertyStatesProgramChange factory function for _BACnetPropertyStatesProgramChange
func NewBACnetPropertyStatesProgramChange( programState BACnetProgramStateTagged , peekedTagHeader BACnetTagHeader ) *_BACnetPropertyStatesProgramChange {
	_result := &_BACnetPropertyStatesProgramChange{
		ProgramState: programState,
    	_BACnetPropertyStates: NewBACnetPropertyStates(peekedTagHeader),
	}
	_result._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetPropertyStatesProgramChange(structType any) BACnetPropertyStatesProgramChange {
    if casted, ok := structType.(BACnetPropertyStatesProgramChange); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetPropertyStatesProgramChange); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetPropertyStatesProgramChange) GetTypeName() string {
	return "BACnetPropertyStatesProgramChange"
}

func (m *_BACnetPropertyStatesProgramChange) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (programState)
	lengthInBits += m.ProgramState.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetPropertyStatesProgramChange) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetPropertyStatesProgramChangeParse(ctx context.Context, theBytes []byte, peekedTagNumber uint8) (BACnetPropertyStatesProgramChange, error) {
	return BACnetPropertyStatesProgramChangeParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), peekedTagNumber)
}

func BACnetPropertyStatesProgramChangeParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, peekedTagNumber uint8) (BACnetPropertyStatesProgramChange, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetPropertyStatesProgramChange"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetPropertyStatesProgramChange")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (programState)
	if pullErr := readBuffer.PullContext("programState"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for programState")
	}
_programState, _programStateErr := BACnetProgramStateTaggedParseWithBuffer(ctx, readBuffer , uint8( peekedTagNumber ) , TagClass( TagClass_CONTEXT_SPECIFIC_TAGS ) )
	if _programStateErr != nil {
		return nil, errors.Wrap(_programStateErr, "Error parsing 'programState' field of BACnetPropertyStatesProgramChange")
	}
	programState := _programState.(BACnetProgramStateTagged)
	if closeErr := readBuffer.CloseContext("programState"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for programState")
	}

	if closeErr := readBuffer.CloseContext("BACnetPropertyStatesProgramChange"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetPropertyStatesProgramChange")
	}

	// Create a partially initialized instance
	_child := &_BACnetPropertyStatesProgramChange{
		_BACnetPropertyStates: &_BACnetPropertyStates{
		},
		ProgramState: programState,
	}
	_child._BACnetPropertyStates._BACnetPropertyStatesChildRequirements = _child
	return _child, nil
}

func (m *_BACnetPropertyStatesProgramChange) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetPropertyStatesProgramChange) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetPropertyStatesProgramChange"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetPropertyStatesProgramChange")
		}

	// Simple Field (programState)
	if pushErr := writeBuffer.PushContext("programState"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for programState")
	}
	_programStateErr := writeBuffer.WriteSerializable(ctx, m.GetProgramState())
	if popErr := writeBuffer.PopContext("programState"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for programState")
	}
	if _programStateErr != nil {
		return errors.Wrap(_programStateErr, "Error serializing 'programState' field")
	}

		if popErr := writeBuffer.PopContext("BACnetPropertyStatesProgramChange"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetPropertyStatesProgramChange")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetPropertyStatesProgramChange) isBACnetPropertyStatesProgramChange() bool {
	return true
}

func (m *_BACnetPropertyStatesProgramChange) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



