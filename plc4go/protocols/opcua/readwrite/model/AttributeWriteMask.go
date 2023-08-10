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

// AttributeWriteMask is an enum
type AttributeWriteMask uint32

type IAttributeWriteMask interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	AttributeWriteMask_attributeWriteMaskNone                    AttributeWriteMask = 0
	AttributeWriteMask_attributeWriteMaskAccessLevel             AttributeWriteMask = 1
	AttributeWriteMask_attributeWriteMaskArrayDimensions         AttributeWriteMask = 2
	AttributeWriteMask_attributeWriteMaskBrowseName              AttributeWriteMask = 4
	AttributeWriteMask_attributeWriteMaskContainsNoLoops         AttributeWriteMask = 8
	AttributeWriteMask_attributeWriteMaskDataType                AttributeWriteMask = 16
	AttributeWriteMask_attributeWriteMaskDescription             AttributeWriteMask = 32
	AttributeWriteMask_attributeWriteMaskDisplayName             AttributeWriteMask = 64
	AttributeWriteMask_attributeWriteMaskEventNotifier           AttributeWriteMask = 128
	AttributeWriteMask_attributeWriteMaskExecutable              AttributeWriteMask = 256
	AttributeWriteMask_attributeWriteMaskHistorizing             AttributeWriteMask = 512
	AttributeWriteMask_attributeWriteMaskInverseName             AttributeWriteMask = 1024
	AttributeWriteMask_attributeWriteMaskIsAbstract              AttributeWriteMask = 2048
	AttributeWriteMask_attributeWriteMaskMinimumSamplingInterval AttributeWriteMask = 4096
	AttributeWriteMask_attributeWriteMaskNodeClass               AttributeWriteMask = 8192
	AttributeWriteMask_attributeWriteMaskNodeId                  AttributeWriteMask = 16384
	AttributeWriteMask_attributeWriteMaskSymmetric               AttributeWriteMask = 32768
	AttributeWriteMask_attributeWriteMaskUserAccessLevel         AttributeWriteMask = 65536
	AttributeWriteMask_attributeWriteMaskUserExecutable          AttributeWriteMask = 131072
	AttributeWriteMask_attributeWriteMaskUserWriteMask           AttributeWriteMask = 262144
	AttributeWriteMask_attributeWriteMaskValueRank               AttributeWriteMask = 524288
	AttributeWriteMask_attributeWriteMaskWriteMask               AttributeWriteMask = 1048576
	AttributeWriteMask_attributeWriteMaskValueForVariableType    AttributeWriteMask = 2097152
	AttributeWriteMask_attributeWriteMaskDataTypeDefinition      AttributeWriteMask = 4194304
	AttributeWriteMask_attributeWriteMaskRolePermissions         AttributeWriteMask = 8388608
	AttributeWriteMask_attributeWriteMaskAccessRestrictions      AttributeWriteMask = 16777216
	AttributeWriteMask_attributeWriteMaskAccessLevelEx           AttributeWriteMask = 33554432
)

var AttributeWriteMaskValues []AttributeWriteMask

func init() {
	_ = errors.New
	AttributeWriteMaskValues = []AttributeWriteMask{
		AttributeWriteMask_attributeWriteMaskNone,
		AttributeWriteMask_attributeWriteMaskAccessLevel,
		AttributeWriteMask_attributeWriteMaskArrayDimensions,
		AttributeWriteMask_attributeWriteMaskBrowseName,
		AttributeWriteMask_attributeWriteMaskContainsNoLoops,
		AttributeWriteMask_attributeWriteMaskDataType,
		AttributeWriteMask_attributeWriteMaskDescription,
		AttributeWriteMask_attributeWriteMaskDisplayName,
		AttributeWriteMask_attributeWriteMaskEventNotifier,
		AttributeWriteMask_attributeWriteMaskExecutable,
		AttributeWriteMask_attributeWriteMaskHistorizing,
		AttributeWriteMask_attributeWriteMaskInverseName,
		AttributeWriteMask_attributeWriteMaskIsAbstract,
		AttributeWriteMask_attributeWriteMaskMinimumSamplingInterval,
		AttributeWriteMask_attributeWriteMaskNodeClass,
		AttributeWriteMask_attributeWriteMaskNodeId,
		AttributeWriteMask_attributeWriteMaskSymmetric,
		AttributeWriteMask_attributeWriteMaskUserAccessLevel,
		AttributeWriteMask_attributeWriteMaskUserExecutable,
		AttributeWriteMask_attributeWriteMaskUserWriteMask,
		AttributeWriteMask_attributeWriteMaskValueRank,
		AttributeWriteMask_attributeWriteMaskWriteMask,
		AttributeWriteMask_attributeWriteMaskValueForVariableType,
		AttributeWriteMask_attributeWriteMaskDataTypeDefinition,
		AttributeWriteMask_attributeWriteMaskRolePermissions,
		AttributeWriteMask_attributeWriteMaskAccessRestrictions,
		AttributeWriteMask_attributeWriteMaskAccessLevelEx,
	}
}

func AttributeWriteMaskByValue(value uint32) (enum AttributeWriteMask, ok bool) {
	switch value {
	case 0:
		return AttributeWriteMask_attributeWriteMaskNone, true
	case 1:
		return AttributeWriteMask_attributeWriteMaskAccessLevel, true
	case 1024:
		return AttributeWriteMask_attributeWriteMaskInverseName, true
	case 1048576:
		return AttributeWriteMask_attributeWriteMaskWriteMask, true
	case 128:
		return AttributeWriteMask_attributeWriteMaskEventNotifier, true
	case 131072:
		return AttributeWriteMask_attributeWriteMaskUserExecutable, true
	case 16:
		return AttributeWriteMask_attributeWriteMaskDataType, true
	case 16384:
		return AttributeWriteMask_attributeWriteMaskNodeId, true
	case 16777216:
		return AttributeWriteMask_attributeWriteMaskAccessRestrictions, true
	case 2:
		return AttributeWriteMask_attributeWriteMaskArrayDimensions, true
	case 2048:
		return AttributeWriteMask_attributeWriteMaskIsAbstract, true
	case 2097152:
		return AttributeWriteMask_attributeWriteMaskValueForVariableType, true
	case 256:
		return AttributeWriteMask_attributeWriteMaskExecutable, true
	case 262144:
		return AttributeWriteMask_attributeWriteMaskUserWriteMask, true
	case 32:
		return AttributeWriteMask_attributeWriteMaskDescription, true
	case 32768:
		return AttributeWriteMask_attributeWriteMaskSymmetric, true
	case 33554432:
		return AttributeWriteMask_attributeWriteMaskAccessLevelEx, true
	case 4:
		return AttributeWriteMask_attributeWriteMaskBrowseName, true
	case 4096:
		return AttributeWriteMask_attributeWriteMaskMinimumSamplingInterval, true
	case 4194304:
		return AttributeWriteMask_attributeWriteMaskDataTypeDefinition, true
	case 512:
		return AttributeWriteMask_attributeWriteMaskHistorizing, true
	case 524288:
		return AttributeWriteMask_attributeWriteMaskValueRank, true
	case 64:
		return AttributeWriteMask_attributeWriteMaskDisplayName, true
	case 65536:
		return AttributeWriteMask_attributeWriteMaskUserAccessLevel, true
	case 8:
		return AttributeWriteMask_attributeWriteMaskContainsNoLoops, true
	case 8192:
		return AttributeWriteMask_attributeWriteMaskNodeClass, true
	case 8388608:
		return AttributeWriteMask_attributeWriteMaskRolePermissions, true
	}
	return 0, false
}

func AttributeWriteMaskByName(value string) (enum AttributeWriteMask, ok bool) {
	switch value {
	case "attributeWriteMaskNone":
		return AttributeWriteMask_attributeWriteMaskNone, true
	case "attributeWriteMaskAccessLevel":
		return AttributeWriteMask_attributeWriteMaskAccessLevel, true
	case "attributeWriteMaskInverseName":
		return AttributeWriteMask_attributeWriteMaskInverseName, true
	case "attributeWriteMaskWriteMask":
		return AttributeWriteMask_attributeWriteMaskWriteMask, true
	case "attributeWriteMaskEventNotifier":
		return AttributeWriteMask_attributeWriteMaskEventNotifier, true
	case "attributeWriteMaskUserExecutable":
		return AttributeWriteMask_attributeWriteMaskUserExecutable, true
	case "attributeWriteMaskDataType":
		return AttributeWriteMask_attributeWriteMaskDataType, true
	case "attributeWriteMaskNodeId":
		return AttributeWriteMask_attributeWriteMaskNodeId, true
	case "attributeWriteMaskAccessRestrictions":
		return AttributeWriteMask_attributeWriteMaskAccessRestrictions, true
	case "attributeWriteMaskArrayDimensions":
		return AttributeWriteMask_attributeWriteMaskArrayDimensions, true
	case "attributeWriteMaskIsAbstract":
		return AttributeWriteMask_attributeWriteMaskIsAbstract, true
	case "attributeWriteMaskValueForVariableType":
		return AttributeWriteMask_attributeWriteMaskValueForVariableType, true
	case "attributeWriteMaskExecutable":
		return AttributeWriteMask_attributeWriteMaskExecutable, true
	case "attributeWriteMaskUserWriteMask":
		return AttributeWriteMask_attributeWriteMaskUserWriteMask, true
	case "attributeWriteMaskDescription":
		return AttributeWriteMask_attributeWriteMaskDescription, true
	case "attributeWriteMaskSymmetric":
		return AttributeWriteMask_attributeWriteMaskSymmetric, true
	case "attributeWriteMaskAccessLevelEx":
		return AttributeWriteMask_attributeWriteMaskAccessLevelEx, true
	case "attributeWriteMaskBrowseName":
		return AttributeWriteMask_attributeWriteMaskBrowseName, true
	case "attributeWriteMaskMinimumSamplingInterval":
		return AttributeWriteMask_attributeWriteMaskMinimumSamplingInterval, true
	case "attributeWriteMaskDataTypeDefinition":
		return AttributeWriteMask_attributeWriteMaskDataTypeDefinition, true
	case "attributeWriteMaskHistorizing":
		return AttributeWriteMask_attributeWriteMaskHistorizing, true
	case "attributeWriteMaskValueRank":
		return AttributeWriteMask_attributeWriteMaskValueRank, true
	case "attributeWriteMaskDisplayName":
		return AttributeWriteMask_attributeWriteMaskDisplayName, true
	case "attributeWriteMaskUserAccessLevel":
		return AttributeWriteMask_attributeWriteMaskUserAccessLevel, true
	case "attributeWriteMaskContainsNoLoops":
		return AttributeWriteMask_attributeWriteMaskContainsNoLoops, true
	case "attributeWriteMaskNodeClass":
		return AttributeWriteMask_attributeWriteMaskNodeClass, true
	case "attributeWriteMaskRolePermissions":
		return AttributeWriteMask_attributeWriteMaskRolePermissions, true
	}
	return 0, false
}

func AttributeWriteMaskKnows(value uint32) bool {
	for _, typeValue := range AttributeWriteMaskValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastAttributeWriteMask(structType any) AttributeWriteMask {
	castFunc := func(typ any) AttributeWriteMask {
		if sAttributeWriteMask, ok := typ.(AttributeWriteMask); ok {
			return sAttributeWriteMask
		}
		return 0
	}
	return castFunc(structType)
}

func (m AttributeWriteMask) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m AttributeWriteMask) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func AttributeWriteMaskParse(ctx context.Context, theBytes []byte) (AttributeWriteMask, error) {
	return AttributeWriteMaskParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func AttributeWriteMaskParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (AttributeWriteMask, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("AttributeWriteMask", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading AttributeWriteMask")
	}
	if enum, ok := AttributeWriteMaskByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for AttributeWriteMask")
		return AttributeWriteMask(val), nil
	} else {
		return enum, nil
	}
}

func (e AttributeWriteMask) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e AttributeWriteMask) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("AttributeWriteMask", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e AttributeWriteMask) PLC4XEnumName() string {
	switch e {
	case AttributeWriteMask_attributeWriteMaskNone:
		return "attributeWriteMaskNone"
	case AttributeWriteMask_attributeWriteMaskAccessLevel:
		return "attributeWriteMaskAccessLevel"
	case AttributeWriteMask_attributeWriteMaskInverseName:
		return "attributeWriteMaskInverseName"
	case AttributeWriteMask_attributeWriteMaskWriteMask:
		return "attributeWriteMaskWriteMask"
	case AttributeWriteMask_attributeWriteMaskEventNotifier:
		return "attributeWriteMaskEventNotifier"
	case AttributeWriteMask_attributeWriteMaskUserExecutable:
		return "attributeWriteMaskUserExecutable"
	case AttributeWriteMask_attributeWriteMaskDataType:
		return "attributeWriteMaskDataType"
	case AttributeWriteMask_attributeWriteMaskNodeId:
		return "attributeWriteMaskNodeId"
	case AttributeWriteMask_attributeWriteMaskAccessRestrictions:
		return "attributeWriteMaskAccessRestrictions"
	case AttributeWriteMask_attributeWriteMaskArrayDimensions:
		return "attributeWriteMaskArrayDimensions"
	case AttributeWriteMask_attributeWriteMaskIsAbstract:
		return "attributeWriteMaskIsAbstract"
	case AttributeWriteMask_attributeWriteMaskValueForVariableType:
		return "attributeWriteMaskValueForVariableType"
	case AttributeWriteMask_attributeWriteMaskExecutable:
		return "attributeWriteMaskExecutable"
	case AttributeWriteMask_attributeWriteMaskUserWriteMask:
		return "attributeWriteMaskUserWriteMask"
	case AttributeWriteMask_attributeWriteMaskDescription:
		return "attributeWriteMaskDescription"
	case AttributeWriteMask_attributeWriteMaskSymmetric:
		return "attributeWriteMaskSymmetric"
	case AttributeWriteMask_attributeWriteMaskAccessLevelEx:
		return "attributeWriteMaskAccessLevelEx"
	case AttributeWriteMask_attributeWriteMaskBrowseName:
		return "attributeWriteMaskBrowseName"
	case AttributeWriteMask_attributeWriteMaskMinimumSamplingInterval:
		return "attributeWriteMaskMinimumSamplingInterval"
	case AttributeWriteMask_attributeWriteMaskDataTypeDefinition:
		return "attributeWriteMaskDataTypeDefinition"
	case AttributeWriteMask_attributeWriteMaskHistorizing:
		return "attributeWriteMaskHistorizing"
	case AttributeWriteMask_attributeWriteMaskValueRank:
		return "attributeWriteMaskValueRank"
	case AttributeWriteMask_attributeWriteMaskDisplayName:
		return "attributeWriteMaskDisplayName"
	case AttributeWriteMask_attributeWriteMaskUserAccessLevel:
		return "attributeWriteMaskUserAccessLevel"
	case AttributeWriteMask_attributeWriteMaskContainsNoLoops:
		return "attributeWriteMaskContainsNoLoops"
	case AttributeWriteMask_attributeWriteMaskNodeClass:
		return "attributeWriteMaskNodeClass"
	case AttributeWriteMask_attributeWriteMaskRolePermissions:
		return "attributeWriteMaskRolePermissions"
	}
	return fmt.Sprintf("Unknown(%v)", uint32(e))
}

func (e AttributeWriteMask) String() string {
	return e.PLC4XEnumName()
}
