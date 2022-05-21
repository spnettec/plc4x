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

// BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable is an enum
type BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable uint8

type IBACnetConfirmedServiceRequestReinitializeDeviceEnableDisable interface {
	Serialize(writeBuffer utils.WriteBuffer) error
}

const (
	BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable_ENABLE             BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable = 0
	BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable_DISABLE            BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable = 1
	BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable_DISABLE_INITIATION BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable = 2
)

var BACnetConfirmedServiceRequestReinitializeDeviceEnableDisableValues []BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable

func init() {
	_ = errors.New
	BACnetConfirmedServiceRequestReinitializeDeviceEnableDisableValues = []BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable{
		BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable_ENABLE,
		BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable_DISABLE,
		BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable_DISABLE_INITIATION,
	}
}

func BACnetConfirmedServiceRequestReinitializeDeviceEnableDisableByValue(value uint8) BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable {
	switch value {
	case 0:
		return BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable_ENABLE
	case 1:
		return BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable_DISABLE
	case 2:
		return BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable_DISABLE_INITIATION
	}
	return 0
}

func BACnetConfirmedServiceRequestReinitializeDeviceEnableDisableByName(value string) BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable {
	switch value {
	case "ENABLE":
		return BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable_ENABLE
	case "DISABLE":
		return BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable_DISABLE
	case "DISABLE_INITIATION":
		return BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable_DISABLE_INITIATION
	}
	return 0
}

func BACnetConfirmedServiceRequestReinitializeDeviceEnableDisableKnows(value uint8) bool {
	for _, typeValue := range BACnetConfirmedServiceRequestReinitializeDeviceEnableDisableValues {
		if uint8(typeValue) == value {
			return true
		}
	}
	return false
}

func CastBACnetConfirmedServiceRequestReinitializeDeviceEnableDisable(structType interface{}) BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable {
	castFunc := func(typ interface{}) BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable {
		if sBACnetConfirmedServiceRequestReinitializeDeviceEnableDisable, ok := typ.(BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable); ok {
			return sBACnetConfirmedServiceRequestReinitializeDeviceEnableDisable
		}
		return 0
	}
	return castFunc(structType)
}

func (m BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable) GetLengthInBits() uint16 {
	return 8
}

func (m BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable) GetLengthInBytes() uint16 {
	return m.GetLengthInBits() / 8
}

func BACnetConfirmedServiceRequestReinitializeDeviceEnableDisableParse(readBuffer utils.ReadBuffer) (BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable, error) {
	val, err := readBuffer.ReadUint8("BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable", 8)
	if err != nil {
		return 0, nil
	}
	return BACnetConfirmedServiceRequestReinitializeDeviceEnableDisableByValue(val), nil
}

func (e BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable) Serialize(writeBuffer utils.WriteBuffer) error {
	return writeBuffer.WriteUint8("BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable", 8, uint8(e), utils.WithAdditionalStringRepresentation(e.name()))
}

func (e BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable) name() string {
	switch e {
	case BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable_ENABLE:
		return "ENABLE"
	case BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable_DISABLE:
		return "DISABLE"
	case BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable_DISABLE_INITIATION:
		return "DISABLE_INITIATION"
	}
	return ""
}

func (e BACnetConfirmedServiceRequestReinitializeDeviceEnableDisable) String() string {
	return e.name()
}
