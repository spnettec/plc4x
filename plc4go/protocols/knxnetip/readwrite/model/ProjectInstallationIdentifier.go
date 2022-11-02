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
	"github.com/apache/plc4x/plc4go/spi/utils"
	"github.com/pkg/errors"
)

	// Code generated by code-generation. DO NOT EDIT.


// ProjectInstallationIdentifier is the corresponding interface of ProjectInstallationIdentifier
type ProjectInstallationIdentifier interface {
	utils.LengthAware
	utils.Serializable
	// GetProjectNumber returns ProjectNumber (property field)
	GetProjectNumber() uint8
	// GetInstallationNumber returns InstallationNumber (property field)
	GetInstallationNumber() uint8
}

// ProjectInstallationIdentifierExactly can be used when we want exactly this type and not a type which fulfills ProjectInstallationIdentifier.
// This is useful for switch cases.
type ProjectInstallationIdentifierExactly interface {
	ProjectInstallationIdentifier
	isProjectInstallationIdentifier() bool
}

// _ProjectInstallationIdentifier is the data-structure of this message
type _ProjectInstallationIdentifier struct {
        ProjectNumber uint8
        InstallationNumber uint8
}


///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_ProjectInstallationIdentifier) GetProjectNumber() uint8 {
	return m.ProjectNumber
}

func (m *_ProjectInstallationIdentifier) GetInstallationNumber() uint8 {
	return m.InstallationNumber
}

///////////////////////
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewProjectInstallationIdentifier factory function for _ProjectInstallationIdentifier
func NewProjectInstallationIdentifier( projectNumber uint8 , installationNumber uint8 ) *_ProjectInstallationIdentifier {
return &_ProjectInstallationIdentifier{ ProjectNumber: projectNumber , InstallationNumber: installationNumber }
}

// Deprecated: use the interface for direct cast
func CastProjectInstallationIdentifier(structType interface{}) ProjectInstallationIdentifier {
    if casted, ok := structType.(ProjectInstallationIdentifier); ok {
		return casted
	}
	if casted, ok := structType.(*ProjectInstallationIdentifier); ok {
		return *casted
	}
	return nil
}

func (m *_ProjectInstallationIdentifier) GetTypeName() string {
	return "ProjectInstallationIdentifier"
}

func (m *_ProjectInstallationIdentifier) GetLengthInBits() uint16 {
	return m.GetLengthInBitsConditional(false)
}

func (m *_ProjectInstallationIdentifier) GetLengthInBitsConditional(lastItem bool) uint16 {
	lengthInBits := uint16(0)

	// Simple field (projectNumber)
	lengthInBits += 8;

	// Simple field (installationNumber)
	lengthInBits += 8;

	return lengthInBits
}


func (m *_ProjectInstallationIdentifier) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func ProjectInstallationIdentifierParse(readBuffer utils.ReadBuffer) (ProjectInstallationIdentifier, error) {
	positionAware := readBuffer
	_ = positionAware
	if pullErr := readBuffer.PullContext("ProjectInstallationIdentifier"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for ProjectInstallationIdentifier")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (projectNumber)
_projectNumber, _projectNumberErr := readBuffer.ReadUint8("projectNumber", 8)
	if _projectNumberErr != nil {
		return nil, errors.Wrap(_projectNumberErr, "Error parsing 'projectNumber' field of ProjectInstallationIdentifier")
	}
	projectNumber := _projectNumber

	// Simple Field (installationNumber)
_installationNumber, _installationNumberErr := readBuffer.ReadUint8("installationNumber", 8)
	if _installationNumberErr != nil {
		return nil, errors.Wrap(_installationNumberErr, "Error parsing 'installationNumber' field of ProjectInstallationIdentifier")
	}
	installationNumber := _installationNumber

	if closeErr := readBuffer.CloseContext("ProjectInstallationIdentifier"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for ProjectInstallationIdentifier")
	}

	// Create the instance
	return &_ProjectInstallationIdentifier{
			ProjectNumber: projectNumber,
			InstallationNumber: installationNumber,
		}, nil
}

func (m *_ProjectInstallationIdentifier) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithByteOrderForByteBasedBuffer(binary.BigEndian), utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes()))) // TODO: get endianness from mspec
	if err := m.SerializeWithWriteBuffer(wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_ProjectInstallationIdentifier) SerializeWithWriteBuffer(writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	if pushErr :=writeBuffer.PushContext("ProjectInstallationIdentifier"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for ProjectInstallationIdentifier")
	}

	// Simple Field (projectNumber)
	projectNumber := uint8(m.GetProjectNumber())
	_projectNumberErr := writeBuffer.WriteUint8("projectNumber", 8, (projectNumber))
	if _projectNumberErr != nil {
		return errors.Wrap(_projectNumberErr, "Error serializing 'projectNumber' field")
	}

	// Simple Field (installationNumber)
	installationNumber := uint8(m.GetInstallationNumber())
	_installationNumberErr := writeBuffer.WriteUint8("installationNumber", 8, (installationNumber))
	if _installationNumberErr != nil {
		return errors.Wrap(_installationNumberErr, "Error serializing 'installationNumber' field")
	}

	if popErr := writeBuffer.PopContext("ProjectInstallationIdentifier"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for ProjectInstallationIdentifier")
	}
	return nil
}


func (m *_ProjectInstallationIdentifier) isProjectInstallationIdentifier() bool {
	return true
}

func (m *_ProjectInstallationIdentifier) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



