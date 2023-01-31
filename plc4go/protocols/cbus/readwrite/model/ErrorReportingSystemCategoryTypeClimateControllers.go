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


// ErrorReportingSystemCategoryTypeClimateControllers is the corresponding interface of ErrorReportingSystemCategoryTypeClimateControllers
type ErrorReportingSystemCategoryTypeClimateControllers interface {
	utils.LengthAware
	utils.Serializable
	ErrorReportingSystemCategoryType
	// GetCategoryForType returns CategoryForType (property field)
	GetCategoryForType() ErrorReportingSystemCategoryTypeForClimateControllers
}

// ErrorReportingSystemCategoryTypeClimateControllersExactly can be used when we want exactly this type and not a type which fulfills ErrorReportingSystemCategoryTypeClimateControllers.
// This is useful for switch cases.
type ErrorReportingSystemCategoryTypeClimateControllersExactly interface {
	ErrorReportingSystemCategoryTypeClimateControllers
	isErrorReportingSystemCategoryTypeClimateControllers() bool
}

// _ErrorReportingSystemCategoryTypeClimateControllers is the data-structure of this message
type _ErrorReportingSystemCategoryTypeClimateControllers struct {
	*_ErrorReportingSystemCategoryType
        CategoryForType ErrorReportingSystemCategoryTypeForClimateControllers
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeClimateControllers)  GetErrorReportingSystemCategoryClass() ErrorReportingSystemCategoryClass {
return ErrorReportingSystemCategoryClass_CLIMATE_CONTROLLERS}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) InitializeParent(parent ErrorReportingSystemCategoryType ) {}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers)  GetParent() ErrorReportingSystemCategoryType {
	return m._ErrorReportingSystemCategoryType
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) GetCategoryForType() ErrorReportingSystemCategoryTypeForClimateControllers {
	return m.CategoryForType
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewErrorReportingSystemCategoryTypeClimateControllers factory function for _ErrorReportingSystemCategoryTypeClimateControllers
func NewErrorReportingSystemCategoryTypeClimateControllers( categoryForType ErrorReportingSystemCategoryTypeForClimateControllers ) *_ErrorReportingSystemCategoryTypeClimateControllers {
	_result := &_ErrorReportingSystemCategoryTypeClimateControllers{
		CategoryForType: categoryForType,
    	_ErrorReportingSystemCategoryType: NewErrorReportingSystemCategoryType(),
	}
	_result._ErrorReportingSystemCategoryType._ErrorReportingSystemCategoryTypeChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastErrorReportingSystemCategoryTypeClimateControllers(structType interface{}) ErrorReportingSystemCategoryTypeClimateControllers {
    if casted, ok := structType.(ErrorReportingSystemCategoryTypeClimateControllers); ok {
		return casted
	}
	if casted, ok := structType.(*ErrorReportingSystemCategoryTypeClimateControllers); ok {
		return *casted
	}
	return nil
}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) GetTypeName() string {
	return "ErrorReportingSystemCategoryTypeClimateControllers"
}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits())

	// Simple field (categoryForType)
	lengthInBits += 4

	return lengthInBits
}


func (m *_ErrorReportingSystemCategoryTypeClimateControllers) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ErrorReportingSystemCategoryTypeClimateControllersParse(theBytes []byte, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (ErrorReportingSystemCategoryTypeClimateControllers, error) {
	return ErrorReportingSystemCategoryTypeClimateControllersParseWithBuffer(utils.NewReadBufferByteBased(theBytes), errorReportingSystemCategoryClass)
}

func ErrorReportingSystemCategoryTypeClimateControllersParseWithBuffer(readBuffer utils.ReadBuffer, errorReportingSystemCategoryClass ErrorReportingSystemCategoryClass) (ErrorReportingSystemCategoryTypeClimateControllers, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ErrorReportingSystemCategoryTypeClimateControllers"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ErrorReportingSystemCategoryTypeClimateControllers")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (categoryForType)
	if pullErr := readBuffer.PullContext("categoryForType"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for categoryForType")
	}
_categoryForType, _categoryForTypeErr := ErrorReportingSystemCategoryTypeForClimateControllersParseWithBuffer(readBuffer)
	if _categoryForTypeErr != nil {
		return nil, errors.Wrap(_categoryForTypeErr, "Error parsing 'categoryForType' field of ErrorReportingSystemCategoryTypeClimateControllers")
	}
	categoryForType := _categoryForType
	if closeErr := readBuffer.CloseContext("categoryForType"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for categoryForType")
	}

	if closeErr := readBuffer.CloseContext("ErrorReportingSystemCategoryTypeClimateControllers"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ErrorReportingSystemCategoryTypeClimateControllers")
	}

	// Create a partially initialized instance
	_child := &_ErrorReportingSystemCategoryTypeClimateControllers{
		_ErrorReportingSystemCategoryType: &_ErrorReportingSystemCategoryType{
		},
		CategoryForType: categoryForType,
	}
	_child._ErrorReportingSystemCategoryType._ErrorReportingSystemCategoryTypeChildRequirements = _child
	return _child, nil
}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes())))
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	ser := func() error {
		if pushErr := writeBuffer.PushContext("ErrorReportingSystemCategoryTypeClimateControllers"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for ErrorReportingSystemCategoryTypeClimateControllers")
		}

	// Simple Field (categoryForType)
	if pushErr := writeBuffer.PushContext("categoryForType"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for categoryForType")
	}
	_categoryForTypeErr := writeBuffer.WriteSerializable(m.GetCategoryForType())
	if popErr := writeBuffer.PopContext("categoryForType"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for categoryForType")
	}
	if _categoryForTypeErr != nil {
		return errors.Wrap(_categoryForTypeErr, "Error serializing 'categoryForType' field")
	}

		if popErr := writeBuffer.PopContext("ErrorReportingSystemCategoryTypeClimateControllers"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for ErrorReportingSystemCategoryTypeClimateControllers")
		}
		return nil
	}
	return m.SerializeParent(writeBuffer, m, ser)
}


func (m *_ErrorReportingSystemCategoryTypeClimateControllers) isErrorReportingSystemCategoryTypeClimateControllers() bool {
	return true
}

func (m *_ErrorReportingSystemCategoryTypeClimateControllers) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



