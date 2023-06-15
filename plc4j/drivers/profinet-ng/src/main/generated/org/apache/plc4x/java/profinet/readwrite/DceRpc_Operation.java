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
package org.apache.plc4x.java.profinet.readwrite;

import java.util.HashMap;
import java.util.Map;

// Code generated by code-generation. DO NOT EDIT.

public enum DceRpc_Operation {
  CONNECT((int) 0x0000),
  RELEASE((int) 0x0001),
  READ((int) 0x0002),
  WRITE((int) 0x0003),
  CONTROL((int) 0x0004),
  READ_IMPLICIT((int) 0x0005);
  private static final Map<Integer, DceRpc_Operation> map;

  static {
    map = new HashMap<>();
    for (DceRpc_Operation value : DceRpc_Operation.values()) {
      map.put((int) value.getValue(), value);
    }
  }

  private final int value;

  DceRpc_Operation(int value) {
    this.value = value;
  }

  public int getValue() {
    return value;
  }

  public static DceRpc_Operation enumForValue(int value) {
    return map.get(value);
  }

  public static Boolean isDefined(int value) {
    return map.containsKey(value);
  }
}
