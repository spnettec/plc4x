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


// IdentifyReplyCommandGAVValuesStored is the corresponding interface of IdentifyReplyCommandGAVValuesStored
type IdentifyReplyCommandGAVValuesStored interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	IdentifyReplyCommand
	// GetValues returns Values (property field)
	GetValues() []byte
}

// IdentifyReplyCommandGAVValuesStoredExactly can be used when we want exactly this type and not a type which fulfills IdentifyReplyCommandGAVValuesStored.
// This is useful for switch cases.
type IdentifyReplyCommandGAVValuesStoredExactly interface {
	IdentifyReplyCommandGAVValuesStored
	isIdentifyReplyCommandGAVValuesStored() bool
}

// _IdentifyReplyCommandGAVValuesStored is the data-structure of this message
type _IdentifyReplyCommandGAVValuesStored struct {
	*_IdentifyReplyCommand
        Values []byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_IdentifyReplyCommandGAVValuesStored)  GetAttribute() Attribute {
return Attribute_GAVValuesStored}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_IdentifyReplyCommandGAVValuesStored) InitializeParent(parent IdentifyReplyCommand ) {}

func (m *_IdentifyReplyCommandGAVValuesStored)  GetParent() IdentifyReplyCommand {
	return m._IdentifyReplyCommand
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_IdentifyReplyCommandGAVValuesStored) GetValues() []byte {
	return m.Values
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewIdentifyReplyCommandGAVValuesStored factory function for _IdentifyReplyCommandGAVValuesStored
func NewIdentifyReplyCommandGAVValuesStored( values []byte , numBytes uint8 ) *_IdentifyReplyCommandGAVValuesStored {
	_result := &_IdentifyReplyCommandGAVValuesStored{
		Values: values,
    	_IdentifyReplyCommand: NewIdentifyReplyCommand(numBytes),
	}
	_result._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastIdentifyReplyCommandGAVValuesStored(structType any) IdentifyReplyCommandGAVValuesStored {
    if casted, ok := structType.(IdentifyReplyCommandGAVValuesStored); ok {
		return casted
	}
	if casted, ok := structType.(*IdentifyReplyCommandGAVValuesStored); ok {
		return *casted
	}
	return nil
}

func (m *_IdentifyReplyCommandGAVValuesStored) GetTypeName() string {
	return "IdentifyReplyCommandGAVValuesStored"
}

func (m *_IdentifyReplyCommandGAVValuesStored) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Array field
	if len(m.Values) > 0 {
		lengthInBits += 8 * uint16(len(m.Values))
	}

	return lengthInBits
}


func (m *_IdentifyReplyCommandGAVValuesStored) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func IdentifyReplyCommandGAVValuesStoredParse(ctx context.Context, theBytes []byte, attribute Attribute, numBytes uint8) (IdentifyReplyCommandGAVValuesStored, error) {
	return IdentifyReplyCommandGAVValuesStoredParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), attribute, numBytes)
}

func IdentifyReplyCommandGAVValuesStoredParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, attribute Attribute, numBytes uint8) (IdentifyReplyCommandGAVValuesStored, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("IdentifyReplyCommandGAVValuesStored"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for IdentifyReplyCommandGAVValuesStored")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos
	// Byte Array field (values)
	numberOfBytesvalues := int(numBytes)
	values, _readArrayErr := readBuffer.ReadByteArray("values", numberOfBytesvalues)
	if _readArrayErr != nil {
		return nil, errors.Wrap(_readArrayErr, "Error parsing 'values' field of IdentifyReplyCommandGAVValuesStored")
	}

	if closeErr := readBuffer.CloseContext("IdentifyReplyCommandGAVValuesStored"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for IdentifyReplyCommandGAVValuesStored")
	}

	// Create a partially initialized instance
	_child := &_IdentifyReplyCommandGAVValuesStored{
		_IdentifyReplyCommand: &_IdentifyReplyCommand{
			NumBytes: numBytes,
		},
		Values: values,
	}
	_child._IdentifyReplyCommand._IdentifyReplyCommandChildRequirements = _child
	return _child, nil
}

func (m *_IdentifyReplyCommandGAVValuesStored) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_IdentifyReplyCommandGAVValuesStored) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("IdentifyReplyCommandGAVValuesStored"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for IdentifyReplyCommandGAVValuesStored")
		}

	// Array Field (values)
	// Byte Array field (values)
	if err := writeBuffer.WriteByteArray("values", m.GetValues()); err != nil {
		return errors.Wrap(err, "Error serializing 'values' field")
	}

		if popErr := writeBuffer.PopContext("IdentifyReplyCommandGAVValuesStored"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for IdentifyReplyCommandGAVValuesStored")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_IdentifyReplyCommandGAVValuesStored) isIdentifyReplyCommandGAVValuesStored() bool {
	return true
}

func (m *_IdentifyReplyCommandGAVValuesStored) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



