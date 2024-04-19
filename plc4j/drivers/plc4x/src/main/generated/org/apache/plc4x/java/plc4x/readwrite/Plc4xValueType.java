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
package org.apache.plc4x.java.plc4x.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum Plc4xValueType {
  NULL((short) 0x00),
  BOOL((short) 0x01),
  BYTE((short) 0x02),
  WORD((short) 0x03),
  DWORD((short) 0x04),
  LWORD((short) 0x05),
  USINT((short) 0x11),
  UINT((short) 0x12),
  UDINT((short) 0x13),
  ULINT((short) 0x14),
  SINT((short) 0x21),
  INT((short) 0x22),
  DINT((short) 0x23),
  LINT((short) 0x24),
  REAL((short) 0x31),
  LREAL((short) 0x32),
  CHAR((short) 0x41),
  WCHAR((short) 0x42),
  STRING((short) 0x43),
  WSTRING((short) 0x44),
  TIME((short) 0x51),
  LTIME((short) 0x52),
  DATE((short) 0x53),
  LDATE((short) 0x54),
  TIME_OF_DAY((short) 0x55),
  LTIME_OF_DAY((short) 0x56),
  DATE_AND_TIME((short) 0x57),
  DATE_AND_LTIME((short) 0x58),
  LDATE_AND_TIME((short) 0x59),
  Struct((short) 0x61),
  List((short) 0x62),
  RAW_BYTE_ARRAY((short) 0x71);
  private static final Map<Short, Plc4xValueType> map;

  static {
    map = new HashMap<>();
    for (Plc4xValueType value : Plc4xValueType.values()) {
      map.put((short) value.getValue(), value);
    }
  }

  private final short value;

  Plc4xValueType(short value) {
    this.value = value;
  }

  public short getValue() {
    return value;
  }

  public static Plc4xValueType enumForValue(short value) {
    return map.get(value);
  }

  public static Boolean isDefined(short value) {
    return map.containsKey(value);
  }
}
