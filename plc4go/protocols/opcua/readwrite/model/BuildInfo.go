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


// BuildInfo is the corresponding interface of BuildInfo
type BuildInfo interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
	ExtensionObjectDefinition
	// GetProductUri returns ProductUri (property field)
	GetProductUri() PascalString
	// GetManufacturerName returns ManufacturerName (property field)
	GetManufacturerName() PascalString
	// GetProductName returns ProductName (property field)
	GetProductName() PascalString
	// GetSoftwareVersion returns SoftwareVersion (property field)
	GetSoftwareVersion() PascalString
	// GetBuildNumber returns BuildNumber (property field)
	GetBuildNumber() PascalString
	// GetBuildDate returns BuildDate (property field)
	GetBuildDate() int64
}

// BuildInfoExactly can be used when we want exactly this type and not a type which fulfills BuildInfo.
// This is useful for switch cases.
type BuildInfoExactly interface {
	BuildInfo
	isBuildInfo() bool
}

// _BuildInfo is the data-structure of this message
type _BuildInfo struct {
	*_ExtensionObjectDefinition
        ProductUri PascalString
        ManufacturerName PascalString
        ProductName PascalString
        SoftwareVersion PascalString
        BuildNumber PascalString
        BuildDate int64
}



///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for discriminator values.
///////////////////////

func (m *_BuildInfo)  GetIdentifier() string {
return "340"}

///////////////////////-1
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////

func (m *_BuildInfo) InitializeParent(parent ExtensionObjectDefinition ) {}

func (m *_BuildInfo)  GetParent() ExtensionObjectDefinition {
	return m._ExtensionObjectDefinition
}
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////
/////////////////////// Accessors for property fields.
///////////////////////

func (m *_BuildInfo) GetProductUri() PascalString {
	return m.ProductUri
}

func (m *_BuildInfo) GetManufacturerName() PascalString {
	return m.ManufacturerName
}

func (m *_BuildInfo) GetProductName() PascalString {
	return m.ProductName
}

func (m *_BuildInfo) GetSoftwareVersion() PascalString {
	return m.SoftwareVersion
}

func (m *_BuildInfo) GetBuildNumber() PascalString {
	return m.BuildNumber
}

func (m *_BuildInfo) GetBuildDate() int64 {
	return m.BuildDate
}

///////////////////////-2
///////////////////////
///////////////////////////////////////////////////////////
///////////////////////////////////////////////////////////


// NewBuildInfo factory function for _BuildInfo
func NewBuildInfo( productUri PascalString , manufacturerName PascalString , productName PascalString , softwareVersion PascalString , buildNumber PascalString , buildDate int64 ) *_BuildInfo {
	_result := &_BuildInfo{
		ProductUri: productUri,
		ManufacturerName: manufacturerName,
		ProductName: productName,
		SoftwareVersion: softwareVersion,
		BuildNumber: buildNumber,
		BuildDate: buildDate,
    	_ExtensionObjectDefinition: NewExtensionObjectDefinition(),
	}
	_result._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _result
	return _result
}

// Deprecated: use the interface for direct cast
func CastBuildInfo(structType any) BuildInfo {
    if casted, ok := structType.(BuildInfo); ok {
		return casted
	}
	if casted, ok := structType.(*BuildInfo); ok {
		return *casted
	}
	return nil
}

func (m *_BuildInfo) GetTypeName() string {
	return "BuildInfo"
}

func (m *_BuildInfo) GetLengthInBits(ctx context.Context) uint16 {
	lengthInBits := uint16(m.GetParentLengthInBits(ctx))

	// Simple field (productUri)
	lengthInBits += m.ProductUri.GetLengthInBits(ctx)

	// Simple field (manufacturerName)
	lengthInBits += m.ManufacturerName.GetLengthInBits(ctx)

	// Simple field (productName)
	lengthInBits += m.ProductName.GetLengthInBits(ctx)

	// Simple field (softwareVersion)
	lengthInBits += m.SoftwareVersion.GetLengthInBits(ctx)

	// Simple field (buildNumber)
	lengthInBits += m.BuildNumber.GetLengthInBits(ctx)

	// Simple field (buildDate)
	lengthInBits += 64;

	return lengthInBits
}


func (m *_BuildInfo) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func BuildInfoParse(ctx context.Context, theBytes []byte, identifier string) (BuildInfo, error) {
	return BuildInfoParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes), identifier)
}

func BuildInfoParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer, identifier string) (BuildInfo, error) {
	positionAware := readBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	if pullErr := readBuffer.PullContext("BuildInfo"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for BuildInfo")
	}
	currentPos := positionAware.GetPos()
	_ = currentPos

	// Simple Field (productUri)
	if pullErr := readBuffer.PullContext("productUri"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for productUri")
	}
_productUri, _productUriErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _productUriErr != nil {
		return nil, errors.Wrap(_productUriErr, "Error parsing 'productUri' field of BuildInfo")
	}
	productUri := _productUri.(PascalString)
	if closeErr := readBuffer.CloseContext("productUri"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for productUri")
	}

	// Simple Field (manufacturerName)
	if pullErr := readBuffer.PullContext("manufacturerName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for manufacturerName")
	}
_manufacturerName, _manufacturerNameErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _manufacturerNameErr != nil {
		return nil, errors.Wrap(_manufacturerNameErr, "Error parsing 'manufacturerName' field of BuildInfo")
	}
	manufacturerName := _manufacturerName.(PascalString)
	if closeErr := readBuffer.CloseContext("manufacturerName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for manufacturerName")
	}

	// Simple Field (productName)
	if pullErr := readBuffer.PullContext("productName"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for productName")
	}
_productName, _productNameErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _productNameErr != nil {
		return nil, errors.Wrap(_productNameErr, "Error parsing 'productName' field of BuildInfo")
	}
	productName := _productName.(PascalString)
	if closeErr := readBuffer.CloseContext("productName"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for productName")
	}

	// Simple Field (softwareVersion)
	if pullErr := readBuffer.PullContext("softwareVersion"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for softwareVersion")
	}
_softwareVersion, _softwareVersionErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _softwareVersionErr != nil {
		return nil, errors.Wrap(_softwareVersionErr, "Error parsing 'softwareVersion' field of BuildInfo")
	}
	softwareVersion := _softwareVersion.(PascalString)
	if closeErr := readBuffer.CloseContext("softwareVersion"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for softwareVersion")
	}

	// Simple Field (buildNumber)
	if pullErr := readBuffer.PullContext("buildNumber"); pullErr != nil {
		return nil, errors.Wrap(pullErr, "Error pulling for buildNumber")
	}
_buildNumber, _buildNumberErr := PascalStringParseWithBuffer(ctx, readBuffer)
	if _buildNumberErr != nil {
		return nil, errors.Wrap(_buildNumberErr, "Error parsing 'buildNumber' field of BuildInfo")
	}
	buildNumber := _buildNumber.(PascalString)
	if closeErr := readBuffer.CloseContext("buildNumber"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for buildNumber")
	}

	// Simple Field (buildDate)
_buildDate, _buildDateErr := readBuffer.ReadInt64("buildDate", 64)
	if _buildDateErr != nil {
		return nil, errors.Wrap(_buildDateErr, "Error parsing 'buildDate' field of BuildInfo")
	}
	buildDate := _buildDate

	if closeErr := readBuffer.CloseContext("BuildInfo"); closeErr != nil {
		return nil, errors.Wrap(closeErr, "Error closing for BuildInfo")
	}

	// Create a partially initialized instance
	_child := &_BuildInfo{
		_ExtensionObjectDefinition: &_ExtensionObjectDefinition{
		},
		ProductUri: productUri,
		ManufacturerName: manufacturerName,
		ProductName: productName,
		SoftwareVersion: softwareVersion,
		BuildNumber: buildNumber,
		BuildDate: buildDate,
	}
	_child._ExtensionObjectDefinition._ExtensionObjectDefinitionChildRequirements = _child
	return _child, nil
}

func (m *_BuildInfo) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased(utils.WithInitialSizeForByteBasedBuffer(int(m.GetLengthInBytes(context.Background()))))
	if err := m.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (m *_BuildInfo) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	positionAware := writeBuffer
	_ = positionAware
	log := zerolog.Ctx(ctx)
	_ = log
	ser := func() error {
		if pushErr := writeBuffer.PushContext("BuildInfo"); pushErr != nil {
			return errors.Wrap(pushErr, "Error pushing for BuildInfo")
		}

	// Simple Field (productUri)
	if pushErr := writeBuffer.PushContext("productUri"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for productUri")
	}
	_productUriErr := writeBuffer.WriteSerializable(ctx, m.GetProductUri())
	if popErr := writeBuffer.PopContext("productUri"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for productUri")
	}
	if _productUriErr != nil {
		return errors.Wrap(_productUriErr, "Error serializing 'productUri' field")
	}

	// Simple Field (manufacturerName)
	if pushErr := writeBuffer.PushContext("manufacturerName"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for manufacturerName")
	}
	_manufacturerNameErr := writeBuffer.WriteSerializable(ctx, m.GetManufacturerName())
	if popErr := writeBuffer.PopContext("manufacturerName"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for manufacturerName")
	}
	if _manufacturerNameErr != nil {
		return errors.Wrap(_manufacturerNameErr, "Error serializing 'manufacturerName' field")
	}

	// Simple Field (productName)
	if pushErr := writeBuffer.PushContext("productName"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for productName")
	}
	_productNameErr := writeBuffer.WriteSerializable(ctx, m.GetProductName())
	if popErr := writeBuffer.PopContext("productName"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for productName")
	}
	if _productNameErr != nil {
		return errors.Wrap(_productNameErr, "Error serializing 'productName' field")
	}

	// Simple Field (softwareVersion)
	if pushErr := writeBuffer.PushContext("softwareVersion"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for softwareVersion")
	}
	_softwareVersionErr := writeBuffer.WriteSerializable(ctx, m.GetSoftwareVersion())
	if popErr := writeBuffer.PopContext("softwareVersion"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for softwareVersion")
	}
	if _softwareVersionErr != nil {
		return errors.Wrap(_softwareVersionErr, "Error serializing 'softwareVersion' field")
	}

	// Simple Field (buildNumber)
	if pushErr := writeBuffer.PushContext("buildNumber"); pushErr != nil {
		return errors.Wrap(pushErr, "Error pushing for buildNumber")
	}
	_buildNumberErr := writeBuffer.WriteSerializable(ctx, m.GetBuildNumber())
	if popErr := writeBuffer.PopContext("buildNumber"); popErr != nil {
		return errors.Wrap(popErr, "Error popping for buildNumber")
	}
	if _buildNumberErr != nil {
		return errors.Wrap(_buildNumberErr, "Error serializing 'buildNumber' field")
	}

	// Simple Field (buildDate)
	buildDate := int64(m.GetBuildDate())
	_buildDateErr := writeBuffer.WriteInt64("buildDate", 64, (buildDate))
	if _buildDateErr != nil {
		return errors.Wrap(_buildDateErr, "Error serializing 'buildDate' field")
	}

		if popErr := writeBuffer.PopContext("BuildInfo"); popErr != nil {
			return errors.Wrap(popErr, "Error popping for BuildInfo")
		}
		return nil
	}
	return m.SerializeParent(ctx, writeBuffer, m, ser)
}


func (m *_BuildInfo) isBuildInfo() bool {
	return true
}

func (m *_BuildInfo) String() string {
	if m == nil {
		return "<nil>"
	}
	writeBuffer := utils.NewWriteBufferBoxBasedWithOptions(true, true)
	if err := writeBuffer.WriteSerializable(context.Background(), m); err != nil {
		return err.Error()
	}
	return writeBuffer.GetBox().String()
}



