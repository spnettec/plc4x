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


// BACnetConfirmedServiceRequestReadRangeRangeByPosition is the corresponding interface of BACnetConfirmedServiceRequestReadRangeRangeByPosition
type BACnetConfirmedServiceRequestReadRangeRangeByPosition interface {
	utils.LengthAware
	utils.Serializable
	BACnetConfirmedServiceRequestReadRangeRange
	// GetReferenceIndex returns ReferenceIndex (property field)
	GetReferenceIndex() BACnetApplicationTagUnsignedInteger
	// GetCount returns Count (property field)
	GetCount() BACnetApplicationTagSignedInteger
}

// BACnetConfirmedServiceRequestReadRangeRangeByPositionExactly can be used when we want exactly this type and not a type which fulfills BACnetConfirmedServiceRequestReadRangeRangeByPosition.
// This is useful for switch cases.
type BACnetConfirmedServiceRequestReadRangeRangeByPositionExactly interface {
	BACnetConfirmedServiceRequestReadRangeRangeByPosition
	isBACnetConfirmedServiceRequestReadRangeRangeByPosition() bool
}

// _BACnetConfirmedServiceRequestReadRangeRangeByPosition is the data-structure of this message
type _BACnetConfirmedServiceRequestReadRangeRangeByPosition struct {
	*_BACnetConfirmedServiceRequestReadRangeRange
        ReferenceIndex BACnetApplicationTagUnsignedInteger
        Count BACnetApplicationTagSignedInteger
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) InitializeParent(parent BACnetConfirmedServiceRequestReadRangeRange , peekedTagHeader BACnetTagHeader , openingTag BACnetOpeningTag , closingTag BACnetClosingTag ) {	m.PeekedTagHeader = peekedTagHeader
	m.OpeningTag = openingTag
	m.ClosingTag = closingTag
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition)  GetParent() BACnetConfirmedServiceRequestReadRangeRange {
	return m._BACnetConfirmedServiceRequestReadRangeRange
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) GetReferenceIndex() BACnetApplicationTagUnsignedInteger {
	return m.ReferenceIndex
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) GetCount() BACnetApplicationTagSignedInteger {
	return m.Count
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBACnetConfirmedServiceRequestReadRangeRangeByPosition factory function for _BACnetConfirmedServiceRequestReadRangeRangeByPosition
func NewBACnetConfirmedServiceRequestReadRangeRangeByPosition( referenceIndex BACnetApplicationTagUnsignedInteger , count BACnetApplicationTagSignedInteger , peekedTagHeader BACnetTagHeader , openingTag BACnetOpeningTag , closingTag BACnetClosingTag ) *_BACnetConfirmedServiceRequestReadRangeRangeByPosition {
	_result := &_BACnetConfirmedServiceRequestReadRangeRangeByPosition{
		ReferenceIndex: referenceIndex,
		Count: count,
    	_BACnetConfirmedServiceRequestReadRangeRange: NewBACnetConfirmedServiceRequestReadRangeRange(peekedTagHeader, openingTag, closingTag),
	}
	_result._BACnetConfirmedServiceRequestReadRangeRange._BACnetConfirmedServiceRequestReadRangeRangeChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBACnetConfirmedServiceRequestReadRangeRangeByPosition(structType interface{}) BACnetConfirmedServiceRequestReadRangeRangeByPosition {
    if casted, ok := structType.(BACnetConfirmedServiceRequestReadRangeRangeByPosition); ok {
		return casted
	}
	if casted, ok := structType.(*BACnetConfirmedServiceRequestReadRangeRangeByPosition); ok {
		return *casted
	}
	return nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) GetTypeName() string {
	return "BACnetConfirmedServiceRequestReadRangeRangeByPosition"
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (referenceIndex)
	lengthInBits += m.ReferenceIndex.GetLengthInBits()

	// Simple field (count)
	lengthInBits += m.Count.GetLengthInBits()

	return lengthInBits
}


func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestReadRangeRangeByPositionParse(readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestReadRangeRangeByPosition, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("BACnetConfirmedServiceRequestReadRangeRangeByPosition"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BACnetConfirmedServiceRequestReadRangeRangeByPosition")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (referenceIndex)
	if pullErr := readBuffer.PullContext("referenceIndex"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for referenceIndex")
	}
_referenceIndex, _referenceIndexErr := BACnetApplicationTagParse(readBuffer)
	if _referenceIndexErr != nil {
		return nil, errors.Wrap(_referenceIndexErr, "Error parsing 'referenceIndex' field of BACnetConfirmedServiceRequestReadRangeRangeByPosition")
	}
	referenceIndex := _referenceIndex.(BACnetApplicationTagUnsignedInteger)
	if closeErr := readBuffer.CloseContext("referenceIndex"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for referenceIndex")
	}

	// Simple Field (count)
	if pullErr := readBuffer.PullContext("count"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for count")
	}
_count, _countErr := BACnetApplicationTagParse(readBuffer)
	if _countErr != nil {
		return nil, errors.Wrap(_countErr, "Error parsing 'count' field of BACnetConfirmedServiceRequestReadRangeRangeByPosition")
	}
	count := _count.(BACnetApplicationTagSignedInteger)
	if closeErr := readBuffer.CloseContext("count"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for count")
	}

	if closeErr := readBuffer.CloseContext("BACnetConfirmedServiceRequestReadRangeRangeByPosition"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BACnetConfirmedServiceRequestReadRangeRangeByPosition")
	}

	// Create a partially initialized instance
	_child := &_BACnetConfirmedServiceRequestReadRangeRangeByPosition{
		_BACnetConfirmedServiceRequestReadRangeRange: &_BACnetConfirmedServiceRequestReadRangeRange{
		},
		ReferenceIndex: referenceIndex,
		Count: count,
	}
	_child._BACnetConfirmedServiceRequestReadRangeRange._BACnetConfirmedServiceRequestReadRangeRangeChildRequirements = _child
	return _child, nil
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BACnetConfirmedServiceRequestReadRangeRangeByPosition"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BACnetConfirmedServiceRequestReadRangeRangeByPosition")
		}

	// Simple Field (referenceIndex)
	if pushErr := writeBuffer.PushContext("referenceIndex"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for referenceIndex")
	}
	_referenceIndexErr := writeBuffer.WriteSerializable(m.GetReferenceIndex())
	if popErr := writeBuffer.PopContext("referenceIndex"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for referenceIndex")
	}
	if _referenceIndexErr != nil {
		return errors.Wrap(_referenceIndexErr, "Error serializing 'referenceIndex' field")
	}

	// Simple Field (count)
	if pushErr := writeBuffer.PushContext("count"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for count")
	}
	_countErr := writeBuffer.WriteSerializable(m.GetCount())
	if popErr := writeBuffer.PopContext("count"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for count")
	}
	if _countErr != nil {
		return errors.Wrap(_countErr, "Error serializing 'count' field")
	}

		if popErr := writeBuffer.PopContext("BACnetConfirmedServiceRequestReadRangeRangeByPosition"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BACnetConfirmedServiceRequestReadRangeRangeByPosition")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) isBACnetConfirmedServiceRequestReadRangeRangeByPosition() bool {
	return true
}

func (m *_BACnetConfirmedServiceRequestReadRangeRangeByPosition) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



