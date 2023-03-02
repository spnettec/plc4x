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


// BACnetEventParameter is the corresponding interface of BACnetEventParameter
type BACnetEventParameter interface {
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetEventParameterExactly can be used when we want exactly this type and not a type which fulfills BACnetEventParameter.
// This is useful for switch cases.
type BACnetEventParameterExactly interface {
	BACnetEventParameter
	isBACnetEventParameter() bool
}

// _BACnetEventParameter is the data-structure of this message
type _BACnetEventParameter struct {
	_BACnetEventParameterChildRequirements
        PeekedTagHeader BACnetTagHeader
}

type _BACnetEventParameterChildRequirements interface {
	utils.Serializable
	GetLengthInBits(ctx context.Context) uint16
	GetPeekedTagNumber() uint8
}


type BACnetEventParameterParent interface {
	SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetEventParameter, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetEventParameterChild interface {
	utils.Serializable
InitializeParent(parent BACnetEventParameter , peekedTagHeader BACnetTagHeader )
	GetParent() *BACnetEventParameter

	GetTypeName() string
	BACnetEventParameter
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetEventParameter) GetPeekedTagHeader() BACnetTagHeader {
	return m.PeekedTagHeader
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_BACnetEventParameter) GetPeekedTagNumber() uint8 {
	ctx := context.Background()
	_ = ctx
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetEventParameter factory function for _BACnetEventParameter
func NewBACnetEventParameter( peekedTagHeader BACnetTagHeader ) *_BACnetEventParameter {
return &_BACnetEventParameter{ PeekedTagHeader: peekedTagHeader }
}

// Deprecated: use the interface for direct cast
func CastBACnetEventParameter(structType interface{}) BACnetEventParameter {
    if casted, ok := structType.(BACnetEventParameter); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetEventParameter); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetEventParameter) GetTypeName() string {
	return "BACnetEventParameter"
}


func (m *_BACnetEventParameter) GetParentLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetEventParameter) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BACnetEventParameterParse(theBytes []byte) (BACnetEventParameter, error) {
	return BACnetEventParameterParseWithBuffer(context.Background(), utils.NewReadBufferByteBased(theBytes))
}

func BACnetEventParameterParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (BACnetEventParameter, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetEventParameter"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetEventParameter")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

				// Peek Field (peekedTagHeader)
				currentPos = positionAware.GetPos()
				if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
					return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
				}
peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(ctx, readBuffer)
				readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetEventParameterChildSerializeRequirement interface {
		BACnetEventParameter
		InitializeParent(BACnetEventParameter,  BACnetTagHeader)
		GetParent() BACnetEventParameter
	}
	var _childTemp interface{}
	var _child BACnetEventParameterChildSerializeRequirement
	var typeSwitchError error
	switch {
case peekedTagNumber == uint8(0) : // BACnetEventParameterChangeOfBitstring
		_childTemp, typeSwitchError = BACnetEventParameterChangeOfBitstringParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(1) : // BACnetEventParameterChangeOfState
		_childTemp, typeSwitchError = BACnetEventParameterChangeOfStateParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(2) : // BACnetEventParameterChangeOfValue
		_childTemp, typeSwitchError = BACnetEventParameterChangeOfValueParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(3) : // BACnetEventParameterCommandFailure
		_childTemp, typeSwitchError = BACnetEventParameterCommandFailureParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(4) : // BACnetEventParameterFloatingLimit
		_childTemp, typeSwitchError = BACnetEventParameterFloatingLimitParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(5) : // BACnetEventParameterOutOfRange
		_childTemp, typeSwitchError = BACnetEventParameterOutOfRangeParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(8) : // BACnetEventParameterChangeOfLifeSavety
		_childTemp, typeSwitchError = BACnetEventParameterChangeOfLifeSavetyParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(9) : // BACnetEventParameterExtended
		_childTemp, typeSwitchError = BACnetEventParameterExtendedParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(10) : // BACnetEventParameterBufferReady
		_childTemp, typeSwitchError = BACnetEventParameterBufferReadyParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(11) : // BACnetEventParameterUnsignedRange
		_childTemp, typeSwitchError = BACnetEventParameterUnsignedRangeParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(13) : // BACnetEventParameterAccessEvent
		_childTemp, typeSwitchError = BACnetEventParameterAccessEventParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(14) : // BACnetEventParameterDoubleOutOfRange
		_childTemp, typeSwitchError = BACnetEventParameterDoubleOutOfRangeParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(15) : // BACnetEventParameterSignedOutOfRange
		_childTemp, typeSwitchError = BACnetEventParameterSignedOutOfRangeParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(16) : // BACnetEventParameterUnsignedOutOfRange
		_childTemp, typeSwitchError = BACnetEventParameterUnsignedOutOfRangeParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(17) : // BACnetEventParameterChangeOfCharacterString
		_childTemp, typeSwitchError = BACnetEventParameterChangeOfCharacterStringParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(18) : // BACnetEventParameterChangeOfStatusFlags
		_childTemp, typeSwitchError = BACnetEventParameterChangeOfStatusFlagsParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(20) : // BACnetEventParameterNone
		_childTemp, typeSwitchError = BACnetEventParameterNoneParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(21) : // BACnetEventParameterChangeOfDiscreteValue
		_childTemp, typeSwitchError = BACnetEventParameterChangeOfDiscreteValueParseWithBuffer(ctx, readBuffer, )
case peekedTagNumber == uint8(22) : // BACnetEventParameterChangeOfTimer
		_childTemp, typeSwitchError = BACnetEventParameterChangeOfTimerParseWithBuffer(ctx, readBuffer, )
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetEventParameter")
	}
	_child = _childTemp.(BACnetEventParameterChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetEventParameter"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetEventParameter")
	}

	// Finish initializing
_child.InitializeParent(_child , peekedTagHeader )
	return _child, nil
}

func (pm *_BACnetEventParameter) SerializeParent(ctx context.Context, writeBuffer utils.WriteBuffer, child BACnetEventParameter, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetEventParameter"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetEventParameter")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual(ctx, "peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetEventParameter"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetEventParameter")
	}
	return nil
}


func (m *_BACnetEventParameter) isBACnetEventParameter() bool {
	return true
}

func (m *_BACnetEventParameter) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



