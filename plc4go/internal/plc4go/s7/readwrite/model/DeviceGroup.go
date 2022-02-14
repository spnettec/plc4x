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

type DeviceGroup int8

type IDeviceGroup interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	DeviceGroup_PG_OR_PC DeviceGroup = 0x01
	DeviceGroup_OS       DeviceGroup = 0x02
	DeviceGroup_OTHERS   DeviceGroup = 0x03
)

var DeviceGroupValues []DeviceGroup

func init() {
	_ = errors.New
	DeviceGroupValues = []DeviceGroup{
		DeviceGroup_PG_OR_PC,
		DeviceGroup_OS,
		DeviceGroup_OTHERS,
	}
}

func DeviceGroupByValue(value int8) DeviceGroup {
	switch value {
	case 0x01:
		return DeviceGroup_PG_OR_PC
	case 0x02:
		return DeviceGroup_OS
	case 0x03:
		return DeviceGroup_OTHERS
	}
	return 0
}

func DeviceGroupByName(value string) DeviceGroup {
	switch value {
	case "PG_OR_PC":
		return DeviceGroup_PG_OR_PC
	case "OS":
		return DeviceGroup_OS
	case "OTHERS":
		return DeviceGroup_OTHERS
	}
	return 0
}

func DeviceGroupKnows(value int8) bool {
	for _, typeValue := range DeviceGroupValues {
		if int8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastDeviceGroup(structType interface{}) DeviceGroup {
	castFunc := func(typ interface{}) DeviceGroup {
		if sDeviceGroup, ok := typ.(DeviceGroup); ok {
			return sDeviceGroup
		}
		return 0
	}
	return castFunc(structType)
}

func (m DeviceGroup) GetLengthInBits() uint16 {
	return 8
}

func (m DeviceGroup) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func DeviceGroupParse(readBuffer utils.ReadBuffer) (DeviceGroup, error) {
	val, err := readBuffer.ReadInt8("DeviceGroup", 8)
	if err != nil {
		return 0, nil
	}
	return DeviceGroupByValue(val), nil
}

func (e DeviceGroup) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteInt8("DeviceGroup", 8, int8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e DeviceGroup) name() string {
	switch e {
	case DeviceGroup_PG_OR_PC:
		return "PG_OR_PC"
	case DeviceGroup_OS:
		return "OS"
	case DeviceGroup_OTHERS:
		return "OTHERS"
	}
	return ""
}

func (e DeviceGroup) String() string {
	return e.name()
}
