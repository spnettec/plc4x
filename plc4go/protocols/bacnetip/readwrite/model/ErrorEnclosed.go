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


// ErrorEnclosed is the corresponding interface of ErrorEnclosed
type ErrorEnclosed interface {
	utils.LengthAware
	utils.Serializable
	// GetOpeningTag returns OpeningTag (property field)
	GetOpeningTag() BACnetOpeningTag
	// GetError returns Error (property field)
	GetError() Error
	// GetClosingTag returns ClosingTag (property field)
	GetClosingTag() BACnetClosingTag
}

// ErrorEnclosedExactly can be used when we want exactly this type and not a type which fulfills ErrorEnclosed.
// This is useful for switch cases.
type ErrorEnclosedExactly interface {
	ErrorEnclosed
	isErrorEnclosed() bool
}

// _ErrorEnclosed is the data-structure of this message
type _ErrorEnclosed struct {
        OpeningTag BACnetOpeningTag
        Error Error
        ClosingTag BACnetClosingTag

	// Arguments.
	TagNumber uint8
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorEnclosed) GetOpeningTag() BACnetOpeningTag {
	return m.OpeningTag
}

func (m *_ErrorEnclosed) GetError() Error {
	return m.Error
}

func (m *_ErrorEnclosed) GetClosingTag() BACnetClosingTag {
	return m.ClosingTag
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewErrorEnclosed factory function for _ErrorEnclosed
func NewErrorEnclosed( openingTag BACnetOpeningTag , error Error , closingTag BACnetClosingTag , tagNumber uint8 ) *_ErrorEnclosed {
return &_ErrorEnclosed{ OpeningTag: openingTag , Error: error , ClosingTag: closingTag , TagNumber: tagNumber }
}

// Deprecated: use the interface for direct cast
func CastErrorEnclosed(structType interface{}) ErrorEnclosed {
    if casted, ok := structType.(ErrorEnclosed); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorEnclosed); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorEnclosed) GetTypeName() string {
	return "ErrorEnclosed"
}

func (m *_ErrorEnclosed) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ErrorEnclosed) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (openingTag)
	lengthInBits += m.OpeningTag.GetLengthInBits()

	// Simple field (error)
	lengthInBits += m.Error.GetLengthInBits()

	// Simple field (closingTag)
	lengthInBits += m.ClosingTag.GetLengthInBits()

	return lengthInBits
}


func (m *_ErrorEnclosed) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ErrorEnclosedParse(theBytes []byte, tagNumber uint8) (ErrorEnclosed, error) {
	return ErrorEnclosedParseWithBuffer(utils.NewReadBufferByteBased(theBytes), tagNumber)
}

func ErrorEnclosedParseWithBuffer(readBuffer utils.ReadBuffer, tagNumber uint8) (ErrorEnclosed, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorEnclosed"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorEnclosed")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (openingTag)
	if pullErr := readBuffer.PullContext("openingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for openingTag")
	}
_openingTag, _openingTagErr := BACnetOpeningTagParseWithBuffer(readBuffer , uint8( tagNumber ) )
	if _openingTagErr != nil {
		return nil, errors.Wrap(_openingTagErr, "Error parsing 'openingTag' field of ErrorEnclosed")
	}
	openingTag := _openingTag.(BACnetOpeningTag)
	if closeErr := readBuffer.CloseContext("openingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for openingTag")
	}

	// Simple Field (error)
	if pullErr := readBuffer.PullContext("error"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for error")
	}
_error, _errorErr := ErrorParseWithBuffer(readBuffer)
	if _errorErr != nil {
		return nil, errors.Wrap(_errorErr, "Error parsing 'error' field of ErrorEnclosed")
	}
	error := _error.(Error)
	if closeErr := readBuffer.CloseContext("error"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for error")
	}

	// Simple Field (closingTag)
	if pullErr := readBuffer.PullContext("closingTag"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for closingTag")
	}
_closingTag, _closingTagErr := BACnetClosingTagParseWithBuffer(readBuffer , uint8( tagNumber ) )
	if _closingTagErr != nil {
		return nil, errors.Wrap(_closingTagErr, "Error parsing 'closingTag' field of ErrorEnclosed")
	}
	closingTag := _closingTag.(BACnetClosingTag)
	if closeErr := readBuffer.CloseContext("closingTag"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for closingTag")
	}

	if closeErr := readBuffer.CloseContext("ErrorEnclosed"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorEnclosed")
	}

	// Create the instance
	return &_ErrorEnclosed{
            TagNumber: tagNumber,
			OpeningTag: openingTag,
			Error: error,
			ClosingTag: closingTag,
		}, nil
}

func (m *_ErrorEnclosed) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorEnclosed) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("ErrorEnclosed"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ErrorEnclosed")
	}

	// Simple Field (openingTag)
	if pushErr := writeBuffer.PushContext("openingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for openingTag")
	}
	_openingTagErr := writeBuffer.WriteSerializable(m.GetOpeningTag())
	if popErr := writeBuffer.PopContext("openingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for openingTag")
	}
	if _openingTagErr != nil {
		return errors.Wrap(_openingTagErr, "Error serializing 'openingTag' field")
	}

	// Simple Field (error)
	if pushErr := writeBuffer.PushContext("error"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for error")
	}
	_errorErr := writeBuffer.WriteSerializable(m.GetError())
	if popErr := writeBuffer.PopContext("error"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for error")
	}
	if _errorErr != nil {
		return errors.Wrap(_errorErr, "Error serializing 'error' field")
	}

	// Simple Field (closingTag)
	if pushErr := writeBuffer.PushContext("closingTag"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for closingTag")
	}
	_closingTagErr := writeBuffer.WriteSerializable(m.GetClosingTag())
	if popErr := writeBuffer.PopContext("closingTag"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for closingTag")
	}
	if _closingTagErr != nil {
		return errors.Wrap(_closingTagErr, "Error serializing 'closingTag' field")
	}

	if popErr := writeBuffer.PopContext("ErrorEnclosed"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ErrorEnclosed")
	}
	return nil
}


////
// Arguments Getter

func (m *_ErrorEnclosed) GetTagNumber() uint8 {
	return m.TagNumber
}
//
////

func (m *_ErrorEnclosed) isErrorEnclosed() bool {
	return true
}

func (m *_ErrorEnclosed) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



