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

public enum OpcuaNodeIdServicesVariableHistorical {
  HistoricalDataConfigurationType_Stepped((int) 2323L),
  HistoricalDataConfigurationType_Definition((int) 2324L),
  HistoricalDataConfigurationType_MaxTimeInterval((int) 2325L),
  HistoricalDataConfigurationType_MinTimeInterval((int) 2326L),
  HistoricalDataConfigurationType_ExceptionDeviation((int) 2327L),
  HistoricalDataConfigurationType_ExceptionDeviationFormat((int) 2328L),
  HistoricalDataConfigurationType_AggregateConfiguration_TreatUncertainAsBad((int) 11168L),
  HistoricalDataConfigurationType_AggregateConfiguration_PercentDataBad((int) 11169L),
  HistoricalDataConfigurationType_AggregateConfiguration_PercentDataGood((int) 11170L),
  HistoricalDataConfigurationType_AggregateConfiguration_UseSlopedExtrapolation((int) 11171L),
  HistoricalEventFilter((int) 11215L),
  HistoricalDataConfigurationType_StartOfArchive((int) 11499L),
  HistoricalDataConfigurationType_StartOfOnlineArchive((int) 11500L),
  HistoricalDataConfigurationType_ServerTimestampSupported((int) 19092L);
  private static final Map<Integer, OpcuaNodeIdServicesVariableHistorical> map;

  static {
    map = new HashMap<>();
    for (OpcuaNodeIdServicesVariableHistorical value :
        OpcuaNodeIdServicesVariableHistorical.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  OpcuaNodeIdServicesVariableHistorical(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static OpcuaNodeIdServicesVariableHistorical enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
