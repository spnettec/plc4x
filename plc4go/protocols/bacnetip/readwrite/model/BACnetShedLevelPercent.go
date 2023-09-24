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


// BACnetShedLevelPercent is the corresponding interface of BACnetShedLevelPercent
type BACnetShedLevelPercent interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	BACnetShedLevel
	// GetPercent returns Percent (property field)
	GetPercent() BACnetContextTagUnsignedInteger
}

// BACnetShedLevelPercentExactly can be used when we want exactly this type and not a type which fulfills BACnetShedLevelPercent.
// This is useful for switch cases.
type BACnetShedLevelPercentExactly interface {
	BACnetShedLevelPercent
	isBACnetShedLevelPercent() bool
}

// _BACnetShedLevelPercent is the data-structure of this message
type _BACnetShedLevelPercent struct {
	*_BACnetShedLevel
        Percent BACnetContextTagUnsignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetShedLevelPercent) InitializeParent(parent BACnetShedLevel , peekedTagHeader BACnetTagHeader ) {	m.PeekedTagHeader = peekedTagHeader
}

func (m *_BACnetShedLevelPercent)  GetParent() BACnetShedLevel {
	return m._BACnetShedLevel
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetShedLevelPercent) GetPercent() BACnetContextTagUnsignedInteger {
	return m.Percent
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetShedLevelPercent factory function for _BACnetShedLevelPercent
func NewBACnetShedLevelPercent( percent BACnetContextTagUnsignedInteger , peekedTagHeader BACnetTagHeader ) *_BACnetShedLevelPercent {
	_result := &_BACnetShedLevelPercent{
		Percent: percent,
    	_BACnetShedLevel: NewBACnetShedLevel(peekedTagHeader),
	}
	_result._BACnetShedLevel._BACnetShedLevelChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetShedLevelPercent(structType any) BACnetShedLevelPercent {
    if casted, ok := structType.(BACnetShedLevelPercent); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetShedLevelPercent); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetShedLevelPercent) GetTypeName() string {
	return "BACnetShedLevelPercent"
}

func (m *_BACnetShedLevelPercent) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (percent)
	lengthInBits += m.Percent.GetLengthInBits(ctx)

	return lengthInBits
}


func (m *_BACnetShedLevelPercent) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetShedLevelPercentParse(ctx context.Context, theBytes []byte) (BACnetShedLevelPercent, error) {
	return BACnetShedLevelPercentParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func BACnetShedLevelPercentParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetShedLevelPercent, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BACnetShedLevelPercent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetShedLevelPercent")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (percent)
	if pullErr := readBuffer.PullContext("percent"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for percent")
	}
_percent, _percentErr := BACnetContextTagParseWithBuffer(ctx, readBuffer , uint8( uint8(0) ) , BACnetDataType( BACnetDataType_UNSIGNED_INTEGER ) )
	if _percentErr != nil {
		return nil, errors.Wrap(_percentErr, "Error parsing 'percent' field of BACnetShedLevelPercent")
	}
	percent := _percent.(BACnetContextTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("percent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for percent")
	}

	if closeErr := readBuffer.CloseContext("BACnetShedLevelPercent"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetShedLevelPercent")
	}

	// Create a partially initialized instance
	_child := &_BACnetShedLevelPercent{
		_BACnetShedLevel: &_BACnetShedLevel{
		},
		Percent: percent,
	}
	_child._BACnetShedLevel._BACnetShedLevelChildRequirements = _child
	return _child, nil
}

func (m *_BACnetShedLevelPercent) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BACnetShedLevelPercent) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetShedLevelPercent"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetShedLevelPercent")
		}

	// Simple Field (percent)
	if pushErr := writeBuffer.PushContext("percent"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for percent")
	}
	_percentErr := writeBuffer.WriteSerializable(ctx, m.GetPercent())
	if popErr := writeBuffer.PopContext("percent"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for percent")
	}
	if _percentErr != nil {
		return errors.Wrap(_percentErr, "Error serializing 'percent' field")
	}

		if popErr := writeBuffer.PopContext("BACnetShedLevelPercent"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetShedLevelPercent")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BACnetShedLevelPercent) isBACnetShedLevelPercent() bool {
	return true
}

func (m *_BACnetShedLevelPercent) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



