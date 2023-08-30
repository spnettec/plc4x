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

// OpcuaNodeIdServicesVariableState is an enum
type OpcuaNodeIdServicesVariableState int32

type IOpcuaNodeIdServicesVariableState interface {
	fmt.Stringer
	utils.LengthAware
	utils.Serializable
}

const (
	OpcuaNodeIdServicesVariableState_StateType_StateNumber                                   OpcuaNodeIdServicesVariableState = 2308
	OpcuaNodeIdServicesVariableState_StateVariableType_Id                                    OpcuaNodeIdServicesVariableState = 2756
	OpcuaNodeIdServicesVariableState_StateVariableType_Name                                  OpcuaNodeIdServicesVariableState = 2757
	OpcuaNodeIdServicesVariableState_StateVariableType_Number                                OpcuaNodeIdServicesVariableState = 2758
	OpcuaNodeIdServicesVariableState_StateVariableType_EffectiveDisplayName                  OpcuaNodeIdServicesVariableState = 2759
	OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState                           OpcuaNodeIdServicesVariableState = 2769
	OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition                         OpcuaNodeIdServicesVariableState = 2770
	OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_Id                        OpcuaNodeIdServicesVariableState = 3720
	OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_Name                      OpcuaNodeIdServicesVariableState = 3721
	OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_Number                    OpcuaNodeIdServicesVariableState = 3722
	OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_EffectiveDisplayName      OpcuaNodeIdServicesVariableState = 3723
	OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_Id                      OpcuaNodeIdServicesVariableState = 3724
	OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_Name                    OpcuaNodeIdServicesVariableState = 3725
	OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_Number                  OpcuaNodeIdServicesVariableState = 3726
	OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_TransitionTime          OpcuaNodeIdServicesVariableState = 3727
	OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_EffectiveTransitionTime OpcuaNodeIdServicesVariableState = 11458
)

var OpcuaNodeIdServicesVariableStateValues []OpcuaNodeIdServicesVariableState

func init() {
	_ = errors.New
	OpcuaNodeIdServicesVariableStateValues = []OpcuaNodeIdServicesVariableState{
		OpcuaNodeIdServicesVariableState_StateType_StateNumber,
		OpcuaNodeIdServicesVariableState_StateVariableType_Id,
		OpcuaNodeIdServicesVariableState_StateVariableType_Name,
		OpcuaNodeIdServicesVariableState_StateVariableType_Number,
		OpcuaNodeIdServicesVariableState_StateVariableType_EffectiveDisplayName,
		OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState,
		OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition,
		OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_Id,
		OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_Name,
		OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_Number,
		OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_EffectiveDisplayName,
		OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_Id,
		OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_Name,
		OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_Number,
		OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_TransitionTime,
		OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_EffectiveTransitionTime,
	}
}

func OpcuaNodeIdServicesVariableStateByValue(value int32) (enum OpcuaNodeIdServicesVariableState, ok bool) {
	switch value {
	case 11458:
		return OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_EffectiveTransitionTime, true
	case 2308:
		return OpcuaNodeIdServicesVariableState_StateType_StateNumber, true
	case 2756:
		return OpcuaNodeIdServicesVariableState_StateVariableType_Id, true
	case 2757:
		return OpcuaNodeIdServicesVariableState_StateVariableType_Name, true
	case 2758:
		return OpcuaNodeIdServicesVariableState_StateVariableType_Number, true
	case 2759:
		return OpcuaNodeIdServicesVariableState_StateVariableType_EffectiveDisplayName, true
	case 2769:
		return OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState, true
	case 2770:
		return OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition, true
	case 3720:
		return OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_Id, true
	case 3721:
		return OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_Name, true
	case 3722:
		return OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_Number, true
	case 3723:
		return OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_EffectiveDisplayName, true
	case 3724:
		return OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_Id, true
	case 3725:
		return OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_Name, true
	case 3726:
		return OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_Number, true
	case 3727:
		return OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_TransitionTime, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableStateByName(value string) (enum OpcuaNodeIdServicesVariableState, ok bool) {
	switch value {
	case "StateMachineType_LastTransition_EffectiveTransitionTime":
		return OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_EffectiveTransitionTime, true
	case "StateType_StateNumber":
		return OpcuaNodeIdServicesVariableState_StateType_StateNumber, true
	case "StateVariableType_Id":
		return OpcuaNodeIdServicesVariableState_StateVariableType_Id, true
	case "StateVariableType_Name":
		return OpcuaNodeIdServicesVariableState_StateVariableType_Name, true
	case "StateVariableType_Number":
		return OpcuaNodeIdServicesVariableState_StateVariableType_Number, true
	case "StateVariableType_EffectiveDisplayName":
		return OpcuaNodeIdServicesVariableState_StateVariableType_EffectiveDisplayName, true
	case "StateMachineType_CurrentState":
		return OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState, true
	case "StateMachineType_LastTransition":
		return OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition, true
	case "StateMachineType_CurrentState_Id":
		return OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_Id, true
	case "StateMachineType_CurrentState_Name":
		return OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_Name, true
	case "StateMachineType_CurrentState_Number":
		return OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_Number, true
	case "StateMachineType_CurrentState_EffectiveDisplayName":
		return OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_EffectiveDisplayName, true
	case "StateMachineType_LastTransition_Id":
		return OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_Id, true
	case "StateMachineType_LastTransition_Name":
		return OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_Name, true
	case "StateMachineType_LastTransition_Number":
		return OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_Number, true
	case "StateMachineType_LastTransition_TransitionTime":
		return OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_TransitionTime, true
	}
	return 0, false
}

func OpcuaNodeIdServicesVariableStateKnows(value int32) bool {
	for _, typeValue := range OpcuaNodeIdServicesVariableStateValues {
		if int32(typeValue) == value {
			return true
		}
	}
	return false
}

func CastOpcuaNodeIdServicesVariableState(structType any) OpcuaNodeIdServicesVariableState {
	castFunc := func(typ any) OpcuaNodeIdServicesVariableState {
		if sOpcuaNodeIdServicesVariableState, ok := typ.(OpcuaNodeIdServicesVariableState); ok {
			return sOpcuaNodeIdServicesVariableState
		}
		return 0
	}
	return castFunc(structType)
}

func (m OpcuaNodeIdServicesVariableState) GetLengthInBits(ctx context.Context) uint16 {
	return 32
}

func (m OpcuaNodeIdServicesVariableState) GetLengthInBytes(ctx context.Context) uint16 {
	return m.GetLengthInBits(ctx) / 8
}

func OpcuaNodeIdServicesVariableStateParse(ctx context.Context, theBytes []byte) (OpcuaNodeIdServicesVariableState, error) {
	return OpcuaNodeIdServicesVariableStateParseWithBuffer(ctx, utils.NewReadBufferByteBased(theBytes))
}

func OpcuaNodeIdServicesVariableStateParseWithBuffer(ctx context.Context, readBuffer utils.ReadBuffer) (OpcuaNodeIdServicesVariableState, error) {
	log := zerolog.Ctx(ctx)
	_ = log
	val, err := readBuffer.ReadInt32("OpcuaNodeIdServicesVariableState", 32)
	if err != nil {
		return 0, errors.Wrap(err, "error reading OpcuaNodeIdServicesVariableState")
	}
	if enum, ok := OpcuaNodeIdServicesVariableStateByValue(val); !ok {
		log.Debug().Interface("val", val).Msg("no value val found for OpcuaNodeIdServicesVariableState")
		return OpcuaNodeIdServicesVariableState(val), nil
	} else {
		return enum, nil
	}
}

func (e OpcuaNodeIdServicesVariableState) Serialize() ([]byte, error) {
	wb := utils.NewWriteBufferByteBased()
	if err := e.SerializeWithWriteBuffer(context.Background(), wb); err != nil {
		return nil, err
	}
	return wb.GetBytes(), nil
}

func (e OpcuaNodeIdServicesVariableState) SerializeWithWriteBuffer(ctx context.Context, writeBuffer utils.WriteBuffer) error {
	log := zerolog.Ctx(ctx)
	_ = log
	return writeBuffer.WriteInt32("OpcuaNodeIdServicesVariableState", 32, int32(e), utils.WithAdditionalStringRepresentation(e.PLC4XEnumName()))
}

// PLC4XEnumName returns the name that is used in code to identify this enum
func (e OpcuaNodeIdServicesVariableState) PLC4XEnumName() string {
	switch e {
	case OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_EffectiveTransitionTime:
		return "StateMachineType_LastTransition_EffectiveTransitionTime"
	case OpcuaNodeIdServicesVariableState_StateType_StateNumber:
		return "StateType_StateNumber"
	case OpcuaNodeIdServicesVariableState_StateVariableType_Id:
		return "StateVariableType_Id"
	case OpcuaNodeIdServicesVariableState_StateVariableType_Name:
		return "StateVariableType_Name"
	case OpcuaNodeIdServicesVariableState_StateVariableType_Number:
		return "StateVariableType_Number"
	case OpcuaNodeIdServicesVariableState_StateVariableType_EffectiveDisplayName:
		return "StateVariableType_EffectiveDisplayName"
	case OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState:
		return "StateMachineType_CurrentState"
	case OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition:
		return "StateMachineType_LastTransition"
	case OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_Id:
		return "StateMachineType_CurrentState_Id"
	case OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_Name:
		return "StateMachineType_CurrentState_Name"
	case OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_Number:
		return "StateMachineType_CurrentState_Number"
	case OpcuaNodeIdServicesVariableState_StateMachineType_CurrentState_EffectiveDisplayName:
		return "StateMachineType_CurrentState_EffectiveDisplayName"
	case OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_Id:
		return "StateMachineType_LastTransition_Id"
	case OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_Name:
		return "StateMachineType_LastTransition_Name"
	case OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_Number:
		return "StateMachineType_LastTransition_Number"
	case OpcuaNodeIdServicesVariableState_StateMachineType_LastTransition_TransitionTime:
		return "StateMachineType_LastTransition_TransitionTime"
	}
	return fmt.Sprintf("Unknown(%v)", int32(e))
}

func (e OpcuaNodeIdServicesVariableState) String() string {
	return e.PLC4XEnumName()
}
