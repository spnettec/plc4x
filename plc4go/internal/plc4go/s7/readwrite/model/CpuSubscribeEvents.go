/*
 * Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
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
	"github.com/apache/plc4x/plc4go/internal/plc4go/spi/utils"
	"github.com/pkg/errors"
)

// Code generated by code-generation. DO NOT EDIT.

type CpuSubscribeEvents uint8

type ICpuSubscribeEvents interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	CpuSubscribeEvents_CPU CpuSubscribeEvents = 0x01
	CpuSubscribeEvents_IM  CpuSubscribeEvents = 0x02
	CpuSubscribeEvents_FM  CpuSubscribeEvents = 0x04
	CpuSubscribeEvents_CP  CpuSubscribeEvents = 0x80
)

var CpuSubscribeEventsValues []CpuSubscribeEvents

func init() {
	_ = errors.New
	CpuSubscribeEventsValues = []CpuSubscribeEvents{
		CpuSubscribeEvents_CPU,
		CpuSubscribeEvents_IM,
		CpuSubscribeEvents_FM,
		CpuSubscribeEvents_CP,
	}
}

func CpuSubscribeEventsByValue(value uint8) CpuSubscribeEvents {
	switch value {
	case 0x01:
		return CpuSubscribeEvents_CPU
	case 0x02:
		return CpuSubscribeEvents_IM
	case 0x04:
		return CpuSubscribeEvents_FM
	case 0x80:
		return CpuSubscribeEvents_CP
	}
	return 0
}

func CpuSubscribeEventsByName(value string) CpuSubscribeEvents {
	switch value {
	case "CPU":
		return CpuSubscribeEvents_CPU
	case "IM":
		return CpuSubscribeEvents_IM
	case "FM":
		return CpuSubscribeEvents_FM
	case "CP":
		return CpuSubscribeEvents_CP
	}
	return 0
}

func CpuSubscribeEventsKnows(value uint8) bool {
	for _, typeValue := range CpuSubscribeEventsValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastCpuSubscribeEvents(structType interface{}) CpuSubscribeEvents {
	castFunc := func(typ interface{}) CpuSubscribeEvents {
		if sCpuSubscribeEvents, ok := typ.(CpuSubscribeEvents); ok {
			return sCpuSubscribeEvents
		}
		return 0
	}
	return castFunc(structType)
}

func (m CpuSubscribeEvents) GetLengthInBits() uint16 {
	return 8
}

func (m CpuSubscribeEvents) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func CpuSubscribeEventsParse(readBuffer utils.ReadBuffer) (CpuSubscribeEvents, error) {
	val, err := readBuffer.ReadUint8("CpuSubscribeEvents", 8)
	if err != nil {
		return 0, nil
	}
	return CpuSubscribeEventsByValue(val), nil
}

func (e CpuSubscribeEvents) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("CpuSubscribeEvents", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e CpuSubscribeEvents) name() string {
	switch e {
	case CpuSubscribeEvents_CPU:
		return "CPU"
	case CpuSubscribeEvents_IM:
		return "IM"
	case CpuSubscribeEvents_FM:
		return "FM"
	case CpuSubscribeEvents_CP:
		return "CP"
	}
	return ""
}

func (e CpuSubscribeEvents) String() string {
	return e.name()
}
