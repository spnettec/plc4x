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


// SecurityDataPasswordEntryStatus is the corresponding interface of SecurityDataPasswordEntryStatus
type SecurityDataPasswordEntryStatus interface {
	utils.LengthAware
	utils.Serializable
	SecurityData
	// GetCode returns Code (property field)
	GetCode() byte
	// GetIsPasswordEntrySucceeded returns IsPasswordEntrySucceeded (virtual field)
	GetIsPasswordEntrySucceeded() bool
	// GetIsPasswordEntryFailed returns IsPasswordEntryFailed (virtual field)
	GetIsPasswordEntryFailed() bool
	// GetIsPasswordEntryDisabled returns IsPasswordEntryDisabled (virtual field)
	GetIsPasswordEntryDisabled() bool
	// GetIsPasswordEntryEnabledAgain returns IsPasswordEntryEnabledAgain (virtual field)
	GetIsPasswordEntryEnabledAgain() bool
	// GetIsReserved returns IsReserved (virtual field)
	GetIsReserved() bool
}

// SecurityDataPasswordEntryStatusExactly can be used when we want exactly this type and not a type which fulfills SecurityDataPasswordEntryStatus.
// This is useful for switch cases.
type SecurityDataPasswordEntryStatusExactly interface {
	SecurityDataPasswordEntryStatus
	isSecurityDataPasswordEntryStatus() bool
}

// _SecurityDataPasswordEntryStatus is the data-structure of this message
type _SecurityDataPasswordEntryStatus struct {
	*_SecurityData
        Code byte
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_SecurityDataPasswordEntryStatus) InitializeParent(parent SecurityData , commandTypeContainer SecurityCommandTypeContainer , argument byte ) {	m.CommandTypeContainer = commandTypeContainer
	m.Argument = argument
}

func (m *_SecurityDataPasswordEntryStatus)  GetParent() SecurityData {
	return m._SecurityData
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_SecurityDataPasswordEntryStatus) GetCode() byte {
	return m.Code
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for virtual fields.
///////////////////////

func (m *_SecurityDataPasswordEntryStatus) GetIsPasswordEntrySucceeded() bool {
	return bool(bool((m.GetCode()) == (0x01)))
}

func (m *_SecurityDataPasswordEntryStatus) GetIsPasswordEntryFailed() bool {
	return bool(bool((m.GetCode()) == (0x02)))
}

func (m *_SecurityDataPasswordEntryStatus) GetIsPasswordEntryDisabled() bool {
	return bool(bool((m.GetCode()) == (0x03)))
}

func (m *_SecurityDataPasswordEntryStatus) GetIsPasswordEntryEnabledAgain() bool {
	return bool(bool((m.GetCode()) == (0x04)))
}

func (m *_SecurityDataPasswordEntryStatus) GetIsReserved() bool {
	return bool(bool((m.GetCode()) >= (0x05)))
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewSecurityDataPasswordEntryStatus factory function for _SecurityDataPasswordEntryStatus
func NewSecurityDataPasswordEntryStatus( code byte , commandTypeContainer SecurityCommandTypeContainer , argument byte ) *_SecurityDataPasswordEntryStatus {
	_result := &_SecurityDataPasswordEntryStatus{
		Code: code,
    	_SecurityData: NewSecurityData(commandTypeContainer, argument),
	}
	_result._SecurityData._SecurityDataChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastSecurityDataPasswordEntryStatus(structType interface{}) SecurityDataPasswordEntryStatus {
    if casted, ok := structType.(SecurityDataPasswordEntryStatus); ok {
		return casted
	}
	if casted, ok := structType.(*SecurityDataPasswordEntryStatus); ok {
		return *casted
	}
	return nil
}

func (m *_SecurityDataPasswordEntryStatus) GetTypeName() string {
	return "SecurityDataPasswordEntryStatus"
}

func (m *_SecurityDataPasswordEntryStatus) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_SecurityDataPasswordEntryStatus) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (code)
	lengthInBits += 8;

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	// A virtual field doesn't have any in- or output.

	return lengthInBits
}


func (m *_SecurityDataPasswordEntryStatus) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func SecurityDataPasswordEntryStatusParse(readBuffer utils.ReadBuffer) (SecurityDataPasswordEntryStatus, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("SecurityDataPasswordEntryStatus"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for SecurityDataPasswordEntryStatus")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (code)
_code, _codeErr := readBuffer.ReadByte("code")
	if _codeErr != nil {
		return nil, errors.Wrap(_codeErr, "Error parsing 'code' field of SecurityDataPasswordEntryStatus")
	}
	code := _code

	// Virtual field
	_isPasswordEntrySucceeded := bool((code) == (0x01))
	isPasswordEntrySucceeded := bool(_isPasswordEntrySucceeded)
	_ = isPasswordEntrySucceeded

	// Virtual field
	_isPasswordEntryFailed := bool((code) == (0x02))
	isPasswordEntryFailed := bool(_isPasswordEntryFailed)
	_ = isPasswordEntryFailed

	// Virtual field
	_isPasswordEntryDisabled := bool((code) == (0x03))
	isPasswordEntryDisabled := bool(_isPasswordEntryDisabled)
	_ = isPasswordEntryDisabled

	// Virtual field
	_isPasswordEntryEnabledAgain := bool((code) == (0x04))
	isPasswordEntryEnabledAgain := bool(_isPasswordEntryEnabledAgain)
	_ = isPasswordEntryEnabledAgain

	// Virtual field
	_isReserved := bool((code) >= (0x05))
	isReserved := bool(_isReserved)
	_ = isReserved

	if closeErr := readBuffer.CloseContext("SecurityDataPasswordEntryStatus"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for SecurityDataPasswordEntryStatus")
	}

	// Create a partially initialized instance
	_child := &_SecurityDataPasswordEntryStatus{
		_SecurityData: &_SecurityData{
		},
		Code: code,
	}
	_child._SecurityData._SecurityDataChildRequirements = _child
	return _child, nil
}

func (m *_SecurityDataPasswordEntryStatus) Serialize(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("SecurityDataPasswordEntryStatus"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for SecurityDataPasswordEntryStatus")
		}

	// Simple Field (code)
	code := byte(m.GetCode())
	_codeErr := writeBuffer.WriteByte("code", (code))
	if _codeErr != nil {
		return errors.Wrap(_codeErr, "Error serializing 'code' field")
	}
	// Virtual field
	if _isPasswordEntrySucceededErr := writeBuffer.WriteVirtual("isPasswordEntrySucceeded", m.GetIsPasswordEntrySucceeded()); _isPasswordEntrySucceededErr != nil {
		return errors.Wrap(_isPasswordEntrySucceededErr, "Error serializing 'isPasswordEntrySucceeded' field")
	}
	// Virtual field
	if _isPasswordEntryFailedErr := writeBuffer.WriteVirtual("isPasswordEntryFailed", m.GetIsPasswordEntryFailed()); _isPasswordEntryFailedErr != nil {
		return errors.Wrap(_isPasswordEntryFailedErr, "Error serializing 'isPasswordEntryFailed' field")
	}
	// Virtual field
	if _isPasswordEntryDisabledErr := writeBuffer.WriteVirtual("isPasswordEntryDisabled", m.GetIsPasswordEntryDisabled()); _isPasswordEntryDisabledErr != nil {
		return errors.Wrap(_isPasswordEntryDisabledErr, "Error serializing 'isPasswordEntryDisabled' field")
	}
	// Virtual field
	if _isPasswordEntryEnabledAgainErr := writeBuffer.WriteVirtual("isPasswordEntryEnabledAgain", m.GetIsPasswordEntryEnabledAgain()); _isPasswordEntryEnabledAgainErr != nil {
		return errors.Wrap(_isPasswordEntryEnabledAgainErr, "Error serializing 'isPasswordEntryEnabledAgain' field")
	}
	// Virtual field
	if _isReservedErr := writeBuffer.WriteVirtual("isReserved", m.GetIsReserved()); _isReservedErr != nil {
		return errors.Wrap(_isReservedErr, "Error serializing 'isReserved' field")
	}

		if popErr := writeBuffer.PopContext("SecurityDataPasswordEntryStatus"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for SecurityDataPasswordEntryStatus")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_SecurityDataPasswordEntryStatus) isSecurityDataPasswordEntryStatus() bool {
	return true
}

func (m *_SecurityDataPasswordEntryStatus) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



