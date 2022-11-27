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
package org.apache.plc4x.java.knxnetip.readwrite;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum DeviceDescriptor {
  TP1_BCU_1_SYSTEM_1_0((int) 0x0010, FirmwareType.SYSTEM_1, DeviceDescriptorMediumType.TP1),
  TP1_BCU_1_SYSTEM_1_1((int) 0x0011, FirmwareType.SYSTEM_1, DeviceDescriptorMediumType.TP1),
  TP1_BCU_1_SYSTEM_1_2((int) 0x0012, FirmwareType.SYSTEM_1, DeviceDescriptorMediumType.TP1),
  TP1_BCU_1_SYSTEM_1_3((int) 0x0013, FirmwareType.SYSTEM_1, DeviceDescriptorMediumType.TP1),
  TP1_BCU_2_SYSTEM_2_0((int) 0x0020, FirmwareType.SYSTEM_2, DeviceDescriptorMediumType.TP1),
  TP1_BCU_2_SYSTEM_2_1((int) 0x0021, FirmwareType.SYSTEM_2, DeviceDescriptorMediumType.TP1),
  TP1_BCU_2_SYSTEM_2_5((int) 0x0025, FirmwareType.SYSTEM_2, DeviceDescriptorMediumType.TP1),
  TP1_SYSTEM_300((int) 0x0300, FirmwareType.SYSTEM_300, DeviceDescriptorMediumType.TP1),
  TP1_BIM_M112_0((int) 0x0700, FirmwareType.SYSTEM_7, DeviceDescriptorMediumType.TP1),
  TP1_BIM_M112_1((int) 0x0701, FirmwareType.SYSTEM_7, DeviceDescriptorMediumType.TP1),
  TP1_BIM_M112_5((int) 0x0705, FirmwareType.SYSTEM_7, DeviceDescriptorMediumType.TP1),
  TP1_SYSTEM_B((int) 0x07B0, FirmwareType.SYSTEM_B, DeviceDescriptorMediumType.TP1),
  TP1_IR_DECODER_0((int) 0x0810, FirmwareType.IR_DECODER, DeviceDescriptorMediumType.TP1),
  TP1_IR_DECODER_1((int) 0x0811, FirmwareType.IR_DECODER, DeviceDescriptorMediumType.TP1),
  TP1_COUPLER_0((int) 0x0910, FirmwareType.COUPLER, DeviceDescriptorMediumType.TP1),
  TP1_COUPLER_1((int) 0x0911, FirmwareType.COUPLER, DeviceDescriptorMediumType.TP1),
  TP1_COUPLER_2((int) 0x0912, FirmwareType.COUPLER, DeviceDescriptorMediumType.TP1),
  TP1_KNXNETIP_ROUTER((int) 0x091A, FirmwareType.COUPLER, DeviceDescriptorMediumType.TP1),
  TP1_NONE_D((int) 0x0AFD, FirmwareType.NONE, DeviceDescriptorMediumType.TP1),
  TP1_NONE_E((int) 0x0AFE, FirmwareType.NONE, DeviceDescriptorMediumType.TP1),
  PL110_BCU_1_2((int) 0x1012, FirmwareType.SYSTEM_1, DeviceDescriptorMediumType.PL110),
  PL110_BCU_1_3((int) 0x1013, FirmwareType.SYSTEM_1, DeviceDescriptorMediumType.PL110),
  PL110_SYSTEM_B((int) 0x17B0, FirmwareType.SYSTEM_B, DeviceDescriptorMediumType.PL110),
  PL110_MEDIA_COUPLER_PL_TP(
      (int) 0x1900, FirmwareType.MEDIA_COUPLER_PL_TP, DeviceDescriptorMediumType.PL110),
  RF_BI_DIRECTIONAL_DEVICES(
      (int) 0x2010, FirmwareType.RF_BI_DIRECTIONAL_DEVICES, DeviceDescriptorMediumType.RF),
  RF_UNI_DIRECTIONAL_DEVICES(
      (int) 0x2110, FirmwareType.RF_UNI_DIRECTIONAL_DEVICES, DeviceDescriptorMediumType.RF),
  TP0_BCU_1((int) 0x3012, FirmwareType.SYSTEM_1, DeviceDescriptorMediumType.TP0),
  PL132_BCU_1((int) 0x4012, FirmwareType.SYSTEM_1, DeviceDescriptorMediumType.PL132),
  KNX_IP_SYSTEM7((int) 0x5705, FirmwareType.SYSTEM_7, DeviceDescriptorMediumType.KNX_IP);
  private static final Map<Integer, DeviceDescriptor> map;

  static {
    map = new HashMap<>();
    for (DeviceDescriptor value : DeviceDescriptor.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private int value;
  private FirmwareType firmwareType;
  private DeviceDescriptorMediumType mediumType;

  DeviceDescriptor(int value, FirmwareType firmwareType, DeviceDescriptorMediumType mediumType) {
    this.value = value;
    this.firmwareType = firmwareType;
    this.mediumType = mediumType;
  }

  public int getValue() {
    return value;
  }

  public FirmwareType getFirmwareType() {
    return firmwareType;
  }

  public static DeviceDescriptor firstEnumForFieldFirmwareType(FirmwareType fieldValue) {
    for (DeviceDescriptor _val : DeviceDescriptor.values()) {
      if (_val.getFirmwareType() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<DeviceDescriptor> enumsForFieldFirmwareType(FirmwareType fieldValue) {
    List<DeviceDescriptor> _values = new ArrayList();
    for (DeviceDescriptor _val : DeviceDescriptor.values()) {
      if (_val.getFirmwareType() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public DeviceDescriptorMediumType getMediumType() {
    return mediumType;
  }

  public static DeviceDescriptor firstEnumForFieldMediumType(
      DeviceDescriptorMediumType fieldValue) {
    for (DeviceDescriptor _val : DeviceDescriptor.values()) {
      if (_val.getMediumType() == fieldValue) {
        return _val;
      }
    }
    return null;
  }

  public static List<DeviceDescriptor> enumsForFieldMediumType(
      DeviceDescriptorMediumType fieldValue) {
    List<DeviceDescriptor> _values = new ArrayList();
    for (DeviceDescriptor _val : DeviceDescriptor.values()) {
      if (_val.getMediumType() == fieldValue) {
        _values.add(_val);
      }
    }
    return _values;
  }

  public static DeviceDescriptor enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
