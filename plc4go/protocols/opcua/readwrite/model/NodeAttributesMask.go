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

// NodeAttributesMask is an enum
type NodeAttributesMask uint32

type INodeAttributesMask interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	NodeAttributesMask_nodeAttributesMaskNone NodeAttributesMask = 0
	NodeAttributesMask_nodeAttributesMaskAccessLevel NodeAttributesMask = 1
	NodeAttributesMask_nodeAttributesMaskArrayDimensions NodeAttributesMask = 2
	NodeAttributesMask_nodeAttributesMaskBrowseName NodeAttributesMask = 4
	NodeAttributesMask_nodeAttributesMaskContainsNoLoops NodeAttributesMask = 8
	NodeAttributesMask_nodeAttributesMaskDataType NodeAttributesMask = 16
	NodeAttributesMask_nodeAttributesMaskDescription NodeAttributesMask = 32
	NodeAttributesMask_nodeAttributesMaskDisplayName NodeAttributesMask = 64
	NodeAttributesMask_nodeAttributesMaskEventNotifier NodeAttributesMask = 128
	NodeAttributesMask_nodeAttributesMaskExecutable NodeAttributesMask = 256
	NodeAttributesMask_nodeAttributesMaskHistorizing NodeAttributesMask = 512
	NodeAttributesMask_nodeAttributesMaskInverseName NodeAttributesMask = 1024
	NodeAttributesMask_nodeAttributesMaskIsAbstract NodeAttributesMask = 2048
	NodeAttributesMask_nodeAttributesMaskMinimumSamplingInterval NodeAttributesMask = 4096
	NodeAttributesMask_nodeAttributesMaskNodeClass NodeAttributesMask = 8192
	NodeAttributesMask_nodeAttributesMaskNodeId NodeAttributesMask = 16384
	NodeAttributesMask_nodeAttributesMaskSymmetric NodeAttributesMask = 32768
	NodeAttributesMask_nodeAttributesMaskUserAccessLevel NodeAttributesMask = 65536
	NodeAttributesMask_nodeAttributesMaskUserExecutable NodeAttributesMask = 131072
	NodeAttributesMask_nodeAttributesMaskUserWriteMask NodeAttributesMask = 262144
	NodeAttributesMask_nodeAttributesMaskValueRank NodeAttributesMask = 524288
	NodeAttributesMask_nodeAttributesMaskWriteMask NodeAttributesMask = 1048576
	NodeAttributesMask_nodeAttributesMaskValue NodeAttributesMask = 2097152
	NodeAttributesMask_nodeAttributesMaskDataTypeDefinition NodeAttributesMask = 4194304
	NodeAttributesMask_nodeAttributesMaskRolePermissions NodeAttributesMask = 8388608
	NodeAttributesMask_nodeAttributesMaskAccessRestrictions NodeAttributesMask = 16777216
	NodeAttributesMask_nodeAttributesMaskAll NodeAttributesMask = 33554431
	NodeAttributesMask_nodeAttributesMaskBaseNode NodeAttributesMask = 26501220
	NodeAttributesMask_nodeAttributesMaskObject NodeAttributesMask = 26501348
	NodeAttributesMask_nodeAttributesMaskObjectType NodeAttributesMask = 26503268
	NodeAttributesMask_nodeAttributesMaskVariable NodeAttributesMask = 26571383
	NodeAttributesMask_nodeAttributesMaskVariableType NodeAttributesMask = 28600438
	NodeAttributesMask_nodeAttributesMaskMethod NodeAttributesMask = 26632548
	NodeAttributesMask_nodeAttributesMaskReferenceType NodeAttributesMask = 26537060
	NodeAttributesMask_nodeAttributesMaskView NodeAttributesMask = 26501356
)

var NodeAttributesMaskValues []NodeAttributesMask

func init() {
	_ = errors.New
	NodeAttributesMaskValues = []NodeAttributesMask {
		NodeAttributesMask_nodeAttributesMaskNone,
		NodeAttributesMask_nodeAttributesMaskAccessLevel,
		NodeAttributesMask_nodeAttributesMaskArrayDimensions,
		NodeAttributesMask_nodeAttributesMaskBrowseName,
		NodeAttributesMask_nodeAttributesMaskContainsNoLoops,
		NodeAttributesMask_nodeAttributesMaskDataType,
		NodeAttributesMask_nodeAttributesMaskDescription,
		NodeAttributesMask_nodeAttributesMaskDisplayName,
		NodeAttributesMask_nodeAttributesMaskEventNotifier,
		NodeAttributesMask_nodeAttributesMaskExecutable,
		NodeAttributesMask_nodeAttributesMaskHistorizing,
		NodeAttributesMask_nodeAttributesMaskInverseName,
		NodeAttributesMask_nodeAttributesMaskIsAbstract,
		NodeAttributesMask_nodeAttributesMaskMinimumSamplingInterval,
		NodeAttributesMask_nodeAttributesMaskNodeClass,
		NodeAttributesMask_nodeAttributesMaskNodeId,
		NodeAttributesMask_nodeAttributesMaskSymmetric,
		NodeAttributesMask_nodeAttributesMaskUserAccessLevel,
		NodeAttributesMask_nodeAttributesMaskUserExecutable,
		NodeAttributesMask_nodeAttributesMaskUserWriteMask,
		NodeAttributesMask_nodeAttributesMaskValueRank,
		NodeAttributesMask_nodeAttributesMaskWriteMask,
		NodeAttributesMask_nodeAttributesMaskValue,
		NodeAttributesMask_nodeAttributesMaskDataTypeDefinition,
		NodeAttributesMask_nodeAttributesMaskRolePermissions,
		NodeAttributesMask_nodeAttributesMaskAccessRestrictions,
		NodeAttributesMask_nodeAttributesMaskAll,
		NodeAttributesMask_nodeAttributesMaskBaseNode,
		NodeAttributesMask_nodeAttributesMaskObject,
		NodeAttributesMask_nodeAttributesMaskObjectType,
		NodeAttributesMask_nodeAttributesMaskVariable,
		NodeAttributesMask_nodeAttributesMaskVariableType,
		NodeAttributesMask_nodeAttributesMaskMethod,
		NodeAttributesMask_nodeAttributesMaskReferenceType,
		NodeAttributesMask_nodeAttributesMaskView,
	}
}

func NodeAttributesMaskByValue(value uint32) (enum NodeAttributesMask, ok bool) {
	switch value {
		case 0:
			return NodeAttributesMask_nodeAttributesMaskNone, true
		case 1:
			return NodeAttributesMask_nodeAttributesMaskAccessLevel, true
		case 1024:
			return NodeAttributesMask_nodeAttributesMaskInverseName, true
		case 1048576:
			return NodeAttributesMask_nodeAttributesMaskWriteMask, true
		case 128:
			return NodeAttributesMask_nodeAttributesMaskEventNotifier, true
		case 131072:
			return NodeAttributesMask_nodeAttributesMaskUserExecutable, true
		case 16:
			return NodeAttributesMask_nodeAttributesMaskDataType, true
		case 16384:
			return NodeAttributesMask_nodeAttributesMaskNodeId, true
		case 16777216:
			return NodeAttributesMask_nodeAttributesMaskAccessRestrictions, true
		case 2:
			return NodeAttributesMask_nodeAttributesMaskArrayDimensions, true
		case 2048:
			return NodeAttributesMask_nodeAttributesMaskIsAbstract, true
		case 2097152:
			return NodeAttributesMask_nodeAttributesMaskValue, true
		case 256:
			return NodeAttributesMask_nodeAttributesMaskExecutable, true
		case 262144:
			return NodeAttributesMask_nodeAttributesMaskUserWriteMask, true
		case 26501220:
			return NodeAttributesMask_nodeAttributesMaskBaseNode, true
		case 26501348:
			return NodeAttributesMask_nodeAttributesMaskObject, true
		case 26501356:
			return NodeAttributesMask_nodeAttributesMaskView, true
		case 26503268:
			return NodeAttributesMask_nodeAttributesMaskObjectType, true
		case 26537060:
			return NodeAttributesMask_nodeAttributesMaskReferenceType, true
		case 26571383:
			return NodeAttributesMask_nodeAttributesMaskVariable, true
		case 26632548:
			return NodeAttributesMask_nodeAttributesMaskMethod, true
		case 28600438:
			return NodeAttributesMask_nodeAttributesMaskVariableType, true
		case 32:
			return NodeAttributesMask_nodeAttributesMaskDescription, true
		case 32768:
			return NodeAttributesMask_nodeAttributesMaskSymmetric, true
		case 33554431:
			return NodeAttributesMask_nodeAttributesMaskAll, true
		case 4:
			return NodeAttributesMask_nodeAttributesMaskBrowseName, true
		case 4096:
			return NodeAttributesMask_nodeAttributesMaskMinimumSamplingInterval, true
		case 4194304:
			return NodeAttributesMask_nodeAttributesMaskDataTypeDefinition, true
		case 512:
			return NodeAttributesMask_nodeAttributesMaskHistorizing, true
		case 524288:
			return NodeAttributesMask_nodeAttributesMaskValueRank, true
		case 64:
			return NodeAttributesMask_nodeAttributesMaskDisplayName, true
		case 65536:
			return NodeAttributesMask_nodeAttributesMaskUserAccessLevel, true
		case 8:
			return NodeAttributesMask_nodeAttributesMaskContainsNoLoops, true
		case 8192:
			return NodeAttributesMask_nodeAttributesMaskNodeClass, true
		case 8388608:
			return NodeAttributesMask_nodeAttributesMaskRolePermissions, true
	}
	return 0, false
}

func NodeAttributesMaskByName(value string) (enum NodeAttributesMask, ok bool) {
	switch value {
	case "nodeAttributesMaskNone":
		return NodeAttributesMask_nodeAttributesMaskNone, true
	case "nodeAttributesMaskAccessLevel":
		return NodeAttributesMask_nodeAttributesMaskAccessLevel, true
	case "nodeAttributesMaskInverseName":
		return NodeAttributesMask_nodeAttributesMaskInverseName, true
	case "nodeAttributesMaskWriteMask":
		return NodeAttributesMask_nodeAttributesMaskWriteMask, true
	case "nodeAttributesMaskEventNotifier":
		return NodeAttributesMask_nodeAttributesMaskEventNotifier, true
	case "nodeAttributesMaskUserExecutable":
		return NodeAttributesMask_nodeAttributesMaskUserExecutable, true
	case "nodeAttributesMaskDataType":
		return NodeAttributesMask_nodeAttributesMaskDataType, true
	case "nodeAttributesMaskNodeId":
		return NodeAttributesMask_nodeAttributesMaskNodeId, true
	case "nodeAttributesMaskAccessRestrictions":
		return NodeAttributesMask_nodeAttributesMaskAccessRestrictions, true
	case "nodeAttributesMaskArrayDimensions":
		return NodeAttributesMask_nodeAttributesMaskArrayDimensions, true
	case "nodeAttributesMaskIsAbstract":
		return NodeAttributesMask_nodeAttributesMaskIsAbstract, true
	case "nodeAttributesMaskValue":
		return NodeAttributesMask_nodeAttributesMaskValue, true
	case "nodeAttributesMaskExecutable":
		return NodeAttributesMask_nodeAttributesMaskExecutable, true
	case "nodeAttributesMaskUserWriteMask":
		return NodeAttributesMask_nodeAttributesMaskUserWriteMask, true
	case "nodeAttributesMaskBaseNode":
		return NodeAttributesMask_nodeAttributesMaskBaseNode, true
	case "nodeAttributesMaskObject":
		return NodeAttributesMask_nodeAttributesMaskObject, true
	case "nodeAttributesMaskView":
		return NodeAttributesMask_nodeAttributesMaskView, true
	case "nodeAttributesMaskObjectType":
		return NodeAttributesMask_nodeAttributesMaskObjectType, true
	case "nodeAttributesMaskReferenceType":
		return NodeAttributesMask_nodeAttributesMaskReferenceType, true
	case "nodeAttributesMaskVariable":
		return NodeAttributesMask_nodeAttributesMaskVariable, true
	case "nodeAttributesMaskMethod":
		return NodeAttributesMask_nodeAttributesMaskMethod, true
	case "nodeAttributesMaskVariableType":
		return NodeAttributesMask_nodeAttributesMaskVariableType, true
	case "nodeAttributesMaskDescription":
		return NodeAttributesMask_nodeAttributesMaskDescription, true
	case "nodeAttributesMaskSymmetric":
		return NodeAttributesMask_nodeAttributesMaskSymmetric, true
	case "nodeAttributesMaskAll":
		return NodeAttributesMask_nodeAttributesMaskAll, true
	case "nodeAttributesMaskBrowseName":
		return NodeAttributesMask_nodeAttributesMaskBrowseName, true
	case "nodeAttributesMaskMinimumSamplingInterval":
		return NodeAttributesMask_nodeAttributesMaskMinimumSamplingInterval, true
	case "nodeAttributesMaskDataTypeDefinition":
		return NodeAttributesMask_nodeAttributesMaskDataTypeDefinition, true
	case "nodeAttributesMaskHistorizing":
		return NodeAttributesMask_nodeAttributesMaskHistorizing, true
	case "nodeAttributesMaskValueRank":
		return NodeAttributesMask_nodeAttributesMaskValueRank, true
	case "nodeAttributesMaskDisplayName":
		return NodeAttributesMask_nodeAttributesMaskDisplayName, true
	case "nodeAttributesMaskUserAccessLevel":
		return NodeAttributesMask_nodeAttributesMaskUserAccessLevel, true
	case "nodeAttributesMaskContainsNoLoops":
		return NodeAttributesMask_nodeAttributesMaskContainsNoLoops, true
	case "nodeAttributesMaskNodeClass":
		return NodeAttributesMask_nodeAttributesMaskNodeClass, true
	case "nodeAttributesMaskRolePermissions":
		return NodeAttributesMask_nodeAttributesMaskRolePermissions, true
	}
	return 0, false
}

func NodeAttributesMaskKnows(value uint32)  bool {
	for _, typeValue := range NodeAttributesMaskValues {
		if uint32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastNodeAttributesMask(structType any) NodeAttributesMask {
	castFunc := func(typ any) NodeAttributesMask {
		if sNodeAttributesMask, ok := typ.(NodeAttributesMask); ok {
			return sNodeAttributesMask
		}
		return 0
	}
	return castFunc(structType)
}

func (m NodeAttributesMask) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m NodeAttributesMask) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func NodeAttributesMaskParse(ctx context.Context, theBytes []byte) (NodeAttributesMask, error) {
	return NodeAttributesMaskParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func NodeAttributesMaskParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (NodeAttributesMask, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadUint32("NodeAttributesMask", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading NodeAttributesMask")
	}
	if enum, ok := NodeAttributesMaskByValue(val); !ok {
		log.Debug().Msgf("no value %x found for RequestType", val)
		return NodeAttributesMask(val), nil
	} else {
		return enum, nil
	}
}

func (e NodeAttributesMask) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e NodeAttributesMask) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteUint32("NodeAttributesMask", 32, uint32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e NodeAttributesMask) PLC4XEnumName() string {
	switch e {
	case NodeAttributesMask_nodeAttributesMaskNone:
		return "nodeAttributesMaskNone"
	case NodeAttributesMask_nodeAttributesMaskAccessLevel:
		return "nodeAttributesMaskAccessLevel"
	case NodeAttributesMask_nodeAttributesMaskInverseName:
		return "nodeAttributesMaskInverseName"
	case NodeAttributesMask_nodeAttributesMaskWriteMask:
		return "nodeAttributesMaskWriteMask"
	case NodeAttributesMask_nodeAttributesMaskEventNotifier:
		return "nodeAttributesMaskEventNotifier"
	case NodeAttributesMask_nodeAttributesMaskUserExecutable:
		return "nodeAttributesMaskUserExecutable"
	case NodeAttributesMask_nodeAttributesMaskDataType:
		return "nodeAttributesMaskDataType"
	case NodeAttributesMask_nodeAttributesMaskNodeId:
		return "nodeAttributesMaskNodeId"
	case NodeAttributesMask_nodeAttributesMaskAccessRestrictions:
		return "nodeAttributesMaskAccessRestrictions"
	case NodeAttributesMask_nodeAttributesMaskArrayDimensions:
		return "nodeAttributesMaskArrayDimensions"
	case NodeAttributesMask_nodeAttributesMaskIsAbstract:
		return "nodeAttributesMaskIsAbstract"
	case NodeAttributesMask_nodeAttributesMaskValue:
		return "nodeAttributesMaskValue"
	case NodeAttributesMask_nodeAttributesMaskExecutable:
		return "nodeAttributesMaskExecutable"
	case NodeAttributesMask_nodeAttributesMaskUserWriteMask:
		return "nodeAttributesMaskUserWriteMask"
	case NodeAttributesMask_nodeAttributesMaskBaseNode:
		return "nodeAttributesMaskBaseNode"
	case NodeAttributesMask_nodeAttributesMaskObject:
		return "nodeAttributesMaskObject"
	case NodeAttributesMask_nodeAttributesMaskView:
		return "nodeAttributesMaskView"
	case NodeAttributesMask_nodeAttributesMaskObjectType:
		return "nodeAttributesMaskObjectType"
	case NodeAttributesMask_nodeAttributesMaskReferenceType:
		return "nodeAttributesMaskReferenceType"
	case NodeAttributesMask_nodeAttributesMaskVariable:
		return "nodeAttributesMaskVariable"
	case NodeAttributesMask_nodeAttributesMaskMethod:
		return "nodeAttributesMaskMethod"
	case NodeAttributesMask_nodeAttributesMaskVariableType:
		return "nodeAttributesMaskVariableType"
	case NodeAttributesMask_nodeAttributesMaskDescription:
		return "nodeAttributesMaskDescription"
	case NodeAttributesMask_nodeAttributesMaskSymmetric:
		return "nodeAttributesMaskSymmetric"
	case NodeAttributesMask_nodeAttributesMaskAll:
		return "nodeAttributesMaskAll"
	case NodeAttributesMask_nodeAttributesMaskBrowseName:
		return "nodeAttributesMaskBrowseName"
	case NodeAttributesMask_nodeAttributesMaskMinimumSamplingInterval:
		return "nodeAttributesMaskMinimumSamplingInterval"
	case NodeAttributesMask_nodeAttributesMaskDataTypeDefinition:
		return "nodeAttributesMaskDataTypeDefinition"
	case NodeAttributesMask_nodeAttributesMaskHistorizing:
		return "nodeAttributesMaskHistorizing"
	case NodeAttributesMask_nodeAttributesMaskValueRank:
		return "nodeAttributesMaskValueRank"
	case NodeAttributesMask_nodeAttributesMaskDisplayName:
		return "nodeAttributesMaskDisplayName"
	case NodeAttributesMask_nodeAttributesMaskUserAccessLevel:
		return "nodeAttributesMaskUserAccessLevel"
	case NodeAttributesMask_nodeAttributesMaskContainsNoLoops:
		return "nodeAttributesMaskContainsNoLoops"
	case NodeAttributesMask_nodeAttributesMaskNodeClass:
		return "nodeAttributesMaskNodeClass"
	case NodeAttributesMask_nodeAttributesMaskRolePermissions:
		return "nodeAttributesMaskRolePermissions"
	}
	return ""
}

func (e NodeAttributesMask) String() string {
	return e.PLC4XEnumName()
}

