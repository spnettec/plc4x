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
package org.apache.plc4x.java.opcua.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum OpcuaNodeIdServicesVariableTransition {
  TransitionType_TransitionNumber((int) 2312L),
  TransitionVariableType_Id((int) 2763L),
  TransitionVariableType_Name((int) 2764L),
  TransitionVariableType_Number((int) 2765L),
  TransitionVariableType_TransitionTime((int) 2766L),
  TransitionEventType_Transition((int) 2774L),
  TransitionEventType_FromState((int) 2775L),
  TransitionEventType_ToState((int) 2776L),
  TransitionEventType_EventId((int) 3737L),
  TransitionEventType_EventType((int) 3738L),
  TransitionEventType_SourceNode((int) 3739L),
  TransitionEventType_SourceName((int) 3740L),
  TransitionEventType_Time((int) 3741L),
  TransitionEventType_ReceiveTime((int) 3742L),
  TransitionEventType_LocalTime((int) 3743L),
  TransitionEventType_Message((int) 3744L),
  TransitionEventType_Severity((int) 3745L),
  TransitionEventType_FromState_Id((int) 3746L),
  TransitionEventType_FromState_Name((int) 3747L),
  TransitionEventType_FromState_Number((int) 3748L),
  TransitionEventType_FromState_EffectiveDisplayName((int) 3749L),
  TransitionEventType_ToState_Id((int) 3750L),
  TransitionEventType_ToState_Name((int) 3751L),
  TransitionEventType_ToState_Number((int) 3752L),
  TransitionEventType_ToState_EffectiveDisplayName((int) 3753L),
  TransitionEventType_Transition_Id((int) 3754L),
  TransitionEventType_Transition_Name((int) 3755L),
  TransitionEventType_Transition_Number((int) 3756L),
  TransitionEventType_Transition_TransitionTime((int) 3757L),
  TransitionVariableType_EffectiveTransitionTime((int) 11456L),
  TransitionEventType_Transition_EffectiveTransitionTime((int) 11460L),
  TransitionEventType_ConditionClassId((int) 31919L),
  TransitionEventType_ConditionClassName((int) 31920L),
  TransitionEventType_ConditionSubClassId((int) 31921L),
  TransitionEventType_ConditionSubClassName((int) 31922L);
  private static final Map<Integer, OpcuaNodeIdServicesVariableTransition> map;

  static {
    map = new HashMap<>();
    for (OpcuaNodeIdServicesVariableTransition value :
        OpcuaNodeIdServicesVariableTransition.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  OpcuaNodeIdServicesVariableTransition(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static OpcuaNodeIdServicesVariableTransition enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
