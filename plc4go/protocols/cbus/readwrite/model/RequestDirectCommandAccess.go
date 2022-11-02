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
	"encoding/binary"
	"fmt"
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
	"io"
)

	// Code generated by code-generation. DO NOT EDIT.


// Constant values.
const RequestDirectCommandAccess_AT byte = 0x40

// RequestDirectCommandAccess is the corresponding interface of RequestDirectCommandAccess
type RequestDirectCommandAccess interface {
	utils.LengthAware
	utils.Serializable
	Request
	// GetCalData returns CalData (property field)
	GetCalData() CALData
	// GetAlpha returns Alpha (property field)
	GetAlpha() Alpha
	// GetCalDataDecoded returns CalDataDecoded (virtual field)
	GetCalDataDecoded() CALData
}

// RequestDirectCommandAccessExactly can be used when we want exactly this type and not a type which fulfills RequestDirectCommandAccess.
// This is useful for switch cases.
type RequestDirectCommandAccessExactly interface {
	RequestDirectCommandAccess
	isRequestDirectCommandAccess() bool
}

// _RequestDirectCommandAccess is the data-structure of this message
type _RequestDirectCommandAccess struct {
	*_Request
        CalData CALData
        Alpha Alpha
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_RequestDirectCommandAccess) InitializeParent(parent Request , peekedByte RequestType , startingCR * RequestType , resetMode * RequestType , secondPeek RequestType , termination RequestTermination ) {	m.PeekedByte = peekedByte
	m.StartingCR = startingCR
	m.ResetMode = resetMode
	m.SecondPeek = secondPeek
	m.Termination = termination
}

func (m *_RequestDirectCommandAccess)  GetParent() Request {
	return m._Request
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_RequestDirectCommandAccess) GetCalData() CALData {
	return m.CalData
}

func (m *_RequestDirectCommandAccess) GetAlpha() Alpha {
	return m.Alpha
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_RequestDirectCommandAccess) GetCalDataDecoded() CALData {
	alpha := m.Alpha
	_ = alpha
	return CastCALData(m.GetCalData())
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for const fields.
///////////////////////

func (m *_RequestDirectCommandAccess) GetAt() byte {
	return RequestDirectCommandAccess_AT
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewRequestDirectCommandAccess factory function for _RequestDirectCommandAccess
func NewRequestDirectCommandAccess( calData CALData , alpha Alpha , peekedByte RequestType , startingCR *RequestType , resetMode *RequestType , secondPeek RequestType , termination RequestTermination , cBusOptions CBusOptions ) *_RequestDirectCommandAccess {
	_result := &_RequestDirectCommandAccess{
		CalData: calData,
		Alpha: alpha,
    	_Request: NewRequest(peekedByte, startingCR, resetMode, secondPeek, termination, cBusOptions),
	}
	_result._Request._RequestChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastRequestDirectCommandAccess(structType interface{}) RequestDirectCommandAccess {
    if casted, ok := structType.(RequestDirectCommandAccess); ok {
		return casted
	}
	if casted, ok := structType.(*RequestDirectCommandAccess); ok {
		return *casted
	}
	return nil
}

func (m *_RequestDirectCommandAccess) GetTypeName() string {
	return "RequestDirectCommandAccess"
}

func (m *_RequestDirectCommandAccess) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_RequestDirectCommandAccess) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Const Field (at)
	lengthInBits += 8

	// Manual Field (calData)
	lengthInBits += uint16(int32((int32(m.GetCalData().GetLengthInBytes()) * int32(int32(2)))) * int32(int32(8)))

	// A virtual field doesn't have any in- or output.

	// Optional Field (alpha)
	if m.Alpha != nil {
		lengthInBits += m.Alpha.GetLengthInBits()
	}

	return lengthInBits
}


func (m *_RequestDirectCommandAccess) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func RequestDirectCommandAccessParse(readBuffer utils.ReadBuffer, cBusOptions CBusOptions) (RequestDirectCommandAccess, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("RequestDirectCommandAccess"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for RequestDirectCommandAccess")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Const Field (at)
	at, _atErr := readBuffer.ReadByte("at")
	if _atErr != nil {
		return nil, errors.Wrap(_atErr, "Error parsing 'at' field of RequestDirectCommandAccess")
	}
	if at != RequestDirectCommandAccess_AT {
		return nil, errors.New("Expected constant value " + fmt.Sprintf("%d", RequestDirectCommandAccess_AT) + " but got " + fmt.Sprintf("%d", at))
	}

	// Manual Field (calData)
	_calData, _calDataErr := ReadCALData(readBuffer)
	if _calDataErr != nil {
		return nil, errors.Wrap(_calDataErr, "Error parsing 'calData' field of RequestDirectCommandAccess")
	}
	var calData CALData
	if _calData != nil {
            calData = _calData.(CALData)
	}

	// Virtual field
	_calDataDecoded := calData
	calDataDecoded := _calDataDecoded
	_ = calDataDecoded

	// Optional Field (alpha) (Can be skipped, if a given expression evaluates to false)
	var alpha Alpha = nil
{
		currentPos = positionAware.GetPos()
		if pullErr := readBuffer.PullContext("alpha"); pullErr != nil {
			return nil, errors.Wrap(pullErr, "Error pulling for alpha")
		}
_val, _err := AlphaParse(readBuffer)
		switch {
		case errors.Is(_err, utils.ParseAssertError{}) || errors.Is(_err, io.EOF):
			Plc4xModelLog.Debug().Err(_err).Msg("Resetting position because optional threw an error")
			readBuffer.Reset(currentPos)
		case _err != nil:
			return nil, errors.Wrap(_err, "Error parsing 'alpha' field of RequestDirectCommandAccess")
		default:
			alpha = _val.(Alpha)
			if closeErr := readBuffer.CloseContext("alpha"); closeErr != nil {
				return nil, errors.Wrap(closeErr, "Error closing for alpha")
			}
		}
	}

	if closeErr := readBuffer.CloseContext("RequestDirectCommandAccess"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for RequestDirectCommandAccess")
	}

	// Create a partially initialized instance
	_child := &_RequestDirectCommandAccess{
		_Request: &_Request{
			CBusOptions: cBusOptions,
		},
		CalData: calData,
		Alpha: alpha,
	}
	_child._Request._RequestChildRequirements = _child
	return _child, nil
}

func (m *_RequestDirectCommandAccess) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_RequestDirectCommandAccess) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("RequestDirectCommandAccess"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for RequestDirectCommandAccess")
		}

	// Const Field (at)
	_atErr := writeBuffer.WriteByte("at", 0x40)
	if _atErr != nil {
		return errors.Wrap(_atErr, "Error serializing 'at' field")
	}

	// Manual Field (calData)
	_calDataErr := WriteCALData(writeBuffer, m.GetCalData())
	if _calDataErr != nil {
		return errors.Wrap(_calDataErr, "Error serializing 'calData' field")
	}
	// Virtual field
	if _calDataDecodedErr := writeBuffer.WriteVirtual("calDataDecoded", m.GetCalDataDecoded()); _calDataDecodedErr != nil {
		return errors.Wrap(_calDataDecodedErr, "Error serializing 'calDataDecoded' field")
	}

	// Optional Field (alpha) (Can be skipped, if the value is null)
	var alpha Alpha = nil
	if m.GetAlpha() != nil {
		if pushErr := writeBuffer.PushContext("alpha"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for alpha")
		}
		alpha = m.GetAlpha()
		_alphaErr := writeBuffer.WriteSerializable(alpha)
		if popErr := writeBuffer.PopContext("alpha"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for alpha")
		}
		if _alphaErr != nil {
			return errors.Wrap(_alphaErr, "Error serializing 'alpha' field")
		}
	}

		if popErr := writeBuffer.PopContext("RequestDirectCommandAccess"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for RequestDirectCommandAccess")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_RequestDirectCommandAccess) isRequestDirectCommandAccess() bool {
	return true
}

func (m *_RequestDirectCommandAccess) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



