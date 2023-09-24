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

// OpcuaNodeIdServicesVariableAudio is an enum
type OpcuaNodeIdServicesVariableAudio int32

type IOpcuaNodeIdServicesVariableAudio interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const(
	OpcuaNodeIdServicesVariableAudio_AudioVariableType_ListId OpcuaNodeIdServicesVariableAudio = 17988
	OpcuaNodeIdServicesVariableAudio_AudioVariableType_AgencyId OpcuaNodeIdServicesVariableAudio = 17989
	OpcuaNodeIdServicesVariableAudio_AudioVariableType_VersionId OpcuaNodeIdServicesVariableAudio = 17990
)

var OpcuaNodeIdServicesVariableAudioValues []OpcuaNodeIdServicesVariableAudio

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableAudioValues = []OpcuaNodeIdServicesVariableAudio {
		OpcuaNodeIdServicesVariableAudio_AudioVariableType_ListId,
		OpcuaNodeIdServicesVariableAudio_AudioVariableType_AgencyId,
		OpcuaNodeIdServicesVariableAudio_AudioVariableType_VersionId,
	}
}

func OpcuaNodeIdServicesVariableAudioByValue(value int32) (enum OpcuaNodeIdServicesVariableAudio, ok bool) {
	switch value {
		case 17988:
			return OpcuaNodeIdServicesVariableAudio_AudioVariableType_ListId, true
		case 17989:
			return OpcuaNodeIdServicesVariableAudio_AudioVariableType_AgencyId, true
		case 17990:
			return OpcuaNodeIdServicesVariableAudio_AudioVariableType_VersionId, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableAudioByName(value string) (enum OpcuaNodeIdServicesVariableAudio, ok bool) {
	switch value {
	case "AudioVariableType_ListId":
		return OpcuaNodeIdServicesVariableAudio_AudioVariableType_ListId, true
	case "AudioVariableType_AgencyId":
		return OpcuaNodeIdServicesVariableAudio_AudioVariableType_AgencyId, true
	case "AudioVariableType_VersionId":
		return OpcuaNodeIdServicesVariableAudio_AudioVariableType_VersionId, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableAudioKnows(value int32)  bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableAudioValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false;
}

func CastOpcuaNodeIdServicesVariableAudio(structType any) OpcuaNodeIdServicesVariableAudio {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableAudio {
		if sOpcuaNodeIdServicesVariableAudio, ok := typ.(OpcuaNodeIdServicesVariableAudio); ok {
			return sOpcuaNodeIdServicesVariableAudio
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableAudio) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableAudio) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableAudioParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableAudio, error) {
	return OpcuaNodeIdServicesVariableAudioParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableAudioParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableAudio, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableAudio", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableAudio")
	}
	if enum, ok := OpcuaNodeIdServicesVariableAudioByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableAudio")
		return OpcuaNodeIdServicesVariableAudio(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableAudio) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableAudio) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableAudio", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableAudio) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableAudio_AudioVariableType_ListId:
		return "AudioVariableType_ListId"
	case OpcuaNodeIdServicesVariableAudio_AudioVariableType_AgencyId:
		return "AudioVariableType_AgencyId"
	case OpcuaNodeIdServicesVariableAudio_AudioVariableType_VersionId:
		return "AudioVariableType_VersionId"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableAudio) String() string {
	return e.PLC4XEnumName()
}

