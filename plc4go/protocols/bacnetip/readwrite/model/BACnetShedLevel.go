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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// BACnetShedLevel is the corresponding interface of BACnetShedLevel
type BACnetShedLevel interface {
	utils.LengthAware
	utils.Serializable
	// GetPeekedTagHeader returns PeekedTagHeader (property field)
	GetPeekedTagHeader() BACnetTagHeader
	// GetPeekedTagNumber returns PeekedTagNumber (virtual field)
	GetPeekedTagNumber() uint8
}

// BACnetShedLevelExactly can be used when we want exactly this type and not a type which fulfills BACnetShedLevel.
// This is useful for switch cases.
type BACnetShedLevelExactly interface {
	BACnetShedLevel
	isBACnetShedLevel() bool
}

// _BACnetShedLevel is the data-structure of this message
type _BACnetShedLevel struct {
	_BACnetShedLevelChildRequirements
        PeekedTagHeader BACnetTagHeader
}

type _BACnetShedLevelChildRequirements interface {
	utils.Serializable
	GetLengthInBits() uint16
	GetLengthInBitsConditional(lastItem bool) uint16
}


type BACnetShedLevelParent interface {
	SerializeParent(writeBuffer utils.WriteBuffer, child BACnetShedLevel, serializeChildFunction func() error) error
	GetTypeName() string
}

type BACnetShedLevelChild interface {
	utils.Serializable
InitializeParent(parent BACnetShedLevel , peekedTagHeader BACnetTagHeader )
	GetParent() *BACnetShedLevel

	GetTypeName() string
	BACnetShedLevel
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetShedLevel) GetPeekedTagHeader() BACnetTagHeader {
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

func (m *_BACnetShedLevel) GetPeekedTagNumber() uint8 {
	return uint8(m.GetPeekedTagHeader().GetActualTagNumber())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetShedLevel factory function for _BACnetShedLevel
func NewBACnetShedLevel( peekedTagHeader BACnetTagHeader ) *_BACnetShedLevel {
return &_BACnetShedLevel{ PeekedTagHeader: peekedTagHeader }
}

// Deprecated: use the interface for direct cast
func CastBACnetShedLevel(structType interface{}) BACnetShedLevel {
    if casted, ok := structType.(BACnetShedLevel); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetShedLevel); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetShedLevel) GetTypeName() string {
	return "BACnetShedLevel"
}



func (m *_BACnetShedLevel) GetParentLengthInBits() uint16 {
	lengthInBits := uint16(0)

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}

func (m *_BACnetShedLevel) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetShedLevelParse(theBytes []byte) (BACnetShedLevel, error) {
	return BACnetShedLevelParseWithBuffer(utils.NewReadBufferByteBased(theBytes))
}

func BACnetShedLevelParseWithBuffer(readBuffer utils.ReadBuffer) (BACnetShedLevel, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetShedLevel"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetShedLevel")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

				// Peek Field (peekedTagHeader)
				currentPos = positionAware.GetPos()
				if pullErr := readBuffer.PullContext("peekedTagHeader"); pullErr != nil {
					return nil, errors.Wrap(pullErr, "Error pulling for peekedTagHeader")
				}
peekedTagHeader, _ := BACnetTagHeaderParseWithBuffer(readBuffer)
				readBuffer.Reset(currentPos)

	// Virtual field
	_peekedTagNumber := peekedTagHeader.GetActualTagNumber()
	peekedTagNumber := uint8(_peekedTagNumber)
	_ = peekedTagNumber

	// Switch Field (Depending on the discriminator values, passes the instantiation to a sub-type)
	type BACnetShedLevelChildSerializeRequirement interface {
		BACnetShedLevel
		InitializeParent(BACnetShedLevel,  BACnetTagHeader)
		GetParent() BACnetShedLevel
	}
	var _childTemp interface{}
	var _child BACnetShedLevelChildSerializeRequirement
	var typeSwitchError error
	switch {
case peekedTagNumber == uint8(0) : // BACnetShedLevelPercent
		_childTemp, typeSwitchError = BACnetShedLevelPercentParseWithBuffer(readBuffer, )
case peekedTagNumber == uint8(1) : // BACnetShedLevelLevel
		_childTemp, typeSwitchError = BACnetShedLevelLevelParseWithBuffer(readBuffer, )
case peekedTagNumber == uint8(2) : // BACnetShedLevelAmount
		_childTemp, typeSwitchError = BACnetShedLevelAmountParseWithBuffer(readBuffer, )
	default:
		typeSwitchError = errors.Errorf("Unmapped type for parameters [peekedTagNumber=%v]", peekedTagNumber)
	}
	if typeSwitchError != nil {
		return nil, errors.Wrap(typeSwitchError, "Error parsing sub-type for type-switch of BACnetShedLevel")
	}
	_child = _childTemp.(BACnetShedLevelChildSerializeRequirement)

	if closeErr := readBuffer.CloseContext("BACnetShedLevel"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetShedLevel")
	}

	// Finish initializing
_child.InitializeParent(_child , peekedTagHeader )
	return _child, nil
}

func (pm *_BACnetShedLevel) SerializeParent(writeBuffer utils.WriteBuffer, child BACnetShedLevel, serializeChildFunction func() error) error {
	// We redirect all calls through client as some methods are only implemented there
	m := child
	_ = m
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("BACnetShedLevel"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for BACnetShedLevel")
	}
	// Virtual field
	if _peekedTagNumberErr := writeBuffer.WriteVirtual("peekedTagNumber", m.GetPeekedTagNumber()); _peekedTagNumberErr != nil {
		return errors.Wrap(_peekedTagNumberErr, "Error serializing 'peekedTagNumber' field")
	}

	// Switch field (Depending on the discriminator values, passes the serialization to a sub-type)
	if _typeSwitchErr := serializeChildFunction(); _typeSwitchErr != nil {
		return errors.Wrap(_typeSwitchErr, "Error serializing sub-type field")
	}

	if popErr := writeBuffer.PopContext("BACnetShedLevel"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for BACnetShedLevel")
	}
	return nil
}


func (m *_BACnetShedLevel) isBACnetShedLevel() bool {
	return true
}

func (m *_BACnetShedLevel) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



